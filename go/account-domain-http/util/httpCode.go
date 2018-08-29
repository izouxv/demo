package util

type reply struct {
	Code   int32       `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

var CodeMap = map[string]reply{
	"Successful":                        {10000, "成功", nil},
	"System_error":                      {10001, "系统异常", nil},

	"Session_is_invalid":                {10002, "Session失效或未登录，请重新登录", nil},
	"Permission_denied":                 {10003, "权限不足，拒绝访问", nil},
	"User_Kick_out":                     {10004, "用户被踢出", nil},
	"Get_context_userInfo_error":        {10005, "登陆信息失效，请重新登录", nil},

	"Body_is_incorrect_or_empty":        {10006, "Body为空或输入有误", nil},

	"Oldpassword_is_incorrect":          {27001, "旧密码不正确", nil},
	"Per_page_is_incorrect_or_empty":    {27002, "Per_page为空或输入有误", nil},
	"Name_is_incorrect_or_empty":        {27003, "Name为空或输入有误", nil},
	"Category_is_incorrect_or_empty":    {27004, "类别为空或输入有误", nil},
	"Did_is_incorrect_or_empty":         {27005, "Did为空或输入有误", nil},
	"Username_is_incorrect_or_empty":    {27006, "Username为空或格式有误", nil},
	"Password_is_incorrect_or_empty":    {27007, "Password为空或输入有误", nil},
	"Account_already_domian":            {27008, "用户已经存在该域中，可以选择修改角色或邀请他人", nil},
	"Equipment_already_exist":           {27009, "设备已存在", nil},
	"Equipment_does_not_exist":          {27010, "设备不存在", nil},
	"Nid_is_incorrect_or_empty":         {27011, "Nid为空或输入有误", nil},
	"DomainName_is_incorrect_or_empty":  {27012, "DomainName为空或输入有误", nil},
	"Mac_is_incorrect_or_empty":         {27013, "Mac为空或输入有误", nil},
	"Page_is_incorrect_or_empty":        {27014, "Page为空或输入有误", nil},
	"Username_is_not_valid":             {27015, "用户名格式不正确", nil},
	"EquipmentLocation_is_nil":          {27016, "设备位置数据为空", nil},
	"Equipment_is_not_nil_in_domain":    {27017, "域中尚有设备，不能删除", nil},
	"Did_is_default":                    {27018, "该域是默认域，不能删除", nil},
	"User_is_not_nil_in_did":            {27019, "域中尚有成员，不能删除", nil},
	"User_is_nil_in_did":                {27020, "域中暂时没有成员", nil},
	"Did_is_limit":                      {27021, "超出域限制", nil},
	"User_does_not_exist":               {27022, "用户不存在", nil},
	"Password_is_incorrect":             {27023, "密码不正确", nil},
	"Can_not_find_domains_by_uid":       {27024, "未查询到该区域", nil},
	"Nickname_is_empty":                 {27025, "nickName为空或输入有误", nil},
	"User_does_exist":                   {27026, "用户已存在", nil},
	"Oldpassword_is_nil":                {27027, "旧密码为空或输入有误", nil},
	"Version_is_incorrect_or_empty":     {27028, "Version为空或输入有误", nil},
	"Description_is_incorrect_or_empty": {27029, "Description为空或输入有误", nil},
	"AppEUI_is_incorrect_or_empty":      {27030, "AppEUI为空或输入有误", nil},
	"Aid_is_incorrect_or_empty":         {27031, "Aid为空或输入有误", nil},
	"Have_node_in_app":                  {27032, "应用中尚有设备，不能删除", nil},
	"DevEUI_is_incorrect_or_empty":      {27033, "devEUI 不能为空", nil},
	"AppKey_is_incorrect_or_empty":      {27034, "appKey不能为空", nil},
	"Node_location_does_not_exist":      {27035, "设备动态信息暂未上传", nil},
	"Application_does_not_exist":        {27036, "应用不存在", nil},
	"New_password_is_nil":               {27037, "新密码为空或输入有误", nil},
	"Token_is_nil":                      {27038, "token不能为空", nil},
	"Token_is_nil_or_invalid":           {27039, "token失效或不存在", nil},
	"Rid_is_incorrect_or_empty":         {27040, "Rid为空或输入有误", nil},
	"uid_is_incorrect_or_empty":         {27041, "被修改人uid为空", nil},
	"active_user_is_failed":             {27042, "账号未激活", nil},
	"DevEUI_is_err":                     {27043, "DevEUI无效", nil},
	"AppKey_is_err":                     {27044, "AppKey无效", nil},
	"AppEUI_is_err":                     {27045, "AppEUI无效", nil},
	"Param_is_incorrect_or_empty ":      {27046, "param为空或输入有误", nil},
	"Value_is_incorrect_or_empty ":      {27047, "Value为空或输入有误", nil},
	"Type_is_incorrect_or_empty ":       {27048, "Type为空或输入有误", nil},
	"DevEUI_is_exist":                   {27049, "DevEUI已经存在或格式错误", nil},
	"AppEUI_is_exist":                   {27050, "AppEUI已经存在或格式错误", nil},
	"Json_is_error":                     {27051, "json格式有误", nil},
	"Data_is_nil":                       {27052, "数据为空", nil},
	"Alert_is_nil":                      {27053, "没有告警", nil},
	"abp_is_incorrect_or_empty":         {27054, "abp不能为空或输入有误", nil},
	"classC_is_incorrect_or_empty":      {27055, "classC不能为空或输入有误", nil},
	"devAddr_is_incorrect_or_empty":     {27056, "DevAddr为空或输入有误", nil},
	"appsKey_is_incorrect_or_empty":     {27057, "appsKey为空或输入有误", nil},
	"nwksKey_is_incorrect_or_empty":     {27058, "nwksKey为空或输入有误", nil},
	"NwsKey_is_err":                     {27059, "nwksKey无效", nil},
	"DevAddr_is_err":                    {27060, "DevAddr无效", nil},
	"Is_Not_ABP":                        {27061, "不是ABP方式", nil},
	"User_does_not_exist_in_did":        {27062, "用户不存在", nil},
	"VersionCode_is_incorrect_or_empty": {27063, "版本号不能为空或输入有误", nil},
	"VersionName_is_incorrect_or_empty": {27064, "版本名称不能为空或输入有误", nil},
	"md5_is_incorrect_or_empty":         {27065, "md5不能为空或输入有误", nil},
	"filename_is_nil":                   {27066, "版本文件名不能为空或输入有误", nil},
	"filelength_is_incorrect_or_empty":  {27067, "版本长度不能为空或输入有误", nil},
	"Vid_is_incorrect_or_empty":         {27068, "vid不能为空或输入有误", nil},
	"File_is_exist":                     {27069, "版本不存在", nil},

	"AdvName_is_incorrect_or_empty":    {27071, "广告名为空或输入有误", nil},
	"State_is_incorrect_or_empty":      {27072, "广告状态参数为空或输入有误", nil},
	"StartTime_is_incorrect_or_empty":  {27073, "广告开始时间为空或输入有误", nil},
	"EndTime_is_incorrect_or_empty":    {27074, "广告结束时间为空或输入有误", nil},
	"Advertiser_is_incorrect_or_empty": {27075, "广告商为空或输入有误", nil},
	"No_Advertisement_Can_Be_Find":     {27076, "广告不存在", nil},
	"Id_is_incorrect_or_empty":         {27077, "ID为空或输入有误", nil},
	"FileUrl_is_incorrect_or_empty":    {27078, "文件Url为空或输入有误", nil},
	"Advertiser_is_newest":             {27079, "广告为最新广告", nil},

	"Domain_Can_Not_Delete":            {27080, "先清空域中数据才能进行删除操作", nil},

	"Status_is_incorrect_or_empty":     {27081, "Status为空或输入有误", nil},
	"Device_is_incorrect_or_empty":     {27082, "Device为空或输入有误", nil},
	"No_RadacatVersion_Can_Be_Find":    {27083, "升级版本不存在", nil},

	"ActionLogType_is_incorrect_or_empty":    {27084, "操作类型为空或输入有误", nil},

	"Fid_is_incorrect_or_empty":              {27085, "fid为空或输入有误", nil},
	"Pid_is_incorrect_or_empty":              {27086, "pid为空或输入有误", nil},

	"TenantId_is_incorrect_or_empty":         {27087, "租户id为空或输入有误", nil},
	"Type_is_incorrect_or_empty":             {27088, "类型为空或输入有误", nil},
	"No_Feedback_Can_Be_Find":                {27089, "反馈不存在", nil},
	"Order_is_incorrect_or_empty":            {27090, "order为空或输入有误", nil},

	"Params_error":                           {27091, "参数为空或输入有误", nil},
	"FileName_is_incorrect_or_empty":         {27092, "文件名为空或输入有误", nil},
	"User_does_not_exist_in_db":              {27093, "不是测试账号", nil},
	"Source_is_incorrect_or_empty":           {27094, "source为空或输入有误", nil},
	"No_DeviceType_Can_Be_Find":              {27095, "设备类型不存在", nil},

	"TestUser_is_exist":                      {27096, "测试账号已存在", nil},
	"RadacatVersion_Can_Be_Find":             {27097, "该版本已存在", nil},
	"Not_exist":                              {27098, "暂不存在", nil},




}

func GetCodeAndMsg(msg string) reply {
	return CodeMap[msg]
}
