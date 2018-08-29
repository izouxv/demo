package com.lanlian.chatting.vo;

import java.util.Map;

/**
 * @author  wdyqxx
 * @version 2017年1月3日 下午1:47:33
 * @explain 此类用于用户上传联系人功能提供对象；
 */
public class UserFriends {
	
	private int uid;//自己的uid
	private Map<String,String> infoMap;//上传的好友信息
	
	public UserFriends() {
		super();
	}

	public UserFriends(int uid, Map<String, String> infoMap) {
		super();
		this.uid = uid;
		this.infoMap = infoMap;
	}

	@Override
	public String toString() {
		return "UserFriends [uid=" + uid + ", infoMap=" + infoMap + "]";
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public Map<String, String> getInfoMap() {
		return infoMap;
	}

	public void setInfoMap(Map<String, String> infoMap) {
		this.infoMap = infoMap;
	}

}
