package service

import (
	"golang.org/x/net/context"
	log "github.com/cihub/seelog"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	. "petfone-rpc/util"
	"petfone-rpc/core"
	"time"
	"sort"
	"strings"
	"github.com/jinzhu/gorm"
)

/**
运动数据信息
 */
type MotionDataRpc struct {
}

type PCoordinateSlice []*pb.Pcoordinate
func (c PCoordinateSlice) Len() int {return len(c)}
func (c PCoordinateSlice) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c PCoordinateSlice) Less(i, j int) bool { return c[i].Nowtime < c[j].Nowtime }

//存储运动数据
func (m *MotionDataRpc) SetExerciseData(ctx context.Context, req *pb.ExerciseDataRequest) (*pb.ExerciseDataReply, error) {
	log.Info("SetExerciseData-req source:", req.Source,",uid:",req.Uid,",pid:",req.Pid,",pdid:",req.Pdid,
		",len:",len(req.GetPcoordinates()))
	if VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.ExerciseDataReply{Code: Source_err_empty}, nil
	}
	if VerifyParamsUInt32(req.GetUid(),req.GetPid(),req.GetPdid()) {
		return &pb.ExerciseDataReply{Code: 33001}, nil
	}
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.ExerciseDataReply{Code: 33012}, nil
		}
		log.Error("SetExerciseData GetUserPetDB err:", err)
		return &pb.ExerciseDataReply{Code: 10001}, nil
	}
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, Did:req.Pdid, DataState:1}
	err = userDevicePo.GetUserDeviceDB()
	log.Info("SetExerciseData-userDevicePo:", userDevicePo)
	if err != nil {
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.ExerciseDataReply{Code: 33012}, nil
		}
		log.Error("SetExerciseData GetUserDeviceDB-err:", err)
		return &pb.ExerciseDataReply{Code: 10001}, nil
	}
	var pCoordinateSlice PCoordinateSlice = req.Pcoordinates
	sort.Stable(pCoordinateSlice)
	coordinates, err := Json.MarshalToString(pCoordinateSlice)
	if err != nil {
		log.Error("SetExerciseData Marshal-err:", err)
		return &pb.ExerciseDataReply{Code: 33006}, nil
	}
	tableId ,err := db.GenerateBigId()
	if err != nil {
		log.Error("SetExerciseData GenerateBigId err:", err)
		return &pb.ExerciseDataReply{Code: 10001}, nil
	}
	nowTime := GetNowTime()
	var reportTime time.Time
	if req.StartTime == 0 {
		reportTime = nowTime
	} else {
		reportTime, _ = Int64ToTime(req.StartTime)
	}
	startTime, _ := Int64ToTime(req.Pcoordinates[0].Nowtime)
	endTime, _ := Int64ToTime(req.Pcoordinates[len(pCoordinateSlice)-1].Nowtime)
	go func(str string,reportTime,nowTime,startTime,endTime time.Time,r *pb.ExerciseDataRequest) {
		remStr := Int32ToRemStr(req.Pid)
		if remStr == "" {
			log.Error("SetExerciseData RemStr:",req.Pid)
			return
		}
		exist := &db.MotionDataPetPo{Pid:req.Pid,Url:req.ImageInfo.Url,StartTime:startTime,EndTime:endTime}
		err = exist.GetMotionDataPetExist(remStr)
		if err != nil && err.Error() != core.ConstStr.NotFound {
			log.Error("SetExerciseData GetMotionDataPetExist err:",err)
			return
		}
		if exist.Id != 0 {
			log.Info("SetExerciseData GetMotionDataPetExist id:",exist.Id)
			return
		}
		if req.ImageInfo != nil {
			motionDataPetPo := &db.MotionDataPetPo{Id:tableId,Uid:req.Uid,Pid:req.Pid,Pdid:req.Pdid,ReportTime:reportTime, StartTime:startTime,
				EndTime:endTime,Calorie:req.Pcals, Url:req.ImageInfo.Url, Width:req.ImageInfo.Width,Height:req.ImageInfo.Height,
				Size:req.ImageInfo.Size,Name:req.ImageInfo.Name,CreationTime:nowTime,DataState:1}
			for kc, vc := range pCoordinateSlice {
				motionDataPetPo.Steps += vc.Steps
				if kc == 0 {
					continue
				}
				if vc.State == 1 {
					motionDataPetPo.StrenuousTimes += pCoordinateSlice[kc].Nowtime - pCoordinateSlice[kc-1].Nowtime
				} else {
					motionDataPetPo.CardioTimes += pCoordinateSlice[kc].Nowtime - pCoordinateSlice[kc-1].Nowtime
				}
			}
			err = motionDataPetPo.SetMotionDataPet(remStr)
			if err != nil {
				log.Error("SetExerciseData SetMotionDataPet err:", err)
			}
		}
		exerp := &db.ExerciseDataPetPo{Id:tableId,Pid:r.Pid,Pdid:r.Pdid,Uid:r.Uid,Cals:r.Pcals,ReportTime:reportTime,
			Coordinates:str,CreationTime:nowTime,DataState: 1}
		exeru := &db.ExerciseDataUserPo{Id:tableId,Udid:r.Udid,Uid:r.Uid, Steps:r.Usteps,Cals:r.Ucals, CreationTime:nowTime,DataState:1}
		if db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
				exerp.SetExerciseDataPet(tx)
				exeru.SetExerciseDataUser(tx)
				return nil
		}) != nil {
			log.Error("SetExerciseData SetExerciseData err", err)
		}
	}(coordinates,reportTime,nowTime,startTime,endTime,req)
	return &pb.ExerciseDataReply{Code:10000}, nil
}

