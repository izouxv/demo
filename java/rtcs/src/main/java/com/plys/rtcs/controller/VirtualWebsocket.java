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
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import javax.annotation.Resource;

import org.apache.log4j.Logger;
import org.springframework.http.HttpHeaders;
import org.springframework.web.socket.AbstractWebSocketMessage;
import org.springframework.web.socket.BinaryMessage;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;

import com.plys.rtcs.mqtt.ClientMQTT;
import com.plys.rtcs.mqtt.ServerMQTT;
import com.plys.rtcs.po.AbsException;
import com.plys.rtcs.po.Proto;
import com.plys.rtcs.po.ProtoMsg;
import com.plys.rtcs.po.User;
import com.plys.rtcs.service.WebMessageService;
import com.plys.rtcs.util.ProtoParse;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月26日 下午6:14:14
 * @$
 * @Administrator
 * @explain
 */

public class VirtualWebsocket implements WebSocketHandler {

	private static Logger log = Logger.getLogger(VirtualWebsocket.class);

	private static Map<String, User> USERS = new ConcurrentHashMap<>();
	
	@Resource
	WebMessageService webMessageService;

	/**
	 * 与系统建立连接
	 */
	@Override
	public void afterConnectionEstablished(WebSocketSession session) throws Exception {
		System.out.println("VirtualWebsocket:"+session.getAttributes());
		
		String sid = session.getAttributes().get("HTTP.SESSION.ID").toString();
		log.info("session:" + sid);
		log.info(session.getUri());
		HttpHeaders httpHandler = session.getHandshakeHeaders();
		log.info(httpHandler);
		// SESSIONS.add(session);
		// map.put(sid, session);
		User user = new User();
		user.setToken(sid);
		user.setSessionId(sid);
		user.setSession(session);
		USERS.put(sid, user);
		// log.info("map-sid:"+map.get(sid));
		// Integer id = SESSIONS.indexOf(session);
		// redis79.setSession(id, session.getId());
		ProtoParse.syst(USERS);
		log.info("成功连接！！！");
	}

