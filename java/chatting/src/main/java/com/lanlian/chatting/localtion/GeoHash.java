/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.localtion;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

import com.lanlian.chatting.po.virtual.NearbyPo;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月17日 上午11:34:16
 * @explain GeoHash实现经纬度的转化
 */

public class GeoHash {

	private NearbyPo localtionPo;

	/**
	 * 1:5009.4km; 2:1252.3km; 3:156.5km; 4:39.1km; 
	 * 5:4.9km; 6:1.2km; 7:152.4m; 8:39.2m;
	 */
	private int hashLength = 8; // 经纬度转化为geohash长度
	private int latLength = 20; // 纬度转化为二进制长度
	private int lngLength = 20; // 经度转化为二进制长度

	private double minLat;// 每格纬度的单位大小
	private double minLng;// 每个经度的倒下
	private static final char[] CHARS = { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'b', 'c', 'd', 'e', 'f',
			'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z' };

	private static HashMap<Character, Integer> CHARSMAP;

	static {
		CHARSMAP = new HashMap<Character, Integer>();
		for (int i = 0; i < CHARS.length; i++) {
			CHARSMAP.put(CHARS[i], i);
		}
	}

	public GeoHash(double lng, double lat) {
		localtionPo = new NearbyPo(lng, lat);
		setMinLatLng();
	}

	public int gethashLength() {
		return hashLength;
	}

	/**
	 * @Author:lulei
	 * @Description: 设置经纬度的最小单位
	 */
	private void setMinLatLng() {
		minLat = NearbyPo.MAXLAT - NearbyPo.MINLAT;
		for (int i = 0; i < latLength; i++) {
			minLat /= 2.0;
		}
		minLng = NearbyPo.MAXLNG - NearbyPo.MINLNG;
		for (int i = 0; i < lngLength; i++) {
			minLng /= 2.0;
		}
	}

	/**
	 * @param length
	 * @return
	 * @Author:lulei
	 * @Description: 设置经纬度转化为geohash长度
	 */
	public boolean sethashLength(int length) {
		if (length < 1) {
			return false;
		}
		hashLength = length;
		latLength = (length * 5) / 2;
		if (length % 2 == 0) {
			lngLength = latLength;
		} else {
			lngLength = latLength + 1;
		}
		setMinLatLng();
		return true;
	}

	/**
	 * 将经纬度转换为Base32
	 * 
	 * @param lat
	 * @param lng
	 * @return
	 * @Author:lulei
	 * @Description: 获取经纬度的base32字符串
	 */
	private String getGeoHashBase32(double longitude, double latitude) {
		boolean[] bools = getGeoBinary(longitude, latitude);
		if (bools == null) {
			return null;
		}
		StringBuffer sb = new StringBuffer();
		for (int i = 0; i < bools.length; i = i + 5) {
			boolean[] base32 = new boolean[5];
			for (int j = 0; j < 5; j++) {
				base32[j] = bools[i + j];
			}
			char cha = getBase32Char(base32);
			if (' ' == cha) {
				return null;
			}
			sb.append(cha);
		}
		return sb.toString();
	}

	/**
	 * @param base32
	 * @return
	 * @Author:lulei
	 * @Description: 将五位二进制转化为base32
	 */
	private char getBase32Char(boolean[] base32) {
		if (base32 == null || base32.length != 5) {
			return ' ';
		}
		int num = 0;
		for (boolean bool : base32) {
			num <<= 1;
			if (bool) {
				num += 1;
			}
		}
		return CHARS[num % CHARS.length];
	}

	/**
	 * @param i
	 * @return
	 * @Author:lulei
	 * @Description: 将数字转化为二进制字符串
	 */
	private String getBase32BinaryString(int i) {
		if (i < 0 || i > 31) {
			return null;
		}
		String str = Integer.toBinaryString(i + 32);
		return str.substring(1);
	}

