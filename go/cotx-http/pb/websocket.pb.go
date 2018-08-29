// Code generated by protoc-gen-go. DO NOT EDIT.
// source: websocket.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	websocket.proto

It has these top-level messages:
	ReqShadow
	PushMessage
	ResShadow
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

type ReqShadow struct {
	UserId int32 `protobuf:"varint,1,opt,name=UserId" json:"UserId"`
}

func (m *ReqShadow) Reset()                    { *m = ReqShadow{} }
func (m *ReqShadow) String() string            { return proto.CompactTextString(m) }
func (*ReqShadow) ProtoMessage()               {}
func (*ReqShadow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqShadow) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type PushMessage struct {
	Code      int32  `protobuf:"varint,1,opt,name=Code" json:"Code"`
	Value     int32  `protobuf:"varint,2,opt,name=Value" json:"Value"`
	State     int32  `protobuf:"varint,3,opt,name=State" json:"State"`
	Describe  string `protobuf:"bytes,4,opt,name=Describe" json:"Describe"`
	GayewayId string `protobuf:"bytes,5,opt,name=GayewayId" json:"GayewayId"`
}

func (m *PushMessage) Reset()                    { *m = PushMessage{} }
func (m *PushMessage) String() string            { return proto.CompactTextString(m) }
func (*PushMessage) ProtoMessage()               {}
func (*PushMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PushMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PushMessage) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PushMessage) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *PushMessage) GetDescribe() string {
	if m != nil {
		return m.Describe
	}
	return ""
}

func (m *PushMessage) GetGayewayId() string {
	if m != nil {
		return m.GayewayId
	}
	return ""
}

type ResShadow struct {
	Message []*PushMessage `protobuf:"bytes,1,rep,name=Message" json:"Message"`
}

func (m *ResShadow) Reset()                    { *m = ResShadow{} }
func (m *ResShadow) String() string            { return proto.CompactTextString(m) }
func (*ResShadow) ProtoMessage()               {}
func (*ResShadow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResShadow) GetMessage() []*PushMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqShadow)(nil), "pb.ReqShadow")
	proto.RegisterType((*PushMessage)(nil), "pb.PushMessage")
	proto.RegisterType((*ResShadow)(nil), "pb.ResShadow")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Websocket service

type WebsocketClient interface {
	GetPushMessage(ctx context.Context, in *ReqShadow, opts ...grpc.CallOption) (*ResShadow, error)
}

type websocketClient struct {
	cc *grpc.ClientConn
}

func NewWebsocketClient(cc *grpc.ClientConn) WebsocketClient {
	return &websocketClient{cc}
}

func (c *websocketClient) GetPushMessage(ctx context.Context, in *ReqShadow, opts ...grpc.CallOption) (*ResShadow, error) {
	out := new(ResShadow)
	err := grpc.Invoke(ctx, "/pb.Websocket/GetPushMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Websocket service

type WebsocketServer interface {
	GetPushMessage(context.Context, *ReqShadow) (*ResShadow, error)
}

func RegisterWebsocketServer(s *grpc.Server, srv WebsocketServer) {
	s.RegisterService(&_Websocket_serviceDesc, srv)
}

func _Websocket_GetPushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqShadow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsocketServer).GetPushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Websocket/GetPushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsocketServer).GetPushMessage(ctx, req.(*ReqShadow))
	}
	return interceptor(ctx, in, info, handler)
}

var _Websocket_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Websocket",
	HandlerType: (*WebsocketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPushMessage",
			Handler:    _Websocket_GetPushMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "websocket.proto",
}

func init() { proto.RegisterFile("websocket.proto", fileDescriptor14) }

var fileDescriptor14= []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4f, 0x4d, 0x2a,
	0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52,
	0x52, 0xe6, 0xe2, 0x0c, 0x4a, 0x2d, 0x0c, 0xce, 0x48, 0x4c, 0xc9, 0x2f, 0x17, 0x12, 0xe3, 0x62,
	0x0b, 0x2d, 0x4e, 0x2d, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2, 0x94,
	0xda, 0x19, 0xb9, 0xb8, 0x03, 0x4a, 0x8b, 0x33, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85,
	0x84, 0xb8, 0x58, 0x9c, 0xf3, 0x53, 0x52, 0xa1, 0xaa, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xb0,
	0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x20, 0x84, 0x03, 0x12, 0x0d, 0x2e, 0x49, 0x2c, 0x49,
	0x95, 0x60, 0x86, 0x88, 0x82, 0x39, 0x42, 0x52, 0x5c, 0x1c, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99,
	0x49, 0xa9, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x0c, 0x17, 0xa7, 0x7b,
	0x62, 0x65, 0x6a, 0x79, 0x62, 0xa5, 0x67, 0x8a, 0x04, 0x2b, 0x58, 0x12, 0x21, 0xa0, 0x64, 0x06,
	0x72, 0x6e, 0x31, 0xd4, 0xb9, 0x9a, 0x5c, 0xec, 0x50, 0x17, 0x49, 0x30, 0x2a, 0x30, 0x6b, 0x70,
	0x1b, 0xf1, 0xeb, 0x15, 0x24, 0xe9, 0x21, 0x39, 0x34, 0x08, 0x26, 0x6f, 0x64, 0xcb, 0xc5, 0x19,
	0x0e, 0xf3, 0xbd, 0x90, 0x01, 0x17, 0x9f, 0x7b, 0x6a, 0x09, 0xb2, 0x87, 0x78, 0x41, 0x1a, 0xe1,
	0xe1, 0x20, 0x05, 0xe5, 0x42, 0xed, 0x51, 0x62, 0x48, 0x62, 0x03, 0x07, 0x98, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xda, 0x71, 0xb6, 0x77, 0x43, 0x01, 0x00, 0x00,
}
