package com.lanlian.chatting.vo;

import java.io.Serializable;

/**
 * @author  wdyqxx
 * @version 2017年1月3日 下午3:43:50
 * @explain 用于接收联系人发送给用户的私信功能；
 */
public class UserMessageNum implements Serializable {
	
	private static final long serialVersionUID = -6689369162693664615L;
	
	private int fpid;//自己的pid
	private int opid;//对方的pid
	private String newest;//最新私信数量
	
	public UserMessageNum() {
		super();
	}
	
	public UserMessageNum(int fpid, int opid, String newest) {
		super();
		this.fpid = fpid;
		this.opid = opid;
		this.newest = newest;
	}
	
	@Override
	public String toString() {
		return "UserMessageNum [fpid=" + fpid + ", opid=" + opid + ", newest=" + newest + "]";
	}

	public int getOpid() {
		return opid;
	}
	public void setOpid(int opid) {
		this.opid = opid;
	}
	public String getNewest() {
		return newest;
	}
	public void setNewest(String newest) {
		this.newest = newest;
	}
	
	public int getFpid() {
		return fpid;
	}

	public void setFpid(int fpid) {
		this.fpid = fpid;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (opid ^ (opid >>> 32));
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj) {
			return true;}
		if (obj == null) {
			return false;}
		if (getClass() != obj.getClass()) {
			return false;}
		UserMessageNum other = (UserMessageNum) obj;
		if (opid != other.opid) {
			return false;}
		return true;
	}
	
	
	
}
