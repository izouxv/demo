// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/device_type.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/device_type.proto
	api/twins.proto

It has these top-level messages:
	Type
	AddDeviceTypeRequest
	AddDeviceTypeReply
	GetDeviceTypesRequest
	GetDeviceTypesReply
	DeleteTwinsBaseTenantRequest
	DeleteTwinsBaseTenantResponse
	GetLastUpdateTimeRequest
	GetLastUpdateTimeResponse
	GetTwinsVersionRequest
	GetTwinsVersionResponse
	AddTwinsBaseTenantRequest
	AddTwinsBaseTenantResponse
	UpdateTwinsBaseTenantRequest
	UpdateTwinsBaseTenantResponse
	GetTwinsBaseTenantRequest
	GetTwinsBaseTenantResponse
	GetTwinsDataPointBaseTenantRequest
	GetTwinsDataPointBaseTenantResponse
	GetLastUpdateTimeBaseTenantRequest
	GetLastUpdateTimeBaseTenantResponse
	GetLoraWANTwinsForGatewayMacBaseTenantRequest
	GetLoraWANTwinsForGatewayMacBaseTenantResponse
*/
package api

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

type Type struct {
	Id         int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType" json:"device_type,omitempty"`
	Status     int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Tid        int64  `protobuf:"varint,4,opt,name=tid" json:"tid,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Type) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Type) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *Type) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Type) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type AddDeviceTypeRequest struct {
	DeviceType *Type `protobuf:"bytes,1,opt,name=deviceType" json:"deviceType,omitempty"`
}

func (m *AddDeviceTypeRequest) Reset()                    { *m = AddDeviceTypeRequest{} }
func (m *AddDeviceTypeRequest) String() string            { return proto.CompactTextString(m) }
func (*AddDeviceTypeRequest) ProtoMessage()               {}
func (*AddDeviceTypeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddDeviceTypeRequest) GetDeviceType() *Type {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

type AddDeviceTypeReply struct {
	ErrorCode  int32 `protobuf:"varint,1,opt,name=errorCode" json:"errorCode,omitempty"`
	DeviceType *Type `protobuf:"bytes,2,opt,name=deviceType" json:"deviceType,omitempty"`
}

func (m *AddDeviceTypeReply) Reset()                    { *m = AddDeviceTypeReply{} }
func (m *AddDeviceTypeReply) String() string            { return proto.CompactTextString(m) }
func (*AddDeviceTypeReply) ProtoMessage()               {}
func (*AddDeviceTypeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddDeviceTypeReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *AddDeviceTypeReply) GetDeviceType() *Type {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

type GetDeviceTypesRequest struct {
	Tid     int64  `protobuf:"varint,1,opt,name=tid" json:"tid,omitempty"`
	Page    int32  `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Count   int32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	OrderBy string `protobuf:"bytes,4,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
}

func (m *GetDeviceTypesRequest) Reset()                    { *m = GetDeviceTypesRequest{} }
func (m *GetDeviceTypesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceTypesRequest) ProtoMessage()               {}
func (*GetDeviceTypesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetDeviceTypesRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *GetDeviceTypesRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetDeviceTypesRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetDeviceTypesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

type GetDeviceTypesReply struct {
	DeviceType []*Type `protobuf:"bytes,1,rep,name=deviceType" json:"deviceType,omitempty"`
	ErrorCode  int32   `protobuf:"varint,2,opt,name=errorCode" json:"errorCode,omitempty"`
	TotalCount int32   `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *GetDeviceTypesReply) Reset()                    { *m = GetDeviceTypesReply{} }
