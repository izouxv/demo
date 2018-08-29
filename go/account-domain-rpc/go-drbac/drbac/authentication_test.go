package drbac

import (
	"account-domain-rpc/module"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_Authentication(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	pool := module.RedisClient(module.Persistence)
	Convey("TestDrbacServer_Authentication", t, func() {
		userToken, isTrue, err := authentication("SuperAdmin1511342126", "123456", tx, pool)
		So(err, ShouldBeNil)
		So(isTrue, ShouldBeTrue)
		fmt.Println(userToken)
	})
}

func Test_ReAuthorization(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	pool := module.RedisClient(module.Nopersistence)
	ds, err := NewDrbacServer(tx, pool)
	if err != nil {
		return
	}
	Convey("测试重新认证，重置redis中用户信息", t, func() {
		userToken, isTrue, err := authentication("SuperAdmin1511342126", "123456", tx, pool)
		So(err, ShouldBeNil)
		So(isTrue, ShouldBeTrue)
		time.Sleep(time.Second * 30)

		oid := userToken.DomainRolePermissions[1].Organization.Oid
		uid := userToken.User.Uid
		did := userToken.DomainRolePermissions[1].Domain.Did
		domain, err := ds.CreateDomainBaseDomain(oid, did, uid, "嘉兴市")
		So(err, ShouldBeNil)
		fmt.Println(domain)
		reAuthentication(userToken.Token, tx, pool)
	})
}
