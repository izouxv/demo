package service

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"strings"
	"petfone-rpc/core"
	"io"
	"strconv"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
)

type PetfoneRpc struct {
	Agents    []*pb.AgentInfo
}

//获取业务信息
func (this *PetfoneRpc) GetPetfoneByUid(ctx context.Context, req *pb.PetfoneRequest) (*pb.PetfoneReply, error) {
	log.Info("GetPetfoneByUid-req:", req)
	if util.VerifyParamsStr(req.GetSource()){
		return &pb.PetfoneReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.PetfoneReply{Code: util.Params_err_empty}, nil
	}
	petfonePo := &db.PetfonePo{Uid: req.Uid, DataState: 1}
	err := petfonePo.GetPetfoneDB()
	if err != nil {
		log.Info("err", err)
		return &pb.PetfoneReply{Code: 10001}, nil
	}
	log.Info("GetPetfoneByUid-petfonePo:", petfonePo)
	return &pb.PetfoneReply{Code: 10000, Radius: petfonePo.Radius,Map:pb.Map(petfonePo.Map)}, nil
}

//修改用户扩展业务信息
func (this *PetfoneRpc) UpdatePetfoneByUid(ctx context.Context, req *pb.PetfoneRequest) (*pb.PetfoneReply, error) {
	log.Info("UpdatePetfoneByUid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetfoneReply{Code: util.Source_err_empty}, nil
	}
	if req.Map == pb.Map_AllMap && req.Radius == 0 {
		return &pb.PetfoneReply{Code: util.Params_err_empty}, nil
	}
	petfonePo := &db.PetfonePo{Uid: req.Uid}
	petfoneP := db.PetfonePo{Radius: req.Radius, UpdateTime: util.GetNowTime(),Map:int32(req.Map)}
	err := petfonePo.UpdatePetfoneDB(petfoneP)
	if err != nil {
		return &pb.PetfoneReply{Code: 10001}, nil
	}
	return &pb.PetfoneReply{Code: 10000}, nil
}

//添加业务信息
func (this *PetfoneRpc) SetPetfoneByUid(ctx context.Context, req *pb.PetfoneRequest) (*pb.PetfoneReply, error) {
	log.Info("SetPetfoneByUid-req:", req)
	return &pb.PetfoneReply{Code: 10000}, nil
}

