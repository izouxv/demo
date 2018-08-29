/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 */

package com.lanlian.chatting.interceptor;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.List;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;
import org.springframework.core.NamedThreadLocal;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.handler.HandlerInterceptorAdapter;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.MetricsConfigService;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.util.CheckApp;
import com.lanlian.chatting.util.IPv4Util;
import com.lanlian.chatting.util.ParameterVerify;

import cz.mallat.uasparser.OnlineUpdater;
import cz.mallat.uasparser.UASparser;
import cz.mallat.uasparser.UserAgentInfo;
import io.grpc.StatusRuntimeException;
import redis.clients.jedis.exceptions.JedisConnectionException;
import redis.clients.jedis.exceptions.JedisException;

/**
 * 程序拦截器，所有请求进行拦截处理
 * 
 * @author Administrator
//WebRequestInterceptor
 *
 */
public class Interceptor extends HandlerInterceptorAdapter {
	
	private static Logger logger =  Logger.getLogger(Interceptor.class);

	@Resource
	private SsoService ssoService;
	
	@Resource
	private MetricsConfigService metricsConfigService;
	
	private List<String> ALLOW_URI;
	
	private NamedThreadLocal<Long> startTimeThreadlocal = new NamedThreadLocal<>("StopWatch-StartTime");

	private static UASparser uaSparser = null;

	/**
	 * 初始化uaSparser对象
	 */
	static {
		try {
			uaSparser = new UASparser(OnlineUpdater.getVendoredInputStream());
		} catch (IOException e) {
			System.out.println("获取客户端信息异常！！！");
			uaSparser = null;
			// log.debug(uaSparser.toString(), e);
		}
	}

