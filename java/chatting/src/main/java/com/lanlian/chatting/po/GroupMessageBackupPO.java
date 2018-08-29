package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午6:32:45
 * @explain 此类用于用户对群信息的操作
 */
public class GroupMessageBackupPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 6331291918434363329L;

	private int uid;// 用户uid
	private int gid;// 群gid
	private int messageId;// 消息内容ID
	private String fileId;// 消息文件ID
	private String fileName;// 消息文件名称
	private double longitude;// 经度
	private double latitude;// 纬度
	private String address;// 对应地址
	private Timestamp createTime;// 消息创建时间
	private Timestamp backupTime;// 消息备份时间
	private Timestamp modifyTime;// 最近修改时间
	private String status;// 该记录状态1：正常；2：禁用；3：已删除；

	public GroupMessageBackupPO() {
		super();
	}

	public GroupMessageBackupPO(int uid, int gid, int messageId, String fileId, String fileName, double longitude,
			double latitude, String address, Timestamp createTime, Timestamp backupTime, Timestamp modifyTime,
			String status) {
		super();
		this.uid = uid;
		this.gid = gid;
		this.messageId = messageId;
		this.fileId = fileId;
		this.fileName = fileName;
		this.longitude = longitude;
		this.latitude = latitude;
		this.address = address;
		this.createTime = createTime;
		this.backupTime = backupTime;
		this.modifyTime = modifyTime;
		this.status = status;
	}

	@Override
	public String toString() {
		return "GroupMessageBackupPO [uid=" + uid + ", gid=" + gid + ", messageId=" + messageId + ", fileId=" + fileId
				+ ", fileName=" + fileName + ", longitude=" + longitude + ", latitude=" + latitude + ", address="
				+ address + ", createTime=" + createTime + ", backupTime=" + backupTime + ", modifyTime=" + modifyTime
				+ ", status=" + status + "]";
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

	public int getMessageId() {
		return messageId;
	}

	public void setMessageId(int messageId) {
		this.messageId = messageId;
	}

	public String getFileId() {
		return fileId;
	}

	public void setFileId(String fileId) {
		this.fileId = fileId;
	}

	public String getFileName() {
		return fileName;
	}

	public void setFileName(String fileName) {
		this.fileName = fileName;
	}

	public double getLongitude() {
		return longitude;
	}

	public void setLongitude(double longitude) {
		this.longitude = longitude;
	}

	public double getLatitude() {
		return latitude;
	}

	public void setLatitude(double latitude) {
		this.latitude = latitude;
	}

	public String getAddress() {
		return address;
	}

	public void setAddress(String address) {
		this.address = address;
	}

	public Timestamp getCreateTime() {
		return createTime;
	}

	public void setCreateTime(Timestamp createTime) {
		this.createTime = createTime;
	}

	public Timestamp getBackupTime() {
		return backupTime;
	}

	public void setBackupTime(Timestamp backupTime) {
		this.backupTime = backupTime;
	}

	public Timestamp getModifyTime() {
		return modifyTime;
	}

	public void setModifyTime(Timestamp modifyTime) {
		this.modifyTime = modifyTime;
	}

	public String getStatus() {
		return status;
	}

	public void setStatus(String status) {
		this.status = status;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + gid;
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
		GroupMessageBackupPO other = (GroupMessageBackupPO) obj;
		if (gid != other.gid) {
			return false;
		}
		return true;
	}
}
