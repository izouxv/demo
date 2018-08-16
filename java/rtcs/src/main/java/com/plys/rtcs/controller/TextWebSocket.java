/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Map;
import java.util.Random;

import org.apache.log4j.Logger;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;

import com.alibaba.fastjson.JSONObject;
import com.plys.rtcs.mqtt.ClientMQTT;
import com.plys.rtcs.mqtt.ServerMQTT;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月7日 下午3:09:13
 * @$
 * @Administrator
 * @explain
 */

public class TextWebSocket extends VirtualWebsocket {

	private Logger log = Logger.getLogger(TextWebSocket.class);

	/**
	 * 消息处理
	 */
	@Override
	public void handleMessage(WebSocketSession session, WebSocketMessage<?> message) {
		// 将消息进行转化，因为是消息是json数据，可能里面包含了发送给某个人的信息，所以需要用json相关的工具类处理之后再封装成TextMessage，
		// 我这儿并没有做处理，消息的封装格式一般有{from:xxxx,to:xxxxx,msg:xxxxx}，来自哪里，发送给谁，什么消息等等
		Map<String, Object> sessions = session.getAttributes();
		String sid = sessions.get("HTTP.SESSION.ID").toString();
		log.info("handleMessage:" + message.getPayload() + ",session:" + sessions);

		// session.sendMessage(message);
		SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd hh:mm:ss");
		String time = format.format(new Date());
		JSONObject jsonObject = new JSONObject();
		jsonObject.put("from", session.getId());
		jsonObject.put("to", "1");
		jsonObject.put("type", 1);
		jsonObject.put("time", time);
		jsonObject.put("info", message.getPayload().toString());
		TextMessage msg = new TextMessage(jsonObject.toString());
		// 给所有用户群发消息
		sendMessagesToUsers(msg);
		// sendFileToUsers(msg);
		// 给指定用户群发消息
		// sendMessageToUser(userId,msg);
		Random random = new Random();
		int i = random.nextInt(3);
		log.info("i:" + i);
		try {
			ClientMQTT.subTopic(sid, "test");
		} catch (Exception e) {
			e.printStackTrace();
		}
		// ServerMQTT.start(jsonObject.toJSONString());
	}

}
