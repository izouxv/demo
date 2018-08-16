/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

import org.apache.log4j.Logger;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.socket.TextMessage;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月7日 下午6:50:34
 * @$
 * @Administrator
 * @explain
 */
@Controller
@RequestMapping("/web")
public class WebsocketController {

	private Logger log = Logger.getLogger(WebsocketController.class);

	/**
	 * 这个注解会从Spring容器拿出Bean
	 */
	@Bean
	public TextWebSocket infoHandler() {
		return new TextWebSocket();
	}

	@ResponseBody
	@RequestMapping("/login")
	public String login(HttpServletRequest request, HttpServletResponse response) throws Exception {
		String username = request.getParameter("username");
		log.info(username + "登录");
		HttpSession session = request.getSession();
		log.info("session:" + session.getId());
		session.setAttribute("username", username);
		// response.sendRedirect("/quicksand/jsp/websocket.jsp");
		return "websocket";
	}

	@RequestMapping("/websocket/send")
	@ResponseBody
	public String send(HttpServletRequest request) {
		String username = request.getParameter("username");
		infoHandler().sendMessageToUser("1", new TextMessage("你好，测试！！！！"));
		return null;
	}

}
