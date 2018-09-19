// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/events/events.proto

/*
Package common_events is a generated protocol buffer package.

It is generated from these files:
	proto/events/events.proto

It has these top-level messages:
	TradeEvent
	TradeEvents
	Auth
	CandleDataRequest
	NewPlanEvent
	Order
	Trigger
	TriggeredOrderEvent
	CompletedOrderEvent
	AbortedOrderEvent
	EngineStartEvent
	DeletedAccountEvent
*/
package common_events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Events
type TradeEvent struct {
	Exchange   string  `protobuf:"bytes,1,opt,name=exchange" json:"exchange,omitempty"`
	Type       string  `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	EventTime  string  `protobuf:"bytes,3,opt,name=eventTime" json:"eventTime,omitempty"`
	MarketName string  `protobuf:"bytes,4,opt,name=marketName" json:"marketName,omitempty"`
	TradeID    string  `protobuf:"bytes,5,opt,name=tradeID" json:"tradeID,omitempty"`
	Price      float64 `protobuf:"fixed64,6,opt,name=price" json:"price,omitempty"`
	Quantity   float64 `protobuf:"fixed64,7,opt,name=quantity" json:"quantity,omitempty"`
	Total      float64 `protobuf:"fixed64,8,opt,name=total" json:"total,omitempty"`
}

func (m *TradeEvent) Reset()                    { *m = TradeEvent{} }
func (m *TradeEvent) String() string            { return proto.CompactTextString(m) }
func (*TradeEvent) ProtoMessage()               {}
func (*TradeEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TradeEvent) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *TradeEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TradeEvent) GetEventTime() string {
	if m != nil {
		return m.EventTime
	}
	return ""
}

func (m *TradeEvent) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *TradeEvent) GetTradeID() string {
	if m != nil {
		return m.TradeID
	}
	return ""
}

func (m *TradeEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TradeEvent) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *TradeEvent) GetTotal() float64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type TradeEvents struct {
	Events []*TradeEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *TradeEvents) Reset()                    { *m = TradeEvents{} }
func (m *TradeEvents) String() string            { return proto.CompactTextString(m) }
func (*TradeEvents) ProtoMessage()               {}
func (*TradeEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TradeEvents) GetEvents() []*TradeEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type Auth struct {
	Key    string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Auth) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Auth) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type CandleDataRequest struct {
	Exchange   string `protobuf:"bytes,1,opt,name=exchange" json:"exchange,omitempty"`
	MarketName string `protobuf:"bytes,2,opt,name=marketName" json:"marketName,omitempty"`
	Interval   string `protobuf:"bytes,3,opt,name=interval" json:"interval,omitempty"`
}

func (m *CandleDataRequest) Reset()                    { *m = CandleDataRequest{} }
func (m *CandleDataRequest) String() string            { return proto.CompactTextString(m) }
func (*CandleDataRequest) ProtoMessage()               {}
func (*CandleDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CandleDataRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *CandleDataRequest) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *CandleDataRequest) GetInterval() string {
	if m != nil {
		return m.Interval
	}
	return ""
}

type NewPlanEvent struct {
	PlanID                string   `protobuf:"bytes,1,opt,name=planID" json:"planID,omitempty"`
	UserID                string   `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	ActiveCurrencySymbol  string   `protobuf:"bytes,3,opt,name=activeCurrencySymbol" json:"activeCurrencySymbol,omitempty"`
	ActiveCurrencyBalance float64  `protobuf:"fixed64,4,opt,name=activeCurrencyBalance" json:"activeCurrencyBalance,omitempty"`
	CloseOnComplete       bool     `protobuf:"varint,5,opt,name=closeOnComplete" json:"closeOnComplete,omitempty"`
	Orders                []*Order `protobuf:"bytes,6,rep,name=orders" json:"orders,omitempty"`
}

func (m *NewPlanEvent) Reset()                    { *m = NewPlanEvent{} }
func (m *NewPlanEvent) String() string            { return proto.CompactTextString(m) }
func (*NewPlanEvent) ProtoMessage()               {}
func (*NewPlanEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NewPlanEvent) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *NewPlanEvent) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *NewPlanEvent) GetActiveCurrencySymbol() string {
	if m != nil {
		return m.ActiveCurrencySymbol
	}
	return ""
}

