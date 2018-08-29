package com.lanlian.chatting.controller;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.net.URLEncoder;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.GroupAndUserPO;
import com.lanlian.chatting.po.GroupInfoPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.po.virtual.TemporaryGroup;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.GroupSettingsService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.GroupInfo;
import com.lanlian.server.redis.RedisClient79;

@Controller
@RequestMapping(value = RequestSetting.GROUP_SETTINGS_PARENT_1_1, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class GroupSettingsController_1_1 extends MyAbstractController {

	private static Logger log = Logger.getLogger(FilesController.class);

	@Autowired
	GroupSettingsService groupSettingsService;

	@Resource
	RedisClient79 redisClient79;

	/**
	 * 群同步
	 * 
	 * @param uid
	 *            群主uid
	 * @param gid
	 *            群id
	 * @param gname
	 *            群名字
	 * @param members
	 *            群成员id
	 * @param avatar
	 *            头像
	 * @param announcement
	 *            群公告
	 * @param longitude
	 *            经度
	 * @param latitude
	 *            纬度
	 * @param createTime
	 *            记录临时群创建的时间
	 * @return 状态码
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_SYNC)
	public String groupSync(@RequestParam(value = "uid", required = true) Integer uid,
			@RequestParam(value = "gid", required = true) Integer gid,
			@RequestParam(value = "gname", required = true) String gname,
			@RequestParam(value = "avatar", required = true) Integer avatar,
			@RequestParam(value = "uids", required = true) String uids,
			@RequestParam(value = "announcement", required = true) String announcement,
			@RequestParam(value = "longitude", required = true) Double longitude,
			@RequestParam(value = "latitude", required = true) Double latitude,
			@RequestParam(value = "createTime", required = true) Long createTime, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {

		log.info("groupSync-gid:" + gid + ",uid:" + uid + ",gname：" + gname + ",uids:" + uids + ",announcement:"
				+ announcement + ",longitude:" + longitude + ",latitude：" + latitude + ",createTime:" + createTime
				+ ",avatar：" + avatar);
		// 进行参数验证
		ParameterVerify.verifyNull(gname, announcement);
		ParameterVerify.verifyNull(gname);
		ParameterVerify.verifyUid(uid);
		ParameterVerify.verifyGid(gid);
		ParameterVerify.verifyCoord(latitude, longitude);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		try {
			gname = URLEncoder.encode(gname.trim(), "UTF-8");
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(21002);
		}
		Timestamp timestamp = new Timestamp(createTime);
		List<Integer> list = conversion(uids);
		list.add(ssoPo.getUid());
		HashSet<Integer> hashSet = new HashSet<>(list);
		list.clear();
		list.addAll(hashSet);
		// 查询群是否存在
		GroupInfoPo groupInfoPo = new GroupInfoPo();
		groupInfoPo.setGid(gid);// 群id
		groupInfoPo.setUid(ssoPo.getUid());
		groupInfoPo = groupSettingsService.findGroupInfo(groupInfoPo);
		log.info("groupInfoPo:" + groupInfoPo);
		// 群不存在
		if (groupInfoPo == null) {
			// 判断redis中的uid未激活群中是否包含gid
			Boolean bool = redisClient79.isExistAndSismember(redisClient79.key9+uid,gid);
			if (!bool) {
				LogUtil.info("redis库中uid = " + uid + " 的未激活群中不包含此gid = " + gid);
				throw new Parameter_Exception(20030);
			}
			GroupInfoPo groupInfoPo1 = new GroupInfoPo();
			groupInfoPo1.setUid(uid);// 群主id
			groupInfoPo1.setGid(gid);// 群id
			groupInfoPo1.setGname(gname);// 群名字
			groupInfoPo1.setAvatar(avatar);// 头像
			groupInfoPo1.setAnnouncement(announcement);// 群公告
			groupInfoPo1.setLongitude(longitude);// 经度
			groupInfoPo1.setLatitude(latitude);// 纬度
			groupInfoPo1.setCreatTime(timestamp);// 临时群创建的时间
			// 创建群并存入成员关系,创建群上报开关0
			groupSettingsService.saveGroupAndUser(groupInfoPo1, list, ssoPo.getSource());
			try {
				groupInfoPo1.setGname(URLDecoder.decode(gname, "UTF-8"));
			} catch (UnsupportedEncodingException e) {
				throw new Parameter_Exception(21002);
			}
			return JSON.toJson(groupInfoPo1);
		}
		// 群存在，更新群
		if (groupInfoPo.getUid() != uid) {
			// 不是群主，还接口
			throw new Parameter_Exception(21005);
		}
		groupInfoPo.setGname(gname);// 群名字
		groupInfoPo.setAnnouncement(announcement);// 群公告
		groupInfoPo.setLongitude(longitude);// 经度
		groupInfoPo.setLatitude(latitude);// 纬度
		groupInfoPo.setCreatTime(timestamp);// 临时群创建的时间
		groupInfoPo.setAvatar(avatar);// 头像
		groupInfoPo = groupSettingsService.syncGroup(groupInfoPo, list, ssoPo.getSource());
		try {
			groupInfoPo.setGname(URLDecoder.decode(gname, "UTF-8"));
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(21002);
		}
		return JSON.toJson(groupInfoPo);
	}

	/**
	 * 将uids转换为json
	 * 
	 * @param uids
	 * @return
	 * @throws Parameter_Exception
	 */
	private List<Integer> conversion(String uids) throws Parameter_Exception {
		List<Integer> list = new ArrayList<>();
		JSONObject jsonObject = new JSONObject();
		try {
			jsonObject = JSONObject.parseObject(uids);
			for (Object object : jsonObject.getJSONArray("uids")) {
				list.add((Integer) object);
			}
		} catch (Exception e) {
			throw new Parameter_Exception(21002);
		}
		return list;
	}

	/**
	 * 踢人出群
	 * 
	 * @param gid
	 *            群id
	 * @param uids
	 *            用户id数组
	 * @param uid
	 *            群主id
	 * @return 状态码
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_KICKMEMBER)
	public String kickMember(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "gid", required = true) int gid,
			@RequestParam(value = "uids", required = true) String uids, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("kickMember-uid:" + uid + ",gid:" + gid + ",uids:" + uids);
		// 进行参数验证
		ParameterVerify.verifyUid(uid);
		ParameterVerify.verifyGid(gid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		List<Integer> list = conversion(uids);
		GroupInfoPo groupInfoPo = new GroupInfoPo();
		groupInfoPo.setGid(gid);
		groupInfoPo.setUid(ssoPo.getUid());
		groupInfoPo = groupSettingsService.findGroupInfo(groupInfoPo);
		log.info("groupInfoPo:" + groupInfoPo);
		if (groupInfoPo == null) {
			return JSON.toJson(21005);
		}
		groupSettingsService.deleteMembers(groupInfoPo, list, ssoPo.getSource());
		return JSON.toJson();
	}

	/**
	 * 群解散
	 * 
	 * @param gid
	 *            群id
	 * @param uid
	 *            群主id
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_DISSOLVE)
	public String groupDissolve(@RequestParam(value = "gid", required = true) int gid,
			@RequestParam(value = "uid", required = true) int uid, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("groupDissolve-uid:" + uid + ",gid:" + gid);
		// 进行参数校验
		ParameterVerify.verifyUid(uid);
		ParameterVerify.verifyGid(gid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		GroupInfoPo groupInfoPo = new GroupInfoPo();
		groupInfoPo.setGid(gid);
		groupInfoPo.setUid(ssoPo.getUid());
		groupInfoPo = groupSettingsService.findGroupInfo(groupInfoPo);
		if (groupInfoPo == null) {
			return JSON.toJson(21005);
		}
		if (groupInfoPo.getUid() != ssoPo.getUid()) {
			throw new Parameter_Exception(21005);
		}
		groupInfoPo.setDataState(2);
		groupInfoPo.setDataModifyTime(getTime());
		groupSettingsService.deleteGroup(groupInfoPo);
		return JSON.toJson();
	}

	/**
	 * 编辑群资料
	 * 
	 * @param tgname
	 *            群名称
	 * @param tgid
	 *            群id
	 * @param pid
	 *            群主id
	 * @param notice
	 *            群公告
	 * @param icon
	 *            群头像
	 * @return 状态码
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_EDIT)
	public String groupEdit(@RequestParam(value = "gid", required = true) int gid,
			@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "gname", required = true) String gname,
			@RequestParam(value = "avatar", required = true) int avatar,
			@RequestParam(value = "announcement", required = true) String announcement, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("groupEdit-uid:" + uid + ",gid:" + gid + ",gname:" + gname + ",avatar:" + avatar + ",announcement:"
				+ announcement);
		// 进行参数校验
		ParameterVerify.verifyNull(gname);
		ParameterVerify.verifyUid(uid);
		ParameterVerify.verifyGid(gid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		try {
			gname = URLEncoder.encode(gname.trim(), "UTF-8");
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(21002);
		}
		GroupInfoPo groupInfoPo = new GroupInfoPo();
		groupInfoPo.setGid(gid);
		groupInfoPo.setUid(ssoPo.getUid());
		groupInfoPo = groupSettingsService.findGroupInfo(groupInfoPo);
		if (groupInfoPo == null) {
			throw new Parameter_Exception(21002);
		}
		if (groupInfoPo.getUid() != ssoPo.getUid()) {
			throw new Parameter_Exception(21005);
		}
		groupInfoPo.setGname(gname);
		groupInfoPo.setAnnouncement(announcement);
		groupInfoPo.setAvatar(avatar);
		groupSettingsService.updateGroupInfo(groupInfoPo);
		return JSON.toJson();
	}

	/**
	 * 用户主动退群
	 * 
	 * @param gid
	 *            群id
	 * @param uid
	 *            用户id
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_EXIT)
	public String exitGroup(@RequestParam(value = "gid", required = true) int gid,
			@RequestParam(value = "uid", required = true) int uid, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("exitGroup-uid:" + uid + "gid:" + gid);
		// 进行参数校验
		ParameterVerify.verifyGid(gid);
		ParameterVerify.verifyUid(uid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		GroupAndUserPO groupAndUserPO = new GroupAndUserPO();
		groupAndUserPO.setGid(gid);
		groupAndUserPO.setUid(ssoPo.getUid());
		groupAndUserPO = groupSettingsService.getGroupUser(groupAndUserPO);
		if (groupAndUserPO == null) {
			return JSON.toJson(21005);
		}
		List<Integer> list = new ArrayList<>();
		list.add(ssoPo.getUid());
		GroupInfoPo groupInfoPo = new GroupInfoPo();
		groupInfoPo.setGid(gid);
		groupSettingsService.deleteMembers(groupInfoPo, list, ssoPo.getSource());
		return JSON.toJson();
	}

	/**
	 * 根据uid查询永久群信息列表
	 * 
	 * @param uid
	 * @return 永久群信息以及所有群成员信息
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_LIST)
	public String findGroupByUidList(@RequestParam(value = "uid", required = true) int uid, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("findGroupByUidList-uid:" + uid);
		// 进行参数校验
		ParameterVerify.verifyUid(uid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		GroupInfoPo groupInfoPo = new GroupInfoPo();
		groupInfoPo.setUid(ssoPo.getUid());
		JSONArray jsonArray = groupSettingsService.findGroupByUidList(groupInfoPo, ssoPo.getSource());
		return JSON.toJson(jsonArray);
	}

	/**
	 * 根据gid查询永久群信息列表
	 * 
	 * @param gids
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_INFO)
	public String findGroupByGidList(HttpServletRequest request, HttpServletResponse response,
			@RequestParam(value = "gid", required = true) int... gids) throws Parameter_Exception {
		LogUtil.info("findGroupByGidList-gids:" + Arrays.toString(gids));
		// 参数校验
		ParameterVerify.verifyGid(gids);
		SsoPo ssoPo = getContextSsoPo(request, response);
		List<GroupInfo> list = new ArrayList<>();
		GroupInfo groupInfo = null;
		for (int gid : gids) {
			groupInfo = new GroupInfo();
			groupInfo.setGid(gid);
			list.add(groupInfo);
		}
		JSONArray jsonArray = groupSettingsService.findGroupByGidList(list, ssoPo.getSource());
		return JSON.toJson(jsonArray);
	}

	/**
	 * 信息上报、临时群备份
	 * 
	 * @param types
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.GROUP_SETTINGS_BODY_TEMPORARY)
	public String temporaryGroup(@PathVariable(value = "types", required = true) String types,
			@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "data", required = true) String data, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("localtion-types:" + types + ":uid," + uid + ",data:" + data);
		TemporaryGroup temporaryGroup = new TemporaryGroup();
		try {
			JSONArray jsonArray = JSONArray.parseArray(data);
			temporaryGroup.setData(jsonArray);
		} catch (Exception e) {
			throw new Parameter_Exception(21002);
		}
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		temporaryGroup.setUid(ssoPo.getUid());
		temporaryGroup.setTime(getTimes());
		switch (types) {
		// 临时群备份
		case "sync":
			groupSettingsService.temporaryGroup(types, temporaryGroup);
			return JSON.toJson();
		// 临时群备份
		case "get":
			groupSettingsService.temporaryGroup(types, temporaryGroup);
			return JSON.toJson(temporaryGroup);
		default:
			return JSON.toJson("别逗！！！错了");
		}
	}

}
