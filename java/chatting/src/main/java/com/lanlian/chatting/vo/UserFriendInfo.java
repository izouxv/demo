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
 * @Title UserFriendInfo.java
 * @Package cn.lanlian.ccat.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月16日 下午4:42:49
 * @explain 
 */

public class UserFriendInfo implements Serializable {
	
	private static final long serialVersionUID = 3281662017277167640L;
	
	private int uid;//用户的uid
	private String note;//备注
	private String nickname;//昵称
	private int gender;//性别1：男/2：女
	private int grade;//等级ID
	private long birthday;//生日
	private int avatar;//头像id（以后会是头像地址）
	private String email;//邮箱
	private String phone;//手机号
	private String signature;//个性签名
	private int creditValues;//信用点
	private String province;//省
	private String city;//市
	private String address;//详细地址
	private int job;//职业ID，对应职位表
	private String isCertification;//是否实名认证
	private int status;//用户状态0：正常；1：注册&没激活；2：注销；3：已删除；4：禁用；
	
	public UserFriendInfo() {
		super();
	}

	public UserFriendInfo(int uid, String note, String nickname, int gender, int grade, long birthday, int avatar,
			String email, String phone, String signature, int creditValues, String province, String city,
			String address, int job, String isCertification, int status) {
		super();
		this.uid = uid;
		this.note = note;
		this.nickname = nickname;
		this.gender = gender;
		this.grade = grade;
		this.birthday = birthday;
		this.avatar = avatar;
		this.email = email;
		this.phone = phone;
		this.signature = signature;
		this.creditValues = creditValues;
		this.province = province;
		this.city = city;
		this.address = address;
		this.job = job;
		this.isCertification = isCertification;
		this.status = status;
	}

	@Override
	public String toString() {
		return "UserFriendInfo [uid=" + uid + ", note=" + note + ", nickname=" + nickname + ", gender=" + gender
				+ ", grade=" + grade + ", birthday=" + birthday + ", avatar=" + avatar + ", email=" + email
				+ ", phone=" + phone + ", signature=" + signature + ", creditValues=" + creditValues + ", province="
				+ province + ", city=" + city + ", address=" + address + ", job=" + job + ", isCertification="
				+ isCertification + ", status=" + status + "]";
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}

	public String getNote() {
		return note;
	}

	public void setNote(String note) {
		this.note = note;
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

	public int getGrade() {
		return grade;
	}

	public void setGrade(int grade) {
		this.grade = grade;
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

	public String getSignature() {
		return signature;
	}

	public void setSignature(String signature) {
		this.signature = signature;
	}

	public int getCreditValues() {
		return creditValues;
	}

	public void setCreditValues(int creditValues) {
		this.creditValues = creditValues;
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

	public int getJob() {
		return job;
	}

	public void setJob(int job) {
		this.job = job;
	}

	public String getIsCertification() {
		return isCertification;
	}

	public void setIsCertification(String isCertification) {
		this.isCertification = isCertification;
	}

	public int getStatus() {
		return status;
	}

	public void setStatus(int status) {
		this.status = status;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (uid ^ (uid >>> 32));
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
		UserFriendInfo other = (UserFriendInfo) obj;
		if (uid != other.uid) {
			return false;}
		return true;
	}
	
}

