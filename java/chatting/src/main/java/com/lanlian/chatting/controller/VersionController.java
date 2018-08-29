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

import java.io.IOException;
import java.util.Arrays;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Fatal_Exception;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.FeedBackClient;
import com.lanlian.chatting.service.VersionInfoService;
import com.lanlian.chatting.util.DataFinals.*;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.Version;
import com.lanlian.rpc.feedback.AddFeedbackReply;
import com.lanlian.server.http.Notification;
import com.lanlian.server.redis.RedisClient79;

/**
 * @Title VersionController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月27日 下午8:47:55
 * @explain
 */
@Controller
@RequestMapping(value = RequestSetting.FILE_PARENT_V, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class VersionController extends MyAbstractController {

	private static Logger log = Logger.getLogger(VersionController.class);

	@Resource
	VersionInfoService versionInfoService;

	/**
	 * 处理上传文件控制器
	 * 
	 * @param device/garcat/tomcatII/
	 * @param name
	 * @param code
	 * @param upMd5
	 * @param description
	 * @param file
	 * @param request
	 * @return
	 * @throws IOException
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_UPLOAD_VERSION)
	public String uploadFileVersion(@PathVariable String device,
			@RequestParam(value = "fileName", required = true) String fileName,
			@RequestParam(value = "file", required = true) MultipartFile file,
			@RequestParam(value = "versionName", required = true) String versionName,
			@RequestParam(value = "versionCode", required = true) String versionCode,
			@RequestParam(value = "upMd5", required = true) String upMd5,
			@RequestParam(value = "description", required = true) String description,
			@RequestParam(value = "description1", required = true) String description1, HttpServletRequest request,
			HttpServletResponse response) throws IOException, Parameter_Exception {
		log.info("uploadFileVersion-device:" + device + ",filename:" + fileName + ",name:" + versionName + ",code:"
				+ versionCode + ",upMd5:" + upMd5 + ",description:" + description + ",description1:" + description1);
		ParameterVerify.verifyNull(versionName, versionCode, upMd5, description);
		device = device.trim();
		versionName = versionName.trim();
		versionCode = versionCode.trim();
		upMd5 = upMd5.trim();
		description = description.trim();
		description1 = description1.trim();
		if (upMd5.trim().length() != 32) {
			throw new Parameter_Exception(21002);
		}
		// 判断单文件MultipartFile是否为空
		if (file == null || file.isEmpty()) {
			throw new Parameter_Exception(21006);
		}
		// 保存版本信息并存储文件
		Version info = new Version();
		info.setDevice(device);
		info.setVersionName(versionName);
		info.setVersionCode(versionCode);
		info.setMd5(upMd5);
		info.setDescription(description);
		info.setDescription1(description1);
		info.setLength(String.valueOf(file.getSize()));
		versionInfoService.saveVersionInfo(info, file);
		log.info(JSON.toJson());
		return JSON.toJson();
	}

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
	@RequestMapping(value = RequestSetting.FILE_BODY_DOWNLOAD_VERSION, method = RequestMethod.GET)
	public String downloadVersion(@PathVariable(value = "action", required = true) String action,
			@PathVariable(value = "device", required = true) String device, HttpServletResponse response)
			throws Parameter_Exception {
		log.info("downloadVersion-device:" + device + ",action：" + action);
		ParameterVerify.verifyNull(action, device);
		// 对请求处理
		Version version = new Version();
		version.setDevice(device);
		version = versionInfoService.findVersionInfo(action, version, response);
		log.info("download-action:" + action + ",version:" + version);
		if (Actions.contains(action) && version != null) {
			return JSON.toJson(version);
		}
		response.setStatus(404);
		return null;
	}

	@Resource
	Notification notification;

	@Resource
	FeedBackClient feedBackClient;

	@Resource
	RedisClient79 redisClient79;

	/**
	 * 反馈意见
	 * 
	 * @throws Fatal_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.FILE_BODY_FEEDBACK, method = RequestMethod.POST)
	public String feedback(@RequestParam(value = "description", required = true) String description,
			@RequestParam(value = "mobileInfo", required = true) String mobileInfo,
			@RequestParam(value = "appInfo", required = true) String appInfo,
			@RequestParam(value = "files", required = false) String[] files,
			@RequestParam(value = "contact", required = false) String contact,
			@RequestParam(value = "deviceInfo", required = false) String deviceInfo,
			@RequestParam(value = "userInfo", required = false) String userInfo,
			@RequestParam(value = "extendInfo", required = false) String extendInfo, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception, Fatal_Exception {
		// 必传参数非空
		log.info("feedback:" + description + "\r\n" + mobileInfo + "\r\n" + appInfo + "\r\n" + contact + "\r\n"
				+ deviceInfo + "\r\n" + userInfo + "\r\n" + extendInfo);
		if (files != null) {
			log.info("files:" + Arrays.asList(files));
		}
		ParameterVerify.verifyNull(description, mobileInfo, appInfo);
		SsoPo ssoPo = getContextSsoPo(request, response);
		AddFeedbackReply reply = feedBackClient.addFeedBack(ssoPo.getSource(), description, mobileInfo, appInfo,
				deviceInfo, userInfo, extendInfo, files, contact);
		notification.sendMail(redisClient79.getString(redisClient79.key14), reply, description, mobileInfo, appInfo,
				deviceInfo, userInfo, extendInfo, files, contact);
		return JSON.toJson();
	}

}
