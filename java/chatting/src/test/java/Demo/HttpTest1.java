/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 */
package Demo;

import java.util.ArrayList;
import java.util.List;

import org.apache.http.HttpResponse;
import org.apache.http.NameValuePair;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.DefaultHttpClient;
import org.apache.http.message.BasicNameValuePair;
import org.apache.http.util.EntityUtils;

import com.lanlian.chatting.util.Base64s;
import com.lanlian.chatting.util.CheckApp;

/**
 * @author wdyqxx
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年2月21日 下午8:27:23
 * @explain: 模拟POST请求；
 */

public class HttpTest1 {

	@SuppressWarnings("unused")
	public static void main(String[] args) throws Exception {

		String aliyun = "http://app.pepfi.com";
		String ip = "http://192.168.0.54";
		String projects = "/chatting";
		String request = "/v1/group_settings";
		String requestbd = "/sync";
		String interfaceId = "/11";
		String url = ip + projects + request + requestbd ;
		HttpPost httpPost = new HttpPost(url);
		DefaultHttpClient httpClient = new DefaultHttpClient();

		long pid = 1000001L;
		String appSecret = Base64s.encryptBASE64(String.valueOf(pid));
		long now = System.currentTimeMillis();
		long now1 = 1490269860242L;
		String nonce = CheckApp.getMD5Hex(pid + String.valueOf(now % 1000));
		String curTime = String.valueOf(now / 1000 / 60);
		String signature = CheckApp.getSignature(appSecret, nonce, curTime);// 参考计算CheckSum的java代码

		// 设置请求的header
		httpPost.addHeader("AppSecret", appSecret.trim());
		httpPost.addHeader("Nonce", nonce);
		httpPost.addHeader("CurTime", curTime);
		httpPost.addHeader("Signature", signature);
		httpPost.addHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8");

		// 设置请求的参数
		List<NameValuePair> nvps = new ArrayList<NameValuePair>();
		nvps.add(new BasicNameValuePair("tgid", "100011350"));
		nvps.add(new BasicNameValuePair("tgname", "嗒嗒打的"));
		nvps.add(new BasicNameValuePair("pid", "1000001"));
		nvps.add(new BasicNameValuePair("members", "[1000001, 1000002,1000003]"));
		nvps.add(new BasicNameValuePair("notice", "望京"));
		nvps.add(new BasicNameValuePair("intro", "望京"));
		nvps.add(new BasicNameValuePair("longitude", "1.1253"));
		nvps.add(new BasicNameValuePair("latitude", "2.1212"));
		nvps.add(new BasicNameValuePair("createTime", "1489909978734"));
		nvps.add(new BasicNameValuePair("icon", "1"));

		httpPost.setEntity(new UrlEncodedFormEntity(nvps, "utf-8"));
		// 执行请求
		HttpResponse response = httpClient.execute(httpPost);
		// 打印执行结果
		System.out.println(EntityUtils.toString(response.getEntity(), "utf-8"));
	}

}
