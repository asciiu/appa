// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asciiu/gomo/plan-service/proto/plan/plan.proto

/*
Package plan is a generated protocol buffer package.

It is generated from these files:
	github.com/asciiu/gomo/plan-service/proto/plan/plan.proto

It has these top-level messages:
	NewPlanRequest
	UpdatePlanRequest
	GetUserPlanRequest
	GetUserPlansRequest
	DeletePlanRequest
	Plan
	PlanWithPagedOrders
	PlanData
	PlanResponse
	PlanWithPagedOrdersResponse
	PlansPageResponse
	PlansPage
*/
package plan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import order "github.com/asciiu/gomo/plan-service/proto/order"

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
type NewPlanRequest struct {
	PlanID          string                   `protobuf:"bytes,1,opt,name=planID" json:"planID"`
	UserID          string                   `protobuf:"bytes,2,opt,name=userID" json:"userID"`
	KeyID           string                   `protobuf:"bytes,3,opt,name=keyID" json:"keyID"`
	PlanTemplateID  string                   `protobuf:"bytes,4,opt,name=planTemplateID" json:"planTemplateID"`
	Exchange        string                   `protobuf:"bytes,5,opt,name=exchange" json:"exchange"`
	MarketName      string                   `protobuf:"bytes,6,opt,name=marketName" json:"marketName"`
	BaseBalance     float64                  `protobuf:"fixed64,7,opt,name=baseBalance" json:"baseBalance"`
	CurrencyBalance float64                  `protobuf:"fixed64,8,opt,name=currencyBalance" json:"currencyBalance"`
	Status          string                   `protobuf:"bytes,9,opt,name=status" json:"status"`
	Orders          []*order.NewOrderRequest `protobuf:"bytes,10,rep,name=orders" json:"orders"`
}

func (m *NewPlanRequest) Reset()                    { *m = NewPlanRequest{} }
func (m *NewPlanRequest) String() string            { return proto.CompactTextString(m) }
func (*NewPlanRequest) ProtoMessage()               {}
func (*NewPlanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewPlanRequest) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *NewPlanRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *NewPlanRequest) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *NewPlanRequest) GetPlanTemplateID() string {
	if m != nil {
		return m.PlanTemplateID
	}
	return ""
}

func (m *NewPlanRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *NewPlanRequest) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *NewPlanRequest) GetBaseBalance() float64 {
	if m != nil {
		return m.BaseBalance
	}
	return 0
}

func (m *NewPlanRequest) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *NewPlanRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *NewPlanRequest) GetOrders() []*order.NewOrderRequest {
	if m != nil {
		return m.Orders
	}
	return nil
}

type UpdatePlanRequest struct {
	PlanID          string                      `protobuf:"bytes,1,opt,name=planID" json:"planID"`
	UserID          string                      `protobuf:"bytes,2,opt,name=userID" json:"userID"`
	Status          string                      `protobuf:"bytes,3,opt,name=status" json:"status"`
	BaseBalance     float64                     `protobuf:"fixed64,4,opt,name=baseBalance" json:"baseBalance"`
	CurrencyBalance float64                     `protobuf:"fixed64,5,opt,name=currencyBalance" json:"currencyBalance"`
	Orders          []*order.UpdateOrderRequest `protobuf:"bytes,6,rep,name=orders" json:"orders"`
}

func (m *UpdatePlanRequest) Reset()                    { *m = UpdatePlanRequest{} }
func (m *UpdatePlanRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePlanRequest) ProtoMessage()               {}
func (*UpdatePlanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdatePlanRequest) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *UpdatePlanRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdatePlanRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UpdatePlanRequest) GetBaseBalance() float64 {
	if m != nil {
		return m.BaseBalance
	}
	return 0
}

func (m *UpdatePlanRequest) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *UpdatePlanRequest) GetOrders() []*order.UpdateOrderRequest {
	if m != nil {
		return m.Orders
	}
	return nil
}

