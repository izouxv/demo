// Code generated by protoc-gen-go. DO NOT EDIT.
// source: setting.proto

package setting

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

type SetUpgradeRequest struct {
	Category      int32  `protobuf:"varint,1,opt,name=Category" json:"Category,omitempty"`
	VersionName   string `protobuf:"bytes,2,opt,name=VersionName" json:"VersionName,omitempty"`
	VersionCode   string `protobuf:"bytes,3,opt,name=VersionCode" json:"VersionCode,omitempty"`
	MD5           string `protobuf:"bytes,4,opt,name=MD5" json:"MD5,omitempty"`
	FileName      string `protobuf:"bytes,5,opt,name=FileName" json:"FileName,omitempty"`
	FileLength    int32  `protobuf:"varint,6,opt,name=FileLength" json:"FileLength,omitempty"`
	DescriptionCN string `protobuf:"bytes,7,opt,name=DescriptionCN" json:"DescriptionCN,omitempty"`
	DescriptionEN string `protobuf:"bytes,8,opt,name=DescriptionEN" json:"DescriptionEN,omitempty"`
	URL           string `protobuf:"bytes,9,opt,name=URL" json:"URL,omitempty"`
}

func (m *SetUpgradeRequest) Reset()                    { *m = SetUpgradeRequest{} }
func (m *SetUpgradeRequest) String() string            { return proto.CompactTextString(m) }
func (*SetUpgradeRequest) ProtoMessage()               {}
func (*SetUpgradeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SetUpgradeRequest) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

func (m *SetUpgradeRequest) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *SetUpgradeRequest) GetVersionCode() string {
	if m != nil {
		return m.VersionCode
	}
	return ""
}

func (m *SetUpgradeRequest) GetMD5() string {
	if m != nil {
		return m.MD5
	}
	return ""
}

func (m *SetUpgradeRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *SetUpgradeRequest) GetFileLength() int32 {
	if m != nil {
		return m.FileLength
	}
	return 0
}

func (m *SetUpgradeRequest) GetDescriptionCN() string {
	if m != nil {
		return m.DescriptionCN
	}
	return ""
}

func (m *SetUpgradeRequest) GetDescriptionEN() string {
	if m != nil {
		return m.DescriptionEN
	}
	return ""
}

