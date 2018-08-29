/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.vo;

import java.io.Serializable;

/**
 * @Title UserInfo.java
 * @Package com.lanlian.chatting.vo
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月12日 下午7:38:39
 * @explain 与前端交互的参数
 */

public class UserInfo implements Serializable {

	private static final long serialVersionUID = 3941282134169637361L;

	private int errorCode;//rpc状态码
	private int uid;//用户的id
	private String username;//用户名
	private String password;//密码
	private String sessionName;//token
	private int loginState;//登录状态
	private int state;//用户状态0：正常；1：注册&没激活；2：注销；3：已删除；4：禁用；
	private String phone;//手机号
	private String email;//邮箱
	private String nickname;//昵称
	private int gender;//性别1：男/2：女
	private long birthday;//生日
	private int avatar;//头像id
	private String signature;//个性签名
	private String province;//省
	private String city;//市
	private String userAddress;//详细地址
	private String isCertification;//是否实名认证
	private int creditValues;//信用点
	private int userPoint;//积分
	private int userJobId;//职业ID，对应职位表
	private int userGradeId;//等级ID
	private int isFirstLogin;//是否第一次登录
	private long lastLoginTime;//最后登录时间
	private long ip;//iP
	private String loginDevice;//设备型号
	private String imei;//唯一标识
	
	public UserInfo() {
		super();
	}

	@Override
	public String toString() {
		return "UserInfo [errorCode=" + errorCode + ", uid=" + uid + ", username=" + username + ", password=" + password
				+ ", sessionName=" + sessionName + ", loginState=" + loginState + ", state=" + state + ", phone="
				+ phone + ", email=" + email + ", nickname=" + nickname + ", gender=" + gender + ", birthday="
				+ birthday + ", avatar=" + avatar + ", signature=" + signature + ", province=" + province + ", city="
				+ city + ", userAddress=" + userAddress + ", isCertification=" + isCertification + ", creditValues="
				+ creditValues + ", userPoint=" + userPoint + ", userJobId=" + userJobId + ", userGradeId="
				+ userGradeId + ", isFirstLogin=" + isFirstLogin + ", lastLoginTime=" + lastLoginTime + ", ip=" + ip
				+ ", loginDevice=" + loginDevice + ", imei=" + imei + "]";
	}

	public UserInfo(int errorCode, int uid, String username, String password, String sessionName, int loginState,
			int state, String phone, String email, String nickname, int gender, long birthday, int avatar,
			String signature, String province, String city, String userAddress, String isCertification,
			int creditValues, int userPoint, int userJobId, int userGradeId, int isFirstLogin, long lastLoginTime,
			long ip, String loginDevice, String imei) {
		super();
		this.errorCode = errorCode;
		this.uid = uid;
		this.username = username;
		this.password = password;
		this.sessionName = sessionName;
		this.loginState = loginState;
		this.state = state;
		this.phone = phone;
		this.email = email;
		this.nickname = nickname;
		this.gender = gender;
		this.birthday = birthday;
		this.avatar = avatar;
		this.signature = signature;
		this.province = province;
		this.city = city;
		this.userAddress = userAddress;
		this.isCertification = isCertification;
		this.creditValues = creditValues;
		this.userPoint = userPoint;
		this.userJobId = userJobId;
		this.userGradeId = userGradeId;
		this.isFirstLogin = isFirstLogin;
		this.lastLoginTime = lastLoginTime;
		this.ip = ip;
		this.loginDevice = loginDevice;
		this.imei = imei;
	}

	public Integer getErrorCode() {
		return errorCode;
	}

	public void setErrorCode(Integer errorCode) {
		this.errorCode = errorCode;
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

	public String getSessionName() {
		return sessionName;
	}

	public void setSessionName(String sessionName) {
		this.sessionName = sessionName;
	}

	public Integer getLoginState() {
		return loginState;
	}

	public void setLoginState(Integer loginState) {
		this.loginState = loginState;
	}

	public Integer getState() {
		return state;
	}

	public void setState(Integer state) {
		this.state = state;
	}

	public String getPhone() {
		return phone;
	}

	public void setPhone(String phone) {
		this.phone = phone;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
	}

	public Integer getGender() {
		return gender;
	}

	public void setGender(Integer gender) {
		this.gender = gender;
	}

	public Long getBirthday() {
		return birthday;
	}

	public void setBirthday(Long birthday) {
		this.birthday = birthday;
	}

	public Integer getAvatar() {
		return avatar;
	}

	public void setAvatar(Integer avatar) {
		this.avatar = avatar;
	}

	public String getSignature() {
		return signature;
	}

	public void setSignature(String signature) {
		this.signature = signature;
	}

	public String getProvince() {
		return province;
	}

	public void setProvince(String province) {
		this.province = province;
	}

	public String getCity() {
		return city;
	}

	public void setCity(String city) {
		this.city = city;
	}

	public String getUserAddress() {
		return userAddress;
	}

	public void setUserAddress(String userAddress) {
		this.userAddress = userAddress;
	}

	public String getIsCertification() {
		return isCertification;
	}

	public void setIsCertification(String isCertification) {
		this.isCertification = isCertification;
	}

	public Integer getCreditValues() {
		return creditValues;
	}

	public void setCreditValues(Integer creditValues) {
		this.creditValues = creditValues;
	}

	public Integer getUserPoint() {
		return userPoint;
	}

	public void setUserPoint(Integer userPoint) {
		this.userPoint = userPoint;
	}

	public Integer getUserJobId() {
		return userJobId;
	}

	public void setUserJobId(Integer userJobId) {
		this.userJobId = userJobId;
	}

	public Integer getUserGradeId() {
		return userGradeId;
	}

	public void setUserGradeId(Integer userGradeId) {
		this.userGradeId = userGradeId;
	}

	public Integer getIsFirstLogin() {
		return isFirstLogin;
	}

	public void setIsFirstLogin(Integer isFirstLogin) {
		this.isFirstLogin = isFirstLogin;
	}

	public Long getLastLoginTime() {
		return lastLoginTime;
	}

	public void setLastLoginTime(Long lastLoginTime) {
		this.lastLoginTime = lastLoginTime;
	}

	public Long getIp() {
		return ip;
	}

	public void setIp(Long ip) {
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
		result = prime * result + uid;
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj) {
			return true;}
		if (obj == null) {
			return false;}
		if (getClass() != obj.getClass()) {
			return false;}
		UserInfo other = (UserInfo) obj;
		if (uid != other.uid) {
			return false;}
		return true;
	}
	
}
