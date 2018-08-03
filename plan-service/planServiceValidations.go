package main

import (
	constOrder "github.com/asciiu/gomo/common/constants/order"
	constPlan "github.com/asciiu/gomo/common/constants/plan"
	constSide "github.com/asciiu/gomo/common/constants/side"
	constStatus "github.com/asciiu/gomo/common/constants/status"
	protoOrder "github.com/asciiu/gomo/plan-service/proto/order"
	"github.com/google/uuid"
)

// New plans can only have a single order with parent order number == 0.
func ValidateSingleRootNode(orderRequests []*protoOrder.NewOrderRequest) bool {
	count := 0
	for _, o := range orderRequests {
		if o.ParentOrderID == uuid.Nil.String() || o.ParentOrderID == "" {
			count += 1
		}
	}
	return count == 1
}

// plan must have at least one order
func ValidateMinOrder(orderRequests []*protoOrder.NewOrderRequest) bool {
	return len(orderRequests) > 0
}

func ValidateConnectedRoutesFromParent(parentOrderID string, orderRequests []*protoOrder.NewOrderRequest) bool {
	orderIDs := make([]string, 0, len(orderRequests)+1)
	orderIDs = append(orderIDs, parentOrderID)

	for _, o := range orderRequests {
		orderIDs = append(orderIDs, o.OrderID)
	}

	for _, o := range orderRequests {
		found := false
		// check connected graph
		for _, n := range orderIDs {
			if o.ParentOrderID == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// All orders must contain triggers
func ValidateOrderTrigger(orderRequests []*protoOrder.NewOrderRequest) bool {
	for _, o := range orderRequests {
		if len(o.Triggers) == 0 {
			return false
		}
	}

	return true
}

// Child orders must have a valid parent order ID
func ValidateChildNodes(orderRequests []*protoOrder.NewOrderRequest) bool {
	for _, o := range orderRequests {
		if o.ParentOrderID == uuid.Nil.String() {
			return false
		}
	}

	return true
}

// limit node count for new requests to 10
func ValidateNodeCount(orderRequests []*protoOrder.NewOrderRequest) bool {
	return len(orderRequests) <= 10
}

// validate non zero currency balance
func ValidateNoneZeroBalance(orderRequests []*protoOrder.NewOrderRequest) bool {
	for _, o := range orderRequests {
		if o.ParentOrderID == uuid.Nil.String() && o.ActiveCurrencyBalance > 0 {
			return true
		}
	}
	return false
}

func ValidatePaperOrders(orderRequests []*protoOrder.NewOrderRequest) bool {
	for _, o := range orderRequests {
		if o.OrderType != constOrder.PaperOrder {
			return false
		}
	}
	return true
}

func ValidateNotPaperOrders(orderRequests []*protoOrder.NewOrderRequest) bool {
	for _, o := range orderRequests {
		if o.OrderType == constOrder.PaperOrder {
			return false
		}
	}
	return true
}

func ValidateOrderType(ot string) bool {
	ots := [...]string{
		constOrder.LimitOrder,
		constOrder.MarketOrder,
		constOrder.PaperOrder,
	}

	for _, ty := range ots {
		if ty == ot {
			return true
		}
	}
	return false
}

func ValidateOrderSide(os string) bool {
	ots := [...]string{
		constSide.Buy,
		constSide.Sell,
	}

	for _, ty := range ots {
		if ty == os {
			return true
		}
	}
	return false
}

// validates user specified plan status
func ValidatePlanInputStatus(pstatus string) bool {
	pistats := [...]string{
		constPlan.Active,
		constPlan.Inactive,
		constPlan.Historic,
	}

	for _, stat := range pistats {
		if stat == pstatus {
			return true
		}
	}
	return false
}

// defines valid input for plan status when updating an executed plan (a.k.a. plan with a filled order)
func ValidatePlanUpdateStatus(pstatus string) bool {
	pistats := [...]string{
		constPlan.Active,
		constPlan.Inactive,
	}

	for _, stat := range pistats {
		if stat == pstatus {
			return true
		}
	}
	return false
}

// New order request cannot overwrite a filled order
func ValidateNonExecutedOrder(porders []*protoOrder.Order, rorders []*protoOrder.NewOrderRequest) bool {
	for _, nor := range rorders {
		for _, po := range porders {
			if nor.OrderID == po.OrderID && po.Status == constStatus.Filled {
				return false
			}
		}
	}
	return true
}
