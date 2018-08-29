package com.lanlian.chatting.util;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

import org.apache.commons.codec.binary.Hex;

public class CheckApp {

	

	// 计算并获取CheckSum
	public static String getSignature(String appSecret, String nonce, String curTime) throws Exception {
		String appKey = Base64s.decryptBASE64(appSecret);
		return encode("sha1", appKey + nonce + curTime);
	}

	private static String encode(String algorithm, String value) {
		try {
			if (value == null) {
				return null;
			}
			MessageDigest messageDigest = MessageDigest.getInstance(algorithm);
			messageDigest.update(value.getBytes());
			return getFormattedText(messageDigest.digest());
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}

	/**
	 * 获取十六进制字符串形式的MD5摘要
	 * @throws NoSuchAlgorithmException
	 */
	public synchronized static String getMD5Hex(String src) throws NoSuchAlgorithmException {
		MessageDigest md5 = MessageDigest.getInstance("MD5");
		byte[] bs = md5.digest(src.getBytes());
		return new String(new Hex().encode(bs));
	}

	private static String getFormattedText(byte[] bytes) {
		int len = bytes.length;
		StringBuilder buf = new StringBuilder(len * 2);
		for (int j = 0; j < len; j++) {
			buf.append(HEX_DIGITS[(bytes[j] >> 4) & 0x0f]);
			buf.append(HEX_DIGITS[bytes[j] & 0x0f]);
		}
		return buf.toString();
	}

	private static final char[] HEX_DIGITS = { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd',
			'e', 'f' };
	
	
	public static void main(String[] args) throws Throwable {
		long pid=111111;
		String appSecret =getMD5Hex(String.valueOf(pid));
		long time=System.currentTimeMillis()/1000;
		String nonce = getMD5Hex(time%1000+""+pid);
		System.out.println("nonce:"+nonce);
		String timestamp1 = String.valueOf(time);
		System.out.println("getSignatunre:" + getSignature(appSecret, nonce, timestamp1));
		
		System.out.println("getMD5:" + getMD5Hex(appSecret + nonce + timestamp1));
		// poi cxf 
	}
	
}