func (m *NewPlanEvent) GetActiveCurrencyBalance() float64 {
	if m != nil {
		return m.ActiveCurrencyBalance
	}
	return 0
}

func (m *NewPlanEvent) GetCloseOnComplete() bool {
	if m != nil {
		return m.CloseOnComplete
	}
	return false
}

func (m *NewPlanEvent) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Order struct {
	OrderID     string     `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	Exchange    string     `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	MarketName  string     `protobuf:"bytes,3,opt,name=marketName" json:"marketName,omitempty"`
	Side        string     `protobuf:"bytes,4,opt,name=side" json:"side,omitempty"`
	LimitPrice  float64    `protobuf:"fixed64,5,opt,name=limitPrice" json:"limitPrice,omitempty"`
	OrderType   string     `protobuf:"bytes,6,opt,name=orderType" json:"orderType,omitempty"`
	OrderStatus string     `protobuf:"bytes,7,opt,name=orderStatus" json:"orderStatus,omitempty"`
	AccountID   string     `protobuf:"bytes,8,opt,name=accountID" json:"accountID,omitempty"`
	KeyPublic   string     `protobuf:"bytes,9,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret   string     `protobuf:"bytes,10,opt,name=keySecret" json:"keySecret,omitempty"`
	Triggers    []*Trigger `protobuf:"bytes,11,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Order) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *Order) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Order) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *Order) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *Order) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

func (m *Order) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *Order) GetOrderStatus() string {
	if m != nil {
		return m.OrderStatus
	}
	return ""
}

func (m *Order) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Order) GetKeyPublic() string {
	if m != nil {
		return m.KeyPublic
	}
	return ""
}

func (m *Order) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *Order) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type Trigger struct {
	TriggerID string   `protobuf:"bytes,1,opt,name=triggerID" json:"triggerID,omitempty"`
	OrderID   string   `protobuf:"bytes,2,opt,name=orderID" json:"orderID,omitempty"`
	Name      string   `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Code      string   `protobuf:"bytes,6,opt,name=code" json:"code,omitempty"`
	Triggered bool     `protobuf:"varint,7,opt,name=triggered" json:"triggered,omitempty"`
	Actions   []string `protobuf:"bytes,8,rep,name=actions" json:"actions,omitempty"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Trigger) GetTriggerID() string {
	if m != nil {
		return m.TriggerID
	}
	return ""
}

func (m *Trigger) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *Trigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Trigger) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Trigger) GetTriggered() bool {
	if m != nil {
		return m.Triggered
	}
	return false
}

func (m *Trigger) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

