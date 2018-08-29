package util

import (
	"bufio"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"math/rand"
	"net/http"
	"petfone-http/pb"
	"petfone-http/po"
	"petfone-http/rpc"
	"time"
	"strconv"
	"github.com/json-iterator/go"
	"fmt"
	"crypto/md5"
	"os"
	"bytes"
	"strings"
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
)

//HTTP获取上下文
func GetContext(req *http.Request) *pb.SsoRequest {
	conSso := req.Context().Value("contextSso").(*pb.SsoRequest)
	return conSso
}

//处理HTTP请求中的数据
func GetHttpData(req *http.Request, types string, t interface{}) bool {
	body, bodyErr := ioutil.ReadAll(bufio.NewReader(req.Body))
	if bodyErr != nil {
		log.Error("GetHttp-ReadAll-bodyErr:", bodyErr)
		return true
	}
	size := len(body)
	log.Info("GetHttp-bodyLen:", size)
	if size < 2 {
		return true
	}
	if size < 400 {
		log.Info("GetHttp-bodyStr:", string(body))
	}
	if err := Json.Unmarshal(body, &t); err != nil {
		log.Info("GetHttp-Unmarshal-err:", err)
		return true
	}
	return false
}

// 对参数进行非空判断
func VerifyParamsStr(params ...string) bool {
	for i := 0; i < len(params); i++ {
		if params[i] == "" {
			return true
		}
	}
	return false
}

// 对参数进行uint32判断
func VerifyParamsUInt32(params ...int32) bool {
	lens := len(params)
	for i := 0; i < lens; i++ {
		if params[i] < 0 {
			return true
		}
	}
	return false
}

// 对参数进行int64判断
func VerifyParamsUInt64(params ...int64) bool {
	lens := len(params)
	for i := 0; i < lens; i++ {
		if params[i] < 0 {
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

//对用户名进行正则验证
func VerifyUsername(username string) bool {
	if REGEXP_MAIL.MatchString(username) {
		return false
	}
	if REGEXP_MOBILE.MatchString(username) {
		return false
	}
	return true
}

//对性别判断
func VerifyGender(gender int32) bool {
	if gender < 0 {
		return true
	}
	switch gender {
	case 0:
	case Gender1:
	case Gender2:
	case Gender3:
	case Gender4:
	default:
		return true
	}
	return false
}

//对密码进行正则验证
func VerifyPassword(str string) bool {
	if REGEXP_PWD.MatchString(str) {
		return false
	}
	return true
}

//对昵称进行正则验证
func VerifyNickname(str string) bool {
	if REGEXP_NICKNAME.MatchString(str) {
		return false
	}
	return true
}

//对设备SN进行正则验证
func VerifySN(sn string) bool {
	if REGEXP_SN.MatchString(sn) {
		return false
	}
	return true
}

//字符串长度校验
func VerifyLenParams(strs ...string) bool {
	for _, v := range strs {
		if len(v) > 80 {
			return true
		}
	}
	return false
}

//验证用户名是否存在   true表示用户存在,false表示用户不存在
func CheckUsername(sso *pb.SsoRequest) bool {
	log.Info("CheckUsername:", sso)
	ssoRC := rpc.SsoRpc(sso, "GetUserByName")
	log.Info("GetUserByName-code:", ssoRC.Code)
	if ssoRC.Code == 10000 {
		return true
	}
	if ssoRC.Code != 33002 {
		return true
	}
	return false
}

//对login/sso进行属性处理
func LoginTssoR(login *po.LoginPo, ssoR *pb.SsoReply) {
	login.Uid = ssoR.Uid
	login.Username = ssoR.Username
	login.Nickname = ssoR.Nickname
	login.Token = ssoR.SessionName
	login.LoginState = ssoR.LoginState
	login.State = ssoR.State
	login.Code = ssoR.Code
}

//对login/account进行属性处理
func LoginTaccountR(login *po.LoginPo, accountR *pb.AccountReply) {
	login.Phone = accountR.Phone
	login.Email = accountR.Email
	login.Gender = accountR.Gender
	login.Birthday = accountR.Birthday
	login.Avatar = accountR.Avatar
	login.Signature = accountR.Signature
	login.Address = accountR.Address
	login.Radius = accountR.Radius
	login.Map = accountR.Map
}

//对sso/ssoR进行属性处理
func SsoTssoR(sso *pb.SsoRequest, ssoR *pb.SsoReply) {
	sso.Uid = ssoR.Uid
	sso.Username = ssoR.Username
	sso.Nickname = ssoR.Nickname
	sso.State = ssoR.State
}

//获取随机字符串
const (
	synchrolock        = 0 //0 释放 1 加锁
	KC_NUMS            = 0 // 纯数字
	KC_LOWER_LETTERS   = 1 // 小写字母
	KC_UPPER_LETTERS   = 2 // 大写字母
	KC_NUMBERS_LETTERS = 3 // 数字、大小写字母
)

//生成随机字节
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

//计算幂
func powerInt(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * powerInt(x, n-1)
}

//时间戳格式化为字符串-分割
func TimeToStr_() string {
	return time.Now().Format("2006-01-02-15-04-05")
}

//时间戳格式化为字符串
func TimeToStr() string {
	return time.Now().Format("2006年01月02 15:04:05")
}

//int to string
func IntToStr(num int) string {
	return strconv.Itoa(num)
}

//int32 to string
func Int32ToStr(num int32) string {
	return strconv.Itoa(int(num))
}

//string to int32
func StrToInt32(str string) (int32, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}

//int64 to int32
func Int64ToInt32(num64 int64) int32 {
	num, _ := strconv.Atoi(strconv.FormatInt(num64, 10))
	return int32(num)
}

//string to int64
func StrToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	return num, err
}

//int64 to string
func Int64ToStr(num64 int64) string {
	return strconv.FormatInt(num64, 10)
}

//float32转str
func FloatToStr(float float64,size int) string {
	return strconv.FormatFloat(float,'E',-1,size)
}

func Int64DifferenceAbs(a, b int64) int64 {
	if a >= b {
		return a-b
	}
	return b-a
}

func PrintStr(a string) string {
	return strings.Replace(fmt.Sprintf("%q",a),"\"","",-1)
}

//字符串拼接
func StrAdd(a ...string) string {
	var bu bytes.Buffer
	for _, v := range a {
		bu.WriteString(v)
	}
	return bu.String()
}

/**
获取当前时间,秒
 */
func GetNowTimeInt64() int64 {
	return time.Now().Unix()
}

//获取当前时间
func GetNowTime() time.Time {
	return time.Now()
}

//获取当前时间毫秒
func GetNowTimeMs() int64 {
	return time.Now().UnixNano()/(1000000)
}

//int64时间转为time
func Int64ToTime(inputTime int64) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	toBeCharge := time.Unix(inputTime, 0).Format(layout)
	theTime, err := time.ParseInLocation(layout, toBeCharge, local)
	if err != nil {
		return time.Now(), err
	}
	return theTime, nil
}

//获取当天的0点时间
func GetZeroTime(timeNow int64) int64 {
	times, _ := Int64ToTime(timeNow)
	t, _ := time.Parse("2006-01-02", times.Format("2006-01-02"))
	return t.Unix() - 28800
}

//计算MD5值
func CalMd5(input []byte) string {
	return fmt.Sprintf("%x", md5.Sum(input)) //md5
}

//判断指定的文件路径是否存在
func ExistsPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

