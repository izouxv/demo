package storage

import (
	"account-domain-rpc/module"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRolePermission_Create(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRolePermission_Create", t, func() {
		rp := RolePermission{Rid: 1, Pid: 1}
		err := rp.Create(tx)
		fmt.Println(err)
	})
}

func TestRolePermission_DeleteByRid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRolePermission_DeleteByRid", t, func() {
		rp := RolePermission{Rid: 1, Pid: 2}
		err := rp.Create(tx)
		fmt.Println(err)
		rp.DeleteByRid(tx)
	})
}

func TestRolePermission_DeleteByPid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRolePermission_DeleteByPid", t, func() {
		rp := RolePermission{Rid: 1, Pid: 2}
		err := rp.Create(tx)
		fmt.Println(err)
		rp.DeleteByPid(tx)
	})
}

func TestRolePermission_GetRolePermissionByRid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestRolePermission_GetRolePermissionByRid", t, func() {
		var rid int32 = 1
		Convey("Create", func() {
			rp := RolePermission{Rid: rid, Pid: 12}
			err := rp.Create(tx)
			So(err, ShouldBeNil)
			rp.Pid = 3
			err = rp.Create(tx)
			So(err, ShouldBeNil)
			rp.Pid = 4
			err = rp.Create(tx)
			So(err, ShouldBeNil)
			rp.Pid = 5
			err = rp.Create(tx)
			So(err, ShouldBeNil)
			rp.Pid = 6
			err = rp.Create(tx)
			So(err, ShouldBeNil)
			rp.Pid = 7
			err = rp.Create(tx)
			So(err, ShouldBeNil)
		})
		rp := RolePermission{Rid: rid}
		ids, err := rp.GetRolePermissionByRid(tx)
		So(err, ShouldBeNil)
		fmt.Println(ids)
	})
}

func TestRolePermission_GetRolePermissionInfoByRid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("测试获取角色对应的权限信息", t, func() {
		rp := RolePermission{Rid: 2}
		rpi, err := rp.GetRolePermissionInfoByRid(tx)
		So(err, ShouldBeNil)
		fmt.Println(rpi)
		for _, v := range rpi {
			fmt.Println(v)
		}
	})
}
