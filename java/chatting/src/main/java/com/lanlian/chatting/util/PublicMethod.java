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

import java.io.BufferedOutputStream;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.lang.reflect.Field;
import java.net.URLDecoder;
import java.net.URLEncoder;
import java.sql.Timestamp;
import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletResponse;

import org.apache.commons.codec.digest.DigestUtils;
import org.apache.commons.io.IOUtils;
import org.apache.commons.lang.StringUtils;
import org.apache.log4j.Logger;
import org.springframework.web.multipart.MultipartFile;

import com.fasterxml.jackson.core.JsonParseException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.binarywang.java.emoji.EmojiConverter;
import com.lanlian.chatting.result.Parameter_Exception;

import net.sf.json.JSONArray;

/**
 * @Title PublicMethod.java
 * @Package com.lanlian.chatting.util
 * @author 王东阳
 * @version V1.0
 * @date 2017年5月10日 下午3:08:01
 * @explain 共用方法类；
 */

public class PublicMethod {

	private static Logger log = Logger.getLogger(PublicMethod.class);

	/**
	 * 校验邮箱验证码的合格性
	 * 
	 * @param code
	 * @param aep
	 * @exception 20013,验证码错误
	 * @exception 20014,验证码失效
	 * @return
	 * @throws Parameter_Exception
	 */
	public static void qualified(String checkCode, String code, Timestamp time, long activeTime)
			throws Parameter_Exception {
		if (!code.equals(checkCode)) {
			throw new Parameter_Exception(20013);
		}
		long value = (activeTime - time.getTime()) / (1000 * 60);
		boolean flag = value > 5;
		if (flag) {
			throw new Parameter_Exception(20014);
		}
	}

	/**
	 * 将Map<Long, String>类型数据转为JSONObject的字符串
	 * 
	 * @param map
	 * @return
	 */
	public static JSONArray mapInt_Json(Map<String, Integer> map) {
		JSONArray jsonArray = JSONArray.fromObject(map);
		return jsonArray;

	}

	/**
	 * 将Map<Long, String>类型数据转为JSONObject的字符串
	 * 
	 * @param map
	 * @return
	 */
	public static JSONArray map_Json(Map<String, String> map) {
		JSONArray jsonArray = JSONArray.fromObject(map);
		return jsonArray;

	}

	/**
	 * 将json的字符串数据转换为map数据
	 * 
	 * @param json
	 * @return
	 * @throws Parameter_Exception
	 */
	public static Map<String, String> json_map(String json) throws Parameter_Exception {
		Map<String, String> map = new HashMap<String, String>();
		ObjectMapper mapper = new ObjectMapper();
		try {
			map = mapper.readValue(json, new TypeReference<HashMap<String, String>>() {
			});
		} catch (JsonParseException e) {
			throw new Parameter_Exception(21002);
		} catch (JsonMappingException e) {
			throw new Parameter_Exception(10002);
		} catch (IOException e) {
			throw new Parameter_Exception(10002);
		}
		return map;
	}

	/**
	 * 将emoji表情替换成""
	 * 
	 * @param source
	 * @return 过滤后的字符串
	 */
	public static String filterEmoji(String source) {
		if (StringUtils.isNotBlank(source)) {
			return source.replaceAll("[\\ud800\\udc00-\\udbff\\udfff\\ud800-\\udfff]|[\u2600-\u27ff]", "");
		} else {
			return source;
		}
	}

	public static void main(String[] arg) {
		try {
			// String text = "This is a smiley \uD83C\uDFA6 face\uD860\uDD5D
			// \uD860\uDE07 \uD860\uDEE2 \uD863\uDCCA \uD863\uDCCD \uD863\uDCD2
			// \uD867\uDD98 ";
			// String text = "️,你好";
			// System.out.println(text);
			// System.out.println(text.replaceAll(
			// "[\\ud83c\\udc00-\\ud83c\\udfff]|[\\ud83d\\udc00-\\ud83d\\udfff]|[\\u2600-\\u27ff]",
			// "*"));
			// System.out.println(filterEmoji(text));

			String a = emojiConverterUnicodeStr("这是：\uD860\uDEE2 \uD863\uDCCA \uD863\uDCCD \uD863\uDCD2");
			System.out.println(a);
			String b = emojiConverterToAlias(a);
			System.out.println(b);
		} catch (Exception ex) {
			ex.printStackTrace();
		}
	}

	private static EmojiConverter emojiConverter = EmojiConverter.getInstance();

	/**
	 * 将emojiStr转为 带有表情的字符
	 * 
	 * @param emojiStr
	 * @return
	 */
	public static String emojiConverterUnicodeStr(String emojiStr) {
		String result = emojiConverter.toUnicode(emojiStr);
		return result;
	}

	/**
	 * 带有表情的字符串转换为编码
	 * 
	 * @param str
	 * @return
	 */
	public static String emojiConverterToAlias(String str) {
		String result = emojiConverter.toAlias(str);
		return result;
	}

