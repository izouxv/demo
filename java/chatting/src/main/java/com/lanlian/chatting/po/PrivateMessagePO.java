package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午4:40:10
 * @explain 此类用于用户之间的私信留言
 */
public class PrivateMessagePO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -1921673953328695703L;

	private long id;// 表id
	private Integer uid;// 发送者的uid
	private Integer ouid;// 接受者的uid
	private String type;// 信息类型：1：文本信息；2：语音；3：图片；
	private String messageInfoId;// 信息id，对应信息内容表
	private Timestamp createTime;// 消息发送时间
	private Timestamp modifyTime;// 修改时间
	private Integer status;// 信息状态0：未读；1：已读；3：已删除；

	public PrivateMessagePO() {
		super();
	}

	public PrivateMessagePO(long id, Integer uid, Integer ouid, String type, String messageInfoId, Timestamp createTime,
			Timestamp modifyTime, int status) {
		super();
		this.id = id;
		this.uid = uid;
		this.ouid = ouid;
		this.type = type;
		this.messageInfoId = messageInfoId;
		this.createTime = createTime;
		this.modifyTime = modifyTime;
		this.status = status;
	}

	@Override
	public String toString() {
		return "PrivateMessagePO [id=" + id + ", uid=" + uid + ", ouid=" + ouid + ", type=" + type + ", messageInfoId="
				+ messageInfoId + ", createTime=" + createTime + ", modifyTime=" + modifyTime + ", status=" + status
				+ "]";
	}

	public long getId() {
		return id;
	}

	public void setId(long id) {
		this.id = id;
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public Integer getOuid() {
		return ouid;
	}

	public void setOuid(Integer ouid) {
		this.ouid = ouid;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getMessageInfoId() {
		return messageInfoId;
	}

	public void setMessageInfoId(String messageInfoId) {
		this.messageInfoId = messageInfoId;
	}

	public Timestamp getCreateTime() {
		return createTime;
	}

	public void setCreateTime(Timestamp createTime) {
		this.createTime = createTime;
	}

	public Timestamp getModifyTime() {
		return modifyTime;
	}

	public void setModifyTime(Timestamp modifyTime) {
		this.modifyTime = modifyTime;
	}

	public Integer getStatus() {
		return status;
	}

	public void setStatus(Integer status) {
		this.status = status;
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
		PrivateMessagePO other = (PrivateMessagePO) obj;
		if (id != other.id) {
			return false;
		}
		return true;
	}

}
