package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @Title DadaGroupPo.java
 * @Package com.lanlian.chatting.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午2:02:46
 * @explain 此类用于永久群数据的属性类
 */
public class DadaGroupPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -746747989192511868L;

	private Integer gid;// 群id
	private Integer uid;// 群主uid
	private Integer upid;// 开启上报的用户id
	private String gname;// 群名称
	private Integer avatar;// 群头像
	private String announcement;// 群公告
	private Timestamp createTime;// 群创建时间
	private Timestamp modifyTime;// 群修改时间
	private Double longitude;// 经度
	private Double latitude;// 纬度
	private String address;// 对应地址
	private Integer groupState;// 上报开关
	private String inviteCode;// 邀请码
	private Timestamp dataCreateTime;// 数据创建时间
	private Timestamp dataModifyTime;// 数据修改时间
	private Integer dataState;// 数据状态
	private Integer newgid;// 返回的新群id

	public DadaGroupPo() {
		super();
	}

	public DadaGroupPo(Integer gid, Integer uid, Integer upid, String gname, Integer avatar, String announcement,
			Timestamp createTime, Timestamp modifyTime, Double longitude, Double latitude, String address,
			Integer groupState, String inviteCode, Timestamp dataCreateTime, Timestamp dataModifyTime,
			Integer dataState, Integer newgid) {
		super();
		this.gid = gid;
		this.uid = uid;
		this.upid = upid;
		this.gname = gname;
		this.avatar = avatar;
		this.announcement = announcement;
		this.createTime = createTime;
		this.modifyTime = modifyTime;
		this.longitude = longitude;
		this.latitude = latitude;
		this.address = address;
		this.groupState = groupState;
		this.inviteCode = inviteCode;
		this.dataCreateTime = dataCreateTime;
		this.dataModifyTime = dataModifyTime;
		this.dataState = dataState;
		this.newgid = newgid;
	}

	@Override
	public String toString() {
		return "DadaGroupPo [gid=" + gid + ", uid=" + uid + ", upid=" + upid + ", gname=" + gname + ", avatar=" + avatar
				+ ", announcement=" + announcement + ", createTime=" + createTime + ", modifyTime=" + modifyTime
				+ ", longitude=" + longitude + ", latitude=" + latitude + ", address=" + address + ", groupState="
				+ groupState + ", inviteCode=" + inviteCode + ", dataCreateTime=" + dataCreateTime + ", dataModifyTime="
				+ dataModifyTime + ", dataState=" + dataState + ", newgid=" + newgid + "]";
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

	public Integer getUpid() {
		return upid;
	}

	public void setUpid(Integer upid) {
		this.upid = upid;
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

	public Double getLongitude() {
		return longitude;
	}

	public void setLongitude(Double longitude) {
		this.longitude = longitude;
	}

	public Double getLatitude() {
		return latitude;
	}

	public void setLatitude(Double latitude) {
		this.latitude = latitude;
	}

	public String getAddress() {
		return address;
	}

	public void setAddress(String address) {
		this.address = address;
	}

	public String getInviteCode() {
		return inviteCode;
	}

	public void setInviteCode(String inviteCode) {
		this.inviteCode = inviteCode;
	}

	public Integer getGroupState() {
		return groupState;
	}

	public void setGroupState(Integer groupState) {
		this.groupState = groupState;
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

	public Integer getDataState() {
		return dataState;
	}

	public void setDataState(Integer dataState) {
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
		result = prime * result + ((uid == null) ? 0 : uid.hashCode());
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
		DadaGroupPo other = (DadaGroupPo) obj;
		if (uid == null) {
			if (other.uid != null) {
				return false;}
		} else if (!uid.equals(other.uid)) {
			return false;
		}
		return true;
	}

}
