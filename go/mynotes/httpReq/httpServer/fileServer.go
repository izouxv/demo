package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

var staticfs = http.FileServer(http.Dir("/"))

func main() {
	http.Handle("/py/", http.StripPrefix("/py/", http.FileServer(http.Dir("E:/py"))))
	//httpReq.HandleFunc("/static/", static)
	//httpReq.HandleFunc("/js/", js)
	http.HandleFunc("/", route)
	http.ListenAndServe(":1789", nil)
}

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:",r.URL)
	fmt.Fprintf(w, "welcome")
	r.Body.Close()
}

//这里可以自行定义安全策略
func static(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("访问静态文件:%s\n", r.URL.Path)
	old := r.URL.Path
	r.URL.Path = strings.Replace(old, "/static", "/client", 1)
	staticfs.ServeHTTP(w, r)
}

//设置单文件访问,不能访问目录
func js(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("不能访问目录:%s\n", r.URL.Path)
	old := r.URL.Path
	name := path.Clean("E:/py" + strings.Replace(old, "/js", "/client", 1))
	info, err := os.Lstat(name)
	if err == nil {
		if !info.IsDir() {
			http.ServeFile(w, r, name)
		} else {
			http.NotFound(w, r)
		}
	} else {
		http.NotFound(w, r)
	}
}
