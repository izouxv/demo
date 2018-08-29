package po

import (
	"petfone-http/pb"
)

//登录：返回用户信息
type LoginPo struct {
	//账户基本资料
	Uid        int32  `json:"uid"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Token      string `json:"token"`
	LoginState int32  `json:"loginState"`
	State      int32  `json:"state"`
	Code       int32  `json:"code"`
	//账户基本属性
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Gender    int32  `json:"gender"`
	Birthday  int64  `json:"birthday"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
	Address   string `json:"address"`
	//账户业务属性
	Radius    int32 `json:"radius"`
	Map       pb.Map `json:"map"`
}

//用户查询用户信息
type AccountFriendPo struct {
	//账户基本资料
	Uid      int32  `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	//账户基本属性
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Gender int32  `json:"gender"`
	Avatar string `json:"avatar"`
}

//修改密码json
type UserPwd struct {
	Action      string `json:"action"`
	Username    string `json:"username"`
	Password    string `json:"pwd"`
	NewPassword string `json:"newpwd"`
	MCode       string `json:"mcode"`
}

//反馈json
type FeedBackPo struct {
	Source      string    `json:"source"`
	Uid         int32     `json:"uid"`
	Types       string    `json:"types"`
	Description string    `json:"description"`
	Contact     string    `json:"contact"`
	Address     []*string `json:"address"`
}

//文件服务json
type FidPo struct {
	Fid string `json:"fid"`
}

//设备信息json
type DevicePo struct {
	Did             int32        `json:"did"`
	Pid             int32        `json:"pid"`
	Types           pb.DeviceTypes	`json:"types"`
	Permit          pb.DevPermit `json:"permit"`
	DeviceName      string       `json:"deviceName"`	//设备名称
	DeviceVersion   string       `json:"deviceVersion"`	//设备版本
	DeviceMac		string       `json:"deviceMac"`	//设备mac
	SoftwareVersion string       `json:"softwareVersion"` //软件名称
	LedModel        int32		 `json:"led_model"`    		//设备灯闪烁模式
	LedColor        int32        `json:"led_color"`    		//设备灯颜色
	LedLight        int32        `json:"led_light"`    		//设备灯亮度
	LedState        int32        `json:"led_state"`    		//宠物设备灯开关状态
	AudioId         int32		 `json:"audio_id"`    		//宠物播放录音的id
}

//通知信息json
type NoticePo struct {
	Id        int32  `json:"id"`
	Froms     int32  `json:"froms"`
	Tos       int32  `json:"tos"`
	Tou       string `json:"tou"`
	Nstate    int32  `json:"nstate"` //3（待确认）-1（确认）
	Types     int32  `json:"types"`  //1.分享设备通知
	Info      string `json:"info"`
	StartTime int64  `json:"startTime"` //创建时间
}

//宠物信息json
type PetinfoPo struct {
	Pid        int32        `json:"pid"`
	Did        int32        `json:"did"`
	Avatar     string       `json:"avatar"`
	Nickname   string       `json:"nickname"`
	Breed      int32        `json:"breed"`
	Gender     int32        `json:"gender"`
	Birthday   int64        `json:"birthday"`
	Weight     float32		`json:"weight"`
	Somatotype int32        `json:"somatotype"`
	Duration	int32		`json:"duration"`
	Brightness	int32		`json:"brightness"`
	CreateTime int64        `json:"create_time"`
	Permit     pb.PetPermit `json:"permit"`
	Trains	[]*pb.PetTrainReply `json:"trains"`
}

//宠物训练信息
type PetTrainDayPo struct {
	Day	[]*pb.PetTrainReply	`json:"day"`
}

//品种信息
type BreedInfoPo struct {
	Id      int32  `json:"id"`
	Name	string `json:"name,omitempty"`
	NameCh	string `json:"nameCh,omitempty"`
	NameEn	string `json:"nameEn,omitempty"`
	Address string `json:"address"`
	Types   int32  `json:"types"`
}
type BreedInfoPoSlice []*BreedInfoPo

func (s BreedInfoPoSlice) Len() int           { return len(s) }
func (s BreedInfoPoSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BreedInfoPoSlice) Less(i, j int) bool { return s[i].Id < s[j].Id }

//常见问题信息
type FaqCommonPo struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Parent int32  `json:"parent"`
	Info   string `json:"info"`
}

//坐标点信息
type Coordinate struct {
	NowTime   int64    `json:"now_time"` //时间点
	Longitude float32  `json:"longitude"`
	Latitude  float32  `json:"latitude"`
	State     pb.State `json:"state"`
}

type CoordinateSlice []*Coordinate

func (s CoordinateSlice) Len() int           { return len(s) }
func (s CoordinateSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s CoordinateSlice) Less(i, j int) bool { return s[i].NowTime < s[j].NowTime }

//宠物运动信息
type DaysExerciseDataPo struct {
	Pid  int32                `json:"pid"`  //宠物id
	Pdid int32                `json:"pdid"` //宠物设备id
	Data []*DayExerciseDataPo `json:"data"`
	Code int32				`json:"code"`
}

//宠物运动属性
type DayExerciseDataPo struct {
	DayTime        	int64         `json:"dayTime"`
	CardioTimes    	int64         `json:"cardioTimes"`    //有氧耗时
	StrenuousTimes 	int64         `json:"strenuousTimes"` //无氧耗时
	Steps          	int32         `json:"steps"`          //运动步数
	Cals           	float32       `json:"cals"`           //消耗卡
	Coordinates    	[]*Coordinate `json:"coordinates"`      //时间与坐标点
}

//宠物运动信息1.1
type MotionDataPo struct {
	Data []*DayMotionDataPo `json:"data"`
	Code int32				`json:"code"`
}
//宠物运动信息
type DayMotionDataPo struct {
	DayTime        		int64         	`json:"day_time"`
	StepsTotal          int32         	`json:"steps_total"`				//运动步数
	CalorieTotal		float32       	`json:"calorie_total"`			//消耗卡
	CardioTimesTotal    int64         	`json:"cardio_times_total"`		//有氧耗时
	StrenuousTimesTotal	int64         	`json:"strenuous_times_total"` //无氧耗时
	CardioDurationMinuteTotal    	int32	`json:"cardio_duration_minute_total"`		//有氧耗时
	StrenuousDurationMinuteTotal	int32	`json:"strenuous_duration_minute_total"` //无氧耗时
	Records				[]*Record		`json:"records"`
}
type Record struct {
	TimeRecord		string		`json:"time_record"`//时间段
	Steps			int32		`json:"steps"`//步数
	Calorie        	float32		`json:"calorie"`//消耗卡
	CardioTimes  	int64		`json:"cardio_times"`//有氧耗时
	StrenuousTimes 	int64		`json:"strenuous_times"`//无氧耗时
	CardioDurationMinute	int32	`json:"cardio_duration_minute"`		//有氧耗时
	StrenuousDurationMinute	int32	`json:"strenuous_duration_minute"` //无氧耗时
	ImageInfo   	*ImageInfo	`json:"image_info"`//图片信息
	Pdid        	int32		`json:"pdid"`
}
type ImageInfo struct {
	Url			string	`json:"url"`		//原图地址
	ViewUrl		string	`json:"view_url"`
	Name		string	`json:"name"`
	Size		int32	`json:"size"`		//大小
	Width		int32	`json:"width"`		//宽度
	Height		int32	`json:"height"`		//高度
	FileUrl		string	`json:"file_url"`   //文件url
}

type SharePo struct {
	Pid       int32           `json:"pid"`
	OwnerInfo pb.AccountReply `json:"owner_info"` //主人信息
	Members   MemberInfo      `json:"members"`    //共享用户信息
}

type MemberInfo struct {
	Member     pb.AccountReply `json:"member"`      //共享用户信息
	CreateTime int64           `json:"create_time"` //共享时间
}

//宠聊返回数据
type ChatMsg struct {
	Types	int32		`json:"types"`
	Output interface{}	`json:"output"`
}
