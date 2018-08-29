///** 
// *<p>开发公司 :		          蓝涟科技 <p>
// *<p>版权所有 :		          蓝涟科技 <p>
// *<p>责任人     :		              王东阳 <p> 
// *<p>网址         :   www.radacat.com <p>
// *<p>邮箱         : wangdy@radact.com <p>
// */
//
//package com.lanlian.chatting.metrics;
//
//import org.slf4j.LoggerFactory;
//import org.springframework.context.annotation.Bean;
//import org.springframework.context.annotation.Configuration;
//
//import com.codahale.metrics.*;
//import com.codahale.metrics.ganglia.GangliaReporter;
//
//import info.ganglia.gmetric4j.gmetric.GMetric;
//import info.ganglia.gmetric4j.gmetric.GMetric.UDPAddressingMode;
//
//import java.io.IOException;
//import java.util.concurrent.TimeUnit;
//
///**
// * @author  王东阳
// * @version V1.0
// * @email wangdy@radacat.com
// * @date 2017年8月7日 上午11:58:37
// * @explain 
// */
//
//@Configuration
//public class MetricConfig {
//	
//	private static final MetricRegistry metricRegistry = new MetricRegistry();
//	
//	public static MetricRegistry config() {
//		return metricRegistry;
//	}
//
//    @Bean
//    public MetricRegistry metrics() {
//        return metricRegistry;
//    }
//
//    /**
//     * Reporter 数据的展现位置
//     *
//     * @param metrics
//     * @return
//     */
//    @Bean
//    public ConsoleReporter consoleReporter(MetricRegistry metrics) {
//        return ConsoleReporter.forRegistry(metrics)
//                .convertRatesTo(TimeUnit.SECONDS)
//                .convertDurationsTo(TimeUnit.MILLISECONDS)
//                .build();
//    }
//
//    @Bean
//    public Slf4jReporter slf4jReporter(MetricRegistry metrics) {
//        return Slf4jReporter.forRegistry(metrics)
//                .outputTo(LoggerFactory.getLogger("demo.metrics"))
//                .convertRatesTo(TimeUnit.SECONDS)
//                .convertDurationsTo(TimeUnit.MILLISECONDS)
//                .build();
//    }
//
//    @Bean
//    public JmxReporter jmxReporter(MetricRegistry metrics) {
//        return JmxReporter.forRegistry(metrics).build();
//    }
//
//    /**
//     * 自定义单位
//     *
//     * @param metrics
//     * @return
//     */
////    @Bean
////    public ListManager listManager(MetricRegistry metrics) {
////        return new ListManager(metrics);
////    }
//
//    /**
//     * TPS 计算器
//     *
//     * @param metrics
//     * @return
//     */
//    @Bean
//    public Meter requestMeter(MetricRegistry metrics) {
//        return metrics.meter("request");
//    }
//
//    /**
//     * 直方图
//     *
//     * @param metrics
//     * @return
//     */
//    @Bean
//    public Histogram responseSizes(MetricRegistry metrics) {
//        return metrics.histogram("response-sizes");
//    }
//
//    /**
//     * 计数器
//     *
//     * @param metrics
//     * @return
//     */
//    @Bean
//    public Counter pendingJobs(MetricRegistry metrics) {
//        return metrics.counter("requestCount---------");
//    }
//
//    /**
//     * 计时器
//     *
//     * @param metrics
//     * @return
//     */
//    @Bean
//    public Timer responses(MetricRegistry metrics) {
//        return metrics.timer("executeTime");
//    }
//    
//    public static void main(String[] args) {
//		try {
//			UDPAddressingMode mode = GMetric.UDPAddressingMode.MULTICAST;
//			GMetric gMetric = new GMetric("192.168.1.6", 9091, mode, 1);
//			System.out.println("--------");
//			Counter counter = metricRegistry.counter("hhhhhhhhhhh");
//			counter.inc(123456);
//			System.out.println("----"+metricRegistry.getCounters().get("hhhhhhhhhhh").getCount());
//			GangliaReporter gangliaReporter = GangliaReporter.forRegistry(metricRegistry)
//					.convertRatesTo(TimeUnit.SECONDS)
//					.convertDurationsTo(TimeUnit.MILLISECONDS).build(gMetric);
//			gangliaReporter.start(1, TimeUnit.SECONDS);
//			System.out.println("--------");
//		} catch (IOException e) {
//			e.printStackTrace();
//		}
//		
//	}
//
//}