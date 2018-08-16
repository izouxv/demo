/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.util.CollectionUtils;
import org.springframework.web.socket.config.annotation.WebSocketConfigurationSupport;
import org.springframework.web.socket.config.annotation.WebSocketConfigurer;
import org.springframework.web.socket.config.annotation.WebSocketHandlerRegistry;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月8日 下午2:35:39
 * @$
 * @Administrator
 * @explain
 */

@Configuration
public class DelegatingWebSocketConfiguration extends WebSocketConfigurationSupport {

	private final List<WebSocketConfigurer> configurers = new ArrayList<>();

	@Autowired(required = false)
	public void setConfigurers(List<WebSocketConfigurer> configurers) {
		System.out.println("Before setConfigurers" + configurers);
		if (!CollectionUtils.isEmpty(configurers)) {
			this.configurers.addAll(configurers);
		}
		System.out.println("Before setConfigurers");
	}

	@Override
	protected void registerWebSocketHandlers(WebSocketHandlerRegistry registry) {
		System.out.println("Before setConfigurers" + registry);
		for (WebSocketConfigurer configurer : this.configurers) {
			configurer.registerWebSocketHandlers(registry);
		}
		System.out.println("Before setConfigurers");
	}
}