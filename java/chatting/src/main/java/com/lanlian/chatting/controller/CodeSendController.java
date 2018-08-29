/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */
package com.lanlian.chatting.controller;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.po.Types;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.util.DataFinals;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @Title CodeSendController.java
 * @Package com.lanlian.ccat.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月18日 下午2:52:41 @explain： 发送验证码
 */
@Controller
@RequestMapping(consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class CodeSendController extends MyAbstractController {

	@Autowired
	SsoService ssoService;

	@Resource
	SsoClient ssoClient;

	/**
	 * 用于用户注册或找回密码的操作，发送给手机或邮箱
	 * 
	 * @param username
	 * @return JSONObject @exception
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.VISITORS_PARENT + RequestSetting.CODE_BODY)
	public String sendCode(@PathVariable String action,
			@RequestParam(value = "username", required = true) String username, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		// 记录请求日志信息
		LogUtil.info("sendCode-action:" + action + ",username:" + username);
		// 进行参数判空验证
		ParameterVerify.usernameNull(username);
		// 进行参数格式验证：1：手机号，2：邮箱
		int un = ParameterVerify.usernameVerify(username);
		SsoPo ssoPoCon = getContextSsoPo(request, response);
		String ip = getContextIP(request, response);
		String dev = getContextDev(request, response);
		SsoPo ssoPo = new SsoPo();
		ssoPo.setSource(ssoPoCon.getSource());
		ssoPo.setUsername(username);
		ssoPo.setIp(ip);
		ssoPo.setLoginDevice(dev);
		// 对手机号或邮箱与注册或找回密码进行判断
		int flag = action(action, un);
		// 手机号发送短信
		if (flag == 1) {
			ssoService.sendMobile(ssoPo, action);
			return JSON.toJson();
		}
		// 邮箱发送邮件
		if (flag == 2) {
			ssoService.sendEmail(ssoPo);
			return JSON.toJson();
		}
		response.setStatus(404);
		return JSON.toJson("别逗！！！！！！！");
	}

	/**
	 * 用于用户注册或找回密码的操作，发送给手机或邮箱
	 * 
	 * @param username
	 * @return JSONObject
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.VISITORS_PARENT_1_1 + RequestSetting.CODE_BODY_1)
	public String sendCode(@RequestParam(value = "types", required = true) Integer types,
			@RequestParam(value = "username", required = true) String username, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		// 记录请求日志信息
		LogUtil.info("captcha-types:" + types + ",username:" + username);
		// 1-手机号;2-邮箱
		// 1-注册;2-找回密码;3-登录
		// 进行参数判空验证
		ParameterVerify.usernameNull(username);
		if (!Types.contains(types)) {
			throw new Parameter_Exception(21002);
		}
		SsoPo ssoPoCon = getContextSsoPo(request, response);
		String ip = getContextIP(request, response);
		String dev = getContextDev(request, response);
		SsoPo ssoPo = new SsoPo();
		ssoPo.setSource(ssoPoCon.getSource());
		ssoPo.setUsername(username);
		ssoPo.setCodeType(types);
		ssoPo.setIp(ip);
		ssoPo.setLoginDevice(dev);
		switch (ParameterVerify.usernameVerify(username)) {
		case 1:
			ssoPo = ssoClient.sendMobile(ssoPo);
			break;
		case 2:
			ssoPo = ssoClient.sendEmail(ssoPo);
			break;
		default:
			ssoPo.setErrorCode(21002);
		}
		if (ssoPo.getErrorCode() == 33008) {
			throw new Parameter_Exception(20008);
		}
		return JSON.toJson(ssoPo.getErrorCode().intValue());
	}

	/******************************/
	/**
	 * 判断注册或找回密码与手机号或邮箱或该用户是否注册
	 * 
	 * @param action
	 * @param un
	 * @param whether
	 * @return
	 * @throws Parameter_Exception
	 *             20008 手机号被注册/邮箱被注册 20009 手机号未注册/邮箱号未注册
	 * @throws Parameter_Exception
	 */
	private int action(String action, int un) throws Parameter_Exception {
		switch (action) {
		// 注册业务
		case DataFinals.REGISTER:
			// 手机号
			if (un == 1) {
				return 1;
			}
			// 邮箱
			if (un == 2) {
				throw new Parameter_Exception(20002);
			}
			// 找回密码
		case DataFinals.RESETPWD:
			// 手机号
			if (un == 1) {
				return 1;
			}
			// 邮箱
			if (un == 2) {
				return 2;
			}
		default:
			return 0;
		}
	}
}