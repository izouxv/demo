/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

/**
 * 
 */
package com.lanlian.chatting.metrics;

import java.util.SortedMap;
import java.util.concurrent.TimeUnit;

import org.springframework.context.annotation.Configuration;

import com.codahale.metrics.ConsoleReporter;
import com.codahale.metrics.Counter;
import com.codahale.metrics.MetricRegistry;
import com.ryantenney.metrics.spring.config.annotation.EnableMetrics;
import com.ryantenney.metrics.spring.config.annotation.MetricsConfigurerAdapter;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月8日 上午10:51:21
 * @explain
 */

@Configuration
@EnableMetrics
public class MetricsConfiguring extends MetricsConfigurerAdapter {


	@Override
	public void configureReporters(MetricRegistry metricRegistry) {
		System.err.println("**********************:" + metricRegistry.getCounters());
		SortedMap<String, Counter> map = metricRegistry.getCounters();
		System.out.println("map:" + map);
		TimeUnit timeUnit = TimeUnit.MINUTES;
		System.out.println("timeUnit:" + timeUnit);
		ConsoleReporter.forRegistry(metricRegistry).build().start(1, timeUnit);
		// PushGateway pushGateway = new PushGateway("");
	}

//	private static final MetricRegistry metricRegistry = new MetricRegistry();
//
//	private static Queue<String> queue = new LinkedBlockingDeque<String>();
//	public static Counter pendingJobs = new Counter();
//
//	public static void addGauges(final String name, String... names) {
//		System.err.println("**********************:" + metricRegistry);
//		// 实例化一个Gauge
//		Gauge<Integer> gauge = new Gauge<Integer>() {
//			@Override
//			public Integer getValue() {
//				queue.add(name);
//				return queue.size();
//			}
//		};
//		metricRegistry.register(MetricRegistry.name(name, names), gauge);
//		queue.add(name);
//	}
//
//	public static void addCounter(final String name, String... names) {
//		System.err.println("**********************:" + metricRegistry.getCounters().size());
//		String path = MetricRegistry.name(name, names);
//		SortedMap<String, Counter> map = metricRegistry.getCounters();
//		boolean flag = map.containsKey(path);
//		System.out.println("flag:" + flag);
//		Counter counter = null;
//		if (flag) {
//			System.out.println("counter1:" + counter);
//			counter = map.get(path);
//			counter.inc();
//		} else {
//			System.out.println("counter2:" + counter);
//			counter = new Counter();
//			counter.inc();
//			metricRegistry.register(path, counter);
//		}
//		SortedMap<String, Counter> maps = metricRegistry.getCounters();
//		for (String str : maps.keySet()) {
//			System.out.println("key:" + str + ",Counter:" + maps.get(str).getCount());
//		}
//	}
//
//	public static void addJob(String job) {
//		pendingJobs.inc();
//		queue.offer(job);
//
//	}
//
//	public static String takeJob() {
//		pendingJobs.dec();
//		return queue.poll();
//	}
//	
//	/**
//	 * 上报到9091
//	 */
//	public synchronized void appear(String url) {
//		Counter counter = null;
//		String key = null;
//		long value = 0L;
//		SortedMap<String, Counter> map = metricRegistry.getCounters();
//		try {
//			if (map.isEmpty()) {
//				return;
//			}
//			LogUtil.info("counter:" + PushGateway.instanceIPGroupingKey());
//			LogUtil.info(System.currentTimeMillis() + ",counter:" + counter.getCount());
//		} catch (Exception e) {
//			LogUtil.error("MetricsConfigServiceImpl - appear:metrics上报信息失败");
//			LogUtil.error(e);
//		} finally {
//			try {
////				PushGateway pg = new PushGateway(url);
////				pg.push(metricRegistry, "counter");
//			} catch (Exception e2) {
//				LogUtil.error("MetricsConfigServiceImpl - appear:metrics信息清除");
//				e2.printStackTrace();
//			}
//		}
//	}

}
