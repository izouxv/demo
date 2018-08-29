package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"cotx-http/result"
	"cotx-http/utils"
	"cotx-http/pb"
	log "github.com/cihub/seelog"
	"fmt"
	"strconv"
	"cotx-http/rpcClient"
	"os"
	"context"
)

func SendInstruction (res http.ResponseWriter,req *http.Request, param httprouter.Params){
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgwset := &pb.ReqInstruction{UserId:userinfo.Uid,GatewayId:gatewayid}
	pbType := utils.GetHttpDataGwSet(req, "application/json;charset=UTF-8", reqgwset)
	log.Info("pbType == :",pbType)
	if pbType != 10000 {
		if pbType == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		if pbType == 30022 {
			result.JsonReply("Instruction_Error",nil,res)
			return
		}
		if pbType == 10001 {
			result.JsonReply("System_error",nil, res)
			return
		}
		if pbType == 20001 {
			result.JsonReply("Body_is_incorrect_or_empty",nil, res)
			return
		}
		if pbType == 30026 {
			result.JsonReply("Set_Error",nil,res)
			return
		}
		if pbType == 30025 {
			result.JsonReply("Value_IS_Error",nil,res)
			return
		}
		if pbType == 30027 {
			result.JsonReply("Gateway_UnResponse",nil,res)
			return
		}
		if pbType == 30033 {
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		}
		result.JsonReply("Unknown_error",nil, res)
		return
	}
	result.JsonReply("Successful",nil,res)
}

