package com.lanlian.chatting.vo;

import java.io.Serializable;
import java.sql.Timestamp;

/** 
 * @Title MessageInfoPage.java
 * @Package com.lanlian.chatting.vo
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午4:36:12
 * @explain 信息实时上报群信息属性类
 */

public class MessageInfoPage implements Serializable {
	
	/**
	 * 序列化
	 */
	private static final long serialVersionUID = -6667316565511500840L;
	
	private Integer mid;
	private Integer uid;
	private Integer avatar;
	private String nickname;
	private Timestamp time;
	private Integer type;
	private String info;
	
	public MessageInfoPage() {
		super();
	}
	
	public MessageInfoPage(Integer mid, Integer uid, Integer avatar, String nickname, Timestamp time, Integer type, String info) {
		super();
		this.mid = mid;
		this.uid = uid;
		this.avatar = avatar;
		this.nickname = nickname;
		this.time = time;
		this.type = type;
		this.info = info;
	}
	public int getMid() {
		return mid;
	}

	public void setMid(Integer mid) {
		this.mid = mid;
	}

	public int getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public int getAvatar() {
		return avatar;
	}

	public void setAvatar(Integer avatar) {
		this.avatar = avatar;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
	}

	public Timestamp getTime() {
		return time;
	}

	public void setTime(Timestamp time) {
		this.time = time;
	}

	public int getType() {
		return type;
	}

	public void setType(Integer type) {
		this.type = type;
	}

	public String getInfo() {
		return info;
	}

	public void setInfo(String info) {
		this.info = info;
	}

	@Override
	public String toString() {
		return "MessageInfoPage [mid=" + mid + ", uid=" + uid + ", avatar=" + avatar + ", nickname=" + nickname
				+ ", time=" + time + ", type=" + type + ", info=" + info + "]";
	}
	
}
