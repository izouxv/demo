/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import java.util.Map;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;

import org.apache.log4j.Logger;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.http.server.ServletServerHttpRequest;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.server.HandshakeInterceptor;
import org.springframework.web.socket.server.support.HttpSessionHandshakeInterceptor;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月7日 下午3:09:23
 * @$
 * @Administrator
 * @explain
 */

public class MyWebSocketInterceptor extends HttpSessionHandshakeInterceptor {

	private Logger log = Logger.getLogger(MyWebSocketInterceptor.class);

	@Override
	public boolean beforeHandshake(ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler,
			Map<String, Object> attributes) {
		try {
			log.info("Before Handshake");
			if (request instanceof ServletServerHttpRequest) {
				log.info("getURI:" + request.getURI().toString() + ":" + request.getHeaders());
				ServletServerHttpRequest servletRequest = (ServletServerHttpRequest) request;
				HttpSession session = servletRequest.getServletRequest().getSession(true);
				log.info("session:" + session.getId() + ",attributes:" + attributes);
				session.setAttribute("token", "hello");
				if (session != null) {
					// 使用userName区分WebSocketHandler，以便定向发送消息
					String userName = (String) session.getAttribute("SESSION_USERNAME");
					System.err.println("userName:" + userName);
					if (userName == null) {
						userName = "default-system";
					}
					attributes.put("WEBSOCKET_USERNAME", userName);
					session.setAttribute("iid", "test2");
				}
			}
			boolean boo = super.beforeHandshake(request, response, wsHandler, attributes);
			log.info("boo:" + boo);
			return boo;
		} catch (Exception e) {
			e.printStackTrace();
			return false;
		}
	}

	/**
	 * 1.2
	 */
	@Override
	public void afterHandshake(ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler,
			Exception ex) {
		try {
			log.info("---------HttpSessionHandshakeInterceptor.afterHandshake-------------");
			// super.afterHandshake(request, response, wsHandler, ex);
		} catch (Exception e) {
			e.printStackTrace();
			log.info("---------HttpSessionHandshakeInterceptor.afterHandshake-------------");
		}
		log.info("---------HttpSessionHandshakeInterceptor.afterHandshake-------------");
	}

}
