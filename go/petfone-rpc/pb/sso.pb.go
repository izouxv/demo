// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/sso.proto

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

// 用户信息
type SsoRequest struct {
	Username             string     `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string     `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	SessionName          string     `protobuf:"bytes,3,opt,name=sessionName,proto3" json:"sessionName,omitempty"`
	Salt                 string     `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	Uid                  int32      `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	State                int32      `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	Exptime              int32      `protobuf:"varint,7,opt,name=exptime,proto3" json:"exptime,omitempty"`
	Nickname             string     `protobuf:"bytes,8,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Token                string     `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	Code                 string     `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	CodeType             int32      `protobuf:"varint,11,opt,name=codeType,proto3" json:"codeType,omitempty"`
	Source               string     `protobuf:"bytes,12,opt,name=source,proto3" json:"source,omitempty"`
	AgentInfo            *AgentInfo `protobuf:"bytes,13,opt,name=agentInfo,proto3" json:"agentInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SsoRequest) Reset()         { *m = SsoRequest{} }
func (m *SsoRequest) String() string { return proto.CompactTextString(m) }
func (*SsoRequest) ProtoMessage()    {}
func (*SsoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sso_78f863fbbc41d0a3, []int{0}
}
func (m *SsoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SsoRequest.Unmarshal(m, b)
}
func (m *SsoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SsoRequest.Marshal(b, m, deterministic)
}
func (dst *SsoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SsoRequest.Merge(dst, src)
}
func (m *SsoRequest) XXX_Size() int {
	return xxx_messageInfo_SsoRequest.Size(m)
}
func (m *SsoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SsoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SsoRequest proto.InternalMessageInfo

func (m *SsoRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SsoRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SsoRequest) GetSessionName() string {
	if m != nil {
		return m.SessionName
	}
	return ""
}

func (m *SsoRequest) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *SsoRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SsoRequest) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SsoRequest) GetExptime() int32 {
	if m != nil {
		return m.Exptime
	}
	return 0
}

func (m *SsoRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *SsoRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SsoRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SsoRequest) GetCodeType() int32 {
	if m != nil {
		return m.CodeType
	}
	return 0
}

func (m *SsoRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *SsoRequest) GetAgentInfo() *AgentInfo {
	if m != nil {
		return m.AgentInfo
	}
	return nil
}

type SsoReply struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	LoginState           int32    `protobuf:"varint,3,opt,name=loginState,proto3" json:"loginState,omitempty"`
	State                int32    `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	Code                 int32    `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
	SessionName          string   `protobuf:"bytes,6,opt,name=sessionName,proto3" json:"sessionName,omitempty"`
	Nickname             string   `protobuf:"bytes,7,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Token                string   `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SsoReply) Reset()         { *m = SsoReply{} }
func (m *SsoReply) String() string { return proto.CompactTextString(m) }
func (*SsoReply) ProtoMessage()    {}
func (*SsoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_sso_78f863fbbc41d0a3, []int{1}
}
func (m *SsoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SsoReply.Unmarshal(m, b)
}
func (m *SsoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SsoReply.Marshal(b, m, deterministic)
}
func (dst *SsoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SsoReply.Merge(dst, src)
}
func (m *SsoReply) XXX_Size() int {
	return xxx_messageInfo_SsoReply.Size(m)
}
func (m *SsoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SsoReply.DiscardUnknown(m)
}

var xxx_messageInfo_SsoReply proto.InternalMessageInfo

func (m *SsoReply) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SsoReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SsoReply) GetLoginState() int32 {
	if m != nil {
		return m.LoginState
	}
	return 0
}

func (m *SsoReply) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SsoReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SsoReply) GetSessionName() string {
	if m != nil {
		return m.SessionName
	}
	return ""
}

func (m *SsoReply) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *SsoReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type MultiSsoRequest struct {
	Ssos                 map[int32]*SsoRequest `protobuf:"bytes,1,rep,name=ssos,proto3" json:"ssos,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Source               string                `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MultiSsoRequest) Reset()         { *m = MultiSsoRequest{} }
