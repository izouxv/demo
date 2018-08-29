// Code generated by protoc-gen-go.
// source: api/service.proto
// DO NOT EDIT!

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

type Service struct {
	Sid                int32  `protobuf:"varint,1,opt,name=Sid" json:"Sid,omitempty"`
	ServiceName        string `protobuf:"bytes,2,opt,name=ServiceName" json:"ServiceName,omitempty"`
	ServiceKey         string `protobuf:"bytes,3,opt,name=ServiceKey" json:"ServiceKey,omitempty"`
	ServiceUrl         string `protobuf:"bytes,4,opt,name=ServiceUrl" json:"ServiceUrl,omitempty"`
	ServiceType        int32  `protobuf:"varint,5,opt,name=ServiceType" json:"ServiceType,omitempty"`
	ServiceTid         int32  `protobuf:"varint,6,opt,name=ServiceTid" json:"ServiceTid,omitempty"`
	CreateTime         int64  `protobuf:"varint,7,opt,name=CreateTime" json:"CreateTime,omitempty"`
	UpdateTime         int64  `protobuf:"varint,8,opt,name=UpdateTime" json:"UpdateTime,omitempty"`
	ServiceDescription string `protobuf:"bytes,9,opt,name=ServiceDescription" json:"ServiceDescription,omitempty"`
	ServiceState       int32  `protobuf:"varint,10,opt,name=ServiceState" json:"ServiceState,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Service) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *Service) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Service) GetServiceKey() string {
	if m != nil {
		return m.ServiceKey
	}
	return ""
}

func (m *Service) GetServiceUrl() string {
	if m != nil {
		return m.ServiceUrl
	}
	return ""
}

func (m *Service) GetServiceType() int32 {
	if m != nil {
		return m.ServiceType
	}
	return 0
}

func (m *Service) GetServiceTid() int32 {
	if m != nil {
		return m.ServiceTid
	}
	return 0
}

func (m *Service) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Service) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Service) GetServiceDescription() string {
	if m != nil {
		return m.ServiceDescription
	}
	return ""
}

func (m *Service) GetServiceState() int32 {
	if m != nil {
		return m.ServiceState
	}
	return 0
}

type AddServiceRequest struct {
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *AddServiceRequest) Reset()                    { *m = AddServiceRequest{} }
func (m *AddServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*AddServiceRequest) ProtoMessage()               {}
func (*AddServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *AddServiceRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type AddServiceResponse struct {
}

func (m *AddServiceResponse) Reset()                    { *m = AddServiceResponse{} }
func (m *AddServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*AddServiceResponse) ProtoMessage()               {}
func (*AddServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

type DeleteServiceRequest struct {
	Sid int32 `protobuf:"varint,1,opt,name=Sid" json:"Sid,omitempty"`
}

func (m *DeleteServiceRequest) Reset()                    { *m = DeleteServiceRequest{} }
func (m *DeleteServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteServiceRequest) ProtoMessage()               {}
func (*DeleteServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *DeleteServiceRequest) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

type DeleteServiceResponse struct {
}

func (m *DeleteServiceResponse) Reset()                    { *m = DeleteServiceResponse{} }
func (m *DeleteServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteServiceResponse) ProtoMessage()               {}
func (*DeleteServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

type UpdateServiceRequest struct {
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *UpdateServiceRequest) Reset()                    { *m = UpdateServiceRequest{} }
func (m *UpdateServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateServiceRequest) ProtoMessage()               {}
func (*UpdateServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *UpdateServiceRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type UpdateServiceResponse struct {
}

func (m *UpdateServiceResponse) Reset()                    { *m = UpdateServiceResponse{} }
func (m *UpdateServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateServiceResponse) ProtoMessage()               {}
func (*UpdateServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

type GetServiceBySidRequest struct {
	Sid int32 `protobuf:"varint,1,opt,name=Sid" json:"Sid,omitempty"`
}

func (m *GetServiceBySidRequest) Reset()                    { *m = GetServiceBySidRequest{} }
func (m *GetServiceBySidRequest) String() string            { return proto.CompactTextString(m) }
func (*GetServiceBySidRequest) ProtoMessage()               {}
func (*GetServiceBySidRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *GetServiceBySidRequest) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

type GetServiceBySidResponse struct {
	Service *Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *GetServiceBySidResponse) Reset()                    { *m = GetServiceBySidResponse{} }
func (m *GetServiceBySidResponse) String() string            { return proto.CompactTextString(m) }
func (*GetServiceBySidResponse) ProtoMessage()               {}
func (*GetServiceBySidResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *GetServiceBySidResponse) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type GetServiceByTidRequest struct {
	Tid int32 `protobuf:"varint,1,opt,name=Tid" json:"Tid,omitempty"`
}

func (m *GetServiceByTidRequest) Reset()                    { *m = GetServiceByTidRequest{} }
func (m *GetServiceByTidRequest) String() string            { return proto.CompactTextString(m) }
func (*GetServiceByTidRequest) ProtoMessage()               {}
func (*GetServiceByTidRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *GetServiceByTidRequest) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type GetServiceByTidResponse struct {
	Service    []*Service `protobuf:"bytes,1,rep,name=service" json:"service,omitempty"`
	TotalCount int32      `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *GetServiceByTidResponse) Reset()                    { *m = GetServiceByTidResponse{} }
