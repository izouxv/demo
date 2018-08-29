package com.lanlian.chatting.serviceImpl;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Set;

import javax.annotation.Resource;

import org.springframework.beans.BeanUtils;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import com.lanlian.chatting.bo.DataBo;
import com.lanlian.chatting.dao.FriendsDao;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.FriendInfoPo;
import com.lanlian.chatting.po.FriendsPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.AccountClient;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.FriendService;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.chatting.vo.UserFriends;

import net.sf.json.JSONArray;

/**
 * @author wdyqxx
 * @version 2017年1月4日 下午2:27:13
 * @explain 此类用于实现用户获取联系人信息的操作；
 */
@Service("userFriendsService")
public class FriendServiceImpl implements FriendService {

	@Resource
	FriendsDao friendsDao;

	@Resource
	AccountClient accountClient;

	@Resource
	SsoClient ssoClient;

	/**
	 * 获取用户联系人信息；
	 * 
	 * @return List<AccountPo>
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public List<FriendInfoPo> findFriends(FriendsPo friendsPo, String source) throws Parameter_Exception {
		// 返回值
		List<FriendInfoPo> friendInfoPos = new ArrayList<>();
		// 根据uid查询到联系人信息
		List<FriendsPo> friendList = friendsDao.findFriends(friendsPo);
		if (friendList.isEmpty()) {
			throw new Parameter_Exception(20023);
		}
		// 遍历friendList
		List<Integer> uids = new ArrayList<>();
		Set<Integer> uidSet = new HashSet<>();
		Map<Integer, String> map = new HashMap<>();
		for (FriendsPo friendsPo2 : friendList) {
			if (friendsPo.getUid1().equals(friendsPo2.getUid1())) {
				uidSet.add(friendsPo2.getUid2());
				map.put(friendsPo2.getUid2(), friendsPo2.getNote1());
			}
			if (friendsPo.getUid1().equals(friendsPo2.getUid2())) {
				uidSet.add(friendsPo2.getUid1());
				map.put(friendsPo2.getUid1(), friendsPo2.getNote2());
			}
		}
		uidSet.remove(friendsPo.getUid1());
		map.remove(friendsPo.getUid1());
		uids.addAll(uidSet);
		if (uids.isEmpty()) {
			// "无数据！请先在云端备份
			throw new Parameter_Exception(20023);
		}
		// 根据查询到的uid，查询联系人信息
		List<AccountPo> accountPos = accountClient.getBatchAllUserInfo(uids, source);
		// 数据聚集
		friendInfoPos = converInfo(accountPos, friendInfoPos, map);
		return friendInfoPos;
	}

	/**
	 * 添加好友；
	 * 
	 * @return
	 * @throws Parameter_Exception
	 */
	@Transactional
	@Override
	public void saveFriends(UserFriends userFriends, DataBo dataBo) throws Parameter_Exception {
		List<FriendsPo> friendsPos = resolveData(userFriends, dataBo);
		LogUtil.info("friendsPos:" + friendsPos);
		if (friendsPos.isEmpty()) {
			throw new Parameter_Exception(21001);
		}
		friendsDao.saveFriends(friendsPos);
	}

