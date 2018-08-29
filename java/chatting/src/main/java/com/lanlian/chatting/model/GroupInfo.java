/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.model;

/** 
 * @Title GroupInfo.java
 * @Package com.lanlian.chatting.model
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午8:14:42
 * @explain 创建群后返回的属性类
 */

public class GroupInfo {
	
	private int gid;//群id
	private Integer groupState;//上报开关
	private String inviteCode;//邀请码
	
	public GroupInfo() {
		super();
	}

	public GroupInfo(int gid, Integer groupState, String inviteCode) {
		super();
		this.gid = gid;
		this.groupState = groupState;
		this.inviteCode = inviteCode;
	}

	@Override
	public String toString() {
		return "GroupInfo [gid=" + gid + ", groupState=" + groupState + ", inviteCode=" + inviteCode + "]";
	}

	public int getGid() {
		return gid;
	}

	public void setGid(int gid) {
		this.gid = gid;
	}

	public Integer getGroupState() {
		return groupState;
	}

	public void setGroupState(Integer groupState) {
		this.groupState = groupState;
	}

	public String getInviteCode() {
		return inviteCode;
	}

	public void setInviteCode(String inviteCode) {
		this.inviteCode = inviteCode;
	}

}

