package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @author wdyqxx
 * @version 2016年12月30日 下午5:44:21
 * @explain 此类用于对用户联系人的操作
 */
public class FriendsPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -3638564619920718763L;

	private Integer uid1;// 用户的uid1
	private String note1;// 用户1对用户2的印象
	private Integer uid2;// 用户的uid1
	private String note2;// 用户2对用户1的印象
	private Integer state;// 信息列表状态1：正常；2：申请；3：删除；

	public FriendsPo() {
		super();
	}

	public FriendsPo(Integer uid1, String note1, Integer uid2, String note2, Integer state) {
		super();
		this.uid1 = uid1;
		this.note1 = note1;
		this.uid2 = uid2;
		this.note2 = note2;
		this.state = state;
	}

	@Override
	public String toString() {
		return "FriendsPo [uid1=" + uid1 + ", note1=" + note1 + ", uid2=" + uid2 + ", note2=" + note2 + ", state="
				+ state + "]";
	}

	public Integer getUid1() {
		return uid1;
	}

	public void setUid1(Integer uid1) {
		this.uid1 = uid1;
	}

	public String getNote1() {
		return note1;
	}

	public void setNote1(String note1) {
		this.note1 = note1;
	}

	public Integer getUid2() {
		return uid2;
	}

	public void setUid2(Integer uid2) {
		this.uid2 = uid2;
	}

	public String getNote2() {
		return note2;
	}

	public void setNote2(String note2) {
		this.note2 = note2;
	}

	public Integer getState() {
		return state;
	}

	public void setState(Integer state) {
		this.state = state;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((uid1 == null) ? 0 : uid1.hashCode());
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
		FriendsPo other = (FriendsPo) obj;
		if (uid1 == null) {
			if (other.uid1 != null) {
				return false;
			}
		} else if (!uid1.equals(other.uid1)) {
			return false;
		}
		return true;
	}

}
