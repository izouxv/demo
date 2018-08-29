package com.lanlian.chatting.serviceImpl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.AccountClient;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.LoginService;

/**
 * 
 * @Title LoginServiceImpl.java
 * @Package cn.lanlian.ccat.serviceImpl
 * @author wangdyqxx
 * @version V1.0
 * @date 2017年3月20日 下午5:41:29
 * @explain: 用于用户登录；
 */
@Service("loginService")
public class LoginServiceImpl implements LoginService {
	
	@Autowired
	SsoClient ssoClient;
	
	@Autowired
	AccountClient accountClient;
	
	/**
	 * ssoRpc
	 */
	@Override
	public SsoPo ssoRpc(SsoPo ssoPo) throws Parameter_Exception {
		// 通过数据库查询用户数据
		ssoPo = ssoClient.login(ssoPo);
		if (ssoPo.getErrorCode() != 10000) {
			throw new Parameter_Exception(ssoPo.getErrorCode());// rpc调用失败
		}
		// 根据用户的状态判断用户能否登录
		if (ssoPo.getState().intValue() == 3) {
			return ssoPo;
		}
		// 用户状态异常，禁止登陆
		throw new Parameter_Exception(21004);
	}
	
	/**
	 * accountRpc
	 * @throws Parameter_Exception 
	 */
	@Override
	public AccountPo accountRpc(AccountPo accountPo) throws Parameter_Exception {
		// 通过account查询用户数据
		accountPo = accountClient.getUserInfo(accountPo);
		if (accountPo.getErrorCode() != 10000) {
			throw new Parameter_Exception(accountPo.getErrorCode());
		}
		if (accountPo.getCreditValues() < 0) {
			throw new Parameter_Exception(20006);// 该用户信用点太低，禁止登陆
		}
		return accountPo;
	}
	
	/**
	 * 用于已登录用户退出登录操作的业务层处理
	 * 
	 * @throws Parameter_Exception
	 */
	@Override
	public void exitUser(SsoPo ssoPo) throws Parameter_Exception {
		ssoPo = ssoClient.logout(ssoPo);
		if (ssoPo.getErrorCode() != 10000) {
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
	}

}
