package com.lanlian.chatting.vo;

import java.io.Serializable;

/**
 * @author  wdyqxx
 * @version 2017年1月3日 下午2:03:15
 * @explain 此类用于用户发送私信功能；
 */
public class UserLetter implements Serializable {
	
	private static final long serialVersionUID = -36143277522335653L;
	
	private int uid;//自己的uid
	private int touid;//对方的touid
	private long sendTime;//发送时间
	private String type;//私信类型
	private String letter;//私信内容
	
	public UserLetter() {
		super();
	}
	
	public UserLetter(int uid, int touid, long sendTime, String type, String letter) {
		super();
		this.uid = uid;
		this.touid = touid;
		this.sendTime = sendTime;
		this.type = type;
		this.letter = letter;
	}
	
	@Override
	public String toString() {
		return "UserLetter [uid=" + uid + ", touid=" + touid + ", sendTime=" + sendTime + ", type=" + type + ", letter="
				+ letter + "]";
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public int getTouid() {
		return touid;
	}
	public void setTouid(int touid) {
		this.touid = touid;
	}
	public long getSendTime() {
		return sendTime;
	}
	public void setSendTime(long sendTime) {
		this.sendTime = sendTime;
	}
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}
	public String getLetter() {
		return letter;
	}
	public void setLetter(String letter) {
		this.letter = letter;
	}
	
}
