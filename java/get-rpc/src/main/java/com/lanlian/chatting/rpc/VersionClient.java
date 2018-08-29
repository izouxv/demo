///** 
// *<p>开发公司：		蓝涟科技 <p>
// *<p>版权所有：		蓝涟科技 <p>
// *<p>责任人：		      王东阳<p> 
// *<p>网址：www.radacat.com <p>
// * @author wangdyqxx
// * @email wangdy@radacat.com
// */
//
//package com.lanlian.chatting.rpc;
//
//import java.util.concurrent.TimeUnit;
//
//import org.apache.log4j.Logger;
//
//import com.lanlian.chatting.result.Parameter_Exception;
//import com.lanlian.rpc.version.GetLatestVersionRequest;
//import com.lanlian.rpc.version.GetLatestVersionRequest.Builder;
//import com.lanlian.rpc.version.GetLatestVersionResponse;
//import com.lanlian.rpc.version.RadacatVersionGrpc;
//import com.lanlian.rpc.version.RadacatVersionGrpc.RadacatVersionBlockingStub;
//
//import io.grpc.ManagedChannel;
//import io.grpc.ManagedChannelBuilder;
//
///**
// * @Title VersionClient.java
// * @Package com.lanlian.chatting.rpc.clientrpc
// * @author 王东阳
// * @version V1.0
// * @date 2018年1月10日 下午20:38:18
// * @explain
// */
//
//public class VersionClient {
//	
//	private static Logger log = Logger.getLogger(VersionClient.class);
//
//	/**
//	 * feedbackChannel初始化请求通道
//	 */
//	private static ManagedChannel channel;
//
//	/**
//	 * feedbackBlockingStub信息
//	 */
//	private static RadacatVersionBlockingStub versionBlockingStub;
//
//	/**
//	 * 构造
//	 * 
//	 * @param host
//	 * @param port
//	 */
//	private VersionClient(String host, int port) {
//		channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext(true).build();
//		versionBlockingStub = RadacatVersionGrpc.newBlockingStub(channel).withDeadlineAfter(60, TimeUnit.DAYS);
//	}
//
//	/**
//	 * 获取一个builder
//	 * 
//	 * @return
//	 */
//	public Builder getBuilder() {
//		Builder builder = GetLatestVersionRequest.newBuilder();
//		return builder;
//	}
//
//	/**
//	 * 获取最新版本
//	 * @param source
//	 * @param device
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public GetLatestVersionResponse getNewVersion(Builder builder) {
//		GetLatestVersionRequest request = builder.build();
//		builder.clear();
//		log.info("getNewVersion-request:" + request);
//		GetLatestVersionResponse rpcReply = versionBlockingStub.getLatestVersion(request);
//		log.info("getNewVersion-rpcReply:" + rpcReply);
//		return rpcReply;
//	}
//
//	/**
//	 * 关闭请求连接
//	 * 
//	 * @throws InterruptedException
//	 */
//	public void shutdown() {
//		try {
//			channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
//		} catch (InterruptedException e) {
//			log.error(e);
//		} finally {
//			channel.shutdownNow();
//		}
//	}
//
//	public static void main(String[] args) throws InterruptedException, Parameter_Exception {
//		VersionClient clientRpc = new VersionClient("120.77.66.96", 7002);
//		Builder builder = clientRpc.getBuilder();
//		builder.setSource("AQIDAA==");
//		builder.setDevice("dacat");
//		clientRpc.getNewVersion(builder);
//		clientRpc.shutdown();
//	}
//
//}
