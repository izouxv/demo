package rpcClient

import (
	"time"
	"cotx-http/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

var rpcClientUpdateDevice pb.UpdateDevceiClient

func NewRpcClinetUpdateDevice()  {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	rpcClientUpdateDevice = pb.NewUpdateDevceiClient(conn)
}
func GetUpdateDeviceRpcClient()pb.UpdateDevceiClient  {
	return rpcClientUpdateDevice
}

func GetNewGatewayVersion(req *pb.ReqGetNewGatewayVersion) (*pb.ResGetNewGatewayVersion ,error) {
	log.Info("rpcclient:start get new gateway version")
	res,err :=  GetUpdateDeviceRpcClient().GetNewGatewayVersion(context.Background(),req)
	return res,err
}

func UpdateGatewayVersion(req *pb.ReqUpdateGatewayVersion)(*pb.ResUpdateGatewayVersion,error)  {
	log.Info("rpcclient:start update gateway version")
	res,err := GetUpdateDeviceRpcClient().UpdateGatewayVersion(context.Background(),req)
	return res,err

}

/*获取网关更新的状态*/
func GetGatewayVersionState(req *pb.ReqGetGatewayVersionState)(*pb.ResGetGatewayVersionState,error)  {
	log.Info("rpcclient:start get gateway version state")
	res,err := GetUpdateDeviceRpcClient().GetGatewayVersionState(context.Background(),req)
	return res,err
}
