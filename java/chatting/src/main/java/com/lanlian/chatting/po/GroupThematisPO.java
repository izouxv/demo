package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @author wdyqxx
 * @version 2017年1月3日 下午5:14:54
 * @explain 此类用于主题群的操作
 */
public class GroupThematisPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -1429457596965680409L;

	private long thematicGroupId;// 主题群id
	private String tgName;// 主题群名称
	private String tgAvatar;// 头像地址
	private String tgDescribe;// 主题群描述
	private String tgAnnouncement;// 主题群公告
	private long createTime;// 创建时间
	private long modifyTime;// 修改时间
	private long tgStatus;// 群状态0：正常；1：被禁用；2：注销；3：已经删除；

	public GroupThematisPO() {
		super();
	}

	public GroupThematisPO(long thematicGroupId, String tgName, String tgAvatar, String tgDescribe,
			String tgAnnouncement, long createTime, long modifyTime, long tgStatus) {
		super();
		this.thematicGroupId = thematicGroupId;
		this.tgName = tgName;
		this.tgAvatar = tgAvatar;
		this.tgDescribe = tgDescribe;
		this.tgAnnouncement = tgAnnouncement;
		this.createTime = createTime;
		this.modifyTime = modifyTime;
		this.tgStatus = tgStatus;
	}

	@Override
	public String toString() {
		return "GroupThematisPOJO [thematicGroupId=" + thematicGroupId + ", tgName=" + tgName + ", tgAvatar=" + tgAvatar
				+ ", tgDescribe=" + tgDescribe + ", tgAnnouncement=" + tgAnnouncement + ", createTime=" + createTime
				+ ", modifyTime=" + modifyTime + ", tgStatus=" + tgStatus + "]";
	}

	public long getThematicGroupId() {
		return thematicGroupId;
	}

	public void setThematicGroupId(long thematicGroupId) {
		this.thematicGroupId = thematicGroupId;
	}

	public String getTgName() {
		return tgName;
	}

	public void setTgName(String tgName) {
		this.tgName = tgName;
	}

	public String getTgAvatar() {
		return tgAvatar;
	}

	public void setTgAvatar(String tgAvatar) {
		this.tgAvatar = tgAvatar;
	}

	public String getTgDescribe() {
		return tgDescribe;
	}

	public void setTgDescribe(String tgDescribe) {
		this.tgDescribe = tgDescribe;
	}

	public String getTgAnnouncement() {
		return tgAnnouncement;
	}

	public void setTgAnnouncement(String tgAnnouncement) {
		this.tgAnnouncement = tgAnnouncement;
	}

	public long getCreateTime() {
		return createTime;
	}

	public void setCreateTime(long createTime) {
		this.createTime = createTime;
	}

	public long getModifyTime() {
		return modifyTime;
	}

	public void setModifyTime(long modifyTime) {
		this.modifyTime = modifyTime;
	}

	public long getTgStatus() {
		return tgStatus;
	}

	public void setTgStatus(long tgStatus) {
		this.tgStatus = tgStatus;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (thematicGroupId ^ (thematicGroupId >>> 32));
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
		GroupThematisPO other = (GroupThematisPO) obj;
		if (thematicGroupId != other.thematicGroupId) {
			return false;
		}
		return true;
	}
}
