package storage

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"account-domain-rpc/module"

	. "account-domain-rpc/go-drbac/common"

	. "github.com/smartystreets/goconvey/convey"
)

func NewSouce() {
	module.NewMysqlClient("192.168.1.6", "3306", "drbac", "root", "Radacat2017")
	module.NewRedisClient(module.Persistence, "192.168.1.6", "6379", 10, 10, "radacat1234")
}

func TestDomain_MarshalJson(t *testing.T) {
	Convey("TestDomain MarshalJson", t, func() {
		d := Domain{Did: 1, DomainName: "helle世界", Pid: 123, CreateTime: time.Now(), UpdateTime: time.Now()}
		data, err := d.MarshalJson()
		ShouldBeNil(err)
		fmt.Println(string(data))
	})
}

func TestDomain_UnmarshalJson(t *testing.T) {
	Convey("TestDomain MarshalJson", t, func() {
		d := Domain{Did: 1, DomainName: "helle世界", Pid: 123, CreateTime: time.Now(), UpdateTime: time.Now()}
		data, err := d.MarshalJson()
		ShouldBeNil(err)
		fmt.Println(string(data))
		domain := Domain{}
		err = domain.UnmarshalJson(data)
		ShouldBeNil(err)
	})
}

func TestDomain_Create(t *testing.T) {
	NewSouce()
	Convey("TestDomain_Create", t, func() {
		d := Domain{DomainName: "test", Pid: 1}
		for i := 0; i < 10; i++ {
			d.Create(module.MysqlClient())
			fmt.Println(d.Did)
		}
	})
}

func TestDomain_DeleteByDID(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain_Delete", t, func() {
		d := Domain{DomainName: "test", Pid: 1}
		d.Create(tx)
		deleteDomain := Domain{Did: d.Did}
		err := deleteDomain.DeleteByDID(tx)
		So(err, ShouldBeNil)
		Convey("Delete error", func() {
			d := Domain{Did: 1111}
			err := d.DeleteByDID(module.MysqlClient())
			fmt.Println(err)
			So(err, ShouldEqual, ErrDoesNotExist)
		})
	})
}

func TestDomain_DeleteByPid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain DeleteByPid", t, func() {
		d := Domain{DomainName: "test", Pid: 1}
		err := d.Create(tx)
		So(err, ShouldBeNil)
		err = d.DeleteByPid(tx)
		So(err, ShouldBeNil)
		Convey("Test DeleteByPid err", func() {
			d := Domain{Pid: 10000}
			err := d.DeleteByPid(tx)
			So(err, ShouldEqual, ErrDoesNotExist)
		})
	})
}

func TestDomain_DeleteByDIDs(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain DeleteByDIDs", t, func() {
		ids := make([]int64, 0)
		for i := 0; i < 10; i++ {
			d := Domain{DomainName: "test", Pid: 1}
			err := d.Create(tx)
			So(err, ShouldBeNil)
			ids = append(ids, d.Did)
		}
		d := Domain{}
		err := d.DeleteByDIDs(ids, tx)
		So(err, ShouldBeNil)

		ids = make([]int64, 0)
		err = d.DeleteByDIDs(ids, tx)
		So(err, ShouldEqual, ErrDoesNotExist)
	})
}

func TestDomain_Update(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain Update", t, func() {
		d := Domain{DomainName: "test", Pid: 1}
		err := d.Create(tx)
		So(err, ShouldBeNil)
		d.DomainName = "update-domain"
		err = d.Update(tx)
		So(err, ShouldBeNil)
		d.Did = 10000
		err = d.Update(tx)
		So(err, ShouldEqual, ErrDoesNotExist)
	})
}

func TestDomain_GetByID(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain GetByID", t, func() {
		d := Domain{DomainName: "test-GetByID", Pid: 1}
		err := d.Create(tx)
		So(err, ShouldBeNil)
		err = d.GetByID(tx)
		So(err, ShouldBeNil)
		So(d.DomainName, ShouldEqual, "test-GetByID")
	})
}