func (m *MultiSsoRequest) String() string { return proto.CompactTextString(m) }
func (*MultiSsoRequest) ProtoMessage()    {}
func (*MultiSsoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sso_78f863fbbc41d0a3, []int{2}
}
func (m *MultiSsoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiSsoRequest.Unmarshal(m, b)
}
func (m *MultiSsoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiSsoRequest.Marshal(b, m, deterministic)
}
func (dst *MultiSsoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSsoRequest.Merge(dst, src)
}
func (m *MultiSsoRequest) XXX_Size() int {
	return xxx_messageInfo_MultiSsoRequest.Size(m)
}
func (m *MultiSsoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSsoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSsoRequest proto.InternalMessageInfo

func (m *MultiSsoRequest) GetSsos() map[int32]*SsoRequest {
	if m != nil {
		return m.Ssos
	}
	return nil
}

func (m *MultiSsoRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type MapSsoReply struct {
	Ssos                 map[int32]*SsoReply `protobuf:"bytes,1,rep,name=ssos,proto3" json:"ssos,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Code                 int32               `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MapSsoReply) Reset()         { *m = MapSsoReply{} }
func (m *MapSsoReply) String() string { return proto.CompactTextString(m) }
func (*MapSsoReply) ProtoMessage()    {}
func (*MapSsoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_sso_78f863fbbc41d0a3, []int{3}
}
func (m *MapSsoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapSsoReply.Unmarshal(m, b)
}
func (m *MapSsoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapSsoReply.Marshal(b, m, deterministic)
}
func (dst *MapSsoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapSsoReply.Merge(dst, src)
}
func (m *MapSsoReply) XXX_Size() int {
	return xxx_messageInfo_MapSsoReply.Size(m)
}
func (m *MapSsoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MapSsoReply.DiscardUnknown(m)
}

var xxx_messageInfo_MapSsoReply proto.InternalMessageInfo

func (m *MapSsoReply) GetSsos() map[int32]*SsoReply {
	if m != nil {
		return m.Ssos
	}
	return nil
}

func (m *MapSsoReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*SsoRequest)(nil), "pb.SsoRequest")
	proto.RegisterType((*SsoReply)(nil), "pb.SsoReply")
	proto.RegisterType((*MultiSsoRequest)(nil), "pb.MultiSsoRequest")
	proto.RegisterMapType((map[int32]*SsoRequest)(nil), "pb.MultiSsoRequest.SsosEntry")
	proto.RegisterType((*MapSsoReply)(nil), "pb.MapSsoReply")
	proto.RegisterMapType((map[int32]*SsoReply)(nil), "pb.MapSsoReply.SsosEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SsoClient is the client API for Sso service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SsoClient interface {
	GetUserInfo(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	Login(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	GetUserByName(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	Add(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	CheckPassword(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	UpdatePassword(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	Logout(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	UpdatePasswordByName(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	FindPasswordByMail(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	ResetPassword(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	UpdateState(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	GetBatchSsoInfos(ctx context.Context, in *MultiSsoRequest, opts ...grpc.CallOption) (*MapSsoReply, error)
	CheckCode(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	ResetPasswordByPhone(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
	SendMobileCode(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error)
}

type ssoClient struct {
	cc *grpc.ClientConn
}

func NewSsoClient(cc *grpc.ClientConn) SsoClient {
	return &ssoClient{cc}
}

func (c *ssoClient) GetUserInfo(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) Login(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) GetUserByName(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) Add(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) CheckPassword(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/CheckPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) UpdatePassword(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) Logout(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) UpdatePasswordByName(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/UpdatePasswordByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) FindPasswordByMail(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/FindPasswordByMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) ResetPassword(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) UpdateState(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) GetBatchSsoInfos(ctx context.Context, in *MultiSsoRequest, opts ...grpc.CallOption) (*MapSsoReply, error) {
	out := new(MapSsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/GetBatchSsoInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) CheckCode(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/CheckCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) ResetPasswordByPhone(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/ResetPasswordByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoClient) SendMobileCode(ctx context.Context, in *SsoRequest, opts ...grpc.CallOption) (*SsoReply, error) {
	out := new(SsoReply)
	err := c.cc.Invoke(ctx, "/pb.Sso/SendMobileCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SsoServer is the server API for Sso service.
type SsoServer interface {
	GetUserInfo(context.Context, *SsoRequest) (*SsoReply, error)
	Login(context.Context, *SsoRequest) (*SsoReply, error)
	GetUserByName(context.Context, *SsoRequest) (*SsoReply, error)
	Add(context.Context, *SsoRequest) (*SsoReply, error)
	CheckPassword(context.Context, *SsoRequest) (*SsoReply, error)
	UpdatePassword(context.Context, *SsoRequest) (*SsoReply, error)
	Logout(context.Context, *SsoRequest) (*SsoReply, error)
	UpdatePasswordByName(context.Context, *SsoRequest) (*SsoReply, error)
	FindPasswordByMail(context.Context, *SsoRequest) (*SsoReply, error)
	ResetPassword(context.Context, *SsoRequest) (*SsoReply, error)
	UpdateState(context.Context, *SsoRequest) (*SsoReply, error)
	GetBatchSsoInfos(context.Context, *MultiSsoRequest) (*MapSsoReply, error)
	CheckCode(context.Context, *SsoRequest) (*SsoReply, error)
	ResetPasswordByPhone(context.Context, *SsoRequest) (*SsoReply, error)
	SendMobileCode(context.Context, *SsoRequest) (*SsoReply, error)
}

func RegisterSsoServer(s *grpc.Server, srv SsoServer) {
	s.RegisterService(&_Sso_serviceDesc, srv)
}

func _Sso_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).GetUserInfo(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).Login(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).GetUserByName(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).Add(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).CheckPassword(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).UpdatePassword(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).Logout(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_UpdatePasswordByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).UpdatePasswordByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/UpdatePasswordByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).UpdatePasswordByName(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_FindPasswordByMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).FindPasswordByMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/FindPasswordByMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).FindPasswordByMail(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).ResetPassword(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).UpdateState(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_GetBatchSsoInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiSsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).GetBatchSsoInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/GetBatchSsoInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).GetBatchSsoInfos(ctx, req.(*MultiSsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_CheckCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).CheckCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/CheckCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).CheckCode(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_ResetPasswordByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).ResetPasswordByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/ResetPasswordByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).ResetPasswordByPhone(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sso_SendMobileCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).SendMobileCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sso/SendMobileCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).SendMobileCode(ctx, req.(*SsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sso_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Sso",
	HandlerType: (*SsoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Sso_GetUserInfo_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Sso_Login_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _Sso_GetUserByName_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Sso_Add_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _Sso_CheckPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Sso_UpdatePassword_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Sso_Logout_Handler,
		},
		{
			MethodName: "UpdatePasswordByName",
			Handler:    _Sso_UpdatePasswordByName_Handler,
		},
		{
			MethodName: "FindPasswordByMail",
			Handler:    _Sso_FindPasswordByMail_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Sso_ResetPassword_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _Sso_UpdateState_Handler,
		},
		{
			MethodName: "GetBatchSsoInfos",
			Handler:    _Sso_GetBatchSsoInfos_Handler,
		},
		{
			MethodName: "CheckCode",
			Handler:    _Sso_CheckCode_Handler,
		},
		{
			MethodName: "ResetPasswordByPhone",
			Handler:    _Sso_ResetPasswordByPhone_Handler,
		},
		{
			MethodName: "SendMobileCode",
			Handler:    _Sso_SendMobileCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/sso.proto",
}

func init() { proto.RegisterFile("pb/sso.proto", fileDescriptor_sso_78f863fbbc41d0a3) }

var fileDescriptor_sso_78f863fbbc41d0a3 = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xeb, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0xd2, 0xcb, 0xda, 0x93, 0xdd, 0x64, 0x26, 0x64, 0x2a, 0x81, 0xaa, 0x08, 0x44, 0x25,
	0xb4, 0x0e, 0x0a, 0x42, 0x68, 0xff, 0xb6, 0x69, 0x4c, 0x48, 0x2b, 0x9a, 0x52, 0xf6, 0x00, 0x69,
	0x73, 0xb6, 0x46, 0xcd, 0x6c, 0x53, 0x3b, 0x40, 0x9e, 0x84, 0x27, 0xe0, 0x95, 0x78, 0x81, 0xbd,
	0x08, 0xb2, 0xbd, 0x34, 0x69, 0x07, 0x34, 0xbf, 0xea, 0x73, 0xf5, 0xf7, 0x7d, 0xe7, 0x38, 0x85,
	0x2d, 0x31, 0x3e, 0x94, 0x92, 0xf7, 0xc5, 0x9c, 0x2b, 0x4e, 0x5c, 0x31, 0xee, 0xec, 0x89, 0xf1,
	0xa1, 0x40, 0x75, 0xcd, 0x19, 0x5a, 0xaf, 0x7f, 0xe7, 0x02, 0x8c, 0x24, 0x0f, 0xf0, 0x6b, 0x8a,
	0x52, 0x91, 0x0e, 0xb4, 0x52, 0x89, 0x73, 0x16, 0xde, 0x22, 0x75, 0xba, 0x4e, 0xaf, 0x1d, 0x2c,
	0x6c, 0x1d, 0x13, 0xa1, 0x94, 0xdf, 0xf9, 0x3c, 0xa2, 0xae, 0x8d, 0xe5, 0x36, 0xe9, 0x82, 0x27,
	0x51, 0xca, 0x98, 0xb3, 0xcf, 0xba, 0xb4, 0x66, 0xc2, 0x65, 0x17, 0x21, 0x50, 0x97, 0x61, 0xa2,
	0x68, 0xdd, 0x84, 0xcc, 0x99, 0xec, 0x41, 0x2d, 0x8d, 0x23, 0xda, 0xe8, 0x3a, 0xbd, 0x46, 0xa0,
	0x8f, 0x64, 0x1f, 0x1a, 0x52, 0x85, 0x0a, 0x69, 0xd3, 0xf8, 0xac, 0x41, 0x28, 0x6c, 0xe2, 0x0f,
	0xa1, 0xe2, 0x5b, 0xa4, 0x9b, 0xc6, 0x9f, 0x9b, 0x1a, 0x13, 0x8b, 0x27, 0x33, 0x83, 0xb7, 0x65,
	0x31, 0xe5, 0xb6, 0xee, 0xa5, 0xf8, 0x0c, 0x19, 0x6d, 0x9b, 0x80, 0x35, 0x34, 0x8e, 0x09, 0x8f,
	0x90, 0x82, 0xc5, 0xa1, 0xcf, 0xba, 0x8b, 0xfe, 0xfd, 0x92, 0x09, 0xa4, 0x9e, 0xb9, 0x60, 0x61,
	0x93, 0xc7, 0xd0, 0x94, 0x3c, 0x9d, 0x4f, 0x90, 0x6e, 0x99, 0x8a, 0x7b, 0x8b, 0xbc, 0x82, 0x76,
	0x78, 0x83, 0x4c, 0x7d, 0x62, 0xd7, 0x9c, 0x6e, 0x77, 0x9d, 0x9e, 0x37, 0xd8, 0xee, 0x8b, 0x71,
	0xff, 0x38, 0x77, 0x06, 0x45, 0xdc, 0xff, 0xed, 0x40, 0xcb, 0xa8, 0x2c, 0x92, 0x2c, 0x67, 0xed,
	0x14, 0xac, 0xcb, 0xaa, 0xbb, 0x2b, 0xaa, 0x3f, 0x03, 0x48, 0xf8, 0x4d, 0xcc, 0x46, 0x46, 0x96,
	0x9a, 0x29, 0x2a, 0x79, 0x0a, 0xc5, 0xea, 0x65, 0xc5, 0x72, 0x96, 0x56, 0x5a, 0xcb, 0x72, 0x65,
	0x46, 0xcd, 0x87, 0x33, 0x2a, 0xab, 0xb9, 0xf9, 0x2f, 0x35, 0x5b, 0x25, 0x35, 0xfd, 0x5f, 0x0e,
	0xec, 0x0e, 0xd3, 0x44, 0xc5, 0xa5, 0x1d, 0x7a, 0x03, 0x75, 0x29, 0xb9, 0xa4, 0x4e, 0xb7, 0xd6,
	0xf3, 0x06, 0x4f, 0xb5, 0x28, 0x2b, 0x29, 0xfd, 0x91, 0xe4, 0xf2, 0x8c, 0xa9, 0x79, 0x16, 0x98,
	0xd4, 0x92, 0xc8, 0x6e, 0x59, 0xe4, 0xce, 0x39, 0xb4, 0x17, 0xa9, 0x5a, 0xb7, 0x19, 0x66, 0xb9,
	0x6e, 0x33, 0xcc, 0xc8, 0x73, 0x68, 0x7c, 0x0b, 0x93, 0xd4, 0x56, 0x79, 0x83, 0x1d, 0x7d, 0x55,
	0x71, 0x4b, 0x60, 0x83, 0x47, 0xee, 0x07, 0xc7, 0xff, 0xe9, 0x80, 0x37, 0x0c, 0xc5, 0x62, 0x06,
	0x07, 0x4b, 0x18, 0x9f, 0x18, 0x8c, 0x45, 0xf8, 0x01, 0xbe, 0x5c, 0x4e, 0xb7, 0x90, 0xb3, 0x73,
	0xf6, 0x7f, 0x6c, 0xfe, 0x32, 0xb6, 0xad, 0x05, 0x36, 0x91, 0x64, 0x25, 0x64, 0x83, 0xbb, 0x06,
	0xd4, 0x46, 0x92, 0x93, 0x03, 0xf0, 0xce, 0x51, 0x5d, 0x49, 0x9c, 0xeb, 0x8d, 0x21, 0x2b, 0x5c,
	0x3a, 0x4b, 0xf5, 0xfe, 0x06, 0x79, 0x09, 0x8d, 0x0b, 0xbd, 0x04, 0x6b, 0x13, 0x0f, 0x61, 0xfb,
	0xbe, 0xef, 0x49, 0x66, 0x86, 0xbc, 0xae, 0xe0, 0x05, 0xd4, 0x8e, 0xa3, 0xa8, 0x4a, 0xdf, 0xd3,
	0x29, 0x4e, 0x66, 0x97, 0xf9, 0x27, 0x60, 0x5d, 0xc1, 0x6b, 0xd8, 0xb9, 0x12, 0x51, 0xa8, 0xb0,
	0x72, 0x45, 0x0f, 0x9a, 0x17, 0xfc, 0x86, 0xa7, 0x6a, 0x6d, 0xe6, 0x7b, 0xd8, 0x5f, 0xee, 0x5d,
	0x91, 0xeb, 0x3b, 0x20, 0x1f, 0x63, 0x16, 0x15, 0x55, 0xc3, 0x30, 0x4e, 0xaa, 0x50, 0x0f, 0x50,
	0xa2, 0xaa, 0x4c, 0xe4, 0x00, 0x3c, 0x0b, 0xcf, 0x3e, 0xd9, 0x75, 0xe9, 0x47, 0xb0, 0x77, 0x8e,
	0xea, 0x24, 0x54, 0x93, 0xe9, 0x48, 0x72, 0xbd, 0x0e, 0x92, 0x3c, 0xfa, 0xcb, 0x33, 0xea, 0xec,
	0xae, 0xec, 0xad, 0xbf, 0xa1, 0x3f, 0x4b, 0x66, 0x2c, 0xa7, 0xfa, 0xc5, 0x57, 0x90, 0x6d, 0x89,
	0xc8, 0x49, 0x76, 0x39, 0xe5, 0x0c, 0xab, 0x8c, 0x72, 0x84, 0x2c, 0x1a, 0xf2, 0x71, 0x9c, 0x60,
	0x95, 0x9b, 0xc6, 0x4d, 0xf3, 0x6f, 0xf3, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x13,
	0x22, 0x3b, 0x93, 0x06, 0x00, 0x00,
}