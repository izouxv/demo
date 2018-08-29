package service

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"net/url"
	"petfone-rpc/db"
	"petfone-rpc/core"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"regexp"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"strconv"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	//un_login = 0	//未登录
	login            = 1 //已登录
	Day       int32  = 24 * 3600
)

var mobileReg = regexp.MustCompile("^((1[3,5,8][0-9])|(14[5,7])|(17[0,1,6,7,8]))\\d{8}$")

type Rpc_sso struct {
}
// 发送验证码和接口访问限制
func limitCount(client redis.Conn,redisFlag string,source string, deadline int,maximum int) (*pb.SsoReply, error) {
	times, err := client.Do("get", source[:2]+util.LimitIp+redisFlag)
	if err != nil {
		log.Error("Get ip times Failed,", err)
		return &pb.SsoReply{Code: util.Code_err}, nil
	}
	if times == nil {
		// 设置次数
		reply, err := client.Do("set", source[:2]+util.LimitIp+redisFlag, 1, "EX", deadline)
		if err != nil || reply == 0 {
			log.Info("Set ip times Failed,", err)
			return &pb.SsoReply{Code: 10001}, nil
		}
	} else {
		count,err:= strconv.Atoi(string(times.([]uint8)))
		log.Infof("已访问次数count(%d)",count)
		if count > maximum {
			log.Errorf("接口已访问次数超限count(%d)",count)
			return &pb.SsoReply{Code: util.Mobile_times_exceed}, nil
		}
		reply, err := client.Do("incrby", source[:2]+util.LimitIp+redisFlag,1)
		if err != nil {
			log.Info("incrby ip times Failed,", err)
			return &pb.SsoReply{Code: 10001}, nil
		}
		log.Debugf("接口已访问次数(%#v)",reply)
	}
	return &pb.SsoReply{Code: 10000}, nil
}

//手机发送验证码
func (this *Rpc_sso) SendMobileCode(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	log.Info("SendMobileCode:",in)
	if in.Username == "" || in.CodeType == 0 || in.Source == "" {
		log.Info("The input is empty! in.Username:", in.Username, ",in.CodeType:", in.CodeType, ", in.Source", in.Source)
		return &pb.SsoReply{Code: 33001}, nil
	}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: 33001}, nil
	}
	s := db.Sso{Username: in.Username}
	err := s.GetByName()
	if err != nil && core.ConstStr.NotFound != err.Error() {
		log.Info("SendMobileCode-err:", err)
		return &pb.SsoReply{Code: util.System_error}, nil
	}
	if s.Id == 0 && in.CodeType == 2 {
		log.Info("cannot find user :" + s.Username)
		return &pb.SsoReply{Code: 33002}, nil
	} else if s.Id != 0 && in.CodeType == 1 {
		log.Info("user already exist:" + s.Username)
		return &pb.SsoReply{Code: 33008}, nil
	}
	//todo 生成code和alisms.SendSms(code)
	code := string(util.Krand(6, util.KC_RAND_KIND_NUM))
	//todo redis save code
	log.Info(in.Username, " start dbclient.RedisClient")
	client := core.RedisClient(6379)
	defer client.Close()

	//// todo 验证ip访问次数
	//reply,err := limitCount(client,in.Ip,in.SsoRequest.Source,1*60,10)
	//if reply.Code != 10000 {
	//	log.Error("ip访问次数异常")
	//	return &pb.SsoReply{Code:util.Ip_times_exceed}, nil
	//}
	//// todo 验证发送次数
	//reply,err = limitCount(client,in.SsoRequest.Username,in.SsoRequest.Source,5*60,5)
	//if reply.Code != 10000 {
	//	log.Error("手机验证码发送次数异常")
	//	return &pb.SsoReply{Code:util.Mobile_times_exceed}, nil
	//}
	if in.CodeType == 1 {
		reply, err := client.Do("set", in.Source[:2]+util.SendMobileCode+in.Username, code, "EX", 5*60)
		if err != nil || reply == 0 {
			log.Info("Set Code Failed,", err)
			return &pb.SsoReply{Code: 33005}, nil
		}
	} else {
		reply, err := client.Do("set", in.Source[:2]+util.SendMobileCode+in.Username, code, "EX", 5*60)
		if err != nil || reply == 0 {
			log.Info("Set Code Failed,", err)
			return &pb.SsoReply{Code: 33005}, nil
		}
	}
	go func() {
		util.SendSms(in.Username,code)
	}()
	return &pb.SsoReply{Code: util.Success}, nil
}

