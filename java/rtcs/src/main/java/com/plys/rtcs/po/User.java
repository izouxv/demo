/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.po;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

import org.eclipse.paho.client.mqttv3.MqttClient;
import org.springframework.web.socket.WebSocketSession;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月29日 上午10:17:45
 * @$
 * @Administrator
 * @explain 
 */

public class User implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1639351564652050266L;
	
	private Integer uid;
	
	private String token;
	
	private String sessionId;
	
	private WebSocketSession session;
	
	private Map<String, Long> topics;
	
	private MqttClient mqttClient;

	public User() {
		super();
	}

	public User(Integer uid, String token, String sessionId, WebSocketSession session, Map<String, Long> topics,
			MqttClient mqttClient) {
		super();
		this.uid = uid;
		this.token = token;
		this.sessionId = sessionId;
		this.session = session;
		this.topics = topics;
		this.mqttClient = mqttClient;
	}

	@Override
	public String toString() {
		return "User [uid=" + uid + ", token=" + token + ", sessionId=" + sessionId + ", session=" + session
				+ ", topics=" + topics + ", mqttClient=" + mqttClient + "]";
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public String getToken() {
		return token;
	}

	public void setToken(String token) {
		this.token = token;
	}

	public String getSessionId() {
		return sessionId;
	}

	public void setSessionId(String sessionId) {
		this.sessionId = sessionId;
	}

	public WebSocketSession getSession() {
		return session;
	}

	public void setSession(WebSocketSession session) {
		this.session = session;
	}

	public Map<String, Long> getTopics() {
		return topics;
	}

	public void setTopics(Map<String, Long> topics) {
		this.topics = topics;
	}

	public MqttClient getMqttClient() {
		return mqttClient;
	}

	public void setMqttClient(MqttClient mqttClient) {
		this.mqttClient = mqttClient;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((uid == null) ? 0 : uid.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		User other = (User) obj;
		if (uid == null) {
			if (other.uid != null)
				return false;
		} else if (!uid.equals(other.uid))
			return false;
		return true;
	}

}

