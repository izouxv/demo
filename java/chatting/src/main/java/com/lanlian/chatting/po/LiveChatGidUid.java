package com.lanlian.chatting.po;
/**
 * @Description: 
 * @author: 李大双
 * @date: 2017年6月30日 下午4:37:30
 * @version: V1.0
 */

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * 微信端用户订阅群属性类
 * 
 * @author Administrator
 *
 */
public class LiveChatGidUid implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 7714184205889977652L;

	private int id;
	private int gid;
	private int uid;
	private String gname;
	private Timestamp createTableTime;
	private Timestamp modifyTableTime;
	private int stateTable;

	public LiveChatGidUid() {
		super();
	}

	public LiveChatGidUid(int id, int gid, int uid, String gname, Timestamp createTableTime, Timestamp modifyTableTime,
			int stateTable) {
		super();
		this.id = id;
		this.gid = gid;
		this.uid = uid;
		this.gname = gname;
		this.createTableTime = createTableTime;
		this.modifyTableTime = modifyTableTime;
		this.stateTable = stateTable;
	}

	@Override
	public String toString() {
		return "LiveChatGidUid [id=" + id + ", gid=" + gid + ", uid=" + uid + ", gname=" + gname + ", createTableTime="
				+ createTableTime + ", modifyTableTime=" + modifyTableTime + ", stateTable=" + stateTable + "]";
	}

	public String getGname() {
		return gname;
	}

	public void setGname(String gname) {
		this.gname = gname;
	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public int getGid() {
		return gid;
	}

	public void setGid(int gid) {
		this.gid = gid;
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public Timestamp getCreateTableTime() {
		return createTableTime;
	}

	public void setCreateTableTime(Timestamp createTableTime) {
		this.createTableTime = createTableTime;
	}

	public Timestamp getModifyTableTime() {
		return modifyTableTime;
	}

	public void setModifyTableTime(Timestamp modifyTableTime) {
		this.modifyTableTime = modifyTableTime;
	}

	public int getStateTable() {
		return stateTable;
	}

	public void setStateTable(int stateTable) {
		this.stateTable = stateTable;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + id;
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
		LiveChatGidUid other = (LiveChatGidUid) obj;
		if (id != other.id) {
			return false;
		}
		return true;
	}
}
