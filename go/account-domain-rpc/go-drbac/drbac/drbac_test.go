package drbac

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"account-domain-rpc/module"

	. "github.com/smartystreets/goconvey/convey"

	"account-domain-rpc/go-drbac/internal/storage"

	. "account-domain-rpc/go-drbac/common"
)

func TestDrbacServer_CreateUserBaseDomain(t *testing.T) {
	NewSouce()
	pool := module.RedisClient(module.Persistence)
	ds, err := NewDrbacServer(module.MysqlClient(), pool)
	fmt.Println(ds)
	if err != nil {
		return
	}
	Convey("TestDrbacServer_CreateUserBaseDomain", t, func() {
		username := "lids@radaca.com" + strconv.FormatInt(time.Now().Unix(), 10)
		password := "123456"
		nickname := "lids"
		var did int64 = 933262636924018688
		rid := DomianAdmin
		ds.CreateUserBaseDomain(username, password, nickname, did, rid)
	})
}

func TestDrbacServer_CreateDomainBaseDomain(t *testing.T) {
	NewSouce()
	pool := module.RedisClient(module.Persistence)
	ds, err := NewDrbacServer(module.MysqlClient(), pool)
	if err != nil {
		return
	}
	Convey("TestDrbacServer CreateDomainBaseDomain", t, func() {
		oid := 2
		did := time.Now().Unix()
		uid := time.Now().Unix()
		domain, err := ds.CreateDomainBaseDomain(int32(oid), int64(did), int64(uid), "CreateDomainBaseDomain")
		ShouldBeNil(err)
		fmt.Println(domain)
	})
}

func TestDrbac(t *testing.T) {
	NewSouce()
	pool := module.RedisClient(module.Persistence)
	ds, err := NewDrbacServer(module.MysqlClient(), pool)
	fmt.Println(ds)
	if err != nil {
		return
	}
	var oid int32 = 3
	user := &storage.User{}
	var uid int64 = 0
	domain := &storage.Domain{}
	Convey("TestDrbac", t, func() {
		//root用户在组织中用户及初始化一个域
		Convey("root用户在组织中用户及初始化一个域", func() {
			username := "SuperAdmin" + strconv.FormatInt(time.Now().Unix(), 10)
			password := "123456"
			nickname := "SuperAdmin"
			user, err = ds.CreateUserBaseOrganization(oid, username, password, nickname)
			So(err, ShouldBeNil)
			fmt.Println(user)
			uid = user.Uid
		})
	})
	//super用户在组织中创建域，及对应域中其他操作
	Convey("super用户在组织中创建域，及对应域中其他操作1", t, func() {
		fmt.Println("=====================1========================")
		//创建一个域
		domain, err = ds.CreateDomainBaseOrganization(oid, user.Uid, "浙江省")
		pid := domain.Did
		So(err, ShouldBeNil)
		//在上一个域中创建一个DomainAdmin 1
		Convey("在上一个域中创建一个DomainAdmin 1", func() {
			username := "DomianAdmin1.1" + strconv.FormatInt(time.Now().Unix(), 10)
			password := "123456"
			nickname := "DomianAdmin"
			var did int64 = domain.Did
			rid := DomianAdmin
			user, err = ds.CreateUserBaseDomain(username, password, nickname, did, rid)
			So(err, ShouldBeNil)
			fmt.Println(user)
			//domainAdmin在他的父域中创建一个子域 1
			func() {
				uid := user.Uid
				domain, err = ds.CreateDomainBaseDomain(oid, pid, uid, "杭州市")
				ShouldBeNil(err)
				fmt.Println(domain)
				//在子域中添加一个applicationAdmin
				func() {
					username = "applicationAdmin 1.1" + strconv.FormatInt(time.Now().Unix(), 10)
					password = "123456"
					nickname = "applicationAdmin"
					var did int64 = pid
					rid := ApplicationAdmin
					user, err = ds.CreateUserBaseDomain(username, password, nickname, did, rid)
					So(err, ShouldBeNil)
				}()
			}()
			//domainAdmin在他的父域中创建一个子域 2
			func() {
				uid := user.Uid
				domain, err = ds.CreateDomainBaseDomain(oid, pid, uid, "温州市")
				ShouldBeNil(err)
				fmt.Println(domain)
				//在子域中添加一个applicationAdmin
				func() {
					username = "applicationAdmin2.1" + strconv.FormatInt(time.Now().Unix(), 10)
					password = "123456"
					nickname = "applicationAdmin"
					var did int64 = domain.Did
					rid := ApplicationAdmin
					user, err = ds.CreateUserBaseDomain(username, password, nickname, did, rid)
					So(err, ShouldBeNil)
				}()
			}()
		})
	})
	//super用户在组织中创建域，及对应域中其他操作
	Convey("super用户在组织中创建域，及对应域中其他操作2", t, func() {
		fmt.Println("==================2===========================")
		//创建一个域
		domain, err = ds.CreateDomainBaseOrganization(oid, uid, "北京市")
		pid := domain.Did
		So(err, ShouldBeNil)
		//在上一个域中创建一个DomainAdmin
		Convey("在上一个域中创建一个DomainAdmin 2", func() {
			username := "DomianAdmin2.1" + strconv.FormatInt(time.Now().Unix(), 10)
			password := "123456"
			nickname := "DomianAdmin"
			var did int64 = domain.Did
			rid := DomianAdmin
			user, err = ds.CreateUserBaseDomain(username, password, nickname, did, rid)
			So(err, ShouldBeNil)
			fmt.Println(user)
			//domainAdmin在他的父域中创建一个子域 1
			time.Sleep(time.Second)
			func() {
				uid := user.Uid
				domain, err = ds.CreateDomainBaseDomain(oid, pid, uid, "朝阳区")
				ShouldBeNil(err)
				fmt.Println(domain)
				//在子域中添加一个applicationAdmin
				func() {
					username = "applicationAdmin2.1" + strconv.FormatInt(time.Now().Unix(), 10)
					password = "123456"
					nickname = "applicationAdmin"
					var did int64 = domain.Did
					rid := ApplicationAdmin
					user, err = ds.CreateUserBaseDomain(username, password, nickname, did, rid)
					So(err, ShouldBeNil)
				}()
			}()
			//domainAdmin在他的父域中创建一个子域 2
			func() {
				domain, err = ds.CreateDomainBaseDomain(oid, pid, uid, "海淀区")
				ShouldBeNil(err)
				fmt.Println(domain)
				//在子域中添加一个applicationAdmin
				func() {
					username = "applicationAdmin2.2" + strconv.FormatInt(time.Now().Unix(), 10)
					password = "123456"
					nickname = "applicationAdmin"
					var did int64 = domain.Did
					rid := ApplicationAdmin
					user, err = ds.CreateUserBaseDomain(username, password, nickname, did, rid)
					So(err, ShouldBeNil)
				}()
				//在子域中添加一个domainAdmin2
				func() {
					username = "domainAdmin2.2" + strconv.FormatInt(time.Now().Unix(), 10)
					password = "123456"
					nickname = "domainAdmin"
					var did int64 = domain.Did
					rid := DomianAdmin
					user, err = ds.CreateUserBaseDomain(username, password, nickname, did, rid)
					So(err, ShouldBeNil)
					func() {
						did := domain.Did
						uid := user.Uid
						domain, err = ds.CreateDomainBaseDomain(oid, did, uid, "中关村")
						ShouldBeNil(err)
					}()
				}()
			}()
		})
	})
}

