// Code generated by protoc-gen-go.
// source: api/domain.proto
// DO NOT EDIT!

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

type Domain struct {
	Did         int32  `protobuf:"varint,1,opt,name=did" json:"did,omitempty"`
	DomainName  string `protobuf:"bytes,2,opt,name=domainName" json:"domainName,omitempty"`
	CreateTime  int64  `protobuf:"varint,3,opt,name=CreateTime" json:"CreateTime,omitempty"`
	UpdateTime  int64  `protobuf:"varint,4,opt,name=UpdateTime" json:"UpdateTime,omitempty"`
	DomainUrl   string `protobuf:"bytes,5,opt,name=domainUrl" json:"domainUrl,omitempty"`
	DomainState int32  `protobuf:"varint,6,opt,name=domainState" json:"domainState,omitempty"`
}

func (m *Domain) Reset()                    { *m = Domain{} }
func (m *Domain) String() string            { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()               {}
func (*Domain) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Domain) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *Domain) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

func (m *Domain) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Domain) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Domain) GetDomainUrl() string {
	if m != nil {
		return m.DomainUrl
	}
	return ""
}

func (m *Domain) GetDomainState() int32 {
	if m != nil {
		return m.DomainState
	}
	return 0
}

type GetDomainRequest struct {
	Did int32 `protobuf:"varint,1,opt,name=did" json:"did,omitempty"`
	Uid int32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetDomainRequest) Reset()                    { *m = GetDomainRequest{} }
