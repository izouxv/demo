/** 
* <p>开发公司：		蓝涟科技 <p>
* <p>版权所有：		蓝涟科技 <p>
* <p>责任人：		      王东阳<p> 
* <p>网址：www.radacat.com <p>
*/
package com.lanlian.chatting.util;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Random;

import org.apache.commons.codec.binary.Hex;

import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @author wdyqxx
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年2月10日 下午1:41:41
 * @explain 此类用于对密码加密的；
 */
public class EncryptPwd {

	/**
	 * 校验密码是否正确
	 * @throws Parameter_Exception 
	 */
	public static boolean verify(String password, String md5) throws Parameter_Exception {
		char[] cs1 = new char[32];
		char[] cs2 = new char[16];
		for (int i = 0; i < 48; i += 3) {
			cs1[i / 3 * 2] = md5.charAt(i);
			cs1[i / 3 * 2 + 1] = md5.charAt(i + 2);
			cs2[i / 3] = md5.charAt(i + 1);
		}
		return getMD5Hex(password + new String(cs2)).equals(new String(cs1));
	}

	/**
	 * 生成含有随机盐的密码
	 * @throws Parameter_Exception 
	 */
	public static String encrypt(String password) throws Parameter_Exception {
		Random random = new Random();
		StringBuilder sb = new StringBuilder(16);
		sb.append(random.nextInt(99999999));
		int len = sb.length();
		System.out.println("sb:" + sb);
		if (len < 16) {
			for (int i = 0; i < 16 - len; i++) {
				sb.append(i);
			}
		}
		password = getMD5Hex(password + sb.toString());
		System.out.println("MD5/password:" + sb.toString() + ",ps:" + password);
		char[] cs = new char[48];
		for (int i = 0; i < 48; i += 3) {
			cs[i] = password.charAt(i / 3 * 2);
			cs[i + 1] = sb.toString().charAt(i / 3);
			cs[i + 2] = password.charAt(i / 3 * 2 + 1);
		}
		return new String(cs);
	}

	/**
	 * 获取十六进制字符串形式的MD5摘要
	 * @throws Parameter_Exception 
	 * 
	 * @throws Exception_md5 md5加密失败
	 */
	public synchronized static String getMD5Hex(String src) throws Parameter_Exception {
		MessageDigest md5 = null;
		try {
			md5 = MessageDigest.getInstance("MD5");
		} catch (NoSuchAlgorithmException e) {
			throw new Parameter_Exception(10002);
		}
		byte[] bs = md5.digest(src.getBytes());
		return new String(new Hex().encode(bs));
	}

	public static void main(String[] args) throws Parameter_Exception {
		String pwd = getMD5Hex("123456");
		System.out.println("pwd:" + pwd);
		String password = encrypt(pwd);
		System.out.println("password:" + password);
		System.out.println(verify("e4949692e039a5121675c94ea0091072483be4d05de6077f", password));

		/*
		 * e10adc3949ba59abbe56e057f20f883e
		 * b7a59741078d809e85829a7a80e419e2fd39841b5906f97f
		 * 59f63dd38e06c7088912cd1be0551352a935142559860b77
		 * 
		 */
	}

}
