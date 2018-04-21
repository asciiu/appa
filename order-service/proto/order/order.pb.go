// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order/order.proto

/*
Package go_micro_srv_order is a generated protocol buffer package.

It is generated from these files:
	proto/order/order.proto

It has these top-level messages:
	OrdersRequest
	OrderRequest
	GetUserOrderRequest
	GetUserOrdersRequest
	RemoveOrderRequest
	Order
	UserOrderData
	UserOrdersData
	OrderResponse
	OrderListResponse
*/
package go_micro_srv_order

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

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
type OrdersRequest struct {
	Orders []*OrderRequest `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *OrdersRequest) Reset()                    { *m = OrdersRequest{} }
func (m *OrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*OrdersRequest) ProtoMessage()               {}
func (*OrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OrdersRequest) GetOrders() []*OrderRequest {
	if m != nil {
		return m.Orders
	}
	return nil
}

type OrderRequest struct {
	OrderId          string  `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
	UserId           string  `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	ApiKeyId         string  `protobuf:"bytes,3,opt,name=apiKeyId" json:"apiKeyId,omitempty"`
	Exchange         string  `protobuf:"bytes,4,opt,name=exchange" json:"exchange,omitempty"`
	MarketName       string  `protobuf:"bytes,5,opt,name=marketName" json:"marketName,omitempty"`
	Side             string  `protobuf:"bytes,6,opt,name=side" json:"side,omitempty"`
	OrderType        string  `protobuf:"bytes,7,opt,name=orderType" json:"orderType,omitempty"`
	BaseQuantity     float64 `protobuf:"fixed64,8,opt,name=baseQuantity" json:"baseQuantity,omitempty"`
	CurrencyQuantity float64 `protobuf:"fixed64,9,opt,name=currencyQuantity" json:"currencyQuantity,omitempty"`
	Conditions       string  `protobuf:"bytes,10,opt,name=conditions" json:"conditions,omitempty"`
	ParentOrderId    string  `protobuf:"bytes,11,opt,name=parentOrderId" json:"parentOrderId,omitempty"`
}

func (m *OrderRequest) Reset()                    { *m = OrderRequest{} }
func (m *OrderRequest) String() string            { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()               {}
func (*OrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OrderRequest) GetApiKeyId() string {
	if m != nil {
		return m.ApiKeyId
	}
	return ""
}

func (m *OrderRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *OrderRequest) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *OrderRequest) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *OrderRequest) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *OrderRequest) GetBaseQuantity() float64 {
	if m != nil {
		return m.BaseQuantity
	}
	return 0
}

func (m *OrderRequest) GetCurrencyQuantity() float64 {
	if m != nil {
		return m.CurrencyQuantity
	}
	return 0
}

func (m *OrderRequest) GetConditions() string {
	if m != nil {
		return m.Conditions
	}
	return ""
}

func (m *OrderRequest) GetParentOrderId() string {
	if m != nil {
		return m.ParentOrderId
	}
	return ""
}

type GetUserOrderRequest struct {
	OrderId string `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
}

func (m *GetUserOrderRequest) Reset()                    { *m = GetUserOrderRequest{} }
func (m *GetUserOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserOrderRequest) ProtoMessage()               {}
func (*GetUserOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserOrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *GetUserOrderRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetUserOrdersRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *GetUserOrdersRequest) Reset()                    { *m = GetUserOrdersRequest{} }
func (m *GetUserOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserOrdersRequest) ProtoMessage()               {}
func (*GetUserOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetUserOrdersRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type RemoveOrderRequest struct {
	OrderId string `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
}

func (m *RemoveOrderRequest) Reset()                    { *m = RemoveOrderRequest{} }
func (m *RemoveOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveOrderRequest) ProtoMessage()               {}
func (*RemoveOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveOrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *RemoveOrderRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// Responses
type Order struct {
	OrderId               string  `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
	UserId                string  `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	ApiKeyId              string  `protobuf:"bytes,3,opt,name=apiKeyId" json:"apiKeyId,omitempty"`
	Exchange              string  `protobuf:"bytes,4,opt,name=exchange" json:"exchange,omitempty"`
	ExchangeOrderId       string  `protobuf:"bytes,5,opt,name=exchangeOrderId" json:"exchangeOrderId,omitempty"`
	ExchangeMarketName    string  `protobuf:"bytes,6,opt,name=exchangeMarketName" json:"exchangeMarketName,omitempty"`
	MarketName            string  `protobuf:"bytes,7,opt,name=marketName" json:"marketName,omitempty"`
	Side                  string  `protobuf:"bytes,8,opt,name=side" json:"side,omitempty"`
	OrderType             string  `protobuf:"bytes,9,opt,name=orderType" json:"orderType,omitempty"`
	BaseQuantity          float64 `protobuf:"fixed64,10,opt,name=baseQuantity" json:"baseQuantity,omitempty"`
	BaseQuantityRemainder float64 `protobuf:"fixed64,11,opt,name=baseQuantityRemainder" json:"baseQuantityRemainder,omitempty"`
	CurrencyQuantity      float64 `protobuf:"fixed64,12,opt,name=currencyQuantity" json:"currencyQuantity,omitempty"`
	Status                string  `protobuf:"bytes,14,opt,name=status" json:"status,omitempty"`
	Conditions            string  `protobuf:"bytes,15,opt,name=conditions" json:"conditions,omitempty"`
	Condition             string  `protobuf:"bytes,16,opt,name=condition" json:"condition,omitempty"`
	ParentOrderId         string  `protobuf:"bytes,13,opt,name=parentOrderId" json:"parentOrderId,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Order) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Order) GetApiKeyId() string {
	if m != nil {
		return m.ApiKeyId
	}
	return ""
}

func (m *Order) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Order) GetExchangeOrderId() string {
	if m != nil {
		return m.ExchangeOrderId
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

func (m *Order) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *Order) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *Order) GetBaseQuantity() float64 {
	if m != nil {
		return m.BaseQuantity
	}
	return 0
}

