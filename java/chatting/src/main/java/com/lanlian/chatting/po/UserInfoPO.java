package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;
import java.util.Date;

/**
 * @author wdyqxx
 * @version 2016年12月30日 下午1:38:05
 * @explain 此类用于用户的个人信息等操作。 属性： pid:用户的id,nickName:昵称,username:注册用户名,
 *          password:密码,salt:加盐,gender:性别1：男/2：女,
 *          birthday:生日,avatar:头像id（以后会是头像地址）, province:省,city:市,address:详细地址,
 *          jobId:职业ID，对应职位表,signature:个性签名,commonEmail:常用邮箱,
 *          realName:真实姓名,identityCard:身份证号码,
 *          isCertification:是否实名认证0：否；1：是；,creditValues:信用点,
 *          point:积分,gradeId:等级,status:用户状态0：正常；1：注册&没激活；2：注销；3：已删除；4：禁用；,
 *          regTime:注册时间,regIp:注册ip,lastLoginTime:最后登录时间,
 *          lastActiveTime:最后活跃时间,gmtModifyTime:最近修改时间,
 *          isFirstLogin:首次登录标记0：非首次（老鸟）0/1：首次登录（新鸟）,
 *          方法：无参构造器，有参构造器，get/set方法，toString方法，hashCode比较PID；
 */
