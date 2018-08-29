/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package com.lanlian.chatting.localtion;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月17日 下午5:09:17
 * @explain 
 */

public class DistanceUtil {
	
	private static final  double EARTH_RADIUS = 6378137;//赤道半径
	
	private static double rad(double d){
	    return d * Math.PI / 180.0;
	}
	
	/**
	 * 计算经纬度两点之间的距离
	 * @param lon1
	 * @param lat1
	 * @param lon2
	 * @param lat2
	 * @return
	 */
	public static double getDistance(double lon1,double lat1,double lon2, double lat2) {
	    double radLat1 = rad(lat1);
	    double radLat2 = rad(lat2);
	    double a = radLat1 - radLat2;
	    double b = rad(lon1) - rad(lon2);
	    double s = 2 *Math.asin(Math.sqrt(Math.pow(Math.sin(a/2),2)+Math.cos(radLat1)*Math.cos(radLat2)*Math.pow(Math.sin(b/2),2))); 
	    s = s * EARTH_RADIUS;    
	   return s;//单位米
	}
	
	public static void main(String[] args) {
		// 116.487829,39.996379 wx4gd9pg
		// 116.471965,40.018844 wx4gdmhq
		//116.48786544799805,3.4332275390625E-4,116.48752212524414
		//39.996328353881836,1.71661376953125E-4,39.99615669250488
		//34.96291571640858
		System.out.println(getDistance(116.48786544799805,39.996328353881836, 116.48752212524414,39.99615669250488));
	}

}

