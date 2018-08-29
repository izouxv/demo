/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package Demo;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.text.ParseException;

import org.apache.commons.codec.digest.DigestUtils;
import org.apache.commons.io.IOUtils;

/**
 * @Title FileMD5.java
 * @Package cn.lanlian.ccat.util
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月28日 下午4:43:42
 * @explain
 */

public class FileMD5 {

	public static String getMd5ByFile(File file) throws FileNotFoundException {
		String md5 = null;
//		FileInputStream in = new FileInputStream(file);
		FileInputStream in = new FileInputStream(file);
		try {
//			MappedByteBuffer byteBuffer = in.getChannel().map(FileChannel.MapMode.READ_ONLY, 0, file.length());
//			MessageDigest md5 = MessageDigest.getInstance("MD5");
//			md5.update(byteBuffer);
//			BigInteger bi = new BigInteger(1, md5.digest());
//			value = bi.toString(16);
			
			md5 = DigestUtils.md5Hex(IOUtils.toByteArray(in));
			IOUtils.closeQuietly(in);
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			if (null != in) {
				try {
					in.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return md5;
	}

	public static void main(String[] args) throws IOException, ParseException {
		//201703281706_20140226193616296.jpg
		String path = "E:\\21342423423\\1.txt";

		String v = getMd5ByFile(new File(path));
		System.out.println("MD5:" + v);

		FileInputStream fis = new FileInputStream(path);
		
		String md5 = DigestUtils.md5Hex(IOUtils.toByteArray(fis));
		IOUtils.closeQuietly(fis);
		System.out.println("MD5:" + md5);

//		System.out.println("MD5:"+DigestUtils.md5Hex("WANGQIUYUN"));
//		System.out.println(System.currentTimeMillis());

	}

}