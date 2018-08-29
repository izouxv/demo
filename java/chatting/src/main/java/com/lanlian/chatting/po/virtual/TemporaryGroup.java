/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.po.virtual;

import java.io.Serializable;

import com.alibaba.fastjson.JSONArray;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月25日 下午3:46:51
 * @explain 临时群备份属性类
 */

public class TemporaryGroup implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 8890612229717055299L;

	private Integer uid;
	private JSONArray data;
	private long time;

	public TemporaryGroup() {
		super();
	}

	public TemporaryGroup(Integer uid, JSONArray data, long time) {
		super();
		this.uid = uid;
		this.data = data;
		this.time = time;
	}

	@Override
	public String toString() {
		return "TemporaryGroup [uid=" + uid + ", data=" + data + ", time=" + time + "]";
	}

	public long getTime() {
		return time;
	}

	public void setTime(long time) {
		this.time = time;
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public JSONArray getData() {
		return data;
	}

	public void setData(JSONArray data) {
		this.data = data;
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
		TemporaryGroup other = (TemporaryGroup) obj;
		if (uid == null) {
			if (other.uid != null) {
				return false;
			}
		} else if (!uid.equals(other.uid)) {
			return false;
		}
		return true;
	}

}