	/**
	 * @param geoHash
	 * @return
	 * @Author:lulei
	 * @Description: 将geoHash转化为二进制字符串
	 */
	private String getGeoHashBinaryString(String geoHash) {
		if (geoHash == null || "".equals(geoHash)) {
			return null;
		}
		StringBuffer sb = new StringBuffer();
		for (int i = 0; i < geoHash.length(); i++) {
			char c = geoHash.charAt(i);
			if (CHARSMAP.containsKey(c)) {
				String cStr = getBase32BinaryString(CHARSMAP.get(c));
				if (cStr != null) {
					sb.append(cStr);
				}
			}
		}
		return sb.toString();
	}

	/**
	 * @param geoHash
	 * @return
	 * @Author:lulei
	 * @Description: 返回geoHash 对应的坐标
	 */
	public NearbyPo getLocation(String geoHash) {
		String geoHashBinaryStr = getGeoHashBinaryString(geoHash);
		if (geoHashBinaryStr == null) {
			return null;
		}
		StringBuffer lat = new StringBuffer();
		StringBuffer lng = new StringBuffer();
		for (int i = 0; i < geoHashBinaryStr.length(); i++) {
			if (i % 2 != 0) {
				lat.append(geoHashBinaryStr.charAt(i));
			} else {
				lng.append(geoHashBinaryStr.charAt(i));
			}
		}
		double latValue = getGeoHashMid(lat.toString(), NearbyPo.MINLAT, NearbyPo.MAXLAT);
		double lngValue = getGeoHashMid(lng.toString(), NearbyPo.MINLNG, NearbyPo.MAXLNG);
		NearbyPo location = new NearbyPo(lngValue, latValue);
		location.setGeoHash(geoHash);
		return location;
	}

	/**
	 * @param binaryStr
	 * @param min
	 * @param max
	 * @return
	 * @Author:lulei
	 * @Description: 返回二进制对应的中间值
	 */
	private double getGeoHashMid(String binaryStr, double min, double max) {
		if (binaryStr == null || binaryStr.length() < 1) {
			return (min + max) / 2.0;
		}
		if (binaryStr.charAt(0) == '1') {
			return getGeoHashMid(binaryStr.substring(1), (min + max) / 2.0, max);
		} else {
			return getGeoHashMid(binaryStr.substring(1), min, (min + max) / 2.0);
		}
	}

	/**
	 * 将经纬度转换为用true/false为代表的二进制boolean数组
	 * 
	 * @param lat
	 * @param lng
	 * @return
	 * @Author:lulei
	 * @Description: 获取坐标的geo二进制字符串
	 */
	private boolean[] getGeoBinary(double lat, double lng) {
		boolean[] latArray = getHashArray(lat, NearbyPo.MINLAT, NearbyPo.MAXLAT, latLength);
		boolean[] lngArray = getHashArray(lng, NearbyPo.MINLNG, NearbyPo.MAXLNG, lngLength);
		return merge(latArray, lngArray);
	}

	/**
	 * @param latArray
	 * @param lngArray
	 * @return
	 * @Author:lulei
	 * @Description: 合并经纬度二进制
	 */
	private boolean[] merge(boolean[] latArray, boolean[] lngArray) {
		if (latArray == null || lngArray == null) {
			return null;
		}
		boolean[] result = new boolean[lngArray.length + latArray.length];
		Arrays.fill(result, false);
		for (int i = 0; i < lngArray.length; i++) {
			result[2 * i] = lngArray[i];
		}
		for (int i = 0; i < latArray.length; i++) {
			result[2 * i + 1] = latArray[i];
		}
		return result;
	}

	/**
	 * @param value
	 * @param min
	 * @param max
	 * @return
	 * @Author:lulei
	 * @Description: 将数字转化为geohash二进制字符串
	 */
	private boolean[] getHashArray(double value, double min, double max, int length) {
		if (value < min || value > max) {
			return null;
		}
		if (length < 1) {
			return null;
		}
		boolean[] result = new boolean[length];
		for (int i = 0; i < length; i++) {
			double mid = (min + max) / 2.0;
			if (value > mid) {
				result[i] = true;
				min = mid;
			} else {
				result[i] = false;
				max = mid;
			}
		}
		return result;
	}