func (m *GetServiceByTidResponse) String() string            { return proto.CompactTextString(m) }
func (*GetServiceByTidResponse) ProtoMessage()               {}
func (*GetServiceByTidResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *GetServiceByTidResponse) GetService() []*Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *GetServiceByTidResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Service)(nil), "api.Service")
	proto.RegisterType((*AddServiceRequest)(nil), "api.AddServiceRequest")
	proto.RegisterType((*AddServiceResponse)(nil), "api.AddServiceResponse")
	proto.RegisterType((*DeleteServiceRequest)(nil), "api.DeleteServiceRequest")
	proto.RegisterType((*DeleteServiceResponse)(nil), "api.DeleteServiceResponse")
	proto.RegisterType((*UpdateServiceRequest)(nil), "api.UpdateServiceRequest")
	proto.RegisterType((*UpdateServiceResponse)(nil), "api.UpdateServiceResponse")
	proto.RegisterType((*GetServiceBySidRequest)(nil), "api.GetServiceBySidRequest")
	proto.RegisterType((*GetServiceBySidResponse)(nil), "api.GetServiceBySidResponse")
	proto.RegisterType((*GetServiceByTidRequest)(nil), "api.GetServiceByTidRequest")
	proto.RegisterType((*GetServiceByTidResponse)(nil), "api.GetServiceByTidResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServiceServer service

type ServiceServerClient interface {
	AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error)
	GetServiceBySid(ctx context.Context, in *GetServiceBySidRequest, opts ...grpc.CallOption) (*GetServiceBySidResponse, error)
	GetServiceByTid(ctx context.Context, in *GetServiceByTidRequest, opts ...grpc.CallOption) (*GetServiceByTidResponse, error)
}

type serviceServerClient struct {
	cc *grpc.ClientConn
}

func NewServiceServerClient(cc *grpc.ClientConn) ServiceServerClient {
	return &serviceServerClient{cc}
}

