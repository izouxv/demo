package storage

import (
	"fmt"
	"testing"

	"account-domain-rpc/module"

	. "github.com/smartystreets/goconvey/convey"

	. "account-domain-rpc/go-drbac/common"
)

func TestPermission_Create(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestPermission Create", t, func() {
		p := Permission{PermissionName: "添加应用", Url: "v1.1/domain/{did}/applicaton", Opt: "post"}
		err := p.Create(tx)
		So(err, ShouldBeNil)
		fmt.Println(p)
	})
}

func TestPermission_DeleteByPID(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestPermission DeleteByPID", t, func() {
		p := Permission{PermissionName: "添加应用", Url: "v1.1/domain/{did}/applicaton", Opt: "post"}
		err := p.Create(tx)
		So(err, ShouldBeNil)
		fmt.Println(p)
		err = p.DeleteByPID(tx)
		So(err, ShouldBeNil)
	})
}

func TestPermission_Update(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestPermission Update", t, func() {
		p := Permission{PermissionName: "添加应用", Url: "v1.1/domain/{did}/applicaton", Opt: "post"}
		err := p.Create(tx)
		So(err, ShouldBeNil)
		fmt.Println(p)
		p.PermissionName = "删除应用"
		p.Url = "v1.1/domain/{did}/applicaton"
		p.Opt = "delete"
		err = p.Update(tx)
		So(err, ShouldBeNil)
	})
}

func TestPermission_GetPermissionByID(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestPermission GetPermissionByID", t, func() {
		p := Permission{PermissionName: "添加应用", Url: "v1.1/domain/{did}/applicaton", Opt: "post"}
		err := p.Create(tx)
		So(err, ShouldBeNil)
		fmt.Println(p)
		err = p.GetPermissionByID(tx)
		So(err, ShouldBeNil)
		Convey("GetPermissionByID nil", func() {
			p.Pid = 1000
			err := p.GetPermissionByID(tx)
			So(err, ShouldEqual, ErrDoesNotExist)
		})
		fmt.Println(p)
	})
}

func TestDomain_GetDomainsByIDs(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestPermission GetPermissionByID", t, func() {
		ids := make([]int32, 0)
		for i := 0; i < 10; i++ {
			p := Permission{PermissionName: "添加应用", Url: "v1.1/domain/{did}/applicaton", Opt: "post"}
			err := p.Create(tx)
			So(err, ShouldBeNil)
			ids = append(ids, p.Pid)
		}
		p := Permission{}
		ps, err := p.GetDomainsByIDs(ids, tx)
		So(err, ShouldBeNil)
		fmt.Println(ps)
	})
}
