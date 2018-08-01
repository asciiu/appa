// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asciiu/gomo/plan-service/proto/order/order.proto

/*
Package order is a generated protocol buffer package.

It is generated from these files:
	github.com/asciiu/gomo/plan-service/proto/order/order.proto

It has these top-level messages:
	TriggerRequest
	NewOrderRequest
	GetUserOrderRequest
	GetUserOrdersRequest
	RemoveOrderRequest
	Order
	Trigger
	UserOrderData
	UserOrdersData
	OrderResponse
	OrderListResponse
	OrdersPage
*/
package order

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

// Requests
type TriggerRequest struct {
	TriggerID         string   `protobuf:"bytes,1,opt,name=triggerID" json:"triggerID"`
	Code              string   `protobuf:"bytes,2,opt,name=code" json:"code"`
	Index             uint32   `protobuf:"varint,3,opt,name=index" json:"index"`
	Name              string   `protobuf:"bytes,4,opt,name=name" json:"name"`
	Title             string   `protobuf:"bytes,5,opt,name=title" json:"title"`
	TriggerTemplateID string   `protobuf:"bytes,6,opt,name=triggerTemplateID" json:"triggerTemplateID"`
	Actions           []string `protobuf:"bytes,7,rep,name=actions" json:"actions"`
}

func (m *TriggerRequest) Reset()                    { *m = TriggerRequest{} }
func (m *TriggerRequest) String() string            { return proto.CompactTextString(m) }
func (*TriggerRequest) ProtoMessage()               {}
func (*TriggerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TriggerRequest) GetTriggerID() string {
	if m != nil {
		return m.TriggerID
	}
	return ""
}

func (m *TriggerRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *TriggerRequest) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TriggerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TriggerRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TriggerRequest) GetTriggerTemplateID() string {
	if m != nil {
		return m.TriggerTemplateID
	}
	return ""
}

func (m *TriggerRequest) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

type NewOrderRequest struct {
	OrderID               string            `protobuf:"bytes,1,opt,name=orderID" json:"orderID"`
	OrderPriority         uint32            `protobuf:"varint,2,opt,name=orderPriority" json:"orderPriority"`
	OrderType             string            `protobuf:"bytes,3,opt,name=orderType" json:"orderType"`
	OrderTemplateID       string            `protobuf:"bytes,4,opt,name=orderTemplateID" json:"orderTemplateID"`
	KeyID                 string            `protobuf:"bytes,5,opt,name=keyID" json:"keyID"`
	ParentOrderID         string            `protobuf:"bytes,6,opt,name=parentOrderID" json:"parentOrderID"`
	Grupo                 string            `protobuf:"bytes,7,opt,name=grupo" json:"grupo"`
	MarketName            string            `protobuf:"bytes,8,opt,name=marketName" json:"marketName"`
	Side                  string            `protobuf:"bytes,9,opt,name=side" json:"side"`
	LimitPrice            float64           `protobuf:"fixed64,10,opt,name=limitPrice" json:"limitPrice"`
	ActiveCurrencyBalance float64           `protobuf:"fixed64,11,opt,name=activeCurrencyBalance" json:"activeCurrencyBalance"`
	Triggers              []*TriggerRequest `protobuf:"bytes,12,rep,name=triggers" json:"triggers"`
}

func (m *NewOrderRequest) Reset()                    { *m = NewOrderRequest{} }
func (m *NewOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*NewOrderRequest) ProtoMessage()               {}
func (*NewOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewOrderRequest) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *NewOrderRequest) GetOrderPriority() uint32 {
	if m != nil {
		return m.OrderPriority
	}
	return 0
}

func (m *NewOrderRequest) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *NewOrderRequest) GetOrderTemplateID() string {
	if m != nil {
		return m.OrderTemplateID
	}
	return ""
}

func (m *NewOrderRequest) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *NewOrderRequest) GetParentOrderID() string {
	if m != nil {
		return m.ParentOrderID
	}
	return ""
}

func (m *NewOrderRequest) GetGrupo() string {
	if m != nil {
		return m.Grupo
	}
	return ""
}

func (m *NewOrderRequest) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *NewOrderRequest) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *NewOrderRequest) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

func (m *NewOrderRequest) GetActiveCurrencyBalance() float64 {
	if m != nil {
		return m.ActiveCurrencyBalance
	}
	return 0
}

