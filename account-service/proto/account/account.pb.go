// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asciiu/gomo/account-service/proto/account/account.proto

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/asciiu/gomo/account-service/proto/account/account.proto

It has these top-level messages:
	NewAccountRequest
	GetAccountKeysRequest
	UpdateAccountRequest
	AccountRequest
	AccountsRequest
	Account
	AccountKey
	UserAccount
	AccountResponse
	UserAccounts
	AccountsResponse
	KeysList
	AccountKeysResponse
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
	AccountType string                   `protobuf:"bytes,6,opt,name=accountType" json:"accountType,omitempty"`
	Title       string                   `protobuf:"bytes,7,opt,name=title" json:"title,omitempty"`
	Balances    []*bal.NewBalanceRequest `protobuf:"bytes,8,rep,name=balances" json:"balances,omitempty"`
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

func (m *NewAccountRequest) GetAccountType() string {
	if m != nil {
		return m.AccountType
	}
	return ""
}

func (m *NewAccountRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NewAccountRequest) GetBalances() []*bal.NewBalanceRequest {
	if m != nil {
		return m.Balances
	}
	return nil
}

type GetAccountKeysRequest struct {
	AccountIDs []string `protobuf:"bytes,1,rep,name=accountIDs" json:"accountIDs,omitempty"`
}

func (m *GetAccountKeysRequest) Reset()                    { *m = GetAccountKeysRequest{} }
func (m *GetAccountKeysRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountKeysRequest) ProtoMessage()               {}
func (*GetAccountKeysRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetAccountKeysRequest) GetAccountIDs() []string {
	if m != nil {
		return m.AccountIDs
	}
	return nil
}

type UpdateAccountRequest struct {
	AccountID   string                   `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	UserID      string                   `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	KeyPublic   string                   `protobuf:"bytes,3,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret   string                   `protobuf:"bytes,4,opt,name=keySecret" json:"keySecret,omitempty"`
	Description string                   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Title       string                   `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Balances    []*bal.NewBalanceRequest `protobuf:"bytes,7,rep,name=balances" json:"balances,omitempty"`
}

