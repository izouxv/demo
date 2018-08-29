package com.lanlian.chatting.dao;

import java.util.List;

import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.GroupAndUserPO;
import com.lanlian.chatting.po.GroupInfoPo;
import com.lanlian.chatting.po.LiveChatGidUid;
import com.lanlian.chatting.vo.GroupInfo;

/** 
 * @Title GroupSettingsDao.java
 * @Package com.lanlian.chatting.dao
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午1:57:37
 * @explain 信息实时上报群消息接口
 */
public interface GroupSettingsDao_1_1 {

	/**
	 * 查询永久群是否存在
	 * @param gid
	 * @return
	 */
	int getGroupExists(int gid);

	/**
	 *  创建群
	 * @param groupSync
	 * @return
	 */
	void saveGroup(GroupInfoPo groupSync);
	
	
	/**
	 *  批量插入，群成员
	 * @param list
	 * @return
	 */
	void saveMembers(List<GroupAndUserPO> list);

	/**
	 * 批量删除，群成员
	 * @param list
	 * @return
	 */
	int deleteMembers(List<GroupAndUserPO> list);

	/**
	 *  查询群主uid
	 * @param groupInfoPo
	 * @return
	 */
	GroupInfoPo findGroupUid(GroupInfoPo groupInfoPo);
	
	/**
	 *  编辑群资料
	 * @param groupInfoPo
	 * @return
	 */
	void updateGroupInfo(GroupInfoPo groupInfoPo);

	/**
	 *  解散群
	 * @param groupInfoPo
	 * @return
	 */
	int deleteGroup(GroupInfoPo groupInfoPo);
	
	/**
	 * 删除群成员
	 * @param list
	 * @return
	 */
	void deleteGroupMembers(GroupAndUserPO groupAndUserPO);

	/**
	 * 查询群指定成员
	 * @param list
	 * @return
	 */
	GroupAndUserPO getGroupMember(GroupAndUserPO groupAndUserPO);
	
	/**
	 *  查询群的用户
	 * @param gid
	 * @return
	 */
	List<Integer> getGroupUser(int gid);
	
	/**
	 *  批量查询群的用户
	 * @param list
	 * @return
	 */
	List<GroupAndUserPO> getGroupsUser(List<GroupInfo> list);

	/**
	 *  根据uid查询永久群信息列表
	 * @param uid
	 * @return
	 */
	List<GroupInfo> findGroupByUidList(GroupInfoPo groupInfoPo);
	
	/**
	 * 根据gid查询永久群信息列表
	 * @param groupInfoPo
	 * @return
	 */
	List<GroupInfo> findGroupByGidList(List<GroupInfo> list);
	
/********************** 群信息实时上报接口  **************************/
	
	/**
	 * 查询群
	 * @param dadaGroupPo
	 * @return DadaGroupPo
	 */
	DadaGroupPo findGroup(DadaGroupPo dadaGroupPo);
	
	/**
	 * 查询实时上报群
	 * @param dadaGroupPo
	 * @return DadaGroupPo
	 */
	DadaGroupPo findDadaGroup(DadaGroupPo dadaGroupPo);
	
	/**
	 * 查询实时上报群的开启人与状态
	 * @param dadaGroupPo
	 * @return
	 */
	DadaGroupPo findDadaGroupSwitchingUser(DadaGroupPo dadaGroupPo);
	
	/**
	 * 创建实时上报群
	 * @param dadaGroupPo
	 * @return 
	 */
	void saveDadaGroup(DadaGroupPo dadaGroupPo);
	
	/**
	 * 存入邀请码
	 * @param dadaGroupPo
	 */
	int saveDadaGroupCode(DadaGroupPo dadaGroupPo);
	
	/**
	 * 创建实时上报群开关
	 * @param dadaGroupPo
	 * @return 
	 */
	void saveDadaGroupSwitching(DadaGroupPo dadaGroupPo);
	
	/**
	 * 修改实时上报群
	 * @param dadaGroupPo
	 */
	void updateDadaGroup(DadaGroupPo dadaGroupPo);
	
	/**
	 * 查询绑定信息
	 * @param liveChatGidUid
	 * @return
	 */
	LiveChatGidUid findBundilingBygid(LiveChatGidUid liveChatGidUid);
	
}
