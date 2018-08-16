/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.plys.rtcs.mqtt;

import java.io.FileInputStream;
import java.io.IOException;
import java.nio.Buffer;
import java.nio.ByteBuffer;

import org.apache.log4j.Logger;
import org.eclipse.paho.client.mqttv3.IMqttDeliveryToken;
import org.eclipse.paho.client.mqttv3.MqttCallback;
import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttConnectOptions;
import org.eclipse.paho.client.mqttv3.MqttDeliveryToken;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;
import org.eclipse.paho.client.mqttv3.MqttPersistenceException;
import org.eclipse.paho.client.mqttv3.MqttTopic;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;
import org.springframework.web.socket.TextMessage;

import com.plys.rtcs.controller.VirtualWebsocket;
import com.plys.rtcs.po.ProtoMsg;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月2日 下午6:10:53
 * @explain 服务器向多个客户端推送主题，即不同客户端可向服务器订阅相同主题
 */

public class ServerMQTT {

	private static final Logger logger = Logger.getLogger(ServerMQTT.class);
	
	// tcp://MQTT安装的服务器地址
	public static final String HOST = "tcp://192.168.1.6:61613";

	private MqttClient client;
	private MqttTopic topic;
	private String userName = "admin";
	private String passWord = "password";


	/**
	 * 构造函数
	 * 
	 * @throws MqttException
	 */
	public ServerMQTT() throws MqttException {
		// MemoryPersistence设置clientid的保存形式，默认为以内存保存
		client = new MqttClient(HOST, "11111111", new MemoryPersistence());
		connect();
	}

	/**
	 * 用来连接服务器
	 */
	private void connect() {
		try {
		MqttConnectOptions options = new MqttConnectOptions();
		options.setCleanSession(false);
		options.setUserName(userName);
		options.setPassword(passWord.toCharArray());
		// 设置超时时间
		options.setConnectionTimeout(10);
		// 设置会话心跳时间
		options.setKeepAliveInterval(20);
		client.setCallback(new MqttCallback() {
			@Override
			public void messageArrived(String topic, MqttMessage message) throws Exception {
				logger.info("server-messageArrived:"+topic);
			}
			@Override
			public void deliveryComplete(IMqttDeliveryToken token) {
				logger.info("server-deliveryComplete:"+token.isComplete());
			}
			@Override
			public void connectionLost(Throwable cause) {
				logger.info("server-connectionLost:"+cause.getMessage());
				try {
					new ServerMQTT();
				} catch (MqttException e) {
					e.printStackTrace();
				}
			}
		});
		client.connect(options);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	private void close() {
		try {
			client.close();
		} catch (MqttException e) {
			e.printStackTrace();
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			try {
				client.close();
			} catch (MqttException e) {
				e.printStackTrace();
			}
		}
	}

	/**
	 * 
	 * @param topic
	 * @param message
	 * @throws MqttPersistenceException
	 * @throws MqttException
	 */
	public void publish(MqttTopic topic, MqttMessage message) throws MqttPersistenceException, MqttException {
		MqttDeliveryToken token = topic.publish(message);
		token.waitForCompletion();
		System.out.println("getResponse:"+token.getResponse());
		System.out.println("message is published completely! " + token.isComplete());
	}

	
	public static void start(ProtoMsg protoMsg) {
		ServerMQTT server = null;
		MqttMessage message = null;
		try {
			server = new ServerMQTT();
			message = new MqttMessage();
			message.setQos(0);
			message.setRetained(true);
			message.setPayload(protoMsg.toString().getBytes());
			String topic = protoMsg.getHead().getDesID().getString(0);
			System.out.println("topic1:"+topic);
			server.topic = server.client.getTopic(topic);
			server.publish(server.topic, message);
			System.out.println("topic2:"+server.topic.getName());
			System.out.println(message.getId() + "------ratained状态" + server.topic.getName());
		} catch (MqttPersistenceException e) {
			e.printStackTrace();
		} catch (MqttException e) {
			e.printStackTrace();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	public static void startBinary(ByteBuffer buffer) {
		ServerMQTT server = null;
		MqttMessage message = null;
		try {
			server = new ServerMQTT();
			message = new MqttMessage();
			message.setQos(0);
			message.setRetained(true);
			message.setPayload(buffer.array());
			String topic = "1";
			System.out.println("topic1:"+topic);
			server.topic = server.client.getTopic(topic);
			server.publish(server.topic, message);
			System.out.println("topic2:"+server.topic.getName());
			System.out.println(message.getId() + "------ratained状态" + server.topic.getName());
		} catch (MqttPersistenceException e) {
			e.printStackTrace();
		} catch (MqttException e) {
			e.printStackTrace();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	/**
	 * 启动入口
	 * 
	 * @param args
	 * @throws Exception
	 */
	public static void main(String[] args) throws Exception {
		ServerMQTT server = new ServerMQTT();
		MqttMessage message = new MqttMessage();
		// for (int i = 1; i <= 10; i++) {
		message.setQos(2);
		message.setRetained(true);
		message.setPayload(sendFile("E:\\aaa.txt"));
		server.publish(server.topic, message);
		message.setPayload("close".getBytes());
		server.publish(server.topic, message);
		System.out.println("times:" + System.currentTimeMillis());
		System.out.println(message.getId() + "------ratained状态" + server.topic.getName());
		message.clearPayload();
		// Thread.sleep(5000);
		// }
	}

	public static byte[] sendFile(String path) throws Exception {
		FileInputStream fs = null;
		try {
			System.out.println("-------");
			fs = new FileInputStream(path);
			// 获取指定文件的长度并用它来创建一个可以存放内容的字节数组
			byte[] content = new byte[fs.available()];
			// 将图片内容读入到字节数组
			fs.read(content);
			// 使用刚才的字节数组创建ByteBuffer对象
			ByteBuffer byteBuffer = ByteBuffer.wrap(content);
			System.out.println("------关闭----------");
			byte[] bytes = byteBuffer.array();
			System.out.println("bytes:" + bytes.length);
			return bytes;
		} catch (IOException e) {
			e.printStackTrace();
			throw new Exception();
		} finally {
			// 关闭文件流对象
			try {
				fs.close();
			} catch (Exception e) {
				e.printStackTrace();
			}
		}
	}

}
