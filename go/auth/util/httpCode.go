package util

type reply struct {
	Code   int32       `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

var CodeMap = map[string]reply{
	"Successful":   {10000, "成功", nil},
	"System_error": {10001, "系统异常", nil},
	//系统
	"Token_is_invalid":         			{10002, "Token失效或未登录，请重新登录", nil},
	"Permission_denied":          			{10003, "权限不足，拒绝访问", nil},
	"User_Kick_Out":              			{10004, "用户被踢出", nil},
	"Body_is_incorrect_or_empty":			{10006, "Body为空或输入有误", nil},
	"InvalidArgument":						{10007, "请求参数有误", nil},

	//账号
	"Username_is_incorrect_or_empty":    	{20001, "Username为空或格式有误", nil},
	"Password_is_incorrect_or_empty":    	{20002, "Password为空或输入有误", nil},
	"Nickname_is_incorrect_or_empty":    	{20003, "Nickname为空或输入有误", nil},
	"Account_Disable_to_Use":    		 	{20004, "账号不可用，请联系管理员", nil},
	"Account_Not_Active":    			 	{20005, "账号未激活，请通过邀请邮件激活", nil},
	"Account_Exception":    			 	{20006, "账号异常，请联系管理员", nil},




}

func GetCodeAndMsg(msg string) reply {
	return CodeMap[msg]
}
