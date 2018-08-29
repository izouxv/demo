package util

import (
	"regexp"
)

//正则验证规则
var (
	REGEXP_MOBILE   = regexp.MustCompile("^1[3|4|5|6|7|8|9][0-9]{9}$")
	REGEXP_MAIL     = regexp.MustCompile("^\\w+([-+._]\\w+)*@\\w+([-.]\\w+)*\\.[a-z]+([-.][a-z]+)*$")
	REGEXP_NICKNAME = regexp.MustCompile("^.{1,15}$")
	REGEXP_PWD      = regexp.MustCompile("^[a-z0-9A-Z]{6,32}$")
	REGEXP_SN       = regexp.MustCompile("^[a-z0-9A-Z]{1,32}$")
)

//常用字符串
const (
	ReqMethodJson   = "application/json;charset=UTF-8"
	ReqMethodBinary = "application/octet-stream"
	path = "/data1/upload/petfone"

	Swagger = "/swagger/"
	Conf = path + "/config/conf.yaml"
	SeeLog = path + "/config/http-seelog.xml"
	FilePath = "/temp/"
)

//性别
const (
	Gender1 = iota + 1
	Gender2
	Gender3
	Gender4
)