	/**
	 * 消息处理
	 * 
	 * @param session
	 * @param message
	 * @throws Exception
	 */
	@Override
	public void handleMessage(WebSocketSession session, WebSocketMessage<?> message) {
		// org.springframework.web.socket.TextMessage
		// org.springframework.web.socket.BinaryMessage
		try {
			System.out.println("???:"+(message instanceof TextMessage));
			System.out.println("???:"+(message instanceof BinaryMessage));
			ProtoMsg protoMsg = null;
			ByteBuffer byteBuffer = null;
			Map<String, Object> sessions = session.getAttributes();
			String sid = sessions.get("HTTP.SESSION.ID").toString();
			if (message instanceof TextMessage) {
				protoMsg = new ProtoMsg();
				protoMsg = (ProtoMsg)Proto.jsonToBean(message.getPayload().toString(), protoMsg);
				protoMsg.getHead().setClient(sid);
				System.out.println("protoMsg:"+protoMsg);
				session.sendMessage(new TextMessage("收到消息".getBytes()));
				webMessageService.parseProto(protoMsg);
//				TextMessage textMessage = new TextMessage(protoMsg.toString().getBytes());
//				sendMessagesToUsers(textMessage);
				return;
			}
			if (message instanceof BinaryMessage) {
				if (message.isLast()) {
					BinaryMessage binary = new BinaryMessage((ByteBuffer) message.getPayload());
					byteBuffer = binary.getPayload();
//					System.out.println("BinaryMessage******"+ProtoParse.byToStr(byteBuffer));
					sendBinaryToUsers(byteBuffer,sid);
					ServerMQTT.startBinary(byteBuffer);
					ClientMQTT.subTopic(sid, "1");
					byteBuffer.clear();
					return;
				}
			}
			this.afterConnectionClosed(session, null);
		} catch (AbsException e) {
			System.out.println("接到的消息无法解析\n"+e.getMessage());
			try {
				session.sendMessage(new TextMessage("无法解析数据.".getBytes()));
			} catch (IOException e1) {
				System.out.println("错误");
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	/**
	 * 错误信息处理并断开连接
	 */
	@Override
	public void handleTransportError(WebSocketSession session, Throwable t) {
		try {
			log.info("handleTransportError：" + session);
			if (session.isOpen()) {
				session.close();
			}
			USERS.remove(session.getAttributes().get("HTTP.SESSION.ID").toString());
			if (t != null) {
				throw t;
			}
		} catch (Throwable e) {
			// TODO: handle exception
			e.printStackTrace();
		}
	}

	/**
	 * 用户断开连接操作
	 */
	@Override
	public void afterConnectionClosed(WebSocketSession session, CloseStatus closeStatus) {
		try {
			log.info("afterConnectionClosed：" + session + ",closeStatus:" + closeStatus);
			if (session.isOpen()) {
				session.close();
			}
			USERS.remove(session.getAttributes().get("HTTP.SESSION.ID").toString());
			log.info("退出成功！");
		} catch (Exception e) {
			// TODO: handle exception
			e.printStackTrace();
		}
	}

	/**
	 * 
	 */
	@Override
	public boolean supportsPartialMessages() {
		return false;
	}
	
	/**
	 * 供其他类使用map
	 * @return
	 */
	public static Map<String, User> getUSERS() {
		return USERS;
	}
	
	public static User getUSER(String uid) {
		if (USERS.containsKey(uid)) {
			return USERS.get(uid);
		}
		return null;
	}
	
	
/**********************************************************************************/



	/**
	 * 给所有的用户发送消息
	 */
	public void sendMessagesToUsers(TextMessage message) {
		try {
			for (String sid : USERS.keySet()) {
				WebSocketSession user = USERS.get(sid).getSession();
				// isOpen()在线就发送
				if (user.isOpen()) {
					user.sendMessage(message);
				}
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	/**
	 * 给用户发送消息
	 * 
	 */
	public static void sendMsgToUsers(String sid, BinaryMessage message) {
		try {
			WebSocketSession user = USERS.get(sid).getSession();
			log.info("session:" + user);
			if (user == null) {
				log.info("sid为空");
				return;
			}
			// isOpen()在线就发送
			if (user.isOpen()) {
				user.sendMessage(message);
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	public void sendBinaryToUsers(ByteBuffer byteBuffer, String sessionid) {
//		FileInputStream fs = null;
		try {
			BinaryMessage binaryMessage = new BinaryMessage(byteBuffer);
			for (String sid : USERS.keySet()) {
				if (sid.equals(sessionid)) {
					break;
				}
				WebSocketSession user = USERS.get(sid).getSession();
				// isOpen()在线就发送
				if (user.isOpen()) {
					user.sendMessage(binaryMessage);
				}
			}
		} catch (IOException e) {
			e.printStackTrace();
		} finally {
			// 关闭文件流对象
//			try {
//				fs.close();
//			} catch (Exception e) {
//				e.printStackTrace();
//			}
		}
	}

	/**
	 * 发送消息给指定的用户
	 */
	public void sendMessageToUser(String id, TextMessage message) {
		try {
			WebSocketSession user = USERS.get(id.toString()).getSession();
			// isOpen()在线就发送
			if (user.isOpen()) {
				user.sendMessage(message);
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
	
//	log.info("-------");
//	fs = new FileInputStream("E:\\aaa.jpg");
//	// 获取指定文件的长度并用它来创建一个可以存放内容的字节数组
//	byte[] content = new byte[fs.available()];
//	// 将图片内容读入到字节数组
//	fs.read(content);
//	// 使用刚才的字节数组创建ByteBuffer对象
//	ByteBuffer byteBuffer = ByteBuffer.wrap(content);
//	// 发送byteBuffer对象到客户端
//	BinaryMessage binaryMessage = new BinaryMessage(byteBuffer);
//	user.sendMessage(binaryMessage);
	
//	// 关闭文件流对象
//	fs.close();
//	log.info("------关闭----------");

}
