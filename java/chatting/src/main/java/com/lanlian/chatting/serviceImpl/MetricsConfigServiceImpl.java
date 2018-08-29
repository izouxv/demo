/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.serviceImpl;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;

import org.apache.commons.lang.math.NumberUtils;
import org.apache.log4j.Logger;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

import com.lanlian.chatting.service.MetricsConfigService;

import io.prometheus.client.CollectorRegistry;
import io.prometheus.client.Counter;
import io.prometheus.client.Gauge;
import io.prometheus.client.exporter.PushGateway;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月14日 上午8:56:57
 * @explain
 */

@Service(value = "metricsConfigServiceImpl")
public class MetricsConfigServiceImpl implements MetricsConfigService {

	private static Logger logger = Logger.getLogger(MetricsConfigServiceImpl.class);
	
	static CollectorRegistry registry;
	static Map<String, Integer> countersMap;
	static {
		registry = new CollectorRegistry();
		countersMap = new HashMap<>();
	}

	void executeBatchJob() {
		CollectorRegistry registry = new CollectorRegistry();
		Gauge duration = Gauge.build().name("my_job_duration").help("duration_seconds").register(registry);
		Gauge.Timer durationTimer = duration.startTimer();
		
		try {
			Gauge lastSuccess = Gauge.build().name("my_batch_job_last_success_unixtime").help("Last time my batch job succeeded, in unixtime.").register(registry);
			lastSuccess.setToCurrentTime();
		} catch (Exception e) {
			logger.error(e);
		} finally {
			durationTimer.setDuration();
			PushGateway pushGateway = new PushGateway("192.168.1.6:9091");
			try {
				pushGateway.pushAdd(registry, "my_batch_job");
			} catch (IOException e) {
				logger.error(e);
			}
		}
		
	}

	/**
	 * 计数
	 */
	@Async
	@Override
	public void interfaceCounter(String uri) {
		try {
			uri = uri.replaceAll("/", "_");
			uri = uri.replaceAll("\\.", "_");
			if (NumberUtils.isNumber(uri.substring(uri.lastIndexOf("_") + 1))) {
				uri = uri.substring(1, uri.lastIndexOf("_")+1);
			} else {
				if (uri.substring(uri.lastIndexOf("_") + 1).length() > 15) {
					uri = uri.substring(1, uri.lastIndexOf("_")+1);
				} else {
					uri = uri.substring(1);
				}
			}
			if (countersMap.containsKey(uri)) {
				countersMap.put(uri, countersMap.get(uri).intValue() + 1);
			} else {
				countersMap.put(uri, 1);
			}
		} catch (Exception e) {
			logger.fatal("interfaceCounter:"+e);
		}
	}

	/**
	 * 上报到9091
	 */
	@Override
	public synchronized void appear(String url) {
		Counter counter = null;
		String key = null;
		int value = 0;
		try {
			if (countersMap.isEmpty()) {
				return;
			}
			logger.info("countersMap:"+countersMap);
			for (Entry<String, Integer> counterMap : countersMap.entrySet()) {
				key = counterMap.getKey();
				value = counterMap.getValue();
				if (key == null || key.trim().isEmpty()) {
					logger.info("数据为空，9091");
					return;
				}
				counter = Counter.build().name(key).help("help").register(registry);
				counter.inc(value);
			}
		} catch (Exception e) {
			logger.fatal("MetricsConfigServiceImpl - appear:metrics上报信息失败:"+e);
		} finally {
			try {
				if (countersMap.isEmpty()) {
					return;
				}
				PushGateway pg = new PushGateway(url);
				pg.pushAdd(registry, "chatting");
			} catch (Exception e2) {
				logger.fatal("MetricsConfigServiceImpl - appear:metrics上报信息失败:"+e2);
			}
			logger.info("metrics clear");
			countersMap.clear();
			registry.clear();
		}
	}

}
