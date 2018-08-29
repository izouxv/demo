package com.lanlian.chatting.service;

import java.util.List;

import com.lanlian.chatting.bo.DataBo;
import com.lanlian.chatting.po.FriendInfoPo;
import com.lanlian.chatting.po.FriendsPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.vo.UserFriends;

/**
 * @author wdyqxx
 * @version 2017年1月4日 下午2:23:35
 * @explain 此接口用于用户好友操作的service层；
 */

public interface FriendService {

	/**
	 * 获取用户联系人信息；
	 * 
	 * @return List<AccountPo>
	 * @throws Parameter_Exception
	 */
	List<FriendInfoPo> findFriends(FriendsPo friendsPo, String source) throws Parameter_Exception;

	/**
	 * 批量添加好友；
	 * 
	 * @return
	 * @throws Parameter_Exception
	 */
	void saveFriends(UserFriends userFriends, DataBo dataBo) throws Parameter_Exception;

	/**
	 * 批量校验好友；
	 * 
	 * @return
	 * @throws Parameter_Exception
	 */
	DataBo checkFriends(UserFriends userFriends, String source) throws Parameter_Exception;

	/**
	 * 批量删除好友
	 * @param uid
	 * @param toUids
	 */
	void deleteMyFriends(int uid, List<Integer> toUids);
	
	/**
	 * 批量修改好友信息
	 * @param list
	 */
	void modifyFriends(List<FriendsPo> list);
	
}
