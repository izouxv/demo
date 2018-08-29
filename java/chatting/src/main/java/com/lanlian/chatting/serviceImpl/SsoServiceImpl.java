/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.serviceImpl;

import javax.annotation.Resource;

import org.apache.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

import com.aliyuncs.DefaultAcsClient;
import com.aliyuncs.IAcsClient;
import com.aliyuncs.dysmsapi.model.v20170525.SendSmsRequest;
import com.aliyuncs.dysmsapi.model.v20170525.SendSmsResponse;
import com.aliyuncs.exceptions.ClientException;
import com.aliyuncs.http.MethodType;
import com.aliyuncs.profile.DefaultProfile;
import com.aliyuncs.profile.IClientProfile;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.rpc.TwinsClient;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.util.DataFinals;
import com.lanlian.chatting.util.SMSparameter;
import com.lanlian.chatting.util.SendCode;
import com.lanlian.server.redis.RedisClient79;

/**
 * @Title SsoServiceImpl.java
 * @Package com.lanlian.chatting.serviceImpl
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月8日 下午3:53:11
 * @explain
 */
@Service("ssoService")
public class SsoServiceImpl implements SsoService {

	@Autowired(required = true)
	SsoClient ssoClient;
	
	@Resource
	RedisClient79 redisClient79;
	
	private static Logger log = Logger.getLogger(SendCode.class);
	
	/**
	 * 检验source
	 * @param ssoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public SsoPo getSource(SsoPo ssoPo) throws Parameter_Exception {
		String value = redisClient79.getString(redisClient79.key13+ssoPo.getSource());
		if (value == null || value.isEmpty()) {
			throw new Parameter_Exception(21002);
		}
		LogUtil.info("SsoServiceImpl-getSource:" + value + ",ssoPo:" + ssoPo);
		return ssoPo;
	}

	/**
	 * 验证头信息
	 * @param request
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public SsoPo verifyToken(SsoPo ssoPo) throws Parameter_Exception {
		ssoPo = ssoClient.getSsoInfo(ssoPo);
		if (10000 != ssoPo.getErrorCode()) {
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
		return ssoPo;
	}

	/**
	 * 检验用户是否被注册,1：无此用户；2：有此用户
	 * 
	 * @param ssoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public int checkUser(SsoPo ssoPo) throws Parameter_Exception {
		ssoPo = ssoClient.getUserByName(ssoPo);
		if (ssoPo.getErrorCode() == 33002) {
			return 1;
		}
		if (ssoPo.getErrorCode() == 10000) {
			return 2;
		}
		throw new Parameter_Exception(ssoPo.getErrorCode());
	}

	/**
	 * 添加用户
	 * 
	 * @param ssoPo
	 * @return
	 * @throws Parameter_Exception
	 */
	@Override
	public SsoPo add(SsoPo ssoPo) throws Parameter_Exception {
		ssoPo = ssoClient.add(ssoPo);
		return ssoPo;
	}

	/**
	 * 找回密码发送邮件
	 * 
	 * @param ssoPo
	 * @throws Parameter_Exception
	 */
	@Override
	public void sendEmail(SsoPo ssoPo) throws Parameter_Exception {
		ssoPo = ssoClient.sendEmail(ssoPo);
		if (10000 != ssoPo.getErrorCode()) {
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
	}

	/**
	 * 发送短信
	 * 
	 * @param ssoPo
	 * @throws Parameter_Exception
	 */
	@Override
	public void sendMobile(SsoPo ssoPo, String action) throws Parameter_Exception {
		switch (action) {
		// 注册业务
		case DataFinals.REGISTER:
			ssoPo = ssoClient.sendMobileCode(ssoPo, 1);
			if (ssoPo.getErrorCode() == 33008) {
				throw new Parameter_Exception(20008);
			}
			break;
		// 找回密码
		case DataFinals.RESETPWD:
			ssoPo = ssoClient.sendMobileCode(ssoPo, 2);
			break;
		default:
			break;
		}
		if (10000 != ssoPo.getErrorCode()) {
			throw new Parameter_Exception(ssoPo.getErrorCode());
		}
	}
	

	
	/**
	 * 调用阿里接口发送短信验证码
	 * @param action
	 * @param username
	 * @param code
	 */
	@Async
	public void aliCode(String action, String username, String code) {
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
	 * 校验短信
	 * 
	 * @param ssoPo
	 * @throws Parameter_Exception
	 */
	@Override
	public void verifyMobile(String username, String code, int action, String source) throws Parameter_Exception {
		ssoClient.verifyMobileCode(username, code, action, source);
	}
	
	@Resource
	TwinsClient client;

	@Override
	@Async
	public void analysis(String json) {
		client.reported(json);
	}

}