func (m *GetDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDomainRequest) ProtoMessage()               {}
func (*GetDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *GetDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *GetDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GetDomainResponse struct {
	Domain *Domain `protobuf:"bytes,1,opt,name=Domain" json:"Domain,omitempty"`
}

func (m *GetDomainResponse) Reset()                    { *m = GetDomainResponse{} }
func (m *GetDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDomainResponse) ProtoMessage()               {}
func (*GetDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *GetDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

type UpdateDomainRequest struct {
	Did        int32  `protobuf:"varint,1,opt,name=did" json:"did,omitempty"`
	Uid        int32  `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	DomainName string `protobuf:"bytes,3,opt,name=domainName" json:"domainName,omitempty"`
}

func (m *UpdateDomainRequest) Reset()                    { *m = UpdateDomainRequest{} }
func (m *UpdateDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDomainRequest) ProtoMessage()               {}
func (*UpdateDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *UpdateDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *UpdateDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateDomainRequest) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

type UpdateDomainResponse struct {
	Domain *Domain `protobuf:"bytes,1,opt,name=Domain" json:"Domain,omitempty"`
}

func (m *UpdateDomainResponse) Reset()                    { *m = UpdateDomainResponse{} }
func (m *UpdateDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateDomainResponse) ProtoMessage()               {}
func (*UpdateDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *UpdateDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

type DeleteDomainRequest struct {
	Did int32 `protobuf:"varint,1,opt,name=did" json:"did,omitempty"`
	Uid int32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *DeleteDomainRequest) Reset()                    { *m = DeleteDomainRequest{} }
func (m *DeleteDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDomainRequest) ProtoMessage()               {}
func (*DeleteDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *DeleteDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *DeleteDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type DeleteDomainResponse struct {
}

func (m *DeleteDomainResponse) Reset()                    { *m = DeleteDomainResponse{} }
func (m *DeleteDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDomainResponse) ProtoMessage()               {}
func (*DeleteDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

type GetUserCountInDomainRequest struct {
	Did int32 `protobuf:"varint,1,opt,name=did" json:"did,omitempty"`
	Uid int32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetUserCountInDomainRequest) Reset()                    { *m = GetUserCountInDomainRequest{} }
func (m *GetUserCountInDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserCountInDomainRequest) ProtoMessage()               {}
func (*GetUserCountInDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *GetUserCountInDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *GetUserCountInDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GetUserCountInDomainResponse struct {
	Count int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *GetUserCountInDomainResponse) Reset()                    { *m = GetUserCountInDomainResponse{} }
func (m *GetUserCountInDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserCountInDomainResponse) ProtoMessage()               {}
func (*GetUserCountInDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *GetUserCountInDomainResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetDomainsRequest struct {
	Uid int32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Did int32 `protobuf:"varint,2,opt,name=did" json:"did,omitempty"`
}

func (m *GetDomainsRequest) Reset()                    { *m = GetDomainsRequest{} }
func (m *GetDomainsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDomainsRequest) ProtoMessage()               {}
func (*GetDomainsRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *GetDomainsRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GetDomainsRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

type GetDomainsResponse struct {
	Domain []*Domain `protobuf:"bytes,1,rep,name=Domain" json:"Domain,omitempty"`
}

func (m *GetDomainsResponse) Reset()                    { *m = GetDomainsResponse{} }
func (m *GetDomainsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDomainsResponse) ProtoMessage()               {}
func (*GetDomainsResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *GetDomainsResponse) GetDomain() []*Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

type GetUserInfoInDomainRequest struct {
	Did   int32 `protobuf:"varint,1,opt,name=did" json:"did,omitempty"`
	Uid   int32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	Page  int32 `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	Count int32 `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *GetUserInfoInDomainRequest) Reset()                    { *m = GetUserInfoInDomainRequest{} }
func (m *GetUserInfoInDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoInDomainRequest) ProtoMessage()               {}
func (*GetUserInfoInDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *GetUserInfoInDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *GetUserInfoInDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GetUserInfoInDomainRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUserInfoInDomainRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetUserInfoInDomainResponse struct {
	UserRoles  []*UserRoles `protobuf:"bytes,1,rep,name=userRoles" json:"userRoles,omitempty"`
	TotalCount int32        `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *GetUserInfoInDomainResponse) Reset()                    { *m = GetUserInfoInDomainResponse{} }
func (m *GetUserInfoInDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoInDomainResponse) ProtoMessage()               {}
func (*GetUserInfoInDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *GetUserInfoInDomainResponse) GetUserRoles() []*UserRoles {
	if m != nil {
		return m.UserRoles
	}
	return nil
}

func (m *GetUserInfoInDomainResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type UpdateUserRoleInDomainRequest struct {
	Uid            int32   `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	UpdateUserID   int32   `protobuf:"varint,2,opt,name=updateUserID" json:"updateUserID,omitempty"`
	UpdateUserRids []int32 `protobuf:"varint,3,rep,packed,name=updateUserRids" json:"updateUserRids,omitempty"`
	Did            int32   `protobuf:"varint,4,opt,name=did" json:"did,omitempty"`
}

func (m *UpdateUserRoleInDomainRequest) Reset()                    { *m = UpdateUserRoleInDomainRequest{} }
func (m *UpdateUserRoleInDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRoleInDomainRequest) ProtoMessage()               {}
func (*UpdateUserRoleInDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *UpdateUserRoleInDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateUserRoleInDomainRequest) GetUpdateUserID() int32 {
	if m != nil {
		return m.UpdateUserID
	}
	return 0
}

func (m *UpdateUserRoleInDomainRequest) GetUpdateUserRids() []int32 {
	if m != nil {
		return m.UpdateUserRids
	}
	return nil
}

func (m *UpdateUserRoleInDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

type UpdateUserRoleInDomainResponse struct {
}

func (m *UpdateUserRoleInDomainResponse) Reset()                    { *m = UpdateUserRoleInDomainResponse{} }
func (m *UpdateUserRoleInDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRoleInDomainResponse) ProtoMessage()               {}
func (*UpdateUserRoleInDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

type DeleteUserInDomainRequest struct {
	Uid          int32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	DeleteUserID int32 `protobuf:"varint,2,opt,name=deleteUserID" json:"deleteUserID,omitempty"`
	Did          int32 `protobuf:"varint,3,opt,name=did" json:"did,omitempty"`
}

func (m *DeleteUserInDomainRequest) Reset()                    { *m = DeleteUserInDomainRequest{} }
func (m *DeleteUserInDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserInDomainRequest) ProtoMessage()               {}
func (*DeleteUserInDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *DeleteUserInDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *DeleteUserInDomainRequest) GetDeleteUserID() int32 {
	if m != nil {
		return m.DeleteUserID
	}
	return 0
}

func (m *DeleteUserInDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

type DeleteUserInDomainResponse struct {
}

func (m *DeleteUserInDomainResponse) Reset()                    { *m = DeleteUserInDomainResponse{} }
func (m *DeleteUserInDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserInDomainResponse) ProtoMessage()               {}
func (*DeleteUserInDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

type AddUserInDomainRequest struct {
	Did             int32   `protobuf:"varint,1,opt,name=did" json:"did,omitempty"`
	Uid             int32   `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	AddUserUsername string  `protobuf:"bytes,3,opt,name=addUserUsername" json:"addUserUsername,omitempty"`
	AddUserNickname string  `protobuf:"bytes,4,opt,name=addUserNickname" json:"addUserNickname,omitempty"`
	AddUserRids     []int32 `protobuf:"varint,5,rep,packed,name=addUserRids" json:"addUserRids,omitempty"`
}

func (m *AddUserInDomainRequest) Reset()                    { *m = AddUserInDomainRequest{} }
func (m *AddUserInDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserInDomainRequest) ProtoMessage()               {}
func (*AddUserInDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

func (m *AddUserInDomainRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

func (m *AddUserInDomainRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AddUserInDomainRequest) GetAddUserUsername() string {
	if m != nil {
		return m.AddUserUsername
	}
	return ""
}

func (m *AddUserInDomainRequest) GetAddUserNickname() string {
	if m != nil {
		return m.AddUserNickname
	}
	return ""
}

func (m *AddUserInDomainRequest) GetAddUserRids() []int32 {
	if m != nil {
		return m.AddUserRids
	}
	return nil
}

type AddUserInDomainResponse struct {
}

func (m *AddUserInDomainResponse) Reset()                    { *m = AddUserInDomainResponse{} }
func (m *AddUserInDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUserInDomainResponse) ProtoMessage()               {}
func (*AddUserInDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

type EnterDomainRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *EnterDomainRequest) Reset()                    { *m = EnterDomainRequest{} }
func (m *EnterDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*EnterDomainRequest) ProtoMessage()               {}
func (*EnterDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

func (m *EnterDomainRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *EnterDomainRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *EnterDomainRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type EnterDomainResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *EnterDomainResponse) Reset()                    { *m = EnterDomainResponse{} }
func (m *EnterDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*EnterDomainResponse) ProtoMessage()               {}
func (*EnterDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{20} }

func (m *EnterDomainResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AddUserTenantACLRequest struct {
	Username string  `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Tids     []int32 `protobuf:"varint,2,rep,packed,name=Tids" json:"Tids,omitempty"`
	Did      int32   `protobuf:"varint,3,opt,name=Did" json:"Did,omitempty"`
}

func (m *AddUserTenantACLRequest) Reset()                    { *m = AddUserTenantACLRequest{} }
func (m *AddUserTenantACLRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserTenantACLRequest) ProtoMessage()               {}
func (*AddUserTenantACLRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{21} }

func (m *AddUserTenantACLRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddUserTenantACLRequest) GetTids() []int32 {
	if m != nil {
		return m.Tids
	}
	return nil
}

func (m *AddUserTenantACLRequest) GetDid() int32 {
	if m != nil {
		return m.Did
	}
	return 0
}

type AddUserTenantACLResponse struct {
}

func (m *AddUserTenantACLResponse) Reset()                    { *m = AddUserTenantACLResponse{} }
func (m *AddUserTenantACLResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUserTenantACLResponse) ProtoMessage()               {}
func (*AddUserTenantACLResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{22} }

func init() {
	proto.RegisterType((*Domain)(nil), "api.Domain")
	proto.RegisterType((*GetDomainRequest)(nil), "api.GetDomainRequest")
	proto.RegisterType((*GetDomainResponse)(nil), "api.GetDomainResponse")
	proto.RegisterType((*UpdateDomainRequest)(nil), "api.UpdateDomainRequest")
	proto.RegisterType((*UpdateDomainResponse)(nil), "api.UpdateDomainResponse")
	proto.RegisterType((*DeleteDomainRequest)(nil), "api.DeleteDomainRequest")
	proto.RegisterType((*DeleteDomainResponse)(nil), "api.DeleteDomainResponse")
	proto.RegisterType((*GetUserCountInDomainRequest)(nil), "api.GetUserCountInDomainRequest")
	proto.RegisterType((*GetUserCountInDomainResponse)(nil), "api.GetUserCountInDomainResponse")
	proto.RegisterType((*GetDomainsRequest)(nil), "api.GetDomainsRequest")
	proto.RegisterType((*GetDomainsResponse)(nil), "api.GetDomainsResponse")
	proto.RegisterType((*GetUserInfoInDomainRequest)(nil), "api.GetUserInfoInDomainRequest")
	proto.RegisterType((*GetUserInfoInDomainResponse)(nil), "api.GetUserInfoInDomainResponse")
	proto.RegisterType((*UpdateUserRoleInDomainRequest)(nil), "api.UpdateUserRoleInDomainRequest")
	proto.RegisterType((*UpdateUserRoleInDomainResponse)(nil), "api.UpdateUserRoleInDomainResponse")
	proto.RegisterType((*DeleteUserInDomainRequest)(nil), "api.DeleteUserInDomainRequest")
	proto.RegisterType((*DeleteUserInDomainResponse)(nil), "api.DeleteUserInDomainResponse")
	proto.RegisterType((*AddUserInDomainRequest)(nil), "api.AddUserInDomainRequest")
	proto.RegisterType((*AddUserInDomainResponse)(nil), "api.AddUserInDomainResponse")
	proto.RegisterType((*EnterDomainRequest)(nil), "api.EnterDomainRequest")
	proto.RegisterType((*EnterDomainResponse)(nil), "api.EnterDomainResponse")
	proto.RegisterType((*AddUserTenantACLRequest)(nil), "api.AddUserTenantACLRequest")
	proto.RegisterType((*AddUserTenantACLResponse)(nil), "api.AddUserTenantACLResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DomainServer service

type DomainServerClient interface {
	GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*GetDomainResponse, error)
	GetDomains(ctx context.Context, in *GetDomainsRequest, opts ...grpc.CallOption) (*GetDomainsResponse, error)
	GetUserCountInDomain(ctx context.Context, in *GetUserCountInDomainRequest, opts ...grpc.CallOption) (*GetUserCountInDomainResponse, error)
	GetUserInfoInDomain(ctx context.Context, in *GetUserInfoInDomainRequest, opts ...grpc.CallOption) (*GetUserInfoInDomainResponse, error)
	UpdateUserRoleInDomain(ctx context.Context, in *UpdateUserRoleInDomainRequest, opts ...grpc.CallOption) (*UpdateUserRoleInDomainResponse, error)
	AddUserInDomain(ctx context.Context, in *AddUserInDomainRequest, opts ...grpc.CallOption) (*AddUserInDomainResponse, error)
	DeleteUserInDomain(ctx context.Context, in *DeleteUserInDomainRequest, opts ...grpc.CallOption) (*DeleteUserInDomainResponse, error)
	EnterDomain(ctx context.Context, in *EnterDomainRequest, opts ...grpc.CallOption) (*EnterDomainResponse, error)
	AddUserTenantACL(ctx context.Context, in *AddUserTenantACLRequest, opts ...grpc.CallOption) (*AddUserTenantACLResponse, error)
}

type domainServerClient struct {
	cc *grpc.ClientConn
}

func NewDomainServerClient(cc *grpc.ClientConn) DomainServerClient {
	return &domainServerClient{cc}
}

func (c *domainServerClient) GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*GetDomainResponse, error) {
	out := new(GetDomainResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/GetDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) GetDomains(ctx context.Context, in *GetDomainsRequest, opts ...grpc.CallOption) (*GetDomainsResponse, error) {
	out := new(GetDomainsResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/GetDomains", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) GetUserCountInDomain(ctx context.Context, in *GetUserCountInDomainRequest, opts ...grpc.CallOption) (*GetUserCountInDomainResponse, error) {
	out := new(GetUserCountInDomainResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/GetUserCountInDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) GetUserInfoInDomain(ctx context.Context, in *GetUserInfoInDomainRequest, opts ...grpc.CallOption) (*GetUserInfoInDomainResponse, error) {
	out := new(GetUserInfoInDomainResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/GetUserInfoInDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) UpdateUserRoleInDomain(ctx context.Context, in *UpdateUserRoleInDomainRequest, opts ...grpc.CallOption) (*UpdateUserRoleInDomainResponse, error) {
	out := new(UpdateUserRoleInDomainResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/UpdateUserRoleInDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) AddUserInDomain(ctx context.Context, in *AddUserInDomainRequest, opts ...grpc.CallOption) (*AddUserInDomainResponse, error) {
	out := new(AddUserInDomainResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/AddUserInDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) DeleteUserInDomain(ctx context.Context, in *DeleteUserInDomainRequest, opts ...grpc.CallOption) (*DeleteUserInDomainResponse, error) {
	out := new(DeleteUserInDomainResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/DeleteUserInDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) EnterDomain(ctx context.Context, in *EnterDomainRequest, opts ...grpc.CallOption) (*EnterDomainResponse, error) {
	out := new(EnterDomainResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/EnterDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServerClient) AddUserTenantACL(ctx context.Context, in *AddUserTenantACLRequest, opts ...grpc.CallOption) (*AddUserTenantACLResponse, error) {
	out := new(AddUserTenantACLResponse)
	err := grpc.Invoke(ctx, "/api.DomainServer/AddUserTenantACL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DomainServer service

type DomainServerServer interface {
	GetDomain(context.Context, *GetDomainRequest) (*GetDomainResponse, error)
	GetDomains(context.Context, *GetDomainsRequest) (*GetDomainsResponse, error)
	GetUserCountInDomain(context.Context, *GetUserCountInDomainRequest) (*GetUserCountInDomainResponse, error)
	GetUserInfoInDomain(context.Context, *GetUserInfoInDomainRequest) (*GetUserInfoInDomainResponse, error)
	UpdateUserRoleInDomain(context.Context, *UpdateUserRoleInDomainRequest) (*UpdateUserRoleInDomainResponse, error)
	AddUserInDomain(context.Context, *AddUserInDomainRequest) (*AddUserInDomainResponse, error)
	DeleteUserInDomain(context.Context, *DeleteUserInDomainRequest) (*DeleteUserInDomainResponse, error)
	EnterDomain(context.Context, *EnterDomainRequest) (*EnterDomainResponse, error)
	AddUserTenantACL(context.Context, *AddUserTenantACLRequest) (*AddUserTenantACLResponse, error)
}

func RegisterDomainServerServer(s *grpc.Server, srv DomainServerServer) {
	s.RegisterService(&_DomainServer_serviceDesc, srv)
}

func _DomainServer_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/GetDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).GetDomain(ctx, req.(*GetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_GetDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).GetDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/GetDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).GetDomains(ctx, req.(*GetDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_GetUserCountInDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCountInDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).GetUserCountInDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/GetUserCountInDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).GetUserCountInDomain(ctx, req.(*GetUserCountInDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_GetUserInfoInDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoInDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).GetUserInfoInDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/GetUserInfoInDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).GetUserInfoInDomain(ctx, req.(*GetUserInfoInDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_UpdateUserRoleInDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleInDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).UpdateUserRoleInDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/UpdateUserRoleInDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).UpdateUserRoleInDomain(ctx, req.(*UpdateUserRoleInDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_AddUserInDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserInDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).AddUserInDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/AddUserInDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).AddUserInDomain(ctx, req.(*AddUserInDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_DeleteUserInDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserInDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).DeleteUserInDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/DeleteUserInDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).DeleteUserInDomain(ctx, req.(*DeleteUserInDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_EnterDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).EnterDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/EnterDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).EnterDomain(ctx, req.(*EnterDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainServer_AddUserTenantACL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserTenantACLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServerServer).AddUserTenantACL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DomainServer/AddUserTenantACL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServerServer).AddUserTenantACL(ctx, req.(*AddUserTenantACLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DomainServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DomainServer",
	HandlerType: (*DomainServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDomain",
			Handler:    _DomainServer_GetDomain_Handler,
		},
		{
			MethodName: "GetDomains",
			Handler:    _DomainServer_GetDomains_Handler,
		},
		{
			MethodName: "GetUserCountInDomain",
			Handler:    _DomainServer_GetUserCountInDomain_Handler,
		},
		{
			MethodName: "GetUserInfoInDomain",
			Handler:    _DomainServer_GetUserInfoInDomain_Handler,
		},
		{
			MethodName: "UpdateUserRoleInDomain",
			Handler:    _DomainServer_UpdateUserRoleInDomain_Handler,
		},
		{
			MethodName: "AddUserInDomain",
			Handler:    _DomainServer_AddUserInDomain_Handler,
		},
		{
			MethodName: "DeleteUserInDomain",
			Handler:    _DomainServer_DeleteUserInDomain_Handler,
		},
		{
			MethodName: "EnterDomain",
			Handler:    _DomainServer_EnterDomain_Handler,
		},
		{
			MethodName: "AddUserTenantACL",
			Handler:    _DomainServer_AddUserTenantACL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/domain.proto",
}

func init() { proto.RegisterFile("api/domain.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xef, 0x6e, 0xd3, 0x3a,
	0x14, 0xbf, 0x59, 0x9a, 0xe9, 0xf6, 0x74, 0xda, 0x7a, 0xdd, 0xde, 0x2e, 0xcb, 0xba, 0x2d, 0xd7,
	0x93, 0xae, 0x2a, 0x81, 0x86, 0x34, 0x10, 0x30, 0x81, 0x84, 0xc6, 0x8a, 0xa6, 0x49, 0x68, 0x12,
	0xd9, 0x2a, 0x04, 0x88, 0x0f, 0x59, 0x63, 0x50, 0x58, 0x97, 0x84, 0xc4, 0x85, 0xc7, 0xe0, 0x5d,
	0xf8, 0xc8, 0x3b, 0xf1, 0x0e, 0x28, 0xb6, 0x63, 0x3b, 0x69, 0x32, 0xd8, 0x3e, 0x54, 0xb2, 0x7f,
	0xe7, 0xdf, 0xef, 0x1c, 0x9f, 0x73, 0x52, 0xe8, 0xfa, 0x49, 0x78, 0x2f, 0x88, 0xaf, 0xfc, 0x30,
	0xda, 0x4b, 0xd2, 0x98, 0xc6, 0xc8, 0xf4, 0x93, 0xd0, 0x61, 0x30, 0x25, 0x91, 0x1f, 0x51, 0x0e,
	0xe3, 0x1f, 0x06, 0x2c, 0x8f, 0x99, 0x1e, 0xea, 0x82, 0x19, 0x84, 0x81, 0x6d, 0xb8, 0xc6, 0xc8,
	0xf2, 0xf2, 0x23, 0xda, 0x06, 0xe0, 0x3e, 0x4e, 0xfd, 0x2b, 0x62, 0x2f, 0xb9, 0xc6, 0xa8, 0xed,
	0x69, 0x48, 0x2e, 0x3f, 0x4a, 0x89, 0x4f, 0xc9, 0x79, 0x78, 0x45, 0x6c, 0xd3, 0x35, 0x46, 0xa6,
	0xa7, 0x21, 0xb9, 0x7c, 0x92, 0x04, 0x85, 0xbc, 0xc5, 0xe5, 0x0a, 0x41, 0x43, 0x68, 0x73, 0x6f,
	0x93, 0x74, 0x66, 0x5b, 0xcc, 0xbd, 0x02, 0x90, 0x0b, 0x1d, 0x7e, 0x39, 0xa3, 0x3e, 0x25, 0xf6,
	0x32, 0xe3, 0xa5, 0x43, 0xf8, 0x21, 0x74, 0x8f, 0x09, 0xe5, 0xf4, 0x3d, 0xf2, 0x79, 0x4e, 0x32,
	0x5a, 0x93, 0x45, 0x17, 0xcc, 0x79, 0x18, 0x30, 0xfa, 0x96, 0x97, 0x1f, 0xf1, 0x63, 0xf8, 0x47,
	0xb3, 0xcb, 0x92, 0x38, 0xca, 0x08, 0xda, 0x2d, 0x0a, 0xc1, 0x6c, 0x3b, 0xfb, 0x9d, 0x3d, 0x3f,
	0x09, 0xf7, 0x84, 0x92, 0x10, 0xe1, 0x37, 0xd0, 0xe3, 0xfc, 0x6f, 0x1c, 0xb4, 0x52, 0x4c, 0xb3,
	0x5a, 0x4c, 0xfc, 0x04, 0xfa, 0x65, 0xd7, 0x37, 0xe1, 0x75, 0x00, 0xbd, 0x31, 0x99, 0x91, 0x5b,
	0xf0, 0xc2, 0x03, 0xe8, 0x97, 0x4d, 0x79, 0x5c, 0x7c, 0x08, 0x9b, 0xc7, 0x84, 0x4e, 0x32, 0x92,
	0x1e, 0xc5, 0xf3, 0x88, 0x9e, 0x44, 0x37, 0x77, 0xfd, 0x00, 0x86, 0xf5, 0x2e, 0x44, 0x6a, 0x7d,
	0xb0, 0xa6, 0xb9, 0x40, 0x78, 0xe1, 0x17, 0xfc, 0x48, 0x7b, 0x9d, 0x4c, 0x0b, 0x37, 0x57, 0xe1,
	0xe6, 0x3c, 0x5c, 0xa0, 0xc2, 0x05, 0x61, 0x80, 0x0f, 0x00, 0xe9, 0x86, 0x35, 0xf5, 0x33, 0x9b,
	0xea, 0xf7, 0x09, 0x1c, 0xc1, 0xf4, 0x24, 0xfa, 0x10, 0xdf, 0x22, 0x57, 0x84, 0xa0, 0x95, 0xf8,
	0x1f, 0xf9, 0xc3, 0x5a, 0x1e, 0x3b, 0xab, 0xfc, 0x5a, 0x7a, 0x7e, 0x97, 0xb2, 0xb0, 0xe5, 0x58,
	0x82, 0xef, 0x5d, 0x68, 0xcf, 0x33, 0x92, 0x7a, 0xf1, 0x8c, 0x64, 0x82, 0xf2, 0x2a, 0xa3, 0x3c,
	0x29, 0x50, 0x4f, 0x29, 0xe4, 0x5d, 0x45, 0x63, 0xea, 0xcf, 0x58, 0x81, 0x05, 0x1f, 0x0d, 0xc1,
	0xdf, 0x0c, 0xd8, 0xe2, 0x6d, 0x55, 0x98, 0xd7, 0x24, 0x57, 0xa9, 0x2c, 0x86, 0x95, 0xb9, 0x34,
	0x39, 0x19, 0x0b, 0xaf, 0x25, 0x0c, 0xfd, 0x0f, 0xab, 0xea, 0xee, 0x85, 0x41, 0x66, 0x9b, 0xae,
	0x39, 0xb2, 0xbc, 0x0a, 0x5a, 0x94, 0xae, 0xa5, 0x5e, 0xc9, 0x85, 0xed, 0x26, 0x42, 0xa2, 0xf3,
	0xa6, 0xb0, 0xc1, 0x3b, 0x92, 0xd7, 0xe8, 0x0f, 0xe8, 0x06, 0x4a, 0x5d, 0xd2, 0xd5, 0xb1, 0x82,
	0x86, 0xa9, 0x68, 0x0c, 0xc1, 0xa9, 0x0b, 0x22, 0x28, 0x7c, 0x37, 0x60, 0x70, 0x18, 0x04, 0x0d,
	0x04, 0x7e, 0xdb, 0x0c, 0x23, 0x58, 0xf3, 0xb9, 0x75, 0xfe, 0x8b, 0xd4, 0xc0, 0x57, 0x61, 0x4d,
	0xf3, 0x34, 0x9c, 0x5e, 0x32, 0xcd, 0x56, 0x49, 0xb3, 0x80, 0xf3, 0x75, 0x28, 0x20, 0x56, 0x6e,
	0x8b, 0x95, 0x5b, 0x87, 0xf0, 0x06, 0xac, 0x2f, 0x70, 0x16, 0xf9, 0x5c, 0x00, 0x7a, 0x11, 0x51,
	0x92, 0x96, 0x53, 0x71, 0xe0, 0xef, 0xc4, 0xcf, 0xb2, 0xaf, 0x71, 0xca, 0xf3, 0x69, 0x7b, 0xf2,
	0x9e, 0xcb, 0xa2, 0x82, 0x11, 0xdf, 0xfc, 0xf2, 0x9e, 0xf7, 0x35, 0x8d, 0x2f, 0x49, 0x24, 0x92,
	0xe2, 0x17, 0x7c, 0x07, 0x7a, 0xa5, 0x18, 0x6a, 0xc8, 0xb9, 0xb2, 0xa1, 0x2b, 0xbf, 0x93, 0x5c,
	0xcf, 0xd9, 0xe7, 0xe8, 0xf0, 0xe8, 0xa5, 0xc6, 0x4a, 0x56, 0x4d, 0xb0, 0x92, 0xe5, 0x42, 0xd0,
	0x3a, 0xcf, 0xb3, 0x5f, 0x62, 0xd9, 0xb3, 0x73, 0x5e, 0xfe, 0xb1, 0x7a, 0xdb, 0x71, 0x18, 0x60,
	0x07, 0xec, 0x45, 0xe7, 0x9c, 0xce, 0xfe, 0x4f, 0x0b, 0x56, 0x38, 0xc3, 0x33, 0x92, 0x7e, 0x21,
	0x29, 0x7a, 0x0a, 0x6d, 0xb9, 0x35, 0xd0, 0xbf, 0x6c, 0xd2, 0xaa, 0x1f, 0x15, 0x67, 0x50, 0x85,
	0x45, 0x59, 0xff, 0x42, 0xcf, 0x00, 0xd4, 0xce, 0x41, 0x15, 0xbd, 0x62, 0x7b, 0x39, 0xeb, 0x0b,
	0xb8, 0x74, 0xf0, 0x1e, 0xfa, 0x75, 0x3b, 0x12, 0xb9, 0x85, 0x49, 0xd3, 0x06, 0x76, 0xfe, 0xbb,
	0x46, 0x43, 0xba, 0x7f, 0x0b, 0xbd, 0x9a, 0x65, 0x83, 0x76, 0x74, 0xdb, 0x9a, 0x95, 0xe7, 0xb8,
	0xcd, 0x0a, 0xd2, 0xf7, 0x14, 0x06, 0xf5, 0x93, 0x8c, 0x30, 0x5f, 0x58, 0xd7, 0xed, 0x1d, 0x67,
	0xf7, 0x5a, 0x1d, 0x19, 0xe4, 0x14, 0xd6, 0x2a, 0x4d, 0x8d, 0x36, 0x99, 0x65, 0xfd, 0x78, 0x3a,
	0xc3, 0x7a, 0xa1, 0xf4, 0xf7, 0x1a, 0xd0, 0xe2, 0xdc, 0xa3, 0x6d, 0xfe, 0x51, 0x68, 0xda, 0x3a,
	0xce, 0x4e, 0xa3, 0x5c, 0x3a, 0x7e, 0x0e, 0x1d, 0xad, 0xfd, 0x11, 0x7f, 0xf2, 0xc5, 0xa1, 0x73,
	0xec, 0x45, 0x81, 0xf4, 0xf1, 0x0a, 0xba, 0xd5, 0xc6, 0x45, 0xa5, 0x84, 0xaa, 0xc3, 0xe2, 0x6c,
	0x35, 0x48, 0x0b, 0x97, 0x17, 0xcb, 0xec, 0x7f, 0xde, 0xfd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x7c, 0xa1, 0x13, 0x12, 0x0a, 0x00, 0x00,
}
