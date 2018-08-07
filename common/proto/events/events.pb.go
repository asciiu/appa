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
	PlanID string   `protobuf:"bytes,3,opt,name=planID" json:"planID,omitempty"`
	UserID string   `protobuf:"bytes,4,opt,name=userID" json:"userID,omitempty"`
	Orders []*Order `protobuf:"bytes,5,rep,name=orders" json:"orders,omitempty"`
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
	KeyID       string     `protobuf:"bytes,8,opt,name=keyID" json:"keyID,omitempty"`
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

func (m *Order) GetKeyID() string {
	if m != nil {
		return m.KeyID
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
	Exchange           string  `protobuf:"bytes,1,opt,name=exchange" json:"exchange,omitempty"`
	OrderID            string  `protobuf:"bytes,2,opt,name=orderID" json:"orderID,omitempty"`
	PlanID             string  `protobuf:"bytes,3,opt,name=planID" json:"planID,omitempty"`
	UserID             string  `protobuf:"bytes,4,opt,name=userID" json:"userID,omitempty"`
	NextOrderID        string  `protobuf:"bytes,5,opt,name=nextOrderID" json:"nextOrderID,omitempty"`
	BaseBalance        float64 `protobuf:"fixed64,6,opt,name=baseBalance" json:"baseBalance,omitempty"`
	CurrencyBalance    float64 `protobuf:"fixed64,7,opt,name=currencyBalance" json:"currencyBalance,omitempty"`
	Quantity           float64 `protobuf:"fixed64,8,opt,name=quantity" json:"quantity,omitempty"`
	KeyID              string  `protobuf:"bytes,9,opt,name=keyID" json:"keyID,omitempty"`
	Key                string  `protobuf:"bytes,10,opt,name=key" json:"key,omitempty"`
	Secret             string  `protobuf:"bytes,11,opt,name=secret" json:"secret,omitempty"`
	MarketName         string  `protobuf:"bytes,12,opt,name=marketName" json:"marketName,omitempty"`
	Side               string  `protobuf:"bytes,13,opt,name=side" json:"side,omitempty"`
	OrderType          string  `protobuf:"bytes,14,opt,name=orderType" json:"orderType,omitempty"`
	Price              float64 `protobuf:"fixed64,15,opt,name=price" json:"price,omitempty"`
	TriggeredPrice     float64 `protobuf:"fixed64,16,opt,name=triggeredPrice" json:"triggeredPrice,omitempty"`
	TriggeredCondition string  `protobuf:"bytes,17,opt,name=triggeredCondition" json:"triggeredCondition,omitempty"`
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

func (m *TriggeredOrderEvent) GetNextOrderID() string {
	if m != nil {
		return m.NextOrderID
	}
	return ""
}

func (m *TriggeredOrderEvent) GetBaseBalance() float64 {
	if m != nil {
		return m.BaseBalance
	}
	return 0
}

func (m *TriggeredOrderEvent) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *TriggeredOrderEvent) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *TriggeredOrderEvent) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *TriggeredOrderEvent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TriggeredOrderEvent) GetSecret() string {
	if m != nil {
		return m.Secret
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

func (m *TriggeredOrderEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
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
	OrderID            string  `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	PlanID             string  `protobuf:"bytes,2,opt,name=planID" json:"planID,omitempty"`
	UserID             string  `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	NextOrderID        string  `protobuf:"bytes,4,opt,name=nextOrderID" json:"nextOrderID,omitempty"`
	BaseBalance        float64 `protobuf:"fixed64,5,opt,name=baseBalance" json:"baseBalance,omitempty"`
	CurrencyBalance    float64 `protobuf:"fixed64,6,opt,name=currencyBalance" json:"currencyBalance,omitempty"`
	Currency           string  `protobuf:"bytes,7,opt,name=currency" json:"currency,omitempty"`
	Side               string  `protobuf:"bytes,8,opt,name=side" json:"side,omitempty"`
	Quantity           float64 `protobuf:"fixed64,9,opt,name=quantity" json:"quantity,omitempty"`
	TriggeredPrice     float64 `protobuf:"fixed64,10,opt,name=triggeredPrice" json:"triggeredPrice,omitempty"`
	TriggeredCondition string  `protobuf:"bytes,11,opt,name=triggeredCondition" json:"triggeredCondition,omitempty"`
	ExchangeOrderID    string  `protobuf:"bytes,12,opt,name=exchangeOrderID" json:"exchangeOrderID,omitempty"`
	ExchangeMarketName string  `protobuf:"bytes,13,opt,name=exchangeMarketName" json:"exchangeMarketName,omitempty"`
	Status             string  `protobuf:"bytes,14,opt,name=status" json:"status,omitempty"`
	Details            string  `protobuf:"bytes,15,opt,name=details" json:"details,omitempty"`
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

func (m *CompletedOrderEvent) GetNextOrderID() string {
	if m != nil {
		return m.NextOrderID
	}
	return ""
}

func (m *CompletedOrderEvent) GetBaseBalance() float64 {
	if m != nil {
		return m.BaseBalance
	}
	return 0
}

func (m *CompletedOrderEvent) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *CompletedOrderEvent) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *CompletedOrderEvent) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *CompletedOrderEvent) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
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
}

