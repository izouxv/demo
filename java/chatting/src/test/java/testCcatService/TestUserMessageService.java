package testCcatService;

import java.io.UnsupportedEncodingException;
import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.UserLetterService;
import com.lanlian.chatting.vo.UserLetter;
import com.lanlian.chatting.vo.UserMessageNum;

import net.sf.json.JSONArray;

/**
 * @author wdyqxx
 * @version 2017年1月4日 上午11:23:25
 * @explain 此类用于测试UserMessageService接口的所有方法；
 */
public class TestUserMessageService {

	ClassPathXmlApplicationContext ctx;
	UserLetterService umService;

	@Before
	public void init() {
		ctx = new ClassPathXmlApplicationContext("spring-service.xml", "spring-controller.xml", "spring-mybatis.xml");
		umService = ctx.getBean("userMessageService", UserLetterService.class);
	}

	@Test
	public void sendPrivateLetteTest() throws Parameter_Exception {
		UserLetter umInfo = new UserLetter();
		umInfo.setUid(1000001);
		umInfo.setTouid(1000002);
		umInfo.setType("1");
		umInfo.setLetter("测试发送私信的service1");
		umService.sendPrivateLette(umInfo);
	}

	@Test
	public void receiveLetteNumTest() throws Parameter_Exception {
		UserMessageNum umn = new UserMessageNum();
		umn.setOpid(1000001);
		JSONArray lists = umService.receiveLetteNum(umn);
		System.out.println(lists);
	}

	@Test
	public void receivePrivateMessageInfoTest() throws Parameter_Exception, UnsupportedEncodingException {
		UserLetter umInfo = new UserLetter();
		umInfo.setUid(1000001);
		umInfo.setTouid(1000002);
		System.out.println(umInfo);
		List<UserLetter> list = umService.receiveLetter(umInfo);
		System.out.println(list);
	}

}