	/**
	 * 对字符串解码
	 * 
	 * @param str
	 * @return
	 * @throws Parameter_Exception
	 */
	public static String decode(String str) throws Parameter_Exception {
		try {
			return URLDecoder.decode(str, "utf-8");
		} catch (Exception e) {
			throw new Parameter_Exception(10002);
		}
	}

	/**
	 * 对字符串编码
	 * 
	 * @param str
	 * @return
	 * @throws Parameter_Exception
	 */
	public static String encode(String str) throws Parameter_Exception {
		try {
			return URLEncoder.encode(str, "utf-8");
		} catch (Exception e) {
			throw new Parameter_Exception(10002);
		}
	}

	/**
	 * 判断bean属性是否为空
	 * 
	 * @param obj
	 * @return
	 * @throws IllegalAccessException
	 */
	public static boolean checkObjFieldIsNull(Object obj) throws IllegalAccessException {
		boolean flag = false;
		for (Field f : obj.getClass().getDeclaredFields()) {
			f.setAccessible(true);
			if (f.get(obj) == null || f.get(obj).equals("")) {
				log.info("checkObjFieldIsNull f:" + f);
				flag = true;
				return flag;
			}
		}
		return flag;
	}

	/**
	 * 下载文件
	 * 
	 * @param response
	 * @param filname
	 * @param filepath
	 * @param manner
	 * @throws Parameter_Exception
	 */
	public synchronized static void downloadFile(HttpServletResponse response, String filname, String filepath,
			int manner) throws Parameter_Exception {
		InputStream ins = null;
		OutputStream outs = null;
		try {
			response.setCharacterEncoding("UTF-8");
			response.setContentType("multipart/form-data");
			if (1 == manner) {
				response.setHeader("Content-Disposition", "attachment;filename=" + URLEncoder.encode(filname, "utf-8"));
			}
			ins = new FileInputStream(new File(filepath));
			response.setHeader("Content-Length", String.valueOf(ins.available()));
			byte[] buffer = new byte[ins.available()];
			ins.read(buffer);
			ins.close();
			outs = new BufferedOutputStream(response.getOutputStream());
			outs.write(buffer);
		} catch (FileNotFoundException e) {
			log.error("downloadFile-err:" + e);
			throw new Parameter_Exception(21002);
		} catch (Exception e) {
			log.error("downloadFile-err:" + e);
			throw new Parameter_Exception(10002);
		} finally {
			// 这里主要关闭
			try {
				ins.close();
				outs.close();
				outs.flush();
			} catch (Exception e) {
				log.error("downloadFile-err:" + e);
			}
		}
	}

	/**
	 * 计算文件md5
	 * 
	 * @param file
	 * @return
	 * @throws FileNotFoundException
	 * @throws Parameter_Exception
	 */
	public static String getMd5ByFile(File file) throws FileNotFoundException, Parameter_Exception {
		String md5 = null;
		FileInputStream in = new FileInputStream(file);
		try {
			md5 = DigestUtils.md5Hex(IOUtils.toByteArray(in));
			IOUtils.closeQuietly(in);
			return md5;
		} catch (Exception e) {
			throw new Parameter_Exception(10002);
		} finally {
			if (in != null) {
				try {
					in.close();
				} catch (IOException e) {
					log.error(e);
				}
			}
		}
	}

	/**
	 * 将文件存入磁盘
	 * 
	 * @param name
	 * @param topic
	 * @param file
	 * @param time
	 * @return
	 * @throws Parameter_Exception
	 */
	public synchronized static String saveFile(String savePath, String name, String md5, MultipartFile file,
			String time) throws Parameter_Exception {
		// 得到上传的文件名
		String filename = file.getOriginalFilename();
		System.out.println("savePath:" + savePath + "," + time + "," + filename);
		File tmpFile = new File(savePath);
		// 判断上传文件的保存目录是否存在
		if (!tmpFile.exists()) {
			// System.out.println(savePath + "目录不存在，创建目录-----");
			tmpFile.mkdirs();
		}
		// 保存完整路径
		String path = savePath + time + "_" + filename;
		// 把文件上传至path的路径
		File localFile = new File(path);
		String md5s = null;
		try {
			file.transferTo(localFile);
			md5s = getMd5ByFile(localFile);
			if (!md5s.equals(md5)) {
				throw new Parameter_Exception(21002);
			}
		} catch (IllegalStateException | IOException e) {
			throw new Parameter_Exception(10002);
		}
		log.info("uploadAdvertisement-localFile:" + localFile);
		return path;
	}

	/**
	 * 将日志追加写入文件
	 * @param filePath
	 * @param context
	 * @throws Exception 
	 */
	public synchronized static void printWriterFile(String filePath, String context) throws Exception {
		BufferedWriter out = null;
		try {
			File f = new File(filePath);
            if (!f.exists()) {
            	f.createNewFile();
            }
			out = new BufferedWriter(new FileWriter(f, true));
			out.write(context);
			out.write("\r\n");
		} catch (Exception e) {
			log.error(e);
			throw e;
		} finally {
			try {
				out.close();
			} catch (IOException e) {
				throw e;
			}
		}
	}

}
