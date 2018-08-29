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
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONArray;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Fatal_Exception;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.DadaGroupMsgService;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.server.http.Wechat;
import com.lanlian.server.redis.RedisClient79;

/**
 * @Title DadaGroupMsgController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午8:27:04
 * @explain 开启群信息实时上报
 */
@Controller
@RequestMapping(value = RequestSetting.GROUP_MESSAGE_PARENT_1_1, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class DadaGroupMsgController_1_1 extends MyAbstractController {

	@Resource
	DadaGroupMsgService dadaGroupMsgService;

	@Resource
	Wechat wechat;

	@Resource
	RedisClient79 redisClient79;

	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_MESSAGE_BODY_SYNC_GID)
	public String dadaGroupMsg(@PathVariable(value = "gid", required = true) int gid,
			@RequestParam(value = "msg", required = true) String msg, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception, Fatal_Exception {
		LogUtil.info("verify_code-gid:" + gid + ",msg:" + msg);
		// 获取上下文
		SsoPo ssoPo = getContextSsoPo(request, response);
		// 查询上报开关
		Integer state = redisClient79.getGidUidSwitchs(gid, ssoPo.getUid());
		if (state.intValue() == 0) {
			throw new Parameter_Exception(21005);
		}
		// 查询存储信息权限
		boolean flag = redisClient79.groupMsgUid(gid, ssoPo.getUid());
		if (flag) {
			DadaGroupPo dadaGroupPo = new DadaGroupPo();
			dadaGroupPo.setGid(gid);
			dadaGroupPo.setUid(ssoPo.getUid());
			JSONArray jsonArray = null;
			try {
				jsonArray = JSONArray.parseArray(msg);
			} catch (Exception e) {
				throw new Parameter_Exception(21002);
			}
			dadaGroupMsgService.saveGroupMessage(jsonArray, ssoPo, dadaGroupPo);
			wechat.send_1_1(dadaGroupPo, ssoPo);
		}
		return JSON.toJson();
	}
}
