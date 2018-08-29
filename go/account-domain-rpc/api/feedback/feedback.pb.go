// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedback.proto

/*
Package feedback is a generated protocol buffer package.

It is generated from these files:
	feedback.proto

It has these top-level messages:
	Feedback
	AddFeedbackRequest
	AddFeedbackReply
	AddFeedbackBaseTenantRequest
	AddFeedbackBaseTenantReply
	GetFeedbacksRequest
	GetFeedbacksReply
	GetFeedbacksByTypeRequest
	GetFeedbacksByTypeReply
	GetFeedbackRequest
	GetFeedbackReply
	DelFeedbackRequest
	DelFeedbackReply
	BatchFeedbackRequest
	BatchFeedbackReply
*/
package feedback

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Feedback struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	DeviceInfo  string `protobuf:"bytes,3,opt,name=deviceInfo" json:"deviceInfo,omitempty"`
	AppInfo     string `protobuf:"bytes,4,opt,name=appInfo" json:"appInfo,omitempty"`
	UserInfo    string `protobuf:"bytes,5,opt,name=userInfo" json:"userInfo,omitempty"`
	MobileInfo  string `protobuf:"bytes,6,opt,name=mobileInfo" json:"mobileInfo,omitempty"`
	ExtendInfo  string `protobuf:"bytes,7,opt,name=extendInfo" json:"extendInfo,omitempty"`
	Files       string `protobuf:"bytes,8,opt,name=files" json:"files,omitempty"`
	Contact     string `protobuf:"bytes,9,opt,name=contact" json:"contact,omitempty"`
	CreateTime  int64  `protobuf:"varint,10,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime  int64  `protobuf:"varint,11,opt,name=updateTime" json:"updateTime,omitempty"`
	Type        int32  `protobuf:"varint,12,opt,name=type" json:"type,omitempty"`
	Tid         int64  `protobuf:"varint,13,opt,name=tid" json:"tid,omitempty"`
}

func (m *Feedback) Reset()                    { *m = Feedback{} }
func (m *Feedback) String() string            { return proto.CompactTextString(m) }
func (*Feedback) ProtoMessage()               {}
func (*Feedback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Feedback) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Feedback) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Feedback) GetDeviceInfo() string {
	if m != nil {
		return m.DeviceInfo
	}
	return ""
}

func (m *Feedback) GetAppInfo() string {
	if m != nil {
		return m.AppInfo
	}
	return ""
}

func (m *Feedback) GetUserInfo() string {
	if m != nil {
		return m.UserInfo
	}
	return ""
}

func (m *Feedback) GetMobileInfo() string {
	if m != nil {
		return m.MobileInfo
	}
	return ""
}

func (m *Feedback) GetExtendInfo() string {
	if m != nil {
		return m.ExtendInfo
	}
	return ""
}

func (m *Feedback) GetFiles() string {
	if m != nil {
		return m.Files
	}
	return ""
}

func (m *Feedback) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func (m *Feedback) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Feedback) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Feedback) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Feedback) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type AddFeedbackRequest struct {
	Source      string   `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	DeviceInfo  string   `protobuf:"bytes,3,opt,name=deviceInfo" json:"deviceInfo,omitempty"`
	AppInfo     string   `protobuf:"bytes,4,opt,name=appInfo" json:"appInfo,omitempty"`
	UserInfo    string   `protobuf:"bytes,5,opt,name=userInfo" json:"userInfo,omitempty"`
	MobileInfo  string   `protobuf:"bytes,6,opt,name=mobileInfo" json:"mobileInfo,omitempty"`
	ExtendInfo  string   `protobuf:"bytes,7,opt,name=extendInfo" json:"extendInfo,omitempty"`
	Files       []string `protobuf:"bytes,8,rep,name=files" json:"files,omitempty"`
	Contact     string   `protobuf:"bytes,9,opt,name=contact" json:"contact,omitempty"`
	CreateTime  int64    `protobuf:"varint,10,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime  int64    `protobuf:"varint,11,opt,name=updateTime" json:"updateTime,omitempty"`
}

func (m *AddFeedbackRequest) Reset()                    { *m = AddFeedbackRequest{} }
func (m *AddFeedbackRequest) String() string            { return proto.CompactTextString(m) }
func (*AddFeedbackRequest) ProtoMessage()               {}
func (*AddFeedbackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddFeedbackRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AddFeedbackRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddFeedbackRequest) GetDeviceInfo() string {
	if m != nil {
		return m.DeviceInfo
	}
	return ""
}

func (m *AddFeedbackRequest) GetAppInfo() string {
	if m != nil {
		return m.AppInfo
	}
	return ""
}

func (m *AddFeedbackRequest) GetUserInfo() string {
	if m != nil {
		return m.UserInfo
	}
	return ""
}

func (m *AddFeedbackRequest) GetMobileInfo() string {
	if m != nil {
		return m.MobileInfo
	}
	return ""
}

func (m *AddFeedbackRequest) GetExtendInfo() string {
	if m != nil {
		return m.ExtendInfo
	}
	return ""
}

func (m *AddFeedbackRequest) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *AddFeedbackRequest) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func (m *AddFeedbackRequest) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *AddFeedbackRequest) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type AddFeedbackReply struct {
	ErrorCode int32 `protobuf:"varint,1,opt,name=errorCode" json:"errorCode,omitempty"`
	Id        int32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *AddFeedbackReply) Reset()                    { *m = AddFeedbackReply{} }
func (m *AddFeedbackReply) String() string            { return proto.CompactTextString(m) }
func (*AddFeedbackReply) ProtoMessage()               {}
func (*AddFeedbackReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddFeedbackReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *AddFeedbackReply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AddFeedbackBaseTenantRequest struct {
	Tid         int64  `protobuf:"varint,12,opt,name=tid" json:"tid,omitempty"`
	Type        int32  `protobuf:"varint,13,opt,name=type" json:"type,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	CreateTime  int64  `protobuf:"varint,10,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime  int64  `protobuf:"varint,11,opt,name=updateTime" json:"updateTime,omitempty"`
	Files       string `protobuf:"bytes,14,opt,name=files" json:"files,omitempty"`
}

