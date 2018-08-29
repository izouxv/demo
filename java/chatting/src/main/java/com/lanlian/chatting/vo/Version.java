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
package com.lanlian.chatting.vo;

import java.io.Serializable;

/** 
 * @Title Version.java
 * @Package cn.lanlian.ccat.entity
 * @author 王东阳
 * @version V1.0
 * @date 2017年3月28日 下午3:31:06
 * @explain 最新版本信息属性类
 */

public class Version implements Serializable {
	
	/**
	 * 实现Serializable的序列化
	 */
	private static final long serialVersionUID = -3044609905161785934L;
	
	private String device;//安装设备标识
	private String versionName;//版本名称
	private String versionCode;//版本号
	private String description;//版本描述
	private String description1;//版本描述-英文
	private String filename;//文件名称
	private String md5;//加密MD5
	private String time;//上传时间 
	private String length;//长度
	private int uid;//用户id
	
	public Version() {
	}

	public Version(String device, String versionName, String versionCode, String description, String description1,
			String filename, String md5, String time, String length, int uid) {
		super();
		this.device = device;
		this.versionName = versionName;
		this.versionCode = versionCode;
		this.description = description;
		this.description1 = description1;
		this.filename = filename;
		this.md5 = md5;
		this.time = time;
		this.length = length;
		this.uid = uid;
	}

	@Override
	public String toString() {
		return "Version [device=" + device + ", versionName=" + versionName + ", versionCode=" + versionCode
				+ ", description=" + description + ", description1=" + description1 + ", filename=" + filename
				+ ", md5=" + md5 + ", time=" + time + ", length=" + length + ", uid=" + uid + "]";
	}

	public String getDescription1() {
		return description1;
	}

	public void setDescription1(String description1) {
		this.description1 = description1;
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

	public String getDescription() {
		return description;
	}

	public void setDescription(String description) {
		this.description = description;
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

	public String getTime() {
		return time;
	}

	public void setTime(String time) {
		this.time = time;
	}

	public String getLength() {
		return length;
	}

	public void setLength(String length) {
		this.length = length;
	}

	public int getUid() {
		return uid;
	}

	public void setUid(int uid) {
		this.uid = uid;
	}
	
	

}

