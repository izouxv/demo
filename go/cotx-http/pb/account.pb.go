// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	account.proto
	device_http.proto
	deviceset.proto
	equipment_warm.proto
	feedback.proto
	gateway.proto
	gateway_addtional.proto
	gateway_set.proto
	gateway_usb.proto
	gateway_user.proto
	sso.proto
	updatedevice.proto

It has these top-level messages:
	AccountRequest
	AccountReply
	MultiAccountRequest
	MapAccountReply
	ResNodeState
	NodeType
	ResNodesType
	ResNodeTypeNum
	NodeDate
	ResNodeDateUp
	ReqNodeMessageByGw
	ReqNodeMessageByNid
	WorkModel
	ResWorkModel
	ResGetDeviceScan
	DeviceEuis
	ReqregistDevice
	DeviceEui
	ResRegistDevice
	ReqRegistNode
	ResRegistNode
	ReqSetDeviceName
	NodeMessage
	ResSetDeviceName
	ReqDeleteDevice
	ResDeleteDevice
	ReqSetDeviceSendModel
	ResSetDeviceSendModel
	ReqSetDeviceReciveModel
	ResSetDeviceReciveModel
	ReqWarm
	Warm
	ResGwWarm
	ResGWNetWarm
	ResNodeWarm
	FeedbackRequest
	FeedbackReply
	MultiFeedbackRequest
	MapFeedbackReply
	ResGwNetState
	ResGwState
	ResGwFile
	ResGwVideo
	ResGwVideos
	ResGwPhoto
	ResGwPhotos
	ResGwUSBStat
	ResGwWifiScan
	ResGwWifiScans
	ResGwUsbWifiScans
	ResGwWifiAddress
	ResGwWifiDNS
	ResGwUsbWifiAddress
	DNSs
	ResGwUsbWifiDNS
	ResGwCableAddress
	ResGwCableDNS
	ResGwMessage
	ResGwLora
	ReqGateway
	ServerIot
	GatewaySwitch
	ResGatewaySwitch
	ResServerIots
	HotSpot
	ResHotSpot
	ResUsbHotSpot
	UsbWifiStat
	ResUsbWifiStat
	GCardStat
	ResUsbGCardStat
	WifiConnecting
	ResWifiConnected
	ReqGetGatewayFile
	GetGatewayFile
	ResGetGatewayFile
	ResPowerModel
	ResMusicSet
	ResPhotoSet
	ResVideoSet
	ResAppEui
	ReqGwAddtional
	BleScan
	ResBleScans
	ResCode
	ReqSwitch
	ReqInstruction
	ReqGwPower
	ReqLora
	ReqVideo
	ReqPhoto
	ReqMusic
	ReqDeletFile
	ReqUpLog
	ReqSetIP
	ReqSetUsbIP
	DNS
	ReqSetDNS
	ReqSetUsbDNS
	ReqSetHotSpot
	ReqSetUsbHotSpot
	ReqSetName
	ReqConnectWifi
	ReqConnectUsbWifi
	ReqUsbSwitch
	ReqIOTServer
	PhotoCode
	ResPhotoCode
	VideoCode
	ResVideoCode
	UsbWifiHotSpotUser
	ResUsbWifiHotSpotUser
	UsbWifiWlan
	ResUsbWifiWlan
	ReqGatewayUsb
	ResWifiHotSpotUser
	UsbWifiScan
	ResUsbWifiScan
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
	SsoRequest
	SsoReply
	MultiSsoRequest
	MapSsoReply
	ReqGetNewGatewayVersion
	GetNewGatewayVersion
	ResGetNewGatewayVersion
	ReqUpdateGatewayVersion
	ResUpdateGatewayVersion
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

