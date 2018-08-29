/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.dao;

import java.util.List;

import com.lanlian.chatting.po.DadaGroupMsgPo;

/**
 * @Title DadaGroupMsgDao.java
 * @Package com.lanlian.chatting.dao
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午1:57:37
 * @explain 信息实时上报群消息接口
 */

public interface DadaGroupMsgDao {

	/**
	 * 将JSONArray的实时群聊记录存入
	 * @param DadaGroupMsgPo
	 */
	void saveGroupMsg(List<DadaGroupMsgPo> list);

}
