//package com.lanlian.chatting.interceptor;
//
//import java.io.File;
//import java.io.FileOutputStream;
//import java.io.PrintWriter;
//import java.sql.Timestamp;
//import java.text.SimpleDateFormat;
//import java.util.concurrent.BlockingQueue;
//import java.util.concurrent.LinkedBlockingQueue;
//
//import org.aspectj.lang.ProceedingJoinPoint;
//import org.aspectj.lang.Signature;
//import org.aspectj.lang.annotation.Around;
//import org.aspectj.lang.annotation.Aspect;
//import org.springframework.beans.factory.annotation.Value;
//import org.springframework.stereotype.Component;
//
//import com.lanlian.chatting.logger.LogUtil;
//
///**
// * @author  wdyqxx
// * @version 2017年1月13日 下午2:43:16
// * @explain 此类用于记录保存业务方法执行的时间；
// * 	aop.properties
// */
//@Aspect
//@Component
//public class TimeLogerAspect {
//	
//	private File file;
//	
//	private Thread writer;
//	/**
//	 * 解析XML,获取类名，根据类名创建Bean
//	 */
//	private BlockingQueue<String> queue=new LinkedBlockingQueue<String>(500);
//	
//	private String filename;
//	
//	@Value("#{aop.filename}")
//	public void setFilename(String filename){
//		this.filename=filename;
//		file=new File(this.filename);
//	}
//	
//	public TimeLogerAspect() {
//		writer = new Thread(){
//			public void run() {
//				while (true) {
//					try {
//						if (queue.isEmpty()) {
//							Thread.sleep(500);
//							continue;
//						}
//						/**
//						 * 当队列中有数据时；
//						 */
//						PrintWriter out=new PrintWriter(new FileOutputStream(file,true));
//						while (!queue.isEmpty()) {
//							String str=queue.poll();
//							out.println(str);
////							System.out.println("TimeLogerAspect()");
//						}
//						out.close();
//					} catch (Exception e) {
//						e.printStackTrace();
//						LogUtil.error(e);
//					}
//				}
//			}
//		};
//		writer.start();
//	}
//	
//	@Around("execution(* cn.lanlian.chatting.serviceImpl.*Service.*(..))")
//	public Object proc(ProceedingJoinPoint joinPoint) throws Throwable{
//		long t1=System.nanoTime();
//		long t2=System.nanoTime();
//		Object val= joinPoint.proceed();
//		//获取类型
//		Signature sig=joinPoint.getSignature();
//		System.out.println(System.currentTimeMillis()+":"+sig+":"+(t2-t1));
//		Timestamp now=new Timestamp(System.currentTimeMillis());
//		SimpleDateFormat sdf=new SimpleDateFormat("yyyy年MM月dd日 hh:mm:ss(sss)");
//		String str="时间："+sdf.format(now)+",返回值类型："+sig+",用时："+(t2-t1)+"ms!!!";
//		queue.offer(str);
//		return val; 
//	}
//	
//	
//}
//
//
//
