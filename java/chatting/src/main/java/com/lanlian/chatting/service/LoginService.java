package com.lanlian.chatting.service;

import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;

/**
 * 用于用户登录操作的业务层接口
 * 
 * @author my
 *
 */
public interface LoginService {
	
	/**
	 * 调用sso-rpc，获取sessionName
	 * @param ssoPo
	 * @return SsoPo
	 * @throws Parameter_Exception
	 */
	SsoPo ssoRpc(SsoPo ssoPo) throws Parameter_Exception;
	
	/**
	 * 调用account-rpc，获取用户信息
	 * @param accountPo
	 * @return AccountPo
	 * @throws Parameter_Exception
	 */
	AccountPo accountRpc(AccountPo accountPo) throws Parameter_Exception;
	
	/**
	 * 用于用户退出操作
	 * @param ssoPo
	 * @throws Parameter_Exception
	 */
	void exitUser(SsoPo ssoPo) throws Parameter_Exception;
}
