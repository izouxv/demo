package core

import (
	cfg "petfone-rpc/config"
	"strconv"
	"sync"
	log "github.com/cihub/seelog"

)

var (
	contextLock sync.Mutex
	ctx         *Context
	isInit      = false
)

type Context struct {
	Config *cfg.Config
	Server *ServiceConfig
	Mysql  *DBConfig
	RpcServer  *RpcConfig
	EtcdServer *ServiceConfig
	Redis []*ServiceConfig
	Metrics    *ServiceConfig
}

func ContextInit(cfgFile string) {
	contextLock.Lock()
	defer contextLock.Unlock()
	if ctx == nil {
		ctx = &Context{
			Server: new(ServiceConfig),
			Mysql:  new(DBConfig),
			//RpcServer: new(RpcConfig),
			//EtcdServer: new(ServiceConfig),
			Redis: make([]*ServiceConfig, 0),
			//Metrics:    new(ServiceConfig),
		}
		ctx.initConfig(cfgFile)
		ctx.parseServer()
		ctx.parseMysql()
		//ctx.initRpc()
		//ctx.parseEtcdServer()
		ctx.parseRedis()
		//ctx.parseMetricsServer()
		ctx.parseConst()
		isInit = true
	}
	log.Info("ContextInit-init:",isInit)
}

func GetContext() *Context {
	return ctx
}

func (c *Context) initConfig(cfgFile string) {
	//$TODO flag -c
	cfgRs := cfg.NewConfigFromFile(cfgFile)
	c.Config = cfgRs
}

func (c *Context) parseServer() {
	if c.Config.GetStruct("petfone-rpc", c.Server) != nil {
		panic("Configure does not have any config for petfone-rpc service!")
	}
}

func (c *Context) parseEtcdServer() {
	c.Config.GetStruct("etcd", c.EtcdServer)
}

func (c *Context) parseMysql() {
	c.Config.GetStruct("mysql", c.Mysql)
	NewMysqlClient(c.Mysql.DBHost, c.Mysql.DBName, c.Mysql.DBUser, c.Mysql.DBPwd, c.Mysql.DBPort)
}

//加载rpc服务
func (c *Context) initRpc() {
	//todo 加载rpc
	c.Config.GetStruct("agent_rpc", &c.RpcServer)
	agentAddress := c.RpcServer.Address + c.RpcServer.Port
	log.Info("Rpc-agentAddress:", agentAddress)
	//todo 初始化rpc
	//rpc_client.AgentRpcInit(agentAddress)
}


func CloseRpc() {
	//rpc_client.AgentRpcClose()
	log.Info("CloseRpc")
}

func (c *Context) parseRedis() {
	c.Config.GetStruct("redis", &c.Redis)
	for _, v := range c.Redis {
		RedisInit((*v).Name, (*v).Host, strconv.Itoa((*v).Port), (*v).MaxIdle, (*v).MaxActive, (*v).Password)
		RedisPing(int32((*v).Port))
	}

}

var ConstStr *Const
var (
	Names = map[int]string{}
	Voices = map[int]string{}
)
func (c *Context) parseConst() {
	c.Config.GetStruct("const", &ConstStr)
	log.Info("init-Const:",ConstStr)
	Names[1] = ConstStr.Name1
	Names[2] = ConstStr.Name2
	Names[3] = ConstStr.Name3
	Voices[1] = ConstStr.Voice1
	Voices[2] = ConstStr.Voice2
	Voices[3] = ConstStr.Voice3
}

//func (c *Context) parseMetricsServer() {
//	c.Config.GetStruct("metrics", &c.Metrics)
//	log.Info("metrics:",c.Metrics)
//}
