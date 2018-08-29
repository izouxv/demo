/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.logger;

import java.io.PrintWriter;
import java.io.StringWriter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * @Title Logger.java
 * @Package com.lanlian.chatting.logger
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月12日 下午3:59:58
 * @explain 输出日志类；
 */

public class LogUtil {

	/**
	 * 使用slf4j创建日志对象
	 */
	private static Logger logger = LoggerFactory.getLogger(LogUtil.class);
	
	private static String getTrace(Throwable e) {
		StringWriter stringWriter= new StringWriter();
		PrintWriter writer= new PrintWriter(stringWriter);
		e.printStackTrace(writer);
		StringBuffer buffer= stringWriter.getBuffer();
		return buffer.toString(); 
	}

	/**
	 * 输出错误级别的日志
	 * 
	 * @param message
	 */
	public static void error(Throwable e) {
		logger.error(getTrace(e));
	}
	
	/**
	 * 输出错误级别的日志
	 * 
	 * @param message
	 */
	public static void error(String str) {
		logger.error(str);
	}

	/**
	 * 输出警告级别的日志
	 * 
	 * @param message
	 */
	public static void warn(Throwable e) {
		logger.warn(e.getMessage());
	}

	/**
	 * 输出debug级别的日志
	 * 
	 * @param message
	 */
	public static void debug(Throwable e) {
		logger.debug(e.getMessage());
	}
	public static void debug(String e) {
		logger.debug(e);
	}

	/**
	 * 输出信息级别的日志
	 * 
	 * @param info
	 */
	public static void info(Object info) {
		logger.info(String.valueOf(info));
	}
	
	
	
}
