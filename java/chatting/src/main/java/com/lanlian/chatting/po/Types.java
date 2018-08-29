/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package com.lanlian.chatting.po;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年2月5日 下午7:47:06
 * @wangdyq
 * @explain 
 */

public enum Types {
	ACCOUNT("account",1),
	PASSWORD("password",2),
	SESSION("session",3);
	private String name;
	private int index;
	private Types(String name, int index) {
		this.index = index;
	}
	public String getName() {
		return this.name;
	}
	public int getIndex() {
		return this.index;
	}

	public static boolean contains(int value) {
		for (Types key : Types.values()) {
			if (key.getIndex() == value) {
				return true;
			}
		}
		return false;
	}

	@Override
	public String toString() {
		return String.valueOf(this.index);
	}
	
	public static void main(String[] args) {
		System.out.println(Types.ACCOUNT.name());
	}
}

