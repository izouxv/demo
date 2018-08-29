package rpc

import(


	pb "auth/api"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	."auth/util"
	"auth/storage"
	"time"
)

type ActionLogServer struct {}

func (al *ActionLogServer) AddActionLog(ctx context.Context, in *pb.AddActionLogRequest) (*pb.AddActionLogResponse, error) {
	log.Info("Start AddActionLog")
	log.Info("AddActionLog in:",in)
	if in.ActionLog.ActionUsername == "" || in.ActionLog.ActionTime == 0 || in.ActionLog.ActionType == 0 || in.ActionLog.ActionName == "" || in.ActionLog.ActionObject == "" {
		log.Errorf("AddActionLog Input is empty, ",in)
		return nil, InvalidArgument
	}
	a := storage.ActionLog{
		ActionUsername:in.ActionLog.ActionUsername,
		ActionTime:in.ActionLog.ActionTime,
		ActionType:in.ActionLog.ActionType,
		ActionName:in.ActionLog.ActionName,
		ActionObject:in.ActionLog.ActionObject,
		CreateTime:time.Now(),
		Tid:in.ActionLog.Tid,
		Did:in.ActionLog.Did,
		}
	if err := a.AddActionLog(); err != nil {
		return nil, SystemError
	}
	return nil, Successful
}

func (al *ActionLogServer) GetActionLogsByType(ctx context.Context, in *pb.GetActionLogsByTypeRequest) (*pb.GetActionLogsByTypeResponse, error) {
	log.Info("Start GetActionLogsByType")
	if in.Page == 0 || in.Count == 0 || in.Type == 0 {
		log.Errorf("GetActionLogsByType Input is empty, ",in)
		return nil, InvalidArgument
	}
	a := storage.ActionLog{ActionType:in.Type}
	als, totalCount, err := a.GetActionLogsByType(in.Page, in.Count)
	if err != nil || len(als) == 0 {
		log.Errorf("GetActionLogsByType Error, ",err)
		return nil,NotFind
	}
	var reply  []*pb.ActionLogInfo
	for i:=0;i<len(als);i++ {
		if als[i].Id != 0 {
			reply = append(reply, &pb.ActionLogInfo{
				Id:als[i].Id,
				ActionUsername:als[i].ActionUsername,
				ActionTime:als[i].ActionTime,
				ActionType:als[i].ActionType,
				ActionName:als[i].ActionName,
				ActionObject:als[i].ActionObject,
			})
		}
 	}
	return &pb.GetActionLogsByTypeResponse{TotalCount:totalCount,ActionLogs:reply}, nil
}

func (al *ActionLogServer) GetActionLogsByUsername(ctx context.Context, in *pb.GetActionLogsByUsernameRequest) (*pb.GetActionLogsByUsernameResponse, error) {
	log.Info("Start GetActionLogsByUsername")
	if in.Page == 0 || in.Count == 0 || in.Username == "" {
		log.Errorf("GetActionLogsByUsername Input is empty, ",in)
		return nil, InvalidArgument
	}
	a := storage.ActionLog{ActionUsername:in.Username}
	als, totalCount, err := a.GetActionLogsByUsername(in.Page, in.Count)
	if err != nil || len(als) == 0 {
		log.Errorf("GetActionLogsByUsername Error, ",err)
		return nil, SystemError
	}
	var reply  []*pb.ActionLogInfo
	for i:=0;i<len(als);i++ {
		if als[i].Id != 0 {
			reply = append(reply, &pb.ActionLogInfo{
				Id:als[i].Id,
				ActionUsername:als[i].ActionUsername,
				ActionTime:als[i].ActionTime,
				ActionType:als[i].ActionType,
				ActionName:als[i].ActionName,
				ActionObject:als[i].ActionObject,
			})
		}
	}
	return &pb.GetActionLogsByUsernameResponse{TotalCount:totalCount,ActionLogs:reply}, nil
}

