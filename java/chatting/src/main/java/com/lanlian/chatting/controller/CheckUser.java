/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.controller;

import org.springframework.beans.factory.annotation.Autowired;
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
import com.wordnik.swagger.annotations.Api;

/**
 * @Title CheckUser.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月14日 下午5:28:29
 * @explain 检查用户名是都被注册
 */
@Controller
@RequestMapping(value = RequestSetting.USER_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
@Api(value = "验证用户名")
public class CheckUser extends MyAbstractController {

	@Autowired
	SsoService ssoService;

	@ResponseBody
	@RequestMapping(value = RequestSetting.USER_BODY_CHECK)
	public String checkUser(@RequestParam(value = "username", required = true) String username)
			throws Parameter_Exception {
		LogUtil.info("checkUser-request-Username：" + username + "}");
		// 用户名判空校验
		ParameterVerify.usernameNull(username);
		// 检验用户名的正则验证
		ParameterVerify.usernameVerify(username);

		// 将数据存入对象进入业务层处理数据
		SsoPo ssoPo = new SsoPo();
		ssoPo.setUsername(username);
		int check = ssoService.checkUser(ssoPo);
		if (check == 2) {
			return JSON.toJson();
		}
		if (check == 1) {
			return JSON.toJson(20009);
		}
		return JSON.toJson(20000);
	}

}
