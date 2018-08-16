/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.plys.rtcs.po;

import com.alibaba.fastjson.JSONObject;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年10月11日 下午6:23:21
 * @explain 
 */

public class ProtoMsg extends Proto<Object> {
	
	private ProtoHead head;
	private JSONObject body;
	
	public ProtoMsg() {
		super();
	}

	public ProtoMsg(ProtoHead head, JSONObject body) {
		super();
		this.head = head;
		this.body = body;
	}

	@Override
	public String toString() {
		return "ProtoMsg [head=" + head + ", body=" + body + "]";
	}

	public ProtoHead getHead() {
		return head;
	}

	public void setHead(ProtoHead head) {
		this.head = head;
	}

	public JSONObject getBody() {
		return body;
	}

	public void setBody(JSONObject body) {
		this.body = body;
	}
}

