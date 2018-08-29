/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.controller;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.RealNameService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @Title RealNameController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年2月19日 下午3:54:55
 * @explain
 */

@Controller
@RequestMapping(value = RequestSetting.USER_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class RealNameController extends MyAbstractController {

	@Autowired
	private RealNameService realNameService;

	/**
	 * 实名认证功能
	 * 
	 * @param uid
	 *            用户嗒嗒id
	 * @param realName
	 *            用户真实姓名
	 * @param idCard
	 *            用户身份证号码
	 * @return String
	 * @throws Parameter_Exception
	 */
	@RequestMapping(value = RequestSetting.USER_BODY_REALNAME)
	@ResponseBody
	public String realName(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "realname", required = true) String realname,
			@RequestParam(value = "idcard", required = true) String idcard, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("realName-uid:" + uid + ",realname：" + realname + ",idcard:" + idcard);
		ParameterVerify.verifyNull(realname, idcard);
		// uid格式校验
		ParameterVerify.verifyUid(uid);
		// 进行判空校验
		realNameNull(realname);
		idCardNull(idcard);
		// 检验用户姓名格式是否合法
		// realNameVerify(realname);
		// 检验用户身份证号格式是否合法
		// idCardVerify(idcard);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		// 数据聚合，进入业务层
		AccountPo accountPo = new AccountPo();
		accountPo.setUid(ssoPo.getUid());
		accountPo.setSource(ssoPo.getSource());
		// realNameService.getCertification(accountPo);
		// accountPo.setUid(uid);
		accountPo.setRealname(realname);
		accountPo.setIdentityCard(idcard);
		realNameService.updateCertification(accountPo);
		LogUtil.info("realName:" + accountPo);
		return JSON.toJson();
	}

	/**
	 * 身份证号码格式
	 * 
	 * @param idCard
	 * @throws Parameter_Exception
	 */
	// private void idCardVerify(String idCard) throws Parameter_Exception {
	//
	// if (!Validator.isIDCard(idCard)) {
	// // 身份证号码格式不合法
	// throw new Parameter_Exception(20020);
	// }
	// System.out.println("year:" + Integer.parseInt(idCard.substring(6, 10)));
	// }

	/**
	 * 身份证号码判空
	 * 
	 * @param idCard
	 * @throws Parameter_Exception
	 */
	private void idCardNull(String idCard) throws Parameter_Exception {
		if (idCard == null || idCard.trim().isEmpty()) {
			// 身份证号码不能为空
			throw new Parameter_Exception(20019);
		}
	}

	/**
	 * 姓名格式
	 * 
	 * @param realName
	 * @throws Parameter_Exception
	 */
	// private void realNameVerify(String realName) throws Parameter_Exception {
	// String regName = "^[\u4e00-\u9fa5]+(·[\u4e00-\u9fa5]+)*$";
	// if (!realName.trim().matches(regName)) {
	// // 姓名格式不合法
	// throw new Parameter_Exception(20018);
	// }
	// }

	/**
	 * 姓名判空
	 * 
	 * @param realName
	 * @throws Parameter_Exception
	 */
	private void realNameNull(String realName) throws Parameter_Exception {
		if (realName == null || realName.trim().isEmpty()) {
			// 姓名不能为空
			throw new Parameter_Exception(20017);
		}
	}

}
