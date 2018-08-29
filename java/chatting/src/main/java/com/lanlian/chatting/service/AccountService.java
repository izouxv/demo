package com.lanlian.chatting.service;

import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @author  wdyqxx
 * @version 2017年1月11日 下午1:15:10
 * @explain 此接口用于用户完善、修改个人信息的service层；
 */
public interface AccountService {
	
	/**
	 * account添加用户
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	void addAccount(AccountPo accountPo) throws Parameter_Exception;
	
	/**
	 * ExInfo添加用户
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	void addExInfo(AccountPo accountPo) throws Parameter_Exception;
	
	/**
	 * account用于完善与修改用户个人信息；
	 * @param userInfo
	 * @return
	 * @throws Parameter_Exception 
	 */
	void accountUserInfo(AccountPo accountPo) throws Parameter_Exception;
	
	
}
