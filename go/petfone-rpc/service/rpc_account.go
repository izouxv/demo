package service

import (
	"fmt"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"net/url"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"petfone-rpc/core"
)

type Rpc_account struct {
}

//基于uid获取用户全部信息（account表）
func (this *Rpc_account) GetAccountInfo(ctx context.Context, in *pb.AccountRequest) (*pb.AccountReply, error) {
	fmt.Println("aaaa:",ctx.Value("test01"))
	if in.Source == "" || in.Uid == 0 {
		log.Info("Input id is null! in.Uid:", in.Uid)
		return &pb.AccountReply{Code: util.Params_err_empty}, nil
	}
	account := &db.Account{Id: in.Uid}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.AccountReply{Code: util.Source_err_empty}, nil
	}
	err := account.GetAccountById()
	if err != nil || account.Nickname == "" {
		log.Info("can not find the account:", in.Uid)
		return &pb.AccountReply{Uid: in.Uid, Code: util.User_unexist}, nil
	}
	signature, _ := url.QueryUnescape(account.Signature)
	nickname, _ := url.QueryUnescape(account.Nickname)
	petfonePo := &db.PetfonePo{Uid: in.Uid, DataState: 1}
	if in.Source == "AgIDAA==" {
		err = petfonePo.GetPetfoneDB()
		if err != nil {
			log.Error("GetAccountInfo GetPetfoneDB err", err)
			return &pb.AccountReply{Code:10000, Uid:account.Id, Nickname:nickname, Gender:account.Gender,
				Birthday:account.Birthday.Unix(), Avatar:core.ConstStr.FileServer+account.Avatar, Signature:signature,
				Address:account.Address, Phone:account.Phone, Email: account.Email}, nil
		}
	}
	log.Info(" GetAccountInfo successful")
	log.Info("### petfonePo = ",petfonePo)
	return &pb.AccountReply{Code:10000, Uid:account.Id, Nickname:nickname, Gender:account.Gender,
		Birthday:account.Birthday.Unix(), Avatar:core.ConstStr.FileServer+account.Avatar, Signature:signature,
		Address:account.Address, Phone:account.Phone, Email: account.Email, Radius:petfonePo.Radius,Map:pb.Map(petfonePo.Map)}, nil
}

//基于uid获取用户基本信息(account表)
func (this *Rpc_account) Show(ctx context.Context, in *pb.AccountRequest) (*pb.AccountReply, error) {
	if in.Source == "" || in.Uid == 0 {
		log.Info("Input id is null! in.Uid:", in.Uid)
		return &pb.AccountReply{Code: util.Params_err_empty}, nil
	}
	a := &db.Account{Id: in.Uid}
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.AccountReply{Code: util.Source_err_empty}, nil
	}
	err := a.GetAccountById()
	if err != nil || a.Nickname == "" {
		log.Info("can not find the account:", in.Uid)
		return &pb.AccountReply{Uid: in.Uid, Code: util.User_unexist}, nil
	}
	signature, _ := url.QueryUnescape(a.Signature)
	nickname, _ := url.QueryUnescape(a.Nickname)
	//fmt.Println(signature)
	return &pb.AccountReply{Code: util.Success,
		Uid:       a.Id,
		Nickname:  nickname,
		Gender:    a.Gender,
		Birthday:  a.Birthday.Unix(),
		Avatar:    core.ConstStr.FileServer+a.Avatar,
		Signature: signature,
		Address:   a.Address,
		Phone:     a.Phone,
		Email:     a.Email}, nil
}

//基于uid插入或修改用户基本信息(account表)
func (this *Rpc_account) UpdateAccountInfo(ctx context.Context, in *pb.AccountRequest) (*pb.AccountReply, error) {
	log.Info("UpdateAccountInfo-req:",in)
	if in.Source == "" || in.Token == "" || util.CheckSource(in.Source) {
		log.Info("Input Source is null! in.Uid:", in.Uid)
		return &pb.AccountReply{Code: util.Source_err_empty}, nil
	}
	if in.Uid <= 0 {
		log.Info("Input id is null! in.Uid:", in.Uid)
		return &pb.AccountReply{Code: util.Params_err_empty}, nil
	}
	account := &db.Account{Id: in.Uid, Nickname: url.QueryEscape(in.Nickname), Gender: in.Gender, Avatar: in.Avatar,
		Signature: url.QueryEscape(in.Signature), Address: in.Address, Phone: in.Phone, Email: in.Email, UpdateTime:util.GetNowTime()}
	if in.Birthday > 0 {
		birthday, err := util.Int64ToTime(in.Birthday)
		if err != nil {
			log.Info("Change Time Failed,", err)
			return &pb.AccountReply{Code: util.User_params_err}, nil
		}
		account.Birthday = birthday
	}
	dbc := core.MysqlClient.Begin()
	//todo 插入数据库
	if account.Id == 0 {
		if errInsert := account.InsertAccount(); errInsert != nil {
			log.Info("err InsertAccount:", errInsert)
			dbc.Rollback()
			return &pb.AccountReply{Code: util.Mysql_err}, nil
		}
		log.Info("InsertAccountInfo Successful! Uid:", account.Id)
		return &pb.AccountReply{Code: util.Success}, nil
	} else {
		if err := account.UpdateAccount(dbc); err != nil {
			log.Info("err UpdateAccount:", err)
			dbc.Rollback()
			return &pb.AccountReply{Code: util.Mysql_err}, nil
		}
	}
	//todo 更新昵称时，更新redis信息
	if in.Nickname != "" {
		sso := &db.Sso{Id: in.Uid}
		err := sso.GetUserInfoById()
		value, errJson := util.Json.Marshal(&pb.SsoReply{Uid: sso.Id, Username: sso.Username, Nickname: sso.Nickname, State: sso.State, LoginState: login})
		if errJson != nil {
			dbc.Rollback()
			log.Info("err json.Marshal value,", errJson)
			return &pb.AccountReply{Code: util.Json_error}, nil
		}
		client := core.RedisClient(6379)
		defer client.Close()
		reply, err := client.Do("set", in.Source[:2]+util.LoginSession+in.Token, value)
		if err != nil || reply == 0 {
			dbc.Rollback()
			log.Info("Get userInfo Failed,", err)
			return &pb.AccountReply{Code: util.Token_err_empty}, nil
		}
	}
	log.Info("UpdateAccountInfo Successful! Uid:", account.Id)
	dbc.Commit()
	if in.Avatar != "" {
		return &pb.AccountReply{Code: util.Success,Avatar:core.ConstStr.FileServer+account.Avatar}, nil
	}
	return &pb.AccountReply{Code: util.Success}, nil
}

