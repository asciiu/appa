// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/activity.proto

/*
Package activity is a generated protocol buffer package.

It is generated from these files:
	proto/activity.proto

It has these top-level messages:
	ActivityRequest
	RecentActivityRequest
	ActivityCountRequest
	UpdateActivityRequest
	Activity
	UserActivityPage
	ActivityPagedResponse
	ActivityData
	ActivityResponse
	ActivityList
	ActivityListResponse
	ActivityCount
	ActivityCountResponse
*/
package activity

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
type ActivityRequest struct {
	UserID   string `protobuf:"bytes,1,opt,name=userID" json:"userID"`
	ObjectID string `protobuf:"bytes,2,opt,name=objectID" json:"objectID"`
	Page     uint32 `protobuf:"varint,3,opt,name=page" json:"page"`
	PageSize uint32 `protobuf:"varint,4,opt,name=pageSize" json:"pageSize"`
}

func (m *ActivityRequest) Reset()                    { *m = ActivityRequest{} }
func (m *ActivityRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivityRequest) ProtoMessage()               {}
func (*ActivityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActivityRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ActivityRequest) GetObjectID() string {
	if m != nil {
		return m.ObjectID
	}
	return ""
}

func (m *ActivityRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ActivityRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type RecentActivityRequest struct {
	ObjectID string `protobuf:"bytes,1,opt,name=objectID" json:"objectID"`
	Count    uint32 `protobuf:"varint,2,opt,name=count" json:"count"`
}

func (m *RecentActivityRequest) Reset()                    { *m = RecentActivityRequest{} }
func (m *RecentActivityRequest) String() string            { return proto.CompactTextString(m) }
func (*RecentActivityRequest) ProtoMessage()               {}
func (*RecentActivityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RecentActivityRequest) GetObjectID() string {
	if m != nil {
		return m.ObjectID
	}
	return ""
}

func (m *RecentActivityRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ActivityCountRequest struct {
	ObjectID string `protobuf:"bytes,1,opt,name=objectID" json:"objectID"`
}

func (m *ActivityCountRequest) Reset()                    { *m = ActivityCountRequest{} }
func (m *ActivityCountRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivityCountRequest) ProtoMessage()               {}
func (*ActivityCountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ActivityCountRequest) GetObjectID() string {
	if m != nil {
		return m.ObjectID
	}
	return ""
}

type UpdateActivityRequest struct {
	ActivityID string `protobuf:"bytes,1,opt,name=activityID" json:"activityID"`
	SeenAt     string `protobuf:"bytes,2,opt,name=seenAt" json:"seenAt"`
	ClickedAt  string `protobuf:"bytes,3,opt,name=clickedAt" json:"clickedAt"`
}

func (m *UpdateActivityRequest) Reset()                    { *m = UpdateActivityRequest{} }
func (m *UpdateActivityRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateActivityRequest) ProtoMessage()               {}
func (*UpdateActivityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateActivityRequest) GetActivityID() string {
	if m != nil {
		return m.ActivityID
	}
	return ""
}

func (m *UpdateActivityRequest) GetSeenAt() string {
	if m != nil {
		return m.SeenAt
	}
	return ""
}

func (m *UpdateActivityRequest) GetClickedAt() string {
	if m != nil {
		return m.ClickedAt
	}
	return ""
}

// Responses
type Activity struct {
	ActivityID  string `protobuf:"bytes,1,opt,name=activityID" json:"activityID"`
	UserID      string `protobuf:"bytes,3,opt,name=userID" json:"userID"`
	Type        string `protobuf:"bytes,2,opt,name=type" json:"type"`
	ObjectID    string `protobuf:"bytes,4,opt,name=objectID" json:"objectID"`
	Title       string `protobuf:"bytes,5,opt,name=title" json:"title"`
	Subtitle    string `protobuf:"bytes,6,opt,name=subtitle" json:"subtitle"`
	Description string `protobuf:"bytes,7,opt,name=description" json:"description"`
	Timestamp   string `protobuf:"bytes,8,opt,name=timestamp" json:"timestamp"`
	ClickedAt   string `protobuf:"bytes,9,opt,name=clickedAt" json:"clickedAt"`
	SeenAt      string `protobuf:"bytes,10,opt,name=seenAt" json:"seenAt"`
}

func (m *Activity) Reset()                    { *m = Activity{} }
func (m *Activity) String() string            { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()               {}
func (*Activity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Activity) GetActivityID() string {
	if m != nil {
		return m.ActivityID
	}
	return ""
}

func (m *Activity) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Activity) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Activity) GetObjectID() string {
	if m != nil {
		return m.ObjectID
	}
	return ""
}

func (m *Activity) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Activity) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *Activity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Activity) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Activity) GetClickedAt() string {
	if m != nil {
		return m.ClickedAt
	}
	return ""
}

