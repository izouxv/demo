package storage

import (
	"account-domain-rpc/module"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRole_Create(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRole Create", t, func() {
		r := Role{RoleName: "super"}
		err := r.Create(tx)
		So(err, ShouldBeNil)
		fmt.Println(r)
	})
}
func TestRole_DeleteByRID(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRole Create", t, func() {
		r := Role{RoleName: "super"}
		err := r.Create(tx)
		So(err, ShouldBeNil)
		fmt.Println(r)
		err = r.DeleteByRID(tx)
		So(err, ShouldBeNil)
		DeleteRoleMapByRid(r.Rid)
	})
}

func TestRole_GetRoleByRid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRole Create", t, func() {
		r := Role{RoleName: "super"}
		err := r.Create(tx)
		So(err, ShouldBeNil)
		fmt.Println(r)
		err = r.GetRoleByRid(tx)
		So(err, ShouldBeNil)
		fmt.Println(r)
	})
}

func TestRole_GetRoles(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRole Create", t, func() {
		r := Role{RoleName: "super"}
		err := r.Create(tx)
		So(err, ShouldBeNil)

		r.Rid = 0
		r.RoleName = "manager"
		err = r.Create(tx)
		So(err, ShouldBeNil)

		r.Rid = 0
		r.RoleName = "readonly"
		err = r.Create(tx)
		So(err, ShouldBeNil)

		fmt.Println(r)
		roles, err := r.GetRoles(tx)
		So(err, ShouldBeNil)
		fmt.Println(roles)
	})
}
func TestGetRoleNameFromRoleMapByRid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRole Create", t, func() {
		r := Role{RoleName: "super"}
		err := r.Create(tx)
		So(err, ShouldBeNil)
		roleName, err := GetRoleNameFromRoleMapByRid(r.Rid, tx)
		ShouldBeNil(err)
		So(roleName, ShouldEqual, "super")
	})
}
