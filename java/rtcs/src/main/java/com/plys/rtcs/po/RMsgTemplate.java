/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.po;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.ObjectOutputStream;
import java.io.Serializable;
import java.util.Arrays;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月13日 下午12:17:24
 * @$
 * @Administrator
 * @explain
 */

public class RMsgTemplate implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 2696992863971979311L;

	// 参数类型
	private Class<?>[] paramTypes;
	// 交换器
	private String exchange;

	private Object[] params;

	// 路由key
	private String routeKey;

	public RMsgTemplate() {
	}

	public RMsgTemplate(String exchange, String routeKey, Object... params) {
		this.params = params;
		this.exchange = exchange;
		this.routeKey = routeKey;
	}

	@SuppressWarnings("rawtypes")
	public RMsgTemplate(String exchange, String routeKey, String methodName, Object... params) {
		this.params = params;
		this.exchange = exchange;
		this.routeKey = routeKey;
		int len = params.length;
		Class[] clazzArray = new Class[len];
		for (int i = 0; i < len; i++) {
			clazzArray[i] = params[i].getClass();
		}
		this.paramTypes = clazzArray;
	}

	public byte[] getSerialBytes() {
		byte[] res = new byte[0];
		ByteArrayOutputStream baos = new ByteArrayOutputStream();
		ObjectOutputStream oos;
		try {
			oos = new ObjectOutputStream(baos);
			oos.writeObject(this);
			oos.close();
			res = baos.toByteArray();
		} catch (IOException e) {
			e.printStackTrace();
		}
		return res;
	}

	public String getRouteKey() {
		return routeKey;
	}

	public String getExchange() {
		return exchange;
	}

	public void setExchange(String exchange) {
		this.exchange = exchange;
	}

	public void setRouteKey(String routeKey) {
		this.routeKey = routeKey;
	}

	public Class<?>[] getParamTypes() {
		return paramTypes;
	}

	public Object[] getParams() {
		return params;
	}

	@Override
	public String toString() {
		return "RMsgTemplate [paramTypes=" + Arrays.toString(paramTypes) + ", exchange=" + exchange + ", params="
				+ Arrays.toString(params) + ", routeKey=" + routeKey + "]";
	}

}
