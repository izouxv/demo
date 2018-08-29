/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.task;

import javax.annotation.Resource;

import org.apache.log4j.Logger;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;

import com.lanlian.chatting.service.MetricsConfigService;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月11日 下午5:37:16
 * @explain
 */

@Service("timer")
public class Timer {
	
	private static Logger LOGGER = Logger.getLogger(Timer.class);
	
	@Resource
	MetricsConfigService metricsConfigService;
	
	@Scheduled(cron = "0 0/5 * * * *")
	public void timingMetrics() {
		LOGGER.info("9091...");
		metricsConfigService.appear("push9091.prometheus.radacat.com:9091");
	}
}
