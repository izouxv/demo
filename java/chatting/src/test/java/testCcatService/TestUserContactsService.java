package testCcatService;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.dao.FriendsDao;
import com.lanlian.chatting.po.FriendInfoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.FriendService;
import com.lanlian.chatting.vo.UserFriends;

/**
 * @author  wdyqxx
 * @version 2017年1月4日 上午11:23:25
 * @explain 此类用于测试UserContactsService接口的所有方法；
 */
public class TestUserContactsService {
	
	ClassPathXmlApplicationContext ctx;
	FriendsDao userContactsDao;
	FriendService ucService;
	
	@Before
	public void init(){
		ctx=new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-controller.xml",
				"spring-mybatis.xml"
				);
		userContactsDao=ctx.getBean("userContactsDao",FriendsDao.class);
		ucService=ctx.getBean("userContactsService",FriendService.class);
	}

	@Test
	public void myFriendsTest() throws Parameter_Exception{
		List<FriendInfoPo> list=ucService.findFriends(null,"");
		System.out.println("联系人信息："+list);
	}
	
	@Test 
	public void backupFriends() throws Parameter_Exception{
		int fpid=1000001;
		int pid1=1000002;
		int pid2=1000003;
		int pid3=1000004;
		String note1="laohua";
		String note2="laobin";
		String note3="laowang";
		UserFriends userContacts=new UserFriends();
		userContacts.setUid(fpid);
		Map<String, String> map=new HashMap<>();
		map.put(pid1+"", note1);
		map.put(pid2+"", note2);
		map.put(pid3+"", note3);
		userContacts.setInfoMap(map);
		System.out.println(userContacts);
		ucService.saveFriends(userContacts, null);
	}
	
	
	
	
	
	
}
