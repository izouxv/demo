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
	Id          int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
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

func (m *Feedback) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
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

func init() {
	proto.RegisterType((*Feedback)(nil), "feedback.Feedback")
	proto.RegisterType((*AddFeedbackRequest)(nil), "feedback.AddFeedbackRequest")
	proto.RegisterType((*AddFeedbackReply)(nil), "feedback.AddFeedbackReply")
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

// Server API for FeedBack service

type FeedBackServer interface {
	AddFeedback(context.Context, *AddFeedbackRequest) (*AddFeedbackReply, error)
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

var _FeedBack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feedback.FeedBack",
	HandlerType: (*FeedBackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeedback",
			Handler:    _FeedBack_AddFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback.proto",
}

func init() { proto.RegisterFile("feedback.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x93, 0x4f, 0x6a, 0xf3, 0x30,
	0x14, 0xc4, 0x3f, 0x3b, 0x5f, 0x12, 0xfb, 0x05, 0x42, 0x11, 0xa5, 0x88, 0x10, 0x8a, 0xc9, 0xca,
	0xab, 0x2c, 0xda, 0x0b, 0xf4, 0x0f, 0x14, 0xb2, 0x15, 0xed, 0x01, 0x1c, 0xe9, 0x19, 0x44, 0x1d,
	0x4b, 0x95, 0xe5, 0xd2, 0x5c, 0xaf, 0x67, 0xe8, 0x81, 0x8a, 0x24, 0x3b, 0x76, 0x29, 0x85, 0x1e,
	0xa0, 0xcb, 0x99, 0x9f, 0x19, 0xc4, 0x78, 0x1e, 0x2c, 0x4b, 0x44, 0xb1, 0x2f, 0xf8, 0xf3, 0x56,
	0x1b, 0x65, 0x15, 0x49, 0x7a, 0xbd, 0x79, 0x8f, 0x21, 0x79, 0xe8, 0x04, 0x59, 0x42, 0x2c, 0x05,
	0x8d, 0xb2, 0x28, 0x9f, 0xb2, 0x58, 0x0a, 0x92, 0xc1, 0x42, 0x60, 0xc3, 0x8d, 0xd4, 0x56, 0xaa,
	0x9a, 0xc6, 0x59, 0x94, 0xa7, 0x6c, 0x6c, 0x91, 0x4b, 0x00, 0x81, 0xaf, 0x92, 0xe3, 0xae, 0x2e,
	0x15, 0x9d, 0xf8, 0x0f, 0x46, 0x0e, 0xa1, 0x30, 0x2f, 0xb4, 0xf6, 0xf0, 0xbf, 0x87, 0xbd, 0x24,
	0x2b, 0x48, 0xda, 0x06, 0x8d, 0x47, 0x53, 0x8f, 0x4e, 0xda, 0xa5, 0x1e, 0xd4, 0x5e, 0x56, 0x21,
	0x75, 0x16, 0x52, 0x07, 0xc7, 0x71, 0x7c, 0xb3, 0x58, 0x0b, 0xcf, 0xe7, 0x81, 0x0f, 0x0e, 0x39,
	0x87, 0x69, 0x29, 0x2b, 0x6c, 0x68, 0x92, 0x4d, 0xf2, 0x94, 0x05, 0xe1, 0xde, 0xc2, 0x55, 0x6d,
	0x0b, 0x6e, 0x69, 0x1a, 0xde, 0xd2, 0x49, 0x97, 0xc7, 0x0d, 0x16, 0x16, 0x1f, 0xe5, 0x01, 0x29,
	0x64, 0x51, 0x3e, 0x61, 0x23, 0xc7, 0xf1, 0x56, 0x8b, 0x9e, 0x2f, 0x02, 0x1f, 0x9c, 0xcd, 0x47,
	0x0c, 0xe4, 0x56, 0x88, 0xbe, 0x47, 0x86, 0x2f, 0x2d, 0x36, 0x96, 0x5c, 0xc0, 0xac, 0x51, 0xad,
	0xe1, 0xe8, 0x2b, 0x4d, 0x59, 0xa7, 0xfe, 0x6a, 0xfd, 0x65, 0xad, 0x37, 0x70, 0xf6, 0xa5, 0x55,
	0x5d, 0x1d, 0xc9, 0x1a, 0x52, 0x34, 0x46, 0x99, 0x7b, 0x25, 0xb0, 0x5b, 0xea, 0x60, 0x74, 0x03,
	0x8e, 0xfb, 0x01, 0x5f, 0x3d, 0x85, 0x71, 0xdf, 0xb9, 0x71, 0xef, 0x60, 0x31, 0x4a, 0x23, 0xeb,
	0xed, 0xe9, 0x26, 0xbe, 0xff, 0xba, 0xd5, 0xea, 0x07, 0xaa, 0xab, 0xe3, 0xe6, 0xdf, 0x7e, 0xe6,
	0xaf, 0xe8, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff, 0x71, 0xf9, 0xfd, 0x5d, 0x57, 0x03, 0x00, 0x00,
}