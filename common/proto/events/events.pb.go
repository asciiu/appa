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

type EngineStartEvent struct {
	EngineID string `protobuf:"bytes,1,opt,name=engineID" json:"engineID,omitempty"`
}

func (m *EngineStartEvent) Reset()                    { *m = EngineStartEvent{} }
func (m *EngineStartEvent) String() string            { return proto.CompactTextString(m) }
func (*EngineStartEvent) ProtoMessage()               {}
func (*EngineStartEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto.RegisterType((*EngineStartEvent)(nil), "common.events.EngineStartEvent")
}

func init() { proto.RegisterFile("proto/events/events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0x92, 0x38, 0xf5, 0x84, 0xd2, 0xb0, 0x42, 0x68, 0x8b, 0x2a, 0x14, 0xf9, 0x29,
	0x4f, 0xe6, 0xdf, 0x05, 0x40, 0x4a, 0x1f, 0xf2, 0x00, 0x95, 0x4c, 0x2e, 0xb0, 0xd8, 0xa3, 0xd6,
	0x0a, 0xde, 0x0d, 0xeb, 0x31, 0xc2, 0x77, 0xe5, 0x04, 0x9c, 0x02, 0xed, 0xec, 0xda, 0x4e, 0xdb,
	0x48, 0x7d, 0xca, 0xfe, 0xbe, 0x19, 0xef, 0xce, 0x7c, 0x9f, 0x02, 0x97, 0x07, 0x6b, 0xc8, 0xbc,
	0xc3, 0xdf, 0xa8, 0xa9, 0x09, 0x3f, 0x19, 0x6b, 0xe2, 0xbc, 0x30, 0x75, 0x6d, 0x74, 0xe6, 0xc5,
	0xf4, 0x6f, 0x04, 0xb0, 0xb3, 0xaa, 0xc4, 0x6b, 0xc7, 0xe2, 0x0d, 0x9c, 0xe1, 0x9f, 0xe2, 0x4e,
	0xe9, 0x5b, 0x94, 0xd1, 0x2a, 0x5a, 0x27, 0xf9, 0xc0, 0x42, 0xc0, 0x94, 0xba, 0x03, 0xca, 0x67,
	0xac, 0xf3, 0x59, 0x5c, 0x41, 0xc2, 0x17, 0xed, 0xaa, 0x1a, 0xe5, 0x84, 0x0b, 0xa3, 0x20, 0xde,
	0x02, 0xd4, 0xca, 0xee, 0x91, 0xbe, 0xa9, 0x1a, 0xe5, 0x94, 0xcb, 0x47, 0x8a, 0x90, 0x30, 0x27,
	0xf7, 0xf6, 0x76, 0x23, 0x67, 0x5c, 0xec, 0x51, 0xbc, 0x82, 0xd9, 0xc1, 0x56, 0x05, 0xca, 0x78,
	0x15, 0xad, 0xa3, 0xdc, 0x83, 0x9b, 0xee, 0x57, 0xab, 0x34, 0x55, 0xd4, 0xc9, 0x39, 0x17, 0x06,
	0x76, 0x5f, 0x90, 0x21, 0xf5, 0x53, 0x9e, 0xf9, 0x2f, 0x18, 0xd2, 0xcf, 0xb0, 0x18, 0xb7, 0x6b,
	0xc4, 0x07, 0x88, 0xfd, 0xde, 0x32, 0x5a, 0x4d, 0xd6, 0x8b, 0x8f, 0x97, 0xd9, 0x3d, 0x37, 0xb2,
	0xb1, 0x37, 0x0f, 0x8d, 0xe9, 0x7b, 0x98, 0x7e, 0x69, 0xe9, 0x4e, 0x2c, 0x61, 0xb2, 0xc7, 0x2e,
	0x98, 0xe2, 0x8e, 0xe2, 0x35, 0xc4, 0x0d, 0x16, 0x16, 0x29, 0x38, 0x12, 0x28, 0xfd, 0x37, 0x01,
	0xb8, 0xb1, 0x25, 0xda, 0xa7, 0x2d, 0x95, 0x30, 0x37, 0xae, 0x73, 0xbb, 0x09, 0x77, 0xf4, 0xe8,
	0x2e, 0x6f, 0x1b, 0x2e, 0x78, 0x57, 0x03, 0xb9, 0x35, 0xf7, 0xd8, 0x6d, 0x37, 0xc1, 0x4d, 0x0f,
	0xfd, 0x70, 0xb3, 0x53, 0xc3, 0xc5, 0xc7, 0xc3, 0x3d, 0x88, 0x64, 0xfe, 0x28, 0x12, 0x01, 0xd3,
	0xa6, 0x2a, 0x91, 0x5d, 0x4c, 0x72, 0x3e, 0xbb, 0x90, 0x79, 0xac, 0x9d, 0x4b, 0x3f, 0xf1, 0x21,
	0x0f, 0x82, 0xdb, 0xaf, 0x68, 0xad, 0x45, 0x5d, 0x74, 0x12, 0xfc, 0x7e, 0x3d, 0xdf, 0x0b, 0x6c,
	0xf1, 0x38, 0x30, 0x1f, 0xf1, 0xf3, 0xe3, 0x88, 0xaf, 0x20, 0x29, 0x8c, 0x2e, 0x2b, 0xaa, 0x8c,
	0x96, 0xe7, 0xfe, 0xad, 0x41, 0x70, 0xd3, 0x0f, 0xd0, 0xc8, 0x17, 0x7e, 0xfa, 0x51, 0x11, 0x6b,
	0xb8, 0xe8, 0xbd, 0xbd, 0x09, 0xbe, 0x5e, 0x70, 0xd3, 0x43, 0x59, 0x64, 0x20, 0x7a, 0xe9, 0xeb,
	0xe8, 0xc7, 0x92, 0x9b, 0x4f, 0x54, 0xd8, 0x4f, 0x52, 0xd4, 0x36, 0xf2, 0x65, 0xf0, 0x93, 0x29,
	0xcd, 0x60, 0x79, 0xad, 0x6f, 0x2b, 0x8d, 0xdf, 0x49, 0x59, 0x1a, 0x13, 0x67, 0x6d, 0xbb, 0x19,
	0x12, 0x0f, 0xfc, 0x23, 0xe6, 0x7f, 0xe1, 0xa7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba, 0x25,
	0x99, 0xae, 0xa2, 0x03, 0x00, 0x00,
}
