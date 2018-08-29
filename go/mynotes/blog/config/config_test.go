package config

import (
	"fmt"
	"testing"
)

//测试读取配置文件
func Test_initConfigFile(t *testing.T) {
	err := InitConfigFile("cfg.toml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Title: ", Config.Title)
	fmt.Println("AppServer: ", Config.AppServer)
	fmt.Println("DataBases: ", Config.DataBases)
	fmt.Println("Servers: ", Config.RpcServers)
	fmt.Println("Client: ", Config.Clients)
}

