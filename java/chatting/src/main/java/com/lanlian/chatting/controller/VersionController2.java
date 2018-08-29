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
package com.lanlian.chatting.controller;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.VersionInfoService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.Version2;
import com.lanlian.rpc.version.GetLatestVersionRequest.Builder;

/**
 * @Title VersionController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月27日 下午8:47:55
 * @explain
 */
@Controller
@RequestMapping(value = RequestSetting.FILE_PARENT_V_1_1, produces = RequestSetting.PRODUCES)
public class VersionController2 extends MyAbstractController {

	private static Logger log = Logger.getLogger(VersionController2.class);

	@Resource
	VersionInfoService versionInfoService;

	/**
	 * 更新版本
	 * 
	 * @param device
	 * @param action
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_NEW_VERSION, method = RequestMethod.GET)
	public String downloadVersion(@PathVariable(value = "dev", required = true) String dev,
			@PathVariable(value = "code", required = true) String code, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		log.info("downloadVersion-dev:" + dev + ",code：" + code);
		ParameterVerify.verifyNull(code, dev);
		SsoPo ssoPo = getContextSsoPo(request, response);
		Builder builder = versionInfoService.getBuilder();
		builder.setSource(ssoPo.getSource());
		builder.setDevice(dev);
		Version2 version = versionInfoService.findNewVersion(builder);
		log.info("download-version:" + version);
		return JSON.toJson(version);
	}

}
