package com.lanlian.chatting.po;

import java.io.Serializable;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午5:58:52
 * @explain 此类用于对等级id的操作
 */
public class GradePO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 5774980304256273412L;

	private long gradeId;// 等级的id
	private String gradeName;// 等级内容

	public GradePO() {
		super();
	}

	public GradePO(long gradeId, String gradeName) {
		super();
		this.gradeId = gradeId;
		this.gradeName = gradeName;
	}

	@Override
	public String toString() {
		return "GradePOJO [gradeId=" + gradeId + ", gradeName=" + gradeName + "]";
	}

	public long getGradeId() {
		return gradeId;
	}

	public void setGradeId(long gradeId) {
		this.gradeId = gradeId;
	}

	public String getGradeName() {
		return gradeName;
	}

	public void setGradeName(String gradeName) {
		this.gradeName = gradeName;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (gradeId ^ (gradeId >>> 32));
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
		GradePO other = (GradePO) obj;
		if (gradeId != other.gradeId) {
			return false;
		}
		return true;
	}
}
