// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway_user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	gateway_user.proto

It has these top-level messages:
	ReqGwUser
	ResGwUser
	ResAccount
	ResAccounts
	ResShowNodeByGw
	ResShowNodesByGw
	ResNodePos
	ResShowAllNodesPos
	ResUserGateway
	ResUserGateways
	ResShowGateway
	ResShowAllGws
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

type ReqGwUser struct {
	UserID        int32  `protobuf:"varint,1,opt,name=UserID" json:"UserID"`
	MAC           string `protobuf:"bytes,2,opt,name=MAC" json:"MAC"`
	GatewayID     string `protobuf:"bytes,3,opt,name=GatewayID" json:"GatewayID"`
	AppEUI        string `protobuf:"bytes,4,opt,name=AppEUI" json:"AppEUI"`
	AccountID     int32  `protobuf:"varint,5,opt,name=AccountID" json:"AccountID"`
	UName         string `protobuf:"bytes,6,opt,name=UName" json:"UName"`
	ApplicationId int64  `protobuf:"varint,7,opt,name=ApplicationId" json:"ApplicationId"`
	BleName       string `protobuf:"bytes,8,opt,name=BleName" json:"BleName"`
}

func (m *ReqGwUser) Reset()                    { *m = ReqGwUser{} }
func (m *ReqGwUser) String() string            { return proto.CompactTextString(m) }
func (*ReqGwUser) ProtoMessage()               {}
func (*ReqGwUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqGwUser) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ReqGwUser) GetMAC() string {
	if m != nil {
		return m.MAC
	}
	return ""
}

func (m *ReqGwUser) GetGatewayID() string {
	if m != nil {
		return m.GatewayID
	}
	return ""
}

func (m *ReqGwUser) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *ReqGwUser) GetAccountID() int32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ReqGwUser) GetUName() string {
	if m != nil {
		return m.UName
	}
	return ""
}

func (m *ReqGwUser) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *ReqGwUser) GetBleName() string {
	if m != nil {
		return m.BleName
	}
	return ""
}

type ResGwUser struct {
	UserID    int32  `protobuf:"varint,1,opt,name=UserID" json:"UserID"`
	MAC       int64  `protobuf:"varint,2,opt,name=MAC" json:"MAC"`
	GatewayID string `protobuf:"bytes,3,opt,name=GatewayID" json:"GatewayID"`
	AppEUI    string `protobuf:"bytes,4,opt,name=AppEUI" json:"AppEUI"`
	AccountID int32  `protobuf:"varint,5,opt,name=AccountID" json:"AccountID"`
	ErrCode   int32  `protobuf:"varint,6,opt,name=ErrCode" json:"ErrCode"`
}

func (m *ResGwUser) Reset()                    { *m = ResGwUser{} }
func (m *ResGwUser) String() string            { return proto.CompactTextString(m) }
func (*ResGwUser) ProtoMessage()               {}
func (*ResGwUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResGwUser) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ResGwUser) GetMAC() int64 {
	if m != nil {
		return m.MAC
	}
	return 0
}

func (m *ResGwUser) GetGatewayID() string {
	if m != nil {
		return m.GatewayID
	}
	return ""
}

func (m *ResGwUser) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *ResGwUser) GetAccountID() int32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ResGwUser) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

type ResAccount struct {
	UserID   int32  `protobuf:"varint,1,opt,name=UserID" json:"UserID"`
	UserName string `protobuf:"bytes,2,opt,name=UserName" json:"UserName"`
	Acatar   int32  `protobuf:"varint,3,opt,name=Acatar" json:"Acatar"`
	NickName string `protobuf:"bytes,4,opt,name=NickName" json:"NickName"`
}

func (m *ResAccount) Reset()                    { *m = ResAccount{} }
func (m *ResAccount) String() string            { return proto.CompactTextString(m) }
func (*ResAccount) ProtoMessage()               {}
func (*ResAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResAccount) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ResAccount) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ResAccount) GetAcatar() int32 {
	if m != nil {
		return m.Acatar
	}
	return 0
}

func (m *ResAccount) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

