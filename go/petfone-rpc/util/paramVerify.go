package util

import (
	log "github.com/cihub/seelog"
	"github.com/json-iterator/go"
)

const (
	path = "/data1/upload/petfone"
	Conf = path + "/config/conf.yaml"
	SeeLog = path + "/config/rpc-seelog.xml"

	//TODO redis-key
	SendMobileCode		 = ":petfone:mobilecode:"    //手机发送验证码
	ResetPasswordByMail	 = ":petfone:resetpwd:mail:" //邮箱重置密码token
	ResetPasswordByPhone = ":petfone:resetpwd:mobile:" //手机重置密码token
	LoginSession         = ":petfone:login:"    //用户登录
	RegisterAgent		 = "petfone:agent"      //注册agent
	PetTrainId	 		 = "Ag:petfone:pet:" 	//用户id
	UDevices			 = "Ag:petfone:uid:dids:"	//用户-设备权限
	DUsers				 = "Ag:petfone:did:uids:"	//设备-用户权限
	UPets				 = "Ag:petfone:uid:pids:"	//用户-宠物权限
	PUsers				 = "Ag:petfone:pid:uids:"	//宠物-用户权限
	PetDeviceShareToken  = "Ag:petfone:share:"	//设备分享token
	UidRedisKey 		 = "Ag:petfone:uid" 	//用户id
	BigIdRedisKey 		 = "Ag:petfone:table:id"		//记录id
	LimitMobil           = ":petfone:mobile:limit"  //限制手机发送验证码的次数
	LimitMobil1           = ":petfone:mobile:ip"  //限制手机发送验证码的次数
	LimitIp              = ":petfone:ip:limit"  //限制IP访问次数
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
 	TagCnKeys = []string{" ~~ "}
 	TagCnValues = []string{"我还没学会呢~~"}
	TagEnKeys = []string{" ~~ "}
	TagEnValues = []string{"I haven't learned yet~~"}
	HttpApiLogFilePrefix = "/var/log/golang/petfone/httpApi.log."
	LogWriter,LogErr = log.NewFileWriter(HttpApiLogFilePrefix)

 	//todo redis命令
	SetStr	= "set"
	GetStr  = "get"
	Incr	= "incr"

	Expire  = "EX"
	Zadd	= "zadd"
	Zscore	= "zscore"
	Zcard	= "zcard"
	Zrem	= "zrem"
	Zrevrangebyscore = "zrevrangebyscore"
)

//对字符串参数判空与去除首尾空格
func VerifyParamsStr(params ...string) bool {
	log.Info("VerifyParamsStr:", params)
	for _, param := range params {
		if param == "" {
			return true
		}
		//strings.Replace(param, " ", "", -1)
		//strings.Replace(param, "\n", "", -1)
	}
	return false
}

// 对参数进行非空判断
func VerifyParamsEmpty(params ...string) bool {
	for i := 0; i < len(params); i++ {
		//log.Info("param:",params[i])
		if params[i] == "" {
			return true
		}
	}
	return false
}

// 对参数进行int32判断
func VerifyParamsUInt32(params ...int32) bool {
	lens := len(params)
	for i := 0; i < lens; i++ {
		if params[i] <= 0 {
			return true
		}
	}
	return false
}

// 对参数进行int64判断
func VerifyParamsUInt64(params ...int64) bool {
	lens := len(params)
	for i := 0; i < lens; i++ {
		if params[i] <= 0 {
			return true
		}
	}
	return false
}

// 对参数进行float32判断
func VerifyParamsFloat32(params ...float32) bool {
	lens := len(params)
	for i := 0; i < lens; i++ {
		if params[i] < 0.0 {
			return true
		}
	}
	return false
}

// 对参数进行float64判断
func VerifyParamsFloat64(params ...float64) bool {
	lens := len(params)
	for i := 0; i < lens; i++ {
		if params[i] < 0.0 {
			return true
		}
	}
	return false
}


//对参数类型判断
func VerifyTypes(t interface{}) int32 {
	switch t.(type) {
	case string:
		return 1
	case int32:
		return 2
	case int64:
		return 3
	case int:
		return 4
	case uint32:
		return 5
	case bool:
		return 6
	default:
		return 0
	}
}
