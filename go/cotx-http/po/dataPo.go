package po

//修改密码json
type UserPwd struct {
	Uid			int32		`json:"uid"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	NewPassword	string		`json:"newPassword"`
	//MCode		string		`json:"mcode"`
}
//重置密码json
type ResetPwd struct {
	Username 	string 		`json:"username"`
	Password	string 		`json:"password"`
	Code    	string  	`json:"code"`
}
type DevEui struct {
	DevEui string   `json:"DevEui"`
}
type DevEuis struct {
	DevEuis *[]DevEui `json:"DevEuis"`
}
