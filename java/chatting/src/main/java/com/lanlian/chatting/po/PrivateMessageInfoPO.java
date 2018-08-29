package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午5:31:06
 * @explain 此类用于返回私信内容
 */
public class PrivateMessageInfoPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1965088331718089290L;

	private long messageInfoId;// 信息内容的id
	private String messageInfo;// 信息内容

	public PrivateMessageInfoPO() {
		super();
	}

	public PrivateMessageInfoPO(long messageInfoId, String messageInfo) {
		super();
		this.messageInfoId = messageInfoId;
		this.messageInfo = messageInfo;
	}

	public long getMessageInfoId() {
		return messageInfoId;
	}

	public void setMessageInfoId(long messageInfoId) {
		this.messageInfoId = messageInfoId;
	}

	public String getMessageInfo() {
		return messageInfo;
	}

	public void setMessageInfo(String messageInfo) {
		this.messageInfo = messageInfo;
	}

	@Override
	public String toString() {
		return "PrivateMessageInfoPOJO [messageInfoId=" + messageInfoId + ", messageInfo=" + messageInfo + "]";
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (messageInfoId ^ (messageInfoId >>> 32));
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj) {
			return true;
		}
		if (obj == null) {
			return false;
		}
		if (getClass() != obj.getClass()) {
			return false;
		}
		PrivateMessageInfoPO other = (PrivateMessageInfoPO) obj;
		if (messageInfoId != other.messageInfoId) {
			return false;
		}
		return true;
	}
}
