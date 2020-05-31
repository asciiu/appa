// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/trade/trade.proto

package trade

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TradeEngine service

type TradeEngineService interface {
	Process(ctx context.Context, in *NewOrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	Cancel(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*StatusResponse, error)
	FindOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	FindUserOrders(ctx context.Context, in *UserOrdersRequest, opts ...client.CallOption) (*OrdersPageResponse, error)
}

type tradeEngineService struct {
	c    client.Client
	name string
}

func NewTradeEngineService(name string, c client.Client) TradeEngineService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "tradeengine"
	}
	return &tradeEngineService{
		c:    c,
		name: name,
	}
}

func (c *tradeEngineService) Process(ctx context.Context, in *NewOrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "TradeEngine.Process", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeEngineService) Cancel(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.name, "TradeEngine.Cancel", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeEngineService) FindOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "TradeEngine.FindOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeEngineService) FindUserOrders(ctx context.Context, in *UserOrdersRequest, opts ...client.CallOption) (*OrdersPageResponse, error) {
	req := c.c.NewRequest(c.name, "TradeEngine.FindUserOrders", in)
	out := new(OrdersPageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TradeEngine service

type TradeEngineHandler interface {
	Process(context.Context, *NewOrderRequest, *OrderResponse) error
	Cancel(context.Context, *OrderRequest, *StatusResponse) error
	FindOrder(context.Context, *OrderRequest, *OrderResponse) error
	FindUserOrders(context.Context, *UserOrdersRequest, *OrdersPageResponse) error
}

func RegisterTradeEngineHandler(s server.Server, hdlr TradeEngineHandler, opts ...server.HandlerOption) error {
	type tradeEngine interface {
		Process(ctx context.Context, in *NewOrderRequest, out *OrderResponse) error
		Cancel(ctx context.Context, in *OrderRequest, out *StatusResponse) error
		FindOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		FindUserOrders(ctx context.Context, in *UserOrdersRequest, out *OrdersPageResponse) error
	}
	type TradeEngine struct {
		tradeEngine
	}
	h := &tradeEngineHandler{hdlr}
	return s.Handle(s.NewHandler(&TradeEngine{h}, opts...))
}

type tradeEngineHandler struct {
	TradeEngineHandler
}

func (h *tradeEngineHandler) Process(ctx context.Context, in *NewOrderRequest, out *OrderResponse) error {
	return h.TradeEngineHandler.Process(ctx, in, out)
}

func (h *tradeEngineHandler) Cancel(ctx context.Context, in *OrderRequest, out *StatusResponse) error {
	return h.TradeEngineHandler.Cancel(ctx, in, out)
}

func (h *tradeEngineHandler) FindOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.TradeEngineHandler.FindOrder(ctx, in, out)
}

func (h *tradeEngineHandler) FindUserOrders(ctx context.Context, in *UserOrdersRequest, out *OrdersPageResponse) error {
	return h.TradeEngineHandler.FindUserOrders(ctx, in, out)
}