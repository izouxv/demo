/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.service;

import java.util.List;

import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.result.Parameter_Exception;

/** 
 * @Title DadaGroupService.java
 * @Package com.lanlian.chatting.service
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午4:36:12
 * @explain 信息实时上报群设置接口业务层
 */

public interface DadaGroupService {
	
	/**
	 * 查询群id是否存在
	 * @param dadaGroupPo
	 * @return DadaGroupPo
	 */
	DadaGroupPo findDadaGroup(DadaGroupPo dadaGroupPo);
	
	/**
	 * 查询实时上报群与开关
	 * @param dadaGroupPo
	 * @throws Parameter_Exception 
	 */
	public DadaGroupPo findDadaGroupSwitchingUser(DadaGroupPo dadaGroupPo) throws Parameter_Exception;
	
	/**
	 * 创建实时上报群
	 * @param dadaGroupPo
	 * @throws Parameter_Exception 
	 */
	DadaGroupPo saveDadaGroup(DadaGroupPo dadaGroupPo, List<Integer> list, String source) throws Parameter_Exception;

	/**
	 * 创建实时上报群开关
	 * 
	 * @param dadaGroupPo
	 * @throws Parameter_Exception
	 */
	void saveDadaGroupSwitching(DadaGroupPo dadaGroupPo) throws Parameter_Exception;
	
	/**
	 * 存入邀请码
	 * @param dadaGroupPo
	 */
	void saveDadaGroupCode(DadaGroupPo dadaGroupPo);
	
	/**
	 * 修改实时上报群开关
	 * @param dadaGroupPo
	 */
	DadaGroupPo updateDadaGroup(DadaGroupPo dadaGroupPo);

	
}

