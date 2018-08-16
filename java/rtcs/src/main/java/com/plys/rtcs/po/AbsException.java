/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.po;

import java.io.Serializable;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月1日 下午3:24:22
 * @$
 * @Administrator
 * @explain 
 */

public class AbsException extends Exception implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 8064273358065867251L;

	public AbsException() {
		super();
	}

	public AbsException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace) {
		super(message, cause, enableSuppression, writableStackTrace);
	}

	public AbsException(String message, Throwable cause) {
		super(message, cause);
	}

	public AbsException(String message) {
		super(message);
	}
	
	public AbsException(Integer message) {
		super(message.toString());
	}

	public AbsException(Throwable cause) {
		super(cause);
	}

}

