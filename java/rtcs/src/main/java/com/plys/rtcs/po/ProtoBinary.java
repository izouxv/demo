/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.po;

import java.io.Serializable;
import java.util.Arrays;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月2日 上午9:37:28
 * @$
 * @Administrator
 * @explain 
 */

public class ProtoBinary implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -4539329511191949950L;
	
	private Integer uid;
	private String topic;
	private Boolean retained;
	private Byte qos;
	private byte[] payload;
	
	public ProtoBinary() {
		super();
	}

	public ProtoBinary(Integer uid, String topic, Boolean retained, Byte qos, byte[] payload) {
		super();
		this.uid = uid;
		this.topic = topic;
		this.retained = retained;
		this.qos = qos;
		this.payload = payload;
	}

	@Override
	public String toString() {
		return "ProtoBinary [uid=" + uid + ", topic=" + topic + ", retained=" + retained + ", qos=" + qos + ", payload="
				+ Arrays.toString(payload) + "]";
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public String getTopic() {
		return topic;
	}

	public void setTopic(String topic) {
		this.topic = topic;
	}

	public Boolean getRetained() {
		return retained;
	}

	public void setRetained(Boolean retained) {
		this.retained = retained;
	}

	public Byte getQos() {
		return qos;
	}

	public void setQos(Byte qos) {
		this.qos = qos;
	}

	public byte[] getPayload() {
		return payload;
	}

	public void setPayload(byte[] payload) {
		this.payload = payload;
	}

}

