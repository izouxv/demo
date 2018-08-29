// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/mangementSso.proto

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

type Sort int32

const (
	Sort_ASC  Sort = 0
	Sort_DESC Sort = 1
)

var Sort_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}
var Sort_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x Sort) String() string {
	return proto.EnumName(Sort_name, int32(x))
}
func (Sort) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PageRequest struct {
	Source string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Count  int32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Order  string `protobuf:"bytes,4,opt,name=order" json:"order,omitempty"`
	Token  string `protobuf:"bytes,5,opt,name=token" json:"token,omitempty"`
	Sort   Sort   `protobuf:"varint,6,opt,name=sort,enum=pb.Sort" json:"sort,omitempty"`
}

func (m *PageRequest) Reset()                    { *m = PageRequest{} }
func (m *PageRequest) String() string            { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()               {}
func (*PageRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *PageRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *PageRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PageRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *PageRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *PageRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PageRequest) GetSort() Sort {
	if m != nil {
		return m.Sort
	}
	return Sort_ASC
}

type PageSsoReply struct {
	MSsos      []*MSsoInfo `protobuf:"bytes,1,rep,name=mSsos" json:"mSsos,omitempty"`
	Code       int32       `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	TotalCount int32       `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *PageSsoReply) Reset()                    { *m = PageSsoReply{} }
func (m *PageSsoReply) String() string            { return proto.CompactTextString(m) }
func (*PageSsoReply) ProtoMessage()               {}
func (*PageSsoReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PageSsoReply) GetMSsos() []*MSsoInfo {
	if m != nil {
		return m.MSsos
	}
	return nil
}

func (m *PageSsoReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PageSsoReply) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type MSsoInfo struct {
	Uid        int32  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Nickname   string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	LoginState int32  `protobuf:"varint,4,opt,name=loginState" json:"loginState,omitempty"`
	State      int32  `protobuf:"varint,5,opt,name=state" json:"state,omitempty"`
	RegTime    int64  `protobuf:"varint,6,opt,name=regTime" json:"regTime,omitempty"`
	RegIP      string `protobuf:"bytes,7,opt,name=regIP" json:"regIP,omitempty"`
	RegAddr    string `protobuf:"bytes,8,opt,name=regAddr" json:"regAddr,omitempty"`
	LoginTime  int64  `protobuf:"varint,9,opt,name=loginTime" json:"loginTime,omitempty"`
	QuitTime   int64  `protobuf:"varint,10,opt,name=quitTime" json:"quitTime,omitempty"`
	NewIP      string `protobuf:"bytes,11,opt,name=newIP" json:"newIP,omitempty"`
	NewAddr    string `protobuf:"bytes,12,opt,name=newAddr" json:"newAddr,omitempty"`
	Token      string `protobuf:"bytes,13,opt,name=token" json:"token,omitempty"`
	DevInfo    string `protobuf:"bytes,14,opt,name=devInfo" json:"devInfo,omitempty"`
	Code       int32  `protobuf:"varint,15,opt,name=code" json:"code,omitempty"`
	Source     string `protobuf:"bytes,16,opt,name=source" json:"source,omitempty"`
}

func (m *MSsoInfo) Reset()                    { *m = MSsoInfo{} }
func (m *MSsoInfo) String() string            { return proto.CompactTextString(m) }
func (*MSsoInfo) ProtoMessage()               {}
func (*MSsoInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *MSsoInfo) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MSsoInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MSsoInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *MSsoInfo) GetLoginState() int32 {
	if m != nil {
		return m.LoginState
	}
	return 0
}

func (m *MSsoInfo) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *MSsoInfo) GetRegTime() int64 {
	if m != nil {
		return m.RegTime
	}
	return 0
}

func (m *MSsoInfo) GetRegIP() string {
	if m != nil {
		return m.RegIP
	}
	return ""
}

func (m *MSsoInfo) GetRegAddr() string {
	if m != nil {
		return m.RegAddr
	}
	return ""
}

func (m *MSsoInfo) GetLoginTime() int64 {
	if m != nil {
		return m.LoginTime
	}
	return 0
}

func (m *MSsoInfo) GetQuitTime() int64 {
	if m != nil {
		return m.QuitTime
	}
	return 0
}

