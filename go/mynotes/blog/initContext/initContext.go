package initContext

import (
	"mynotes/blog/config"
	"mynotes/blog/logger"
	"flag"
	"mynotes/blog/database"
	"mynotes/blog/router"
	"fmt"
)

var (
	cfgPath = flag.String("c", "config/cfg.toml", "config file")
	db = flag.String("db", "mongo", "database")
)

//初始化
func init() {
	flag.Parse()
	fmt.Println("aaa")
	err := config.InitConfigFile(*cfgPath)
	if err != nil {
		panic("InitContext error :"+err.Error())
		return
	}
	logger.InitLogger(&config.Config.LogPath)
	err = database.NewMongoDB(config.Config.DataBases[*db])
	if err != nil {
		panic("InitContext error :"+err.Error())
		return
	}
}

func Run() {
	router.RegisterRouter()
}

func CloseConn() {
	err := router.HttpServer.Close()
	if err != nil {
		logger.Error(err)
	}
	database.CloseMongo()
}



