package testCcatDao;

import java.util.ArrayList;
import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.dao.UserMessageDao;
import com.lanlian.chatting.po.PrivateMessageInfoPO;
import com.lanlian.chatting.po.PrivateMessagePO;


/**
 * @author  wdyqxx
 * @version 2017年1月6日 上午10:00:37
 * @explain 此类用于测试UserMessageDao接口的所有方法；
 */

public class TestUserMessageDao {
	
	ClassPathXmlApplicationContext ctx;
	UserMessageDao umDao;
	
	@Before
	public void init(){
		ctx=new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-mybatis.xml"
				);
		umDao=ctx.getBean("userMessageDao",UserMessageDao.class);
	}
	
	@Test
	public void  sendPrivateLetteInfoIdTest(){
		String messageInfo="测试发送私信内容6";
		PrivateMessageInfoPO pmipojo=new PrivateMessageInfoPO();
		pmipojo.setMessageInfo(messageInfo);
		umDao.sendPrivateLetteInfoId(pmipojo);
		System.out.println(pmipojo);
	}
	
	@Test
	public void sendPrivateLetteTest(){
		int fpid=1000001;
		int opid=1000002;
		String type="1";
		String messageInfoId="71843";
		PrivateMessagePO pmpojo=new PrivateMessagePO();
		pmpojo.setOuid(fpid);
		pmpojo.setUid(opid);
		pmpojo.setType(type);
		pmpojo.setMessageInfoId(messageInfoId);
		int plid=umDao.sendPrivateLette(pmpojo);
		System.out.println(plid);
	}
	
	@Test
	public void receiveLetteByIdTest(){	
		PrivateMessagePO pmpojo=new PrivateMessagePO();
		pmpojo.setOuid(1000001);
		System.out.println(pmpojo);
		List<PrivateMessagePO> list=umDao.receiveLetteById(pmpojo);
		for (PrivateMessagePO pmpojo1 : list) {
			System.out.println("pmpojo1:"+pmpojo1);
		}
		
	}
	
	@Test
	public void receiveLetteInfoTest(){	
		List<Long> as = new ArrayList<>();
		for (int i = 0; i < 4; i++) {
//			PrivateMessageInfoPO po = new PrivateMessageInfoPO();
//			po.setMessageInfoId(206+i);
			as.add((long) (206+i));
		}
		List<PrivateMessageInfoPO> asas=umDao.receiveLetteInfo(as);
		System.out.println(asas);
		
	}
	
	
	
}




