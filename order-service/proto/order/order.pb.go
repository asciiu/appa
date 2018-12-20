// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order/order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Requests
type OrderRequest struct {
	OrderID              string   `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{0}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *OrderRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type UserOrdersRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Page                 uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             uint32   `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserOrdersRequest) Reset()         { *m = UserOrdersRequest{} }
func (m *UserOrdersRequest) String() string { return proto.CompactTextString(m) }
func (*UserOrdersRequest) ProtoMessage()    {}
func (*UserOrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{1}
}

func (m *UserOrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOrdersRequest.Unmarshal(m, b)
}
func (m *UserOrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOrdersRequest.Marshal(b, m, deterministic)
}
func (m *UserOrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrdersRequest.Merge(m, src)
}
func (m *UserOrdersRequest) XXX_Size() int {
	return xxx_messageInfo_UserOrdersRequest.Size(m)
}
func (m *UserOrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrdersRequest proto.InternalMessageInfo

func (m *UserOrdersRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserOrdersRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *UserOrdersRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *UserOrdersRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type NewOrderRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	MarketName           string   `protobuf:"bytes,2,opt,name=marketName,proto3" json:"marketName,omitempty"`
	Side                 string   `protobuf:"bytes,3,opt,name=side,proto3" json:"side,omitempty"`
	Size                 float64  `protobuf:"fixed64,4,opt,name=size,proto3" json:"size,omitempty"`
	Price                float64  `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewOrderRequest) Reset()         { *m = NewOrderRequest{} }
func (m *NewOrderRequest) String() string { return proto.CompactTextString(m) }
func (*NewOrderRequest) ProtoMessage()    {}
func (*NewOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{2}
}

func (m *NewOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewOrderRequest.Unmarshal(m, b)
}
func (m *NewOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewOrderRequest.Marshal(b, m, deterministic)
}
func (m *NewOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewOrderRequest.Merge(m, src)
}
func (m *NewOrderRequest) XXX_Size() int {
	return xxx_messageInfo_NewOrderRequest.Size(m)
}
func (m *NewOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewOrderRequest proto.InternalMessageInfo

func (m *NewOrderRequest) GetUserID() string {
	if m != nil {
		return m.UserID
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

func (m *NewOrderRequest) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *NewOrderRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *NewOrderRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// Responses
type Order struct {
	OrderID              string   `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	MarketName           string   `protobuf:"bytes,3,opt,name=marketName,proto3" json:"marketName,omitempty"`
	Side                 string   `protobuf:"bytes,4,opt,name=side,proto3" json:"side,omitempty"`
	Size                 float64  `protobuf:"fixed64,5,opt,name=size,proto3" json:"size,omitempty"`
	Price                float64  `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedOn            string   `protobuf:"bytes,9,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	UpdatedOn            string   `protobuf:"bytes,10,opt,name=updatedOn,proto3" json:"updatedOn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{3}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *Order) GetUserID() string {
	if m != nil {
		return m.UserID
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

func (m *Order) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Order) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
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

type OrderData struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderData) Reset()         { *m = OrderData{} }
func (m *OrderData) String() string { return proto.CompactTextString(m) }
func (*OrderData) ProtoMessage()    {}
func (*OrderData) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{4}
}

func (m *OrderData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderData.Unmarshal(m, b)
}
func (m *OrderData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderData.Marshal(b, m, deterministic)
}
func (m *OrderData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderData.Merge(m, src)
}
func (m *OrderData) XXX_Size() int {
	return xxx_messageInfo_OrderData.Size(m)
}
func (m *OrderData) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderData.DiscardUnknown(m)
}

var xxx_messageInfo_OrderData proto.InternalMessageInfo

func (m *OrderData) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type OrdersPage struct {
	Page                 uint32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             uint32   `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Total                uint32   `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Orders               []*Order `protobuf:"bytes,4,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrdersPage) Reset()         { *m = OrdersPage{} }
func (m *OrdersPage) String() string { return proto.CompactTextString(m) }
func (*OrdersPage) ProtoMessage()    {}
func (*OrdersPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{5}
}

func (m *OrdersPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdersPage.Unmarshal(m, b)
}
func (m *OrdersPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdersPage.Marshal(b, m, deterministic)
}
func (m *OrdersPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersPage.Merge(m, src)
}
func (m *OrdersPage) XXX_Size() int {
	return xxx_messageInfo_OrdersPage.Size(m)
}
func (m *OrdersPage) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersPage.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersPage proto.InternalMessageInfo

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

type OrdersPageResponse struct {
	Status               string      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *OrdersPage `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OrdersPageResponse) Reset()         { *m = OrdersPageResponse{} }