//查询运动数据1.0
func (m *MotionDataRpc) GetExerciseDataPet(ctx context.Context, req *pb.ExerciseDataRequest) (*pb.MotionDataMapReply, error) {
	log.Info("GetExerciseDataPet-req:", req)
	if VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.MotionDataMapReply{Code: Source_err_empty}, nil
	}
	if VerifyParamsUInt32(req.GetPid()) {
		return &pb.MotionDataMapReply{Code: Source_err_empty}, nil
	}
	exer := &db.ExerciseDataPetPo{Pid:req.GetPid(), DataState:1}
	exers, err := exer.GetExerciseDataPet(req.StartTime, req.EndTime)
	log.Info("GetExerciseDataPet-exer:", len(exers))
	if err != nil {
		log.Info("GetExerciseDataPet-err", err)
		return &pb.MotionDataMapReply{Code: 10001}, nil
	}
	if len(exers) == 0 {
		return &pb.MotionDataMapReply{Code: 33013}, nil
	}
	var days []*pb.DayExerciseDataReply
	startZeroTime := GetZeroTime(req.StartTime)
	for startZeroTime <= req.EndTime {
		nextZeroTime := startZeroTime + 86400
		dayExerciseDataReply := &pb.DayExerciseDataReply{}
		for _, v := range exers {
			dayExerciseDataReply.DayTime = startZeroTime
			if v.ReportTime.Unix() >= startZeroTime && v.ReportTime.Unix() < nextZeroTime {
				var text []*pb.Pcoordinate
				if Json.Unmarshal([]byte(v.Coordinates), &text) != nil {
					log.Info("GetExerciseDataPet-err", err)
					return &pb.MotionDataMapReply{Code: 10001}, nil
				}
				lens := len(text)-1
				dayExerciseDataReply.Calorie += v.Cals
				var coordinates []*pb.Coordinate
				for kc, vc := range text {
					dayExerciseDataReply.Steps += vc.Steps
					//todo 处理坐标点与状态
					if kc == 0 || kc == lens || (kc)%9 == 0 {
						if vc.Longitude != 0 || vc.Latitude != 0 {
							coordinates = append(coordinates,
								&pb.Coordinate{State: vc.State, Longitude: vc.Longitude, Latitude: vc.Latitude, NowTime: vc.Nowtime})
						} else {
							if kc == 0 {
								coordinates = append(coordinates,
									&pb.Coordinate{State: text[kc+1].State,
										Longitude: text[kc+1].Longitude, Latitude: text[kc+1].Latitude, NowTime: text[kc+1].Nowtime})
							} else {
								coordinates = append(coordinates,
									&pb.Coordinate{State: text[kc-1].State,
										Longitude: text[kc-1].Longitude, Latitude: text[kc-1].Latitude, NowTime: text[kc-1].Nowtime})
							}
						}
					}
					if kc == 0 {
						continue
					}
					if vc.State == 1 {
						dayExerciseDataReply.StrenuousTimes += text[kc].Nowtime - text[kc-1].Nowtime
					} else {
						dayExerciseDataReply.CardioTimes += text[kc].Nowtime - text[kc-1].Nowtime
					}
				}
				dayExerciseDataReply.Coordinates = append(dayExerciseDataReply.Coordinates,coordinates...)
			}
		}
		days = append(days,dayExerciseDataReply)
		startZeroTime = nextZeroTime
	}
	return &pb.MotionDataMapReply{Code: 10000, Pid:req.GetPid(), Pdid:exers[0].Pdid, Data: days}, nil
}

