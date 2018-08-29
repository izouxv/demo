package com.lanlian.chatting.util;

import javax.mail.MessagingException;

import org.apache.log4j.Logger;
import org.springframework.scheduling.annotation.Async;

import com.aliyuncs.DefaultAcsClient;
import com.aliyuncs.IAcsClient;
import com.aliyuncs.dysmsapi.model.v20170525.SendSmsRequest;
import com.aliyuncs.dysmsapi.model.v20170525.SendSmsResponse;
import com.aliyuncs.exceptions.ClientException;
import com.aliyuncs.exceptions.ServerException;
import com.aliyuncs.http.MethodType;
import com.aliyuncs.profile.DefaultProfile;
import com.aliyuncs.profile.IClientProfile;
import com.lanlian.chatting.result.Parameter_Exception;

/**
 * 
 * @author wangdyq
 *
 */
public class SendCode {

	private static Logger log = Logger.getLogger(SendCode.class);
	
	/**
	 * @throws Parameter_Exception 调用阿里大于接口发送短信验证码
	 * 
	 * @param username @param activeCode @throws
	 */
	@Async
	public void aliAuthCode(String action, String username, String code) {
		try {
			// 设置超时时间-可自行调整
			System.setProperty("sun.net.client.defaultConnectTimeout", "10000");
			System.setProperty("sun.net.client.defaultReadTimeout", "10000");
			// 初始化ascClient需要的几个参数
			// 短信API产品名称（短信产品名固定，无需修改）
			final String product = "Dysmsapi";
			// 短信API产品域名（接口地址固定，无需修改）
			final String domain = SMSparameter.URL;
			//accessKeyId,accessKeySecret
			final String accessKeyId = SMSparameter.ACCESSKEYID;
			final String accessKeySecret = SMSparameter.ACCESSKEYSECRET;
			// 初始化ascClient,暂时不支持多region（请勿修改）
			IClientProfile profile = DefaultProfile.getProfile("cn-hangzhou", accessKeyId, accessKeySecret);
			DefaultProfile.addEndpoint("cn-hangzhou", "cn-hangzhou", product, domain);
			IAcsClient acsClient = new DefaultAcsClient(profile);
			// 组装请求对象
			SendSmsRequest request = new SendSmsRequest();
			// 使用post提交
			request.setMethod(MethodType.POST);
			// 必填:待发送手机号。支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式
			request.setPhoneNumbers(username);
			// 必填:短信签名-可在短信控制台中找到
			request.setSignName(SMSparameter.NAME);
			// 必填:短信模板-可在短信控制台中找到
			switch (action) {
			case "reg":
				request.setTemplateCode(SMSparameter.CODEREG);
				break;
			case "reset":
				request.setTemplateCode(SMSparameter.CODERESET);
				break;
			}
			// 可选:模板中的变量替换JSON串,如模板内容为"亲爱的${name},您的验证码为${code}"时,此处的值为
			// 友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\\r\\n,否则会导致JSON在服务端解析失败
			request.setTemplateParam("{\"code\":\""+code+"\"}");
			// 可选-上行短信扩展码(扩展码字段控制在7位或以下，无特殊需求用户请忽略此字段)
			// request.setSmsUpExtendCode("90997");
			// 可选:outId为提供给业务方扩展字段,最终在短信回执消息中将此值带回给调用者
			request.setOutId("yourOutId");
			// 请求失败这里会抛ClientException异常
			SendSmsResponse sendSmsResponse = acsClient.getAcsResponse(request);
			log.info("SendCode:"+sendSmsResponse.getCode()+","+sendSmsResponse.getMessage());
			if (sendSmsResponse.getCode() != null && sendSmsResponse.getCode().equals("OK")) {
				// 请求成功
			}
		} catch (ClientException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}
	
	/**
	 * 对邮箱发送验证码
	 * 
	 * @param username
	 * @param authCode
	 * @throws Parameter_Exception
	 * @throws MessagingException
	 * @throws Exception
	 */
	@SuppressWarnings("unused")
	private void authCodeEmail(String username, String authCode) throws Parameter_Exception {
		try {
			MailSender mailInfo = new MailSender();
			mailInfo.setMailServerHost("mail.radacat.com");
			mailInfo.setMailServerPort("25");
			mailInfo.setValidate(true);
			mailInfo.setUserName("wangdy@radacat.com");// 邮箱账号
			mailInfo.setPassword("dongyang");// 邮箱密码
			mailInfo.setFromAddress("wangdy@radacat.com");
			mailInfo.setToAddress(username);
			mailInfo.setSubject("注册激活");
			mailInfo.setContent("您的验证码为:" + authCode + ",5分钟内输入有效");
			SimpleMailSender sms = new SimpleMailSender();
			sms.sendTextMail(mailInfo);// 发送文体格式
		} catch (Exception e) {
			// "邮件发送失败
			throw new Parameter_Exception(10005);
		}
	}


	public static void main(String[] args) throws ServerException, ClientException {
		// 设置超时时间-可自行调整
		System.setProperty("sun.net.client.defaultConnectTimeout", "10000");
		System.setProperty("sun.net.client.defaultReadTimeout", "10000");
		// 初始化ascClient需要的几个参数
		// 短信API产品名称（短信产品名固定，无需修改）
		final String product = "Dysmsapi";
		// 短信API产品域名（接口地址固定，无需修改）
		final String domain = "dysmsapi.aliyuncs.com";
		//accessKeyId,accessKeySecret
		final String accessKeyId = "LTAI3pqHeu2ht5p5";
		final String accessKeySecret = "N87mmOHNA3q2AoubRFRaYW32CqUFVp";
		// 初始化ascClient,暂时不支持多region（请勿修改）
		IClientProfile profile = DefaultProfile.getProfile("cn-hangzhou", accessKeyId, accessKeySecret);
		DefaultProfile.addEndpoint("cn-hangzhou", "cn-hangzhou", product, domain);
		IAcsClient acsClient = new DefaultAcsClient(profile);
		// 组装请求对象
		SendSmsRequest request = new SendSmsRequest();
		// 使用post提交
		request.setMethod(MethodType.POST);
		// 必填:待发送手机号。支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式
		request.setPhoneNumbers("17600117962");
		// 必填:短信签名-可在短信控制台中找到
		request.setSignName(SMSparameter.NAME);
		// 必填:短信模板-可在短信控制台中找到
		request.setTemplateCode(SMSparameter.CODEREG);
		// 可选:模板中的变量替换JSON串,如模板内容为"亲爱的${name},您的验证码为${code}"时,此处的值为
		// 友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\\r\\n,否则会导致JSON在服务端解析失败
		request.setTemplateParam("{\"code\":\""+"123456"+"\"}");
		// 可选-上行短信扩展码(扩展码字段控制在7位或以下，无特殊需求用户请忽略此字段)
		// request.setSmsUpExtendCode("90997");
		// 可选:outId为提供给业务方扩展字段,最终在短信回执消息中将此值带回给调用者
		request.setOutId("yourOutId");
		// 请求失败这里会抛ClientException异常
		SendSmsResponse sendSmsResponse = acsClient.getAcsResponse(request);
		System.out.println("1111:"+sendSmsResponse.getCode()+","+sendSmsResponse.getMessage());
		if (sendSmsResponse.getCode() != null && sendSmsResponse.getCode().equals("OK")) {
			// 请求成功
		}
	}

}