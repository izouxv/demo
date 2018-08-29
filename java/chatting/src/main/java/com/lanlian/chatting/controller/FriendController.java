package com.lanlian.chatting.controller;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONException;
import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.bo.DataBo;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.FriendInfoPo;
import com.lanlian.chatting.po.FriendsPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.FriendService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.UserFriends;

/**
 * @author wdyqxx
 * @version 2017年1月4日 下午8:12:03
 * @explain 此类为控制层，用于用户好友信息的操作；
 */
@Controller
@RequestMapping(value = RequestSetting.FRIENDS_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class FriendController extends MyAbstractController {

	@Resource
	FriendService friendsService;

	/**
	 * 获取好友信息；
	 * 
	 * @param uid
	 * @return String
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FRIENDS_BODY_ACQUIRE)
	public String acquireFriends(@RequestParam(value = "uid", required = true) int uid, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("acquireFriends-uid:" + uid);
		// 进行参数格式验证
		ParameterVerify.verifyUid(uid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		// 进入业务层
		FriendsPo friendsPo = new FriendsPo();
		friendsPo.setUid1(ssoPo.getUid());
		List<FriendInfoPo> list = friendsService.findFriends(friendsPo, ssoPo.getSource());
		return JSON.toJson(list);
	}

	/**
	 * 同步联系人
	 * 
	 * @param uid
	 * @param infos
	 * @return
	 * @throws Parameter_Exception
	 */
	@RequestMapping(value = RequestSetting.FRIENDS_BODY_SAVE)
	@ResponseBody
	public String syncFriends(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "infos", required = true) String infos, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("syncFriends-uid:" + uid + ",infos:" + infos);
		// 进行参数判空校验
		ParameterVerify.verifyNull(infos);
		// 进行参数格式校验
		ParameterVerify.verifyUid(uid);
		// 将数据放入到相应的对象；
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		UserFriends userFriends = parseBackupFriends(ssoPo.getUid(), infos);
		// 校验好友
		DataBo dataBo = friendsService.checkFriends(userFriends, ssoPo.getSource());
		// 添加好友
		friendsService.saveFriends(userFriends, dataBo);
		LogUtil.info(JSON.toJson(dataBo.getJson()));
		return JSON.toJson(dataBo.getJson());
	}

	/**
	 * 删除好友
	 * 
	 * @param uid
	 * @param uids
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FRIENDS_BODY_DELETE)
	public String deleteFriends(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "uids", required = true) String uids, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("deleteMyFriends-uid:" + uid + ",toUids:" + uids);
		// 进行参数格式校验
		ParameterVerify.verifyUid(uid);
		List<Integer> list = new ArrayList<>();
		try {
			JSONObject jsonObject = JSONObject.parseObject(uids);
			for (Object object : jsonObject.getJSONArray("uids")) {
				list.add((Integer) object);
			}
		} catch (Exception e) {
			throw new Parameter_Exception(21002);
		}
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		friendsService.deleteMyFriends(ssoPo.getUid(), list);
		return JSON.toJson();
	}

	/**
	 * 修改好友信息
	 * 
	 * @param uid
	 * @param infos
	 * @return
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FRIENDS_BODY_UPDATE)
	public String updateFriends(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "infos", required = true) String infos, HttpServletRequest request,
			HttpServletResponse response) {
		LogUtil.info("updateFriends-uid:" + uid + ",infos:" + infos);
		List<FriendsPo> list = new ArrayList<>();
		JSONObject jsonObject = new JSONObject();
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		try {
			jsonObject = JSONObject.parseObject(infos);
			FriendsPo friendsPo = null;
			for (String key : jsonObject.keySet()) {
				friendsPo = new FriendsPo();
				int toUid = Integer.parseInt(key);
				String note = jsonObject.getString(key);
				if (ssoPo.getUid() < toUid) {
					friendsPo.setUid1(ssoPo.getUid());
					friendsPo.setNote1(note);
					friendsPo.setUid2(toUid);
					list.add(friendsPo);
				}
				if (ssoPo.getUid() > toUid) {
					friendsPo.setUid1(toUid);
					friendsPo.setUid2(ssoPo.getUid());
					friendsPo.setNote2(note);
					list.add(friendsPo);
				}
			}
		} catch (JSONException e) {
			return JSON.toJson(21002);
		} catch (Exception e) {
			return JSON.toJson(21002);
		}
		friendsService.modifyFriends(list);
		return JSON.toJson();
	}

	/**
	 * ===========以下为私有方法===========================
	 */

	/**
	 * 将数据放入到相应的对象；
	 * 
	 * @param backup
	 * @return UserFriends
	 * @throws Parameter_Exception
	 */
	private UserFriends parseBackupFriends(int uid, String infos) throws Parameter_Exception {
		UserFriends userFriends = new UserFriends();
		Map<String, String> map = PublicMethod.json_map(infos);
		userFriends.setUid(uid);
		userFriends.setInfoMap(map);
		return userFriends;
	}

}
