package drbac

import (
	"account-domain-rpc/module"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func Test_Authorization(t *testing.T) {
	NewSouce()
	pool := module.RedisClient(module.Persistence)
	Convey("认证权限", t, func() {
		userToken, isTrue := authorization(1, "domain", "post", "adasf", pool)
		So(isTrue, ShouldBeFalse)
		fmt.Println(userToken)
	})
}

func Test_MatchString(t *testing.T) {
	match, _ := regexp.MatchString("/v1.0/domains/([0-9])+/nodes", "/v1.0/domains/12/nodes")
	fmt.Println(match)
}
