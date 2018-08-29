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
package com.lanlian.chatting.result;

import java.io.Serializable;

import com.alibaba.fastjson.JSONObject;
import com.alibaba.fastjson.serializer.SerializerFeature;

/**
 * @Title JSON.java
 * @Package com.lanlian.ccat.json
 * @author 王东阳
 * @version V1.0
 * @date 2017年2月21日 下午5:20:15
 * @explain 封装返回结果类（泛型T）
 */

public class JSON<T> implements Serializable {

	private static final long serialVersionUID = 5267628520350537080L;

	/**
	 * 成功的请求，返回成功提示，无返回结果
	 * 
	 * @param code
	 * @return
	 */
	public static <T> String toJson() {
		return JSONObject.toJSONString(new JSON<>(), SerializerFeature.WriteMapNullValue);
	}

	/**
	 * 客户端请求成功，返回结果
	 * 
	 * @param code
	 * @param result
	 * @return <T> String
	 */
	public static <T> String toJson(T result) {
		return JSONObject.toJSONString(new JSON<>(result), SerializerFeature.WriteMapNullValue);
	}

	/**
	 * 请求失败，友情提示
	 * 
	 * @param error
	 * @return
	 */
	public static String toJson(int error) {
		return JSONObject.toJSONString(new JSON<>(error), SerializerFeature.WriteMapNullValue);
	}

	/**
	 * 成功时，返回状态码：10000
	 */
	private static final int SUCCESS = 10000;

	/**
	 * 返回状态码：code
	 */
	private int code;

	/**
	 * 返回请求结果信息
	 */
	private String msg;

	/**
	 * 返回结果集：result
	 */
	private T result;

	/**
	 * 请求成功，无需返回结果
	 */
	public JSON() {
		this.code = SUCCESS;
		this.msg = ErrorCode.map.get(SUCCESS);
	}

	/**
	 * 请求成功，需要返回结果
	 * 
	 * @param result
	 */
	public JSON(T result) {
		this.code = SUCCESS;
		this.msg = ErrorCode.map.get(SUCCESS);
		this.result = result;
	}
	
	/**
	 * 请求失败，友情错误提示
	 * 
	 * @param code
	 */
	public JSON(int code) {
		this.code = code;
		this.msg = ErrorCode.map.get(code);
	}
	

	@Override
	public String toString() {
		return "JSON [code=" + code + ", msg=" + msg + ", result=" + result + "]";
	}

	public int getCode() {
		return code;
	}

	public void setCode(int code) {
		this.code = code;
	}

	public String getMsg() {
		return msg;
	}

	public void setMsg(String msg) {
		this.msg = msg;
	}

	public T getResult() {
		return result;
	}

	public void setResult(T result) {
		this.result = result;
	}

}
