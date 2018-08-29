package api

import (
	pb "auth/api"
	"auth/go-drbac/common"
	"auth/go-drbac/drbac"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"strconv"
	"time"
	"github.com/gin-gonic/gin/json"
)

type DomainServer struct {
	DrbacServer *drbac.DrbacServer
}

/*GetUserInfoInDomain
获取域中成员
*/
func (this *DomainServer) GetUserInfoInDomain(ctx context.Context, in *pb.GetUserInfoInDomainRequest) (*pb.GetUserInfoInDomainResponse, error) {
	log.Info("------------GetUserInfoInDomain------------------")
	log.Info("GetUserInfoInDomain Input Info, page:", in.Page, ", count:", in.Count)
	if in.Page == 0 {
		log.Info("GetUserInfoInDomain Input is Empty")
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
		did   int32
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["did"]; ok && val[0] != "" {
		didInt, _ := strconv.Atoi(val[0])
		did = int32(didInt)
	}
	log.Info("token:", token)
	log.Info("did:", did)
	if did == 0 || token == "" {
		log.Error("tid or token is empty")
		return nil, InvalidArgument
	}
	userToken := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userToken)
	if userToken == nil || userToken.DomainRoleResource == nil || userToken.DomainRoleResource.Domain == nil {
		return nil, TokenIsInvalid
	}
	if userToken.DomainRoleResource.Domain.Did == did {
		start := time.Now()
		userRoles, totalCount, err := this.DrbacServer.GetUserRolesInDomain(did, in.Page, in.Count)
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
		return &pb.GetUserInfoInDomainResponse{UserRoles: reply, TotalCount: totalCount}, nil
	} else {
		log.Info("用户不存在改域中 uid is %d, tid is %d", userToken.User.Uid, did)
		return nil, UserDoesNotExist
	}
}

/*UpdateUserRoleInTenant
修改租户中成员角色
*/
func (this *DomainServer) UpdateUserRoleInDomain(ctx context.Context, in *pb.UpdateUserRoleInDomainRequest) (*pb.UpdateUserRoleInDomainResponse, error) {
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
		did   int32
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["did"]; ok && val[0] != "" {
		didInt, _ := strconv.Atoi(val[0])
		did = int32(didInt)
	}
	log.Info("token:", token)
	log.Info("did:", did)
	if did == 0 || token == "" {
		log.Error("did or token is empty")
		return nil, InvalidArgument
	}
	userToken := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userToken)
	if &userToken == nil {
		return nil, SystemError
	}
	if userToken.DomainRoleResource.Domain.Did == did {
		if userToken.User.Uid == in.UpdateUserID {
			log.Errorf("不能修改自己 uid is %d,deleteUserId is %d", userToken.User.Uid, in.UpdateUserID)
			return nil, PermissionDenied
		}
		user, err := this.DrbacServer.GetUserByUid(in.UpdateUserID)
		if err != nil || user.Username == "" {
			log.Errorf("修改用户不存在,", err)
			return nil, UserDoesNotExist
		}
		start := time.Now()
		if err := this.DrbacServer.UpdateUserRoleInDomain(in.UpdateUserID, did, in.UpdateUserRids, token); err != nil {

			log.Errorf("UpdateUserRole is system error %s,uid is %d,updateUserId is %d,updateUserRoleId is %s", err, userToken.User.Uid, in.UpdateUserID, in.UpdateUserRids)
			return nil, SystemError
		}
		log.Info("调用时间：", time.Now().Sub(start))
		return nil, Successful
	} else {
		log.Errorf("UpdateUserRole is error Domain_Does_Not_Exist ,uid is %d,updateUserId is %d,updateUserRoleId is %s", userToken.User.Uid, in.UpdateUserID, in.UpdateUserRids)
		return nil, PermissionDenied
	}
}

/*DeleteUserInDomain
删除租户中成员
*/
func (this *DomainServer) DeleteUserInDomain(ctx context.Context, in *pb.DeleteUserInDomainRequest) (*pb.DeleteUserInDomainResponse, error) {
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
		did   int32
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["did"]; ok && val[0] != "" {
		didInt, _ := strconv.Atoi(val[0])
		did = int32(didInt)
	}
	log.Info("token:", token)
	log.Info("did:", did)
	if did == 0 || token == "" {
		log.Error("did or token is empty")
		return nil, InvalidArgument
	}
	userToken := this.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userToken)
	if &userToken == nil {
		return nil, SystemError
	}

	if userToken.User.Uid == in.DeleteUserID {
		log.Errorf("不能删除自己 uid is %d,deleteUserId is %d", userToken.User.Uid, in.DeleteUserID)
		return nil, PermissionDenied
	}

	if userToken.DomainRoleResource.Domain.Did == did {
		start := time.Now()
		err := this.DrbacServer.DeleteUserInDomain(in.DeleteUserID, did)
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
		log.Info("用户不存在该域中 uid is %d, did is %d", userToken.User.Uid)
		return nil, PermissionDenied
	}
}

