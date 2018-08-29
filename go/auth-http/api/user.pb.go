// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/user.proto

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

type UpdateUserInfoRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Tid                  int32    `protobuf:"varint,3,opt,name=tid,proto3" json:"tid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfoRequest) Reset()         { *m = UpdateUserInfoRequest{} }
func (m *UpdateUserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfoRequest) ProtoMessage()    {}
func (*UpdateUserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{0}
}
func (m *UpdateUserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfoRequest.Unmarshal(m, b)
}
func (m *UpdateUserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfoRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfoRequest.Merge(dst, src)
}
func (m *UpdateUserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfoRequest.Size(m)
}
func (m *UpdateUserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfoRequest proto.InternalMessageInfo

func (m *UpdateUserInfoRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateUserInfoRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UpdateUserInfoRequest) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

type UpdateUserInfoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfoResponse) Reset()         { *m = UpdateUserInfoResponse{} }
func (m *UpdateUserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfoResponse) ProtoMessage()    {}
func (*UpdateUserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{1}
}
func (m *UpdateUserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfoResponse.Unmarshal(m, b)
}
func (m *UpdateUserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfoResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateUserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfoResponse.Merge(dst, src)
}
func (m *UpdateUserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfoResponse.Size(m)
}
func (m *UpdateUserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfoResponse proto.InternalMessageInfo

type UpdateNicknameAndPasswordRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Tid                  int32    `protobuf:"varint,3,opt,name=tid,proto3" json:"tid,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword          string   `protobuf:"bytes,5,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNicknameAndPasswordRequest) Reset()         { *m = UpdateNicknameAndPasswordRequest{} }
func (m *UpdateNicknameAndPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNicknameAndPasswordRequest) ProtoMessage()    {}
func (*UpdateNicknameAndPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{2}
}
func (m *UpdateNicknameAndPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNicknameAndPasswordRequest.Unmarshal(m, b)
}
func (m *UpdateNicknameAndPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNicknameAndPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateNicknameAndPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNicknameAndPasswordRequest.Merge(dst, src)
}
func (m *UpdateNicknameAndPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNicknameAndPasswordRequest.Size(m)
}
func (m *UpdateNicknameAndPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNicknameAndPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNicknameAndPasswordRequest proto.InternalMessageInfo

func (m *UpdateNicknameAndPasswordRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateNicknameAndPasswordRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UpdateNicknameAndPasswordRequest) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *UpdateNicknameAndPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdateNicknameAndPasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type UpdateNicknameAndPasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNicknameAndPasswordResponse) Reset()         { *m = UpdateNicknameAndPasswordResponse{} }
func (m *UpdateNicknameAndPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateNicknameAndPasswordResponse) ProtoMessage()    {}
func (*UpdateNicknameAndPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{3}
}
func (m *UpdateNicknameAndPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNicknameAndPasswordResponse.Unmarshal(m, b)
}
func (m *UpdateNicknameAndPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNicknameAndPasswordResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateNicknameAndPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNicknameAndPasswordResponse.Merge(dst, src)
}
func (m *UpdateNicknameAndPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateNicknameAndPasswordResponse.Size(m)
}
func (m *UpdateNicknameAndPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNicknameAndPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNicknameAndPasswordResponse proto.InternalMessageInfo

type UpdatePasswordRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	Tid                  int32    `protobuf:"varint,4,opt,name=tid,proto3" json:"tid,omitempty"`
	Did                  int32    `protobuf:"varint,5,opt,name=did,proto3" json:"did,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordRequest) Reset()         { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()    {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{4}
}
func (m *UpdatePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordRequest.Unmarshal(m, b)
}
func (m *UpdatePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordRequest.Merge(dst, src)
}
func (m *UpdatePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordRequest.Size(m)
}
func (m *UpdatePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordRequest proto.InternalMessageInfo

func (m *UpdatePasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdatePasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *UpdatePasswordRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

type UpdatePasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordResponse) Reset()         { *m = UpdatePasswordResponse{} }
func (m *UpdatePasswordResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordResponse) ProtoMessage()    {}
func (*UpdatePasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{5}
}
func (m *UpdatePasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordResponse.Unmarshal(m, b)
}
func (m *UpdatePasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordResponse.Marshal(b, m, deterministic)
}
func (dst *UpdatePasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordResponse.Merge(dst, src)
}
func (m *UpdatePasswordResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordResponse.Size(m)
}
func (m *UpdatePasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordResponse proto.InternalMessageInfo

type FindPasswordRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Tid                  int32    `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
	Did                  int32    `protobuf:"varint,3,opt,name=did,proto3" json:"did,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPasswordRequest) Reset()         { *m = FindPasswordRequest{} }
func (m *FindPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*FindPasswordRequest) ProtoMessage()    {}
func (*FindPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{6}
}
func (m *FindPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPasswordRequest.Unmarshal(m, b)
}
func (m *FindPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *FindPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPasswordRequest.Merge(dst, src)
}
func (m *FindPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_FindPasswordRequest.Size(m)
}
func (m *FindPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPasswordRequest proto.InternalMessageInfo

func (m *FindPasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FindPasswordRequest) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *FindPasswordRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

type FindPasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPasswordResponse) Reset()         { *m = FindPasswordResponse{} }
func (m *FindPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*FindPasswordResponse) ProtoMessage()    {}
func (*FindPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{7}
}
func (m *FindPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPasswordResponse.Unmarshal(m, b)
}
func (m *FindPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPasswordResponse.Marshal(b, m, deterministic)
}
func (dst *FindPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPasswordResponse.Merge(dst, src)
}
func (m *FindPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_FindPasswordResponse.Size(m)
}
func (m *FindPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindPasswordResponse proto.InternalMessageInfo

type ResetPasswordRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordRequest) Reset()         { *m = ResetPasswordRequest{} }
func (m *ResetPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordRequest) ProtoMessage()    {}
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{8}
}
func (m *ResetPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordRequest.Unmarshal(m, b)
}
func (m *ResetPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *ResetPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordRequest.Merge(dst, src)
}
func (m *ResetPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordRequest.Size(m)
}
func (m *ResetPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordRequest proto.InternalMessageInfo

func (m *ResetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPasswordRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ResetPasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordResponse) Reset()         { *m = ResetPasswordResponse{} }
func (m *ResetPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordResponse) ProtoMessage()    {}
func (*ResetPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{9}
}
func (m *ResetPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordResponse.Unmarshal(m, b)
}
func (m *ResetPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordResponse.Marshal(b, m, deterministic)
}
func (dst *ResetPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordResponse.Merge(dst, src)
}
func (m *ResetPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordResponse.Size(m)
}
func (m *ResetPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordResponse proto.InternalMessageInfo

type UpdateUserStateRequest struct {
	UpdateUid            int32    `protobuf:"varint,2,opt,name=updateUid,proto3" json:"updateUid,omitempty"`
	UpdateState          int32    `protobuf:"varint,3,opt,name=updateState,proto3" json:"updateState,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserStateRequest) Reset()         { *m = UpdateUserStateRequest{} }
func (m *UpdateUserStateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserStateRequest) ProtoMessage()    {}
func (*UpdateUserStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{10}
}
func (m *UpdateUserStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserStateRequest.Unmarshal(m, b)
}
func (m *UpdateUserStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserStateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserStateRequest.Merge(dst, src)
}
func (m *UpdateUserStateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserStateRequest.Size(m)
}
func (m *UpdateUserStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserStateRequest proto.InternalMessageInfo

func (m *UpdateUserStateRequest) GetUpdateUid() int32 {
	if m != nil {
		return m.UpdateUid
	}
	return 0
}

func (m *UpdateUserStateRequest) GetUpdateState() int32 {
	if m != nil {
		return m.UpdateState
	}
	return 0
}

type UpdateUserStateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserStateResponse) Reset()         { *m = UpdateUserStateResponse{} }
func (m *UpdateUserStateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserStateResponse) ProtoMessage()    {}
func (*UpdateUserStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_91973effd4c5c16e, []int{11}
}
func (m *UpdateUserStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserStateResponse.Unmarshal(m, b)
}
func (m *UpdateUserStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserStateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateUserStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserStateResponse.Merge(dst, src)
}
func (m *UpdateUserStateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserStateResponse.Size(m)
}
func (m *UpdateUserStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserStateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateUserInfoRequest)(nil), "api.UpdateUserInfoRequest")
	proto.RegisterType((*UpdateUserInfoResponse)(nil), "api.UpdateUserInfoResponse")
	proto.RegisterType((*UpdateNicknameAndPasswordRequest)(nil), "api.UpdateNicknameAndPasswordRequest")
	proto.RegisterType((*UpdateNicknameAndPasswordResponse)(nil), "api.UpdateNicknameAndPasswordResponse")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "api.UpdatePasswordRequest")
	proto.RegisterType((*UpdatePasswordResponse)(nil), "api.UpdatePasswordResponse")
	proto.RegisterType((*FindPasswordRequest)(nil), "api.FindPasswordRequest")
	proto.RegisterType((*FindPasswordResponse)(nil), "api.FindPasswordResponse")
	proto.RegisterType((*ResetPasswordRequest)(nil), "api.ResetPasswordRequest")
	proto.RegisterType((*ResetPasswordResponse)(nil), "api.ResetPasswordResponse")
	proto.RegisterType((*UpdateUserStateRequest)(nil), "api.UpdateUserStateRequest")
	proto.RegisterType((*UpdateUserStateResponse)(nil), "api.UpdateUserStateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServerClient is the client API for UserServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServerClient interface {
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	UpdateNicknameAndPassword(ctx context.Context, in *UpdateNicknameAndPasswordRequest, opts ...grpc.CallOption) (*UpdateNicknameAndPasswordResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	FindPassword(ctx context.Context, in *FindPasswordRequest, opts ...grpc.CallOption) (*FindPasswordResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	UpdateUserState(ctx context.Context, in *UpdateUserStateRequest, opts ...grpc.CallOption) (*UpdateUserStateResponse, error)
}

type userServerClient struct {
	cc *grpc.ClientConn
}

func NewUserServerClient(cc *grpc.ClientConn) UserServerClient {
	return &userServerClient{cc}
}

func (c *userServerClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/api.UserServer/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) UpdateNicknameAndPassword(ctx context.Context, in *UpdateNicknameAndPasswordRequest, opts ...grpc.CallOption) (*UpdateNicknameAndPasswordResponse, error) {
	out := new(UpdateNicknameAndPasswordResponse)
	err := c.cc.Invoke(ctx, "/api.UserServer/UpdateNicknameAndPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/api.UserServer/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) FindPassword(ctx context.Context, in *FindPasswordRequest, opts ...grpc.CallOption) (*FindPasswordResponse, error) {
	out := new(FindPasswordResponse)
	err := c.cc.Invoke(ctx, "/api.UserServer/FindPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/api.UserServer/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) UpdateUserState(ctx context.Context, in *UpdateUserStateRequest, opts ...grpc.CallOption) (*UpdateUserStateResponse, error) {
	out := new(UpdateUserStateResponse)
	err := c.cc.Invoke(ctx, "/api.UserServer/UpdateUserState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServer is the server API for UserServer service.
type UserServerServer interface {
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	UpdateNicknameAndPassword(context.Context, *UpdateNicknameAndPasswordRequest) (*UpdateNicknameAndPasswordResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	FindPassword(context.Context, *FindPasswordRequest) (*FindPasswordResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	UpdateUserState(context.Context, *UpdateUserStateRequest) (*UpdateUserStateResponse, error)
}

func RegisterUserServerServer(s *grpc.Server, srv UserServerServer) {
	s.RegisterService(&_UserServer_serviceDesc, srv)
}

func _UserServer_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserServer/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_UpdateNicknameAndPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNicknameAndPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).UpdateNicknameAndPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserServer/UpdateNicknameAndPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).UpdateNicknameAndPassword(ctx, req.(*UpdateNicknameAndPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserServer/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_FindPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).FindPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserServer/FindPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).FindPassword(ctx, req.(*FindPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserServer/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_UpdateUserState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).UpdateUserState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserServer/UpdateUserState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).UpdateUserState(ctx, req.(*UpdateUserStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserServer",
	HandlerType: (*UserServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserServer_UpdateUserInfo_Handler,
		},
		{
			MethodName: "UpdateNicknameAndPassword",
			Handler:    _UserServer_UpdateNicknameAndPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserServer_UpdatePassword_Handler,
		},
		{
			MethodName: "FindPassword",
			Handler:    _UserServer_FindPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserServer_ResetPassword_Handler,
		},
		{
			MethodName: "UpdateUserState",
			Handler:    _UserServer_UpdateUserState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user.proto",
}

func init() { proto.RegisterFile("api/user.proto", fileDescriptor_user_91973effd4c5c16e) }

var fileDescriptor_user_91973effd4c5c16e = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x2d, 0xc6, 0x54, 0xf6, 0xb4, 0x75, 0xad, 0xad, 0x3f, 0xf0, 0xda, 0x07, 0x4a, 0xd5, 0xca,
	0x27, 0x57, 0x6a, 0x7f, 0x41, 0x0f, 0x8d, 0x1c, 0x45, 0xb2, 0x22, 0x22, 0x2b, 0xb9, 0x92, 0xb0,
	0x91, 0x56, 0x4e, 0x80, 0xb0, 0x4b, 0xfc, 0x4b, 0x72, 0xcf, 0x9f, 0xcc, 0x3d, 0x82, 0x5d, 0x96,
	0xaf, 0xb5, 0x73, 0xc9, 0x8d, 0xd9, 0x37, 0xfb, 0xde, 0x9b, 0xc7, 0x00, 0x0c, 0xfc, 0x98, 0xfe,
	0x4e, 0x19, 0x49, 0x56, 0x71, 0x12, 0xf1, 0x08, 0x99, 0x7e, 0x4c, 0xdd, 0x4b, 0x18, 0x6f, 0xe3,
	0xc0, 0xe7, 0x64, 0xcb, 0x48, 0x72, 0x1a, 0xde, 0x46, 0x1e, 0x79, 0x48, 0x09, 0xe3, 0x68, 0x08,
	0x66, 0x4a, 0x03, 0xdb, 0x70, 0x8c, 0xa5, 0xe5, 0x65, 0x8f, 0x08, 0x43, 0x2f, 0xa4, 0x37, 0xbb,
	0xd0, 0xbf, 0x27, 0x76, 0xc7, 0x31, 0x96, 0x7d, 0x4f, 0xd5, 0x59, 0x37, 0xa7, 0x81, 0x6d, 0x8a,
	0x6e, 0x4e, 0x03, 0xd7, 0x86, 0x49, 0x93, 0x98, 0xc5, 0x51, 0xc8, 0x88, 0xfb, 0x6c, 0x80, 0x23,
	0xa0, 0x8d, 0xbc, 0xfe, 0x2f, 0x0c, 0xce, 0x7d, 0xc6, 0xf6, 0x51, 0x12, 0xbc, 0x93, 0x7c, 0xd6,
	0x1d, 0x4b, 0x4a, 0xbb, 0x2b, 0xba, 0x8b, 0x1a, 0x39, 0xf0, 0x29, 0x24, 0xfb, 0x42, 0xd1, 0xb6,
	0x72, 0xb8, 0x7a, 0xe4, 0xfe, 0x80, 0xef, 0x47, 0x1c, 0xca, 0x39, 0x9e, 0x8c, 0x22, 0xbb, 0xa6,
	0x79, 0x0c, 0xbd, 0x2c, 0xe7, 0xdc, 0xaa, 0x21, 0xc4, 0x8b, 0xba, 0x66, 0xac, 0x73, 0xdc, 0x98,
	0xd9, 0x32, 0x56, 0x0c, 0xda, 0x2d, 0x07, 0x1d, 0x82, 0x19, 0x50, 0x31, 0x84, 0xe5, 0x65, 0x8f,
	0x65, 0xf2, 0x2d, 0xc7, 0x5b, 0xf8, 0x76, 0x42, 0xdb, 0x59, 0x1f, 0xb3, 0x2b, 0x05, 0x3b, 0x2d,
	0x41, 0xb3, 0x14, 0x9c, 0xc0, 0xa8, 0x4e, 0x2b, 0xe5, 0xd6, 0x30, 0xf2, 0x08, 0x23, 0x5c, 0xa3,
	0xa7, 0x22, 0x30, 0x1a, 0x11, 0x8c, 0xc0, 0xe2, 0xd1, 0x8e, 0x84, 0x32, 0x1b, 0x51, 0xb8, 0x53,
	0x18, 0x37, 0x98, 0xa4, 0xc4, 0x55, 0x75, 0xcb, 0x2e, 0xb8, 0xcf, 0x49, 0x21, 0xb2, 0x80, 0x7e,
	0x2a, 0x10, 0x65, 0xbf, 0x3c, 0xc8, 0x92, 0x16, 0x45, 0x7e, 0x47, 0x0e, 0x53, 0x3d, 0x72, 0x67,
	0x30, 0x6d, 0x31, 0x0b, 0xd1, 0x3f, 0x2f, 0x26, 0x40, 0x7e, 0x4a, 0x92, 0x47, 0x92, 0xa0, 0x33,
	0x18, 0xd4, 0x37, 0x1d, 0xe1, 0x95, 0x1f, 0xd3, 0x95, 0xf6, 0xbb, 0xc2, 0x73, 0x2d, 0x26, 0xc7,
	0xf9, 0x80, 0xee, 0x60, 0x76, 0x70, 0xf3, 0xd0, 0xcf, 0xca, 0xdd, 0xc3, 0xdf, 0x0e, 0xfe, 0xf5,
	0x56, 0x9b, 0x52, 0x53, 0xd6, 0x95, 0x44, 0xd5, 0x7a, 0x93, 0x77, 0xae, 0xc5, 0x14, 0xd9, 0x7f,
	0xf8, 0x5c, 0x5d, 0x03, 0x64, 0xe7, 0xed, 0x9a, 0x85, 0xc3, 0x33, 0x0d, 0xa2, 0x68, 0xd6, 0xf0,
	0xa5, 0xf6, 0xae, 0x91, 0xe8, 0xd6, 0x6d, 0x12, 0xc6, 0x3a, 0x48, 0x31, 0x6d, 0xe0, 0x6b, 0xe3,
	0x15, 0xa2, 0x66, 0xfa, 0xd5, 0x95, 0xc1, 0x0b, 0x3d, 0x58, 0xf0, 0x5d, 0x7f, 0xcc, 0xff, 0x9b,
	0x7f, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x82, 0x2e, 0xf1, 0x49, 0x05, 0x00, 0x00,
}
