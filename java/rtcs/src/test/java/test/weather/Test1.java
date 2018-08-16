/** 
 *<p>开发公司 :		           蓝涟科技<p>
 *<p>版权所有 :		           蓝涟科技<p>
 *<p>责任人     :		               王东阳<p> 
 *<p>网址         :    www.radacat.com<p>
 *<p>邮箱         : wangdy@radacat.com<p>
 */

package test.weather;

import java.io.InputStream;
import java.io.OutputStream;
import java.io.OutputStreamWriter;
import java.net.URL;
import java.net.URLConnection;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;

import org.w3c.dom.Document;
import org.w3c.dom.Node;
import org.w3c.dom.NodeList;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月13日 上午10:22:57
 * @$
 * @Administrator
 * @explain  
 * http://www.webxml.com.cn/WebServices/WeatherWS.asmx
 * http://www.cnblogs.com/jason-star/archive/2012/09/25/2702032.html
 */

public final class Test1 {

	/**
	 * 获取SOAP的请求头，并替换其中的标志符号为用户输入的城市
	 * 
	 * @param city
	 *            用户输入的城市名称
	 * @return 客户将要发送给服务器的SOAP请求
	 */
	private static String getSoapRequest(String city) {
		StringBuilder sb = new StringBuilder();
		sb.append("<?xml version=\"1.0\" encoding=\"utf-8\"?>"
				+ "<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" "
				+ "xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" "
				+ "xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">"
				+ "<soap:Body>    <getWeather xmlns=\"http://WebXml.com.cn/\">" + "<theCityCode>" + city
				+ "</theCityCode>    </getWeather>" + "</soap:Body></soap:Envelope>");
		return sb.toString();
	}

	/**
	 * 用户把SOAP请求发送给服务器端，并返回服务器点返回的输入流
	 * 
	 * @param city
	 *            用户输入的城市名称
	 * @return 服务器端返回的输入流，供客户端读取
	 * @throws Exception
	 */
	private static InputStream getSoapInputStream(String city) throws Exception {
		try {
			String soap = getSoapRequest(city);
			if (soap == null) {
				return null;
			}
			URL url = new URL("http://www.webxml.com.cn/WebServices/WeatherWS.asmx");
			URLConnection conn = url.openConnection();
			conn.setUseCaches(false);
			conn.setDoInput(true);
			conn.setDoOutput(true);

			conn.setRequestProperty("Content-Length", Integer.toString(soap.length()));
			conn.setRequestProperty("Content-Type", "text/xml; charset=utf-8");
			conn.setRequestProperty("SOAPAction", "http://WebXml.com.cn/getWeather");

			OutputStream os = conn.getOutputStream();
			OutputStreamWriter osw = new OutputStreamWriter(os, "utf-8");
			osw.write(soap);
			osw.flush();
			osw.close();
			InputStream is = conn.getInputStream();
			return is;
		} catch (Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	/**
	 * 对服务器端返回的XML进行解析
	 * 
	 * @param city
	 *            用户输入的城市名称
	 * @return 字符串 用#分割
	 */
	public static String getWeather(String city) {
		try {
			Document doc;
			DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
			dbf.setNamespaceAware(true);
			DocumentBuilder db = dbf.newDocumentBuilder();
			InputStream is = getSoapInputStream(city);
			doc = db.parse(is);
			NodeList nl = doc.getElementsByTagName("string");
			StringBuffer sb = new StringBuffer();
			for (int count = 0; count < nl.getLength(); count++) {
				Node n = nl.item(count);
				if ("查询结果为空！".equals(n.getFirstChild().getNodeValue())) {
					sb = new StringBuffer("\n");
					break;
				}
				sb.append(n.getFirstChild().getNodeValue() + "\n");
			}
			is.close();
			return sb.toString();
		} catch (Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	/**
	 * 测试
	 * 
	 * @param args
	 * @throws Exception
	 */
	public static void main(String[] args) throws Exception {
		String weatherInfo = getWeather("北京");
		System.out.println(weatherInfo);
	}

}