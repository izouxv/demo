package rpc

import (
	pb "account-domain-rpc/api/user"
	"account-domain-rpc/module"
	"account-domain-rpc/storage"
	"account-domain-rpc/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"time"
)

type TestUserServer struct{}

func (fd *TestUserServer) AddTestUser(ctx context.Context, in *pb.AddTestUserReq) (*pb.AddTestUserRes, error) {
	log.Infof("Start AddTestUser %#v",in)
	if in.TestUser.Tid == 0 || in.TestUser.Username == "" {
		log.Infof("输入参数异常 tid:(%d) username :(%s)", in.TestUser.Tid, in.TestUser.Username)
		return &pb.AddTestUserRes{ErrorCode: util.Input_parameter_error}, nil
	}
	tu := &storage.TestUser{
		Tid: in.TestUser.Tid,
		UserName:  in.TestUser.Username}
	testUser ,_:=  tu.CheckTestUserIsExist(module.MysqlClient())
	if testUser != nil {
		log.Info("白名单账号已存在")
		return &pb.AddTestUserRes{ErrorCode: util.TestUser_is_exist}, nil
	}
	if err := tu.CreateTestUser(module.MysqlClient()); err != nil {
		log.Infof("创建测试账号异常,error:(%s),testuser :(%s)", err, tu)
		return &pb.AddTestUserRes{ErrorCode: util.System_error}, nil
	}
	log.Debug("测试账号信息:", tu)
	return &pb.AddTestUserRes{ErrorCode: util.Successfull,TestUser:returnTserUser(tu)}, nil
}

func (fd *TestUserServer) DelTestUser(ctx context.Context, in *pb.DelTestUserReq) (*pb.DelTestUserRes, error) {
	log.Infof("Start DelTestUser %#v", in)
	if in.Id == 0 || in.Tid == 0 {
		log.Infof("输入参数异常 tid:(%s),id :(%d)", in.Tid, in.Id)
		return &pb.DelTestUserRes{ErrorCode: util.Input_parameter_error}, nil
	}
	tu := &storage.TestUser{Id:  in.Id, Tid:in.Tid}
	if err := tu.DelTestUser(module.MysqlClient()); err != nil {
		log.Infof("删除测试账号异常,error:(%s), testuser :(%#v)", err, tu)
		return &pb.DelTestUserRes{ErrorCode: util.System_error}, err
	}
	return &pb.DelTestUserRes{ErrorCode: util.Successfull}, nil
}

func (fd *TestUserServer) GetTestUsers(ctx context.Context, in *pb.GetTestUsersReq) (*pb.GetTestUsersRes, error) {
	log.Infof("Start GetTestUsers %#v",in)
	if in.Tid == 0 {
		log.Infof("输入参数异常 page :(%s),count :(%s),tid :(%d)", in.Page, in.Count,in.Tid)
		return &pb.GetTestUsersRes{ErrorCode: util.Input_parameter_error}, nil
	}
	tu := &storage.TestUser{Tid:in.Tid}
	tus, totalCount, err := tu.GetTestUsers(module.MysqlClient(), in.Count, in.Page)
	if err != nil {
		log.Infof("分页获取测试账号异常.error is (%s),page is (%d),count is (%d)  ", err, in.Page, in.Count)
		return &pb.GetTestUsersRes{}, err
	}
	tuResp := make([]*pb.TestUser, 0)
	for _, v := range tus {
		tuResp = append(tuResp, returnTserUser(v))
		}
	log.Debug("返回数据:",tuResp)
	return &pb.GetTestUsersRes{ErrorCode: util.Successfull, TestUser:tuResp, TotalCount: totalCount}, nil
}

func (fd *TestUserServer) PutTestUser(ctx context.Context, in *pb.PutTestUserReq) (*pb.PutTestUserRes, error) {
	log.Infof("Start PutTestUser %#v", in)
	if in.TestUser == nil || in.TestUser.Tid == 0 || in.TestUser.Username == "" {
		log.Infof("输入参数错误 %#v", in)
		return &pb.PutTestUserRes{ErrorCode: util.Input_parameter_error}, nil
	}
	tu := &storage.TestUser{Tid:in.TestUser.Tid,Id:in.TestUser.Id,UserName:in.TestUser.Username,UpdateTime:time.Now()}
	testUser ,_:=  tu.CheckTestUserIsExistPut(module.MysqlClient())
	if testUser != nil {
		log.Info("白名单账号已存在")
		return &pb.PutTestUserRes{ErrorCode: util.TestUser_is_exist}, nil
	}
	if errs := tu.UpdateTestUser(module.MysqlClient());errs != nil {
		log.Errorf("修改测试账号异常,error:(%s)", errs)
		return &pb.PutTestUserRes{ErrorCode: util.System_error}, nil
	}
	return &pb.PutTestUserRes{ErrorCode: util.Successfull,TestUser:returnTserUser(tu)}, nil
}

func (fd *TestUserServer) GetTestUser(ctx context.Context, in *pb.GetTestUserReq) (*pb.GetTestUserRes, error) {
	log.Infof("Start GetTestUser %#v", in)
	if in.Id == 0 || in.Tid == 0 {
		log.Infof("输入参数异常 tid:(%s),id :(%d)", in.Tid, in.Id)
		return &pb.GetTestUserRes{ErrorCode: util.Input_parameter_error}, nil
	}
	tu := &storage.TestUser{Id: in.Id, Tid:in.Tid}
	if err := tu.GetTestUser(module.MysqlClient()); err != nil {
		if err == storage.TestUserNotExists{
			log.Info("获取单个测试账号不存在")
			return &pb.GetTestUserRes{ErrorCode: util.TestUser_not_exist}, nil
		}
		log.Infof("获取单个测试账号异常,error:(%s)", err)
		return &pb.GetTestUserRes{ErrorCode: util.System_error}, err
	}
	return &pb.GetTestUserRes{ErrorCode: util.Successfull,TestUser:returnTserUser(tu)}, nil
}

func (fd *TestUserServer) GetUserByUsername(ctx context.Context, in *pb.GetUserByUsernameReq) (*pb.GetUserByUsernameRes, error) {
	log.Infof("Start GetUserByUsername %#v", in)
	if in.Tid == 0 || in.Username == "" {
		log.Infof("输入参数异常 tid:(%s),username :(%s)", in.Tid, in.Username)
		return &pb.GetUserByUsernameRes{ErrorCode: util.Input_parameter_error}, nil
	}
	tu := &storage.TestUser{Tid:in.Tid, UserName:in.Username}
	if err := tu.GetTestUserByUsername(module.MysqlClient()); err != nil {
		if err == storage.TestUserNotExists{
			return &pb.GetUserByUsernameRes{ErrorCode: util.TestUser_not_exist}, nil
		}
		log.Infof("用户名获取测试账号异常,error:(%s)", err)
		return &pb.GetUserByUsernameRes{ErrorCode: util.System_error}, err
	}
	return &pb.GetUserByUsernameRes{ErrorCode: util.Successfull}, nil
}

func returnTserUser(user *storage.TestUser)(tu *pb.TestUser){
	tu = &pb.TestUser{
		Id:user.Id,
		Tid:user.Tid,
		Username:user.UserName,
		CreateTime:user.CreateTime.Unix(),
		UpdateTime:user.UpdateTime.Unix(),
	}
	return
}