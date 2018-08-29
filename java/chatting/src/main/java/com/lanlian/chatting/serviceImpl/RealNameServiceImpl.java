package com.lanlian.chatting.serviceImpl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.AccountClient;
import com.lanlian.chatting.service.RealNameService;

/**
 * 用于处理用户实名认证业务
 * 
 * @author my
 *
 */
@Service("realNameServiceImpl")
public class RealNameServiceImpl implements RealNameService {

	@Autowired
	public AccountClient accountClient;
	
	/**
	 * 查询实名信息
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	@Override
	public void getCertification(AccountPo accountPo) throws Parameter_Exception {

		// 检验用户是否已经进行过实名认证
		accountPo = accountClient.getCertification(accountPo);
		if (33010 == accountPo.getErrorCode()) {
			throw new Parameter_Exception(accountPo.getErrorCode());
		}
		if (10000 != accountPo.getErrorCode()) {
			throw new Parameter_Exception(accountPo.getErrorCode());
		}
		if (accountPo.getIsCertification() == 2) {
			// 该账户已被实名认证
			throw new Parameter_Exception(20022);
		}
	}
	
	/**
	 * 实名认证
	 * @param accountPo
	 * @throws Parameter_Exception
	 */
	@Override
	public void updateCertification(AccountPo accountPo) throws Parameter_Exception {
		// 调用实名认证接口判断用户姓名,身份证号码,手机号码是否一致
		
		// 通过第三方接口判断,保存用户实名认证信息
		accountPo = accountClient.updateCertification(accountPo);
		if (10000 != accountPo.getErrorCode()) {
			throw new Parameter_Exception(accountPo.getErrorCode());
		}
	}

}
