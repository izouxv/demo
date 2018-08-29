/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.service;

import javax.servlet.http.HttpServletResponse;

import org.springframework.web.multipart.MultipartFile;

import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.vo.Version;
import com.lanlian.chatting.vo.Version2;
import com.lanlian.rpc.version.GetLatestVersionRequest.Builder;

/** 
 * @Title VersionInfoService.java
 * @Package cn.lanlian.ccat.service
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月28日 下午3:47:58
 * @explain 用于存取版本文件的信息
 */

public interface VersionInfoService {
	
	Builder getBuilder();
	
	/**
	 * 上传最新版本
	 * @param version
	 * @param file
	 * @throws Parameter_Exception
	 */
	void saveVersionInfo(Version version, MultipartFile file) throws Parameter_Exception;
	
	/**
	 * 下载最新版本
	 * @param action
	 * @param version
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	Version findVersionInfo(String action, Version version,
			HttpServletResponse response) throws Parameter_Exception;
	
	/**
	 * 获取最新版本信息
	 * @param builder
	 * @return
	 * @throws Parameter_Exception
	 */
	Version2 findNewVersion(Builder builder) throws Parameter_Exception;
	
}

