package core

import (
	cfg "account-domain-http/config"
	"sync"
	"strconv"
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
		ctx.parseEtcdServer()
		ctx.parseMetricsServer()
		ctx.parseRedis()
		isInit = true
	}
}

func GetContext() *Context {
	return ctx
}

func (c *Context) initConfig(cfgFile string) {
	//TODO flag -c
	cfgRs, _ := cfg.NewConfigFromFile(cfgFile)
	c.Config = cfgRs
}

func (c *Context) parseServer() {
	err := c.Config.GetStruct("account-domain-http", c.Server)
	if err != nil {
		panic("Configure does not have any config for account-domain!")
	}
}

func (c *Context) parseEtcdServer() {
	c.Config.GetStruct("etcd", c.EtcdServer)
}

func (c *Context) parseMetricsServer() {
	c.Config.GetStruct("metrics", c.Metrics)
}

func (c *Context) parseRedis() {
	c.Config.GetStruct("redis", &c.Redis)
	for _, v := range c.Redis {
		NewRedisClient((*v).Name, (*v).Host, strconv.Itoa((*v).Port), (*v).MaxIdle, (*v).MaxActive, (*v).Password)
	}
}

