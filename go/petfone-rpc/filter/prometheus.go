package filter

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
)

func PrometheusInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//methods := info.FullMethod
		//metrics.AddCounter(GetKey(methods)+"_count",1)
		//metrics.AddCounter("total_count", 1) //total_count
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func GetKey(methed string) string {
	//var key string
	//switch methed {
	//case "/pb.Sso/GetUserInfo":
	//	key = strings.Replace(methed,"/","_",-1)[1:]
	//	fmt.Println(key)
	//case "/pb.Sso/Login":
	//	key = strings.Replace(methed,"/","_",-1)[1:]
	//	fmt.Println(key)
	//case "/pb.Account/Show":
	//	key = strings.Replace(methed,"/","_",-1)[1:]
	//	fmt.Println(key)
	//default:
	//	key = "other"
	//}
	return strings.Replace(methed, "/", "_", -1)[1:]
}
