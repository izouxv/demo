package result

//swagger:response reserror
type reply struct {
	//返回的响应码
	//
	Code	int32				`json:"code"`
	//响应码代表信息
	//
	Msg		string				`json:"msg"`
	//响应返回的参数，如果响应码不等于10000的话表示响应错误信息，响应返回参数为空
	//
	Result 	interface{}		`    json:"result"`
}

var CodeMap = map[string]reply {
	//成功
	"Successful"								: 	{10000, "成功", 							nil} ,
	//系统异常
	"System_error"								:	{10001, "系统繁忙", 						nil} ,
	"Unknown_error"								:	{10002, "未知错误",                    		nil} ,
	""											:	{10003, "服务接口维护（接口停用或维护）",	nil} ,
	//HTTP异常
	"Body_is_incorrect_or_empty"				:	{20001, "Body为空或输入有误", 				nil} ,

	//user参数异常
	"Source_is_incorrect_or_empty"				:	{30001, "Source为空或输入有误", 			nil} ,
	"Username_is_incorrect_or_empty"			:	{30002, "用户名为空或输入有误", 			nil} ,
	"Password_is_incorrect_or_empty"			:	{30003, "密码为空或输入有误", 				nil} ,
	"Account_does_not_existed"					:	{30004, "账号不存在", 						nil} ,
	"Account_does_not_activated"				:	{30005, "账号未激活", 						nil} ,
	"Nickname_is_incorrect_or_empty"			:	{30006, "昵称为空或输入有误", 				nil} ,
	"MobileCode_is_incorrect_or_empty"			:	{30007, "验证码错误", 			nil} ,
	"Account_already_existed"					:	{30008, "账号已存在", 						nil} ,
	"OldPassword_is_incorrect_or_empty"			:	{30009, "旧密码为空或输入有误", 			nil} ,
	"NewPassword_is_incorrect_or_empty"			:	{30010, "新密码为空或输入有误", 			nil} ,
	"Token_is_empty"							:	{30011, "token为空", 						nil} ,
	"Token_is_incorrect_or_lose_efficacy"		:	{30012, "token输入有误或失效", 				nil} ,
	"Session_is_lose_efficacy"					:	{30013, "登陆超时", 						nil} ,
	"Password_mismatched"						:	{30014, "密码错误", 						nil} ,
	"MobileCode_is_incorrect_or_lose_efficacy"  :   {30015, "验证码错误", 				nil} ,
	"CodeType_is_incorrect_or_empty"    		:   {30016, "CodeType为空或输入有误", 			nil} ,
	"Parameter_format_error"					:	{30017, "参数不合法", 						nil} ,
   "Account_abnormality"                        :   {30017,"帐号异常",nil},
    "Gateway_havebeen_binding"                  :   {30018,"网关已经被绑定了",nil},
    "Account_Auth"                              :   {30019,"该帐号已经被该用户授权过了",nil},
    "Operation_DB_Error"                        :   {30020,"内部数据错误",nil},
    "NOFind_Gateway"                            :   {30021,"未找到该网关",nil},
    "Instruction_Error"                         :   {30022,"指令码错误",nil},
    "NoFind_node_deveui"                        :   {30023,"未找到需要绑定的终端",nil},
    "Value_IS_Error"                            :   {30025,"参数错误",nil},
    "Set_Error"                                 :   {30026,"下发设置失败(网关当前可能处于离网络状态)",nil},
    "Gateway_UnResponse"                        :   {30027,"网关未响应",nil},
    "Parameter_is_null"                         :   {30028,"参数为空",nil},
    "Update_File_Error"                         :   {30029,"文件上传失败",nil},
    "File_IsNot_Exit"                           :   {30030,"请求文件不存或者网关未上传该文件到服务器",nil},
    "Gateway_Camera_Use"                        :   {30031,"网关摄像头被占用",nil},
    "No_Update_Device"                          :   {30032,"未发现新版本",nil},
    "SysTem_ERROR"                              :   {30033,"系统异常",nil},
    "Insufficient_permissions"                  :   {30034,"权限不足不能执行此操作",nil},
 }

func GetCodeAndMsg(msg string) reply {
	return CodeMap[msg]
}

