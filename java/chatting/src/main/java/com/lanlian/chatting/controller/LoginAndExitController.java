package com.lanlian.chatting.controller;

import java.util.Set;

import javax.annotation.Resource;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.model.Login;
import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Fatal_Exception;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.LoginService;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.rpc.sso.AgentInfo;
import com.lanlian.rpc.sso.SsoReply;
import com.lanlian.rpc.sso.SsoRequest;
import com.lanlian.rpc.sso.SsoRequest.Builder;
import com.lanlian.server.redis.RedisClient79;

@Controller
@RequestMapping(consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class LoginAndExitController extends MyAbstractController {

	private static Logger log = Logger.getLogger(LoginAndExitController.class);

	@Autowired
	private LoginService loginService;

	@Resource
	RedisClient79 redisClient79;

	@Resource
	SsoService ssoService;

	/**
	 * 登录功能
	 * 
	 * @param <T>
	 * 
	 * @param username
	 *            用户登录用户名
	 * @param password
	 *            用户登录密码
	 * @param device
	 *            用户登录设备
	 * @param request
	 *            Http请求
	 * @return 状态码
	 * @throws Parameter_Exception
	 * @throws Fatal_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.USER_PARENT + RequestSetting.USER_BODY_LOGIN)
	public String login(@RequestParam(value = "username", required = true) String username,
			@RequestParam(value = "password", required = true) String password,
			@RequestParam(value = "device", required = false) String device,
			@RequestParam(value = "imei", required = false) String imei, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception, Fatal_Exception {
		log.info("login-Username：" + username + ",Password：" + password + ",Device:" + device + ",Imei:" + imei);
		// 用户名判空校验
		ParameterVerify.usernameNull(username);
		// 检验用户名的正则验证
		ParameterVerify.usernameVerify(username);
		// 检验密码的合理性
		ParameterVerify.pwdNull(password);
		// 检验用户密码的正则表达验证
		ParameterVerify.pwdVerify(password);
		// 声明返回模型
		Login login = new Login();
		// 将数据存入对象进入业务层处理数据
		SsoPo ssoPoCon = getContextSsoPo(request, response);
		String ip = getContextIP(request, response);
		String dev = getContextDev(request, response);
		SsoPo ssoPo = new SsoPo();
		ssoPo.setUsername(username);
		ssoPo.setPassword(password);
		ssoPo.setSource(ssoPoCon.getSource());
		ssoPo.setIp(ip);
		ssoPo.setLoginDevice(dev);
		if (imei != null) {
			// imeiVerify(imei);
			ssoPo.setImei(imei);
		}
		ssoPo = loginService.ssoRpc(ssoPo);
		AccountPo accountPo = new AccountPo();
		accountPo.setUid(ssoPo.getUid());
		accountPo.setSource(ssoPoCon.getSource());
		accountPo = loginService.accountRpc(accountPo);
		// 将数据结果汇集并返回
		login.setUid(ssoPo.getUid());
		login.setUsername(ssoPo.getUsername());
		login.setLoginState(ssoPo.getLoginState());
		login.setState(ssoPo.getState());
		login.setErrorCode(ssoPo.getErrorCode());
		login.setToken(ssoPo.getSessionName());
		login.setNickname(ssoPo.getNickname());
		if (accountPo != null) {
			// BeanUtils.copyProperties(accountPo, login);
			login.setPhone(accountPo.getPhone());
			login.setEmail(accountPo.getEmail());
			login.setGender(accountPo.getGender());
			login.setBirthday(accountPo.getBirthday());
			login.setAvatar(accountPo.getAvatar());
			login.setSignature(accountPo.getSignature());
			login.setProvince(accountPo.getProvince());
			login.setCity(accountPo.getCity());
			login.setAddress(accountPo.getUserAddress());
			login.setRealname(accountPo.getRealname());
			login.setIdentityCard(accountPo.getIdentityCard());
			login.setCreditValues(accountPo.getCreditValues());
			login.setPoint(accountPo.getUserPoint());
			login.setJob(accountPo.getUserJobId());
			login.setGrade(accountPo.getUserGradeId());
			login.setErrorCode(accountPo.getErrorCode());
		}
		Set<Integer> allTGId = redisClient79.getAllTGId(ssoPo.getUid());
		if (allTGId.size() != 16) {
			throw new Fatal_Exception("从redis中获取到的可用群id不等于16：" + allTGId.size());
		}
		login.setGids(allTGId);
		return JSON.toJson(login);
	}

	/**
	 * 用于用户退出登录操作
	 * 
	 * @param pid
	 *            用户嗒嗒id
	 * @param imei
	 *            用户手机串码
	 * @return 状态码
	 * @throws Parameter_Exception
	 * @throws ServletException
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.USER_PARENT + RequestSetting.USER_BODY_EXIT)
	public String exit(HttpServletRequest request, HttpServletResponse response) throws Parameter_Exception {
		SsoPo ssoPo = getContextSsoPo(request, response);
		ssoPo.setSource(ssoPo.getSource());
		ssoPo.setSessionName(ssoPo.getSessionName());
		ssoPo = ssoService.verifyToken(ssoPo);
		log.info("exit-token:" + ssoPo);
		loginService.exitUser(ssoPo);
		return JSON.toJson();
	}

	@Resource
	SsoClient ssoClient;

	/**
	 * 快捷登录功能
	 * 
	 * @param <T>
	 * 
	 * @param username
	 *            用户登录用户名
	 * @param password
	 *            用户登录密码
	 * @param device
	 *            用户登录设备
	 * @param request
	 *            Http请求
	 * @return 状态码
	 * @throws Parameter_Exception
	 * @throws Fatal_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.SESSIONS)
	public String sessions(@RequestParam(value = "areacode", required = false) String areacode, // 区号
			@RequestParam(value = "username", required = true) String username, // 用户名
			@RequestParam(value = "captcha", required = false) String captcha, // 验证码
			@RequestParam(value = "password", required = false) String password, // 密码
			@RequestParam(value = "device", required = false) String device, // 设备名称
			@RequestParam(value = "imei", required = false) String imei, // 手机串号
			HttpServletRequest request, HttpServletResponse response) throws Parameter_Exception, Fatal_Exception {
		log.info("login-areacode:" + areacode + ",username:" + username + ",captcha:" + captcha + ",password:"
				+ password + ",device:" + device + ",imei:" + imei);
		// 用户名判空校验
		ParameterVerify.usernameNull(username);
		// 检验用户名的正则验证
		ParameterVerify.usernameVerify(username);
		// 检验验证码或密码
		if ((captcha == null || captcha.trim().isEmpty()) && (password == null || password.trim().isEmpty())) {
			throw new Parameter_Exception(21001);
		}
		// 声明返回模型
		Login login = new Login();
		// 将数据存入对象进入业务层处理数据
		SsoPo ssoPoCon = getContextSsoPo(request, response);
		String ip = getContextIP(request, response);
		String dev = getContextDev(request, response);
		Builder builder = SsoRequest.newBuilder();
		builder.setUsername(username);
		if (captcha != null && !captcha.trim().isEmpty()) {
			ParameterVerify.checkCodeVerify(captcha);
			builder.setCode(captcha);
		}
		if (password != null && !password.trim().isEmpty()) {
			ParameterVerify.pwdVerify(password);
			builder.setPassword(password);
		}
		builder.setSource(ssoPoCon.getSource());
		com.lanlian.rpc.sso.AgentInfo.Builder abuilder = AgentInfo.newBuilder();
		abuilder.setIp(ip);
		abuilder.setDevInfo(dev);
		builder.setAgentInfo(abuilder.build());
		SsoReply ssoRes = ssoClient.loginAll(builder);
		if (ssoRes.getErrorCode() != 10000) {
			throw new Parameter_Exception(ssoRes.getErrorCode());
		}
		AccountPo accountPo = new AccountPo();
		accountPo.setUid(ssoRes.getUid());
		accountPo.setSource(ssoPoCon.getSource());
		accountPo = loginService.accountRpc(accountPo);
		// 将数据结果聚合并返回
		login.setUid(ssoRes.getUid());
		login.setUsername(ssoRes.getUsername());
		login.setLoginState(ssoRes.getState());
		login.setState(ssoRes.getState());
		login.setErrorCode(ssoRes.getErrorCode());
		login.setToken(ssoRes.getSessionName());
		login.setNickname(ssoRes.getNickname());
		if (accountPo != null) {
			login.setPhone(accountPo.getPhone());
			login.setEmail(accountPo.getEmail());
			login.setGender(accountPo.getGender());
			login.setBirthday(accountPo.getBirthday());
			login.setAvatar(accountPo.getAvatar());
			login.setSignature(accountPo.getSignature());
			login.setProvince(accountPo.getProvince());
			login.setCity(accountPo.getCity());
			login.setAddress(accountPo.getUserAddress());
			login.setRealname(accountPo.getRealname());
			login.setIdentityCard(accountPo.getIdentityCard());
			login.setCreditValues(accountPo.getCreditValues());
			login.setPoint(accountPo.getUserPoint());
			login.setJob(accountPo.getUserJobId());
			login.setGrade(accountPo.getUserGradeId());
			login.setErrorCode(accountPo.getErrorCode());
		}
		Set<Integer> allTGId = redisClient79.getAllTGId(ssoRes.getUid());
		if (allTGId.size() != 16) {
			throw new Fatal_Exception("从redis中获取到的可用群id不等于16：" + allTGId.size());
		}
		login.setGids(allTGId);
		return JSON.toJson(login);
	}

}
