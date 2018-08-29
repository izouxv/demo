package Demo;

import java.util.Random;

/**
 * @Title DadaGroupController.java
 * @Package com.lanlian.chatting.controller
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午8:27:04
 * @explain 开启群信息实时上报
 * 邀请码生成器，算法原理：
 * 1) 获取id: 1127738
 * 2) 使用自定义进制转为：gpm6
 * 3) 转为字符串，并在后面加'o'字符：gpm6o
 * 4）在后面随机产生若干个随机数字字符：gpm6o7
 * 转为自定义进制后就不会出现o这个字符，然后在后面加个'o'，这样就确定唯一性。
 * 最后在后面产生一些随机字符进行补全
 */
public class InviteCode {

	/** 自定义进制(0,1没有加入,容易与o,l混淆) */
	private static final char[] chars = new char[] { 'q', 'w', 'e', '8', 'a', 's', '2', 'd', 'z', 'x', '9', 'c', '7', 'p',
			'5', 'i', 'k', '3', 'm', 'j', 'u', 'f', 'r', '4', 'v', 'y', 'l', 't', 'n', '6', 'b', 'g', 'h' };

	
	/** (不能与自定义进制有重复) */
	private static final char b = 'o';

	/** 进制长度 */
	private static final int binLen = chars.length;

	/** 序列最小长度 */
	private static final int s = 4;

	/**
	 * 根据ID生成六位随机码
	 * 
	 * @param id
	 *            ID
	 * @return 随机码
	 */
	public static String getRandomCode(Integer id) {
		char[] buf = new char[32];
		int charPos = 32;
		while ((id / binLen) > 0) {
			int ind = id % binLen;
			// System.out.println(num + "-->" + ind);
			buf[--charPos] = chars[ind];
			id /= binLen;
		}
		buf[--charPos] = chars[(int) (id % binLen)];
		// System.out.println(num + "-->" + num % binLen);
		String str = new String(buf, charPos, (32 - charPos));
		// 不够长度的自动随机补全
		if (str.length() < s) {
			StringBuilder sb = new StringBuilder();
			sb.append(b);
			Random rnd = new Random();
			for (int i = 1; i < s - str.length(); i++) {
				sb.append(chars[rnd.nextInt(binLen)]);
			}
			str += sb.toString();
		}
		return str;
	}

	/**
	 * code转id
	 * @param code
	 * @return
	 */
	public static int codeToId(String code) {
		char chs[] = code.toCharArray();
		int res = 0;
		for (int i = 0; i < chs.length; i++) {
			int ind = 0;
			for (int j = 0; j < binLen; j++) {
				if (chs[i] == chars[j]) {
					ind = j;
					break;
				}
			}
			if (chs[i] == b) {
				break;
			}
			if (i > 0) {
				res = res * binLen + ind;
			} else {
				res = ind;
			}
			System.out.println(ind + "-->" + res);
		}
		return res;
	}

	public static void main(String[] args) throws Exception {
		for (int i = 0; i < 100; i++) {
			if (i%10 == 0) {
				System.out.println();
			}
			System.out.print(getRandomCode(100000000 + i) + ",");
		}
	}


}