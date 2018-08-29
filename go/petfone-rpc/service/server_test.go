package service_test

import (
	"context"
	"testing"
	"fmt"
	"github.com/cihub/seelog"
	"google.golang.org/grpc"
	"petfone-rpc/pb"
	"regexp"
)

var (
	address = "192.168.1.178:7005"
)


//todo 测试宠物功能
func TestPetinfoRpc_GetPetInfoByUid(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewPetInfoClient(conn)
	req := &pb.PetInfoRequest{Source:"AgIDAA==",Uid:1000001,Pid:1,Avatar:"2f2ad8cc6d7aa2cdd85d6d46e72e5938",Birthday:1529484587,Breed:1,
		Nickname:"ww",Gender:1,Weight:1.1,Somatotype:1}
	//r, e := c.GetPetInfoByUid(context.Background(), req)
	//for _, v := range r.Petinfos {
	//	fmt.Println("v:",v)
	//	fmt.Println("vt:",v.Trains)
	//}
	//r, e := c.GetPetInfoByPid(context.Background(), req)
	//fmt.Println(r)
	r, e := c.SetPetInfo(context.Background(), req)
	fmt.Println(r)
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试宠物与设备功能
func TestPetinfoRpc_SetDevicePet(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewPetInfoClient(conn)
	req := &pb.PetInfoRequest{Source: "AgIDAA==", Uid: 1000001, Pid: 1,Did:1}
	r, e := c.GetPetInfoByPid(context.Background(), req)
	fmt.Println(r)
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试宠物训练功能
func TestPetinfoRpc_GetPetTrainByPid(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewPetInfoClient(conn)
	req := &pb.PetTrainRequest{Source:"AgIDAA==",Uid:1000001,Pid:4,StartTime:1522984455,EndTime:1525708800}
	r, e := c.GetPetTrainByPid(context.Background(), req)
	fmt.Println("r:", r.Code)
	for _, v := range r.SliceTrains {
		fmt.Println("v:",v.Trains)
	}
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 运动数据
func TestExerciseDatasRpc_GeDaysExerciseDataPet(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewExerciseDataClient(conn)
	req := &pb.ExerciseDataRequest{Source:"AgIDAA==",Pid:1,Uid:1000001,StartTime:1531065600,EndTime:1531105600}
	r, e := c.GetMotionDataPetByTime(context.Background(), req)
	fmt.Println("r:", r.Code)
	for _, v := range r.Data {
		fmt.Println("v:",v)
	}
	if e != nil {
		fmt.Println("e:", e)
	}
	//dayTime:1531065600 cardioTimes:1030 cals:0.1 steps:47
}

//todo 测试分享功能
func TestShareRpc_GetShare(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewShareManageClient(conn)
	req := &pb.ShareRequest{Source:"AgIDAA==",OwnerUid:1000002,MemberUid:1000003,Pids:[]int32{1}}
	r, e := c.GetShare(context.Background(), req)
	//r, e := c.SetShare(context.Background(), req)
	//r, e := c.DeleteShare(context.Background(), req)
	fmt.Println("r:", r.Code)
	for _, v := range r.Shares {
		fmt.Println("v:",v)
	}
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试分页查询sso
func TestRpc_msso_GetPageSsoInfos(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewMSsoClient(conn)
	req := &pb.PageRequest{Source:"AgIDAA==",Page:1,Count:10,Sort:pb.Sort_ASC}
	r, e := c.GetPageSsoInfos(context.Background(), req)
	//r, e := c.SetShare(context.Background(), req)
	fmt.Println("r:", r.Code,",",r.TotalCount)
	for _, v := range r.MSsos {
		fmt.Println("v:",v)
	}
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试宠聊功能
func TestPetfoneRpc_GetPetChatByPid(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewPetfoneClient(conn)
	req := &pb.PetChatRequest{Source:"AgIDAA==",Uid:1000001,Pid:1,Input:"movement?",Language:pb.Language_En}
	r, e := c.GetPetChatByPid(context.Background(), req)
	fmt.Println(r.Code)
	for _, v := range r.ChatMsgs {
		seelog.Info("Output1:",v.Output1)
		seelog.Info("Output2:",v.Output2)
	}
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试设备功能
func TestDeviceRpc_Device(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewDevicesClient(conn)
	req := &pb.DeviceRequest{Source:"AgIDAA==",Uid:1000001,Sn:"test001"}
	r, e := c.GetDeviceSn(context.Background(), req)
	//for _, v := range r.Petinfos {
	//	fmt.Println("v:",v)
	//	fmt.Println("vt:",v.Trains)
	//}
	fmt.Println(r)
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试分页查询设备
func TestRpc_msso_GetPageDevices(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewMSsoClient(conn)
	//req := &pb.PageRequest{Source:"AgIDAA==",Page:1,Count:10,Sort:pb.Sort_ASC}
	//r, e := c.GetPageDevices(context.Background(), req)
	req := &pb.DeviceRequest{Source:"AgIDAA==",Did:1,DeviceMac:"",Sn:"TEST00000000000"}
	r, e := c.SearchDevice(context.Background(), req)
	fmt.Println("r:", r)
	//fmt.Println("r:", r.Code,",",r.TotalCount)
	//for _, v := range r.Devices {
	//	fmt.Println("v:",v)
	//}
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试品种
//func TestFileRpc_GetBreeds(t *testing.T) {
//	fmt.Println("rrrrrrr")
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		fmt.Println("err:", err)
//	}
//	defer func() {
//		conn.Close()
//	}()
//	c := pb.NewFilesClient(conn)
//	req := &pb.FilesRequest{Source:"AgIDAA==",Id:1,Number:180,Types:1}
//	r, e := c.GetBreeds(context.Background(), req)
//	if e != nil {
//		fmt.Println("e:", e)
//	}
//	fmt.Println(r.Code)
//	reqMap := &pb.FilesMapRequest{Source:"AgIDAA=="}
//	var files []*pb.FilesRequest
//	for _, v := range r.GetFiles() {
//		//fmt.Println("v:",v.Name)
//		info := method.GetBaikeDog(v.Name)
//		//fmt.Println("--",info)
//		for _,kv := range keys {
//			var aa []rune
//			for _,v := range []rune(info) {
//				if v == 160 {
//					continue
//				}
//				aa = append(aa,v)
//			}
//			info = compressStr(strings.Replace(string(aa),kv,",",-1))
//
//		}
//		file := &pb.FilesRequest{Id:v.Id, Describe:info}
//		files = append(files,file)
//	}
//	reqMap.Files = files
//	r, e = c.SetBreeds(context.Background(), reqMap)
//	if e != nil {
//		fmt.Println("e:", e)
//	}
//	fmt.Println(r.Code)
//}
var keys = []string{"中文学名","别称","界","门","亚门","纲","亚纲",
	"目","亚目","亚科","属","种","亚种",
	"脊索动物","脊椎动物","哺乳","真兽","食肉","裂脚","犬科","犬亚科","犬属",
	" "}
func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//空白符正则
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

//todo 测试分页查询设备
func TestRpc_msso_SearchDevice(t *testing.T) {
	fmt.Println("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewMSsoClient(conn)
	req := &pb.DeviceRequest{Source:"AgIDAA==",Sn:"TEST00000000000X"}
	r, e := c.SearchDevice(context.Background(), req)
	fmt.Println("r:", r)
	if e != nil {
		fmt.Println("e:", e)
	}
}
