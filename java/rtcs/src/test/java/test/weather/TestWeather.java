/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package test.weather;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月13日 上午10:13:50
 * @$
 * @Administrator
 * @explain 
 */

public class TestWeather {
	/**
	 *  Java通过webservice获取天气预报情况
	 * @param args
	 */
	public static void main(String[] args) {
		WeatherWebService weatherWebService = new WeatherWebService();
		WeatherWebServiceSoap soap = weatherWebService.getWeatherWebServiceSoap();
		System.out.println("----  info  ------");
		ArrayOfString info = soap.getWeatherbyCityName("上海");
		for (String str : info.getString()) {
			System.out.println("天气："+str);
		}
		
	}

}

