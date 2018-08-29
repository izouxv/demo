package com.lanlian.chatting.dao;

import java.util.List;

import com.lanlian.chatting.po.FriendsPo;

/**
 * @author  wdyqxx
 * @version 2017年1月4日 上午10:50:39
 * @explain 此接口用于用户好友的增删改查信息操作的Dao层；
 */
public interface FriendsDao {
	
	/**
	 * 查询好友
	 * @param FriendsPo
	 * @return List<FriendsPo>
	 */
	List<FriendsPo> findFriends(FriendsPo friendsPo);
	
	/**
	 * 验证好友
	 * @param list
	 * @return
	 */
	List<FriendsPo> verifyFriends(List<FriendsPo> list);
	
	/**
	 * 添加好友
	 * @param List<FriendsPo>
	 * @return 
	 */
	void saveFriends(List<FriendsPo> list);
	
	/**
	 * 修改好友印象
	 * @param List<FriendsPo>
	 * @return 
	 */
	int modifyFriends(List<FriendsPo> list);
	
	/**
	 * 删除好友
	 * @param List<FriendsPo>
	 * @return 
	 */
	void deleteFriends(List<FriendsPo> list);
	
	
}
