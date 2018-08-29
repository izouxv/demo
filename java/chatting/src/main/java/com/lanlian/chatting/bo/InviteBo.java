package com.lanlian.chatting.bo;
/**
 * @Description: TODO
 * @author: 李大双
 * @date: 2017年6月30日 下午3:26:06
 * @version: V1.0
 */
public class InviteBo {
	//APP端开启上报后获得的邀请码
	private String inviteCode;
	//调用接口时微信用户的ID
	private int uid;
	public InviteBo() {
	}
	public InviteBo(String inviteCode, int uid) {
		this.inviteCode = inviteCode;
		this.uid = uid;
	}
	public String getInviteCode() {
		return inviteCode;
	}
	public void setInviteCode(String inviteCode) {
		this.inviteCode = inviteCode;
	}
	public int getUid() {
		return uid;
	}
	public void setUid(int uid) {
		this.uid = uid;
	}
	@Override
	public String toString() {
		return "InviteBo [inviteCode=" + inviteCode + ", uid=" + uid + "]";
	}
}
