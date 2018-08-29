package storage

import (
	"fmt"
	"testing"
	"time"

	"account-domain-rpc/module"

	. "github.com/smartystreets/goconvey/convey"

	. "account-domain-rpc/go-drbac/common"
)

func TestUserDomainRole_Create(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole Create", t, func() {
		udr := UserDomainRole{Uid: time.Now().Unix(), Did: 1111, Rid: 5}
		err := udr.Create(tx)
		So(err, ShouldBeNil)
		udr.Did = 11
		udr.Rid = 1
		err = udr.Create(tx)
		So(err, ShouldBeNil)
		udr.Did = 12
		udr.Rid = 2
		err = udr.Create(tx)
		So(err, ShouldBeNil)
		udr.Did = 13
		udr.Rid = 1
		err = udr.Create(tx)
		So(err, ShouldBeNil)
	})
}

func TestUserDomainRole_DeleteByUid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole DeleteByUid", t, func() {
		uid := time.Now().Unix()
		udr := UserDomainRole{Uid: uid, Did: 1111, Rid: 1}
		err := udr.Create(tx)
		So(err, ShouldBeNil)
		udr.DeleteByUid(tx)
	})
}
func TestUserDomainRole_DeleteByDid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole DeleteByDid", t, func() {
		uid := time.Now().Unix()
		udr := UserDomainRole{Uid: uid, Did: 1111, Rid: 1}
		err := udr.Create(tx)
		So(err, ShouldBeNil)
		udr.DeleteByDid(tx)
	})
}
func TestUserDomainRole_DeleteByRid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole DeleteByDid", t, func() {
		uid := time.Now().Unix()
		udr := UserDomainRole{Uid: uid, Did: 1111, Rid: 1}
		err := udr.Create(tx)
		So(err, ShouldBeNil)
		udr.DeleteByRid(tx)
	})
}
func TestUserDomainRole_DeleteByUidAndDid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole DeleteByDid", t, func() {
		uid := time.Now().Unix()
		udr := UserDomainRole{Uid: uid, Did: 1111, Rid: 1}
		err := udr.Create(tx)
		So(err, ShouldBeNil)
		err = udr.DeleteByUidAndDid(tx)
		So(err, ShouldBeNil)

		time.Sleep(time.Second)
		uid = time.Now().Unix()
		udr = UserDomainRole{Uid: uid, Did: 1111, Rid: 1, IsDefault: true}
		err = udr.Create(tx)
		So(err, ShouldBeNil)
		err = udr.DeleteByUidAndDid(tx)
		So(err, ShouldEqual, ErrDoesNotExist)
	})
}
func TestUserDomainRole_GetUserDomainRoleByUid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole DeleteByDid", t, func() {
		uid := time.Now().Unix()
		Convey("TestUserDomainRole create", func() {
			udr := UserDomainRole{Uid: uid, Did: 1111, Rid: 5}
			err := udr.Create(tx)
			So(err, ShouldBeNil)
			udr.Did = 11
			udr.Rid = 1
			err = udr.Create(tx)
			So(err, ShouldBeNil)
			udr.Did = 12
			udr.Rid = 2
			err = udr.Create(tx)
			So(err, ShouldBeNil)
			udr.Did = 13
			udr.Rid = 1
			err = udr.Create(tx)
			So(err, ShouldBeNil)
		})
		udr := UserDomainRole{Uid: uid, Did: 1111, Rid: 1}
		udrs, err := udr.GetUserDomainRoleByUid(tx)
		So(err, ShouldBeNil)
		fmt.Println(udrs)
		for _, v := range udrs {
			fmt.Println(*v)
		}
	})
}
func TestUserDomainRole_GetUserDomainRoleByDid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole DeleteByDid", t, func() {
		var did int64 = 100
		Convey("TestUserDomainRole create", func() {
			udr := UserDomainRole{Uid: time.Now().Unix(), Did: did, Rid: 5}
			err := udr.Create(tx)
			So(err, ShouldBeNil)
			udr.Uid = time.Now().Unix() + 1
			err = udr.Create(tx)
			So(err, ShouldBeNil)
			udr.Uid = time.Now().Unix() + 2
			err = udr.Create(tx)
			So(err, ShouldBeNil)
			udr.Uid = time.Now().Unix() + 3
			err = udr.Create(tx)
			So(err, ShouldBeNil)
		})
		udr := UserDomainRole{Did: 1111}
		udrs, err := udr.GetUserDomainRoleByDid(tx)
		So(err, ShouldBeNil)
		fmt.Println(udrs)
		for _, v := range udrs {
			fmt.Println(*v)
		}
	})
}

func TestUserDomainRole_UpdateUserRoleByDidAndUid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole DeleteByDid", t, func() {
		uid := time.Now().Unix()
		var did int64 = 100
		udr := UserDomainRole{Uid: uid, Did: did, Rid: 5}
		err := udr.Create(tx)
		So(err, ShouldBeNil)
		udr.Rid = 3
		err = udr.UpdateUserRoleByDidAndUid(tx)
		So(err, ShouldBeNil)
	})
}

func TestUserDomainRole_GetUserDomainMaxRoleByUid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUserDomainRole GetUserDomainMaxRoleByUid", t, func() {
		uid := time.Now().Unix()
		var did int64 = 100
		udr := UserDomainRole{Uid: uid, Did: did, Rid: 5}
		err := udr.Create(tx)
		So(err, ShouldBeNil)

		udr = UserDomainRole{Uid: uid, Did: 101, Rid: 4}
		err = udr.Create(tx)
		So(err, ShouldBeNil)

		udr = UserDomainRole{Uid: uid, Did: 102, Rid: 4}
		err = udr.Create(tx)
		So(err, ShouldBeNil)

		max, err := udr.GetUserDomainMaxRoleByUid(tx)
		So(max, ShouldEqual, 4)
	})
}

func TestUserDomainRole_GetUserRoleByDid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("关联查询用户角色域信息", t, func() {
		udr := UserDomainRole{Did: 933262636978544640}
		userRoles, err := udr.GetUserRoleByDid(tx)
		So(err, ShouldBeNil)
		fmt.Println(userRoles)
		for _, v := range userRoles {
			fmt.Println(v)
			fmt.Println(v.Username)
			fmt.Println(v.RoleName)
		}
	})
}
