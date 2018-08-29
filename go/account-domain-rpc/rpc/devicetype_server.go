package rpc

import(
	pb "account-domain-rpc/api"
    log "github.com/cihub/seelog"
     "golang.org/x/net/context"
     "account-domain-rpc/util"
     "account-domain-rpc/storage"
     "account-domain-rpc/module"
     "account-domain-rpc/common"
     )


type DeviceTypeServer struct{}

func (d *DeviceTypeServer) GetDeviceTypes(ctx context.Context, in *pb.GetDeviceTypesRequest) (*pb.GetDeviceTypesReply, error) {
	log.Infof("Start GetDeviceTypes %#v",in)
	if in.Tid == 0{
		log.Infof("输入参数异常 tid:(%d)", in.Tid)
		return &pb.GetDeviceTypesReply{ErrorCode: util.Input_parameter_error}, nil
	}
	dt  := &storage.DeviceType{Tid:in.Tid}
	dts, totalCount, err := dt.GetDeviceTypes(module.MysqlClient(), in.Count, in.Page,in.OrderBy)
	if err != nil {
		if err == common.ErrDoesNotExist{
			return &pb.GetDeviceTypesReply{ErrorCode: util.DeviceType_Does_Not_Exist}, nil
		}
		log.Infof("分页获取设备类型异常:", err)
		return &pb.GetDeviceTypesReply{ErrorCode: util.System_error}, err
	}
	dtRes := make([]*pb.Type,0)
	for _, v := range dts {
		dtRes = append(dtRes, &pb.Type{
			Id:          v.Id,
			Tid:         v.Tid,
			DeviceType:  v.DeviceType,
			Status:      v.Status,
		})
	}
	log.Debug(dtRes)
	return &pb.GetDeviceTypesReply{TotalCount:totalCount,ErrorCode:util.Successfull,DeviceType:dtRes}, nil
}
func (d *DeviceTypeServer) AddDeviceType(ctx context.Context, in *pb.AddDeviceTypeRequest) (*pb.AddDeviceTypeReply, error) {
	log.Infof("Start AddDeviceType %#v",in.DeviceType)
	if in.DeviceType == nil || in.DeviceType.Tid == 0 {
		log.Info("输入参数异常")
		return &pb.AddDeviceTypeReply{ErrorCode: util.Input_parameter_error}, nil
	}
	dt  := &storage.DeviceType{
		Tid:in.DeviceType.Tid,
		DeviceType:in.DeviceType.DeviceType,
		Status:in.DeviceType.Status}
	if err := dt.CreateDeviceType(module.MysqlClient());err != nil {
		log.Infof("添加设备类型异常:", err)
		return &pb.AddDeviceTypeReply{ErrorCode: util.System_error}, err
	}
	return &pb.AddDeviceTypeReply{ErrorCode:util.Successfull,DeviceType:returnDeviceType(dt)}, nil
}

func returnDeviceType (dt *storage.DeviceType) (resp *pb.Type) {
	resp = &pb.Type{
		Id:dt.Id,
		Tid:dt.Tid,
		DeviceType:dt.DeviceType,
		Status:dt.Status,
	}
	return
}