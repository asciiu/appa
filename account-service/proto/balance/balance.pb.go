// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asciiu/gomo/account-service/proto/balance/balance.proto

/*
Package balance is a generated protocol buffer package.

It is generated from these files:
	github.com/asciiu/gomo/account-service/proto/balance/balance.proto

It has these top-level messages:
	NewBalanceRequest
	Balance
*/
package balance

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
type NewBalanceRequest struct {
	CurrencySymbol  string  `protobuf:"bytes,1,opt,name=currencySymbol" json:"currencySymbol"`
	CurrencyBalance float64 `protobuf:"fixed64,2,opt,name=currencyBalance" json:"currencyBalance"`
}

func (m *NewBalanceRequest) Reset()                    { *m = NewBalanceRequest{} }
func (m *NewBalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*NewBalanceRequest) ProtoMessage()               {}
func (*NewBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewBalanceRequest) GetCurrencySymbol() string {
	if m != nil {
		return m.CurrencySymbol
	}
	return ""
}

func (m *NewBalanceRequest) GetCurrencyBalance() float64 {
	if m != nil {
		return m.CurrencyBalance
	}
	return 0
}

// Responses
type Balance struct {
	BalanceID         string  `protobuf:"bytes,1,opt,name=balanceID" json:"balanceID"`
	UserID            string  `protobuf:"bytes,2,opt,name=userID" json:"userID"`
	AccountID         string  `protobuf:"bytes,3,opt,name=accountID" json:"accountID"`
	CurrencySymbol    string  `protobuf:"bytes,4,opt,name=currencySymbol" json:"currencySymbol"`
	Available         float64 `protobuf:"fixed64,5,opt,name=available" json:"available"`
	Locked            float64 `protobuf:"fixed64,6,opt,name=locked" json:"locked"`
	ExchangeTotal     float64 `protobuf:"fixed64,7,opt,name=exchangeTotal" json:"exchangeTotal"`
	ExchangeAvailable float64 `protobuf:"fixed64,8,opt,name=exchangeAvailable" json:"exchangeAvailable"`
	ExchangeLocked    float64 `protobuf:"fixed64,9,opt,name=exchangeLocked" json:"exchangeLocked"`
	CreatedOn         string  `protobuf:"bytes,10,opt,name=created_on,json=createdOn" json:"created_on"`
	UpdatedOn         string  `protobuf:"bytes,11,opt,name=updated_on,json=updatedOn" json:"updated_on"`
}

func (m *Balance) Reset()                    { *m = Balance{} }
func (m *Balance) String() string            { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()               {}
func (*Balance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Balance) GetBalanceID() string {
	if m != nil {
		return m.BalanceID
	}
	return ""
}

func (m *Balance) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Balance) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Balance) GetCurrencySymbol() string {
	if m != nil {
		return m.CurrencySymbol
	}
	return ""
}

func (m *Balance) GetAvailable() float64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *Balance) GetLocked() float64 {
	if m != nil {
		return m.Locked
	}
	return 0
}

func (m *Balance) GetExchangeTotal() float64 {
	if m != nil {
		return m.ExchangeTotal
	}
	return 0
}

func (m *Balance) GetExchangeAvailable() float64 {
	if m != nil {
		return m.ExchangeAvailable
	}
	return 0
}

func (m *Balance) GetExchangeLocked() float64 {
	if m != nil {
		return m.ExchangeLocked
	}
	return 0
}

func (m *Balance) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *Balance) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func init() {
	proto.RegisterType((*NewBalanceRequest)(nil), "balance.NewBalanceRequest")
	proto.RegisterType((*Balance)(nil), "balance.Balance")
}

func init() {
	proto.RegisterFile("github.com/asciiu/gomo/account-service/proto/balance/balance.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0xd4, 0xcd, 0x3e, 0x51, 0x59, 0x0e, 0x92, 0x83, 0xc2, 0x18, 0x22, 0x3b, 0xe8,
	0x7a, 0xf0, 0x2f, 0x70, 0xf4, 0x52, 0x10, 0x07, 0xd5, 0xbb, 0xa4, 0xe9, 0xa3, 0x2b, 0xa6, 0xc9,
	0x4c, 0x93, 0xe9, 0xfe, 0x70, 0xef, 0xd2, 0x34, 0xdd, 0xb0, 0xee, 0xd4, 0xbe, 0xdf, 0x97, 0x7c,
	0xdf, 0xf7, 0x08, 0x2c, 0x8a, 0xd2, 0xac, 0x6c, 0x36, 0xe7, 0xaa, 0x8a, 0x58, 0xcd, 0xcb, 0xd2,
	0x46, 0x85, 0xaa, 0x54, 0xc4, 0x38, 0x57, 0x56, 0x9a, 0x87, 0x1a, 0xf5, 0xa6, 0xe4, 0x18, 0xad,
	0xb5, 0x32, 0x2a, 0xca, 0x98, 0x60, 0x92, 0x63, 0xf7, 0x9d, 0x3b, 0x4a, 0x46, 0x7e, 0x9c, 0x22,
	0x8c, 0x5f, 0xf0, 0x6b, 0xd1, 0x4e, 0x29, 0x7e, 0x5a, 0xac, 0x0d, 0xb9, 0x83, 0x0b, 0x6e, 0xb5,
	0x46, 0xc9, 0xb7, 0xaf, 0xdb, 0x2a, 0x53, 0x82, 0x06, 0x93, 0x60, 0x16, 0xa6, 0x3d, 0x4a, 0x66,
	0x70, 0xd9, 0x11, 0xef, 0x40, 0x07, 0x93, 0x60, 0x16, 0xa4, 0x7d, 0x3c, 0xfd, 0x19, 0xc0, 0xc8,
	0xff, 0x93, 0x6b, 0x08, 0x7d, 0x7a, 0x12, 0x7b, 0xe3, 0x3d, 0x20, 0x57, 0x30, 0xb4, 0x35, 0xea,
	0x24, 0x76, 0x56, 0x61, 0xea, 0xa7, 0xe6, 0x96, 0x5f, 0x30, 0x89, 0xe9, 0x51, 0x7b, 0x6b, 0x07,
	0x0e, 0x34, 0x3e, 0x3e, 0xd8, 0xb8, 0x71, 0xd9, 0xb0, 0x52, 0xb0, 0x4c, 0x20, 0x3d, 0x71, 0x5d,
	0xf7, 0xa0, 0xc9, 0x16, 0x8a, 0x7f, 0x60, 0x4e, 0x87, 0x4e, 0xf2, 0x13, 0xb9, 0x85, 0x73, 0xfc,
	0xe6, 0x2b, 0x26, 0x0b, 0x7c, 0x53, 0x86, 0x09, 0x3a, 0x72, 0xf2, 0x5f, 0x48, 0xee, 0x61, 0xdc,
	0x81, 0xa7, 0x5d, 0xc6, 0xa9, 0x3b, 0xf9, 0x5f, 0x68, 0x1a, 0x77, 0xf0, 0xb9, 0xcd, 0x0c, 0xdd,
	0xd1, 0x1e, 0x25, 0x37, 0x00, 0x5c, 0x23, 0x33, 0x98, 0xbf, 0x2b, 0x49, 0xa1, 0x5d, 0xdc, 0x93,
	0xa5, 0x6c, 0x64, 0xbb, 0xce, 0x3b, 0xf9, 0xac, 0x95, 0x3d, 0x59, 0xca, 0x6c, 0xe8, 0x9e, 0xfb,
	0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xcf, 0x23, 0xa5, 0x34, 0x02, 0x00, 0x00,
}