/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/** 
 * @Title Parameter_Exception.java
 * @Package com.lanlian.chatting.exception
 * @author 王东阳
 * @version V1
 * @date 2017年5月19日 上午10:22:05
 * @explain 
 */
package com.lanlian.chatting.result;

import java.io.Serializable;

/** 
 * @Title Parameter_Exception.java
 * @Package com.lanlian.chatting.exception
 * @author 王东阳
 * @version V1.0.3
 * @date 2017年4月19日 上午10:22:05
 * @explain 抛出异常,根据传进的int类型参数code，返回提示
 */

public class Parameter_Exception extends Exception implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 5072791519792842329L;

	public Parameter_Exception() {
		super();
	}

	public Parameter_Exception(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace) {
		super(message, cause, enableSuppression, writableStackTrace);
	}

	public Parameter_Exception(String message, Throwable cause) {
		super(message, cause);
	}

	public Parameter_Exception(String message) {
		super(message);
	}
	
	public Parameter_Exception(Integer code) {
		super(code.toString());
	}

	public Parameter_Exception(Throwable cause) {
		super(cause);
	}
	
}

