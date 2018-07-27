package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	balances "github.com/asciiu/gomo/balance-service/proto/balance"
	"github.com/asciiu/gomo/common/constants/key"
	orderConstants "github.com/asciiu/gomo/common/constants/order"
	"github.com/asciiu/gomo/common/constants/plan"
	"github.com/asciiu/gomo/common/constants/response"
	"github.com/asciiu/gomo/common/constants/side"
	"github.com/asciiu/gomo/common/constants/status"
	evt "github.com/asciiu/gomo/common/proto/events"
	keys "github.com/asciiu/gomo/key-service/proto/key"
	planRepo "github.com/asciiu/gomo/plan-service/db/sql"
	protoOrder "github.com/asciiu/gomo/plan-service/proto/order"
	protoPlan "github.com/asciiu/gomo/plan-service/proto/plan"
	"github.com/google/uuid"
	"github.com/lib/pq"
	micro "github.com/micro/go-micro"
)

// PlanService ...
type PlanService struct {
	DB        *sql.DB
	Client    balances.BalanceServiceClient
	KeyClient keys.KeyServiceClient
	OrderPub  micro.Publisher
}

// private: This is where the order events are published to the rest of the system
// this function should only be callable from within the PlanService. When a plan is
// published the first order of the plan will be emmitted as an ActiveOrderEvent to the
// system.
//
// VERY IMPORTANT: Only send Plans where the first order plan's orders is the next order to active.
// That is to say. DO NOT load a plan where the first order in the orders array has been filled. Fuck
// it, I'm going to implement a check here to ensure this never happens.
func (service *PlanService) publishPlan(ctx context.Context, plan *protoPlan.Plan, isRevision bool) error {
	// the first plan order will always be the active one
	planOrder := plan.Orders[0]

	// only pub plan if the next plan order is active or inactive
	// we do not pub plan orders that have been filled, failed, or aborted
	// reexecuting those plan orders would be very bad!
	if planOrder.Status != status.Active && planOrder.Status != status.Inactive {
		return nil
	}
	triggers := make([]*evt.Trigger, 0)
	for _, t := range planOrder.Triggers {
		trig := evt.Trigger{
			TriggerID: t.TriggerID,
			OrderID:   t.OrderID,
			Name:      t.Name,
			Code:      t.Code,
			Triggered: t.Triggered,
			Actions:   t.Actions,
		}
		triggers = append(triggers, &trig)
	}

	// convert order to order event
	// activeOrder := evt.ActiveOrderEvent{
	// 	//Exchange:        plan.Exchange,
	// 	OrderID: planOrder.OrderID,
	// 	PlanID:  plan.PlanID,
	// 	UserID:  plan.UserID,
	// 	//BaseBalance:     plan.BaseBalance,
	// 	//CurrencyBalance: plan.CurrencyBalance,
	// 	BalancePercent: planOrder.PercentBalance,
	// 	KeyID:          plan.KeyID,
	// 	Key:            plan.Key,
	// 	Secret:         plan.KeySecret,
	// 	//MarketName:      plan.MarketName,
	// 	Side:      planOrder.Side,
	// 	OrderType: planOrder.OrderType,
	// 	Price:     planOrder.LimitPrice,
	// 	//NextOrderID:     planOrder.NextOrderID,
	// 	Revision:    isRevision,
	// 	OrderStatus: planOrder.Status,
	// 	Triggers:    triggers,
	// }

	// if err := service.OrderPub.Publish(context.Background(), &activeOrder); err != nil {
	// 	return fmt.Errorf("publish error: %s -- ActiveOrderEvent %+v", err, &activeOrder)
	// }
	//log.Printf("publish active order -- %+v\n", &activeOrder)
	return nil
}

// private: validateBalance
func (service *PlanService) validateBalance(ctx context.Context, currency string, balance float64, userID string, apikeyID string) (bool, error) {
	balRequest := balances.GetUserBalanceRequest{
		UserID:   userID,
		KeyID:    apikeyID,
		Currency: currency,
	}

	balResponse, err := service.Client.GetUserBalance(ctx, &balRequest)
	if err != nil {
		return false, fmt.Errorf("ecountered error from GetUserBalance: %s", err.Error())
	}

	if balResponse.Data.Balance.Available < balance {
		return false, nil
	}
	return true, nil
}

