/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.rpc;

import java.util.concurrent.TimeUnit;

import com.lanlian.rpc.feedback.AddFeedbackReply;
import com.lanlian.rpc.feedback.AddFeedbackRequest;
import com.lanlian.rpc.feedback.AddFeedbackRequest.Builder;
import com.lanlian.rpc.feedback.FeedBackGrpc;
import com.lanlian.rpc.feedback.FeedBackGrpc.FeedBackBlockingStub;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

/**
 * @Title FeedBackClient.java
 * @Package com.lanlian.chatting.rpc.clientrpc
 * @author 王东阳
 * @version V1.0
 * @date 2018年1月10日 下午20:38:18
 * @explain
 */

public class FeedBackClient {
	
	/**
	 * feedbackChannel初始化请求通道
	 */
	private static ManagedChannel feedbackChannel;

	/**
	 * feedbackBlockingStub信息
	 */
	private static FeedBackBlockingStub feedbackBlockingStub;

	/**
	 * 构造
	 * 
	 * @param host
	 * @param port
	 */
	private FeedBackClient(String host, int port) {
		feedbackChannel = ManagedChannelBuilder.forAddress(host, port).usePlaintext(true).build();
		feedbackBlockingStub = FeedBackGrpc.newBlockingStub(feedbackChannel).withDeadlineAfter(60, TimeUnit.DAYS);
	}

	/**
	 * 获取一个builder
	 * 
	 * @return
	 */
	private Builder getBuilder() {
		Builder builder = AddFeedbackRequest.newBuilder();
		return builder;
	}

	/**
	 * 意见反馈
	 * @param statis
	 */
	public AddFeedbackReply addFeedBack(String source, String description,String mobileInfo, String appInfo, String deviceInfo, 
			String userInfo, String extendInfo,String[] files,String contact) {
		try {
			Builder builder = getBuilder();
			builder.setSource(source);
			builder.setDescription(description);
			builder.setMobileInfo(mobileInfo);
			builder.setAppInfo(appInfo);
			if (files != null) {
				for (String file : files) {
					builder.addFiles(file);				
				}				
			}
			if (contact != null) {
				builder.setContact(contact);				
			}
			if (deviceInfo != null) {
				builder.setDeviceInfo(deviceInfo);				
			}
			if (userInfo != null) {
				builder.setUserInfo(userInfo);				
			}
			if (extendInfo != null) {
				builder.setExtendInfo(extendInfo);				
			}
			AddFeedbackRequest request = builder.build();
			builder.clear();
			System.out.println("addFeedBack-request:" + request);
			AddFeedbackReply rpcReply = feedbackBlockingStub.addFeedback(request);
			System.out.println("addFeedBack-rpcReply:" + rpcReply);
			return rpcReply;
		} catch (Exception e) {
			System.out.println(e);
			return null;
		}
	}

	/**
	 * 关闭请求连接
	 * 
	 * @throws InterruptedException
	 */
	public void shutdown() {
		try {
			feedbackChannel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
		} catch (InterruptedException e) {
			System.out.println(e);
		} finally {
			feedbackChannel.shutdownNow();
		}
	}

	public static void main(String[] args) throws InterruptedException {
		FeedBackClient clientRpc = new FeedBackClient("192.168.1.51", 7002);
		System.out.println(clientRpc.addFeedBack("AQIDAA==", "test", "", "test", ""," test", "logFileled='http://file.radacat.com:88/v1.0/file/8a0a4719c3b66b2b926ffe5b1eede605'", new String[]{"http://file.radacat.com:88/v1.0/file/8a0a4719c3b66b2b926ffe5b1eede605"}, ""));
		clientRpc.shutdown();
	}

}
