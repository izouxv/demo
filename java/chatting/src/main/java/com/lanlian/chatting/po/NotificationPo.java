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
 * @Title NotificationPo.java
 * @Package com.lanlian.chatting.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月26日 上午11:02:51
 * @explain Notification的详细地址信息；
 */
public class NotificationPo {

	private String host;// 地址
	private String port;// 端口
	private String server;// 服务
	private String version;// 版本
	private String mobile;// 手机发送

	public NotificationPo() {
		super();
	}

	public NotificationPo(String host, String port, String server, String version, String mobile) {
		super();
		this.host = host;
		this.port = port;
		this.server = server;
		this.version = version;
		this.mobile = mobile;
	}

	@Override
	public String toString() {
		return "NotificationPo [host=" + host + ", port=" + port + ", server=" + server + ", version=" + version
				+ ", mobile=" + mobile + "]";
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

	public String getMobile() {
		return mobile;
	}

	public void setMobile(String mobile) {
		this.mobile = mobile;
	}

}