// LoadPlanOrder will activate an order (i.e. send a plan order) to the execution engine to process.
func (service *PlanService) LoadPlanOrder(ctx context.Context, plan *protoPlan.Plan, isRevision bool) error {

	// planOrder := plan.Orders[0]
	// currencies := strings.Split(plan.MarketName, "-")
	// // default market currency
	// currency := currencies[0]
	// balance := plan.CurrencyBalance
	// if planOrder.Side == side.Buy {
	// 	// buy uses base currency
	// 	currency = currencies[1]
	// 	balance = plan.BaseBalance
	// }

	// if err := service.validateBalance(ctx, currency, balance, plan.UserID, plan.KeyID); err != nil {
	// 	return err
	// }

	// if err := service.publishPlan(ctx, plan, isRevision); err != nil {
	// 	return err
	// }

	return nil
}

func (service *PlanService) fetchKeys(keyIDs []string) ([]*keys.Key, error) {
	request := keys.GetKeysRequest{
		KeyIDs: keyIDs}

	r, _ := service.KeyClient.GetKeys(context.Background(), &request)
	if r.Status != response.Success {
		if r.Status == response.Fail {
			return nil, fmt.Errorf(r.Message)
		}
		if r.Status == response.Error {
			return nil, fmt.Errorf(r.Message)
		}
		if r.Status == response.Nonentity {
			return nil, fmt.Errorf("invalid keys")
		}
	}

	return r.Data.Keys, nil
}

