package com.lanlian.chatting.bo;
/**
 * @Description: TODO
 * @author: 李大双
 * @date: 2017年6月30日 下午2:59:47
 * @version: V1.0
 */
public class MessagePageBo {
	//群编号
	int gid;
	//开始id，向下查询
	int startid;
	//结束id,向上查询  startid,endid 二选一
	int endid;
	//取出个数
	int count;
	
	public MessagePageBo() {
		super();
	}
	public MessagePageBo(int gid, int startid, int endid, int count) {
		super();
		this.gid = gid;
		this.startid = startid;
		this.endid = endid;
		this.count = count;
	}

	public int getGid() {
		return gid;
	}
	public void setGid(int gid) {
		this.gid = gid;
	}
	public int getStartid() {
		return startid;
	}
	public void setStartid(int startid) {
		this.startid = startid;
	}
	public int getEndid() {
		return endid;
	}
	public void setEndid(int endid) {
		this.endid = endid;
	}
	public int getCount() {
		return count;
	}
	public void setCount(int count) {
		this.count = count;
	}
	
}
