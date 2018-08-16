/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Map;
import java.util.Map.Entry;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

import org.springframework.core.NamedThreadLocal;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import cz.mallat.uasparser.OnlineUpdater;
import cz.mallat.uasparser.UASparser;
import cz.mallat.uasparser.UserAgentInfo;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月22日 上午9:24:45
 * @$
 * @Administrator
 * @explain 对所有HTTP请求进行拦截处理
 */

public class MyInterceptor implements HandlerInterceptor {

	// @Resource
	// private SsoService ssoService;

	// @Resource
	// private MetricsConfigService metricsConfigService;

	private static final NamedThreadLocal<Long> STARTTIMETHREADLOCAL = new NamedThreadLocal<>("StopWatch-StartTime");

	private static UASparser UASPARSER = null;

	/**
	 * 初始化uaSparser对象
	 */
	static {
		try {
			UASPARSER = new UASparser(OnlineUpdater.getVendoredInputStream());
		} catch (IOException e) {
			System.out.println("获取客户端信息异常！！！");
			UASPARSER = null;
			// log.debug(uaSparser.toString(), e);
		}
	}

	/**
	 * 进入 Handler方法之前执行,用于身份认证、身份授权； 比如身份认证，如果认证通不过表示当前用户没有登陆，需要此方法拦截不再向下执行；
	 */
	@Override
	public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
		try {
			// 1、开始时间
			long beginTime = System.currentTimeMillis();
			// 线程绑定变量（该数据只有当前请求的线程可见）
			STARTTIMETHREADLOCAL.set(beginTime);
			request.setAttribute("times", System.currentTimeMillis());
			UserAgentInfo userAgentInfo = UASPARSER.parse(request.getHeader("User-Agent"));
			System.out.println("操作系统名称：" + userAgentInfo.getOsFamily() + ",操作系统：" + userAgentInfo.getOsName()
					+ ",浏览器名称：" + userAgentInfo.getUaFamily() + ",浏览器版本：" + userAgentInfo.getBrowserVersionInfo()
					+ ",设备类型：" + userAgentInfo.getDeviceType() + ",浏览器:" + userAgentInfo.getUaName() + ",类型："
					+ userAgentInfo.getType() + ",编码方式:" + request.getCharacterEncoding() + request.getContentType());

			HttpSession session = request.getSession();

			String url = request.getRequestURI();
			String referer = request.getHeader("Referer");
			request.getHeader("");

			// if (url.contains("/web/text")) {
			// return false;
			// }
			session.setAttribute("id", "test");
			// metrics-counter
			// metricsConfigService.interfaceCounter(url);

			// 获取请求的url,判断url是否是公开地址
			// 处理请求头的信息
			// verifyhander(request, uid);
			System.out.println("request-url:" + url + ",referer:" + referer + ",session:" + session.getId());
			return true;
		} catch (IOException e) {
			try {
				response.setCharacterEncoding("UTF-8");
				response.setContentType("application/json;charset=UTF-8");
				PrintWriter pw = response.getWriter();
				pw.println(Integer.parseInt(e.getMessage()));
				pw.flush();
			} catch (IOException e1) {
				return false;
			}
			return false;
		} catch (Exception e) {
			return false;
		}

	}

	/**
	 * 
	 */
	@Override
	public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler,
			ModelAndView modelAndView) {
	}

	/**
	 * @see org.springframework.web.servlet.HandlerInterceptor#afterCompletion
	 *      (javax.servlet.http.HttpServletRequest,
	 *      javax.servlet.http.HttpServletResponse, java.lang.Object,
	 *      java.lang.Exception)
	 */
	@Override
	public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object object,
			Exception exception) throws Exception {
		try {
			Long times = (Long) request.getAttribute("times");
			System.out.println("time-consuming:" + (System.currentTimeMillis() - times.longValue()));
			// 2、结束时间
			long endTime = System.currentTimeMillis();
			// 得到线程绑定的局部变量（开始时间）
			long beginTime = STARTTIMETHREADLOCAL.get();
			// 3、消耗的时间
			long consumeTime = endTime - beginTime;
			// 此处认为处理时间超过1000毫秒的请求为慢请求
			if (consumeTime > 100) {
				System.out.println(String.format("%s consume %d millis", request.getRequestURI(), consumeTime));
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