func (m *GetDeviceTypesReply) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceTypesReply) ProtoMessage()               {}
func (*GetDeviceTypesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetDeviceTypesReply) GetDeviceType() []*Type {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

func (m *GetDeviceTypesReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GetDeviceTypesReply) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Type)(nil), "api.Type")
	proto.RegisterType((*AddDeviceTypeRequest)(nil), "api.AddDeviceTypeRequest")
	proto.RegisterType((*AddDeviceTypeReply)(nil), "api.AddDeviceTypeReply")
	proto.RegisterType((*GetDeviceTypesRequest)(nil), "api.GetDeviceTypesRequest")
	proto.RegisterType((*GetDeviceTypesReply)(nil), "api.GetDeviceTypesReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeviceType service

type DeviceTypeClient interface {
	GetDeviceTypes(ctx context.Context, in *GetDeviceTypesRequest, opts ...grpc.CallOption) (*GetDeviceTypesReply, error)
	AddDeviceType(ctx context.Context, in *AddDeviceTypeRequest, opts ...grpc.CallOption) (*AddDeviceTypeReply, error)
}

type deviceTypeClient struct {
	cc *grpc.ClientConn
}

func NewDeviceTypeClient(cc *grpc.ClientConn) DeviceTypeClient {
	return &deviceTypeClient{cc}
}

func (c *deviceTypeClient) GetDeviceTypes(ctx context.Context, in *GetDeviceTypesRequest, opts ...grpc.CallOption) (*GetDeviceTypesReply, error) {
	out := new(GetDeviceTypesReply)
	err := grpc.Invoke(ctx, "/api.DeviceType/GetDeviceTypes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceTypeClient) AddDeviceType(ctx context.Context, in *AddDeviceTypeRequest, opts ...grpc.CallOption) (*AddDeviceTypeReply, error) {
	out := new(AddDeviceTypeReply)
	err := grpc.Invoke(ctx, "/api.DeviceType/AddDeviceType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceType service

type DeviceTypeServer interface {
	GetDeviceTypes(context.Context, *GetDeviceTypesRequest) (*GetDeviceTypesReply, error)
	AddDeviceType(context.Context, *AddDeviceTypeRequest) (*AddDeviceTypeReply, error)
}

func RegisterDeviceTypeServer(s *grpc.Server, srv DeviceTypeServer) {
	s.RegisterService(&_DeviceType_serviceDesc, srv)
}

func _DeviceType_GetDeviceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceTypeServer).GetDeviceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceType/GetDeviceTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceTypeServer).GetDeviceTypes(ctx, req.(*GetDeviceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceType_AddDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceTypeServer).AddDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceType/AddDeviceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceTypeServer).AddDeviceType(ctx, req.(*AddDeviceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceType_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceType",
	HandlerType: (*DeviceTypeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceTypes",
			Handler:    _DeviceType_GetDeviceTypes_Handler,
		},
		{
			MethodName: "AddDeviceType",
			Handler:    _DeviceType_AddDeviceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/device_type.proto",
}

func init() { proto.RegisterFile("api/device_type.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x4e, 0x02, 0x31,
	0x10, 0xb6, 0xbb, 0x80, 0xee, 0x10, 0x89, 0x19, 0x41, 0x17, 0x62, 0x94, 0xec, 0x09, 0x2f, 0x98,
	0xe0, 0x13, 0x20, 0x1a, 0x3d, 0x37, 0x5e, 0x0d, 0x29, 0xb4, 0x31, 0x4d, 0x08, 0xad, 0xdd, 0x62,
	0xd2, 0x8b, 0x6f, 0xe2, 0xbb, 0x1a, 0xba, 0x2b, 0xfb, 0x93, 0xd5, 0x5b, 0x67, 0xa6, 0xf3, 0xfd,
	0xb5, 0x30, 0x60, 0x5a, 0xde, 0x71, 0xf1, 0x29, 0xd7, 0x62, 0x69, 0x9d, 0x16, 0x53, 0x6d, 0x94,
	0x55, 0x18, 0x32, 0x2d, 0x13, 0x06, 0xad, 0x57, 0xa7, 0x05, 0xf6, 0x20, 0x90, 0x3c, 0x26, 0x63,
	0x32, 0x69, 0xd3, 0x40, 0x72, 0xbc, 0x81, 0x6e, 0x69, 0x23, 0x0e, 0xc6, 0x64, 0x12, 0x51, 0xc8,
	0x5a, 0x7e, 0xe1, 0x02, 0x3a, 0xa9, 0x65, 0x76, 0x97, 0xc6, 0xa1, 0x5f, 0xca, 0x2b, 0x3c, 0x83,
	0xd0, 0x4a, 0x1e, 0xb7, 0xc6, 0x64, 0x12, 0xd2, 0xfd, 0x31, 0x99, 0x43, 0x7f, 0xce, 0xf9, 0xe3,
	0x61, 0x95, 0x8a, 0x8f, 0x9d, 0x48, 0x2d, 0xde, 0x42, 0x09, 0xcf, 0x53, 0x77, 0x67, 0xd1, 0x94,
	0x69, 0x39, 0xf5, 0xb7, 0x4a, 0xc3, 0xe4, 0x0d, 0xb0, 0x06, 0xa1, 0x37, 0x0e, 0xaf, 0x20, 0x12,
	0xc6, 0x28, 0xb3, 0x50, 0x5c, 0xe4, 0xd2, 0x8b, 0x46, 0x0d, 0x3e, 0xf8, 0x0f, 0x7e, 0x0b, 0x83,
	0x67, 0x61, 0x0b, 0xf8, 0xf4, 0x57, 0x62, 0x6e, 0x86, 0x1c, 0xcc, 0x20, 0x42, 0x4b, 0xb3, 0xf7,
	0x0c, 0xaf, 0x4d, 0xfd, 0x19, 0xfb, 0xd0, 0x5e, 0xab, 0xdd, 0xd6, 0xe6, 0x49, 0x64, 0x05, 0x0e,
	0xe1, 0x44, 0x19, 0x2e, 0xcc, 0x72, 0xe5, 0x7c, 0x1a, 0x11, 0x3d, 0xf6, 0xf5, 0x83, 0x4b, 0xbe,
	0xe0, 0xbc, 0xce, 0xb7, 0xf7, 0x53, 0x0f, 0x24, 0xfc, 0x53, 0x71, 0xd5, 0x7a, 0x50, 0xb7, 0x7e,
	0x0d, 0x60, 0x95, 0x65, 0x9b, 0x45, 0x49, 0x55, 0xa9, 0x33, 0xfb, 0x26, 0x00, 0x05, 0x3b, 0xbe,
	0x40, 0xaf, 0x2a, 0x07, 0x47, 0x9e, 0xb5, 0x31, 0x93, 0x51, 0xdc, 0x38, 0xd3, 0x1b, 0x97, 0x1c,
	0xe1, 0x13, 0x9c, 0x56, 0xde, 0x09, 0x87, 0xfe, 0x72, 0xd3, 0xf3, 0x8f, 0x2e, 0x9b, 0x46, 0x1e,
	0x66, 0xd5, 0xf1, 0x1f, 0xf4, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x44, 0x43, 0x63, 0x0f, 0xb9,
	0x02, 0x00, 0x00,
}