//查询宠聊信息
func (this *PetfoneRpc) GetPetChatByPid(ctx context.Context, req *pb.PetChatRequest) (*pb.PetChatReply, error) {
	log.Info("GetPetChatByPid-req:", req)
	if util.VerifyParamsStr(req.GetSource()){
		return &pb.PetChatReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid(),req.Pid) {
		return &pb.PetChatReply{Code: util.Params_err_empty}, nil
	}
	if util.VerifyParamsStr(req.GetInput(),req.Language.String()){
		return &pb.PetChatReply{Code: util.Params_err_empty}, nil
	}
	nowTime := util.GetNowTime()
	//todo 输入内容
	go func() {
		chatMsgPo := db.PetChatMsgPo{
			MsgSource:1,MsgAbout:util.Int32ToStr(req.Uid)+","+util.Int32ToStr(req.Pid)+","+req.Language.String(),
			Msg:req.Input,CreationTime:nowTime}
		chatMsgPo.SetMsg()
	}()
	//校验权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetChatReply{Code: 33012}, nil
		}
		log.Error("GetPetChatByPid UpdatePetInfoByPid-err:", err)
		return &pb.PetChatReply{Code: 10001}, nil
	}
	//声明变量
	zeroTime := util.GetZeroTime(nowTime.Unix())
	petInfoPo := &db.PetInfoPo{Pid: req.Pid, DataState: 1}
	breeds := &db.BreedInfoPo{Types:1}
	petTrainPo := &db.PetTrainPo{Pid:req.Pid}
	var exers []*db.ExerciseDataPetPo
	var chatMsgs []*pb.ChatMsg
	var tagKeys, tagValues []string
	if req.Language == pb.Language_Cn {
		tagKeys = util.TagCnKeys
		tagValues = util.TagCnValues
	} else if req.Language == pb.Language_En {
		tagKeys = util.TagEnKeys
		tagValues = util.TagEnValues
	} else {
		log.Error("GetPetChatByPid Language err")
		return &pb.PetChatReply{Code: 10001}, nil
	}
	//查询宠物训练信息
	petTrainPos, err := petTrainPo.GetPetTrainsDB(zeroTime,nowTime.Unix())
	if err != nil {
		log.Error("GetPetChatByPid GetPetTrainsDB err:", err)
		return &pb.PetChatReply{Code: 10001}, nil
	}
	if len(petTrainPos) == 0 {
		log.Error("GetPetChatByPid petTrainPos len empty")
		return &pb.PetChatReply{Code: 33013}, nil
	}
	for _,v := range petTrainPos {
		tagKeys = append(tagKeys,v.Name)
	}
	k := MatchingKey(req.GetInput(),tagKeys)
	chatMsg := &pb.ChatMsg{Types:1}
	log.Info("GetPetChatByPid MatchingKey:",k)
	switch {
	case k <= 0:
		chatMsg.Output1 = tagValues[k]
	case k >= len(tagKeys)-3: //查询宠物训练信息
		chatMsg.Output1 = util.StrReplaceStar(
			util.StrReplaceStar(tagValues[9],tagKeys[k],1), util.Int32ToStr(petTrainPos[k-10].Counter),-1)
	case k >= 1 && k <= 4: //查询宠物属性
		err = petInfoPo.GetPetInfoDB()
		if err != nil {
			log.Info("GetPetChatByPid GetPetInfoDB err:", err)
			return &pb.PetChatReply{Code: 10001}, nil
		}
		if k==1 {
			chatMsg.Output1 = tagValues[k]
			break
		}
		if k==2 {
			breeds.Id = petInfoPo.Breed
			err = breeds.GetBreedinfoDB()
			if err != nil {
				log.Info("GetPetChatByPid GetBreedinfoDB err:", err)
				return &pb.PetChatReply{Code: 10001}, nil
			}
			if req.Language == pb.Language_Cn {
				chatMsg.Output1 = util.StrReplaceStar(tagValues[k],breeds.NameCh,-1)
			} else {
				chatMsg.Output1 = util.StrReplaceStar(tagValues[k],breeds.NameEn,-1)
			}
			break
		}
		if k==3 {
			ageStr := ""
			age := util.GetNowTime().Year()-petInfoPo.Birthday.Year()
			if age == 0 {
				if req.Language == pb.Language_En {
					ageStr = "not even a"
				} else {
					ageStr = "还不满1"
				}

			} else {
				ageStr = util.Int32ToStr(int32(age))
			}
			chatMsg.Output1 = util.StrReplaceStar(tagValues[k],ageStr,-1)
			break
		}
		if k==4 {
			chatMsg.Output1 = util.StrReplaceStar(tagValues[k],util.FloatToStr(float64(petInfoPo.Weight),32),-1)
			break
		}
	case k >= 5 && k <= 7: //查询宠物运动信息
		exer := &db.ExerciseDataPetPo{Pid: req.GetPid(), DataState: 1}
		exers, err = exer.GetExerciseDataPet(zeroTime, nowTime.Unix())
		log.Info("GetExerciseDataPet-exer:", len(exers))
		if err != nil {
			log.Info("GetPetChatByPid GetExerciseDataPet err:", err)
			return &pb.PetChatReply{Code: 10001}, nil
		}
		if len(exers) == 0 {
			chatMsg.Output1 = tagValues[8]
			break
		}
		if k==5 {
			chatMsg.Types = 4
			chatMsgs = append(chatMsgs,&pb.ChatMsg{Types : 1,Output1:tagValues[k]})
			var pCoordinates []*pb.Pcoordinate
			util.Json.Unmarshal([]byte(exers[len(exers)-1].Coordinates), &pCoordinates)
			chatMsg.Output2 = pCoordinates[len(pCoordinates)-1:]
			break
		}
		if k==6 { //todo 计算消耗
			var steps int32; var cals float32;var duration int64
			for _, v := range exers {
				cals += v.Cals
				var coordinates []*pb.Pcoordinate
				util.Json.Unmarshal([]byte(v.Coordinates), &coordinates)
				for kc, vc := range coordinates { //todo 叠加步数
					steps += vc.Steps
					if kc != 0 {
						duration += vc.Nowtime - coordinates[kc-1].Nowtime
					}
				}
			}
			chatMsg.Output1 =
				util.StrReplaceStar(
					util.StrReplaceStar(
						util.StrReplaceStar(
							util.StrReplaceStar(tagValues[k], util.Int32ToStr(steps),1),
						fmt.Sprintf("%.1f",cals),1),
					util.Int64ToStr(duration/60),1),
				util.Int64ToStr(duration%60), 1)
			break
		}
		if k==7 {
			chatMsg.Types = 3
			chatMsgs = append(chatMsgs,&pb.ChatMsg{Types:1,Output1:tagValues[k]})
			var pCoordinates []*pb.Pcoordinate
			for _, v := range exers {
				var coordinates []*pb.Pcoordinate
				util.Json.Unmarshal([]byte(v.Coordinates), &coordinates)
				pCoordinates = append(pCoordinates,coordinates...)
			}
			chatMsg.Output2 = pCoordinates
		}
	}
	chatMsgs = append(chatMsgs,chatMsg)
	return &pb.PetChatReply{Code: 10000,ChatMsgs:chatMsgs}, nil
}