func (m *Activity) GetSeenAt() string {
	if m != nil {
		return m.SeenAt
	}
	return ""
}

type UserActivityPage struct {
	Page     uint32      `protobuf:"varint,1,opt,name=page" json:"page"`
	PageSize uint32      `protobuf:"varint,2,opt,name=pageSize" json:"pageSize"`
	Total    uint32      `protobuf:"varint,3,opt,name=total" json:"total"`
	Activity []*Activity `protobuf:"bytes,4,rep,name=activity" json:"activity"`
}

func (m *UserActivityPage) Reset()                    { *m = UserActivityPage{} }
func (m *UserActivityPage) String() string            { return proto.CompactTextString(m) }
func (*UserActivityPage) ProtoMessage()               {}
func (*UserActivityPage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserActivityPage) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *UserActivityPage) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *UserActivityPage) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *UserActivityPage) GetActivity() []*Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

type ActivityPagedResponse struct {
	Status  string            `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string            `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *UserActivityPage `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *ActivityPagedResponse) Reset()                    { *m = ActivityPagedResponse{} }
func (m *ActivityPagedResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivityPagedResponse) ProtoMessage()               {}
func (*ActivityPagedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ActivityPagedResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ActivityPagedResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ActivityPagedResponse) GetData() *UserActivityPage {
	if m != nil {
		return m.Data
	}
	return nil
}

type ActivityData struct {
	Activity *Activity `protobuf:"bytes,1,opt,name=activity" json:"activity"`
}

func (m *ActivityData) Reset()                    { *m = ActivityData{} }
func (m *ActivityData) String() string            { return proto.CompactTextString(m) }
func (*ActivityData) ProtoMessage()               {}
func (*ActivityData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ActivityData) GetActivity() *Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

type ActivityResponse struct {
	Status  string        `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string        `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *ActivityData `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *ActivityResponse) Reset()                    { *m = ActivityResponse{} }
func (m *ActivityResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivityResponse) ProtoMessage()               {}
func (*ActivityResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ActivityResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ActivityResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ActivityResponse) GetData() *ActivityData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ActivityList struct {
	Activity []*Activity `protobuf:"bytes,1,rep,name=activity" json:"activity"`
}

func (m *ActivityList) Reset()                    { *m = ActivityList{} }
func (m *ActivityList) String() string            { return proto.CompactTextString(m) }
func (*ActivityList) ProtoMessage()               {}
func (*ActivityList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ActivityList) GetActivity() []*Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

type ActivityListResponse struct {
	Status  string        `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string        `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *ActivityList `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *ActivityListResponse) Reset()                    { *m = ActivityListResponse{} }
func (m *ActivityListResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivityListResponse) ProtoMessage()               {}
func (*ActivityListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ActivityListResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ActivityListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ActivityListResponse) GetData() *ActivityList {
	if m != nil {
		return m.Data
	}
	return nil
}

type ActivityCount struct {
	Count uint32 `protobuf:"varint,1,opt,name=count" json:"count"`
}

func (m *ActivityCount) Reset()                    { *m = ActivityCount{} }
func (m *ActivityCount) String() string            { return proto.CompactTextString(m) }
func (*ActivityCount) ProtoMessage()               {}
func (*ActivityCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ActivityCount) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ActivityCountResponse struct {
	Status  string         `protobuf:"bytes,1,opt,name=status" json:"status"`
	Message string         `protobuf:"bytes,2,opt,name=message" json:"message"`
	Data    *ActivityCount `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *ActivityCountResponse) Reset()                    { *m = ActivityCountResponse{} }
