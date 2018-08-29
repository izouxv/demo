package drbac

import (
	"account-domain-rpc/module"
	"fmt"
	"github.com/gin-gonic/gin/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func NewSouce() {
	module.NewMysqlClient("192.168.1.6", "3306", "domain", "root", "Radacat2017")
	module.NewRedisClient(module.Persistence, "192.168.1.6", "6379", 10, 10, "radacat1234")
	module.NewRedisClient(module.Nopersistence, "192.168.1.6", "6380", 10, 10, "radacat1234")
}

func TestGetDomainRolePermissionByUid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	//redis := module.RedisClient(module.Persistence)
	Convey("TestGetDomainRolePermissionByUid", t, func() {
		drp, err := getDomainRolePermissionByUid(933262636907241472, tx)
		So(err, ShouldBeNil)
		for _, v := range drp {
			fmt.Println("---------第一级")
			fmt.Println("isDefault:", v.IsDefaultDomain)
			fmt.Printf("%+v\n", v.Domain)
			fmt.Printf("%+v\n", v.Role)
			fmt.Printf("%+v\n", v.Organization)
			printChildren(v.Children)
		}
		resp, err := json.Marshal(drp)
		fmt.Println("json格式", string(resp))
	})
}

func printChildren(trees []*DomainTree) {
	for _, v := range trees {
		fmt.Printf("%+v\n", v.Domain)
		printChildren(v.Children)
	}
}
