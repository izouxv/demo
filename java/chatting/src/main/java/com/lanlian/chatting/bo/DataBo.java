/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.bo;

import java.util.List;

import com.lanlian.chatting.po.FriendsPo;

/** 
 * @Title DataBo.java
 * @Package com.lanlian.chatting.bo
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月27日 下午8:07:12
 * @explain 业务之间传输数据的对象
 */

public class DataBo {
	
	/**
	 * FriendsPo的集合
	 */
	private List<FriendsPo> friendsPos;
	
	/**
	 * uid的集合
	 */
	private List<Integer> uidList;
	
	/**
	 * 返回的异常数据
	 */
	private String json;
	
	public DataBo() {
		super();
	}

	public DataBo(List<FriendsPo> friendsPos, List<Integer> uidList, String json) {
		super();
		this.friendsPos = friendsPos;
		this.uidList = uidList;
		this.json = json;
	}

	@Override
	public String toString() {
		return "DataBo [friendsPos=" + friendsPos + ", uidList=" + uidList + ", json=" + json + "]";
	}

	public List<FriendsPo> getFriendsPos() {
		return friendsPos;
	}

	public void setFriendsPos(List<FriendsPo> friendsPos) {
		this.friendsPos = friendsPos;
	}

	public List<Integer> getUidList() {
		return uidList;
	}

	public void setUidList(List<Integer> uidList) {
		this.uidList = uidList;
	}

	public String getJson() {
		return json;
	}

	public void setJson(String json) {
		this.json = json;
	}
	
}

