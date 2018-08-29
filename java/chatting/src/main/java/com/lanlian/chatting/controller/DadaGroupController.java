/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.controller;

import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.List;

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

import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.model.GroupSwitching;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.DadaGroupService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @Title DadaGroupController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午8:27:04
 * @explain 开启群信息实时上报
 */
@Controller
@RequestMapping(value = RequestSetting.GROUP_SETTINGS_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class DadaGroupController extends MyAbstractController {

	@Resource
	@Qualifier("dadaGroupServiceImpl")
	DadaGroupService dadaGroupService;

	/**
	 * 开启实时上传群
	 * 
	 * @param gid
	 * @param state
	 * @param gName
	 * @param avatar
	 * @param announcement
	 * @param creatTime
	 * @param members
	 * @param longitude
	 * @param latitude
	 * @param request
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 * @throws UnsupportedEncodingException
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_SYNC_GID)
	public String dadaGroup(@PathVariable(value = "gid", required = true) int gid,
			@RequestParam(value = "state", required = true) Integer state,
			@RequestParam(value = "gName", required = true) String gName,
			@RequestParam(value = "avatar", required = true) Integer avatar,
			@RequestParam(value = "announcement", required = true) String announcement,
			@RequestParam(value = "creatTime", required = true) long creatTime,
			@RequestParam(value = "members", required = true) String members,
			@RequestParam(value = "longitude", required = true) Double longitude,
			@RequestParam(value = "latitude", required = true) Double latitude, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("dadaGroup-gid:" + gid + ",gName:" + gName + ",avatar:" + avatar + ",announcement:" + announcement
				+ ",creatTime:" + creatTime + ",members:" + members + ",longitude:" + longitude + ",latitude:"
				+ latitude + ",state:" + state);
		// 获取上下文
		SsoPo ssoPo = getContextSsoPo(request, response);
		//
		ParameterVerify.verifyNull(gName);
		DadaGroupPo dadaGroupPo = new DadaGroupPo();
		dadaGroupPo.setGid(gid);
		dadaGroupPo.setUid(ssoPo.getUid());
		// 查询群
		dadaGroupPo = dadaGroupService.findDadaGroup(dadaGroupPo);
		// System.out.println("gid:" + gid + "," + dadaGroupPo);
		// 该群不存在
		if (dadaGroupPo == null) {
			DadaGroupPo dadaGroupPoNew = new DadaGroupPo();
			dadaGroupPoNew.setGid(gid);
			dadaGroupPoNew.setUid(ssoPo.getUid());
			dadaGroupPoNew.setGname(gName);
			dadaGroupPoNew.setAvatar(avatar);
			dadaGroupPoNew.setAnnouncement(announcement);
			dadaGroupPoNew.setCreateTime(new Timestamp(creatTime));
			dadaGroupPoNew.setLongitude(longitude);
			dadaGroupPoNew.setLatitude(latitude);
			dadaGroupPoNew.setGroupState(state);
			dadaGroupPoNew.setDataCreateTime(getTime());
			dadaGroupPoNew.setModifyTime(getTime());
			dadaGroupPoNew.setDataState(1);
			// members转为jsonarray
			List<Integer> list = new ArrayList<>();
			JSONObject jsonObject = JSONObject.parseObject(members);
			for (Object object : jsonObject.getJSONArray("uids")) {
				list.add((Integer) object);
			}
			// list.add(ssoPo.getUid());
			// 创建群并开启上报开关
			dadaGroupPoNew = dadaGroupService.saveDadaGroup(dadaGroupPoNew, list, ssoPo.getSource());
			// 创建邀请码并存储
			dadaGroupPoNew.setInviteCode(getRandomCode(dadaGroupPoNew.getGid()));
			dadaGroupService.saveDadaGroupCode(dadaGroupPoNew);
			// 创建返回模型
			GroupSwitching groupInfo = new GroupSwitching();
			groupInfo.setGid(dadaGroupPoNew.getGid());
			groupInfo.setInviteCode(dadaGroupPoNew.getInviteCode());
			groupInfo.setGroupState(dadaGroupPoNew.getGroupState());
			return JSON.toJson(groupInfo);
		}
		// 群存在,关闭或开启
		switch (state) {
		case 0:
			if (dadaGroupPo.getUpid() != ssoPo.getUid()) {
				throw new Parameter_Exception(21005);
			}
			actionGroupState(state, dadaGroupPo);
			return JSON.toJson();
		case 1:
			if (dadaGroupPo.getUpid() != ssoPo.getUid() && dadaGroupPo.getGroupState() == 0) {
				dadaGroupPo.setUpid(ssoPo.getUid());
				dadaGroupService.saveDadaGroupSwitching(dadaGroupPo);
				// 创建返回模型
				GroupSwitching groupInfo = new GroupSwitching();
				groupInfo.setGid(dadaGroupPo.getGid());
				groupInfo.setInviteCode(dadaGroupPo.getInviteCode());
				groupInfo.setGroupState(dadaGroupPo.getGroupState());
				return JSON.toJson(groupInfo);
			}
			dadaGroupPo = actionGroupState(state, dadaGroupPo);
			// 创建返回模型
			GroupSwitching groupInfo = new GroupSwitching();
			groupInfo.setGid(dadaGroupPo.getGid());
			groupInfo.setInviteCode(dadaGroupPo.getInviteCode());
			groupInfo.setGroupState(dadaGroupPo.getGroupState());
			return JSON.toJson(groupInfo);
		default:
			response.setStatus(404);
			return JSON.toJson("别逗........");
		}
	}

	/**
	 * 根据action修改群开关或返回异常
	 * 
	 * @param state
	 * @param dadaGroupPo
	 * @throws Parameter_Exception
	 */
	private DadaGroupPo actionGroupState(Integer state, DadaGroupPo dadaGroupPo) throws Parameter_Exception {
		// // 判断action冲突
		// if (dadaGroupPo.getGroupState() == state) {
		// throw new Parameter_Exception(21005);
		// }
		// 修改上报开关
		dadaGroupPo.setGroupState(state);
		return dadaGroupService.updateDadaGroup(dadaGroupPo);
	}

}