func TestDomain_GetDomainsByPid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain GetDomainsByIDs", t, func() {
		d := Domain{DomainName: "test", Pid: 0}
		err := d.Create(tx)
		So(err, ShouldBeNil)
		var Pid = d.Did
		for i := 0; i < 10; i++ {
			d := Domain{DomainName: "test", Pid: Pid}
			err := d.Create(tx)
			So(err, ShouldBeNil)
		}
		d = Domain{Pid: Pid}
		domians, err := d.GetDomainsByPid(tx)
		So(err, ShouldBeNil)
		fmt.Println(domians)
		So(len(domians), ShouldEqual, 10)
	})
}

func TestDomain_GetDomainCountByPid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("测试基于父域查询子域数量", t, func() {
		domain := Domain{Pid: 933262636978544640}
		count, err := domain.GetDomainCountByPid(tx)
		So(err, ShouldBeNil)
		fmt.Println(count)
	})
}

func TestDomain_IsExistDomainByPid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("测试判断父域下是否有子域", t, func() {
		domain := Domain{Pid: 933262636978544640}
		exist := domain.IsExistDomainByPid(tx)
		So(exist, ShouldBeTrue)
	})
}

func TestDomain_GetPidByID(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain GetPidByID", t, func() {
		var Pid int64 = 1
		d := Domain{DomainName: "test", Pid: Pid}
		err := d.Create(tx)
		So(err, ShouldBeNil)
		pid, err := d.GetPidByDID(tx)
		So(err, ShouldBeNil)
		So(pid, ShouldEqual, Pid)
	})
}

func TestDomain_GetIDsByPid(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain_GetIDsByPid", t, func() {
		d := Domain{DomainName: "test", Pid: 1}
		err := d.Create(tx)
		So(err, ShouldBeNil)
		ids, err := d.GetIDsByPid(tx)
		So(err, ShouldBeNil)
		fmt.Println(ids)
	})
}

func TestDomain_GetDomainDepth(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain GetDomainDepth", t, func() {
		d := Domain{DomainName: "0", Pid: 0}
		err := d.Create(tx)
		So(err, ShouldBeNil)
		var Pid = d.Did
		for i := 0; i < 4; i++ {
			d := Domain{DomainName: "test" + strconv.Itoa(i), Pid: Pid}
			err := d.Create(tx)
			Pid = d.Did
			So(err, ShouldBeNil)
		}
		d.Did = Pid
		depth, err := d.GetDomainDepth(tx)
		fmt.Println(depth)
	})
}

func TestDomain_IsExceedMaxDepth(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	Convey("TestDomain IsExceedMaxDepth", t, func() {
		var Pid int64
		//第一级域
		fooDomain := Domain{Did: Pid}
		isExceed, err := fooDomain.IsExceedMaxDepth(tx)
		So(err, ShouldEqual, ErrDoesNotExist)
		d := Domain{DomainName: "浙江省", Pid: Pid}
		err = d.Create(tx)
		So(err, ShouldBeNil)

		//第二级域
		fooDomain = Domain{Did: d.Did}
		isExceed, err = fooDomain.IsExceedMaxDepth(tx)
		So(err, ShouldBeNil)
		So(isExceed, ShouldBeFalse)
		d.Pid = fooDomain.Did
		d.DomainName = "杭州市"
		err = d.Create(tx)
		So(err, ShouldBeNil)

		////第三级域
		//fooDomain = Domain{Did:d.Did}
		//isExceed,err = fooDomain.IsExceedMaxDepth(tx)
		//d.Pid = fooDomain.Did
		//d.DomainName = "余杭区"
		//err = d.Create(tx)
		//So(err,ShouldBeNil)
		//
		////第四级域
		//d.Pid = Pid
		//d.DomainName = "仁和街道"
		//d.Create(tx)
		//Pid = d.Did
		//isExceed,err = d.IsExceedMaxDepth(tx)
		//So(err,ShouldBeNil)
		//So(isExceed,ShouldBeFalse)
		//
		////第五级域
		//d.Pid = Pid
		//d.Create(tx)
		//Pid = d.Did
		//isExceed,err = d.IsExceedMaxDepth(tx)
		//So(err,ShouldBeNil)
		//So(isExceed,ShouldBeTrue)
	})
}