//查询运动数据1.1
func (m *MotionDataRpc) GetMotionDataPetByTime(ctx context.Context, req *pb.ExerciseDataRequest) (*pb.MotionDataMapReply, error) {
	log.Info("GetMotionDataPetByTime-req:", req)
	if VerifyParamsStr(req.GetSource()) {
		return &pb.MotionDataMapReply{Code: Source_err_empty}, nil
	}
	if VerifyParamsUInt32(req.GetUid(),req.GetPid()) {
		return &pb.MotionDataMapReply{Code: 33001}, nil
	}
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.MotionDataMapReply{Code: 33012}, nil
		}
		log.Error("GetMotionDataPetByTime GetUserPetDB err:", err)
		return &pb.MotionDataMapReply{Code: 10001}, nil
	}
	motionDataPo := db.MotionDataPetPo{Pid:req.Pid}
	motionDataPos,err := motionDataPo.GetMotionDataPetByTime(Int32ToRemStr(req.Pid),req.StartTime,req.EndTime)
	log.Info("GetMotionDataPetByTime GetMotionDataPetByTime:", len(motionDataPos))
	if err != nil {
		log.Info("GetExerciseDataPet-err", err)
		return &pb.MotionDataMapReply{Code:10001}, nil
	}
	if len(motionDataPos) == 0 {
		return &pb.MotionDataMapReply{Code:33013}, nil
	}
	var days []*pb.DayExerciseDataReply
	startZeroTime := GetZeroTime(req.StartTime)
	for startZeroTime <= req.EndTime {
		nextZeroTime := startZeroTime+86400
		var records []*pb.Record
		day := &pb.DayExerciseDataReply{DayTime:startZeroTime,Types:2}
		for _, v := range motionDataPos {
			if v.ReportTime.Unix() >= startZeroTime && v.ReportTime.Unix() < nextZeroTime {
				strenuousMinute := v.StrenuousTimes/60
				minute := CountMinute(v.StartTime,v.EndTime)
				cardioDurationMinute := int32(minute-strenuousMinute)
				records = append(records,&pb.Record{
					TimeRecord:strings.Split(Int64ToTimeStr(v.StartTime.Unix())," ")[1][:5]+ "-"+
						strings.Split(Int64ToTimeStr(v.EndTime.Unix())," ")[1][:5],
					Steps:v.Steps,Calorie:v.Calorie,Pdid:v.Pdid,
					CardioTimes:v.CardioTimes, StrenuousTimes:v.StrenuousTimes,
					CardioDurationMinute:cardioDurationMinute,StrenuousDurationMinute:int32(strenuousMinute),
					ImageInfo:&pb.ImageInfo{Url:v.Url,Name:v.Name,Size:v.Size,Width:v.Width,Height:v.Height}})
				day.Steps += v.Steps
				day.Calorie += v.Calorie
				day.CardioTimes += v.CardioTimes
				day.StrenuousTimes += v.StrenuousTimes
				day.CardioDurationMinute += cardioDurationMinute
				day.StrenuousDurationMinute += int32(strenuousMinute)
			}
		}
		day.Records = records
		days = append(days,day)
		startZeroTime = nextZeroTime
	}
	return &pb.MotionDataMapReply{Code:10000,Data:days}, nil
}