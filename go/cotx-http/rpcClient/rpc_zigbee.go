package rpcClient

import (
	"time"
	"cotx-http/pb"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var rpcClientZigbee pb.Rpc_ZigbeeClient

func NewRpcClinetZigbee()  {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	rpcClientZigbee = pb.NewRpc_ZigbeeClient(conn)
}
func GetZigbeeRpcClient()pb.Rpc_ZigbeeClient  {
	return rpcClientZigbee
}

/*获取网关zigbee的信息*/
func GetZigbee(req *pb.ReqGetZigbee)(*pb.ResGetZigbee,error)  {
	log.Info("rpc_zigbee:start get zigbee")
	res,err := GetZigbeeRpcClient().GetZigbee(context.Background(),req)
	return res,err
}

/*更新网关zigbee的状态*/
func UpdateZigbee(req *pb.ReqUpdateZigbee)(*pb.ResUpdateZigbee,error)  {
	log.Info("rpc_zigbee:start update zigbee")
	res,err := GetZigbeeRpcClient().UpdateZigbee(context.Background(),req)
	return res,err
}