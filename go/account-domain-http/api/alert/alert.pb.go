// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alert.proto

/*
Package alert is a generated protocol buffer package.

It is generated from these files:
	alert.proto

It has these top-level messages:
	AlertCountRequest
	AlertCountReply
	AlertRequest
	AlertNodeReply
	AlertReply
*/
package alert

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

type AlertCountRequest struct {
	Did int64 `protobuf:"varint,1,opt,name=Did" json:"Did,omitempty"`
}

func (m *AlertCountRequest) Reset()                    { *m = AlertCountRequest{} }
func (m *AlertCountRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertCountRequest) ProtoMessage()               {}
func (*AlertCountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AlertCountRequest) GetDid() int64 {
	if m != nil {
		return m.Did
	}
	return 0
}

type AlertCountReply struct {
	AlertCount int32 `protobuf:"varint,1,opt,name=AlertCount" json:"AlertCount,omitempty"`
	ErrorCode  int32 `protobuf:"varint,2,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *AlertCountReply) Reset()                    { *m = AlertCountReply{} }
func (m *AlertCountReply) String() string            { return proto.CompactTextString(m) }
func (*AlertCountReply) ProtoMessage()               {}
func (*AlertCountReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AlertCountReply) GetAlertCount() int32 {
	if m != nil {
		return m.AlertCount
	}
	return 0
}

func (m *AlertCountReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type AlertRequest struct {
	Did   int64 `protobuf:"varint,1,opt,name=Did" json:"Did,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=Page" json:"Page,omitempty"`
	Count int32 `protobuf:"varint,3,opt,name=Count" json:"Count,omitempty"`
}

func (m *AlertRequest) Reset()                    { *m = AlertRequest{} }
func (m *AlertRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertRequest) ProtoMessage()               {}
func (*AlertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AlertRequest) GetDid() int64 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *AlertRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AlertRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type AlertNodeReply struct {
	Nid        int64  `protobuf:"varint,1,opt,name=Nid" json:"Nid,omitempty"`
	Info       string `protobuf:"bytes,2,opt,name=Info" json:"Info,omitempty"`
	CreateTime int64  `protobuf:"varint,3,opt,name=CreateTime" json:"CreateTime,omitempty"`
}

func (m *AlertNodeReply) Reset()                    { *m = AlertNodeReply{} }
func (m *AlertNodeReply) String() string            { return proto.CompactTextString(m) }
func (*AlertNodeReply) ProtoMessage()               {}
func (*AlertNodeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AlertNodeReply) GetNid() int64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

func (m *AlertNodeReply) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *AlertNodeReply) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type AlertReply struct {
	NodeAlert  []*AlertNodeReply `protobuf:"bytes,1,rep,name=NodeAlert" json:"NodeAlert,omitempty"`
	AlertTotal int32             `protobuf:"varint,2,opt,name=AlertTotal" json:"AlertTotal,omitempty"`
	ErrorCode  int32             `protobuf:"varint,3,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *AlertReply) Reset()                    { *m = AlertReply{} }
func (m *AlertReply) String() string            { return proto.CompactTextString(m) }
func (*AlertReply) ProtoMessage()               {}
func (*AlertReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AlertReply) GetNodeAlert() []*AlertNodeReply {
	if m != nil {
		return m.NodeAlert
	}
	return nil
}

func (m *AlertReply) GetAlertTotal() int32 {
	if m != nil {
		return m.AlertTotal
	}
	return 0
}

func (m *AlertReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*AlertCountRequest)(nil), "alert.AlertCountRequest")
	proto.RegisterType((*AlertCountReply)(nil), "alert.AlertCountReply")
	proto.RegisterType((*AlertRequest)(nil), "alert.AlertRequest")
	proto.RegisterType((*AlertNodeReply)(nil), "alert.AlertNodeReply")
	proto.RegisterType((*AlertReply)(nil), "alert.AlertReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Alert service

type AlertClient interface {
	GetCount(ctx context.Context, in *AlertCountRequest, opts ...grpc.CallOption) (*AlertCountReply, error)
	GetAlerts(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertReply, error)
}

type alertClient struct {
	cc *grpc.ClientConn
}

func NewAlertClient(cc *grpc.ClientConn) AlertClient {
	return &alertClient{cc}
}

func (c *alertClient) GetCount(ctx context.Context, in *AlertCountRequest, opts ...grpc.CallOption) (*AlertCountReply, error) {
	out := new(AlertCountReply)
	err := grpc.Invoke(ctx, "/alert.Alert/GetCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertClient) GetAlerts(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertReply, error) {
	out := new(AlertReply)
	err := grpc.Invoke(ctx, "/alert.Alert/GetAlerts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Alert service

type AlertServer interface {
	GetCount(context.Context, *AlertCountRequest) (*AlertCountReply, error)
	GetAlerts(context.Context, *AlertRequest) (*AlertReply, error)
}

func RegisterAlertServer(s *grpc.Server, srv AlertServer) {
	s.RegisterService(&_Alert_serviceDesc, srv)
}

func _Alert_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alert/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).GetCount(ctx, req.(*AlertCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alert_GetAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).GetAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alert/GetAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).GetAlerts(ctx, req.(*AlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alert_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.Alert",
	HandlerType: (*AlertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCount",
			Handler:    _Alert_GetCount_Handler,
		},
		{
			MethodName: "GetAlerts",
			Handler:    _Alert_GetAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alert.proto",
}

func init() { proto.RegisterFile("alert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x8d, 0x6b, 0xc4, 0x4c, 0x45, 0xdb, 0xf1, 0x0f, 0x41, 0x44, 0xca, 0x82, 0xd0, 0x53,
	0x0f, 0x2d, 0x5e, 0x05, 0x89, 0x52, 0xf4, 0x10, 0x65, 0x29, 0xde, 0x23, 0x19, 0xa5, 0x10, 0xbb,
	0x71, 0xbb, 0x3d, 0xf4, 0xa2, 0x5f, 0x5d, 0x32, 0x9b, 0x98, 0x8d, 0xe8, 0x6d, 0xf6, 0xb7, 0x33,
	0xef, 0x65, 0x5e, 0x16, 0x7a, 0x59, 0x41, 0xc6, 0x8e, 0x4b, 0xa3, 0xad, 0xc6, 0x90, 0x0f, 0xf2,
	0x12, 0x06, 0x37, 0x55, 0x91, 0xe8, 0xf5, 0xd2, 0x2a, 0xfa, 0x58, 0xd3, 0xca, 0x62, 0x1f, 0xc4,
	0xed, 0x22, 0x8f, 0x83, 0x61, 0x30, 0x12, 0xaa, 0x2a, 0xe5, 0x23, 0x1c, 0xfa, 0x6d, 0x65, 0xb1,
	0xc1, 0x0b, 0x80, 0x16, 0x71, 0x6f, 0xa8, 0x3c, 0x82, 0xe7, 0x10, 0xdd, 0x19, 0xa3, 0x4d, 0xa2,
	0x73, 0x8a, 0xb7, 0xf9, 0xba, 0x05, 0xf2, 0x01, 0xf6, 0xb9, 0xf7, 0x5f, 0x4b, 0x44, 0xd8, 0x79,
	0xca, 0xde, 0x9a, 0x51, 0xae, 0xf1, 0x18, 0x42, 0x67, 0x27, 0x18, 0xba, 0x83, 0x7c, 0x86, 0x03,
	0xd6, 0x4a, 0x75, 0x4e, 0xee, 0xdb, 0xfa, 0x20, 0xd2, 0x56, 0x2d, 0x75, 0x6a, 0xf7, 0xcb, 0x57,
	0xcd, 0x6a, 0x91, 0xe2, 0xba, 0xda, 0x20, 0x31, 0x94, 0x59, 0x9a, 0x2f, 0xde, 0x89, 0x25, 0x85,
	0xf2, 0x88, 0xfc, 0xaa, 0x37, 0x74, 0x9a, 0x53, 0x88, 0x2a, 0x03, 0x26, 0x71, 0x30, 0x14, 0xa3,
	0xde, 0xe4, 0x64, 0xec, 0x12, 0xed, 0xba, 0xab, 0xb6, 0xef, 0x27, 0xa4, 0xb9, 0xb6, 0x59, 0x51,
	0xaf, 0xe2, 0x91, 0x6e, 0x48, 0xe2, 0x57, 0x48, 0x93, 0x4f, 0x08, 0x9d, 0xcc, 0x35, 0xec, 0xcd,
	0xa8, 0xce, 0x35, 0xf6, 0x4d, 0xfd, 0xdf, 0x76, 0x76, 0xfa, 0xc7, 0x4d, 0x59, 0x6c, 0xe4, 0x16,
	0x5e, 0x41, 0x34, 0x23, 0xcb, 0x7c, 0x85, 0x47, 0x7e, 0x5b, 0x33, 0x3b, 0xe8, 0x42, 0x1e, 0x7b,
	0xd9, 0xe5, 0xa7, 0x32, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x32, 0xb4, 0x85, 0xd9, 0x39, 0x02,
	0x00, 0x00,
}