public class UserInfoPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 410734197657535238L;

	private int uid;// 用户的uid
	private String username;// 用户名
	private String password;// 密码
	private String salt;// 加盐
	private int status;// 用户状态1：正常；2：禁用；3：注销；4：删除；
	private String nickname;// 昵称
	private String realName;// 真实姓名
	private int isCertification;// 是否实名认证1：是；2：否；
	private String identityCard;// 身份证号码
	private int gender;// 性别1：男/2：女/0:无
	private Date birthday;// 生日
	private int avatar;// 头像id（以后会是头像地址）
	private String province;// 省
	private String city;// 市
	private String signature;// 个性签名
	private long regIp;// 注册ip
	private String address;// 详细地址
	private int jobId;// 职业ID，对应职位表
	private String commonEmail;// 常用邮箱
	private int creditValues;// 信用点
	private int point;// 积分
	private int gradeId;// 等级ID
	private long lastIp;// 最后登录ip
	private String isFirstLogin;// 首次登录标记1：首次登录（新鸟）；2：非首次（老鸟）
	private Timestamp regTime;// 注册时间
	private Timestamp lastLoginTime;// 最后登录时间
	private Timestamp lastActive;// 最后活跃时间
	private Timestamp modifyTime;// 最近修改时间

	public UserInfoPO() {
		super();
	}

	public UserInfoPO(int uid, String nickname, String username, String password, String salt, int gender,
			Date birthday, int avatar, String province, String city, String address, int jobId, String signature,
			String commonEmail, String realName, String identityCard, int isCertification, int creditValues, int point,
			int gradeId, int status, Timestamp regTime, long regIp, Timestamp lastLoginTime, long lastIp,
			Timestamp lastActive, Timestamp modifyTime, String isFirstLogin) {
		super();
		this.uid = uid;
		this.nickname = nickname;
		this.username = username;
		this.password = password;
		this.salt = salt;
		this.gender = gender;
		this.birthday = birthday;
		this.avatar = avatar;
		this.province = province;
		this.city = city;
		this.address = address;
		this.jobId = jobId;
		this.signature = signature;
		this.commonEmail = commonEmail;
		this.realName = realName;
		this.identityCard = identityCard;
		this.isCertification = isCertification;
		this.creditValues = creditValues;
		this.point = point;
		this.gradeId = gradeId;
		this.status = status;
		this.regTime = regTime;
		this.regIp = regIp;
		this.lastLoginTime = lastLoginTime;
		this.lastIp = lastIp;
		this.lastActive = lastActive;
		this.modifyTime = modifyTime;
		this.isFirstLogin = isFirstLogin;
	}

	@Override
	public String toString() {
		return "UserInfoPO [uid=" + uid + ", nickname=" + nickname + ", username=" + username + ", password=" + password
				+ ", salt=" + salt + ", gender=" + gender + ", birthday=" + birthday + ", avatar=" + avatar
				+ ", province=" + province + ", city=" + city + ", address=" + address + ", jobId=" + jobId
				+ ", signature=" + signature + ", commonEmail=" + commonEmail + ", realName=" + realName
				+ ", identityCard=" + identityCard + ", isCertification=" + isCertification + ", creditValues="
				+ creditValues + ", point=" + point + ", gradeId=" + gradeId + ", status=" + status + ", regTime="
				+ regTime + ", regIp=" + regIp + ", lastLoginTime=" + lastLoginTime + ", lastIp=" + lastIp
				+ ", lastActive=" + lastActive + ", modifyTime=" + modifyTime + ", isFirstLogin=" + isFirstLogin + "]";
	}

	public int getPid() {
		return uid;
	}

	public void setPid(int uid) {
		this.uid = uid;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
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

	public String getSalt() {
		return salt;
	}

	public void setSalt(String salt) {
		this.salt = salt;
	}

	public int getGender() {
		return gender;
	}

	public void setGender(int gender) {
		this.gender = gender;
	}

	public Date getBirthday() {
		return birthday;
	}

	public void setBirthday(Date birthday) {
		this.birthday = birthday;
	}

	public int getAvatar() {
		return avatar;
	}

	public void setAvatar(int avatar) {
		this.avatar = avatar;
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

	public String getAddress() {
		return address;
	}

	public void setAddress(String address) {
		this.address = address;
	}

	public int getJobId() {
		return jobId;
	}

	public void setJobId(int jobId) {
		this.jobId = jobId;
	}

	public String getSignature() {
		return signature;
	}

	public void setSignature(String signature) {
		this.signature = signature;
	}

	public String getCommonEmail() {
		return commonEmail;
	}

	public void setCommonEmail(String commonEmail) {
		this.commonEmail = commonEmail;
	}

	public String getRealName() {
		return realName;
	}

	public void setRealName(String realName) {
		this.realName = realName;
	}

	public String getIdentityCard() {
		return identityCard;
	}

	public void setIdentityCard(String identityCard) {
		this.identityCard = identityCard;
	}

	public int getIsCertification() {
		return isCertification;
	}

	public void setIsCertification(int isCertification) {
		this.isCertification = isCertification;
	}

	public int getCreditValues() {
		return creditValues;
	}

	public void setCreditValues(int creditValues) {
		this.creditValues = creditValues;
	}

	public int getPoint() {
		return point;
	}

	public void setPoint(int point) {
		this.point = point;
	}

	public int getGradeId() {
		return gradeId;
	}

	public void setGradeId(int gradeId) {
		this.gradeId = gradeId;
	}

	public int getStatus() {
		return status;
	}

	public void setStatus(int status) {
		this.status = status;
	}

	public Timestamp getRegTime() {
		return regTime;
	}

	public void setRegTime(Timestamp regTime) {
		this.regTime = regTime;
	}

	public long getRegIp() {
		return regIp;
	}

	public void setRegIp(long regIp) {
		this.regIp = regIp;
	}

	public Timestamp getLastLoginTime() {
		return lastLoginTime;
	}

	public void setLastLoginTime(Timestamp lastLoginTime) {
		this.lastLoginTime = lastLoginTime;
	}

	public long getLastIp() {
		return lastIp;
	}

	public void setLastIp(long lastIp) {
		this.lastIp = lastIp;
	}

	public Timestamp getLastActive() {
		return lastActive;
	}

	public void setLastActive(Timestamp lastActive) {
		this.lastActive = lastActive;
	}

	public Timestamp getModifyTime() {
		return modifyTime;
	}

	public void setModifyTime(Timestamp modifyTime) {
		this.modifyTime = modifyTime;
	}

	public String getIsFirstLogin() {
		return isFirstLogin;
	}

	public void setIsFirstLogin(String isFirstLogin) {
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
		UserInfoPO other = (UserInfoPO) obj;
		if (uid != other.uid) {
			return false;
		}
		return true;
	}
}
