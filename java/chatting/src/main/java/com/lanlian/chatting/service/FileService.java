/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.lanlian.chatting.service;

import com.lanlian.chatting.po.AdverPo;
import com.lanlian.chatting.po.AdvertisementPo;
import com.lanlian.chatting.result.Parameter_Exception;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月15日 上午10:21:48
 * @explain 文件设置的业务层
 */

public interface FileService {

	/**
	 * 广告信息存入
	 * @param advertisementPo
	 * @return
	 */
	void setAdvertisement(AdvertisementPo advertisementPo);

	/**
	 * 获取广告信息
	 * @param advertisementPo
	 * @return
	 * @throws Parameter_Exception 
	 */
	AdvertisementPo getAdvertisement(AdvertisementPo advertisementPo) throws Parameter_Exception;
	
	/**
	 * 修改广告信息
	 * @param advertisementPo
	 * @return
	 */
	AdvertisementPo updateAdvertisement(AdvertisementPo advertisementPo);

	/**
	 * 获取新的广告信息
	 * @param advertisementPo
	 * @return
	 * @throws Parameter_Exception
	 */
	AdverPo getNewAdver(AdverPo adverPo) throws Parameter_Exception;


	
}

