// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/account.proto

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

type AccountRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salt                 string   `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	State                int32    `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Nickname             string   `protobuf:"bytes,8,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender               int32    `protobuf:"varint,9,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday             int64    `protobuf:"varint,10,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Avatar               string   `protobuf:"bytes,11,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Signature            string   `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`
	Address              string   `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
	Source               string   `protobuf:"bytes,14,opt,name=source,proto3" json:"source,omitempty"`
	Token                string   `protobuf:"bytes,15,opt,name=token,proto3" json:"token,omitempty"`
	Radius               int32    `protobuf:"varint,16,opt,name=radius,proto3" json:"radius,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_a65bb3278ede60a4, []int{0}
}
func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (dst *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(dst, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AccountRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AccountRequest) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *AccountRequest) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *AccountRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AccountRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AccountRequest) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *AccountRequest) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *AccountRequest) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *AccountRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AccountRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccountRequest) GetRadius() int32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

type AccountReply struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salt                 string   `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	State                int32    `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Nickname             string   `protobuf:"bytes,8,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender               int32    `protobuf:"varint,9,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday             int64    `protobuf:"varint,10,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Avatar               string   `protobuf:"bytes,11,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Signature            string   `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`
	Address              string   `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
	Source               string   `protobuf:"bytes,14,opt,name=source,proto3" json:"source,omitempty"`
	Code                 int32    `protobuf:"varint,15,opt,name=code,proto3" json:"code,omitempty"`
	Radius               int32    `protobuf:"varint,16,opt,name=radius,proto3" json:"radius,omitempty"`
	Map                  Map      `protobuf:"varint,17,opt,name=map,proto3,enum=pb.Map" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountReply) Reset()         { *m = AccountReply{} }
func (m *AccountReply) String() string { return proto.CompactTextString(m) }
func (*AccountReply) ProtoMessage()    {}
func (*AccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_a65bb3278ede60a4, []int{1}
}
func (m *AccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountReply.Unmarshal(m, b)
}
func (m *AccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountReply.Marshal(b, m, deterministic)
}
func (dst *AccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountReply.Merge(dst, src)
}
func (m *AccountReply) XXX_Size() int {
	return xxx_messageInfo_AccountReply.Size(m)
}
func (m *AccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccountReply proto.InternalMessageInfo

func (m *AccountReply) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AccountReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountReply) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AccountReply) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *AccountReply) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *AccountReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountReply) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AccountReply) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AccountReply) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *AccountReply) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *AccountReply) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *AccountReply) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AccountReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountReply) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AccountReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AccountReply) GetRadius() int32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *AccountReply) GetMap() Map {
	if m != nil {
		return m.Map
	}
	return Map_AllMap
}

type MultiAccountRequest struct {
	Accounts             map[int32]*AccountRequest `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Source               string                    `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MultiAccountRequest) Reset()         { *m = MultiAccountRequest{} }