func (m *OrdersPageResponse) String() string { return proto.CompactTextString(m) }
func (*OrdersPageResponse) ProtoMessage()    {}
func (*OrdersPageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{6}
}

func (m *OrdersPageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdersPageResponse.Unmarshal(m, b)
}
func (m *OrdersPageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdersPageResponse.Marshal(b, m, deterministic)
}
func (m *OrdersPageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersPageResponse.Merge(m, src)
}
func (m *OrdersPageResponse) XXX_Size() int {
	return xxx_messageInfo_OrdersPageResponse.Size(m)
}
func (m *OrdersPageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersPageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersPageResponse proto.InternalMessageInfo

func (m *OrdersPageResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrdersPageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OrdersPageResponse) GetData() *OrdersPage {
	if m != nil {
		return m.Data
	}
	return nil
}

type OrderResponse struct {
	Status               string     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *OrderData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OrderResponse) Reset()         { *m = OrderResponse{} }
func (m *OrderResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()    {}
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{7}
}

func (m *OrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResponse.Unmarshal(m, b)
}
func (m *OrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResponse.Marshal(b, m, deterministic)
}
func (m *OrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResponse.Merge(m, src)
}
func (m *OrderResponse) XXX_Size() int {
	return xxx_messageInfo_OrderResponse.Size(m)
}
func (m *OrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResponse proto.InternalMessageInfo

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

func (m *OrderResponse) GetData() *OrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

type StatusResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{8}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderRequest)(nil), "OrderRequest")
	proto.RegisterType((*UserOrdersRequest)(nil), "UserOrdersRequest")
	proto.RegisterType((*NewOrderRequest)(nil), "NewOrderRequest")
	proto.RegisterType((*Order)(nil), "Order")
	proto.RegisterType((*OrderData)(nil), "OrderData")
	proto.RegisterType((*OrdersPage)(nil), "OrdersPage")
	proto.RegisterType((*OrdersPageResponse)(nil), "OrdersPageResponse")
	proto.RegisterType((*OrderResponse)(nil), "OrderResponse")
	proto.RegisterType((*StatusResponse)(nil), "StatusResponse")
}

func init() { proto.RegisterFile("proto/order/order.proto", fileDescriptor_986e030a471601a2) }

