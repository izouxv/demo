package core

import (
	"sync"
	"strconv"
	log "github.com/cihub/seelog"
	"petfone-http/config"
	"petfone-http/rpc"
)

//启动的配置服务
type Context struct {
	Config     *config.Config
	Server     *ServiceConfig
	Redis      []*ServiceConfig
	Mysql      *DBConfig
	RpcServer  *RpcConfig
	EtcdServer *ServiceConfig
	Metrics    *ServiceConfig
}

type RpcConfig struct {
	Address string
	Port    string
}

//服务配置属性
type ServiceConfig struct {
	Name      string
	Host      string
	Port      int
	MaxIdle   int
	MaxActive int
	Enable    bool
	Interval  int
	Password  string
}

//数据库配置属性
type DBConfig struct {
	DBHost      string
	DBPort      int
	DBUsername  string
	DBPassword  string
	DBName      string
	DBMaxIdle   int
	DBMaxActive int
}

type Const struct {
	//todo 默认头像ID
	UserAvatar	string
	PetAvatar   string
	//todo 宠物默认训练项
	Name1       string
	Name2       string
	Name3       string
	Voice1      string
	Voice2      string
	Voice3      string
	//todo 文件服务地址
	FileServer  string
	//todo 图片服务地址
	ImageServer	string
	//todo 通知服务地址
	NoticeServer  string
	//todo mysql异常字符串
	NotFound    string
}

var (
	contextLock sync.Mutex
	ctx         *Context
)

func GetContext() *Context {
	return ctx
}

//加载petfone服务启动的前置条件
func ContextInit(cfgFile string) {
	//加载数据库服务
	contextLock.Lock()
	defer contextLock.Unlock()
	if ctx == nil {
		ctx = &Context{
			Server:    new(ServiceConfig),
			Redis:     make([]*ServiceConfig, 0),
			RpcServer: new(RpcConfig),
			//Mysql:  new(DBConfig),
			//EtcdServer	:	new(ServiceConfig),
			//Metrics   	:	new(ServiceConfig),
		}
		ctx.parseConfig(cfgFile)
		ctx.initServer()
		ctx.initRedis()
		ctx.initRpc()
		//ctx.initMysql()
		//ctx.parseEtcdServer()
		//ctx.parseMetricsServer()
		ctx.parseConst()
	}
	log.Info("Context-init")
}

//将从配置文件读取的信息应用
func (c *Context) parseConfig(cfgFile string) {
	//todo flag -c
	cfgRs := config.NewConfigFromFile(cfgFile)
	c.Config = cfgRs
}

//加载Server配置
func (c *Context) initServer() {
	err := c.Config.GetStruct("petfone-http", c.Server)
	if err != nil {
		panic("petfone initServer err:"+err.Error())
	}
	log.Info("server config finish ...")
}

//加载初始化redis配置
func (c *Context) initRedis() {
	c.Config.GetStruct("redis", &c.Redis)
	for _, value := range c.Redis {
		redis := *value
		RedisConfig(redis.Name, redis.Host, strconv.Itoa(redis.Port), redis.MaxIdle, redis.MaxActive, redis.Password)
	}
}

//加载rpc服务
func (c *Context) initRpc() {
	//todo 加载项圈rpc
	c.Config.GetStruct("rpc", &c.RpcServer)
	address := c.RpcServer.Address + c.RpcServer.Port
	log.Info("Rpc-petfoneAddress:", address)
	//todo 加载管理后台rpc
	c.Config.GetStruct("admin_rpc", &c.RpcServer)
	adminAddress := c.RpcServer.Address + c.RpcServer.Port
	log.Info("Rpc-adminAddress:", adminAddress)
	//todo 加载设备数据rpc
	c.Config.GetStruct("agent_rpc", &c.RpcServer)
	agentAddress := c.RpcServer.Address + c.RpcServer.Port
	log.Info("Rpc-agentAddress:", agentAddress)
	//todo 初始化rpc
	rpc.SsoRpcInit(address)
	rpc.AccountRpcInit(address)
	rpc.FaqCommonRpcInit(address)
	rpc.PetfoneRpcInit(address)
	rpc.DevicesRpcInit(address)
	rpc.PetInfoRpcInit(address)
	rpc.ShareRpcInit(address)
	rpc.ExerciseRpcInit(address)
	rpc.NoticeRpcInit(address)
	rpc.FilesRpcInit(address)
	go func() {
		rpc.FeedBackRpcInit(adminAddress)
		rpc.AdverRpcInit(adminAddress)
		rpc.VersionRpcInit(adminAddress)
		rpc.AgentRpcInit(agentAddress)
	}()
}

var ConstStr *Const
func (c *Context) parseConst() {
	c.Config.GetStruct("const", &ConstStr)
}

func CloseRpc() {
	rpc.AccountRpcClose()
	rpc.FaqRpcClose()
	rpc.PetfoneRpcClose()
	rpc.DeviceRpcClose()
	rpc.PetInfoRpcClose()
	rpc.ShareRpcClose()
	rpc.SsoRpcClose()
	rpc.ExerciseRpcClose()
	rpc.NoticeRpcClose()
	rpc.FilesRpcClose()
	rpc.FeedBackRpcClose()
	rpc.AdverRpcClose()
	rpc.VersionRpcClose()
	rpc.AgentRpcClose()
	log.Info("CloseRpc")
}

//加载etcd配置
func (c *Context) parseEtcdServer() {
	c.Config.GetStruct("etcd", c.EtcdServer)
}

//加载metrics配置
func (c *Context) parseMetricsServer() {
	c.Config.GetStruct("metrics", c.Metrics)
}
