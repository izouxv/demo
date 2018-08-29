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
import com.lanlian.chatting.service.UpdatePwdService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @Title ResetController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0.3
 * @date 2017年5月19日 下午6:11:49
 * @explain 用于用户手机号重置密码
 */
@Controller
@RequestMapping(value = RequestSetting.VISITORS_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class ResetController extends MyAbstractController {

	@Resource
	UpdatePwdService updatePwdService;

	@Resource
	SsoService ssoService;

	/**
	 * 用户通过验证码验证后，重置密码功能
	 * 
	 * @param username
	 *            用户注册手机号码
	 * @param checkCode
	 *            验证码
	 * @param passwordNew
	 *            用户重置新密码
	 * @return 状态码
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.RESET_BODY)
	public String resetPwd(@RequestParam(value = "username", required = true) String username,
			@RequestParam(value = "checkCode", required = true) String checkCode,
			@RequestParam(value = "password", required = true) String password, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("resetPwd-username:" + username + ",checkCode:" + checkCode + ",password:" + password);
		ParameterVerify.verifyNull(username, checkCode, password);
		// 用户名判空
		ParameterVerify.usernameNull(username);
		// 用户名格式检验
		int flag = ParameterVerify.usernameVerify(username);
		if (flag == 2) {
			throw new Parameter_Exception(20002);
		}
		// 验证码判空/格式检验
		ParameterVerify.checkCodeVerify(checkCode);
		// 密码判空
		ParameterVerify.pwdNull(password);
		// 密码格式检验
		ParameterVerify.pwdVerify(password);
		SsoPo ssoPoCon = getContextSsoPo(request, response);
		SsoPo ssoPo = new SsoPo();
		ssoPo.setSource(ssoPoCon.getSource());
		// 检验验证码正确性
		ssoService.verifyMobile(username, checkCode, 2, ssoPo.getSource());

		// 进入业务层
		ssoPo.setUsername(username);
		ssoPo.setPassword(password);
		ssoPo.setSalt(getSalt());
		updatePwdService.resetPassword(ssoPo);
		return JSON.toJson();
	}

}
