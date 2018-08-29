package logger

import (
	"testing"
	"time"
	"fmt"
)

func Test_log(t *testing.T) {
	//loggerFile := &LoggerFile{DebugFile:"log/debug.log",InfoFile:"log/info.log",WarnFile:"log/warn.log",ErrorFile:"log/err.log"}
	//fmt.Println(*loggerFile)
	//loggerFile := &LoggerFile{}
	InitLogger(nil)
	time1 := time.Now()
	for i := 0; i < 10; i++ {
		Debug("11111111111", "aaaaaaaaaaaa", "啊啊啊啊啊啊啊啊")
		Info("2222222222", "bbbbbbbbbbbb", "巴巴爸爸版本巴巴爸爸版本")
		Warn("3333333333333", "cccccccccccccc", "草草草草草错错错错错")
		Error("44444444444444444", "ddddddddddddd", "点点滴滴大多多多多多")
	}
	//for i := 0; i < 10000; i++ {
	//	seelog.Info("2222222222", "bbbbbbbbbbbb", "巴巴爸爸版本巴巴爸爸版本")
	//	seelog.Debug("11111111111", "aaaaaaaaaaaa", "啊啊啊啊啊啊啊啊")
	//	seelog.Warn("3333333333333", "cccccccccccccc", "草草草草草错错错错错")
	//	seelog.Error("44444444444444444", "ddddddddddddd", "点点滴滴大多多多多多")
	//}
	a := time.Now().Sub(time1)
	fmt.Println(a) //601.4005ms 662.4382ms
}
