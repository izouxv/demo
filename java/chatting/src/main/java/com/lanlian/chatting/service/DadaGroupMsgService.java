/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.service;

import com.alibaba.fastjson.JSONArray;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;

/** 
 * @Title DadaGroupMsgService.java
 * @Package com.lanlian.chatting.service
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午4:36:12
 * @explain 信息实时上报群设置接口业务层
 */

public interface DadaGroupMsgService {
	
	/**
	 * 查询打开实时群信息上报开关的用户信息
	 * @return
	 */
	DadaGroupPo findDadaGroupSwitchingUser(DadaGroupPo dadaGroupPo);
	
	/**
	 * 将同步微信信息保存
	 * @param groupMessageInfo
	 * @param ssoPo
	 * @param dadaGroupPo
	 */
	void saveGroupMessage(JSONArray jsonArray,SsoPo ssoPo, DadaGroupPo dadaGroupPo)throws Parameter_Exception;	
	
	
}