//查询宠聊关键词信息
func (this *PetfoneRpc) GetPetChatKey(ctx context.Context, req *pb.PetChatRequest) (*pb.PetChatKeysReply, error) {
	log.Info("GetPetChatKey-req:", req)
	if util.VerifyParamsStr(req.GetSource()){
		return &pb.PetChatKeysReply{Code: util.Source_err_empty}, nil
	}
	petChatKeyPo := &db.PetChatKeyPo{}
	petChatKeyPos, err := petChatKeyPo.GetPetChatKeysDB()
	if err != nil {
		log.Info("GetPetChatKey err", err)
		return &pb.PetChatKeysReply{Code: 10001}, nil
	}
	var keys []*pb.PetChatKey
	for _, v := range petChatKeyPos {
		keys = append(keys,&pb.PetChatKey{Id:v.Id, InfoCn:v.InfoCn, InfoEn:v.InfoEn})
	}
	return &pb.PetChatKeysReply{Code: 10000, ChatKeys:keys}, nil
}

//用户操作日志
func (this *PetfoneRpc) SetActionLog(stream pb.Petfone_SetActionLogServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Info("SetActionLog read done")
				return nil
			}
			log.Error("SetActionLog err",err)
			return err
		}
		log.Info("SetActionLog in:",in)
		go func(in *pb.AgentInfo) {
			nowTime,_ := util.Int64ToTime(in.CreateTime)
			if nowTime.Unix() == 0 {
				log.Info("SetActionLog nowTime:",nowTime)
				nowTime = util.GetNowTime()
			}
			sso := db.Sso{Id:in.Uid, LastLoginTime:nowTime}
			var addr string
			var ip int64
			if in.DevInfo != "" && in.Ip != "" {
				ip, _ = util.IPToInt64(in.Ip)
				chanIp := make(chan string)
				go util.IPToAddr(chanIp,in.Ip)
				select {
				case <-time.After(time.Second*3):
				case addr = <-chanIp:
					sso.LastLoginAddr = addr
				}
				sso.LastLoginIp = ip
				sso.LastLoginDevInfo = in.DevInfo
			}
			if in.Path == "/petfone/v1.0/sessions" || in.Path == "/petfone/v1.0/reg" {
				err := db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
					sso.UpdateSso(tx)
					return nil
				})
				if err != nil {
					log.Error("SetActionLog UpdateSso err:",err)
				}
			}
			util.LogWriter.Write([]byte(fmt.Sprintln(in.Uid,in.Path,in.Method,in.Code,
				in.Duration,util.TimeToStrCha2(nowTime),addr,ip,in.DevInfo)))
		}(in)
	}
}

func MatchingKey(input string, tagKeys []string) int {
	//flag := unicode.Is(unicode.Scripts["Han"], []rune(input)[0])
	for k,v := range tagKeys {
		if strings.Contains(input,v) {
			return k
		}
	}
	return 0
}

//限制IP访问次数
func (this *PetfoneRpc) CheckPetfoneIp(ctx context.Context, req *pb.CheckPetfoneIpRequest) (*pb.CheckPetfoneIpResponse, error) {
	log.Info("CheckPetfoneIp...", req)
	client := core.RedisClient(6379)
	times, err := client.Do("get", req.Source[:2]+util.LimitIp+req.Ip)
	if err != nil {
		log.Error("Get petfone ip times Failed,", err)
		return &pb.CheckPetfoneIpResponse{Code: util.Code_err}, nil
	}
	defer client.Close()
	if times == nil {
		// 设置次数
		reply, err := client.Do("set", req.Source[:2]+util.LimitIp+req.Ip, 1, "EX", 1*10)
		if err != nil || reply == 0 {
			log.Info("Set petfone ip times Failed,", err)
			return &pb.CheckPetfoneIpResponse{Code: 33006}, nil
		}
	} else {
		count,err:= strconv.Atoi(string(times.([]uint8)))
		if count > 10 {
			log.Errorf("IP访问次数超限制count(%d)",count)
			return &pb.CheckPetfoneIpResponse{Code: util.Ip_times_exceed}, nil
		}
		reply, err := client.Do("incrby", req.Source[:2]+util.LimitIp+req.Ip,1)
		if err != nil {
			log.Info("incrby petfone ip times Failed,", err)
			return &pb.CheckPetfoneIpResponse{Code: 33006}, nil
		}
		log.Debugf("reply(%#v)",reply)
	}
	return &pb.CheckPetfoneIpResponse{Code: 10000}, nil
}
