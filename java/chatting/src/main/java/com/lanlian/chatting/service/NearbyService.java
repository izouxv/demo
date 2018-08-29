/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.service;

import java.util.List;

import com.lanlian.chatting.po.virtual.NearbyPo;
import com.lanlian.chatting.po.virtual.NearbyGroupInfoPo;
import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月19日 下午7:25:08
 * @explain 
 */

public interface NearbyService {
	
	/**
	 * 生成一个临时用户id
	 * @param nearbyPo
	 * @return
	 * @throws Parameter_Exception
	 */
	void getTid(NearbyPo nearbyPo) throws Parameter_Exception;
	
	/**
	 * 生成一个群id
	 * @return
	 * @throws Parameter_Exception 
	 */
	void getTgid(NearbyPo nearbyPo) throws Parameter_Exception;
	
	/**
	 * 开启用户信息并获取附近人的信息
	 * @param localtionPo
	 * @return
	 * @throws Parameter_Exception
	 */
	List<NearbyPo> nearbyLocaltion(NearbyPo localtionPo) throws Parameter_Exception;
	
	/**
	 * 关闭用户信息并获取附近人的信息
	 * @param localtionPo
	 * @return
	 * @throws Parameter_Exception
	 */
	void deleteNearbyLocaltion(NearbyPo localtionPo) throws Parameter_Exception;

	/**
	 * 对附近的人发送群操作消息
	 * @param nearbyPo
	 * @return 
	 * @throws Parameter_Exception
	 */
	NearbyPo sendNearbyInfo(NearbyGroupInfoPo nearbyGroupInfoPo, NearbyPo nearbyPo, int action) throws Parameter_Exception;
	
	/**
	 * 对附近的人发送消息
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	void sendNearbyInfo(NearbyGroupInfoPo nearbyGroupInfoPo) throws Parameter_Exception;
	
	/**
	 * 收取自己的消息
	 * 
	 * @param nearbyPo
	 * @throws Parameter_Exception 
	 */
	List<NearbyGroupInfoPo> getNearbyInfo(NearbyGroupInfoPo nearbyPo) throws Parameter_Exception;
	
}

