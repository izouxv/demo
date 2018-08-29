
package com.lanlian.chatting.controller;

import java.util.ArrayList;
import java.util.List;

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
import com.lanlian.chatting.po.GroupMessageBackupPO;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.GroupMessageService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.GroupMessageInfo;

@Controller
@RequestMapping(value = RequestSetting.GROUP_MESSAGE_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class GroupMessageController extends MyAbstractController {

	@Resource
	private GroupMessageService messageService;

	/**
	 * 用户聊天消息上传
	 * 
	 * @param message
	 *            消息集合message
	 * @param uid
	 *            用户的id
	 * @param tgid
	 *            永久群id
	 * @param createTime
	 *            聊天开始时间
	 * @param longitude
	 *            开始聊天的经度
	 * @param latitude
	 *            开始聊天的纬度
	 * @param address
	 *            聊天的地址
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_MESSAGE_BODY_SYNC)
	public String messageSync(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "gid", required = true) int gid,
			@RequestParam(value = "createTime", required = true) long createTime,
			@RequestParam(value = "msg", required = true) String msg,
			@RequestParam(value = "fileId", required = false) String fileId,
			@RequestParam(value = "fileName", required = false) String fileName,
			@RequestParam(value = "longitude", required = true) double longitude,
			@RequestParam(value = "latitude", required = true) double latitude,
			@RequestParam(value = "address", required = true) String address, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {

		LogUtil.info("请求参数-msg：" + msg + ",fileId" + fileId + ",uid：" + uid + ",gid：" + gid + ",longitude：" + longitude
				+ ",latitude：" + latitude + ",address:" + address);

		// if (fileId.length() != 24) {
		// return JSON.toJson(21002);
		// }

		ParameterVerify.verifyNull(msg);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		GroupMessageBackupPO gmbpo = new GroupMessageBackupPO();
		GroupMessageInfo messageinfo = new GroupMessageInfo();
		messageinfo.setMessageInfo(msg);
		messageService.messageList(messageinfo);
		gmbpo.setUid(ssoPo.getUid());
		gmbpo.setGid(gid);
		gmbpo.setMessageId(messageinfo.getMessageInfoId());
		gmbpo.setFileId(fileId);
		gmbpo.setFileName(fileName);
		gmbpo.setLatitude(latitude);
		gmbpo.setLongitude(longitude);
		gmbpo.setAddress(address);
		gmbpo.setCreateTime(getTime());
		gmbpo.setBackupTime(getTime());
		gmbpo.setModifyTime(getTime());
		messageService.messagesUpload(gmbpo);
		return JSON.toJson();
	}

	/**
	 * 获取聊天记录列表
	 * 
	 * @param uid
	 *            用户uid
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_MESSAGE_BODY_LIST)
	public String groupChatList(@RequestParam(value = "uid", required = true) int uid, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("groupChatList-uid:" + uid);
		ParameterVerify.verifyUid(uid);
		List<GroupMessageBackupPO> list = new ArrayList<>();
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		list = messageService.theMessageList(ssoPo.getUid());
		return JSON.toJson(list);
	}

	/**
	 * 以id获取一条群聊记录的内容
	 * 
	 * @param id
	 *            消息的id
	 * @return 状态码
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_MESSAGE_BODY_INFO)
	public String messageInfo(@PathVariable(value = "id", required = true) int id, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		if (id <= 0) {
			// 参数错误
			throw new Parameter_Exception(21002);
		}
		SsoPo ssoPo = getContextSsoPo(request, response);
		GroupMessageInfo info = new GroupMessageInfo();
		info = messageService.messageUpdown(ssoPo.getUid(), id);
		if (info == null) {
			return JSON.toJson(21002);
		}
		JSONArray jsonArray = JSONArray.parseArray(info.getMessageInfo());
		return JSON.toJson(jsonArray);
	}

}