func (m *Order) GetBaseQuantityRemainder() float64 {
	if m != nil {
		return m.BaseQuantityRemainder
	}
	return 0
}

func (m *Order) GetCurrencyQuantity() float64 {
	if m != nil {
		return m.CurrencyQuantity
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Order) GetConditions() string {
	if m != nil {
		return m.Conditions
	}
	return ""
}

func (m *Order) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *Order) GetParentOrderId() string {
	if m != nil {
		return m.ParentOrderId
	}
	return ""
}

type UserOrderData struct {
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
}

func (m *UserOrderData) Reset()                    { *m = UserOrderData{} }
func (m *UserOrderData) String() string            { return proto.CompactTextString(m) }
func (*UserOrderData) ProtoMessage()               {}
func (*UserOrderData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*UserOrdersData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*OrderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*OrderListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

func init() {
	proto.RegisterType((*OrdersRequest)(nil), "go.micro.srv.order.OrdersRequest")
	proto.RegisterType((*OrderRequest)(nil), "go.micro.srv.order.OrderRequest")
	proto.RegisterType((*GetUserOrderRequest)(nil), "go.micro.srv.order.GetUserOrderRequest")
	proto.RegisterType((*GetUserOrdersRequest)(nil), "go.micro.srv.order.GetUserOrdersRequest")
	proto.RegisterType((*RemoveOrderRequest)(nil), "go.micro.srv.order.RemoveOrderRequest")
	proto.RegisterType((*Order)(nil), "go.micro.srv.order.Order")
	proto.RegisterType((*UserOrderData)(nil), "go.micro.srv.order.UserOrderData")
	proto.RegisterType((*UserOrdersData)(nil), "go.micro.srv.order.UserOrdersData")
	proto.RegisterType((*OrderResponse)(nil), "go.micro.srv.order.OrderResponse")
	proto.RegisterType((*OrderListResponse)(nil), "go.micro.srv.order.OrderListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for OrderService service

type OrderServiceClient interface {
	AddOrders(ctx context.Context, in *OrdersRequest, opts ...client.CallOption) (*OrderListResponse, error)
	GetUserOrder(ctx context.Context, in *GetUserOrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, opts ...client.CallOption) (*OrderListResponse, error)
	RemoveOrder(ctx context.Context, in *RemoveOrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	UpdateOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
}

type orderServiceClient struct {
	c           client.Client
	serviceName string
}

func NewOrderServiceClient(serviceName string, c client.Client) OrderServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.order"
	}
	return &orderServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *orderServiceClient) AddOrders(ctx context.Context, in *OrdersRequest, opts ...client.CallOption) (*OrderListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.AddOrders", in)
	out := new(OrderListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetUserOrder(ctx context.Context, in *GetUserOrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.GetUserOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, opts ...client.CallOption) (*OrderListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.GetUserOrders", in)
	out := new(OrderListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) RemoveOrder(ctx context.Context, in *RemoveOrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.RemoveOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.UpdateOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	AddOrders(context.Context, *OrdersRequest, *OrderListResponse) error
	GetUserOrder(context.Context, *GetUserOrderRequest, *OrderResponse) error
	GetUserOrders(context.Context, *GetUserOrdersRequest, *OrderListResponse) error
	RemoveOrder(context.Context, *RemoveOrderRequest, *OrderResponse) error
	UpdateOrder(context.Context, *OrderRequest, *OrderResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&OrderService{hdlr}, opts...))
}

type OrderService struct {
	OrderServiceHandler
}

func (h *OrderService) AddOrders(ctx context.Context, in *OrdersRequest, out *OrderListResponse) error {
	return h.OrderServiceHandler.AddOrders(ctx, in, out)
}

func (h *OrderService) GetUserOrder(ctx context.Context, in *GetUserOrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.GetUserOrder(ctx, in, out)
}

func (h *OrderService) GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, out *OrderListResponse) error {
	return h.OrderServiceHandler.GetUserOrders(ctx, in, out)
}

func (h *OrderService) RemoveOrder(ctx context.Context, in *RemoveOrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.RemoveOrder(ctx, in, out)
}

func (h *OrderService) UpdateOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.UpdateOrder(ctx, in, out)
}

func init() { proto.RegisterFile("proto/order/order.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5b, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0xdc, 0xa6, 0x49, 0x33, 0x49, 0xda, 0x7e, 0x03, 0x94, 0xa5, 0xaa, 0x50, 0x6a, 0x71,
	0x89, 0x78, 0x70, 0x45, 0xb8, 0x88, 0x47, 0x10, 0x88, 0xaa, 0xe2, 0x52, 0x61, 0x5a, 0x21, 0x21,
	0x5e, 0xb6, 0xf1, 0xa8, 0x58, 0x28, 0xb6, 0xd9, 0xdd, 0x54, 0xcd, 0x03, 0xbf, 0x85, 0x1f, 0xc6,
	0x0b, 0x3f, 0x05, 0x75, 0x7c, 0x4f, 0xec, 0xb6, 0x28, 0x12, 0x2f, 0x91, 0xe7, 0x9c, 0xb9, 0xec,
	0x8e, 0x8f, 0x8e, 0x03, 0x37, 0x23, 0x15, 0x9a, 0x70, 0x37, 0x54, 0x1e, 0xa9, 0xf8, 0xd7, 0x61,
	0x04, 0xf1, 0x24, 0x74, 0xc6, 0xfe, 0x48, 0x85, 0x8e, 0x56, 0xa7, 0x0e, 0x33, 0xf6, 0x3e, 0xf4,
	0x0e, 0xce, 0x1f, 0xb4, 0x4b, 0xdf, 0x27, 0xa4, 0x0d, 0x3e, 0x83, 0x26, 0x33, 0x5a, 0x58, 0xfd,
	0xe5, 0x41, 0x67, 0xd8, 0x77, 0xe6, 0xab, 0x1c, 0x2e, 0x49, 0x2a, 0xdc, 0x24, 0xdf, 0xfe, 0xbd,
	0x04, 0xdd, 0x22, 0x81, 0x02, 0x5a, 0x4c, 0xed, 0x7b, 0xc2, 0xea, 0x5b, 0x83, 0xb6, 0x9b, 0x86,
	0xb8, 0x09, 0xcd, 0x89, 0x66, 0x62, 0x89, 0x89, 0x24, 0xc2, 0x2d, 0x58, 0x95, 0x91, 0xff, 0x86,
	0xa6, 0xfb, 0x9e, 0x58, 0x66, 0x26, 0x8b, 0xcf, 0x39, 0x3a, 0x1b, 0x7d, 0x95, 0xc1, 0x09, 0x89,
	0x46, 0xcc, 0xa5, 0x31, 0xde, 0x06, 0x18, 0x4b, 0xf5, 0x8d, 0xcc, 0x7b, 0x39, 0x26, 0xb1, 0xc2,
	0x6c, 0x01, 0x41, 0x84, 0x86, 0xf6, 0x3d, 0x12, 0x4d, 0x66, 0xf8, 0x19, 0xb7, 0xa1, 0xcd, 0xc7,
	0x39, 0x9c, 0x46, 0x24, 0x5a, 0x4c, 0xe4, 0x00, 0xda, 0xd0, 0x3d, 0x96, 0x9a, 0x3e, 0x4c, 0x64,
	0x60, 0x7c, 0x33, 0x15, 0xab, 0x7d, 0x6b, 0x60, 0xb9, 0x25, 0x0c, 0x1f, 0xc0, 0xc6, 0x68, 0xa2,
	0x14, 0x05, 0xa3, 0x69, 0x96, 0xd7, 0xe6, 0xbc, 0x39, 0xfc, 0xfc, 0x84, 0xa3, 0x30, 0xf0, 0x7c,
	0xe3, 0x87, 0x81, 0x16, 0x10, 0x9f, 0x30, 0x47, 0xf0, 0x0e, 0xf4, 0x22, 0xa9, 0x28, 0x30, 0x07,
	0xc9, 0xc6, 0x3a, 0x9c, 0x52, 0x06, 0xed, 0x3d, 0xb8, 0xb6, 0x47, 0xe6, 0x48, 0x93, 0x5a, 0x6c,
	0xd1, 0xb6, 0x03, 0xd7, 0x8b, 0x8d, 0xb2, 0xb7, 0x9f, 0xe7, 0x5b, 0xa5, 0xfc, 0xd7, 0x80, 0x2e,
	0x8d, 0xc3, 0x53, 0x5a, 0x70, 0xee, 0xcf, 0x06, 0xac, 0x70, 0x8b, 0x7f, 0x28, 0x8e, 0x01, 0xac,
	0xa7, 0xcf, 0xe9, 0x72, 0x63, 0x85, 0xcc, 0xc2, 0xe8, 0x00, 0xa6, 0xd0, 0xbb, 0x5c, 0x4e, 0xb1,
	0x68, 0x2a, 0x98, 0x19, 0xd9, 0xb5, 0x6a, 0x65, 0xb7, 0x5a, 0x27, 0xbb, 0xf6, 0x65, 0xb2, 0x83,
	0x0a, 0xd9, 0x3d, 0x86, 0x1b, 0xc5, 0xd8, 0xa5, 0xb1, 0xf4, 0x03, 0x8f, 0x14, 0x4b, 0xc6, 0x72,
	0xab, 0xc9, 0x4a, 0xb1, 0x76, 0x6b, 0xc4, 0xba, 0x09, 0x4d, 0x6d, 0xa4, 0x99, 0x68, 0xb1, 0x16,
	0xbf, 0x81, 0x38, 0x9a, 0x11, 0xf1, 0xfa, 0x9c, 0x88, 0xb7, 0xa1, 0x9d, 0x45, 0x62, 0x23, 0xbe,
	0x5b, 0x06, 0xcc, 0x4b, 0xbc, 0x57, 0x25, 0xf1, 0xe7, 0xd0, 0xcb, 0x64, 0xf9, 0x4a, 0x1a, 0x89,
	0xbb, 0xb0, 0xc2, 0xfb, 0x61, 0x99, 0x74, 0x86, 0xb7, 0xea, 0xfd, 0x28, 0xce, 0xb3, 0x5f, 0xc2,
	0x5a, 0x2e, 0x6c, 0x6e, 0xf1, 0x70, 0xc6, 0xd3, 0x2e, 0xe8, 0x91, 0x9a, 0xd9, 0x59, 0xe2, 0x8b,
	0x2e, 0xe9, 0x28, 0x0c, 0x34, 0x15, 0x76, 0x62, 0x95, 0x76, 0x22, 0xa0, 0x35, 0x26, 0xad, 0xe5,
	0x09, 0x25, 0x72, 0x4d, 0x43, 0x7c, 0x02, 0x0d, 0x4f, 0x1a, 0xc9, 0x5a, 0xed, 0x0c, 0x77, 0xaa,
	0x66, 0x96, 0x6e, 0xea, 0x72, 0xba, 0xfd, 0x03, 0xfe, 0x67, 0xe8, 0xad, 0xaf, 0xcd, 0x02, 0xd3,
	0x9f, 0x96, 0xa6, 0xdb, 0x17, 0x4e, 0xd7, 0xf9, 0xf8, 0xe1, 0xaf, 0xe5, 0xc4, 0xc5, 0x3f, 0x92,
	0x3a, 0xf5, 0x47, 0x84, 0x9f, 0xa0, 0xfd, 0xc2, 0xf3, 0xe2, 0x3c, 0xdc, 0xa9, 0xdd, 0x5c, 0x6a,
	0x21, 0x5b, 0x77, 0x6b, 0x53, 0x8a, 0x37, 0xb2, 0xff, 0xc3, 0x2f, 0xd0, 0x2d, 0x7a, 0x10, 0xde,
	0xaf, 0x2a, 0xac, 0xb0, 0xbb, 0xad, 0x9d, 0x0b, 0x3e, 0x49, 0x59, 0xf7, 0x63, 0xe8, 0x95, 0x1c,
	0x0e, 0x07, 0x97, 0xb5, 0xff, 0xfb, 0x1b, 0x7c, 0x86, 0x4e, 0xc1, 0x15, 0xf1, 0x5e, 0x55, 0xdd,
	0xbc, 0x6d, 0x5e, 0xed, 0xfc, 0x87, 0xd0, 0x39, 0x8a, 0x3c, 0x69, 0x92, 0xde, 0x97, 0x7e, 0x86,
	0xaf, 0xd4, 0xf5, 0xb8, 0xc9, 0xff, 0x04, 0x1e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x9a,
	0x9d, 0x86, 0x24, 0x08, 0x00, 0x00,
}
