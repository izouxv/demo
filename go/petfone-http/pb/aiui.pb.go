// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/aiui.proto

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

type SemanticsType int32

const (
	SemanticsType_All   SemanticsType = 0
	SemanticsType_Text  SemanticsType = 1
	SemanticsType_Audio SemanticsType = 2
)

var SemanticsType_name = map[int32]string{
	0: "All",
	1: "Text",
	2: "Audio",
}
var SemanticsType_value = map[string]int32{
	"All":   0,
	"Text":  1,
	"Audio": 2,
}

func (x SemanticsType) String() string {
	return proto.EnumName(SemanticsType_name, int32(x))
}
func (SemanticsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aiui_8fa1794057b4cc63, []int{0}
}

type LanguageType int32

const (
	LanguageType_All_language LanguageType = 0
	LanguageType_Chiness      LanguageType = 1
	LanguageType_English      LanguageType = 2
)

var LanguageType_name = map[int32]string{
	0: "All_language",
	1: "Chiness",
	2: "English",
}
var LanguageType_value = map[string]int32{
	"All_language": 0,
	"Chiness":      1,
	"English":      2,
}

func (x LanguageType) String() string {
	return proto.EnumName(LanguageType_name, int32(x))
}
func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aiui_8fa1794057b4cc63, []int{1}
}

