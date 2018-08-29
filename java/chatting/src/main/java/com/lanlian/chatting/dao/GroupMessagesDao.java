package com.lanlian.chatting.dao;

import java.util.List;

import com.lanlian.chatting.po.GroupMessageBackupPO;
import com.lanlian.chatting.vo.GroupMessageInfo;

public interface GroupMessagesDao {
	
	/**
	 * 插入聊天数据
	 * @param groupMessageRecord
	 * @return
	 */
	int messagesUpload(GroupMessageBackupPO groupMessageRecord);
	
	/**
	 * 插入数据
	 * @param messagesList
	 * @return
	 */
	int messageList(GroupMessageInfo messagesList);
	
	/**
	 * 获取数据
	 * @param id
	 * @return
	 */
	GroupMessageInfo messageUpdown(GroupMessageBackupPO groupMessageBackupPO);
	
	/**
	 * 查询用户列表
	 * @param gid
	 * @return
	 */
	List<GroupMessageBackupPO> theMessageList(long gid);
	
	
}
