package testCcatService;


import org.junit.Before;
import org.junit.Ignore;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.SsoService;
import com.lanlian.chatting.service.AccountService;

/**
 * @author  wdyqxx
 * @version 2017年1月4日 上午11:23:25
 * @explain 此类用于测试UserInfoService接口的所有方法；
 */
@SuppressWarnings("unused")
public class TestUserInfoService {
	
	ClassPathXmlApplicationContext ctx;
	AccountService uiService;
	SsoService ssoService;
	
	@Before
	public void init(){
		ctx=new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-controller.xml",
				"spring-mybatis.xml"
				);
		uiService=ctx.getBean("userInfoService",AccountService.class);
		ssoService=ctx.getBean("ssoService", SsoService.class);
	}
	
	
	@Test
	public void ssoTest() throws Parameter_Exception{
		SsoPo ssoPo = new SsoPo();
		System.out.println(ssoPo);
	}
	
	
	
	
	
}
