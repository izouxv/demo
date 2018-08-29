package main

import (
	"flag"
	log "github.com/cihub/seelog"
	"cotx-http/core"
	"cotx-http/rpcClient"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"cotx-http/redis"
	"cotx-http/result"
	"net"
	"strconv"
	"cotx-http/utils"
	"cotx-http/pb"
	"context"
	"cotx-http/controller"
	"fmt"
	"cotx-http/permission"
	"cotx-http/po"
)


var (
	cfgFile *string = flag.String("config", "config/conf.yaml", "config path")
	seelog *string = flag.String("seelog","config/seelog.xml","log config path")
	NotificationIp string = "notification.radacat.com"
)

func main()  {
	log.Info("start --cotx-http-- ")
	flag.Parse()
	initLoger()
	initWebsocketChan()
	core.ContextInit(*cfgFile)
	/*初始化rpc*/
	rpcInit()
	/*初始话redis*/
	permission.RedisInit()
	redis.InitRedisConn()
	/*启动推送服务*/
	go websocket()
	router := httprouter.New()
	//判断用户是否存在
	router.POST(urlJudgeUsername    ,Auth(controller.JudgeUsername))
	//注册
	router.POST(urlRegister			, Auth(controller.Register))
	//登入
	router.POST(urlSession			, Auth(controller.Login))
	//登出
	router.DELETE(urlSession		, Auth(controller.Logout))
	//获取用户信息
	router.POST(urlGetUserInfo		, Auth(controller.GetUserInfo))
	//跟新用户信息
	router.PUT(urlUpdateUser		    ,Auth(controller.UpdateUserInfo))
	//修改密码
	router.PUT(urlUpdatePwd			    ,Auth(controller.UpdatePassword))
	//发送验证码(注册时、找回密码时)
	router.POST(urlSendCode		     	,Auth(controller.SendCode))
	//手机重置密码（未登录状态）
	router.POST(urlResetMobilePwd	    ,Auth(controller.ResetPassword))
	router.POST(urlResetMailPwd		    ,Auth(controller.SendMail))
	router.POST(urlFeedback		     	,Auth(controller.Feedback))
	router.POST(urlAddgateway           ,Auth(controller.RegistrationGateway))
	router.POST(urlAuthoriseAccount     ,Auth(controller.BingAccountWithGw))
	router.DELETE(urldeletAuthoriseAccount,Auth(controller.DeletAuthoriseAccount))
	router.GET(urlshowAuthAccount       ,Auth(controller.ShowAuthoriseAccount))
	router.DELETE(urldeletBindgateway   ,Auth(controller.UnwoundGateway))
	router.POST(urlvalidationgateway    ,Auth(controller.ValidationGateway))
	router.POST(urlvalidatioGwAccounr   ,Auth(controller.ValidationGatewayAccount))
	router.GET(urlgwblescans            ,Auth(controller.GetBleScans))
	router.GET(urlShowgateways          ,Auth(controller.ShowAllUserGws))
	router.GET(urlshowAllNodes          ,Auth(controller.ShowAllNodesPos))
	router.GET(urlNodesByID             ,Auth(controller.ShowNodesByGateway))
	router.GET(urlAllgws                ,Auth(controller.ShowAllGws))
	router.GET(urlgwusbwifistat         ,Auth(controller.GetGatewayUsbWifiStat))
	router.GET(urlgwusbgcardstat        ,Auth(controller.GetGatewayUsbGCardStat))
	router.GET(urlgwwificonnecting      ,Auth(controller.GetGatewayWifi))
	router.GET(urlgwusbwifiscan         ,Auth(controller.GetUsbNumWifiScan))
	router.GET(urlusbHotspotuser        ,Auth(controller.GetUsbHotSpotUser))
	router.GET(urlgwwifihotspotuser     ,Auth(controller.GetWifiHotSpotUser))
	router.GET(urlusbwifiwlan           ,Auth(controller.GetUsbWifiWlan))
    router.GET(urlGetGatewayNetState    ,Auth(controller.GetGatewayNetState))
    router.GET(urlgwswitch              ,Auth(controller.GetGatewaySwitch))
    router.GET(urlgwlora                ,Auth(controller.GetGatewayLora))
    router.GET(urlgatewayState          ,Auth(controller.GetGatewayState))
    router.GET(urlgatewayfile           ,Auth(controller.GetGatewayFileState))
    router.GET(urlgatewayvideos         ,Auth(controller.GetGatewayVideos))
    router.GET(urlgatewayphotos         ,Auth(controller.GetGatewayPhotos))
    router.GET(urlgatewayusb            ,Auth(controller.GetGatewayUSBStat))
    router.GET(urlgwwifiscan            ,Auth(controller.GetGatewayWifiScans))
    router.GET(urlgwwifiaddress         ,Auth(controller.GetGatewayWifiAddress))
    router.GET(urlgwwifiDNS             ,Auth(controller.GetGatewayWifiDNS))
    router.GET(urlgwusbwifiaddr         ,Auth(controller.GetGatewayUsbWifiAddress))
    router.GET(urlgwusbwifidns          ,Auth(controller.GetGatewayUsbWifiDNS))
    router.GET(urlgwcableaddress        ,Auth(controller.GetGatewayCableAddress))
    router.GET(urlgwcabledns            ,Auth(controller.GetGatewayCableDNS))
    router.GET(urlgwmessgae             ,Auth(controller.GetGatewayMessage))
    router.POST(urlgwset                ,Auth(controller.SendInstruction))
    router.GET(urlgwhotspot             ,Auth(controller.GetGatewayHotSpot))
    router.GET(urlgwusbhotspot          ,Auth(controller.GetGatewayUsbHotSpot))
    router.GET(urlgwpowermodel          ,Auth(controller.GetPowerModelSet))
    router.GET(urlgwMedia               ,Auth(controller.GetGatewayMedia))
    router.GET(urlgwAppEui              ,Auth(controller.GetAppEui))
    router.GET(urlnodedate              ,Auth(controller.GetNodeDateUpByNid))
    router.GET(urlgetdeveuis            ,Auth(controller.GetDeviceEuis))
    router.GET(urlgwseveriot            ,Auth(controller.GetGatewayServerIots))
    router.POST(urlregistnodes          ,Auth(controller.RegistDevice))
    router.POST(urlregistnode           ,Auth(controller.RegistNode))
    router.GET(urlgetInstructionState   ,Auth(controller.GetInstructionState))
    router.GET(urlgetNodeTypeList       ,Auth(controller.GetNodeTypeListByGatewayId))
	router.GET(urlgetNodeStatList       ,Auth(controller.GetNodeStatsByGatewayIdAndConnectState))
    router.GET(urlgetNodeStatebyType    ,Auth(controller.GetNodeStateByGatewayIdAndType))
    router.GET(urlgetGatewayTrafficStat ,Auth(controller.GetGatewayTrafficStatByGatewayId))
    router.GET(urlgetWarnStat           ,Auth(controller.GetWarnStat))
    router.GET(urlgetNodeStateById      ,Auth(controller.GetNodeStateByGatewayId))
    router.POST(urlupdatenodename       ,Auth(controller.UpdateDeviceName))
    router.DELETE(urldeletenode         ,Auth(controller.DeleteDevice))
    router.POST(urlsetiotserver         ,Auth(controller.SetIOTServer))
    router.GET(urlgetgatewayfile        ,Auth(controller.GetGatewayFile))
    router.POST(urlSetgatewayPhoto      ,Auth(controller.SetGatewayPhoto))
    router.POST(urlSetgatewayVideo      ,Auth(controller.SetGatewayVideo))
    router.PUT(urllorawandevice          ,Auth(controller.AddDevice))
    router.POST(urllorawandevice         ,Auth(controller.UpdateDevice))
    router.GET(urlGetLorawanDevice       ,Auth(controller.GetDevice))
    router.POST(urllorawandevices        ,Auth(controller.GetDeviceList))
    router.PUT(urlPutLorawanDevices      ,Auth(controller.AddDevices))
    router.DELETE(urlDeleteLorawanDevice ,Auth(controller.DeleteDeviceByDeveui))
    router.GET(urlGetGatewayVersion      ,Auth(controller.GetNewGatewayVersion))
    router.POST(urlUpdateGatewayVersion  ,Auth(controller.UpdateGatewayVersion ))
    router.GET(urlGetGatewayVersionState ,Auth(controller.GetGatewayVersionState))
    router.GET(urlGetZigbee              ,Auth(controller.GetZigbee))
    router.POST(urlupdateZigbee          ,Auth(controller.UpdateZigbee))
    router.GET(urlGetVpn                 ,Auth(controller.GetVpn))
    router.POST(urlUpdateVpn             ,Auth(controller.UpdateVpn))
    router.POST(urlpinggateway           ,Auth(controller.PingGateway))
    router.PUT(urlUploadFile             ,Auth(controller.AddFile))
    router.PUT(urlAddFeedBack            ,Auth(controller.AddFeedback))
    router.GET(urlGetSoftWareList        ,Auth(controller.GetSoftWareByversionType)) //基于类型分页获取相应类型的应用
    router.POST(urlInstallSofet          ,Auth(controller.InstallSoftware)) //基于应用id安装应用

	interceptor := make(Intercepor)
	address := net.JoinHostPort(core.GetContext().Server.Host, strconv.Itoa(core.GetContext().Server.Port))
	log.Info(address)
	interceptor[address] = router

	log.Info("http.ListenAndServe",http.ListenAndServe(address, router))
}