type GetUserPlanRequest struct {
	PlanID   string `protobuf:"bytes,1,opt,name=planID" json:"planID"`
	UserID   string `protobuf:"bytes,2,opt,name=userID" json:"userID"`
	Page     uint32 `protobuf:"varint,3,opt,name=page" json:"page"`
	PageSize uint32 `protobuf:"varint,4,opt,name=pageSize" json:"pageSize"`
}

func (m *GetUserPlanRequest) Reset()                    { *m = GetUserPlanRequest{} }
func (m *GetUserPlanRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserPlanRequest) ProtoMessage()               {}
func (*GetUserPlanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserPlanRequest) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *GetUserPlanRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetUserPlanRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUserPlanRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type GetUserPlansRequest struct {
	UserID     string `protobuf:"bytes,1,opt,name=userID" json:"userID"`
	Exchange   string `protobuf:"bytes,2,opt,name=exchange" json:"exchange"`
	MarketName string `protobuf:"bytes,3,opt,name=marketName" json:"marketName"`
	Status     string `protobuf:"bytes,4,opt,name=status" json:"status"`
	Page       uint32 `protobuf:"varint,5,opt,name=page" json:"page"`
	PageSize   uint32 `protobuf:"varint,6,opt,name=pageSize" json:"pageSize"`
}

func (m *GetUserPlansRequest) Reset()                    { *m = GetUserPlansRequest{} }
func (m *GetUserPlansRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserPlansRequest) ProtoMessage()               {}
func (*GetUserPlansRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetUserPlansRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetUserPlansRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *GetUserPlansRequest) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *GetUserPlansRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetUserPlansRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUserPlansRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type DeletePlanRequest struct {
	PlanID string `protobuf:"bytes,1,opt,name=planID" json:"planID"`
	UserID string `protobuf:"bytes,2,opt,name=userID" json:"userID"`
}

func (m *DeletePlanRequest) Reset()                    { *m = DeletePlanRequest{} }
func (m *DeletePlanRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePlanRequest) ProtoMessage()               {}
func (*DeletePlanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeletePlanRequest) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *DeletePlanRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

// Responses
type Plan struct {
	PlanID                string         `protobuf:"bytes,1,opt,name=planID" json:"planID"`
	PlanTemplateID        string         `protobuf:"bytes,2,opt,name=planTemplateID" json:"planTemplateID"`
	UserID                string         `protobuf:"bytes,3,opt,name=userID" json:"userID"`
	KeyID                 string         `protobuf:"bytes,4,opt,name=keyID" json:"keyID"`
	Key                   string         `protobuf:"bytes,5,opt,name=key" json:"key"`
	KeySecret             string         `protobuf:"bytes,6,opt,name=keySecret" json:"keySecret"`
	KeyDescription        string         `protobuf:"bytes,7,opt,name=keyDescription" json:"keyDescription"`
	LastExecutedPlanDepth uint32         `protobuf:"varint,8,opt,name=lastExecutedPlanDepth" json:"lastExecutedPlanDepth"`
	LastExecutedOrderID   string         `protobuf:"bytes,9,opt,name=lastExecutedOrderID" json:"lastExecutedOrderID"`
	Exchange              string         `protobuf:"bytes,10,opt,name=exchange" json:"exchange"`
	ExchangeMarketName    string         `protobuf:"bytes,11,opt,name=exchangeMarketName" json:"exchangeMarketName"`
	MarketName            string         `protobuf:"bytes,12,opt,name=marketName" json:"marketName"`
	BaseBalance           float64        `protobuf:"fixed64,13,opt,name=baseBalance" json:"baseBalance"`
	CurrencyBalance       float64        `protobuf:"fixed64,14,opt,name=currencyBalance" json:"currencyBalance"`
	Status                string         `protobuf:"bytes,15,opt,name=status" json:"status"`
	CreatedOn             string         `protobuf:"bytes,16,opt,name=createdOn" json:"createdOn"`
	UpdatedOn             string         `protobuf:"bytes,17,opt,name=updatedOn" json:"updatedOn"`
	Orders                []*order.Order `protobuf:"bytes,18,rep,name=orders" json:"orders"`
}