func (m *MSsoInfo) GetNewIP() string {
	if m != nil {
		return m.NewIP
	}
	return ""
}

func (m *MSsoInfo) GetNewAddr() string {
	if m != nil {
		return m.NewAddr
	}
	return ""
}

func (m *MSsoInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MSsoInfo) GetDevInfo() string {
	if m != nil {
		return m.DevInfo
	}
	return ""
}

func (m *MSsoInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MSsoInfo) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type MSsoReply struct {
	Code   int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
}

func (m *MSsoReply) Reset()                    { *m = MSsoReply{} }
func (m *MSsoReply) String() string            { return proto.CompactTextString(m) }
func (*MSsoReply) ProtoMessage()               {}
func (*MSsoReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *MSsoReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MSsoReply) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func init() {
	proto.RegisterType((*PageRequest)(nil), "pb.PageRequest")
	proto.RegisterType((*PageSsoReply)(nil), "pb.PageSsoReply")
	proto.RegisterType((*MSsoInfo)(nil), "pb.MSsoInfo")
	proto.RegisterType((*MSsoReply)(nil), "pb.MSsoReply")
	proto.RegisterEnum("pb.Sort", Sort_name, Sort_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MSso service

type MSsoClient interface {
	GetPageSsoInfos(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageSsoReply, error)
	SearchSsoInfo(ctx context.Context, in *MSsoInfo, opts ...grpc.CallOption) (*MSsoInfo, error)
	DeleteAccount(ctx context.Context, in *MSsoInfo, opts ...grpc.CallOption) (*MSsoReply, error)
	GetPageDevices(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*BatchDeviceRe, error)
	SearchDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error)
}

type mSsoClient struct {
	cc *grpc.ClientConn
}

func NewMSsoClient(cc *grpc.ClientConn) MSsoClient {
	return &mSsoClient{cc}
}

func (c *mSsoClient) GetPageSsoInfos(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageSsoReply, error) {
	out := new(PageSsoReply)
	err := grpc.Invoke(ctx, "/pb.MSso/GetPageSsoInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSsoClient) SearchSsoInfo(ctx context.Context, in *MSsoInfo, opts ...grpc.CallOption) (*MSsoInfo, error) {
	out := new(MSsoInfo)
	err := grpc.Invoke(ctx, "/pb.MSso/SearchSsoInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSsoClient) DeleteAccount(ctx context.Context, in *MSsoInfo, opts ...grpc.CallOption) (*MSsoReply, error) {
	out := new(MSsoReply)
	err := grpc.Invoke(ctx, "/pb.MSso/DeleteAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSsoClient) GetPageDevices(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*BatchDeviceRe, error) {
	out := new(BatchDeviceRe)
	err := grpc.Invoke(ctx, "/pb.MSso/GetPageDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSsoClient) SearchDevice(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*DeviceReply, error) {
	out := new(DeviceReply)
	err := grpc.Invoke(ctx, "/pb.MSso/SearchDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MSso service

type MSsoServer interface {
	GetPageSsoInfos(context.Context, *PageRequest) (*PageSsoReply, error)
	SearchSsoInfo(context.Context, *MSsoInfo) (*MSsoInfo, error)
	DeleteAccount(context.Context, *MSsoInfo) (*MSsoReply, error)
	GetPageDevices(context.Context, *PageRequest) (*BatchDeviceRe, error)
	SearchDevice(context.Context, *DeviceRequest) (*DeviceReply, error)
}

func RegisterMSsoServer(s *grpc.Server, srv MSsoServer) {
	s.RegisterService(&_MSso_serviceDesc, srv)
}

func _MSso_GetPageSsoInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSsoServer).GetPageSsoInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MSso/GetPageSsoInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSsoServer).GetPageSsoInfos(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSso_SearchSsoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MSsoInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSsoServer).SearchSsoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MSso/SearchSsoInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSsoServer).SearchSsoInfo(ctx, req.(*MSsoInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSso_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MSsoInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSsoServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MSso/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSsoServer).DeleteAccount(ctx, req.(*MSsoInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSso_GetPageDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSsoServer).GetPageDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MSso/GetPageDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSsoServer).GetPageDevices(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSso_SearchDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSsoServer).SearchDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MSso/SearchDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSsoServer).SearchDevice(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MSso_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MSso",
	HandlerType: (*MSsoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPageSsoInfos",
			Handler:    _MSso_GetPageSsoInfos_Handler,
		},
		{
			MethodName: "SearchSsoInfo",
			Handler:    _MSso_SearchSsoInfo_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _MSso_DeleteAccount_Handler,
		},
		{
			MethodName: "GetPageDevices",
			Handler:    _MSso_GetPageDevices_Handler,
		},
		{
			MethodName: "SearchDevice",
			Handler:    _MSso_SearchDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/mangementSso.proto",
}

func init() { proto.RegisterFile("pb/mangementSso.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdb, 0x6e, 0xda, 0x40,
	0x10, 0xc5, 0xd8, 0x0e, 0x30, 0x5c, 0xbb, 0x6a, 0xab, 0x2d, 0x8a, 0x2a, 0xe4, 0x27, 0x54, 0xa9,
	0x44, 0xa2, 0x55, 0xfa, 0x4c, 0xa1, 0xaa, 0x78, 0x88, 0x84, 0xec, 0xfe, 0x80, 0xb1, 0x27, 0x8e,
	0x15, 0xf0, 0x3a, 0xf6, 0x3a, 0x51, 0xbf, 0xa1, 0x5f, 0xd0, 0x7e, 0x6d, 0xb5, 0xb3, 0x5e, 0x30,
	0x52, 0xde, 0xf6, 0x9c, 0x99, 0x9d, 0x39, 0x9e, 0x39, 0x6b, 0x78, 0x97, 0xef, 0x6f, 0x8e, 0x61,
	0x96, 0xe0, 0x11, 0x33, 0x19, 0x94, 0x62, 0x91, 0x17, 0x42, 0x0a, 0xd6, 0xce, 0xf7, 0xd3, 0x71,
	0xbe, 0xbf, 0x89, 0xf1, 0x39, 0x8d, 0x50, 0x93, 0xde, 0x5f, 0x0b, 0xfa, 0xbb, 0x30, 0x41, 0x1f,
	0x9f, 0x2a, 0x2c, 0x25, 0x7b, 0x0f, 0x57, 0xa5, 0xa8, 0x8a, 0x08, 0xb9, 0x35, 0xb3, 0xe6, 0x3d,
	0xbf, 0x46, 0x8c, 0x81, 0x93, 0x87, 0x09, 0xf2, 0xf6, 0xcc, 0x9a, 0xbb, 0x3e, 0x9d, 0xd9, 0x5b,
	0x70, 0x23, 0x51, 0x65, 0x92, 0xdb, 0x44, 0x6a, 0xa0, 0x58, 0x51, 0xc4, 0x58, 0x70, 0x87, 0x0a,
	0x68, 0xa0, 0x58, 0x29, 0x1e, 0x31, 0xe3, 0xae, 0x66, 0x09, 0xb0, 0x6b, 0x70, 0x4a, 0x51, 0x48,
	0x7e, 0x35, 0xb3, 0xe6, 0xa3, 0x65, 0x77, 0x91, 0xef, 0x17, 0x81, 0x28, 0xa4, 0x4f, 0xac, 0x77,
	0x0f, 0x03, 0x25, 0x2d, 0x28, 0x85, 0x8f, 0xf9, 0xe1, 0x37, 0xf3, 0xc0, 0x3d, 0x06, 0xa5, 0x28,
	0xb9, 0x35, 0xb3, 0xe7, 0xfd, 0xe5, 0x40, 0xa5, 0xdf, 0x05, 0xa5, 0xd8, 0x66, 0xf7, 0xc2, 0xd7,
	0x21, 0xa5, 0x33, 0x12, 0xf1, 0x49, 0xa7, 0x3a, 0xb3, 0x8f, 0x00, 0x52, 0xc8, 0xf0, 0xb0, 0x6e,
	0x88, 0x6d, 0x30, 0xde, 0x3f, 0x1b, 0xba, 0xa6, 0x0e, 0x9b, 0x80, 0x5d, 0xa5, 0x31, 0x7d, 0xbd,
	0xeb, 0xab, 0x23, 0x9b, 0x42, 0xb7, 0x2a, 0xb1, 0xc8, 0xc2, 0xa3, 0x2e, 0xdb, 0xf3, 0x4f, 0x58,
	0xc5, 0xb2, 0x34, 0x7a, 0xa4, 0x98, 0xad, 0x63, 0x06, 0xab, 0xb6, 0x07, 0x91, 0xa4, 0x59, 0x20,
	0x43, 0x89, 0x34, 0x0d, 0xd7, 0x6f, 0x30, 0x6a, 0x24, 0x25, 0x85, 0x5c, 0x3d, 0x3e, 0x02, 0x8c,
	0x43, 0xa7, 0xc0, 0xe4, 0x57, 0x7a, 0x44, 0x9a, 0x8a, 0xed, 0x1b, 0xa8, 0xf2, 0x0b, 0x4c, 0xb6,
	0x3b, 0xde, 0xd1, 0x23, 0x24, 0x50, 0xe7, 0xaf, 0xe2, 0xb8, 0xe0, 0x5d, 0xe2, 0x0d, 0x64, 0xd7,
	0xd0, 0xa3, 0x6e, 0x54, 0xab, 0x47, 0xb5, 0xce, 0x84, 0x52, 0xfe, 0x54, 0xa5, 0x92, 0x82, 0x40,
	0xc1, 0x13, 0x56, 0x9d, 0x32, 0x7c, 0xd9, 0xee, 0x78, 0x5f, 0x77, 0x22, 0xa0, 0x3a, 0x65, 0xf8,
	0x42, 0x9d, 0x06, 0xba, 0x53, 0x0d, 0xcf, 0xcb, 0x1d, 0x36, 0x97, 0xcb, 0xa1, 0x13, 0xe3, 0xb3,
	0x1a, 0x2a, 0x1f, 0xe9, 0xfc, 0x1a, 0x9e, 0x96, 0x34, 0x6e, 0x2c, 0xe9, 0x6c, 0xbc, 0x49, 0xd3,
	0x78, 0xde, 0x37, 0xe8, 0xdd, 0x9d, 0x1c, 0x60, 0x2e, 0x5a, 0xaf, 0x5e, 0x6c, 0x37, 0x2f, 0x7e,
	0xfa, 0x00, 0x8e, 0xf2, 0x12, 0xeb, 0x80, 0xbd, 0x0a, 0xd6, 0x93, 0x16, 0xeb, 0x82, 0xb3, 0xf9,
	0x11, 0xac, 0x27, 0xd6, 0xf2, 0x4f, 0x1b, 0x1c, 0x55, 0x94, 0xdd, 0xc2, 0xf8, 0x27, 0xca, 0xda,
	0x64, 0x4a, 0x5a, 0xc9, 0xc6, 0xca, 0x55, 0x8d, 0x17, 0x31, 0x9d, 0x18, 0xc2, 0xa8, 0xf0, 0x5a,
	0xec, 0x33, 0x0c, 0x03, 0x0c, 0x8b, 0xe8, 0xc1, 0xb8, 0xe6, 0xc2, 0x8b, 0xd3, 0x0b, 0xe4, 0xb5,
	0xd8, 0x02, 0x86, 0x1b, 0x3c, 0xa0, 0xc4, 0x55, 0xa4, 0xdf, 0xc8, 0x65, 0xfa, 0xd0, 0x20, 0x53,
	0xfe, 0x16, 0x46, 0xb5, 0xac, 0x0d, 0xbd, 0xd5, 0x57, 0x54, 0xbd, 0x51, 0xc4, 0xf7, 0x50, 0x46,
	0x0f, 0x3a, 0xc5, 0x47, 0xaf, 0xc5, 0xbe, 0xc2, 0x40, 0xcb, 0xd2, 0x1c, 0xa3, 0x24, 0x13, 0xd7,
	0xf7, 0xc6, 0x4d, 0x8a, 0xba, 0xed, 0xaf, 0xe8, 0x4f, 0xf0, 0xe5, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb6, 0xbc, 0xca, 0xa2, 0x37, 0x04, 0x00, 0x00,
}