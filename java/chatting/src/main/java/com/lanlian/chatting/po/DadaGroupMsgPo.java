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
 * @Title DadaGroupMsgPo.java
 * @Package com.lanlian.chatting.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午3:00:10
 * @explain 群信息实时上报消息属性类
 */

public class DadaGroupMsgPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -7779129046926289806L;

	private Integer mid;// 消息id
	private Integer upuid;
	private Integer gid;// 群id
	private Integer uid;// 用户id
	private Integer type;// 类型
	private String info;// 内容
	private Timestamp sendTime;// 发送时间
	private Timestamp createTime;// 表记录创建时间
	private Timestamp updateTime;// 表修改时间
	private Integer stateTable;// 消息状态

	public DadaGroupMsgPo() {
		super();
	}

	public DadaGroupMsgPo(Integer upuid, Integer mid, Integer gid, Integer uid, Integer type, String info,
			Timestamp sendTime, Timestamp createTime, Timestamp updateTime, Integer stateTable) {
		super();
		this.upuid = upuid;
		this.mid = mid;
		this.gid = gid;
		this.uid = uid;
		this.type = type;
		this.info = info;
		this.sendTime = sendTime;
		this.createTime = createTime;
		this.updateTime = updateTime;
		this.stateTable = stateTable;
	}

	@Override
	public String toString() {
		return "LiveChatMessagePo [upuid" + upuid + ",mid=" + mid + ", gid=" + gid + ", uid=" + uid + ", type=" + type
				+ ", info=" + info + ", sendTime=" + sendTime + ", createTime=" + createTime + ", updateTime="
				+ updateTime + ", stateTable=" + stateTable + "]";
	}

	public Integer getUpuid() {
		return upuid;
	}

	public void setUpuid(Integer upuid) {
		this.upuid = upuid;
	}

	public Integer getMid() {
		return mid;
	}

	public void setMid(Integer mid) {
		this.mid = mid;
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

	public Integer getType() {
		return type;
	}

	public void setType(Integer type) {
		this.type = type;
	}

	public String getInfo() {
		return info;
	}

	public void setInfo(String info) {
		this.info = info;
	}

	public Timestamp getSendTime() {
		return sendTime;
	}

	public void setSendTime(Timestamp sendTime) {
		this.sendTime = sendTime;
	}

	public Timestamp getCreateTime() {
		return createTime;
	}

	public void setCreateTime(Timestamp createTime) {
		this.createTime = createTime;
	}

	public Timestamp getUpdateTime() {
		return updateTime;
	}

	public void setUpdateTime(Timestamp updateTime) {
		this.updateTime = updateTime;
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
		result = prime * result + ((mid == null) ? 0 : mid.hashCode());
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
		DadaGroupMsgPo other = (DadaGroupMsgPo) obj;
		if (mid == null) {
			if (other.mid != null) {
				return false;}
		} else if (!mid.equals(other.mid)) {
			return false;
		}
		return true;
	}

}
