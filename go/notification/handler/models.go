package handler

/*接收数据*/
type PayloadEmail struct {
	Aid            int64  `json:"aid"`
	Tid            int64  `json:"tid"`
	DeviceId       string `json:"deviceId"`
	AlarmName      string `json:"alarmName"`
	AlarmId        int64  `json:"alarmId"`
	AlarmEmail     string `json:"alarmEmail"`
	ALarmTemId     int64  `json:"alarmTemId"`
	AlarmTime      string  `json:"alarmTime"`
	ReceiveTime    string  `json:"receiveTime"`
	Trigger        string  `json:"trigger"`
	AlarmDes       string  `json:"alarmDesc"`
	Msg            map[string]interface{} `json:"msg"`
}




