/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.server.http;

import java.io.IOException;
import java.io.InputStream;
import java.io.UnsupportedEncodingException;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

import org.apache.http.HttpResponse;
import org.apache.http.NameValuePair;
import org.apache.http.ParseException;
import org.apache.http.client.ClientProtocolException;
import org.apache.http.client.HttpClient;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.DefaultHttpClient;
import org.apache.http.util.EntityUtils;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.NotificationPo;
import com.lanlian.chatting.result.Fatal_Exception;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.rpc.feedback.AddFeedbackReply;

/**
 * @Title Notification.java
 * @Package com.lanlian.server.http
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月14日 上午9:37:08
 * @explain
 */
@Service("notification")
public class Notification {

	public static void main(String[] args) throws IOException, Fatal_Exception {
		Notification notification = new Notification();
		notification.sendMail("",null,"test", "", "", "", "", "",new String[]{"test01"}, "");
	}
	private String url = "http://192.168.1.51:7023/v1.0/notices";
	private HttpPost httpPost = null;

	/**
	 * 发送邮件提醒
	 * @param description
	 * @param mobileInfo
	 * @param appInfo
	 * @param files
	 * @param contact
	 * @param deviceInfo
	 * @param userInfo
	 * @param extendInfo
	 * @throws Fatal_Exception 
	 */
	@Async
	public void sendMail(String mail, AddFeedbackReply reply, String description,String mobileInfo, String appInfo, String deviceInfo, 
			String userInfo, String extendInfo,String[] files,String contact) throws Fatal_Exception {
		try {
			String fileStr = "";
			if (files != null) {
				for (String file : files) {
					fileStr += file;
					fileStr += " ; ";
				}
			}
			if (contact == null) {
				contact = "";
			} else {
				contact += " ;";
			}
			if (deviceInfo == null) {
				deviceInfo = "";
			} else {
				deviceInfo += " ;";
			}
			if (userInfo == null) {
				userInfo = "";
			} else {
				userInfo += " ;";
			}
			if (extendInfo == null) {
				extendInfo = "";
			} else {
				extendInfo += " ;";
			}
			String info = "<pre>联系方式："+	description+	"</pre>"+
				"<pre>描述信息："+	mobileInfo+		"</pre>"+
				"<pre>手机信息："+	appInfo+		"</pre>"+
				"<pre>应用信息："+	contact+		"</pre>"+
				"<pre>设备信息："+	deviceInfo+		"</pre>"+
				"<pre>用户信息："+	userInfo+		"</pre>"+
				"<pre>文件信息："+	fileStr+		"</pre>"+
				"<pre>扩展信息："+	extendInfo+		"</pre>";
			JSONObject content = new JSONObject();
			content.put("tos", mail);
			content.put("subject", "APP 建议与反馈功能-工单ID："+reply.getId());
			content.put("content", info);
			JSONObject json = new JSONObject();
			json.put("tem_id", 5);
			json.put("email_addr", mail);
			json.put("send_data", content);
			LogUtil.info("***json***"+json);
			HttpClient httpClient = new DefaultHttpClient();
			httpPost = new HttpPost(url);
			// 设置请求的header
			httpPost.addHeader("Content-Type", RequestSetting.PRODUCES);
			// 设置请求的参数
			StringEntity data = new StringEntity(json.toJSONString(), "utf-8");
			httpPost.setEntity(data);
			// 执行请求
			HttpResponse response = httpClient.execute(httpPost);
			String resStr = EntityUtils.toString(response.getEntity(), "utf-8");
			if (response.getStatusLine().getStatusCode() != 200) {
				throw new Fatal_Exception("群消息通知微信失败，"+response.getStatusLine().getStatusCode()+",resStr:"+resStr);
			}
			LogUtil.info("sendNotice:"+resStr);
		} catch (Exception e) {
			throw new Fatal_Exception(e);
		} finally {
			httpPost.abort();
		}
	}
}