//批量查询用户基本信息
func (this *Rpc_account) GetBatchAccountInfo(ctx context.Context, in *pb.MultiAccountRequest) (*pb.MapAccountReply, error) {
	if in.Source == "" || len(in.Accounts) == 0 {
		log.Info("Input id is null! in.Uid:")
		return &pb.MapAccountReply{Code: util.Params_err_empty}, nil
	}
	var uids []int32 = make([]int32, 0)
	for k := range in.Accounts {
		if k != 0 {
			uids = append(uids, k)
		}
	}
	log.Info(uids)
	log.Info("Uids length:", len(uids))
	userMap := make(map[int32]*pb.AccountReply)
	var a db.Account
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.MapAccountReply{Code: util.Source_err_empty}, nil
	}
	rs, err := a.GetBatchAccount(uids)
	if err != nil || rs == nil {
		log.Info("errGetBatchAccountInfo")
		return &pb.MapAccountReply{Code: util.User_unexist}, nil
	}
	for i := 0; i < len(rs); i++ {
		if rs[i].Id != 0 {
			birthday := rs[i].Birthday.Unix() * 1000
			signature, _ := url.QueryUnescape(rs[i].Signature)
			nickname, _ := url.QueryUnescape(rs[i].Nickname)
			userMap[rs[i].Id] = &pb.AccountReply{Uid: rs[i].Id, Gender: rs[i].Gender, Birthday: birthday, Avatar: core.ConstStr.FileServer+rs[i].Avatar,
				Signature: signature, Nickname: nickname, Address: rs[i].Address, Phone: rs[i].Phone, Email: rs[i].Email}
		}
	}
	return &pb.MapAccountReply{Code: util.Success, Accounts: userMap}, nil
}

//批量查询用户所有信息
func (this *Rpc_account) GetBatchAllUserInfo(ctx context.Context, in *pb.MultiAccountRequest) (*pb.MapAccountReply, error) {
	var uids []int32 = make([]int32, 0)
	if len(in.Accounts) == 0 {
		log.Info("in.in.Accounts is empty")
		return &pb.MapAccountReply{Code: util.Params_err_empty}, nil
	}
	for k := range in.Accounts {
		if k != 0 {
			uids = append(uids, k)
		}
	}
	log.Info(uids)
	log.Info("Uids length:", len(uids))
	userMap := make(map[int32]*pb.AccountReply)
	var a db.Account
	if util.CheckSource(in.Source) {
		log.Info("Source_is_incorrent_or_empty, in.Source:", in.Source)
		return &pb.MapAccountReply{Code: util.Source_err_empty}, nil
	}
	rs, err := a.GetBatchAccount(uids)
	if err != nil || rs == nil {
		log.Info("errGetBatchAccountInfo")
		return &pb.MapAccountReply{Code: util.User_unexist}, nil
	}

	fmt.Println(rs)
	for i := 0; i < len(rs); i++ {
		if rs[i].Id != 0 {
			birthday := rs[i].Birthday.Unix() * 1000
			signature, _ := url.QueryUnescape(rs[i].Signature)
			nickname, _ := url.QueryUnescape(rs[i].Nickname)
			userMap[rs[i].Id] = &pb.AccountReply{Uid: rs[i].Id, Gender: rs[i].Gender, Birthday: birthday, Avatar: core.ConstStr.FileServer+rs[i].Avatar,
				Signature: signature, Nickname: nickname,
				Address: rs[i].Address, Phone: rs[i].Phone, Email: rs[i].Email}
		}
	}
	fmt.Println(userMap)
	return &pb.MapAccountReply{Code: util.Success, Accounts: userMap}, nil
}