type TriggeredOrderEvent struct {
	Exchange              string  `protobuf:"bytes,1,opt,name=exchange" json:"exchange,omitempty"`
	OrderID               string  `protobuf:"bytes,2,opt,name=orderID" json:"orderID,omitempty"`
	PlanID                string  `protobuf:"bytes,3,opt,name=planID" json:"planID,omitempty"`
	UserID                string  `protobuf:"bytes,4,opt,name=userID" json:"userID,omitempty"`
	AccountID             string  `protobuf:"bytes,5,opt,name=accountID" json:"accountID,omitempty"`
	ActiveCurrencySymbol  string  `protobuf:"bytes,6,opt,name=activeCurrencySymbol" json:"activeCurrencySymbol,omitempty"`
	ActiveCurrencyBalance float64 `protobuf:"fixed64,7,opt,name=activeCurrencyBalance" json:"activeCurrencyBalance,omitempty"`
	Quantity              float64 `protobuf:"fixed64,8,opt,name=quantity" json:"quantity,omitempty"`
	KeyPublic             string  `protobuf:"bytes,9,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret             string  `protobuf:"bytes,10,opt,name=keySecret" json:"keySecret,omitempty"`
	MarketName            string  `protobuf:"bytes,11,opt,name=marketName" json:"marketName,omitempty"`
	Side                  string  `protobuf:"bytes,12,opt,name=side" json:"side,omitempty"`
	OrderType             string  `protobuf:"bytes,13,opt,name=orderType" json:"orderType,omitempty"`
	LimitPrice            float64 `protobuf:"fixed64,14,opt,name=limitPrice" json:"limitPrice,omitempty"`
	TriggerID             string  `protobuf:"bytes,15,opt,name=triggerID" json:"triggerID,omitempty"`
	TriggeredPrice        float64 `protobuf:"fixed64,16,opt,name=triggeredPrice" json:"triggeredPrice,omitempty"`
	TriggeredCondition    string  `protobuf:"bytes,17,opt,name=triggeredCondition" json:"triggeredCondition,omitempty"`
}

func (m *TriggeredOrderEvent) Reset()                    { *m = TriggeredOrderEvent{} }
func (m *TriggeredOrderEvent) String() string            { return proto.CompactTextString(m) }
func (*TriggeredOrderEvent) ProtoMessage()               {}
func (*TriggeredOrderEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TriggeredOrderEvent) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *TriggeredOrderEvent) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *TriggeredOrderEvent) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *TriggeredOrderEvent) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TriggeredOrderEvent) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *TriggeredOrderEvent) GetActiveCurrencySymbol() string {
	if m != nil {
		return m.ActiveCurrencySymbol
	}
	return ""
}

func (m *TriggeredOrderEvent) GetActiveCurrencyBalance() float64 {
	if m != nil {
		return m.ActiveCurrencyBalance
	}
	return 0
}

func (m *TriggeredOrderEvent) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *TriggeredOrderEvent) GetKeyPublic() string {
	if m != nil {
		return m.KeyPublic
	}
	return ""
}

func (m *TriggeredOrderEvent) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *TriggeredOrderEvent) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *TriggeredOrderEvent) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *TriggeredOrderEvent) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *TriggeredOrderEvent) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

func (m *TriggeredOrderEvent) GetTriggerID() string {
	if m != nil {
		return m.TriggerID
	}
	return ""
}

func (m *TriggeredOrderEvent) GetTriggeredPrice() float64 {
	if m != nil {
		return m.TriggeredPrice
	}
	return 0
}

func (m *TriggeredOrderEvent) GetTriggeredCondition() string {
	if m != nil {
		return m.TriggeredCondition
	}
	return ""
}

type CompletedOrderEvent struct {
	OrderID                  string  `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	PlanID                   string  `protobuf:"bytes,2,opt,name=planID" json:"planID,omitempty"`
	UserID                   string  `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	Exchange                 string  `protobuf:"bytes,4,opt,name=exchange" json:"exchange,omitempty"`
	MarketName               string  `protobuf:"bytes,5,opt,name=marketName" json:"marketName,omitempty"`
	Side                     string  `protobuf:"bytes,6,opt,name=side" json:"side,omitempty"`
	AccountID                string  `protobuf:"bytes,7,opt,name=accountID" json:"accountID,omitempty"`
	InitialCurrencySymbol    string  `protobuf:"bytes,8,opt,name=initialCurrencySymbol" json:"initialCurrencySymbol,omitempty"`
	InitialCurrencyBalance   float64 `protobuf:"fixed64,9,opt,name=initialCurrencyBalance" json:"initialCurrencyBalance,omitempty"`
	InitialCurrencyTraded    float64 `protobuf:"fixed64,10,opt,name=initialCurrencyTraded" json:"initialCurrencyTraded,omitempty"`
	InitialCurrencyRemainder float64 `protobuf:"fixed64,11,opt,name=initialCurrencyRemainder" json:"initialCurrencyRemainder,omitempty"`
	InitialCurrencyPrice     float64 `protobuf:"fixed64,12,opt,name=initialCurrencyPrice" json:"initialCurrencyPrice,omitempty"`
	FinalCurrencySymbol      string  `protobuf:"bytes,13,opt,name=finalCurrencySymbol" json:"finalCurrencySymbol,omitempty"`
	FinalCurrencyBalance     float64 `protobuf:"fixed64,14,opt,name=finalCurrencyBalance" json:"finalCurrencyBalance,omitempty"`
	FeeCurrencySymbol        string  `protobuf:"bytes,15,opt,name=feeCurrencySymbol" json:"feeCurrencySymbol,omitempty"`
	FeeCurrencyAmount        float64 `protobuf:"fixed64,16,opt,name=feeCurrencyAmount" json:"feeCurrencyAmount,omitempty"`
	TriggerID                string  `protobuf:"bytes,17,opt,name=triggerID" json:"triggerID,omitempty"`
	TriggeredPrice           float64 `protobuf:"fixed64,18,opt,name=triggeredPrice" json:"triggeredPrice,omitempty"`
	TriggeredCondition       string  `protobuf:"bytes,19,opt,name=triggeredCondition" json:"triggeredCondition,omitempty"`
	ExchangeOrderID          string  `protobuf:"bytes,20,opt,name=exchangeOrderID" json:"exchangeOrderID,omitempty"`
	ExchangeMarketName       string  `protobuf:"bytes,21,opt,name=exchangeMarketName" json:"exchangeMarketName,omitempty"`
	ExchangeTime             string  `protobuf:"bytes,22,opt,name=exchangeTime" json:"exchangeTime,omitempty"`
	Status                   string  `protobuf:"bytes,23,opt,name=status" json:"status,omitempty"`
	Details                  string  `protobuf:"bytes,24,opt,name=details" json:"details,omitempty"`
	CloseOnComplete          bool    `protobuf:"varint,25,opt,name=closeOnComplete" json:"closeOnComplete,omitempty"`
}

func (m *CompletedOrderEvent) Reset()                    { *m = CompletedOrderEvent{} }
func (m *CompletedOrderEvent) String() string            { return proto.CompactTextString(m) }
func (*CompletedOrderEvent) ProtoMessage()               {}
func (*CompletedOrderEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CompletedOrderEvent) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *CompletedOrderEvent) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *CompletedOrderEvent) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CompletedOrderEvent) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *CompletedOrderEvent) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *CompletedOrderEvent) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *CompletedOrderEvent) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *CompletedOrderEvent) GetInitialCurrencySymbol() string {
	if m != nil {
		return m.InitialCurrencySymbol
	}
	return ""
}

func (m *CompletedOrderEvent) GetInitialCurrencyBalance() float64 {
	if m != nil {
		return m.InitialCurrencyBalance
	}
	return 0
}

func (m *CompletedOrderEvent) GetInitialCurrencyTraded() float64 {
	if m != nil {
		return m.InitialCurrencyTraded
	}
	return 0
}

func (m *CompletedOrderEvent) GetInitialCurrencyRemainder() float64 {
	if m != nil {
		return m.InitialCurrencyRemainder
	}
	return 0
}

func (m *CompletedOrderEvent) GetInitialCurrencyPrice() float64 {
	if m != nil {
		return m.InitialCurrencyPrice
	}
	return 0
}

func (m *CompletedOrderEvent) GetFinalCurrencySymbol() string {
	if m != nil {
		return m.FinalCurrencySymbol
	}
	return ""
}

func (m *CompletedOrderEvent) GetFinalCurrencyBalance() float64 {
	if m != nil {
		return m.FinalCurrencyBalance
	}
	return 0
}

func (m *CompletedOrderEvent) GetFeeCurrencySymbol() string {
	if m != nil {
		return m.FeeCurrencySymbol
	}
	return ""
}

func (m *CompletedOrderEvent) GetFeeCurrencyAmount() float64 {
	if m != nil {
		return m.FeeCurrencyAmount
	}
	return 0
}

func (m *CompletedOrderEvent) GetTriggerID() string {
	if m != nil {
		return m.TriggerID
	}
	return ""
}

func (m *CompletedOrderEvent) GetTriggeredPrice() float64 {
	if m != nil {
		return m.TriggeredPrice
	}
	return 0
}

func (m *CompletedOrderEvent) GetTriggeredCondition() string {
	if m != nil {
		return m.TriggeredCondition
	}
	return ""
}

func (m *CompletedOrderEvent) GetExchangeOrderID() string {
	if m != nil {
		return m.ExchangeOrderID
	}
	return ""
}

func (m *CompletedOrderEvent) GetExchangeMarketName() string {
	if m != nil {
		return m.ExchangeMarketName
	}
	return ""
}

func (m *CompletedOrderEvent) GetExchangeTime() string {
	if m != nil {
		return m.ExchangeTime
	}
	return ""
}

func (m *CompletedOrderEvent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CompletedOrderEvent) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *CompletedOrderEvent) GetCloseOnComplete() bool {
	if m != nil {
		return m.CloseOnComplete
	}
	return false
}

type AbortedOrderEvent struct {
	OrderID string `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	PlanID  string `protobuf:"bytes,2,opt,name=planID" json:"planID,omitempty"`
}

