package rpc

import (
	"time"

	. "account-domain-rpc/common"

	"account-domain-rpc/storage"

	pb "account-domain-rpc/api/adv"

	"account-domain-rpc/util"

	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
)

type AdvertisementServer struct{}

func (adv *AdvertisementServer) NewAdvertisement(ctx context.Context, in *pb.AdvertisementRequest) (*pb.AdvertisementReply, error) {
	log.Infof("start NewAdvertisement %#v",in)
	if in.Name == "" || in.StartTime == 0 || in.EndTime == 0 || in.Advertiser == "" ||
		in.Md5 == ""  || in.Tid == 0{
		log.Infof("参数异常 name (%s)  id (%d)  StartTime (%d)  AdvUrl (%s) EndTime (%d) Advertiser (%s)   md5 (%s)", in.Name, in.Id, in.StartTime, in.AdvUrl, in.EndTime, in.Advertiser, in.Md5)
		return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	start, err := ChangeTime(in.StartTime)
	if err != nil {
		log.Debugf("参数StartTime转换时间有误 err:", err)
		return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	end, err := ChangeTime(in.EndTime)
	if err != nil {
		log.Debugf("参数EndTime转换时间有误 err:", err)
		return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	a := &storage.Advertisement{
		Id:         in.Id,
		Tid:        in.Tid,
		Name:       in.Name,
		StartTime:  start,
		EndTime:    end,
		Advertiser: in.Advertiser,
		State:      1,
		AdvUrl:     in.AdvUrl,
		FileName:   in.FileName,
		FileUrl:    "http://file.radacat.com:88/v1.0/file/" + in.FileUrl,
		Md5:        in.Md5,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	log.Debugf("insert db before Advertisement:%#v", a)
	if err = a.NewAdvertisement(); err != nil {
		log.Info("增加广告异常 NewAdvertisement err:", err)
		return &pb.AdvertisementReply{ErrorCode: util.System_error}, nil
	}
	return &pb.AdvertisementReply{ErrorCode: util.Successfull}, nil
}

func (adv *AdvertisementServer) UpdateAdvertisement(ctx context.Context, in *pb.AdvertisementRequest) (*pb.AdvertisementReply, error) {
	log.Infof("start UpdateAdvertisement %#v",in)
	if in.Id == 0 || in.Tid == 0{
		log.Infof("参数异常 id (%d) tid (%d)", in.Id, in.Tid)
		return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	if in.FileUrl != "" {
		in.FileUrl  =  "http://file.radacat.com:88/v1.0/file/" + in.FileUrl
	}
	a := &storage.Advertisement{
		Id:         in.Id,
		Tid:        in.Tid,
		Name:       in.Name,
		Md5:        in.Md5,
		Advertiser: in.Advertiser,
		State:      in.State,
		FileName:   in.FileName,
		FileUrl:    in.FileUrl,
		AdvUrl:     in.AdvUrl,
		UpdateTime: time.Now(),
	}
	log.Debug("update db before Advertisement:", a)
	if in.StartTime != 0 {
		start, err := ChangeTime(in.StartTime)
		if err != nil {
			log.Debugf("参数StartTime转换时间有误 err:", err)
			return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
		}
		a.StartTime = start
	}
	if in.EndTime != 0 {
		end, err := ChangeTime(in.EndTime)
		if err != nil {
			log.Debugf("参数StartTime转换时间有误 err:", err)
			return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
		}
		a.EndTime = end
	}
	if err := a.UpdateAdvertisement(); err != nil {
		if err == ErrDoesNotExist {
			return &pb.AdvertisementReply{ErrorCode: util.Advertisement_not_exist}, nil
		}
		log.Info("更新广告异常 UpdateAdvertisement error :", err)
		return &pb.AdvertisementReply{ErrorCode: util.System_error}, nil
	}
	return &pb.AdvertisementReply{ErrorCode: util.Successfull}, nil
}
//todo chatting调的rpc方法,需要对source做处理,转换为租户id
func (adv *AdvertisementServer) GetAdvertisement(ctx context.Context, in *pb.AdvertisementRequest) (*pb.AdvertisementReply, error) {
	log.Infof("start GetAdvertisement %#v",in)
	if in.Source == ""{
		log.Infof("GetAdvertisement input, Source (%s) startTime (%d) ", in.Source,in.StartTime)
		return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	tid,ok := util.SourceToTid[in.Source]
	log.Infof("source (%s)转 tid (%d):", in.Source, tid)
	if  !ok {
		log.Infof("source (%s) to tid (%d) not success  ", in.Source, tid)
		return  &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	a := &storage.Advertisement{}
	if err := a.GetLatestdAdvertisement(tid); err != nil {
		log.Infof("获取单条广告异常 Get advertisement from db error (%s)", err)
		return &pb.AdvertisementReply{ErrorCode: util.System_error}, nil
	}
	if a.Id == 0 {
		return &pb.AdvertisementReply{ErrorCode: util.No_Adv_Can_Be_Find}, nil
	}
	return &pb.AdvertisementReply{
		ErrorCode:  util.Successfull,
		Id:         a.Id,
		Tid:        a.Tid,
		Name:       a.Name,
		Md5:        a.Md5,
		State:      a.State,
		StartTime:  a.StartTime.Unix(),
		EndTime:    a.EndTime.Unix(),
		Advertiser: a.Advertiser,
		AdvUrl:     a.AdvUrl,
		FileName:   a.FileName,
		FileUrl:    a.FileUrl,
	}, nil
}

func (adv *AdvertisementServer) GetAllAdvertisement(ctx context.Context, in *pb.AdvertisementRequest) (*pb.MapAdvertisementReply, error) {
	log.Infof("satrt GetAllAdvertisement %#v",in)
	if in.Page == 0 || in.Count == 0 || in.Tid == 0 {
		log.Infof("参数异常 Page (%d) count (%d) tid (%d)", in.Page, in.Count, in.Tid)
		return &pb.MapAdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	advReply := []*pb.AdvertisementReply{}
	a := &storage.Advertisement{Tid:in.Tid}
	advs, totalCount, err := a.GetAllAdvertisement(in.Page, in.Count)
	if err != nil {
		log.Info("批量获取广告异常 GetAllAdvertisement err,", err)
		return &pb.MapAdvertisementReply{ErrorCode: util.System_error}, nil
	}
	if len(advs) == 0 {
		log.Info("GetAllAdvertisement no advertisement has been find")
		return &pb.MapAdvertisementReply{ErrorCode: util.NO_Advertisement_Can_Be_Find}, nil
	}
	for i := 0; i < len(advs); i++ {
		if advs[i].Id != 0 {
			advReply = append(advReply, &pb.AdvertisementReply{Id: advs[i].Id,Tid:advs[i].Tid,Name: advs[i].Name, Md5: advs[i].Md5, State: advs[i].State,
				StartTime: advs[i].StartTime.Unix(), EndTime: advs[i].EndTime.Unix(), Advertiser: advs[i].Advertiser, AdvUrl: advs[i].AdvUrl,FileName:advs[i].FileName, FileUrl: advs[i].FileUrl})
		}
	}
	return &pb.MapAdvertisementReply{ErrorCode: util.Successfull, TotalCount: totalCount, Advs: advReply}, nil
}

func (adv *AdvertisementServer) DelAdvertisement(ctx context.Context, in *pb.AdvertisementRequest) (*pb.AdvertisementReply, error) {
	log.Infof("start DelAdvertisement tid (%d) id (%d)",in.Tid,in.Id)
	if in.Id == 0 || in.Tid == 0 {
		log.Infof("参数异常id (%d) tid (%d)", in.Id, in.Tid)
		return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	a := &storage.Advertisement{Id: in.Id,Tid:in.Tid}
	err := a.DelAdvertisement()
	if err != nil {
		log.Info("DelAdvertisement err:", err)
		if err == ErrDoesNotExist {
			return &pb.AdvertisementReply{ErrorCode: util.Advertisement_not_exist}, nil
		}
		return &pb.AdvertisementReply{ErrorCode: util.System_error}, nil
	}
	return &pb.AdvertisementReply{ErrorCode: util.Successfull}, nil
}

func (adv *AdvertisementServer) GetOneAdvertisement(ctx context.Context, in *pb.AdvertisementRequest) (*pb.AdvertisementReply, error) {
	log.Infof("start GetOneAdvertisement %#v",in)
	if in.Tid == 0 || in.Id == 0 {
		log.Infof("GetOneAdvertisement input, tid (%s) id (%d) ", in.Tid,in.Id)
		return &pb.AdvertisementReply{ErrorCode: util.Input_parameter_error}, nil
	}
	a := &storage.Advertisement{Tid:in.Tid,Id:in.Id}
	if err := a.GetAdvertisement(); err != nil {
		if err == ErrDoesNotExist{
			return &pb.AdvertisementReply{ErrorCode: util.NO_Advertisement_Can_Be_Find}, nil
		}
		log.Infof("获取单条广告异常 Get advertisement from db error (%s)", err)
		return &pb.AdvertisementReply{ErrorCode: util.System_error}, nil
	}
	return &pb.AdvertisementReply{
		ErrorCode:  util.Successfull,
		Id:         a.Id,
		Tid:        a.Tid,
		Name:       a.Name,
		Md5:        a.Md5,
		State:      a.State,
		StartTime:  a.StartTime.Unix(),
		EndTime:    a.EndTime.Unix(),
		Advertiser: a.Advertiser,
		AdvUrl:     a.AdvUrl,
		FileName:   a.FileName,
		FileUrl:    a.FileUrl,
	}, nil
}

/*ChangeTime   转换时间戳*/
func ChangeTime(inputTime int64) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	toBeCharge := time.Unix(inputTime, 0).Format(layout)
	theTime, err := time.ParseInLocation(layout, toBeCharge, local)
	if err != nil {
		return time.Now(), err
	}
	return theTime, nil
}