func (m *ActivityCountResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivityCountResponse) ProtoMessage()               {}
func (*ActivityCountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ActivityCountResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ActivityCountResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ActivityCountResponse) GetData() *ActivityCount {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivityRequest)(nil), "activity.ActivityRequest")
	proto.RegisterType((*RecentActivityRequest)(nil), "activity.RecentActivityRequest")
	proto.RegisterType((*ActivityCountRequest)(nil), "activity.ActivityCountRequest")
	proto.RegisterType((*UpdateActivityRequest)(nil), "activity.UpdateActivityRequest")
	proto.RegisterType((*Activity)(nil), "activity.Activity")
	proto.RegisterType((*UserActivityPage)(nil), "activity.UserActivityPage")
	proto.RegisterType((*ActivityPagedResponse)(nil), "activity.ActivityPagedResponse")
	proto.RegisterType((*ActivityData)(nil), "activity.ActivityData")
	proto.RegisterType((*ActivityResponse)(nil), "activity.ActivityResponse")
	proto.RegisterType((*ActivityList)(nil), "activity.ActivityList")
	proto.RegisterType((*ActivityListResponse)(nil), "activity.ActivityListResponse")
	proto.RegisterType((*ActivityCount)(nil), "activity.ActivityCount")
	proto.RegisterType((*ActivityCountResponse)(nil), "activity.ActivityCountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ActivityBulletin service

type ActivityBulletinClient interface {
	FindUserActivity(ctx context.Context, in *ActivityRequest, opts ...client.CallOption) (*ActivityPagedResponse, error)
	FindMostRecentActivity(ctx context.Context, in *RecentActivityRequest, opts ...client.CallOption) (*ActivityListResponse, error)
	FindActivityCount(ctx context.Context, in *ActivityCountRequest, opts ...client.CallOption) (*ActivityCountResponse, error)
	UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...client.CallOption) (*ActivityResponse, error)
}

type activityBulletinClient struct {
	c           client.Client
	serviceName string
}