func (m *AbortedOrderEvent) Reset()                    { *m = AbortedOrderEvent{} }
func (m *AbortedOrderEvent) String() string            { return proto.CompactTextString(m) }
func (*AbortedOrderEvent) ProtoMessage()               {}
func (*AbortedOrderEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AbortedOrderEvent) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *AbortedOrderEvent) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

type EngineStartEvent struct {
	EngineID string `protobuf:"bytes,1,opt,name=engineID" json:"engineID,omitempty"`
}

func (m *EngineStartEvent) Reset()                    { *m = EngineStartEvent{} }
func (m *EngineStartEvent) String() string            { return proto.CompactTextString(m) }
func (*EngineStartEvent) ProtoMessage()               {}
func (*EngineStartEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *EngineStartEvent) GetEngineID() string {
	if m != nil {
		return m.EngineID
	}
	return ""
}

type DeletedAccountEvent struct {
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
}

func (m *DeletedAccountEvent) Reset()                    { *m = DeletedAccountEvent{} }
func (m *DeletedAccountEvent) String() string            { return proto.CompactTextString(m) }
func (*DeletedAccountEvent) ProtoMessage()               {}
func (*DeletedAccountEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeletedAccountEvent) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func init() {
	proto.RegisterType((*TradeEvent)(nil), "common.events.TradeEvent")
	proto.RegisterType((*TradeEvents)(nil), "common.events.TradeEvents")
	proto.RegisterType((*Auth)(nil), "common.events.Auth")
	proto.RegisterType((*CandleDataRequest)(nil), "common.events.CandleDataRequest")
	proto.RegisterType((*NewPlanEvent)(nil), "common.events.NewPlanEvent")
	proto.RegisterType((*Order)(nil), "common.events.Order")
	proto.RegisterType((*Trigger)(nil), "common.events.Trigger")
	proto.RegisterType((*TriggeredOrderEvent)(nil), "common.events.TriggeredOrderEvent")
	proto.RegisterType((*CompletedOrderEvent)(nil), "common.events.CompletedOrderEvent")
	proto.RegisterType((*AbortedOrderEvent)(nil), "common.events.AbortedOrderEvent")
	proto.RegisterType((*EngineStartEvent)(nil), "common.events.EngineStartEvent")
	proto.RegisterType((*DeletedAccountEvent)(nil), "common.events.DeletedAccountEvent")
}

