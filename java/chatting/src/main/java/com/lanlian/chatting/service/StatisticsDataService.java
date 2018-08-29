/** 
 *<p>开发公司：	鹏联优思 <p>
 *<p>版权所有：	鹏联优思 <p>
 *<p>责任人：	王东阳    <p> 
 *<p>网址：www.penslink.com <p>
 */

package com.lanlian.chatting.service;

import java.io.IOException;

import com.lanlian.chatting.model.UserHeartbeatsModel;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.po.UserHeartbeatPO;
import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @author  wangdyq
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年8月2日 下午2:14:48
 * @explain 统计信息业务
 */

public interface StatisticsDataService {
	
	/**
	 * 上报心与获取心跳
	 * @param userHeartbeatPO
	 * @param info 
	 * @return
	 * @throws Parameter_Exception 
	 * @throws IOException 
	 */
	UserHeartbeatsModel heartbeat(SsoPo ssoPo,UserHeartbeatPO userHeartbeatPO, String info) throws Parameter_Exception;
	
}
