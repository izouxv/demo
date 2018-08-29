package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午3:26:04
 * @explain 此类用于永久群数据的操作
 */
public class GroupUpgradePO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1984997137689458792L;

	private int pgid;// 群id
	private int uid;// 群主uid
	private String pgName;// 群名称
	private int avatar;// 群头像
	private String describe;// 群描述
	private String announcement;// 群公告
	private Timestamp creatTime;// 群创建时间
	private Timestamp modifyTime;// 最近修改时间
	private double longitude;// 创建时经度
	private double latitude;// 纬度
	private String groupStatus;// 群状态0：正常；3：已删除；

	public GroupUpgradePO() {
		super();
	}

	public GroupUpgradePO(int pgid, int uid, String pgName, int avatar, String describe, String announcement,
			Timestamp creatTime, Timestamp modifyTime, double longitude, double latitude, String groupStatus) {
		super();
		this.pgid = pgid;
		this.uid = uid;
		this.pgName = pgName;
		this.avatar = avatar;
		this.describe = describe;
		this.announcement = announcement;
		this.creatTime = creatTime;
		this.modifyTime = modifyTime;
		this.longitude = longitude;
		this.latitude = latitude;
		this.groupStatus = groupStatus;
	}

	@Override
	public String toString() {
		return "GroupUpgradePO [pgid=" + pgid + ", uid=" + uid + ", pgName=" + pgName + ", avatar=" + avatar
				+ ", describe=" + describe + ", announcement=" + announcement + ", creatTime=" + creatTime
				+ ", modifyTime=" + modifyTime + ", longitude=" + longitude + ", latitude=" + latitude
				+ ", groupStatus=" + groupStatus + "]";
	}

	public int getPgid() {
		return pgid;
	}

	public void setPgid(int pgid) {
		this.pgid = pgid;
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public String getPgName() {
		return pgName;
	}

	public void setPgName(String pgName) {
		this.pgName = pgName;
	}

	public int getAvatar() {
		return avatar;
	}

	public void setAvatar(int avatar) {
		this.avatar = avatar;
	}

	public String getDescribe() {
		return describe;
	}

	public void setDescribe(String describe) {
		this.describe = describe;
	}

	public String getAnnouncement() {
		return announcement;
	}

	public void setAnnouncement(String announcement) {
		this.announcement = announcement;
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

	public String getGroupStatus() {
		return groupStatus;
	}

	public void setGroupStatus(String groupStatus) {
		this.groupStatus = groupStatus;
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
		GroupUpgradePO other = (GroupUpgradePO) obj;
		if (uid != other.uid) {
			return false;
		}
		return true;
	}
}
