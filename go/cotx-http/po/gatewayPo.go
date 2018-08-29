package po

import "cotx-http/pb"

type UserGateway struct {
	GatewayId string   `json:"gateway_id"`
	AppId     string   `json:"app_id"`
} 
type AuthoriseAccount struct {
     Uid        int32     `json:"uid"`
     UserName   string    `json:"user_name"`
     Avatar     int32     `json:"avatar"`
     NickName   string    `json:"nick_name"`
}
//swagger:parameters  reqbandgw  valibandgw
type ReqBandGw struct {
	//用户与网关绑定需要的参数
	//扫描到的网关的唯一标识
	//
	//unique:true
	//in: body
	 MAC             string    `json:"mac"`
	//unique:true
	//in: body
	ApplicationId    int64     `json:"ApplicationId"`

}
type ResBandGw struct {
    GatewayID   int32  `json:"GatewayID"`
    AppEUI      string `json:"AppEUI"`
}
//swagger:response    resbandgw
type reply struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	ResBandGw	        `json:"result"`
}
//swagger:parameters  reqauth reqvaliauth
type ReqAuthAccount struct {
	//网关的id
	//
	//unique:true
	//in    :body
	GatewayID   int32     `json:"GatewayID"`
	//账户名称
	//
	//unique:true
	// in   :body
	UName       string    `json:"UName"`
}
//swagger:parameters  reqdeletaccount
type  ReqDeletAccount struct {
	//帐号名称
	//
	//unique:true
	//in     :body
	UName  string     `json:"UName"`
}
//swagger:parameters    requnwoundaccount  getgwposs
type RequnwoundAccount struct {
	//网关id
	//
	//unique:true
	//in    : body
	GatewayID   int32  `json:"GatewayID"`
}
//swagger:parameters NULL NULL1 NULL2 NULL3 NULL4 NULL5 NULL6 NULL7 NULL8 NULL9 NULL10 NULL11 NULL12 NULL13 NULL14 NULL15 NULL16 NULL17 NULL18 NULL19 NULL20
type NULL struct {
	//NULL
}
//swagger:response respos
type replyPos struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResUserGateways	        `json:"result"`
}
//swagger:response resAccounts
type replyAccounts struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResAccounts	        `json:"result"`
}
//swagger:response resGwNetState
type replyGwNetState struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwNetState	        `json:"result"`
}
//swagger:response resGwState
type replyGwState struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwState	      `json:"result"`
}
//swagger:response resGwFile
type replyGwFile struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwFile	      `json:"result"`
}
//swagger:response resGwVideos
type replyGwVideos struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwVideos	      `json:"result"`
}
//swagger:response resGwPhotos
type replyGwPhotos struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwPhotos	      `json:"result"`
}
//swagger:response resWifiScan
type replyGwWifiScan struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwWifiScans	  `json:"result"`
}
//swagger:response resWifiAddress
type replyGwWifiAddress struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwWifiAddress	  `json:"result"`
}
//swagger:response resWifiDNS
type replyGwWifiDNS struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwWifiDNS	  `json:"result"`
}
//swagger:response resCableAddress
type replyGwCableAddress struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwCableAddress	  `json:"result"`
}
//swagger:response resCableDNS
type replyGwCableDNS struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwCableDNS	  `json:"result"`
}
//swagger:response resGwMessage
type replyGwMessage struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwMessage	  `json:"result"`
}
//swagger:response resGwLora
type replyGwLora struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResGwLora	  `json:"result"`
}
//swagger:response resGwPowerModel
type replyGwPowerModel struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResPowerModel	  `json:"result"`
}
//swagger:response resGwMusicSet
type replyGwMusicSet struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
}
//swagger:response resGwPhotoSet
type replyGwPhotoSet struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
}
//swagger:response resGwVideoSet
type replyGwVideoSet struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
}
//swagger:response resGwAppEui
type replyGwAppEui struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空,响应码等于10000，有返回参数的话会显示在result中
	//
	Result 	pb.ResAppEui	  `json:"result"`
}