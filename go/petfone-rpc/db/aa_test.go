package db

import (
	"flag"
	log "github.com/cihub/seelog"
	"context"
	"testing"
	"time"
	"fmt"
	"google.golang.org/grpc"
	"petfone-rpc/core"
	"petfone-rpc/pb"
	"github.com/jinzhu/gorm"
)

var (
	address = "192.168.1.178:7005"
	ncfgFile = flag.String("config", "../config/conf.yaml", "config path")
)

//todo 测试消息通知
func TestTestPicPo_GetPic(t *testing.T) {
	flag.Parse()
	core.ContextInit(*ncfgFile)
	testPo := &TestPicPo{Id:16}//,State:2,Pic:[]byte("aaaaaaaa")
	err := testPo.GetPic()
	//err := testPo.SetPic()
	log.Info("err:",err)
	log.Info("testPo:",string(testPo.Pic))
}

//todo 测试消息通知
func TestNoticePo_UpdateNoticePo(t *testing.T) {
	flag.Parse()
	core.ContextInit(*ncfgFile)

	f := NoticePo{
		Id:			4,
		To:			1000002,
		State:		1,
		UpdateTime:	time.Time{},
	}
	log.Info("sss:",f.UpdateNoticePo())
	log.Info(f)
}

//todo 测试分享功能
func TestShareUserPo_SetShareDB(t *testing.T) {
	fmt.Print("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Print("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewShareManageClient(conn)
	req := &pb.ShareRequest{Source:"AgIDAA=="}
	//r, e := c.SetShare(context.Background(), req)
	r, e := c.DeleteShare(context.Background(), req)
	fmt.Print("r:",r)
	if e != nil {
		fmt.Println("e:", e)
	}
}

//todo 测试分享功能
func TestPetChatPo_GetPetChatKeyDB(t *testing.T) {
	fmt.Print("rrrrrrr")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Print("err:", err)
	}
	defer func() {
		conn.Close()
	}()
	c := pb.NewShareManageClient(conn)
	req := &pb.ShareRequest{Source:"AgIDAA=="}
	//r, e := c.SetShare(context.Background(), req)
	r, e := c.DeleteShare(context.Background(), req)
	fmt.Print("r:",r)
	if e != nil {
		fmt.Println("e:", e)
	}
}

// todo 导入宠端设备数据到表device_train
func Test1(t *testing.T) {
	//Init()
	db,err:=gorm.Open("mysql","root:Radacat2017@tcp(192.168.1.6:3306)/petfone?charset=utf8")
	if err != nil {
		log.Infof("err(%#v)",err)
		return
	}
	fmt.Println("....")
	// 1.查询所有宠物训练数据
	petTrainPo := make([]*PetTrainPo, 100)
	err = db.Table("pet_train").Find(&petTrainPo).Error
	if err != nil {
		log.Infof("err(%#v)", err)
	}
	fmt.Print(petTrainPo)
	fmt.Print(len(petTrainPo))
	for index, value := range petTrainPo {
		log.Infof("index(%#v),value(%#v)", index, value)
	}
	// 2.查询所有的设备
	device := make([]*DevicePo,100)
	err = db.Table("device").Find(&device).Error
	if err != nil {
		log.Infof("get devices err(%#v)",err)
	}
	deviceTrain := make([]*DeviceTrainPo,0)
	for index,value := range device{
		log.Infof("device index(%#v),value(%#v)", index, value)
		// 3.查找设备有没有被绑定

	}
	// 3.
	defer db.Close()
}