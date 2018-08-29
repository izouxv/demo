package com.lanlian.chatting.serviceImpl;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import javax.annotation.Resource;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.controller.MyAbstractController;
import com.lanlian.chatting.dao.GroupSettingsDao;
import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.GroupAndUserPO;
import com.lanlian.chatting.po.GroupInfoPo;
import com.lanlian.chatting.po.virtual.TemporaryGroup;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.AccountClient;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.GroupSettingsService;
import com.lanlian.chatting.service.UserLetterService;
import com.lanlian.chatting.util.DataFinals;
import com.lanlian.chatting.vo.GroupInfo;
import com.lanlian.chatting.vo.UserLetter;
import com.lanlian.server.redis.RedisClient79;

@Service("groupSettingsServiceImpl")
public class GroupSettingsServiceImpl extends MyAbstractController implements GroupSettingsService {

	@Resource
	GroupSettingsDao groupSettingsDao;

	@Resource
	UserLetterService letter;

	@Resource
	AccountClient accountClient;
	
	@Resource
	SsoClient ssoClient;
	
	@Resource
	RedisClient79 redisClient79;
	
	/**
	 * 查询群信息
	 * 
	 * @param groupInfoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public GroupInfoPo findGroupInfo(GroupInfoPo groupInfoPo) throws Parameter_Exception {
		System.err.println("findGroupInfo-groupInfoPo:"+groupInfoPo);
		groupInfoPo = groupSettingsDao.findGroupUid(groupInfoPo);
		// System.out.println("findGroupInfo-groupInfoPo:"+groupInfoPo);
		return groupInfoPo;
	}
	
	/**
	 * 创建群并存入群用户关系
	 * @param groupInfoPo
	 * @param list
	 * @throws Parameter_Exception 
	 */
	@Transactional
	@Override
	public void saveGroupAndUser(GroupInfoPo groupInfoPo,List<Integer> list, String source) throws Parameter_Exception {
		groupInfoPo.setDataCreateTime(getTime());
		groupInfoPo.setDataModifyTime(getTime());
		groupInfoPo.setDataState(1);
		// System.out.println("saveGroupAndUser-groupInfoPo/:" + groupInfoPo);
		groupSettingsDao.saveGroup(groupInfoPo);
		// System.out.println("saveGroupAndUser-groupInfoPo//:"+groupInfoPo);

		int flag1 = list.size();
		list = ssoClient.getBatchSsoInfos(list, source);
		int flag2 = list.size();
		if (flag1 != flag2) {
			throw new Parameter_Exception(20009);
		}

		List<GroupAndUserPO> list2 = new ArrayList<>();
		GroupAndUserPO groupAndUserPO = null;
		for (Integer uid : list) {
			groupAndUserPO = new GroupAndUserPO();
			groupAndUserPO.setGid(groupInfoPo.getGid());
			groupAndUserPO.setUid(uid);
			groupAndUserPO.setCreatTime(getTime());
			list2.add(groupAndUserPO);
		}
		// System.out.println("saveGroupAndUser-list2//:"+list2);
		groupSettingsDao.saveMembers(list2);
	}

	/**
	 * 同步群
	 * 
	 * @throws Parameter_Exception
	 * @throws InterruptedException
	 */
	@Transactional
	@Override
	public GroupInfoPo syncGroup(GroupInfoPo groupSync, List<Integer> list, String source) throws Parameter_Exception {

		int flag1 = list.size();
		list = ssoClient.getBatchSsoInfos(list, source);
		int flag2 = list.size();
		if (flag1 != flag2) {
			throw new Parameter_Exception(20009);
		}

		Set<Integer> tableUids = new HashSet<>(); // 数据库中的群成员
		tableUids.addAll(groupSettingsDao.getGroupUser(groupSync.getGid()));

		Set<Integer> reqUids = new HashSet<>(); // 请求中上传的群成员
		reqUids.addAll(list);

		// 比对后将要插入数据库的群成员
		Set<Integer> insertTable = new HashSet<>();

		// 比对后将要从数据库中删除的群成员
		Set<Integer> deleteTable = new HashSet<>();

		// 更新群信息
		groupSettingsDao.updateGroupInfo(groupSync);

		// 批量删除群成员
		deleteTable.addAll(tableUids);
		deleteTable.removeAll(reqUids);
		deleteTable.remove(groupSync.getUid());
		if (deleteTable != null && !deleteTable.isEmpty()) {
			List<GroupAndUserPO> userList = uisToPo(groupSync, deleteTable, 2);
			groupSettingsDao.deleteMembers(userList);
		}

		insertTable.addAll(reqUids);
		insertTable.removeAll(tableUids);
		// 批量插入群成员
		if (insertTable != null && !insertTable.isEmpty()) {
			List<GroupAndUserPO> userList = uisToPo(groupSync, insertTable, 1);
			groupSettingsDao.saveMembers(userList);
		}

		// 对删除的成员发送私信
		UserLetter userLetter = deleteMember(groupSync, deleteTable);
		// 对添加的成员发送私信
		insertTable.clear();
		// System.out.println("删除成员:"+resultToDelete);

		insertTable.addAll(reqUids);
		insertTable.remove(groupSync.getUid());
		insertTable.removeAll(tableUids);
		// System.out.println("添加成员:"+resultToInsert);
		for (int opid : insertTable) {
			userLetter.setTouid(opid);
			userLetter.setLetter(groupSync.getUid() + DataFinals.PUBLIC + DataFinals.ADD_GROUP);
			letter.sendPrivateLette(userLetter);
		}
		return groupSync;
	}

