/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.service;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;

/** 
 * @Title SsoService.java
 * @Package com.lanlian.chatting.service
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月8日 下午3:46:12
 * @explain 
 */

public interface SsoService {
	
	
	/**
	 * 检验source
	 * @param ssoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	public SsoPo getSource(SsoPo ssoPo) throws Parameter_Exception;
	
	/**
	 * 验证头信息
	 * @param request
	 * @return
	 * @throws Parameter_Exception
	 */
	public SsoPo verifyToken(SsoPo ssoPo) throws Parameter_Exception;
	
	/**
	 * 检验用户是否被注册
	 * @param ssoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	int checkUser(SsoPo ssoPo) throws Parameter_Exception;
	
	/**
	 * 添加用户
	 * @param ssoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	SsoPo add(SsoPo ssoPo) throws Parameter_Exception;
	
	/**
	 * 找回密码发送邮件
	 * @param ssoPo
	 * @throws Parameter_Exception 
	 */
	void sendEmail(SsoPo ssoPo) throws Parameter_Exception;
	
	/**
	 * 发送短信
	 * @param ssoPo
	 * @throws Parameter_Exception 
	 */
	void sendMobile(SsoPo ssoPo, String action) throws Parameter_Exception;
	
	/**
	 * 校验短信
	 * 
	 * @param ssoPo
	 * @throws Parameter_Exception
	 */
	void verifyMobile(String username, String code, int action, String source) throws Parameter_Exception;

	/**
	 * 
	 * @param json
	 */
	void analysis(String json);
}

