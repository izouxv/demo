package service

import (
	"testing"
	"golang.org/x/net/context"
	"petfone-rpc/pb"
	log "github.com/cihub/seelog"
)

func TestRpc_msso_SearchDevice(t *testing.T) {
	Init()
	rpc_msso := Rpc_msso{}
	request := pb.DeviceRequest{Source:"AgIDAA==",Sn:"TEST00000000000M"}
	reply,err := rpc_msso.SearchDevice(context.Background(),&request)
	log.Info("err: ",err)
	log.Info("reply: ",reply)
}


