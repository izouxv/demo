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
 * @Title SMSparameter.java
 * @Package com.lanlian.chatting.util
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月6日 下午3:59:19
 * @explain 
 */

public class Mailparameter {
	
	/**
	 * 阿里大于发送短信：URL接口
	 */
	public static final String URL = "http://gw.api.taobao.com/router/rest";
	
	/**
	 * 阿里大于发送短信：App Key
	 */
	public static final String APPKEY = "23783317";
	
	/**
	 * 阿里大于发送短信：App Secret
	 */
	public static final String SECRET = "693e543b4f66bef183678a1eb69843c0";
	
	/**
	 * 阿里大于发送短信：短信类型
	 */
	public static final String SMS_TYPE = "normal";
	
	/**
	 * 阿里大于发送短信：短信签名
	 */
	public static final String NAME= "胡虹";
	
	/**
	 * 阿里大于发送短信：短信模板ID
	 */
	public static final String CODE = "SMS_63860763";
	
	/**
	 * 阿里大于发送短信：短信模板变量
	 */
	public static String sms_param(String username, String auth_code) {
		final String sms_param = "{\"name\":" + "\"" + username + "\"," + "\"authcode\":" + "\"" + auth_code + "\"" + "}";
		return sms_param;
		
	}
	
	
}

