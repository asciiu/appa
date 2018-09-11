package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	binance "github.com/asciiu/go-binance"
	protoBalance "github.com/asciiu/gomo/binance-service/proto/balance"
	constExt "github.com/asciiu/gomo/common/constants/exchange"
	constRes "github.com/asciiu/gomo/common/constants/response"
	protoEvt "github.com/asciiu/gomo/common/proto/events"
	constPlan "github.com/asciiu/gomo/plan-service/constants"
	kitLog "github.com/go-kit/kit/log"
	micro "github.com/micro/go-micro"
)

type BinanceService struct {
	CompletedPub micro.Publisher
}

func (service *BinanceService) HandleFillOrder(ctx context.Context, triggerEvent *protoEvt.TriggeredOrderEvent) error {
	// ignore events not binance
	// perhaps we can have this handler only receive binance triggers but for the sake of
	// simplicity when adding new exchanges let's just have each exchange service do a check
	// on the exchange
	if triggerEvent.Exchange != constExt.Binance {
		return nil
	}

	go func() {
		var logger kitLog.Logger
		logger = kitLog.NewLogfmtLogger(os.Stdout)
		logger = kitLog.With(logger, "time", kitLog.DefaultTimestampUTC, "caller", kitLog.DefaultCaller)

		hmacSigner := &binance.HmacSigner{
			Key: []byte(triggerEvent.Secret),
		}

		binanceService := binance.NewAPIService(
			"https://www.binance.com",
			triggerEvent.Key,
			hmacSigner,
			logger,
			ctx,
		)
		b := binance.NewBinance(binanceService)

		// Buy-limit: plan.baseBalance / planOrder.Price
		// Buy-market: plan.baseBalance / trigger.Price (can only determine this at trigger time)
		// Sell-limit: currencyBalance
		// Sell-market: currencyBalance

		// binance expects the symbol to be formatted as a single word: e.g. BNBBTC
		symbol := strings.Replace(triggerEvent.MarketName, "-", "", 1)
		// buy or sell
		ellado := binance.SideBuy
		if triggerEvent.Side == constPlan.Sell {
			ellado = binance.SideSell
		}
		// order type can be market or limit
		orderType := binance.TypeMarket
		if triggerEvent.OrderType == constPlan.LimitOrder {
			orderType = binance.TypeLimit
		}

		completedEvent := protoEvt.CompletedOrderEvent{
			UserID:             triggerEvent.UserID,
			PlanID:             triggerEvent.PlanID,
			OrderID:            triggerEvent.OrderID,
			Side:               triggerEvent.Side,
			TriggeredPrice:     triggerEvent.TriggeredPrice,
			TriggeredCondition: triggerEvent.TriggeredCondition,
		}
		// https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md
		// Limit type orders require a price
		processedOrder, err := b.NewOrder(binance.NewOrderRequest{
			Symbol:      symbol,
			Quantity:    triggerEvent.Quantity,
			Side:        ellado,
			Price:       triggerEvent.Price,
			TimeInForce: binance.GTC,
			Type:        orderType,
			Timestamp:   time.Now(),
		})

		if err != nil {
			log.Printf("failed binance order call -- event: %+v\n", triggerEvent)
			completedEvent.Status = constPlan.Failed
			completedEvent.Details = err.Error()
		}

		completedEvent.Status = constPlan.Filled

		if err := service.CompletedPub.Publish(ctx, &completedEvent); err != nil {
			log.Println("publish err: ", err.Error())
		}

		log.Printf("processed order -- %+v\n", processedOrder)
	}()
	return nil
}

// Retrieve the exchange balances
func (service *BinanceService) GetBalances(ctx context.Context, req *protoBalance.BalanceRequest, res *protoBalance.BalancesResponse) error {

	var logger kitLog.Logger
	logger = kitLog.NewLogfmtLogger(os.Stdout)
	logger = kitLog.With(logger, "time", kitLog.DefaultTimestampUTC, "caller", kitLog.DefaultCaller)

	hmacSigner := &binance.HmacSigner{
		Key: []byte(req.KeySecret),
	}

	binanceService := binance.NewAPIService(
		"https://www.binance.com",
		req.KeyPublic,
		hmacSigner,
		logger,
		ctx,
	)
	b := binance.NewBinance(binanceService)
	request := binance.AccountRequest{
		RecvWindow: time.Duration(2) * time.Second,
		Timestamp:  time.Now(),
	}

	account, err := b.Account(request)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "nvalid"):
			res.Status = constRes.Fail
			res.Message = "invalid keys"
			return nil
		default:
			res.Status = constRes.Error
			res.Message = err.Error()
			return nil
		}
	}

	balances := make([]*protoBalance.Balance, 0)
	for _, balance := range account.Balances {
		bal := protoBalance.Balance{
			CurrencySymbol: balance.Asset,
			Free:           balance.Free,
			Locked:         balance.Locked,
		}
		balances = append(balances, &bal)
	}

	res.Status = constRes.Success
	res.Data = &protoBalance.BalanceList{
		Balances: balances,
	}

	return nil

}