// The request message containing the account's id etcd.
type AccountRequest struct {
	Uid             int32  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Username        string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Salt            string `protobuf:"bytes,4,opt,name=salt" json:"salt,omitempty"`
	State           int32  `protobuf:"varint,5,opt,name=state" json:"state,omitempty"`
	Email           string `protobuf:"bytes,6,opt,name=email" json:"email,omitempty"`
	Phone           string `protobuf:"bytes,7,opt,name=phone" json:"phone,omitempty"`
	Nickname        string `protobuf:"bytes,8,opt,name=nickname" json:"nickname,omitempty"`
	Realname        string `protobuf:"bytes,9,opt,name=realname" json:"realname,omitempty"`
	IsCertification int32  `protobuf:"varint,10,opt,name=isCertification" json:"isCertification,omitempty"`
	IdentityCard    string `protobuf:"bytes,11,opt,name=identityCard" json:"identityCard,omitempty"`
	Gender          int32  `protobuf:"varint,12,opt,name=gender" json:"gender,omitempty"`
	Birthday        int64  `protobuf:"varint,13,opt,name=birthday" json:"birthday,omitempty"`
	Avatar          int32  `protobuf:"varint,14,opt,name=avatar" json:"avatar,omitempty"`
	Province        string `protobuf:"bytes,15,opt,name=province" json:"province,omitempty"`
	City            string `protobuf:"bytes,16,opt,name=city" json:"city,omitempty"`
	Signature       string `protobuf:"bytes,17,opt,name=signature" json:"signature,omitempty"`
	RegIp           int64  `protobuf:"varint,18,opt,name=regIp" json:"regIp,omitempty"`
	UserAddress     string `protobuf:"bytes,19,opt,name=userAddress" json:"userAddress,omitempty"`
	UserJobId       int32  `protobuf:"varint,20,opt,name=userJobId" json:"userJobId,omitempty"`
	CreditValues    int32  `protobuf:"varint,21,opt,name=creditValues" json:"creditValues,omitempty"`
	UserPoint       int32  `protobuf:"varint,22,opt,name=userPoint" json:"userPoint,omitempty"`
	UserGradeId     int32  `protobuf:"varint,23,opt,name=userGradeId" json:"userGradeId,omitempty"`
	LastLoginIp     int64  `protobuf:"varint,24,opt,name=lastLoginIp" json:"lastLoginIp,omitempty"`
	IsFirstLogin    int32  `protobuf:"varint,25,opt,name=isFirstLogin" json:"isFirstLogin,omitempty"`
	Source          string `protobuf:"bytes,26,opt,name=source" json:"source,omitempty"`
	Radius          int32  `protobuf:"varint,27,opt,name=radius" json:"radius,omitempty"`
	Language        string `protobuf:"bytes,28,opt,name=language" json:"language,omitempty"`
}

func (m *AccountRequest) Reset()                    { *m = AccountRequest{} }
func (m *AccountRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()               {}
func (*AccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *AccountRequest) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *AccountRequest) GetIsCertification() int32 {
	if m != nil {
		return m.IsCertification
	}
	return 0
}

