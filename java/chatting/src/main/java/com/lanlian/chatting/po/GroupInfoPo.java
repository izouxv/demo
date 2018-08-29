package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午3:26:04
 * @explain 此类用于永久群数据的操作
 */
public class GroupInfoPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 4719055033038215007L;

	private int gid;// 群id
	private int uid;// 群主uid
	private String gname;// 群名称
	private int avatar;// 群头像
	private String describe;// 群描述
	private String announcement;// 群公告
	private Timestamp creatTime;// 群创建时间
	private Timestamp modifyTime;// 最近修改时间
	private double longitude;// 经度
	private double latitude;// 纬度
	private Timestamp dataCreateTime;// 表记录创建时间
	private Timestamp dataModifyTime;// 表记录修改时间
	private int dataState;// 表记录状态
	private Integer newgid;// 新的临时群id(此群升级时，返回新的临时群id)

	public GroupInfoPo() {
		super();
	}

	public GroupInfoPo(int gid, int uid, String gname, int avatar, String describe, String announcement,
			Timestamp creatTime, Timestamp modifyTime, double longitude, double latitude, Timestamp dataCreateTime,
			Timestamp dataModifyTime, int dataState, Integer newgid) {
		super();
		this.gid = gid;
		this.uid = uid;
		this.gname = gname;
		this.avatar = avatar;
		this.describe = describe;
		this.announcement = announcement;
		this.creatTime = creatTime;
		this.modifyTime = modifyTime;
		this.longitude = longitude;
		this.latitude = latitude;
		this.dataCreateTime = dataCreateTime;
		this.dataModifyTime = dataModifyTime;
		this.dataState = dataState;
		this.newgid = newgid;
	}

	@Override
	public String toString() {
		return "GroupInfoPo [gid=" + gid + ", uid=" + uid + ", gname=" + gname + ", avatar=" + avatar + ", describe="
				+ describe + ", announcement=" + announcement + ", creatTime=" + creatTime + ", modifyTime="
				+ modifyTime + ", longitude=" + longitude + ", latitude=" + latitude + ", dataCreateTime="
				+ dataCreateTime + ", dataModifyTime=" + dataModifyTime + ", dataState=" + dataState + ", newgid="
				+ newgid + "]";
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

	public String getGname() {
		return gname;
	}

	public void setGname(String gname) {
		this.gname = gname;
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

	public Timestamp getDataCreateTime() {
		return dataCreateTime;
	}

	public void setDataCreateTime(Timestamp dataCreateTime) {
		this.dataCreateTime = dataCreateTime;
	}

	public Timestamp getDataModifyTime() {
		return dataModifyTime;
	}

	public void setDataModifyTime(Timestamp dataModifyTime) {
		this.dataModifyTime = dataModifyTime;
	}

	public int getDataState() {
		return dataState;
	}

	public void setDataState(int dataState) {
		this.dataState = dataState;
	}

	public Integer getNewgid() {
		return newgid;
	}

	public void setNewgid(Integer newgid) {
		this.newgid = newgid;
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
		GroupInfoPo other = (GroupInfoPo) obj;
		if (gid != other.gid) {
			return false;
		}
		return true;
	}
}
