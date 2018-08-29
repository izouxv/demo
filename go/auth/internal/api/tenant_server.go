package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"

	. "auth/common"
	"auth/go-drbac/common"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
	"github.com/gin-gonic/gin/json"
)

/*AuthServer rpc服务*/
type TenantServer struct {
	DrbacServer *drbac.DrbacServer
}

/*创建租户*/
func (this *TenantServer) AddTenant(ctx context.Context, in *pb.AddTenantRequest) (*pb.AddTenantResponse, error) {
	log.Info("Start Func AddTenant")
	log.Info("AddTenant Request Info:",in)
	if in.TenantName == "" || in.TenantURL == "" || in.Email == "" {
		log.Info("Request Info is Empty")
		return nil, InvalidArgument
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	if in.TenantName == "" || in.Did == 0 {
		log.Error("Func AddTenant Request is empty, in.TenantName:(&s), in.Did:(&d)", in.TenantName, in.Did)
		return nil, InvalidArgument
	}

	if in.Pid != 0 {
		depth, err := this.DrbacServer.GetTenantDepth(in.Pid)
		if err != nil {
			log.Error("GetDomainDepth Error, ", err)
			return nil, NotFind
		}
		if depth > MaxDepth {
			log.Errorf("深度超过最大深度")
			return nil, PermissionDenied
		}
	}
	//创建租户
	tenantInfo := drbac.TenantInfo{
		TenantName:  in.TenantName,
		TenantURL:   in.TenantURL,
		Description: in.Description,
		Email:       in.Email,
		Phone:       in.Phone,
		Contacts:    in.Contacts,
		Icon:		 in.Icon,
		Logo:		 in.Logo,
	}
	_, err := this.DrbacServer.CreateTenantBaseTenant(in.Did, in.Pid, tenantInfo, token, userTokenTenants.User.Nickname)
	if err != nil {
		log.Errorf("create Domain Pid is %d,基于父租户创建租户 error is : %s \n ,domainName is %s,Oid is %d", in.Pid, err, in.TenantName)
		return nil, SystemError
	}
	return nil, Successful
}

/*再次邀请租户*/
func (this *TenantServer) InviteUnactivatedTenant(ctx context.Context, in *pb.InviteUnactivatedTenantRequest) (*pb.InviteUnactivatedTenantResponse, error) {
	log.Info("------------InviteUnactivatedTenant----------------")
	log.Info("InviteUnactivatedTenant Request Info: ",&in)

	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants.User:", userTokenTenants.User)
	log.Info("UserTokenTenants:", userTokenTenants.TenantRoleResource)
	if &userTokenTenants == nil {
		return nil, SystemError
	}

	if in.Tid == 0 {
		log.Info("InviteUnactivatedTenant Input Is Empty")
		return nil, InvalidArgument
	}
	//todo 获取租户信息
	tenant,err := this.DrbacServer.GetTenant(in.Tid)
	if err != nil || tenant.TenantName == ""{
		log.Info("GetTenant Error, ",err)
		return nil, TidIsIncorrectOrEmpty
	}
	if tenant.TenantState != 1 {
		log.Info("Tid:",in.Tid,",租户已激活，无需再次邀请")
		return nil, TenantAlreadyActivated
	}
	user, err := this.DrbacServer.GetUserByUsernameAndTid(in.Username, in.Tid)
	if err != nil || user.Uid == 0{
		log.Info("GetUserByUsernameAndTid Error, ",err)
		return nil,UserDoesNotExist
	}
	//todo 发送邀请邮件
	sendData := SendData{
		To:       user.Username,
		Nickname: userTokenTenants.User.Nickname,
		Token:    user.Token,
		Url:      tenant.TenantURL,
	}
	mailInfo := MailInfo{
		TemId:     3,
		EmailAddr: user.Username,
		SendData:  sendData,
	}
	req, err := json.Marshal(mailInfo)
	if err != nil {
		log.Info("json.Marshal(mailInfo) Error,", err)
		return nil, err
	}
	res := HttpSendMailPost(req)
	log.Info("HttpSendMailPost:", res)
	return nil, Successful
}



/*UpdateDomain
修改租户
*/
func (this *TenantServer) UpdateTenant(ctx context.Context, in *pb.UpdateTenantRequest) (*pb.UpdateTenantResponse, error) {
	log.Info("-------UpdateDomain----------")
	if in.TenantInfo.Tid == 0 {
		log.Errorf("input error,input is: %v#", in)
		return nil, InvalidArgument
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userToken := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserToken:", userToken)
	if &userToken == nil {
		return nil, TokenIsInvalid
	}
	start := time.Now()
	tenantInfo := drbac.TenantInfo{
		Tid:         in.TenantInfo.Tid,
		TenantName:  in.TenantInfo.TenantName,
		TenantURL:   in.TenantInfo.TenantURL,
		Description: in.TenantInfo.Description,
		Contacts:    in.TenantInfo.Contacts,
		Email:       in.TenantInfo.Email,
		Phone:       in.TenantInfo.Phone,
		Icon:		 in.TenantInfo.Icon,
		Logo:		 in.TenantInfo.Logo,
	}
	_, err := this.DrbacServer.UpdateTenantInfo(tenantInfo, token)
	log.Info("调用时间：", time.Now().Sub(start))
	if err != nil {
		log.Errorf("UpdateTenantInfo Error,", err)
		return nil, SystemError
	}
	return nil, Successful
}

/*UpdateDomain
修改租户状态
*/
func (this *TenantServer) UpdateTenantState(ctx context.Context, in *pb.UpdateTenantStateRequest) (*pb.UpdateTenantStateResponse, error) {
	log.Info("-------UpdateTenantState----------")
	if in.Tid == 0 || in.State == 0{
		log.Errorf("input error,input is: %v#", in)
		return nil, InvalidArgument
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userToken := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserToken:", userToken)
	if &userToken == nil {
		return nil, TokenIsInvalid
	}
	start := time.Now()

	_, err := this.DrbacServer.UpdateTenantState(in.Tid, in.State, token)
	log.Info("调用时间：", time.Now().Sub(start))
	if err != nil {
		log.Errorf("UpdateTenantInfo Error,", err)
		return nil, SystemError
	}
	return nil, Successful
}


/*GetTenant
获取租户信息
*/
func (this *TenantServer) GetTenant(ctx context.Context, in *pb.GetTenantRequest) (*pb.GetTenantResponse, error) {
	if in.Tid == 0 {
		log.Errorf("input error,input is: %v#", in)
		return nil, status.Errorf(codes.InvalidArgument, "%v")
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	isTrue := false
	for _, v := range userTokenTenants.TenantRoleResource {
		if v.Tenant.Tid == in.Tid {
			isTrue = true
			break
		}
	}
	if isTrue {
		domain, err := this.DrbacServer.GetTenant(in.Tid)
		if err != nil {
			log.Errorf("系统异常 GetDomains is error %s,uid is %d,did is %d", err, in.Tid)
			return nil, status.Errorf(codes.Unknown, "%v")
		}
		if err != nil {
			log.Errorf("系统异常，error is %s", err)
		}
		userCount, err := this.DrbacServer.GetUserCountBaseTenant(domain.Tid)
		if err != nil {
			log.Errorf("系统异常，error is %s", err)
		}
		domainResp := &pb.Tenant{
			Tid:          domain.Tid,
			TenantName:   domain.TenantName,
			Pid:          domain.Pid,
			CreateTime:   domain.CreateTime.Unix(),
			UpdateTime:   domain.UpdateTime.Unix(),
			TenantExtend: &pb.TenantExtend{UserCount: userCount},

		}
		log.Info("获取域成功%#v", domainResp)
		return &pb.GetTenantResponse{Tenant: domainResp}, nil
	} else {
		log.Info("用户不存在改域中 uid is %d, did is %d", in.Tid)
		return nil, status.Errorf(codes.NotFound, "%v")
	}
}

/*GetTenants
获取域下租户列表
*/
func (this *TenantServer) GetTenants(ctx context.Context, in *pb.GetTenantsRequest) (*pb.GetTenantsResponse, error) {
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
		did   int32
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["did"]; ok && val[0] != "" {
		didInt, _ := strconv.Atoi(val[0])
		did = int32(didInt)
	}
	log.Info("did:", did)
	if did == 0 || token == "" {
		log.Error("did or token is empty")
		return nil, InvalidArgument
	}
	log.Info("token:", token)
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants.User)
	log.Info("UserTokenTenants:", userTokenTenants.DomainRoleResource)
	if userTokenTenants == nil {
		return nil, SystemError
	}
	if userTokenTenants.DomainRoleResource == nil {
		return nil, PermissionDenied
	}
	log.Info("UserTokenTenants.domain:", userTokenTenants.DomainRoleResource.Domain)
	if userTokenTenants.DomainRoleResource.Domain.Did == did {
		tenants, err := this.DrbacServer.GetTenants(did,userTokenTenants.User.Uid)
		if err != nil {
			log.Errorf("系统异常 GetDomains is error %s", err)
			return nil, SystemError
		}
		var resp []*pb.TenantInfo
		for _, v := range tenants {
			tenant := &pb.TenantInfo{
				Tid:         v.Tid,
				TenantName:  v.TenantName,
				TenantURL:   v.TenantURL,
				Description: v.Description,
				Contacts:    v.Contacts,
				Email:       v.Email,
				Phone:       v.Phone,
				State:       v.TenantState,
				CreateTime:  v.CreateTime.Unix(),
				Icon:		 v.Icon,
				Logo:		 v.Logo,
			}
			resp = append(resp, tenant)
		}

		log.Info("获取域成功%#v", resp)
		return &pb.GetTenantsResponse{Tenants: resp}, nil
	} else {
		log.Info("用户不存在改域中 uid is %d, did is %d", userTokenTenants.User.Uid, did)
		return nil, PermissionDenied
	}
}

/*DeleteDomain
删除租户
*/
func (this *TenantServer) DeleteTenant(ctx context.Context, in *pb.DeleteTenantRequest) (*pb.DeleteTenantResponse, error) {
	log.Info("---------DeleteTenant------------")
	if in.Tid == 0 {
		log.Errorf("input error,input is: %v#", in)
		return nil, InvalidArgument
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}

	start := time.Now()
	err := this.DrbacServer.DeleteTenant(in.Tid, token)
	log.Info("调用时间：", time.Now().Sub(start))
	if err != nil {
		if err == common.ErrCanNotDelete {
			log.Errorf("无法删除域，其下有资源")
			return nil, CanNotDeleteTenant
		}
		log.Errorf("系统异常 error is %s", err)
		return nil, SystemError
	}
	log.Infof("删除域成功,did is %d", in.Tid)
	return nil,Successful

}

/*GetUserCountInDomain
获取域中成员数量
*/
func (this *TenantServer) GetUserCountInTenant(ctx context.Context, in *pb.GetUserCountInTenantRequest) (*pb.GetUserCountInTenantResponse, error) {
	log.Info("------------GetUserCountInDomain------------------")
	if in.Tid == 0 || in.Uid == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v")
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	isTrue := false
	for _, v := range userTokenTenants.TenantRoleResource {
		if v.Tenant.Tid == in.Tid {
			isTrue = true
			break
		}
	}
	if isTrue {
		start := time.Now()
		count, err := this.DrbacServer.GetUserCountBaseTenant(in.Tid)
		log.Info("调用时间：", time.Now().Sub(start))
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "%v")
		}
		return &pb.GetUserCountInTenantResponse{Count: count}, nil
	} else {
		log.Info("用户不存在改域中 uid is %d, did is %d", in.Uid, in.Tid)
		return nil, status.Errorf(codes.NotFound, "%v")
	}
}

/*GetUserInfoInDomain
获取租户中成员
*/
func (this *TenantServer) GetUserInfoInTenant(ctx context.Context, in *pb.GetUserInfoInTenantRequest) (*pb.GetUserInfoInTenantResponse, error) {
	log.Info("------------GetUserInfoInTenant------------------")
	log.Info("GetUserInfoInTenant Input Info, page:", in.Page, ", count:", in.Count)
	if in.Page == 0 {
		log.Info("GetUserInfoInTenant Input is Empty")
		return nil,InvalidArgument
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
		tid   int32
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["tid"]; ok && val[0] != "" {
		tidInt, _ := strconv.Atoi(val[0])
		tid = int32(tidInt)
	}
	log.Info("token:", token)
	log.Info("tid:", tid)
	if tid == 0 || token == "" {
		log.Error("tid or token is empty")
		return nil, InvalidArgument
	}
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	isTrue := true
	//for _, v := range userTokenTenants.TenantRolePermissions {
	//	if v.Tenant.Tid == tid {
	//		isTrue = true
	//		break
	//	}
	//}

	if isTrue {
		start := time.Now()
		userRoles,totalCount, err := this.DrbacServer.GetUserRoles(tid, in.Page, in.Count)
		log.Info("调用时间：", time.Now().Sub(start))
		if err != nil {
			return nil, InvalidArgument
		}
		var reply []*pb.UserRoles
		for _, v := range userRoles {
			var roles []*pb.Role
			for _, v2 := range v.Roles {
				roles = append(roles, &pb.Role{Rid: v2.Rid, RoleName: v2.RoleName})
			}
			reply = append(reply, &pb.UserRoles{
				User: &pb.User{
					Uid:        v.User.Uid,
					Username:   v.User.Username,
					Nickname:   v.User.Nickname,
					State:      v.User.State,
					CreateTime: v.User.CreateTime.Unix()},
				Roles: roles,
			})
		}
		return &pb.GetUserInfoInTenantResponse{UserRoles: reply,TotalCount:totalCount}, nil
	} else {
		log.Info("用户不存在该域中 uid is %d, tid is %d", userTokenTenants.User.Uid, tid)
		return nil, UserDoesNotExist
	}
}

/*UpdateUserRoleInTenant
修改租户中成员角色
*/
func (this *TenantServer) UpdateUserRoleInTenant(ctx context.Context, in *pb.UpdateUserRoleInTenantRequest) (*pb.UpdateUserRoleInTenantResponse, error) {
	log.Info("------------UpdateUserRoleInDomain-----------------")
	if in.UpdateUserID == 0 {
		return nil, InvalidArgument
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
		tid   int32
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["tid"]; ok && val[0] != "" {
		tidInt, _ := strconv.Atoi(val[0])
		tid = int32(tidInt)
	}
	log.Info("token:", token)
	log.Info("tid:", tid)
	if tid == 0 || token == "" {
		log.Error("tid or token is empty")
		return nil, InvalidArgument
	}
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	isTrue := true
	//for _, v := range userTokenTenants.TenantRolePermissions {
	//	if v.Tenant.Tid == tid {
	//		isTrue = true
	//		break
	//	}
	//}
	if isTrue {
		if userTokenTenants.User.Uid == in.UpdateUserID {
			log.Errorf("不能修改自己 uid is %d,deleteUserId is %d", userTokenTenants.User.Uid, in.UpdateUserID)
			return nil, PermissionDenied
		}
		user, err := this.DrbacServer.GetUserByUid(in.UpdateUserID)
		if err != nil || user.Username == "" {
			log.Errorf("修改用户不存在,", err)
			return nil, UserDoesNotExist
		}
		start := time.Now()
		if err := this.DrbacServer.UpdateUserRole(in.UpdateUserID, tid, in.UpdateUserRids, token); err != nil {

			log.Errorf("UpdateUserRole is system error %s,uid is %d,updateUserId is %d,updateUserRoleId is %s", err, userTokenTenants.User.Uid, in.UpdateUserID, in.UpdateUserRids)
			return nil, SystemError
		}
		log.Info("调用时间：", time.Now().Sub(start))
		return nil, Successful
	} else {
		log.Errorf("UpdateUserRole is error Domain_Does_Not_Exist ,uid is %d,updateUserId is %d,updateUserRoleId is %s", userTokenTenants.User.Uid, in.UpdateUserID, in.UpdateUserRids)
		return nil, PermissionDenied
	}
}

/*DeleteUserInDomain
删除租户中成员
*/
func (this *TenantServer) DeleteUserInTenant(ctx context.Context, in *pb.DeleteUserInTenantRequest) (*pb.DeleteUserInTenantResponse, error) {
	if in.DeleteUserID == 0 {
		log.Errorf("输入参数异常，%#v", in)
		return nil, InvalidArgument
	}

	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
		tid   int32
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["tid"]; ok && val[0] != "" {
		tidInt, _ := strconv.Atoi(val[0])
		tid = int32(tidInt)
	}
	log.Info("token:", token)
	log.Info("tid:", tid)
	if tid == 0 || token == "" {
		log.Error("tid or token is empty")
		return nil, InvalidArgument
	}
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}

	if userTokenTenants.User.Uid == in.DeleteUserID {
		log.Errorf("不能删除自己 uid is %d,deleteUserId is %d", userTokenTenants.User.Uid, in.DeleteUserID)
		return nil, PermissionDenied
	}
	isTrue := true
	//for _, v := range userTokenTenants.TenantRolePermissions {
	//	if v.Tenant.Tid == tid {
	//		isTrue = true
	//		break
	//	}
	//}
	if isTrue {
		start := time.Now()
		err := this.DrbacServer.DeleteUser(in.DeleteUserID, tid)
		log.Info("调用时间：", time.Now().Sub(start))
		if err != nil {
			if err == common.ErrDoesNotExist {
				log.Errorf("删除用户不存在")
				return nil, UserDoesNotExist
			}
			log.Errorf("系统异常 error is %s", err)
			return nil, SystemError
		}
		log.Info("删除域中用户成功")
		return nil, Successful
	} else {
		log.Info("用户不存在该域中 uid is %d, did is %d", userTokenTenants.User.Uid)
		return nil, PermissionDenied
	}
}

/*AddUserInTenant
邀请用户
*/
func (this *TenantServer) AddUserInTenant(ctx context.Context, in *pb.AddUserInTenantRequest) (*pb.AddUserInTenantResponse, error) {
	log.Info("------------AddUserInDomain-----------------")
	if in.Tid == 0 || in.AddUserUsername == "" || in.AddUserNickname == "" {
		log.Errorf("len(in.AddUserRids):", len(in.AddUserRids))
		log.Errorf("input error tid is %d,AddUserRids is %d,AddUserUsername is %s,AddUserNickname is %s", in.Tid, in.AddUserRids, in.AddUserUsername, in.AddUserNickname)
		return nil, InvalidArgument
	}

	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val   []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:", token)
	userTokenTenants := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants.User:", userTokenTenants.User)
	log.Info("UserTokenTenants:", userTokenTenants.TenantRoleResource)
	if &userTokenTenants == nil {
		return nil, SystemError
	}

	//todo 获取租户信息
	log.Info("in.tid:",in.Tid)
	tenant, err := this.DrbacServer.GetTenant(in.Tid)
	if tenant.TenantURL == "" || err != nil {
		log.Error("GetTenantByTid Err, ", err)
		return nil, TidIsIncorrectOrEmpty
	}
	log.Info("tenant.TenantURL:",tenant.TenantURL)

	//todo 判断角色是否存在
	for _, rid := range in.AddUserRids {
		if _, err := this.DrbacServer.GetRoleByRid(rid); err != nil {
			log.Errorf("角色不存在")
			return nil, InvalidArgument
		}
	}
	//todo 判断用户操作权限
	//isTrue := false
	//for _, v := range userTokenTenants.TenantRolePermissions {
	//	if v.Tenant.Tid == in.Tid {
	//		isTrue = true
	//		break
	//	}
	//}
	//if !isTrue {
	//	log.Info("用户不存在该租户中 uid is (%d) , did is (%d)", userTokenTenants.User.Uid, in.Tid)
	//	return nil, PermissionDenied
	//}
	start := time.Now()
	yourself, err := this.DrbacServer.GetUserByUid(userTokenTenants.User.Uid)
	log.Info("调用时间：", time.Now().Sub(start))
	if err != nil {
		log.Errorf("系统异常：error is %s", err)
		return nil, SystemError
	}
	user, err := this.DrbacServer.GetUserByUsernameAndTid(in.AddUserUsername,in.Tid)
	log.Infof("uid is (%s),修改角色uid is (%s)", userTokenTenants.User.Uid, user.Uid)

	//todo 准备邮件信息
	sendData := SendData{
		To:       user.Username,
		Nickname: yourself.Nickname,
		Token:    user.Token,
		Url:      tenant.TenantURL,
	}
	mailInfo := MailInfo{
		TemId:     3,
		EmailAddr: user.Username,
		SendData:  sendData,
	}
	req, err := json.Marshal(mailInfo)
	if err != nil {
		log.Info("json.Marshal(mailInfo) Error,", err)
		return nil, err
	}

	//todo 判断用户是否已存在
	if user.Uid != 0 {
		//todo 判断是否已在邀请的租户中 1是 重新邀请  2不是，插入utr关系 邀请
		isTrue := this.DrbacServer.IsExistedInTenant(user.Uid, in.Tid)
		if isTrue {
			if user.State == 1 {
				res := HttpSendMailPost(req)
				log.Info("HttpSendMailPost:", res)
				return nil, Successful
			} else {
				if userTokenTenants.User.Uid == user.Uid {
					log.Errorf("不能添加自己 uid is %d,AddUserUsername is %s", in.AddUserUsername)
					return nil, UserAlreadyExist
				} else {
					log.Errorf("用户已存在 AddUserUsername is %s", in.AddUserUsername)
					return nil, UserAlreadyExist
				}
			}
		} else {
			if err := this.DrbacServer.AddUdr(user.Uid, in.Tid, in.AddUserRids); err != nil {
				res := HttpSendMailPost(req)
				log.Info("HttpSendMailPost:", res)
				return nil, Successful
			}
		}
	}
	//todo 被邀请人信息直接存数据库
	emailToken := string(Krand(32, KC_RAND_KIND_ALL))
	if err := this.DrbacServer.AddUser(in.AddUserUsername, in.AddUserNickname, emailToken, in.Tid, in.AddUserRids); err != nil {
		log.Error("AddUser Error, ", err)
		return nil, SystemError
	}
	//todo 发送邀请邮件
	sendData = SendData{
		To:       in.AddUserUsername,
		Nickname: yourself.Nickname,
		Token:    emailToken,
		Url:      tenant.TenantURL,
	}
	mailInfo = MailInfo{
		TemId:     3,
		EmailAddr: in.AddUserUsername,
		SendData:  sendData,
	}
	req, err = json.Marshal(mailInfo)
	if err != nil {
		log.Info("json.Marshal(mailInfo) Error,", err)
		return nil, err
	}
	res := HttpSendMailPost(req)
	log.Info("HttpSendMailPost:", res)
	return nil, Successful
}

/*EnterDomain
用户同意进入租户中
*/
func (this *TenantServer) EnterTenant(ctx context.Context, in *pb.EnterTenantRequest) (*pb.EnterTenantResponse, error) {
	log.Info("------------EnterDomain----------------")
	if in.Token == "" || in.Password == "" {
		log.Error("Input InvalidArgument, in.Token:", in.Token, " in.Password: ", in.Password)
		return nil, InvalidArgument
	}
	//todo 根据token查询邀请是否存在
	user, err := this.DrbacServer.GetUserByToken(in.Token)
	if err != nil || user.Uid == 0 {
		log.Error("GetUserByToken Error,", err)
		return nil, UserDoesNotExist
	}
	if user.Tid == 0 {
		log.Info("用户不在租户中")
		return nil,UserDoesNotExist
	}
	log.Info("userinfo:", user)
	//todo 判断账号是否可用
	if user.State != 1 {
		log.Infof("账号状态异常，state：", user.State)
		switch user.State {
		case 2:
			return nil, AccountDisableToUse
		default:
			return nil, AccountException
		}
	}
	//todo 完善用户信息
	ra, err := this.DrbacServer.UpdateUser(user.Uid, in.Password)
	if err != nil || ra == 0 {
		log.Error("UpdateUser Error,", err, ", RowsAffected:", ra)
		return nil, SystemError
	}
	//todo 修改租户状态
	tenant,err := this.DrbacServer.GetTenant(user.Tid)
	if err != nil || tenant.TenantName == "" {
		log.Info("GetTenant Error,",err)
		return nil, TidIsIncorrectOrEmpty
	}
	if tenant.TenantState == 1 {
		_ ,err = this.DrbacServer.UpdateTenantState(user.Tid,3,"")
		if err != nil {
			log.Info("UpdateTenantState Error,",err)
			return nil, TidIsIncorrectOrEmpty
		}
	}
	return nil, Successful
}



func (this *TenantServer) GetTidByUrl(ctx context.Context, in *pb.GetTidByUrlRequest) (*pb.GetTidByUrlResponse, error) {
	if in.Url == "" {
		log.Error("GetTidByUrl in.url is empty, in.Url:",in.Url)
		return nil,InvalidArgument
	}
	tenant, err := this.DrbacServer.GetTidByUrl(in.Url)
	if err != nil {
		log.Error("GetTidByUrl Error,",err)
		return nil,URLDoesNotExist
	}
	return &pb.GetTidByUrlResponse{Tid:tenant.Tid, Icon:tenant.Icon, Logo:tenant.Logo},nil
}

func (this *TenantServer) GetDidByTid(ctx context.Context, in *pb.GetDidByTidRequest) (*pb.GetDidByTidResponse, error) {
	if in.Tid == 0 {
		log.Error("GetDidByTid in.Tid is empty, in.Tid:",in.Tid)
		return nil,InvalidArgument
	}
	did, err := this.DrbacServer.GetDidByTid(in.Tid)
	if err != nil || did == 0{
		log.Error("GetTidByUrl Error,",err)
		return nil,SystemError
	}
	return &pb.GetDidByTidResponse{Did:did},nil
}

