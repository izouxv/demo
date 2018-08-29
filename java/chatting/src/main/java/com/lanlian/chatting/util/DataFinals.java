package com.lanlian.chatting.util;

/**
 * 
 * @Title DataFinals.java
 * @Package com.lanlian.chatting.util
 * @author wangdyqxx
 * @version V1.0
 * @date 2017年3月22日 上午9:20:21
 * @explain: 全局常量类；
 */
public class DataFinals {

	/**
	 * 嗒嗒账号pid
	 */
	public static final int DADA = 1000000;// 嗒嗒通知账号

	public static final String PUBLIC = "邀请你";

	public static final String ADD_GROUP = "加入群聊";

	public static final String REMOVE_GROUP = "你已被管理员移出群";

	public static final String DISSPLVE_GROUP = "已被群主解散";

	/**
	 * 整型数字
	 */
	public enum number {
		zero(0), one(1), two(2), three(3), four(4), five(5), six(6);
		private int num;

		private number(int num) {
			this.num = num;
		}

		public int getValue() {
			return this.num;
		}

		public static boolean contains(int num) {
			for (number nums : number.values()) {
				if (nums.getValue() == num) {
					return true;
				}
			}
			return false;
		}

		@Override
		public String toString() {
			return String.valueOf(this.num);
		}
	}

	/**
	 * 性别:0.不变;1.男;2.女;3.无;
	 */
	public enum Gender {
		constant(0), man(1), woman(2), nothing(3);
		private int num;

		private Gender(int num) {
			this.num = num;
		}

		public int getGender() {
			return this.num;
		}

		@Override
		public String toString() {
			return String.valueOf(this.num);
		}

		public static boolean contains(int num) {
			for (number nums : number.values()) {
				if (nums.getValue() == num) {
					return true;
				}
			}
			return false;
		}
	}

	/**
	 * 接口
	 */
	public enum Interfaces {
		register("register"), findPwd("findpwd");
		private String inter;

		private Interfaces(String interfaces) {
			this.inter = interfaces;
		}

		public String getValue() {
			return this.inter;
		}

		@Override
		public final String toString() {
			return this.inter;
		}

		public static boolean contains(String inters) {
			for (Interfaces inter : Interfaces.values()) {
				if (inter.getValue() == inters) {
					return true;
				}
			}
			return false;
		}
	}

	/**
	 * 模块接口序列号:注册
	 */
	public static final String REGISTER = "register";

	/**
	 * 模块接口序列号:找回密码
	 */
	public static final String RESETPWD = "resetpwd";

	/**
	 * 操作
	 */
	public static enum Actions {
		find("find"), get("get");
		private String action;
		private Actions(String action) {
			this.action = action;
		}
		public String getActions() {
			return action;
		}

		/**
		 * 根据位置判断
		 * @param le
		 * @param action
		 * @return
		 */
		public static boolean equals(int le, String action) {
			if (Actions.values()[le].toString().equals(action)) {
				return true;
			}
			return false;
		}

		/**
		 * 判断是否包含于
		 * 
		 * @param action
		 * @return
		 */
		public static boolean contains(String action) {
			for (Actions param : Actions.values()) {
				if (param.getActions().equals(action)) {
					return true;
				}
			}
			return false;
		}
	}

	/**
	 * 设备名称
	 */
	public static enum Devices {
		dacat("dacat"), dacat_demo("dacat-Demo"), garcat("garcat"), tomcatI("tomcatI"), tomcatII("tomcatII");
		private String device;

		private Devices(String device) {
			this.device = device;
		}

		public String getDevice() {
			return device;
		}

		public static boolean contains(String device) {
			for (Devices param : Devices.values()) {
				if (param.getDevice().equals(device)) {
					return true;
				}
			}
			return false;
		}
	}

}

enum Interfaces1 {
	register, findPwd, REGISTER;
}
