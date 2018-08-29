/** 
 *<p>开发公司：	鹏联优思 <p>
 *<p>版权所有：	鹏联优思 <p>
 *<p>责任人：	王东阳    <p> 
 *<p>网址：www.penslink.com <p>
 */

package com.lanlian.chatting.model;

import java.io.Serializable;
import java.util.List;

import com.alibaba.fastjson.annotation.JSONField;

/**
 * @author  wangdyq
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年8月2日 上午10:50:11
 * @explain 用户获取心跳信息
 */

public class UserHeartbeatsModel implements Serializable {

	private static final long serialVersionUID = -1424917671952785533L;
	
	@JSONField(name="group_info")
	private String groupInfo;//群信息
	
	@JSONField(name="heartbeats")
	private List<UserHeartbeatModel> heartbeats;
	
	public UserHeartbeatsModel() {}

	public UserHeartbeatsModel(String groupInfo, List<UserHeartbeatModel> heartbeats) {
		super();
		this.groupInfo = groupInfo;
		this.heartbeats = heartbeats;
	}

	@Override
	public String toString() {
		return "UserHeartbeatsModel [groupInfo=" + groupInfo + ", heartbeats=" + heartbeats + "]";
	}

	public String getGroupInfo() {
		return groupInfo;
	}

	public void setGroupInfo(String groupInfo) {
		this.groupInfo = groupInfo;
	}

	public List<UserHeartbeatModel> getHeartbeats() {
		return heartbeats;
	}

	public void setHeartbeats(List<UserHeartbeatModel> heartbeats) {
		this.heartbeats = heartbeats;
	}

}
