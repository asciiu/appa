// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asciiu/gomo/plan-service/proto/order/order.proto

/*
Package order is a generated protocol buffer package.

It is generated from these files:
	github.com/asciiu/gomo/plan-service/proto/order/order.proto

It has these top-level messages:
	TriggerRequest
	UpdateOrderRequest
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
	TriggerID         string   `protobuf:"bytes,1,opt,name=triggerID" json:"triggerID,omitempty"`
	Name              string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Code              string   `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	TriggerTemplateID string   `protobuf:"bytes,4,opt,name=triggerTemplateID" json:"triggerTemplateID,omitempty"`
	Actions           []string `protobuf:"bytes,5,rep,name=actions" json:"actions,omitempty"`
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

func (m *TriggerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TriggerRequest) GetCode() string {
	if m != nil {
		return m.Code
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

type UpdateOrderRequest struct {
	OrderID         string            `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	OrderType       string            `protobuf:"bytes,2,opt,name=orderType" json:"orderType,omitempty"`
	OrderTemplateID string            `protobuf:"bytes,3,opt,name=orderTemplateID" json:"orderTemplateID,omitempty"`
	KeyID           string            `protobuf:"bytes,4,opt,name=keyID" json:"keyID,omitempty"`
	ParentOrderID   string            `protobuf:"bytes,5,opt,name=parentOrderID" json:"parentOrderID,omitempty"`
	MarketName      string            `protobuf:"bytes,6,opt,name=marketName" json:"marketName,omitempty"`
	Side            string            `protobuf:"bytes,7,opt,name=side" json:"side,omitempty"`
	LimitPrice      float64           `protobuf:"fixed64,8,opt,name=limitPrice" json:"limitPrice,omitempty"`
	CurrencyBalance float64           `protobuf:"fixed64,9,opt,name=currencyBalance" json:"currencyBalance,omitempty"`
	Action          string            `protobuf:"bytes,10,opt,name=action" json:"action,omitempty"`
	Triggers        []*TriggerRequest `protobuf:"bytes,11,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *UpdateOrderRequest) Reset()                    { *m = UpdateOrderRequest{} }
func (m *UpdateOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateOrderRequest) ProtoMessage()               {}
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateOrderRequest) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *UpdateOrderRequest) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *UpdateOrderRequest) GetOrderTemplateID() string {
	if m != nil {
		return m.OrderTemplateID
	}
	return ""
}

func (m *UpdateOrderRequest) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *UpdateOrderRequest) GetParentOrderID() string {
	if m != nil {
		return m.ParentOrderID
	}
	return ""
}

func (m *UpdateOrderRequest) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *UpdateOrderRequest) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *UpdateOrderRequest) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

func (m *UpdateOrderRequest) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *UpdateOrderRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *UpdateOrderRequest) GetTriggers() []*TriggerRequest {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type NewOrderRequest struct {
	OrderID         string            `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	OrderType       string            `protobuf:"bytes,2,opt,name=orderType" json:"orderType,omitempty"`
	OrderTemplateID string            `protobuf:"bytes,3,opt,name=orderTemplateID" json:"orderTemplateID,omitempty"`
	KeyID           string            `protobuf:"bytes,4,opt,name=keyID" json:"keyID,omitempty"`
	ParentOrderID   string            `protobuf:"bytes,5,opt,name=parentOrderID" json:"parentOrderID,omitempty"`
	MarketName      string            `protobuf:"bytes,6,opt,name=marketName" json:"marketName,omitempty"`
	Side            string            `protobuf:"bytes,7,opt,name=side" json:"side,omitempty"`
	LimitPrice      float64           `protobuf:"fixed64,8,opt,name=limitPrice" json:"limitPrice,omitempty"`
	CurrencyBalance float64           `protobuf:"fixed64,9,opt,name=currencyBalance" json:"currencyBalance,omitempty"`
	Triggers        []*TriggerRequest `protobuf:"bytes,10,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *NewOrderRequest) Reset()                    { *m = NewOrderRequest{} }
func (m *NewOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*NewOrderRequest) ProtoMessage()               {}
func (*NewOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NewOrderRequest) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
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

func (m *NewOrderRequest) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
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
	OrderID string `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	UserID  string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *GetUserOrderRequest) Reset()                    { *m = GetUserOrderRequest{} }
func (m *GetUserOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserOrderRequest) ProtoMessage()               {}
func (*GetUserOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
}

func (m *GetUserOrdersRequest) Reset()                    { *m = GetUserOrdersRequest{} }
func (m *GetUserOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserOrdersRequest) ProtoMessage()               {}
func (*GetUserOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetUserOrdersRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type RemoveOrderRequest struct {
	OrderID string `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	UserID  string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *RemoveOrderRequest) Reset()                    { *m = RemoveOrderRequest{} }
func (m *RemoveOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveOrderRequest) ProtoMessage()               {}
func (*RemoveOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
	OrderID            string     `protobuf:"bytes,1,opt,name=orderID" json:"orderID,omitempty"`
	ParentOrderID      string     `protobuf:"bytes,2,opt,name=parentOrderID" json:"parentOrderID,omitempty"`
	PlanID             string     `protobuf:"bytes,3,opt,name=planID" json:"planID,omitempty"`
	PlanDepth          uint32     `protobuf:"varint,4,opt,name=planDepth" json:"planDepth,omitempty"`
	OrderTemplateID    string     `protobuf:"bytes,5,opt,name=orderTemplateID" json:"orderTemplateID,omitempty"`
	KeyID              string     `protobuf:"bytes,6,opt,name=keyID" json:"keyID,omitempty"`
	KeyPublic          string     `protobuf:"bytes,7,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret          string     `protobuf:"bytes,8,opt,name=keySecret" json:"keySecret,omitempty"`
	KeyDescription     string     `protobuf:"bytes,9,opt,name=keyDescription" json:"keyDescription,omitempty"`
	OrderType          string     `protobuf:"bytes,10,opt,name=orderType" json:"orderType,omitempty"`
	Side               string     `protobuf:"bytes,11,opt,name=side" json:"side,omitempty"`
	LimitPrice         float64    `protobuf:"fixed64,12,opt,name=limitPrice" json:"limitPrice,omitempty"`
	Exchange           string     `protobuf:"bytes,13,opt,name=exchange" json:"exchange,omitempty"`
	ExchangeMarketName string     `protobuf:"bytes,14,opt,name=exchangeMarketName" json:"exchangeMarketName,omitempty"`
	MarketName         string     `protobuf:"bytes,15,opt,name=marketName" json:"marketName,omitempty"`
	CurrencySymbol     string     `protobuf:"bytes,16,opt,name=currencySymbol" json:"currencySymbol,omitempty"`
	CurrencyBalance    float64    `protobuf:"fixed64,17,opt,name=currencyBalance" json:"currencyBalance,omitempty"`
	CurrencyTraded     float64    `protobuf:"fixed64,18,opt,name=currencyTraded" json:"currencyTraded,omitempty"`
	Status             string     `protobuf:"bytes,19,opt,name=status" json:"status,omitempty"`
	CreatedOn          string     `protobuf:"bytes,20,opt,name=createdOn" json:"createdOn,omitempty"`
	UpdatedOn          string     `protobuf:"bytes,21,opt,name=updatedOn" json:"updatedOn,omitempty"`
	Triggers           []*Trigger `protobuf:"bytes,22,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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

func (m *Order) GetCurrencySymbol() string {
	if m != nil {
		return m.CurrencySymbol
	}
	return ""
}

func (m *Order) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *Order) GetCurrencyTraded() float64 {
	if m != nil {
		return m.CurrencyTraded
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
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
	TriggerID         string   `protobuf:"bytes,1,opt,name=triggerID" json:"triggerID,omitempty"`
	OrderID           string   `protobuf:"bytes,2,opt,name=orderID" json:"orderID,omitempty"`
	TriggerNumber     uint32   `protobuf:"varint,3,opt,name=triggerNumber" json:"triggerNumber,omitempty"`
	TriggerTemplateID string   `protobuf:"bytes,4,opt,name=triggerTemplateID" json:"triggerTemplateID,omitempty"`
	Name              string   `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Code              string   `protobuf:"bytes,6,opt,name=code" json:"code,omitempty"`
	Triggered         bool     `protobuf:"varint,7,opt,name=triggered" json:"triggered,omitempty"`
	CreatedOn         string   `protobuf:"bytes,8,opt,name=createdOn" json:"createdOn,omitempty"`
	UpdatedOn         string   `protobuf:"bytes,9,opt,name=updatedOn" json:"updatedOn,omitempty"`
	Actions           []string `protobuf:"bytes,10,rep,name=actions" json:"actions,omitempty"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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

func (m *Trigger) GetTriggerNumber() uint32 {
	if m != nil {
		return m.TriggerNumber
	}
	return 0
}

func (m *Trigger) GetTriggerTemplateID() string {
	if m != nil {
		return m.TriggerTemplateID
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
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
}

func (m *UserOrderData) Reset()                    { *m = UserOrderData{} }
func (m *UserOrderData) String() string            { return proto.CompactTextString(m) }
func (*UserOrderData) ProtoMessage()               {}
func (*UserOrderData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UserOrderData) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type UserOrdersData struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *UserOrdersData) Reset()                    { *m = UserOrdersData{} }
func (m *UserOrdersData) String() string            { return proto.CompactTextString(m) }
func (*UserOrdersData) ProtoMessage()               {}
func (*UserOrdersData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UserOrdersData) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type OrderResponse struct {
	Status  string         `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Data    *UserOrderData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *OrderResponse) Reset()                    { *m = OrderResponse{} }
func (m *OrderResponse) String() string            { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()               {}
func (*OrderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
	Status  string          `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Data    *UserOrdersData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *OrderListResponse) Reset()                    { *m = OrderListResponse{} }
func (m *OrderListResponse) String() string            { return proto.CompactTextString(m) }
func (*OrderListResponse) ProtoMessage()               {}
func (*OrderListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
	Page     uint32   `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	PageSize uint32   `protobuf:"varint,2,opt,name=pageSize" json:"pageSize,omitempty"`
	Total    uint32   `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	Orders   []*Order `protobuf:"bytes,4,rep,name=orders" json:"orders,omitempty"`
}

func (m *OrdersPage) Reset()                    { *m = OrdersPage{} }
func (m *OrdersPage) String() string            { return proto.CompactTextString(m) }
func (*OrdersPage) ProtoMessage()               {}
func (*OrdersPage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

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
	proto.RegisterType((*UpdateOrderRequest)(nil), "order.UpdateOrderRequest")
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
	// 851 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x65, 0x89, 0x92, 0x46, 0x91, 0x5c, 0x6f, 0x6c, 0x63, 0x51, 0x04, 0x85, 0x21, 0x04,
	0x81, 0x5a, 0xb4, 0x12, 0xea, 0x00, 0xbd, 0xf4, 0x56, 0x08, 0x0d, 0x02, 0xb4, 0xb6, 0x41, 0x2b,
	0x0f, 0xb0, 0x22, 0x07, 0x32, 0x21, 0xf1, 0xa7, 0xbb, 0xcb, 0x34, 0xea, 0xc3, 0xf4, 0x1d, 0xfa,
	0x38, 0xbd, 0xf5, 0x3d, 0x7a, 0x29, 0x76, 0x96, 0xff, 0x52, 0x12, 0x17, 0xb9, 0xe6, 0x22, 0xed,
	0x7c, 0x33, 0xbb, 0x33, 0xfc, 0xf8, 0xcd, 0x0e, 0xe1, 0xc7, 0x4d, 0xa8, 0x1f, 0xb2, 0xf5, 0xdc,
	0x4f, 0xa2, 0x85, 0x50, 0x7e, 0x18, 0x66, 0x8b, 0x4d, 0x12, 0x25, 0x8b, 0x74, 0x27, 0xe2, 0xef,
	0x14, 0xca, 0xb7, 0xa1, 0x8f, 0x8b, 0x54, 0x26, 0x3a, 0x59, 0x24, 0x32, 0x40, 0x69, 0x7f, 0xe7,
	0x84, 0xb0, 0x1e, 0x19, 0xd3, 0x3f, 0x1d, 0x98, 0xac, 0x64, 0xb8, 0xd9, 0xa0, 0xf4, 0xf0, 0xb7,
	0x0c, 0x95, 0x66, 0xcf, 0x60, 0xa8, 0x2d, 0xf2, 0x7a, 0xc9, 0x9d, 0x2b, 0x67, 0x36, 0xf4, 0x2a,
	0x80, 0x31, 0xe8, 0xc6, 0x22, 0x42, 0xde, 0x21, 0x07, 0xad, 0x0d, 0xe6, 0x27, 0x01, 0xf2, 0x13,
	0x8b, 0x99, 0x35, 0xfb, 0x16, 0xce, 0xf2, 0x4d, 0x2b, 0x8c, 0xd2, 0x9d, 0xd0, 0xf8, 0x7a, 0xc9,
	0xbb, 0x14, 0x70, 0xe8, 0x60, 0x1c, 0xfa, 0xc2, 0xd7, 0x61, 0x12, 0x2b, 0xde, 0xbb, 0x3a, 0x99,
	0x0d, 0xbd, 0xc2, 0x9c, 0xfe, 0xdb, 0x01, 0xf6, 0x26, 0x0d, 0x84, 0xc6, 0x5b, 0x53, 0x70, 0x51,
	0x24, 0x87, 0x3e, 0x3d, 0x40, 0x59, 0x62, 0x61, 0x9a, 0xf2, 0x69, 0xb9, 0xda, 0xa7, 0x45, 0x95,
	0x15, 0xc0, 0x66, 0x70, 0x6a, 0x8d, 0xaa, 0x28, 0x5b, 0x75, 0x1b, 0x66, 0xe7, 0xd0, 0xdb, 0xe2,
	0xbe, 0x2c, 0xda, 0x1a, 0xec, 0x39, 0x8c, 0x53, 0x21, 0x31, 0xd6, 0xb7, 0x79, 0xf6, 0x1e, 0x79,
	0x9b, 0x20, 0xfb, 0x0a, 0x20, 0x12, 0x72, 0x8b, 0xfa, 0xc6, 0x50, 0xe5, 0x52, 0x48, 0x0d, 0x31,
	0x84, 0xa9, 0x30, 0x40, 0xde, 0xb7, 0x84, 0x99, 0xb5, 0xd9, 0xb3, 0x0b, 0xa3, 0x50, 0xdf, 0xc9,
	0xd0, 0x47, 0x3e, 0xb8, 0x72, 0x66, 0x8e, 0x57, 0x43, 0x4c, 0xe5, 0x7e, 0x26, 0x25, 0xc6, 0xfe,
	0xfe, 0x27, 0xb1, 0x13, 0xb1, 0x8f, 0x7c, 0x48, 0x41, 0x6d, 0x98, 0x5d, 0x82, 0x6b, 0xd9, 0xe3,
	0x40, 0xe7, 0xe7, 0x16, 0xfb, 0x1e, 0x06, 0x39, 0xf3, 0x8a, 0x8f, 0xae, 0x4e, 0x66, 0xa3, 0xeb,
	0x8b, 0xb9, 0x95, 0x44, 0x53, 0x01, 0x5e, 0x19, 0x36, 0xfd, 0xa7, 0x03, 0xa7, 0x37, 0xf8, 0xfb,
	0x67, 0xea, 0x1f, 0x49, 0x7d, 0x9d, 0x62, 0x78, 0x1c, 0xc5, 0xaf, 0xe0, 0xe9, 0x2b, 0xd4, 0x6f,
	0x14, 0xca, 0x47, 0xb2, 0x7c, 0x09, 0x6e, 0xa6, 0xc8, 0x61, 0x29, 0xce, 0xad, 0xe9, 0x1c, 0xce,
	0xeb, 0x07, 0xa9, 0xe2, 0xa4, 0x2a, 0xde, 0x69, 0xc4, 0xff, 0x0c, 0xcc, 0xc3, 0x28, 0x79, 0x8b,
	0x9f, 0x98, 0xf7, 0xef, 0x1e, 0xf4, 0xe8, 0x88, 0x0f, 0xec, 0x3d, 0x78, 0x77, 0x9d, 0x63, 0xef,
	0xee, 0x12, 0x5c, 0x73, 0x77, 0x95, 0xc2, 0xc8, 0x2d, 0xa3, 0x2b, 0xb3, 0x5a, 0x62, 0xaa, 0x1f,
	0x48, 0x13, 0x63, 0xaf, 0x02, 0x8e, 0xe9, 0xaa, 0xf7, 0x11, 0x5d, 0xb9, 0x75, 0x5d, 0x3d, 0x83,
	0xe1, 0x16, 0xf7, 0x77, 0xd9, 0x7a, 0x17, 0xfa, 0xb9, 0x2c, 0x2a, 0x20, 0xf7, 0xde, 0xa3, 0x2f,
	0x51, 0x93, 0x34, 0xac, 0xd7, 0x02, 0xec, 0x05, 0x4c, 0xb6, 0xb8, 0x5f, 0xa2, 0xf2, 0x65, 0x98,
	0x52, 0xcb, 0x0d, 0x29, 0xa4, 0x85, 0x36, 0x3b, 0x03, 0xda, 0x9d, 0x51, 0x68, 0x72, 0xf4, 0x5e,
	0x4d, 0x3e, 0x39, 0xd0, 0xe4, 0x97, 0x30, 0xc0, 0x77, 0xfe, 0x83, 0x88, 0x37, 0xc8, 0xc7, 0xb4,
	0xaf, 0xb4, 0xd9, 0x1c, 0x58, 0xb1, 0xfe, 0xb5, 0xea, 0x85, 0x09, 0x45, 0x1d, 0xf1, 0xb4, 0x7a,
	0xe6, 0xf4, 0xa0, 0x67, 0x5e, 0xc0, 0xa4, 0x10, 0xfa, 0xfd, 0x3e, 0x5a, 0x27, 0x3b, 0xfe, 0x85,
	0x7d, 0xca, 0x26, 0x7a, 0xac, 0x4f, 0xce, 0x8e, 0xf7, 0x49, 0xed, 0xc4, 0x95, 0x14, 0x01, 0x06,
	0x9c, 0x51, 0x60, 0x0b, 0x35, 0x8a, 0x50, 0x5a, 0xe8, 0x4c, 0xf1, 0xa7, 0x56, 0x11, 0xd6, 0x32,
	0x7c, 0xfa, 0x12, 0x85, 0xc6, 0xe0, 0x36, 0xe6, 0xe7, 0x96, 0xcf, 0x12, 0x30, 0xde, 0x8c, 0x46,
	0x86, 0xf1, 0x5e, 0x58, 0x6f, 0x09, 0xb0, 0x6f, 0x6a, 0x3d, 0x7a, 0x49, 0x3d, 0x3a, 0x69, 0xf5,
	0x68, 0xd5, 0x9c, 0x7f, 0x75, 0xa0, 0x9f, 0xa3, 0x1f, 0x99, 0x8b, 0x35, 0xed, 0x77, 0x0e, 0xb4,
	0x9f, 0x87, 0xdd, 0x64, 0xd1, 0x1a, 0x25, 0x89, 0x7b, 0xec, 0x35, 0xc1, 0xff, 0x39, 0x2f, 0x8b,
	0x29, 0xdc, 0x3b, 0x32, 0x85, 0xdd, 0xda, 0x14, 0xae, 0x6a, 0xc6, 0x80, 0xb4, 0x3d, 0xf0, 0x2a,
	0xa0, 0xc9, 0xe2, 0xe0, 0x83, 0x2c, 0x0e, 0xdb, 0x2c, 0xd6, 0x26, 0x36, 0x34, 0x27, 0xf6, 0x4b,
	0x18, 0x97, 0x97, 0xd0, 0x52, 0x68, 0xc1, 0xa6, 0x60, 0x3f, 0x36, 0x88, 0xb4, 0xd1, 0xf5, 0x93,
	0x9c, 0x6d, 0x7b, 0xed, 0xe4, 0xdf, 0x21, 0x3f, 0xc0, 0xa4, 0xba, 0xb9, 0x68, 0xd7, 0x73, 0x70,
	0xc9, 0xa5, 0xb8, 0x43, 0x2f, 0xa9, 0xb9, 0x2d, 0xf7, 0x4d, 0xb7, 0x30, 0xce, 0xaf, 0x2f, 0x95,
	0x26, 0xb1, 0xc2, 0x9a, 0x62, 0x9c, 0x86, 0x62, 0x38, 0xf4, 0x23, 0x54, 0x4a, 0x6c, 0x8a, 0xc9,
	0x54, 0x98, 0x6c, 0x06, 0xdd, 0x40, 0x68, 0x41, 0xaf, 0x65, 0x74, 0x7d, 0x9e, 0xa7, 0x69, 0x3c,
	0x82, 0x47, 0x11, 0xd3, 0x14, 0xce, 0x08, 0xfa, 0x25, 0x54, 0xfa, 0x13, 0x12, 0x7e, 0xdd, 0x48,
	0x78, 0xd1, 0x4e, 0xa8, 0x6a, 0x19, 0xdf, 0x01, 0x58, 0xec, 0xce, 0x6c, 0x64, 0xd0, 0x4d, 0xcd,
	0x79, 0x0e, 0x09, 0x88, 0xd6, 0xe6, 0x1e, 0x30, 0xff, 0xf7, 0xe1, 0x1f, 0x36, 0xcf, 0xd8, 0x2b,
	0x6d, 0x73, 0xdf, 0xe9, 0x44, 0x8b, 0x5d, 0xae, 0x38, 0x6b, 0xd4, 0x88, 0xed, 0xbe, 0x9f, 0xd8,
	0xb5, 0x4b, 0x9f, 0x89, 0x2f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd6, 0x51, 0x1c, 0x65,
	0x0a, 0x00, 0x00,
}