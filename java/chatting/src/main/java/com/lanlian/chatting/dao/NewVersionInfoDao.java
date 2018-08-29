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
package com.lanlian.chatting.dao;

import com.lanlian.chatting.po.VersionPO;

/** 
 * @Title NewVersionInfoDao.java
 * @Package cn.lanlian.ccat.dao
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月28日 下午3:29:38
 * @explain 用于存取版本文件的信息
 */

public interface NewVersionInfoDao {
	
	/**
	 * 存储最新版本信息
	 * @param version
	 * @return int
	 */
	int saveVersionInfo(VersionPO ver);
	
	/**
	 * 查询最新版本信息
	 * @param version
	 * @return Version
	 */
	VersionPO findVersionInfo(VersionPO ver);
	
	
	

}

