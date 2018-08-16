// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tenant.proto

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

type GetDidByTidRequest struct {
	Tid int64 `protobuf:"varint,1,opt,name=Tid" json:"Tid,omitempty"`
}

func (m *GetDidByTidRequest) Reset()                    { *m = GetDidByTidRequest{} }
func (m *GetDidByTidRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDidByTidRequest) ProtoMessage()               {}
func (*GetDidByTidRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetDidByTidRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type GetDidByTidResponse struct {
	Did int64 `protobuf:"varint,2,opt,name=Did" json:"Did,omitempty"`
}

func (m *GetDidByTidResponse) Reset()                    { *m = GetDidByTidResponse{} }
func (m *GetDidByTidResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDidByTidResponse) ProtoMessage()               {}
func (*GetDidByTidResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetDidByTidResponse) GetDid() int64 {
	if m != nil {
		return m.Did
	}
	return 0
}

func init() {
	proto.RegisterType((*GetDidByTidRequest)(nil), "api.GetDidByTidRequest")
	proto.RegisterType((*GetDidByTidResponse)(nil), "api.GetDidByTidResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TenantServer service

type TenantServerClient interface {
	GetDidByTid(ctx context.Context, in *GetDidByTidRequest, opts ...grpc.CallOption) (*GetDidByTidResponse, error)
}

type tenantServerClient struct {
	cc *grpc.ClientConn
}

func NewTenantServerClient(cc *grpc.ClientConn) TenantServerClient {
	return &tenantServerClient{cc}
}

func (c *tenantServerClient) GetDidByTid(ctx context.Context, in *GetDidByTidRequest, opts ...grpc.CallOption) (*GetDidByTidResponse, error) {
	out := new(GetDidByTidResponse)
	err := grpc.Invoke(ctx, "/api.TenantServer/GetDidByTid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TenantServer service

type TenantServerServer interface {
	GetDidByTid(context.Context, *GetDidByTidRequest) (*GetDidByTidResponse, error)
}

func RegisterTenantServerServer(s *grpc.Server, srv TenantServerServer) {
	s.RegisterService(&_TenantServer_serviceDesc, srv)
}

func _TenantServer_GetDidByTid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDidByTidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServerServer).GetDidByTid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TenantServer/GetDidByTid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServerServer).GetDidByTid(ctx, req.(*GetDidByTidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TenantServer",
	HandlerType: (*TenantServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDidByTid",
			Handler:    _TenantServer_GetDidByTid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant.proto",
}

func init() { proto.RegisterFile("tenant.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x49, 0xcd, 0x4b,
	0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52, 0xe3,
	0x12, 0x72, 0x4f, 0x2d, 0x71, 0xc9, 0x4c, 0x71, 0xaa, 0x0c, 0xc9, 0x4c, 0x09, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x0e, 0xc9, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0e, 0x02, 0x31, 0x95, 0xd4, 0xb9, 0x84, 0x51, 0xd4, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82,
	0x14, 0xba, 0x64, 0xa6, 0x48, 0x30, 0x41, 0x14, 0xba, 0x64, 0xa6, 0x18, 0x05, 0x71, 0xf1, 0x84,
	0x80, 0x6d, 0x09, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x12, 0x72, 0xe2, 0xe2, 0x46, 0xd2, 0x28, 0x24,
	0xae, 0x97, 0x58, 0x90, 0xa9, 0x87, 0x69, 0xa5, 0x94, 0x04, 0xa6, 0x04, 0xc4, 0x0e, 0x25, 0x86,
	0x24, 0x36, 0xb0, 0x83, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x90, 0x2d, 0x22, 0xc0,
	0x00, 0x00, 0x00,
}
