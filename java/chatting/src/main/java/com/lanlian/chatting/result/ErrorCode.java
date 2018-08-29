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
package com.lanlian.chatting.result;

import java.util.HashMap;
import java.util.Map;

/** 
 * @Title ErrorCode.java
 * @Package com.lanlian.chatting.result
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月8日 下午5:15:38
 * @explain 状态码与提示信息；
 */

public class ErrorCode {
	
	public static Map<Integer, String> map = new HashMap<Integer, String>();
	
	//读取配置文件信息
//	private static final ResourceBundle bundle = ResourceBundle.getBundle("codemsg");
	
	/**
	 * 状态码与提示放入map
	 */
	static {
		
		for (Result result : Result.values()) {
			int code = result.code;
			String msg = result.msg;
			map.put(code, msg);
		}
	}
	
	public enum Result {
		
		//成功
		C10000(10000, "成功(OK)"),
		C20000(20000, "网络异常，请稍后再试"),// Error（错误）
		
		//服务端错误
		C10001(10001, "网络繁忙"),
		C10002(10002, "网络繁忙"),
		C10003(10003, "系统接口维护"),
		C10004(10004, "短信运营商繁忙"),
		C10005(10005, "邮件服务器繁忙"),
		C32001(32001, "网络繁忙"),
		C32003(32003, "网络繁忙"),
		C33001(33001, "网络繁忙"),
		C33003(33003, "网络繁忙"),
		C33006(33006, "网络繁忙"),

		//用户输入错误
		C20001(20001, "请填写用户名"),
		C20002(20002, "手机号或邮箱格式错误"),
		C20003(20003, "请填写密码"),
		C20004(20004, "密码格式错误（6-16位的数字或字母）"),
		C20005(20005, "用户名或密码错误"),
		C20006(20006, "该用户信用太低，禁止登陆"),
		C20007(20007, "该用户状态异常，禁止登陆"),
		C20008(20008, "该用户名已被注册"),
		C20009(20009, "该用户名不存在"),
		C20010(20010, "短信发送次数达到当日上限"),
		C20011(20011, "请填写验证码"),
		C20012(20012, "验证码格式错误"),
		C20013(20013, "验证码错误"),
		C20014(20014, "验证码失效"),
		C20015(20015, "请填写昵称"),
		C20016(20016, "请填写正确的昵称"),
		C20017(20017, "姓名不能为空"),
		C20018(20018, "请填写正确的姓名"),
		C20019(20019, "身份证号码不能为空"),
		C20020(20020, "身份证号码格式错误"),
		C20021(20021, "姓名与身份证号码不符"),
		C20022(20022, "该账户已被实名认证"),
		C20023(20023, "无数据，需要先备份"),
		C20024(20024, "密码错误"),
		C20025(20025, "请填写正确邮箱"),
		C20026(20026, "请填写正确用户名"),
		C20027(20027, "请先登陆后，再操作"),
		C20028(20028, "邀请码错误"),
		C20029(20029, "已绑定该邀请码"),
		C20030(20030, "该用户无操作权限"),
		C23019(23019, "token为空"),
		C32002(32002, "用户名错误或密码错误"),
		C33002(33002, "未查到该用户信息"),
		C33004(33004, "验证码错误"),
		C33005(33005, "无该token"),
		C33007(33007, "账号未激活"),
		C33008(33008, "该用户名已存在"),
		C33009(33009, "密码错误"),
		C33010(33010, "source为空或有误"),

		//客户端错误
		C21001(21001, "参数不能为空"),
		C21002(21002, "参数格式不合法"),
		C21003(21003, "记录被封禁"),
		C21004(21004, "用户被警告"),
		C21005(21005, "非法操作或没有权限"),
		C21006(21006, "文件不能为空"),
		C21007(21007, "文件ID错误"),
		C21008(21008, "无资源");
		
		private int code;
 		private String msg;
		
		private Result(int code,String msg) {
			this.code = code;
			this.msg = msg;
		}
		
		public int getCode() {
			return this.code;
		}
		
		public String getMsg() {
			return this.msg;
		}
		
		@Override
		public String toString() {
			return this.code + ":" + this.msg;
		}
		
	}
	
	
	
}