func (m *Plan) Reset()                    { *m = Plan{} }
func (m *Plan) String() string            { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()               {}
func (*Plan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Plan) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *Plan) GetPlanTemplateID() string {
	if m != nil {
		return m.PlanTemplateID
	}
	return ""
}

func (m *Plan) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Plan) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *Plan) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Plan) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *Plan) GetKeyDescription() string {
	if m != nil {
		return m.KeyDescription
	}
	return ""
}

func (m *Plan) GetLastExecutedPlanDepth() uint32 {
	if m != nil {
		return m.LastExecutedPlanDepth
	}
	return 0
}

func (m *Plan) GetLastExecutedOrderID() string {
	if m != nil {
		return m.LastExecutedOrderID
	}
	return ""
}

func (m *Plan) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Plan) GetExchangeMarketName() string {
	if m != nil {
		return m.ExchangeMarketName
	}
	return ""
}

func (m *Plan) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *Plan) GetBaseBalance() float64 {
	if m != nil {
		return m.BaseBalance
	}
	return 0
}

func (m *Plan) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *Plan) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Plan) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *Plan) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *Plan) GetOrders() []*order.Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type PlanWithPagedOrders struct {
	PlanID             string            `protobuf:"bytes,1,opt,name=planID" json:"planID"`
	PlanTemplateID     string            `protobuf:"bytes,2,opt,name=planTemplateID" json:"planTemplateID"`
	UserID             string            `protobuf:"bytes,3,opt,name=userID" json:"userID"`
	KeyID              string            `protobuf:"bytes,4,opt,name=keyID" json:"keyID"`
	Key                string            `protobuf:"bytes,5,opt,name=key" json:"key"`
	KeySecret          string            `protobuf:"bytes,6,opt,name=keySecret" json:"keySecret"`
	KeyDescription     string            `protobuf:"bytes,7,opt,name=keyDescription" json:"keyDescription"`
	Exchange           string            `protobuf:"bytes,8,opt,name=exchange" json:"exchange"`
	ExchangeMarketName string            `protobuf:"bytes,9,opt,name=exchangeMarketName" json:"exchangeMarketName"`
	MarketName         string            `protobuf:"bytes,10,opt,name=marketName" json:"marketName"`
	BaseBalance        float64           `protobuf:"fixed64,11,opt,name=baseBalance" json:"baseBalance"`
	CurrencyBalance    float64           `protobuf:"fixed64,12,opt,name=currencyBalance" json:"currencyBalance"`
	Status             string            `protobuf:"bytes,13,opt,name=status" json:"status"`
	CreatedOn          string            `protobuf:"bytes,14,opt,name=createdOn" json:"createdOn"`
	UpdatedOn          string            `protobuf:"bytes,15,opt,name=updatedOn" json:"updatedOn"`
	OrdersPage         *order.OrdersPage `protobuf:"bytes,16,opt,name=ordersPage" json:"ordersPage"`
}

func (m *PlanWithPagedOrders) Reset()                    { *m = PlanWithPagedOrders{} }
func (m *PlanWithPagedOrders) String() string            { return proto.CompactTextString(m) }
func (*PlanWithPagedOrders) ProtoMessage()               {}
func (*PlanWithPagedOrders) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PlanWithPagedOrders) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *PlanWithPagedOrders) GetPlanTemplateID() string {
	if m != nil {
		return m.PlanTemplateID
	}
	return ""
}

func (m *PlanWithPagedOrders) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *PlanWithPagedOrders) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *PlanWithPagedOrders) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PlanWithPagedOrders) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *PlanWithPagedOrders) GetKeyDescription() string {
	if m != nil {
		return m.KeyDescription
	}
	return ""
}

func (m *PlanWithPagedOrders) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *PlanWithPagedOrders) GetExchangeMarketName() string {
	if m != nil {
		return m.ExchangeMarketName
	}
	return ""
}

func (m *PlanWithPagedOrders) GetMarketName() string {
	if m != nil {
		return m.MarketName
	}
	return ""
}

