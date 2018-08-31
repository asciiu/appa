// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asciiu/gomo/account-service/proto/account/account.proto

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/asciiu/gomo/account-service/proto/account/account.proto

It has these top-level messages:
	NewAccountRequest
	UpdateAccountRequest
	AccountRequest
	AccountsRequest
	Account
	UserAccount
	AccountResponse
	UserAccounts
	AccountsResponse
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bal "github.com/asciiu/gomo/account-service/proto/balance"

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
type NewAccountRequest struct {
	UserID      string                   `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Exchange    string                   `protobuf:"bytes,2,opt,name=exchange" json:"exchange,omitempty"`
	KeyPublic   string                   `protobuf:"bytes,3,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret   string                   `protobuf:"bytes,4,opt,name=keySecret" json:"keySecret,omitempty"`
	Description string                   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Balances    []*bal.NewBalanceRequest `protobuf:"bytes,6,rep,name=balances" json:"balances,omitempty"`
}

func (m *NewAccountRequest) Reset()                    { *m = NewAccountRequest{} }
func (m *NewAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*NewAccountRequest) ProtoMessage()               {}
func (*NewAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewAccountRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *NewAccountRequest) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *NewAccountRequest) GetKeyPublic() string {
	if m != nil {
		return m.KeyPublic
	}
	return ""
}

func (m *NewAccountRequest) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *NewAccountRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NewAccountRequest) GetBalances() []*bal.NewBalanceRequest {
	if m != nil {
		return m.Balances
	}
	return nil
}

type UpdateAccountRequest struct {
	AccountID   string                   `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	UserID      string                   `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	KeyPublic   string                   `protobuf:"bytes,3,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret   string                   `protobuf:"bytes,4,opt,name=keySecret" json:"keySecret,omitempty"`
	Description string                   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Balances    []*bal.NewBalanceRequest `protobuf:"bytes,6,rep,name=balances" json:"balances,omitempty"`
}

func (m *UpdateAccountRequest) Reset()                    { *m = UpdateAccountRequest{} }
func (m *UpdateAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccountRequest) ProtoMessage()               {}
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateAccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *UpdateAccountRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateAccountRequest) GetKeyPublic() string {
	if m != nil {
		return m.KeyPublic
	}
	return ""
}

func (m *UpdateAccountRequest) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *UpdateAccountRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateAccountRequest) GetBalances() []*bal.NewBalanceRequest {
	if m != nil {
		return m.Balances
	}
	return nil
}

type AccountRequest struct {
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
}

func (m *AccountRequest) Reset()                    { *m = AccountRequest{} }
func (m *AccountRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()               {}
func (*AccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type AccountsRequest struct {
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
}

func (m *AccountsRequest) Reset()                    { *m = AccountsRequest{} }
func (m *AccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountsRequest) ProtoMessage()               {}
func (*AccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AccountsRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

// Responses
type Account struct {
	AccountID   string         `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	UserID      string         `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	Exchange    string         `protobuf:"bytes,3,opt,name=exchange" json:"exchange,omitempty"`
	KeyPublic   string         `protobuf:"bytes,4,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret   string         `protobuf:"bytes,5,opt,name=keySecret" json:"keySecret,omitempty"`
	Title       string         `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Description string         `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	Status      string         `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	CreatedOn   string         `protobuf:"bytes,9,opt,name=createdOn" json:"createdOn,omitempty"`
	UpdatedOn   string         `protobuf:"bytes,10,opt,name=updatedOn" json:"updatedOn,omitempty"`
	Balances    []*bal.Balance `protobuf:"bytes,11,rep,name=balances" json:"balances,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Account) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Account) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Account) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Account) GetKeyPublic() string {
	if m != nil {
		return m.KeyPublic
	}
	return ""
}

func (m *Account) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *Account) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Account) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Account) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Account) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *Account) GetUpdatedOn() string {
	if m != nil {
		return m.UpdatedOn
	}
	return ""
}

func (m *Account) GetBalances() []*bal.Balance {
	if m != nil {
		return m.Balances
	}
	return nil
}

type UserAccount struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *UserAccount) Reset()                    { *m = UserAccount{} }
func (m *UserAccount) String() string            { return proto.CompactTextString(m) }
func (*UserAccount) ProtoMessage()               {}
func (*UserAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserAccount) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AccountResponse struct {
	Status  string       `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Data    *UserAccount `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *AccountResponse) Reset()                    { *m = AccountResponse{} }
