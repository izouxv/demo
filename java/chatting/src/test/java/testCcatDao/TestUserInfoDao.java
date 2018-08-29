package testCcatDao;

import java.util.Date;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.dao.UserInfoDao;
import com.lanlian.chatting.po.UserInfoPO;
import com.lanlian.chatting.util.DataFinals.Gender;

/**
 * @author  wdyqxx
 * @version 2017年1月4日 上午11:23:25
 * @explain 此类用于测试UserContactsDao接口的所有方法；
 */
public class TestUserInfoDao {
	
	ClassPathXmlApplicationContext ctx;
	UserInfoDao uiDao;
	
	@Before
	public void init(){
		ctx=new ClassPathXmlApplicationContext(
				"spring-mybatis.xml"
				);
		uiDao=ctx.getBean("userInfoDao",UserInfoDao.class);
	}
	
	@Test
	public void findJobIdTest(){
//		JobPOJO job=new JobPOJO();
//		job.setJobName("工程师");
//		job=uiDao.findJobId(job);
//		System.out.println(job);
	}
	
	@Test
	public void completeUserInfoTest(){
		UserInfoPO uipojo=new UserInfoPO();
		uipojo.setPid(1000001);
		uipojo.setNickname("老王");
		uipojo.setGender(Gender.man.getGender());
		uipojo.setBirthday(new Date(System.currentTimeMillis()));
		uiDao.completeUserInfo(uipojo);
	}
	
	
	
	
	
}