func (m *NewOrderRequest) GetTriggers() []*TriggerRequest {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type GetUserOrderRequest struct {
	OrderID string `protobuf:"bytes,1,opt,name=orderID" json:"orderID"`
	UserID  string `protobuf:"bytes,2,opt,name=userID" json:"userID"`
}

func (m *GetUserOrderRequest) Reset()                    { *m = GetUserOrderRequest{} }
func (m *GetUserOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserOrderRequest) ProtoMessage()               {}
func (*GetUserOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserOrderRequest) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *GetUserOrderRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type GetUserOrdersRequest struct {
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID"`
}

func (m *GetUserOrdersRequest) Reset()                    { *m = GetUserOrdersRequest{} }
func (m *GetUserOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserOrdersRequest) ProtoMessage()               {}
func (*GetUserOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetUserOrdersRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type RemoveOrderRequest struct {
	OrderID string `protobuf:"bytes,1,opt,name=orderID" json:"orderID"`
	UserID  string `protobuf:"bytes,2,opt,name=userID" json:"userID"`
}

func (m *RemoveOrderRequest) Reset()                    { *m = RemoveOrderRequest{} }
func (m *RemoveOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveOrderRequest) ProtoMessage()               {}
func (*RemoveOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveOrderRequest) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *RemoveOrderRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Order struct {
	OrderID               string     `protobuf:"bytes,1,opt,name=orderID" json:"orderID"`
	ParentOrderID         string     `protobuf:"bytes,2,opt,name=parentOrderID" json:"parentOrderID"`
	PlanID                string     `protobuf:"bytes,3,opt,name=planID" json:"planID"`
	PlanDepth             uint32     `protobuf:"varint,4,opt,name=planDepth" json:"planDepth"`
	OrderTemplateID       string     `protobuf:"bytes,5,opt,name=orderTemplateID" json:"orderTemplateID"`
	KeyID                 string     `protobuf:"bytes,6,opt,name=keyID" json:"keyID"`
	KeyPublic             string     `protobuf:"bytes,7,opt,name=keyPublic" json:"keyPublic"`
	KeySecret             string     `protobuf:"bytes,8,opt,name=keySecret" json:"keySecret"`
	KeyDescription        string     `protobuf:"bytes,9,opt,name=keyDescription" json:"keyDescription"`
	OrderPriority         uint32     `protobuf:"varint,10,opt,name=orderPriority" json:"orderPriority"`
	OrderType             string     `protobuf:"bytes,11,opt,name=orderType" json:"orderType"`
	Side                  string     `protobuf:"bytes,12,opt,name=side" json:"side"`
	LimitPrice            float64    `protobuf:"fixed64,13,opt,name=limitPrice" json:"limitPrice"`
	Exchange              string     `protobuf:"bytes,14,opt,name=exchange" json:"exchange"`
	ExchangeMarketName    string     `protobuf:"bytes,15,opt,name=exchangeMarketName" json:"exchangeMarketName"`
	MarketName            string     `protobuf:"bytes,16,opt,name=marketName" json:"marketName"`
	ActiveCurrencySymbol  string     `protobuf:"bytes,17,opt,name=activeCurrencySymbol" json:"activeCurrencySymbol"`
	ActiveCurrencyBalance float64    `protobuf:"fixed64,18,opt,name=activeCurrencyBalance" json:"activeCurrencyBalance"`
	ActiveCurrencyTraded  float64    `protobuf:"fixed64,19,opt,name=activeCurrencyTraded" json:"activeCurrencyTraded"`
	Status                string     `protobuf:"bytes,20,opt,name=status" json:"status"`
	Grupo                 string     `protobuf:"bytes,21,opt,name=grupo" json:"grupo"`
	CreatedOn             string     `protobuf:"bytes,22,opt,name=createdOn" json:"createdOn"`
	UpdatedOn             string     `protobuf:"bytes,23,opt,name=updatedOn" json:"updatedOn"`
	Triggers              []*Trigger `protobuf:"bytes,24,rep,name=triggers" json:"triggers"`
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

func (m *Order) GetParentOrderID() string {
	if m != nil {
		return m.ParentOrderID
	}
	return ""
}

func (m *Order) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *Order) GetPlanDepth() uint32 {
	if m != nil {
		return m.PlanDepth
	}
	return 0
}

func (m *Order) GetOrderTemplateID() string {
	if m != nil {
		return m.OrderTemplateID
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

func (m *Order) GetKeyDescription() string {
	if m != nil {
		return m.KeyDescription
	}
	return ""
}

func (m *Order) GetOrderPriority() uint32 {
	if m != nil {
		return m.OrderPriority
	}
	return 0
}

func (m *Order) GetOrderType() string {
	if m != nil {
		return m.OrderType
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

func (m *Order) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Order) GetExchangeMarketName() string {
	if m != nil {
		return m.ExchangeMarketName
	}
	return ""
}

func (m *Order) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *Order) GetActiveCurrencySymbol() string {
	if m != nil {
		return m.ActiveCurrencySymbol
	}
	return ""
}

func (m *Order) GetActiveCurrencyBalance() float64 {
	if m != nil {
		return m.ActiveCurrencyBalance
	}
	return 0
}

func (m *Order) GetActiveCurrencyTraded() float64 {
	if m != nil {
		return m.ActiveCurrencyTraded
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Order) GetGrupo() string {
	if m != nil {
		return m.Grupo
	}
	return ""
}

func (m *Order) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *Order) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
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
	TriggerID         string   `protobuf:"bytes,1,opt,name=triggerID" json:"triggerID"`
	TriggerTemplateID string   `protobuf:"bytes,2,opt,name=triggerTemplateID" json:"triggerTemplateID"`
	OrderID           string   `protobuf:"bytes,3,opt,name=orderID" json:"orderID"`
	Index             uint32   `protobuf:"varint,4,opt,name=index" json:"index"`
	Title             string   `protobuf:"bytes,5,opt,name=title" json:"title"`
	Name              string   `protobuf:"bytes,6,opt,name=name" json:"name"`
	Code              string   `protobuf:"bytes,7,opt,name=code" json:"code"`
	Triggered         bool     `protobuf:"varint,8,opt,name=triggered" json:"triggered"`
	CreatedOn         string   `protobuf:"bytes,9,opt,name=createdOn" json:"createdOn"`
	UpdatedOn         string   `protobuf:"bytes,10,opt,name=updatedOn" json:"updatedOn"`
	Actions           []string `protobuf:"bytes,11,rep,name=actions" json:"actions"`
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

func (m *Trigger) GetTriggerTemplateID() string {
	if m != nil {
		return m.TriggerTemplateID
	}
	return ""
}

func (m *Trigger) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *Trigger) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Trigger) GetTitle() string {
	if m != nil {
		return m.Title
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

func (m *Trigger) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *Trigger) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *Trigger) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

// Responses
type UserOrderData struct {
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order"`
}

func (m *UserOrderData) Reset()                    { *m = UserOrderData{} }
func (m *UserOrderData) String() string            { return proto.CompactTextString(m) }
func (*UserOrderData) ProtoMessage()               {}
func (*UserOrderData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserOrderData) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type UserOrdersData struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders"`
}

func (m *UserOrdersData) Reset()                    { *m = UserOrdersData{} }
func (m *UserOrdersData) String() string            { return proto.CompactTextString(m) }
func (*UserOrdersData) ProtoMessage()               {}
func (*UserOrdersData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UserOrdersData) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type OrderResponse struct {
	Status  string         `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string         `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *UserOrderData `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *OrderResponse) Reset()                    { *m = OrderResponse{} }
func (m *OrderResponse) String() string            { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()               {}
func (*OrderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *OrderResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrderResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OrderResponse) GetData() *UserOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

type OrderListResponse struct {
	Status  string          `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string          `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *UserOrdersData `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *OrderListResponse) Reset()                    { *m = OrderListResponse{} }
func (m *OrderListResponse) String() string            { return proto.CompactTextString(m) }
func (*OrderListResponse) ProtoMessage()               {}
func (*OrderListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *OrderListResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrderListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OrderListResponse) GetData() *UserOrdersData {
	if m != nil {
		return m.Data
	}
	return nil
}

type OrdersPage struct {
	Page     uint32   `protobuf:"varint,1,opt,name=page" json:"page"`
	PageSize uint32   `protobuf:"varint,2,opt,name=pageSize" json:"pageSize"`
	Total    uint32   `protobuf:"varint,3,opt,name=total" json:"total"`
	Orders   []*Order `protobuf:"bytes,4,rep,name=orders" json:"orders"`
}

func (m *OrdersPage) Reset()                    { *m = OrdersPage{} }
func (m *OrdersPage) String() string            { return proto.CompactTextString(m) }
func (*OrdersPage) ProtoMessage()               {}
func (*OrdersPage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *OrdersPage) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *OrdersPage) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *OrdersPage) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *OrdersPage) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*TriggerRequest)(nil), "order.TriggerRequest")
	proto.RegisterType((*NewOrderRequest)(nil), "order.NewOrderRequest")
	proto.RegisterType((*GetUserOrderRequest)(nil), "order.GetUserOrderRequest")
	proto.RegisterType((*GetUserOrdersRequest)(nil), "order.GetUserOrdersRequest")
	proto.RegisterType((*RemoveOrderRequest)(nil), "order.RemoveOrderRequest")
	proto.RegisterType((*Order)(nil), "order.Order")
	proto.RegisterType((*Trigger)(nil), "order.Trigger")
	proto.RegisterType((*UserOrderData)(nil), "order.UserOrderData")
	proto.RegisterType((*UserOrdersData)(nil), "order.UserOrdersData")
	proto.RegisterType((*OrderResponse)(nil), "order.OrderResponse")
	proto.RegisterType((*OrderListResponse)(nil), "order.OrderListResponse")
	proto.RegisterType((*OrdersPage)(nil), "order.OrdersPage")
}

