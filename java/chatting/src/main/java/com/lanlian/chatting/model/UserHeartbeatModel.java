/** 
 *<p>开发公司：	鹏联优思 <p>
 *<p>版权所有：	鹏联优思 <p>
 *<p>责任人：	王东阳    <p> 
 *<p>网址：www.penslink.com <p>
 */

package com.lanlian.chatting.model;

import java.io.Serializable;

import com.alibaba.fastjson.annotation.JSONField;

/**
 * @author  wangdyq
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年8月2日 上午10:50:11
 * @explain 用户上报心跳信息
 */

public class UserHeartbeatModel implements Serializable {

	private static final long serialVersionUID = 995184346032174833L;

	private Integer uid;
	
	@JSONField(name="group_id")
	private Integer groupId;//群id
	
	@JSONField(name="owner_id")
	private Integer ownerId;
	
	@JSONField(name="group_type")
	private Integer groupType;
	
	@JSONField(name="group_mode")
	private Integer groupMode;
	
	@JSONField(name="chan_id")
	private Integer chanId;
	
	@JSONField(name="room_id")
	private Integer roomId;
	
	private Double lat;
	private Double lng;
	
	@JSONField(name="device_status")
	private Integer deviceStatus;
	
	@JSONField(name="in_time")
	private String inTime;
	
	@JSONField(name="now_time")
	private String nowTime;

	public UserHeartbeatModel() {}

	public UserHeartbeatModel(Integer uid, Integer groupId, Integer ownerId, Integer groupType, Integer groupMode,
			Integer chanId, Integer roomId, Double lat, Double lng, Integer deviceStatus, String inTime,
			String nowTime) {
		super();
		this.uid = uid;
		this.groupId = groupId;
		this.ownerId = ownerId;
		this.groupType = groupType;
		this.groupMode = groupMode;
		this.chanId = chanId;
		this.roomId = roomId;
		this.lat = lat;
		this.lng = lng;
		this.deviceStatus = deviceStatus;
		this.inTime = inTime;
		this.nowTime = nowTime;
	}

	@Override
	public String toString() {
		return "UserHeartbeatModel [uid=" + uid + ", groupId=" + groupId + ", ownerId=" + ownerId + ", groupType="
				+ groupType + ", groupMode=" + groupMode + ", chanId=" + chanId + ", roomId=" + roomId + ", lat=" + lat
				+ ", lng=" + lng + ", deviceStatus=" + deviceStatus + ", inTime=" + inTime + ", nowTime=" + nowTime
				+ "]";
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((chanId == null) ? 0 : chanId.hashCode());
		result = prime * result + ((groupId == null) ? 0 : groupId.hashCode());
		result = prime * result + ((roomId == null) ? 0 : roomId.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		UserHeartbeatModel other = (UserHeartbeatModel) obj;
		if (chanId == null) {
			if (other.chanId != null)
				return false;
		} else if (!chanId.equals(other.chanId))
			return false;
		if (groupId == null) {
			if (other.groupId != null)
				return false;
		} else if (!groupId.equals(other.groupId))
			return false;
		if (roomId == null) {
			if (other.roomId != null)
				return false;
		} else if (!roomId.equals(other.roomId))
			return false;
		return true;
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public Integer getGroupId() {
		return groupId;
	}

	public void setGroupId(Integer groupId) {
		this.groupId = groupId;
	}

	public Integer getOwnerId() {
		return ownerId;
	}

	public void setOwnerId(Integer ownerId) {
		this.ownerId = ownerId;
	}

	public Integer getGroupType() {
		return groupType;
	}

	public void setGroupType(Integer groupType) {
		this.groupType = groupType;
	}

	public Integer getGroupMode() {
		return groupMode;
	}

	public void setGroupMode(Integer groupMode) {
		this.groupMode = groupMode;
	}

	public Integer getChanId() {
		return chanId;
	}

	public void setChanId(Integer chanId) {
		this.chanId = chanId;
	}

	public Integer getRoomId() {
		return roomId;
	}

	public void setRoomId(Integer roomId) {
		this.roomId = roomId;
	}

	public Double getLat() {
		return lat;
	}

	public void setLat(Double lat) {
		this.lat = lat;
	}

	public Double getLng() {
		return lng;
	}

	public void setLng(Double lng) {
		this.lng = lng;
	}

	public Integer getDeviceStatus() {
		return deviceStatus;
	}

	public void setDeviceStatus(Integer deviceStatus) {
		this.deviceStatus = deviceStatus;
	}

	public String getInTime() {
		return inTime;
	}

	public void setInTime(String inTime) {
		this.inTime = inTime;
	}

	public String getNowTime() {
		return nowTime;
	}

	public void setNowTime(String nowTime) {
		this.nowTime = nowTime;
	}

}
