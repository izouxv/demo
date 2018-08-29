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
package com.lanlian.chatting.service;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @Title UpdatePwdService.java
 * @Package cn.lanlian.chatting.service
 * @author 王东阳
 * @version V1.0.3
 * @date 2017年5月3日 上午11:08:22
 * @explain 变更密码
 */

public interface UpdatePwdService {
	
	/**
	 * 校验密码
	 * 
	 * @throws Parameter_Exception
	 */
	void checkPassword(SsoPo ssoPo) throws Parameter_Exception;

	/**
	 * 修改密码
	 * 
	 * @throws Parameter_Exception
	 */
	void updatePassword(SsoPo ssoPo) throws Parameter_Exception;
	
	/**
	 * 重置密码
	 * @param ssoPo
	 * @throws Parameter_Exception
	 */
	void resetPassword(SsoPo ssoPo) throws Parameter_Exception;
	
	
}
