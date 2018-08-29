package testCcatDao;

import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.bo.MessagePageBo;
import com.lanlian.chatting.dao.WeiChatGroupSettingsDao;
import com.lanlian.chatting.vo.MessageInfoPage;

/** 
* @author  yfq
* @version 2017年3月27日 上午11:31:29
*  xxx.... 
*/
public class TestGroupMsgDao {
	
	ClassPathXmlApplicationContext ctx;
	WeiChatGroupSettingsDao weiChatGroupSettingsDao;
	
	@Before
	public void init(){
		ctx = new ClassPathXmlApplicationContext(
				"spring-mybatis.xml"
				);
		weiChatGroupSettingsDao = ctx.getBean("weiChatGroupSettingsDao",WeiChatGroupSettingsDao.class);
	}
	
	@Test
	public void saveGroupMemdersTest(){
		MessagePageBo messagePageBo = new MessagePageBo();
		messagePageBo.setGid(100000001);
		messagePageBo.setCount(100);
		List<MessageInfoPage> list = weiChatGroupSettingsDao.findMessageInfo(messagePageBo);
		System.out.println(list);
	}
} 



