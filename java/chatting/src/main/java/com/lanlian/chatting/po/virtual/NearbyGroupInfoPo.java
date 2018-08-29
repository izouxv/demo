/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.po.virtual;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月20日 下午4:47:52
 * @explain
 */

public class NearbyGroupInfoPo {

	private Integer types;// 消息类型
	private Integer tuid;// 自己的id
	private Integer toid;// 对方的id
	private Long time;// 消息发送时间
	private Integer type;// 信息类型：1：文本信息；2：语音；3：图片；
	private String info;// 信息内容

	public NearbyGroupInfoPo() {
		super();
	}

	public NearbyGroupInfoPo(Integer types, Integer tuid, Integer toid, Long time, Integer type, String info) {
		super();
		this.types = types;
		this.tuid = tuid;
		this.toid = toid;
		this.time = time;
		this.type = type;
		this.info = info;
	}

	@Override
	public String toString() {
		return "NearbyGroupInfoPo [types=" + types + ", tuid=" + tuid + ", toid=" + toid + ", time=" + time + ", type="
				+ type + ", info=" + info + "]";
	}

	public Integer getTypes() {
		return types;
	}

	public void setTypes(Integer types) {
		this.types = types;
	}

	public Integer getTuid() {
		return tuid;
	}

	public void setTuid(Integer tuid) {
		this.tuid = tuid;
	}

	public Integer getToid() {
		return toid;
	}

	public void setToid(Integer toid) {
		this.toid = toid;
	}

	public Long getTime() {
		return time;
	}

	public void setTime(Long time) {
		this.time = time;
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

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((tuid == null) ? 0 : tuid.hashCode());
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
		NearbyGroupInfoPo other = (NearbyGroupInfoPo) obj;
		if (tuid == null) {
			if (other.tuid != null) {
				return false;
			}
		} else if (!tuid.equals(other.tuid)) {
			return false;
		}
		return true;
	}

}