type ResAccounts struct {
	Ra      []*ResAccount `protobuf:"bytes,1,rep,name=Ra" json:"Ra"`
	ErrCode int32         `protobuf:"varint,2,opt,name=ErrCode" json:"ErrCode"`
}

func (m *ResAccounts) Reset()                    { *m = ResAccounts{} }
func (m *ResAccounts) String() string            { return proto.CompactTextString(m) }
func (*ResAccounts) ProtoMessage()               {}
func (*ResAccounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResAccounts) GetRa() []*ResAccount {
	if m != nil {
		return m.Ra
	}
	return nil
}

func (m *ResAccounts) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

type ResShowNodeByGw struct {
	// 终端的id
	NodeID string `protobuf:"bytes,1,opt,name=NodeID" json:"NodeID"`
	// 终端的昵称
	NodeName  string `protobuf:"bytes,2,opt,name=NodeName" json:"NodeName"`
	NodeState int32  `protobuf:"varint,3,opt,name=NodeState" json:"NodeState"`
}

func (m *ResShowNodeByGw) Reset()                    { *m = ResShowNodeByGw{} }
func (m *ResShowNodeByGw) String() string            { return proto.CompactTextString(m) }
func (*ResShowNodeByGw) ProtoMessage()               {}
func (*ResShowNodeByGw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResShowNodeByGw) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *ResShowNodeByGw) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ResShowNodeByGw) GetNodeState() int32 {
	if m != nil {
		return m.NodeState
	}
	return 0
}

type ResShowNodesByGw struct {
	AllNodes []*ResShowNodeByGw `protobuf:"bytes,1,rep,name=AllNodes" json:"AllNodes"`
	ErrCode  int32              `protobuf:"varint,2,opt,name=ErrCode" json:"ErrCode"`
}

func (m *ResShowNodesByGw) Reset()                    { *m = ResShowNodesByGw{} }
func (m *ResShowNodesByGw) String() string            { return proto.CompactTextString(m) }
func (*ResShowNodesByGw) ProtoMessage()               {}
func (*ResShowNodesByGw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResShowNodesByGw) GetAllNodes() []*ResShowNodeByGw {
	if m != nil {
		return m.AllNodes
	}
	return nil
}

func (m *ResShowNodesByGw) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

type ResNodePos struct {
	// 网关id
	GatewayId string `protobuf:"bytes,8,opt,name=GatewayId" json:"GatewayId"`
	// 终端的id
	NodeID string `protobuf:"bytes,1,opt,name=NodeID" json:"NodeID"`
	// 终端的昵称
	NodeName string `protobuf:"bytes,2,opt,name=NodeName" json:"NodeName"`
	// 终端类型（0:不可移动终端1:可移动终端）
	NodeType string `protobuf:"bytes,6,opt,name=NodeType" json:"NodeType"`
	NTyepe   int32  `protobuf:"varint,7,opt,name=NTyepe" json:"NTyepe"`
	// 终端的经度
	Longitude float64 `protobuf:"fixed64,3,opt,name=Longitude" json:"Longitude"`
	// 终端的纬度
	Latitude float64 `protobuf:"fixed64,4,opt,name=Latitude" json:"Latitude"`
	// 终端的海拔高度
	Altitude  float64 `protobuf:"fixed64,5,opt,name=Altitude" json:"Altitude"`
	NodeState int32   `protobuf:"varint,9,opt,name=NodeState" json:"NodeState"`
}

func (m *ResNodePos) Reset()                    { *m = ResNodePos{} }
func (m *ResNodePos) String() string            { return proto.CompactTextString(m) }
func (*ResNodePos) ProtoMessage()               {}
func (*ResNodePos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResNodePos) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *ResNodePos) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *ResNodePos) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ResNodePos) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *ResNodePos) GetNTyepe() int32 {
	if m != nil {
		return m.NTyepe
	}
	return 0
}

func (m *ResNodePos) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *ResNodePos) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ResNodePos) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *ResNodePos) GetNodeState() int32 {
	if m != nil {
		return m.NodeState
	}
	return 0
}

