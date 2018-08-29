/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/**
 * 
 */
package com.lanlian.chatting.util;

/** 
 * @Title RequestSetting.java
 * @Package com.lanlian.chatting.util
 * @author 王东阳
 * @version v1.0
 * @date 2017年4月7日 下午7:12:16
 * @explain chatting项目的Request常量、URI路径；
 */

public class RequestSetting {
	
	/**
	 * HTTP请求的提交内容类型:consumes
	 */
	public static final String CONSUMES = "application/x-www-form-urlencoded;charset=UTF-8";
	
	/**
	 * HTTP请求的返回内容类型:produces
	 */
	public static final String PRODUCES = "application/json;charset=UTF-8";
	
	
	/**
	 * 项目版本号
	 */
	public static final String VERSION = "/v1.0";
	
	public static final String VERSION_1_1 = "/v1.1";
	
	
	/**===================   消息通知服务      =============================*/
	
	/**
	 * NoticeController的类路径
	 */
	public static final String NOTICE_PARENT = VERSION + "/notice";
	/**
	 * NoticeController的hello方法路径
	 */
	public static final String NOTICE_BODY_NOTICE = "/dada_notice";
	
	
	/**======================   游客服务     =============================*/
	
	/**
	 * 游客父路径
	 */
	public static final String VISITORS_PARENT = VERSION + "/visitors";
	/**
	 * 游客父路径
	 */
	public static final String VISITORS_PARENT_1_1 = VERSION_1_1 + "/visitors";
	/**
	 * CodeSendController的active方法路径
	 */
	public static final String CODE_BODY_1 = "/send_code";
	/**
	 * CodeSendController的active方法路径
	 */
	public static final String CODE_BODY = "/send_code/{action}";
	/**
	 * CodeVerifyController的verifyCode方法路径
	 */
	public static final String VERIFY_BODY = "/verify_code";
	/**
	 * RegisterController的register方法路径
	 */
	public static final String REGISTER_BODY = "/register";
	/**
	 * ResetController的resetPwd方法路径
	 */
	public static final String RESET_BODY = "/reset_pwd";
	
	/******************* 游客与用户的虚拟信道服务   *********************/
	
	/**
	 * NearbyController的localtion方法路径
	 */
	public static final String NEARBY_BODY = "/nearby/{types}/{imei}";
	/**
	 * NearbyController的nearbyLocaltion方法路径
	 */
	public static final String NEARBY_LOCALTION_BODY = "/nearby/{types}";
	/**
	 * NearbyController的sendNearby方法路径
	 */
	public static final String NEARBY_SEND_BODY = "/{types}/{tuid}/send/{toid}";
	/**
	 * NearbyController的getNearby方法路径
	 */
	public static final String NEARBY_GET_BODY = "/pop/{tuid}";
	
	
	/**======================   用户服务     =============================*/
	/**
	 * user的请求，类路径
	 */
	public static final String USER_PARENT =  VERSION + "/user";
	/**
	 * user的请求，类路径
	 */
	public static final String SESSIONS =  VERSION + "/sessions";
	/**
	 * user的checkUser方法路径
	 */
	public static final String USER_BODY_CHECK = "/check";
	/**
	 * user的login方法路径
	 */
	public static final String USER_BODY_LOGIN = "/login";
	/**
	 * user的exit方法路径
	 */
	public static final String USER_BODY_EXIT = "/exit";
	/**
	 * user的basicInfo方法路径
	 */
	public static final String USER_BODY_INFO = "/basic_info";
	/**
	 * user的changePwd方法路径
	 */
	public static final String USER_BODY_CHANGE = "/pwd/{action:[a-z]{4,6}}";
	/**
	 * user的realName方法路径
	 */
	public static final String USER_BODY_REALNAME = "/real_name";
	
	
	/**====================   联系人服务        =============================*/
	
	/**
	 * FriendController的类路径
	 */
	public static final String FRIENDS_PARENT = VERSION + "/friend";
	/**
	 * FriendController的acquireFriends方法路径
	 */
	public static final String FRIENDS_BODY_ACQUIRE = "/acquire";
	/**
	 * FriendController的saveFriends方法路径
	 */
	public static final String FRIENDS_BODY_SAVE = "/sync";
	/**
	 * FriendController的deleteFriends方法路径
	 */
	public static final String FRIENDS_BODY_DELETE = "/delete";
	/**
	 * FriendController的updateFriends方法路径
	 */
	public static final String FRIENDS_BODY_UPDATE = "/update";
	
	
	/**=====================  私信服务       =============================*/
	
	/**
	 * UserLetterController的类路径
	 */
	public static final String LETTER_PARENT = VERSION + "/letter";
	/**
	 * UserLetterController的sendPrivateLette方法路径
	 */
	public static final String LETTER_BODY_SEND = "/send/{touid}";
	/**
	 * UserLetterController的getLetter方法路径
	 */
	public static final String LETTER_BODY_GET = "/get/{touid}";
	
	
	/**==================   文件服务         ==============================*/
	
