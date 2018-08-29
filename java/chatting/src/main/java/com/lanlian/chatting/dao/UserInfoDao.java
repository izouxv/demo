package com.lanlian.chatting.dao;

import java.util.List;

import com.lanlian.chatting.po.FriendInfoPo;
import com.lanlian.chatting.po.UserInfoPO;

/**
 * @author  wdyqxx
 * @version 2017年1月10日 下午5:36:17
 * @explain 
 */

public interface UserInfoDao {
	
	/**
	 * 查询用户帐号的当前状态；
	 * @param userInfoPO
	 * @return 
	 */
	UserInfoPO findUidByUserStatus(UserInfoPO userInfoPO);
	
	/**
	 * 用于完善与修改用户信息；
	 * @param userInfo
	 * @return
	 */
	int completeUserInfo(UserInfoPO userInfo);
	
	/**
	 * findPidByUserInfo：以saveFriendsList查询出的数据解析，
	 * 得到每个联系人的pid，查询联系人的信息；
	 * @param UserInfoPOJO的pid
	 * @return UserInfoPOJO
	 */
	UserInfoPO findPidByUserInfo(UserInfoPO pid);
	
	/**
	 * 以saveFriendsList查询出的数据解析，得到每个联系人的pid，查询联系人的信息；
	 * @param userFriendInfo
	 * @return List<UserFriendInfoPO>
	 */
	List<FriendInfoPo> findPidByFriendsInfo(FriendInfoPo userFriendInfo);
	
	
}
