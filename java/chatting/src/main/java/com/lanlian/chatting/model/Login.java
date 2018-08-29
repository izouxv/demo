/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.model;

import java.io.Serializable;
import java.util.Set;

/**
 * @Title UserInfo.java
 * @Package com.lanlian.chatting.vo
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月12日 下午7:38:39
 * @explain 返回到前端的参数
 */

public class Login implements Serializable {

	private static final long serialVersionUID = 3941282134169637361L;

	private int errorCode;// rpc状态码
	private int uid;// 用户的id
	private String username;// 用户名
	private String token;// token
	private int loginState;// 登录状态
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
	private String address;// 详细地址
	private String realname;// 真实姓名
	private String identityCard;// 身份证号码
	private int isCertification;// 是否实名认证
	private int creditValues;// 信用点
	private int point;// 积分
	private int job;// 职业ID，对应职位表
	private int grade;// 等级ID
	private Set<Integer> gids; // 16个临时群id

	public Login() {
		super();
	}

	@Override
	public String toString() {
		return "Login [errorCode=" + errorCode + ", uid=" + uid + ", username=" + username + ", token=" + token
				+ ", loginState=" + loginState + ", state=" + state + ", phone=" + phone + ", email=" + email
				+ ", nickname=" + nickname + ", gender=" + gender + ", birthday=" + birthday + ", avatar=" + avatar
				+ ", signature=" + signature + ", province=" + province + ", city=" + city + ", address=" + address
				+ ", realname=" + realname + ", identityCard=" + identityCard + ", isCertification=" + isCertification
				+ ", creditValues=" + creditValues + ", point=" + point + ", job=" + job + ", grade=" + grade
				+ ", gids=" + gids + "]";
	}

	public Login(int errorCode, int uid, String username, String token, int loginState, int state, String phone,
			String email, String nickname, int gender, long birthday, int avatar, String signature, String province,
			String city, String address, String realname, String identityCard, int isCertification, int creditValues,
			int point, int job, int grade, Set<Integer> gids) {
		super();
		this.errorCode = errorCode;
		this.uid = uid;
		this.username = username;
		this.token = token;
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
		this.address = address;
		this.realname = realname;
		this.identityCard = identityCard;
		this.isCertification = isCertification;
		this.creditValues = creditValues;
		this.point = point;
		this.job = job;
		this.grade = grade;
		this.gids = gids;
	}

	public int getErrorCode() {
		return errorCode;
	}

	public Set<Integer> getGids() {
		return gids;
	}

	public void setGids(Set<Integer> gids) {
		this.gids = gids;
	}

	public void setErrorCode(int errorCode) {
		this.errorCode = errorCode;
	}

	public int getUid() {
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

	public String getToken() {
		return token;
	}

	public void setToken(String token) {
		this.token = token;
	}

	public int getLoginState() {
		return loginState;
	}

	public void setLoginState(int loginState) {
		this.loginState = loginState;
	}

	public int getState() {
		return state;
	}

	public void setState(int state) {
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

	public String getAddress() {
		return address;
	}

	public void setAddress(String address) {
		this.address = address;
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

	public int getJob() {
		return job;
	}

	public void setJob(int job) {
		this.job = job;
	}

	public int getGrade() {
		return grade;
	}

	public void setGrade(int grade) {
		this.grade = grade;
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
		Login other = (Login) obj;
		if (uid != other.uid) {
			return false;
		}
		return true;
	}

}