	/**
	 * VersionController的类路径
	 */
	public static final String FILE_PARENT_V = VERSION + "/file";
	/**
	 * VersionController2的类路径
	 */
	public static final String FILE_PARENT_V_1_1 = VERSION + "/version";
	/**
	 * FileController的类路径
	 */
	public static final String FILE_PARENT_F = VERSION + "/files";
	/**
	 * VersionController的uploadFiles方法路径
	 */
	public static final String FILE_BODY_UPLOAD_VERSION = "/put/{device}";
	/**
	 * VersionController的download方法路径
	 */
	public static final String FILE_BODY_DOWNLOAD_VERSION = "/{action:[a-z]{1,5}}/{device}";
	/**
	 * VersionController1.1的newVersion方法路径
	 */
	public static final String FILE_BODY_NEW_VERSION = "/{dev}/{code}";
	/**
	 * FilesController的uploadFile方法路径
	 */
	public static final String FILE_BODY_UPLOAD = "/upload/feedback";
	/**
	 * FilesController的download方法路径
	 */
	public static final String FILE_BODY_DOWNLOAD = "/download/{id:[a-z0-9]{1,24}}";
	/**
	 * FilesController的uploadFileLog方法路径
	 */
	public static final String FILE_BODY_UPLOAD_LOG = "/upload_log/{source}";
	/**
	 * FilesController的feedback方法路径
	 */
	public static final String FILE_BODY_FEEDBACK = "/feedback";
	
	/**=====================    广告接口    ===========================*/
	/**
	 * FilesController的uploadFile方法路径
	 */
	public static final String FILE_BODY_ADVERTISEMENT = "/upload/advertisement/{name}";
	/**
	 * FilesController的getAdvertisement方法路径
	 */
	public static final String FILE_BODY_GET_ADVERTISEMENT = "/advertisement/get";
	/**
	 * FilesController的getAdver方法路径
	 */
	public static final String FILE_BODY_GET_ADVER = "/adver";
	/**
	 * FilesController的downAdvertisement方法路径
	 */
	public static final String FILE_BODY_DOWNLOAD_ADVERTISEMENT = "/advertisement/get/{file}";
	
	/**=======================    实时群信息服务        =======================*/
	
	/**
	 * DadaGroupController的dadaGroup方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_SYNC_GID = "/sync/gid/{gid}";
	
	/**
	 * DadaGroupMsgController的dadaGroupMsg方法路径
	 */
	public static final String GROUP_MESSAGE_BODY_SYNC_GID = "/info/gid/{gid}";
	
	
	/**=======================    群资料服务        =======================*/
	
	/**
	 * GroupSettingsController的类路径
	 */
	public static final String GROUP_SETTINGS_PARENT = VERSION + "/group_settings";
	public static final String GROUP_SETTINGS_PARENT_1_1 = VERSION_1_1 + "/group_settings";
	/**
	 * GroupSettingsController的syncGroup方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_SYNC = "/sync";
	/**
	 * GroupSettingsController的deleteGroup方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_DISSOLVE = "/dissolve";
	/**
	 * GroupSettingsController的groupEdit方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_EDIT = "/edit";
	/**
	 * GroupSettingsController的kickMember方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_KICKMEMBER = "/kick_member";
	/**
	 * GroupSettingsController的exitGroup方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_EXIT = "/quit";
	/**
	 * GroupSettingsController的findGroupByUidList方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_LIST = "/list";
	/**
	 * GroupSettingsController的findGroupByGidList方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_INFO = "/group_info";
	/**
	 * GroupSettingsController的temporaryGroup方法路径
	 */
	public static final String GROUP_SETTINGS_BODY_TEMPORARY = "/temporary/{types}";
	
	
	
	/**========================   群信息服务      ===========================*/
	
	/**
	 * GroupMessageController的类路径
	 */
	public static final String GROUP_MESSAGE_PARENT = VERSION + "/group_message";
	public static final String GROUP_MESSAGE_PARENT_1_1 = VERSION_1_1 + "/group_message";
	/**
	 * GroupMessageController的messageSync方法路径
	 */
	public static final String GROUP_MESSAGE_BODY_SYNC = "/sync";
	/**
	 * GroupMessageController的messageList方法路径
	 */
	public static final String GROUP_MESSAGE_BODY_LIST = "/list";
	/**
	 * GroupMessageController的messageInfo方法路径
	 */
	public static final String GROUP_MESSAGE_BODY_INFO = "/{id}/info";
	
	/*********************  统计信息上报        ************************/
	/**
	 * localtionDataController的localtion方法路径
	 */
	public static final String LOCALTIONDATA_BODY = "/statistics/{types}";
	public static final String STATISTICS	 = VERSION_1_1 + "/statistics";
	public static final String DEVINFO_BODY = "/devinfo";
	public static final String HEARTBEAT_BODY	 = "/heartbeat";
	
	
	/**
	 * 微信调用接口
	 */
	public static final String GROUP_WECHAT_PARENT = VERSION + "/groups";
	public static final String GROUP_WECHAT_PARENT_1_1 = VERSION_1_1 + "/groups";
	public static final String GROUP_WECHAT_BODY_SUB = "/group_settings";
	public static final String GROUP_WECHAT_BODY_UNSUB = "/unsub/{gid}";
	public static final String GROUP_WECHAT_BODY_MESSAGE = "/group_message";
	
	
	
	/** 暂停接口 start============================== */
	public static final String GROUP_SETTINGS_BODY_add_member = "/add_member";
	public static final String GROUP_SETTINGS_BODY_userInformationQuery = "/query";
	public static final String GROUP_SETTINGS_BODY_getGroupUsers = "/groupusers";
	/**暂停接口 end==============================*/
	


}