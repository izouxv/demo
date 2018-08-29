/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.controller;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月19日 下午5:39:03
 * @explain
 */
@Controller
@RequestMapping(value = RequestSetting.USER_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class StatisticsData extends MyAbstractController {
	
	@Resource
	SsoService ssoService;

	/**
	 * 信息上报、统计
	 * 
	 * @param types
	 * @param imei
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.LOCALTIONDATA_BODY)
	public String localtion(
			@PathVariable(value = "types", required = true) String types,
			@RequestParam(value = "data", required = true) String data, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		LogUtil.info("localtion-types:" + types);
		SsoPo ssoPo = getContextSsoPo(request, response);
		boolean flag = ("chatting_equipment_location".equals(types) || "chatting_equipment_dynamic".equals(types)
				|| "chatting_equipment_static".equals(types));
		if (flag) {
			JSONObject jsonObject = null;
			try {
				jsonObject = JSONObject.parseObject(data);
				jsonObject.put("uid", ssoPo.getUid());
				jsonObject.put("source", ssoPo.getSource());
			} catch (Exception e) {
				throw new Parameter_Exception(21002);
			}
			ssoService.analysis(jsonObject.toJSONString());
			return JSON.toJson();
		}
		return JSON.toJson(21002);
	}
	/*
	 * String name = null;//TOMCAT 设备名称 String sn = null;// 234512334467//设备序列号
	 * String model = null;//RADACAT//设备型号 String manufacturer = null;//北京蓝涟科技//生产厂商
	 * String frequency = null;//23HZ~60HZ//频段支持 String mac =
	 * null;//ef:52:e4:a2:25:89//蓝牙MAC地址 String version = null;//V1.0.0//固件版本 String
	 * md5 = null;//jfieh32432ji8hbddw//MD5加密 String electricity = null;//78//电池剩余电量
	 * String temperature = null;//40//温度 String rssi = null;//jife//环境RSSI值 String
	 * buffer = null;//38//收发的缓存区状态 String srceen = null;//t//屏幕工作是否正常 String fsk =
	 * null;//无线模块FSK是否正常 String lora = null;//lora是否正常
	 */
}