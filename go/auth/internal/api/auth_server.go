package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"

	"auth/storage"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"regexp"
	"strconv"
	"time"
)

/*AuthServer rpc服务*/
type AuthServer struct {
	DrbacServer *drbac.DrbacServer
}

//权限校验
func (this *AuthServer) Authorization(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	log.Info("------------Start Authorization----------------")
	log.Info("in:", in)
	//todo 校验参数
	if in.Tid == 0 || in.Token == "" || in.Url == "" || in.Opt == "" {
		log.Error("Authorization Input InvalidArgument, in:", in)
		return nil, InvalidArgument
	}
	userToken := this.DrbacServer.GetAuthorizationInfo(in.Token)
	if userToken == nil {
		log.Errorf("用户未登录，token失效 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	log.Info("in.Tid:", in.Tid)
	log.Info("userToken.User.Tid :", userToken.User.Tid)
	if userToken.User.Tid != in.Tid {
		log.Info("用户不在该租户中，权限不足")
		return nil, PermissionDenied
	}
	//log.Info("---校验成功---")
	//return nil,Successful
	if in.Tid == 100002 || in.Tid == 100003 || in.Tid == 100004 {
		//todo radacat后台鉴权
		if userToken.User.Username == "" || userToken.DomainRoleResource.Domain.Did == 0 || len(userToken.TenantRoleResource) == 0 {
			log.Errorf("用户未登录，userToken信息为空 token is %s", in.Token)
			return nil, TokenIsInvalid
		}
		//todo 比较tid
		for _, v := range userToken.TenantRoleResource {
			log.Info("v.Tenant.Tid:", v.Tenant.Tid)
			if v.Tenant.Tid == in.Tid {
				//todo 校验白名单
				if this.DrbacServer.CheckWhitelist(in.Url, in.Opt) {
					return nil, Successful
				}
				//todo 校验url权限
				log.Info("url不在白名单，开始校验")
				for _, v := range userToken.DomainRoleResource.Resource {
					if v.ResUrl != "" {
						log.Info("permissions URL:", v.ResUrl, " OPT:", v.ResOpt)
						match, _ := regexp.MatchString(v.ResUrl, in.Url)
						if match {
							log.Info("URL匹配成功")
						}
						if match && v.ResOpt == in.Opt {
							log.Info("OPT匹配成功")
							log.Info("url校验通过")
							return nil, Successful
						}
					}
				}
				log.Info("url校验未通过，权限不足")
				return nil, PermissionDenied
			}
		}
		log.Info("userToken中无租户信息，鉴权失败")
		return nil, TokenIsInvalid
	}
	//todo 正常流程
	//todo token获取用户租户信息
	if userToken.User.Username == "" || len(userToken.TenantRoleResource) == 0 {
		log.Errorf("用户未登录，userToken信息为空 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	//todo 比较tid
	log.Info("userToken.TenantRolePermissions[0].Tenant.Tid:", userToken.TenantRoleResource[0].Tenant.Tid)
	if userToken.TenantRoleResource[0].Tenant.Tid != in.Tid {
		log.Info("tid与该账号所在租户不匹配")
		return nil, PermissionDenied
	}
	//todo 校验白名单
	if this.DrbacServer.CheckWhitelist(in.Url, in.Opt) {
		return nil, Successful
	}
	//todo 校验url权限
	log.Info("url不在白名单，开始校验")
	for _, v := range userToken.TenantRoleResource[0].Resource {
		if v.ResUrl != "" {
			log.Info("permissions URL:", v.ResUrl, " OPT:", v.ResOpt)
			match, _ := regexp.MatchString(v.ResUrl, in.Url)
			if match {
				log.Info("URL匹配成功")
			}
			if match && v.ResOpt == in.Opt {
				log.Info("OPT匹配成功")
				log.Info("url校验通过")
				log.Info("Username:", userToken.User.Username)
				return &pb.AuthorizationResponse{Username: userToken.User.Username}, nil
				//return nil,Successful
			}
		}
	}
	log.Info("url校验未通过，权限不足")
	return nil, PermissionDenied
}

func (this *AuthServer) AuthorizationWithDid(ctx context.Context, in *pb.AuthorizationWithDidRequest) (*pb.AuthorizationWithDidResponse, error) {
	log.Info("------------Start Authorization----------------")
	log.Info("in:", in)
	//todo 校验参数
	if in.Did == 0 || in.Token == "" || in.Url == "" || in.Opt == "" {
		log.Error("Authorization Input InvalidArgument, in:", in)
		return nil, InvalidArgument
	}
	userToken := this.DrbacServer.GetAuthorizationInfo(in.Token)
	log.Info("in.Did:", in.Did)
	if userToken == nil {
		log.Errorf("用户未登录，token失效 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	log.Info("userToken.User.Did :", userToken.User.Did)
	if userToken.User.Did != in.Did {
		log.Info("用户不在该域中，权限不足")
		return nil, PermissionDenied
	}
	//log.Info("---校验成功---")
	//return nil,Successful

	//todo 正常流程
	//todo token获取用户租户信息
	if userToken.User.Username == "" || userToken.DomainRoleResource == nil {
		log.Errorf("用户未登录，userToken信息为空 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	//todo 比较did
	log.Info("userToken.DomainRoleResource.Domain.Did:", userToken.DomainRoleResource.Domain.Did)
	if userToken.DomainRoleResource.Domain.Did != in.Did {
		log.Info("Did与该账号所在租户不匹配")
		return nil, PermissionDenied
	}
	//todo 校验白名单
	if this.DrbacServer.CheckWhitelist(in.Url, in.Opt) {
		return nil, Successful
	}
	//todo 校验url权限
	log.Info("url不在白名单，开始校验")
	for _, v := range userToken.DomainRoleResource.Resource {
		if v.ResUrl != "" {
			log.Info("permissions URL:", v.ResUrl, " OPT:", v.ResOpt)
			match, _ := regexp.MatchString(v.ResUrl, in.Url)
			if match {
				log.Info("URL匹配成功")
			}
			if match && v.ResOpt == in.Opt {
				log.Info("OPT匹配成功")
				log.Info("url校验通过")
				return nil, Successful
			}
		}
	}
	log.Info("url校验未通过，权限不足")
	return nil, PermissionDenied
}

//登录指定租户
func (this *AuthServer) AuthenticationWithTid(ctx context.Context, in *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	log.Info("------------Authentication----------------")
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val       []string
		tidString string
		tid       int32
	)
	if val, ok = md["tid"]; ok {
		tidString = val[0]
	}
	if tidString != "" {
		tidInt, _ := strconv.Atoi(tidString)
		tid = int32(tidInt)
	}
	log.Info("tid:", tid)
	tenatInfo, err := this.DrbacServer.GetTenant(tid)
	if err != nil || tenatInfo.TenantState != 3 {
		log.Error("租户异常或已停用，ERR：", err, "TenantState:", tenatInfo.TenantState)
		return nil, TenantSystemError
	}
	if in.Username == "" || in.Password == "" {
		log.Errorf("input parameter error,username is (%s),password is (%s)", in.Username, in.Password)
		return nil, InvalidArgument
	}
	userTokenTenants, isAuth, state, err := this.DrbacServer.AuthenticationWithTid(in.Username, in.Password, tid)
	if err != nil {
		log.Errorf("用户不存在 username is %s", in.Username)
		return nil, UserDoesNotExist
	}
	if state != 3 {
		log.Infof("账号状态异常，state：", state)
		switch state {
		case 1:
			return nil, AccountNotActive
		case 2:
			return nil, AccountDisableToUse
		default:
			return nil, AccountException
		}
	}
	if !isAuth {
		log.Errorf("用户密码错误 username is %s,password is %s", in.Username, in.Password)
		return nil, UsernameAndPasswordError
	}
	ut := handlerUserTokenTenants(userTokenTenants)
	log.Infof("认证成功 username is %#v", ut)
	return &pb.AuthenticationResponse{UserToken: ut}, nil
}

//登录指定域
func (this *AuthServer) AuthenticationWithDid(ctx context.Context, in *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	log.Info("------------AuthenticationWithDid----------------")
	if in.Username == "" || in.Password == "" {
		log.Errorf("input parameter error,username is (%s),password is (%s)", in.Username, in.Password)
		return nil, InvalidArgument
	}
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, SystemError
	}
	var (
		val []string
		did int32
	)
	if val, ok = md["did"]; ok && val[0] != "" {
		didInt, _ := strconv.Atoi(val[0])
		did = int32(didInt)
	}
	log.Info("did:", did)
	if did == 0 {
		log.Error("did or token is empty")
		return nil, InvalidArgument
	}
	userTokenDomain, isAuth, state, err := this.DrbacServer.AuthenticationWithDid(in.Username, in.Password, did)
	if err != nil {
		log.Errorf("用户不存在 username is %s", in.Username)
		return nil, UserDoesNotExist
	}
	if state != 3 {
		log.Infof("账号状态异常，state：", state)
		switch state {
		case 1:
			return nil, AccountNotActive
		case 2:
			return nil, AccountDisableToUse
		default:
			return nil, AccountException
		}
	}
	if !isAuth {
		log.Errorf("用户密码错误 username is %s,password is %s", in.Username, in.Password)
		return nil, UsernameAndPasswordError
	}

	var roles []*pb.Role
	for _, role := range userTokenDomain.DomainRoleResource.Role {
		roles = append(roles, &pb.Role{
			Rid:        role.Rid,
			RoleName:   role.RoleName,
			CreateTime: role.CreateTime.Unix(),
			UpdateTime: role.UpdateTime.Unix(),
		})
	}

	userToken := &pb.UserToken{
		User: &pb.User{
			Uid:        userTokenDomain.User.Uid,
			Username:   userTokenDomain.User.Username,
			State:      userTokenDomain.User.State,
			Nickname:   userTokenDomain.User.Nickname,
			LoginState: 1,
			CreateTime: userTokenDomain.User.CreateTime.Unix(),
			UpdateTime: userTokenDomain.User.UpdateTime.Unix(),
		},
		Token: userTokenDomain.Token,
		DomainRole: &pb.DomainRole{
			Did:        userTokenDomain.DomainRoleResource.Domain.Did,
			DomainName: userTokenDomain.DomainRoleResource.Domain.DomainName,
			Role:       roles,
		},
	}
	log.Infof("认证成功 username is %#v", userToken)
	return &pb.AuthenticationResponse{UserToken: userToken}, nil
}

/*处理多租户返回信息*/
func handlerUserTokenTenants(userTokenTenants *drbac.UserTokenTenants) *pb.UserToken {
	user := &pb.User{
		Uid:        userTokenTenants.User.Uid,
		Username:   userTokenTenants.User.Username,
		State:      userTokenTenants.User.State,
		Nickname:   userTokenTenants.User.Nickname,
		LoginState: 1,
		CreateTime: userTokenTenants.User.CreateTime.Unix(),
		UpdateTime: userTokenTenants.User.UpdateTime.Unix(),
	}
	trees := handlerTenantsRoleTree(userTokenTenants)
	ut := &pb.UserToken{Token: userTokenTenants.Token, User: user, TenantRoleTree: trees}
	return ut
}

/*处理多租户返回信息*/
func handlerTenantsRoleTree(userTokenTenants *drbac.UserTokenTenants) []*pb.TenantRoleTree {
	resp := make([]*pb.TenantRoleTree, 0)
	for _, trp := range userTokenTenants.TenantRoleResource {
		tenant := &pb.Tenant{
			Tid:        trp.Tenant.Tid,
			TenantName: trp.Tenant.TenantName,
			CreateTime: trp.Tenant.CreateTime.Unix(),
		}
		var roles []*pb.Role
		for _, role := range trp.Role {
			roles = append(roles, &pb.Role{
				Rid:        role.Rid,
				RoleName:   role.RoleName,
				CreateTime: role.CreateTime.Unix(),
				UpdateTime: role.UpdateTime.Unix(),
			})
		}
		tree := &pb.TenantRoleTree{Tenant: tenant, IsDefaultTenant: trp.IsDefaultTenant, Role: roles}
		resp = append(resp, tree)
	}
	return resp
}

/*handlerDomainTree
递归处理域树
*/
func handlerTenantTree(trees []*drbac.TenantTree) []*pb.TenantTree {

	tenantTrees := make([]*pb.TenantTree, 0)
	for _, v := range trees {
		tenant := &pb.Tenant{Tid: v.Tenant.Tid, TenantName: v.Tenant.TenantName, Pid: v.Tenant.Pid, CreateTime: v.Tenant.CreateTime.Unix(), UpdateTime: v.Tenant.UpdateTime.Unix()}
		tenantTree := &pb.TenantTree{Tenant: tenant}
		tenantTree.Children = handlerTenantTree(v.Children)
		tenantTrees = append(tenantTrees, tenantTree)
	}
	return tenantTrees
}

/*GetAuthorizationInfo
获取用户信息
*/
func (this *AuthServer) GetAuthorizationInfo(ctx context.Context, in *pb.GetAuthorizationInfoRequest) (*pb.GetAuthorizationInfoResponse, error) {
	if in.Token == "" {
		log.Errorf("input parameter error,token is (%s)", in.Token)
		return nil, InvalidArgument
	}
	userToken := this.DrbacServer.GetAuthorizationInfo(in.Token)
	if userToken == nil {
		log.Errorf("用户未登录，token失效 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	if userToken.User.Username == "" {
		log.Errorf("用户未登录，token失效 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	var roles []*pb.Role
	if userToken.DomainRoleResource != nil {
		log.Info("DomainRolePermission.Domain:", userToken.DomainRoleResource.Domain)
		log.Info("DomainRolePermission.Resource:", userToken.DomainRoleResource.Resource)
		for _, role := range userToken.DomainRoleResource.Role {
			roles = append(roles, &pb.Role{
				Rid:        role.Rid,
				RoleName:   role.RoleName,
				CreateTime: role.CreateTime.Unix(),
				UpdateTime: role.UpdateTime.Unix(),
			})
		}
	} else {
		if userToken.TenantRoleResource == nil {
			log.Info("in.Token:", in.Token)
			return nil, TokenIsInvalid
		}
		for _, role := range userToken.TenantRoleResource[0].Role {
			roles = append(roles, &pb.Role{
				Rid:        role.Rid,
				RoleName:   role.RoleName,
				CreateTime: role.CreateTime.Unix(),
				UpdateTime: role.UpdateTime.Unix(),
			})
		}
	}
	reply := &pb.UserToken{
		User: &pb.User{
			Uid:        userToken.User.Uid,
			Username:   userToken.User.Username,
			State:      userToken.User.State,
			Nickname:   userToken.User.Nickname,
			LoginState: 1,
			CreateTime: userToken.User.CreateTime.Unix(),
			UpdateTime: userToken.User.UpdateTime.Unix(),
		},
		Token: userToken.Token,
	}
	if len(userToken.TenantRoleResource) != 0 {
		reply.TenantRoleTree = append(reply.TenantRoleTree, &pb.TenantRoleTree{Tenant: &pb.Tenant{
			Tid:        userToken.TenantRoleResource[0].Tenant.Tid,
			TenantName: userToken.TenantRoleResource[0].Tenant.TenantName,
			CreateTime: userToken.TenantRoleResource[0].Tenant.CreateTime.Unix(),
		}})
	}
	log.Infof("获取用户信息成功%#v", reply.User)
	return &pb.GetAuthorizationInfoResponse{UserToken: reply}, nil
}

/*Logout
退出登录
*/
func (this *AuthServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	log.Info("Start Logout")
	userToken := this.DrbacServer.GetAuthorizationInfo(in.Token)
	if userToken == nil {
		log.Errorf("退出登录成功，用户未登录，token失效 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	if userToken.User.Username == "" {
		log.Errorf("退出登录成功，用户未登录，token失效 token is %s", in.Token)
		return nil, TokenIsInvalid
	}
	err := this.DrbacServer.LogoutToken(in.Token)
	if err != nil {
		log.Errorf("退出登录出错 error is %s,token is %s", err, in.Token)
	}
	a := storage.ActionLog{
		ActionUsername: userToken.User.Username,
		ActionTime:     time.Now().Unix(),
		ActionType:     1,
		ActionName:     "登出",
		ActionObject:   "成功",
		CreateTime:     time.Now(),
	}
	if len(userToken.TenantRoleResource) != 0 {
		a.Tid = userToken.TenantRoleResource[0].Tenant.Tid
	}
	if userToken.DomainRoleResource != nil {
		if userToken.DomainRoleResource.Domain != nil {
			a.Did = userToken.DomainRoleResource.Domain.Did
		}
	}
	if err := a.AddActionLog(); err != nil {
		log.Error("退出登录，添加操作日志失败，ERROR:", err)
	}
	return nil, Successful
}