// AddPlans returns error to conform to protobuf def, but the error will always be returned as nil.
// Can't return an error with a response object - response object is returned as nil when error is non nil.
// Therefore, return error in response object. MarketName example: ADA-BTC where BTC is base.
func (service *PlanService) NewPlan(ctx context.Context, req *protoPlan.NewPlanRequest, res *protoPlan.PlanResponse) error {

	switch {
	case !ValidatePlanInputStatus(req.Status):
		res.Status = response.Fail
		res.Message = "plan status must be active, inactive, or historic"
		return nil
	case !ValidateSingleRootNode(req.Orders):
		res.Status = response.Fail
		res.Message = "multiple root nodes found, only one is allowed"
		return nil
	case !ValidateConnectedRoutes(req.Orders):
		res.Status = response.Fail
		res.Message = "an order does not have a valid parent_order_id in your request"
		return nil
	case !ValidateNodeCount(req.Orders):
		res.Status = response.Fail
		res.Message = "you can only post 10 inactive nodes at a time!"
		return nil
	case !ValidateNoneZeroBalance(req.Orders):
		res.Status = response.Fail
		res.Message = "non zero currency balance required for root order!"
		return nil
	case !ValidateUniformOrderType(req.Orders):
		res.Status = response.Fail
		res.Message = "all orders must be of the same order type as root"
		return nil
	}

	// fetch all order keys
	keyIDs := make([]string, 0, len(req.Orders))
	for _, or := range req.Orders {
		keyIDs = append(keyIDs, or.KeyID)
	}
	kys, err := service.fetchKeys(keyIDs)
	if err != nil {
		res.Status = response.Error
		res.Message = fmt.Sprintf("ecountered error when fetching keys: %s", err.Error())
		return nil
	}

	none := uuid.Nil.String()
	planID := uuid.New()
	now := string(pq.FormatTimestamp(time.Now().UTC()))
	newOrders := make([]*protoOrder.Order, 0, len(req.Orders))
	exchange := ""

	for _, or := range req.Orders {
		orderStatus := status.Inactive
		depth := uint32(1)

		if or.MarketName == "" || or.KeyID == "" {
			res.Status = response.Fail
			res.Message = "missing marketName/keyID for order"
			return nil
		}
		if !strings.Contains(or.MarketName, "-") {
			res.Status = response.Fail
			res.Message = "marketName must be currency-base: e.g. ADA-BTC"
			return nil
		}
		if !ValidateOrderType(or.OrderType) {
			res.Status = response.Fail
			res.Message = "market, limit, or paper required for order type"
			return nil
		}
		if !ValidateOrderSide(or.Side) {
			res.Status = response.Fail
			res.Message = "buy or sell required for order side"
			return nil
		}

		// compute the depth for the order
		if or.ParentOrderID != none {
			for _, o := range newOrders {
				if o.OrderID == or.ParentOrderID {
					depth = o.PlanDepth + 1
					break
				}
			}
		}

		if or.ParentOrderID == none && req.Status == plan.Active {
			orderStatus = status.Active
		}

		// assign exchange name from key
		for _, ky := range kys {
			if ky.KeyID == or.KeyID {
				exchange = ky.Exchange

				if ky.Status != key.Verified {
					res.Status = response.Fail
					res.Message = "using an unverified key!"
					return nil

				}
			}
		}

		// collect triggers for this order
		triggers := make([]*protoOrder.Trigger, 0, len(or.Triggers))
		for j, cond := range or.Triggers {
			triggerID := uuid.New()
			trigger := protoOrder.Trigger{
				TriggerID:         triggerID.String(),
				TriggerNumber:     uint32(j),
				TriggerTemplateID: cond.TriggerTemplateID,
				OrderID:           or.OrderID,
				Name:              cond.Name,
				Code:              cond.Code,
				Actions:           cond.Actions,
				Triggered:         false,
				CreatedOn:         now,
				UpdatedOn:         now,
			}
			triggers = append(triggers, &trigger)
		}

		// market name will be Currency-Base: ADA-BTC
		symbolPair := strings.Split(or.MarketName, "-")
		symbol := symbolPair[1]
		if or.Side == side.Sell {
			symbol = symbolPair[0]
		}

		order := protoOrder.Order{
			KeyID:                 or.KeyID,
			OrderID:               or.OrderID,
			OrderPriority:         or.OrderPriority,
			OrderType:             or.OrderType,
			OrderTemplateID:       or.OrderTemplateID,
			ParentOrderID:         or.ParentOrderID,
			PlanID:                planID.String(),
			PlanDepth:             depth,
			Side:                  or.Side,
			LimitPrice:            or.LimitPrice,
			Exchange:              exchange,
			MarketName:            or.MarketName,
			ActiveCurrencySymbol:  symbol,
			ActiveCurrencyBalance: or.ActiveCurrencyBalance,
			Status:                orderStatus,
			Triggers:              triggers,
			CreatedOn:             now,
			UpdatedOn:             now,
		}
		newOrders = append(newOrders, &order)
	}

	currencySymbol := newOrders[0].ActiveCurrencySymbol
	currencyBalance := newOrders[0].ActiveCurrencyBalance
	keyID := newOrders[0].KeyID

	if newOrders[0].OrderType != orderConstants.PaperOrder {
		validBalance, err := service.validateBalance(ctx, currencySymbol, currencyBalance, req.UserID, keyID)
		if err != nil {
			res.Status = response.Error
			res.Message = fmt.Sprintf("failed to validate the currency balance for %s: %s", currencySymbol, err.Error())
			return nil
		}
		if !validBalance {
			res.Status = response.Fail
			res.Message = fmt.Sprintf("insufficient %s balance, %.8f requested", currencySymbol, currencyBalance)
			return nil

		}
	}

	pln := protoPlan.Plan{
		PlanID:                planID.String(),
		PlanTemplateID:        req.PlanTemplateID,
		UserID:                req.UserID,
		ActiveCurrencySymbol:  newOrders[0].ActiveCurrencySymbol,
		ActiveCurrencyBalance: newOrders[0].ActiveCurrencyBalance,
		Exchange:              newOrders[0].Exchange,
		MarketName:            newOrders[0].MarketName,
		LastExecutedPlanDepth: 0,
		LastExecutedOrderID:   none,
		Orders:                newOrders,
		Status:                req.Status,
		CloseOnComplete:       req.CloseOnComplete,
		CreatedOn:             now,
		UpdatedOn:             now,
	}

	error := planRepo.InsertPlan(service.DB, &pln)
	if error != nil {
		res.Status = response.Error
		res.Message = "NewPlan error: " + error.Error()
		return nil
	}

	// activate first plan order if plan is active
	if pln.Status == plan.Active {
		// send key and secret with plan
		//pln.Key = ky.Key
		//pln.KeySecret = ky.Secret

		// this is a new plan
		if err := service.publishPlan(ctx, &pln, false); err != nil {
			// TODO return a warning here
			res.Status = response.Error
			res.Message = "could not publish first order: " + err.Error()
			return nil
		}
	}

	res.Status = response.Success
	res.Data = &protoPlan.PlanData{Plan: &pln}

	return nil
}