func TestDrbacServer_Authentication(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	pool := module.RedisClient(module.Nopersistence)
	ds, err := NewDrbacServer(tx, pool)
	if err != nil {
		return
	}
	Convey("认证成功", t, func() {
		userToken, isTrue, err := ds.Authentication("SuperAdmin1511598477", "123456")
		So(err, ShouldBeNil)
		So(isTrue, ShouldBeTrue)
		fmt.Println(userToken)
		b, _ := json.Marshal(userToken)
		fmt.Println(string(b))

		Convey("用户名不存在", func() {
			_, isTrue, err := ds.Authentication("SuperAdmin", "123456")
			So(isTrue, ShouldBeFalse)
			So(err, ShouldEqual, ErrDoesNotExist)
		})

		Convey("密码错误，认证失败", func() {
			_, isTrue, err := ds.Authentication("SuperAdmin1511598477", "12345678")
			So(isTrue, ShouldBeFalse)
			So(err, ShouldBeNil)
		})
	})
}

func TestDrbacServer_Authorization(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	pool := module.RedisClient(module.Nopersistence)
	ds, err := NewDrbacServer(tx, pool)
	if err != nil {
		return
	}
	Convey("认证+授权", t, func() {
		//userToken, isTrue, err := ds.Authentication("SuperAdmin1511342126", "123456")
		//So(err, ShouldBeNil)
		//So(isTrue, ShouldBeTrue)
		//fmt.Println(userToken)

		userToken, isTrue := ds.Authorization(933262636924018688, "/session", "post", "78953ee6ab778c035ea3679c651e950e")
		So(isTrue, ShouldBeTrue)
		fmt.Println(userToken)

		userToken, isTrue = ds.Authorization(933262636924018688, "/session", "post", "78953ee6ab778c035ea3679c65")
		So(isTrue, ShouldBeFalse)
		fmt.Println(userToken)
	})
}

//获取域中的用户信息
func TestDrbacServer_GetUserBaseDomain(t *testing.T) {
	NewSouce()
	tx := module.MysqlClient()
	pool := module.RedisClient(module.Nopersistence)
	ds, err := NewDrbacServer(tx, pool)
	if err != nil {
		return
	}
	Convey("获取域中的用户信息", t, func() {
		domainUser, err := ds.GetUserBaseDomain(933262636978544640)
		So(err, ShouldBeNil)
		fmt.Println(domainUser)
		for _, v := range domainUser.UserRoleInfos {
			fmt.Println(v.User)
			fmt.Println(v.Role)
			for _, w := range v.Permissions {
				fmt.Println(w.Url, " ", w.Opt)
			}
		}
	})
}
