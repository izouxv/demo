package com.lanlian.chatting.rpc;
/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

import java.io.BufferedOutputStream;
import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.concurrent.TimeUnit;

import org.apache.log4j.Logger;

import com.lanlian.rpc.adver.AdvertisementGrpc;
import com.lanlian.rpc.adver.AdvertisementGrpc.AdvertisementBlockingStub;
import com.lanlian.rpc.adver.AdvertisementReply;
import com.lanlian.rpc.adver.AdvertisementRequest;
import com.lanlian.rpc.adver.AdvertisementRequest.Builder;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

/**
 * 获取广告信息
 * 
 * @author wangdyq
 *
 */
public class AdverClient {

	private static Logger logger = Logger.getLogger(AdverClient.class);

	private static ManagedChannel channel;

	private static AdvertisementBlockingStub blockingStub;

	/**
	 * 构造
	 * 
	 * @param host
	 * @param port
	 */
	private AdverClient(String host, int port) {
		channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext(true).build();
		blockingStub = AdvertisementGrpc.newBlockingStub(channel).withDeadlineAfter(60, TimeUnit.DAYS);
	}

	/**
	 * 获取一个builder
	 * 
	 * @return
	 */
	public Builder getBuilder() {
		Builder builder = AdvertisementRequest.newBuilder();
		return builder;
	}

	/**
	 * 获取最新版本
	 * 
	 * @param source
	 * @param username
	 * @return
	 * @throws Parameter_Exception
	 */
	public AdvertisementReply getAdver(Builder builder) {
		AdvertisementRequest request = builder.build();
		builder.clear();
		logger.info("getNewVersion-request:" + request);
		AdvertisementReply rpcReply = blockingStub.getAdvertisement(request);
		logger.info("getNewVersion-rpcReply:" + rpcReply);
		return rpcReply;
	}

	/**
	 * 关闭请求连接
	 * 
	 * @throws InterruptedException
	 */
	public void shutdown() {
		try {
			channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
		} catch (InterruptedException e) {
			System.out.println(e);
		} finally {
			channel.shutdownNow();
		}
	}

	public static void main(String[] args) throws InterruptedException {
		AdverClient clientRpc = new AdverClient("192.168.1.51", 7002);
		Builder builder = clientRpc.getBuilder();
		builder.setSource("AQIDAA==");
		clientRpc.getAdver(builder);
		clientRpc.shutdown();
	}

	public void writeFile(byte[] bfile, String filePath, String fileName) {
		BufferedOutputStream bos = null;
		File file = null;
		try {
			file = new File(filePath + "/" + fileName);
			System.err.println(filePath + "/" + fileName);
			File dir = new File(filePath);
			if (!dir.exists() && dir.isDirectory()) {// 判断文件目录是否存在
				dir.mkdirs();
			}
			bos = new BufferedOutputStream(new FileOutputStream(file, true));
			bos.write(bfile);
			bos.write("\r\n".getBytes());
			bos.flush();
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			if (bos != null) {
				try {
					bos.close();
				} catch (IOException e1) {
					e1.printStackTrace();
				}
			}

		}
	}

}
