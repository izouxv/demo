package com.lanlian.chatting.controller;

import java.io.UnsupportedEncodingException;
import java.util.List;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.bo.InviteBo;
import com.lanlian.chatting.bo.MessagePageBo;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.WeChatGroupSettingsService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.BundlingGroupVo;
import com.lanlian.chatting.vo.MessageInfoPage;
import com.lanlian.server.redis.RedisClient79;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年6月19日 上午10:24:38
 * @explain
 */
@Controller
@RequestMapping(value = RequestSetting.GROUP_WECHAT_PARENT_1_1)
public class WeChatController_1_1 extends MyAbstractController {

	@Autowired
	WeChatGroupSettingsService weChatGroupSettingsService;

	@Resource
	RedisClient79 redisClient;

	/**
	 * 取消订阅群信息
	 * 
	 * @param gid
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_WECHAT_BODY_UNSUB, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
	public String groupSettings(@PathVariable(value = "gid", required = true) int gid) throws Parameter_Exception {
		// 进行参数验证
		ParameterVerify.verifyGid(gid);
		// 取消订阅群信息
		redisClient.deleteSubscribeGroupId(gid);
		return JSON.toJson();
	}

	/**
	 * 订阅群信息
	 * 
	 * @param inviteCode
	 *            邀请码
	 * @return
	 * @throws Parameter_Exception
	 * @throws UnsupportedEncodingException
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_WECHAT_BODY_SUB, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
	public String subscribe(@RequestParam(value = "inviteCode", required = true) String inviteCode,
			HttpServletRequest request, HttpServletResponse response)
			throws Parameter_Exception, UnsupportedEncodingException {
		SsoPo ssoPo = getContextSsoPo(request, response);
		// 订阅群信息
		InviteBo inviteBo = new InviteBo(inviteCode, ssoPo.getUid());
		BundlingGroupVo bundling = weChatGroupSettingsService.bundling(inviteBo);
		redisClient.saveSubscribeGroupId(bundling.getGid());
		return JSON.toJson(bundling);
	}

	/**
	 * 查询实时群上报信息
	 * 
	 * @param gid
	 * @param startid
	 * @param endid
	 * @param count
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_WECHAT_BODY_MESSAGE, consumes = RequestSetting.CONSUMES, method = RequestMethod.GET, produces = RequestSetting.PRODUCES)
	public String getMessage(@RequestParam(value = "gid", required = true) int gid,
			@RequestParam(value = "startid", required = false) int startid,
			@RequestParam(value = "endid", required = false) int endid,
			@RequestParam(value = "count", required = true) int count, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		// 参数校验，count值不能设置太大
		LogUtil.info("getMessage-gid:" + gid + ",startid:" + startid + ",endid:" + endid + ",count:" + count);
		SsoPo ssoPo = getContextSsoPo(request, response);
		MessagePageBo messagePageBo = new MessagePageBo(gid, startid, endid, count);
		messagePageBo.setGid(gid);
		messagePageBo.setCount(count);

		List<MessageInfoPage> messages = weChatGroupSettingsService.getMessage(messagePageBo, ssoPo.getSource());
		return JSON.toJson(messages);
	}
}
