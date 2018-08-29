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
import com.lanlian.chatting.po.Types;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.util.IPv4Util;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @Title CodeVerifyController.java
 * @Package com.lanlian.ccat.controller
 * @author 王东阳
 * @version V1.0.3
 * @date 2017年3月18日 下午2:08:07 @explain： 校验验证码
 */
@Controller
@RequestMapping(method = RequestMethod.POST, consumes = RequestSetting.CONSUMES, produces = RequestSetting.PRODUCES)
public class CodeVerifyController extends MyAbstractController {

	@Resource
	SsoService ssoService;

	@Resource
	SsoClient ssoClient;

	/**
	 * 验证激活码，通过验证后可以进行密码的重置
	 * 
	 * @param username
	 *            手机号码
	 * @param checkCode
	 *            验证码
	 * @return String
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.VISITORS_PARENT + RequestSetting.VERIFY_BODY)
	public String verifyCode(@RequestParam(value = "username", required = true) String username,
			@RequestParam(value = "checkCode", required = true) String checkCode, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("verify_code-username:" + username + ",chackCode:" + checkCode);
		// 进行参数校验
		ParameterVerify.usernameNull(username);
		int flag = ParameterVerify.usernameVerify(username);
		ParameterVerify.checkCodeVerify(checkCode);
		SsoPo ssoPo = getContextSsoPo(request, response);
		String ip = IPv4Util.getIP(request);
		LogUtil.info("ip:" + ip);
		switch (flag) {
		case 1:
			LogUtil.info("username:" + username + ",chackCode:" + checkCode);
			// 校验验证码的正确性
			ssoService.verifyMobile(username, checkCode, 1, ssoPo.getSource());
			return JSON.toJson();
		default:
			// 用户名格式不合法
			throw new Parameter_Exception(20002);
		}
	}

	/**
	 * 验证激活码，通过验证后可以进行密码的重置
	 * 
	 * @param types
	 * @param username
	 * @param code
	 * @param request
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.VISITORS_PARENT_1_1 + RequestSetting.VERIFY_BODY)
	public String verifyCode(@RequestParam(value = "types", required = true) Integer types,
			@RequestParam(value = "username", required = true) String username,
			@RequestParam(value = "code", required = true) String code, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("verify_code-types:" + types + ",username:" + username + ",code:" + code);
		// 进行参数校验
		if (!Types.contains(types)) {
			throw new Parameter_Exception(21002);
		}
		ParameterVerify.usernameNull(username);
		ParameterVerify.verifyPhone(username);
		ParameterVerify.checkCodeVerify(code);
		SsoPo ssoPo = getContextSsoPo(request, response);
		ssoClient.verifyMobileCode(username, code, types, ssoPo.getSource());
		return JSON.toJson();
	}

}
