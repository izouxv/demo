package rpcClient

import (
	"time"
	"cotx-http/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

var rpcClientVpn pb.RpcVpnClient

func NewRpcClinetVpn()  {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	rpcClientVpn = pb.NewRpcVpnClient(conn)
}
func GetVpnRpcClient()pb.RpcVpnClient {
	return rpcClientVpn
}

/*获取网关的vpn信息*/
func GetVpn(req *pb.ReqGetVpn)(*pb.ResGetVpn,error)  {
	log.Info("vpn:start get gateway vpn")
	res,err := GetVpnRpcClient().GetVpn(context.Background(),req)
	if err != nil {
		log.Error("vpn:rpc_err",err)
	}
	return res,err
}

/*更新网关的vpn信息*/

func UpdateVpn(req *pb.ReqUpdateVpn)(*pb.ResUpdateVpn,error)  {
  log.Info("vpn:start update gateway vpn")
  res,err := GetVpnRpcClient().UpdateVpn(context.Background(),req)
	if err != nil {
		log.Error("vpn:rpc_err",err)
	}
	return res,err
}

/*网关的ping 指令*/
func PingGateway(req *pb.ReqPingGateway)(*pb.ResPingGateway,error)  {
	log.Info("ping:start ping gateway vpn")
	res,err := GetVpnRpcClient().PingGateway(context.Background(),req)
	if err != nil {
		log.Error("ping:rpc_err",err)
	}
	return res,err
}
