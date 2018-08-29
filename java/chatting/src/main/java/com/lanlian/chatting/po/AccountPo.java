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
 * @Title AccountPo.java
 * @Package cn.lanlian.ccat.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年5月16日 下午4:42:49
 * @explain account-rpc
 */

public class AccountPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 2026141043155159269L;

	private int errorCode;// rpc状态码
	private String source;// 来源
	private int uid;// 用户的id
	private String username;// 用户名
	private int state;// 用户状态0：正常；1：注册&没激活；2：注销；3：已删除；4：禁用；
	private String phone;// 手机号
	private String email;// 邮箱
	private String nickname;// 昵称
	private int gender;// 性别1：男/2：女
	private long birthday;// 生日
	private int avatar;// 头像id
	private String signature;// 个性签名
	private String province;// 省
	private String city;// 市
	private String userAddress;// 详细地址
	private String realname;// 真实姓名
	private String identityCard;// 身份证号码
	private int isCertification;// 是否实名认证
	private int creditValues;// 信用点
	private int userPoint;// 积分
	private int userJobId;// 职业ID，对应职位表
	private int userGradeId;// 等级ID
	private int isFirstLogin;// 是否第一次登录
	private long lastLoginTime;// 最后登录时间
	private long lastLoginIp;// 最后登录IP
	private long lastActive;// 最后活跃时间
	private long regTime;// 注册时间
	private long regIp;// 注册IP
	private long userModify;// 最后修改时间
	private long createTime;// 创建时间

	public AccountPo() {
		super();
	}

	public AccountPo(int errorCode, String source, int uid, String username, int state, String phone, String email,
			String nickname, int gender, long birthday, int avatar, String signature, String province, String city,
			String userAddress, String realname, String identityCard, int isCertification, int creditValues,
			int userPoint, int userJobId, int userGradeId, int isFirstLogin, long lastLoginTime, long lastLoginIp,
			long lastActive, long regTime, long regIp, long userModify, long createTime) {
		super();
		this.errorCode = errorCode;
		this.source = source;
		this.uid = uid;
		this.username = username;
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
		this.realname = realname;
		this.identityCard = identityCard;
		this.isCertification = isCertification;
		this.creditValues = creditValues;
		this.userPoint = userPoint;
		this.userJobId = userJobId;
		this.userGradeId = userGradeId;
		this.isFirstLogin = isFirstLogin;
		this.lastLoginTime = lastLoginTime;
		this.lastLoginIp = lastLoginIp;
		this.lastActive = lastActive;
		this.regTime = regTime;
		this.regIp = regIp;
		this.userModify = userModify;
		this.createTime = createTime;
	}

	@Override
	public String toString() {
		return "AccountPo [errorCode=" + errorCode + ", source=" + source + ", uid=" + uid + ", username=" + username
				+ ", state=" + state + ", phone=" + phone + ", email=" + email + ", nickname=" + nickname + ", gender="
				+ gender + ", birthday=" + birthday + ", avatar=" + avatar + ", signature=" + signature + ", province="
				+ province + ", city=" + city + ", userAddress=" + userAddress + ", realname=" + realname
				+ ", identityCard=" + identityCard + ", isCertification=" + isCertification + ", creditValues="
				+ creditValues + ", userPoint=" + userPoint + ", userJobId=" + userJobId + ", userGradeId="
				+ userGradeId + ", isFirstLogin=" + isFirstLogin + ", lastLoginTime=" + lastLoginTime + ", lastLoginIp="
				+ lastLoginIp + ", lastActive=" + lastActive + ", regTime=" + regTime + ", regIp=" + regIp
				+ ", userModify=" + userModify + ", createTime=" + createTime + "]";
	}

	public String getSource() {
		return source;
	}

	public void setSource(String source) {
		this.source = source;
	}

	public String getRealname() {
		return realname;
	}

	public void setRealname(String realname) {
		this.realname = realname;
	}

	public String getIdentityCard() {
		return identityCard;
	}

	public void setIdentityCard(String identityCard) {
		this.identityCard = identityCard;
	}

	public Integer getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public String getUsername() {
		return username;
	}

	public void setUsername(String username) {
		this.username = username;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	public String getPhone() {
		return phone;
	}

	public void setPhone(String phone) {
		this.phone = phone;
	}

	public int getState() {
		return state;
	}

	public void setState(int state) {
		this.state = state;
	}

	public long getLastLoginTime() {
		return lastLoginTime;
	}

	public void setLastLoginTime(long lastLoginTime) {
		this.lastLoginTime = lastLoginTime;
	}

	public long getCreateTime() {
		return createTime;
	}

	public void setCreateTime(long createTime) {
		this.createTime = createTime;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
	}

	public int getIsCertification() {
		return isCertification;
	}

	public void setIsCertification(int isCertification) {
		this.isCertification = isCertification;
	}

	public int getErrorCode() {
		return errorCode;
	}

	public void setErrorCode(int errorCode) {
		this.errorCode = errorCode;
	}

	public int getGender() {
		return gender;
	}

	public void setGender(int gender) {
		this.gender = gender;
	}

	public long getBirthday() {
		return birthday;
	}

	public void setBirthday(long birthday) {
		this.birthday = birthday;
	}

	public int getAvatar() {
		return avatar;
	}

	public void setAvatar(int avatar) {
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

	public int getUserJobId() {
		return userJobId;
	}

	public void setUserJobId(int userJobId) {
		this.userJobId = userJobId;
	}

	public int getCreditValues() {
		return creditValues;
	}

	public void setCreditValues(int creditValues) {
		this.creditValues = creditValues;
	}

	public int getUserPoint() {
		return userPoint;
	}

	public void setUserPoint(int userPoint) {
		this.userPoint = userPoint;
	}

	public int getUserGradeId() {
		return userGradeId;
	}

	public void setUserGradeId(int userGradeId) {
		this.userGradeId = userGradeId;
	}

	public long getRegTime() {
		return regTime;
	}

	public void setRegTime(long regTime) {
		this.regTime = regTime;
	}

	public long getRegIp() {
		return regIp;
	}

	public void setRegIp(long regIp) {
		this.regIp = regIp;
	}

	public long getLastLoginIp() {
		return lastLoginIp;
	}

	public void setLastLoginIp(long lastLoginIp) {
		this.lastLoginIp = lastLoginIp;
	}

	public long getLastActive() {
		return lastActive;
	}

	public void setLastActive(long lastActive) {
		this.lastActive = lastActive;
	}

	public long getUserModify() {
		return userModify;
	}

	public void setUserModify(long userModify) {
		this.userModify = userModify;
	}

	public int getIsFirstLogin() {
		return isFirstLogin;
	}

	public void setIsFirstLogin(int isFirstLogin) {
		this.isFirstLogin = isFirstLogin;
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
			return true;
		}
		if (obj == null) {
			return false;
		}
		if (getClass() != obj.getClass()) {
			return false;
		}
		AccountPo other = (AccountPo) obj;
		if (uid != other.uid) {
			return false;
		}
		return true;
	}

}