func (m *SetUpgradeRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type SetUpgradeResponse struct {
	ErrorCode int32 `protobuf:"varint,1,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *SetUpgradeResponse) Reset()                    { *m = SetUpgradeResponse{} }
func (m *SetUpgradeResponse) String() string            { return proto.CompactTextString(m) }
func (*SetUpgradeResponse) ProtoMessage()               {}
func (*SetUpgradeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SetUpgradeResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type UpgradeCategory struct {
	Id   int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *UpgradeCategory) Reset()                    { *m = UpgradeCategory{} }
func (m *UpgradeCategory) String() string            { return proto.CompactTextString(m) }
func (*UpgradeCategory) ProtoMessage()               {}
func (*UpgradeCategory) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *UpgradeCategory) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpgradeCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpgrageCategoryResponse struct {
	ErrorCode int32 `protobuf:"varint,1,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *UpgrageCategoryResponse) Reset()                    { *m = UpgrageCategoryResponse{} }
func (m *UpgrageCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpgrageCategoryResponse) ProtoMessage()               {}
func (*UpgrageCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *UpgrageCategoryResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type GetUpgradeCategoryRequest struct {
}

func (m *GetUpgradeCategoryRequest) Reset()                    { *m = GetUpgradeCategoryRequest{} }
func (m *GetUpgradeCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUpgradeCategoryRequest) ProtoMessage()               {}
func (*GetUpgradeCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type GetUpgradeCategoryResponse struct {
	UC        []*UpgradeCategory `protobuf:"bytes,1,rep,name=UC" json:"UC,omitempty"`
	ErrorCode int32              `protobuf:"varint,2,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *GetUpgradeCategoryResponse) Reset()                    { *m = GetUpgradeCategoryResponse{} }
func (m *GetUpgradeCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUpgradeCategoryResponse) ProtoMessage()               {}
func (*GetUpgradeCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetUpgradeCategoryResponse) GetUC() []*UpgradeCategory {
	if m != nil {
		return m.UC
	}
	return nil
}

func (m *GetUpgradeCategoryResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type GetUpgradeInfosRequest struct {
	Count    int32 `protobuf:"varint,1,opt,name=Count" json:"Count,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Category int32 `protobuf:"varint,3,opt,name=Category" json:"Category,omitempty"`
}

func (m *GetUpgradeInfosRequest) Reset()                    { *m = GetUpgradeInfosRequest{} }
func (m *GetUpgradeInfosRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUpgradeInfosRequest) ProtoMessage()               {}
func (*GetUpgradeInfosRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *GetUpgradeInfosRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetUpgradeInfosRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUpgradeInfosRequest) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

type UpgradeInfo struct {
	Category      int32  `protobuf:"varint,1,opt,name=Category" json:"Category,omitempty"`
	VersionName   string `protobuf:"bytes,2,opt,name=VersionName" json:"VersionName,omitempty"`
	VersionCode   string `protobuf:"bytes,3,opt,name=VersionCode" json:"VersionCode,omitempty"`
	MD5           string `protobuf:"bytes,4,opt,name=MD5" json:"MD5,omitempty"`
	FileName      string `protobuf:"bytes,5,opt,name=FileName" json:"FileName,omitempty"`
	FileLength    int32  `protobuf:"varint,6,opt,name=FileLength" json:"FileLength,omitempty"`
	DescriptionCN string `protobuf:"bytes,7,opt,name=DescriptionCN" json:"DescriptionCN,omitempty"`
	DescriptionEN string `protobuf:"bytes,8,opt,name=DescriptionEN" json:"DescriptionEN,omitempty"`
	CreateTime    int64  `protobuf:"varint,9,opt,name=CreateTime" json:"CreateTime,omitempty"`
	Id            int32  `protobuf:"varint,10,opt,name=Id" json:"Id,omitempty"`
	CategoryName  string `protobuf:"bytes,11,opt,name=CategoryName" json:"CategoryName,omitempty"`
}

func (m *UpgradeInfo) Reset()                    { *m = UpgradeInfo{} }
func (m *UpgradeInfo) String() string            { return proto.CompactTextString(m) }
func (*UpgradeInfo) ProtoMessage()               {}
func (*UpgradeInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *UpgradeInfo) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

func (m *UpgradeInfo) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *UpgradeInfo) GetVersionCode() string {
	if m != nil {
		return m.VersionCode
	}
	return ""
}

func (m *UpgradeInfo) GetMD5() string {
	if m != nil {
		return m.MD5
	}
	return ""
}

func (m *UpgradeInfo) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *UpgradeInfo) GetFileLength() int32 {
	if m != nil {
		return m.FileLength
	}
	return 0
}

func (m *UpgradeInfo) GetDescriptionCN() string {
	if m != nil {
		return m.DescriptionCN
	}
	return ""
}

func (m *UpgradeInfo) GetDescriptionEN() string {
	if m != nil {
		return m.DescriptionEN
	}
	return ""
}

func (m *UpgradeInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UpgradeInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpgradeInfo) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

type GetUpgradeInfosResponse struct {
	UIS        []*UpgradeInfo `protobuf:"bytes,1,rep,name=UIS" json:"UIS,omitempty"`
	TotalCount int32          `protobuf:"varint,2,opt,name=TotalCount" json:"TotalCount,omitempty"`
	ErrorCode  int32          `protobuf:"varint,3,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *GetUpgradeInfosResponse) Reset()                    { *m = GetUpgradeInfosResponse{} }
func (m *GetUpgradeInfosResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUpgradeInfosResponse) ProtoMessage()               {}
func (*GetUpgradeInfosResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *GetUpgradeInfosResponse) GetUIS() []*UpgradeInfo {
	if m != nil {
		return m.UIS
	}
	return nil
}

func (m *GetUpgradeInfosResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *GetUpgradeInfosResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type DeleteUpgradeInfoRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
}

func (m *DeleteUpgradeInfoRequest) Reset()                    { *m = DeleteUpgradeInfoRequest{} }
func (m *DeleteUpgradeInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUpgradeInfoRequest) ProtoMessage()               {}
func (*DeleteUpgradeInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *DeleteUpgradeInfoRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteUpgradeInfoResponse struct {
	ErrorCode int32 `protobuf:"varint,1,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *DeleteUpgradeInfoResponse) Reset()                    { *m = DeleteUpgradeInfoResponse{} }
func (m *DeleteUpgradeInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteUpgradeInfoResponse) ProtoMessage()               {}
func (*DeleteUpgradeInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *DeleteUpgradeInfoResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type UpdateUpgradeInfoRequest struct {
	UI *UpgradeInfo `protobuf:"bytes,1,opt,name=UI" json:"UI,omitempty"`
}

func (m *UpdateUpgradeInfoRequest) Reset()                    { *m = UpdateUpgradeInfoRequest{} }
func (m *UpdateUpgradeInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUpgradeInfoRequest) ProtoMessage()               {}
func (*UpdateUpgradeInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *UpdateUpgradeInfoRequest) GetUI() *UpgradeInfo {
	if m != nil {
		return m.UI
	}
	return nil
}

type UpdataUpgradeInfoResponse struct {
	ErrorCode int32 `protobuf:"varint,1,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *UpdataUpgradeInfoResponse) Reset()                    { *m = UpdataUpgradeInfoResponse{} }
func (m *UpdataUpgradeInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdataUpgradeInfoResponse) ProtoMessage()               {}
func (*UpdataUpgradeInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *UpdataUpgradeInfoResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type NodeAdv struct {
	Id  int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Md5 string `protobuf:"bytes,2,opt,name=Md5" json:"Md5,omitempty"`
	Url string `protobuf:"bytes,3,opt,name=Url" json:"Url,omitempty"`
}

func (m *NodeAdv) Reset()                    { *m = NodeAdv{} }
func (m *NodeAdv) String() string            { return proto.CompactTextString(m) }
func (*NodeAdv) ProtoMessage()               {}
func (*NodeAdv) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *NodeAdv) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NodeAdv) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *NodeAdv) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type NodeAdvInfoRequest struct {
	DevEUI  string     `protobuf:"bytes,1,opt,name=DevEUI" json:"DevEUI,omitempty"`
	NodeAdv []*NodeAdv `protobuf:"bytes,2,rep,name=NodeAdv" json:"NodeAdv,omitempty"`
}

func (m *NodeAdvInfoRequest) Reset()                    { *m = NodeAdvInfoRequest{} }
func (m *NodeAdvInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*NodeAdvInfoRequest) ProtoMessage()               {}
func (*NodeAdvInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *NodeAdvInfoRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *NodeAdvInfoRequest) GetNodeAdv() []*NodeAdv {
	if m != nil {
		return m.NodeAdv
	}
	return nil
}

type NodeAdvInfoResponse struct {
	DevEUI    string     `protobuf:"bytes,1,opt,name=DevEUI" json:"DevEUI,omitempty"`
	NodeAdv   []*NodeAdv `protobuf:"bytes,2,rep,name=NodeAdv" json:"NodeAdv,omitempty"`
	ErrorCode int32      `protobuf:"varint,3,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
}

func (m *NodeAdvInfoResponse) Reset()                    { *m = NodeAdvInfoResponse{} }
func (m *NodeAdvInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeAdvInfoResponse) ProtoMessage()               {}
func (*NodeAdvInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *NodeAdvInfoResponse) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *NodeAdvInfoResponse) GetNodeAdv() []*NodeAdv {
	if m != nil {
		return m.NodeAdv
	}
	return nil
}

func (m *NodeAdvInfoResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*SetUpgradeRequest)(nil), "setting.SetUpgradeRequest")
	proto.RegisterType((*SetUpgradeResponse)(nil), "setting.SetUpgradeResponse")
	proto.RegisterType((*UpgradeCategory)(nil), "setting.UpgradeCategory")
	proto.RegisterType((*UpgrageCategoryResponse)(nil), "setting.UpgrageCategoryResponse")
	proto.RegisterType((*GetUpgradeCategoryRequest)(nil), "setting.GetUpgradeCategoryRequest")
	proto.RegisterType((*GetUpgradeCategoryResponse)(nil), "setting.GetUpgradeCategoryResponse")
	proto.RegisterType((*GetUpgradeInfosRequest)(nil), "setting.GetUpgradeInfosRequest")
	proto.RegisterType((*UpgradeInfo)(nil), "setting.UpgradeInfo")
	proto.RegisterType((*GetUpgradeInfosResponse)(nil), "setting.GetUpgradeInfosResponse")
	proto.RegisterType((*DeleteUpgradeInfoRequest)(nil), "setting.DeleteUpgradeInfoRequest")
	proto.RegisterType((*DeleteUpgradeInfoResponse)(nil), "setting.DeleteUpgradeInfoResponse")
	proto.RegisterType((*UpdateUpgradeInfoRequest)(nil), "setting.UpdateUpgradeInfoRequest")
	proto.RegisterType((*UpdataUpgradeInfoResponse)(nil), "setting.UpdataUpgradeInfoResponse")
	proto.RegisterType((*NodeAdv)(nil), "setting.NodeAdv")
	proto.RegisterType((*NodeAdvInfoRequest)(nil), "setting.NodeAdvInfoRequest")
	proto.RegisterType((*NodeAdvInfoResponse)(nil), "setting.NodeAdvInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Setting service

type SettingClient interface {
	SetUpgradeInfo(ctx context.Context, in *SetUpgradeRequest, opts ...grpc.CallOption) (*SetUpgradeResponse, error)
	GetUpgradeCategory(ctx context.Context, in *GetUpgradeCategoryRequest, opts ...grpc.CallOption) (*GetUpgradeCategoryResponse, error)
	GetUpgradeInfos(ctx context.Context, in *GetUpgradeInfosRequest, opts ...grpc.CallOption) (*GetUpgradeInfosResponse, error)
	DeleteUpgradeInfo(ctx context.Context, in *DeleteUpgradeInfoRequest, opts ...grpc.CallOption) (*DeleteUpgradeInfoResponse, error)
	UpdateUpgradeInfo(ctx context.Context, in *UpdateUpgradeInfoRequest, opts ...grpc.CallOption) (*UpdataUpgradeInfoResponse, error)
	NodeAdvInfo(ctx context.Context, in *NodeAdvInfoRequest, opts ...grpc.CallOption) (*NodeAdvInfoResponse, error)
}

type settingClient struct {
	cc *grpc.ClientConn
}

func NewSettingClient(cc *grpc.ClientConn) SettingClient {
	return &settingClient{cc}
}

func (c *settingClient) SetUpgradeInfo(ctx context.Context, in *SetUpgradeRequest, opts ...grpc.CallOption) (*SetUpgradeResponse, error) {
	out := new(SetUpgradeResponse)
	err := grpc.Invoke(ctx, "/setting.Setting/SetUpgradeInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) GetUpgradeCategory(ctx context.Context, in *GetUpgradeCategoryRequest, opts ...grpc.CallOption) (*GetUpgradeCategoryResponse, error) {
	out := new(GetUpgradeCategoryResponse)
	err := grpc.Invoke(ctx, "/setting.Setting/GetUpgradeCategory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) GetUpgradeInfos(ctx context.Context, in *GetUpgradeInfosRequest, opts ...grpc.CallOption) (*GetUpgradeInfosResponse, error) {
	out := new(GetUpgradeInfosResponse)
	err := grpc.Invoke(ctx, "/setting.Setting/GetUpgradeInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) DeleteUpgradeInfo(ctx context.Context, in *DeleteUpgradeInfoRequest, opts ...grpc.CallOption) (*DeleteUpgradeInfoResponse, error) {
	out := new(DeleteUpgradeInfoResponse)
	err := grpc.Invoke(ctx, "/setting.Setting/DeleteUpgradeInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) UpdateUpgradeInfo(ctx context.Context, in *UpdateUpgradeInfoRequest, opts ...grpc.CallOption) (*UpdataUpgradeInfoResponse, error) {
	out := new(UpdataUpgradeInfoResponse)
	err := grpc.Invoke(ctx, "/setting.Setting/UpdateUpgradeInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) NodeAdvInfo(ctx context.Context, in *NodeAdvInfoRequest, opts ...grpc.CallOption) (*NodeAdvInfoResponse, error) {
	out := new(NodeAdvInfoResponse)
	err := grpc.Invoke(ctx, "/setting.Setting/NodeAdvInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Setting service

type SettingServer interface {
	SetUpgradeInfo(context.Context, *SetUpgradeRequest) (*SetUpgradeResponse, error)
	GetUpgradeCategory(context.Context, *GetUpgradeCategoryRequest) (*GetUpgradeCategoryResponse, error)
	GetUpgradeInfos(context.Context, *GetUpgradeInfosRequest) (*GetUpgradeInfosResponse, error)
	DeleteUpgradeInfo(context.Context, *DeleteUpgradeInfoRequest) (*DeleteUpgradeInfoResponse, error)
	UpdateUpgradeInfo(context.Context, *UpdateUpgradeInfoRequest) (*UpdataUpgradeInfoResponse, error)
	NodeAdvInfo(context.Context, *NodeAdvInfoRequest) (*NodeAdvInfoResponse, error)
}

func RegisterSettingServer(s *grpc.Server, srv SettingServer) {
	s.RegisterService(&_Setting_serviceDesc, srv)
}

func _Setting_SetUpgradeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).SetUpgradeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/setting.Setting/SetUpgradeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).SetUpgradeInfo(ctx, req.(*SetUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_GetUpgradeCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpgradeCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).GetUpgradeCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/setting.Setting/GetUpgradeCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).GetUpgradeCategory(ctx, req.(*GetUpgradeCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_GetUpgradeInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpgradeInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).GetUpgradeInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/setting.Setting/GetUpgradeInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).GetUpgradeInfos(ctx, req.(*GetUpgradeInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_DeleteUpgradeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUpgradeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).DeleteUpgradeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/setting.Setting/DeleteUpgradeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).DeleteUpgradeInfo(ctx, req.(*DeleteUpgradeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_UpdateUpgradeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUpgradeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).UpdateUpgradeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/setting.Setting/UpdateUpgradeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).UpdateUpgradeInfo(ctx, req.(*UpdateUpgradeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_NodeAdvInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeAdvInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).NodeAdvInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/setting.Setting/NodeAdvInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).NodeAdvInfo(ctx, req.(*NodeAdvInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Setting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "setting.Setting",
	HandlerType: (*SettingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUpgradeInfo",
			Handler:    _Setting_SetUpgradeInfo_Handler,
		},
		{
			MethodName: "GetUpgradeCategory",
			Handler:    _Setting_GetUpgradeCategory_Handler,
		},
		{
			MethodName: "GetUpgradeInfos",
			Handler:    _Setting_GetUpgradeInfos_Handler,
		},
		{
			MethodName: "DeleteUpgradeInfo",
			Handler:    _Setting_DeleteUpgradeInfo_Handler,
		},
		{
			MethodName: "UpdateUpgradeInfo",
			Handler:    _Setting_UpdateUpgradeInfo_Handler,
		},
		{
			MethodName: "NodeAdvInfo",
			Handler:    _Setting_NodeAdvInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "setting.proto",
}

func init() { proto.RegisterFile("setting.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0x93, 0x6d, 0x5d, 0x5f, 0xd9, 0x2f, 0x33, 0x6d, 0x5e, 0x36, 0x8d, 0x62, 0x26, 0x54,
	0xed, 0xb0, 0xc3, 0xd0, 0x84, 0x38, 0x20, 0x81, 0xda, 0x82, 0x02, 0xa3, 0x87, 0x74, 0x99, 0x38,
	0x20, 0x50, 0x20, 0x26, 0x54, 0xea, 0x92, 0x90, 0x78, 0x43, 0x9c, 0xf8, 0x33, 0x38, 0xf2, 0x87,
	0xf1, 0xcf, 0x20, 0x3b, 0x8e, 0xe7, 0xfc, 0xaa, 0x86, 0xc4, 0x8d, 0xdb, 0xf3, 0xcb, 0xf7, 0xbe,
	0xf7, 0xd9, 0xef, 0xb3, 0x5b, 0x58, 0x4d, 0x29, 0x63, 0xd3, 0x30, 0x38, 0x8e, 0x93, 0x88, 0x45,
	0xa8, 0x2d, 0x97, 0xe4, 0x97, 0x01, 0x9b, 0x13, 0xca, 0xdc, 0x38, 0x48, 0x3c, 0x9f, 0x3a, 0xf4,
	0xeb, 0x15, 0x4d, 0x19, 0xb2, 0x60, 0x65, 0xe0, 0x31, 0x1a, 0x44, 0xc9, 0x77, 0xdc, 0xea, 0xb5,
	0xfa, 0x4b, 0x8e, 0x5a, 0xa3, 0x1e, 0x74, 0x2f, 0x68, 0x92, 0x4e, 0xa3, 0x70, 0xec, 0x5d, 0x52,
	0x6c, 0xf4, 0x5a, 0xfd, 0x8e, 0xa3, 0xa7, 0x34, 0xc4, 0x20, 0xf2, 0x29, 0x36, 0x0b, 0x08, 0x9e,
	0x42, 0x1b, 0x60, 0xbe, 0x19, 0x9e, 0xe2, 0x45, 0xf1, 0x85, 0x87, 0xbc, 0xe3, 0x8b, 0xe9, 0x8c,
	0x0a, 0xca, 0x25, 0x91, 0x56, 0x6b, 0x74, 0x00, 0xc0, 0xe3, 0x33, 0x1a, 0x06, 0xec, 0x0b, 0x5e,
	0x16, 0x7a, 0xb4, 0x0c, 0x3a, 0x84, 0xd5, 0x21, 0x4d, 0x3f, 0x25, 0xd3, 0x98, 0xf1, 0x06, 0x63,
	0xdc, 0x16, 0x04, 0xc5, 0x64, 0x09, 0x35, 0x1a, 0xe3, 0x95, 0x0a, 0x6a, 0x34, 0xe6, 0xca, 0x5c,
	0xe7, 0x0c, 0x77, 0x32, 0x65, 0xae, 0x73, 0x46, 0x4e, 0x00, 0xe9, 0x07, 0x94, 0xc6, 0x51, 0x98,
	0x52, 0xb4, 0x0f, 0x9d, 0x51, 0x92, 0x44, 0x89, 0xd8, 0x61, 0x76, 0x44, 0x37, 0x09, 0x72, 0x0a,
	0xeb, 0xb2, 0x40, 0x1d, 0xdb, 0x1a, 0x18, 0xb6, 0x2f, 0x91, 0x86, 0xed, 0x23, 0x04, 0x8b, 0xda,
	0xf9, 0x89, 0x98, 0x3c, 0x86, 0x1d, 0x51, 0x16, 0xa8, 0xb2, 0x5b, 0xf6, 0xdb, 0x83, 0xdd, 0x97,
	0x4a, 0xe3, 0x4d, 0xad, 0x18, 0x26, 0xf1, 0xc1, 0xaa, 0xfb, 0x28, 0x89, 0xfb, 0x60, 0xb8, 0x03,
	0xdc, 0xea, 0x99, 0xfd, 0xee, 0x09, 0x3e, 0xce, 0x5d, 0x52, 0x46, 0x1b, 0xee, 0xa0, 0x28, 0xc1,
	0x28, 0x4b, 0x78, 0x0f, 0xdb, 0x37, 0x5d, 0xec, 0xf0, 0x73, 0x94, 0xe6, 0x66, 0xda, 0x82, 0xa5,
	0x41, 0x74, 0x15, 0x32, 0x29, 0x3b, 0x5b, 0xf0, 0xfd, 0xc7, 0x5e, 0x90, 0x13, 0x89, 0xb8, 0x60,
	0x3b, 0xb3, 0x68, 0x3b, 0xf2, 0xdb, 0x80, 0xae, 0xc6, 0xfe, 0x1f, 0x5b, 0xf4, 0x00, 0x60, 0x90,
	0x50, 0x8f, 0xd1, 0xf3, 0xe9, 0x25, 0x15, 0x4e, 0x35, 0x1d, 0x2d, 0x23, 0x9d, 0x06, 0xca, 0x69,
	0x04, 0xee, 0xe4, 0x27, 0x23, 0xb4, 0x77, 0x05, 0x69, 0x21, 0x47, 0x7e, 0xc0, 0x4e, 0x65, 0x7a,
	0xd2, 0x20, 0x0f, 0xc1, 0x74, 0xed, 0x89, 0x74, 0xc8, 0x56, 0xd9, 0x21, 0x1c, 0xeb, 0x70, 0x00,
	0x97, 0x75, 0x1e, 0x31, 0x6f, 0x96, 0xcd, 0x3a, 0x1b, 0xab, 0x96, 0x29, 0xda, 0xc7, 0x2c, 0xdb,
	0xe7, 0x08, 0xf0, 0x90, 0xce, 0x28, 0xa3, 0x3a, 0xaf, 0x34, 0x50, 0xe9, 0xea, 0x90, 0x27, 0xb0,
	0x5b, 0x83, 0xbd, 0xd5, 0x45, 0x79, 0x06, 0xd8, 0x8d, 0x7d, 0xaf, 0xb6, 0xcd, 0x21, 0x18, 0xae,
	0x2d, 0x4a, 0x9a, 0xf6, 0x69, 0xb8, 0x36, 0x6f, 0x2e, 0x18, 0xbc, 0xbf, 0x6f, 0xfe, 0x14, 0xda,
	0xe3, 0xc8, 0xa7, 0xcf, 0xfd, 0xeb, 0xca, 0x6b, 0xc0, 0xdd, 0xe6, 0x9f, 0x4a, 0xa7, 0xf2, 0x50,
	0x3c, 0x44, 0xc9, 0x4c, 0x3a, 0x93, 0x87, 0xe4, 0x2d, 0x20, 0x59, 0xae, 0xab, 0xde, 0x86, 0xe5,
	0x21, 0xbd, 0x1e, 0x49, 0xe5, 0x1d, 0x47, 0xae, 0xd0, 0x91, 0x6a, 0x86, 0x0d, 0x31, 0xba, 0x0d,
	0xb5, 0x25, 0x99, 0x77, 0x72, 0x00, 0xf9, 0x06, 0x77, 0x0b, 0xcc, 0x72, 0x37, 0xff, 0x80, 0x7a,
	0xfe, 0xd4, 0x4f, 0x7e, 0x2e, 0x42, 0x7b, 0x92, 0x95, 0xa2, 0xd7, 0xb0, 0x36, 0x29, 0x58, 0x10,
	0x59, 0x8a, 0xb6, 0xf2, 0x0b, 0x65, 0xed, 0xd5, 0x7e, 0xcb, 0x84, 0x93, 0x05, 0xf4, 0x01, 0x50,
	0xf5, 0xcd, 0x43, 0x44, 0x15, 0x35, 0xbe, 0x96, 0xd6, 0x83, 0xb9, 0x18, 0xd5, 0xe0, 0x02, 0xd6,
	0x4b, 0x17, 0x06, 0xdd, 0xab, 0xa9, 0xd4, 0x1f, 0x42, 0xab, 0xd7, 0x0c, 0x50, 0xbc, 0xef, 0x60,
	0xb3, 0xe2, 0x6d, 0x74, 0x5f, 0x15, 0x36, 0xdd, 0x11, 0x8b, 0xcc, 0x83, 0xe8, 0xec, 0x15, 0xfb,
	0x6b, 0xec, 0x4d, 0x57, 0x43, 0x63, 0x6f, 0xf4, 0x3e, 0x59, 0x40, 0xaf, 0xa0, 0xab, 0xd9, 0x08,
	0xed, 0x95, 0x5d, 0xa1, 0x33, 0xee, 0xd7, 0x7f, 0xcc, 0xb9, 0x3e, 0x2e, 0x8b, 0xff, 0x29, 0x8f,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x97, 0x36, 0xff, 0xb8, 0x08, 0x00, 0x00,
}