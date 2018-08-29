// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/device.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/device.proto
	pb/mangementSso.proto

It has these top-level messages:
	DeviceRequest
	DeviceReply
	BatchDeviceRe
	PageRequest
	PageSsoReply
	MSsoInfo
	MSsoReply
*/
package pb

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

type Types int32

const (
	Types_Pet  Types = 0
	Types_User Types = 1
)

var Types_name = map[int32]string{
	0: "Pet",
	1: "User",
}
var Types_value = map[string]int32{
	"Pet":  0,
	"User": 1,
}

func (x Types) String() string {
	return proto.EnumName(Types_name, int32(x))
}
func (Types) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DevPermit int32

const (
	DevPermit_DevPermit0 DevPermit = 0
	DevPermit_DevPermit1 DevPermit = 1
)

var DevPermit_name = map[int32]string{
	0: "DevPermit0",
	1: "DevPermit1",
}
var DevPermit_value = map[string]int32{
	"DevPermit0": 0,
	"DevPermit1": 1,
}

func (x DevPermit) String() string {
	return proto.EnumName(DevPermit_name, int32(x))
}
func (DevPermit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DeviceRequest struct {
	Source          string    `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Did             int32     `protobuf:"varint,2,opt,name=did" json:"did,omitempty"`
	Uid             int32     `protobuf:"varint,3,opt,name=uid" json:"uid,omitempty"`
	Pid             int32     `protobuf:"varint,4,opt,name=pid" json:"pid,omitempty"`
	Touid           int32     `protobuf:"varint,5,opt,name=touid" json:"touid,omitempty"`
	Sn              string    `protobuf:"bytes,6,opt,name=sn" json:"sn,omitempty"`
	DeviceMac       string    `protobuf:"bytes,7,opt,name=deviceMac" json:"deviceMac,omitempty"`
	DeviceName      string    `protobuf:"bytes,8,opt,name=deviceName" json:"deviceName,omitempty"`
	DeviceVersion   string    `protobuf:"bytes,9,opt,name=deviceVersion" json:"deviceVersion,omitempty"`
	SoftwareVersion string    `protobuf:"bytes,10,opt,name=softwareVersion" json:"softwareVersion,omitempty"`
	Permit          DevPermit `protobuf:"varint,11,opt,name=permit,enum=pb.DevPermit" json:"permit,omitempty"`
	Types           Types     `protobuf:"varint,12,opt,name=types,enum=pb.Types" json:"types,omitempty"`
	Isdel           uint32    `protobuf:"varint,13,opt,name=isdel" json:"isdel,omitempty"`
	Input           string    `protobuf:"bytes,14,opt,name=input" json:"input,omitempty"`
}

func (m *DeviceRequest) Reset()                    { *m = DeviceRequest{} }
func (m *DeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceRequest) ProtoMessage()               {}
func (*DeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeviceRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *DeviceRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *DeviceRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *DeviceRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *DeviceRequest) GetTouid() int32 {
	if m != nil {
		return m.Touid
	}
	return 0
}

func (m *DeviceRequest) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *DeviceRequest) GetDeviceMac() string {
	if m != nil {
		return m.DeviceMac
	}
	return ""
}

func (m *DeviceRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *DeviceRequest) GetDeviceVersion() string {
	if m != nil {
		return m.DeviceVersion
	}
	return ""
}

func (m *DeviceRequest) GetSoftwareVersion() string {
	if m != nil {
		return m.SoftwareVersion
	}
	return ""
}

func (m *DeviceRequest) GetPermit() DevPermit {
	if m != nil {
		return m.Permit
	}
	return DevPermit_DevPermit0
}

func (m *DeviceRequest) GetTypes() Types {
	if m != nil {
		return m.Types
	}
	return Types_Pet
}

func (m *DeviceRequest) GetIsdel() uint32 {
	if m != nil {
		return m.Isdel
	}
	return 0
}

func (m *DeviceRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type DeviceReply struct {
	Source          string    `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Did             int32     `protobuf:"varint,2,opt,name=did" json:"did,omitempty"`
	Uid             int32     `protobuf:"varint,3,opt,name=uid" json:"uid,omitempty"`
	Pid             int32     `protobuf:"varint,4,opt,name=pid" json:"pid,omitempty"`
	Touid           int32     `protobuf:"varint,5,opt,name=touid" json:"touid,omitempty"`
	Sn              string    `protobuf:"bytes,6,opt,name=sn" json:"sn,omitempty"`
	DeviceMac       string    `protobuf:"bytes,7,opt,name=deviceMac" json:"deviceMac,omitempty"`
	DeviceName      string    `protobuf:"bytes,8,opt,name=deviceName" json:"deviceName,omitempty"`
	DeviceVersion   string    `protobuf:"bytes,9,opt,name=deviceVersion" json:"deviceVersion,omitempty"`
	SoftwareVersion string    `protobuf:"bytes,10,opt,name=softwareVersion" json:"softwareVersion,omitempty"`
	Permit          DevPermit `protobuf:"varint,11,opt,name=permit,enum=pb.DevPermit" json:"permit,omitempty"`
	Types           Types     `protobuf:"varint,12,opt,name=types,enum=pb.Types" json:"types,omitempty"`
	Isdel           uint32    `protobuf:"varint,13,opt,name=isdel" json:"isdel,omitempty"`
	ShareUrl        string    `protobuf:"bytes,14,opt,name=shareUrl" json:"shareUrl,omitempty"`
	Code            int32     `protobuf:"varint,15,opt,name=code" json:"code,omitempty"`
}

func (m *DeviceReply) Reset()                    { *m = DeviceReply{} }
func (m *DeviceReply) String() string            { return proto.CompactTextString(m) }
func (*DeviceReply) ProtoMessage()               {}
func (*DeviceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceReply) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *DeviceReply) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *DeviceReply) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *DeviceReply) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *DeviceReply) GetTouid() int32 {
	if m != nil {
		return m.Touid
	}
	return 0
}

