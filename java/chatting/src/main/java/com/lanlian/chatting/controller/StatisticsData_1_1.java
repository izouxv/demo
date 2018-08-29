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

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONException;
import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.po.UserHeartbeatPO;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.service.StatisticsDataService;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @author 王东阳
 * @version V1.1
 * @email wangdy@radacat.com
 * @date 2018年1月30日 下午4:39:03
 * @explain
 */
@Controller
@RequestMapping(value = RequestSetting.STATISTICS,consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class StatisticsData_1_1 extends MyAbstractController {
	
	private static Logger logger = LoggerFactory.getLogger(StatisticsData_1_1.class);

	@Resource
	SsoService ssoService;
	
	@Resource
	StatisticsDataService statisticsDataService;

	/**
	 * 信息统计
	 * 
	 * @param types
	 * @param data
	 * @param request
	 * @param response
	 * @return
	 * @throws Parameter_Exception
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.DEVINFO_BODY)
	public String statistics(@RequestParam(value = "data", required = true) String data, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception {
		logger.info("statistics-data:" + data);
		try {
			SsoPo ssoPo = getContextSsoPo(request, response);
			JSONObject jsonObject = JSONObject.parseObject(data);
			if (!jsonObject.containsKey("sn")) {
				throw new Parameter_Exception(21002);
			}
			jsonObject.put("source", ssoPo.getSource());
			String token = request.getHeader("token");
			if (token != null) {
				ssoPo.setSessionName(token);
				ssoPo = ssoService.verifyToken(ssoPo);
				jsonObject.put("uid", ssoPo.getUid());
			}
			ssoService.analysis(jsonObject.toJSONString());
			return JSON.toJson();
		} catch (Exception e) {
			throw e;
		}
	}
	
	/**
	 * 上报心跳
	 * @param data
	 * @param request
	 * @param response
	 * @return
	 * @throws Exception 
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.HEARTBEAT_BODY)
	public String localtion(
			@RequestParam(value = "data", required = true) String data,
			@RequestParam(value = "info", required = false) String info,
			HttpServletRequest request, HttpServletResponse response) throws Exception {
		logger.info("localtion-data:" + data);
		try {
			SsoPo ssoPo = getContextSsoPo(request, response);
			UserHeartbeatPO userHeartbeatPO = JSONObject.parseObject(data, UserHeartbeatPO.class);
			if (PublicMethod.checkObjFieldIsNull(userHeartbeatPO)) {
				logger.info("localtion bean is null, userHeartbeatPO:",userHeartbeatPO);
				return JSON.toJson(21002);
			}
			return JSON.toJson(statisticsDataService.heartbeat(ssoPo, userHeartbeatPO,info));
		} catch (JSONException e) {
			throw new Parameter_Exception(21002);
		} catch (Exception e) {
			throw e;
		}
	}
}
