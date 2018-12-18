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

type UserOrdersRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
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

type NewOrderRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	MarketName           string   `protobuf:"bytes,2,opt,name=marketName,proto3" json:"marketName,omitempty"`
	Side                 string   `protobuf:"bytes,3,opt,name=side,proto3" json:"side,omitempty"`
	Size                 float64  `protobuf:"fixed64,4,opt,name=size,proto3" json:"size,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
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
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Status               string   `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	CreatedOn            string   `protobuf:"bytes,8,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	UpdatedOn            string   `protobuf:"bytes,9,opt,name=updatedOn,proto3" json:"updatedOn,omitempty"`
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

type OrderList struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderList) Reset()         { *m = OrderList{} }
func (m *OrderList) String() string { return proto.CompactTextString(m) }
func (*OrderList) ProtoMessage()    {}
func (*OrderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{5}
}

func (m *OrderList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderList.Unmarshal(m, b)
}
func (m *OrderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderList.Marshal(b, m, deterministic)
}
func (m *OrderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderList.Merge(m, src)
}
func (m *OrderList) XXX_Size() int {
	return xxx_messageInfo_OrderList.Size(m)
}
func (m *OrderList) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderList.DiscardUnknown(m)
}

var xxx_messageInfo_OrderList proto.InternalMessageInfo

func (m *OrderList) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type OrdersResponse struct {
	Status               string     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *OrderList `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OrdersResponse) Reset()         { *m = OrdersResponse{} }
func (m *OrdersResponse) String() string { return proto.CompactTextString(m) }
func (*OrdersResponse) ProtoMessage()    {}
func (*OrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{6}
}