	/**
	 * 信息放入list集合
	 * 
	 * @param groupSync
	 * @param uids
	 * @return
	 */
	private List<GroupAndUserPO> uisToPo(GroupInfoPo groupSync, Set<Integer> uids, int state) {
		List<GroupAndUserPO> userList = new ArrayList<>();
		GroupAndUserPO users = null;
		for (int uid : uids) {
			users = new GroupAndUserPO();
			users.setGid(groupSync.getGid());
			users.setUid(uid);
			users.setDataState(state);
			userList.add(users);
		}
		return userList;
	}

	/**
	 * 给被踢出群的发送私信
	 * 
	 * @param groupUpgrade
	 * @param resultToDelete
	 * @return
	 * @throws Parameter_Exception
	 */
	private UserLetter deleteMember(GroupInfoPo groupUpgrade, Set<Integer> resultToDelete) throws Parameter_Exception {
		UserLetter userLetter = new UserLetter();
		userLetter.setUid(DataFinals.DADA);
		userLetter.setType("1");
		userLetter.setSendTime(System.currentTimeMillis());
		for (Integer opid : resultToDelete) {
			userLetter.setTouid(opid);
			userLetter.setLetter(DataFinals.REMOVE_GROUP + "，由" + groupUpgrade.getUid() + "处理");
			letter.sendPrivateLette(userLetter);
		}
		return userLetter;
	}

	/**
	 * 编辑群资料
	 */
	@Transactional
	@Override
	public void updateGroupInfo(GroupInfoPo groupInfoPo) {
		groupSettingsDao.updateGroupInfo(groupInfoPo);
	}

	/**
	 * 批量删除群成员
	 * 
	 * @throws Parameter_Exception
	 */
	@Transactional
	@Override
	public void deleteMembers(GroupInfoPo groupInfoPo, List<Integer> list, String source) throws Parameter_Exception {
		int flag1 = list.size();
		list = ssoClient.getBatchSsoInfos(list, source);
		int flag2 = list.size();
		if (flag1 != flag2) {
			throw new Parameter_Exception(20009);
		}
		Set<Integer> uids = new HashSet<>();
		uids.addAll(list);
		List<GroupAndUserPO> userList = uisToPo(groupInfoPo, uids, 2);
		groupSettingsDao.deleteMembers(userList);
	}

	/**
	 * 群解散
	 */
	@Transactional
	@Override
	public void deleteGroup(GroupInfoPo groupSync) {
		GroupAndUserPO groupAndUserPo = new GroupAndUserPO();
		groupAndUserPo.setGid(groupSync.getGid());
		groupAndUserPo.setDataState(2);
		groupSettingsDao.deleteGroup(groupSync);
		groupSettingsDao.deleteGroupMembers(groupAndUserPo);
	}

