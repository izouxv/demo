package upgrade

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"

	. "account-domain-rpc/common"
	"account-domain-rpc/module"
	"account-domain-rpc/service/air"
	"account-domain-rpc/storage"
	"account-domain-rpc/util"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	DefalutDomian      int64 = 4
	DefalutApplication int64 = 1
)

const (
	Switch1 int32 = iota
	Switch2
)

// 对应版本（主、次）的枚举
type VersionType int32

const (
	Version1 VersionType = iota
	Version2
)

// 版本类型的枚举（空气净化器，网关，node）
type UpgradeEnum int32

const (
	AirMain UpgradeEnum = 1 + iota
	AirConnection
)

// 其他类型的版本名称，后续可以增加
var UpgradeType map[UpgradeEnum]string = map[UpgradeEnum]string{
	AirMain:       "空气净化器主控板",
	AirConnection: "空气净化器通信模块",
}

// HandlerUpgrade
func HandlerUpgrade(bytes []byte) (string, error) {
	um, nv, err := getNodeAndUpgradeModele(bytes)
	if err != nil {
		log.Error(err)
		return "", err
	}
	go heartbeat(nv.Nid)
	resp, err := upgarade(nv, &um.Data)
	if err != nil {
		return "", err
	}
	//控制下发的量
	if !controlFlow() {
		return "", errors.New("control Flow error wait")
	}
	return resp, nil
}

// getNodeAndUpgradeModele获取设备版本信息及上传版本
func getNodeAndUpgradeModele(decodeBytes []byte) (*UpgradeModele, *storage.NodeVersion, error) {
	um := UpgradeModele{}
	if err := um.UnmarshalBinary(decodeBytes); err != nil {
		log.Error(decodeBytes, err)
		return nil, nil, err
	}
	log.Info(fmt.Printf("%#v\n", um))
	devEUI := um.Data.DevEUI[:]
	node := storage.Node{DevEUI: devEUI}
	if err := node.GetNodeByDevEUI(); err != nil {
		// 插入node信息
		node.Name = "空气净化器"
		node.ApplicationID = DefalutApplication
		node.Version = showVersion(um.Data.MainVersion, um.Data.ComVersion, AirMain)
		node.Category = Air
		node.Mac = showMac(devEUI)
		node.Upgrade = false
		node.Is_active = true
		node.Did = DefalutDomian
		node.Description = "默认添加空气净化器"
		node.AppKey = []byte{0, 0, 0, 0, 0, 0, 0, 0}
		node.RXDelay = 0
		node.RX1DROffset = 0
		node.RXWindow = 0
		node.RX2DR = 0
		node.RelaxFCnt = false
		node.ADRInterval = 0
		node.InstallationMargin = 0
		node.IsABP = false
		node.IsClassC = true
		node.UseApplicationSettings = true
		err = storage.Transaction(module.MysqlClient(), func(tx *gorm.DB) error {
			err = node.Create(tx)
			if err != nil {
				log.Errorf("创建node异常，error is %s,node is %#v", err, node)
				return err
			}
			/*todo 创建设备的其他事物处理*/
			acl := &storage.Acl{Nid: node.Nid, Username: ShowAclUsernameByDevEUI(node.DevEUI), Pw: ShowAclPasswordByDevEUI(node.DevEUI), Topic: ShowAclTopicByDevEUI(node.DevEUI)}
			err = acl.CreateAcl(tx)
			if err != nil {
				log.Errorf("创建acl异常，error is %s,acl is %#v", err, acl)
				return err
			}
			nv := &storage.NodeVersion{Nid: node.Nid}
			err = nv.Create(tx)
			if err != nil {
				log.Errorf("创建nodeVersion 异常，error is %s,nodeVersion is %#v", err, nv)
				return err
			}
			return nil
		})
		if err != nil {
			return nil, nil, err
		}

		log.Info("Upgrade Add Node %+v \n", node)
	}
	if node.Version != showVersion(um.Data.MainVersion, um.Data.ComVersion, AirMain) {
		node.Version = showVersion(um.Data.MainVersion, um.Data.ComVersion, AirMain)
		if err := node.UpdateNodeUpgrade(); err != nil {
			log.Errorf("Update Node version error %+v \n", node)
			return nil, nil, err
		}
	}
	log.Infof("node info  %+v \n", node)
	nv := storage.NodeVersion{Nid: node.Nid}
	if err := nv.GetByNid(); err != nil {
		return nil, nil, err
	}
	log.Infof("node Version +v\n", nv)
	return &um, &nv, nil
}

// heartbeat设备在线状态心跳
func heartbeat(id int64) {
	client := module.RedisClient(module.Persistence).Get()
	defer client.Close()
	node := storage.Node{Nid: id}
	err := node.GetNodeByNid()
	key := fmt.Sprintf("%s:%d:%d", module.HeartbeatKeyPrefixSet, node.Did, node.ApplicationID)
	resp, err := client.Do("ZADD", key, time.Now().Unix(), id)
	if err != nil {
		log.Error(err)
	}
	log.Info(resp)
}