func (m *PlanWithPagedOrders) GetBaseBalance() float64 {
	if m != nil {
		return m.BaseBalance
	}
	return 0
}

func (m *PlanWithPagedOrders) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

func (m *PlanWithPagedOrders) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PlanWithPagedOrders) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *PlanWithPagedOrders) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *PlanWithPagedOrders) GetOrdersPage() *order.OrdersPage {
	if m != nil {
		return m.OrdersPage
	}
	return nil
}

type PlanData struct {
	Plan *Plan `protobuf:"bytes,1,opt,name=plan" json:"plan"`
}

func (m *PlanData) Reset()                    { *m = PlanData{} }
func (m *PlanData) String() string            { return proto.CompactTextString(m) }
func (*PlanData) ProtoMessage()               {}
func (*PlanData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PlanData) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

type PlanResponse struct {
	Status  string    `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string    `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *PlanData `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *PlanResponse) Reset()                    { *m = PlanResponse{} }
func (m *PlanResponse) String() string            { return proto.CompactTextString(m) }
func (*PlanResponse) ProtoMessage()               {}
func (*PlanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PlanResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PlanResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PlanResponse) GetData() *PlanData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PlanWithPagedOrdersResponse struct {
	Status  string               `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string               `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *PlanWithPagedOrders `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *PlanWithPagedOrdersResponse) Reset()                    { *m = PlanWithPagedOrdersResponse{} }
func (m *PlanWithPagedOrdersResponse) String() string            { return proto.CompactTextString(m) }
func (*PlanWithPagedOrdersResponse) ProtoMessage()               {}
func (*PlanWithPagedOrdersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PlanWithPagedOrdersResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PlanWithPagedOrdersResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PlanWithPagedOrdersResponse) GetData() *PlanWithPagedOrders {
	if m != nil {
		return m.Data
	}
	return nil
}

