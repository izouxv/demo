/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.controller;

import java.io.FileInputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.text.SimpleDateFormat;
import java.util.Date;

import javax.annotation.Resource;
import javax.websocket.server.PathParam;

import org.apache.log4j.Logger;
import org.springframework.http.HttpHeaders;
import org.springframework.web.socket.BinaryMessage;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.PongMessage;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;
import org.springframework.web.socket.handler.AbstractWebSocketHandler;

import com.alibaba.fastjson.JSONObject;
import com.plys.rtcs.po.RMsgTemplate;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月7日 下午3:09:13
 * @$
 * @Administrator
 * @explain
 */

public class BinaryWebSocket extends VirtualWebsocket {

	private Logger log = Logger.getLogger(BinaryWebSocket.class);

	// public static final List<WebSocketSession> SESSIONS =
	// Collections.synchronizedList(new ArrayList<WebSocketSession>());

	/**
	 * 与系统建立连接
	 */
	@Override
	public void afterConnectionEstablished(WebSocketSession session) throws Exception {
		// log.info("session:"+session + ":" +session.getAttributes());
		// HttpHeaders httpHandler = session.getHandshakeHeaders();
		// log.info(httpHandler);
		// TextWebSocket.SESSIONS.add(session);
		// session.sendMessage(new TextMessage("connect"));
		// log.info("成功连接！！！");
		System.out.println("BinaryWebSocket-afterConnectionEstablished");
		super.afterConnectionEstablished(session);
	}

	/**
	 * 消息处理
	 */
	@Override
	public void handleMessage(WebSocketSession session, WebSocketMessage<?> message) {
		try {
			// 将消息进行转化，因为是消息是json数据，可能里面包含了发送给某个人的信息，所以需要用json相关的工具类处理之后再封装成TextMessage，
			// 我这儿并没有做处理，消息的封装格式一般有{from:xxxx,to:xxxxx,msg:xxxxx}，来自哪里，发送给谁，什么消息等等
			log.info("session:" + session.getAttributes() + ",Message:" + message.getPayloadLength() + ","
					+ message.getClass().getTypeName());
			SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd hh:mm:ss");
			String time = format.format(new Date());
			TextMessage text = new TextMessage(time);
			session.sendMessage(text);
			session.sendMessage(message);

			// JSONObject jsonObject = new JSONObject();
			// jsonObject.put("from", session.getId());
			// jsonObject.put("to", "1");
			// jsonObject.put("type", 1);
			// jsonObject.put("time", time);
			// jsonObject.put("info", message);
			// RMsgTemplate rMsgTemplate = new RMsgTemplate("leo.pay.fanout.exchange",
			// "rtcs.online.group.settings", message);

			 BinaryMessage msg = new BinaryMessage((ByteBuffer) message.getPayload());
			// 给所有用户群发消息
			// sendFileToUsers(conver(message));
			System.out.println("----------------" + message.getPayloadLength() + session.getAttributes());
			super.handleMessage(session, message);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (Exception e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}

	/****************************************************************/

	private BinaryMessage conver(WebSocketMessage<?> message) throws Exception {
		try {
			BinaryMessage binaryMessage = (BinaryMessage) message;
			return binaryMessage;
		} catch (Exception e) {
			e.printStackTrace();
			throw new Exception();
		}
	}

}
