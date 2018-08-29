package config

import (
	"time"
	"sync"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"fmt"
)

var (
	Config *configToml
	configLock = new(sync.RWMutex)
)

//读取配置文件返回配置参数
func InitConfigFile(filePath string) error {
	if Config != nil {
		return nil
	}
	configLock.Lock()
	defer configLock.Unlock()
	if Config != nil {
		return errors.New("InitConfigFile Config not nil")
	}
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		return errors.New("InitConfigFile err:"+err.Error())
	}
	fmt.Println("qqq:",Config)
	return nil
}

//配置参数模块
type configToml struct {
	Title		string				`toml:"title"`
	LogPath		LogPath				`toml:"log_path"`
	AppServer	appServer			`toml:"blog_http"`
	DataBases	map[string]Database	`toml:"databases"`
	RpcServers 	map[string]RpcServer`toml:"rpc_servers"`
	Clients 	Clients				`toml:"clients"`
}

type appServer struct {
	AppName string	`toml:"name"`
	Addr	string
	Use		string
	Short	string
	Version	string
	Time	time.Time
}

type LogPath struct {
	Debug   string	// debug日志输出文件
	Info    string 	// info日志输出文件
	Warn	string 	// warn日志输出文件
	Error	string 	// error日志输出文件
}

type Database struct {
	Addr	string
	Name	string
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type RpcServer struct {
	Addr  	string
}

type Clients struct {
	Data  [][]interface{}
	Hosts []string
}

