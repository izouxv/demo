package com.lanlian.chatting.service;

import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.result.Parameter_Exception;

public interface RealNameService {
	
	/**
	 * 查询实名信息
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	void getCertification(AccountPo accountPo) throws Parameter_Exception;
	
	/**
	 * 实名认证
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	void updateCertification(AccountPo accountPo) throws Parameter_Exception;
}
