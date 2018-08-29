/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/**
 * 
 */
package com.lanlian.chatting.controller;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.UpdatePwdService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @Title UserPwdController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0.3
 * @date 2017年5月3日 上午11:39:58
 * @explain 用户的密码修改
 */
@Controller
@RequestMapping(value = RequestSetting.USER_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class UserPwdController extends MyAbstractController {

	@Resource
	UpdatePwdService updatePwdService;

	@ResponseBody
	@RequestMapping(value = RequestSetting.USER_BODY_CHANGE)
	public String changePwd(@PathVariable(value = "action", required = true) String action,
			@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "pwd", required = true) String pwd,
			@RequestParam(value = "newPwd", required = false) String newPwd, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("changePwd-action:" + action + ",uid:" + uid + ",pwd:" + pwd + ",newPwd:" + newPwd);
		// 对参数进行验证
		ParameterVerify.verifyUid(uid);
		ParameterVerify.pwdNull(pwd);
		ParameterVerify.pwdVerify(pwd);
		SsoPo ssoPoCon = getContextSsoPo(request, response);
		if (ssoPoCon.getUid() != uid) {
			return JSON.toJson(21002);
		}
		SsoPo ssoPo = new SsoPo();
		ssoPo.setUid(ssoPoCon.getUid());
		ssoPo.setPassword(pwd);
		ssoPo.setSource(ssoPoCon.getSource());
		switch (action) {
		case "check":
			// 校验密码
			updatePwdService.checkPassword(ssoPo);
			return JSON.toJson();
		case "update":
			// 校验密码
			updatePwdService.checkPassword(ssoPo);
			// 进入修改密码
			ParameterVerify.pwdNull(newPwd);
			ParameterVerify.pwdVerify(newPwd);
			ssoPo.setUid(ssoPoCon.getUid());
			ssoPo.setPassword(newPwd);
			ssoPo.setSalt(getSalt());
			ssoPo.setSessionName(ssoPoCon.getSessionName());
			updatePwdService.updatePassword(ssoPo);
			return JSON.toJson();
		default:
			// "没有权限"
			throw new Parameter_Exception(21005);
		}
	}

}
