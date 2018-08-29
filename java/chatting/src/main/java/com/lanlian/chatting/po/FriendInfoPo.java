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
 * @Title FriendInfoPo.java
 * @Package cn.lanlian.chatting.po
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月16日 下午4:42:49
 * @explain 好友信息
 */

public class FriendInfoPo implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1025378902686507791L;

	private int uid;// 用户的id
	private String note;// 备注
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
	private int isCertification;// 是否实名认证
	private int creditValues;// 信用点
	private int userPoint;// 积分
	private int userJobId;// 职业ID，对应职位表
	private int userGradeId;// 等级ID

	public FriendInfoPo() {
		super();
	}

	public FriendInfoPo(int uid, String note, String phone, String email, String nickname, int gender, long birthday,
			int avatar, String signature, String province, String city, String userAddress, int isCertification,
			int creditValues, int userPoint, int userJobId, int userGradeId) {
		super();
		this.uid = uid;
		this.note = note;
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
	}

	@Override
	public String toString() {
		return "FriendInfoPo [uid=" + uid + ", note=" + note + ", phone=" + phone + ", email=" + email + ", nickname="
				+ nickname + ", gender=" + gender + ", birthday=" + birthday + ", avatar=" + avatar + ", signature="
				+ signature + ", province=" + province + ", city=" + city + ", userAddress=" + userAddress
				+ ", isCertification=" + isCertification + ", creditValues=" + creditValues + ", userPoint=" + userPoint
				+ ", userJobId=" + userJobId + ", userGradeId=" + userGradeId + "]";
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

	public String getUserAddress() {
		return userAddress;
	}

	public void setUserAddress(String userAddress) {
		this.userAddress = userAddress;
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

	public int getUserPoint() {
		return userPoint;
	}

	public void setUserPoint(int userPoint) {
		this.userPoint = userPoint;
	}

	public int getUserJobId() {
		return userJobId;
	}

	public void setUserJobId(int userJobId) {
		this.userJobId = userJobId;
	}

	public int getUserGradeId() {
		return userGradeId;
	}

	public void setUserGradeId(int userGradeId) {
		this.userGradeId = userGradeId;
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
		FriendInfoPo other = (FriendInfoPo) obj;
		if (uid != other.uid) {
			return false;
		}
		return true;
	}

}