func init() { proto.RegisterFile("proto/events/events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 968 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcb, 0x6e, 0xe3, 0x36,
	0x14, 0x85, 0xfc, 0x90, 0xed, 0xeb, 0xcc, 0xc3, 0x74, 0xe2, 0x32, 0x83, 0xa2, 0x30, 0xb4, 0x28,
	0xb2, 0x18, 0xb8, 0xd3, 0xcc, 0xa0, 0x8b, 0xae, 0x9a, 0xc6, 0x59, 0x64, 0xd1, 0x49, 0xa0, 0xe4,
	0x07, 0x18, 0x89, 0x93, 0x21, 0x22, 0x51, 0x1e, 0x89, 0x4e, 0xeb, 0x5d, 0x7f, 0xa5, 0xdf, 0xd2,
	0xdf, 0x28, 0xd0, 0x8f, 0xe8, 0x0f, 0x14, 0xbc, 0xa4, 0x9e, 0x96, 0x32, 0x68, 0x67, 0x65, 0xdd,
	0x73, 0xc8, 0x4b, 0xeb, 0xf0, 0x1c, 0x8a, 0x70, 0xbc, 0x49, 0x13, 0x95, 0x7c, 0xc7, 0x1f, 0xb9,
	0x54, 0x99, 0xfd, 0x59, 0x21, 0x46, 0x9e, 0x05, 0x49, 0x1c, 0x27, 0x72, 0x65, 0x40, 0xef, 0x2f,
	0x07, 0xe0, 0x36, 0x65, 0x21, 0xbf, 0xd0, 0x35, 0x79, 0x05, 0x63, 0xfe, 0x5b, 0xf0, 0x91, 0xc9,
	0x7b, 0x4e, 0x9d, 0xa5, 0x73, 0x32, 0xf1, 0x8b, 0x9a, 0x10, 0x18, 0xa8, 0xdd, 0x86, 0xd3, 0x1e,
	0xe2, 0xf8, 0x4c, 0xbe, 0x86, 0x09, 0x36, 0xba, 0x15, 0x31, 0xa7, 0x7d, 0x24, 0x4a, 0x80, 0x7c,
	0x03, 0x10, 0xb3, 0xf4, 0x81, 0xab, 0xf7, 0x2c, 0xe6, 0x74, 0x80, 0x74, 0x05, 0x21, 0x14, 0x46,
	0x4a, 0xaf, 0x7d, 0xb9, 0xa6, 0x43, 0x24, 0xf3, 0x92, 0x1c, 0xc2, 0x70, 0x93, 0x8a, 0x80, 0x53,
	0x77, 0xe9, 0x9c, 0x38, 0xbe, 0x29, 0xf4, 0xbf, 0xfb, 0xb4, 0x65, 0x52, 0x09, 0xb5, 0xa3, 0x23,
	0x24, 0x8a, 0x5a, 0xcf, 0x50, 0x89, 0x62, 0x11, 0x1d, 0x9b, 0x19, 0x58, 0x78, 0x3f, 0xc1, 0xb4,
	0x7c, 0xbb, 0x8c, 0x7c, 0x0f, 0xae, 0x79, 0x6f, 0xea, 0x2c, 0xfb, 0x27, 0xd3, 0xd3, 0xe3, 0x55,
	0x4d, 0x8d, 0x55, 0x39, 0xd6, 0xb7, 0x03, 0xbd, 0x37, 0x30, 0x38, 0xdb, 0xaa, 0x8f, 0xe4, 0x25,
	0xf4, 0x1f, 0xf8, 0xce, 0x8a, 0xa2, 0x1f, 0xc9, 0x02, 0xdc, 0x8c, 0x07, 0x29, 0x57, 0x56, 0x11,
	0x5b, 0x79, 0x0f, 0x30, 0x3b, 0x67, 0x32, 0x8c, 0xf8, 0x9a, 0x29, 0xe6, 0xf3, 0x4f, 0x5b, 0x9e,
	0x3d, 0x2d, 0x6c, 0x5d, 0xa6, 0xde, 0x9e, 0x4c, 0xaf, 0x60, 0x2c, 0xa4, 0xe2, 0xe9, 0x23, 0x8b,
	0xac, 0xc6, 0x45, 0xed, 0xfd, 0xde, 0x83, 0x83, 0xf7, 0xfc, 0xd7, 0xeb, 0x88, 0x49, 0xb3, 0x83,
	0x0b, 0x70, 0x37, 0x11, 0x93, 0x97, 0x6b, 0xbb, 0x8c, 0xad, 0x34, 0xbe, 0xcd, 0x78, 0x7a, 0xb9,
	0xce, 0xff, 0xad, 0xa9, 0xc8, 0x29, 0x1c, 0xb2, 0x40, 0x89, 0x47, 0x7e, 0xbe, 0x4d, 0x53, 0x2e,
	0x83, 0xdd, 0xcd, 0x2e, 0xbe, 0x4b, 0xf2, 0x85, 0x5a, 0x39, 0xf2, 0x0e, 0x8e, 0xea, 0xf8, 0xcf,
	0x2c, 0x62, 0x32, 0x30, 0x5b, 0xec, 0xf8, 0xed, 0x24, 0x39, 0x81, 0x17, 0x41, 0x94, 0x64, 0xfc,
	0x4a, 0x9e, 0x27, 0xf1, 0x26, 0xe2, 0x8a, 0xe3, 0xae, 0x8f, 0xfd, 0x26, 0x4c, 0x5e, 0x83, 0x9b,
	0xa4, 0x21, 0x4f, 0x33, 0xea, 0xe2, 0x36, 0x1d, 0x36, 0xb6, 0xe9, 0x4a, 0x93, 0xbe, 0x1d, 0xe3,
	0xfd, 0xdd, 0x83, 0x21, 0x22, 0xda, 0x4f, 0x88, 0x15, 0x2f, 0x9f, 0x97, 0x35, 0xf9, 0x7b, 0x4f,
	0xca, 0xdf, 0xdf, 0x93, 0x9f, 0xc0, 0x20, 0x13, 0x61, 0xee, 0x5f, 0x7c, 0xd6, 0x73, 0x22, 0x11,
	0x0b, 0x75, 0x8d, 0x26, 0x1d, 0xe2, 0x6b, 0x57, 0x10, 0x9d, 0x0b, 0x5c, 0xfa, 0x56, 0x07, 0xc6,
	0x35, 0xb9, 0x28, 0x00, 0xb2, 0x84, 0x29, 0x16, 0x37, 0x8a, 0xa9, 0x6d, 0x86, 0x56, 0x9e, 0xf8,
	0x55, 0x48, 0xcf, 0x67, 0x41, 0x90, 0x6c, 0xa5, 0xba, 0x5c, 0xa3, 0xa3, 0x27, 0x7e, 0x09, 0x68,
	0xf6, 0x81, 0xef, 0xae, 0xb7, 0x77, 0x91, 0x08, 0xe8, 0xc4, 0xb0, 0x05, 0x60, 0xd9, 0x1b, 0x63,
	0x4d, 0x28, 0x58, 0x03, 0x90, 0x53, 0x18, 0xab, 0x54, 0xdc, 0xdf, 0x6b, 0x75, 0xa7, 0xa8, 0xee,
	0x62, 0x2f, 0x04, 0x48, 0xfb, 0xc5, 0x38, 0xef, 0x0f, 0x07, 0x46, 0x16, 0xd5, 0xdd, 0x2d, 0x5e,
	0xa8, 0x5c, 0x02, 0xd5, 0x1d, 0xe8, 0xd5, 0x77, 0x80, 0xc0, 0x40, 0x6a, 0x7d, 0x4d, 0xd0, 0xf1,
	0x59, 0x63, 0x41, 0x12, 0xe6, 0x02, 0xe1, 0x73, 0xa5, 0x3f, 0x0f, 0x51, 0x99, 0xb1, 0x5f, 0x02,
	0xba, 0xbf, 0x36, 0x57, 0x22, 0x33, 0x3a, 0x5e, 0xf6, 0x75, 0x7f, 0x5b, 0x7a, 0x7f, 0x0e, 0x60,
	0x7e, 0x9b, 0x8f, 0x43, 0x3b, 0x7c, 0xfe, 0x44, 0xeb, 0xfe, 0xb7, 0x65, 0x8a, 0xfa, 0x1d, 0x29,
	0x1a, 0xd4, 0x52, 0x54, 0xdb, 0xaf, 0x61, 0x73, 0xbf, 0xba, 0x32, 0xe6, 0xfe, 0x9f, 0x8c, 0x8d,
	0x9e, 0xca, 0x58, 0xf5, 0x84, 0x1c, 0x37, 0x4e, 0xc8, 0x2f, 0x71, 0x4d, 0x3d, 0x23, 0xd3, 0xce,
	0x8c, 0x1c, 0x54, 0x32, 0x52, 0xcb, 0xc0, 0xb3, 0x66, 0x06, 0xea, 0x09, 0x7a, 0xde, 0x96, 0xa0,
	0xd2, 0x67, 0x2f, 0x9a, 0x3e, 0xfb, 0x16, 0x9e, 0x17, 0xa6, 0x30, 0x1d, 0x5e, 0x62, 0x87, 0x06,
	0x4a, 0x56, 0x40, 0x0a, 0xe4, 0x3c, 0x91, 0xa1, 0xd0, 0x66, 0xa1, 0x33, 0x6c, 0xd7, 0xc2, 0x78,
	0xff, 0x8c, 0x60, 0x9e, 0x1f, 0x43, 0x55, 0x17, 0x75, 0x9f, 0x2c, 0xa5, 0x53, 0x7a, 0x1d, 0x4e,
	0xe9, 0xd7, 0x9c, 0x52, 0xf5, 0xe3, 0xe0, 0xc9, 0x93, 0x68, 0xd8, 0xa9, 0xb2, 0x5b, 0x57, 0xb9,
	0x74, 0xde, 0xa8, 0xe9, 0xbc, 0x77, 0x70, 0x24, 0xa4, 0x50, 0x82, 0x45, 0x0d, 0xeb, 0x99, 0x33,
	0xa5, 0x9d, 0x24, 0x3f, 0xc0, 0xa2, 0x41, 0xe4, 0xe6, 0x9b, 0xa0, 0xca, 0x1d, 0x6c, 0xcb, 0x6a,
	0xf8, 0x41, 0x0d, 0xd1, 0x4f, 0x8e, 0xdf, 0x4e, 0x92, 0x1f, 0x81, 0x36, 0x08, 0x9f, 0xc7, 0x4c,
	0xc8, 0x90, 0xa7, 0xe8, 0x34, 0xc7, 0xef, 0xe4, 0x75, 0xb2, 0x1a, 0x9c, 0x71, 0xc3, 0x01, 0xce,
	0x6b, 0xe5, 0xc8, 0x1b, 0x98, 0x7f, 0x10, 0x72, 0x4f, 0x11, 0xe3, 0xd0, 0x36, 0x4a, 0xaf, 0x52,
	0x83, 0x73, 0x35, 0x8c, 0x6b, 0x5b, 0x39, 0xf2, 0x1a, 0x66, 0x1f, 0x78, 0x33, 0xf0, 0xc6, 0xc7,
	0xfb, 0x44, 0x63, 0xf4, 0x59, 0xac, 0xb7, 0xcf, 0x5a, 0x7a, 0x9f, 0xa8, 0x67, 0x63, 0xf6, 0xf9,
	0x6c, 0x90, 0xff, 0x90, 0x8d, 0x79, 0x57, 0x36, 0xf4, 0xf7, 0x3b, 0x77, 0xea, 0x95, 0xcd, 0xc2,
	0x21, 0x0e, 0x6e, 0xc2, 0xba, 0x73, 0x0e, 0xfd, 0x52, 0xfa, 0xf9, 0xc8, 0x74, 0xde, 0x67, 0x88,
	0x07, 0x07, 0x39, 0x8a, 0x17, 0xc9, 0x05, 0x8e, 0xac, 0x61, 0x78, 0xdb, 0x32, 0x9f, 0xcb, 0xaf,
	0xec, 0x6d, 0xcb, 0x7c, 0x29, 0x29, 0x8c, 0x42, 0xae, 0x98, 0x88, 0x32, 0x4a, 0x4d, 0x32, 0x6d,
	0xd9, 0x76, 0xdf, 0x38, 0x6e, 0xbd, 0x6f, 0x78, 0x17, 0x30, 0x3b, 0xbb, 0x4b, 0xd2, 0x2f, 0x8c,
	0xbc, 0xb7, 0x82, 0x97, 0x17, 0xf2, 0x5e, 0x48, 0x7e, 0xa3, 0x58, 0xaa, 0xca, 0xcf, 0x0f, 0x62,
	0x45, 0x9b, 0xa2, 0xf6, 0xde, 0xc2, 0x7c, 0xcd, 0xf1, 0xa4, 0x39, 0x33, 0x81, 0x35, 0x53, 0x6a,
	0x89, 0x76, 0x1a, 0x89, 0xbe, 0x73, 0xf1, 0x1a, 0xff, 0xf6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0xe5, 0xff, 0x05, 0xe3, 0x0b, 0x00, 0x00,
}