func (m *AccountResponse) String() string            { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()               {}
func (*AccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AccountResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AccountResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AccountResponse) GetData() *UserAccount {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserAccounts struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *UserAccounts) Reset()                    { *m = UserAccounts{} }
func (m *UserAccounts) String() string            { return proto.CompactTextString(m) }
func (*UserAccounts) ProtoMessage()               {}
func (*UserAccounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserAccounts) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type AccountsResponse struct {
	Status  string        `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Data    *UserAccounts `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *AccountsResponse) Reset()                    { *m = AccountsResponse{} }
func (m *AccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*AccountsResponse) ProtoMessage()               {}
func (*AccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AccountsResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AccountsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AccountsResponse) GetData() *UserAccounts {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*NewAccountRequest)(nil), "account.NewAccountRequest")
	proto.RegisterType((*UpdateAccountRequest)(nil), "account.UpdateAccountRequest")
	proto.RegisterType((*AccountRequest)(nil), "account.AccountRequest")
	proto.RegisterType((*AccountsRequest)(nil), "account.AccountsRequest")
	proto.RegisterType((*Account)(nil), "account.Account")
	proto.RegisterType((*UserAccount)(nil), "account.UserAccount")
	proto.RegisterType((*AccountResponse)(nil), "account.AccountResponse")
	proto.RegisterType((*UserAccounts)(nil), "account.UserAccounts")
	proto.RegisterType((*AccountsResponse)(nil), "account.AccountsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AccountService service

type AccountServiceClient interface {
	AddAccount(ctx context.Context, in *NewAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	AddAccountBalance(ctx context.Context, in *bal.NewBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	DeleteAccount(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetAccounts(ctx context.Context, in *AccountsRequest, opts ...client.CallOption) (*AccountsResponse, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetAccountBalance(ctx context.Context, in *bal.BalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	UpdateAccountBalance(ctx context.Context, in *bal.UpdateBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	ValidateAccountBalance(ctx context.Context, in *bal.ValidateBalanceRequest, opts ...client.CallOption) (*bal.ValidateBalanceResponse, error)
}

type accountServiceClient struct {
	c           client.Client
	serviceName string
}

func NewAccountServiceClient(serviceName string, c client.Client) AccountServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "account"
	}
	return &accountServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountServiceClient) AddAccount(ctx context.Context, in *NewAccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.AddAccount", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) AddAccountBalance(ctx context.Context, in *bal.NewBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.AddAccountBalance", in)
	out := new(bal.BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.DeleteAccount", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccounts(ctx context.Context, in *AccountsRequest, opts ...client.CallOption) (*AccountsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.GetAccounts", in)
	out := new(AccountsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.GetAccount", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountBalance(ctx context.Context, in *bal.BalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.GetAccountBalance", in)
	out := new(bal.BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.UpdateAccount", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccountBalance(ctx context.Context, in *bal.UpdateBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.UpdateAccountBalance", in)
	out := new(bal.BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ValidateAccountBalance(ctx context.Context, in *bal.ValidateBalanceRequest, opts ...client.CallOption) (*bal.ValidateBalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.ValidateAccountBalance", in)
	out := new(bal.ValidateBalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceHandler interface {
	AddAccount(context.Context, *NewAccountRequest, *AccountResponse) error
	AddAccountBalance(context.Context, *bal.NewBalanceRequest, *bal.BalanceResponse) error
	DeleteAccount(context.Context, *AccountRequest, *AccountResponse) error
	GetAccounts(context.Context, *AccountsRequest, *AccountsResponse) error
	GetAccount(context.Context, *AccountRequest, *AccountResponse) error
	GetAccountBalance(context.Context, *bal.BalanceRequest, *bal.BalanceResponse) error
	UpdateAccount(context.Context, *UpdateAccountRequest, *AccountResponse) error
	UpdateAccountBalance(context.Context, *bal.UpdateBalanceRequest, *bal.BalanceResponse) error
	ValidateAccountBalance(context.Context, *bal.ValidateBalanceRequest, *bal.ValidateBalanceResponse) error
}

func RegisterAccountServiceHandler(s server.Server, hdlr AccountServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AccountService{hdlr}, opts...))
}

type AccountService struct {
	AccountServiceHandler
}

func (h *AccountService) AddAccount(ctx context.Context, in *NewAccountRequest, out *AccountResponse) error {
	return h.AccountServiceHandler.AddAccount(ctx, in, out)
}

func (h *AccountService) AddAccountBalance(ctx context.Context, in *bal.NewBalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.AddAccountBalance(ctx, in, out)
}

func (h *AccountService) DeleteAccount(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.AccountServiceHandler.DeleteAccount(ctx, in, out)
}

func (h *AccountService) GetAccounts(ctx context.Context, in *AccountsRequest, out *AccountsResponse) error {
	return h.AccountServiceHandler.GetAccounts(ctx, in, out)
}

func (h *AccountService) GetAccount(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.AccountServiceHandler.GetAccount(ctx, in, out)
}

func (h *AccountService) GetAccountBalance(ctx context.Context, in *bal.BalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.GetAccountBalance(ctx, in, out)
}

func (h *AccountService) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, out *AccountResponse) error {
	return h.AccountServiceHandler.UpdateAccount(ctx, in, out)
}

func (h *AccountService) UpdateAccountBalance(ctx context.Context, in *bal.UpdateBalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.UpdateAccountBalance(ctx, in, out)
}

func (h *AccountService) ValidateAccountBalance(ctx context.Context, in *bal.ValidateBalanceRequest, out *bal.ValidateBalanceResponse) error {
	return h.AccountServiceHandler.ValidateAccountBalance(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/asciiu/gomo/account-service/proto/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x26, 0x4d, 0xf3, 0x37, 0x6e, 0xa1, 0x59, 0x42, 0x70, 0x4d, 0x90, 0x22, 0x9f, 0x52, 0x04,
	0x8e, 0x14, 0x4e, 0x48, 0x1c, 0x48, 0x89, 0x84, 0xca, 0xa1, 0xa0, 0x54, 0xe1, 0xbe, 0xb1, 0x47,
	0xa9, 0x85, 0x63, 0x07, 0xef, 0x9a, 0xc2, 0xab, 0xf0, 0x4e, 0x3c, 0x00, 0x17, 0x9e, 0x05, 0x65,
	0xd7, 0xbb, 0xb6, 0x13, 0x93, 0xd2, 0x9c, 0x38, 0x45, 0x33, 0xdf, 0x78, 0xe6, 0x9b, 0xf9, 0x76,
	0x26, 0x70, 0xbe, 0xf0, 0xf9, 0x75, 0x32, 0x77, 0xdc, 0x68, 0x39, 0xa4, 0xcc, 0xf5, 0xfd, 0x64,
	0xb8, 0x88, 0x96, 0xd1, 0x90, 0xba, 0x6e, 0x94, 0x84, 0xfc, 0x05, 0xc3, 0xf8, 0xab, 0xef, 0xe2,
	0x70, 0x15, 0x47, 0x5c, 0x7b, 0xd5, 0xaf, 0x23, 0xbc, 0xa4, 0x91, 0x9a, 0xd6, 0xdd, 0x92, 0xcd,
	0x69, 0x40, 0x43, 0x17, 0xd5, 0xaf, 0x4c, 0x66, 0xff, 0xaa, 0x40, 0xfb, 0x12, 0x6f, 0xc6, 0xf2,
	0x93, 0x29, 0x7e, 0x49, 0x90, 0x71, 0xd2, 0x85, 0x7a, 0xc2, 0x30, 0xbe, 0x98, 0x98, 0x95, 0x7e,
	0x65, 0xd0, 0x9a, 0xa6, 0x16, 0xb1, 0xa0, 0x89, 0xdf, 0xdc, 0x6b, 0x1a, 0x2e, 0xd0, 0x3c, 0x10,
	0x88, 0xb6, 0x49, 0x0f, 0x5a, 0x9f, 0xf1, 0xfb, 0xc7, 0x64, 0x1e, 0xf8, 0xae, 0x59, 0x15, 0x60,
	0xe6, 0x48, 0xd1, 0x2b, 0x74, 0x63, 0xe4, 0xe6, 0xa1, 0x46, 0xa5, 0x83, 0xf4, 0xc1, 0xf0, 0x90,
	0xb9, 0xb1, 0xbf, 0xe2, 0x7e, 0x14, 0x9a, 0x35, 0x81, 0xe7, 0x5d, 0x64, 0x04, 0xcd, 0x94, 0x38,
	0x33, 0xeb, 0xfd, 0xea, 0xc0, 0x18, 0x75, 0x9d, 0x39, 0x0d, 0x9c, 0x4b, 0xbc, 0x39, 0x97, 0xfe,
	0x94, 0xfb, 0x54, 0xc7, 0xd9, 0xbf, 0x2b, 0xd0, 0x99, 0xad, 0x3c, 0xca, 0x71, 0xa3, 0xbd, 0x1e,
	0xb4, 0xd2, 0x19, 0xe9, 0x0e, 0x33, 0x47, 0xae, 0xf9, 0x83, 0x42, 0xf3, 0xff, 0x5f, 0x83, 0x0e,
	0xdc, 0xbf, 0x4b, 0x67, 0xf6, 0x19, 0x3c, 0x48, 0xe3, 0xd9, 0x2d, 0x4a, 0xdb, 0x3f, 0x0f, 0xa0,
	0x91, 0xc6, 0xee, 0x39, 0xae, 0xfc, 0x5b, 0xa9, 0xee, 0x7a, 0x2b, 0x87, 0x3b, 0x47, 0x59, 0xdb,
	0x1c, 0x65, 0x07, 0x6a, 0xdc, 0xe7, 0x01, 0x9a, 0x75, 0x81, 0x48, 0x63, 0x73, 0xc0, 0x8d, 0xed,
	0x01, 0x77, 0xa1, 0xce, 0x38, 0xe5, 0x09, 0x33, 0x9b, 0x92, 0xa7, 0xb4, 0xd6, 0xd5, 0xdc, 0x18,
	0x29, 0x47, 0xef, 0x43, 0x68, 0xb6, 0x64, 0x35, 0xed, 0x58, 0xa3, 0x89, 0x78, 0x42, 0x6b, 0x14,
	0x24, 0xaa, 0x1d, 0x64, 0x90, 0x13, 0xcd, 0x10, 0xa2, 0x1d, 0x09, 0xd1, 0x94, 0x62, 0x99, 0x54,
	0xaf, 0xc0, 0x98, 0x31, 0x8c, 0xd5, 0x48, 0x9f, 0x81, 0xda, 0x62, 0x31, 0x50, 0x63, 0x74, 0xe2,
	0xa8, 0x25, 0x57, 0x8a, 0xaa, 0x00, 0x7b, 0xa9, 0x55, 0x9b, 0x22, 0x5b, 0x45, 0x21, 0xc3, 0x5c,
	0x2f, 0x95, 0x42, 0x2f, 0x26, 0x34, 0x96, 0xc8, 0x18, 0xd5, 0xeb, 0xa9, 0x4c, 0x32, 0x80, 0x43,
	0x8f, 0x72, 0x2a, 0x94, 0x30, 0x46, 0x1d, 0x5d, 0x2d, 0x47, 0x6a, 0x2a, 0x22, 0xec, 0xd7, 0x70,
	0x94, 0x73, 0x32, 0xf2, 0x1c, 0x9a, 0x69, 0xf0, 0xba, 0x5a, 0xb5, 0x94, 0xab, 0x8e, 0xb0, 0x23,
	0x38, 0xc9, 0x9e, 0xd8, 0xde, 0x6c, 0xcf, 0x0a, 0x6c, 0x1f, 0x95, 0xb1, 0x65, 0x92, 0xee, 0xe8,
	0x47, 0x4d, 0x2f, 0xc1, 0x95, 0xbc, 0x77, 0x64, 0x02, 0x30, 0xf6, 0x3c, 0x35, 0x6a, 0x4b, 0x7f,
	0xbd, 0x75, 0xe7, 0x2c, 0x73, 0xab, 0x93, 0x94, 0xb3, 0x7d, 0x8f, 0xbc, 0x85, 0x76, 0x96, 0x25,
	0x15, 0x94, 0xfc, 0x65, 0x27, 0xad, 0x4e, 0x41, 0xf6, 0x2c, 0xc9, 0x04, 0x8e, 0x27, 0x18, 0xa0,
	0xbe, 0x40, 0xe4, 0xf1, 0x76, 0xc5, 0xdb, 0xa9, 0x4c, 0xc0, 0x78, 0x87, 0x5c, 0x2b, 0xb2, 0x15,
	0xaa, 0xb6, 0xd9, 0x3a, 0x2d, 0x41, 0x74, 0x96, 0x31, 0x40, 0x96, 0x65, 0x3f, 0x22, 0x6f, 0xa0,
	0x9d, 0xa5, 0x50, 0x33, 0x79, 0x58, 0xec, 0x7d, 0xf7, 0x40, 0xde, 0xc3, 0x71, 0xe1, 0x24, 0x93,
	0xa7, 0x99, 0xb8, 0x25, 0xa7, 0x7a, 0x27, 0x9b, 0x8b, 0x8d, 0xf3, 0xae, 0x08, 0x9d, 0x8a, 0xda,
	0x12, 0xfa, 0x47, 0x5a, 0x33, 0xe8, 0x7e, 0xa2, 0x81, 0x5f, 0x92, 0xec, 0x89, 0xf8, 0x42, 0x81,
	0x1b, 0xe9, 0x7a, 0xe5, 0xa0, 0x4a, 0x3b, 0xaf, 0x8b, 0x3f, 0xd9, 0x97, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x16, 0xa3, 0x94, 0xe1, 0xf7, 0x07, 0x00, 0x00,
}