// GetUserPlan returns error to conform to protobuf def, but the error will always be returned as nil.
// Can't return an error with a response object - response object is returned as nil when error is non nil.
// Therefore, return error in response object.
func (service *PlanService) GetUserPlan(ctx context.Context, req *protoPlan.GetUserPlanRequest, res *protoPlan.PlanResponse) error {
	plan, error := planRepo.FindPlanOrders(service.DB, req)

	switch {
	case error == sql.ErrNoRows:
		res.Status = response.Nonentity
		res.Message = fmt.Sprintf("planID not found %s", req.PlanID)
	case error != nil:
		res.Status = response.Error
		res.Message = error.Error()
	// case plan.totalDepth < req.PlanDepth:
	// 	res.Status = response.Nonentity
	// 	res.Message = "plan depth out of bounds, max depth is %s"
	case error == nil:
		res.Status = response.Success
		res.Data = &protoPlan.PlanData{Plan: plan}
	}

	return nil
}

// GetUserPlans returns error to conform to protobuf def, but the error will always be returned as nil.
// Can't return an error with a response object - response object is returned as nil when error is non nil.
// Therefore, return error in response object.
func (service *PlanService) GetUserPlans(ctx context.Context, req *protoPlan.GetUserPlansRequest, res *protoPlan.PlansPageResponse) error {

	var page *protoPlan.PlansPage
	var err error

	switch {
	case req.MarketName == "" && req.Exchange != "":
		// search by userID, exchange, status when no marketName
		page, err = planRepo.FindUserExchangePlansWithStatus(service.DB, req.UserID, req.Status, req.Exchange, req.Page, req.PageSize)
	case req.MarketName != "" && req.Exchange != "":
		// search by userID, exchange, marketName, status
		page, err = planRepo.FindUserExchangePlansWithStatus(service.DB, req.UserID, req.Status, req.Exchange, req.Page, req.PageSize)
	default:
		// search by userID and status
		page, err = planRepo.FindUserPlansWithStatus(service.DB, req.UserID, req.Status, req.Page, req.PageSize)
	}

	switch {
	case err == nil:
		res.Status = response.Success
		res.Data = page
	default:
		res.Status = response.Error
		res.Message = err.Error()
	}

	return nil
}