//HTTP接口地址
const (
	//项目名，版本号
	serverName        = "/cotx"
	version_1     	  = "/v1.0"
	urlSendCode   	  = serverName + version_1 + "/code/:action"
	urlRegister   	  = serverName + version_1 + "/register"
	urlResetMobilePwd = serverName + version_1 + "/reset/mobile"

	urlSession    	  = serverName + version_1 +"/session"
	urlUpdateUser 	  = serverName + version_1 + "/update/user/:uid"
    urlGetUserInfo 	  = serverName + version_1 + "/info"
	urlUpdatePwd  	  = serverName + version_1 + "/update/pwd/:uid"

	urlFeedback   	  = serverName + version_1 + "/feedback"
	urlResetMailPwd   = serverName + version_1 + "/mailbox/:action"

	urlJudgeUsername  = serverName + version_1 + "/user/judge"
	                           /*todo 网关与帐号*/
	 urlAddgateway            = serverName   + version_1   +    "/user/gateway"
	 urlAuthoriseAccount      = serverName   + version_1   +    "/user/account"
	 urldeletAuthoriseAccount = serverName   + version_1   +    "/user/account"
	 urlshowAuthAccount       = serverName   + version_1   +    "/user/accounts/:gwid"
	 urldeletBindgateway      = serverName   + version_1   +    "/user/gateway"
	 urlvalidationgateway     = serverName   + version_1   +    "/user/gateway/compare"
     urlvalidatioGwAccounr    = serverName   + version_1   +    "/user/account/compare"
	 urlAllgws                = serverName   + version_1   +    "/user/gateways"
	 urlShowgateways          = serverName   + version_1   +    "/gateway/position"

	                               /* todo 网关状态属性*/
     urlGetGatewayNetState    = serverName   + version_1   +    "/gateways/:gwid/net"
     urlgatewayState          = serverName   + version_1   +    "/gateways/:gwid/state"
	 urlgwmessgae             = serverName   + version_1   +    "/gateways/:gwid"

	                              /*todo 网关usb状态*/
	urlgatewayusb            = serverName    + version_1   +    "/gateways/:gwid/usb"
	urlgwusbwifiaddr          = serverName   + version_1   +    "/gateways/:gwid/usbs/:usb/wifi"
	urlgwusbwifidns           = serverName   + version_1   +    "/gateways/:gwid/usbs/:usb/wifi/dns"
	urlgwusbhotspot           = serverName   + version_1   +    "/gateways/:gwid/usbs/:usb/hotspot"
	urlgwusbwifiscan          = serverName   + version_1   +    "/gateways/:gwid/usbs/:usb/wifiscan"
	urlgwusbwifistat          = serverName   + version_1   +    "/gateways/:gwid/usb/wifistat"
	urlgwusbgcardstat         = serverName   + version_1   +    "/gateways/:gwid/usb/gcardstat"
	urlusbHotspotuser         = serverName   + version_1   +    "/gateways/:gwid/usbs/:usb/hotspotusers"
	urlusbwifiwlan            = serverName   + version_1   +    "/gateways/:gwid/usbs/:usb/wifiwlan"
	                            /*todo 网关多媒体*/
	urlgatewayfile           = serverName   + version_1   +    "/gateways/:gwid/file"
	urlgatewayvideos         = serverName   + version_1   +    "/gateways/:gwid/file/videos"
	urlgatewayphotos         = serverName   + version_1   +    "/gateways/:gwid/file/photos"
	urlgwMedia               = serverName   + version_1   +    "/gateways/:gwid/media"
	                            /*todo 网关网络接口*/
     urlgwwifiscan            = serverName   + version_1   +    "/gateways/:gwid/wifiscan"
     urlgwwifiaddress         = serverName   + version_1   +    "/gateways/:gwid/wifi"
     urlgwwifiDNS             = serverName   + version_1   +    "/gateways/:gwid/wifi/dns"
     urlgwhotspot             = serverName   + version_1   +    "/gateways/:gwid/hotspot"
     urlgwcableaddress        = serverName   + version_1   +    "/gateways/:gwid/Wired_network"
     urlgwcabledns            = serverName   + version_1   +    "/gateways/:gwid/Wired_network/dns"
     urlgwlora                = serverName   + version_1   +    "/gateways/:gwid/lorawan"
     urlgwpowermodel          = serverName   + version_1   +    "/gateways/:gwid/powermodel"
     urlgwseveriot            = serverName   + version_1   +    "/gateways/:gwid/iots"
     urlgwAppEui              = serverName   + version_1   +    "/gateways/:gwid/appeui"
     urlgwwificonnecting      = serverName   + version_1   +    "/gateways/:gwid/wifi/connection"
     urlgwswitch              = serverName   + version_1   +    "/gateways/:gwid/switch"
     urlgwwifihotspotuser     = serverName   + version_1   +    "/gateways/:gwid/hotspotusers"
     urlgwblescans            = serverName   + version_1   +    "/gateways/:gwid/blescans"
                                /*todo 网关统计接口*/
    urlgetNodeStateById       = serverName   + version_1   +     "/gateways/:gwid/nodes/stat"
	urlgetNodeTypeList        = serverName   + version_1   +     "/gateways/:gwid/nodes/typelist"
	urlgetNodeStatebyType     = serverName   + version_1   +     "/gateways/:gwid/nodes/stat/:type"
	urlgetNodeStatList        = serverName   + version_1   +     "/gateways/:gwid/nodes/state/:type"
	urlgetGatewayTrafficStat  = serverName   + version_1   +     "/gateways/:gwid/traffics/stat"
	urlgetWarnStat            = serverName   + version_1   +     "/gateways/:gwid/nodes/warns/stat"
                                /*todo 网关 设置*/
	urlgwset                  = serverName   + version_1   +    "/gateways/:gwid/instruction"
	urlgetInstructionState    = serverName   + version_1   +    "/gateways/:gwid/instruction/:instruction"
	                            /*todo 应用安装*/
	urlGetSoftWareList        = serverName   + version_1   +    "/gateways/:gwid/softwares/:type"
	urlInstallSofet           = serverName   + version_1   +    "/gateways/:gwid/software"
//终端接口
     urlshowAllNodes          = serverName   + version_1   +     "/node/position"
     urlNodesByID             = serverName   + version_1   +     "/gateways/:gwid/nodes"
     urlnodedate              = serverName   + version_1   +     "/nodes/:nid/data"
     urlgetdeveuis            = serverName   + version_1   +     "/gateways/:gwid/nodes/deveuis"
     urlregistnodes           = serverName   + version_1   +     "/gateways/:gwid/nodes"
     urlregistnode            = serverName   + version_1   +     "/gateways/:gwid/node"
     //终端名称更新
     urlupdatenodename        = serverName   + version_1    +     "/gateways/:gwid/nodes/:nid/node"
     //删除终端
     urldeletenode            = serverName   +  version_1   +     "/gateways/:gwid/nodes/:nid/node"
     //设置云平台
     urlsetiotserver          = serverName   + version_1    +     "/gateways/:gwid/IotServer"
     //上传文件的指令
     urlgetgatewayfile        = serverName   + version_1    +     "/gateways/:gwid/upfile/:name"
     //设置手动拍照
     urlSetgatewayPhoto       = serverName    + version_1   +     "/gateways/:gwid/camera/photo"
     //设置自动拍照
     urlSetgatewayVideo       = serverName    + version_1   +     "/gateways/:gwid/camera/video"
     //网关升级的操作
     urlGetGatewayVersion      = serverName  + version_1  +        "/gateways/:gwid/version/:type"
     urlUpdateGatewayVersion   = serverName  +  version_1 +        "/gateways/:gwid/version"
     urlGetGatewayVersionState = serverName  +  version_1 +        "/gateways/:gwid/versionstate/:type"
     //网关zigbee操作
     urlGetZigbee              = serverName   + version_1  +       "/gateways/:gwid/zigbee"
     urlupdateZigbee           = serverName   + version_1  +       "/gateways/:gwid/zigbee"
     //网关zigbee操作
     urlGetVpn                 = serverName   + version_1  +        "/gateways/:gwid/vpn"
     urlUpdateVpn              = serverName   + version_1  +        "/gateways/:gwid/vpn"

     //网关ping功能
     urlpinggateway            = serverName    + version_1  +       "/gateways/:gwid/ping"

	 //lora 终端操作接口
	 urllorawandevice          = serverName   + version_1   +      "/lorawan/device"
	 urllorawandevices         = serverName   + version_1   +      "/lorawan/devices"
	 urlPutLorawanDevices      = serverName   + version_1   +      "/lorawan/devices"
	 urlDeleteLorawanDevice    = serverName   + version_1   +      "/lorawan/device/:deveui"
	 urlGetLorawanDevice       = serverName   + version_1   +      "/lorawan/device/:deveui"
	 /*app问题与反馈*/
	 urlUploadFile             = serverName   + version_1   +      "/user/file"
	 urlAddFeedBack            = serverName   + version_1   +      "/user/feedback"
)