	/**
	 * 进入 Handler方法之前执行,用于身份认证、身份授权； 比如身份认证，如果认证通不过表示当前用户没有登陆，需要此方法拦截不再向下执行；
	 */
	@Override
	public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
		try {
			startTimeThreadlocal.set(System.currentTimeMillis());//线程绑定变量（该数据只有当前请求的线程可见）
			String url = request.getServletPath();
			String ip = IPv4Util.getIP(request);
			String agent = request.getHeader("User-Agent");
			UserAgentInfo userAgentInfo = Interceptor.uaSparser.parse(agent);
			logger.info("request-url:" + url + ",ip:" + ip+
					"操作系统名称：" + userAgentInfo.getOsFamily() + ",操作系统：" + userAgentInfo.getOsName() 
					+ ",浏览器名称："+ userAgentInfo.getUaFamily() + ",浏览器版本：" + userAgentInfo.getBrowserVersionInfo() 
					+ ",设备类型：" + userAgentInfo.getDeviceType() + ",浏览器:" + userAgentInfo.getUaName() 
					+ ",类型：" + userAgentInfo.getType() + ",编码方式:" + request.getCharacterEncoding() + request.getContentType());
			request.setAttribute("ip", ip);
			request.setAttribute("dev", agent);
			// metrics-counter
			String metrics = request.getRequestURI();
			metricsConfigService.interfaceCounter(metrics);
			// 获取请求的url,判断url是否是公开地址
			SsoPo ssoPo = new SsoPo();
			ssoPo = verifySource(ssoPo, request);
			ssoService.getSource(ssoPo);
			if (openInterface(url)) {
				request.setAttribute("SsoPo", ssoPo);
				return true;
			}
			logger.info("get handinfo ......");
			//处理token信息
			verifyToken(ssoPo, request);
			return true;
		} catch (Parameter_Exception e) {
			logger.info(e.getMessage());
			return tryError(response, Integer.parseInt(e.getMessage()),e);
		} catch (StatusRuntimeException e) {
			logger.fatal("拦截器错误---preHandle---StatusRuntimeException:" + e);
			return tryError(response, 10002,e);
		} catch (JedisConnectionException e) {
			logger.fatal("拦截器错误---preHandle---JedisConnectionException:" + e);
			return tryError(response, 10002,e);
		} catch (JedisException e) {
			logger.fatal("拦截器错误---preHandle---JedisException:" + e);
			return tryError(response, 10002,e);
		} catch (Exception e) {
			logger.error("拦截器错误---preHandle---Exception:" + e);
			return tryError(response, 10001,e);
		}
	}
	
	/**
	 * 拦截器发生异常时捕捉并返回code
	 * @param response
	 * @param code
	 * @param e
	 * @return
	 */
	private Boolean tryError(HttpServletResponse response, int code,Exception e) {
		try {
			response.setCharacterEncoding("UTF-8");
			response.setContentType("application/json;charset=UTF-8");
			PrintWriter pw = response.getWriter();
			pw.println(JSON.toJson(code));
			pw.flush();
		} catch (Exception e1) {
			return false;
		}
		return false;
	}

	/**
	 * 检查请求头信息，并处理
	 * @param request
	 * @param uid
	 * @return
	 * @throws Parameter_Exception
	 */
	private void verifyToken(SsoPo ssoPo, HttpServletRequest request) throws Parameter_Exception {
		// 获取请求的url,判断url是否是半公开接口
		if (semiOenInterface(request.getServletPath())) {
			logger.info("semiOenInterface-sso:" + ssoPo);
			request.setAttribute("SsoPo", ssoPo);
			return;
		}
		// 兼容处理
		String token = getToken(request);
		ssoPo.setSessionName(token);
		ssoPo = ssoService.verifyToken(ssoPo);
		request.setAttribute("SsoPo", ssoPo);
		logger.info("Interceptor-sso:" + ssoPo);
	}

	/**
	 * 检查source
	 * @param ssoPo
	 * @param request
	 * @return
	 */
	private SsoPo verifySource(SsoPo ssoPo, HttpServletRequest request) {
		String source = request.getHeader("source");
		logger.info("verifyhander-source:" + source);
		// 兼容处理
		if (source == null || source.trim().isEmpty() || !source.equals("BAIDAA==")) {
			source = "AQIDAA==";
		}
		ssoPo.setSource(source);
		return ssoPo;
	}
	
	/**
	 * 获取token值
	 * @param req
	 * @return
	 * @throws Parameter_Exception
	 */
	public static String getToken(HttpServletRequest req) throws Parameter_Exception {
		String token = req.getHeader("token");
		if (token == null || token.trim().isEmpty()) {
			token = req.getHeader("radacat_app");
			if (token == null || token.trim().isEmpty()) {
				token = req.getHeader("radacat_wechat");
				if (token == null || token.trim().isEmpty()) {
					throw new Parameter_Exception(23019);
				}
			}
		}
		return token;
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
			request.removeAttribute("SsoPo");
	        long consumeTime = System.currentTimeMillis() - startTimeThreadlocal.get();//消耗时间=当前时间-线程绑定是局部变量时间
	        if(consumeTime > 100) {//此处认为处理时间超过1000毫秒的请求为慢请求
	        	logger.info(String.format("http-times : %s : consume %d ms", request.getRequestURI(), consumeTime));
	        }
		} catch (Exception e) {
			logger.error(e);
		}
	}

	/**
	 * 校验签名
	 * 
	 * @param appSecret
	 * @param nonce
	 * @param curTime
	 * @param signature
	 * @return
	 */
	@SuppressWarnings("unused")
	private boolean checkSignature(String appSecret, String nonce, String curTime, String signature) {
		try {
			// 对取出的参数进行判空校验
			ParameterVerify.verifyNull(new String[] { appSecret, nonce, curTime, signature });
			// 对客户端与服务器进行时间差校验
			boolean mistiming = timeMistiming(curTime);
			if (!mistiming) {
				return false;
			}
			// 对Signature进行运算，返回true，则通过
			if (!CheckApp.getSignature(appSecret, nonce, curTime).equals(signature)) {
				return false;
			}
			return true;
		} catch (Parameter_Exception e) {
			return false;
		} catch (Exception e) {
			return false;
		}
	}

	/**
	 * 对时间差进行校验，十分钟之内；
	 * 
	 * @param curTime
	 * @return
	 */
	private boolean timeMistiming(String curTime) {
		long now = System.currentTimeMillis() / 1000 / 60;
		long time = Long.parseLong(curTime);
		boolean flag1 = now >= time;
		boolean flag2 = (0 <= (now - time)) && ((now - time) < 10);
		if (flag1 && flag2) {
			return true;
		}
		return false;
	}
	
	/**
	 * allow_uri
	 * @return
	 */
	public List<String> getALLOW_URI() {
		return ALLOW_URI;
	}

	public void setALLOW_URI(List<String> aLLOW_URI) {
		ALLOW_URI = aLLOW_URI;
	}

	/**
	 * 公开接口
	 * @param url
	 * @return
	 */
	private boolean openInterface(String url) {
		if (url.contains("/file/")) {
			return true;
		}
		if (url.contains("/demo/")) {
			return true;
		}
		if (url.contains("/api/")) {
			return true;
		}
		if (url.contains("/api-docs/")) {
			return true;
		}
		return false;
	}
	
	/**
	 * 半公开接口
	 * @param url
	 * @return
	 */
	private boolean semiOenInterface(String url) {
		if (url.contains("/user/login")) {
			return true;
		}
		if (url.contains("/sessions")) {
			return true;
		}
		if (url.contains("/visitors/")) {
			return true;
		}
		if (url.contains("/captcha")) {
			return true;
		}
		if (url.contains("/v1.0/user/statistics/")) {
			return true;
		}
		if (url.contains("/v1.1/statistics/devinfo")) {
			return true;
		}
		if (url.contains("/files/advertisement/get")) {
			return true;
		}
		if (url.contains("/files/adver")) {
			return true;
		}
		if (url.contains("/v1.0/version/")) {
			return true;
		}
		return false;
	}

}
