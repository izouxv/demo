package com.lanlian.chatting.controller;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * 该类用于用户的注册操作
 * 
 * @author my
 */
@Controller
@RequestMapping(value = RequestSetting.VISITORS_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class RegisterController extends MyAbstractController {

	@Resource
	SsoService ssoService;

	/**
	 * 用户注册功能
	 * 
	 * @param username
	 * @param password
	 * @param nickname
	 * @param checkCode
	 * @return 注册信息；
	 * @throws Parameter_Exception
	 */
	@RequestMapping(value = RequestSetting.REGISTER_BODY)
	@ResponseBody
	public String register(@RequestParam(value = "username", required = true) String username,
			@RequestParam(value = "password", required = true) String password,
			@RequestParam(value = "nickname", required = true) String nickname,
			@RequestParam(value = "checkCode", required = false) String checkCode, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("register-username:" + username + ",password:" + password + ",nickname:" + nickname + ",checkCode:"
				+ checkCode);
		ParameterVerify.verifyNull(username, password, nickname);
		// 进行用户名判空
		ParameterVerify.usernameNull(username);
		// 检验用户名格式
		int un = ParameterVerify.usernameVerify(username);
		// 检验昵称
		ParameterVerify.nicknameVerify(nickname);
		// 进行密码判空
		ParameterVerify.pwdNull(password);
		// 检验密码格式
		ParameterVerify.pwdVerify(password.trim());
		SsoPo ssoPoCon = getContextSsoPo(request, response);
		String ip = getContextIP(request, response);
		String dev = getContextDev(request, response);
		SsoPo ssoPo = new SsoPo();
		ssoPo.setSource(ssoPoCon.getSource());
		// 检验验证码
		if (un == 1) {
			ParameterVerify.checkCodeVerify(checkCode);
			// 校验验证码的正确性
			ssoService.verifyMobile(username, checkCode, 1, ssoPo.getSource());
		}
		// 检验用户是否被注册
		ssoPo.setUsername(username);
		int checkUser = ssoService.checkUser(ssoPo);
		if (checkUser == 2) {
			throw new Parameter_Exception(20008);
		}
		// sso中添加用户
		ssoPo.setUsername(username);
		ssoPo.setPassword(password);
		ssoPo.setNickname(nickname);
		ssoPo.setSalt(getSalt());
		ssoPo.setState(3);
		ssoPo.setIp(ip);
		ssoPo.setLoginDevice(dev);
		ssoPo = ssoService.add(ssoPo);
		if (ssoPo.getErrorCode().intValue() != 10000) {
			if (ssoPo.getErrorCode().intValue() == 33005) {
				throw new Parameter_Exception(20014);
			}
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
		return JSON.toJson();
	}

}
