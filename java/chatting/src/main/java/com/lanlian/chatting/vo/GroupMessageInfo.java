package com.lanlian.chatting.vo;

import java.io.Serializable;

public class GroupMessageInfo implements Serializable {
	
	/**
	 * 
	 */
	private static final long serialVersionUID = 2465199653696777053L;
	
	private int messageInfoId;   //消息id
	private  String messageInfo; //消息的内容
	
	
	
	public GroupMessageInfo() {
		super();
	}
	


	public GroupMessageInfo(int messageInfoId, String messageInfo) {
		super();
		this.messageInfoId = messageInfoId;
		this.messageInfo = messageInfo;
	}



	@Override
	public String toString() {
		return "GroupMessageInfo [messageInfoId=" + messageInfoId + ", messageInfo=" + messageInfo + "]";
	}

	public int getMessageInfoId() {
		return messageInfoId;
	}

	public void setMessageInfoId(int messageInfoId) {
		this.messageInfoId = messageInfoId;
	}

	public String getMessageInfo() {
		return messageInfo;
	}

	public void setMessageInfo(String messageInfo) {
		this.messageInfo = messageInfo;
	}
	
	
}
