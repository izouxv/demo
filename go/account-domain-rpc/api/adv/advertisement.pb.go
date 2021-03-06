// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/adv/advertisement.proto

/*
Package adv is a generated protocol buffer package.

It is generated from these files:
	api/adv/advertisement.proto

It has these top-level messages:
	AdvertisementRequest
	AdvertisementReply
	MapAdvertisementReply
*/
package adv

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

type AdvertisementRequest struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Md5        string `protobuf:"bytes,2,opt,name=md5" json:"md5,omitempty"`
	StartTime  int64  `protobuf:"varint,3,opt,name=startTime" json:"startTime,omitempty"`
	EndTime    int64  `protobuf:"varint,4,opt,name=endTime" json:"endTime,omitempty"`
	Source     string `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	FileName   string `protobuf:"bytes,6,opt,name=fileName" json:"fileName,omitempty"`
	FileUrl    string `protobuf:"bytes,7,opt,name=fileUrl" json:"fileUrl,omitempty"`
	State      int32  `protobuf:"varint,8,opt,name=state" json:"state,omitempty"`
	Advertiser string `protobuf:"bytes,9,opt,name=advertiser" json:"advertiser,omitempty"`
	Id         int32  `protobuf:"varint,10,opt,name=id" json:"id,omitempty"`
	Signature  string `protobuf:"bytes,11,opt,name=signature" json:"signature,omitempty"`
	AdvUrl     string `protobuf:"bytes,12,opt,name=advUrl" json:"advUrl,omitempty"`
	Page       int32  `protobuf:"varint,13,opt,name=page" json:"page,omitempty"`
	Count      int32  `protobuf:"varint,14,opt,name=count" json:"count,omitempty"`
	Tid        int64  `protobuf:"varint,15,opt,name=tid" json:"tid,omitempty"`
}

func (m *AdvertisementRequest) Reset()                    { *m = AdvertisementRequest{} }
func (m *AdvertisementRequest) String() string            { return proto.CompactTextString(m) }
func (*AdvertisementRequest) ProtoMessage()               {}
func (*AdvertisementRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AdvertisementRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AdvertisementRequest) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *AdvertisementRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AdvertisementRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *AdvertisementRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AdvertisementRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *AdvertisementRequest) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *AdvertisementRequest) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *AdvertisementRequest) GetAdvertiser() string {
	if m != nil {
		return m.Advertiser
	}
	return ""
}

func (m *AdvertisementRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AdvertisementRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AdvertisementRequest) GetAdvUrl() string {
	if m != nil {
		return m.AdvUrl
	}
	return ""
}

func (m *AdvertisementRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AdvertisementRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *AdvertisementRequest) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type AdvertisementReply struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Md5        string `protobuf:"bytes,2,opt,name=md5" json:"md5,omitempty"`
	StartTime  int64  `protobuf:"varint,3,opt,name=startTime" json:"startTime,omitempty"`
	EndTime    int64  `protobuf:"varint,4,opt,name=endTime" json:"endTime,omitempty"`
	Source     string `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	FileUrl    string `protobuf:"bytes,6,opt,name=fileUrl" json:"fileUrl,omitempty"`
	State      int32  `protobuf:"varint,7,opt,name=state" json:"state,omitempty"`
	Advertiser string `protobuf:"bytes,8,opt,name=advertiser" json:"advertiser,omitempty"`
	ErrorCode  int32  `protobuf:"varint,9,opt,name=errorCode" json:"errorCode,omitempty"`
	Signature  string `protobuf:"bytes,10,opt,name=signature" json:"signature,omitempty"`
	Id         int32  `protobuf:"varint,11,opt,name=id" json:"id,omitempty"`
	AdvUrl     string `protobuf:"bytes,12,opt,name=advUrl" json:"advUrl,omitempty"`
	Tid        int64  `protobuf:"varint,13,opt,name=tid" json:"tid,omitempty"`
	FileName   string `protobuf:"bytes,14,opt,name=fileName" json:"fileName,omitempty"`
}

