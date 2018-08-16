/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package test.mqtt;

import java.io.FileInputStream;
import java.io.IOException;
import java.nio.ByteBuffer;

import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttConnectOptions;
import org.eclipse.paho.client.mqttv3.MqttDeliveryToken;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;
import org.eclipse.paho.client.mqttv3.MqttPersistenceException;
import org.eclipse.paho.client.mqttv3.MqttTopic;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月2日 下午6:10:53
 * @explain 服务器向多个客户端推送主题，即不同客户端可向服务器订阅相同主题
 */

public class ServerMQTT {

	// tcp://MQTT安装的服务器地址
	public static final String HOST = "tcp://192.168.1.6:61613";
	// 定义一个主题
	public static final String[] TOPIC = {"test","test1"};
	// 定义MQTT的ID，可在MQTT服务配置中指定
	private static final String CLIENTID = "server11";

	private MqttClient client;
	private MqttTopic topic11;
	private String userName = "admin";
	private String passWord = "password";

	private MqttMessage message;

	/**
	 * 构造函数
	 * 
	 * @throws MqttException
	 */
	public ServerMQTT() throws MqttException {
		// MemoryPersistence设置clientid的保存形式，默认为以内存保存
		client = new MqttClient(HOST, CLIENTID, new MemoryPersistence());
		connect();
	}

	/**
	 * 用来连接服务器
	 */
	private void connect() {
		MqttConnectOptions options = new MqttConnectOptions();
		options.setCleanSession(false);
		options.setUserName(userName);
		options.setPassword(passWord.toCharArray());
		// 设置超时时间
		options.setConnectionTimeout(10);
		// 设置会话心跳时间
		options.setKeepAliveInterval(20);
		try {
			client.setCallback(new PushCallback());
			client.connect(options);
			for (String str : TOPIC) {
				topic11 = client.getTopic(str);				
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	private void close() {
		try {
			client.close();
		} catch (MqttException e) {
			e.printStackTrace();
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
		System.out.println("times:"+System.currentTimeMillis());
		System.out.println(token.getResponse());
		System.out.println("message is published completely! " + token.isComplete());
	}

	/**
	 * 启动入口
	 * 
	 * @param args
	 * @throws Exception 
	 */
	public static void main(String[] args) throws Exception {
		ServerMQTT server = new ServerMQTT();
		server.message = new MqttMessage();
//		for (int i = 1; i <= 10; i++) {
			server.message.setQos(2);
			server.message.setRetained(true);
			server.message.setPayload(sendFile("E:\\aaa.txt"));
			server.publish(server.topic11, server.message);
			server.message.setPayload("close".getBytes());
			server.publish(server.topic11, server.message);
			System.out.println("times:"+System.currentTimeMillis());
			System.out.println(server.message.getId() + "------ratained状态"+server.topic11.getName());
			server.message.clearPayload();
//			Thread.sleep(5000);
//		}
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
			System.out.println("bytes:"+bytes.length);
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