/*AddUserInTenant
邀请用户
*/
func (this *DomainServer) AddUserInDomain(ctx context.Context, in *pb.AddUserInDomainRequest) (*pb.AddUserInDomainResponse, error) {
	log.Info("------------AddUserInDomain-----------------")
	if in.Did == 0 || in.AddUserUsername == "" || in.AddUserNickname == "" {
		log.Errorf("len(in.AddUserRids):", len(in.AddUserRids))
		log.Errorf("input error tid is %d,AddUserRids is %d,AddUserUsername is %s,AddUserNickname is %s", in.Did, in.AddUserRids, in.AddUserUsername, in.AddUserNickname)
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
	log.Info("UserTokenTenants.User:", userToken.User)
	log.Info("UserTokenTenants:", userToken.TenantRoleResource)
	if &userToken == nil {
		return nil, SystemError
	}

	//todo 获取域信息
	log.Info("in.did:",in.Did)
	domain, err := this.DrbacServer.GetDomainByDid(in.Did)
	if domain.DomainUrl == "" || err != nil {
		log.Error("GetDomainByDid Err, ", err)
		return nil, DidIsIncorrectOrEmpty
	}
	log.Info("domain.DomainUrl:",domain.DomainUrl)

	//todo 判断角色是否存在
	for _, rid := range in.AddUserRids {
		if _, err := this.DrbacServer.GetRoleByRid(rid); err != nil {
			log.Errorf("角色不存在")
			return nil, InvalidArgument
		}
	}
	//todo 判断用户操作权限
	if userToken.DomainRoleResource.Domain.Did != in.Did {
		log.Info("用户不存在该租户中 uid is (%d) , did is (%d)", userToken.User.Uid, in.Did)
		return nil, PermissionDenied
	}
	start := time.Now()
	yourself, err := this.DrbacServer.GetUserByUid(userToken.User.Uid)
	log.Info("调用时间：", time.Now().Sub(start))
	if err != nil {
		log.Errorf("系统异常：error is %s", err)
		return nil, SystemError
	}
	user, err := this.DrbacServer.GetUserByUsernameAndDid(in.AddUserUsername, in.Did)
	log.Info("GetUserInfo:",user)
	log.Info("UserToken:",user.Token)
	//todo 准备邮件信息
	sendData := SendData{
		To:       user.Username,
		Nickname: yourself.Nickname,
		Token:    user.Token,
		Url:      domain.DomainUrl,
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
		//todo 判断是否已在邀请的租户中 1是 重新邀请  2不是，插入udr关系 邀请
		isTrue := this.DrbacServer.IsExistedInDomain(user.Uid, in.Did)
		if isTrue {
			if user.State == 1 {
				res := HttpSendMailPost(req)
				log.Info("HttpSendMailPost:", res)
				return nil, Successful
			} else {
				if userToken.User.Uid == user.Uid {
					log.Errorf("不能添加自己 uid is %d,AddUserUsername is %s", in.AddUserUsername)
					return nil, UserAlreadyExist
				} else {
					log.Errorf("用户已存在 AddUserUsername is %s", in.AddUserUsername)
					return nil, UserAlreadyExist
				}
			}
		} else {
			if err := this.DrbacServer.AddUserInDomain(user.Uid, in.Did, in.AddUserRids); err != nil {
				res := HttpSendMailPost(req)
				log.Info("HttpSendMailPost:", res)
				return nil, Successful
			}
		}
	}
	//todo 被邀请人信息直接存数据库
	emailToken := string(Krand(32, KC_RAND_KIND_ALL))
	if err := this.DrbacServer.AddDomainUser(in.AddUserUsername, in.AddUserNickname, emailToken, in.Did, in.AddUserRids); err != nil {
		log.Error("AddUser Error, ", err)
		return nil, SystemError
	}

	//todo 发送邀请邮件
	//todo 准备邮件信息
	sendData = SendData{
		To:       in.AddUserUsername,
		Nickname: yourself.Nickname,
		Token:    emailToken,
		Url:      domain.DomainUrl,
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
func (this *DomainServer) EnterDomain(ctx context.Context, in *pb.EnterDomainRequest) (*pb.EnterDomainResponse, error) {
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
	return nil, Successful
}

func (this *DomainServer) GetUserCountInDomain(ctx context.Context, in *pb.GetUserCountInDomainRequest) (*pb.GetUserCountInDomainResponse, error) {
	return nil, nil
}

func (this *DomainServer) GetDomain(ctx context.Context, in *pb.GetDomainRequest) (*pb.GetDomainResponse, error) {
	return nil, nil
}

func (this *DomainServer) GetDomains(ctx context.Context, in *pb.GetDomainsRequest) (*pb.GetDomainsResponse, error) {
	domains,err := this.DrbacServer.GetDomains()
	if err != nil {
		log.Error("GetDomains Error")
		return nil,SystemError
	}
	log.Info("len(domains)",len(domains))
	var res []*pb.Domain
	for _,v := range domains {
		log.Info(v.Did)
		if v.Did != 0 {
			res = append(res, &pb.Domain{
				Did:v.Did,
				DomainName:v.DomainName,
				DomainUrl:v.DomainUrl,
				DomainState:v.DomainState,
			})
		}
	}
	log.Info(res)
	return &pb.GetDomainsResponse{Domain:res}, nil
}

func (this *DomainServer) AddUserTenantACL(ctx context.Context, in *pb.AddUserTenantACLRequest) (*pb.AddUserTenantACLResponse, error) {
	log.Info("------------AddUserTenantACL----------------")
	if in.Username == "" || len(in.Tids) == 0 {
		log.Error("Input InvalidArgument, in.Username:", in.Username, " in.Tids: ", in.Tids)
		return nil, InvalidArgument
	}
	err := this.DrbacServer.AddUserTenantACL(in.Username, in.Did, in.Tids)
	if err != nil {
		log.Info("AddUserTenantACL Error,", err)
		return nil, SystemError
	}
	return nil, Successful
}