func init() { proto.RegisterFile("proto/events/events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4b, 0x6e, 0xdb, 0x3a,
	0x14, 0x85, 0x7f, 0xb2, 0x74, 0x9d, 0x2f, 0x13, 0x04, 0x4c, 0xf0, 0xf0, 0x60, 0x68, 0xf0, 0xe0,
	0xc1, 0x83, 0xdb, 0xa6, 0x1b, 0x68, 0x1a, 0x67, 0xe0, 0x41, 0x93, 0x40, 0xf1, 0x06, 0x68, 0x89,
	0x70, 0x04, 0xcb, 0x94, 0x23, 0xd1, 0x69, 0xbc, 0x80, 0x2e, 0xa4, 0xab, 0xea, 0xa4, 0x5b, 0x29,
	0x50, 0xf0, 0x92, 0xa2, 0x64, 0x3b, 0x4e, 0x1a, 0x74, 0x64, 0x9e, 0x73, 0xa9, 0x2b, 0xf1, 0x9c,
	0xc3, 0x0b, 0xc3, 0xe9, 0x3c, 0x4b, 0x65, 0xfa, 0x8e, 0x3f, 0x72, 0x21, 0x73, 0xf3, 0xd3, 0x47,
	0x8e, 0xec, 0x86, 0xe9, 0x6c, 0x96, 0x8a, 0xbe, 0x26, 0xfd, 0x9f, 0x35, 0x80, 0x51, 0xc6, 0x22,
	0x7e, 0xa5, 0x30, 0x39, 0x03, 0x97, 0x3f, 0x85, 0xf7, 0x4c, 0x4c, 0x38, 0xad, 0x75, 0x6b, 0x3d,
	0x2f, 0xb0, 0x98, 0x10, 0x68, 0xca, 0xe5, 0x9c, 0xd3, 0x3a, 0xf2, 0xb8, 0x26, 0xff, 0x80, 0x87,
	0x8d, 0x46, 0xf1, 0x8c, 0xd3, 0x06, 0x16, 0x4a, 0x82, 0xfc, 0x0b, 0x30, 0x63, 0xd9, 0x94, 0xcb,
	0x6b, 0x36, 0xe3, 0xb4, 0x89, 0xe5, 0x0a, 0x43, 0x28, 0xb4, 0xa5, 0x7a, 0xf7, 0x70, 0x40, 0x5b,
	0x58, 0x2c, 0x20, 0x39, 0x86, 0xd6, 0x3c, 0x8b, 0x43, 0x4e, 0x9d, 0x6e, 0xad, 0x57, 0x0b, 0x34,
	0x50, 0x5f, 0xf7, 0xb0, 0x60, 0x42, 0xc6, 0x72, 0x49, 0xdb, 0x58, 0xb0, 0x58, 0x3d, 0x21, 0x53,
	0xc9, 0x12, 0xea, 0xea, 0x27, 0x10, 0xf8, 0x9f, 0xa0, 0x53, 0x9e, 0x2e, 0x27, 0x1f, 0xc0, 0xd1,
	0xe7, 0xa6, 0xb5, 0x6e, 0xa3, 0xd7, 0x39, 0x3f, 0xed, 0xaf, 0xa8, 0xd1, 0x2f, 0xf7, 0x06, 0x66,
	0xa3, 0xff, 0x1e, 0x9a, 0x17, 0x0b, 0x79, 0x4f, 0x0e, 0xa0, 0x31, 0xe5, 0x4b, 0x23, 0x8a, 0x5a,
	0x92, 0x13, 0x70, 0x72, 0x1e, 0x66, 0x5c, 0x1a, 0x45, 0x0c, 0xf2, 0xa7, 0x70, 0x78, 0xc9, 0x44,
	0x94, 0xf0, 0x01, 0x93, 0x2c, 0xe0, 0x0f, 0x0b, 0x9e, 0xbf, 0x2c, 0xec, 0xaa, 0x4c, 0xf5, 0x0d,
	0x99, 0xce, 0xc0, 0x8d, 0x85, 0xe4, 0xd9, 0x23, 0x4b, 0x8c, 0xc6, 0x16, 0xfb, 0x09, 0xec, 0x5c,
	0xf3, 0xaf, 0xb7, 0x09, 0x13, 0xda, 0xc0, 0x13, 0x70, 0xe6, 0x09, 0x13, 0xc3, 0x81, 0xd9, 0x69,
	0x90, 0xe2, 0x17, 0x39, 0xcf, 0x86, 0x03, 0x63, 0x83, 0x41, 0xe4, 0x7f, 0x70, 0xd2, 0x2c, 0xe2,
	0x59, 0x4e, 0x5b, 0xa8, 0xc8, 0xf1, 0x9a, 0x22, 0x37, 0xaa, 0x18, 0x98, 0x3d, 0xfe, 0x8f, 0x3a,
	0xb4, 0x90, 0x51, 0xd6, 0x21, 0x37, 0x1c, 0x98, 0xe3, 0x14, 0x70, 0xe5, 0xa4, 0xf5, 0x17, 0x4f,
	0xda, 0xd8, 0x38, 0x29, 0x81, 0x66, 0x1e, 0x47, 0x45, 0x54, 0x70, 0xad, 0x9e, 0x49, 0xe2, 0x59,
	0x2c, 0x6f, 0x31, 0x0f, 0x2d, 0x74, 0xb7, 0xc2, 0xa8, 0x08, 0xe2, 0xab, 0x47, 0x2a, 0x9b, 0x8e,
	0x8e, 0xa0, 0x25, 0x48, 0x17, 0x3a, 0x08, 0xee, 0x24, 0x93, 0x8b, 0x1c, 0x53, 0xe3, 0x05, 0x55,
	0x4a, 0x05, 0x67, 0xca, 0x97, 0xc3, 0x01, 0x06, 0xc7, 0x0b, 0x34, 0x50, 0x5d, 0xa7, 0x7c, 0x79,
	0xbb, 0x18, 0x27, 0x71, 0x48, 0x3d, 0xdd, 0xd5, 0x12, 0xa6, 0x7a, 0xa7, 0xdd, 0x07, 0x5b, 0xd5,
	0x04, 0x39, 0x07, 0x57, 0x66, 0xf1, 0x64, 0xa2, 0x54, 0xed, 0xa0, 0xaa, 0x27, 0x1b, 0x39, 0xc3,
	0x72, 0x60, 0xf7, 0xf9, 0xdf, 0x6b, 0xd0, 0x36, 0xac, 0xea, 0x6e, 0x78, 0xab, 0x6e, 0x49, 0x54,
	0x95, 0xaf, 0xaf, 0x2a, 0x4f, 0xa0, 0x29, 0x94, 0xae, 0xfa, 0x2e, 0xe1, 0x5a, 0x71, 0x61, 0x1a,
	0x15, 0xc2, 0xe0, 0xba, 0xd2, 0x9f, 0x47, 0xa8, 0x88, 0x1b, 0x94, 0x84, 0xea, 0xcf, 0x42, 0x19,
	0xa7, 0x22, 0xa7, 0x6e, 0xb7, 0xa1, 0xfa, 0x1b, 0xe8, 0x7f, 0x6b, 0xc2, 0xd1, 0xa8, 0xd8, 0x87,
	0x31, 0x78, 0x7d, 0x68, 0x6c, 0xff, 0xda, 0xb7, 0x26, 0xb5, 0x0b, 0x1d, 0xc1, 0x9f, 0xe4, 0x8d,
	0xe9, 0xa6, 0x0f, 0x59, 0xa5, 0xd4, 0x8e, 0x31, 0xcb, 0xf9, 0x67, 0x96, 0x30, 0x61, 0x47, 0x47,
	0x95, 0x22, 0x3d, 0xd8, 0x0f, 0x17, 0x59, 0xc6, 0x45, 0xb8, 0x2c, 0x76, 0xe9, 0x39, 0xb2, 0x4e,
	0xaf, 0x8c, 0x1a, 0x77, 0x73, 0xd4, 0xe8, 0xc4, 0x78, 0xd5, 0xc4, 0x98, 0x01, 0x01, 0xcf, 0x0d,
	0x88, 0x4e, 0x75, 0x40, 0xac, 0xdd, 0x82, 0x9d, 0xad, 0xb7, 0x60, 0xb7, 0x72, 0x0b, 0x56, 0x52,
	0xbe, 0xb7, 0x9e, 0x72, 0x3b, 0x2e, 0xf7, 0xab, 0xe3, 0xf2, 0x3f, 0xd8, 0xb3, 0xb6, 0xea, 0xdb,
	0x73, 0x80, 0xe5, 0x35, 0x96, 0xf4, 0x81, 0x58, 0xe6, 0x32, 0x15, 0x51, 0xac, 0xec, 0xa6, 0x87,
	0xf8, 0x92, 0x67, 0x2a, 0xfe, 0xaf, 0x06, 0x1c, 0x5d, 0xa6, 0xb3, 0x79, 0xc2, 0xe5, 0x4a, 0x0e,
	0xb6, 0xcf, 0x84, 0xd2, 0xeb, 0xfa, 0x16, 0xaf, 0x1b, 0x2f, 0x79, 0xdd, 0x7c, 0xd5, 0xeb, 0xd6,
	0x1f, 0x79, 0xed, 0x6c, 0xf5, 0xba, 0xa0, 0xcc, 0x80, 0xb0, 0xd8, 0x7a, 0xe1, 0x56, 0xbc, 0xa8,
	0x66, 0xc3, 0x5b, 0xcb, 0xc6, 0xa6, 0xe6, 0xf0, 0x06, 0xcd, 0x3b, 0xdb, 0x34, 0x57, 0xa7, 0x29,
	0xee, 0x54, 0xa1, 0x8a, 0x0e, 0xce, 0x3a, 0xad, 0x3a, 0x17, 0xd4, 0x97, 0x32, 0x65, 0x3a, 0x4b,
	0xcf, 0x54, 0x30, 0xa5, 0x7a, 0x38, 0xee, 0x99, 0x94, 0xea, 0xb9, 0x48, 0xa1, 0x1d, 0x71, 0xc9,
	0xe2, 0x24, 0xc7, 0x54, 0x79, 0x41, 0x01, 0xfd, 0x2b, 0x38, 0xbc, 0x18, 0xa7, 0xd9, 0x5f, 0x9a,
	0xef, 0xf7, 0xe1, 0xe0, 0x4a, 0x4c, 0x62, 0xc1, 0xef, 0x24, 0xcb, 0x64, 0x39, 0x4a, 0x90, 0xb3,
	0x6d, 0x2c, 0x1e, 0x3b, 0xf8, 0x07, 0xe6, 0xe3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x36,
	0x33, 0x8c, 0xdd, 0x08, 0x00, 0x00,
}
