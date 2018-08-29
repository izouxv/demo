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

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONArray;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Fatal_Exception;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.DadaGroupMsgService;
import com.lanlian.chatting.service.DadaGroupService;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.server.http.Wechat;

/**
 * @Title DadaGroupMsgController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午8:27:04
 * @explain 开启群信息实时上报
 */
@Controller
@RequestMapping(value = RequestSetting.GROUP_MESSAGE_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class DadaGroupMsgController extends MyAbstractController {

	@Resource
	@Qualifier("dadaGroupMsgServiceImpl")
	DadaGroupMsgService dadaGroupMsgService;

	@Resource
	@Qualifier("dadaGroupServiceImpl")
	DadaGroupService dadaGroupService;

	@Resource
	Wechat wechat;

	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_MESSAGE_BODY_SYNC_GID)
	public String dadaGroupMsg(@PathVariable(value = "gid", required = true) int gid,
			@RequestParam(value = "msg", required = true) String msg, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception, Fatal_Exception {
		// 获取上下文
		SsoPo ssoPo = getContextSsoPo(request, response);
		//
		DadaGroupPo dadaGroupPo = new DadaGroupPo();
		dadaGroupPo.setGid(gid);
		dadaGroupPo.setUid(ssoPo.getUid());
		dadaGroupPo = dadaGroupService.findDadaGroupSwitchingUser(dadaGroupPo);
		if (dadaGroupPo == null) {
			throw new Parameter_Exception(21005);
		}
		if (dadaGroupPo.getUpid() != ssoPo.getUid()) {
			throw new Parameter_Exception(21005);
		}
		JSONArray jsonArray = JSONArray.parseArray(msg);
		dadaGroupMsgService.saveGroupMessage(jsonArray, ssoPo, dadaGroupPo);
		wechat.send(dadaGroupPo, ssoPo);
		return JSON.toJson();
	}

}
