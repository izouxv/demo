package com.lanlian.chatting.po;

import java.io.Serializable;
import java.sql.Timestamp;

/**
 * @author wdyqxx
 * @version 2017年1月2日 下午5:58:59
 * @explain 此类用于对等级与积分的操作
 */
public class GradePointPO implements Serializable {

	/**
	 * 
	 */
	private static final long serialVersionUID = -1385991373391512754L;

	private long gradePointId;// 等级与积分表的id
	private String point;// 积分
	private String grade;// 等级
	private Timestamp createTime;// 创建时间
	private Timestamp modifyTime;// 最近修改时间

	public GradePointPO() {
		super();
	}

	public GradePointPO(long gradePointId, String point, String grade, Timestamp createTime, Timestamp modifyTime) {
		super();
		this.gradePointId = gradePointId;
		this.point = point;
		this.grade = grade;
		this.createTime = createTime;
		this.modifyTime = modifyTime;
	}

	@Override
	public String toString() {
		return "GradePointPOJO [gradePointId=" + gradePointId + ", point=" + point + ", grade=" + grade
				+ ", createTime=" + createTime + ", modifyTime=" + modifyTime + "]";
	}

	public long getGradePointId() {
		return gradePointId;
	}

	public void setGradePointId(long gradePointId) {
		this.gradePointId = gradePointId;
	}

	public String getPoint() {
		return point;
	}

	public void setPoint(String point) {
		this.point = point;
	}

	public String getGrade() {
		return grade;
	}

	public void setGrade(String grade) {
		this.grade = grade;
	}

	public Timestamp getCreateTime() {
		return createTime;
	}

	public void setCreateTime(Timestamp createTime) {
		this.createTime = createTime;
	}

	public Timestamp getModifyTime() {
		return modifyTime;
	}

	public void setModifyTime(Timestamp modifyTime) {
		this.modifyTime = modifyTime;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + (int) (gradePointId ^ (gradePointId >>> 32));
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
		GradePointPO other = (GradePointPO) obj;
		if (gradePointId != other.gradePointId) {
			return false;
		}
		return true;
	}

}
