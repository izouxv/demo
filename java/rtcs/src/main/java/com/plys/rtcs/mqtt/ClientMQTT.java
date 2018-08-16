/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.plys.rtcs.mqtt;

import java.util.concurrent.ScheduledExecutorService;

import org.apache.log4j.Logger;
import org.eclipse.paho.client.mqttv3.IMqttDeliveryToken;
import org.eclipse.paho.client.mqttv3.MqttCallback;
import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttConnectOptions;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;
import org.eclipse.paho.client.mqttv3.MqttTopic;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;
import org.springframework.web.socket.BinaryMessage;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;

import com.plys.rtcs.controller.BinaryWebSocket;
import com.plys.rtcs.controller.VirtualWebsocket;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月2日 下午6:23:22
 * @explain
 */

public class ClientMQTT {

	private static Logger log = Logger.getLogger(ClientMQTT.class);

	public static final String HOST = "tcp://192.168.1.6:61613";

	private static String userName = "admin";
	private static String passWord = "password";

	private static MqttClient start(final String clientid) {
		try {
			// host为主机名，clientid即连接MQTT的客户端ID，一般以唯一标识符表示，MemoryPersistence设置clientid的保存形式，默认为以内存保存
			MqttClient mqttClient = new MqttClient(HOST, clientid, new MemoryPersistence());
			MqttConnectOptions options = mqttConnectConfig();
			// 设置回调
			mqttClient.setCallback(new MqttCallback() {
				@Override
				public void messageArrived(String topic, MqttMessage message) throws Exception {
					log.info("topic:" + topic + ",id:" + clientid);
					log.info("内容 : " + message.getPayload().length);
//					log.info("内容 : " + new String(message.getPayload(), "utf-8"));
					BinaryMessage textMessage = new BinaryMessage(message.getPayload());
					VirtualWebsocket.sendMsgToUsers(clientid, textMessage);
				}
				@Override
				public void deliveryComplete(IMqttDeliveryToken token) {
					log.info(token);
				}
				@Override
				public void connectionLost(Throwable cause) {
					log.info(cause);
				}
			});
			mqttClient.connect(options);
			VirtualWebsocket.getUSERS().get(clientid).setMqttClient(mqttClient);
			return mqttClient;
		} catch (Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	/**
	 * 设置MQTT的连接参数
	 * @return
	 */
	private static MqttConnectOptions mqttConnectConfig() {
		// MQTT的连接设置
		MqttConnectOptions options = new MqttConnectOptions();
		// 设置是否清空session,这里如果设置为false表示服务器会保留客户端的连接记录，这里设置为true表示每次连接到服务器都以新的身份连接
		options.setCleanSession(true);
		// 设置连接的用户名
		options.setUserName(userName);
		// 设置连接的密码
		options.setPassword(passWord.toCharArray());
		// 设置超时时间 单位为秒
		options.setConnectionTimeout(30);
		// 设置会话心跳时间 单位为秒 服务器会每隔1.5*20秒的时间向客户端发送个消息判断客户端是否在线，但这个方法并没有重连的机制
		options.setKeepAliveInterval(60);
		return options;
	}

	public static void subTopic(String uid, String topic) {
		try {
			MqttClient mqttClient = VirtualWebsocket.getUSER(uid).getMqttClient();
			boolean falg = (mqttClient == null || !mqttClient.isConnected());
			log.info("falg:" + falg);
			// 消息质量
			int[] qos = { 1,1 };
			// 订阅主题
			String[] topics = { topic,topic+1 };
			if (falg) {
				mqttClient = start(uid);
				MqttTopic mqttTopic = mqttClient.getTopic(topic);
				log.info("topic:" + mqttTopic);
				// setWill方法，如果项目中需要知道客户端是否掉线可以调用该方法。设置最终端口的通知消息
				// options.setWill(topic, "close".getBytes(), 2, true);
				mqttClient.subscribe(topics, qos);
				return;
			}
			String sids = mqttClient.getClientId();
			log.info("sid:" + sids);
			MqttTopic mqttTopic = mqttClient.getTopic(topic);
			log.info("topic:" + mqttTopic);
			mqttClient.subscribe(topics, qos);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	public static void main(String[] args) throws MqttException, InterruptedException {
		MqttClient mqttClient = new MqttClient("", "");
		MqttClient mqttclient = start("client125");
		log.info("mqttclient:" + mqttclient);
		Thread.sleep(5000);
		mqttclient.disconnect();
	}
}
