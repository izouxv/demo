package core

import (
	cfg "cotx-http/config"
	"cotx-http/redis"
	"strconv"
	"sync"
)
const (
	CfgFile = "config/conf.yaml"
)

var (
	contextLock sync.Mutex
	ctx         *Context
	isInit      = false
)

type Context struct {
	Config     *cfg.Config
	Server     *ServiceConfig
	Redis      []*ServiceConfig
	EtcdServer *ServiceConfig
	Metrics    *ServiceConfig
}

func ContextInit(cfgFile string) {
	contextLock.Lock()
	defer contextLock.Unlock()
	if ctx == nil {
		ctx = &Context{
			Server:     new(ServiceConfig),
			Redis:      make([]*ServiceConfig, 0),
			EtcdServer: new(ServiceConfig),
			Metrics:    new(ServiceConfig),
		}
		ctx.initConfig(cfgFile)
		ctx.parseServer()
		ctx.parseRedis()
		ctx.parseEtcdServer()
		ctx.parseMetricsServer()
		isInit = true
	}
}

func GetContext() *Context {
	return ctx
}

func (c *Context) initConfig(cfgFile string) {
	//$TODO flag -c
	cfgRs, _ := cfg.NewConfigFromFile(cfgFile)
	c.Config = cfgRs
}

func (c *Context) parseServer() {
    err := c.Config.GetStruct("cotx-http", c.Server)
	if err != nil {
		panic("Configure does not have any config for sso-http!")
	}
}

func (c *Context) parseRedis() {
	c.Config.GetStruct("redis", &c.Redis)
	//TODO 实例化redisClient
	for _, v := range c.Redis {
		redis.NewRedisClient((*v).Name, (*v).Host, strconv.Itoa((*v).Port), (*v).MaxIdle, (*v).MaxActive, (*v).Password)
	}
}

func (c *Context) parseEtcdServer() {
	c.Config.GetStruct("etcd", c.EtcdServer)
}

func (c *Context) parseMetricsServer() {
	c.Config.GetStruct("metrics", c.Metrics)
}
