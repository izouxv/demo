/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.po;

import java.io.Serializable;

import com.alibaba.fastjson.JSONArray;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月1日 上午9:16:17
 * @$
 * @Administrator
 * @explain 
 */

public class ProtoHead implements Serializable {
	
	/**
	 * 
	 */
	private static final long serialVersionUID = -8599028041118066173L;
	
	//协议版本
	private String pv;
	//客户端标识
	private String client;
	//协议类型
	private Byte type;
	//子协议
	private Byte subType;
	//分片标志
	private Byte flage;
	//目标个数
	private Byte scMulitca;
	//目标ID
	private JSONArray desID;
	//来源ID
	private Integer sourceID;
	//消息体长度
	private Integer length;
	
	public ProtoHead() {
		super();
	}
	
	public ProtoHead(String pv, String client, Byte type, Byte subType, Byte flage, Byte scMulitca, JSONArray desID,
			Integer sourceID, Integer length) {
		super();
		this.pv = pv;
		this.client = client;
		this.type = type;
		this.subType = subType;
		this.flage = flage;
		this.scMulitca = scMulitca;
		this.desID = desID;
		this.sourceID = sourceID;
		this.length = length;
	}
	
	@Override
	public String toString() {
		return "ProtoHead [pv=" + pv + ", client=" + client + ", type=" + type + ", subType=" + subType + ", flage="
				+ flage + ", scMulitca=" + scMulitca + ", desID=" + desID + ", sourceID=" + sourceID + ", length="
				+ length + "]";
	}
	
	public String getPv() {
		return pv;
	}
	public void setPv(String pv) {
		this.pv = pv;
	}
	public String getClient() {
		return client;
	}
	public void setClient(String client) {
		this.client = client;
	}
	public Byte getType() {
		return type;
	}
	public void setType(Byte type) {
		this.type = type;
	}
	public Byte getSubType() {
		return subType;
	}
	public void setSubType(Byte subType) {
		this.subType = subType;
	}
	public Byte getFlage() {
		return flage;
	}
	public void setFlage(Byte flage) {
		this.flage = flage;
	}
	public Byte getScMulitca() {
		return scMulitca;
	}
	public void setScMulitca(Byte scMulitca) {
		this.scMulitca = scMulitca;
	}
	public JSONArray getDesID() {
		return desID;
	}
	public void setDesID(JSONArray desID) {
		this.desID = desID;
	}
	public Integer getSourceID() {
		return sourceID;
	}
	public void setSourceID(Integer sourceID) {
		this.sourceID = sourceID;
	}
	public Integer getLength() {
		return length;
	}
	public void setLength(Integer length) {
		this.length = length;
	}

}

