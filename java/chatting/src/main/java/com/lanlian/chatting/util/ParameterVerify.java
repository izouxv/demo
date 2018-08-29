/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 */

package com.lanlian.chatting.util;

import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.util.DataFinals.Gender;

/**
 * @author wdyqxx
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年2月16日 下午2:34:10
 * @explain 此类用于对共用参数的判空与格式校验；
 */

public class ParameterVerify {

	/**
	 * 用户名校验与鉴别；
	 * 
	 * @exception 20001
	 *                用户名为空
	 * @exception 20002
	 *                用户名格式错误
	 * @param parameters
	 * @throws Parameter_Exception
	 */
	public static int username(String username) throws Parameter_Exception {
		if (username == null || username.trim().isEmpty()) {
			throw new Parameter_Exception(20001);
		}
		if (Validator.isMobile(username)) {
			return 1;
		}
		if (Validator.isEmail(username)) {
			return 2;
		}
		throw new Parameter_Exception(20002);
	}

	/**
	 * 用户名判空校验；
	 * 
	 * @exception 20001
	 *                用户名为空
	 * @param parameters
	 * @throws Parameter_Exception
	 */
	public static void usernameNull(String username) throws Parameter_Exception {
		if (username == null || username.trim().isEmpty()) {
			throw new Parameter_Exception(20001);
		}
	}

	/**
	 * 用户名格式校验:1为手机号，2为邮箱号
	 * 
	 * @param username
	 * @throws Parameter_Exception
	 * @exception 20002
	 *                用户名格式错误
	 */
	public static int usernameVerify(String username) throws Parameter_Exception {
		if (Validator.isMobile(username)) {
			return 1;
		}
		if (Validator.isEmail(username)) {
			return 2;
		}
		throw new Parameter_Exception(20002);
	}

	/**
	 * 验证码格式校验
	 * 
	 * @param checkCode
	 * @throws Parameter_Exception
	 * @exception 20011
	 *                验证码为空
	 * @exception 20012
	 *                格式不合法
	 */
	public static void checkCodeVerify(String checkCode) throws Parameter_Exception {
		// 判空
		if (checkCode == null || checkCode.trim().isEmpty()) {
			throw new Parameter_Exception(20011);
		}
		// 校验验证码格式是否合法
		String regCode = "^\\d{6}$";
		if (!checkCode.trim().matches(regCode)) {
			throw new Parameter_Exception(20012);
		}
	}

	/**
	 * 昵称判空与格式校验
	 * 
	 * @param nickname
	 * @throws Parameter_Exception
	 * @throws 20015
	 *             昵称不能为空;20016 昵称格式错误;
	 */
	public static void nicknameVerify(String nickname) throws Parameter_Exception {
		if (nickname == null || nickname.trim().isEmpty()) {
			throw new Parameter_Exception(20015);
		}
		byte[] nicknames = nickname.getBytes();
		if (nicknames.length > 12 || nicknames.length < 1) {
			throw new Parameter_Exception(20016);
		}
	}

	/**
	 * 性别格式判断
	 * 
	 * @param gender
	 * @throws Parameter_Exception
	 *             21002 性别错误
	 */
	public static void genderVerify(int gender) throws Parameter_Exception {
		if (!Gender.contains(gender)) {
			// 参数格式错误
			throw new Parameter_Exception(21002);
		}
	}

	/**
	 * 密码判空校验
	 * 
	 * @param nickname
	 * @throws Parameter_Exception
	 * @throws 20003
	 *             密码不能为空
	 */
	public static void pwdNull(String password) throws Parameter_Exception {
		if (password == null || password.trim().isEmpty()) {
			throw new Parameter_Exception(20003);
		}
	}

	/**
	 * 密码格式校验
	 * 
	 * @param password
	 * @throws Parameter_Exception
	 * @throws 20004
	 *             密码格式错误
	 */
	public static void pwdVerify(String password) throws Parameter_Exception {
		if (!Validator.isPassword(password)) {
			throw new Parameter_Exception(20004);
		}
	}

	/**
	 * 设备型号校验
	 * 
	 * @param device
	 * @throws Parameter_Exception
	 * @throws 21001
	 *             设备型号为空
	 */
	public static void deviceNull(String device) throws Parameter_Exception {
		if (device == null || device.trim().isEmpty()) {
			throw new Parameter_Exception(21001);
		}
	}

