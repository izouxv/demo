package util

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/trace"
	_"net/http/pprof"
	"runtime/debug"
)

func main() {
	//开启强大的分析器
	go Runpprof()
	//以下运行具体代码
}

//运行pprof分析器
func Runpprof() {
	//运行trace
	http.HandleFunc("/start", traces)
	//停止trace
	http.HandleFunc("/stop", traceStop)
	//手动GC
	http.HandleFunc("/start_gc", startGc)
	http.HandleFunc("/stop_gc", stopGc)
	//网站开始监听
	fmt.Println("ListenAndServe:",http.ListenAndServe(":8080", nil))
}

//手动GC
func startGc(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	w.Write([]byte("startGc"))
}

//手动停止GC
func stopGc(w http.ResponseWriter, r *http.Request) {
	debug.SetGCPercent(-1)
	w.Write([]byte("stopGc"))
}

//运行trace
func traces(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("TrancStart"))
	fmt.Println("StartTrancs")
}

//停止trace
func traceStop(w http.ResponseWriter, r *http.Request) {
	trace.Stop()
	w.Write([]byte("TrancStop"))
	fmt.Println("StopTrancs")
}