func (m *AddFeedbackBaseTenantRequest) Reset()                    { *m = AddFeedbackBaseTenantRequest{} }
func (m *AddFeedbackBaseTenantRequest) String() string            { return proto.CompactTextString(m) }
func (*AddFeedbackBaseTenantRequest) ProtoMessage()               {}
func (*AddFeedbackBaseTenantRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddFeedbackBaseTenantRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *AddFeedbackBaseTenantRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AddFeedbackBaseTenantRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddFeedbackBaseTenantRequest) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *AddFeedbackBaseTenantRequest) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *AddFeedbackBaseTenantRequest) GetFiles() string {
	if m != nil {
		return m.Files
	}
	return ""
}

type AddFeedbackBaseTenantReply struct {
	ErrorCode int32 `protobuf:"varint,1,opt,name=errorCode" json:"errorCode,omitempty"`
}

func (m *AddFeedbackBaseTenantReply) Reset()                    { *m = AddFeedbackBaseTenantReply{} }
func (m *AddFeedbackBaseTenantReply) String() string            { return proto.CompactTextString(m) }
func (*AddFeedbackBaseTenantReply) ProtoMessage()               {}
func (*AddFeedbackBaseTenantReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddFeedbackBaseTenantReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type GetFeedbacksRequest struct {
	Tid   int64 `protobuf:"varint,1,opt,name=tid" json:"tid,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Count int32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
}

func (m *GetFeedbacksRequest) Reset()                    { *m = GetFeedbacksRequest{} }
func (m *GetFeedbacksRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFeedbacksRequest) ProtoMessage()               {}
func (*GetFeedbacksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetFeedbacksRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *GetFeedbacksRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetFeedbacksRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetFeedbacksReply struct {
	FeedBack   []*Feedback `protobuf:"bytes,1,rep,name=feedBack" json:"feedBack,omitempty"`
	ErrorCode  int32       `protobuf:"varint,2,opt,name=errorCode" json:"errorCode,omitempty"`
	TotalCount int32       `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *GetFeedbacksReply) Reset()                    { *m = GetFeedbacksReply{} }
func (m *GetFeedbacksReply) String() string            { return proto.CompactTextString(m) }
func (*GetFeedbacksReply) ProtoMessage()               {}
func (*GetFeedbacksReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetFeedbacksReply) GetFeedBack() []*Feedback {
	if m != nil {
		return m.FeedBack
	}
	return nil
}

func (m *GetFeedbacksReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GetFeedbacksReply) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type GetFeedbacksByTypeRequest struct {
	Tid   int64 `protobuf:"varint,1,opt,name=tid" json:"tid,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Count int32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Type  int32 `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
}

func (m *GetFeedbacksByTypeRequest) Reset()                    { *m = GetFeedbacksByTypeRequest{} }
func (m *GetFeedbacksByTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFeedbacksByTypeRequest) ProtoMessage()               {}
func (*GetFeedbacksByTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetFeedbacksByTypeRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *GetFeedbacksByTypeRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetFeedbacksByTypeRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetFeedbacksByTypeRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type GetFeedbacksByTypeReply struct {
	FeedBack   []*Feedback `protobuf:"bytes,1,rep,name=feedBack" json:"feedBack,omitempty"`
	ErrorCode  int32       `protobuf:"varint,2,opt,name=errorCode" json:"errorCode,omitempty"`
	TotalCount int32       `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *GetFeedbacksByTypeReply) Reset()                    { *m = GetFeedbacksByTypeReply{} }
func (m *GetFeedbacksByTypeReply) String() string            { return proto.CompactTextString(m) }
func (*GetFeedbacksByTypeReply) ProtoMessage()               {}
func (*GetFeedbacksByTypeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetFeedbacksByTypeReply) GetFeedBack() []*Feedback {
	if m != nil {
		return m.FeedBack
	}
	return nil
}

func (m *GetFeedbacksByTypeReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GetFeedbacksByTypeReply) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type GetFeedbackRequest struct {
	Id  int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Tid int64 `protobuf:"varint,2,opt,name=tid" json:"tid,omitempty"`
}

func (m *GetFeedbackRequest) Reset()                    { *m = GetFeedbackRequest{} }
func (m *GetFeedbackRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFeedbackRequest) ProtoMessage()               {}
func (*GetFeedbackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetFeedbackRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetFeedbackRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type GetFeedbackReply struct {
	FeedBack  *Feedback `protobuf:"bytes,1,opt,name=feedBack" json:"feedBack,omitempty"`
	ErrorCode int32     `protobuf:"varint,2,opt,name=errorCode" json:"errorCode,omitempty"`
}

func (m *GetFeedbackReply) Reset()                    { *m = GetFeedbackReply{} }
func (m *GetFeedbackReply) String() string            { return proto.CompactTextString(m) }
func (*GetFeedbackReply) ProtoMessage()               {}
func (*GetFeedbackReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetFeedbackReply) GetFeedBack() *Feedback {
	if m != nil {
		return m.FeedBack
	}
	return nil
}

func (m *GetFeedbackReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type DelFeedbackRequest struct {
	Tid int64   `protobuf:"varint,1,opt,name=tid" json:"tid,omitempty"`
	Ids []int32 `protobuf:"varint,2,rep,packed,name=ids" json:"ids,omitempty"`
}

func (m *DelFeedbackRequest) Reset()                    { *m = DelFeedbackRequest{} }
func (m *DelFeedbackRequest) String() string            { return proto.CompactTextString(m) }
func (*DelFeedbackRequest) ProtoMessage()               {}
func (*DelFeedbackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DelFeedbackRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *DelFeedbackRequest) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type DelFeedbackReply struct {
	ErrorCode int32 `protobuf:"varint,1,opt,name=errorCode" json:"errorCode,omitempty"`
}

func (m *DelFeedbackReply) Reset()                    { *m = DelFeedbackReply{} }
func (m *DelFeedbackReply) String() string            { return proto.CompactTextString(m) }
func (*DelFeedbackReply) ProtoMessage()               {}
func (*DelFeedbackReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DelFeedbackReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type BatchFeedbackRequest struct {
	Tid int64   `protobuf:"varint,1,opt,name=tid" json:"tid,omitempty"`
	Ids []int32 `protobuf:"varint,2,rep,packed,name=ids" json:"ids,omitempty"`
}

func (m *BatchFeedbackRequest) Reset()                    { *m = BatchFeedbackRequest{} }
func (m *BatchFeedbackRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchFeedbackRequest) ProtoMessage()               {}
func (*BatchFeedbackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *BatchFeedbackRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *BatchFeedbackRequest) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type BatchFeedbackReply struct {
	ErrorCode int32       `protobuf:"varint,1,opt,name=errorCode" json:"errorCode,omitempty"`
	FeedBack  []*Feedback `protobuf:"bytes,3,rep,name=feedBack" json:"feedBack,omitempty"`
}

func (m *BatchFeedbackReply) Reset()                    { *m = BatchFeedbackReply{} }
func (m *BatchFeedbackReply) String() string            { return proto.CompactTextString(m) }
func (*BatchFeedbackReply) ProtoMessage()               {}
func (*BatchFeedbackReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *BatchFeedbackReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *BatchFeedbackReply) GetFeedBack() []*Feedback {
	if m != nil {
		return m.FeedBack
	}
	return nil
}

func init() {
	proto.RegisterType((*Feedback)(nil), "feedback.Feedback")
	proto.RegisterType((*AddFeedbackRequest)(nil), "feedback.AddFeedbackRequest")
	proto.RegisterType((*AddFeedbackReply)(nil), "feedback.AddFeedbackReply")
	proto.RegisterType((*AddFeedbackBaseTenantRequest)(nil), "feedback.AddFeedbackBaseTenantRequest")
	proto.RegisterType((*AddFeedbackBaseTenantReply)(nil), "feedback.AddFeedbackBaseTenantReply")
	proto.RegisterType((*GetFeedbacksRequest)(nil), "feedback.GetFeedbacksRequest")
	proto.RegisterType((*GetFeedbacksReply)(nil), "feedback.GetFeedbacksReply")
	proto.RegisterType((*GetFeedbacksByTypeRequest)(nil), "feedback.GetFeedbacksByTypeRequest")
	proto.RegisterType((*GetFeedbacksByTypeReply)(nil), "feedback.GetFeedbacksByTypeReply")
	proto.RegisterType((*GetFeedbackRequest)(nil), "feedback.GetFeedbackRequest")
	proto.RegisterType((*GetFeedbackReply)(nil), "feedback.GetFeedbackReply")
	proto.RegisterType((*DelFeedbackRequest)(nil), "feedback.DelFeedbackRequest")
	proto.RegisterType((*DelFeedbackReply)(nil), "feedback.DelFeedbackReply")
	proto.RegisterType((*BatchFeedbackRequest)(nil), "feedback.BatchFeedbackRequest")
	proto.RegisterType((*BatchFeedbackReply)(nil), "feedback.BatchFeedbackReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FeedBack service

type FeedBackClient interface {
	AddFeedback(ctx context.Context, in *AddFeedbackRequest, opts ...grpc.CallOption) (*AddFeedbackReply, error)
	AddFeedbackBaseTenant(ctx context.Context, in *AddFeedbackBaseTenantRequest, opts ...grpc.CallOption) (*AddFeedbackBaseTenantReply, error)
	GetFeedbacks(ctx context.Context, in *GetFeedbacksRequest, opts ...grpc.CallOption) (*GetFeedbacksReply, error)
	GetFeedbacksByType(ctx context.Context, in *GetFeedbacksByTypeRequest, opts ...grpc.CallOption) (*GetFeedbacksByTypeReply, error)
	GetFeedback(ctx context.Context, in *GetFeedbackRequest, opts ...grpc.CallOption) (*GetFeedbackReply, error)
	DelFeedback(ctx context.Context, in *DelFeedbackRequest, opts ...grpc.CallOption) (*DelFeedbackReply, error)
	BatchFeedback(ctx context.Context, in *BatchFeedbackRequest, opts ...grpc.CallOption) (*BatchFeedbackReply, error)
}

type feedBackClient struct {
	cc *grpc.ClientConn
}

func NewFeedBackClient(cc *grpc.ClientConn) FeedBackClient {
	return &feedBackClient{cc}
}

func (c *feedBackClient) AddFeedback(ctx context.Context, in *AddFeedbackRequest, opts ...grpc.CallOption) (*AddFeedbackReply, error) {
	out := new(AddFeedbackReply)
	err := grpc.Invoke(ctx, "/feedback.FeedBack/AddFeedback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) AddFeedbackBaseTenant(ctx context.Context, in *AddFeedbackBaseTenantRequest, opts ...grpc.CallOption) (*AddFeedbackBaseTenantReply, error) {
	out := new(AddFeedbackBaseTenantReply)
	err := grpc.Invoke(ctx, "/feedback.FeedBack/AddFeedbackBaseTenant", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) GetFeedbacks(ctx context.Context, in *GetFeedbacksRequest, opts ...grpc.CallOption) (*GetFeedbacksReply, error) {
	out := new(GetFeedbacksReply)
	err := grpc.Invoke(ctx, "/feedback.FeedBack/GetFeedbacks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) GetFeedbacksByType(ctx context.Context, in *GetFeedbacksByTypeRequest, opts ...grpc.CallOption) (*GetFeedbacksByTypeReply, error) {
	out := new(GetFeedbacksByTypeReply)
	err := grpc.Invoke(ctx, "/feedback.FeedBack/GetFeedbacksByType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) GetFeedback(ctx context.Context, in *GetFeedbackRequest, opts ...grpc.CallOption) (*GetFeedbackReply, error) {
	out := new(GetFeedbackReply)
	err := grpc.Invoke(ctx, "/feedback.FeedBack/GetFeedback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) DelFeedback(ctx context.Context, in *DelFeedbackRequest, opts ...grpc.CallOption) (*DelFeedbackReply, error) {
	out := new(DelFeedbackReply)
	err := grpc.Invoke(ctx, "/feedback.FeedBack/DelFeedback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) BatchFeedback(ctx context.Context, in *BatchFeedbackRequest, opts ...grpc.CallOption) (*BatchFeedbackReply, error) {
	out := new(BatchFeedbackReply)
	err := grpc.Invoke(ctx, "/feedback.FeedBack/BatchFeedback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FeedBack service

type FeedBackServer interface {
	AddFeedback(context.Context, *AddFeedbackRequest) (*AddFeedbackReply, error)
	AddFeedbackBaseTenant(context.Context, *AddFeedbackBaseTenantRequest) (*AddFeedbackBaseTenantReply, error)
	GetFeedbacks(context.Context, *GetFeedbacksRequest) (*GetFeedbacksReply, error)
	GetFeedbacksByType(context.Context, *GetFeedbacksByTypeRequest) (*GetFeedbacksByTypeReply, error)
	GetFeedback(context.Context, *GetFeedbackRequest) (*GetFeedbackReply, error)
	DelFeedback(context.Context, *DelFeedbackRequest) (*DelFeedbackReply, error)
	BatchFeedback(context.Context, *BatchFeedbackRequest) (*BatchFeedbackReply, error)
}

func RegisterFeedBackServer(s *grpc.Server, srv FeedBackServer) {
	s.RegisterService(&_FeedBack_serviceDesc, srv)
}

func _FeedBack_AddFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).AddFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedBack/AddFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).AddFeedback(ctx, req.(*AddFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_AddFeedbackBaseTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedbackBaseTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).AddFeedbackBaseTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedBack/AddFeedbackBaseTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).AddFeedbackBaseTenant(ctx, req.(*AddFeedbackBaseTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_GetFeedbacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbacksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).GetFeedbacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedBack/GetFeedbacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).GetFeedbacks(ctx, req.(*GetFeedbacksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_GetFeedbacksByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbacksByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).GetFeedbacksByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedBack/GetFeedbacksByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).GetFeedbacksByType(ctx, req.(*GetFeedbacksByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_GetFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).GetFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedBack/GetFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).GetFeedback(ctx, req.(*GetFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_DelFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).DelFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedBack/DelFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).DelFeedback(ctx, req.(*DelFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_BatchFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).BatchFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedBack/BatchFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).BatchFeedback(ctx, req.(*BatchFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedBack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feedback.FeedBack",
	HandlerType: (*FeedBackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeedback",
			Handler:    _FeedBack_AddFeedback_Handler,
		},
		{
			MethodName: "AddFeedbackBaseTenant",
			Handler:    _FeedBack_AddFeedbackBaseTenant_Handler,
		},
		{
			MethodName: "GetFeedbacks",
			Handler:    _FeedBack_GetFeedbacks_Handler,
		},
		{
			MethodName: "GetFeedbacksByType",
			Handler:    _FeedBack_GetFeedbacksByType_Handler,
		},
		{
			MethodName: "GetFeedback",
			Handler:    _FeedBack_GetFeedback_Handler,
		},
		{
			MethodName: "DelFeedback",
			Handler:    _FeedBack_DelFeedback_Handler,
		},
		{
			MethodName: "BatchFeedback",
			Handler:    _FeedBack_BatchFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback.proto",
}

func init() { proto.RegisterFile("feedback.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xed, 0xa4, 0x4d, 0x26, 0x6d, 0x15, 0x86, 0x02, 0xc6, 0x84, 0x2a, 0x2c, 0x08, 0xe5,
	0x54, 0xa1, 0x22, 0x21, 0xd4, 0x13, 0xa4, 0x08, 0x54, 0x09, 0x09, 0x61, 0xf5, 0x8c, 0x70, 0xec,
	0x69, 0xb1, 0x92, 0xda, 0xc6, 0xde, 0x20, 0x72, 0xe4, 0xc4, 0x4b, 0xf1, 0x16, 0xf0, 0x10, 0x3c,
	0x06, 0xf2, 0xae, 0x7f, 0xd6, 0x8e, 0x43, 0x10, 0x45, 0x3d, 0x70, 0xdb, 0x99, 0x6f, 0x67, 0x76,
	0xf6, 0x9b, 0x6f, 0xc7, 0x86, 0xdd, 0x33, 0x22, 0x6f, 0xe2, 0xb8, 0xd3, 0x83, 0x28, 0x0e, 0x79,
	0x88, 0x9d, 0xdc, 0x66, 0x3f, 0x75, 0xe8, 0xbc, 0xcc, 0x0c, 0xdc, 0x05, 0xdd, 0xf7, 0x4c, 0x6d,
	0xa8, 0x8d, 0xda, 0xb6, 0xee, 0x7b, 0x38, 0x84, 0x9e, 0x47, 0x89, 0x1b, 0xfb, 0x11, 0xf7, 0xc3,
	0xc0, 0xd4, 0x87, 0xda, 0xa8, 0x6b, 0xab, 0x2e, 0xdc, 0x07, 0xf0, 0xe8, 0x93, 0xef, 0xd2, 0x49,
	0x70, 0x16, 0x9a, 0x86, 0xd8, 0xa0, 0x78, 0xd0, 0x84, 0x2d, 0x27, 0x8a, 0x04, 0xd8, 0x12, 0x60,
	0x6e, 0xa2, 0x05, 0x9d, 0x79, 0x42, 0xb1, 0x80, 0xda, 0x02, 0x2a, 0xec, 0x34, 0xeb, 0x45, 0x38,
	0xf1, 0x67, 0x32, 0xeb, 0xa6, 0xcc, 0x5a, 0x7a, 0x52, 0x9c, 0x3e, 0x73, 0x0a, 0x3c, 0x81, 0x6f,
	0x49, 0xbc, 0xf4, 0xe0, 0x1e, 0xb4, 0xcf, 0xfc, 0x19, 0x25, 0x66, 0x47, 0x40, 0xd2, 0x48, 0x6b,
	0x71, 0xc3, 0x80, 0x3b, 0x2e, 0x37, 0xbb, 0xb2, 0x96, 0xcc, 0x4c, 0xf3, 0xb9, 0x31, 0x39, 0x9c,
	0x4e, 0xfd, 0x0b, 0x32, 0x61, 0xa8, 0x8d, 0x0c, 0x5b, 0xf1, 0xa4, 0xf8, 0x3c, 0xf2, 0x72, 0xbc,
	0x27, 0xf1, 0xd2, 0x83, 0x08, 0x2d, 0xbe, 0x88, 0xc8, 0xdc, 0x16, 0xcc, 0x89, 0x35, 0xf6, 0xc1,
	0xe0, 0xbe, 0x67, 0xee, 0x88, 0xcd, 0xe9, 0x92, 0xfd, 0xd0, 0x01, 0x9f, 0x7b, 0x5e, 0xce, 0xb6,
	0x4d, 0x1f, 0xe7, 0x94, 0x70, 0xbc, 0x09, 0x9b, 0x49, 0x38, 0x8f, 0x5d, 0x12, 0xc4, 0x77, 0xed,
	0xcc, 0xfa, 0x1f, 0xc8, 0x37, 0xae, 0x80, 0x7c, 0xf6, 0x0c, 0xfa, 0x15, 0x56, 0xa3, 0xd9, 0x02,
	0x07, 0xd0, 0xa5, 0x38, 0x0e, 0xe3, 0xe3, 0xd0, 0xa3, 0x4c, 0xcf, 0xa5, 0x23, 0x93, 0xb9, 0x9e,
	0xcb, 0x9c, 0x7d, 0xd3, 0x60, 0xa0, 0xa4, 0x18, 0x3b, 0x09, 0x9d, 0x52, 0xe0, 0x04, 0x3c, 0x6f,
	0x51, 0xd6, 0xcb, 0xed, 0xa2, 0x97, 0x45, 0xc7, 0x77, 0x94, 0x8e, 0xff, 0x51, 0xc3, 0x2e, 0xa5,
	0xb3, 0x82, 0xda, 0x5d, 0x45, 0xd7, 0xec, 0x08, 0xac, 0x15, 0xd5, 0xaf, 0xa5, 0x82, 0xbd, 0x85,
	0xeb, 0xaf, 0x88, 0xe7, 0xb1, 0x49, 0xed, 0xc2, 0x5a, 0xe5, 0xc2, 0x91, 0x73, 0x4e, 0x19, 0x6b,
	0x62, 0x9d, 0x96, 0xe3, 0x86, 0xf3, 0x80, 0x0b, 0xe9, 0xb5, 0x6d, 0x69, 0xb0, 0x2f, 0x1a, 0x5c,
	0xab, 0xe6, 0x4c, 0xcb, 0x38, 0x00, 0x31, 0x73, 0xc6, 0x8e, 0x3b, 0x35, 0xb5, 0xa1, 0x31, 0xea,
	0x1d, 0xe2, 0x41, 0x31, 0x94, 0x8a, 0xe6, 0x15, 0x7b, 0xaa, 0x65, 0xeb, 0xf5, 0x0e, 0xee, 0x03,
	0xf0, 0x90, 0x3b, 0xb3, 0x63, 0xe5, 0x78, 0xc5, 0xc3, 0xa6, 0x70, 0x5b, 0x2d, 0x61, 0xbc, 0x38,
	0x5d, 0x44, 0xf4, 0x0f, 0x2e, 0x57, 0xf4, 0xbd, 0x55, 0xf6, 0x9d, 0x7d, 0xd5, 0xe0, 0x56, 0xd3,
	0x69, 0x57, 0x7f, 0xed, 0x27, 0x80, 0x4a, 0x21, 0xf9, 0x7d, 0xeb, 0x53, 0x3d, 0xbb, 0xbf, 0x5e,
	0x4e, 0xa6, 0xf7, 0xd0, 0xaf, 0xc4, 0x2d, 0x57, 0xae, 0x5d, 0xae, 0x72, 0xf6, 0x14, 0xf0, 0x05,
	0xcd, 0xea, 0x95, 0x2d, 0x77, 0xa2, 0x0f, 0x86, 0xef, 0x25, 0xa6, 0x3e, 0x34, 0x46, 0x6d, 0x3b,
	0x5d, 0xb2, 0x47, 0xd0, 0xaf, 0x44, 0xae, 0xd7, 0xf4, 0x11, 0xec, 0x8d, 0x1d, 0xee, 0x7e, 0xf8,
	0x9b, 0xd3, 0x26, 0x80, 0xb5, 0xd8, 0xf5, 0xe3, 0x44, 0x65, 0xca, 0x58, 0xdf, 0xe3, 0xc3, 0xef,
	0x2d, 0xf9, 0xc9, 0x15, 0xb4, 0x9d, 0x40, 0x4f, 0x79, 0xbc, 0x38, 0x28, 0x23, 0x97, 0x3f, 0x15,
	0x96, 0xb5, 0x02, 0x8d, 0x66, 0x0b, 0xb6, 0x81, 0xe7, 0x70, 0xa3, 0x71, 0x0e, 0xe0, 0xc3, 0xc6,
	0xb0, 0xa5, 0x31, 0x67, 0x3d, 0x58, 0xbb, 0x4f, 0x1e, 0xf4, 0x1a, 0xb6, 0x55, 0xbd, 0xe3, 0xdd,
	0x32, 0xae, 0x61, 0x98, 0x58, 0x77, 0x56, 0xc1, 0x32, 0xdb, 0xbb, 0x8a, 0x68, 0xb3, 0xd7, 0x83,
	0xf7, 0x9b, 0x83, 0x2a, 0x2f, 0xd9, 0xba, 0xf7, 0xfb, 0x4d, 0x32, 0xff, 0x09, 0xf4, 0x14, 0x50,
	0x65, 0x78, 0xf9, 0xad, 0xa8, 0x0c, 0xd7, 0x5f, 0x84, 0x4c, 0xa5, 0x68, 0x51, 0x4d, 0xb5, 0x2c,
	0x6e, 0x35, 0x55, 0x5d, 0xc0, 0x6c, 0x03, 0xdf, 0xc0, 0x4e, 0x45, 0x68, 0xb8, 0x5f, 0x6e, 0x6f,
	0x52, 0xaf, 0x35, 0x58, 0x89, 0x8b, 0x84, 0x93, 0x4d, 0xf1, 0x67, 0xf7, 0xf8, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x39, 0x90, 0x6a, 0x82, 0xeb, 0x09, 0x00, 0x00,
}
