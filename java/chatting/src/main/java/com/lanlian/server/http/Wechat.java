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
import java.net.URLDecoder;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

import javax.annotation.Resource;

import org.apache.http.HttpResponse;
import org.apache.http.NameValuePair;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.DefaultHttpClient;
import org.apache.http.message.BasicNameValuePair;
import org.apache.http.util.EntityUtils;
import org.apache.log4j.Logger;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

import com.lanlian.chatting.dao.GroupSettingsDao;
import com.lanlian.chatting.dao.GroupSettingsDao_1_1;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.LiveChatGidUid;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.po.WechatPo;
import com.lanlian.chatting.result.Fatal_Exception;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @Title Wechat.java
 * @Package com.lanlian.chatting.redis
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月14日 上午9:37:08
 * @explain
 */
@Service("wechat")
public class Wechat {

	private static Logger logger = Logger.getLogger(Wechat.class);
	
	public static void main(String[] args) throws IOException {
		Wechat wechat = new Wechat();
		System.out.println(wechat.url);
	}

	private String host = "http://wechat.radacat.com";
	private String port = "80";
	private String server = "/wechat";
	private String version = "/v1.0";
	private String pathMsgP = "/message";
	private String pathMsgB = "/api/send_notice";
	private String hp = ":";
	private String url = "";

	/**
	 * 获取Wechat的地址信息
	 * 
	 * @return
	 */
	public WechatPo wechatAddress() {
		WechatPo wechatPo = new WechatPo();
		Properties properties = new Properties();
		InputStream in = Wechat.class.getResourceAsStream("/server.properties");
		try {
			properties.load(in);
			wechatPo.setHost(properties.getProperty("wechat.host"));
			wechatPo.setPort(properties.getProperty("wechat.port"));
			wechatPo.setServer(properties.getProperty("wechat.server"));
			wechatPo.setVersion(properties.getProperty("wechat.version"));
			wechatPo.setPathMsgP(properties.getProperty("wechat.pathMsgP"));
			wechatPo.setPathMsgB(properties.getProperty("wechat.pathMsgB"));
			System.out.println("try-wechatPo:" + wechatPo);
			return wechatPo;
		} catch (Exception e) {
			LogUtil.error(e);
			wechatPo.setHost(host);
			wechatPo.setPort(port);
			wechatPo.setServer(server);
			wechatPo.setVersion(version);
			wechatPo.setPathMsgP(pathMsgP);
			wechatPo.setPathMsgB(pathMsgB);
			System.out.println("catch-wechatPo:" + wechatPo);
			return wechatPo;
		}
	}

	@Resource
	GroupSettingsDao groupSettingsDao;
	@Resource
	GroupSettingsDao_1_1 groupSettingsDao_1_1;
	
	HttpPost httpPost = null;

	@Async
	public void send(DadaGroupPo dadaGroupPo, SsoPo ssoPo) throws Fatal_Exception {
		try {
			LiveChatGidUid liveChatGidUid = new LiveChatGidUid();
			liveChatGidUid.setGid(dadaGroupPo.getGid());
			liveChatGidUid = groupSettingsDao.findBundilingBygid(liveChatGidUid);
			if (liveChatGidUid == null) {
				return;
			}
			WechatPo wechatPo = wechatAddress();
			if (url.equals("")) {
				url = url+wechatPo.getHost()+hp+wechatPo.getPort()+wechatPo.getServer()+wechatPo.getVersion()+wechatPo.getPathMsgP()+wechatPo.getPathMsgB();
			}
			LogUtil.info("wechatPo-url" + url.toString());
			DefaultHttpClient httpClient = new DefaultHttpClient();
			httpPost = new HttpPost(url.toString());
			// 设置请求的header
			httpPost.addHeader("Content-Type", RequestSetting.CONSUMES);
			// 设置请求的参数
			List<NameValuePair> nvps = new ArrayList<NameValuePair>();
			nvps.add(new BasicNameValuePair("access_token", "fwG2gemk9d"));
			liveChatGidUid.setGname(URLDecoder.decode(liveChatGidUid.getGname(), "UTF-8"));
			nvps.add(new BasicNameValuePair("gname", liveChatGidUid.getGname()));
			nvps.add(new BasicNameValuePair("username", ssoPo.getUsername()));
			nvps.add(new BasicNameValuePair("gid", dadaGroupPo.getGid().toString()));
			httpPost.setEntity(new UrlEncodedFormEntity(nvps, "utf-8"));
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
	
	@Async
	public void send_1_1(DadaGroupPo dadaGroupPo, SsoPo ssoPo) throws Fatal_Exception {
		logger.info("Wechat --- send_1_1 --- dadaGroupPo:"+dadaGroupPo+",ssoPo:"+ssoPo);
		try {
			LiveChatGidUid liveChatGidUid = new LiveChatGidUid();
			liveChatGidUid.setGid(dadaGroupPo.getGid());
			liveChatGidUid = groupSettingsDao_1_1.findBundilingBygid(liveChatGidUid);
			if (liveChatGidUid == null) {
				return;
			}
			liveChatGidUid.setGname(URLDecoder.decode(liveChatGidUid.getGname(), "UTF-8"));
			WechatPo wechatPo = wechatAddress();
			if (url.equals("")) {
				url = url + wechatPo.getHost() + hp + wechatPo.getPort() + wechatPo.getServer() + wechatPo.getVersion()
						+ wechatPo.getPathMsgP() + wechatPo.getPathMsgB();
			}
			LogUtil.info("wechatPo-url:" + url.toString());
			httpPost = new HttpPost(url.toString());
			DefaultHttpClient httpClient = new DefaultHttpClient();

			// 设置请求的header
			httpPost.addHeader("Content-Type", RequestSetting.CONSUMES);
			// 设置请求的参数
			List<NameValuePair> nvps = new ArrayList<NameValuePair>();
			nvps.add(new BasicNameValuePair("access_token", "fwG2gemk9d"));
			nvps.add(new BasicNameValuePair("gname", liveChatGidUid.getGname()));
			nvps.add(new BasicNameValuePair("username", ssoPo.getUsername()));
			nvps.add(new BasicNameValuePair("gid", dadaGroupPo.getGid().toString()));
			httpPost.setEntity(new UrlEncodedFormEntity(nvps, "utf-8"));
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
