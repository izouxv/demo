/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import java.security.Principal;
import java.util.HashMap;
import java.util.Map;

import javax.servlet.ServletContext;
import javax.servlet.ServletRequestEvent;
import javax.servlet.ServletRequestListener;
import javax.servlet.annotation.WebListener;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;

import org.springframework.context.ApplicationListener;
import org.springframework.messaging.simp.stomp.StompCommand;
import org.springframework.messaging.simp.stomp.StompHeaderAccessor;
import org.springframework.messaging.support.MessageHeaderAccessor;
import org.springframework.web.context.ContextLoader;
import org.springframework.web.context.WebApplicationContext;
import org.springframework.web.socket.messaging.SessionConnectedEvent;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月20日 下午4:27:25
 * @$
 * @Administrator
 * @explain
 */

@WebListener
public class RequestListener implements ServletRequestListener {

	@Override
	public void requestDestroyed(ServletRequestEvent sre) {
		// TODO Auto-generated method stub

	}

	@Override
	public void requestInitialized(ServletRequestEvent sre) {
		// TODO Auto-generated method stub

	}

	// @Override
	// public void requestDestroyed(ServletRequestEvent sre) {
	// HttpServletRequest request = (HttpServletRequest)sre.getServletRequest();
	// HttpSession session = request.getSession();
	// System.out.println("dest-req-ses:"+session.getId());
	// }

	/**
	 * 初始化
	 */
	// @Override
	// public void requestInitialized(ServletRequestEvent sre) {
	// HttpServletRequest request = (HttpServletRequest)sre.getServletRequest();
	//// System.out.println("ses:"+request.getSession().getId());
	// HttpSession session = request.getSession();
	// System.out.println("init-req-ses:"+session.getId());
	// session.setAttribute("id", "qwer");
	// }

}
