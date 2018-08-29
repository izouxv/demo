package util

type reply struct {
	Code   int32       `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

const (
	Successful    = "Successful"
	SystemError   = "System_error"
	UploadFail    = "upload_fail"
	DeleteFail    = "delete_fail"
)

var CodeMap = map[string]reply{
	Successful:   {10000, "成功", nil},
	SystemError:  {10001, "系统异常", nil},
	UploadFail:   {39001, "上传文件失败",nil},
	DeleteFail:   {39002, "删除文件失败",nil},
}

func GetCodeAndMsg(msg string) reply {
	return CodeMap[msg]
}
