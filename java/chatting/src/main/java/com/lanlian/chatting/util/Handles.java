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
package com.lanlian.chatting.util;

/**
 * @Title Handle.java
 * @Package com.lanlian.chatting.util
 * @author 王东阳
 * @version V1.0
 * @date 2017年5月8日 上午11:15:21
 * @explain
 */

@SuppressWarnings("unused")
public enum Handles {

	REGISTER(11), FINDPWD(18), ZERO(0), ONE(1), TWO(2), THREE(3), FOUR(4), FIVE(5);

	private int value = 0;

	private Handles(int value) {
		this.value = value;
	}

	public Handles valueOf(int value) {
		switch (value) {
		case 11:
			return REGISTER;
		default:
			return null;
		}
	}

}

enum MyEnum {
	Normal, SECOND_ELEMENT;
	public static void main(String[] args) {
		
		
	}
}

