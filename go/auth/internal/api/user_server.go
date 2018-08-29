package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"auth/go-drbac/common"
	"auth/module"
	"auth/util"
	"github.com/garyburd/redigo/redis"
	"google.golang.org/grpc/metadata"
	."auth/util"
	"strconv"
	"auth/config"
)

type UserServer struct {
	DrbacServer *drbac.DrbacServer
}

/*UpdateNickname
修改用户昵称 nickname
*/
func (u *UserServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil,SystemError
	}
	var (
		val []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:",token)
	userTokenTenants := u.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:",userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	//todo 正常流程
	if in.Nickname == ""{
		log.Errorf("输入参数错误，%#v",in)
		return nil, InvalidArgument
	}
	err := u.DrbacServer.UpdateNickname(userTokenTenants.User.Uid,in.Nickname,token)
	if err != nil {
		log.Errorf("系统异常，error is %s",err)
		return nil, SystemError
	}
	log.Info("修改用户信息成功，uid is %d,nickname is %s",in.Uid,in.Nickname)
	return nil,Successful
}
/*UpdatePassword
修改用户密码
*/
func (u *UserServer) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	//todo ctx拿数据
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil,SystemError
	}
	var (
		val []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:",token)
	userTokenTenants := u.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:",userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	//todo 正常流程
	if in.Password == "" || in.NewPassword == ""{
		log.Errorf("输入参数错误，%#v",in)
		return nil, InvalidArgument
	}
	isAuth,state,err := u.DrbacServer.Authentication(userTokenTenants.User.Uid,in.Password)
	if err != nil {
		log.Errorf("系统异常，error is %s",err)
		return nil, SystemError
	}
	if state != 3 {
		log.Errorf("认证失败，账号不可用")
		return nil, AccountDisableToUse
	}
	if !isAuth {
		log.Errorf("认证失败，用户名或密码错误")
		return nil, UsernameAndPasswordError
	}
	err = u.DrbacServer.UpdatePassword(userTokenTenants.User.Uid,in.NewPassword)
	if err != nil {
		log.Errorf("系统异常，error is %s",err)
		return nil, SystemError
	}
	log.Info("修改密码成功 username is %s",in.Username)
	return nil, Successful
}
/*FindPassword
找回密码
*/
func (u *UserServer) FindPassword(ctx context.Context, in *pb.FindPasswordRequest) (*pb.FindPasswordResponse, error) {
	if in.Username == "" {
		log.Errorf("输入参数错误，%#v",in)
		return nil, InvalidArgument
	}
	if in.Did != 0 {
		user,err :=u.DrbacServer.GetUserByUsernameAndDid(in.Username,in.Did)
		if err != nil {
			if err == common.ErrDoesNotExist {
				log.Errorf("用户名不存在 username is %s",in.Username)
				return nil, UsernameIsIncorrectOrEmpty
			}
			log.Errorf("系统异常，error is %s",err)
		}
		client := config.C.Redis.Pool.Get()
		defer client.Close()
		token := string(util.Krand(32, util.KC_RAND_KIND_ALL))
		if _, err = client.Do("set", module.FindpwReidsKeyPrefix+token, user.Uid, "EX", 1*3600); err != nil {
			log.Errorf("系统异常，FindPassword Set Redis Token error %s ,username is %s",err,user.Username)
			return nil, SystemError
		}
		log.Info("in.Did:",in.Did)
		var mailModel Model
		switch in.Did {
		case 100002:
			mailModel = CotxInvitationMail
		case 100001:
			mailModel = RadacatInvitationMail
		case 100003:
			mailModel = PensLinkInvitationMail
		default:
			log.Error("找不到did对应的邮件模板")
			return nil,SystemError
		}
		go util.SendMail(user.Username,[]string{"token:"+token,"nickname:"+user.Nickname},mailModel)
		log.Info("找回密码成功，username is %s",in.Username)
		return nil, Successful
	}else if in.Tid != 0{
		user,err :=u.DrbacServer.GetUserByUsernameAndTid(in.Username,in.Tid)
		if err != nil {
			if err == common.ErrDoesNotExist {
				log.Errorf("用户名不存在 username is %s",in.Username)
				return nil, UsernameIsIncorrectOrEmpty
			}
			log.Errorf("系统异常，error is %s",err)
		}
		client := config.C.Redis.Pool.Get()
		defer client.Close()
		token := string(util.Krand(32, util.KC_RAND_KIND_ALL))
		if _, err = client.Do("set", module.FindpwReidsKeyPrefix+token, user.Uid, "EX", 1*3600); err != nil {
			log.Errorf("系统异常，FindPassword Set Redis Token error %s ,Uid is %s",err,user.Uid)
			return nil, SystemError
		}
		//通过tid获取did
		tenant,err := u.DrbacServer.GetTenant(in.Tid)
		if tenant.Tid == 0 || err != nil {
			log.Error("GetTenant Error,",err)
			return nil, SystemError
		}
		log.Info("tenant.Did:",tenant.Did)
		var mailModel Model
		switch tenant.Did {
		case 100002:
			mailModel = CotxInvitationMail
		case 100001:
			mailModel = RadacatInvitationMail
		case 100003:
			mailModel = PensLinkInvitationMail
		default:
			log.Error("找不到did对应的邮件模板")
			return nil,SystemError
		}
		go util.SendMail(user.Username,[]string{"token:"+token,"nickname:"+user.Nickname},mailModel)
		log.Info("找回密码成功，username is %s",in.Username)
		return nil, Successful
	}else {
		log.Errorf("输入参数错误，%#v",in)
		return nil, InvalidArgument
	}
}
/*ResetPassword
重置密码
*/
func (u *UserServer) ResetPassword(ctx context.Context, in *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	if in.Password == "" || in.Token == "" {
		log.Errorf("输入参数错误，%#v",in)
		return nil, InvalidArgument
	}
	client := config.C.Redis.Pool.Get()
	defer client.Close()
	uidString, err := redis.String(client.Do("GET", module.FindpwReidsKeyPrefix+in.Token))
	client.Do("DEL",module.FindpwReidsKeyPrefix+in.Token)
	if err != nil || uidString == "" {
		log.Errorf("token失效，ResetPassword GET Username By Token Error %s,token is %s",err,in.Token)
		return nil, PasswordTokenIsIncorrectOrEmpty
	}
	log.Info("uidString:",uidString)
	uidInt,err := strconv.Atoi(uidString)
	if err != nil {
		log.Error("uid转化失败")
		return nil, PasswordTokenIsIncorrectOrEmpty
	}
	uid := int32(uidInt)
	user,err := u.DrbacServer.GetUserByUid(uid)
	if err != nil {
		if err == common.ErrDoesNotExist {
			log.Errorf("用户不存在ErrDoesNotExist token is %s , uid is %s",in.Token,uid)
			return nil, UserDoesNotExist
		}
		log.Errorf("系统异常 error is %s",err)
		return nil, SystemError
	}
	err = u.DrbacServer.UpdatePassword(user.Uid,in.Password)
	if err != nil {
		log.Errorf("系统异常，Reset Password UpdatePassword Error is %s,uid is %s,password is %s,token is %s",err,uid,in.Password,in.Token)
		return nil, SystemError
	}
	log.Infof("重置密码成功，%#v",user)
	return nil, Successful
}

