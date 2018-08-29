package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"cotx-http/result"
	"fmt"
	log "github.com/cihub/seelog"
	"cotx-http/rpcClient"
	"cotx-http/utils"
	"cotx-http/pb"
	"strconv"
	"golang.org/x/net/context"
)

func GetNodeDateUpByNid(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	nodeid := param.ByName("nid")
	fmt.Println("nodeid======",nodeid)
	reqnode := &pb.ReqNodeMessageByNid{NodeId:nodeid,UserId:userinfo.Uid}
	//rpc 调用
	resnode ,err:=rpcClient.GetNodeDateUpByNid(reqnode)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	fmt.Println("++++++++++++++++++CODE=:",resnode.ErrCode)
	if resnode.ErrCode != 10000 {
		switch resnode.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resnode.Data,res)
}
func GetDeviceEuis(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	nodeid  := param.ByName("gwid")
	typeString := req.FormValue("type")
	typeInt ,_ := strconv.Atoi(typeString)
	reqnode := &pb.ReqNodeMessageByNid{GatewayId:nodeid,UserId:userinfo.Uid,Type:int32(typeInt)}
	//rpc 调用
	resnode ,err:=rpcClient.GetDeviceEuis(reqnode)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	fmt.Println("++++++++++++++++++CODE=:",resnode.ErrCode)
	if resnode.ErrCode != 10000 {
		switch resnode.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30023:
			result.JsonReply("NoFind_node_deveui",nil,res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return

		}
	}
	result.JsonReply("Successful",resnode.DevEuis,res)
}
func RegistDevice(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqnode := &pb.ReqregistDevice{GatewayId:gatewayid,UserId:userinfo.Uid}
	pbType := utils.GetHttpData(req, "application/json;charset=UTF-8", reqnode)

	log.Info("判断后台是否获取deveuis参数：",reqnode)
	if pbType != 10000 {
		if pbType == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		if pbType == 30022 {
			result.JsonReply("Instruction_Error",nil,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	//rpc 调用
	resnode ,err:=rpcClient.RegistDevices(reqnode)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	fmt.Println("++++++++++++++++++CODE=:",resnode.ErrCode)
	if resnode.ErrCode != 10000 {
		switch resnode.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resnode,res)
}
func RegistNode(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqnode := &pb.ReqRegistNode{GatewayId:gatewayid,UserId:userinfo.Uid}
	pbType := utils.GetHttpData(req, "application/json;charset=UTF-8", reqnode)
	if pbType != 10000 {
		if pbType == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		if pbType == 30022 {
			result.JsonReply("Instruction_Error",nil,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	//rpc 调用
	resnode ,err:=rpcClient.RegistNode(reqnode)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	fmt.Println("++++++++++++++++++CODE=:",resnode.ErrCode)
	if resnode.ErrCode != 10000 {
		switch resnode.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resnode,res)
}

func AddDevice(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var reqDevice = new(pb.PutDeviceRequest)
	var device = new(pb.LoraDevice)
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &device)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	log.Debugf("前端传来的参数:%+v",device)
	reqDevice.Device = device
	reqDevice.UserId = userinfo.Uid
	resDevice := rpcClient.AddDevice(reqDevice)
	if resDevice.ErrCode != 10000 {
		switch resDevice.ErrCode  {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resDevice.Device,res)
}


func GetDevice(res http.ResponseWriter ,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var reqDevice = new(pb.GetDeviceRequest)
	reqDevice.UserId = userinfo.Uid
	deveui := param.ByName("deveui")
	reqDevice.Deveui = deveui
	resDevice := rpcClient.GetDevice(reqDevice)
	if resDevice.ErrCode != 10000 {
		switch resDevice.ErrCode  {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resDevice.Device,res)
}
func UpdateDevice(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var reqDevice = new(pb.PostDeviceRequest)
	var device = new(pb.LoraDevice)
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &device)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	reqDevice.Device = device
	reqDevice.UserId = userinfo.Uid
	resDevice := rpcClient.UpdateDevice(reqDevice)
	if resDevice.ErrCode != 10000 {
		switch resDevice.ErrCode  {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resDevice.Device,res)
}

func GetDeviceList(res http.ResponseWriter ,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var reqDevice = new(pb.GetDeviceListRequest)
	var deveuis = make([]string,0)
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &deveuis)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	reqDevice.Deveuis = deveuis
	reqDevice.UserId = userinfo.Uid
	resDevice := rpcClient.GetDeviceListByDeveuis(reqDevice)
	if resDevice.ErrCode != 10000 {
		switch resDevice.ErrCode  {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resDevice.Devices,res)
}

/*基于deveui 删除终端*/
func DeleteDeviceByDeveui(res http.ResponseWriter ,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var reqDevice = new(pb.DeleteDeviceRequest)
    deveui := param.ByName("deveui")
    reqDevice.Deveui = deveui
	reqDevice.UserId = userinfo.Uid
	resDevice := rpcClient.DeleteDeviceByDeveui(reqDevice)
	if resDevice.ErrCode != 10000 {
		switch resDevice.ErrCode  {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resDevice,res)
}

func AddDevices(res http.ResponseWriter ,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var reqDevice = new(pb.AddDevicesRequest)
	var devices = make([]*pb.LoraDevice,0)
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &devices)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	reqDevice.Devices = devices
	reqDevice.UserId = userinfo.Uid
	resDevice := rpcClient.AddDevices(reqDevice)
	if resDevice.ErrCode != 10000 {
		switch resDevice.ErrCode  {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resDevice.Devices,res)
}
func  UpdateDeviceName(res http.ResponseWriter,req *http.Request,param httprouter.Params) {
	userinfo, ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :", userinfo)

	gatewayid := param.ByName("gwid")
	nodeid := param.ByName("nid")

	reqgateway := &pb.ReqSetDeviceName{UserId: userinfo.Uid, GatewayId: gatewayid, NodeId: nodeid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality", nil, res)
		return
	}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &reqgateway)
	fmt.Println("errCode = ", errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	resgateway ,err:= rpcClient.GetDeviceRpcClient().UpdateDeviceName(context.Background(),reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway", nil, res)
			return
		case 30020:
			result.JsonReply("Operation_DB_Error", nil, res)
			return
		case 30017:
			result.JsonReply("Parameter_format_error", nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful", resgateway.NodeMessage, res)
}

/*删除设备*/
func  DeleteDevice(res http.ResponseWriter,req *http.Request,param httprouter.Params) {
	userinfo, ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :", userinfo)
	gatewayid := param.ByName("gwid")
	nodeid := param.ByName("nid")

	reqgateway := &pb.ReqDeleteDevice{UserId: userinfo.Uid, GatewayId: gatewayid, NodeId: nodeid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality", nil, res)
		return
	}

	resgateway,err := rpcClient.GetDeviceRpcClient().DeleteDevice(context.Background(),reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway", nil, res)
			return
		case 30020:
			result.JsonReply("Operation_DB_Error", nil, res)
			return
		case 30017:
			result.JsonReply("Parameter_format_error", nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful", resgateway, res)
}
/*基于网关的id获取当前网关下的所有终端类型*/
func GetNodeTypeListByGatewayId(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	var getNodeTypeListByGatewayIdRequest = new(pb.GetNodeTypeByGatewayIdRequest)
	getNodeTypeListByGatewayIdRequest.GatewayId = gatewayid
	getNodeTypeListByGatewayIdRequest.UserId = userinfo.Uid
	getNodeTypeListByGatewayIdResponse,err := rpcClient.GetDeviceRpcClient().GetNodeTypeByGatewayId(context.Background(),getNodeTypeListByGatewayIdRequest)
	if err != nil || getNodeTypeListByGatewayIdResponse.ErrCode != 10000 {
		switch getNodeTypeListByGatewayIdResponse.ErrCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway", nil, res)
			return
		case 30020:
			result.JsonReply("Operation_DB_Error", nil, res)
			return
		case 30017:
			result.JsonReply("Parameter_format_error", nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful", getNodeTypeListByGatewayIdResponse.NodeType, res)
}

func GetNodeStateByGatewayIdAndType(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid  := param.ByName("gwid")
	nodeType  := param.ByName("type")
	var getNodeStateByGatewayIdAndTypeRequest = new(pb.GetNodeStateByGatewayIdAndTypeRequest)
	getNodeStateByGatewayIdAndTypeRequest.UserId = userinfo.Uid
	getNodeStateByGatewayIdAndTypeRequest.GatewayId = gatewayid
	getNodeStateByGatewayIdAndTypeRequest.NodeType = nodeType
	getNodeStateByGatewayIdAndTypeResponse,err :=  rpcClient.GetDeviceRpcClient().GetNodeStateByGatewayIdAndNodeType(context.Background(),getNodeStateByGatewayIdAndTypeRequest)
	if err != nil || getNodeStateByGatewayIdAndTypeResponse.ErrCode != 10000 {
		switch getNodeStateByGatewayIdAndTypeResponse.ErrCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway", nil, res)
			return
		case 30020:
			result.JsonReply("Operation_DB_Error", nil, res)
			return
		case 30017:
			result.JsonReply("Parameter_format_error", nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful",getNodeStateByGatewayIdAndTypeResponse.NodeState, res)
}
/*基于网关id,终端类型和终端的连接状态*/
func GetNodeStatsByGatewayIdAndConnectState(res http.ResponseWriter,req *http.Request,param httprouter.Params) {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	nodeType := param.ByName("type")
	connectStateString := req.FormValue("state")
	connectState,_ := strconv.Atoi(connectStateString)
	var getNodeStatListByGatewayIdAndConnectStateRequest  = new(pb.GetNodeStatListByGatewayIdAndConnectStateRequest)
	getNodeStatListByGatewayIdAndConnectStateRequest.GatewayId = gatewayid
	getNodeStatListByGatewayIdAndConnectStateRequest.UserId = userinfo.Uid
	getNodeStatListByGatewayIdAndConnectStateRequest.ConnectState = int32(connectState)
	getNodeStatListByGatewayIdAndConnectStateRequest.NodeType = nodeType
	getNodeStatListByGatewayIdAndConnectStateResponse,err := rpcClient.GetDeviceRpcClient().GetNodeStatsByGatewayIdAndConnectState(context.Background(),getNodeStatListByGatewayIdAndConnectStateRequest)
	if err != nil || getNodeStatListByGatewayIdAndConnectStateResponse.ErrCode != 10000 {
		switch getNodeStatListByGatewayIdAndConnectStateResponse.ErrCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway", nil, res)
			return
		case 30020:
			result.JsonReply("Operation_DB_Error", nil, res)
			return
		case 30017:
			result.JsonReply("Parameter_format_error", nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful",getNodeStatListByGatewayIdAndConnectStateResponse.NodeStats, res)
}

func GetNodeStateByGatewayId(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	var getNodeStateByGatewayIdRequest = new(pb.GetNodeStateByGatewayIdRequest)
	getNodeStateByGatewayIdRequest.UserId = userinfo.Uid
	getNodeStateByGatewayIdRequest.GatewayId = gatewayid
	getNodeStateByGatewayIdResponse,err := rpcClient.GetDeviceRpcClient().GetNodeStateByGatewayId(context.Background(),getNodeStateByGatewayIdRequest)
	if err != nil || getNodeStateByGatewayIdResponse.ErrCode != 10000 {
		switch getNodeStateByGatewayIdResponse.ErrCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway", nil, res)
			return
		case 30020:
			result.JsonReply("Operation_DB_Error", nil, res)
			return
		case 30017:
			result.JsonReply("Parameter_format_error", nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful",getNodeStateByGatewayIdResponse, res)
}