func initLoger() {
	logger, err := log.LoggerFromConfigAsFile(*seelog)
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
}

func rpcInit() {
	rpcClient.NewSsoRpcClient()
	rpcClient.NewAccountRpcClient()
	rpcClient.NewGwUserRpcClient()
	rpcClient.NewGatewayRpcClient()
	rpcClient.NewGatewaySetRpcClient()
	rpcClient.NewRpcClinetDevice()
	rpcClient.NewGatewayUsbRpcClient()
	rpcClient.NewRpcClinetUpdateDevice()
	rpcClient.NewRpcClinetZigbee()
	rpcClient.NewRpcClinetVpn()
	rpcClient.NewWebscoketRpcClient()
	rpcClient.NewFileClient()
	rpcClient.NewFeedbackClient()
	rpcClient.NewRpcClinetSoftware()
}

func  Auth(handle httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Println("params",params)
		log.Info("Start Auth....")
		source := req.Header.Get("source")
		log.Info("source:",source)
		if source == ""{
			log.Error("Get source Failed , ")
			result.JsonReply("Source_is_incorrect_or_empty", nil, res)
			return
		}
		//对半公开接口进行判断
		url :=req.URL.Path
		urlFlag := utils.VerifyHalfOpenAddress(url, req.Method)
		fmt.Println("urlFlag = ", urlFlag)
		if urlFlag {
			req = req.WithContext(context.WithValue(req.Context(), "contextSso", &pb.SsoRequest{Source:source}))
			handle(res, req, params)
			log.Info("Half_open_address urlFlag = ", urlFlag)
			return
		}

		token := req.Header.Get("token")
		if token == "" {
			log.Info("Get Token Failed,")
			result.JsonReply("Token_is_empty",nil, res)
			return
		}

		//调用rpc进行token校验
		reply := rpcClient.GetUserInfo(&pb.SsoRequest{Source:source, SessionName:token})
		reply.SessionName = token
		log.Info("Check Token Reply:  ",reply)
		if reply.ErrorCode != 10000 {
			log.Error("Token:", token, " Check  Failed, ", reply.ErrorCode)
			result.JsonReply("Token_is_incorrect_or_lose_efficacy",nil, res)
			return
		}
		handle(res, req.WithContext(context.WithValue(req.Context(), "userInfo", reply)), params)
	}
}

type Intercepor map[string]http.Handler

func (hs Intercepor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	log.Info("in source :",r.Header.Get("source"))
	client := redis.RedisClient("persistence").Get()
	defer client.Close()
	res, err := client.Do("hget", r.Header.Get("source"))
	if err != nil || res == nil{
		log.Info("Source_is_incorrect_or_empty")
		result.JsonReply("Source_is_incorrect_or_empty", nil, w)
		return
	}
	if handler := hs[r.Host]; handler != nil {
		//TODO filter
		handler.ServeHTTP(w, r)
	} else {
		// Handle host names for which no handler is registered
		http.Error(w, "Forbidden", 403) // Or Redirect?
	}
}

func initWebsocketChan()  {
	 controller.WebsocketChan = po.WebscoketChan{
	 	   Websocketchan:make(chan map[string]po.WebSocketMessgage),
	 	   Websocketmap:make(map[string]po.WebSocketMessgage),
	 }
}

func websocket()  {
	http.HandleFunc("/websocket",controller.WebsocketHandle)
	http.ListenAndServe(":7009",nil)
}