func (m *AccountRequest) GetIdentityCard() string {
	if m != nil {
		return m.IdentityCard
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

func (m *AccountRequest) GetAvatar() int32 {
	if m != nil {
		return m.Avatar
	}
	return 0
}

func (m *AccountRequest) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *AccountRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *AccountRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AccountRequest) GetRegIp() int64 {
	if m != nil {
		return m.RegIp
	}
	return 0
}

func (m *AccountRequest) GetUserAddress() string {
	if m != nil {
		return m.UserAddress
	}
	return ""
}

func (m *AccountRequest) GetUserJobId() int32 {
	if m != nil {
		return m.UserJobId
	}
	return 0
}

func (m *AccountRequest) GetCreditValues() int32 {
	if m != nil {
		return m.CreditValues
	}
	return 0
}

func (m *AccountRequest) GetUserPoint() int32 {
	if m != nil {
		return m.UserPoint
	}
	return 0
}

func (m *AccountRequest) GetUserGradeId() int32 {
	if m != nil {
		return m.UserGradeId
	}
	return 0
}

func (m *AccountRequest) GetLastLoginIp() int64 {
	if m != nil {
		return m.LastLoginIp
	}
	return 0
}

func (m *AccountRequest) GetIsFirstLogin() int32 {
	if m != nil {
		return m.IsFirstLogin
	}
	return 0
}

func (m *AccountRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AccountRequest) GetRadius() int32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *AccountRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

// The response message containing the id
type AccountReply struct {
	Uid             int32  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Username        string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email           string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone           string `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	State           int32  `protobuf:"varint,5,opt,name=state" json:"state,omitempty"`
	LastLoginTime   int64  `protobuf:"varint,6,opt,name=lastLoginTime" json:"lastLoginTime,omitempty"`
	CreateTime      int64  `protobuf:"varint,7,opt,name=createTime" json:"createTime,omitempty"`
	Nickname        string `protobuf:"bytes,8,opt,name=nickname" json:"nickname,omitempty"`
	Realname        string `protobuf:"bytes,9,opt,name=realname" json:"realname,omitempty"`
	IsCertification int32  `protobuf:"varint,10,opt,name=isCertification" json:"isCertification,omitempty"`
	IdentityCard    string `protobuf:"bytes,11,opt,name=identityCard" json:"identityCard,omitempty"`
	ErrorCode       int32  `protobuf:"varint,12,opt,name=errorCode" json:"errorCode,omitempty"`
	Gender          int32  `protobuf:"varint,13,opt,name=gender" json:"gender,omitempty"`
	Birthday        int64  `protobuf:"varint,14,opt,name=birthday" json:"birthday,omitempty"`
	Avatar          int32  `protobuf:"varint,15,opt,name=avatar" json:"avatar,omitempty"`
	Province        string `protobuf:"bytes,16,opt,name=province" json:"province,omitempty"`
	City            string `protobuf:"bytes,17,opt,name=city" json:"city,omitempty"`
	Signature       string `protobuf:"bytes,18,opt,name=signature" json:"signature,omitempty"`
	UserAddress     string `protobuf:"bytes,19,opt,name=userAddress" json:"userAddress,omitempty"`
	UserJobId       int32  `protobuf:"varint,20,opt,name=userJobId" json:"userJobId,omitempty"`
	CreditValues    int32  `protobuf:"varint,21,opt,name=creditValues" json:"creditValues,omitempty"`
	UserPoint       int32  `protobuf:"varint,22,opt,name=userPoint" json:"userPoint,omitempty"`
	UserGradeId     int32  `protobuf:"varint,23,opt,name=userGradeId" json:"userGradeId,omitempty"`
	RegTime         int64  `protobuf:"varint,24,opt,name=regTime" json:"regTime,omitempty"`
	RegIp           int64  `protobuf:"varint,25,opt,name=regIp" json:"regIp,omitempty"`
	LastLoginIp     int64  `protobuf:"varint,26,opt,name=lastLoginIp" json:"lastLoginIp,omitempty"`
	LastActive      int64  `protobuf:"varint,27,opt,name=lastActive" json:"lastActive,omitempty"`
	UserModify      int64  `protobuf:"varint,28,opt,name=userModify" json:"userModify,omitempty"`
	IsFirstLogin    int32  `protobuf:"varint,29,opt,name=isFirstLogin" json:"isFirstLogin,omitempty"`
	Radius          int32  `protobuf:"varint,30,opt,name=radius" json:"radius,omitempty"`
	Language        string `protobuf:"bytes,31,opt,name=language" json:"language,omitempty"`
}

func (m *AccountReply) Reset()                    { *m = AccountReply{} }
func (m *AccountReply) String() string            { return proto.CompactTextString(m) }
func (*AccountReply) ProtoMessage()               {}
func (*AccountReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func (m *AccountReply) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *AccountReply) GetLastLoginTime() int64 {
	if m != nil {
		return m.LastLoginTime
	}
	return 0
}

func (m *AccountReply) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *AccountReply) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AccountReply) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *AccountReply) GetIsCertification() int32 {
	if m != nil {
		return m.IsCertification
	}
	return 0
}

func (m *AccountReply) GetIdentityCard() string {
	if m != nil {
		return m.IdentityCard
	}
	return ""
}

func (m *AccountReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
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

func (m *AccountReply) GetAvatar() int32 {
	if m != nil {
		return m.Avatar
	}
	return 0
}

func (m *AccountReply) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *AccountReply) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *AccountReply) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AccountReply) GetUserAddress() string {
	if m != nil {
		return m.UserAddress
	}
	return ""
}

func (m *AccountReply) GetUserJobId() int32 {
	if m != nil {
		return m.UserJobId
	}
	return 0
}

func (m *AccountReply) GetCreditValues() int32 {
	if m != nil {
		return m.CreditValues
	}
	return 0
}

func (m *AccountReply) GetUserPoint() int32 {
	if m != nil {
		return m.UserPoint
	}
	return 0
}

func (m *AccountReply) GetUserGradeId() int32 {
	if m != nil {
		return m.UserGradeId
	}
	return 0
}

func (m *AccountReply) GetRegTime() int64 {
	if m != nil {
		return m.RegTime
	}
	return 0
}

func (m *AccountReply) GetRegIp() int64 {
	if m != nil {
		return m.RegIp
	}
	return 0
}

func (m *AccountReply) GetLastLoginIp() int64 {
	if m != nil {
		return m.LastLoginIp
	}
	return 0
}

func (m *AccountReply) GetLastActive() int64 {
	if m != nil {
		return m.LastActive
	}
	return 0
}

func (m *AccountReply) GetUserModify() int64 {
	if m != nil {
		return m.UserModify
	}
	return 0
}

func (m *AccountReply) GetIsFirstLogin() int32 {
	if m != nil {
		return m.IsFirstLogin
	}
	return 0
}

func (m *AccountReply) GetRadius() int32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *AccountReply) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type MultiAccountRequest struct {
	Accounts map[int32]*AccountRequest `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Source   string                    `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
}

func (m *MultiAccountRequest) Reset()                    { *m = MultiAccountRequest{} }
func (m *MultiAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiAccountRequest) ProtoMessage()               {}
func (*MultiAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	Accounts  map[int32]*AccountReply `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ErrorCode int32                   `protobuf:"varint,2,opt,name=errorCode" json:"errorCode,omitempty"`
}

func (m *MapAccountReply) Reset()                    { *m = MapAccountReply{} }
func (m *MapAccountReply) String() string            { return proto.CompactTextString(m) }
func (*MapAccountReply) ProtoMessage()               {}
func (*MapAccountReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MapAccountReply) GetAccounts() map[int32]*AccountReply {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *MapAccountReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountRequest)(nil), "pb.AccountRequest")
	proto.RegisterType((*AccountReply)(nil), "pb.AccountReply")
	proto.RegisterType((*MultiAccountRequest)(nil), "pb.MultiAccountRequest")
	proto.RegisterType((*MapAccountReply)(nil), "pb.MapAccountReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Account service

type AccountClient interface {
	UpdateExInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	Show(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetUserInfoById(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetUserInfoAll(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	UpdateAccountInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	UpdateCertification(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetCertification(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	UpdateUserValues(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetUserValues(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetBatchAccountInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error)
	GetBatchExInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error)
	GetBatchAllUserInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) UpdateExInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/UpdateExInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Show(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/Show", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUserInfoById(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/GetUserInfoById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUserInfoAll(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/GetUserInfoAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateAccountInfo(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/UpdateAccountInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateCertification(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/UpdateCertification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetCertification(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/GetCertification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateUserValues(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/UpdateUserValues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUserValues(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/GetUserValues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetBatchAccountInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error) {
	out := new(MapAccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/GetBatchAccountInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetBatchExInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error) {
	out := new(MapAccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/GetBatchExInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetBatchAllUserInfo(ctx context.Context, in *MultiAccountRequest, opts ...grpc.CallOption) (*MapAccountReply, error) {
	out := new(MapAccountReply)
	err := grpc.Invoke(ctx, "/pb.Account/GetBatchAllUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountServer interface {
	UpdateExInfo(context.Context, *AccountRequest) (*AccountReply, error)
	Show(context.Context, *AccountRequest) (*AccountReply, error)
	GetUserInfoById(context.Context, *AccountRequest) (*AccountReply, error)
	GetUserInfoAll(context.Context, *AccountRequest) (*AccountReply, error)
	UpdateAccountInfo(context.Context, *AccountRequest) (*AccountReply, error)
	UpdateCertification(context.Context, *AccountRequest) (*AccountReply, error)
	GetCertification(context.Context, *AccountRequest) (*AccountReply, error)
	UpdateUserValues(context.Context, *AccountRequest) (*AccountReply, error)
	GetUserValues(context.Context, *AccountRequest) (*AccountReply, error)
	GetBatchAccountInfo(context.Context, *MultiAccountRequest) (*MapAccountReply, error)
	GetBatchExInfo(context.Context, *MultiAccountRequest) (*MapAccountReply, error)
	GetBatchAllUserInfo(context.Context, *MultiAccountRequest) (*MapAccountReply, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_UpdateExInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateExInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/UpdateExInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateExInfo(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _Account_GetUserInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUserInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetUserInfoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUserInfoById(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetUserInfoAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUserInfoAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetUserInfoAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUserInfoAll(ctx, req.(*AccountRequest))
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

func _Account_UpdateCertification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateCertification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/UpdateCertification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateCertification(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetCertification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetCertification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetCertification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetCertification(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UpdateUserValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateUserValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/UpdateUserValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateUserValues(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetUserValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUserValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetUserValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUserValues(ctx, req.(*AccountRequest))
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

func _Account_GetBatchExInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetBatchExInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetBatchExInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetBatchExInfo(ctx, req.(*MultiAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetBatchAllUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetBatchAllUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/GetBatchAllUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetBatchAllUserInfo(ctx, req.(*MultiAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateExInfo",
			Handler:    _Account_UpdateExInfo_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Account_Show_Handler,
		},
		{
			MethodName: "GetUserInfoById",
			Handler:    _Account_GetUserInfoById_Handler,
		},
		{
			MethodName: "GetUserInfoAll",
			Handler:    _Account_GetUserInfoAll_Handler,
		},
		{
			MethodName: "UpdateAccountInfo",
			Handler:    _Account_UpdateAccountInfo_Handler,
		},
		{
			MethodName: "UpdateCertification",
			Handler:    _Account_UpdateCertification_Handler,
		},
		{
			MethodName: "GetCertification",
			Handler:    _Account_GetCertification_Handler,
		},
		{
			MethodName: "UpdateUserValues",
			Handler:    _Account_UpdateUserValues_Handler,
		},
		{
			MethodName: "GetUserValues",
			Handler:    _Account_GetUserValues_Handler,
		},
		{
			MethodName: "GetBatchAccountInfo",
			Handler:    _Account_GetBatchAccountInfo_Handler,
		},
		{
			MethodName: "GetBatchExInfo",
			Handler:    _Account_GetBatchExInfo_Handler,
		},
		{
			MethodName: "GetBatchAllUserInfo",
			Handler:    _Account_GetBatchAllUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 889 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xde, 0x24, 0x4d, 0xd3, 0x9e, 0x36, 0x4d, 0x3a, 0x59, 0x76, 0x67, 0x43, 0x29, 0x21, 0x02,
	0x94, 0xab, 0x5c, 0x14, 0x09, 0x96, 0x45, 0x2b, 0x91, 0xad, 0x4a, 0x15, 0x44, 0x05, 0x32, 0x2c,
	0xf7, 0x13, 0xcf, 0x69, 0x3a, 0xaa, 0x6b, 0x9b, 0x99, 0x71, 0x17, 0xbf, 0x05, 0x0f, 0xc0, 0xab,
	0x70, 0xc1, 0xb3, 0xf0, 0x22, 0x68, 0x66, 0xec, 0xc4, 0x4e, 0x62, 0x91, 0x70, 0x05, 0x77, 0xfe,
	0xbe, 0xf3, 0x33, 0xe7, 0x9c, 0x99, 0xef, 0x24, 0xd0, 0x66, 0xbe, 0x1f, 0x25, 0xa1, 0x1e, 0xc7,
	0x32, 0xd2, 0x11, 0xa9, 0xc7, 0xb3, 0xe1, 0xef, 0xfb, 0x70, 0x32, 0x71, 0xac, 0x87, 0xbf, 0x24,
	0xa8, 0x34, 0xe9, 0x42, 0x23, 0x11, 0x9c, 0xd6, 0x06, 0xb5, 0x51, 0xd3, 0x33, 0x9f, 0xa4, 0x0f,
	0x07, 0x89, 0x42, 0x19, 0xb2, 0x07, 0xa4, 0xf5, 0x41, 0x6d, 0x74, 0xe8, 0x2d, 0xb0, 0xb1, 0xc5,
	0x4c, 0xa9, 0x77, 0x91, 0xe4, 0xb4, 0xe1, 0x6c, 0x39, 0x26, 0x04, 0xf6, 0x14, 0x0b, 0x34, 0xdd,
	0xb3, 0xbc, 0xfd, 0x26, 0x4f, 0xa1, 0xa9, 0x34, 0xd3, 0x48, 0x9b, 0x36, 0xbf, 0x03, 0x86, 0xc5,
	0x07, 0x26, 0x02, 0xba, 0x6f, 0x5d, 0x1d, 0x30, 0x6c, 0x7c, 0x17, 0x85, 0x48, 0x5b, 0x8e, 0xb5,
	0xc0, 0x9c, 0x18, 0x0a, 0xff, 0xde, 0x56, 0x73, 0xe0, 0x4e, 0xcc, 0xb1, 0xb1, 0x49, 0x64, 0x81,
	0xb5, 0x1d, 0x3a, 0x5b, 0x8e, 0xc9, 0x08, 0x3a, 0x42, 0x5d, 0xa2, 0xd4, 0xe2, 0x56, 0xf8, 0x4c,
	0x8b, 0x28, 0xa4, 0x60, 0x6b, 0x58, 0xa5, 0xc9, 0x10, 0x8e, 0x05, 0xc7, 0x50, 0x0b, 0x9d, 0x5e,
	0x32, 0xc9, 0xe9, 0x91, 0xcd, 0x54, 0xe2, 0xc8, 0x33, 0xd8, 0x9f, 0x63, 0xc8, 0x51, 0xd2, 0x63,
	0x9b, 0x24, 0x43, 0xa6, 0x82, 0x99, 0x90, 0xfa, 0x8e, 0xb3, 0x94, 0xb6, 0x07, 0xb5, 0x51, 0xc3,
	0x5b, 0x60, 0x13, 0xc3, 0x1e, 0x99, 0x66, 0x92, 0x9e, 0xb8, 0x18, 0x87, 0xec, 0x0c, 0x65, 0xf4,
	0x28, 0x42, 0x1f, 0x69, 0x27, 0x9b, 0x61, 0x86, 0xcd, 0x0c, 0x7d, 0xa1, 0x53, 0xda, 0x75, 0x33,
	0x34, 0xdf, 0xe4, 0x0c, 0x0e, 0x95, 0x98, 0x87, 0x4c, 0x27, 0x12, 0xe9, 0xa9, 0x35, 0x2c, 0x09,
	0x33, 0x35, 0x89, 0xf3, 0x69, 0x4c, 0x89, 0x3d, 0xde, 0x01, 0x32, 0x80, 0x23, 0x73, 0x67, 0x13,
	0xce, 0x25, 0x2a, 0x45, 0x7b, 0x36, 0xaa, 0x48, 0x99, 0xac, 0x06, 0x7e, 0x1b, 0xcd, 0xa6, 0x9c,
	0x3e, 0xb5, 0x05, 0x2e, 0x09, 0x33, 0x13, 0x5f, 0x22, 0x17, 0xfa, 0x67, 0x16, 0x24, 0xa8, 0xe8,
	0x7b, 0xd6, 0xa1, 0xc4, 0xe5, 0x19, 0x7e, 0x88, 0x44, 0xa8, 0xe9, 0xb3, 0x65, 0x06, 0x4b, 0xe4,
	0x15, 0x5c, 0x4b, 0xc6, 0x71, 0xca, 0xe9, 0x73, 0x6b, 0x2f, 0x52, 0xc6, 0x23, 0x60, 0x4a, 0x7f,
	0x17, 0xcd, 0x45, 0x38, 0x8d, 0x29, 0xb5, 0xf5, 0x17, 0x29, 0x7b, 0x33, 0xea, 0x1b, 0x21, 0x33,
	0x86, 0xbe, 0x70, 0x55, 0x14, 0x39, 0x33, 0x65, 0x15, 0x25, 0xd2, 0x47, 0xda, 0xb7, 0x4d, 0x66,
	0xc8, 0xf0, 0x92, 0x71, 0x91, 0x28, 0xfa, 0xbe, 0x9b, 0xbe, 0x43, 0x66, 0xfa, 0x01, 0x0b, 0xe7,
	0x09, 0x9b, 0x23, 0x3d, 0x73, 0xd3, 0xcf, 0xf1, 0xf0, 0xb7, 0x16, 0x1c, 0x2f, 0xe4, 0x11, 0x07,
	0xe9, 0x8e, 0xe2, 0x58, 0x3c, 0xeb, 0xc6, 0xc6, 0x67, 0xbd, 0x57, 0x7c, 0xd6, 0x9b, 0x85, 0xf1,
	0x31, 0xb4, 0x17, 0xfd, 0xff, 0x24, 0x1e, 0xd0, 0x0a, 0xa4, 0xe1, 0x95, 0x49, 0x72, 0x0e, 0xe0,
	0x4b, 0x64, 0x1a, 0xad, 0x4b, 0xcb, 0xba, 0x14, 0x98, 0xff, 0x88, 0x64, 0xce, 0xe0, 0x10, 0xa5,
	0x8c, 0xe4, 0x65, 0xc4, 0x31, 0x53, 0xcd, 0x92, 0x28, 0x08, 0xaa, 0x5d, 0x29, 0xa8, 0x93, 0x4a,
	0x41, 0x75, 0x2a, 0x05, 0xd5, 0xad, 0x10, 0xd4, 0x69, 0x95, 0xa0, 0xc8, 0xaa, 0xa0, 0xfe, 0x0f,
	0xd2, 0xa1, 0xd0, 0x92, 0x38, 0xb7, 0xd7, 0xef, 0x64, 0x93, 0xc3, 0xe5, 0x3a, 0x78, 0xb1, 0xb2,
	0x0e, 0x8a, 0x52, 0xeb, 0xaf, 0x4b, 0xed, 0x1c, 0xc0, 0xc0, 0x89, 0xaf, 0xc5, 0x23, 0x5a, 0xc9,
	0x34, 0xbc, 0x02, 0x63, 0xec, 0xa6, 0x80, 0x9b, 0x88, 0x8b, 0xdb, 0xd4, 0x0a, 0xa7, 0xe1, 0x15,
	0x98, 0x35, 0xa9, 0x7e, 0xb0, 0x59, 0xaa, 0x99, 0x24, 0xcf, 0x2b, 0x25, 0xf9, 0xe1, 0x8a, 0x24,
	0xff, 0xac, 0x41, 0xef, 0x26, 0x09, 0xb4, 0x58, 0xf9, 0xd9, 0x9a, 0xc0, 0x41, 0xf6, 0xf3, 0xa6,
	0x68, 0x6d, 0xd0, 0x18, 0x1d, 0x5d, 0x7c, 0x32, 0x8e, 0x67, 0xe3, 0x0d, 0xae, 0xe3, 0x0c, 0xaa,
	0xab, 0x50, 0xcb, 0xd4, 0x5b, 0x84, 0x15, 0x36, 0x47, 0xbd, 0xb8, 0x39, 0xfa, 0xdf, 0x43, 0xbb,
	0x14, 0x62, 0xb6, 0xc0, 0x3d, 0xa6, 0xf9, 0x16, 0xb8, 0xc7, 0x94, 0x8c, 0xa0, 0xf9, 0x68, 0x6e,
	0xd2, 0x46, 0x1e, 0x5d, 0x10, 0x73, 0x74, 0xf9, 0x54, 0xcf, 0x39, 0xbc, 0xaa, 0xbf, 0xac, 0x0d,
	0xff, 0xa8, 0x41, 0xe7, 0x86, 0xc5, 0xa5, 0xcd, 0xf2, 0x7a, 0xad, 0xfe, 0x8f, 0x6c, 0xfd, 0x65,
	0xb7, 0xca, 0xda, 0x4b, 0xe2, 0xaa, 0xaf, 0x88, 0xab, 0x7f, 0xf3, 0xcf, 0x1d, 0x7c, 0x5a, 0xee,
	0xa0, 0x5b, 0xea, 0x20, 0x0e, 0xd2, 0x42, 0xfd, 0x17, 0x7f, 0x35, 0xa1, 0x95, 0xd9, 0xc8, 0xe7,
	0x70, 0xfc, 0x36, 0xe6, 0x4c, 0xe3, 0xd5, 0xaf, 0xd3, 0xf0, 0x36, 0x22, 0x1b, 0x5a, 0xef, 0xaf,
	0x25, 0x1b, 0x3e, 0x21, 0x63, 0xd8, 0xfb, 0xf1, 0x2e, 0x7a, 0xb7, 0xb5, 0xff, 0x97, 0xd0, 0xb9,
	0x46, 0xfd, 0x56, 0xa1, 0x34, 0xc7, 0xbc, 0x49, 0xa7, 0x7c, 0xeb, 0xd0, 0x97, 0x70, 0x52, 0x08,
	0x9d, 0x04, 0xc1, 0xd6, 0x91, 0x5f, 0xc1, 0xa9, 0x6b, 0x2e, 0xe3, 0x77, 0xea, 0xf0, 0x35, 0xf4,
	0x5c, 0x70, 0x79, 0x55, 0x6e, 0x1b, 0xfe, 0x0a, 0xba, 0xd7, 0xa8, 0xff, 0x75, 0xac, 0x3b, 0xda,
	0x34, 0x9d, 0xad, 0x98, 0x6d, 0x63, 0xbf, 0x80, 0x76, 0x36, 0xad, 0x1d, 0x03, 0xaf, 0xa0, 0x77,
	0x8d, 0xfa, 0x0d, 0xd3, 0xfe, 0x5d, 0x71, 0x5c, 0xcf, 0x2b, 0x64, 0xd8, 0xef, 0x6d, 0x78, 0xdf,
	0xc3, 0x27, 0xe4, 0x6b, 0x7b, 0x5b, 0x36, 0x4d, 0xf6, 0xa4, 0x76, 0xcd, 0x50, 0x2c, 0x24, 0x08,
	0xf2, 0x7b, 0xdf, 0x35, 0xcd, 0x6c, 0xdf, 0xfe, 0x4d, 0xfe, 0xec, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x74, 0x94, 0xa0, 0x37, 0x0b, 0x00, 0x00,
}
