/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @Title SsoPo.java
 * @Package com.lanlian.chatting.entity
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月8日 下午2:50:09
 * @explain 验证登录信息的属性类；
 */

public class SsoPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 207476572311371315L;

	private Integer uid;// id
	private String username;// 用户名
	private String password;// 密码
	private String salt;// 盐值
	private String nickname;// 昵称
	private String sessionName;// token
	private Integer state;// 用户状态
	private Integer loginState;// 登录状态
	private Integer errorCode;
	private String ip;// iP
	private String loginDevice;// 设备型号
	private String imei;// 唯一标识
	private Integer codeType;// 短信类型
	private String source;// 来源

	public SsoPo() {
		super();
	}

	public SsoPo(Integer uid, String username, String password, String salt, String nickname, String sessionName, Integer state,
			Integer loginState, Integer errorCode, String ip, String loginDevice, String imei, Integer codeType, String source) {
		super();
		this.uid = uid;
		this.username = username;
		this.password = password;
		this.salt = salt;
		this.nickname = nickname;
		this.sessionName = sessionName;
		this.state = state;
		this.loginState = loginState;
		this.errorCode = errorCode;
		this.ip = ip;
		this.loginDevice = loginDevice;
		this.imei = imei;
		this.codeType = codeType;
		this.source = source;
	}

	@Override
	public String toString() {
		return "SsoPo [uid=" + uid + ", username=" + username + ", password=" + password + ", salt=" + salt
				+ ", nickname=" + nickname + ", sessionName=" + sessionName + ", state=" + state + ", loginState="
				+ loginState + ", errorCode=" + errorCode + ", ip=" + ip + ", loginDevice=" + loginDevice + ", imei="
				+ imei + ", codeType=" + codeType + ", source=" + source + "]";
	}

	public String getSource() {
		return source;
	}

	public void setSource(String source) {
		this.source = source;
	}

	public Integer getCodeType() {
		return codeType;
	}

	public void setCodeType(Integer codeType) {
		this.codeType = codeType;
	}

	public String getSalt() {
		return salt;
	}

	public void setSalt(String salt) {
		this.salt = salt;
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(Integer uid) {
		this.uid = uid;
	}

	public String getUsername() {
		return username;
	}

	public void setUsername(String username) {
		this.username = username;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
	}

	public String getSessionName() {
		return sessionName;
	}

	public void setSessionName(String sessionName) {
		this.sessionName = sessionName;
	}

	public Integer getState() {
		return state;
	}

	public void setState(Integer state) {
		this.state = state;
	}

	public Integer getLoginState() {
		return loginState;
	}

	public void setLoginState(Integer loginState) {
		this.loginState = loginState;
	}

	public Integer getErrorCode() {
		return errorCode;
	}

	public void setErrorCode(Integer errorCode) {
		this.errorCode = errorCode;
	}

	public String getIp() {
		return ip;
	}

	public void setIp(String ip) {
		this.ip = ip;
	}

	public String getLoginDevice() {
		return loginDevice;
	}

	public void setLoginDevice(String loginDevice) {
		this.loginDevice = loginDevice;
	}

	public String getImei() {
		return imei;
	}

	public void setImei(String imei) {
		this.imei = imei;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((username == null) ? 0 : username.hashCode());
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
		SsoPo other = (SsoPo) obj;
		if (username == null) {
			if (other.username != null) {
				return false;
			}
		} else if (!username.equals(other.username)) {
			return false;
		}
		return true;
	}
}