//func (al *ActionLogServer) GetActionLogsByRole(ctx context.Context, in *pb.GetActionLogsByRoleRequest) (*pb.GetActionLogsByRoleResponse, error) {
//	log.Info("Start GetActionLogsByRole")
//	if in.Page == 0 || in.Count == 0 || in.Role == 0 {
//		log.Errorf("GetActionLogsByRole Input is empty, ",in)
//		return &pb.GetActionLogsByRoleResponse{ErrorCode:util.Input_parameter_error}, nil
//	}
//	a := storage.ActionLog{ActionRole:in.Role}
//	als, totalCount, err := a.GetActionLogsByRole(in.Page, in.Count)
//	if err != nil || len(als) == 0 {
//		log.Errorf("GetActionLogsByRole Error, ",err)
//		return &pb.GetActionLogsByRoleResponse{ErrorCode:util.System_error}, nil
//	}
//	var reply  []*pb.ActionLogInfo
//	for i:=0;i<len(als);i++ {
//		if als[i].Id != 0 {
//			reply = append(reply, &pb.ActionLogInfo{
//				Id:als[i].Id,
//				ActionUsername:als[i].ActionUsername,
//				ActionRole:als[i].ActionRole,
//				ActionTime:als[i].ActionTime,
//				ActionType:als[i].ActionType,
//				ActionName:als[i].ActionName,
//				ActionObject:als[i].ActionObject,
//			})
//		}
//	}
//	return &pb.GetActionLogsByRoleResponse{ErrorCode:util.Successfull,TotalCount:totalCount,ActionLogs:reply}, nil
//}

