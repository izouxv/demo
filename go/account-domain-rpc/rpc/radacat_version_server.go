package rpc

import (
	"golang.org/x/net/context"
	pb "account-domain-rpc/api/setting"
	"account-domain-rpc/storage"
	"account-domain-rpc/util"
	"account-domain-rpc/module"
	log "github.com/cihub/seelog"
	"time"
)

type RadacatVersionServer struct{}


//添加新版本
func (r *RadacatVersionServer) AddNewVersion(ctx context.Context, in *pb.AddNewVersionRequest) (*pb.AddNewVersionResponse, error) {
	log.Info("	Start AddNewVersion %#v", in)
	if in.Version.Device == "" || in.Version.Tid == 0 || in.Version.Md5 == "" || in.Version.Path == "" || in.Version.Length == 0 || in.Version.Status == 0{
		log.Infof("参数异常 device (%s) tid (%d) md5 (%s) path (%s) length (%d) status (%d) ", in.Version.Device, in.Version.Tid, in.Version.Md5, in.Version.Path, in.Version.Length, in.Version.Status)
		return &pb.AddNewVersionResponse{ErrorCode: util.Input_parameter_error}, nil
	}
	rv := &storage.RadacatVersion{
		Tid:           in.Version.Tid,
		Device:        in.Version.Device,
		VersionName:   in.Version.VersionName,
		VersionCode:   in.Version.VersionCode,
		FileName:      in.Version.FileName,
		FileLength:    in.Version.Length,
		MD5:           in.Version.Md5,
		URL:           in.Version.Path,
		DescriptionCN: in.Version.DescriptionCn,
		DescriptionEN: in.Version.DescriptionEn,
		Status:        in.Version.Status,
	}
	version,_ := rv.GetVersionByVersionCode()
	if version != nil {
		log.Info("版本已存在")
		return &pb.AddNewVersionResponse{ErrorCode: util.Version_is_exist}, nil
	}
	if err := rv.Create(); err != nil {
		log.Errorf("AddNewVersion Create Error: %s", err)
		return &pb.AddNewVersionResponse{ErrorCode: util.System_error}, nil
	}
	return &pb.AddNewVersionResponse{ErrorCode: util.Successfull}, nil
}

//查询所有版本
func (r *RadacatVersionServer) GetAllVersions(ctx context.Context, in *pb.GetAllVersionsRequest) (*pb.GetAllVersionsResponse, error) {
	log.Info("Start GetAllVersions %#v", in)
	if in.Count == 0 || in.Page == 0 || in.Tid == 0 {
		log.Infof("参数异常 count: %s, page: %s tid : %d", in.Count, in.Page, in.Tid)
		return &pb.GetAllVersionsResponse{ErrorCode: util.Input_parameter_error}, nil
	}
	rv := &storage.RadacatVersion{Tid:in.Tid}
	rvs, totalCount, err := rv.GetAllVersions(in.Page, in.Count)
	log.Info("版本：",rvs,"总数：",totalCount)
	if err != nil {
		log.Error("GetAllVersions Error: ", err)
		return &pb.GetAllVersionsResponse{ErrorCode: util.System_error}, nil
	}

	//if len(rvs) == 0 {
	//	log.Info("未找到版本")
	//	return &pb.GetAllVersionsResponse{ErrorCode: util.NO_RadacatVersion_Can_Be_Find}, nil
	//}
	var rs []*pb.Version
	for _, v := range rvs {
		if v.Id != 0 {
			rs = append(rs, &pb.Version{
				Id:            v.Id,
				Device:        v.Device,
				VersionName:   v.VersionName,
				VersionCode:   v.VersionCode,
				FileName:      v.FileName,
				Length:        v.FileLength,
				Md5:           v.MD5,
				Path:          v.URL,
				DescriptionCn: v.DescriptionCN,
				DescriptionEn: v.DescriptionEN,
				Status:        v.Status,
				CreateTime:    v.CreateTime.Unix(),
				UpdateTime:    v.UpdateTime.Unix(),
			})
		}
	}
	return &pb.GetAllVersionsResponse{ErrorCode: util.Successfull, Versions: rs, TotalCount: totalCount}, nil
}

