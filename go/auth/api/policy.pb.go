// Code generated by protoc-gen-go.
// source: api/policy.proto
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

type Policy struct {
	Pid             int32   `protobuf:"varint,1,opt,name=Pid" json:"Pid,omitempty"`
	PolicyName      string  `protobuf:"bytes,2,opt,name=PolicyName" json:"PolicyName,omitempty"`
	PolicyType      int32   `protobuf:"varint,3,opt,name=PolicyType" json:"PolicyType,omitempty"`
	PolicyCycle     int32   `protobuf:"varint,4,opt,name=PolicyCycle" json:"PolicyCycle,omitempty"`
	PolicyFeeType   int32   `protobuf:"varint,5,opt,name=PolicyFeeType" json:"PolicyFeeType,omitempty"`
	PolicyUnitPrice float32 `protobuf:"fixed32,6,opt,name=PolicyUnitPrice" json:"PolicyUnitPrice,omitempty"`
	PolicyUnitType  int32   `protobuf:"varint,7,opt,name=PolicyUnitType" json:"PolicyUnitType,omitempty"`
	PolicyUnitCount int32   `protobuf:"varint,8,opt,name=PolicyUnitCount" json:"PolicyUnitCount,omitempty"`
	PolicySid       int32   `protobuf:"varint,9,opt,name=PolicySid" json:"PolicySid,omitempty"`
	CreateTime      int64   `protobuf:"varint,10,opt,name=CreateTime" json:"CreateTime,omitempty"`
	UpdateTime      int64   `protobuf:"varint,11,opt,name=UpdateTime" json:"UpdateTime,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Policy) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Policy) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Policy) GetPolicyType() int32 {
	if m != nil {
		return m.PolicyType
	}
	return 0
}

func (m *Policy) GetPolicyCycle() int32 {
	if m != nil {
		return m.PolicyCycle
	}
	return 0
}

func (m *Policy) GetPolicyFeeType() int32 {
	if m != nil {
		return m.PolicyFeeType
	}
	return 0
}

func (m *Policy) GetPolicyUnitPrice() float32 {
	if m != nil {
		return m.PolicyUnitPrice
	}
	return 0
}

func (m *Policy) GetPolicyUnitType() int32 {
	if m != nil {
		return m.PolicyUnitType
	}
	return 0
}

func (m *Policy) GetPolicyUnitCount() int32 {
	if m != nil {
		return m.PolicyUnitCount
	}
	return 0
}

func (m *Policy) GetPolicySid() int32 {
	if m != nil {
		return m.PolicySid
	}
	return 0
}

func (m *Policy) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Policy) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type AddPolicyRequest struct {
	Policy *Policy `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
}

