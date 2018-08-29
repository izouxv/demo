/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.serviceImpl;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;

import javax.annotation.Resource;

import org.springframework.stereotype.Service;

import com.lanlian.chatting.controller.MyAbstractController;
import com.lanlian.chatting.localtion.DistanceUtil;
import com.lanlian.chatting.localtion.GeoHash;
import com.lanlian.chatting.po.virtual.NearbyPo;
import com.lanlian.chatting.po.virtual.NearbyGroupInfoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.NearbyService;
import com.lanlian.server.redis.RedisClient79;
import com.lanlian.server.redis.RedisClient80;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月19日 下午7:25:58
 * @explain
 */
@Service(value = "nearbyServiceImpl")
public class NearbyServiceImpl extends MyAbstractController implements NearbyService {

	@Resource
	RedisClient79 redisClient79;

	@Resource
	RedisClient80 redisClient80;

	/**
	 * 返回一个临时用户id
	 * 
	 * @param nearbyPo
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public void getTid(NearbyPo nearbyPo) throws Parameter_Exception {
		Integer tid = redisClient79.getTid();
		nearbyPo.setTuid(tid);
	}

	/**
	 * 生成一个临时群id
	 * 
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public void getTgid(NearbyPo nearbyPo) throws Parameter_Exception {
		Integer tgid = redisClient79.getGid();
		nearbyPo.setTgid(tgid);
		nearbyPo.setOwnerid(nearbyPo.getTuid());
		redisClient79.setGroupId(nearbyPo);
		List<Integer> list = new ArrayList<>();
		list.add(nearbyPo.getTuid());
		nearbyPo.setTuids(list);
		redisClient79.setGroupMember(nearbyPo);
	}

	/**
	 * 存入用户信息并获取附近人的信息
	 * 
	 * @param nearbyPo
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public List<NearbyPo> nearbyLocaltion(NearbyPo nearbyPo) throws Parameter_Exception {
		GeoHash geoHash = new GeoHash(nearbyPo.getLongitude(), nearbyPo.getLatitude());
		geoHash.sethashLength(5);
		String geoHash5 = geoHash.getGeoHashBase32();
		nearbyPo.setGeoHash(geoHash5);
		// 存储用户当前信息
		redisClient80.saveNearbyToKey(nearbyPo);
		redisClient80.saveNearbyToSets(nearbyPo);
		// 获取用户附近的信息
		List<String> list = geoHash.getGeoHashBase32For9();
		String[] geohashs = list.toArray(new String[list.size()]);
		Set<String> set = redisClient80.findNearbyToSets(geohashs);
		list.clear();
		list.addAll(set);
		List<NearbyPo> list2 = redisClient80.findNearbyToKey(list);
		List<NearbyPo> list3 = new ArrayList<>();
		for (NearbyPo nearby : list2) {
			if (nearbyPo.getTuid().equals(nearby.getTuid())) {
				continue;
			}
			double distance = DistanceUtil.getDistance(nearbyPo.getLongitude(), nearbyPo.getLatitude(),
					nearby.getLongitude(), nearby.getLatitude());
			if (distance < 30000) {
				nearby.setDistance(distance);
				list3.add(nearby);
			}
		}
		return list3;
	}

	/**
	 * 关闭用户信息并获取附近人的信息
	 * 
	 * @param localtionPo
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public void deleteNearbyLocaltion(NearbyPo nearbyPo) throws Parameter_Exception {
		redisClient80.deleteNearbyToKey(nearbyPo.getTuid());
	}

	/**
	 * 对附近的人发送群操作消息
	 * 
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	@Override
	public NearbyPo sendNearbyInfo(NearbyGroupInfoPo nearbyGroupInfoPo, NearbyPo nearbyPo, int action)
			throws Parameter_Exception {
		nearbyGroupInfoPo.setTime(getTimes());
		redisClient79.findGroupId(nearbyPo);
		switch (action) {
		case 2:
			if (!nearbyPo.getOwnerid().equals(nearbyPo.getTuid())) {
				throw new Parameter_Exception(21005);
			}
			String info2 = nearbyPo.getTuid() + ":" + nearbyPo.getTgid();
			nearbyGroupInfoPo.setTypes(action);
			nearbyGroupInfoPo.setTuid(nearbyPo.getTuid());
			nearbyGroupInfoPo.setType(1);
			nearbyGroupInfoPo.setInfo(info2);
			List<Integer> tuids2 = nearbyPo.getTuids();
			for (int i = 0; i < tuids2.size(); i++) {
				nearbyGroupInfoPo.setToid(tuids2.get(i));
				redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			}
			break;
		case 3:
			String info3 = nearbyPo.getTuid() + ":" + nearbyPo.getTgid();
			nearbyGroupInfoPo.setTypes(action);
			nearbyGroupInfoPo.setTuid(nearbyPo.getTuid());
			nearbyGroupInfoPo.setToid(nearbyPo.getOwnerid());
			nearbyGroupInfoPo.setType(1);
			nearbyGroupInfoPo.setInfo(info3);
			redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			List<Integer> tuids3 = new ArrayList<>();
			tuids3.add(nearbyPo.getTuid());
			nearbyPo.setTuids(tuids3);
			redisClient79.setGroupMember(nearbyPo);
			findGroup(nearbyPo);
			return nearbyPo;
		case 4:
			String info4 = nearbyPo.getTuid() + ":" + nearbyPo.getTgid();
			nearbyGroupInfoPo.setTypes(action);
			nearbyGroupInfoPo.setTuid(nearbyPo.getTuid());
			nearbyGroupInfoPo.setToid(nearbyPo.getOwnerid());
			nearbyGroupInfoPo.setType(1);
			nearbyGroupInfoPo.setInfo(info4);
			redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			break;
		case 5:
			if (!nearbyPo.getOwnerid().equals(nearbyPo.getTuid())) {
				throw new Parameter_Exception(21005);
			}
			String info5 = nearbyPo.getTuid() + ":" + nearbyPo.getTgid();
			nearbyGroupInfoPo.setTypes(action);
			nearbyGroupInfoPo.setTuid(nearbyPo.getTuid());
			nearbyGroupInfoPo.setToid(nearbyPo.getTuids().get(0));
			nearbyGroupInfoPo.setType(1);
			nearbyGroupInfoPo.setInfo(info5);
			redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			redisClient79.setGroupMember(nearbyPo);
			break;
		case 6:
			if (!nearbyPo.getOwnerid().equals(nearbyPo.getTuid())) {
				throw new Parameter_Exception(21005);
			}
			String info6 = nearbyPo.getTgid() + ":" + nearbyPo.getTuids().get(0);
			nearbyGroupInfoPo.setTypes(action);
			nearbyGroupInfoPo.setType(1);
			nearbyGroupInfoPo.setInfo(info6);
			nearbyGroupInfoPo.setTuid(nearbyPo.getTuid());
			List<String> list = redisClient79.findGroupMember(nearbyPo);
			if (list.size() == 1) {
				throw new Parameter_Exception(21005);
			}
			for (String touid : list) {
				nearbyGroupInfoPo.setToid(Integer.valueOf(touid));
				redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			}
			redisClient79.deleteGroupMember(nearbyPo);
			break;
		case 7:
			nearbyGroupInfoPo.setTypes(action);
			nearbyGroupInfoPo.setType(1);
			nearbyGroupInfoPo.setTuid(nearbyPo.getTuid());
			List<String> list7 = redisClient79.findGroupMember(nearbyPo);
			for (String touid : list7) {
				String info7 = nearbyPo.getTuid() + ":" + nearbyPo.getTgid();
				nearbyGroupInfoPo.setInfo(info7);
				nearbyGroupInfoPo.setToid(Integer.valueOf(touid));
				redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			}
			redisClient79.deleteGroupMember(nearbyPo);
			break;
		case 8:
			if (!nearbyPo.getOwnerid().equals(nearbyPo.getTuid())) {
				throw new Parameter_Exception(21005);
			}
			String info8 = "群主解散了群";
			nearbyGroupInfoPo.setTypes(action);
			nearbyGroupInfoPo.setType(1);
			nearbyGroupInfoPo.setInfo(info8);
			nearbyGroupInfoPo.setTuid(nearbyPo.getTuid());
			List<String> list8 = redisClient79.findGroupMember(nearbyPo);
			for (String touid : list8) {
				nearbyGroupInfoPo.setToid(Integer.valueOf(touid));
				redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			}
			redisClient79.deleteGroupId(nearbyPo.getTgid());
			break;
		case 9:
			if (!nearbyPo.getOwnerid().equals(nearbyPo.getTuid())) {
				throw new Parameter_Exception(21005);
			}
			redisClient79.updateGroupId(nearbyPo);
			break;
		case 10:
			findGroup(nearbyPo);
			return nearbyPo;
		case 11:
			List<String> list11 = redisClient79.findGroupMember(nearbyPo);
			for (String str : list11) {
				nearbyGroupInfoPo.setToid(Integer.valueOf(str));
				if (nearbyGroupInfoPo.getTuid().equals(nearbyGroupInfoPo.getToid())) {
					continue;
				}
				redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
			}
			return null;
		default:
			break;
		}
		return null;
	}

	/**
	 * 查询群信息
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	private void findGroup(NearbyPo nearbyPo) throws Parameter_Exception {
		redisClient79.findGroupInfo(nearbyPo);
		List<String> list10 = redisClient79.findGroupMember(nearbyPo);
		nearbyPo.setNearbyPos(redisClient80.findNearbyToKey(list10));
	}

	/**
	 * 对附近的人发送消息
	 * 
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	@Override
	public void sendNearbyInfo(NearbyGroupInfoPo nearbyGroupInfoPo) throws Parameter_Exception {
		redisClient80.sendNearbyInfoToKey(nearbyGroupInfoPo);
	}

	/**
	 * 收取自己的消息
	 * 
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	@Override
	public List<NearbyGroupInfoPo> getNearbyInfo(NearbyGroupInfoPo nearbyPo) throws Parameter_Exception {
		List<NearbyGroupInfoPo> list = redisClient80.findNearbyInfoToKey(nearbyPo);
		return list;
	}

}
