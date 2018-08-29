/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/** 
 * @Title AbstractController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1
 * @date 2017年5月18日 下午7:12:57
 * @explain 
 */
package com.lanlian.chatting.controller;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.net.URLEncoder;
import java.sql.Timestamp;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;
import org.mybatis.spring.MyBatisSystemException;
import org.springframework.web.HttpRequestMethodNotSupportedException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.mvc.AbstractController;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Fatal_Exception;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.util.InviteCode;
import com.mysql.jdbc.exceptions.MySQLSyntaxErrorException;

import io.grpc.StatusRuntimeException;
import redis.clients.jedis.exceptions.JedisConnectionException;
import redis.clients.jedis.exceptions.JedisException;

/**
 * @Title AbstractController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @param Throwable
 *            e
 * @date 2017年3月06日 上午06:06:06
 * @explain 对子类的统一处理
 */

@ControllerAdvice
public class MyAbstractController extends AbstractController {

	private static Logger logger = Logger.getLogger(MyAbstractController.class);

	/**
	 * 当前时间
	 */
	private Timestamp time = new Timestamp(0);

	public Timestamp getTime() {
		time.setTime(System.currentTimeMillis());
		return time;
	}

	public Long getTimes() {
		return System.currentTimeMillis();
	}

	public void setTime(Long now) {
		time.setTime(now);
	}

	@Override
	protected ModelAndView handleRequestInternal(HttpServletRequest request, HttpServletResponse response)
			throws Exception {
		if (isFormSubmission(request)) {
		}
		return null;
	}

	private Boolean isFormSubmission(HttpServletRequest request) {
		if ("POST".equals(request.getMethod())) {
			return true;
		}
		return false;
	}

	/**
	 * 请求身份信息
	 */
	public SsoPo getContextSsoPo(HttpServletRequest request, HttpServletResponse response) {
		return (SsoPo) request.getAttribute("SsoPo");
	}

	/**
	 * 请求代理信息
	 */
	public String getContextIP(HttpServletRequest request, HttpServletResponse response) {
		return (String) request.getAttribute("ip");
	}

	/**
	 * 请求代理信息
	 */
	public String getContextDev(HttpServletRequest request, HttpServletResponse response) {
		return (String) request.getAttribute("dev");
	}

	/**
	 * 进行url编码
	 * 
	 * @param value
	 * @return
	 * @throws UnsupportedEncodingException
	 */
	public String encode(String value) throws UnsupportedEncodingException {
		return URLEncoder.encode(value, "UTF-8");
	}

	/**
	 * 进行url解码
	 * 
	 * @param value
	 * @return
	 * @throws UnsupportedEncodingException
	 */
	public String decode(String value) throws UnsupportedEncodingException {
		return URLDecoder.decode(value, "UTF-8");
	}

	/**
	 * 字符数组
	 */
	public static String[] chars = new String[] { "a", "b", "c", "d", "e", "f", "g", "h", "j", "k", "m", "n", "o", "p",
			"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C",
			"D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z" };

	/**
	 * 返回id
	 * 
	 * @return
	 */
	public int getId(String code) {
		return InviteCode.codeToId(code);
	}

	/**
	 * 返回随机码
	 * 
	 * @return
	 */
	public String getRandomCode(Integer id) {
		return InviteCode.getRandomCode(id);
	}
	// public String getRandomCode() {
	// StringBuffer shortBuffer = new StringBuffer();
	// String uuid = UUID.randomUUID().toString().replace("-", "");
	// for (int i = 0; i < 8; i++) {
	// String str = uuid.substring(i * 4, i * 4 + 4);
	// int x = Integer.parseInt(str, 16);
	// shortBuffer.append(chars[x % 0x3E]);
	// }
	// return shortBuffer.toString();
	// }

	/**
	 * 生成随机盐
	 * 
	 * @return
	 */
	public synchronized static String getSalt() {
		String code = "";
		for (int i = 5; i >= 0; i--) {
			int a = (int) ((Math.random() * chars.length));
			code = code + chars[a];
		}
		return code;
	}

	/***********************************
	 * 异常捕捉
	 ********************************************/

	/**
	 * 捕捉异常与返回状态码
	 * 
	 * @param e
	 * @return
	 */
	@ResponseBody
	@ExceptionHandler(Throwable.class)
	private String verifyfo(Throwable e) {
		// 请求的json格式错误或json中缺少属性
		if (e.getMessage().contains("Could not read JSON")) {
			return JSON.toJson(21002);
		}
		// request请求需要的参数格式错误
		if (e.getMessage().contains("Failed to convert value of type")) {
			return JSON.toJson(21002);
		}
		// 缺少request请求需要的参数
		if (e.getMessage().contains("Required")) {
			return JSON.toJson(21001);
		}
		if (e.getMessage().contains("Failed to invoke handler method")) {
			return JSON.toJson(20000);
		}
		logger.fatal(e);
		return JSON.toJson(10002);
	}

	@ResponseBody
	@ExceptionHandler(Parameter_Exception.class)
	private String verify(Parameter_Exception e) {
		logger.info("verify-Parameter_Exception:" + e);
		return JSON.toJson(Integer.parseInt(e.getMessage()));
	}
	
	@ResponseBody
	@ExceptionHandler(HttpRequestMethodNotSupportedException.class)
	private String verify(HttpRequestMethodNotSupportedException e) {
		logger.info("HttpRequestMethodNotSupportedException:" + e);
		return JSON.toJson(21002);
	}

	@ResponseBody
	@ExceptionHandler(StatusRuntimeException.class)
	private String verify(StatusRuntimeException e) {
		logger.fatal("Grpc连接错误---StatusRuntimeException:" + e.getMessage());
		return JSON.toJson(10002);
	}

	@ResponseBody
	@ExceptionHandler(JedisConnectionException.class)
	private String verify(JedisConnectionException e) {
		logger.fatal("Redis连接错误---JedisConnectionException:" + e.getMessage());
		return JSON.toJson(10002);
	}
	
	@ResponseBody
	@ExceptionHandler(JedisException.class)
	private String verify(JedisException e) {
		logger.fatal("Redis连接错误---JedisException:" + e.getMessage());
		return JSON.toJson(10002);
	}

	@ResponseBody
	@ExceptionHandler(MyBatisSystemException.class)
	private String verify(MyBatisSystemException e) {
		logger.fatal("MySql连接错误---MyBatisSystemException:" + e.getMessage());
		return JSON.toJson(10002);
	}
	
	@ResponseBody
	@ExceptionHandler(MySQLSyntaxErrorException.class)
	private String verify(MySQLSyntaxErrorException e) {
		logger.fatal("MySql连接错误---MySQLSyntaxErrorException:" + e.getMessage());
		return JSON.toJson(10002);
	}
	
	@ResponseBody
	@ExceptionHandler(Fatal_Exception.class)
	private String verify(Fatal_Exception e) {
		logger.fatal("HTTP调用错误---Fatal_Exception:" + e.getMessage());
		return JSON.toJson(10002);
	}

}
