package com.lanlian.chatting.dao;

import java.util.List;

import com.lanlian.chatting.po.PrivateMessageInfoPO;
import com.lanlian.chatting.po.PrivateMessagePO;

/**
 * @author  wdyqxx
 * @version 2017年1月6日 上午9:12:52
 * @explain 此类为dao层接口，用于用户对联系人发送、接收私信操作；
 */

public interface UserMessageDao {
	
	/**
	 * 用户发送私信：此方法用于记录用户对联系人发送的私信功能，返回记录的id，
	 * 并将插入记录的id存入PrivateMessageInfoPOJO的属性id；
	 * @param pmpojo
	 * @return 
	 */
	void sendPrivateLetteInfoId(PrivateMessageInfoPO pmipojo);
	
	/**
	 * 用户发送私信：此方法将sendPrivateLetteInfoId方法的
	 * PrivateMessageInfoPOJO属性id放入PrivateMessagePOJO中，
	 * 然后将对象存入数据库，等待查看；
	 * @param pmpojo
	 * @return
	 */
	int sendPrivateLette(PrivateMessagePO pmpojo);
	
	/**
	 * 用户接收私信
	 * @param pmpo
	 * @return int
	 */
	List<PrivateMessagePO> receiveLetteById(PrivateMessagePO pmpo);
	
	/**
	 * 用户接收私信：此方法将PrivateMessagePOJO的pid做为参数查询private_info表，
	 * 返回指定用户发送的私信id；
	 * @param pmpojo
	 * @return int
	 */
	List<PrivateMessagePO> receiveLetteUidById(PrivateMessagePO pmpojo);
	
	int modifyLetteStatus(PrivateMessagePO pmpojo);
	
	/**
	 * 用户接收私信：此方法将PrivateMessagePOJO的fpid与opid做为参数查询private_info表，
	 * 将得到的PrivateMessagePOJO属性messageInfoId作为参数查询private_lette_info，
	 * 得到用户的私信内容PrivateMessageInfoPOJO；
	 * @param pmipojo
	 * @return PrivateMessageInfoPOJO
	 */
	List<PrivateMessageInfoPO> receiveLetteInfo(List<Long> list);
	
	
	
	
	
	
	
}