func (m *MultiAccountRequest) String() string { return proto.CompactTextString(m) }
func (*MultiAccountRequest) ProtoMessage()    {}
func (*MultiAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_a65bb3278ede60a4, []int{2}
}
func (m *MultiAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiAccountRequest.Unmarshal(m, b)
}
func (m *MultiAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiAccountRequest.Marshal(b, m, deterministic)
}
func (dst *MultiAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiAccountRequest.Merge(dst, src)
}
func (m *MultiAccountRequest) XXX_Size() int {
	return xxx_messageInfo_MultiAccountRequest.Size(m)
}
func (m *MultiAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiAccountRequest proto.InternalMessageInfo

func (m *MultiAccountRequest) GetAccounts() map[int32]*AccountRequest {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *MultiAccountRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type MapAccountReply struct {
	Accounts             map[int32]*AccountReply `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Code                 int32                   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MapAccountReply) Reset()         { *m = MapAccountReply{} }
func (m *MapAccountReply) String() string { return proto.CompactTextString(m) }
func (*MapAccountReply) ProtoMessage()    {}
func (*MapAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_a65bb3278ede60a4, []int{3}
}
func (m *MapAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapAccountReply.Unmarshal(m, b)
}
func (m *MapAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapAccountReply.Marshal(b, m, deterministic)
}
func (dst *MapAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapAccountReply.Merge(dst, src)
}
func (m *MapAccountReply) XXX_Size() int {
	return xxx_messageInfo_MapAccountReply.Size(m)
}
func (m *MapAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MapAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_MapAccountReply proto.InternalMessageInfo

func (m *MapAccountReply) GetAccounts() map[int32]*AccountReply {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *MapAccountReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountRequest)(nil), "pb.AccountRequest")
	proto.RegisterType((*AccountReply)(nil), "pb.AccountReply")
	proto.RegisterType((*MultiAccountRequest)(nil), "pb.MultiAccountRequest")
	proto.RegisterMapType((map[int32]*AccountRequest)(nil), "pb.MultiAccountRequest.AccountsEntry")
	proto.RegisterType((*MapAccountReply)(nil), "pb.MapAccountReply")
	proto.RegisterMapType((map[int32]*AccountReply)(nil), "pb.MapAccountReply.AccountsEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Show(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	UpdateAccountInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetAccountInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetBatchAccountInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Show(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := c.cc.Invoke(ctx, "/pb.Account/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateAccountInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := c.cc.Invoke(ctx, "/pb.Account/UpdateAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccountInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := c.cc.Invoke(ctx, "/pb.Account/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetBatchAccountInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error) {
	out := new(MapAccountReply)
	err := c.cc.Invoke(ctx, "/pb.Account/GetBatchAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Show(context.Context, *AccountRequest) (*AccountReply, error)
	UpdateAccountInfo(context.Context, *AccountRequest) (*AccountReply, error)
	GetAccountInfo(context.Context, *AccountRequest) (*AccountReply, error)
	GetBatchAccountInfo(context.Context, *MultiAccountRequest) (*MapAccountReply, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Show(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UpdateAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/UpdateAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateAccountInfo(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccountInfo(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetBatchAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetBatchAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetBatchAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetBatchAccountInfo(ctx, req.(*MultiAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _Account_Show_Handler,
		},
		{
			MethodName: "UpdateAccountInfo",
			Handler:    _Account_UpdateAccountInfo_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _Account_GetAccountInfo_Handler,
		},
		{
			MethodName: "GetBatchAccountInfo",
			Handler:    _Account_GetBatchAccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/account.proto",
}

func init() { proto.RegisterFile("pb/account.proto", fileDescriptor_account_a65bb3278ede60a4) }

var fileDescriptor_account_a65bb3278ede60a4 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0xe3, 0xfc, 0x9d, 0xb4, 0x69, 0xba, 0xf9, 0xe9, 0xc7, 0x62, 0x71, 0x08, 0x96, 0x40,
	0x3e, 0x05, 0x29, 0x5c, 0x2a, 0x10, 0x87, 0x22, 0x55, 0x15, 0x87, 0x08, 0xc9, 0x88, 0x07, 0xd8,
	0xd8, 0xd3, 0xc6, 0x8a, 0xe3, 0x5d, 0x76, 0xd7, 0xad, 0xfc, 0x22, 0xbc, 0x05, 0x27, 0x9e, 0x80,
	0xa7, 0xe2, 0x8a, 0x76, 0xd7, 0x49, 0x13, 0xd2, 0x08, 0x84, 0xc4, 0x8d, 0xdb, 0x7e, 0xdf, 0xcc,
	0x37, 0xfb, 0xcd, 0xce, 0xd8, 0x30, 0x14, 0xf3, 0x17, 0x2c, 0x49, 0x78, 0x59, 0xe8, 0x89, 0x90,
	0x5c, 0x73, 0xd2, 0x10, 0xf3, 0xc0, 0xb0, 0x02, 0xf5, 0x35, 0x2f, 0xd0, 0xb1, 0xe1, 0x67, 0x1f,
	0x06, 0x17, 0x2e, 0x2f, 0xc6, 0x4f, 0x25, 0x2a, 0x4d, 0x86, 0xe0, 0x97, 0x59, 0x4a, 0xbd, 0xb1,
	0x17, 0xb5, 0x62, 0x73, 0x24, 0x01, 0x74, 0x4b, 0x85, 0xb2, 0x60, 0x2b, 0xa4, 0x8d, 0xb1, 0x17,
	0xf5, 0xe2, 0x0d, 0x36, 0x31, 0xc1, 0x94, 0xba, 0xe3, 0x32, 0xa5, 0xbe, 0x8b, 0xad, 0x31, 0x21,
	0xd0, 0x54, 0x2c, 0xd7, 0xb4, 0x69, 0x79, 0x7b, 0x26, 0xff, 0x41, 0x4b, 0x69, 0xa6, 0x91, 0xb6,
	0x6c, 0x7d, 0x07, 0x0c, 0x8b, 0x2b, 0x96, 0xe5, 0xb4, 0x6d, 0x53, 0x1d, 0x30, 0xac, 0x58, 0xf0,
	0x02, 0x69, 0xc7, 0xb1, 0x16, 0x98, 0x1b, 0x8b, 0x2c, 0x59, 0x5a, 0x37, 0x5d, 0x77, 0xe3, 0x1a,
	0x93, 0xff, 0xa1, 0x7d, 0x83, 0x45, 0x8a, 0x92, 0xf6, 0x6c, 0xf9, 0x1a, 0x19, 0xcd, 0x3c, 0x93,
	0x7a, 0x91, 0xb2, 0x8a, 0xc2, 0xd8, 0x8b, 0xfc, 0x78, 0x83, 0x8d, 0x86, 0xdd, 0x32, 0xcd, 0x24,
	0xed, 0xdb, 0x6a, 0x35, 0x22, 0x4f, 0xa0, 0xa7, 0xb2, 0x9b, 0x82, 0xe9, 0x52, 0x22, 0x3d, 0xb6,
	0xa1, 0x7b, 0x82, 0x50, 0xe8, 0xb0, 0x34, 0x95, 0xa8, 0x14, 0x3d, 0xb1, 0xb1, 0x35, 0x34, 0xf5,
	0x14, 0x2f, 0x65, 0x82, 0x74, 0xe0, 0xea, 0x39, 0x64, 0xba, 0xd1, 0x7c, 0x89, 0x05, 0x3d, 0x75,
	0xdd, 0x58, 0x60, 0xb2, 0x25, 0x4b, 0xb3, 0x52, 0xd1, 0xa1, 0x73, 0xec, 0x50, 0xf8, 0xc5, 0x87,
	0xe3, 0xcd, 0x60, 0x44, 0x5e, 0xfd, 0x1b, 0xcb, 0x5f, 0x1a, 0x0b, 0x81, 0x66, 0xc2, 0x53, 0xb4,
	0x53, 0x69, 0xc5, 0xf6, 0x7c, 0x68, 0x28, 0xe4, 0x31, 0xf8, 0x2b, 0x26, 0xe8, 0xd9, 0xd8, 0x8b,
	0x06, 0xd3, 0xce, 0x44, 0xcc, 0x27, 0x33, 0x26, 0x62, 0xc3, 0x85, 0xdf, 0x3c, 0x18, 0xcd, 0xca,
	0x5c, 0x67, 0x3f, 0x7d, 0x4d, 0x17, 0xd0, 0xad, 0xbf, 0x43, 0x45, 0xbd, 0xb1, 0x1f, 0xf5, 0xa7,
	0xcf, 0xac, 0x6e, 0x3f, 0x75, 0x52, 0x43, 0x75, 0x59, 0x68, 0x59, 0xc5, 0x1b, 0xd9, 0x96, 0xf3,
	0xc6, 0xb6, 0xf3, 0xe0, 0x3d, 0x9c, 0xec, 0x48, 0xcc, 0x8a, 0x2c, 0xb1, 0x5a, 0xaf, 0xc8, 0x12,
	0x2b, 0x12, 0x41, 0xeb, 0x96, 0xe5, 0xa5, 0x53, 0xf6, 0xa7, 0xc4, 0x5c, 0xbd, 0x7b, 0x6b, 0xec,
	0x12, 0x5e, 0x35, 0xce, 0xbd, 0xf0, 0xab, 0x07, 0xa7, 0x33, 0x26, 0x76, 0xd6, 0xee, 0xcd, 0x9e,
	0xff, 0xa7, 0x75, 0xdf, 0xdb, 0x69, 0x07, 0xbd, 0xaf, 0x5f, 0xb7, 0x71, 0xff, 0xba, 0xc1, 0xec,
	0xd7, 0xbe, 0x9f, 0xef, 0xfa, 0x1e, 0xee, 0xf8, 0x16, 0x79, 0xb5, 0xe5, 0x7a, 0xfa, 0xdd, 0x83,
	0x4e, 0x1d, 0x23, 0x13, 0x68, 0x7e, 0x58, 0xf0, 0x3b, 0xf2, 0x40, 0xa3, 0xc1, 0x5e, 0x91, 0xf0,
	0x88, 0xbc, 0x86, 0xb3, 0x8f, 0x22, 0x65, 0x1a, 0x6b, 0xfe, 0x5d, 0x71, 0xcd, 0x7f, 0x5b, 0x7c,
	0x0e, 0x83, 0x2b, 0xd4, 0x7f, 0xa2, 0xbc, 0x84, 0xd1, 0x15, 0xea, 0xb7, 0x4c, 0x27, 0x8b, 0x6d,
	0xf9, 0xa3, 0x03, 0x9b, 0x11, 0x8c, 0x1e, 0x78, 0xf2, 0xf0, 0x68, 0xde, 0xb6, 0xff, 0xf0, 0x97,
	0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x6b, 0xda, 0x42, 0xed, 0x05, 0x00, 0x00,
}