func (c *serviceServerClient) AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error) {
	out := new(AddServiceResponse)
	err := grpc.Invoke(ctx, "/api.ServiceServer/AddService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServerClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	out := new(DeleteServiceResponse)
	err := grpc.Invoke(ctx, "/api.ServiceServer/DeleteService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServerClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error) {
	out := new(UpdateServiceResponse)
	err := grpc.Invoke(ctx, "/api.ServiceServer/UpdateService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServerClient) GetServiceBySid(ctx context.Context, in *GetServiceBySidRequest, opts ...grpc.CallOption) (*GetServiceBySidResponse, error) {
	out := new(GetServiceBySidResponse)
	err := grpc.Invoke(ctx, "/api.ServiceServer/GetServiceBySid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServerClient) GetServiceByTid(ctx context.Context, in *GetServiceByTidRequest, opts ...grpc.CallOption) (*GetServiceByTidResponse, error) {
	out := new(GetServiceByTidResponse)
	err := grpc.Invoke(ctx, "/api.ServiceServer/GetServiceByTid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceServer service

type ServiceServerServer interface {
	AddService(context.Context, *AddServiceRequest) (*AddServiceResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error)
	GetServiceBySid(context.Context, *GetServiceBySidRequest) (*GetServiceBySidResponse, error)
	GetServiceByTid(context.Context, *GetServiceByTidRequest) (*GetServiceByTidResponse, error)
}

func RegisterServiceServerServer(s *grpc.Server, srv ServiceServerServer) {
	s.RegisterService(&_ServiceServer_serviceDesc, srv)
}

func _ServiceServer_AddService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServerServer).AddService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceServer/AddService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServerServer).AddService(ctx, req.(*AddServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceServer_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServerServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceServer/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServerServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceServer_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServerServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceServer/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServerServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceServer_GetServiceBySid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceBySidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServerServer).GetServiceBySid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceServer/GetServiceBySid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServerServer).GetServiceBySid(ctx, req.(*GetServiceBySidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceServer_GetServiceByTid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceByTidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServerServer).GetServiceByTid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceServer/GetServiceByTid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServerServer).GetServiceByTid(ctx, req.(*GetServiceByTidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ServiceServer",
	HandlerType: (*ServiceServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddService",
			Handler:    _ServiceServer_AddService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _ServiceServer_DeleteService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _ServiceServer_UpdateService_Handler,
		},
		{
			MethodName: "GetServiceBySid",
			Handler:    _ServiceServer_GetServiceBySid_Handler,
		},
		{
			MethodName: "GetServiceByTid",
			Handler:    _ServiceServer_GetServiceByTid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}

func init() { proto.RegisterFile("api/service.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x51, 0x6f, 0xd3, 0x30,
	0x14, 0x85, 0xd7, 0x85, 0xae, 0xec, 0x6e, 0x13, 0xcc, 0xea, 0x56, 0x13, 0x10, 0x8a, 0xf2, 0x80,
	0x22, 0x1e, 0x82, 0x34, 0x1e, 0x91, 0x40, 0x63, 0x93, 0x40, 0x42, 0xda, 0x43, 0xe2, 0xfe, 0x00,
	0xb3, 0xdc, 0x07, 0x4b, 0x59, 0x63, 0x1c, 0xaf, 0x52, 0xdf, 0xf9, 0x87, 0xfc, 0x21, 0x14, 0xc7,
	0x51, 0x9d, 0xc4, 0x41, 0x68, 0x6f, 0xed, 0x39, 0xd7, 0xdf, 0xb9, 0xcd, 0x71, 0x03, 0xe7, 0x5c,
	0x8a, 0x0f, 0x35, 0xaa, 0xad, 0xb8, 0xc7, 0x54, 0xaa, 0x4a, 0x57, 0x24, 0xe0, 0x52, 0xc4, 0x7f,
	0x0e, 0x61, 0x91, 0xb7, 0x32, 0x79, 0x09, 0x41, 0x2e, 0x0a, 0x3a, 0x8b, 0x66, 0xc9, 0x3c, 0x6b,
	0x3e, 0x92, 0x08, 0x4e, 0xac, 0x79, 0xc7, 0x1f, 0x90, 0x1e, 0x46, 0xb3, 0xe4, 0x38, 0x73, 0x25,
	0xf2, 0x16, 0xc0, 0x7e, 0xfd, 0x81, 0x3b, 0x1a, 0x98, 0x01, 0x47, 0x71, 0xfc, 0xb5, 0x2a, 0xe9,
	0xb3, 0x9e, 0xbf, 0x56, 0xa5, 0x93, 0xc0, 0x76, 0x12, 0xe9, 0xdc, 0x64, 0xbb, 0x92, 0x43, 0x60,
	0xa2, 0xa0, 0x47, 0x66, 0xc0, 0x51, 0x1a, 0xff, 0x46, 0x21, 0xd7, 0xc8, 0xc4, 0x03, 0xd2, 0x45,
	0x34, 0x4b, 0x82, 0xcc, 0x51, 0x1a, 0x7f, 0x2d, 0x8b, 0xce, 0x7f, 0xde, 0xfa, 0x7b, 0x85, 0xa4,
	0x40, 0x2c, 0xed, 0x16, 0xeb, 0x7b, 0x25, 0xa4, 0x16, 0xd5, 0x86, 0x1e, 0x9b, 0x4d, 0x3d, 0x0e,
	0x89, 0xe1, 0xd4, 0xaa, 0xb9, 0xe6, 0x1a, 0x29, 0x98, 0x8d, 0x7a, 0x5a, 0xfc, 0x09, 0xce, 0xaf,
	0x8b, 0xc2, 0x4a, 0x19, 0xfe, 0x7a, 0xc4, 0x5a, 0x93, 0x77, 0xb0, 0xb0, 0x05, 0x98, 0x47, 0x7c,
	0x72, 0x75, 0x9a, 0x72, 0x29, 0xd2, 0x6e, 0xaa, 0x33, 0xe3, 0x25, 0x10, 0xf7, 0x70, 0x2d, 0xab,
	0x4d, 0x8d, 0x71, 0x02, 0xcb, 0x5b, 0x2c, 0x51, 0xe3, 0x80, 0x3a, 0x2a, 0x2d, 0x5e, 0xc1, 0xc5,
	0x60, 0xd2, 0x22, 0x3e, 0xc3, 0xb2, 0xfd, 0xdd, 0x4f, 0x5c, 0x6c, 0x05, 0x17, 0x83, 0xf3, 0x16,
	0xfc, 0x1e, 0x2e, 0xbf, 0xa1, 0xb6, 0xea, 0xd7, 0x5d, 0x2e, 0x8a, 0xe9, 0xed, 0xae, 0x61, 0x35,
	0x9a, 0x6d, 0x31, 0xff, 0xbd, 0xc7, 0x20, 0x8e, 0xf5, 0xe2, 0xd8, 0x3e, 0x8e, 0x89, 0x22, 0xe6,
	0xfd, 0x38, 0x36, 0x15, 0x17, 0x4c, 0xc6, 0x35, 0x17, 0x48, 0x57, 0x9a, 0x97, 0x37, 0xd5, 0xe3,
	0x46, 0x9b, 0xff, 0xc0, 0x3c, 0x73, 0x94, 0xab, 0xdf, 0x01, 0x9c, 0x75, 0xed, 0xa3, 0xda, 0xa2,
	0x22, 0x5f, 0x00, 0xf6, 0x0d, 0x92, 0x4b, 0x83, 0x1d, 0xdd, 0x87, 0x70, 0x35, 0xd2, 0xed, 0xe3,
	0x3c, 0x20, 0xdf, 0xe1, 0xac, 0x57, 0x21, 0x79, 0x65, 0x66, 0x7d, 0x17, 0x20, 0x0c, 0x7d, 0x96,
	0x4b, 0xea, 0x75, 0x66, 0x49, 0xbe, 0x7b, 0x60, 0x49, 0xfe, 0x8a, 0x0f, 0xc8, 0x1d, 0xbc, 0x18,
	0x14, 0x47, 0x5e, 0x9b, 0x03, 0xfe, 0xea, 0xc3, 0x37, 0x7e, 0x73, 0x8a, 0xc7, 0xbc, 0x3c, 0xf6,
	0x2f, 0x1e, 0x73, 0x79, 0x3f, 0x8f, 0xcc, 0x5b, 0xed, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x54, 0x9b, 0x6e, 0x5f, 0xea, 0x04, 0x00, 0x00,
}