func (al *ActionLogServer) GetActionLogs(ctx context.Context, in *pb.GetActionLogsRequest) (*pb.GetActionLogsResponse, error) {
	log.Info("Start GetAllActionLogs")
	log.Infof("GetAllActionLogs InputInfo:",in)
	if in.Page == 0 || in.Count == 0{
		log.Errorf("GetAllActionLogs Input is empty, ",in)
		return nil, InvalidArgument
	}
	a := storage.ActionLog{Tid:in.Tid,Did:in.Did,ActionType:in.Type,ActionUsername:in.Username,}
	var als []storage.ActionLog
	var totalCount int32
	var err error
		als, totalCount, err = a.GetActionLogs(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	//if in.Type != 0 && in.Username != "" {
	//	a.ActionType = in.Type
	//	a.ActionUsername = in.Username
	//	als, totalCount, err = a.GetActionLogsByUsernameAndType(in.Page, in.Count)
	//	if err != nil || len(als) == 0 {
	//		log.Errorf("GetAllActionLogs Error, ",err)
	//		return nil, NotFind
	//	}
	//}else if in.Type == 0 && in.Username == "" {
	//	als, totalCount, err = a.GetAllActionLogs(in.Page, in.Count)
	//	if err != nil || len(als) == 0 {
	//		log.Errorf("GetAllActionLogs Error, ",err)
	//		return nil, NotFind
	//	}
	//}else if in.Type != 0 && in.Username == "" {
	//	a.ActionType = in.Type
	//	als, totalCount, err = a.GetActionLogsByType(in.Page, in.Count)
	//	if err != nil || len(als) == 0 {
	//		log.Errorf("GetAllActionLogs Error, ",err)
	//		return nil, NotFind
	//	}
	//}else if in.Type == 0 && in.Username != "" {
	//	a.ActionUsername = in.Username
	//	als, totalCount, err = a.GetActionLogsByUsername(in.Page, in.Count)
	//	if err != nil || len(als) == 0 {
	//		log.Errorf("GetAllActionLogs Error, ",err)
	//		return nil, NotFind
	//	}
	//}
	var reply  []*pb.ActionLogInfo
	for i:=0;i<len(als);i++ {
		if als[i].Id != 0 {
			reply = append(reply, &pb.ActionLogInfo{
				Id:als[i].Id,
				ActionUsername:als[i].ActionUsername,
				ActionTime:als[i].ActionTime,
				ActionType:als[i].ActionType,
				ActionName:als[i].ActionName,
				ActionObject:als[i].ActionObject,
			})
		}
	}
	return &pb.GetActionLogsResponse{TotalCount:totalCount,ActionLogs:reply}, nil
}

func (al *ActionLogServer) GetActionLogsByTid(ctx context.Context, in *pb.GetActionLogsByTidRequest) (*pb.GetActionLogsByTidResponse, error) {
	log.Info("Start GetAllActionLogs")
	log.Infof("GetAllActionLogs InputInfo:",in)
	if in.Page == 0 || in.Count == 0{
		log.Errorf("GetAllActionLogs Input is empty, ",in)
		return nil, InvalidArgument
	}
	a := storage.ActionLog{}
	var als []storage.ActionLog
	var totalCount int32
	var err error
	if in.Type != 0 && in.Username != "" {
		a.ActionType = in.Type
		a.ActionUsername = in.Username
		als, totalCount, err = a.GetActionLogsByUsernameAndType(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}else if in.Type == 0 && in.Username == "" {
		als, totalCount, err = a.GetAllActionLogs(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}else if in.Type != 0 && in.Username == "" {
		a.ActionType = in.Type
		als, totalCount, err = a.GetActionLogsByType(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}else if in.Type == 0 && in.Username != "" {
		a.ActionUsername = in.Username
		als, totalCount, err = a.GetActionLogsByUsername(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}
	var reply  []*pb.ActionLogInfo
	for i:=0;i<len(als);i++ {
		if als[i].Id != 0 {
			reply = append(reply, &pb.ActionLogInfo{
				Id:als[i].Id,
				ActionUsername:als[i].ActionUsername,
				ActionTime:als[i].ActionTime,
				ActionType:als[i].ActionType,
				ActionName:als[i].ActionName,
				ActionObject:als[i].ActionObject,
			})
		}
	}
	return &pb.GetActionLogsByTidResponse{TotalCount:totalCount,ActionLogs:reply}, nil
}

func (al *ActionLogServer) GetActionLogsByDid(ctx context.Context, in *pb.GetActionLogsByDidRequest) (*pb.GetActionLogsByDidResponse, error) {
	log.Info("Start GetAllActionLogs")
	log.Infof("GetAllActionLogs InputInfo:",in)
	if in.Page == 0 || in.Count == 0{
		log.Errorf("GetAllActionLogs Input is empty, ",in)
		return nil, InvalidArgument
	}
	a := storage.ActionLog{}
	var als []storage.ActionLog
	var totalCount int32
	var err error
	if in.Type != 0 && in.Username != "" {
		a.ActionType = in.Type
		a.ActionUsername = in.Username
		als, totalCount, err = a.GetActionLogsByUsernameAndType(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}else if in.Type == 0 && in.Username == "" {
		als, totalCount, err = a.GetAllActionLogs(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}else if in.Type != 0 && in.Username == "" {
		a.ActionType = in.Type
		als, totalCount, err = a.GetActionLogsByType(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}else if in.Type == 0 && in.Username != "" {
		a.ActionUsername = in.Username
		als, totalCount, err = a.GetActionLogsByUsername(in.Page, in.Count)
		if err != nil || len(als) == 0 {
			log.Errorf("GetAllActionLogs Error, ",err)
			return nil, NotFind
		}
	}
	var reply  []*pb.ActionLogInfo
	for i:=0;i<len(als);i++ {
		if als[i].Id != 0 {
			reply = append(reply, &pb.ActionLogInfo{
				Id:als[i].Id,
				ActionUsername:als[i].ActionUsername,
				ActionTime:als[i].ActionTime,
				ActionType:als[i].ActionType,
				ActionName:als[i].ActionName,
				ActionObject:als[i].ActionObject,
			})
		}
	}
	return &pb.GetActionLogsByDidResponse{TotalCount:totalCount,ActionLogs:reply}, nil
}
