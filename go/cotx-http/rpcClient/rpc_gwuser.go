package rpcClient

import (
	"cotx-http/pb"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"time"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:7007"
)

var (
	gwuserRpcClient pb.GwUserClient
)

func NewGwUserRpcClient() {
     ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
     conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	gwuserRpcClient = pb.NewGwUserClient(conn)
}
func GetGwUserRpcClient() pb.GwUserClient {
	return gwuserRpcClient
}
//绑定网关
func BindGateway(req *pb.ReqGwUser) (*pb.ResGwUser,error){
	log.Info("rpc/binging gateway")
	resgatewayuser,err:= GetGwUserRpcClient().AddGateway(context.Background(),req)
	return resgatewayuser,err
}
//用户授权
func BingAccount(req *pb.ReqGwUser)(*pb.ResGwUser ,error) {
	log.Info("rpc/authotise account")
	resgwuser,err:= GetGwUserRpcClient().AddGatewayAccount(context.Background(),req)
	return resgwuser,err
}
//删除授权用户
func DeletAuthoriseAccount(req *pb.ReqGwUser)(*pb.ResGwUser,error)  {
	log.Info("rpc/delet authorise account")
	resgwuser,err := GetGwUserRpcClient().DeletGatewayAccount(context.Background(),req)
	return resgwuser,err
}
//展示授权用户
func ShowAuthoiseAccount(req *pb.ReqGwUser)(*pb.ResAccounts ,error) {
	log.Info("rpc/show authorise account")
	resgwuser,err:= GetGwUserRpcClient().GetGatewayAccoount(context.Background(),req)
	return resgwuser,err
}
//解除网关的绑定关系
func UnwoundBindGateway(req *pb.ReqGwUser)(*pb.ResGwUser ,error) {
	log.Info("rpc/delet binding gateway ")
	resgwuser ,err:= GetGwUserRpcClient().DeletGateway(context.Background(),req)
	return resgwuser,err
}
//验证网关是否被绑定
func ValidationGateway(req *pb.ReqGwUser)(*pb.ResGwUser ,error) {
	log.Info("rpc/validation gateway")
	resgwuser,err:= GetGwUserRpcClient().ValidationGateway(context.Background(),req)
	return resgwuser,err
}
// 验证用户是否已经把该网关授权给该帐号
func ValidationGatewayAccount(req *pb.ReqGwUser)(*pb.ResGwUser,error){
	log.Info("rpc/validation gateway account")
	resgwuser,err:= GetGwUserRpcClient().ValidationGatewayAccount(context.Background(),req)
	return resgwuser,err
}
//展示用户所有的网关信息
func ShowAllUserGateway(req *pb.ReqGwUser)(*pb.ResUserGateways,error){
	log.Info("rpc/show all wuer_gateways")
	resallgws,err:= GetGwUserRpcClient().GetAllGateways(context.Background(),req)
	return resallgws,err
}

//
func ShowNodesByGateway(gwreq *pb.ReqGwUser)(*pb.ResShowNodesByGw,error)  {
	log.Info("rpc/show nodes by gateway")
	gwres,err := GetGwUserRpcClient().ShowNodesByGateway(context.Background(),gwreq)
	return gwres,err
}

//
func ShowAllNodesPos(gwreq *pb.ReqGwUser)(*pb.ResShowAllNodesPos,error)  {
	log.Info("rpc/ show all nodes postion")
	gwres,err:= GetGwUserRpcClient().ShowAllNodesPos(context.Background(),gwreq)
	return gwres,err
}
// 
func ShowAllGws(gwreq *pb.ReqGwUser)(*pb.ResShowAllGws,error)  {
	log.Info("start \rpc show all gateways ")
	gwres,err := GetGwUserRpcClient().ShowAllGws(context.Background(),gwreq)
	return gwres,err
}