	/**
	 * 对参数进行判空校验；
	 * 
	 * @exception 21001,参数不能为空
	 * @param parameters
	 * @throws Parameter_Exception
	 */
	public static void verifyNull(String... parameters) throws Parameter_Exception {
		for (String parameter : parameters) {
			if (parameter == null || parameter.trim().isEmpty()) {
				throw new Parameter_Exception(21001);
			}
		}
	}

	/**
	 * 验证参数是否符合uid范围；
	 * 
	 * @exception 21002,参数格式不合法
	 * @param uids      
	 * @throws Parameter_Exception
	 */
	public static void verifyUid(int... uids) throws Parameter_Exception {
		for (int uid : uids) {
			if ((Math.pow(10, 6) + 1) > uid || uid > Math.pow(10, 8)) {
				throw new Parameter_Exception(21002);
			}
		}
	}

	/**
	 * 验证参数是否符合gid范围
	 * 
	 * @exception 21002,参数格式不合法
	 * @param gid
	 * @throws Parameter_Exception
	 */
	public static void verifyGid(int... gids) throws Parameter_Exception {
		for (int gid : gids) {
			if (Math.pow(10, 4) >= gid || gid >= Math.pow(10, 10)) {
//				continue;
				throw new Parameter_Exception(21002);
			}
		}
	}

	/**
	 * 验证参数是否为Integer类型
	 * 
	 * @param str
	 * @throws 21002,参数格式错误
	 * @return Integer
	 * @throws Parameter_Exception
	 */
	public static int verifyLong(String str) throws Parameter_Exception {
		try {
			int pid = Integer.parseInt(str);
			return pid;
		} catch (NumberFormatException e) {
			throw new Parameter_Exception(21002);
		}
	}

	/**
	 * 验证参数是否为正整数
	 * 
	 * @param str
	 * @throws 21002
	 *             参数格式错误
	 * @return
	 * @throws Parameter_Exception
	 */
	public static void verifyIntegerPositive(Integer... str) throws Parameter_Exception {
		try {
			for (int i : str) {
				if (i < 0) {
					throw new Parameter_Exception(21002);
				}
			}
		} catch (NumberFormatException e) {
			throw new Parameter_Exception(21002);
		}
	}
	
	/**
	 * 对邮箱格式判断
	 * @param email
	 * @throws Parameter_Exception 20025 邮箱格式错误
	 */
	public static void verifyEmail(String email) throws Parameter_Exception {
		if (!Validator.isEmail(email)) {
			throw new Parameter_Exception(20025);
		}
	}
	
	/**
	 * 对手机号格式判断
	 * @param phone
	 * @throws Parameter_Exception 20025 手机号码格式错误
	 */
	public static void verifyPhone(String phone) throws Parameter_Exception {
		if (!Validator.isMobile(phone)) {
			throw new Parameter_Exception(20026);
		}
	}

	/**
	 * 校验经纬度；
	 * 
	 * @exception 21002
	 *                参数格式不合法
	 * @param str
	 * @return boolean
	 * @throws Parameter_Exception
	 */
	public static void verifyCoord(double... strs) throws Parameter_Exception {
		for (double str : strs) {
			if (-180 >= str || 180 <= str) {
				throw new Parameter_Exception(21002);
			}
		}
	}
	
	/**
	 * 设备唯一标识校验
	 * 
	 * @param imei
	 *             设备唯一标识
	 * @throws Parameter_Exception 20005
	 */
	public static void verifyImei(String imei) throws Parameter_Exception {
		// 校验imei的正则验证
		String regImei = "(\\d{15})|([0-9A-Z]{32})";
		if (!imei.trim().matches(regImei) || imei.trim().isEmpty()) {
			throw new Parameter_Exception(20005);// 设备唯一标识格式不合法
		}
	}
	
	/**
	 * 验证token为空
	 * @param token
	 * @throws Parameter_Exception
	 */
	public static void verifyToken(String token) throws Parameter_Exception {
		if (token == null || token.trim().isEmpty()) {
			throw new Parameter_Exception(23019);
		}
	}
}
