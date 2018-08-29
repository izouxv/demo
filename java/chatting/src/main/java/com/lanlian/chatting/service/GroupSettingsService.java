package com.lanlian.chatting.service;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.lanlian.chatting.po.GroupAndUserPO;
import com.lanlian.chatting.po.GroupInfoPo;
import com.lanlian.chatting.po.virtual.TemporaryGroup;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.vo.GroupInfo;

public interface GroupSettingsService {
	
	/**
	 * 查询群信息
	 * @param groupInfoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	GroupInfoPo findGroupInfo(GroupInfoPo groupInfoPo) throws Parameter_Exception;

	/**
	 * 创建群并存入群用户关系
	 * @param groupInfoPo
	 * @param list
	 * @throws Parameter_Exception 
	 */
	void saveGroupAndUser(GroupInfoPo groupInfoPo, List<Integer> list, String source) throws Parameter_Exception;
	
	/**
	 * 同步群
	 * @param groupUpgrade
	 * @param members
	 * @return
	 * @throws Parameter_Exception
	 */
	GroupInfoPo syncGroup(GroupInfoPo groupUpgrade,List<Integer> list, String source) throws Parameter_Exception;
	
	/**
	 * 踢人出群，用户主动退群
	 * @param groupInfoPo
	 * @param list
	 * @return
	 * @throws Parameter_Exception
	 */
	void deleteMembers(GroupInfoPo groupInfoPo, List<Integer> list, String source) throws Parameter_Exception;
	
	/**
	 * 解散群，
	 * @param groupInfoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	void deleteGroup(GroupInfoPo groupInfoPo) throws Parameter_Exception;
	
	/**
	 * 编辑群资料
	 * @param groupInfoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	void updateGroupInfo(GroupInfoPo groupInfoPo) throws Parameter_Exception;
	
	
	/**
	 * 根据uid查询永久群信息列表
	 * @param uid
	 * @return
	 * @throws Parameter_Exception
	 */
	JSONArray findGroupByUidList(GroupInfoPo groupInfoPo, String source) throws Parameter_Exception;
	
	/**
	 * 根据gid查询永久群信息列表
	 * 
	 * @throws Parameter_Exception
	 */
	JSONArray findGroupByGidList(List<GroupInfo> list, String source) throws Parameter_Exception;

	/**
	 * 临时群同步
	 * @param types
	 * @param temporaryGroup
	 * @throws Parameter_Exception
	 */
	void temporaryGroup(String types, TemporaryGroup temporaryGroup) throws Parameter_Exception;
	
	/**
	 * 查询群成员
	 * 
	 * @param list
	 * @return
	 * @throws Parameter_Exception
	 */
	JSONArray getGroupUserByRpcToJson(List<GroupInfo> list, String source) throws Parameter_Exception;

	/**
	 * 查询群指定群员
	 * @param groupAndUserPO
	 * @return
	 */
	GroupAndUserPO getGroupUser(GroupAndUserPO groupAndUserPO);
}
