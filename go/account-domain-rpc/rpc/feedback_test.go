package rpc

import (
	pb "account-domain-rpc/api/feedback"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"testing"
)

func TestAddFeedbackServer_GetNodes(t *testing.T) {
	NewSource()
	Convey("测试添加反馈", t, func() {
		rule := FeedBackServer{}
		//f := &pb.Feedback{UserInfo:"用户",Description:"用户描述",AppInfo:"app描述",MobileInfo:"mobile描述 ",ExtendInfo:"扩展描述",Files:files, Contact:"lvxx@radacat.com",DeviceInfo:"设备描述"}
		rep, err := rule.AddFeedback(context.Background(), &pb.AddFeedbackRequest{Source:"AQIDAA==",Description:"111",AppInfo:"test",MobileInfo:"test"})
		fmt.Printf("rep :(%s),err :(%s)", rep, err)
	})
}
func TestGetFeedbacksServer_GetFeedbacks(t *testing.T) {
	NewSource()
	Convey("分页查询反馈", t, func() {
		/*	//f := &pb.GetFeedbacksRequest{Page:1,Count:3}
			rule	 := FeedBackServer{}
			//rep ,err := rule.(context.Background(), &pb.AddFeedbackRequest{})
			fmt.Printf("rep :(%s),err :(%s)",rep,err)*/
	})
}