func NewActivityBulletinClient(serviceName string, c client.Client) ActivityBulletinClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "activity"
	}
	return &activityBulletinClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *activityBulletinClient) FindUserActivity(ctx context.Context, in *ActivityRequest, opts ...client.CallOption) (*ActivityPagedResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ActivityBulletin.FindUserActivity", in)
	out := new(ActivityPagedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityBulletinClient) FindMostRecentActivity(ctx context.Context, in *RecentActivityRequest, opts ...client.CallOption) (*ActivityListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ActivityBulletin.FindMostRecentActivity", in)
	out := new(ActivityListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityBulletinClient) FindActivityCount(ctx context.Context, in *ActivityCountRequest, opts ...client.CallOption) (*ActivityCountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ActivityBulletin.FindActivityCount", in)
	out := new(ActivityCountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityBulletinClient) UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...client.CallOption) (*ActivityResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ActivityBulletin.UpdateActivity", in)
	out := new(ActivityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ActivityBulletin service

type ActivityBulletinHandler interface {
	FindUserActivity(context.Context, *ActivityRequest, *ActivityPagedResponse) error
	FindMostRecentActivity(context.Context, *RecentActivityRequest, *ActivityListResponse) error
	FindActivityCount(context.Context, *ActivityCountRequest, *ActivityCountResponse) error
	UpdateActivity(context.Context, *UpdateActivityRequest, *ActivityResponse) error
}

func RegisterActivityBulletinHandler(s server.Server, hdlr ActivityBulletinHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ActivityBulletin{hdlr}, opts...))
}

type ActivityBulletin struct {
	ActivityBulletinHandler
}

func (h *ActivityBulletin) FindUserActivity(ctx context.Context, in *ActivityRequest, out *ActivityPagedResponse) error {
	return h.ActivityBulletinHandler.FindUserActivity(ctx, in, out)
}

func (h *ActivityBulletin) FindMostRecentActivity(ctx context.Context, in *RecentActivityRequest, out *ActivityListResponse) error {
	return h.ActivityBulletinHandler.FindMostRecentActivity(ctx, in, out)
}

func (h *ActivityBulletin) FindActivityCount(ctx context.Context, in *ActivityCountRequest, out *ActivityCountResponse) error {
	return h.ActivityBulletinHandler.FindActivityCount(ctx, in, out)
}

func (h *ActivityBulletin) UpdateActivity(ctx context.Context, in *UpdateActivityRequest, out *ActivityResponse) error {
	return h.ActivityBulletinHandler.UpdateActivity(ctx, in, out)
}

func init() { proto.RegisterFile("proto/activity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xdd, 0xb4, 0xd9, 0x6e, 0x7b, 0x6b, 0xb5, 0x0e, 0x6d, 0x8d, 0x41, 0x6a, 0x09, 0x08, 0x45,
	0xa1, 0x42, 0x7d, 0x17, 0xaa, 0x45, 0x28, 0x28, 0x2e, 0x91, 0xf5, 0xc5, 0xa7, 0x34, 0x19, 0x24,
	0xda, 0x26, 0xb1, 0x73, 0xb3, 0x50, 0xbf, 0xc0, 0x6f, 0xf0, 0x13, 0xfd, 0x0a, 0x99, 0x99, 0x4c,
	0x32, 0x93, 0xb6, 0xba, 0x2c, 0xec, 0x53, 0x73, 0xee, 0xcc, 0xdc, 0x73, 0xee, 0x99, 0x3b, 0xb7,
	0x30, 0xc8, 0x76, 0x29, 0xa6, 0x2f, 0x83, 0x10, 0xe3, 0xeb, 0x18, 0xf7, 0x33, 0x01, 0x49, 0x5b,
	0x61, 0x2f, 0x87, 0x07, 0x8b, 0xe2, 0xdb, 0xa7, 0x3f, 0x72, 0xca, 0x90, 0x8c, 0xa0, 0x95, 0x33,
	0xba, 0x5b, 0x2d, 0x1d, 0x6b, 0x62, 0x4d, 0x3b, 0x7e, 0x81, 0x88, 0x0b, 0xed, 0x74, 0xfd, 0x8d,
	0x86, 0xb8, 0x5a, 0x3a, 0x0d, 0xb1, 0x52, 0x62, 0x42, 0xc0, 0xce, 0x82, 0xaf, 0xd4, 0x69, 0x4e,
	0xac, 0x69, 0xcf, 0x17, 0xdf, 0x7c, 0x3f, 0xff, 0xfd, 0x14, 0xff, 0xa4, 0x8e, 0x2d, 0xe2, 0x25,
	0xf6, 0x56, 0x30, 0xf4, 0x69, 0x48, 0x13, 0xac, 0x93, 0xeb, 0x24, 0x56, 0x8d, 0x64, 0x00, 0xe7,
	0x61, 0x9a, 0x27, 0x28, 0xd8, 0x7b, 0xbe, 0x04, 0xde, 0x1c, 0x06, 0x2a, 0xc9, 0x5b, 0x1e, 0xb8,
	0x41, 0x26, 0x6f, 0x0b, 0xc3, 0xab, 0x2c, 0x0a, 0x90, 0xd6, 0xe9, 0xc7, 0x00, 0xca, 0x9a, 0xf2,
	0x98, 0x16, 0xe1, 0xde, 0x30, 0x4a, 0x93, 0x05, 0x16, 0x0e, 0x14, 0x88, 0x3c, 0x81, 0x4e, 0xb8,
	0x89, 0xc3, 0xef, 0x34, 0x5a, 0xa0, 0x30, 0xa1, 0xe3, 0x57, 0x01, 0xef, 0x77, 0x03, 0xda, 0x8a,
	0xe9, 0x26, 0x14, 0x85, 0xfd, 0x4d, 0xc3, 0x7e, 0x02, 0x36, 0xee, 0x33, 0x5a, 0x10, 0x8b, 0x6f,
	0xa3, 0x46, 0xfb, 0xd0, 0x2d, 0x8c, 0x71, 0x43, 0x9d, 0x73, 0xb1, 0x20, 0x01, 0x3f, 0xc1, 0xf2,
	0xb5, 0x5c, 0x68, 0xc9, 0x13, 0x0a, 0x93, 0x09, 0x74, 0x23, 0xca, 0xc2, 0x5d, 0x9c, 0x61, 0x9c,
	0x26, 0xce, 0x85, 0x58, 0xd6, 0x43, 0xbc, 0x4c, 0x8c, 0xb7, 0x94, 0x61, 0xb0, 0xcd, 0x9c, 0xb6,
	0x2c, 0xb3, 0x0c, 0x98, 0x26, 0x74, 0x6a, 0x26, 0x68, 0xd6, 0x81, 0x6e, 0x9d, 0xf7, 0xcb, 0x82,
	0xfe, 0x15, 0xa3, 0x3b, 0x65, 0xd0, 0x25, 0xef, 0x1d, 0xd5, 0x4f, 0xd6, 0x89, 0x7e, 0x6a, 0x98,
	0xfd, 0x24, 0x8a, 0x4d, 0x31, 0xd8, 0x14, 0x0d, 0x28, 0x01, 0x99, 0x41, 0xd9, 0xe8, 0x8e, 0x3d,
	0x69, 0x4e, 0xbb, 0x73, 0x32, 0x2b, 0x5f, 0x42, 0x79, 0xf5, 0xd5, 0x63, 0xd8, 0xc3, 0x50, 0x57,
	0x11, 0xf9, 0x94, 0x65, 0x69, 0xc2, 0xa8, 0xd0, 0x8e, 0x01, 0xe6, 0x4c, 0x3d, 0x09, 0x89, 0x88,
	0x03, 0x17, 0x5b, 0xca, 0x18, 0x57, 0x2a, 0xaf, 0x45, 0x41, 0x32, 0x03, 0x3b, 0x0a, 0x30, 0x10,
	0x7a, 0xba, 0x73, 0xb7, 0xa2, 0xad, 0x97, 0xea, 0x8b, 0x7d, 0xde, 0x6b, 0xb8, 0xa7, 0xa2, 0xcb,
	0x00, 0x03, 0x43, 0xba, 0x25, 0x72, 0xfc, 0x5b, 0x7a, 0x06, 0xfd, 0xaa, 0x97, 0x6f, 0xad, 0xfa,
	0xb9, 0xa1, 0x7a, 0x74, 0xc8, 0xc8, 0xb5, 0x1d, 0x2a, 0x7e, 0x1f, 0x33, 0xac, 0x29, 0xfe, 0xbf,
	0xd9, 0x58, 0xbd, 0x5b, 0x7e, 0xfe, 0x2e, 0x55, 0x8b, 0xfc, 0x52, 0xf5, 0x33, 0xe8, 0x19, 0xd3,
	0xa2, 0x1a, 0x2a, 0x96, 0x3e, 0x54, 0xae, 0xab, 0x4e, 0x28, 0x86, 0xca, 0xad, 0xd5, 0xbd, 0x30,
	0xd4, 0x3d, 0x3a, 0x54, 0x27, 0x09, 0xc4, 0xa6, 0xf9, 0x9f, 0x46, 0x75, 0x8f, 0x6f, 0xf2, 0xcd,
	0x86, 0x62, 0x9c, 0x90, 0x4b, 0xe8, 0xbf, 0x8b, 0x93, 0x48, 0xef, 0x1c, 0xf2, 0xf8, 0x88, 0xb7,
	0x72, 0x86, 0xb9, 0x4f, 0x0f, 0x97, 0x8c, 0x6e, 0xf6, 0xce, 0xc8, 0x17, 0x18, 0xf1, 0x8c, 0x1f,
	0x52, 0xee, 0xbb, 0x3e, 0x86, 0x89, 0x76, 0xf8, 0xe8, 0x80, 0x76, 0xc7, 0x27, 0xec, 0xad, 0x92,
	0x7f, 0x86, 0x87, 0x3c, 0xb9, 0x69, 0xf3, 0xf8, 0x54, 0xdd, 0xa7, 0x45, 0x1b, 0xc6, 0x7b, 0x67,
	0xe4, 0x23, 0xdc, 0x37, 0x87, 0xb6, 0x2e, 0xf6, 0xe8, 0x38, 0x77, 0xdd, 0x63, 0x2e, 0xa9, 0x84,
	0xeb, 0x96, 0xf8, 0x33, 0x7c, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xf9, 0x7a, 0x48, 0x24,
	0x07, 0x00, 0x00,
}
