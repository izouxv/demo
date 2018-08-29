package util

import (
	"strconv"
	"time"
	"sync"
	"sort"
	"strings"
	"fmt"
	log "github.com/cihub/seelog"
	"bytes"
)

func StrAdd(a ...string) string {
	var bu bytes.Buffer
	for _, v := range a {
		bu.WriteString(v)
	}
	return bu.String()
}

//int32 to string
func Int32ToStr(num int32) string {
	return strconv.Itoa(int(num))
}

//string to int32
func StrToInt32(str string) (int32, error) {
	num ,err := strconv.Atoi(str)
	if err != nil {
		return  0,err
	}
	return int32(num),nil
}

//string to int64
func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

//int64 to int32
func Int64ToInt32(num64 int64) int32 {
	num,_ := strconv.Atoi(strconv.FormatInt(num64,10))
	return int32(num)
}

//int64 to string
func Int64ToStr(num64 int64) string {
	return strconv.FormatInt(num64, 10)
}

//int64 to string
func IntToStr(num int) string {
	return strconv.Itoa(num)
}

//float to Str
func FloatToStr(float float64,size int) string {
	return strconv.FormatFloat(float,'f',-1,size)
}

//float to Str
func FloatToStrSize(float float64, size int) string {
	return fmt.Sprintf("%."+IntToStr(size)+"f",float)
}

//时间戳格式化为字符串
func TimeToStr() string {
	return time.Now().Format("2006-01-02-15-04-05")
}

//时间戳格式化为字符串
func NowTimeToStr() string {
	return time.Now().Format("2006年01月02 15:04:05")
}
func TimeToStrCha2(now time.Time) string {
	return now.Format("2006-01-02 15:04:05")
}

//int64时间转为time
func Int64ToTime(inputTime int64) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Local")
	toBeCharge := time.Unix(inputTime, 0).Format(layout)
	theTime, err := time.ParseInLocation(layout, toBeCharge, local)
	if err != nil {
		log.Error("Int64ToTime err:",err)
		return time.Now(), err
	}
	return theTime, nil
}

//int64时间转为字符串
func Int64ToTimeStr(inputTime int64) string {
	return time.Unix(inputTime,0).Format("2006-01-02 15:04:05")
}

//获取当前时间,time
func GetNowTime() time.Time {
	return time.Now()
}

//获取当前时间,time
func GetNowTimeSecond() int64 {
	return time.Now().Unix()
}

//获取当前时间,time
func GetNowTimeMs() int64 {
	return time.Now().UnixNano()/(1000000)
}

//获取当天的0点时间
func GetZeroTime(timeNow int64) int64 {
	times, _ := Int64ToTime(timeNow)
	t, _ := time.Parse("2006-01-02", times.Format("2006-01-02"))
	return t.Unix() - 28800
}

func CountMinute(startTime, endTime time.Time) int64 {
	sub := int64(endTime.Sub(startTime).Seconds())
	duration := sub/60
	//if sub%60 > 0 {
	//	duration++
	//}
	return duration
}

/**--------------------------------------------*/

//自定义set
type Set struct {
	maps map[int32]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		maps: map[int32]bool{},
	}
}

func (s *Set) Add(item int32) {
	s.Lock()
	defer s.Unlock()
	s.maps[item] = true
}

func (s *Set) Adds(items []int32) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items{
		s.maps[item] = true
	}
}

func (s *Set) Remove(item int32) {
	s.Lock()
	defer s.Unlock()
	delete(s.maps, item)
}

func (s *Set) Hash(item int32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.maps[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.maps = map[int32]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Set) List() []int32 {
	s.RLock()
	defer s.RUnlock()
	list := []int32{}
	for item := range s.maps {
		list = append(list, item)
	}
	return list
}

func (s *Set) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.maps {
		list = append(list, int(item))
	}
	sort.Ints(list)
	return list
}

//替换字符串所有的双星号
func StrReplaceStar(new, old string, n int) string {
	return strings.Replace(new,"**",old,n)
}

//数字取余
func Int32ToRemStr(number int32) string {
	return Int32ToStr(number%10)
}


