package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @author wdyqxx
 * @version 2017年1月3日 下午5:53:40
 * @explain
 */
public class GroupMessageBackupInfoPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 3759606044103400527L;

	private long mbiId;// 群消息id
	private String messagesList;// 群消息内容

	public GroupMessageBackupInfoPO() {
		super();
	}

	public GroupMessageBackupInfoPO(long mbiId, String messagesList) {
		super();
		this.mbiId = mbiId;
		this.messagesList = messagesList;
	}

	@Override
	public String toString() {
		return "GroupMessageBackupInfoPOJO [mbiId=" + mbiId + ", messagesList=" + messagesList + "]";
	}

	public long getMbiId() {
		return mbiId;
	}

	public void setMbiId(long mbiId) {
		this.mbiId = mbiId;
	}

	public String getMessagesList() {
		return messagesList;
	}

	public void setMessagesList(String messagesList) {
		this.messagesList = messagesList;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (mbiId ^ (mbiId >>> 32));
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
		GroupMessageBackupInfoPO other = (GroupMessageBackupInfoPO) obj;
		if (mbiId != other.mbiId) {
			return false;
		}
		return true;
	}
}
