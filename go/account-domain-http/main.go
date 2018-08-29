package main

import (
	. "account-domain-http/controller/adv"
	. "account-domain-http/controller/feedback"
	. "account-domain-http/controller/setting"
	. "account-domain-http/controller/user"
	. "account-domain-http/controller/asset"
	. "account-domain-http/controller"
	"account-domain-http/core"
	. "account-domain-http/filter"
	"account-domain-http/rpc"
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"strconv"
	"os"
	"os/signal"
	"time"
)

var (
	cfgFile *string = flag.String("config", "config/conf.yaml", "config path")
	seelog  *string = flag.String("seelog", "config/seelog.xml", "log config path")
)

type Intercepor map[string]http.Handler

func (hs Intercepor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	if handler := hs[r.Host]; handler != nil {
		//TODO filter
		fmt.Println("r.RequestURI :", r.RequestURI)
	} else {
		// Handle host names for wich no handler is registered
		http.Error(w, "Forbidden", 403) // Or Redirect?
	}
}

func initLoger() {
	logger, err := log.LoggerFromConfigAsFile(*seelog)
	if err != nil {
		log.Error("err parsing config log file:", err)
		return
	}
	log.ReplaceLogger(logger)
}

func rpcInit() {
	rpc.NewAdvRpcClient()
	rpc.NewRadacatVersionRpcClient()
	rpc.NewFeedbackRpcClient()
	rpc.NewMUserRpcClient()
	rpc.NewAuthRpcClient()
	rpc.NewMUserPetRpcClient()
	rpc.NewAssetRpcClient()
	rpc.NewDevTypeRpcClient()
	rpc.NewTwinsRpcClient()
	rpc.NewTeetUserRpcClient()
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,)
	flag.Parse()
	initLoger()
	log.Debug("start account-domain-http server")
	core.ContextInit(*cfgFile)
	rpcInit()
	router := httprouter.New()

/*----------------------基于租户接口-----------------------------*/
	router.POST("/v1.1/tenant/:tid/radacat/versions",                    Auth(AddNewVersion))           //增加版本
	router.PUT("/v1.1/tenant/:tid/radacat/versions/:id",                 Auth(UpdateVersion))           //修改版本
	router.GET("/v1.1/tenant/:tid/radacat/versions",                     Auth(GetAllVersions))          //分页获取版本
	router.DELETE("/v1.1/tenant/:tid/radacat/versions/:id",              Auth(DeleteVersion))           //删除版本
	router.GET("/v1.1/tenant/:tid/radacat/versions/:id",                 Auth(GetVersion))                     //获取单个版本

	router.GET( "/v1.1/tenant/:tid/accounts",                            Auth(GetPageAccounts))                //分页获取用户
	router.GET( "/v1.1/tenant/:tid/accounts/:uid",                       Auth(GetAccountByUid))                //按uid查询用户

	//router.GET( "/v1.1/tenant/:tid/accounts",                               Auth(GetAccountByUserName))
	router.POST("/v1.1/tenant/:tid/advertisements",                      Auth(NewAdvertisement))                //增加广告
	router.PUT("/v1.1/tenant/:tid/advertisements/:id",                   Auth(UpdateAdvertisement))             //修改广告
	router.GET("/v1.1/tenant/:tid/advertisements",                       Auth(GetAdvertisements))               //获取广告列表
	router.DELETE("/v1.1/tenant/:tid/advertisements/:id",                Auth(DeleteAdvertisement))             //删除广告
	router.GET("/v1.1/tenant/:tid/advertisements/:id",                   Auth(GetAdvertisement))

	router.POST("/v1.1/tenant/:tid/feedbacks",                           (AddFeedbackBaseTenant))                  //增加工单
	router.GET( "/v1.1/tenant/:tid/feedbacks",                           Auth(GetFeedbacksBaseTenant))             //获取工单列表
	router.GET( "/v1.1/tenant/:tid/feedbacks/:id",                       Auth(GetFeedbackBaseTenant))              //获取工单详情
	router.DELETE( "/v1.1/tenant/:tid/feedbacks",                        Auth(DelFeedbackBaseTenant))              //批量删除工单


	router.GET( "/v1.1/tenant/:tid/assets",                              Auth(GetAssetsForKeywordBaseTenant))      //模糊查询设备

	router.GET( "/v1.1/tenant/:tid/devices/type",                        Auth(GetDeviceTypes))                     //获取设备类型
	router.POST( "/v1.1/tenant/:tid/devices/type",                       Auth(AddDeviceType))                      //增加设备类型

	router.DELETE( "/v1.0/account/:username",                            DeleteTestUser)

	router.GET( "/v1.1/tenant/:tid/devices/info/:aid",                   Auth(GetAssetsBaseTenant))

	router.POST("/v1.1/tenant/:tid/users",                               Auth(AddTestUser))                         //增加测试账号
	router.GET( "/v1.1/tenant/:tid/users",                               Auth(GetTestUsers))                        //获取测试账号
	router.GET( "/v1.1/tenant/:tid/users/:id",                           Auth(GetTestUser))                         //获取单个测试账号
	router.DELETE( "/v1.1/tenant/:tid/users/:id",                        Auth(DelTestUser))                         //删除测试账号
	router.PUT( "/v1.1/tenant/:tid/users/:id",                           Auth(UpdateTestUser))                      //修改测试账号

	//拦截
	interceptor := make(Intercepor)
	address := net.JoinHostPort(core.GetContext().Server.Host, strconv.Itoa(core.GetContext().Server.Port))
	htp := &http.Server{Addr: address, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second, Handler: router, //TLSConfig		: tls.Config{},
	}
	interceptor[address] = router
	//监听指定信号 第一个参数表示接收信号的管道
	//第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号
	go func() {
		for s := range signalChan {
			log.Info("close conn...")
			closeRpc()
			htp.Close()
			log.Info("Go signal:", s)
		}
	}()
	log.Info(htp.ListenAndServe())

}

func closeRpc () {
	rpc.AdvRpcClientClose()
	rpc.FeedbackRpcClientClose()
	rpc.AssetRpcClientClose()
	rpc.AuthRpcClientClose()
	rpc.RadacatVersionRpcClientClose()
	rpc.UserPetClientClose()
	rpc.MuserRpcClientClose()
	rpc.DevTypeRpcClientClose()
	rpc.TwinsRpcClientClose()
}

