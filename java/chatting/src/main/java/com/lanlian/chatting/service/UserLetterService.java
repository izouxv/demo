package com.lanlian.chatting.service;

import java.io.UnsupportedEncodingException;
import java.util.List;

import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.vo.UserLetter;
import com.lanlian.chatting.vo.UserMessageNum;

import net.sf.json.JSONArray;

/**
 * @author  wdyqxx
 * @version 2017年1月6日 上午11:17:33
 * @explain 此接口为业务处理层，用于处理用户发送私信操作；
 */

public interface UserLetterService {
	
	/**
	 * 此方法用于用户发送私信功能；
	 * @param um
	 * @throws Parameter_Exception 
	 */
	void sendPrivateLette(UserLetter um) throws Parameter_Exception;
	
	/**
	 * 此方法用于用户传入自己的pid，返回用户每个联系人未读取的通知数量；
	 * @param um
	 * @return
	 * @throws Parameter_Exception 
	 */
	JSONArray receiveLetteNum(UserMessageNum um) throws Parameter_Exception;
	
	/**
	 * 获取私信内容
	 * @param um
	 * @return
	 * @throws Parameter_Exception
	 * @throws UnsupportedEncodingException 
	 */
	List<UserLetter> receiveLetter(UserLetter um) throws Parameter_Exception, UnsupportedEncodingException;
	
}