func init() {
	proto.RegisterFile("github.com/asciiu/gomo/plan-service/proto/order/order.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 889 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0x97, 0xdb, 0xfc, 0x69, 0x26, 0x4d, 0x4a, 0xf7, 0xd2, 0xb2, 0x42, 0x15, 0x8a, 0xa2, 0x13,
	0x0a, 0x08, 0x12, 0xd1, 0x43, 0x7c, 0xe1, 0x1b, 0x44, 0x9c, 0x2a, 0xc1, 0xb5, 0x72, 0xcb, 0x03,
	0x6c, 0xed, 0x51, 0xba, 0x4a, 0x62, 0x9b, 0xdd, 0x75, 0x69, 0x78, 0x22, 0xde, 0x00, 0xf1, 0x0a,
	0x3c, 0x15, 0xda, 0xd9, 0x75, 0x6c, 0x27, 0x69, 0x8b, 0x74, 0x5f, 0x92, 0x9d, 0xdf, 0xcc, 0x78,
	0x3c, 0xb3, 0xbf, 0x99, 0x31, 0xfc, 0x30, 0x97, 0xe6, 0x21, 0xbf, 0x9f, 0x44, 0xe9, 0x6a, 0x2a,
	0x74, 0x24, 0x65, 0x3e, 0x9d, 0xa7, 0xab, 0x74, 0x9a, 0x2d, 0x45, 0xf2, 0x8d, 0x46, 0xf5, 0x28,
	0x23, 0x9c, 0x66, 0x2a, 0x35, 0xe9, 0x34, 0x55, 0x31, 0x2a, 0xf7, 0x3b, 0x21, 0x84, 0x35, 0x49,
	0x18, 0xfd, 0x1b, 0x40, 0xff, 0x4e, 0xc9, 0xf9, 0x1c, 0x55, 0x88, 0xbf, 0xe7, 0xa8, 0x0d, 0xbb,
	0x80, 0x8e, 0x71, 0xc8, 0xd5, 0x8c, 0x07, 0xc3, 0x60, 0xdc, 0x09, 0x4b, 0x80, 0x31, 0x68, 0x44,
	0x69, 0x8c, 0xfc, 0x80, 0x14, 0x74, 0x66, 0x03, 0x68, 0xca, 0x24, 0xc6, 0x27, 0x7e, 0x38, 0x0c,
	0xc6, 0xbd, 0xd0, 0x09, 0xd6, 0x32, 0x11, 0x2b, 0xe4, 0x0d, 0x67, 0x69, 0xcf, 0xd6, 0xd2, 0x48,
	0xb3, 0x44, 0xde, 0x24, 0xd0, 0x09, 0xec, 0x6b, 0x38, 0xf5, 0x01, 0xee, 0x70, 0x95, 0x2d, 0x85,
	0xc1, 0xab, 0x19, 0x6f, 0x91, 0xc5, 0xae, 0x82, 0x71, 0x68, 0x8b, 0xc8, 0xc8, 0x34, 0xd1, 0xbc,
	0x3d, 0x3c, 0x1c, 0x77, 0xc2, 0x42, 0x1c, 0xfd, 0x75, 0x08, 0x27, 0x1f, 0xf0, 0x8f, 0x6b, 0x9b,
	0x59, 0x91, 0x0d, 0x87, 0x36, 0x65, 0xba, 0xc9, 0xa5, 0x10, 0xd9, 0x5b, 0xe8, 0xd1, 0xf1, 0x46,
	0xc9, 0x54, 0x49, 0xb3, 0xa6, 0x94, 0x7a, 0x61, 0x1d, 0xb4, 0xd5, 0x20, 0xe0, 0x6e, 0x9d, 0x21,
	0xe5, 0xd7, 0x09, 0x4b, 0x80, 0x8d, 0xe1, 0xc4, 0x09, 0xe5, 0x7b, 0xbb, 0x74, 0xb7, 0x61, 0x9b,
	0xf9, 0x02, 0xd7, 0x57, 0xb3, 0x22, 0x73, 0x12, 0xec, 0x3b, 0x64, 0x42, 0x61, 0x62, 0xae, 0xfd,
	0x3b, 0xba, 0xac, 0xeb, 0xa0, 0xf5, 0x9d, 0xab, 0x3c, 0x4b, 0x79, 0xdb, 0xf9, 0x92, 0xc0, 0x3e,
	0x07, 0x58, 0x09, 0xb5, 0x40, 0xf3, 0xc1, 0x56, 0xf9, 0x88, 0x54, 0x15, 0xc4, 0xd6, 0x5f, 0xcb,
	0x18, 0x79, 0xc7, 0xd5, 0xdf, 0x9e, 0xad, 0xcf, 0x52, 0xae, 0xa4, 0xb9, 0x51, 0x32, 0x42, 0x0e,
	0xc3, 0x60, 0x1c, 0x84, 0x15, 0x84, 0x7d, 0x07, 0x67, 0xb6, 0x98, 0x8f, 0xf8, 0x53, 0xae, 0x14,
	0x26, 0xd1, 0xfa, 0x47, 0xb1, 0x14, 0x49, 0x84, 0xbc, 0x4b, 0xa6, 0xfb, 0x95, 0xec, 0x5b, 0x38,
	0xf2, 0xd7, 0xa4, 0xf9, 0xf1, 0xf0, 0x70, 0xdc, 0xbd, 0x3c, 0x9b, 0x38, 0xae, 0xd5, 0xa9, 0x15,
	0x6e, 0xcc, 0x46, 0xef, 0xe1, 0xcd, 0x7b, 0x34, 0xbf, 0x69, 0x54, 0xff, 0xf3, 0xb6, 0xce, 0xa1,
	0x95, 0x6b, 0x52, 0x38, 0xe6, 0x79, 0x69, 0x34, 0x81, 0x41, 0xf5, 0x41, 0xba, 0x78, 0x52, 0x69,
	0x1f, 0xd4, 0xec, 0x7f, 0x06, 0x16, 0xe2, 0x2a, 0x7d, 0xc4, 0x8f, 0x8c, 0xfb, 0x4f, 0x0b, 0x9a,
	0xf4, 0x88, 0x97, 0x19, 0x56, 0xbf, 0xdd, 0x83, 0x7d, 0xb7, 0x7b, 0x0e, 0x2d, 0xdb, 0xb1, 0x57,
	0x33, 0x4f, 0x2f, 0x2f, 0x59, 0xe6, 0xd9, 0xd3, 0x0c, 0x33, 0xf3, 0x40, 0xac, 0xea, 0x85, 0x25,
	0xb0, 0x8f, 0x79, 0xcd, 0x57, 0x98, 0xd7, 0xaa, 0x32, 0xef, 0x02, 0x3a, 0x0b, 0x5c, 0xdf, 0xe4,
	0xf7, 0x4b, 0x19, 0x79, 0x5e, 0x95, 0x80, 0xd7, 0xde, 0x62, 0xa4, 0xd0, 0x78, 0x6a, 0x95, 0x00,
	0xfb, 0x02, 0xfa, 0x0b, 0x5c, 0xcf, 0x50, 0x47, 0x4a, 0x66, 0xb6, 0xf5, 0x3c, 0xc7, 0xb6, 0xd0,
	0xdd, 0x0e, 0x83, 0x57, 0x3b, 0xac, 0xbb, 0xdd, 0x61, 0x05, 0x8b, 0x8f, 0x9f, 0x65, 0x71, 0x6f,
	0x87, 0xc5, 0x9f, 0xc1, 0x11, 0x3e, 0x45, 0x0f, 0x22, 0x99, 0x23, 0xef, 0x93, 0xdf, 0x46, 0x66,
	0x13, 0x60, 0xc5, 0xf9, 0xd7, 0xb2, 0x7b, 0x4e, 0xc8, 0x6a, 0x8f, 0x66, 0xab, 0xcb, 0x3e, 0xd9,
	0xe9, 0xb2, 0x4b, 0x18, 0xd4, 0x9b, 0xe2, 0x76, 0xbd, 0xba, 0x4f, 0x97, 0xfc, 0x94, 0x2c, 0xf7,
	0xea, 0x9e, 0xef, 0x32, 0xf6, 0x52, 0x97, 0xed, 0x44, 0xba, 0x53, 0x22, 0xc6, 0x98, 0xbf, 0x21,
	0xa7, 0xbd, 0x3a, 0xcb, 0x2d, 0x6d, 0x84, 0xc9, 0x35, 0x1f, 0x38, 0x6e, 0x39, 0xa9, 0x9c, 0x28,
	0x67, 0xd5, 0x89, 0x72, 0x01, 0x9d, 0x48, 0xa1, 0x30, 0x18, 0x5f, 0x27, 0xfc, 0xdc, 0xdd, 0xc4,
	0x06, 0xb0, 0xda, 0x3c, 0x8b, 0xbd, 0xf6, 0x53, 0xa7, 0xdd, 0x00, 0xec, 0xab, 0xca, 0x0c, 0xe0,
	0x34, 0x03, 0xfa, 0x5b, 0x33, 0xa0, 0x6c, 0xfe, 0xbf, 0x0f, 0xa0, 0xed, 0xd1, 0x57, 0xb6, 0xcd,
	0xde, 0xcd, 0x70, 0xf0, 0xc2, 0x66, 0x28, 0x3a, 0xf1, 0xb0, 0xde, 0x89, 0x9b, 0x0d, 0xd5, 0xa8,
	0x6e, 0xa8, 0xfd, 0xdb, 0xa8, 0xd8, 0x5b, 0xad, 0xca, 0xde, 0x2a, 0xb6, 0x5e, 0xbb, 0xb2, 0xf5,
	0xca, 0x37, 0xc7, 0x98, 0x7a, 0xe4, 0x28, 0x2c, 0x81, 0x7a, 0x2d, 0x3b, 0x2f, 0xd6, 0x12, 0xb6,
	0x6b, 0x59, 0xd9, 0x70, 0xdd, 0xfa, 0x86, 0x7b, 0x07, 0xbd, 0xcd, 0xa8, 0x9b, 0x09, 0x23, 0xd8,
	0x08, 0xdc, 0x22, 0xa7, 0xd2, 0x75, 0x2f, 0x8f, 0x7d, 0xcd, 0xdd, 0x70, 0xf3, 0x3b, 0xfe, 0x7b,
	0xe8, 0x97, 0xf3, 0x91, 0xbc, 0xde, 0x42, 0x8b, 0x54, 0x9a, 0x07, 0x74, 0x55, 0x75, 0x37, 0xaf,
	0x1b, 0x2d, 0xa0, 0xe7, 0x87, 0xa4, 0xce, 0xd2, 0x44, 0x63, 0x85, 0x4d, 0x41, 0x8d, 0x4d, 0x1c,
	0xda, 0x2b, 0xd4, 0x5a, 0xcc, 0x8b, 0xcf, 0x82, 0x42, 0x64, 0x63, 0x68, 0xc4, 0xc2, 0x08, 0xba,
	0x8e, 0xee, 0xe5, 0xc0, 0x87, 0xa9, 0xa5, 0x10, 0x92, 0xc5, 0x28, 0x83, 0x53, 0x82, 0x7e, 0x91,
	0xda, 0x7c, 0x44, 0xc0, 0x2f, 0x6b, 0x01, 0xcf, 0xb6, 0x03, 0xea, 0x4a, 0xc4, 0x27, 0x00, 0x87,
	0xdd, 0x58, 0x47, 0x06, 0x8d, 0xcc, 0x3e, 0x2f, 0x20, 0x82, 0xd0, 0xd9, 0xce, 0x11, 0xfb, 0x7f,
	0x2b, 0xff, 0x44, 0xff, 0x71, 0xb0, 0x91, 0x89, 0x3b, 0xa9, 0x11, 0xcb, 0xe2, 0x9b, 0x87, 0x84,
	0x4a, 0x61, 0x1b, 0xcf, 0x17, 0xf6, 0xbe, 0x45, 0x9f, 0x60, 0xef, 0xfe, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xd5, 0xd4, 0x4c, 0x79, 0xc1, 0x09, 0x00, 0x00,
}
