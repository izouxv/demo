package com.lanlian.chatting.vo;

import java.sql.Timestamp;

import com.alibaba.fastjson.JSONArray;

/**
 * @author wangdy
 * @version 2017年3月27日 上午10:19:06
 * @explain 此类用于群信息修改相关操作。
 */
public class GroupInfo {

	private Integer gid; // 群ID
	private Integer uid;//群主id
	private String gname;// 群名称
	private Integer avatar;// 群头像
	private String announcement;// 群公告
	private Timestamp creatTime;// 群创建时间
	private String longitude;// 创建时经度
	private String latitude;// 纬度
	private JSONArray users; // 所有群成员信息Map.toString
	
	public GroupInfo() {
		super();
	}

	public GroupInfo(Integer gid, Integer uid, String gname, Integer avatar, String announcement, Timestamp creatTime,
			String longitude, String latitude, JSONArray users) {
		super();
		this.gid = gid;
		this.uid = uid;
		this.gname = gname;
		this.avatar = avatar;
		this.announcement = announcement;
		this.creatTime = creatTime;
		this.longitude = longitude;
		this.latitude = latitude;
		this.users = users;
	}

	@Override
	public String toString() {
		return "GroupInfo [gid=" + gid + ", uid=" + uid + ", gname=" + gname + ", avatar=" + avatar + ", announcement="
				+ announcement + ", creatTime=" + creatTime + ", longitude=" + longitude + ", latitude=" + latitude
				+ ", users=" + users + "]";
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public Integer getGid() {
		return gid;
	}

	public void setGid(Integer gid) {
		this.gid = gid;
	}

	public String getGname() {
		return gname;
	}

	public void setGname(String gname) {
		this.gname = gname;
	}

	public Integer getAvatar() {
		return avatar;
	}

	public void setAvatar(Integer avatar) {
		this.avatar = avatar;
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

	public String getLongitude() {
		return longitude;
	}

	public void setLongitude(String longitude) {
		this.longitude = longitude;
	}

	public String getLatitude() {
		return latitude;
	}

	public void setLatitude(String latitude) {
		this.latitude = latitude;
	}

	public JSONArray getUsers() {
		return users;
	}

	public void setUsers(JSONArray users) {
		this.users = users;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((gid == null) ? 0 : gid.hashCode());
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
		GroupInfo other = (GroupInfo) obj;
		if (gid == null) {
			if (other.gid != null) {
				return false;}
		} else if (!gid.equals(other.gid)) {
			return false;}
		return true;
	}
}
