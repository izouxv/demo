package rpc

import (
	"sync"
	"time"
	pb "account-domain-http/api/user/pb"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//const userPetFoneServer = "120.76.54.242:7005"
//const userPetFoneServer = "127.0.0.1:7005"
const userPetFoneServer = "rpc.petfone.penslink.com:7005"


var (
	userPetOnce       sync.Once
	userPetRpcClient  pb.MSsoClient
	userPetConn      *grpc.ClientConn
	err               error
)

func MUserPetRpcClient() pb.MSsoClient {
	return userPetRpcClient
}

func NewMUserPetRpcClient() {
	userPetOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		userPetConn, err = grpc.DialContext(ctx, userPetFoneServer, grpc.WithInsecure())
		if err != nil || userPetConn == nil {
			log.Error(err)
			return
		}
		userPetRpcClient = pb.NewMSsoClient(userPetConn)
	})
}
func UserPetClientClose() {
	if userPetConn != nil {
		userPetConn.Close()
	}
}


func GetDeviceBaseTenant(req *pb.DeviceRequest) (response *pb.DeviceReply, err error) {
	log.Infof("Start rpc_petfone GetDeviceBaseTenant: req(%#v)",req)
	response, err = MUserPetRpcClient().SearchDevice(context.Background(),req)
	return
}
