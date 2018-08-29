package service

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"petfone-rpc/core"
)

type FileRpc struct {
}


//添加文件信息
func (this *FileRpc) SetFile(ctx context.Context, req *pb.FilesRequest) (*pb.FilesReply, error) {
	log.Info("SetNotice-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.FilesReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsStr(req.Md5,req.Address) {
		return &pb.FilesReply{Code: util.Params_err_empty}, nil
	}
	//todo 数据库set
	files := &db.FilesPo{Md5:req.Md5, Uid:req.Uid, Address:req.Address, CreationTime:util.GetNowTime()}
	err := files.SetFile()
	if err != nil {
		log.Info("SetNotice-err:", err)
		return &pb.FilesReply{Code: 10001}, nil
	}
	return &pb.FilesReply{Code: 10000}, nil
}

//获取文件信息
func (this *FileRpc) GetFile(ctx context.Context, req *pb.FilesRequest) (*pb.FilesReply, error) {
	log.Info("GetDevices-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.FilesReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsStr(req.GetMd5()) {
		return &pb.FilesReply{Code: util.Params_err_empty}, nil
	}
	files := &db.FilesPo{Md5:req.GetMd5()}
	err := files.GetFile()
	if err != nil {
		if err.Error() == core.ConstStr.NotFound {
			log.Info("err:record not found")
			return &pb.FilesReply{Code: 33013}, nil
		}
		log.Info("err:", err)
		return &pb.FilesReply{Code: 10001}, nil
	}
	log.Info("files:", files)
	return &pb.FilesReply{Code: 10000, Md5:files.Md5, Uid:files.Uid, Address:files.Address, Times:files.CreationTime.Unix()}, nil
}

//批量添加文件信息
func (this *FileRpc) SetFiles(ctx context.Context, req *pb.FilesMapRequest) (*pb.FilesMapReply, error) {
	log.Info("SetNotice-req:", req)
	return &pb.FilesMapReply{Code: 10000}, nil
}

//批量获取文件信息
func (this *FileRpc) GetFiles(ctx context.Context, req *pb.FilesMapRequest) (*pb.FilesMapReply, error) {
	log.Info("GetDevices-req:", req)
	return &pb.FilesMapReply{Code: 10000}, nil
}

//更新品种信息
func (this *FileRpc) SetBreeds(ctx context.Context, req *pb.FilesMapRequest) (*pb.FilesMapReply, error) {
	log.Info("GetDevices-req:", len(req.Files))
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.FilesMapReply{Code: util.Source_err_empty}, nil
	}
	dbc := core.MysqlClient.Begin()
	for _,v := range req.Files {
		breed := &db.BreedInfoPo{Id:v.Id,Info:v.Describe}
		err := breed.UpdateInfoDB("breedinfo",dbc)
		if err != nil {
			dbc.Rollback()
			log.Info("err:", err)
			return &pb.FilesMapReply{Code: 10001}, nil
		}
	}
	dbc.Commit()
	return &pb.FilesMapReply{Code: 10000}, nil
}

//获取品种信息
func (this *FileRpc) GetBreeds(ctx context.Context, req *pb.FilesRequest) (*pb.FilesMapReply, error) {
	log.Info("GetDevices-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.FilesMapReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.Id,req.Types,req.Number) {
		return &pb.FilesMapReply{Code: util.Params_err_empty}, nil
	}
	breeds := &db.BreedInfoPo{Id:req.Id,Types:req.Types}
	breedsRe, err := breeds.GetBreedinfosDB(req.Id+req.Number)
	if err != nil {
		if err.Error() == core.ConstStr.NotFound {
			log.Info("err:record not found")
			return &pb.FilesMapReply{Code: 33013}, nil
		}
		log.Info("err:", err)
		return &pb.FilesMapReply{Code: 10001}, nil
	}
	if len(breedsRe) == 0 {
		log.Info("err:", err)
		return &pb.FilesMapReply{Code: 33013}, nil
	}
	log.Info("breedsRe:", len(breedsRe))
	fileMap := make(map[int32]*pb.FilesReply)
	for k, v := range breedsRe {
		fileMap[int32(k)] = &pb.FilesReply{Id:v.Id, Name:v.NameCh, NameCh:v.NameCh, NameEn:v.NameEn,
		Types:v.Types, Address:v.Address,Describe:v.Info}
	}
	return &pb.FilesMapReply{Code: 10000, Files:fileMap}, nil
}