func (m *OrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdersResponse.Unmarshal(m, b)
}
func (m *OrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdersResponse.Marshal(b, m, deterministic)
}
func (m *OrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersResponse.Merge(m, src)
}
func (m *OrdersResponse) XXX_Size() int {
	return xxx_messageInfo_OrdersResponse.Size(m)
}
func (m *OrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersResponse proto.InternalMessageInfo

func (m *OrdersResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrdersResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OrdersResponse) GetData() *OrderList {
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
	proto.RegisterType((*OrderList)(nil), "OrderList")
	proto.RegisterType((*OrdersResponse)(nil), "OrdersResponse")
	proto.RegisterType((*OrderResponse)(nil), "OrderResponse")
	proto.RegisterType((*StatusResponse)(nil), "StatusResponse")
}

func init() { proto.RegisterFile("proto/order/order.proto", fileDescriptor_986e030a471601a2) }

var fileDescriptor_986e030a471601a2 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xad, 0xb7, 0x49, 0xba, 0x99, 0xb2, 0x29, 0xf8, 0x00, 0xd6, 0x6a, 0x55, 0x55, 0x3e, 0x15,
	0x2d, 0xf2, 0x4a, 0xe5, 0xc0, 0x19, 0x58, 0x21, 0x21, 0xa1, 0x5d, 0x29, 0x2b, 0x3e, 0xc0, 0x5b,
	0x8f, 0x50, 0x04, 0x4d, 0x42, 0xec, 0x80, 0xe0, 0xce, 0x2f, 0xf2, 0x0f, 0xfc, 0x05, 0xca, 0x38,
	0x69, 0x13, 0xda, 0x72, 0x40, 0x5c, 0x2a, 0xcf, 0x7b, 0xaf, 0x7e, 0x6f, 0x3c, 0x76, 0xe0, 0x49,
	0x59, 0x15, 0xae, 0xb8, 0x2a, 0x2a, 0x83, 0x95, 0xff, 0x55, 0x84, 0xc8, 0x25, 0x3c, 0xb8, 0x6d,
	0xca, 0x14, 0x3f, 0xd7, 0x68, 0x1d, 0x17, 0x30, 0x21, 0xfa, 0xed, 0xb5, 0x60, 0x0b, 0xb6, 0x8c,
	0xd3, 0xae, 0x94, 0x97, 0xf0, 0xe8, 0xbd, 0xc5, 0x8a, 0xd4, 0xb6, 0x93, 0x3f, 0x86, 0xa8, 0xb6,
	0x3d, 0x75, 0x5b, 0xc9, 0x1f, 0x0c, 0x66, 0x37, 0xf8, 0x75, 0xb0, 0xf5, 0x11, 0x2d, 0x9f, 0x03,
	0x6c, 0x74, 0xf5, 0x11, 0xdd, 0x8d, 0xde, 0xa0, 0x38, 0x21, 0xae, 0x87, 0x70, 0x0e, 0x81, 0xcd,
	0x0c, 0x8a, 0x31, 0x31, 0xb4, 0xf6, 0xd8, 0x77, 0x14, 0xc1, 0x82, 0x2d, 0x59, 0x4a, 0xeb, 0x06,
	0x73, 0xdf, 0x4a, 0x14, 0xa1, 0xd7, 0x35, 0x6b, 0xf9, 0x8b, 0x41, 0x48, 0x21, 0x8e, 0x37, 0xd6,
	0xcb, 0x75, 0xf2, 0x97, 0x5c, 0xe3, 0xa3, 0xb9, 0x82, 0x03, 0xb9, 0xc2, 0x03, 0xb9, 0xa2, 0x5d,
	0xae, 0xc6, 0xd3, 0x3a, 0xed, 0x6a, 0x2b, 0x26, 0xde, 0xd3, 0x57, 0xfc, 0x02, 0xe2, 0x75, 0x85,
	0xda, 0xa1, 0xb9, 0xcd, 0xc5, 0x29, 0x51, 0x3b, 0xa0, 0x61, 0xeb, 0xd2, 0xb4, 0x6c, 0xec, 0xd9,
	0x2d, 0x20, 0x9f, 0x42, 0x4c, 0xad, 0x5e, 0x6b, 0xa7, 0xf9, 0x05, 0x84, 0xd4, 0x1f, 0x35, 0x3b,
	0x5d, 0x45, 0xca, 0x8f, 0xc2, 0x83, 0xf2, 0xb2, 0x95, 0xbe, 0xcb, 0xac, 0xe3, 0x73, 0x88, 0x08,
	0xb5, 0x82, 0x2d, 0xc6, 0x3d, 0x6d, 0x8b, 0xca, 0x7b, 0x48, 0xba, 0xa1, 0xdb, 0xb2, 0xc8, 0x6d,
	0x3f, 0x3d, 0x1b, 0xa4, 0x17, 0x30, 0xd9, 0xa0, 0xb5, 0xfa, 0x43, 0x37, 0xc6, 0xae, 0xe4, 0x73,
	0x08, 0x8c, 0x76, 0x9a, 0x4e, 0x71, 0xba, 0x02, 0xb5, 0x75, 0x4f, 0x09, 0x97, 0x1a, 0xce, 0xda,
	0xbb, 0xf2, 0x7f, 0x2d, 0x9a, 0xb3, 0x68, 0x2d, 0x5e, 0x41, 0x72, 0x47, 0x7b, 0xfc, 0xbb, 0xc7,
	0xea, 0x27, 0x6b, 0x9f, 0xcb, 0x1d, 0x56, 0x5f, 0xb2, 0x35, 0x72, 0x05, 0xa7, 0x2f, 0x8d, 0xf1,
	0x37, 0xec, 0xa1, 0xfa, 0xe3, 0xc6, 0x9f, 0x27, 0x6a, 0xd0, 0x94, 0x1c, 0xf1, 0x2b, 0x98, 0xbe,
	0xd6, 0xf9, 0x1a, 0x3f, 0xf9, 0xbf, 0x9c, 0xa9, 0x81, 0x7e, 0xa6, 0x86, 0x09, 0xe5, 0x88, 0x3f,
	0x83, 0xf8, 0x4d, 0x96, 0x9b, 0x83, 0xf2, 0xfd, 0xed, 0x5f, 0x40, 0xd2, 0xa8, 0x77, 0xef, 0x94,
	0x73, 0xb5, 0xf7, 0x68, 0xcf, 0x67, 0x6a, 0x38, 0x4f, 0x39, 0xba, 0x8f, 0xe8, 0x6b, 0xf0, 0xfc,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x1f, 0xad, 0x3c, 0x28, 0x04, 0x00, 0x00,
}
