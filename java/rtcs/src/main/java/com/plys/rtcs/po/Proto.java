/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.po;

import com.alibaba.fastjson.JSON;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月1日 上午11:05:08
 * @$
 * @Administrator
 * @explain 
 */

public abstract class Proto<T> {
	
	public static <T> Object jsonToBean(String msg,T t) throws AbsException {
		try {
			System.out.println("jsonToBean:"+msg);
			if (t instanceof ProtoMsg) {
				return JSON.parseObject(msg,ProtoMsg.class);
			}
			return JSON.parseObject(msg,t.getClass());
		} catch (Exception e) {
			throw new AbsException(21002);
		}
	}
	public static <T> String beanToJson(T t) throws AbsException {
		try {
			System.out.println("beanToJson:"+t);
			return JSON.toJSONString(t);
		} catch (Exception e) {
			throw new AbsException(21002);
		}
	}

}

