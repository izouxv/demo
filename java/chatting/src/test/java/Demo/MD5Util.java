package Demo;

import java.security.MessageDigest;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class MD5Util {
	public final static String MD5(String s) {
		char hexDigits[] = { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F' };
		try {
			byte[] btInput = s.getBytes();
			for (byte b : btInput) {
				System.out.print("b:" + b + ",");
			}
			// 获得MD5摘要算法的 MessageDigest 对象
			MessageDigest mdInst = MessageDigest.getInstance("MD5");
			System.out.println("mdInst:" + mdInst);
			// 使用指定的字节更新摘要
			mdInst.update(btInput);
			System.out.println("mdInst:" + mdInst);
			// 获得密文
			byte[] md = mdInst.digest();
			for (byte b : md) {
				System.out.println("b:" + b);
			}
			// 把密文转换成十六进制的字符串形式
			int j = md.length;
			char str[] = new char[j * 2];
			int k = 0;
			for (int i = 0; i < j; i++) {
				byte byte0 = md[i];
				str[k++] = hexDigits[byte0 >>> 4 & 0xf];
				str[k++] = hexDigits[byte0 & 0xf];
			}
			return new String(str);
		} catch (Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	public static void main(String[] args) throws ParseException {
		Date date = new Date(0L);
		SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss sss");
		Date da = sdf.parse("1000-00-00 00:00:00 001");
		String time = sdf.format(date);
		
		System.out.println(da.getTime() + ",date:"+time);
	}

}