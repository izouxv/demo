package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午6:49:58
 * @explain 此类用于用户与群之间的关系操作
 */
public class GroupAndUserPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -6163209952681482411L;

	private long id;// 记录id
	private int uid;// 用户uid
	private int gid;// 永久群pgid
	private Timestamp creatTime;// 创建时间
	private Timestamp modifyTime;// 最近修改时间
	private int dataState;// 用户与群的关系状态

	public GroupAndUserPO() {
		super();
	}

	public GroupAndUserPO(long id, int uid, int gid, Timestamp creatTime, Timestamp modifyTime, int dataState) {
		super();
		this.id = id;
		this.uid = uid;
		this.gid = gid;
		this.creatTime = creatTime;
		this.modifyTime = modifyTime;
		this.dataState = dataState;
	}

	@Override
	public String toString() {
		return "GroupAndUserPOJO [id=" + id + ", uid=" + uid + ", gid=" + gid + ", creatTime=" + creatTime
				+ ", modifyTime=" + modifyTime + ", dataState=" + dataState + "]";
	}

	public long getId() {
		return id;
	}

	public void setId(long id) {
		this.id = id;
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public int getGid() {
		return gid;
	}

	public void setGid(int gid) {
		this.gid = gid;
	}

	public Timestamp getCreatTime() {
		return creatTime;
	}

	public void setCreatTime(Timestamp creatTime) {
		this.creatTime = creatTime;
	}

	public Timestamp getModifyTime() {
		return modifyTime;
	}

	public void setModifyTime(Timestamp modifyTime) {
		this.modifyTime = modifyTime;
	}

	public int getDataState() {
		return dataState;
	}

	public void setDataState(int dataState) {
		this.dataState = dataState;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (id ^ (id >>> 32));
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
		GroupAndUserPO other = (GroupAndUserPO) obj;
		if (id != other.id) {
			return false;
		}
		return true;
	}
}
