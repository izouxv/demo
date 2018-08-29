package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb"
	"sync"
	"time"
)

var (
	exerciseOnce      sync.Once
	exerciseRpcClient pb.ExerciseDataClient
	exerciseConn *grpc.ClientConn
)

//初始化ExerciseRpc
func ExerciseRpcInit(address string) {
	exerciseOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		exerciseConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("ExerciseRpcInit:", err)
			panic(err)
		}
		exerciseRpcClient = pb.NewExerciseDataClient(exerciseConn)
	})
}

//结束Rpc
func ExerciseRpcClose() {
	if exerciseConn != nil {
		exerciseConn.Close()
	}
}

//调用exerciseRpc
func ExerciseRpc(exercise *pb.ExerciseDataRequest, method string) *pb.ExerciseDataReply {
	log.Info("ExerciseRpc-exercise:", exercise.Source,exercise.Uid,exercise.Pid,exercise.Pdid, len(exercise.Pcoordinates))
	var err error
	var exerciseReply *pb.ExerciseDataReply
	switch method {
	//添加信息
	case "SetExerciseData":
		exerciseReply, err = exerciseRpcClient.SetExerciseData(context.Background(), exercise)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("ExerciseRpc-Error", err)
		exerciseReply.Code = 10001
	}
	log.Info("ExerciseRpc-exerciseReply:", exerciseReply)
	return exerciseReply
}

//调用exerciseRpc
func MotionDataMapRpc(exercise *pb.ExerciseDataRequest, method string) *pb.MotionDataMapReply {
	log.Info("MotionDataMapRpc exercise:", exercise)
	var err error
	var exerciseMapReply *pb.MotionDataMapReply
	switch method {
	//获取信息
	case "GetExerciseDataPet":
		exerciseMapReply, err = exerciseRpcClient.GetExerciseDataPet(context.Background(), exercise)
	case "GetMotionDataPetByTime":
		exerciseMapReply, err = exerciseRpcClient.GetMotionDataPetByTime(context.Background(), exercise)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("MotionDataMapRpc Error:", err)
		exerciseMapReply.Code = 10001
	}
	return exerciseMapReply
}
