package com.lanlian.chatting.dao;

import java.util.List;

import com.lanlian.chatting.bo.MessagePageBo;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.LiveChatGidUid;
import com.lanlian.chatting.vo.MessageInfoPage;

/**
 * @Description: TODO
 * @author: 李大双
 * @date: 2017年6月30日 下午4:11:15
 * @version: V1.0
 */
public interface WeiChatGroupSettingsDao_1_1 {
	//绑定群
	void bundlingGP(LiveChatGidUid liveChatGidUid);
	//查询绑定表
	int findBundilingGP(LiveChatGidUid liveChatGidUid);
	//跟据邀请码查询群
	DadaGroupPo findGroupInfoByinviteCode(DadaGroupPo liveChatGroupPo);
	//默认第一次查询消息
	List<MessageInfoPage> findMessageInfo(MessagePageBo messagePageBo);
	//向上查询消息
	List<MessageInfoPage> findMessageInfoUp(MessagePageBo messagePageBo);
	//向下查询
	List<MessageInfoPage> findMessageInfoDown(MessagePageBo messagePageBo);
	//查询所有信息总数
	int findAllMsgCount();
}