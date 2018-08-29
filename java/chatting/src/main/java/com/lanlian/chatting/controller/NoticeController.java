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
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.UserLetterService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.UserMessageNum;

import net.sf.json.JSONArray;

/**
 * @Title NoticeController.java
 * @Package cn.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年2月18日 上午10:02:12
 * @explain 消息通知功能（前端登录后的轮询接口）
 */
@Controller
@RequestMapping(value = RequestSetting.NOTICE_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class NoticeController extends MyAbstractController {

	@Resource
	UserLetterService letterService;

	@ResponseBody
	@RequestMapping(value = RequestSetting.NOTICE_BODY_NOTICE)
	public String dadaNotice(@RequestParam(value = "uid", required = true) int uid, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("dadaNotice-uid:" + uid);
		// 进行参数格式校验
		ParameterVerify.verifyUid(uid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		UserMessageNum num = new UserMessageNum();
		num.setOpid(ssoPo.getUid());
		JSONArray jsonArray = letterService.receiveLetteNum(num);
		LogUtil.info(JSON.toJson(jsonArray));
		return JSON.toJson(jsonArray);
	}

}