// We can delete plans that have no filled orders and that are inactive. This becomes an abort plan
// if the plan status is active.
func (service *PlanService) DeletePlan(ctx context.Context, req *protoPlan.DeletePlanRequest, res *protoPlan.PlanResponse) error {
	// pln, err := planRepo.FindPlanSummary(service.DB, req.PlanID)
	// switch {
	// case err == sql.ErrNoRows:
	// 	res.Status = response.Nonentity
	// 	res.Message = fmt.Sprintf("planID not found %s", req.PlanID)
	// 	return nil

	// case err != nil:
	// 	res.Status = response.Error
	// 	res.Message = fmt.Sprintf("unexpected error in DeletePlan: %s", err.Error())
	// 	return nil

	// default:

	// 	switch {
	// 	case pln.ActiveOrderNumber == 0 && pln.Status != plan.Active:
	// 		// we can safely delete this plan from the system because the plan is not in memory
	// 		// (i.e. not active) and the first order of the plan has not been executed
	// 		err = planRepo.DeletePlan(service.DB, req.PlanID)
	// 		if err != nil {
	// 			res.Status = response.Error
	// 			res.Message = err.Error()
	// 		} else {
	// 			pln.Status = plan.Deleted
	// 			res.Status = response.Success
	// 			res.Data = &protoPlan.PlanData{
	// 				Plan: pln,
	// 			}
	// 		}

	// 	case pln.Status == plan.Active:
	// 		pln.Status = plan.PendingAbort
	// 		_, err = planRepo.UpdatePlanStatus(service.DB, req.PlanID, pln.Status)
	// 		if err != nil {
	// 			res.Status = response.Error
	// 			res.Message = err.Error()
	// 			return nil
	// 		}

	// 		// set the plan order status to aborted we are going to use
	// 		// this status in the execution engine to remove order from memory
	// 		pln.Orders[0].Status = status.Aborted
	// 		// publish this revision to the system so the plan order can be removed from execution
	// 		if err := service.publishPlan(ctx, pln, true); err != nil {
	// 			res.Status = response.Error
	// 			res.Message = fmt.Sprintf("failed to remove active plan order from execution: %s", err.Error())
	// 			return nil
	// 		}

	// 		res.Status = response.Success
	// 		res.Data = &protoPlan.PlanData{
	// 			Plan: pln,
	// 		}

	// 	default:
	// 		// what's this?
	// 	}
	// }
	return nil
}

func NewOrder(planID, exchange, orderStatus, timestamp string, ur *protoOrder.UpdateOrderRequest) *protoOrder.Order {
	// collect triggers for this order
	triggers := make([]*protoOrder.Trigger, 0, len(ur.Triggers))
	for j, cond := range ur.Triggers {
		triggerID := uuid.New()
		trigger := protoOrder.Trigger{
			TriggerID:         triggerID.String(),
			TriggerNumber:     uint32(j),
			TriggerTemplateID: cond.TriggerTemplateID,
			OrderID:           ur.OrderID,
			Name:              cond.Name,
			Code:              cond.Code,
			Actions:           cond.Actions,
			Triggered:         false,
			CreatedOn:         timestamp,
			UpdatedOn:         timestamp,
		}
		triggers = append(triggers, &trigger)
	}

	// market name will be Currency-Base: ADA-BTC
	symbolPair := strings.Split(ur.MarketName, "-")
	symbol := symbolPair[1]
	if ur.Side == side.Sell {
		symbol = symbolPair[0]
	}

	return &protoOrder.Order{
		KeyID:                 ur.KeyID,
		OrderID:               ur.OrderID,
		OrderPriority:         ur.OrderPriority,
		OrderType:             ur.OrderType,
		OrderTemplateID:       ur.OrderTemplateID,
		ParentOrderID:         ur.ParentOrderID,
		PlanID:                planID,
		Side:                  ur.Side,
		LimitPrice:            ur.LimitPrice,
		Exchange:              exchange,
		MarketName:            ur.MarketName,
		ActiveCurrencySymbol:  symbol,
		ActiveCurrencyBalance: ur.ActiveCurrencyBalance,
		Status:                orderStatus,
		Triggers:              triggers,
		CreatedOn:             timestamp,
		UpdatedOn:             timestamp,
	}
}

