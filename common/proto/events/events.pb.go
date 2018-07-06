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
	OrderEvent
	ActivateOrderEvent
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

type OrderEvent struct {
	Exchange           string  `protobuf:"bytes,1,opt,name=exchange" json:"exchange,omitempty"`
	OrderID            string  `protobuf:"bytes,2,opt,name=orderID" json:"orderID,omitempty"`
	UserID             string  `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	KeyID              string  `protobuf:"bytes,4,opt,name=keyID" json:"keyID,omitempty"`
	Key                string  `protobuf:"bytes,5,opt,name=key" json:"key,omitempty"`
	Secret             string  `protobuf:"bytes,6,opt,name=secret" json:"secret,omitempty"`
	MarketName         string  `protobuf:"bytes,7,opt,name=marketName" json:"marketName,omitempty"`
	Side               string  `protobuf:"bytes,8,opt,name=side" json:"side,omitempty"`
	OrderType          string  `protobuf:"bytes,9,opt,name=orderType" json:"orderType,omitempty"`
	Currency           string  `protobuf:"bytes,10,opt,name=currency" json:"currency,omitempty"`
	Quantity           float64 `protobuf:"fixed64,11,opt,name=quantity" json:"quantity,omitempty"`
	Price              float64 `protobuf:"fixed64,12,opt,name=price" json:"price,omitempty"`
	Condition          string  `protobuf:"bytes,13,opt,name=condition" json:"condition,omitempty"`
	Conditions         string  `protobuf:"bytes,14,opt,name=conditions" json:"conditions,omitempty"`
	ExchangeOrderID    string  `protobuf:"bytes,15,opt,name=exchangeOrderID" json:"exchangeOrderID,omitempty"`
	ExchangeMarketName string  `protobuf:"bytes,16,opt,name=exchangeMarketName" json:"exchangeMarketName,omitempty"`
	Status             string  `protobuf:"bytes,17,opt,name=status" json:"status,omitempty"`
}

func (m *OrderEvent) Reset()                    { *m = OrderEvent{} }
func (m *OrderEvent) String() string            { return proto.CompactTextString(m) }
func (*OrderEvent) ProtoMessage()               {}
func (*OrderEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OrderEvent) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *OrderEvent) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *OrderEvent) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *OrderEvent) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *OrderEvent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *OrderEvent) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *OrderEvent) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *OrderEvent) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *OrderEvent) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *OrderEvent) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *OrderEvent) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderEvent) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *OrderEvent) GetConditions() string {
	if m != nil {
		return m.Conditions
	}
	return ""
}

func (m *OrderEvent) GetExchangeOrderID() string {
	if m != nil {
		return m.ExchangeOrderID
	}
	return ""
}

func (m *OrderEvent) GetExchangeMarketName() string {
	if m != nil {
		return m.ExchangeMarketName
	}
	return ""
}

func (m *OrderEvent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ActivateOrderEvent struct {
	Exchange        string  `protobuf:"bytes,1,opt,name=exchange" json:"exchange,omitempty"`
	OrderID         string  `protobuf:"bytes,2,opt,name=orderID" json:"orderID,omitempty"`
	PlanID          string  `protobuf:"bytes,3,opt,name=planID" json:"planID,omitempty"`
	UserID          string  `protobuf:"bytes,4,opt,name=userID" json:"userID,omitempty"`
	BaseBalance     float64 `protobuf:"fixed64,5,opt,name=baseBalance" json:"baseBalance,omitempty"`
	CurrencyBalance float64 `protobuf:"fixed64,6,opt,name=currencyBalance" json:"currencyBalance,omitempty"`
	KeyID           string  `protobuf:"bytes,7,opt,name=keyID" json:"keyID,omitempty"`
	Key             string  `protobuf:"bytes,8,opt,name=key" json:"key,omitempty"`
	Secret          string  `protobuf:"bytes,9,opt,name=secret" json:"secret,omitempty"`
	MarketName      string  `protobuf:"bytes,10,opt,name=marketName" json:"marketName,omitempty"`
	Side            string  `protobuf:"bytes,11,opt,name=side" json:"side,omitempty"`
	OrderType       string  `protobuf:"bytes,12,opt,name=orderType" json:"orderType,omitempty"`
	Price           float64 `protobuf:"fixed64,14,opt,name=price" json:"price,omitempty"`
	Conditions      string  `protobuf:"bytes,16,opt,name=conditions" json:"conditions,omitempty"`
	Status          string  `protobuf:"bytes,19,opt,name=status" json:"status,omitempty"`
}

func (m *ActivateOrderEvent) Reset()                    { *m = ActivateOrderEvent{} }
func (m *ActivateOrderEvent) String() string            { return proto.CompactTextString(m) }
func (*ActivateOrderEvent) ProtoMessage()               {}
func (*ActivateOrderEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ActivateOrderEvent) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *ActivateOrderEvent) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *ActivateOrderEvent) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *ActivateOrderEvent) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ActivateOrderEvent) GetBaseBalance() float64 {
	if m != nil {
		return m.BaseBalance
	}
	return 0
}

func (m *ActivateOrderEvent) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *ActivateOrderEvent) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *ActivateOrderEvent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ActivateOrderEvent) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ActivateOrderEvent) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *ActivateOrderEvent) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *ActivateOrderEvent) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *ActivateOrderEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ActivateOrderEvent) GetConditions() string {
	if m != nil {
		return m.Conditions
	}
	return ""
}

func (m *ActivateOrderEvent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type EngineStartEvent struct {
	EngineID string `protobuf:"bytes,1,opt,name=engineID" json:"engineID,omitempty"`
}

func (m *EngineStartEvent) Reset()                    { *m = EngineStartEvent{} }
func (m *EngineStartEvent) String() string            { return proto.CompactTextString(m) }
func (*EngineStartEvent) ProtoMessage()               {}
func (*EngineStartEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
	proto.RegisterType((*OrderEvent)(nil), "common.events.OrderEvent")
	proto.RegisterType((*ActivateOrderEvent)(nil), "common.events.ActivateOrderEvent")
	proto.RegisterType((*EngineStartEvent)(nil), "common.events.EngineStartEvent")
}

func init() { proto.RegisterFile("proto/events/events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x6e, 0xdb, 0x30,
	0x10, 0x85, 0x6a, 0xc7, 0xb6, 0x46, 0xf9, 0xb8, 0x6c, 0x51, 0x30, 0x45, 0x50, 0x18, 0x5a, 0x79,
	0xa5, 0xfe, 0x2e, 0xd0, 0x14, 0xce, 0xc2, 0x8b, 0x36, 0x80, 0xea, 0x0b, 0x30, 0xf2, 0x20, 0x11,
	0x6c, 0x51, 0x2e, 0x45, 0x05, 0xd5, 0x85, 0x7a, 0xaa, 0x9e, 0xa0, 0xa7, 0x28, 0x38, 0xa4, 0x7e,
	0xb6, 0xd3, 0x6e, 0xba, 0x32, 0xdf, 0x9b, 0x11, 0x35, 0xf3, 0xde, 0xb3, 0xe0, 0x72, 0xa7, 0x72,
	0x9d, 0xbf, 0xc5, 0x47, 0x94, 0xba, 0x70, 0x3f, 0x11, 0x71, 0xec, 0x2c, 0xc9, 0xb3, 0x2c, 0x97,
	0x91, 0x25, 0xc3, 0x5f, 0x1e, 0xc0, 0x4a, 0x89, 0x35, 0xde, 0x18, 0xcc, 0x5e, 0xc3, 0x04, 0x7f,
	0x24, 0x0f, 0x42, 0xde, 0x23, 0xf7, 0x66, 0xde, 0xdc, 0x8f, 0x1b, 0xcc, 0x18, 0x0c, 0x75, 0xb5,
	0x43, 0xfe, 0x8c, 0x78, 0x3a, 0xb3, 0x2b, 0xf0, 0xe9, 0xa2, 0x55, 0x9a, 0x21, 0x1f, 0x50, 0xa1,
	0x25, 0xd8, 0x1b, 0x80, 0x4c, 0xa8, 0x0d, 0xea, 0xaf, 0x22, 0x43, 0x3e, 0xa4, 0x72, 0x87, 0x61,
	0x1c, 0xc6, 0xda, 0xbc, 0x7b, 0xb9, 0xe0, 0x27, 0x54, 0xac, 0x21, 0x7b, 0x09, 0x27, 0x3b, 0x95,
	0x26, 0xc8, 0x47, 0x33, 0x6f, 0xee, 0xc5, 0x16, 0x98, 0xe9, 0xbe, 0x97, 0x42, 0xea, 0x54, 0x57,
	0x7c, 0x4c, 0x85, 0x06, 0x9b, 0x27, 0x74, 0xae, 0xc5, 0x96, 0x4f, 0xec, 0x13, 0x04, 0xc2, 0x4f,
	0x10, 0xb4, 0xdb, 0x15, 0xec, 0x3d, 0x8c, 0xec, 0xde, 0xdc, 0x9b, 0x0d, 0xe6, 0xc1, 0x87, 0xcb,
	0xa8, 0xa7, 0x46, 0xd4, 0xf6, 0xc6, 0xae, 0x31, 0x7c, 0x07, 0xc3, 0xeb, 0x52, 0x3f, 0xb0, 0x29,
	0x0c, 0x36, 0x58, 0x39, 0x51, 0xcc, 0x91, 0xbd, 0x82, 0x51, 0x81, 0x89, 0x42, 0xed, 0x14, 0x71,
	0x28, 0xfc, 0x3d, 0x00, 0xb8, 0x55, 0x6b, 0x54, 0xff, 0x96, 0x94, 0xc3, 0x38, 0x37, 0x9d, 0xcb,
	0x85, 0xbb, 0xa3, 0x86, 0xe6, 0xf2, 0xb2, 0xa0, 0x82, 0x55, 0xd5, 0x21, 0xb3, 0xe6, 0x06, 0xab,
	0xe5, 0xc2, 0xa9, 0x69, 0x41, 0x3d, 0xdc, 0xc9, 0xb1, 0xe1, 0x46, 0xdd, 0xe1, 0xf6, 0x2c, 0x19,
	0x1f, 0x58, 0xc2, 0x60, 0x58, 0xa4, 0x6b, 0x24, 0x15, 0xfd, 0x98, 0xce, 0xc6, 0x64, 0x1a, 0x6b,
	0x65, 0xdc, 0xf7, 0xad, 0xc9, 0x0d, 0x61, 0xf6, 0x4b, 0x4a, 0xa5, 0x50, 0x26, 0x15, 0x07, 0xbb,
	0x5f, 0x8d, 0x7b, 0x86, 0x05, 0x87, 0x86, 0x59, 0x8b, 0x4f, 0xbb, 0x16, 0x5f, 0x81, 0x9f, 0xe4,
	0x72, 0x9d, 0xea, 0x34, 0x97, 0xfc, 0xcc, 0xbe, 0xab, 0x21, 0xcc, 0xf4, 0x0d, 0x28, 0xf8, 0xb9,
	0x9d, 0xbe, 0x65, 0xd8, 0x1c, 0x2e, 0x6a, 0x6d, 0x6f, 0x9d, 0xae, 0x17, 0xd4, 0xb4, 0x4f, 0xb3,
	0x08, 0x58, 0x4d, 0x7d, 0x69, 0xf5, 0x98, 0x52, 0xf3, 0x91, 0x0a, 0xe9, 0xa9, 0x85, 0x2e, 0x0b,
	0xfe, 0xdc, 0xe9, 0x49, 0x28, 0xfc, 0x39, 0x00, 0x76, 0x9d, 0xe8, 0xf4, 0x51, 0x68, 0xfc, 0x1f,
	0xa6, 0xef, 0xb6, 0x42, 0xb6, 0xa6, 0x5b, 0xd4, 0x09, 0xc3, 0xb0, 0x17, 0x86, 0x19, 0x04, 0x77,
	0xa2, 0xc0, 0xcf, 0x62, 0x2b, 0x64, 0x82, 0x64, 0xbf, 0x17, 0x77, 0x29, 0x23, 0x48, 0x6d, 0x46,
	0xdd, 0x65, 0xff, 0x51, 0xfb, 0x74, 0x1b, 0xac, 0xf1, 0x91, 0x60, 0x4d, 0x8e, 0x05, 0xcb, 0xff,
	0x4b, 0xb0, 0xe0, 0xc9, 0x60, 0x05, 0x4f, 0x05, 0xeb, 0x74, 0x3f, 0x58, 0x4d, 0x40, 0xce, 0xbb,
	0x01, 0xe9, 0x47, 0x60, 0x7a, 0x10, 0x81, 0xd6, 0xa8, 0x17, 0x3d, 0xa3, 0x22, 0x98, 0xde, 0xc8,
	0xfb, 0x54, 0xe2, 0x37, 0x2d, 0x94, 0x6e, 0x5d, 0x22, 0x6e, 0xb9, 0x68, 0x5c, 0x72, 0xf8, 0x6e,
	0x44, 0x9f, 0xcb, 0x8f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x07, 0x43, 0xa0, 0x61, 0x4b, 0x05,
	0x00, 0x00,
}