func (m *DeviceReply) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *DeviceReply) GetDeviceMac() string {
	if m != nil {
		return m.DeviceMac
	}
	return ""
}

func (m *DeviceReply) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *DeviceReply) GetDeviceVersion() string {
	if m != nil {
		return m.DeviceVersion
	}
	return ""
}

func (m *DeviceReply) GetSoftwareVersion() string {
	if m != nil {
		return m.SoftwareVersion
	}
	return ""
}

func (m *DeviceReply) GetPermit() DevPermit {
	if m != nil {
		return m.Permit
	}
	return DevPermit_DevPermit0
}

func (m *DeviceReply) GetTypes() Types {
	if m != nil {
		return m.Types
	}
	return Types_Pet
}

func (m *DeviceReply) GetIsdel() uint32 {
	if m != nil {
		return m.Isdel
	}
	return 0
}

func (m *DeviceReply) GetShareUrl() string {
	if m != nil {
		return m.ShareUrl
	}
	return ""
}

func (m *DeviceReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

// 批量返回
type BatchDeviceRe struct {
	Source     string         `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Code       int32          `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Devices    []*DeviceReply `protobuf:"bytes,3,rep,name=devices" json:"devices,omitempty"`
	TotalCount int32          `protobuf:"varint,4,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *BatchDeviceRe) Reset()                    { *m = BatchDeviceRe{} }
func (m *BatchDeviceRe) String() string            { return proto.CompactTextString(m) }
func (*BatchDeviceRe) ProtoMessage()               {}
func (*BatchDeviceRe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BatchDeviceRe) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *BatchDeviceRe) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchDeviceRe) GetDevices() []*DeviceReply {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *BatchDeviceRe) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceRequest)(nil), "pb.DeviceRequest")
	proto.RegisterType((*DeviceReply)(nil), "pb.DeviceReply")
	proto.RegisterType((*BatchDeviceRe)(nil), "pb.BatchDeviceRe")
	proto.RegisterEnum("pb.Types", Types_name, Types_value)
	proto.RegisterEnum("pb.DevPermit", DevPermit_name, DevPermit_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Devices service

type DevicesClient interface {
	VerificationDeviceBySn(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error)
	SetDeviceBySn(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error)
	DeleteDeviceByDid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error)
	UpdateDeviceByDid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error)
	GetDevicesByUid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*BatchDeviceRe, error)
	GetDevicesByDid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error)
	GetDeviceSn(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error)
}

type devicesClient struct {
	cc *grpc.ClientConn
}

func NewDevicesClient(cc *grpc.ClientConn) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) VerificationDeviceBySn(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error) {
	out := new(DeviceReply)
	err := grpc.Invoke(ctx, "/pb.Devices/VerificationDeviceBySn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) SetDeviceBySn(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error) {
	out := new(DeviceReply)
	err := grpc.Invoke(ctx, "/pb.Devices/SetDeviceBySn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) DeleteDeviceByDid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error) {
	out := new(DeviceReply)
	err := grpc.Invoke(ctx, "/pb.Devices/DeleteDeviceByDid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) UpdateDeviceByDid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error) {
	out := new(DeviceReply)
	err := grpc.Invoke(ctx, "/pb.Devices/UpdateDeviceByDid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDevicesByUid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*BatchDeviceRe, error) {
	out := new(BatchDeviceRe)
	err := grpc.Invoke(ctx, "/pb.Devices/GetDevicesByUid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDevicesByDid(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error) {
	out := new(DeviceReply)
	err := grpc.Invoke(ctx, "/pb.Devices/GetDevicesByDid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDeviceSn(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error) {
	out := new(DeviceReply)
	err := grpc.Invoke(ctx, "/pb.Devices/GetDeviceSn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Devices service

type DevicesServer interface {
	VerificationDeviceBySn(context.Context, *DeviceRequest) (*DeviceReply, error)
	SetDeviceBySn(context.Context, *DeviceRequest) (*DeviceReply, error)
	DeleteDeviceByDid(context.Context, *DeviceRequest) (*DeviceReply, error)
	UpdateDeviceByDid(context.Context, *DeviceRequest) (*DeviceReply, error)
	GetDevicesByUid(context.Context, *DeviceRequest) (*BatchDeviceRe, error)
	GetDevicesByDid(context.Context, *DeviceRequest) (*DeviceReply, error)
	GetDeviceSn(context.Context, *DeviceRequest) (*DeviceReply, error)
}

func RegisterDevicesServer(s *grpc.Server, srv DevicesServer) {
	s.RegisterService(&_Devices_serviceDesc, srv)
}

func _Devices_VerificationDeviceBySn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).VerificationDeviceBySn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Devices/VerificationDeviceBySn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).VerificationDeviceBySn(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_SetDeviceBySn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).SetDeviceBySn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Devices/SetDeviceBySn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).SetDeviceBySn(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_DeleteDeviceByDid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).DeleteDeviceByDid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Devices/DeleteDeviceByDid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).DeleteDeviceByDid(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_UpdateDeviceByDid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).UpdateDeviceByDid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Devices/UpdateDeviceByDid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).UpdateDeviceByDid(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDevicesByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDevicesByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Devices/GetDevicesByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDevicesByUid(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDevicesByDid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDevicesByDid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Devices/GetDevicesByDid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDevicesByDid(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDeviceSn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDeviceSn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Devices/GetDeviceSn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDeviceSn(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Devices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerificationDeviceBySn",
			Handler:    _Devices_VerificationDeviceBySn_Handler,
		},
		{
			MethodName: "SetDeviceBySn",
			Handler:    _Devices_SetDeviceBySn_Handler,
		},
		{
			MethodName: "DeleteDeviceByDid",
			Handler:    _Devices_DeleteDeviceByDid_Handler,
		},
		{
			MethodName: "UpdateDeviceByDid",
			Handler:    _Devices_UpdateDeviceByDid_Handler,
		},
		{
			MethodName: "GetDevicesByUid",
			Handler:    _Devices_GetDevicesByUid_Handler,
		},
		{
			MethodName: "GetDevicesByDid",
			Handler:    _Devices_GetDevicesByDid_Handler,
		},
		{
			MethodName: "GetDeviceSn",
			Handler:    _Devices_GetDeviceSn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/device.proto",
}

func init() { proto.RegisterFile("pb/device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x9b, 0xa6, 0x7f, 0x4f, 0x4d, 0xdb, 0x3d, 0xc8, 0x32, 0x14, 0xd1, 0x52, 0x14, 0xe2,
	0x0a, 0x55, 0xbb, 0x88, 0x78, 0xe3, 0x45, 0x2d, 0x78, 0xa5, 0x2c, 0x59, 0xbb, 0xf7, 0x69, 0xe6,
	0x2c, 0x3b, 0xd0, 0xcd, 0x8c, 0x99, 0xc9, 0x4a, 0x1f, 0x40, 0xdf, 0x42, 0x7c, 0x55, 0xc9, 0x4c,
	0xd2, 0x6d, 0x17, 0x85, 0x2a, 0x78, 0xe7, 0x5d, 0xbe, 0xdf, 0x9c, 0x2f, 0x73, 0xce, 0xf9, 0x60,
	0x60, 0xa0, 0x56, 0xcf, 0x39, 0xdd, 0x88, 0x84, 0xa6, 0x2a, 0x93, 0x46, 0x62, 0x5d, 0xad, 0x26,
	0x5f, 0x7d, 0x08, 0x16, 0x16, 0x46, 0xf4, 0x39, 0x27, 0x6d, 0xf0, 0x18, 0x5a, 0x5a, 0xe6, 0x59,
	0x42, 0xcc, 0x1b, 0x7b, 0x61, 0x37, 0x2a, 0x15, 0x0e, 0xc1, 0xe7, 0x82, 0xb3, 0xfa, 0xd8, 0x0b,
	0x9b, 0x51, 0xf1, 0x59, 0x90, 0x5c, 0x70, 0xe6, 0x3b, 0x92, 0x3b, 0xa2, 0x04, 0x67, 0x0d, 0x47,
	0x94, 0xe0, 0x78, 0x1f, 0x9a, 0x46, 0x16, 0x55, 0x4d, 0xcb, 0x9c, 0xc0, 0x3e, 0xd4, 0x75, 0xca,
	0x5a, 0xf6, 0xff, 0x75, 0x9d, 0xe2, 0x03, 0xe8, 0xba, 0xce, 0x3e, 0xc4, 0x09, 0x6b, 0x5b, 0x7c,
	0x0b, 0xf0, 0x21, 0x80, 0x13, 0x1f, 0xe3, 0x6b, 0x62, 0x1d, 0x7b, 0xbc, 0x43, 0xf0, 0x31, 0x04,
	0x4e, 0x5d, 0x50, 0xa6, 0x85, 0x4c, 0x59, 0xd7, 0x96, 0xec, 0x43, 0x0c, 0x61, 0xa0, 0xe5, 0xa5,
	0xf9, 0x12, 0x67, 0xdb, 0x3a, 0xb0, 0x75, 0x77, 0x31, 0x3e, 0x81, 0x96, 0xa2, 0xec, 0x5a, 0x18,
	0xd6, 0x1b, 0x7b, 0x61, 0x7f, 0x16, 0x4c, 0xd5, 0x6a, 0xba, 0xa0, 0x9b, 0x33, 0x0b, 0xa3, 0xf2,
	0x10, 0x1f, 0x41, 0xd3, 0x6c, 0x14, 0x69, 0x76, 0xcf, 0x56, 0x75, 0x8b, 0xaa, 0x4f, 0x05, 0x88,
	0x1c, 0x2f, 0x66, 0x17, 0x9a, 0xd3, 0x9a, 0x05, 0x63, 0x2f, 0x0c, 0x22, 0x27, 0x2c, 0x4d, 0x55,
	0x6e, 0x58, 0xdf, 0xde, 0xee, 0xc4, 0xe4, 0xbb, 0x0f, 0xbd, 0x2a, 0x07, 0xb5, 0xde, 0xfc, 0x4f,
	0xe1, 0x1f, 0xa7, 0x30, 0x82, 0x8e, 0xbe, 0x8a, 0x33, 0x5a, 0x66, 0xeb, 0x32, 0x88, 0xad, 0x46,
	0x84, 0x46, 0x22, 0x39, 0xb1, 0x81, 0x5d, 0x96, 0xfd, 0x9e, 0x7c, 0xf3, 0x20, 0x98, 0xc7, 0x26,
	0xb9, 0xaa, 0x42, 0xfa, 0x6d, 0x42, 0x95, 0xbb, 0x7e, 0xeb, 0xc6, 0xa7, 0xd0, 0x76, 0x6b, 0xd0,
	0xcc, 0x1f, 0xfb, 0x61, 0x6f, 0x36, 0x28, 0x87, 0xa9, 0xf2, 0x8e, 0xaa, 0xf3, 0x62, 0xcd, 0x46,
	0x9a, 0x78, 0xfd, 0x4e, 0xe6, 0xa9, 0x29, 0x33, 0xdc, 0x21, 0x27, 0x23, 0x68, 0xda, 0xf1, 0xb0,
	0x0d, 0xfe, 0x19, 0x99, 0x61, 0x0d, 0x3b, 0xd0, 0x58, 0x6a, 0xca, 0x86, 0xde, 0xc9, 0x33, 0xe8,
	0x6e, 0x17, 0x84, 0x7d, 0x80, 0xad, 0x78, 0x31, 0xac, 0xed, 0xe9, 0x97, 0x43, 0x6f, 0xf6, 0xc3,
	0x87, 0xf6, 0xa2, 0xbc, 0xf4, 0x2d, 0x1c, 0x5f, 0x50, 0x26, 0x2e, 0x45, 0x12, 0x1b, 0x21, 0x53,
	0x87, 0xe7, 0x9b, 0xf3, 0x14, 0x8f, 0x76, 0x1b, 0xb5, 0x0f, 0xc4, 0xe8, 0x6e, 0xef, 0x93, 0x1a,
	0xbe, 0x82, 0xe0, 0x9c, 0xcc, 0x1f, 0xdb, 0xde, 0xc0, 0xd1, 0x82, 0xd6, 0x64, 0xa8, 0x72, 0x2e,
	0x04, 0x3f, 0xdc, 0xba, 0x54, 0x3c, 0xfe, 0x3b, 0xeb, 0xe0, 0x7d, 0xd5, 0xac, 0x9e, 0x6f, 0x96,
	0xbf, 0x36, 0x5a, 0xb4, 0x97, 0xf8, 0xa4, 0x86, 0xaf, 0xf7, 0xad, 0x87, 0xdf, 0x79, 0x0a, 0xbd,
	0xad, 0xf1, 0xd0, 0xf5, 0xac, 0x5a, 0xf6, 0x99, 0x3e, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x11,
	0x6b, 0xf5, 0x6d, 0xb9, 0x05, 0x00, 0x00,
}