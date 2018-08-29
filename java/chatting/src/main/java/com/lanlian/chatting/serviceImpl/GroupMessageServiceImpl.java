package com.lanlian.chatting.serviceImpl;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import com.lanlian.chatting.dao.GroupMessagesDao;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.GroupMessageBackupPO;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.GroupMessageService;
import com.lanlian.chatting.vo.GroupMessageInfo;

@Service(value = "messageService")
public class GroupMessageServiceImpl implements GroupMessageService {

	@Autowired
	private GroupMessagesDao messagesDao;// 调用消息的接口

	/**
	 * 消息上传的实现
	 */
	@Transactional(propagation = Propagation.REQUIRED)
	@Override
	public void messagesUpload(GroupMessageBackupPO groupMessageRecord) {
		messagesDao.messagesUpload(groupMessageRecord);
	}

	/**
	 * 消息的存储的方法实现
	 */
	@Transactional(propagation = Propagation.REQUIRED)
	@Override
	public void messageList(GroupMessageInfo messagesList) {
		messagesDao.messageList(messagesList);
	}

	/**
	 * 获取消息列表的实现
	 * @throws Parameter_Exception 
	 */
	@Transactional(readOnly = true)
	@Override
	public List<GroupMessageBackupPO> theMessageList(int uid) throws Parameter_Exception {
		List<GroupMessageBackupPO> list = new ArrayList<>();
		list = messagesDao.theMessageList(uid);
		if (list.size() < 1) {
			//无数据
			throw new Parameter_Exception(20023);
		}
		return list;
	}

	/**
	 * 根据消息id下载消息
	 */
	@Transactional(propagation = Propagation.REQUIRED)
	@Override
	public GroupMessageInfo messageUpdown(Integer uid, int messageId) {
		GroupMessageBackupPO groupMessageBackupPO = new GroupMessageBackupPO();
		groupMessageBackupPO.setUid(uid);
		groupMessageBackupPO.setMessageId(messageId);
		GroupMessageInfo messages = messagesDao.messageUpdown(groupMessageBackupPO);
		LogUtil.info("msg:"+messages);
		return messages;
	}

}