type GetTextSemanticsRequest struct {
	Types                SemanticsType `protobuf:"varint,1,opt,name=types,proto3,enum=pb.SemanticsType" json:"types,omitempty"`
	Language             LanguageType  `protobuf:"varint,2,opt,name=language,proto3,enum=pb.LanguageType" json:"language,omitempty"`
	Input                string        `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetTextSemanticsRequest) Reset()         { *m = GetTextSemanticsRequest{} }
func (m *GetTextSemanticsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTextSemanticsRequest) ProtoMessage()    {}
func (*GetTextSemanticsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aiui_8fa1794057b4cc63, []int{0}
}
func (m *GetTextSemanticsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTextSemanticsRequest.Unmarshal(m, b)
}
func (m *GetTextSemanticsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTextSemanticsRequest.Marshal(b, m, deterministic)
}
func (dst *GetTextSemanticsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTextSemanticsRequest.Merge(dst, src)
}
func (m *GetTextSemanticsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTextSemanticsRequest.Size(m)
}
func (m *GetTextSemanticsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTextSemanticsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTextSemanticsRequest proto.InternalMessageInfo

func (m *GetTextSemanticsRequest) GetTypes() SemanticsType {
	if m != nil {
		return m.Types
	}
	return SemanticsType_All
}

func (m *GetTextSemanticsRequest) GetLanguage() LanguageType {
	if m != nil {
		return m.Language
	}
	return LanguageType_All_language
}

func (m *GetTextSemanticsRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type GetTextSemanticsResponse struct {
	Semantics            string   `protobuf:"bytes,1,opt,name=Semantics,proto3" json:"Semantics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTextSemanticsResponse) Reset()         { *m = GetTextSemanticsResponse{} }
func (m *GetTextSemanticsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTextSemanticsResponse) ProtoMessage()    {}
func (*GetTextSemanticsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aiui_8fa1794057b4cc63, []int{1}
}
func (m *GetTextSemanticsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTextSemanticsResponse.Unmarshal(m, b)
}
func (m *GetTextSemanticsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTextSemanticsResponse.Marshal(b, m, deterministic)
}
func (dst *GetTextSemanticsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTextSemanticsResponse.Merge(dst, src)
}
func (m *GetTextSemanticsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTextSemanticsResponse.Size(m)
}
func (m *GetTextSemanticsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTextSemanticsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTextSemanticsResponse proto.InternalMessageInfo

func (m *GetTextSemanticsResponse) GetSemantics() string {
	if m != nil {
		return m.Semantics
	}
	return ""
}

type GetAudioSemanticsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAudioSemanticsRequest) Reset()         { *m = GetAudioSemanticsRequest{} }
func (m *GetAudioSemanticsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAudioSemanticsRequest) ProtoMessage()    {}
func (*GetAudioSemanticsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aiui_8fa1794057b4cc63, []int{2}
}
func (m *GetAudioSemanticsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAudioSemanticsRequest.Unmarshal(m, b)
}
func (m *GetAudioSemanticsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAudioSemanticsRequest.Marshal(b, m, deterministic)
}
func (dst *GetAudioSemanticsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAudioSemanticsRequest.Merge(dst, src)
}
func (m *GetAudioSemanticsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAudioSemanticsRequest.Size(m)
}
func (m *GetAudioSemanticsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAudioSemanticsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAudioSemanticsRequest proto.InternalMessageInfo

type GetAudioSemanticsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAudioSemanticsResponse) Reset()         { *m = GetAudioSemanticsResponse{} }
func (m *GetAudioSemanticsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAudioSemanticsResponse) ProtoMessage()    {}
func (*GetAudioSemanticsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aiui_8fa1794057b4cc63, []int{3}
}
func (m *GetAudioSemanticsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAudioSemanticsResponse.Unmarshal(m, b)
}
func (m *GetAudioSemanticsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAudioSemanticsResponse.Marshal(b, m, deterministic)
}
func (dst *GetAudioSemanticsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAudioSemanticsResponse.Merge(dst, src)
}
func (m *GetAudioSemanticsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAudioSemanticsResponse.Size(m)
}
func (m *GetAudioSemanticsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAudioSemanticsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAudioSemanticsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTextSemanticsRequest)(nil), "pb.GetTextSemanticsRequest")
	proto.RegisterType((*GetTextSemanticsResponse)(nil), "pb.GetTextSemanticsResponse")
	proto.RegisterType((*GetAudioSemanticsRequest)(nil), "pb.GetAudioSemanticsRequest")
	proto.RegisterType((*GetAudioSemanticsResponse)(nil), "pb.GetAudioSemanticsResponse")
	proto.RegisterEnum("pb.SemanticsType", SemanticsType_name, SemanticsType_value)
	proto.RegisterEnum("pb.LanguageType", LanguageType_name, LanguageType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AiuiServerClient is the client API for AiuiServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AiuiServerClient interface {
	GetTextSemantics(ctx context.Context, in *GetTextSemanticsRequest, opts ...grpc.CallOption) (*GetTextSemanticsResponse, error)
	GetAudioSemantics(ctx context.Context, in *GetAudioSemanticsRequest, opts ...grpc.CallOption) (*GetAudioSemanticsResponse, error)
}

type aiuiServerClient struct {
	cc *grpc.ClientConn
}

func NewAiuiServerClient(cc *grpc.ClientConn) AiuiServerClient {
	return &aiuiServerClient{cc}
}

func (c *aiuiServerClient) GetTextSemantics(ctx context.Context, in *GetTextSemanticsRequest, opts ...grpc.CallOption) (*GetTextSemanticsResponse, error) {
	out := new(GetTextSemanticsResponse)
	err := c.cc.Invoke(ctx, "/pb.AiuiServer/GetTextSemantics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aiuiServerClient) GetAudioSemantics(ctx context.Context, in *GetAudioSemanticsRequest, opts ...grpc.CallOption) (*GetAudioSemanticsResponse, error) {
	out := new(GetAudioSemanticsResponse)
	err := c.cc.Invoke(ctx, "/pb.AiuiServer/GetAudioSemantics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AiuiServerServer is the server API for AiuiServer service.
type AiuiServerServer interface {
	GetTextSemantics(context.Context, *GetTextSemanticsRequest) (*GetTextSemanticsResponse, error)
	GetAudioSemantics(context.Context, *GetAudioSemanticsRequest) (*GetAudioSemanticsResponse, error)
}

func RegisterAiuiServerServer(s *grpc.Server, srv AiuiServerServer) {
	s.RegisterService(&_AiuiServer_serviceDesc, srv)
}

func _AiuiServer_GetTextSemantics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTextSemanticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiuiServerServer).GetTextSemantics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AiuiServer/GetTextSemantics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiuiServerServer).GetTextSemantics(ctx, req.(*GetTextSemanticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AiuiServer_GetAudioSemantics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAudioSemanticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiuiServerServer).GetAudioSemantics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AiuiServer/GetAudioSemantics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiuiServerServer).GetAudioSemantics(ctx, req.(*GetAudioSemanticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AiuiServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AiuiServer",
	HandlerType: (*AiuiServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTextSemantics",
			Handler:    _AiuiServer_GetTextSemantics_Handler,
		},
		{
			MethodName: "GetAudioSemantics",
			Handler:    _AiuiServer_GetAudioSemantics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/aiui.proto",
}

func init() { proto.RegisterFile("pb/aiui.proto", fileDescriptor_aiui_8fa1794057b4cc63) }

var fileDescriptor_aiui_8fa1794057b4cc63 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xd9, 0x22, 0x42, 0x47, 0x30, 0xcb, 0xc4, 0xc4, 0x0a, 0x98, 0x90, 0x5e, 0x24, 0x44,
	0x31, 0xc1, 0x8b, 0xf1, 0xd6, 0x18, 0xe3, 0xc5, 0xc4, 0xa4, 0x70, 0x37, 0xad, 0x4e, 0x60, 0x93,
	0xda, 0xae, 0xec, 0xae, 0x91, 0x17, 0xf0, 0x7d, 0x7c, 0x43, 0xd3, 0xad, 0xa0, 0x42, 0x39, 0xce,
	0xfc, 0xff, 0x4c, 0xfe, 0x6f, 0x76, 0xa1, 0x25, 0xe3, 0xcb, 0x48, 0x18, 0x31, 0x92, 0x8b, 0x4c,
	0x67, 0xe8, 0xc8, 0xd8, 0xff, 0x64, 0x70, 0x7c, 0x4f, 0x7a, 0x4a, 0x1f, 0x7a, 0x42, 0xaf, 0x51,
	0xaa, 0xc5, 0xb3, 0x0a, 0xe9, 0xcd, 0x90, 0xd2, 0x78, 0x06, 0x35, 0xbd, 0x94, 0xa4, 0x3c, 0xd6,
	0x67, 0x83, 0xc3, 0x71, 0x7b, 0x24, 0xe3, 0xd1, 0xda, 0x34, 0x5d, 0x4a, 0x0a, 0x0b, 0x1d, 0xcf,
	0xa1, 0x91, 0x44, 0xe9, 0xcc, 0x44, 0x33, 0xf2, 0x1c, 0xeb, 0xe5, 0xb9, 0xf7, 0xe1, 0xa7, 0x67,
	0xad, 0x6b, 0x07, 0x1e, 0x41, 0x4d, 0xa4, 0xd2, 0x68, 0xaf, 0xda, 0x67, 0x03, 0x37, 0x2c, 0x0a,
	0xff, 0x1a, 0xbc, 0xed, 0x1c, 0x4a, 0x66, 0xa9, 0x22, 0xec, 0x81, 0xbb, 0x6e, 0xda, 0x30, 0x6e,
	0xf8, 0xdb, 0xf0, 0x3b, 0x76, 0x32, 0x30, 0x2f, 0x22, 0xdb, 0x44, 0xf0, 0xbb, 0x70, 0x52, 0xa2,
	0x15, 0x6b, 0x87, 0x17, 0xd0, 0xfa, 0x87, 0x83, 0x75, 0xa8, 0x06, 0x49, 0xc2, 0x2b, 0xd8, 0x80,
	0xbd, 0x3c, 0x09, 0x67, 0xe8, 0x42, 0xcd, 0x4e, 0x73, 0x67, 0x78, 0x03, 0xcd, 0xbf, 0x44, 0xc8,
	0xa1, 0x19, 0x24, 0xc9, 0xd3, 0x8a, 0x8b, 0x57, 0xf0, 0x00, 0xea, 0xb7, 0x73, 0x91, 0x92, 0x52,
	0x9c, 0xe5, 0xc5, 0x5d, 0x3a, 0x4b, 0x84, 0x9a, 0x73, 0x67, 0xfc, 0xc5, 0x00, 0x02, 0x61, 0xc4,
	0x84, 0x16, 0xef, 0xb4, 0xc0, 0x47, 0xe0, 0x9b, 0xb0, 0xd8, 0xcd, 0x4f, 0xb6, 0xe3, 0x29, 0x3a,
	0xbd, 0x72, 0xb1, 0x00, 0xf1, 0x2b, 0x18, 0x42, 0x7b, 0x8b, 0x13, 0x57, 0x43, 0xa5, 0xa7, 0xe9,
	0x9c, 0xee, 0x50, 0x57, 0x3b, 0xe3, 0x7d, 0xfb, 0x4b, 0xae, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xa8, 0xd1, 0xd2, 0x6d, 0x36, 0x02, 0x00, 0x00,
}
