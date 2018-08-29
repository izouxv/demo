/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package Demo;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.text.DateFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.TimeZone;

import javax.xml.transform.Transformer;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.alibaba.fastjson.serializer.SerializerFeature;

import net.sf.cglib.core.CollectionUtils;

/**
 * @Title TestMain.java
 * @Package Demo
 * @author 王东阳
 * @version V1.0
 * @date 2017年5月23日 上午9:18:40
 * @explain
 */

public class TestMain {

	// public static void main(String[] args) throws ParseException {
	// Date date = new Date(1412151214000L);
	// SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
	// sdf.parse("1900-01-01 00:00:00");//
	// System.out.println(sdf.format(date)+ "," +sdf.parse("1900-01-01
	// 00:00:00").getTime());
	// System.out.println(System.currentTimeMillis());
	// }

	public static void getUTCTimeStr() throws ParseException {
		// 1、取得本地时间：
		Calendar cal = Calendar.getInstance();
		// 2、取得时间偏移量：
		int zoneOffset = cal.get(java.util.Calendar.ZONE_OFFSET);
		// 3、取得夏令时差：
		int dstOffset = cal.get(java.util.Calendar.DST_OFFSET);
		// 4、从本地时间里扣除这些差量，即可以取得UTC时间：
		cal.add(java.util.Calendar.MILLISECOND, -(zoneOffset + dstOffset));
		System.out.println("aaa:"+cal.getTime());
	}

	public static void main(String[] args) throws UnsupportedEncodingException, ParseException {
		getUTCTimeStr();
		/**
		 * QuoteFieldNames———-输出key时是否使用双引号,默认为true
		 * WriteMapNullValue——–是否输出值为null的字段,默认为false
		 * WriteNullNumberAsZero—-数值字段如果为null,输出为0,而非null
		 * WriteNullListAsEmpty—–List字段如果为null,输出为[],而非null
		 * WriteNullStringAsEmpty—字符类型字段如果为null,输出为"",而非null
		 * WriteNullBooleanAsFalse–Boolean字段如果为null,输出为false,而非null
		 */

		/**
		 * Map<String, Object> jsonMap = new HashMap<>(); jsonMap.put("a", 1);
		 * jsonMap.put("b", ""); jsonMap.put("c", null); jsonMap.put("d", "json");
		 * 
		 * String str = JSONObject.toJSONString(jsonMap,
		 * SerializerFeature.WriteMapNullValue); // ①忽略null输出 System.out.println(str);
		 * 
		 * String str2 = JSONObject.toJSONString(jsonMap,
		 * SerializerFeature.WriteMapNullValue); // ② System.out.println(str2);
		 * 
		 * String json =
		 * "{\"fail\":null,\"updateTimestamp\":\"1484096131863\",\"productName\":\"json测试\"}";
		 * // ③忽略null输出 System.out.println(JSON.parse(json)); // ④
		 * System.out.println(JSONObject.toJSON(json));
		 * 
		 * //测试CollectionUtils System.out.println(System.currentTimeMillis());
		 * Set<String> set = new HashSet<>(); Set<Integer> set2 = new HashSet<>();
		 * Integer gid = 1; for (int i = 0; i < 16; i++) { set2.add(gid); gid++; }
		 * CollectionUtils.collect(set2, new Transformer() { public String
		 * transform(Object input) { return input.toString(); } }, set);
		 * System.out.println(set); System.out.println(System.currentTimeMillis());
		 * 
		 * //测试url编码后的长度 String aa = "%E6%B5%8B%E8%AF%952123"; aa =
		 * URLDecoder.decode(aa, "utf-8"); System.out.println(aa + "," +
		 * aa.getBytes().length); int[][] arrays = { // {1,2,3,4}, // {2,3,4,5}, //
		 * {3,4,5,6}, // {4,5,6,7}}; {3,4,7}, {4,7,9}, {7,12,14}, {10,13,17}};
		 * System.out.println(FindArray2(13, arrays));
		 */
	}

	/**
	 * 运行时间：202ms 占用内存：17308k
	 */

	/**
	 * 依次无规律递增
	 * 
	 * @param target
	 * @param array
	 * @return
	 */
	public static boolean FindArray2(int target, int[][] array) {
		int row = array.length;
		int col = array[0].length;
		if (row == 0 || col == 0 || array[0][0] > target || array[row - 1][col - 1] < target) {
			return false;
		}
		for (int i = 0; i < row; i++) {
			if (array[i][col - 1] >= target) {
				for (int j = 0; j < col; j++) {
					if (target == array[i][j]) {
						return true;
					}
				}
			}

		}
		return false;
	}

	/**
	 * 依次规律递增
	 * 
	 * @param target
	 * @param array
	 * @return
	 */
	public static boolean FindArray1(int target, int[][] array) {
		int row1V = array[0][1] - array[0][0];
		int row2V = array[1][0] - array[1][1];
		int width = array[0].length;
		int height = array.length;
		if ((array[height - 1][0] + row1V * width) >= target) {
			if ((target - array[0][0]) % row1V == 0 || (target - array[0][0]) % row2V == 0) {
				return true;
			}
		}
		return false;
	}

}