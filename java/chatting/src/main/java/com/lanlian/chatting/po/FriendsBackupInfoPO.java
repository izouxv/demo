package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @author wdyqxx
 * @version 2017年1月3日 下午5:57:08
 * @explain fbiId:好友列表的id； friendsInfo:备份的好友信息；
 */
public class FriendsBackupInfoPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -7636732032591496567L;

	private long fbiId;// 好友列表id
	private String friendsInfo;// 好友信息

	public FriendsBackupInfoPO() {
		super();
	}

	public FriendsBackupInfoPO(long fbiId, String friendsInfo) {
		super();
		this.fbiId = fbiId;
		this.friendsInfo = friendsInfo;
	}

	@Override
	public String toString() {
		return "FriendsBackupInfoPO [fbiId=" + fbiId + ", friendsInfo=" + friendsInfo + "]";
	}

	public long getFbiId() {
		return fbiId;
	}

	public void setFbiId(long fbiId) {
		this.fbiId = fbiId;
	}

	public String getFriendsInfo() {
		return friendsInfo;
	}

	public void setFriendsInfo(String friendsInfo) {
		this.friendsInfo = friendsInfo;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (fbiId ^ (fbiId >>> 32));
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
		FriendsBackupInfoPO other = (FriendsBackupInfoPO) obj;
		if (fbiId != other.fbiId) {
			return false;
		}
		return true;
	}
}
