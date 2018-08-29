package util

const Successfull int32 = 10000
const System_error int32 = 10001

const (
	Input_parameter_error        int32 = 37001 + iota //1
	Token_inavailable                                 //5
	AppId_is_not_exist                                //11
	Have_node_in_app                                  //13
	AppKey_is_err                                     //15
	Node_is_not_have                                  //17
	Delete_is_not_have                                //18
	DevEUI_is_exist                                   //输入的devEUI已存在		  26
	AppEUI_is_exist                                   //输入的appEUI已存在		  27
	Alert_is_not_have                                 //28
	Upgrade_Not_Have                                  //34
	NO_Advertisement_Can_Be_Find                      //35
	Advertisement_not_exist                           //36
	User_Does_Not_Exist
	User_Does_Already_Exist
	Authentication_Fail
	Not_Permission
	User_Does_Not_Login
	User_Was_Kicked_Out
	Domain_Does_Not_Exist
	Domain_Can_Not_Delete    //无法删除域，其下有子域或应用或设备
	Advertisement_Not_Update //广告暂未更新
	NwsKey_is_err
	AppsKey_is_err
	DevAddr_is_err
	DevEUI_is_err
	Domain_Expand_MaxPath
	Node_Adv_Does_Not_Exist //
	User_Does_Not_Delete
	User_Does_Not_Update
	NO_RadacatVersion_Can_Be_Find
	No_Feedback_Can_Be_Find
	No_Adv_Can_Be_Find   // 37033
	DeviceType_Does_Not_Exist
	TestUser_is_exist
	TestUser_not_exist
	Version_is_exist
)