	/**
	 * 求所在坐标点及周围点组成的九个
	 * 
	 * @return
	 */
	public List<String> getGeoHashBase32For9() {
		double leftLat = localtionPo.getLatitude() - minLat;
		double rightLat = localtionPo.getLatitude() + minLat;
		double upLng = localtionPo.getLongitude() - minLng;
		double downLng = localtionPo.getLongitude() + minLng;
		List<String> base32For9 = new ArrayList<String>();
		// 左侧从上到下 3个
		String leftUp = getGeoHashBase32(leftLat, upLng);
		if (!(leftUp == null || "".equals(leftUp))) {
			base32For9.add(leftUp);
		}
		String leftMid = getGeoHashBase32(leftLat, localtionPo.getLongitude());
		if (!(leftMid == null || "".equals(leftMid))) {
			base32For9.add(leftMid);
		}
		String leftDown = getGeoHashBase32(leftLat, downLng);
		if (!(leftDown == null || "".equals(leftDown))) {
			base32For9.add(leftDown);
		}
		// 中间从上到下 3个
		String midUp = getGeoHashBase32(localtionPo.getLatitude(), upLng);
		if (!(midUp == null || "".equals(midUp))) {
			base32For9.add(midUp);
		}
		String midMid = getGeoHashBase32(localtionPo.getLatitude(), localtionPo.getLongitude());
		if (!(midMid == null || "".equals(midMid))) {
			base32For9.add(midMid);
		}
		String midDown = getGeoHashBase32(localtionPo.getLatitude(), downLng);
		if (!(midDown == null || "".equals(midDown))) {
			base32For9.add(midDown);
		}
		// 右侧从上到下 3个
		String rightUp = getGeoHashBase32(rightLat, upLng);
		if (!(rightUp == null || "".equals(rightUp))) {
			base32For9.add(rightUp);
		}
		String rightMid = getGeoHashBase32(rightLat, localtionPo.getLongitude());
		if (!(rightMid == null || "".equals(rightMid))) {
			base32For9.add(rightMid);
		}
		String rightDown = getGeoHashBase32(rightLat, downLng);
		if (!(rightDown == null || "".equals(rightDown))) {
			base32For9.add(rightDown);
		}
		return base32For9;
	}

	/**
	 * @Description: 获取经纬度的base32字符串
	 * @return
	 */
	public String getGeoHashBase32() {
		return getGeoHashBase32(localtionPo.getLatitude(), localtionPo.getLongitude());
	}

	public static void main(String[] args) {
		// 116.487829,39.996379 wx4gd9pg
		// 116.471965,40.018844 wx4gdmhq
		GeoHash g = new GeoHash(116.487829, 39.996379);
		g.sethashLength(5);
		String geoHash = g.getGeoHashBase32();
		System.out.println(geoHash);
		List<String> geoHash9 = g.getGeoHashBase32For9();
		System.out.println(geoHash9);

		NearbyPo localtionPo = g.getLocation(geoHash);
		System.out.println(localtionPo);
		System.out.println(new GeoHash(localtionPo.getLongitude(), localtionPo.getLatitude()).getGeoHashBase32());
		System.out.println(localtionPo.getLongitude() + "," + g.minLng + "," + (localtionPo.getLongitude() - g.minLng));
		System.out.println(localtionPo.getLatitude() + "," + g.minLat + "," + (localtionPo.getLatitude() - g.minLat));
		System.out.println(DistanceUtil.getDistance(localtionPo.getLongitude(), localtionPo.getLatitude(),
				localtionPo.getLongitude() - g.minLng, localtionPo.getLatitude() - g.minLat));
	}

}