//通过用户名查询用户信息（验证账号重复）
func (this *Rpc_sso) GetUserByName(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	log.Info("Start RPC GetUserByName ")
	if in.Source == "" || in.Username == "" {
		log.Info("username check failed!")
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	s := &db.Sso{Username: in.Username}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	err := s.GetByName()
	if err != nil || s.Id == 0 {
		log.Info("cannot find user :" + s.Username)
		return &pb.SsoReply{Code: util.User_unexist}, nil
	}
	log.Info("username:" + s.Username + "already exists")
	nickname, _ := url.QueryUnescape(s.Nickname)
	return &pb.SsoReply{Code: 10000, Uid: s.Id, Username: s.Username, State: s.State, Nickname: nickname}, nil
}

//用户注册
func (this *Rpc_sso) Add(ctx context.Context, req *pb.SsoRequest) (*pb.SsoReply, error) {
	log.Info("Add-req:", req)
	if util.VerifyParamsStr(req.Source,req.Username,req.Password,req.Salt,req.Nickname) {
		log.Info("Add VerifyParamsStr input is empty")
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	if util.CheckSource(req.Source) {
		log.Info("Add CheckSource err:", req.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	sso := &db.Sso{Username: req.Username}
	if err := sso.GetByName(); err == nil {
		log.Info("Add-GetByName-err:", req.Source)
		return &pb.SsoReply{Code: util.User_exists}, nil
	}
	req.Password = util.EncryptWithSalt(req.Password, req.Salt)
	node, err := db.NewNode(1)
	if err != nil {
		log.Error("Failed to Get Uid", err)
		return &pb.SsoReply{Code: util.System_error}, nil
	}
	uid, err := node.Generate()
	if err != nil {
		log.Error("Add Generate-err:",err)
		return &pb.SsoReply{Code: util.System_error}, nil
	}
	nowTime := util.GetNowTime()
	sso = &db.Sso{Id: uid, Username: req.Username, Password: req.Password, Salt: req.Salt, State: req.State,
		Nickname: url.QueryEscape(req.Nickname), UpdateTime: nowTime, CreateTime: nowTime, LastLoginTime:nowTime}
	if req.AgentInfo != nil && req.AgentInfo.Ip != "" {
		ip, err := util.IPToInt64(req.AgentInfo.Ip)
		if err == nil {
			chanIp := make(chan string)
			go util.IPToAddr(chanIp,req.AgentInfo.Ip)
			select {
			case <-time.After(time.Second*3):
			case addr := <-chanIp:
				sso.RegAddr = addr
				sso.LastLoginAddr = addr
			}
		}
		sso.RegIp = ip
		sso.LastLoginIp = ip
		sso.LastLoginDevInfo = req.AgentInfo.DevInfo
		if err != nil {
			log.Info("ip-err:",err)
		}
	}
	//todo account插入手机或邮箱
	account := &db.Account{Id: sso.Id, Avatar:core.ConstStr.UserAvatar}
	if mobileReg.MatchString(req.Username) {
		account.Phone = req.Username
	} else {
		account.Email = req.Username
	}
	//todo user function 插入记录
	petfone := &db.PetfonePo{Uid:sso.Id, Radius:50, CreationTime:nowTime, UpdateTime:nowTime, DataState:1}
	//todo 执行sql
	err = db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
		sso.Insert(tx)
		account.UpdateAccount(tx)
		petfone.SetPetfoneDB(tx)
		return nil
	})
	return &pb.SsoReply{Code: util.Success}, nil
}

//用户登录
func (this *Rpc_sso) Login(ctx context.Context, req *pb.SsoRequest) (*pb.SsoReply, error) {
	log.Info("Login-req:", req)
	if util.VerifyParamsStr(req.Source, req.Password, req.Username) || util.CheckSource(req.Source) {
		log.Info("Login VerifyParamsStr is empty or CheckSource err")
		return &pb.SsoReply{Code: 33001}, nil
	}
	s := &db.Sso{Username: req.Username}
	err := s.GetByName()
	if err != nil || s.Id == 0 {
		log.Info("cannot find user :" + s.Username)
		return &pb.SsoReply{Code: util.User_unexist}, nil
	}
	if s.State != 3 {
		log.Info("Username:", s.Username, " Account Unactivated!")
		return &pb.SsoReply{Code: util.Account_unactivated}, nil
	}
	password := util.EncryptWithSalt(req.Password, s.Salt)
	//todo 密码错误
	if password != s.Password {
		log.Info("username:" + s.Username + ": passworld mismatched")
		return &pb.SsoReply{Code: util.Password_err}, nil
	}
	//todo 密码正确
	nickname, _ := url.QueryUnescape(s.Nickname)
	value, err := util.Json.Marshal(&pb.SsoReply{Uid: s.Id, Username: s.Username, Nickname: nickname,
	State: s.State, LoginState: login})
	if err != nil {
		log.Info("err json.Marshal value,", err)
		return &pb.SsoReply{Code: util.Json_error}, nil
	}
	sessionName := (&util.Session{Uid: s.Id,Username:s.Username}).GenerateSession()
	client := core.RedisClient(6380)
	defer client.Close()
	source := util.Source{}
	source.UnBase64(req.Source)
	log.Info("source:",source.Client)
	reply := ""
	if source.Client == 1 {
		reply, err = redis.String(client.Do("set", req.Source[:2]+util.LoginSession+sessionName, value))
	} else {
		reply, err = redis.String(client.Do("set", req.Source[:2]+util.LoginSession+sessionName, value, "EX", Day))
	}
	if err != nil || reply != "OK" {
		log.Info("SSO login Redis set failed!", err)
		return &pb.SsoReply{Code: util.Token_err_empty}, nil
	}
	ssoRe :=  &pb.SsoReply{Code: util.Success, Uid: s.Id, Username: s.Username, Nickname: nickname,
		State: s.State, LoginState: login, SessionName: sessionName, Token:sessionName}
	log.Info("login successful : ",ssoRe)
	return ssoRe, nil
}

//从Redis中获取用户基本信息
func (this *Rpc_sso) GetUserInfo(ctx context.Context, req *pb.SsoRequest) (*pb.SsoReply, error) {
	log.Info("GetUserInfo:", req)
	if req.Source == "" || req.SessionName == "" {
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	var ssoReply = pb.SsoReply{}
	client := core.RedisClient(6381)
	defer client.Close()
	userInfo, err := client.Do("get", req.Source[:2]+util.LoginSession+req.SessionName)
	if err != nil || userInfo == nil {
		log.Info("Get userInfo Failed,", err)
		return &pb.SsoReply{Code: util.Token_err_empty}, nil
	}
	errJson := util.Json.Unmarshal(userInfo.([]uint8), &ssoReply)
	if errJson != nil {
		log.Info("err json.Unmarshal:", err)
		return &pb.SsoReply{Code: util.Json_error}, nil
	}
	fmt.Println("ssoReply:", ssoReply)
	//todo session续租
	so := util.Source{}
	so.UnBase64(req.Source)
	if so.Client == 2 {
		reply, errSession := client.Do("set", req.Source[:2]+util.LoginSession+req.SessionName, userInfo, "EX", Day)
		if errSession != nil || reply == 0 {
			log.Info("SSO login Redis set failed!", errSession)
			return &pb.SsoReply{Code: util.Token_err_empty}, nil
		}
	} else if so.Client == 3 {
		reply, errSession := client.Do("set", req.Source[:2]+util.LoginSession+req.SessionName, userInfo, "EX", 7*Day)
		if errSession != nil || reply == 0 {
			log.Info("SSO login Redis set failed!", errSession)
			return &pb.SsoReply{Code: util.Token_err_empty}, nil
		}
	}

	return &pb.SsoReply{Code: util.Success, Uid: ssoReply.Uid, Username: ssoReply.Username, Nickname: ssoReply.Nickname, State: ssoReply.State, LoginState: login}, nil
}

//验证密码
func (this *Rpc_sso) CheckPassword(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	if in.Source == "" || in.Uid == 0 || in.Password == "" {
		log.Info("the input is empty! in.Uid:", in.Uid, " in.Password:", in.Password)
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	check := &db.Sso{Id: in.Uid}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	err := check.GetUserInfoById()
	if err != nil || check.Username == "" {
		return &pb.SsoReply{Code: util.User_unexist}, nil
	}
	if util.EncryptWithSalt(in.Password, check.Salt) != check.Password {
		log.Info("password mismatched")
		return &pb.SsoReply{Code: util.Password_err}, nil
	}
	return &pb.SsoReply{Code: util.Success}, nil
}

//修改密码
func (this *Rpc_sso) UpdatePassword(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	if in.Source == "" || in.Uid == 0 || in.Password == "" || in.Salt == "" || in.SessionName == "" {
		log.Info("the input is empty! in.Uid:", in.Uid, " in.Password:", in.Password, " in.Salt:", in.Salt, " in.SessionName:", in.SessionName)
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	s := &db.Sso{Id: in.Uid, Password: util.EncryptWithSalt(in.Password, in.Salt), Salt: in.Salt}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	if err := s.UpdatePasswordAndSalt(); err != nil {
		log.Info("errUpdatePasswoedAndSalt:", err)
		return &pb.SsoReply{Code: util.Mysql_err}, nil
	}
	//退出登录
	//client := core.RedisClient(6380)
	//defer client.Close()
	//userInfo, err := client.Do("get", in.Source[:2]+util.LoginSession+in.SessionName)
	//if err != nil || userInfo == nil {
	//	log.Info("Get userInfo Failed,", err)
	//	return &pb.SsoReply{Code: util.Success}, nil
	//}
	//delUserInfo, err := client.Do("del", in.Source[:2]+util.LoginSession+in.SessionName)
	//if err != nil || delUserInfo == nil {
	//	log.Info("Del userInfo Failed,", err)
	//	return &pb.SsoReply{Code: util.Success}, nil
	//}
	//var ssoReply = pb.SsoReply{}
	//errJson := util.Json.Unmarshal(userInfo.([]uint8), &ssoReply)
	//if errJson != nil {
	//	log.Info("err json.Unmarshal:", err)
	//	return &pb.SsoReply{Code: util.Json_error}, nil
	//}
	log.Info("UpdatePassword updatePassword successed")
	return &pb.SsoReply{Code: util.Success}, nil
}

//邮箱找回密码
func (this *Rpc_sso) FindPasswordByMail(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	if in.Source == "" || in.Username == "" {
		log.Info("FindPasswordByMail input is empty")
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	sso := &db.Sso{Username: in.Username}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	err := sso.GetByName()
	if err != nil || sso.Id == 0 {
		log.Info("Cannot Find User:", in.Username)
		return &pb.SsoReply{Code: util.User_unexist}, nil
	}
	//发送邮件
	go func(sso *db.Sso,source string) {
		sso.Nickname, err = url.QueryUnescape(sso.Nickname)
		if err != nil {
			log.Error("FindPasswordByMail QueryUnescape err:", err);return
		}
		client := core.RedisClient(6379)
		defer client.Close()
		token := string(util.Krand(32, util.KC_RAND_KIND_ALL))
		if _, err := client.Do("set", source+util.ResetPasswordByMail+sso.Username, token, "EX", 1*3600); err != nil {
			log.Error("FindPasswordByMail redis err:", err);return
		}
		util.SendMail(sso.Username,sso.Nickname,token)
	}(sso,in.Source[:2])
	return &pb.SsoReply{Code: util.Success}, nil
}

//邮箱重置密码
func (this *Rpc_sso) ResetPassword(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	if in.Source == "" || in.Username == "" || in.Password == "" || in.Token == "" || in.Salt == "" {
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	//todo check token available
	log.Info(in.Username, " start dbclient.RedisClient")
	client := core.RedisClient(6379)
	defer client.Close()
	token, err := redis.String(client.Do("get", in.Source[:2]+util.ResetPasswordByMail+in.Username))
	if err != nil {
		log.Info("Get token Failed,", err)
		return &pb.SsoReply{Code: util.System_error}, nil
	}
	if token == "" {
		log.Info("user token empty")
		return &pb.SsoReply{Code: util.Token_err_empty}, nil
	}
	log.Info("in.Token:", in.Token,",token:",token)
	if token != in.Token {
		return &pb.SsoReply{Code: util.Token_err_empty}, nil
	}
	s := db.Sso{Username: in.Username}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	err = s.GetByName()
	if err != nil || s.Id == 0 {
		log.Info("Cannot Find User:", in.Username)
		return &pb.SsoReply{Code: util.User_unexist}, nil
	}
	s.Password = util.EncryptWithSalt(in.Password, in.Salt)
	s.Salt = in.Salt
	if err := s.UpdatePasswordAndSaltByName(); err != nil {
		return &pb.SsoReply{Code: util.Mysql_err}, nil
	}
	return &pb.SsoReply{Code: util.Success}, nil
}

//用户登出
func (this *Rpc_sso) Logout(ctx context.Context, req *pb.SsoRequest) (*pb.SsoReply, error) {
	if req.Source == "" || req.SessionName == "" {
		log.Info("sessionName is empty! req.SessionName:", req.SessionName)
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	client := core.RedisClient(6380)
	defer client.Close()
	userInfo, err := client.Do("get", req.Source[:2]+util.LoginSession+req.SessionName)
	if err != nil || userInfo == nil {
		log.Info("Get userInfo Failed,", err)
		return &pb.SsoReply{Code: util.Success}, nil
	}
	delUserInfo, err := client.Do("del", req.Source[:2]+util.LoginSession+req.SessionName)
	if err != nil || delUserInfo == nil {
		log.Info("Del userInfo Failed,", err)
		return &pb.SsoReply{Code: util.Success}, nil
	}
	var s = pb.SsoReply{}
	errJson := util.Json.Unmarshal(userInfo.([]uint8), &s)
	if errJson != nil {
		log.Info("err json.Unmarshal:", err)
		return &pb.SsoReply{Code: util.Json_error}, nil
	}
	log.Info("Uid:", s.Uid, " logout successed")
	return &pb.SsoReply{Code: util.Success}, nil
}

//通过用户名修改密码
func (this *Rpc_sso) UpdatePasswordByName(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	if in.Source == "" || in.Username == "" || in.Password == "" || in.Salt == "" {
		log.Info("The input is empty in.Username:", in.Username, " in.Password:", in.Password, " in.Salt:", in.Salt)
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	s := db.Sso{Username: in.Username, Password: util.EncryptWithSalt(in.Password, in.Salt), Salt: in.Salt}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	err := s.UpdatePasswordAndSaltByName()
	if err != nil {
		log.Info("err UpdatePasswordAndSaltByName:", err)
		return &pb.SsoReply{Code: util.Mysql_err}, nil
	}
	return &pb.SsoReply{Code: util.Success}, nil
}

//修改账号状态
func (this *Rpc_sso) UpdateState(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	if in.Uid == 0 || in.State == 0 {
		log.Info("The input is empty! in.Uid:", in.Uid, " in.State:", in.State)
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	s := &db.Sso{Id: in.Uid, State: in.State}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	if err := s.UpdateSso(core.MysqlClient); err != nil {
		dbc := core.MysqlClient.Begin()
		if err := s.UpdateSso(dbc); err != nil {
			dbc.Rollback()
			log.Info("changeStateFailed:", s)
			return &pb.SsoReply{Code: util.Mysql_err}, nil
		}
		dbc.Commit()
	}
	return &pb.SsoReply{Code: util.Success}, nil
}

//批量查询sso信息
func (this *Rpc_sso) GetBatchSsoInfos(ctx context.Context, in *pb.MultiSsoRequest) (*pb.MapSsoReply, error) {
	userMap := make(map[int32]*pb.SsoReply)
	if in.Source == "" || len(in.Ssos) == 0 {
		log.Info("in.Ssos is empty")
		return &pb.MapSsoReply{Code: util.Params_err_empty}, nil
	}
	var uids []int32 = make([]int32, 0)
	for k := range in.Ssos {
		if k != 0 {
			uids = append(uids, k)
		}
	}
	log.Info("Uids:", uids)
	s := db.Sso{}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.MapSsoReply{Code: util.Source_err_empty}, nil
	}
	rs, err := s.GetBatchSsoInfo(uids)
	if err != nil {
		return &pb.MapSsoReply{Code: util.System_error}, nil
	}
	if len(rs) == 0 {
		log.Info("errGetBatchSsoInfo")
		return &pb.MapSsoReply{Code: util.User_unexist}, nil
	}
	for i := 0; i < len(rs); i++ {
		if (rs)[i].Id != 0 {
			nickname, _ := url.QueryUnescape(rs[i].Nickname)
			userMap[(rs)[i].Id] = &pb.SsoReply{Uid: (rs)[i].Id, Username: (rs)[i].Username, Nickname: nickname, State: (rs)[i].State}
		}
	}
	return &pb.MapSsoReply{Code: util.Success, Ssos: userMap}, nil
}

//校验验证码
func (this *Rpc_sso) CheckCode(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	log.Info("CheckCode:", in)
	if in.Source == "" || in.Username == "" || in.Code == "" || in.CodeType == 0 {
		log.Info("Func CheckCode Input is Empty")
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	//todo check user
	s := db.Sso{Username: in.Username}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	s.GetByName()
	if s.Id != 0 && in.CodeType == 1 {
		log.Info("user already exist:" + in.Username)
		return &pb.SsoReply{Code: util.User_exists}, nil
	}
	if s.Id == 0 && in.CodeType == 2 {
		log.Info("cannot find user :" + in.Username)
		return &pb.SsoReply{Code: util.User_unexist}, nil
	}
	//todo check code
	log.Info(in.Username, " start dbclient.RedisClient")
	client := core.RedisClient(6379)
	defer client.Close()
	if in.CodeType == 1 {
		//todo 注册流程
		code, err := client.Do("get", in.Source[:2]+util.SendMobileCode+in.Username)
		if err != nil || code == nil {
			log.Info("Get Code Failed,", err)
			return &pb.SsoReply{Code: util.Code_err}, nil
		}
		coder :=string(code.([]uint8))
		log.Info("code:", coder)
		log.Info("in.Code:", in.Code)
		if in.Code != coder {
			return &pb.SsoReply{Code: util.Code_err}, nil
		}
		return &pb.SsoReply{Code: util.Success}, nil
	} else {
		//todo 找回密码流程
		code, err := client.Do("get", in.Source[:2]+util.SendMobileCode+in.Username)
		if err != nil || code == nil {
			log.Info("Get Code Failed,", err)
			return &pb.SsoReply{Code: 33004}, nil
		}
		coder :=string(code.([]uint8))
		log.Info("code:", coder)
		log.Info("in.Code:", in.Code)
		if in.Code != coder {
			return &pb.SsoReply{Code: util.Code_err}, nil
		}
		token := string(util.Krand(32, util.KC_RAND_KIND_ALL))
		_, err = client.Do("set", in.Source[:2]+util.ResetPasswordByPhone+in.Username, token, "EX", 5*60)
		if err != nil {
			log.Info("Func CheckCode Set Token Failed,", err)
			return &pb.SsoReply{Code: util.System_error}, nil
		}
		return &pb.SsoReply{Code: util.Success, Token: token}, nil
	}
}

//重置密码
func (this *Rpc_sso) ResetPasswordByPhone(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {
	//todo check input
	if in.Source == "" || in.Username == "" || in.Token == "" || in.Password == "" {
		log.Info("Func ResetPasswordByPhone Input is Empty")
		return &pb.SsoReply{Code: util.Params_err_empty}, nil
	}
	// todo check token
	log.Info(in.Username, " start RedisClient")
	client := core.RedisClient(6379)
	defer client.Close()
	token, err := client.Do("get", in.Source[:2]+util.ResetPasswordByPhone+in.Username)
	if token == nil || err != nil {
		log.Info("Get token Failed,", err)
		return &pb.SsoReply{Code: util.System_error}, nil
	}
	log.Info("token:", string(token.([]uint8)))
	log.Info("in.Token:", in.Token)
	if string(token.([]uint8)) != in.Token {
		return &pb.SsoReply{Code: util.Token_err_empty}, nil
	}
	//todo update password
	salt := string(util.Krand(6, util.KC_RAND_KIND_ALL))
	resetPwd := db.Sso{Username: in.Username, Password: util.EncryptWithSalt(in.Password, salt), Salt: salt}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.SsoReply{Code: util.Source_err_empty}, nil
	}
	if err := resetPwd.UpdatePasswordAndSaltByName(); err != nil {
		log.Info("Username:", in.Username, " UpdatePasswordAndSaltByName Failed,", err)
		return &pb.SsoReply{Code: util.System_error}, nil
	}
	return &pb.SsoReply{Code: util.Success}, nil
}

//投诉与建议
func (this *Rpc_sso) Feedback(ctx context.Context, in *pb.SsoRequest) (*pb.SsoReply, error) {

	return nil, nil
}
