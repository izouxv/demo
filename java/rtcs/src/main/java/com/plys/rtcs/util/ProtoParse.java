/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.util;

import java.nio.ByteBuffer;
import java.nio.CharBuffer;
import java.nio.charset.Charset;
import java.util.Map;

import org.springframework.web.socket.BinaryMessage;
import org.springframework.web.socket.WebSocketMessage;

import com.alibaba.fastjson.JSON;
import com.plys.rtcs.po.AbsException;
import com.plys.rtcs.po.User;


/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月1日 上午9:11:34
 * @$
 * @Administrator
 * @explain 
 */

public class ProtoParse {


//	public static <T> Object jsonToBean(String msg,T t) {
//		return JSON.parseObject(msg,t.getClass());
//	}
//	public static <T> String beanToJson(T t) {
//		return JSON.toJSONString(t);
//	}
	
	/**
	 * 
	 * @param byteBuffer
	 * @return
	 * @throws AbsException
	 */
	public static String byToStr(ByteBuffer byteBuffer) throws AbsException {
		try {
			Charset charset = Charset.forName("UTF-8");
			CharBuffer charBuffer = charset.decode(byteBuffer);
//			String str = charBuffer.toString();
//			System.out.println("charBuffer-"+charBuffer);
			return charBuffer.toString();
		} catch (Exception e) {
			throw new AbsException("解析binary异常");
		}
	}
	
	public static void syst(Map<String, User> users) {
		for (String key : users.keySet()) {
			System.out.println("key:"+users.get(key));
		}
	}

}

