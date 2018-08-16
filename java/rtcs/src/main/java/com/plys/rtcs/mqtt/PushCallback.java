/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.plys.rtcs.mqtt;

import org.eclipse.paho.client.mqttv3.IMqttDeliveryToken;
import org.eclipse.paho.client.mqttv3.MqttCallback;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年8月2日 下午6:24:02
 * @explain 发布消息的回调类
 * 
 *          必须实现MqttCallback的接口并实现对应的相关接口方法CallBack 类将实现 MqttCallBack。
 *          每个客户机标识都需要一个回调实例。在此示例中，构造函数传递客户机标识以另存为实例数据。
 *          在回调中，将它用来标识已经启动了该回调的哪个实例。 必须在回调类中实现三个方法：
 * 
 *          public void messageArrived(MqttTopic topic, MqttMessage
 *          message)接收已经预订的发布。
 * 
 *          public void connectionLost(Throwable cause)在断开连接时调用。
 * 
 *          public void deliveryComplete(MqttDeliveryToken token)) 接收到已经发布的 QoS
 *          1 或 QoS 2 消息的传递令牌时调用。 由 MqttClient.connect 激活此回调。
 * 
 */
public class PushCallback implements MqttCallback {

	@Override
	public void connectionLost(Throwable cause) {
		System.out.println("连接断开，重连");
	}

	@Override
	public void deliveryComplete(IMqttDeliveryToken token) {
		System.out.println("交付完成----" + token.getMessageId());
	}

	@Override
	public void messageArrived(String topic, MqttMessage message) throws Exception {
		System.out.println("times:" + System.currentTimeMillis());
		// subscribe得到的消息
		System.out.println("主题 : " + topic);
		System.out.println("接收消息: " + message.getId() + "," + message.isDuplicate() + "," + message.isRetained());
		System.out.println("Qos : " + message.getQos());
		System.out.println("size:" + message.getPayload().length + ",内容 : " + new String(message.getPayload(), "utf-8"));

	}

}