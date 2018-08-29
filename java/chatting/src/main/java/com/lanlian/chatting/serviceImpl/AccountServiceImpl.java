package com.lanlian.chatting.serviceImpl;

import javax.annotation.Resource;

import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.AccountClient;
import com.lanlian.chatting.service.AccountService;

/**
 * @author wdyqxx
 * @version 2017年1月11日 下午1:33:45
 * @explain
 */
@Service("accountService")
public class AccountServiceImpl implements AccountService {
	
	@Resource
	AccountClient accountClient;
	
	/**
	 * account添加用户
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	@Override
	public void addAccount(AccountPo accountPo) throws Parameter_Exception {
		accountPo = accountClient.addAccountInfo(accountPo);
		if (10000 != accountPo.getErrorCode()) {
			//该用户状态异常
			throw new Parameter_Exception(accountPo.getErrorCode());
		}
	}
	
	/**
	 * ExInfo添加用户
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	@Async
	@Override
	public void addExInfo(AccountPo accountPo) {
			accountPo = accountClient.updateExInfo(accountPo);
			if (10000 != accountPo.getErrorCode()) {
				//该用户状态异常
				LogUtil.error(accountPo.toString());
			}
	}
	
	
	/**
	 * 修改用户信息
	 */
	@Override
	public void accountUserInfo(AccountPo accountPo) throws Parameter_Exception {
		accountPo = accountClient.updateAccountInfo(accountPo);
		if (10000 != accountPo.getErrorCode()) {
			throw new Parameter_Exception(accountPo.getErrorCode());
		}
	}


}
