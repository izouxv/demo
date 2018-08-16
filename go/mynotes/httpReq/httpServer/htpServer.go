package main

import (
	"net/http"
	"io"
	"encoding/json"
	"expvar"
	"fmt"
	"runtime"
	"time"
)


func main() {
	//http.HandleFunc("/a",router)
	//http.HandleFunc("/debug/vars", vars)
	http.ListenAndServe(":8888",nil)
}

func router(res http.ResponseWriter, req *http.Request)  {
	io.Copy(res,req.Body)
	res.Write([]byte("1111111111"))
}

//进程运行信息
func vars(res http.ResponseWriter, req *http.Request)  {
	// GetCurrentRunningStats 返回当前运行信息
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	first := true
	report := func(key string, value interface{}) {
		if !first {
			res.Write([]byte(",\n"))
		}
		first = false
		if str, ok := value.(string); ok {
			res.Write([]byte(fmt.Sprintf("%q: %q", key, str)))
		} else {
			res.Write([]byte(fmt.Sprintf("%q: %v", key, value)))
		}
	}
	res.Write([]byte("{\n"))
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	res.Write([]byte("\n}\n"))

	//res.String(http.StatusOK, "")
}


var CuMemoryPtr *map[string]string
var BTCMemoryPtr *map[string]string


// 开始时间
var start = time.Now()

// calculateUptime 计算运行时间
func calculateUptime() interface{} {
	return time.Since(start).String()
}

// currentGoVersion 当前 Golang 版本
func currentGoVersion() interface{} {
	return runtime.Version()
}

// getNumCPUs 获取 CPU 核心数量
func getNumCPUs() interface{} {
	return runtime.NumCPU()
}

// getGoOS 当前系统类型
func getGoOS() interface{} {
	return runtime.GOOS
}

// getNumGoroutins 当前 goroutine 数量
func getNumGoroutins() interface{} {
	return runtime.NumGoroutine()
}

// getNumCgoCall CGo 调用次数
func getNumCgoCall() interface{} {
	return runtime.NumCgoCall()
}

// 业务特定的内存数据
func getCuMemoryMap() interface{} {
	if CuMemoryPtr == nil {
		return 0
	} else {
		return len(*CuMemoryPtr)
	}
}
// 业务特定的内存数据
func getBTCMemoryMap() interface{} {
	if BTCMemoryPtr == nil {
		return 0
	} else {
		return len(*BTCMemoryPtr)
	}
}

var lastPause uint32

// getLastGCPauseTime 获取上次 GC 的暂停时间
func getLastGCPauseTime() interface{} {
	var gcPause uint64
	ms := new(runtime.MemStats)

	statString := expvar.Get("memstats").String()
	if statString != "" {
		json.Unmarshal([]byte(statString), ms)

		if lastPause == 0 || lastPause != ms.NumGC {
			gcPause = ms.PauseNs[(ms.NumGC+255)%256]
			lastPause = ms.NumGC
		}
	}

	return gcPause
}

func init() {   //这些都是我自定义的变量，发布到expvar中，每次请求接口，expvar会自动去获取这些变量，并返回给我
	expvar.Publish("运行时间", expvar.Func(calculateUptime))
	expvar.Publish("version", expvar.Func(currentGoVersion))
	expvar.Publish("cores", expvar.Func(getNumCPUs))
	expvar.Publish("os", expvar.Func(getGoOS))
	expvar.Publish("cgo", expvar.Func(getNumCgoCall))
	expvar.Publish("goroutine", expvar.Func(getNumGoroutins))
	expvar.Publish("gcpause", expvar.Func(getLastGCPauseTime))
	expvar.Publish("CuMemory", expvar.Func(getCuMemoryMap))
	expvar.Publish("BTCMemory", expvar.Func(getBTCMemoryMap))
}