/*UpdateUserState
修改用户状态
*/
func (u *UserServer) UpdateUserState(ctx context.Context, in *pb.UpdateUserStateRequest) (*pb.UpdateUserStateResponse, error) {
	if in.UpdateUid == 0 || in.UpdateState == 0{
		log.Errorf("输入参数错误，%#v",in)
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
		tid int64
		did int64
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["tid"]; ok && val[0] != ""{
		tid, _ = strconv.ParseInt(val[0], 10, 64)
	}
	if val, ok = md["did"]; ok && val[0] != ""{
		did, _ = strconv.ParseInt(val[0], 10, 64)
	}
	log.Info("token:", token)
	log.Info("tid:",tid)
	log.Info("did:",did)
	if tid == 0 && did == 0 {
		log.Error("did or tid is empty")
		return nil, InvalidArgument
	}
	if token == "" {
		log.Error("token is empty")
		return nil, InvalidArgument
	}
	userTokenTenants := u.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:", userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}


	if userTokenTenants.User.Uid == in.UpdateUid {
		log.Error("不能修改自己的状态")
		return nil, PermissionDenied
	}

	if user,err := u.DrbacServer.GetUserByUid(in.UpdateUid); err != nil || user.Username == "" {
		log.Error("uid:",in.UpdateUid," does not exist")
		return nil, UserDoesNotExist
	}

	ra, err := u.DrbacServer.UpdateUserState(in.UpdateUid, in.UpdateState)
	if err != nil || ra == 0 {
		log.Errorf("系统异常，error is %s, RowsAffected is %d",err,ra)
		return nil, SystemError
	}
	log.Info("修改状态成功 Uid is %d,State is %d",in.UpdateUid,in.UpdateState)
	return nil, Successful
}

/*UpdateNicknameAndPassword
修改用户昵称和密码
*/
func (u *UserServer) UpdateNicknameAndPassword(ctx context.Context, in *pb.UpdateNicknameAndPasswordRequest) (*pb.UpdateNicknameAndPasswordResponse, error) {
	//todo ctx拿参数
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil,SystemError
	}
	var (
		val []string
		token string
	)
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	log.Info("token:",token)
	userTokenTenants := u.DrbacServer.GetAuthorizationInfo(token)
	log.Info("UserTokenTenants:",userTokenTenants)
	if &userTokenTenants == nil {
		return nil, SystemError
	}
	//todo 正常流程
	if in.Nickname == "" || in.Password == "" || in.NewPassword == ""{
		log.Errorf("输入参数错误，%#v",in)
		return nil,InvalidArgument
	}
	user,err := u.DrbacServer.GetUserByUid(userTokenTenants.User.Uid)
	log.Info("UpdateNicknameAndPassword UserInfo:",user)
	if err != nil {
		if err == common.ErrDoesNotExist {
			log.Errorf("GetUserByUid 用户不存在，uid is %d",in.Uid)
			return nil,UserDoesNotExist
		}
		log.Errorf("GetUserByUid 系统异常，error is %s",err)
		return nil,SystemError
	}
	isAuth,_,err := u.DrbacServer.Authentication(user.Uid,in.Password)
	if err != nil {
		log.Errorf("Authentication 系统异常，error is %s",err)
		return nil,SystemError
	}
	if !isAuth {
		log.Errorf("认证失败，用户名或密码错误")
		return nil,UsernameAndPasswordError
	}
	err = u.DrbacServer.UpdateNicknameAndPassword(token,user.Uid, in.Nickname, in.NewPassword)
	if err != nil {
		log.Errorf("UpdateNicknameAndPassword 系统异常，error is %s",err)
		return nil, SystemError
	}
	log.Info("修改密码成功 username is %s",user.Username)
	return nil,Successful
}
