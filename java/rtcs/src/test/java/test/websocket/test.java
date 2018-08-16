/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package test.websocket;

import static org.junit.Assert.fail;

import java.io.IOException;
import java.lang.reflect.Type;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.ByteBuffer;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

import javax.websocket.CloseReason;
import javax.websocket.ContainerProvider;
import javax.websocket.DeploymentException;
import javax.websocket.SendHandler;
import javax.websocket.SendResult;
import javax.websocket.Session;
import javax.websocket.WebSocketContainer;

import org.springframework.messaging.converter.MappingJackson2MessageConverter;
import org.springframework.messaging.converter.StringMessageConverter;
import org.springframework.messaging.simp.stomp.StompCommand;
import org.springframework.messaging.simp.stomp.StompFrameHandler;
import org.springframework.messaging.simp.stomp.StompHeaders;
import org.springframework.messaging.simp.stomp.StompSession;
import org.springframework.messaging.simp.stomp.StompSessionHandler;
import org.springframework.messaging.simp.stomp.StompSessionHandlerAdapter;
import org.springframework.scheduling.concurrent.ThreadPoolTaskScheduler;
import org.springframework.util.StopWatch;
import org.springframework.util.concurrent.ListenableFuture;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.WebSocketSession;
import org.springframework.web.socket.client.WebSocketClient;
import org.springframework.web.socket.client.standard.StandardWebSocketClient;
import org.springframework.web.socket.messaging.WebSocketStompClient;
import org.springframework.web.socket.sockjs.client.SockJsClient;
import org.springframework.web.socket.sockjs.client.Transport;
import org.springframework.web.socket.sockjs.client.TransportRequest;
import org.springframework.web.socket.sockjs.client.WebSocketTransport;
import org.springframework.web.socket.sockjs.transport.TransportType;

import test.websocket.StompWebSocketLoadTestClient.ConsumerStompSessionHandler;

import javax.websocket.CloseReason.CloseCodes;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月7日 下午5:30:11
 * @$
 * @Administrator
 * @explain
 */

public class test {

//	public static void main(String[] aaa) {
//		try {
//			String uri = "ws://192.168.1.69:8080/rtcs/app/sendMessage";
//			StandardWebSocketClient standardWebSocketClient = new StandardWebSocketClient();
//			List<Transport> transports = new ArrayList<>();
//			transports.add(new WebSocketTransport(standardWebSocketClient));
//			WebSocketClient webSocketClient = new SockJsClient(transports);
//			WebSocketStompClient webSocketStompClient = new WebSocketStompClient(webSocketClient);
//			webSocketStompClient.setMessageConverter(new MappingJackson2MessageConverter());  
//	        ThreadPoolTaskScheduler taskScheduler = new ThreadPoolTaskScheduler();  
//	        taskScheduler.afterPropertiesSet();  
//	        webSocketStompClient.setTaskScheduler(taskScheduler); // for heartbeats  
//	  
//	        String url = "ws://localhost:8080/gs-guide-websocket";
//			
//		} catch (Exception e) {
//			e.printStackTrace();
//		}
//	}
	
	
	/**
	 * 测试连接websocket
	 */
	public static void main(String[] args) {
		Session session = null;
		try {
			String uri = "ws://192.168.1.69:8080/rtcs/hello";
			WebSocketContainer container = ContainerProvider.getWebSocketContainer();
			// 连接会话
			session = container.connectToServer(Client.class, new URI(uri));
			//设置格式
			//换行
			char lf = 10;
			//消息结尾标记，必须
			char nl = 0;
			//组装消息
			StringBuffer sb = new StringBuffer();
			//请求命令策略
			sb.append("SEND").append(lf);
			//请求资源
			sb.append("destination:/app/send").append(lf);
			// 消息体的长度
			sb.append("content-length:14").append(lf).append(lf); 
			// 消息体
			sb.append("{\"name\":\"123\"}").append(nl); 
			System.out.println("消息:"+sb.toString());
			// 发送文本消息
			session.getBasicRemote().sendText(sb.toString());
			// 发送二进制流
//			ByteBuffer buffer = ByteBuffer.allocate(1024);
//			buffer.put("hahahaah".getBytes("utf-8"));
//			buffer.flip();
//			session.getBasicRemote().sendBinary(buffer, true);
//			session.getBasicRemote().sendBinary(buffer);
			Thread.sleep(1000);
			session.close(new CloseReason(CloseCodes.NORMAL_CLOSURE, "session close"));
		} catch (DeploymentException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		} catch (URISyntaxException e) {
			e.printStackTrace();
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
	}
	
	
	

}
