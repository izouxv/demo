/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年11月15日 上午10:26:04
 * @explain 广告属性类
 */

public class AdverPo implements Serializable {

	
	/**
	 * 
	 */
	private static final long serialVersionUID = 7033455584876605726L;
	
	private int code;
	private int id;
	private String name;
	private String source;
	private String md5;
	private String fileUrl;
	private String advertiseUrl;
	private Long startTime;
	private Long endTime;
	
	public AdverPo() {
		super();
	}

	public AdverPo(int code, int id, String name, String source, String md5, String fileUrl, String advertiseUrl,
			Long startTime, Long endTime) {
		super();
		this.code = code;
		this.id = id;
		this.name = name;
		this.source = source;
		this.md5 = md5;
		this.fileUrl = fileUrl;
		this.advertiseUrl = advertiseUrl;
		this.startTime = startTime;
		this.endTime = endTime;
	}

	@Override
	public String toString() {
		return "AdverPo [code=" + code + ", id=" + id + ", name=" + name + ", source=" + source + ", md5=" + md5
				+ ", fileUrl=" + fileUrl + ", advertiseUrl=" + advertiseUrl + ", startTime=" + startTime + ", endTime="
				+ endTime + "]";
	}

	public int getCode() {
		return code;
	}

	public void setCode(int code) {
		this.code = code;
	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getSource() {
		return source;
	}

	public void setSource(String source) {
		this.source = source;
	}

	public String getMd5() {
		return md5;
	}

	public void setMd5(String md5) {
		this.md5 = md5;
	}

	public String getFileUrl() {
		return fileUrl;
	}

	public void setFileUrl(String fileUrl) {
		this.fileUrl = fileUrl;
	}

	public String getAdvertiseUrl() {
		return advertiseUrl;
	}

	public void setAdvertiseUrl(String advertiseUrl) {
		this.advertiseUrl = advertiseUrl;
	}

	public Long getStartTime() {
		return startTime;
	}

	public void setStartTime(Long startTime) {
		this.startTime = startTime;
	}

	public Long getEndTime() {
		return endTime;
	}

	public void setEndTime(Long endTime) {
		this.endTime = endTime;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + id;
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
		AdverPo other = (AdverPo) obj;
		if (id != other.id)
			return false;
		return true;
	}

}