type ResShowAllNodesPos struct {
	AllNodesPos []*ResNodePos `protobuf:"bytes,1,rep,name=AllNodesPos" json:"AllNodesPos"`
	ErrCode     int32         `protobuf:"varint,2,opt,name=ErrCode" json:"ErrCode"`
}

func (m *ResShowAllNodesPos) Reset()                    { *m = ResShowAllNodesPos{} }
func (m *ResShowAllNodesPos) String() string            { return proto.CompactTextString(m) }
func (*ResShowAllNodesPos) ProtoMessage()               {}
func (*ResShowAllNodesPos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResShowAllNodesPos) GetAllNodesPos() []*ResNodePos {
	if m != nil {
		return m.AllNodesPos
	}
	return nil
}

func (m *ResShowAllNodesPos) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

type ResUserGateway struct {
	GatewayID   string  `protobuf:"bytes,1,opt,name=GatewayID" json:"GatewayID"`
	GatewayName string  `protobuf:"bytes,2,opt,name=GatewayName" json:"GatewayName"`
	GatewayType string  `protobuf:"bytes,3,opt,name=GatewayType" json:"GatewayType"`
	GWType      int32   `protobuf:"varint,7,opt,name=GWType" json:"GWType"`
	Longitude   float64 `protobuf:"fixed64,4,opt,name=Longitude" json:"Longitude"`
	Latitude    float64 `protobuf:"fixed64,5,opt,name=Latitude" json:"Latitude"`
	Altitude    float64 `protobuf:"fixed64,6,opt,name=Altitude" json:"Altitude"`
}

func (m *ResUserGateway) Reset()                    { *m = ResUserGateway{} }
func (m *ResUserGateway) String() string            { return proto.CompactTextString(m) }
func (*ResUserGateway) ProtoMessage()               {}
func (*ResUserGateway) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ResUserGateway) GetGatewayID() string {
	if m != nil {
		return m.GatewayID
	}
	return ""
}

func (m *ResUserGateway) GetGatewayName() string {
	if m != nil {
		return m.GatewayName
	}
	return ""
}

func (m *ResUserGateway) GetGatewayType() string {
	if m != nil {
		return m.GatewayType
	}
	return ""
}

func (m *ResUserGateway) GetGWType() int32 {
	if m != nil {
		return m.GWType
	}
	return 0
}

func (m *ResUserGateway) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *ResUserGateway) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ResUserGateway) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

type ResUserGateways struct {
	RUGS    []*ResUserGateway `protobuf:"bytes,1,rep,name=RUGS" json:"RUGS"`
	ErrCode int32             `protobuf:"varint,2,opt,name=ErrCode" json:"ErrCode"`
}

func (m *ResUserGateways) Reset()                    { *m = ResUserGateways{} }
func (m *ResUserGateways) String() string            { return proto.CompactTextString(m) }
func (*ResUserGateways) ProtoMessage()               {}
func (*ResUserGateways) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ResUserGateways) GetRUGS() []*ResUserGateway {
	if m != nil {
		return m.RUGS
	}
	return nil
}

func (m *ResUserGateways) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

type ResShowGateway struct {
	GatewayID   string `protobuf:"bytes,1,opt,name=GatewayID" json:"GatewayID"`
	GatewayName string `protobuf:"bytes,2,opt,name=GatewayName" json:"GatewayName"`
	DeviceNum   int32  `protobuf:"varint,3,opt,name=DeviceNum" json:"DeviceNum"`
	BleState    int32  `protobuf:"varint,4,opt,name=BleState" json:"BleState"`
	NetServer   int32  `protobuf:"varint,5,opt,name=NetServer" json:"NetServer"`
	AppId       string `protobuf:"bytes,6,opt,name=AppId" json:"AppId"`
	BleName     string `protobuf:"bytes,7,opt,name=BleName" json:"BleName"`
	Permissions int32  `protobuf:"varint,8,opt,name=Permissions" json:"Permissions"`
}