//通过device查询最新版本信息         todo:该接口由chatting调用,source转tid查询
func (r *RadacatVersionServer) GetLatestVersion(ctx context.Context, in *pb.GetLatestVersionRequest) (*pb.GetLatestVersionResponse, error) {
	log.Infof("Start GetLatestVersion %#v",in)
	if in.Device == "" || in.Source == "" {
		log.Infof("参数异常 in.Device(%s) source (%s) Username(%s) ", in.Device, in.Source,in.Username)
		return &pb.GetLatestVersionResponse{ErrorCode: util.Input_parameter_error}, nil
	}
	if _, ok := util.SourceToTid[in.Source]; !ok {
		log.Infof("source 有误 source (%s) ", in.Source)
		return &pb.GetLatestVersionResponse{ErrorCode: util.Input_parameter_error}, nil
	}
	rv := &storage.RadacatVersion{Device: in.Device,Tid: util.SourceToTid[in.Source]}
	tu := &storage.TestUser{Tid:util.SourceToTid[in.Source],UserName:in.Username}
	if err := tu.GetTestUserByUsername(module.MysqlClient());err != nil {
		log.Debug("测试账号中是否存在此账号",err)
	switch err{
		/*case storage.TestUserAlreadyExists:
			log.Info("是测试账号,返回现在最新版本")
			if err := rv.GetLatestVersion();err != nil {
				log.Debug("查询测试账号本版 err:",err)
				if err == storage.RadacatVersionDoesNotExist{
					return &pb.GetLatestVersionResponse{ErrorCode: util.NO_RadacatVersion_Can_Be_Find}, nil
				}
			}
			log.Debug("返回测试账号版本:",ReturnVersion(rv))
			return &pb.GetLatestVersionResponse{ErrorCode: util.Successfull, Version:ReturnVersion(rv)},  nil*/
	   case storage.TestUserNotExists:
		   log.Info("不是测试账号,返回release最新版本")
		   if err := rv.GetLatestVersionRelease();err != nil {
			   log.Debug("查询非测试账号本版 err:",err)
			   if err == storage.RadacatVersionDoesNotExist{
				   return &pb.GetLatestVersionResponse{ErrorCode: util.NO_RadacatVersion_Can_Be_Find}, nil
			   }
			   log.Debug("返回非测试账号版本:",ReturnVersion(rv))
			   return &pb.GetLatestVersionResponse{ErrorCode: util.System_error}, nil
		   }
		   return &pb.GetLatestVersionResponse{ErrorCode: util.Successfull, Version: ReturnVersion(rv)},  nil
	    default :
		    log.Info("获取测试账号有误：",err)
			return &pb.GetLatestVersionResponse{ErrorCode: util.System_error}, nil
	   }
	}
	log.Info("是测试账号,返回现在最新版本")
	if err := rv.GetLatestVersion();err != nil {
		log.Debug("查询测试账号本版 err:",err)
		if err == storage.RadacatVersionDoesNotExist{
			return &pb.GetLatestVersionResponse{ErrorCode: util.NO_RadacatVersion_Can_Be_Find}, nil
		}
	}
	log.Debug("返回测试账号版本:",ReturnVersion(rv))
	return &pb.GetLatestVersionResponse{ErrorCode: util.Successfull, Version:ReturnVersion(rv)},  nil
}

