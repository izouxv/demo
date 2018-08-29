package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * 
 * @author yfq
 * @version 2017年3月29日 下午1:53:45
 * @explain 此类用于群用户信息查询的相关操作
 */

public class GroupUserInfoPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 6556698032898659681L;

	private int uid;// 用户的id
	private String nickname;// 昵称
	private int gender;// 性别1：男/2：女
	private int avatar;// 头像id（以后会是头像地址）
	private String signature;// 个性签名

	@Override
	public String toString() {
		return "GroupUsersInfo [uid=" + uid + ", nickname=" + nickname + ", gender=" + gender + ", avatar=" + avatar
				+ ", signature=" + signature + "]";
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
	}

	public int getGender() {
		return gender;
	}

	public void setGender(int gender) {
		this.gender = gender;
	}

	public int getAvatar() {
		return avatar;
	}

	public void setAvatar(int avatar) {
		this.avatar = avatar;
	}

	public String getSignature() {
		return signature;
	}

	public void setSignature(String signature) {
		this.signature = signature;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + uid;
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
		GroupUserInfoPO other = (GroupUserInfoPO) obj;
		if (uid != other.uid) {
			return false;
		}
		return true;
	}
}