// upgarade 升级逻辑，返回对应的升级文件内容的编号
func upgarade(nv *storage.NodeVersion, um *UpgradeDataModel) (string, error) {
	// 升级的版本对应的
	resp := ""
	var err error
	u := &storage.Upgrade{}
	ok := false
	if nv.Switch1 {
		// 主版本开关开启，优先升级
		log.Info("主版本开关开启，优先升级")
		u, ok = isUpgrade(nv.Version1, um.MainVersion)
		if ok {
			log.Infof("对应数据库中的升级信息: #v", u)
			resp, err = returnResource(u, Version2)
			if err != nil {
				return "", err
			}
			return resp, nil
		} else {
			// 关闭开关
			log.Info("版本号相同，已经完成升级任务，关闭开关1")
			if u != nil {
				cid := u.Category
				closeSwitch(nv, um, Switch1, UpgradeEnum(cid))
			}
		}
	} else if nv.Switch2 {
		// 次版本开关开启
		log.Info("次版本开关开启,升级")
		u, ok = isUpgrade(nv.Version2, um.ComVersion)
		if ok {
			log.Infof("对应数据库中的升级信息: #v", u)
			resp, err = returnResource(u, Version1)
			if err != nil {
				return "", err
			}
			return resp, nil
		} else {
			// 关闭开关
			log.Info("版本号相同，已经完成升级任务，关闭开关2")
			if u != nil {
				cid := u.Category
				closeSwitch(nv, um, Switch2, UpgradeEnum(cid))
			}
		}
	} else {
		log.Errorf("开关都关闭,不做操作")
		// 开关都关闭,不做操作
		return resp, errors.New("all switch close")
	}
	return resp, errors.New("close switch")
}

// isUpgrade判断版本是否可以升级
func isUpgrade(version int32, newVersion [3]byte) (*storage.Upgrade, bool) {
	u, err := storage.GetUpgrade(version)
	if err != nil || u == nil {
		log.Errorf("获取版本为空")
		return nil, false
	}
	oldVersion, err := showVersionCode(u.VersionCode)
	log.Infof("上传的版本号：%d，数据库中的版本：%s", newVersion, u.VersionCode)
	if oldVersion != newVersion {
		// 可以升级or回退办版本
		return u, true
	}
	log.Errorf("版本号相同，不需要升级")
	return u, false
}

// returnResource获取升级资源
func returnResource(u *storage.Upgrade, ut VersionType) (string, error) {
	version, err := showVersionCode(u.VersionCode)
	if err != nil {
		log.Errorf("没有对应的版本号")
		return "", err
	}
	urm := &UpgradeResultModel{Type: air.UpgradeOrderAck}
	urdm := UpgradeResultDataModel{UpgradeType: byte(ut)}
	urdm.URL = []byte(u.URL)
	urdm.URLLen = byte(len(urdm.URL))
	urdm.Version = version
	md5, err := util.Md5StringtoByte(u.MD5)
	if err != nil {
		return "", err
	}
	for i, v := range md5 {
		urdm.MD5[i] = v
	}
	urm.Data = urdm
	urm.SetLen()
	if err := urm.SetCRC(); err != nil {
		return "", err
	}
	log.Infof("对应升级的主次版本：%d", ut)
	log.Infof("升级文件下载地址: [%s]", u.URL)
	resp, err := urm.String()
	if err != nil {
		return "", err
	}
	return resp, err
}

// closeSwitch关闭升级开关
func closeSwitch(nv *storage.NodeVersion, um *UpgradeDataModel, flag int32, enum UpgradeEnum) error {
	var err error
	switch flag {
	case Switch1:
		err = nv.CloseSwitch1()
	case Switch2:
		err = nv.CloseSwitch2()
	}
	if err != nil {
		return err
	}
	node := storage.Node{Nid: nv.Nid, Version: showVersion(um.MainVersion, um.ComVersion, enum), Upgrade: false}
	err = node.UpdateNodeUpgrade()
	if err != nil {
		return err
	}
	return nil
}

// showVersion显示版本号的格式
func showVersion(version1, version2 [3]byte, enum UpgradeEnum) string {
	switch enum {
	case AirConnection, AirMain:
		return fmt.Sprintf("控制板:%d.%d.%d 通信模块:%d.%d.%d",
			version1[0], version1[1], version1[2],
			version2[0], version2[1], version2[2])
	}
	return ""
}

// showMac显示mac的格式[48:76:88:99:00:00]
func showMac(devEUI []byte) string {
	mac := ""
	for i, v := range devEUI[:] {
		if i > 5 {
			return mac
		}
		mac += hex.EncodeToString([]byte{v})
		if i < 5 {
			mac += ":"
		}
	}
	return mac
}

// showVersionCode将字符类型的版本号转换为byte类型
func showVersionCode(versionCode string) ([3]byte, error) {
	var out [3]byte
	v := strings.Split(versionCode, ".")
	if len(v) != 3 {
		return out, errors.New("version error")
	}
	for i, v := range v {
		u, err := strconv.Atoi(v)
		if err != nil {
			return out, err
		}
		out[i] = byte(u)
	}
	return out, nil
}

// controlFlow控制下载量
func controlFlow() bool {
	client := module.RedisClient(module.Nopersistence).Get()
	defer client.Close()
	count, err := redis.Int(client.Do("Get", module.UpgradeCount))
	if err != nil && err != redis.ErrNil {
		log.Error(err)
		return false
	}
	if count > 50 {
		log.Error("wait for upgrade")
		return false
	}
	client.Send("MULTI")
	client.Send("INCR", module.UpgradeCount)
	client.Send("EXPIRE", module.UpgradeCount, 1) //s
	if _, err = client.Do("EXEC"); err != nil {
		log.Error(err)
		return false
	}
	return true
}