// UpdatePlan returns error to conform to protobuf def, but the error will always be returned as nil.
// Can't return an error with a response object - response object is returned as nil when error is non nil.
// Therefore, return error in response object.
func (service *PlanService) UpdatePlan(ctx context.Context, req *protoPlan.UpdatePlanRequest, res *protoPlan.PlanResponse) error {

	// load current state of plan
	// let's assume the plan has been paused while editing
	pln, err := planRepo.FindPlanWithUnexecutedOrders(service.DB, req.PlanID)

	pln.PlanTemplateID = req.PlanTemplateID
	pln.Status = req.Status
	pln.CloseOnComplete = req.CloseOnComplete

	// fetch all order keys
	keyIDs := make([]string, 0, len(req.Orders))
	for _, or := range req.Orders {
		keyIDs = append(keyIDs, or.KeyID)
	}
	kys, err := service.fetchKeys(keyIDs)
	if err != nil {
		res.Status = response.Error
		res.Message = fmt.Sprintf("ecountered error when fetching keys: %s", err.Error())
		return nil
	}

	currentTimestamp := string(pq.FormatTimestamp(time.Now().UTC()))
	newOrders := make([]*protoOrder.Order, 0)
	allTriggers := make([]*protoOrder.Trigger, 0)
	deleteOrderIDs := make([]string, 0)
	txn, err := service.DB.Begin()
	if err != nil {
		return err
	}

	for _, o := range req.Orders {

		switch o.Action {
		case orderConstants.NewOrder:
			// assign exchange name from key
			for _, ky := range kys {
				if ky.KeyID == o.KeyID {
					if ky.Status != key.Verified {
						res.Status = response.Fail
						res.Message = "using an unverified key!"
						return nil

					}
					newOrder := NewOrder(
						pln.PlanID,
						ky.Exchange,
						status.Inactive,
						currentTimestamp,
						o)

					pln.Orders = append(pln.Orders, newOrder)

					// bulk insert orders
					newOrders = append(newOrders, newOrder)

					// bulk insert triggers
					allTriggers = append(allTriggers, newOrder.Triggers...)
				}
			}

		case orderConstants.UpdateOrder:
			// assign exchange name from key
			for _, ky := range kys {
				if ky.KeyID == o.KeyID {
					if ky.Status != key.Verified {
						res.Status = response.Fail
						res.Message = "using an unverified key!"
						return nil

					}

					for _, uo := range pln.Orders {
						if uo.OrderID == o.OrderID {
							symbolPair := strings.Split(o.MarketName, "-")
							symbol := symbolPair[1]
							if o.Side == side.Sell {
								symbol = symbolPair[0]
							}

							uo.Exchange = ky.Exchange
							uo.KeyID = ky.KeyID
							uo.ParentOrderID = o.ParentOrderID
							//uo.PlanDepth = o.PlanDepth
							uo.MarketName = o.MarketName
							uo.ActiveCurrencySymbol = symbol
							uo.ActiveCurrencyBalance = o.ActiveCurrencyBalance
							uo.OrderPriority = o.OrderPriority
							uo.OrderTemplateID = o.OrderTemplateID
							uo.Side = o.Side
							uo.LimitPrice = o.LimitPrice

							//planRepo.UpdateOrder(txn, ctx, &uo)

							// assumes the trigger requests were updated also, therefore, drop the previous triggers
							planRepo.DeleteTriggersWithOrderID(txn, ctx, uo.OrderID)

							triggers := make([]*protoOrder.Trigger, 0, len(o.Triggers))
							for j, cond := range o.Triggers {
								triggerID := uuid.New()
								trigger := protoOrder.Trigger{
									TriggerID:         triggerID.String(),
									TriggerNumber:     uint32(j),
									TriggerTemplateID: cond.TriggerTemplateID,
									OrderID:           o.OrderID,
									Name:              cond.Name,
									Code:              cond.Code,
									Actions:           cond.Actions,
									Triggered:         false,
									CreatedOn:         currentTimestamp,
									UpdatedOn:         currentTimestamp,
								}
								triggers = append(triggers, &trigger)
							}
							uo.Triggers = triggers
							// append to bulk insert of triggers
							allTriggers = append(allTriggers, triggers...)
							// update the order
						}
					}
				}
			}
		case orderConstants.DeleteOrder:
			for i, do := range pln.Orders {
				if do.OrderID == o.OrderID {
					deleteOrderIDs = append(deleteOrderIDs, o.OrderID)
					// swap with last order and delete it from the plan
					pln.Orders[len(pln.Orders)-1], pln.Orders[i] = pln.Orders[i], pln.Orders[len(pln.Orders)-1]
					pln.Orders = pln.Orders[:len(pln.Orders)-1]
				}
			}
		}
	}

	planRepo.DeleteOrders(txn, ctx, deleteOrderIDs)
	planRepo.InsertOrders(txn, newOrders)
	planRepo.InsertTriggers(txn, allTriggers)

	res.Status = response.Success
	res.Data = &protoPlan.PlanData{Plan: pln}

	return nil
}