func (m *AdvertisementReply) Reset()                    { *m = AdvertisementReply{} }
func (m *AdvertisementReply) String() string            { return proto.CompactTextString(m) }
func (*AdvertisementReply) ProtoMessage()               {}
func (*AdvertisementReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AdvertisementReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AdvertisementReply) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *AdvertisementReply) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AdvertisementReply) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *AdvertisementReply) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AdvertisementReply) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *AdvertisementReply) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *AdvertisementReply) GetAdvertiser() string {
	if m != nil {
		return m.Advertiser
	}
	return ""
}

func (m *AdvertisementReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *AdvertisementReply) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AdvertisementReply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AdvertisementReply) GetAdvUrl() string {
	if m != nil {
		return m.AdvUrl
	}
	return ""
}

func (m *AdvertisementReply) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *AdvertisementReply) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type MapAdvertisementReply struct {
	Advs       []*AdvertisementReply `protobuf:"bytes,1,rep,name=advs" json:"advs,omitempty"`
	ErrorCode  int32                 `protobuf:"varint,2,opt,name=errorCode" json:"errorCode,omitempty"`
	TotalCount int32                 `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *MapAdvertisementReply) Reset()                    { *m = MapAdvertisementReply{} }
func (m *MapAdvertisementReply) String() string            { return proto.CompactTextString(m) }
func (*MapAdvertisementReply) ProtoMessage()               {}
func (*MapAdvertisementReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MapAdvertisementReply) GetAdvs() []*AdvertisementReply {
	if m != nil {
		return m.Advs
	}
	return nil
}

func (m *MapAdvertisementReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *MapAdvertisementReply) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*AdvertisementRequest)(nil), "adv.AdvertisementRequest")
	proto.RegisterType((*AdvertisementReply)(nil), "adv.AdvertisementReply")
	proto.RegisterType((*MapAdvertisementReply)(nil), "adv.MapAdvertisementReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Advertisement service

type AdvertisementClient interface {
	NewAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error)
	UpdateAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error)
	GetAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error)
	GetAllAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*MapAdvertisementReply, error)
	DelAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error)
	GetOneAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error)
}

type advertisementClient struct {
	cc *grpc.ClientConn
}

func NewAdvertisementClient(cc *grpc.ClientConn) AdvertisementClient {
	return &advertisementClient{cc}
}

func (c *advertisementClient) NewAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error) {
	out := new(AdvertisementReply)
	err := grpc.Invoke(ctx, "/adv.Advertisement/NewAdvertisement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementClient) UpdateAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error) {
	out := new(AdvertisementReply)
	err := grpc.Invoke(ctx, "/adv.Advertisement/UpdateAdvertisement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementClient) GetAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error) {
	out := new(AdvertisementReply)
	err := grpc.Invoke(ctx, "/adv.Advertisement/GetAdvertisement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementClient) GetAllAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*MapAdvertisementReply, error) {
	out := new(MapAdvertisementReply)
	err := grpc.Invoke(ctx, "/adv.Advertisement/GetAllAdvertisement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementClient) DelAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error) {
	out := new(AdvertisementReply)
	err := grpc.Invoke(ctx, "/adv.Advertisement/DelAdvertisement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementClient) GetOneAdvertisement(ctx context.Context, in *AdvertisementRequest, opts ...grpc.CallOption) (*AdvertisementReply, error) {
	out := new(AdvertisementReply)
	err := grpc.Invoke(ctx, "/adv.Advertisement/GetOneAdvertisement", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Advertisement service

type AdvertisementServer interface {
	NewAdvertisement(context.Context, *AdvertisementRequest) (*AdvertisementReply, error)
	UpdateAdvertisement(context.Context, *AdvertisementRequest) (*AdvertisementReply, error)
	GetAdvertisement(context.Context, *AdvertisementRequest) (*AdvertisementReply, error)
	GetAllAdvertisement(context.Context, *AdvertisementRequest) (*MapAdvertisementReply, error)
	DelAdvertisement(context.Context, *AdvertisementRequest) (*AdvertisementReply, error)
	GetOneAdvertisement(context.Context, *AdvertisementRequest) (*AdvertisementReply, error)
}

func RegisterAdvertisementServer(s *grpc.Server, srv AdvertisementServer) {
	s.RegisterService(&_Advertisement_serviceDesc, srv)
}

func _Advertisement_NewAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).NewAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adv.Advertisement/NewAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).NewAdvertisement(ctx, req.(*AdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertisement_UpdateAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).UpdateAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adv.Advertisement/UpdateAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).UpdateAdvertisement(ctx, req.(*AdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertisement_GetAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).GetAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adv.Advertisement/GetAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).GetAdvertisement(ctx, req.(*AdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertisement_GetAllAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).GetAllAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adv.Advertisement/GetAllAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).GetAllAdvertisement(ctx, req.(*AdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertisement_DelAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).DelAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adv.Advertisement/DelAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).DelAdvertisement(ctx, req.(*AdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertisement_GetOneAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).GetOneAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adv.Advertisement/GetOneAdvertisement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).GetOneAdvertisement(ctx, req.(*AdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Advertisement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adv.Advertisement",
	HandlerType: (*AdvertisementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAdvertisement",
			Handler:    _Advertisement_NewAdvertisement_Handler,
		},
		{
			MethodName: "UpdateAdvertisement",
			Handler:    _Advertisement_UpdateAdvertisement_Handler,
		},
		{
			MethodName: "GetAdvertisement",
			Handler:    _Advertisement_GetAdvertisement_Handler,
		},
		{
			MethodName: "GetAllAdvertisement",
			Handler:    _Advertisement_GetAllAdvertisement_Handler,
		},
		{
			MethodName: "DelAdvertisement",
			Handler:    _Advertisement_DelAdvertisement_Handler,
		},
		{
			MethodName: "GetOneAdvertisement",
			Handler:    _Advertisement_GetOneAdvertisement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/adv/advertisement.proto",
}

func init() { proto.RegisterFile("api/adv/advertisement.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0x4d, 0xd2, 0xcf, 0x53, 0x5b, 0xcb, 0xec, 0xaa, 0x63, 0x5d, 0x96, 0xd2, 0xab, 0x82,
	0x50, 0x61, 0xc5, 0x07, 0x90, 0x15, 0x76, 0x41, 0xac, 0x10, 0xdc, 0x07, 0x18, 0x77, 0x8e, 0xcb,
	0x40, 0x9a, 0xc4, 0xc9, 0x69, 0xc4, 0x5b, 0x5f, 0xd2, 0x97, 0xf0, 0x09, 0xbc, 0x92, 0x39, 0xd9,
	0x4d, 0x93, 0xb6, 0x01, 0x2f, 0x02, 0x5e, 0x14, 0xce, 0xc7, 0xcc, 0x7f, 0xfe, 0xe7, 0x37, 0xd3,
	0xc0, 0x4b, 0x95, 0x9a, 0xd7, 0x4a, 0xe7, 0xee, 0x87, 0x96, 0x4c, 0x86, 0x1b, 0x8c, 0x69, 0x95,
	0xda, 0x84, 0x12, 0x11, 0x28, 0x9d, 0x2f, 0xfe, 0xf8, 0x70, 0xfa, 0xae, 0xda, 0x0c, 0xf1, 0xdb,
	0x16, 0x33, 0x12, 0x02, 0x3a, 0xb1, 0xda, 0xa0, 0xf4, 0xe6, 0xde, 0x72, 0x18, 0x72, 0x2c, 0xa6,
	0x10, 0x6c, 0xf4, 0x5b, 0xe9, 0x73, 0xc9, 0x85, 0xe2, 0x0c, 0x86, 0x19, 0x29, 0x4b, 0x9f, 0xcd,
	0x06, 0x65, 0x30, 0xf7, 0x96, 0x41, 0xb8, 0x2b, 0x08, 0x09, 0x7d, 0x8c, 0x35, 0xf7, 0x3a, 0xdc,
	0x7b, 0x48, 0xc5, 0x33, 0xe8, 0x65, 0xc9, 0xd6, 0xde, 0xa2, 0xec, 0xb2, 0xd8, 0x7d, 0x26, 0x66,
	0x30, 0xf8, 0x6a, 0x22, 0x5c, 0xbb, 0x93, 0x7b, 0xdc, 0x29, 0x73, 0xa7, 0xe6, 0xe2, 0x1b, 0x1b,
	0xc9, 0x3e, 0xb7, 0x1e, 0x52, 0x71, 0x0a, 0xdd, 0x8c, 0x14, 0xa1, 0x1c, 0xcc, 0xbd, 0x65, 0x37,
	0x2c, 0x12, 0x71, 0x0e, 0x50, 0x8e, 0x6d, 0xe5, 0x90, 0xb7, 0x54, 0x2a, 0x62, 0x02, 0xbe, 0xd1,
	0x12, 0x78, 0x8b, 0x6f, 0x34, 0xcf, 0x62, 0xee, 0x62, 0x45, 0x5b, 0x8b, 0x72, 0xc4, 0xcb, 0x77,
	0x05, 0xe7, 0x58, 0xe9, 0xdc, 0x1d, 0xfe, 0xb8, 0x70, 0x5c, 0x64, 0x8e, 0x53, 0xaa, 0xee, 0x50,
	0x8e, 0x59, 0x87, 0x63, 0xe7, 0xe7, 0x36, 0xd9, 0xc6, 0x24, 0x27, 0x85, 0x1f, 0x4e, 0x1c, 0x3d,
	0x32, 0x5a, 0x3e, 0x61, 0x12, 0x2e, 0x5c, 0xfc, 0xf6, 0x41, 0xec, 0xc1, 0x4f, 0xa3, 0x1f, 0xff,
	0x09, 0x7d, 0x05, 0x6f, 0xaf, 0x01, 0x6f, 0xbf, 0x19, 0xef, 0xe0, 0x00, 0xef, 0x19, 0x0c, 0xd1,
	0xda, 0xc4, 0x5e, 0x26, 0x1a, 0x99, 0x7e, 0x37, 0xdc, 0x15, 0xea, 0xb0, 0x61, 0x1f, 0x76, 0x71,
	0x35, 0xa3, 0xf2, 0x6a, 0x9a, 0xe0, 0xdf, 0x23, 0x1d, 0x97, 0x48, 0x6b, 0x0f, 0x68, 0x52, 0x7f,
	0x40, 0x8b, 0x9f, 0x1e, 0x3c, 0xfd, 0xa8, 0xd2, 0x23, 0xc4, 0x5f, 0x41, 0x47, 0xe9, 0x3c, 0x93,
	0xde, 0x3c, 0x58, 0x8e, 0x2e, 0x9e, 0xaf, 0x94, 0xce, 0x57, 0x87, 0xcb, 0x42, 0x5e, 0x54, 0x1f,
	0xcc, 0xdf, 0x1f, 0xec, 0x1c, 0x80, 0x12, 0x52, 0xd1, 0x25, 0x3f, 0x80, 0x80, 0xdb, 0x95, 0xca,
	0xc5, 0xaf, 0x00, 0xc6, 0x35, 0x69, 0x71, 0x0d, 0xd3, 0x35, 0x7e, 0xaf, 0xd7, 0x5e, 0x1c, 0xb3,
	0xc0, 0x7f, 0xcc, 0x59, 0x93, 0xbb, 0xc5, 0x23, 0xf1, 0x01, 0x4e, 0x6e, 0x52, 0xad, 0x08, 0xdb,
	0x10, 0xbb, 0x86, 0xe9, 0x15, 0x52, 0x1b, 0x4a, 0x6b, 0x38, 0x71, 0x4a, 0x51, 0xf4, 0xcf, 0x62,
	0x33, 0x6e, 0x1d, 0xbd, 0xab, 0xc2, 0xd9, 0x7b, 0x8c, 0x5a, 0x02, 0x76, 0x85, 0xf4, 0x29, 0x6e,
	0x03, 0xd8, 0x97, 0x1e, 0x7f, 0x56, 0xdf, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x00, 0xc5, 0x1c,
	0x02, 0x75, 0x05, 0x00, 0x00,
}
