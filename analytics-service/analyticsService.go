package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	repoAnalytics "github.com/asciiu/gomo/analytics-service/db/sql"
	protoAnalytics "github.com/asciiu/gomo/analytics-service/proto/analytics"
	constRes "github.com/asciiu/gomo/common/constants/response"
	evt "github.com/asciiu/gomo/common/proto/events"
	"github.com/lib/pq"
	micro "github.com/micro/go-micro"
)

// Processor will process orders
type AnalyticsService struct {
	sync.RWMutex
	DB               *sql.DB
	MarketClosePrice map[Market]float64

	MarketCandles map[Market]string
	ProcessQueue  map[Market]float64
	CandlePub     micro.Publisher
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func (service *AnalyticsService) Ticker() error {
	fmt.Println("ticker started: ", time.Now())
	for {
		time.Sleep(2 * time.Second)

		for k := range service.ProcessQueue {
			service.Lock()

			service.MarketCandles[k] = "ding"
			delete(service.ProcessQueue, k)

			service.Unlock()
			fmt.Println(k)
			break
		}

		// get market from collection
		// send out candle request if no candles
		// otherwise ignore
		//if length > 0 {
		//	market := processor.MarketClosePrice[0]
		//	fmt.Println(market)
		//}
	}
	return nil
}

// ProcessEvent will process ExchangeEvents. These events are published from the exchange sockets.
func (service *AnalyticsService) HandleTradeEvent(payload *evt.TradeEvents) error {
	// record close price for the market
	for _, event := range payload.Events {
		market := Market{
			Exchange: event.Exchange,
			Name:     event.MarketName,
		}

		service.RLock()
		_, ok1 := service.MarketCandles[market]
		_, ok2 := service.ProcessQueue[market]
		service.RUnlock()

		if !ok1 && !ok2 {
			service.Lock()
			service.ProcessQueue[market] = event.Price
			service.Unlock()
		}

		//found := false
		// for _, m := range processor.MarketQueue {
		// 	if m == marketName {
		// 		found = true
		// 	}
		// }
		// if !found {
		// 	processor.MarketQueue = append(processor.MarketQueue, marketName)
		// }
	}

	return nil
}

// This function was formly known as the Amigoni special. It has been refined by yours truely - Axl Codes.
func (service *AnalyticsService) ConvertCurrency(ctx context.Context, req *protoAnalytics.ConversionRequest, res *protoAnalytics.ConversionResponse) error {
	var rate, reverse, fromRate, toRate float64
	from := req.From
	to := req.To
	atTime, _ := time.Parse(time.RFC3339, req.AtTimestamp)
	trunctTime := atTime.Truncate(time.Duration(5) * time.Minute)
	trunctTimeStr := string(pq.FormatTimestamp(trunctTime))

	//Case in which to and from are the same i.e. BTCBTC
	if from == to {
		res.Status = constRes.Success
		res.Data = &protoAnalytics.ConversionAmount{
			ConvertedAmount: req.FromAmount,
		}
		return nil
	}

	// find all prices here for given time
	markets, err := repoAnalytics.FindExchangeRates(service.DB, req.Exchange, trunctTimeStr)
	if err != nil {
		res.Status = constRes.Error
		res.Message = err.Error()
		return nil
	}

	for _, market := range markets {
		if market.MarketName == from+"-"+to {
			//Simple Case where the rate exists i.e. ADA-BTC
			rate = market.ClosedAtPrice
			break
		}
		if market.MarketName == to+"-"+from {
			// reverse case exists BTC-ADA
			reverse = 1 / market.ClosedAtPrice
		}
		//if market.MarketName == from+"-BTC" {
		//	// indirect from rate
		//	fromRate = market.ClosedAtPrice
		//}
		//if market.MarketName == to+"-BTC" {
		//	// indirect to rate
		//	toRate = market.ClosedAtPrice
		//}
		if market.MarketName == from+"-BTC" {
			// indirect from rate
			fromRate = market.ClosedAtPrice
		}
		if market.MarketName == "BTC-"+to {
			// indirect to rate
			toRate = market.ClosedAtPrice
		}
	}

	switch {
	case rate == 0 && reverse != 0:
		rate = reverse
	case rate == 0 && reverse == 0:
		// direct rate doesn't exist so going through BTC to convert i.e ADAXVG
		//rate = fromRate / toRate
		rate = fromRate * toRate
	}

	//Ti.API.trace("Convert: "+from+" to "+to+" = "+rate+" "+fromExchange);
	res.Status = constRes.Success
	res.Data = &protoAnalytics.ConversionAmount{
		ConvertedAmount: rate * req.FromAmount,
	}
	return nil
}