var fileDescriptor_986e030a471601a2 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xad, 0xd7, 0x26, 0x6d, 0x6e, 0x69, 0x07, 0x66, 0x82, 0xa8, 0x9a, 0x0a, 0xf2, 0x13, 0x48,
	0xc8, 0x93, 0xca, 0x23, 0x2f, 0x7c, 0x4c, 0x48, 0xbc, 0x6c, 0x28, 0x15, 0x3f, 0xc0, 0x34, 0x57,
	0x53, 0xc5, 0xda, 0x04, 0xdb, 0x05, 0xb1, 0xbf, 0xc2, 0x3f, 0xe3, 0xcf, 0x80, 0x7c, 0xed, 0x34,
	0x49, 0x3f, 0x1e, 0xd8, 0xcb, 0xe6, 0x7b, 0xee, 0xcd, 0x39, 0x27, 0x37, 0xa7, 0x86, 0xa7, 0xa5,
	0x2e, 0x6c, 0x71, 0x51, 0xe8, 0x1c, 0xb5, 0xff, 0x2b, 0x09, 0x11, 0x6f, 0xe1, 0xc1, 0xb5, 0x2b,
	0x33, 0xfc, 0xbe, 0x41, 0x63, 0x79, 0x0a, 0x7d, 0x6a, 0x7f, 0xba, 0x4c, 0xd9, 0x73, 0xf6, 0x22,
	0xc9, 0xaa, 0x92, 0x3f, 0x81, 0x78, 0x63, 0xa8, 0x71, 0x42, 0x8d, 0x50, 0x09, 0x03, 0x8f, 0xbe,
	0x18, 0xd4, 0xc4, 0x62, 0x2a, 0x9a, 0x7a, 0x98, 0x35, 0x87, 0x39, 0x87, 0x5e, 0xa9, 0x6e, 0x90,
	0x28, 0x46, 0x19, 0x9d, 0xf9, 0x04, 0x06, 0xee, 0xff, 0x7c, 0x79, 0x87, 0x69, 0x97, 0xf0, 0x6d,
	0xed, 0x78, 0x8c, 0x55, 0x76, 0x63, 0xd2, 0x9e, 0xe7, 0xf1, 0x95, 0xf8, 0xcd, 0xe0, 0xf4, 0x0a,
	0x7f, 0xb6, 0xac, 0x1f, 0xd3, 0x9c, 0x02, 0xac, 0x94, 0xfe, 0x86, 0xf6, 0x4a, 0xad, 0x30, 0x98,
	0x6f, 0x20, 0xce, 0x93, 0x59, 0xe6, 0x5e, 0x3b, 0xc9, 0xe8, 0xec, 0xb1, 0x3b, 0x24, 0x55, 0x96,
	0xd1, 0x99, 0x9f, 0x41, 0x54, 0xea, 0xe5, 0x02, 0xd3, 0x88, 0x40, 0x5f, 0xb8, 0x49, 0xfb, 0xab,
	0xc4, 0x34, 0xf6, 0x4f, 0xbb, 0xb3, 0xf8, 0xcb, 0x20, 0x22, 0x6b, 0xff, 0xbf, 0xce, 0x1d, 0xb7,
	0xdd, 0xa3, 0x6e, 0x7b, 0x07, 0xdc, 0x46, 0x87, 0xdc, 0xc6, 0x87, 0xdc, 0xf6, 0x6b, 0xb7, 0x8d,
	0x1d, 0x0f, 0x9a, 0x3b, 0xe6, 0xe7, 0x90, 0x2c, 0x34, 0x2a, 0x8b, 0xf9, 0xf5, 0x3a, 0x4d, 0xa8,
	0x55, 0x03, 0xae, 0xbb, 0x29, 0xf3, 0xd0, 0x05, 0xdf, 0xdd, 0x02, 0xe2, 0x25, 0x24, 0xb4, 0x80,
	0x4b, 0x65, 0x15, 0x3f, 0x87, 0x88, 0xde, 0x9a, 0x56, 0x30, 0x9c, 0xc5, 0xd2, 0x7f, 0x36, 0x0f,
	0x0a, 0x0d, 0xe0, 0xb3, 0xf3, 0xd9, 0x85, 0xa1, 0x0a, 0x08, 0x3b, 0x12, 0x90, 0x93, 0x9d, 0x80,
	0x9c, 0x41, 0x64, 0x0b, 0xab, 0x6e, 0x43, 0x72, 0x7c, 0xc1, 0xa7, 0x10, 0x13, 0xb9, 0x8b, 0x4d,
	0xb7, 0x21, 0x19, 0x50, 0x71, 0x03, 0xbc, 0xd6, 0xcc, 0xd0, 0x94, 0xc5, 0xda, 0x34, 0x17, 0xc1,
	0x5a, 0x8b, 0x48, 0xa1, 0xbf, 0x42, 0x63, 0xaa, 0xdc, 0x26, 0x59, 0x55, 0xf2, 0x67, 0xd0, 0xcb,
	0x95, 0x55, 0x24, 0x3e, 0x9c, 0x0d, 0x65, 0x83, 0x94, 0x1a, 0x42, 0xc1, 0x28, 0x64, 0xf4, 0xde,
	0x1a, 0xd3, 0x96, 0x06, 0xc8, 0xed, 0x5e, 0x83, 0xc4, 0x7b, 0x18, 0xcf, 0x89, 0xe3, 0xfe, 0x1a,
	0xb3, 0x3f, 0x2c, 0x5c, 0x03, 0x73, 0xd4, 0x3f, 0x5c, 0x4e, 0x24, 0x0c, 0xde, 0xe5, 0xb9, 0xcf,
	0xf0, 0x43, 0xb9, 0xf3, 0x4b, 0x9b, 0x8c, 0x65, 0xeb, 0xa5, 0x44, 0x87, 0x5f, 0xc0, 0xf0, 0x83,
	0x5a, 0x2f, 0xf0, 0xd6, 0x3f, 0x32, 0x92, 0xad, 0xf9, 0x53, 0xd9, 0x76, 0x28, 0x3a, 0xfc, 0x15,
	0x24, 0x1f, 0x97, 0xeb, 0xfc, 0xe0, 0xf8, 0x3e, 0xfd, 0x1b, 0x18, 0xbb, 0xe9, 0xfa, 0x9e, 0xe1,
	0x5c, 0xee, 0x5d, 0x3a, 0x93, 0xc7, 0x72, 0xff, 0xa3, 0x8a, 0xce, 0xd7, 0x98, 0x6e, 0xba, 0xd7,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0x98, 0x00, 0x8b, 0x04, 0x05, 0x00, 0x00,
}
