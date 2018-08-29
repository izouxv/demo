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
 * @Title Version2.java
 * @Package cn.lanlian.chatting.vo
 * @author 王东阳
 * @version V1.1
 * @date 2018年3月12日 下午3:31:06
 * @explain 最新版本信息属性类
 */

public class Version2 implements Serializable {
	
	/**
	 * 实现Serializable的序列化
	 */
	private static final long serialVersionUID = -5818129803479141119L;

	private String device;//设备标识
	private String versionName;//版本名称
	private String versionCode;//版本号
	private String DescriptionCn;//版本描述
	private String DescriptionEn;//版本描述-英文
	private String filename;//文件名称
	private String md5;//文件MD5
	private String time;//上传时间
	private Long length;//长度
	private String url;//文件地址
	
	public Version2() {
	}

	public Version2(String device, String versionName, String versionCode, String descriptionCn, String descriptionEn,
			String filename, String md5, String time, Long length, String url) {
		super();
		this.device = device;
		this.versionName = versionName;
		this.versionCode = versionCode;
		DescriptionCn = descriptionCn;
		DescriptionEn = descriptionEn;
		this.filename = filename;
		this.md5 = md5;
		this.time = time;
		this.length = length;
		this.url = url;
	}

	@Override
	public String toString() {
		return "Version2 [device=" + device + ", versionName=" + versionName + ", versionCode=" + versionCode
				+ ", DescriptionCn=" + DescriptionCn + ", DescriptionEn=" + DescriptionEn + ", filename=" + filename
				+ ", md5=" + md5 + ", time=" + time + ", length=" + length + ", url=" + url + "]";
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

	public String getDescriptionCn() {
		return DescriptionCn;
	}

	public void setDescriptionCn(String descriptionCn) {
		DescriptionCn = descriptionCn;
	}

	public String getDescriptionEn() {
		return DescriptionEn;
	}

	public void setDescriptionEn(String descriptionEn) {
		DescriptionEn = descriptionEn;
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

	public Long getLength() {
		return length;
	}

	public void setLength(Long length) {
		this.length = length;
	}

	public String getUrl() {
		return url;
	}

	public void setUrl(String url) {
		this.url = url;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((md5 == null) ? 0 : md5.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Version2 other = (Version2) obj;
		if (md5 == null) {
			if (other.md5 != null)
				return false;
		} else if (!md5.equals(other.md5))
			return false;
		return true;
	}
	
}