type PlansPageResponse struct {
	Status  string     `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string     `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *PlansPage `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *PlansPageResponse) Reset()                    { *m = PlansPageResponse{} }
func (m *PlansPageResponse) String() string            { return proto.CompactTextString(m) }
func (*PlansPageResponse) ProtoMessage()               {}
func (*PlansPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PlansPageResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PlansPageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PlansPageResponse) GetData() *PlansPage {
	if m != nil {
		return m.Data
	}
	return nil
}

type PlansPage struct {
	Page     uint32  `protobuf:"varint,1,opt,name=page" json:"page"`
	PageSize uint32  `protobuf:"varint,2,opt,name=pageSize" json:"pageSize"`
	Total    uint32  `protobuf:"varint,3,opt,name=total" json:"total"`
	Plans    []*Plan `protobuf:"bytes,4,rep,name=plans" json:"plans"`
}

func (m *PlansPage) Reset()                    { *m = PlansPage{} }
func (m *PlansPage) String() string            { return proto.CompactTextString(m) }
func (*PlansPage) ProtoMessage()               {}
func (*PlansPage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PlansPage) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PlansPage) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PlansPage) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PlansPage) GetPlans() []*Plan {
	if m != nil {
		return m.Plans
	}
	return nil
}

func init() {
	proto.RegisterType((*NewPlanRequest)(nil), "plan.NewPlanRequest")
	proto.RegisterType((*UpdatePlanRequest)(nil), "plan.UpdatePlanRequest")
	proto.RegisterType((*GetUserPlanRequest)(nil), "plan.GetUserPlanRequest")
	proto.RegisterType((*GetUserPlansRequest)(nil), "plan.GetUserPlansRequest")
	proto.RegisterType((*DeletePlanRequest)(nil), "plan.DeletePlanRequest")
	proto.RegisterType((*Plan)(nil), "plan.Plan")
	proto.RegisterType((*PlanWithPagedOrders)(nil), "plan.PlanWithPagedOrders")
	proto.RegisterType((*PlanData)(nil), "plan.PlanData")
	proto.RegisterType((*PlanResponse)(nil), "plan.PlanResponse")
	proto.RegisterType((*PlanWithPagedOrdersResponse)(nil), "plan.PlanWithPagedOrdersResponse")
	proto.RegisterType((*PlansPageResponse)(nil), "plan.PlansPageResponse")
	proto.RegisterType((*PlansPage)(nil), "plan.PlansPage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PlanService service

type PlanServiceClient interface {
	NewPlan(ctx context.Context, in *NewPlanRequest, opts ...client.CallOption) (*PlanResponse, error)
	GetUserPlan(ctx context.Context, in *GetUserPlanRequest, opts ...client.CallOption) (*PlanWithPagedOrdersResponse, error)
	GetUserPlans(ctx context.Context, in *GetUserPlansRequest, opts ...client.CallOption) (*PlansPageResponse, error)
	DeletePlan(ctx context.Context, in *DeletePlanRequest, opts ...client.CallOption) (*PlanResponse, error)
	UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...client.CallOption) (*PlanResponse, error)
}

type planServiceClient struct {
	c           client.Client
	serviceName string
}

func NewPlanServiceClient(serviceName string, c client.Client) PlanServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "plan"
	}
	return &planServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *planServiceClient) NewPlan(ctx context.Context, in *NewPlanRequest, opts ...client.CallOption) (*PlanResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.NewPlan", in)
	out := new(PlanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) GetUserPlan(ctx context.Context, in *GetUserPlanRequest, opts ...client.CallOption) (*PlanWithPagedOrdersResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.GetUserPlan", in)
	out := new(PlanWithPagedOrdersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) GetUserPlans(ctx context.Context, in *GetUserPlansRequest, opts ...client.CallOption) (*PlansPageResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.GetUserPlans", in)
	out := new(PlansPageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) DeletePlan(ctx context.Context, in *DeletePlanRequest, opts ...client.CallOption) (*PlanResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.DeletePlan", in)
	out := new(PlanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...client.CallOption) (*PlanResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PlanService.UpdatePlan", in)
	out := new(PlanResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlanService service

type PlanServiceHandler interface {
	NewPlan(context.Context, *NewPlanRequest, *PlanResponse) error
	GetUserPlan(context.Context, *GetUserPlanRequest, *PlanWithPagedOrdersResponse) error
	GetUserPlans(context.Context, *GetUserPlansRequest, *PlansPageResponse) error
	DeletePlan(context.Context, *DeletePlanRequest, *PlanResponse) error
	UpdatePlan(context.Context, *UpdatePlanRequest, *PlanResponse) error
}

func RegisterPlanServiceHandler(s server.Server, hdlr PlanServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PlanService{hdlr}, opts...))
}

type PlanService struct {
	PlanServiceHandler
}

func (h *PlanService) NewPlan(ctx context.Context, in *NewPlanRequest, out *PlanResponse) error {
	return h.PlanServiceHandler.NewPlan(ctx, in, out)
}

func (h *PlanService) GetUserPlan(ctx context.Context, in *GetUserPlanRequest, out *PlanWithPagedOrdersResponse) error {
	return h.PlanServiceHandler.GetUserPlan(ctx, in, out)
}

func (h *PlanService) GetUserPlans(ctx context.Context, in *GetUserPlansRequest, out *PlansPageResponse) error {
	return h.PlanServiceHandler.GetUserPlans(ctx, in, out)
}

func (h *PlanService) DeletePlan(ctx context.Context, in *DeletePlanRequest, out *PlanResponse) error {
	return h.PlanServiceHandler.DeletePlan(ctx, in, out)
}

func (h *PlanService) UpdatePlan(ctx context.Context, in *UpdatePlanRequest, out *PlanResponse) error {
	return h.PlanServiceHandler.UpdatePlan(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/asciiu/gomo/plan-service/proto/plan/plan.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 894 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xc6, 0x89, 0x93, 0x99, 0x94, 0xf3, 0xb3, 0xe9, 0x59, 0x16, 0x4f, 0x58, 0xad, 0x82, 0x41,
	0x28, 0x42, 0xda, 0x84, 0x0d, 0x70, 0x40, 0x73, 0x03, 0x23, 0x14, 0x09, 0x86, 0x95, 0x87, 0x15,
	0xe7, 0x1e, 0xa7, 0x94, 0x98, 0x24, 0xb6, 0x71, 0xb7, 0xd9, 0x0d, 0x07, 0x1e, 0x84, 0xb7, 0xe0,
	0x45, 0x38, 0x70, 0xe4, 0x01, 0x78, 0x0e, 0xd4, 0xdd, 0x8e, 0xff, 0xe2, 0x4c, 0x22, 0x72, 0xdb,
	0x4b, 0xe2, 0xaa, 0xea, 0xae, 0xea, 0xaa, 0xfe, 0xbe, 0xae, 0x82, 0x2f, 0x17, 0x1e, 0x5f, 0xc6,
	0xf7, 0x63, 0x37, 0xd8, 0x4c, 0x28, 0x73, 0x3d, 0x2f, 0x9e, 0x2c, 0x82, 0x4d, 0x30, 0x09, 0xd7,
	0xd4, 0x7f, 0xce, 0x30, 0xfa, 0xd5, 0x73, 0x71, 0x12, 0x46, 0x01, 0x57, 0x2a, 0xf9, 0x33, 0x96,
	0x32, 0xd1, 0xc5, 0xf7, 0xe0, 0xe6, 0x74, 0x07, 0x41, 0x34, 0xc7, 0x48, 0xfd, 0x2a, 0x17, 0xd6,
	0x5f, 0x35, 0xe8, 0xde, 0xe2, 0xeb, 0x97, 0x6b, 0xea, 0x3b, 0xf8, 0x4b, 0x8c, 0x8c, 0x93, 0x27,
	0xd0, 0x14, 0x5b, 0x67, 0xb6, 0xa9, 0x0d, 0xb5, 0x51, 0xcb, 0x49, 0x24, 0xa1, 0x8f, 0x19, 0x46,
	0x33, 0xdb, 0xac, 0x29, 0xbd, 0x92, 0xc8, 0x63, 0x68, 0xac, 0x70, 0x3b, 0xb3, 0xcd, 0xba, 0x54,
	0x2b, 0x81, 0x7c, 0x0c, 0x5d, 0xb1, 0xef, 0x47, 0xdc, 0x84, 0x6b, 0xca, 0x71, 0x66, 0x9b, 0xba,
	0x34, 0x97, 0xb4, 0x64, 0x00, 0x97, 0xf8, 0xc6, 0x5d, 0x52, 0x7f, 0x81, 0x66, 0x43, 0xae, 0x48,
	0x65, 0xf2, 0x0c, 0x60, 0x43, 0xa3, 0x15, 0xf2, 0x5b, 0xba, 0x41, 0xb3, 0x29, 0xad, 0x39, 0x0d,
	0x19, 0x82, 0x71, 0x4f, 0x19, 0x7e, 0x45, 0xd7, 0xd4, 0x77, 0xd1, 0xbc, 0x18, 0x6a, 0x23, 0xcd,
	0xc9, 0xab, 0xc8, 0x08, 0x7a, 0x6e, 0x1c, 0x45, 0xe8, 0xbb, 0xdb, 0xdd, 0xaa, 0x4b, 0xb9, 0xaa,
	0xac, 0x16, 0xd9, 0x31, 0x4e, 0x79, 0xcc, 0xcc, 0x96, 0xca, 0x4e, 0x49, 0x64, 0x0c, 0x4d, 0x59,
	0x2f, 0x66, 0xc2, 0xb0, 0x3e, 0x32, 0xa6, 0x4f, 0xc6, 0xaa, 0x7c, 0xb7, 0xf8, 0xfa, 0x07, 0xf1,
	0x91, 0x54, 0xcd, 0x49, 0x56, 0x59, 0xff, 0x68, 0xd0, 0x7f, 0x15, 0xce, 0x29, 0xc7, 0x73, 0x6a,
	0x9a, 0x9d, 0xa6, 0x5e, 0x38, 0x4d, 0x29, 0x63, 0xfd, 0xa4, 0x8c, 0x1b, 0xd5, 0x19, 0xbf, 0x48,
	0x33, 0x6b, 0xca, 0xcc, 0xae, 0x93, 0xcc, 0xd4, 0xe9, 0x2b, 0x93, 0xe3, 0x40, 0xbe, 0x45, 0xfe,
	0x8a, 0x61, 0x74, 0x4e, 0x72, 0x04, 0xf4, 0x90, 0x2e, 0x50, 0xa6, 0xd6, 0x71, 0xe4, 0xb7, 0x80,
	0x81, 0xf8, 0xbf, 0xf3, 0x7e, 0x53, 0x59, 0x75, 0x9c, 0x54, 0xb6, 0xfe, 0xd4, 0xe0, 0x2a, 0x17,
	0x96, 0xe5, 0xe2, 0x26, 0xfe, 0xb5, 0x82, 0xff, 0x3c, 0xa4, 0x6a, 0x0f, 0x42, 0xaa, 0xbe, 0x07,
	0xa9, 0xac, 0xf0, 0x7a, 0xa1, 0xf0, 0xbb, 0x33, 0x37, 0x0e, 0x9c, 0xb9, 0x59, 0x3a, 0xf3, 0xd7,
	0xd0, 0xb7, 0x71, 0x8d, 0x67, 0xa1, 0xc0, 0xfa, 0x57, 0x07, 0x5d, 0xec, 0x3f, 0xb8, 0x71, 0x9f,
	0x64, 0xb5, 0x4a, 0x92, 0x65, 0x01, 0xea, 0xd5, 0xd4, 0xd5, 0xf3, 0xd4, 0x7d, 0x04, 0xf5, 0x15,
	0x6e, 0x13, 0x36, 0x8a, 0x4f, 0xf2, 0x14, 0x5a, 0x2b, 0xdc, 0xde, 0xa1, 0x1b, 0x21, 0x4f, 0x78,
	0x98, 0x29, 0xc4, 0x29, 0x56, 0xb8, 0xb5, 0x91, 0xb9, 0x91, 0x17, 0x72, 0x2f, 0xf0, 0x25, 0x13,
	0x5b, 0x4e, 0x49, 0x4b, 0x3e, 0x87, 0x77, 0xd7, 0x94, 0xf1, 0x6f, 0xde, 0xa0, 0x1b, 0x73, 0x9c,
	0x8b, 0xcc, 0x6c, 0x0c, 0xf9, 0x52, 0x52, 0xb2, 0xe3, 0x54, 0x1b, 0xc9, 0xa7, 0x70, 0x95, 0x37,
	0x48, 0x5c, 0xce, 0xec, 0x84, 0xa5, 0x55, 0xa6, 0xc2, 0xfd, 0x43, 0xe9, 0xfe, 0xc7, 0x40, 0x76,
	0xdf, 0xdf, 0x67, 0x38, 0x30, 0xe4, 0xaa, 0x0a, 0x4b, 0x09, 0x2f, 0xed, 0x63, 0x4f, 0x50, 0xe7,
	0x24, 0x42, 0x76, 0x8f, 0x3d, 0x41, 0xbd, 0x02, 0xf6, 0x9e, 0x42, 0xcb, 0x8d, 0x90, 0x8a, 0x0c,
	0x7d, 0xf3, 0x91, 0xaa, 0x7e, 0xaa, 0x10, 0xd6, 0x58, 0x32, 0x56, 0x58, 0xfb, 0xca, 0x9a, 0x2a,
	0xc8, 0x47, 0x29, 0xc9, 0x89, 0x24, 0x79, 0x3b, 0x21, 0xb9, 0xa2, 0xf7, 0x8e, 0xd7, 0x7f, 0xe8,
	0x70, 0x25, 0x2a, 0xfe, 0x93, 0xc7, 0x97, 0x2f, 0xe9, 0x22, 0x29, 0x25, 0x7b, 0x4b, 0x70, 0x97,
	0xc7, 0xc3, 0xe5, 0x49, 0x78, 0x68, 0x9d, 0x88, 0x07, 0x38, 0x86, 0x07, 0xe3, 0x24, 0x3c, 0xb4,
	0x8f, 0xe1, 0xa1, 0x73, 0x18, 0x0f, 0xdd, 0x07, 0xf1, 0xd0, 0x2b, 0xe3, 0xe1, 0x05, 0x80, 0xba,
	0x73, 0x71, 0xcd, 0x12, 0x4c, 0xc6, 0xb4, 0x9f, 0xc7, 0x84, 0x34, 0x38, 0xb9, 0x45, 0xd6, 0x27,
	0x70, 0x29, 0xd9, 0x48, 0x39, 0x25, 0xcf, 0x40, 0xce, 0x1c, 0x12, 0x0e, 0xc6, 0x14, 0xc6, 0x72,
	0x18, 0x91, 0x4f, 0x9c, 0xd4, 0x5b, 0x73, 0x68, 0xab, 0x07, 0x8f, 0x85, 0x81, 0xcf, 0xf2, 0x29,
	0x68, 0x85, 0x14, 0x4c, 0xb8, 0xd8, 0x20, 0x63, 0x34, 0x7d, 0xa1, 0x77, 0x22, 0xb1, 0x40, 0x9f,
	0x53, 0x4e, 0x25, 0x60, 0x8c, 0x69, 0x37, 0x8b, 0x20, 0xe2, 0x3b, 0xd2, 0x66, 0xfd, 0x0e, 0xef,
	0x57, 0xa0, 0xf5, 0x8c, 0xa0, 0xcf, 0x0b, 0x41, 0xaf, 0xb3, 0xa0, 0xe5, 0x10, 0x2a, 0xfe, 0xcf,
	0xd0, 0x97, 0x8d, 0x48, 0x96, 0xea, 0xff, 0x47, 0xfd, 0xb0, 0x10, 0xb5, 0x97, 0x45, 0x55, 0x8e,
	0x55, 0x2c, 0x06, 0xad, 0x54, 0x95, 0x76, 0x21, 0xed, 0x40, 0x17, 0xaa, 0x15, 0xbb, 0x90, 0xe0,
	0x19, 0x0f, 0x38, 0x5d, 0x27, 0xad, 0x56, 0x09, 0x64, 0x08, 0x0d, 0x11, 0x4a, 0xb4, 0xb8, 0x7a,
	0xe9, 0x16, 0x95, 0x61, 0xfa, 0x77, 0x0d, 0x0c, 0x21, 0xdf, 0xa9, 0xe9, 0x91, 0x7c, 0x01, 0x17,
	0xc9, 0x90, 0x48, 0x1e, 0xab, 0xd5, 0xc5, 0x99, 0x71, 0x40, 0x72, 0x3e, 0x92, 0x82, 0x58, 0xef,
	0x90, 0xef, 0xc0, 0xc8, 0xf5, 0x6d, 0x62, 0xaa, 0x45, 0xfb, 0x13, 0xc4, 0xe0, 0x83, 0xc3, 0x15,
	0xcf, 0xbc, 0xd9, 0xd0, 0xce, 0x4f, 0x01, 0xe4, 0x7a, 0xcf, 0xdd, 0x6e, 0x32, 0x18, 0xbc, 0x57,
	0xae, 0x65, 0xe6, 0xe5, 0x06, 0x20, 0x6b, 0xcc, 0x24, 0x59, 0xb8, 0xd7, 0xaa, 0x0f, 0x24, 0x74,
	0x03, 0x90, 0xcd, 0x76, 0xbb, 0xcd, 0x7b, 0xd3, 0x5e, 0xf5, 0xe6, 0xfb, 0xa6, 0x9c, 0xb8, 0x3f,
	0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x99, 0x38, 0xa6, 0x3e, 0xf1, 0x0b, 0x00, 0x00,
}
