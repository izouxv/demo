/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.po.virtual;

import java.util.List;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月18日 下午5:08:23
 * @explain 虚拟信道业务类
 */

public class NearbyPo {

	public static final double MINLAT = -90;// 最小纬度

	public static final double MAXLAT = 90;// 最大纬度

	public static final double MINLNG = -180;// 最小经度

	public static final double MAXLNG = 180;// 最大经度

	private Integer tuid;// 用户id
	private String nickname;// 昵称
	private Integer avatar;// 头像id
	private Integer gender;// 性别1：男/2：女
	private Integer age;// 年龄
	private String signature;// 个性签名
	private String imei;// 用户设备标识
	private Double longitude;// 经度[-180,180]
	private Double latitude;// 纬度[-90,90]

	private Integer tgid;// 临时群id
	private Integer ownerid;// 拥有者id
	private String notice;// 群通知
	private List<Integer> tuids;// 群成员

	private List<NearbyPo> nearbyPos;// 附近的人信息

	private Double distance;// 距离
	private String geoHash;// hash字符串

	public NearbyPo() {
		super();
	}

	public NearbyPo(Double longitude, Double latitude) {
		super();
		this.longitude = longitude;
		this.latitude = latitude;
	}

	public NearbyPo(Integer tuid, String nickname, Integer avatar, Integer gender, Integer age, String signature,
			String imei, Double longitude, Double latitude, Integer tgid, Integer ownerid, String notice,
			List<Integer> tuids, List<NearbyPo> nearbyPos, Double distance, String geoHash) {
		super();
		this.tuid = tuid;
		this.nickname = nickname;
		this.avatar = avatar;
		this.gender = gender;
		this.age = age;
		this.signature = signature;
		this.imei = imei;
		this.longitude = longitude;
		this.latitude = latitude;
		this.tgid = tgid;
		this.ownerid = ownerid;
		this.notice = notice;
		this.tuids = tuids;
		this.nearbyPos = nearbyPos;
		this.distance = distance;
		this.geoHash = geoHash;
	}

	@Override
	public String toString() {
		return "NearbyPo [tuid=" + tuid + ", nickname=" + nickname + ", avatar=" + avatar + ", gender=" + gender
				+ ", age=" + age + ", signature=" + signature + ", imei=" + imei + ", longitude=" + longitude
				+ ", latitude=" + latitude + ", tgid=" + tgid + ", ownerid=" + ownerid + ", notice=" + notice
				+ ", tuids=" + tuids + ", nearbyPos=" + nearbyPos + ", distance=" + distance + ", geoHash=" + geoHash
				+ "]";
	}

	public Integer getOwnerid() {
		return ownerid;
	}

	public void setOwnerid(Integer ownerid) {
		this.ownerid = ownerid;
	}

	public Integer getTgid() {
		return tgid;
	}

	public void setTgid(Integer tgid) {
		this.tgid = tgid;
	}

	public String getNotice() {
		return notice;
	}

	public void setNotice(String notice) {
		this.notice = notice;
	}

	public List<Integer> getTuids() {
		return tuids;
	}

	public void setTuids(List<Integer> tuids) {
		this.tuids = tuids;
	}

	public List<NearbyPo> getNearbyPos() {
		return nearbyPos;
	}

	public void setNearbyPos(List<NearbyPo> nearbyPos) {
		this.nearbyPos = nearbyPos;
	}

	public Integer getTuid() {
		return tuid;
	}

	public void setTuid(Integer tuid) {
		this.tuid = tuid;
	}

	public String getNickname() {
		return nickname;
	}

	public void setNickname(String nickname) {
		this.nickname = nickname;
	}

	public Integer getAvatar() {
		return avatar;
	}

	public void setAvatar(Integer avatar) {
		this.avatar = avatar;
	}

	public Integer getGender() {
		return gender;
	}

	public void setGender(Integer gender) {
		this.gender = gender;
	}

	public Integer getAge() {
		return age;
	}

	public void setAge(Integer age) {
		this.age = age;
	}

	public String getSignature() {
		return signature;
	}

	public void setSignature(String signature) {
		this.signature = signature;
	}

	public String getImei() {
		return imei;
	}

	public void setImei(String imei) {
		this.imei = imei;
	}

	public Double getLongitude() {
		return longitude;
	}

	public void setLongitude(Double longitude) {
		this.longitude = longitude;
	}

	public Double getLatitude() {
		return latitude;
	}

	public void setLatitude(Double latitude) {
		this.latitude = latitude;
	}

	public Double getDistance() {
		return distance;
	}

	public void setDistance(Double distance) {
		this.distance = distance;
	}

	public String getGeoHash() {
		return geoHash;
	}

	public void setGeoHash(String geoHash) {
		this.geoHash = geoHash;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((tuid == null) ? 0 : tuid.hashCode());
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
		NearbyPo other = (NearbyPo) obj;
		if (tuid == null) {
			if (other.tuid != null) {
				return false;
			}
		} else if (!tuid.equals(other.tuid)) {
			return false;
		}
		return true;
	}

}