func (m *ResShowGateway) Reset()                    { *m = ResShowGateway{} }
func (m *ResShowGateway) String() string            { return proto.CompactTextString(m) }
func (*ResShowGateway) ProtoMessage()               {}
func (*ResShowGateway) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ResShowGateway) GetGatewayID() string {
	if m != nil {
		return m.GatewayID
	}
	return ""
}

func (m *ResShowGateway) GetGatewayName() string {
	if m != nil {
		return m.GatewayName
	}
	return ""
}

func (m *ResShowGateway) GetDeviceNum() int32 {
	if m != nil {
		return m.DeviceNum
	}
	return 0
}

func (m *ResShowGateway) GetBleState() int32 {
	if m != nil {
		return m.BleState
	}
	return 0
}

func (m *ResShowGateway) GetNetServer() int32 {
	if m != nil {
		return m.NetServer
	}
	return 0
}

func (m *ResShowGateway) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ResShowGateway) GetBleName() string {
	if m != nil {
		return m.BleName
	}
	return ""
}

func (m *ResShowGateway) GetPermissions() int32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

type ResShowAllGws struct {
	ShowAllGws []*ResShowGateway `protobuf:"bytes,1,rep,name=ShowAllGws" json:"ShowAllGws"`
	ErrCode    int32             `protobuf:"varint,2,opt,name=ErrCode" json:"ErrCode"`
}

func (m *ResShowAllGws) Reset()                    { *m = ResShowAllGws{} }
func (m *ResShowAllGws) String() string            { return proto.CompactTextString(m) }
func (*ResShowAllGws) ProtoMessage()               {}
func (*ResShowAllGws) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ResShowAllGws) GetShowAllGws() []*ResShowGateway {
	if m != nil {
		return m.ShowAllGws
	}
	return nil
}

func (m *ResShowAllGws) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func init() {
	proto.RegisterType((*ReqGwUser)(nil), "pb.ReqGwUser")
	proto.RegisterType((*ResGwUser)(nil), "pb.ResGwUser")
	proto.RegisterType((*ResAccount)(nil), "pb.ResAccount")
	proto.RegisterType((*ResAccounts)(nil), "pb.ResAccounts")
	proto.RegisterType((*ResShowNodeByGw)(nil), "pb.ResShowNodeByGw")
	proto.RegisterType((*ResShowNodesByGw)(nil), "pb.ResShowNodesByGw")
	proto.RegisterType((*ResNodePos)(nil), "pb.ResNodePos")
	proto.RegisterType((*ResShowAllNodesPos)(nil), "pb.ResShowAllNodesPos")
	proto.RegisterType((*ResUserGateway)(nil), "pb.ResUserGateway")
	proto.RegisterType((*ResUserGateways)(nil), "pb.ResUserGateways")
	proto.RegisterType((*ResShowGateway)(nil), "pb.ResShowGateway")
	proto.RegisterType((*ResShowAllGws)(nil), "pb.ResShowAllGws")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GwUser service

type GwUserClient interface {
	AddGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error)
	AddGatewayAccount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error)
	DeletGatewayAccount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error)
	GetGatewayAccoount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResAccounts, error)
	DeletGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error)
	ValidationGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error)
	ValidationGatewayAccount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error)
	GetAllGateways(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResUserGateways, error)
	ShowNodesByGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResShowNodesByGw, error)
	ShowAllNodesPos(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResShowAllNodesPos, error)
	ShowAllGws(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResShowAllGws, error)
}

type gwUserClient struct {
	cc *grpc.ClientConn
}

func NewGwUserClient(cc *grpc.ClientConn) GwUserClient {
	return &gwUserClient{cc}
}

