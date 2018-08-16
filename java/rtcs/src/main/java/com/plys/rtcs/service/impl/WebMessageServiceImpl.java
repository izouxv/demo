/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.service.impl;

import java.io.IOException;

import org.eclipse.paho.client.mqttv3.MqttException;
import org.springframework.stereotype.Service;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;

import com.plys.rtcs.controller.VirtualWebsocket;
import com.plys.rtcs.mqtt.ClientMQTT;
import com.plys.rtcs.mqtt.ServerMQTT;
import com.plys.rtcs.po.Proto;
import com.plys.rtcs.po.ProtoHead;
import com.plys.rtcs.po.ProtoMsg;
import com.plys.rtcs.po.User;
import com.plys.rtcs.service.WebMessageService;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月1日 下午6:21:24
 * @$
 * @Administrator
 * @explain 
 */

@Service("webMessageServiceImpl")
public class WebMessageServiceImpl implements WebMessageService {

	@Override
	public <T> void parseProto(Proto<T> proto) {
		try {
			ProtoMsg protoMsg = (ProtoMsg) proto;
			ProtoHead protoHead = protoMsg.getHead();
			if ("0.0.1".equals(protoHead.getPv())) {
				System.out.println("协议版本错误");
			}
			if ("123456".equals(protoHead.getClient())) {
				System.out.println("客户端错误");
			}
			//子类型判断
//			subTypeParse(protoHead);
			//订阅主题
			pubTopic(protoMsg);
			String uid = protoHead.getClient();
			System.out.println("uid:"+uid);
			ClientMQTT.subTopic(uid, protoHead.getDesID().getString(0).toString());
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	private void subTypeParse(ProtoHead protoHead) {
		String type = protoHead.getType().toString() + protoHead.getSubType();
		switch (type) {
		case "":
			
			break;

		default:
			break;
		}
	}
	
	/**
	 * 单聊
	 * @param protoMsg
	 */
	public void sendMsgToUser(ProtoMsg protoMsg) {
		try {
			User user = VirtualWebsocket.getUSER(protoMsg.getHead().getDesID().getString(0));
			WebSocketSession webSocketSession = user.getSession();
			TextMessage textMessage = new TextMessage(protoMsg.toString().getBytes());
			webSocketSession.sendMessage(textMessage);
		} catch (IOException e) {
			e.printStackTrace();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	/**
	 * 获取成员
	 */
	public void getMember() {
		
	}
	
	/**
	 * 
	 * 发布主题
	 * @param protoMsg
	 */
	public void pubTopic(ProtoMsg protoMsg) {
		ServerMQTT.start(protoMsg);
	}
	
	/**
	 * 订阅主题
	 * @param protoMsg
	 * @param user
	 */
	public void subTopic(ProtoMsg protoMsg) {
		
	}
	
	/**
	 * 向主题发小消息
	 * @param protoMsg
	 */
	public void sendMsgToTopic(ProtoMsg protoMsg) {
		
	}

}

