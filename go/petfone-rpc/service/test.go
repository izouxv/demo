package service

import (
	"petfone-rpc/core"
	"petfone-rpc/util"
	"flag"
)

func Init()  {
	cfgFile := flag.String("config", util.Conf, "config path")
	core.ContextInit(*cfgFile)
}