func (c *gwUserClient) AddGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error) {
	out := new(ResGwUser)
	err := grpc.Invoke(ctx, "/pb.GwUser/AddGateway", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) AddGatewayAccount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error) {
	out := new(ResGwUser)
	err := grpc.Invoke(ctx, "/pb.GwUser/AddGatewayAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) DeletGatewayAccount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error) {
	out := new(ResGwUser)
	err := grpc.Invoke(ctx, "/pb.GwUser/DeletGatewayAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) GetGatewayAccoount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResAccounts, error) {
	out := new(ResAccounts)
	err := grpc.Invoke(ctx, "/pb.GwUser/GetGatewayAccoount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) DeletGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error) {
	out := new(ResGwUser)
	err := grpc.Invoke(ctx, "/pb.GwUser/DeletGateway", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) ValidationGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error) {
	out := new(ResGwUser)
	err := grpc.Invoke(ctx, "/pb.GwUser/ValidationGateway", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) ValidationGatewayAccount(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResGwUser, error) {
	out := new(ResGwUser)
	err := grpc.Invoke(ctx, "/pb.GwUser/ValidationGatewayAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) GetAllGateways(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResUserGateways, error) {
	out := new(ResUserGateways)
	err := grpc.Invoke(ctx, "/pb.GwUser/GetAllGateways", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) ShowNodesByGateway(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResShowNodesByGw, error) {
	out := new(ResShowNodesByGw)
	err := grpc.Invoke(ctx, "/pb.GwUser/ShowNodesByGateway", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) ShowAllNodesPos(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResShowAllNodesPos, error) {
	out := new(ResShowAllNodesPos)
	err := grpc.Invoke(ctx, "/pb.GwUser/ShowAllNodesPos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gwUserClient) ShowAllGws(ctx context.Context, in *ReqGwUser, opts ...grpc.CallOption) (*ResShowAllGws, error) {
	out := new(ResShowAllGws)
	err := grpc.Invoke(ctx, "/pb.GwUser/ShowAllGws", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GwUser service

type GwUserServer interface {
	AddGateway(context.Context, *ReqGwUser) (*ResGwUser, error)
	AddGatewayAccount(context.Context, *ReqGwUser) (*ResGwUser, error)
	DeletGatewayAccount(context.Context, *ReqGwUser) (*ResGwUser, error)
	GetGatewayAccoount(context.Context, *ReqGwUser) (*ResAccounts, error)
	DeletGateway(context.Context, *ReqGwUser) (*ResGwUser, error)
	ValidationGateway(context.Context, *ReqGwUser) (*ResGwUser, error)
	ValidationGatewayAccount(context.Context, *ReqGwUser) (*ResGwUser, error)
	GetAllGateways(context.Context, *ReqGwUser) (*ResUserGateways, error)
	ShowNodesByGateway(context.Context, *ReqGwUser) (*ResShowNodesByGw, error)
	ShowAllNodesPos(context.Context, *ReqGwUser) (*ResShowAllNodesPos, error)
	ShowAllGws(context.Context, *ReqGwUser) (*ResShowAllGws, error)
}

func RegisterGwUserServer(s *grpc.Server, srv GwUserServer) {
	s.RegisterService(&_GwUser_serviceDesc, srv)
}

func _GwUser_AddGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).AddGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/AddGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).AddGateway(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_AddGatewayAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).AddGatewayAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/AddGatewayAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).AddGatewayAccount(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_DeletGatewayAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).DeletGatewayAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/DeletGatewayAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).DeletGatewayAccount(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_GetGatewayAccoount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).GetGatewayAccoount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/GetGatewayAccoount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).GetGatewayAccoount(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_DeletGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).DeletGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/DeletGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).DeletGateway(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_ValidationGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).ValidationGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/ValidationGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).ValidationGateway(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_ValidationGatewayAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).ValidationGatewayAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/ValidationGatewayAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).ValidationGatewayAccount(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_GetAllGateways_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).GetAllGateways(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/GetAllGateways",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).GetAllGateways(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_ShowNodesByGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).ShowNodesByGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/ShowNodesByGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).ShowNodesByGateway(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_ShowAllNodesPos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).ShowAllNodesPos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/ShowAllNodesPos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).ShowAllNodesPos(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GwUser_ShowAllGws_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGwUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GwUserServer).ShowAllGws(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GwUser/ShowAllGws",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GwUserServer).ShowAllGws(ctx, req.(*ReqGwUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _GwUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GwUser",
	HandlerType: (*GwUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGateway",
			Handler:    _GwUser_AddGateway_Handler,
		},
		{
			MethodName: "AddGatewayAccount",
			Handler:    _GwUser_AddGatewayAccount_Handler,
		},
		{
			MethodName: "DeletGatewayAccount",
			Handler:    _GwUser_DeletGatewayAccount_Handler,
		},
		{
			MethodName: "GetGatewayAccoount",
			Handler:    _GwUser_GetGatewayAccoount_Handler,
		},
		{
			MethodName: "DeletGateway",
			Handler:    _GwUser_DeletGateway_Handler,
		},
		{
			MethodName: "ValidationGateway",
			Handler:    _GwUser_ValidationGateway_Handler,
		},
		{
			MethodName: "ValidationGatewayAccount",
			Handler:    _GwUser_ValidationGatewayAccount_Handler,
		},
		{
			MethodName: "GetAllGateways",
			Handler:    _GwUser_GetAllGateways_Handler,
		},
		{
			MethodName: "ShowNodesByGateway",
			Handler:    _GwUser_ShowNodesByGateway_Handler,
		},
		{
			MethodName: "ShowAllNodesPos",
			Handler:    _GwUser_ShowAllNodesPos_Handler,
		},
		{
			MethodName: "ShowAllGws",
			Handler:    _GwUser_ShowAllGws_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway_user.proto",
}

func init() { proto.RegisterFile("gateway_user.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x8e, 0x93, 0x38, 0x69, 0x26, 0xa7, 0x7f, 0xdb, 0xaa, 0xb2, 0xaa, 0xea, 0x28, 0xb2, 0x8e,
	0x8e, 0x72, 0x81, 0x42, 0xd5, 0x0a, 0x2e, 0xe0, 0xca, 0x6d, 0x2a, 0x2b, 0x52, 0x89, 0xaa, 0x4d,
	0x03, 0x57, 0x15, 0xb8, 0xf1, 0xaa, 0x58, 0xb8, 0xb1, 0xf1, 0x3a, 0x8d, 0xf2, 0x02, 0xbc, 0x08,
	0xef, 0x05, 0x3c, 0x07, 0x57, 0x68, 0xc7, 0x6b, 0x7b, 0xed, 0xb6, 0x21, 0x20, 0xb8, 0x8a, 0xbf,
	0x99, 0xf9, 0x76, 0x67, 0xbe, 0x9d, 0x9d, 0x0d, 0x90, 0x1b, 0x27, 0x66, 0x73, 0x67, 0xf1, 0x76,
	0xc6, 0x59, 0xd4, 0x0b, 0xa3, 0x20, 0x0e, 0x48, 0x35, 0xbc, 0x36, 0xbf, 0x6a, 0xd0, 0xa2, 0xec,
	0xa3, 0x3d, 0x1f, 0x73, 0x16, 0x91, 0x3d, 0x68, 0x88, 0xdf, 0x41, 0xdf, 0xd0, 0x3a, 0x5a, 0x57,
	0xa7, 0x12, 0x91, 0x2d, 0xa8, 0xbd, 0xb2, 0x4e, 0x8d, 0x6a, 0x47, 0xeb, 0xb6, 0xa8, 0xf8, 0x24,
	0x07, 0xd0, 0xb2, 0x93, 0x15, 0x07, 0x7d, 0xa3, 0x86, 0xf6, 0xdc, 0x20, 0xd6, 0xb1, 0xc2, 0xf0,
	0x6c, 0x3c, 0x30, 0xea, 0xe8, 0x92, 0x48, 0xb0, 0xac, 0xc9, 0x24, 0x98, 0x4d, 0xe3, 0x41, 0xdf,
	0xd0, 0x71, 0x8b, 0xdc, 0x40, 0x76, 0x41, 0x1f, 0x0f, 0x9d, 0x5b, 0x66, 0x34, 0x90, 0x94, 0x00,
	0xf2, 0x1f, 0xac, 0x5b, 0x61, 0xe8, 0x7b, 0x13, 0x27, 0xf6, 0x82, 0xe9, 0xc0, 0x35, 0x9a, 0x1d,
	0xad, 0x5b, 0xa3, 0x45, 0x23, 0x31, 0xa0, 0x79, 0xe2, 0x33, 0x64, 0xaf, 0x21, 0x3b, 0x85, 0xe6,
	0x67, 0xac, 0x90, 0xaf, 0x5e, 0x61, 0xed, 0xef, 0x55, 0x68, 0x40, 0xf3, 0x2c, 0x8a, 0x4e, 0x03,
	0x37, 0xa9, 0x51, 0xa7, 0x29, 0x34, 0x63, 0x00, 0xca, 0xb8, 0x8c, 0x7c, 0x34, 0xcb, 0x7d, 0x58,
	0x13, 0x5f, 0x58, 0x66, 0x72, 0x18, 0x19, 0xc6, 0x8c, 0x26, 0x4e, 0xec, 0x44, 0x98, 0xac, 0x4e,
	0x25, 0x12, 0x9c, 0xa1, 0x37, 0xf9, 0x80, 0x9c, 0x24, 0xd7, 0x0c, 0x9b, 0x36, 0xb4, 0xf3, 0x5d,
	0x39, 0xf9, 0x17, 0xaa, 0xd4, 0x31, 0xb4, 0x4e, 0xad, 0xdb, 0x3e, 0xda, 0xe8, 0x85, 0xd7, 0xbd,
	0xdc, 0x49, 0xab, 0xd4, 0x51, 0xd3, 0xaf, 0x16, 0xd3, 0x9f, 0xc0, 0x26, 0x65, 0x7c, 0xf4, 0x3e,
	0x98, 0x0f, 0x03, 0x97, 0x9d, 0x2c, 0xec, 0xb9, 0xc8, 0x47, 0x7c, 0xcb, 0x1a, 0x5a, 0x54, 0x22,
	0xcc, 0x27, 0x70, 0x99, 0x5a, 0x43, 0x8a, 0x85, 0x7a, 0xe2, 0x7b, 0x14, 0x3b, 0x31, 0x93, 0x65,
	0xe4, 0x06, 0xf3, 0x0a, 0xb6, 0x94, 0x4d, 0x38, 0xee, 0xf2, 0x14, 0xd6, 0x2c, 0xdf, 0x47, 0x2c,
	0x13, 0xdf, 0x91, 0x89, 0xab, 0xc9, 0xd0, 0x2c, 0x68, 0x49, 0x0d, 0x9f, 0xaa, 0x78, 0x06, 0x22,
	0xec, 0x22, 0xe0, 0xea, 0xf9, 0xbb, 0xb2, 0xa7, 0x72, 0xc3, 0x6f, 0x55, 0x27, 0x7d, 0x97, 0x8b,
	0x30, 0x6d, 0xf1, 0x0c, 0xe3, 0x7a, 0x97, 0x0b, 0x16, 0x32, 0x6c, 0x6f, 0x9d, 0x4a, 0x24, 0xb2,
	0x38, 0x0f, 0xa6, 0x37, 0x5e, 0x3c, 0x73, 0x13, 0x45, 0x34, 0x9a, 0x1b, 0xc4, 0x8a, 0xe7, 0x4e,
	0x9c, 0x38, 0xeb, 0xe8, 0xcc, 0xb0, 0xf0, 0x59, 0xbe, 0xf4, 0xe9, 0x89, 0x2f, 0xc5, 0x45, 0x9d,
	0x5b, 0x65, 0x9d, 0xdf, 0x01, 0x91, 0xfa, 0xa5, 0xaa, 0x09, 0x3d, 0x0e, 0xa1, 0xad, 0xc0, 0x52,
	0x97, 0x48, 0xd1, 0xa8, 0x1a, 0xb2, 0x44, 0xea, 0x6f, 0x1a, 0x6c, 0x50, 0xc6, 0x45, 0xef, 0x4a,
	0x49, 0x8b, 0xd7, 0x4d, 0x2b, 0x5f, 0xb7, 0x0e, 0xb4, 0x25, 0x50, 0x94, 0x55, 0x4d, 0x4a, 0x04,
	0xea, 0x5b, 0x2b, 0x44, 0xa4, 0x12, 0xdb, 0x6f, 0xd0, 0x29, 0x25, 0x4e, 0x50, 0x51, 0xe2, 0xfa,
	0x32, 0x89, 0xf5, 0x25, 0x12, 0x37, 0x8a, 0x12, 0x9b, 0x23, 0xbc, 0x11, 0x4a, 0x85, 0x9c, 0xfc,
	0x0f, 0x75, 0x3a, 0xb6, 0x47, 0x52, 0x3a, 0x22, 0xa5, 0x53, 0x42, 0x28, 0xfa, 0x97, 0xe8, 0xf6,
	0x3d, 0xd1, 0x4d, 0x1c, 0xcd, 0x9f, 0xd2, 0xed, 0x00, 0x5a, 0x7d, 0x76, 0xe7, 0x4d, 0xd8, 0x70,
	0x76, 0x9b, 0x5e, 0xb9, 0xcc, 0x20, 0x2a, 0x3c, 0xf1, 0x65, 0x9f, 0xd4, 0xd1, 0x99, 0x61, 0x6c,
	0x22, 0x16, 0x8f, 0x58, 0x74, 0xc7, 0xa2, 0x74, 0xd4, 0x65, 0x06, 0x31, 0xcc, 0xad, 0x30, 0x1c,
	0xb8, 0xe9, 0x30, 0x47, 0xa0, 0x8e, 0xe9, 0x66, 0x61, 0x4c, 0x8b, 0x4c, 0x2f, 0x58, 0x74, 0xeb,
	0x71, 0xee, 0x05, 0x53, 0x8e, 0x17, 0x4e, 0xa7, 0xaa, 0xc9, 0xbc, 0x82, 0xf5, 0xbc, 0x2d, 0xed,
	0x39, 0x27, 0x47, 0x00, 0x39, 0x2a, 0xa9, 0xaa, 0x48, 0x44, 0x95, 0xa8, 0xc7, 0xb5, 0x3d, 0xfa,
	0x52, 0x87, 0x86, 0x7c, 0x24, 0x9e, 0x00, 0x58, 0xae, 0x9b, 0x2a, 0xbc, 0x9e, 0x2c, 0x29, 0xdf,
	0xc8, 0x7d, 0x09, 0xe5, 0x83, 0x62, 0x56, 0xc8, 0x31, 0x6c, 0xe7, 0xd1, 0xe9, 0x04, 0xff, 0x19,
	0xe9, 0x19, 0xec, 0xf4, 0x99, 0xcf, 0xe2, 0x5f, 0xa4, 0x3d, 0x07, 0x62, 0x17, 0x48, 0x0f, 0xb1,
	0x36, 0x8b, 0xa3, 0x9b, 0x9b, 0x15, 0xd2, 0x83, 0x7f, 0xd4, 0xed, 0x56, 0xa9, 0xe9, 0xb5, 0xe3,
	0x7b, 0x2e, 0x3e, 0xaf, 0xab, 0x92, 0x5e, 0x80, 0x71, 0x8f, 0xb4, 0x7a, 0x61, 0x1b, 0x36, 0x8b,
	0xc5, 0x21, 0xa5, 0xb7, 0xa5, 0xc4, 0xd8, 0xb9, 0x7f, 0x5d, 0x44, 0x61, 0x2f, 0x81, 0xa8, 0x0f,
	0xc2, 0xc3, 0x99, 0xee, 0x96, 0x9e, 0x04, 0x7c, 0x3a, 0x30, 0xe1, 0xcd, 0xf2, 0x94, 0x2b, 0x31,
	0xf7, 0x14, 0xa6, 0x12, 0x66, 0x56, 0xc8, 0xa1, 0xda, 0x7c, 0x65, 0xda, 0x76, 0x91, 0x66, 0xcf,
	0xb9, 0x59, 0xb9, 0x6e, 0xe0, 0xbf, 0xae, 0xe3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xc9,
	0x26, 0x9d, 0x8b, 0x09, 0x00, 0x00,
}
