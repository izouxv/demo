/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/**
 * 
 */
package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @Title Version.java
 * @Package cn.lanlian.ccat.entity
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月28日 下午3:31:06
 * @explain 最新版本信息属性类
 */

public class VersionPO implements Serializable {

	/**
	 * 实现Serializable的序列化
	 */
	private static final long serialVersionUID = -3044609905161785934L;

	private long id;// id
	private String device;// 安装设备标识
	private String versionName;// 版本名称
	private String versionCode;// 版本号
	private String filename;// 文件名称
	private String md5;// 加密MD5
	private String description;// 版本描述
	private String description1;// 版本描述-英文
	private Timestamp time;// 上传时间
	private String length;// 长度
	private String path;// 文件路径
	private int status;// 版本状态
	private long adminPid;//
	private String adminUsername;//
	private String adminPWD;//
	private String adminRealName;//
	private String adminIDcard;//

	public VersionPO() {
		super();
	}

	public VersionPO(long id, String device, String versionName, String versionCode, String filename, String md5,
			String description, String description1, Timestamp time, String length, String path, int status,
			long adminPid, String adminUsername, String adminPWD, String adminRealName, String adminIDcard) {
		super();
		this.id = id;
		this.device = device;
		this.versionName = versionName;
		this.versionCode = versionCode;
		this.filename = filename;
		this.md5 = md5;
		this.description = description;
		this.description1 = description1;
		this.time = time;
		this.length = length;
		this.path = path;
		this.status = status;
		this.adminPid = adminPid;
		this.adminUsername = adminUsername;
		this.adminPWD = adminPWD;
		this.adminRealName = adminRealName;
		this.adminIDcard = adminIDcard;
	}

	@Override
	public String toString() {
		return "VersionPO [id=" + id + ", device=" + device + ", versionName=" + versionName + ", versionCode="
				+ versionCode + ", filename=" + filename + ", md5=" + md5 + ", description=" + description
				+ ", description1=" + description1 + ", time=" + time + ", length=" + length + ", path=" + path
				+ ", status=" + status + ", adminPid=" + adminPid + ", adminUsername=" + adminUsername + ", adminPWD="
				+ adminPWD + ", adminRealName=" + adminRealName + ", adminIDcard=" + adminIDcard + "]";
	}

	public String getDescription1() {
		return description1;
	}

	public void setDescription1(String description1) {
		this.description1 = description1;
	}

	public long getId() {
		return id;
	}

	public void setId(long id) {
		this.id = id;
	}

	public String getDevice() {
		return device;
	}

	public void setDevice(String device) {
		this.device = device;
	}

	public String getVersionName() {
		return versionName;
	}

	public void setVersionName(String versionName) {
		this.versionName = versionName;
	}

	public String getVersionCode() {
		return versionCode;
	}

	public void setVersionCode(String versionCode) {
		this.versionCode = versionCode;
	}

	public String getFilename() {
		return filename;
	}

	public void setFilename(String filename) {
		this.filename = filename;
	}

	public String getMd5() {
		return md5;
	}

	public void setMd5(String md5) {
		this.md5 = md5;
	}

	public String getDescription() {
		return description;
	}

	public void setDescription(String description) {
		this.description = description;
	}

	public Timestamp getTime() {
		return time;
	}

	public void setTime(Timestamp time) {
		this.time = time;
	}

	public String getLength() {
		return length;
	}

	public void setLength(String length) {
		this.length = length;
	}

	public String getPath() {
		return path;
	}

	public void setPath(String path) {
		this.path = path;
	}

	public int getStatus() {
		return status;
	}

	public void setStatus(int status) {
		this.status = status;
	}

	public long getAdminPid() {
		return adminPid;
	}

	public void setAdminPid(long adminPid) {
		this.adminPid = adminPid;
	}

	public String getAdminUsername() {
		return adminUsername;
	}

	public void setAdminUsername(String adminUsername) {
		this.adminUsername = adminUsername;
	}

	public String getAdminPWD() {
		return adminPWD;
	}

	public void setAdminPWD(String adminPWD) {
		this.adminPWD = adminPWD;
	}

	public String getAdminRealName() {
		return adminRealName;
	}

	public void setAdminRealName(String adminRealName) {
		this.adminRealName = adminRealName;
	}

	public String getAdminIDcard() {
		return adminIDcard;
	}

	public void setAdminIDcard(String adminIDcard) {
		this.adminIDcard = adminIDcard;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (id ^ (id >>> 32));
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
		VersionPO other = (VersionPO) obj;
		if (id != other.id) {
			return false;
		}
		return true;
	}
}