/*设置第三方云平台信息*/
func SetIOTServer(w http.ResponseWriter,r *http.Request,param httprouter.Params)  {
	log.Info("http-router/start set iot server")
	userinfo ,ok := r.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,w)
		return
	}
	log.Info("当前登录帐号信息",userinfo)
	gatewaid := param.ByName("gwid")

	//创建rpc对象实例
	reqIOTServer := &pb.ReqIOTServer{}
	reqIOTServer.UserId = userinfo.Uid
	reqIOTServer.GatewayId = gatewaid

	//获取云平台id，判断需要的设置的云平台类型
	bodyType := r.Header.Get("Content-Type")
	log.Infof("接口使用的请求头类型是 %s",bodyType)
	IotIdString := r.FormValue("IotId")
	if IotIdString == "" {
		log.Error("获取iotid参数为空")
		result.JsonReply("Parameter_is_null",nil,w)
		return
	}
	IotId,_ := strconv.Atoi(IotIdString)
	log.Info("要设置的云平台的id是：",IotId)
    reqIOTServer.IOTId = int32(IotId)

	//根据云平台id获取响应的的信息 并对空值记性判断和处理
	if IotId == 1 {
		log.Info("设置ttn云平台")

		//获取域名或者ip
		hostname := r.FormValue("HostName")
		//if hostname == "" {
		//	log.Error("获取hostname参数错误")
		//	result.JsonReply("Parameter_is_null",nil,w)
		//}
		reqIOTServer.HostName = hostname

		//获取端口号
		portString := r.FormValue("Port")
		port ,_ := strconv.Atoi(portString)
		//if port == 0 {
		//	log.Error("获取port参数错误")
		//	result.JsonReply("Parameter_is_null",nil,w)
		//	return
		//}
		reqIOTServer.Port = int32(port)

		//获取相应的网关mac地址 （这个参数是需要具体考虑下是不是只展示就不用传了）
		gwMAC := r.FormValue("GwMac")
		//if gwMAC == "" {
		//	log.Error("获取gwMac参数错误")
		//	result.JsonReply("Parameter_is_null",nil,w)
		//	return
		//}
		reqIOTServer.GWMac = gwMAC

		//rpc调用后台操作
		resIOTServer ,err:= rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			result.JsonReply("SysTem_ERROR",nil, w)
			return
		}
		//判断返回结果
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return
	}

	/*aws云平台信息设置*/
	if IotId  == 2{
		 log.Info("设置AWS云平台的信息")
		 /*创建文件目录*/
		FilePath1 := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload")
		if !utils.CheckFileIsExist(FilePath1) {
			os.Mkdir(FilePath1,os.ModePerm)
		}

		//验证二级目录存不存在
		FilePath2 := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d",userinfo.Uid)
		if !utils.CheckFileIsExist(FilePath2) {
			os.Mkdir(FilePath2,os.ModePerm)
		}

		// 验证三级目录
		FilePath3 := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/%s/aws",userinfo.Uid,gatewaid)
		if !utils.CheckFileIsExist(FilePath3) {
			os.Mkdir(FilePath3,os.ModePerm)
		}

		//获取域名或者ip
		hostname := r.FormValue("HostName")
		if hostname == "" {
			log.Error("aws/获取hostname参数失败")
			result.JsonReply("Parameter_is_null",nil,w)
			return
		}
		reqIOTServer.HostName = hostname
		/*获取根证书*/
		rootFile ,awsRootFileHead,err := r.FormFile("RootCertificate")
		if err != nil {
			log.Error("aws/获取根证书参数失败")
			result.JsonReply("Update_File_Error",nil,w)
			return
		}
        defer rootFile.Close()

        //生成文件路径
        rootFilePath := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/%s/aws/%s",userinfo.Uid,gatewaid,awsRootFileHead.Filename)//线上文件路径
		//rootFilePath := fmt.Sprintf("D://FilerRoot//%d//aws//%s",userinfo.Uid,awsRootFileHead.Filename)//线下路径
        //调用工具类进行文件操作
        code := utils.CopyFileTo(rootFilePath,rootFile)
		if code != 200 {
			log.Errorf("上传根证书失败，code== %d",code)
			result.JsonReply("Update_File_Error",nil,w)
			return
		}
		rootPath := fmt.Sprintf("http://gateway.cotxnetworks.com/upload/%d/%s/aws/%s",userinfo.Uid,gatewaid,awsRootFileHead.Filename)
		reqIOTServer.RootCertificate = rootPath

		/*获取客户端证书*/
		clientFile,awsClientFileHead ,err := r.FormFile("ClientCertificate")
		if err != nil {
			log.Error("aws/获取客户端证书参数失败")
			result.JsonReply("Update_File_Error",nil,w)
			return
		}
		defer clientFile.Close()

		//生成文件路径
		clientFilePath :=fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/%s/aws/%s",userinfo.Uid,gatewaid,awsClientFileHead.Filename)//线上路径
		//clientFilePath :=fmt.Sprintf("D://FilerRoot//%d//aws//%s",userinfo.Uid,awsClientFileHead.Filename)// 本地测试路径
		//调用工具类操作文件
		code = utils.CopyFileTo(clientFilePath,clientFile)
		if code != 200 {
			log.Errorf("上传客户端证书失败，code==%d",code)
			result.JsonReply("Update_File_Error",nil,w)
			return
		}
		clientPath := fmt.Sprintf("http://gateway.cotxnetworks.com/upload/%d/%s/aws/%s",userinfo.Uid,gatewaid,awsClientFileHead.Filename)
		reqIOTServer.ClientCertificate = clientPath

		/*获取客户端密钥*/
		clientFileKey,awsClientFileKeyHead,err := r.FormFile("ClientKey")
		if err != nil {
			log.Error("aws/获取客户端密钥参数失败")
			result.JsonReply("Update_File_Error",nil,w)
			return
		}
		defer clientFileKey.Close()
		clientFileKeyPath := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/%s/aws/%s",userinfo.Uid,gatewaid,awsClientFileKeyHead.Filename)//线上路径
		//clientFileKeyPath := fmt.Sprintf("D://FilerRoot//%d//aws//%s",userinfo.Uid,awsClientFileKeyHead.Filename)//本地测试路径
		// 调用工具类操作文件
		code = utils.CopyFileTo(clientFileKeyPath,clientFileKey)
		if code != 200 {
			log.Info("上传客户端密钥失败",code)
			result.JsonReply("Update_File_Error",nil,w)
			return
		}
		clientKeyPath := fmt.Sprintf("http://gateway.cotxnetworks.com/upload/%d/%s/aws/%s",userinfo.Uid,gatewaid,awsClientFileKeyHead.Filename)
		reqIOTServer.ClientKey = clientKeyPath
		devicename := r.FormValue("devicename")
		reqIOTServer.DeviceName = devicename
		/*实现完成rpc调用*/
		resIOTServer ,err := rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			result.JsonReply("SysTem_ERROR",nil, w)
			return
		}
		//结果判断返回
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return

	}
	/*设置阿里云*/
	if IotId == 4 {
		log.Info("iotserver/开始设置阿里云")
		productKey := r.FormValue("productkey")
		if productKey == "" {
			result.JsonReply("Parameter_format_error",nil,w)
			return
		}
		deviceName := r.FormValue("devicename")
		if deviceName == "" {
			result.JsonReply("Parameter_format_error",nil,w)
			return
		}
		deviceKey := r.FormValue("devicekey")
		if deviceKey  == ""{
			result.JsonReply("Parameter_format_error",nil,w)
			return
		}
		nodeKey := r.FormValue("nodekey")
		if nodeKey  == ""{
			result.JsonReply("Parameter_format_error",nil,w)
			return
		}
		//rpc调用
		reqIOTServer.ProductKey = productKey
		reqIOTServer.DeviceName = deviceName
		reqIOTServer.DeviceKey  = deviceKey
		reqIOTServer.NodeKey    =  nodeKey
		reqIOTServer.IotNum = int32(IotId)
		log.Infof("设置阿里云参数 productkey:%s ,devicename:%s ,devicekey:%s  , nodekey:%s  ",productKey,deviceName,deviceKey,nodeKey)
		resIOTServer,err := rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			log.Error("IOT/rpc 调用失败",err)
			result.JsonReply("System_error",nil,w)
			return
		}
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return
	}

	/*本地云平台设置*/
	if IotId == 7 {
		log.Info("iotserver/开始设置本地云平台")
		resIOTServer,err := rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			log.Error("IOT/rpc 调用失败",err)
			result.JsonReply("System_error",nil,w)
			return
		}
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return

	}
	/*百度云平台设置*/
	if IotId == 6 {
		log.Info("iotserver/开始设置百度云平台")
		hostname := r.FormValue("hostname")
		username := r.FormValue("username")
		password := r.FormValue("passwoed")
		topic     := r.FormValue("topic")
		reqIOTServer.UserName = username
		reqIOTServer.HostName = hostname
		reqIOTServer.Password = password
		reqIOTServer.DeviceKey = topic
		resIOTServer,err := rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			log.Error("IOT/rpc 调用失败",err)
			result.JsonReply("System_error",nil,w)
			return
		}
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return
	}
	/*onenet平台设置*/
	if IotId == 6 {
		log.Info("iotserver/开始设置onenet平台")
		productId := r.FormValue("productId")
		registerCode := r.FormValue("registerCode")
		reqIOTServer.ProductKey = productId
		reqIOTServer.DeviceKey = registerCode
		resIOTServer,err := rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			log.Error("IOT/rpc 调用失败",err)
			result.JsonReply("System_error",nil,w)
			return
		}
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return

	}
	if IotId == 8 {
		hostname := r.FormValue("HostName")
		portString := r.FormValue("Port")
		reqIOTServer.HostName = hostname
		port ,_ := strconv.Atoi(portString)
		reqIOTServer.Port = int32(port)
		log.Infof("设置自定义云平台hostname= %s,port = %d",hostname,port)
		resIOTServer,err := rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			log.Error("IOT/rpc 调用失败",err)
			result.JsonReply("System_error",nil,w)
			return
		}
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return

	}

	//设置alay 云平台
	if IotId == 9 {
		confFile,confFileHeard,err := r.FormFile("conf")
		if err != nil {
			result.JsonReply("System_error",nil,w)
			return
		}
		defer confFile.Close()
		//验证二级目录存不存在
		FilePath2 := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/",userinfo.Uid)
		if !utils.CheckFileIsExist(FilePath2) {
			log.Infof("文件路径不存在,将创建文件 %s",FilePath2)
			os.Mkdir(FilePath2,os.ModePerm)
		}

		FilePath4 := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/%s/",userinfo.Uid,gatewaid)
		if !utils.CheckFileIsExist(FilePath4) {
			os.Mkdir(FilePath4,os.ModePerm)
		}
		// 验证三级目录
		FilePath3 := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/%s/alay/",userinfo.Uid,gatewaid)
		if !utils.CheckFileIsExist(FilePath3) {
			os.Mkdir(FilePath3,os.ModePerm)
		}

		confFilePath := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/%d/%s/alay/%s",userinfo.Uid,gatewaid,confFileHeard.Filename)
		code := utils.CopyFileTo(confFilePath,confFile)
		if code!= 200 {
			log.Info("上传alay配置文件失败",code)
			result.JsonReply("Update_File_Error",nil,w)
			return
		}
		confFileUrl := fmt.Sprintf("http://gateway.cotxnetworks.com/upload/%d/alay/%s",userinfo.Uid,confFileHeard.Filename)
		reqIOTServer.ClientCertificate = confFileUrl
		resIOTServer,err := rpcClient.SetIOTServer(reqIOTServer)
		if err != nil {
			log.Error("IOT/rpc 调用失败",err)
			result.JsonReply("System_error",nil,w)
			return
		}
		if resIOTServer.ErrCode != 10000 {
			if resIOTServer.ErrCode == 404 {
				result.ResCode(http.StatusNotFound,w)
				return
			}
			if resIOTServer.ErrCode == 30022 {
				result.JsonReply("Instruction_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 10001 {
				result.JsonReply("System_error",nil, w)
				return
			}
			if resIOTServer.ErrCode == 20001 {
				result.JsonReply("Body_is_incorrect_or_empty",nil, w)
				return
			}
			if resIOTServer.ErrCode == 30026 {
				result.JsonReply("Set_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30025 {
				result.JsonReply("Value_IS_Error",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30027 {
				result.JsonReply("Gateway_UnResponse",nil,w)
				return
			}
			if resIOTServer.ErrCode == 30033 {
				result.JsonReply("SysTem_ERROR",nil, w)
				return
			}
			result.JsonReply("Unknown_error",nil, w)
			return
		}
		result.JsonReply("Successful",nil, w)
		return

	}
	//参数不合法结果返回
	result.JsonReply("Parameter_format_error",nil,w)
}

/*设置手动拍照*/
func SetGatewayPhoto(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	log.Info("http/start set gateway photo")
	instruction := &pb.ReqInstruction{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", instruction)
	fmt.Println("errCode = ", errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	//上下文中的用户信息
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}

	//获取网关信息
	gatewayid := param.ByName("gwid")
	instruction.GatewayId= gatewayid
	instruction.UserId=userinfo.Uid

	//rpc调用
	resgateway ,err:= rpcClient.SetGatewayPhoto(instruction)
	if err != nil || resgateway.ErrCode != 10000 {
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
	result.JsonReply("Successful",resgateway.ErrCode, res)
}

/*设置自动摄像*/
func SetGatewayVideo(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	log.Info("http/start set gateway video")
	setSwitch := &pb.ReqSwitch{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", setSwitch)
	fmt.Println("errCode = ", errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	//上下文中的用户信息
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return

	}
	//获取网关信息
	gatewayid := param.ByName("gwid")
	setSwitch.GatewayId= gatewayid
	setSwitch.UserId=userinfo.Uid
	//rpc调用
	resgateway ,err:= rpcClient.SetGatewayVideo(setSwitch)
	if err != nil || resgateway.ErrCode != 10000 {
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
	result.JsonReply("Successful",resgateway.VideoCode, res)
}

func GetInstructionState(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	instruction := param.ByName("instruction")
	var getInstructionStateRequest = new(pb.GetInstructionStateRequest)
	getInstructionStateRequest.GatewayId = gatewayid
    getInstructionStateRequest.Instruction = instruction
    getInstructionStateRequest.UserId = userinfo.Uid
    getInstructionResponse,_:= rpcClient.GetGatewaySetRpcClient().GetInstructionState(context.Background(),getInstructionStateRequest)
	if getInstructionResponse.ErrCode != 10000 {
		switch getInstructionResponse.ErrCode {
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		case 30027:
			result.JsonReply("Gateway_UnResponse",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",getInstructionResponse.InstructionState,res)
}