func (m *AddPolicyRequest) Reset()                    { *m = AddPolicyRequest{} }
func (m *AddPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPolicyRequest) ProtoMessage()               {}
func (*AddPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *AddPolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type AddPolicyResponse struct {
}

func (m *AddPolicyResponse) Reset()                    { *m = AddPolicyResponse{} }
func (m *AddPolicyResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPolicyResponse) ProtoMessage()               {}
func (*AddPolicyResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

type DeletePolicyByPidRequest struct {
	Pid int32 `protobuf:"varint,1,opt,name=Pid" json:"Pid,omitempty"`
}

func (m *DeletePolicyByPidRequest) Reset()                    { *m = DeletePolicyByPidRequest{} }
func (m *DeletePolicyByPidRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePolicyByPidRequest) ProtoMessage()               {}
func (*DeletePolicyByPidRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DeletePolicyByPidRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type DeletePolicyByPidResponse struct {
}

func (m *DeletePolicyByPidResponse) Reset()                    { *m = DeletePolicyByPidResponse{} }
func (m *DeletePolicyByPidResponse) String() string            { return proto.CompactTextString(m) }
func (*DeletePolicyByPidResponse) ProtoMessage()               {}
func (*DeletePolicyByPidResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

type DeletePolicyBySidRequest struct {
	PolicySid int32 `protobuf:"varint,1,opt,name=PolicySid" json:"PolicySid,omitempty"`
}

func (m *DeletePolicyBySidRequest) Reset()                    { *m = DeletePolicyBySidRequest{} }
func (m *DeletePolicyBySidRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePolicyBySidRequest) ProtoMessage()               {}
func (*DeletePolicyBySidRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *DeletePolicyBySidRequest) GetPolicySid() int32 {
	if m != nil {
		return m.PolicySid
	}
	return 0
}

type DeletePolicyBySidResponse struct {
}

func (m *DeletePolicyBySidResponse) Reset()                    { *m = DeletePolicyBySidResponse{} }
func (m *DeletePolicyBySidResponse) String() string            { return proto.CompactTextString(m) }
func (*DeletePolicyBySidResponse) ProtoMessage()               {}
func (*DeletePolicyBySidResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

type UpdatePolicyRequest struct {
	Policy *Policy `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
}

func (m *UpdatePolicyRequest) Reset()                    { *m = UpdatePolicyRequest{} }
func (m *UpdatePolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePolicyRequest) ProtoMessage()               {}
func (*UpdatePolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *UpdatePolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type UpdatePolicyResponse struct {
}

func (m *UpdatePolicyResponse) Reset()                    { *m = UpdatePolicyResponse{} }
func (m *UpdatePolicyResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdatePolicyResponse) ProtoMessage()               {}
func (*UpdatePolicyResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

type GetPolicyByPidRequest struct {
	Pid int32 `protobuf:"varint,1,opt,name=Pid" json:"Pid,omitempty"`
}

func (m *GetPolicyByPidRequest) Reset()                    { *m = GetPolicyByPidRequest{} }
func (m *GetPolicyByPidRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPolicyByPidRequest) ProtoMessage()               {}
func (*GetPolicyByPidRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *GetPolicyByPidRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type GetPolicyByPidResponse struct {
	Policy *Policy `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
}

func (m *GetPolicyByPidResponse) Reset()                    { *m = GetPolicyByPidResponse{} }
func (m *GetPolicyByPidResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPolicyByPidResponse) ProtoMessage()               {}
func (*GetPolicyByPidResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

func (m *GetPolicyByPidResponse) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type GetPolicyBySidRequest struct {
	Sid int32 `protobuf:"varint,1,opt,name=Sid" json:"Sid,omitempty"`
}

func (m *GetPolicyBySidRequest) Reset()                    { *m = GetPolicyBySidRequest{} }
func (m *GetPolicyBySidRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPolicyBySidRequest) ProtoMessage()               {}
func (*GetPolicyBySidRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{11} }

func (m *GetPolicyBySidRequest) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

type GetPolicyBySidResponse struct {
	Policy     []*Policy `protobuf:"bytes,1,rep,name=policy" json:"policy,omitempty"`
	TotalCount int32     `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *GetPolicyBySidResponse) Reset()                    { *m = GetPolicyBySidResponse{} }
func (m *GetPolicyBySidResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPolicyBySidResponse) ProtoMessage()               {}
func (*GetPolicyBySidResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{12} }

func (m *GetPolicyBySidResponse) GetPolicy() []*Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *GetPolicyBySidResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Policy)(nil), "api.Policy")
	proto.RegisterType((*AddPolicyRequest)(nil), "api.AddPolicyRequest")
	proto.RegisterType((*AddPolicyResponse)(nil), "api.AddPolicyResponse")
	proto.RegisterType((*DeletePolicyByPidRequest)(nil), "api.DeletePolicyByPidRequest")
	proto.RegisterType((*DeletePolicyByPidResponse)(nil), "api.DeletePolicyByPidResponse")
	proto.RegisterType((*DeletePolicyBySidRequest)(nil), "api.DeletePolicyBySidRequest")
	proto.RegisterType((*DeletePolicyBySidResponse)(nil), "api.DeletePolicyBySidResponse")
	proto.RegisterType((*UpdatePolicyRequest)(nil), "api.UpdatePolicyRequest")
	proto.RegisterType((*UpdatePolicyResponse)(nil), "api.UpdatePolicyResponse")
	proto.RegisterType((*GetPolicyByPidRequest)(nil), "api.GetPolicyByPidRequest")
	proto.RegisterType((*GetPolicyByPidResponse)(nil), "api.GetPolicyByPidResponse")
	proto.RegisterType((*GetPolicyBySidRequest)(nil), "api.GetPolicyBySidRequest")
	proto.RegisterType((*GetPolicyBySidResponse)(nil), "api.GetPolicyBySidResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PolicyServer service

type PolicyServerClient interface {
	AddPolicy(ctx context.Context, in *AddPolicyRequest, opts ...grpc.CallOption) (*AddPolicyResponse, error)
	DeletePolicyByPid(ctx context.Context, in *DeletePolicyByPidRequest, opts ...grpc.CallOption) (*DeletePolicyByPidResponse, error)
	DeletePolicyBySid(ctx context.Context, in *DeletePolicyBySidRequest, opts ...grpc.CallOption) (*DeletePolicyBySidResponse, error)
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error)
	GetPolicyByPid(ctx context.Context, in *GetPolicyByPidRequest, opts ...grpc.CallOption) (*GetPolicyByPidResponse, error)
	GetPolicyBySid(ctx context.Context, in *GetPolicyBySidRequest, opts ...grpc.CallOption) (*GetPolicyBySidResponse, error)
}

type policyServerClient struct {
	cc *grpc.ClientConn
}

func NewPolicyServerClient(cc *grpc.ClientConn) PolicyServerClient {
	return &policyServerClient{cc}
}

func (c *policyServerClient) AddPolicy(ctx context.Context, in *AddPolicyRequest, opts ...grpc.CallOption) (*AddPolicyResponse, error) {
	out := new(AddPolicyResponse)
	err := grpc.Invoke(ctx, "/api.PolicyServer/AddPolicy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServerClient) DeletePolicyByPid(ctx context.Context, in *DeletePolicyByPidRequest, opts ...grpc.CallOption) (*DeletePolicyByPidResponse, error) {
	out := new(DeletePolicyByPidResponse)
	err := grpc.Invoke(ctx, "/api.PolicyServer/DeletePolicyByPid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServerClient) DeletePolicyBySid(ctx context.Context, in *DeletePolicyBySidRequest, opts ...grpc.CallOption) (*DeletePolicyBySidResponse, error) {
	out := new(DeletePolicyBySidResponse)
	err := grpc.Invoke(ctx, "/api.PolicyServer/DeletePolicyBySid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServerClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error) {
	out := new(UpdatePolicyResponse)
	err := grpc.Invoke(ctx, "/api.PolicyServer/UpdatePolicy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServerClient) GetPolicyByPid(ctx context.Context, in *GetPolicyByPidRequest, opts ...grpc.CallOption) (*GetPolicyByPidResponse, error) {
	out := new(GetPolicyByPidResponse)
	err := grpc.Invoke(ctx, "/api.PolicyServer/GetPolicyByPid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServerClient) GetPolicyBySid(ctx context.Context, in *GetPolicyBySidRequest, opts ...grpc.CallOption) (*GetPolicyBySidResponse, error) {
	out := new(GetPolicyBySidResponse)
	err := grpc.Invoke(ctx, "/api.PolicyServer/GetPolicyBySid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PolicyServer service

type PolicyServerServer interface {
	AddPolicy(context.Context, *AddPolicyRequest) (*AddPolicyResponse, error)
	DeletePolicyByPid(context.Context, *DeletePolicyByPidRequest) (*DeletePolicyByPidResponse, error)
	DeletePolicyBySid(context.Context, *DeletePolicyBySidRequest) (*DeletePolicyBySidResponse, error)
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error)
	GetPolicyByPid(context.Context, *GetPolicyByPidRequest) (*GetPolicyByPidResponse, error)
	GetPolicyBySid(context.Context, *GetPolicyBySidRequest) (*GetPolicyBySidResponse, error)
}

func RegisterPolicyServerServer(s *grpc.Server, srv PolicyServerServer) {
	s.RegisterService(&_PolicyServer_serviceDesc, srv)
}

func _PolicyServer_AddPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServerServer).AddPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PolicyServer/AddPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServerServer).AddPolicy(ctx, req.(*AddPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyServer_DeletePolicyByPid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyByPidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServerServer).DeletePolicyByPid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PolicyServer/DeletePolicyByPid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServerServer).DeletePolicyByPid(ctx, req.(*DeletePolicyByPidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyServer_DeletePolicyBySid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyBySidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServerServer).DeletePolicyBySid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PolicyServer/DeletePolicyBySid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServerServer).DeletePolicyBySid(ctx, req.(*DeletePolicyBySidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyServer_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServerServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PolicyServer/UpdatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServerServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyServer_GetPolicyByPid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyByPidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServerServer).GetPolicyByPid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PolicyServer/GetPolicyByPid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServerServer).GetPolicyByPid(ctx, req.(*GetPolicyByPidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyServer_GetPolicyBySid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyBySidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServerServer).GetPolicyBySid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PolicyServer/GetPolicyBySid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServerServer).GetPolicyBySid(ctx, req.(*GetPolicyBySidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolicyServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PolicyServer",
	HandlerType: (*PolicyServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPolicy",
			Handler:    _PolicyServer_AddPolicy_Handler,
		},
		{
			MethodName: "DeletePolicyByPid",
			Handler:    _PolicyServer_DeletePolicyByPid_Handler,
		},
		{
			MethodName: "DeletePolicyBySid",
			Handler:    _PolicyServer_DeletePolicyBySid_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _PolicyServer_UpdatePolicy_Handler,
		},
		{
			MethodName: "GetPolicyByPid",
			Handler:    _PolicyServer_GetPolicyByPid_Handler,
		},
		{
			MethodName: "GetPolicyBySid",
			Handler:    _PolicyServer_GetPolicyBySid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/policy.proto",
}

func init() { proto.RegisterFile("api/policy.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x49, 0xc3, 0x0a, 0x7d, 0x33, 0x46, 0xe7, 0xb1, 0xca, 0xcb, 0x60, 0x8a, 0x02, 0x42,
	0x41, 0x42, 0x45, 0x1a, 0x07, 0x10, 0x82, 0x03, 0x94, 0x3f, 0x07, 0x24, 0x14, 0xc5, 0xdd, 0x91,
	0x43, 0x68, 0xde, 0x83, 0xa5, 0xac, 0x31, 0xa9, 0x87, 0x94, 0x8f, 0xca, 0x77, 0xe1, 0x80, 0x62,
	0x67, 0x4b, 0xec, 0x26, 0xd3, 0xb4, 0x5b, 0xf2, 0x3c, 0x8f, 0x7f, 0xae, 0x5f, 0x3f, 0x0d, 0x4c,
	0x53, 0xc1, 0x5f, 0x89, 0x22, 0xe7, 0xab, 0x6a, 0x2e, 0xca, 0x42, 0x16, 0xc4, 0x4d, 0x05, 0x0f,
	0xff, 0x8d, 0x60, 0x1c, 0x2b, 0x95, 0x4c, 0xc1, 0x8d, 0x79, 0x46, 0x9d, 0xc0, 0x89, 0x76, 0x92,
	0xfa, 0x91, 0x9c, 0x00, 0x68, 0xef, 0x47, 0x7a, 0x8e, 0x74, 0x14, 0x38, 0xd1, 0x24, 0xe9, 0x28,
	0xad, 0xbf, 0xac, 0x04, 0x52, 0x57, 0x2d, 0xec, 0x28, 0x24, 0x00, 0x4f, 0xbf, 0x2d, 0xaa, 0x55,
	0x8e, 0xf4, 0xae, 0x0a, 0x74, 0x25, 0xf2, 0x0c, 0x1e, 0xe8, 0xd7, 0xaf, 0x88, 0x0a, 0xb2, 0xa3,
	0x32, 0xa6, 0x48, 0x22, 0x78, 0xa8, 0x85, 0xb3, 0x35, 0x97, 0x71, 0xc9, 0x57, 0x48, 0xc7, 0x81,
	0x13, 0x8d, 0x12, 0x5b, 0x26, 0xcf, 0x61, 0xaf, 0x95, 0x14, 0xf0, 0x9e, 0x02, 0x5a, 0xaa, 0x49,
	0x5c, 0x14, 0x17, 0x6b, 0x49, 0xef, 0xab, 0xa0, 0x2d, 0x93, 0xc7, 0x30, 0xd1, 0x12, 0xe3, 0x19,
	0x9d, 0xa8, 0x4c, 0x2b, 0xd4, 0x13, 0x58, 0x94, 0x98, 0x4a, 0x5c, 0xf2, 0x73, 0xa4, 0x10, 0x38,
	0x91, 0x9b, 0x74, 0x94, 0xda, 0x3f, 0x13, 0xd9, 0xa5, 0xef, 0x69, 0xbf, 0x55, 0xc2, 0x37, 0x30,
	0xfd, 0x98, 0x65, 0x9a, 0x97, 0xe0, 0xef, 0x0b, 0xdc, 0x48, 0xf2, 0x14, 0xc6, 0xfa, 0x9e, 0xd4,
	0x55, 0x78, 0xa7, 0xde, 0x3c, 0x15, 0x7c, 0xde, 0x64, 0x1a, 0x2b, 0x3c, 0x80, 0xfd, 0xce, 0xc2,
	0x8d, 0x28, 0xd6, 0x1b, 0x0c, 0x5f, 0x02, 0xfd, 0x8c, 0x39, 0x4a, 0xd4, 0xfa, 0xa7, 0x2a, 0xe6,
	0xd9, 0x25, 0x75, 0xeb, 0x76, 0xc3, 0x63, 0x38, 0xea, 0x49, 0x37, 0xa8, 0xb7, 0x36, 0x8a, 0xb5,
	0x28, 0x63, 0x24, 0x8e, 0x35, 0x92, 0x6d, 0x2c, 0xeb, 0x60, 0xdf, 0xc1, 0x81, 0x3e, 0xfd, 0x2d,
	0x8e, 0x3c, 0x83, 0x47, 0xe6, 0xda, 0x86, 0xf9, 0x02, 0x0e, 0xbf, 0xa1, 0xbc, 0xd1, 0x91, 0x3f,
	0xc0, 0xcc, 0x8e, 0x6a, 0xc8, 0xcd, 0x7e, 0x81, 0xb9, 0x13, 0x33, 0x76, 0x6a, 0x67, 0x51, 0x3f,
	0x86, 0x3f, 0x8d, 0x9d, 0xd8, 0xc0, 0x4e, 0xee, 0xc0, 0x4e, 0x75, 0x6f, 0x64, 0x21, 0xd3, 0x5c,
	0x57, 0x73, 0xa4, 0xff, 0x59, 0xad, 0x72, 0xfa, 0xd7, 0x85, 0xdd, 0x66, 0xe4, 0x58, 0xfe, 0xc1,
	0x92, 0xbc, 0x87, 0xc9, 0x55, 0x1f, 0xc8, 0xa1, 0x42, 0xda, 0xc5, 0xf2, 0x67, 0xb6, 0xdc, 0x0c,
	0xf0, 0x0e, 0x59, 0xc2, 0xfe, 0x56, 0x15, 0xc8, 0x13, 0x15, 0x1f, 0x2a, 0x94, 0x7f, 0x32, 0x64,
	0x0f, 0x53, 0xd9, 0x00, 0x95, 0x5d, 0x4f, 0x65, 0x06, 0xf5, 0x0b, 0xec, 0x76, 0x6b, 0x40, 0xa8,
	0x5a, 0xd1, 0xd3, 0x2a, 0xff, 0xa8, 0xc7, 0xb9, 0xc2, 0x7c, 0x87, 0x3d, 0xb3, 0x0a, 0xc4, 0x57,
	0xf1, 0xde, 0x2a, 0xf9, 0xc7, 0xbd, 0xde, 0x00, 0x8c, 0xf5, 0xc1, 0xd8, 0x35, 0x30, 0xe3, 0x80,
	0xbf, 0xc6, 0xea, 0xf3, 0xfc, 0xfa, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x4a, 0x38, 0xfd,
	0xb2, 0x05, 0x00, 0x00,
}
