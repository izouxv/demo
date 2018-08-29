/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */
package com.lanlian.chatting.po;

/**
 * @Title WechatPo.java
 * @Package com.lanlian.chatting.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月26日 上午11:02:51
 * @explain WechatPo的详细地址信息；
 */
public class WechatPo {

	private String host;// 地址
	private String port;// 端口
	private String server;// 服务
	private String version;// 版本
	private String pathMsgP;// 父的地址
	private String pathMsgB;// body地址

	public WechatPo() {
		super();
	}

	public WechatPo(String host, String port, String server, String version, String pathMsgP, String pathMsgB) {
		super();
		this.host = host;
		this.port = port;
		this.server = server;
		this.version = version;
		this.pathMsgP = pathMsgP;
		this.pathMsgB = pathMsgB;
	}

	@Override
	public String toString() {
		return "WechatPo [host=" + host + ", port=" + port + ", server=" + server + ", version=" + version
				+ ", pathMsgP=" + pathMsgP + ", pathMsgB=" + pathMsgB + "]";
	}

	public String getHost() {
		return host;
	}

	public void setHost(String host) {
		this.host = host;
	}

	public String getPort() {
		return port;
	}

	public void setPort(String port) {
		this.port = port;
	}

	public String getServer() {
		return server;
	}

	public void setServer(String server) {
		this.server = server;
	}

	public String getVersion() {
		return version;
	}

	public void setVersion(String version) {
		this.version = version;
	}

	public String getPathMsgP() {
		return pathMsgP;
	}

	public void setPathMsgP(String pathMsgP) {
		this.pathMsgP = pathMsgP;
	}

	public String getPathMsgB() {
		return pathMsgB;
	}

	public void setPathMsgB(String pathMsgB) {
		this.pathMsgB = pathMsgB;
	}
}