//通过id修改版本信息
func (r *RadacatVersionServer) UpdateVersion(ctx context.Context, in *pb.UpdateVersionRequest) (*pb.UpdateVersionResponse, error) {
	log.Infof("Start UpdateVersion %#v", in)
	if in.Version.Id == 0 || in.Version.Tid == 0 {
		log.Infof("参数有误 Id (%d) tid (%d) ", in.Version.Id, in.Version.Tid)
		return &pb.UpdateVersionResponse{ErrorCode: util.Input_parameter_error}, nil
	}
	rv := &storage.RadacatVersion{
		Id:            in.Version.Id,
		Tid:           in.Version.Tid,
		Device:        in.Version.Device,
		VersionName:   in.Version.VersionName,
		VersionCode:   in.Version.VersionCode,
		DescriptionCN: in.Version.DescriptionCn,
		DescriptionEN: in.Version.DescriptionEn,
		Status:        in.Version.Status,
		UpdateTime:    time.Now(),
	}
    /*if in.Version.FileName != "" {
		rv.FileName = in.Version.FileName
		rv.URL = "http://file.radacat.com:88/v1.0/file/" + in.Version.FileName
	}*/
	log.Debug("修改的device:",in.Version.Device)
	version,_ := rv.GetVersionByVersionCodePut()
	if version != nil {
		log.Info("版本已存在")
		return &pb.UpdateVersionResponse{ErrorCode: util.Version_is_exist}, nil
	}
	if err := rv.Update(); err != nil {
		log.Error("UpdateVersion Error: ", err)
		return &pb.UpdateVersionResponse{ErrorCode: util.System_error}, nil
	}
	return &pb.UpdateVersionResponse{ErrorCode: util.Successfull}, nil
}

//删除版本信息
func (r *RadacatVersionServer) DeleteVersion(ctx context.Context, in *pb.DeleteVersionRequest) (*pb.DeleteVersionResponse, error) {
	log.Infof("Start DeleteVersion %#v", in)
	if in.Id == 0 || in.Tid == 0 {
		log.Infof("参数异常 Id (%s) tid (%d) ", in.Id, in.Tid)
		return &pb.DeleteVersionResponse{ErrorCode: util.Input_parameter_error}, nil
	}
	rv := &storage.RadacatVersion{Id:in.Id,Tid:in.Tid}
	if err := rv.DeleteVersion(); err != nil {
		log.Error("DeleteVersion Error:", err)
		return &pb.DeleteVersionResponse{ErrorCode: util.NO_RadacatVersion_Can_Be_Find}, nil
	}
	return &pb.DeleteVersionResponse{ErrorCode: util.Successfull}, nil
}

//获取版本信息
func (r *RadacatVersionServer) GetVersion(ctx context.Context, in *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	log.Infof("Start GetVersion %#v", in)
	if in.Id == 0 || in.Tid == 0 {
		log.Infof("参数异常 Id (%s) tid (%d) ", in.Id, in.Tid)
		return &pb.GetVersionResponse{ErrorCode: util.Input_parameter_error}, nil
	}
	rv := &storage.RadacatVersion{Id:in.Id,Tid:in.Tid}
	version ,err := rv.GetVersion()
	if err != nil  {
		log.Error("GetVersion Error:", err)
		return &pb.GetVersionResponse{ErrorCode: util.NO_RadacatVersion_Can_Be_Find}, nil
	}
	return &pb.GetVersionResponse{ErrorCode: util.Successfull,Version:&pb.Version{
		Id:version.Id,
		Tid:version.Tid,
		Device:version.Device,
		FileName:version.FileName,
		VersionName:version.VersionName,
		VersionCode:version.VersionCode,
		DescriptionCn:version.DescriptionCN,
		DescriptionEn:version.DescriptionEN,
		CreateTime:version.CreateTime.Unix(),
		UpdateTime:version.UpdateTime.Unix(),
		Path:version.URL,
		Status:version.Status,
	}}, nil
}

func ReturnVersion (rv *storage.RadacatVersion) (v *pb.Version){
	v = &pb.Version{
		Id:            rv.Id,
		Device:        rv.Device,
		Md5:           rv.MD5,
		FileName:      rv.FileName,
		VersionName:   rv.VersionName,
		VersionCode:   rv.VersionCode,
		Path:          rv.URL,
		Length:        rv.FileLength,
		DescriptionCn: rv.DescriptionCN,
		DescriptionEn: rv.DescriptionEN,
		CreateTime:    rv.CreateTime.Unix(),
		UpdateTime:    rv.UpdateTime.Unix(),
		Status:        rv.Status}
	return  v
}