func (m *UpdateAccountRequest) Reset()                    { *m = UpdateAccountRequest{} }
func (m *UpdateAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccountRequest) ProtoMessage()               {}
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *UpdateAccountRequest) GetTitle() string {
	if m != nil {
		return m.Title
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
	UserID    string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *AccountRequest) Reset()                    { *m = AccountRequest{} }
func (m *AccountRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()               {}
func (*AccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AccountRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *AccountRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type AccountsRequest struct {
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
}

func (m *AccountsRequest) Reset()                    { *m = AccountsRequest{} }
func (m *AccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountsRequest) ProtoMessage()               {}
func (*AccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AccountsRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

// Responses
type Account struct {
	AccountID   string         `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	AccountType string         `protobuf:"bytes,12,opt,name=accountType" json:"accountType,omitempty"`
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
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Account) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Account) GetAccountType() string {
	if m != nil {
		return m.AccountType
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

type AccountKey struct {
	AccountID   string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	AccountType string `protobuf:"bytes,2,opt,name=accountType" json:"accountType,omitempty"`
	UserID      string `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	Exchange    string `protobuf:"bytes,4,opt,name=exchange" json:"exchange,omitempty"`
	KeyPublic   string `protobuf:"bytes,5,opt,name=keyPublic" json:"keyPublic,omitempty"`
	KeySecret   string `protobuf:"bytes,6,opt,name=keySecret" json:"keySecret,omitempty"`
	Status      string `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
}

func (m *AccountKey) Reset()                    { *m = AccountKey{} }
func (m *AccountKey) String() string            { return proto.CompactTextString(m) }
func (*AccountKey) ProtoMessage()               {}
func (*AccountKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AccountKey) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *AccountKey) GetAccountType() string {
	if m != nil {
		return m.AccountType
	}
	return ""
}

func (m *AccountKey) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AccountKey) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *AccountKey) GetKeyPublic() string {
	if m != nil {
		return m.KeyPublic
	}
	return ""
}

func (m *AccountKey) GetKeySecret() string {
	if m != nil {
		return m.KeySecret
	}
	return ""
}

func (m *AccountKey) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UserAccount struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *UserAccount) Reset()                    { *m = UserAccount{} }
func (m *UserAccount) String() string            { return proto.CompactTextString(m) }
func (*UserAccount) ProtoMessage()               {}
func (*UserAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*AccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*UserAccounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
func (*AccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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

type KeysList struct {
	Keys []*AccountKey `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (m *KeysList) Reset()                    { *m = KeysList{} }
func (m *KeysList) String() string            { return proto.CompactTextString(m) }
func (*KeysList) ProtoMessage()               {}
func (*KeysList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *KeysList) GetKeys() []*AccountKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

type AccountKeysResponse struct {
	Status  string    `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Data    *KeysList `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *AccountKeysResponse) Reset()                    { *m = AccountKeysResponse{} }
func (m *AccountKeysResponse) String() string            { return proto.CompactTextString(m) }
func (*AccountKeysResponse) ProtoMessage()               {}
func (*AccountKeysResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AccountKeysResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AccountKeysResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AccountKeysResponse) GetData() *KeysList {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*NewAccountRequest)(nil), "account.NewAccountRequest")
	proto.RegisterType((*GetAccountKeysRequest)(nil), "account.GetAccountKeysRequest")
	proto.RegisterType((*UpdateAccountRequest)(nil), "account.UpdateAccountRequest")
	proto.RegisterType((*AccountRequest)(nil), "account.AccountRequest")
	proto.RegisterType((*AccountsRequest)(nil), "account.AccountsRequest")
	proto.RegisterType((*Account)(nil), "account.Account")
	proto.RegisterType((*AccountKey)(nil), "account.AccountKey")
	proto.RegisterType((*UserAccount)(nil), "account.UserAccount")
	proto.RegisterType((*AccountResponse)(nil), "account.AccountResponse")
	proto.RegisterType((*UserAccounts)(nil), "account.UserAccounts")
	proto.RegisterType((*AccountsResponse)(nil), "account.AccountsResponse")
	proto.RegisterType((*KeysList)(nil), "account.KeysList")
	proto.RegisterType((*AccountKeysResponse)(nil), "account.AccountKeysResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AccountService service

type AccountServiceClient interface {
	AddAccount(ctx context.Context, in *NewAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	ChangeAvailableBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	ChangeLockedBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	DeleteAccount(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetAccounts(ctx context.Context, in *AccountsRequest, opts ...client.CallOption) (*AccountsResponse, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetAccountKeys(ctx context.Context, in *GetAccountKeysRequest, opts ...client.CallOption) (*AccountKeysResponse, error)
	GetAccountBalance(ctx context.Context, in *bal.BalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	LockBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	UnlockBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error)
	ResyncAccounts(ctx context.Context, in *AccountsRequest, opts ...client.CallOption) (*AccountsResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	ValidateAvailableBalance(ctx context.Context, in *bal.ValidateBalanceRequest, opts ...client.CallOption) (*bal.ValidateBalanceResponse, error)
	ValidateLockedBalance(ctx context.Context, in *bal.ValidateBalanceRequest, opts ...client.CallOption) (*bal.ValidateBalanceResponse, error)
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

func (c *accountServiceClient) ChangeAvailableBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.ChangeAvailableBalance", in)
	out := new(bal.BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ChangeLockedBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.ChangeLockedBalance", in)
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

func (c *accountServiceClient) GetAccountKeys(ctx context.Context, in *GetAccountKeysRequest, opts ...client.CallOption) (*AccountKeysResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.GetAccountKeys", in)
	out := new(AccountKeysResponse)
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

func (c *accountServiceClient) LockBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.LockBalance", in)
	out := new(bal.BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UnlockBalance(ctx context.Context, in *bal.ChangeBalanceRequest, opts ...client.CallOption) (*bal.BalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.UnlockBalance", in)
	out := new(bal.BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ResyncAccounts(ctx context.Context, in *AccountsRequest, opts ...client.CallOption) (*AccountsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.ResyncAccounts", in)
	out := new(AccountsResponse)
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

func (c *accountServiceClient) ValidateAvailableBalance(ctx context.Context, in *bal.ValidateBalanceRequest, opts ...client.CallOption) (*bal.ValidateBalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.ValidateAvailableBalance", in)
	out := new(bal.ValidateBalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ValidateLockedBalance(ctx context.Context, in *bal.ValidateBalanceRequest, opts ...client.CallOption) (*bal.ValidateBalanceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.ValidateLockedBalance", in)
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
	ChangeAvailableBalance(context.Context, *bal.ChangeBalanceRequest, *bal.BalanceResponse) error
	ChangeLockedBalance(context.Context, *bal.ChangeBalanceRequest, *bal.BalanceResponse) error
	DeleteAccount(context.Context, *AccountRequest, *AccountResponse) error
	GetAccounts(context.Context, *AccountsRequest, *AccountsResponse) error
	GetAccount(context.Context, *AccountRequest, *AccountResponse) error
	GetAccountKeys(context.Context, *GetAccountKeysRequest, *AccountKeysResponse) error
	GetAccountBalance(context.Context, *bal.BalanceRequest, *bal.BalanceResponse) error
	LockBalance(context.Context, *bal.ChangeBalanceRequest, *bal.BalanceResponse) error
	UnlockBalance(context.Context, *bal.ChangeBalanceRequest, *bal.BalanceResponse) error
	ResyncAccounts(context.Context, *AccountsRequest, *AccountsResponse) error
	UpdateAccount(context.Context, *UpdateAccountRequest, *AccountResponse) error
	ValidateAvailableBalance(context.Context, *bal.ValidateBalanceRequest, *bal.ValidateBalanceResponse) error
	ValidateLockedBalance(context.Context, *bal.ValidateBalanceRequest, *bal.ValidateBalanceResponse) error
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

func (h *AccountService) ChangeAvailableBalance(ctx context.Context, in *bal.ChangeBalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.ChangeAvailableBalance(ctx, in, out)
}

func (h *AccountService) ChangeLockedBalance(ctx context.Context, in *bal.ChangeBalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.ChangeLockedBalance(ctx, in, out)
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

func (h *AccountService) GetAccountKeys(ctx context.Context, in *GetAccountKeysRequest, out *AccountKeysResponse) error {
	return h.AccountServiceHandler.GetAccountKeys(ctx, in, out)
}

func (h *AccountService) GetAccountBalance(ctx context.Context, in *bal.BalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.GetAccountBalance(ctx, in, out)
}

func (h *AccountService) LockBalance(ctx context.Context, in *bal.ChangeBalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.LockBalance(ctx, in, out)
}

func (h *AccountService) UnlockBalance(ctx context.Context, in *bal.ChangeBalanceRequest, out *bal.BalanceResponse) error {
	return h.AccountServiceHandler.UnlockBalance(ctx, in, out)
}

func (h *AccountService) ResyncAccounts(ctx context.Context, in *AccountsRequest, out *AccountsResponse) error {
	return h.AccountServiceHandler.ResyncAccounts(ctx, in, out)
}

func (h *AccountService) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, out *AccountResponse) error {
	return h.AccountServiceHandler.UpdateAccount(ctx, in, out)
}

func (h *AccountService) ValidateAvailableBalance(ctx context.Context, in *bal.ValidateBalanceRequest, out *bal.ValidateBalanceResponse) error {
	return h.AccountServiceHandler.ValidateAvailableBalance(ctx, in, out)
}

func (h *AccountService) ValidateLockedBalance(ctx context.Context, in *bal.ValidateBalanceRequest, out *bal.ValidateBalanceResponse) error {
	return h.AccountServiceHandler.ValidateLockedBalance(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/asciiu/gomo/account-service/proto/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5f, 0x4f, 0xdb, 0x48,
	0x10, 0xbf, 0xfc, 0x4f, 0xc6, 0xc0, 0xc1, 0x26, 0xe4, 0x4c, 0x2e, 0x87, 0x22, 0x4b, 0xa7, 0x0b,
	0xa7, 0xbb, 0x44, 0x0a, 0x0f, 0xa7, 0x93, 0xee, 0x81, 0x70, 0x51, 0x29, 0x05, 0x51, 0x64, 0xa0,
	0x7d, 0x76, 0x9c, 0x11, 0x58, 0x38, 0x76, 0x9a, 0xb5, 0xa1, 0xf9, 0x18, 0x7d, 0xed, 0xe7, 0xea,
	0xf7, 0xe8, 0x4b, 0x3f, 0x40, 0x95, 0xf5, 0xee, 0xfa, 0x4f, 0x4c, 0x28, 0xa4, 0x7d, 0xb2, 0x76,
	0x66, 0x76, 0xe6, 0x37, 0xbf, 0xf9, 0xb3, 0x86, 0xc3, 0x6b, 0xcb, 0xbb, 0xf1, 0x87, 0x1d, 0xd3,
	0x1d, 0x77, 0x0d, 0x6a, 0x5a, 0x96, 0xdf, 0xbd, 0x76, 0xc7, 0x6e, 0xd7, 0x30, 0x4d, 0xd7, 0x77,
	0xbc, 0xbf, 0x29, 0x4e, 0xef, 0x2c, 0x13, 0xbb, 0x93, 0xa9, 0xeb, 0x49, 0xa9, 0xf8, 0x76, 0x98,
	0x94, 0x94, 0xf8, 0xb1, 0xf1, 0x34, 0x67, 0x43, 0xc3, 0x36, 0x1c, 0x13, 0xc5, 0x37, 0x70, 0xa6,
	0x7d, 0xc8, 0xc2, 0xd6, 0x19, 0xde, 0xf7, 0x83, 0x2b, 0x3a, 0xbe, 0xf3, 0x91, 0x7a, 0xa4, 0x0e,
	0x45, 0x9f, 0xe2, 0xf4, 0x78, 0xa0, 0x66, 0x5a, 0x99, 0x76, 0x45, 0xe7, 0x27, 0xd2, 0x80, 0x32,
	0xbe, 0x37, 0x6f, 0x0c, 0xe7, 0x1a, 0xd5, 0x2c, 0xd3, 0xc8, 0x33, 0x69, 0x42, 0xe5, 0x16, 0x67,
	0xe7, 0xfe, 0xd0, 0xb6, 0x4c, 0x35, 0xc7, 0x94, 0xa1, 0x80, 0x6b, 0x2f, 0xd0, 0x9c, 0xa2, 0xa7,
	0xe6, 0xa5, 0x36, 0x10, 0x90, 0x16, 0x28, 0x23, 0xa4, 0xe6, 0xd4, 0x9a, 0x78, 0x96, 0xeb, 0xa8,
	0x05, 0xa6, 0x8f, 0x8a, 0xe6, 0x16, 0x3c, 0xad, 0xcb, 0xd9, 0x04, 0xd5, 0x62, 0x60, 0x11, 0x11,
	0x91, 0x1a, 0x14, 0x3c, 0xcb, 0xb3, 0x51, 0x2d, 0x31, 0x5d, 0x70, 0x20, 0x3d, 0x28, 0xf3, 0x84,
	0xa9, 0x5a, 0x6e, 0xe5, 0xda, 0x4a, 0xaf, 0xde, 0x19, 0x1a, 0x76, 0xe7, 0x0c, 0xef, 0x0f, 0x03,
	0x39, 0xcf, 0x59, 0x97, 0x76, 0xda, 0x3f, 0xb0, 0x7d, 0x84, 0x1e, 0xa7, 0xe4, 0x04, 0x67, 0x54,
	0xd0, 0xb2, 0x0b, 0xc0, 0x23, 0x1e, 0x0f, 0xa8, 0x9a, 0x69, 0xe5, 0xda, 0x15, 0x3d, 0x22, 0xd1,
	0xbe, 0x64, 0xa0, 0x76, 0x35, 0x19, 0x19, 0x1e, 0x26, 0xf8, 0x6c, 0x42, 0x45, 0x9a, 0x71, 0x4a,
	0x43, 0x41, 0x84, 0xed, 0x6c, 0x8c, 0xed, 0x1f, 0xcb, 0xa8, 0xe4, 0xab, 0xf8, 0x10, 0x5f, 0xa5,
	0x6f, 0xe4, 0xeb, 0x05, 0x6c, 0x7c, 0x8f, 0x7c, 0xb5, 0x3d, 0xf8, 0x99, 0xfb, 0xa1, 0x8f, 0x34,
	0xa2, 0xf6, 0x39, 0x0b, 0x25, 0x6e, 0xfb, 0x48, 0xb0, 0x44, 0xe3, 0xac, 0x2d, 0x36, 0xce, 0x43,
	0xf4, 0x47, 0x9b, 0x3d, 0xb7, 0xac, 0xd9, 0xf3, 0x4b, 0x4b, 0x53, 0x48, 0x96, 0x26, 0x9d, 0xf8,
	0x44, 0xc1, 0x4a, 0x8b, 0x05, 0xab, 0x43, 0x91, 0x7a, 0x86, 0xe7, 0xcf, 0x1b, 0x99, 0xe1, 0x0c,
	0x4e, 0xf3, 0x68, 0xe6, 0x14, 0x0d, 0x0f, 0x47, 0xaf, 0x1d, 0xb5, 0x12, 0x44, 0x93, 0x82, 0xb9,
	0xd6, 0x67, 0x2d, 0x39, 0xd7, 0x42, 0xa0, 0x95, 0x02, 0xd2, 0x8e, 0x94, 0x5b, 0x61, 0xe5, 0x5e,
	0x63, 0xe5, 0x16, 0xb5, 0x0e, 0x8b, 0xfc, 0x29, 0x03, 0x10, 0x8e, 0xc4, 0xd3, 0x48, 0xcf, 0x2e,
	0x23, 0x3d, 0xf7, 0x20, 0xe9, 0xf9, 0x65, 0xa4, 0x17, 0x96, 0x92, 0x5e, 0x4c, 0x92, 0x1e, 0x92,
	0x57, 0x8a, 0x92, 0xa7, 0xfd, 0x0b, 0xca, 0x15, 0xc5, 0xa9, 0xe8, 0xa5, 0x3f, 0x41, 0x6c, 0x57,
	0x96, 0x94, 0xd2, 0xdb, 0xec, 0x88, 0xe5, 0x2b, 0x5a, 0x5c, 0x18, 0x68, 0x63, 0xd9, 0xae, 0x3a,
	0xd2, 0x89, 0xeb, 0x50, 0x8c, 0x44, 0xc9, 0xc4, 0x4a, 0xa4, 0x42, 0x69, 0x8c, 0x94, 0x1a, 0x72,
	0x6d, 0x8a, 0x23, 0x69, 0x43, 0x7e, 0x64, 0x78, 0x06, 0x63, 0x41, 0xe9, 0xd5, 0x64, 0xb4, 0x08,
	0x28, 0x9d, 0x59, 0x68, 0xff, 0xc1, 0x5a, 0x44, 0x48, 0xc9, 0x5f, 0x50, 0xe6, 0xc6, 0xc1, 0x2a,
	0x4a, 0xc3, 0x2a, 0x2d, 0x34, 0x17, 0x36, 0xc3, 0xd9, 0x7a, 0x36, 0xda, 0xbd, 0x18, 0xda, 0xed,
	0x34, 0xb4, 0x94, 0xc3, 0xdd, 0x87, 0xf2, 0x7c, 0x75, 0x9e, 0x5a, 0xd4, 0x23, 0x7f, 0x40, 0xfe,
	0x16, 0x67, 0x02, 0x66, 0x35, 0x09, 0xf3, 0x04, 0x67, 0x3a, 0x33, 0xd0, 0x1c, 0xa8, 0xc6, 0xd6,
	0xee, 0xb3, 0x81, 0xfe, 0x1e, 0x03, 0xba, 0x25, 0x23, 0x0a, 0x48, 0x01, 0xc8, 0xde, 0xc7, 0xb2,
	0x5c, 0x5d, 0x17, 0xc1, 0x63, 0x49, 0x06, 0x00, 0xfd, 0xd1, 0x48, 0xf4, 0x43, 0x43, 0xde, 0x5c,
	0x78, 0x24, 0x1b, 0xea, 0x02, 0xdd, 0x1c, 0xaf, 0xf6, 0x13, 0x39, 0x81, 0xfa, 0xff, 0xac, 0x69,
	0xfb, 0x77, 0x86, 0x65, 0x1b, 0x43, 0x1b, 0xf9, 0x44, 0x91, 0x1d, 0x36, 0x5f, 0x81, 0x32, 0xbe,
	0x51, 0x1b, 0xb5, 0xd8, 0xe8, 0x85, 0xce, 0x5e, 0x42, 0x35, 0xb0, 0x3f, 0x75, 0xcd, 0x5b, 0x1c,
	0xad, 0xe0, 0x69, 0x00, 0xeb, 0x03, 0xb4, 0x51, 0xbe, 0x4f, 0xe4, 0x97, 0xc5, 0x1c, 0x1e, 0x4f,
	0x6e, 0x00, 0x4a, 0xf8, 0x3e, 0x52, 0xb2, 0x60, 0x2a, 0xb6, 0x77, 0x63, 0x27, 0x45, 0x23, 0xbd,
	0xf4, 0x01, 0x42, 0x2f, 0xcf, 0x03, 0x72, 0x0e, 0x1b, 0xf1, 0x87, 0x9a, 0xec, 0x4a, 0xeb, 0xd4,
	0x17, 0xbc, 0xd1, 0x4c, 0xe9, 0xbd, 0x28, 0xa8, 0x03, 0xd8, 0x0a, 0x2f, 0x0a, 0xa2, 0xab, 0x71,
	0x36, 0x97, 0x53, 0x7c, 0x00, 0xca, 0xbc, 0x4c, 0x2b, 0x14, 0xe9, 0x10, 0xd6, 0xaf, 0x1c, 0x7b,
	0x35, 0x1f, 0x47, 0xb0, 0xa1, 0x23, 0x9d, 0x39, 0xe6, 0xaa, 0x55, 0x7a, 0x05, 0xeb, 0xb1, 0x3f,
	0x1a, 0xf2, 0x5b, 0x38, 0xf4, 0x29, 0x7f, 0x3a, 0x4b, 0xcb, 0xf5, 0x16, 0xd4, 0x37, 0x86, 0x6d,
	0xb1, 0x5b, 0xc9, 0xb1, 0xf8, 0x95, 0x25, 0x22, 0xd4, 0x89, 0x2c, 0x9b, 0xe9, 0x4a, 0xe9, 0xf8,
	0x12, 0xb6, 0x85, 0x32, 0x3e, 0x22, 0xab, 0x78, 0x1d, 0x16, 0xd9, 0x1f, 0xf2, 0xfe, 0xd7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x46, 0xed, 0x6d, 0x10, 0xb4, 0x0b, 0x00, 0x00,
}
