package rpc

import ("testing"
	"petfone-http/pb"
	log "github.com/cihub/seelog"

)

func TestGetTextSemantics(t *testing.T) {
	AiuiRpcInit("192.168.1.11:7007")
	request := &pb.GetTextSemanticsRequest{}
	reply,err := GetTextSemantics(request)
	log.Errorf("err(%#v)",err)
	log.Errorf("reply(%#v)",reply)
}