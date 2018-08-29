/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.controller;

import java.util.ArrayList;
import java.util.List;

import javax.annotation.Resource;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.virtual.NearbyPo;
import com.lanlian.chatting.po.virtual.NearbyGroupInfoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.NearbyService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月19日 下午5:39:03
 * @explain
 */
@Controller
@RequestMapping(value = RequestSetting.VISITORS_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class NearbyController extends MyAbstractController {

	@Resource
	NearbyService nearbyService;

	/**
	 * 回传位置信息与获取tid
	 * 
	 * @param action
	 * @param imei
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.NEARBY_BODY)
	public String localtion(@PathVariable(value = "types", required = true) String types,
			@PathVariable(value = "imei", required = true) String imei,
			@RequestParam(value = "channel", required = true) String channel) throws Parameter_Exception {
		LogUtil.info("localtion-types:" + types + ",imei:" + imei + ",channel:" + channel);
		// 声明变量
		Integer tuid = 0;// 用户id
		String nickname = null;// 昵称
		Integer avatar = 0;// 头像id
		Integer gender = 0;// 性别1:男/2:女/0:无
		Integer age = 0;// 年龄
		String signature = null;// 个性签名
		Double longitude = 0.0;
		Double latitude = 0.0;
		try {
			JSONObject jsonObject = JSONObject.parseObject(channel);
			tuid = jsonObject.getInteger("tuid");
			nickname = jsonObject.getString("nickname");
			avatar = jsonObject.getInteger("avatar");
			gender = jsonObject.getInteger("gender");
			age = jsonObject.getInteger("age");
			signature = jsonObject.getString("signature");
			longitude = jsonObject.getDouble("longitude");// 经度[-180,180]
			latitude = jsonObject.getDouble("latitude");// 纬度[-90,90]
		} catch (Exception e) {
			throw new Parameter_Exception(21002);
		}
		ParameterVerify.verifyNull(nickname);
		ParameterVerify.verifyCoord(longitude, latitude);
		// 在localtion中存值
		NearbyPo nearbyPo = new NearbyPo();
		nearbyPo.setTuid(tuid);
		nearbyPo.setNickname(nickname);
		nearbyPo.setAvatar(avatar);
		nearbyPo.setGender(gender);
		nearbyPo.setAge(age);
		nearbyPo.setSignature(signature);
		nearbyPo.setImei(imei);
		nearbyPo.setLongitude(longitude);
		nearbyPo.setLatitude(latitude);
		switch (types) {
		// 获取临时id
		case "get":
			ParameterVerify.verifyImei(imei);
			nearbyService.getTid(nearbyPo);
			List<NearbyPo> list = nearbyService.nearbyLocaltion(nearbyPo);
			nearbyPo.setNearbyPos(list);
			LogUtil.info("nearbyLocaltion-types:" + types + ",nearbyPo:" + nearbyPo);
			return JSON.toJson(nearbyPo);
		// 更新位置
		case "update":
			ParameterVerify.verifyIntegerPositive(tuid);
			List<NearbyPo> listup = nearbyService.nearbyLocaltion(nearbyPo);
			nearbyPo.setNearbyPos(listup);
			LogUtil.info("nearbyLocaltion-types:" + types + ",nearbyPo:" + nearbyPo);
			return JSON.toJson(nearbyPo);
		// 清除位置信息并关闭虚拟信道
		case "delete":
			nearbyService.deleteNearbyLocaltion(nearbyPo);
			LogUtil.info("nearbyLocaltion-types:" + types + "nearbyPo:" + nearbyPo);
			return JSON.toJson();
		default:
			break;
		}
		return null;
	}

	/**
	 * 群操作
	 * 
	 * @param types
	 * @param channel
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.NEARBY_LOCALTION_BODY)
	public String nearbyLocaltion(@PathVariable(value = "types", required = true) int types,
			@RequestParam(value = "channel", required = true) String channel) throws Parameter_Exception {
		LogUtil.info("nearbyLocaltion-action:" + types + ",channel:" + channel);
		// 声明变量
		Integer tuid = 0;// 用户id
		String nickname = null;// 昵称
		Integer avatar = 0;// 头像id
		Integer gender = 0;// 性别1:男/2:女/0:无
		Integer age = 0;// 年龄
		String signature = null;// 个性签名
		Double longitude = 0.0;// 经度
		Double latitude = 0.0;// 纬度
		Integer tgid = 0;// 临时群id
		Integer ownerid = 0;// 群拥有者id
		String notice = null;// 群通知
		List<Integer> tuids = new ArrayList<>();
		try {
			JSONObject jsonObject = JSONObject.parseObject(channel);
			tuid = jsonObject.getInteger("tuid");// 自己的id
			nickname = jsonObject.getString("nickname");
			avatar = jsonObject.getInteger("avatar");
			gender = jsonObject.getInteger("gender");
			age = jsonObject.getInteger("age");
			signature = jsonObject.getString("signature");
			longitude = jsonObject.getDouble("longitude");// 经度[-180,180]
			latitude = jsonObject.getDouble("latitude");// 纬度[-90,90]

			tgid = jsonObject.getInteger("tgid");// 临时群id
			ownerid = jsonObject.getInteger("ownerid");// 临时群主id
			notice = jsonObject.getString("notice");
			JSONArray jsonArray = jsonObject.getJSONArray("tuids");// 对方的id
			for (Object object : jsonArray) {
				tuids.add((Integer) object);
			}
		} catch (Exception e) {
			throw new Parameter_Exception(21002);
		}
		ParameterVerify.verifyNull(nickname);
		ParameterVerify.verifyIntegerPositive(tuid, tgid, avatar, gender, age);
		ParameterVerify.verifyIntegerPositive(tuids.toArray(new Integer[tuids.size()]));
		ParameterVerify.verifyCoord(longitude, latitude);
		// 在nearbyPo中存值
		NearbyPo nearbyPo = new NearbyPo();
		nearbyPo.setTuid(tuid);
		nearbyPo.setNickname(nickname);
		nearbyPo.setAvatar(avatar);
		nearbyPo.setGender(gender);
		nearbyPo.setAge(age);
		nearbyPo.setSignature(signature);
		nearbyPo.setLongitude(longitude);
		nearbyPo.setLatitude(latitude);
		nearbyPo.setTgid(tgid);
		nearbyPo.setOwnerid(ownerid);
		nearbyPo.setNotice(notice);
		nearbyPo.setTuids(tuids);
		LogUtil.info("nearbyLocaltion-nearbyPo:" + nearbyPo);
		// 根据操作类型判断
		NearbyGroupInfoPo nearbyGroupInfoPo;
		switch (types) {
		// 获取群id
		case 1:
			nearbyService.getTgid(nearbyPo);
			return JSON.toJson(nearbyPo);
		// 邀请入群
		case 2:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		// 响应邀请入群
		case 3:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson(nearbyPo);
		// 请求加群
		case 4:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		// 响应请求加群
		case 5:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		// 群主踢人出群
		case 6:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		// 成员退群
		case 7:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		// 群主解散群
		case 8:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		// 编辑群信息
		case 9:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		// 拉取群资料
		case 10:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson(nearbyPo);
		default:
			nearbyGroupInfoPo = new NearbyGroupInfoPo();
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		}

	}

	/**
	 * 对附近的人发送群操作或聊天信息
	 * 
	 * @param imei
	 * @param toimei
	 * @param type
	 * @param info
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.NEARBY_SEND_BODY)
	public String sendNearby(@PathVariable(value = "types", required = true) int types,
			@PathVariable(value = "tuid", required = true) int tuid,
			@PathVariable(value = "toid", required = true) int toid,
			@RequestParam(value = "type", required = true) int type,
			@RequestParam(value = "info", required = true) String info) throws Parameter_Exception {

		LogUtil.info("sendNearby-tuid:" + tuid + ",toid:" + toid + ",type:" + type + ",info:" + info);
		// 进行参数判空校验
		ParameterVerify.verifyNull(info);
		// 进行参数格式校验
		ParameterVerify.verifyIntegerPositive(tuid, toid, type);

		NearbyPo nearbyPo = new NearbyPo();
		nearbyPo.setTgid(toid);
		NearbyGroupInfoPo nearbyGroupInfoPo = new NearbyGroupInfoPo();
		nearbyGroupInfoPo.setTypes(types);
		nearbyGroupInfoPo.setTuid(tuid);
		nearbyGroupInfoPo.setTime(getTimes());
		nearbyGroupInfoPo.setType(type);
		nearbyGroupInfoPo.setInfo(info);
		switch (types) {
		case 11:
			nearbyService.sendNearbyInfo(nearbyGroupInfoPo, nearbyPo, types);
			return JSON.toJson();
		default:
			break;
		}
		return JSON.toJson("别逗！！！没成功");
	}

	/**
	 * 接收自己的信息箱
	 * 
	 * @param imei
	 * @param toimei
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.NEARBY_GET_BODY)
	public String getNearby(@PathVariable(value = "tuid", required = true) int tuid) throws Parameter_Exception {

		LogUtil.info("getNearby-tuid:" + tuid);
		// 进行参数格式校验
		ParameterVerify.verifyIntegerPositive(tuid);
		NearbyGroupInfoPo nearbyPo = new NearbyGroupInfoPo();
		nearbyPo.setTuid(tuid);
		List<NearbyGroupInfoPo> list = nearbyService.getNearbyInfo(nearbyPo);
		return JSON.toJson(list);
	}

}