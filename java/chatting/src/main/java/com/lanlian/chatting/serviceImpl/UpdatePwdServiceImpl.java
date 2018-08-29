/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/**
 * 
 */
package com.lanlian.chatting.serviceImpl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.UpdatePwdService;

/**
 * @Title UpdatePwdServiceImpl.java
 * @Package cn.lanlian.chatting.service
 * @author 王东阳
 * @version V1.0
 * @date 2017年5月3日 上午11:08:22
 * @explain
 */
@Service(value = "updatePwdService")
public class UpdatePwdServiceImpl implements UpdatePwdService {

	
	@Autowired(required=true)
	SsoClient ssoClient;
	
	/**
	 * 校验密码
	 * 
	 * @throws Parameter_Exception
	 */
	@Override
	public void checkPassword(SsoPo ssoPo) throws Parameter_Exception {
		//校验密码
		ssoClient.checkPassword(ssoPo);
		if (10000 != ssoPo.getErrorCode()) {
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
	}

	/**
	 * 修改密码
	 * @throws Parameter_Exception
	 */
	@Override
	public void updatePassword(SsoPo ssoPo) throws Parameter_Exception {
		ssoClient.updatePassword(ssoPo);
		if (10000 != ssoPo.getErrorCode()) {
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
	}
	
	/**
	 * 重置密码
	 */
	@Override
	public void resetPassword(SsoPo ssoPo) throws Parameter_Exception {
		ssoClient.updatePasswordByName(ssoPo);
		if (10000 != ssoPo.getErrorCode()) {
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
	}

}
