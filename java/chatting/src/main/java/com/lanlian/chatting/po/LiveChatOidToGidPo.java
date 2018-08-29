/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @Title LiveChatOidToGidPo.java
 * @Package com.lanlian.chatting.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午3:23:36
 * @explain 实时上报群属性与微信用户id关系
 */

public class LiveChatOidToGidPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 9103826480465713900L;

	private Integer id;// 记录id
	private Integer gid;// 群id
	private Integer uid;// 微信用户id
	private String openid;// 微信id
	private Timestamp createTableTime;// 表记录创建时间
	private Timestamp modifyTableTime;// 表记录修改时间
	private Integer stateTable;// 表记录状态

	public LiveChatOidToGidPo() {
		super();
	}

	public LiveChatOidToGidPo(Integer id, Integer gid, Integer uid, String openid, Timestamp createTableTime,
			Timestamp modifyTableTime, Integer stateTable) {
		super();
		this.id = id;
		this.gid = gid;
		this.uid = uid;
		this.openid = openid;
		this.createTableTime = createTableTime;
		this.modifyTableTime = modifyTableTime;
		this.stateTable = stateTable;
	}

	@Override
	public String toString() {
		return "LiveChatOidToGidPo [id=" + id + ", gid=" + gid + ", uid=" + uid + ", openid=" + openid
				+ ", createTableTime=" + createTableTime + ", modifyTableTime=" + modifyTableTime + ", stateTable="
				+ stateTable + "]";
	}

	public Integer getId() {
		return id;
	}

	public void setId(Integer id) {
		this.id = id;
	}

	public Integer getGid() {
		return gid;
	}

	public void setGid(Integer gid) {
		this.gid = gid;
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public String getOpenid() {
		return openid;
	}

	public void setOpenid(String openid) {
		this.openid = openid;
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

	public Integer getStateTable() {
		return stateTable;
	}

	public void setStateTable(Integer stateTable) {
		this.stateTable = stateTable;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((id == null) ? 0 : id.hashCode());
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
		LiveChatOidToGidPo other = (LiveChatOidToGidPo) obj;
		if (id == null) {
			if (other.id != null) {
				return false;
			}
		} else if (!id.equals(other.id)) {
			return false;
		}
		return true;
	}

}
