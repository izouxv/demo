/** 
* <p>开发公司：		蓝涟科技 <p>
* <p>版权所有：		蓝涟科技 <p>
* <p>责任人：		      王东阳<p> 
* <p>网址：www.radacat.com <p>
*/
package com.lanlian.chatting.util;

import java.io.UnsupportedEncodingException;

import org.apache.commons.codec.binary.Base64;

import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @author  wdyqxx
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年2月10日 上午11:43:00
 * @explain 
 */
public class Base64s {
	
	/**
	 * BASE64解密
	 * @param key
	 * @return byte[]
	 * @throws Parameter_Exception 
	 * @throws Exception
	 */
	public static String decryptBASE64(String key) throws Parameter_Exception {
		final byte[] input = key.getBytes();
		byte[] output = Base64.decodeBase64(input);
		String base64;
		try {
			base64 = new String(output,"utf-8");
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(10001);
		}
		return base64;
	}

	/**
	 * BASE64加密
	 * @param key
	 * @return String
	 * @throws Parameter_Exception 
	 * @throws Exception
	 */
	public static String encryptBASE64(String key) throws Parameter_Exception {
		final byte[] input = key.getBytes();
		byte[] output = Base64.encodeBase64(input);
		String base64;
		try {
			base64 = new String(output,"utf-8");
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(10001);
		}
		return base64;
	}
	
	
	public static void main(String[] args) {
		try {
//			String aaa= "/var/log/chatting/advertisement/test/2017_11_16_18_44_aaa.jpg";
//			String str = PublicMethod.encode(aaa);
//			System.out.println(str);
//			String sad=encryptBASE64(aaa);
//			System.out.println(sad);
			String asd=decryptBASE64("L3Zhci9sb2cvY2hhdHRpbmcvYWR2ZXJ0aXNlbWVudC/ohb7orq8vMjAxN18xMV8xN18xOV80N19kZGQuanBn");
			System.out.println(new String(asd));
		} catch (Exception e) {
			e.printStackTrace();
		}
	}


	

}