	/**
	 * 根据uid查询永久群信息列表
	 * 
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public JSONArray findGroupByUidList(GroupInfoPo groupInfoPo, String source) throws Parameter_Exception {
		// 查询用户的群列表
		List<GroupInfo> list = new ArrayList<>();
		list = groupSettingsDao.findGroupByUidList(groupInfoPo);
		if (list.isEmpty()) {
			return null;
		}
		JSONArray groupJsonArray = getGroupUserByRpcToJson(list,source);
		return groupJsonArray;
	}

	/**
	 * 根据gid查询永久群信息列表
	 * 
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public JSONArray findGroupByGidList(List<GroupInfo> list, String source) throws Parameter_Exception {
		// 查询的群列表
		List<GroupInfo> groupInfos = groupSettingsDao.findGroupByGidList(list);
		if (groupInfos.isEmpty()) {
			throw new Parameter_Exception(21002);
		}
		JSONArray groupJsonArray = getGroupUserByRpcToJson(groupInfos, source);
		return groupJsonArray;
	}
	
	/**
	 * 查询群指定成员 ****弃用
	 * 
	 * @param list
	 * @return
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public GroupAndUserPO getGroupUser(GroupAndUserPO groupAndUserPO) {
		return null;
	}

	/**
	 * 查询群成员
	 * 
	 * @param list
	 * @return
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public JSONArray getGroupUserByRpcToJson(List<GroupInfo> list, String source) throws Parameter_Exception {
		// 查询群的成员id
		List<GroupAndUserPO> groupAndUserPOs = groupSettingsDao.getGroupsUser(list);
		if (groupAndUserPOs.isEmpty()) {
			return null;
		}
		// 用户的id存入set去重
		Set<Integer> uidSet = new HashSet<>();
		for (GroupAndUserPO uid : groupAndUserPOs) {
			uidSet.add(uid.getUid());
		}
		List<Integer> uids = new ArrayList<>(uidSet);
		List<AccountPo> accountPos = accountClient.getBatchAllUserInfo(uids, source);

		JSONArray groupJsonArray = new JSONArray();
		// 遍历群
		JSONArray userJsonArray = null;
		try {
			for (GroupInfo groupInfo : list) {
				groupInfo.setGname(URLDecoder.decode(groupInfo.getGname(), "UTF-8"));
				userJsonArray = new JSONArray();
				// 遍历群成员
				for (GroupAndUserPO groupAndUserPO : groupAndUserPOs) {
					// 遍历的群id等于用户的群id
					if (groupInfo.getGid() == groupAndUserPO.getGid()) {
						// 遍历获取到的用户信息
						for (AccountPo accountPo : accountPos) {
							if (groupAndUserPO.getUid() == accountPo.getUid()) {
								JSONObject jsonObject = BeanToJson(accountPo);
								userJsonArray.add(jsonObject);
							}
						}
					}
				}
				groupInfo.setUsers(userJsonArray);
				groupJsonArray.add(groupInfo);
			}
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(21002);
		}
		return groupJsonArray;
	}

	/**
	 * 临时群备份
	 * 
	 * @param types
	 * @param temporaryGroup
	 * @throws Parameter_Exception
	 */
	@Override
	public void temporaryGroup(String types, TemporaryGroup temporaryGroup) throws Parameter_Exception {
		switch (types) {
		case "sync":
			redisClient79.setTemporaryGroup(temporaryGroup);
			return;
		case "get":
			redisClient79.findTemporaryGroup(temporaryGroup);
			return;
		default:
			throw new Parameter_Exception(21002);
		}
	}
	
	
/*************************   私有方法    *********************************************************/

	/**
	 * account属性值放入jsonObject
	 * 
	 * @param accountPo
	 * @return
	 */
	private JSONObject BeanToJson(AccountPo accountPo) {
		JSONObject jsonObject = new JSONObject();
		jsonObject.put("uid", accountPo.getUid());
		jsonObject.put("nickname", accountPo.getNickname());
		jsonObject.put("gender", accountPo.getGender());
		jsonObject.put("birthday", accountPo.getBirthday());
		jsonObject.put("avatar", accountPo.getAvatar());
		jsonObject.put("signature", accountPo.getSignature());
		jsonObject.put("phone", accountPo.getPhone());
		jsonObject.put("email", accountPo.getEmail());
		jsonObject.put("province", accountPo.getProvince());
		jsonObject.put("city", accountPo.getCity());
		jsonObject.put("isCertification", accountPo.getIsCertification());
		jsonObject.put("userJobId", accountPo.getUserJobId());
		jsonObject.put("userGradeId", accountPo.getUserGradeId());
		jsonObject.put("creditValues", accountPo.getCreditValues());
		return jsonObject;
	}

}
