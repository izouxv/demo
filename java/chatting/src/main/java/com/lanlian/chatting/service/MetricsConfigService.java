/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.service;

import org.springframework.scheduling.annotation.Async;

import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月14日 上午8:54:44
 * @explain 
 */

public interface MetricsConfigService {
	
	@Async
	void interfaceCounter(String uri) throws Parameter_Exception;
	
	void appear(String string);

}