	/**
	 * 校验好友；
	 * 
	 * @return
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public DataBo checkFriends(UserFriends userFriends, String source) throws Parameter_Exception {
		DataBo dataBo = new DataBo();
		// 返回异常联系人信息
		Map<String, String> resultMap = new HashMap<>();
		// 解析用户上传的联系人
		Map<String, String> map = userFriends.getInfoMap();
		// 从map中取出uid，查询该联系人信息
		List<Integer> uidList = new ArrayList<>();
		// 存入数据库的集合
		List<FriendsPo> friendsPos = new ArrayList<>();
		// 遍历
		for (String uids : map.keySet()) {
			if (String.valueOf(userFriends.getUid()).equals(uids)) {
				resultMap.put(uids, map.get(uids));
				continue;
			}
			uidList.add(Integer.parseInt(uids));
		}
		// 去自己之后为空，直接返回
		if (uidList.isEmpty()) {
			throw new Parameter_Exception(21005);
		}
		// 验证好友关系
		bldata(userFriends.getUid(), uidList, friendsPos);
		// 取出用户的好友信息
		friendsPos = friendsDao.verifyFriends(friendsPos);
		// 遍历并区分
		for (FriendsPo friendsPo : friendsPos) {
			Iterator<Integer> iterator = uidList.iterator();
			while (iterator.hasNext()) {
				Integer uid = iterator.next();
				if (uid.equals(friendsPo.getUid1()) || uid.equals(friendsPo.getUid2())) {
					resultMap.put(uid.toString(), map.get(uid.toString()));
					// 将已经存在的好友关系的uid去掉
					iterator.remove();
				}
			}
		}
		// 清空friendsPos集合
		friendsPos.clear();
		// 去重之后为空，直接返回
		if (uidList.isEmpty()) {
			throw new Parameter_Exception(21005);
		}
		int flag1 = uidList.size();
		// 当双方不是好友关系，调用rpc
		uidList = ssoClient.getBatchSsoInfos(uidList, source);
		int flag2 = uidList.size();
		if (flag1 != flag2) {
			throw new Parameter_Exception(20009);
		}
		// 异常信息转换
		JSONArray jsonArray = PublicMethod.map_Json(resultMap);
		// 返回信息
		dataBo.setJson(jsonArray.toString());
		dataBo.setUidList(uidList);
		return dataBo;
	}

	/**
	 * 批量删除好友
	 */
	@Transactional
	@Override
	public void deleteMyFriends(int uid, List<Integer> uids) {
		List<FriendsPo> friendsPos = new ArrayList<>();
		// 将信息重组
		bldata(uid, uids, friendsPos);
		friendsDao.deleteFriends(friendsPos);
	}

	/**
	 * 批量修改好友信息
	 */
	@Transactional
	@Override
	public void modifyFriends(List<FriendsPo> list) {
		friendsDao.modifyFriends(list);
	}

	/** ========================= 私有方法 ====================================== */

	/**
	 * 将查询好友需要的数据重组
	 * 
	 * @param uid
	 * @param uidList
	 * @param friendsPos
	 */
	private void bldata(int uid, List<Integer> uidList, List<FriendsPo> friendsPos) {
		FriendsPo friendsPo1 = null;
		for (Integer toUid : uidList) {
			friendsPo1 = new FriendsPo();
			if (uid < toUid) {
				friendsPo1.setUid1(uid);
				friendsPo1.setUid2(toUid);
				friendsPos.add(friendsPo1);
			}
			if (uid > toUid) {
				friendsPo1.setUid1(toUid);
				friendsPo1.setUid2(uid);
				friendsPos.add(friendsPo1);
			}
		}
	}

	/**
	 * 将json数据解析并处理
	 * 
	 * @param userFriends
	 * @param dataBo
	 * @return
	 */
	private List<FriendsPo> resolveData(UserFriends userFriends, DataBo dataBo) {
		// 解析用户上传的联系人
		Map<String, String> map = userFriends.getInfoMap();
		List<FriendsPo> friendsPos = new ArrayList<>();
		// 将用户状态正常的联系人放入list，异常的联系人放入resultMap
		FriendsPo friendsPo = null;
		for (Integer uid : dataBo.getUidList()) {
			friendsPo = new FriendsPo();
			if (userFriends.getUid() < uid) {
				friendsPo.setUid1(userFriends.getUid());
				friendsPo.setNote1(map.get(uid.toString()).trim());
				friendsPo.setUid2(uid);
				friendsPos.add(friendsPo);
			}
			if (userFriends.getUid() > uid) {
				friendsPo.setUid1(uid);
				friendsPo.setUid2(userFriends.getUid());
				friendsPo.setNote2(map.get(uid.toString()).trim());
				friendsPos.add(friendsPo);
			}
		}
		return friendsPos;
	}

	/**
	 * 属性数据转换
	 * 
	 * @param accountPos
	 * @param friendList
	 * @param map
	 * @return
	 */
	private List<FriendInfoPo> converInfo(List<AccountPo> accountPos, List<FriendInfoPo> friendInfoPos,
			Map<Integer, String> map) {
		// 当联系人信息状态异常，则只返回用户异常信息与联系人uid；
		FriendInfoPo friendInfoPo = null;
		for (AccountPo accountPo : accountPos) {
			friendInfoPo = new FriendInfoPo();
			BeanUtils.copyProperties(accountPo, friendInfoPo);
			friendInfoPo.setNote(map.get(accountPo.getUid()).trim());
			friendInfoPos.add(friendInfoPo);
		}
		return friendInfoPos;
	}

}
