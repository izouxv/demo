/** 
 *<p>开发公司 :		           蓝涟科技 <p>
 *<p>版权所有 :		           蓝涟科技 <p>
 *<p>责任人     :		               王东阳 <p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.config.annotation.EnableWebSocket;
import org.springframework.web.socket.config.annotation.WebSocketConfigurer;
import org.springframework.web.socket.config.annotation.WebSocketHandlerRegistry;
import org.springframework.web.socket.server.standard.ServletServerContainerFactoryBean;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月7日 下午2:44:59
 * @explain
 */

@Configuration
@EnableWebMvc
@EnableWebSocket
public class WebsocketConfig extends WebMvcConfigurerAdapter implements WebSocketConfigurer {

	@Bean
	public WebSocketHandler textWebSocket() {
		return new TextWebSocket();
	}

	@Bean
	public WebSocketHandler binaryWebSocket() {
		return new BinaryWebSocket();
	}

	@Bean
	public VirtualWebsocket abstractWebsocket() {
		return new VirtualWebsocket();
	}

	@Bean
	public MyWebSocketInterceptor myInterceptor() {
		return new MyWebSocketInterceptor();
	}

	@Bean
	public ServletServerContainerFactoryBean createWebSocketContainer() {
		ServletServerContainerFactoryBean container = new ServletServerContainerFactoryBean();
		container.setMaxTextMessageBufferSize(10485760);
		container.setMaxBinaryMessageBufferSize(10485760);
		return container;
	}

	@Override
	public void registerWebSocketHandlers(WebSocketHandlerRegistry registry) {
		// abstract
		registry.addHandler(abstractWebsocket(), "/web/abstract").setAllowedOrigins("*")
				.addInterceptors(myInterceptor());
		registry.addHandler(textWebSocket(), "/sockjs/abstract").setAllowedOrigins("*").addInterceptors(myInterceptor())
				.withSockJS();

		// // text
		// registry.addHandler(textWebSocket(), "/web/text").setAllowedOrigins("*")
		// .addInterceptors(myInterceptor());
		// registry.addHandler(textWebSocket(), "/sockjs/text").setAllowedOrigins("*")
		// .addInterceptors(myInterceptor()).withSockJS();
		// // binary
		// registry.addHandler(binaryWebSocket(), "/web/binary").setAllowedOrigins("*")
		// .addInterceptors(myInterceptor());
		// registry.addHandler(binaryWebSocket(),
		// "/sockjs/binary").setAllowedOrigins("*")
		// .addInterceptors(myInterceptor()).withSockJS();
	}

}
