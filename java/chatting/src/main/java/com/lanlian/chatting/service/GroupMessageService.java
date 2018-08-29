package com.lanlian.chatting.service;

import java.util.List;

import com.lanlian.chatting.po.GroupMessageBackupPO;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.vo.GroupMessageInfo;


public interface GroupMessageService {

	/**
	 *消息上传
	 * @param groupMessageRecord  消息类
	 * @return 
	 */
	public void messagesUpload(GroupMessageBackupPO groupMessageRecord);
	/**
	 * 消息方法
	 * @param messagesList
	 * @return
	 */
	public void messageList(GroupMessageInfo messagesList);
	
	/**
	 * 消息下载
	 * @param id  消息id
	 * @return
	 */
	public GroupMessageInfo messageUpdown(Integer uid, int id);
	
	/**
	 * 获取消息列表
	 * @param uid 用户id
	 * @return   消息列表
	 * @throws Parameter_Exception 
	 */
	public List<GroupMessageBackupPO>  theMessageList(int uid) throws Parameter_Exception;
	
	
	
}
