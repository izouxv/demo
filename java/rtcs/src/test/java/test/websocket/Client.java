/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package test.websocket;

import java.nio.ByteBuffer;

import javax.websocket.ClientEndpoint;
import javax.websocket.OnError;
import javax.websocket.OnMessage;
import javax.websocket.OnOpen;
import javax.websocket.Session;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月7日 下午5:31:07
 * @$
 * @Administrator
 * @explain
 */

@ClientEndpoint
public class Client {

	@OnOpen
	public void onOpen(Session session) {
		System.out.println("Connected to endpoint: " + session.getRequestParameterMap());
	}

	@OnMessage
	public void onMessage(String message) {
		System.out.println(message);
	}

	@OnError
	public void onError(Throwable t) {
		t.printStackTrace();
	}
}
