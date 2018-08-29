package storage

import (
	"fmt"
	"testing"
	"time"

	"account-domain-rpc/module"

	. "github.com/smartystreets/goconvey/convey"

	"account-domain-rpc/util"

	. "account-domain-rpc/go-drbac/common"
)

func TestUser_Create(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUser Create", t, func() {
		u := User{
			Username: "lids@radacat.com",
			Nickname: "lids",
			Password: "123456",
		}
		err := u.Create(tx)
		So(err, ShouldEqual, ErrAlreadyExists)
		fmt.Println(u)
		Convey("Create DoesNotExists", func() {
			u := User{
				Username: "lids@radacat" + string(util.Krand(20, util.KC_RAND_KIND_ALL)) + ".com",
				Nickname: "lids",
				Password: "123456",
			}
			err := u.Create(tx)
			So(err, ShouldBeNil)
		})
	})
}
func TestUser_GetUserByUsername(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUser GetUserByUsername Create", t, func() {
		u := User{
			Username: "lids@radacat" + string(util.Krand(20, util.KC_RAND_KIND_ALL)) + ".com",
			Nickname: "lids",
			Password: "123456",
		}
		err := u.Create(tx)
		ShouldBeNil(err)
		err = u.GetUserByUsername(tx)
		ShouldBeNil(err)
	})
}
func TestUser_GetUserByUID(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUser GetUserByUID", t, func() {
		u := User{
			Username: "lids@radacat" + string(util.Krand(20, util.KC_RAND_KIND_ALL)) + ".com",
			Nickname: "lids",
			Password: "123456",
		}
		err := u.Create(tx)
		ShouldBeNil(err)
		err = u.GetUserByUID(tx)
		ShouldBeNil(err)
	})
}
func TestUser_GetUserByUIDs(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	ids := make([]int64, 0)
	Convey("TestUser GetUserByUID", t, func() {
		for i := 0; i < 10; i++ {
			u := User{
				Username: "lids@radacat" + string(util.Krand(20, util.KC_RAND_KIND_ALL)) + ".com",
				Nickname: "lids",
				Password: "123456",
			}
			err := u.Create(tx)
			ShouldBeNil(err)
			ids = append(ids, u.Uid)
		}
		u := User{}
		users, err := u.GetUserByUIDs(ids, tx)
		ShouldBeNil(err)
		So(len(users), ShouldEqual, 10)
	})
}
func TestUser_UpdatePassword(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUser GetUserByUID", t, func() {
		u := User{
			Username: "lids@radacat" + string(util.Krand(20, util.KC_RAND_KIND_ALL)) + ".com",
			Nickname: "lids",
			Password: "123456",
		}
		err := u.Create(tx)
		ShouldBeNil(err)
		u.Password = "654321"
		time.Sleep(time.Second * 5)
		err = u.UpdatePassword(tx)
		ShouldBeNil(err)
	})
}
func TestUser_IsTruePassword(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestUser GetUserByUID", t, func() {
		u := User{
			Username: "lids@radacat" + string(util.Krand(20, util.KC_RAND_KIND_ALL)) + ".com",
			Nickname: "lids",
			Password: "123456",
		}
		err := u.Create(tx)
		ShouldBeNil(err)
		isTrue, err := u.IsTruePassword("123456", tx)
		ShouldBeNil(err)
		So(isTrue, ShouldBeTrue)
		u.Password = "654321"
		err = u.UpdatePassword(tx)
		ShouldBeNil(err)
		isTrue, err = u.IsTruePassword("654321", tx)
		ShouldBeNil(err)
		So(isTrue, ShouldBeTrue)
	})
}
