  package utils

import "regexp"

//正则验证规则
var (
	//正则验证规则
	REGEXP_MOBILE	= regexp.MustCompile("^1[3|4|5|7|8][0-9]{9}$")
	REGEXP_MAIL		= regexp.MustCompile("^\\w+([-+._]\\w+)*@\\w+([-.]\\w+)*\\.[a-z]+([-.][a-z]+)*$")
	REGEXP_NICKNAME	= regexp.MustCompile("^.{1,12}$")
	REGEXP_PWD		= regexp.MustCompile("^[a-z0-9A-Z]{6,32}$")
)
//获取随机字符串
const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)
//指令码设置
const (
	GatewayUpdata   = 1010//网关的升级指令
	CloseGateway    = 1011 //网关关机
	DisConnect_Net  = 1012//断开当前网络
	WifiCard        = 1013//
	WifiCardScan    = 1014
	ConnectWifi     = 1015
    GatewayReset    = 1016
    NBIOTSet        = 1017
    NFCSet          = 1018
    PowerModel      = 1019
    LoraSet         = 1020
    TakeVideo       = 1021
    AutoVideo       = 1022
    TakePhoto       = 1023
    AutoPhoto       = 1024
    AutoMusic       = 1025
    TakeMusic       = 1026
    DeletFile       = 1027
    SetSSH          = 1028
    UpLog           = 1029
    BackUp          = 1030
    RevertBackUp    = 1031
    RevertGateway   = 1032
    SetIP           = 1033
    SetDNS          = 1034
    HotSpot         = 1035
    SetHotSpot      = 1036
    GatewayName     = 1037
    PowerSwitch     = 1039
    UpdateGatewayING = 1042
    photo           = 1048
    SetUsbIp        = 1049
    SetUsbDns       = 1050
    SetwifiIp       = 1051
    SetWifiDns      = 1052
    Lora1301        = 1053
    Usbwifiscan     = 1054
    Usbwifi         = 1055
    UsbGCard        = 1056
    UsbconnectionWifi = 1057
    UsbHotSpot        = 1058
	UsbSetHotSpot     = 1059
	UnConnectWifi     = 1060
	UnConnectUsbWifi  = 1061
	deleteBackUpFile  =  1062
	setGatewaySysWarn = 1063 //系统告警设置
	setGatewayNetWarn = 1064//网关状态告警设置
	setBleScan        = 1065

)
