/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.service;

import com.plys.rtcs.po.Proto;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月1日 下午6:18:35
 * @$
 * @Administrator
 * @explain 
 */

public interface WebMessageService {

	<T> void parseProto(Proto<T> proto);
	
}

