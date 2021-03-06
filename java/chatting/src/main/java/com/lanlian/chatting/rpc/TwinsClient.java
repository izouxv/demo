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

import org.apache.log4j.Logger;

import com.lanlian.rpc.twins.AddTwinsAgentRequest;
import com.lanlian.rpc.twins.AddTwinsAgentRequest.Builder;
import com.lanlian.rpc.twins.AddTwinsAgentResponse;
import com.lanlian.rpc.twins.TwinsAgentServerGrpc;
import com.lanlian.rpc.twins.TwinsAgentServerGrpc.TwinsAgentServerBlockingStub;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

/**
 * @Title TwinsClient.java
 * @Package com.lanlian.chatting.rpc
 * @author 王东阳
 * @version V1.0
 * @date 2018年1月31日 下午2:38:18
 * @explain
 */

public class TwinsClient {

	private static Logger logger = Logger.getLogger(TwinsClient.class);

	private static ManagedChannel channel = null;
	private static TwinsAgentServerBlockingStub blockingStub = null;

	/**
	 * 构造
	 * 
	 * @param host
	 * @param port
	 * @throws InterruptedException 
	 */
	private TwinsClient(String host, int port) throws InterruptedException {
		channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext(true).build();
	}

	/**
	 * 获取一个builder
	 * 
	 * @return
	 */
	private Builder getBuilder() {
		blockingStub = TwinsAgentServerGrpc.newBlockingStub(channel);
		Builder builder = AddTwinsAgentRequest.newBuilder();
		return builder;
	}

	/**
	 * 统计数据上报
	 * 
	 * @param value
	 */
	public void reported(String value) {
		try {
			Builder builder = getBuilder();
			builder.setReported(value);
			AddTwinsAgentRequest request = builder.build();
			builder.clear();
			logger.info("reported-request:" + request);
			AddTwinsAgentResponse reply = blockingStub.addTwinsAgent(request);
			logger.info("reported-addTwins:" + reply);
		} catch (Exception e) {
			logger.fatal("reported-e:" + e);
		}
	}

	/**
	 * 关闭请求连接
	 * 
	 * @throws InterruptedException
	 */
	public void shutdown() {
		try {
			channel.shutdown().awaitTermination(50, TimeUnit.MILLISECONDS);
		} catch (InterruptedException e) {
			logger.error(e);
		} finally {
			channel.shutdownNow();
		}
	}

	/**
	 * 120.76.54.242
	 * 
	 * @param args
	 * @throws InterruptedException
	 */
	public static void main(String[] args) throws InterruptedException {
		TwinsClient clientRpc = new TwinsClient("120.76.54.242", 7011);
		String statis = "{ \"source\":\"AQIDAA==\", \"uid\":12345600, \"sn\":\"L0104AA174\","
				+ "  \"name\":\"Test\", \"model\":\"Tomcat11\", \"manufacturer\":\"北京蓝涟科技有限责任公司\","
				+ "  \"frequency\":\"40HZ\", \"mac\":\"ab:33:7f:f5:c3\","
				+ "  \"md5\":\"e10adc3949ba59abbe56e057f20f883e\"}";
		clientRpc.reported(statis);
		clientRpc.shutdown();
	}

}
