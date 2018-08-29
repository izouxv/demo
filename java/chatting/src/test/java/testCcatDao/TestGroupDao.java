package testCcatDao;

import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.dao.GroupSettingsDao;
import com.lanlian.chatting.po.GroupAndUserPO;
import com.lanlian.chatting.po.GroupInfoPo;

/** 
* @author  yfq
* @version 2017年3月27日 上午11:31:29
*  xxx.... 
*/
public class TestGroupDao {
	
	ClassPathXmlApplicationContext ctx;
	GroupSettingsDao groupSettingsDao;
	
	@Before
	public void init(){
		ctx = new ClassPathXmlApplicationContext(
				"spring-mybatis.xml"
				);
		groupSettingsDao = ctx.getBean("groupSettingsDao",GroupSettingsDao.class);
	}
	
	@Test
	public void saveGroupSyncTest(){
		GroupInfoPo groupInfoPo = new GroupInfoPo();
		groupInfoPo.setUid(1000001);
		groupInfoPo.setGname("1523");
		groupInfoPo.setAvatar(1);
//		groupSettingsDao.syncGroup(groupInfoPo);
		System.out.println(groupInfoPo);
	}
	
	@Test
	public void saveGroupMemdersTest(){
		GroupAndUserPO groupAndUserPO = new GroupAndUserPO();
		groupAndUserPO.setUid(123456);
		groupAndUserPO.setGid(100000001);
		groupAndUserPO.setCreatTime(new Timestamp(System.currentTimeMillis()));
		
		GroupAndUserPO groupAndUserPO1 = new GroupAndUserPO();
		groupAndUserPO1.setUid(123457);
		groupAndUserPO1.setGid(100000001);
		groupAndUserPO1.setCreatTime(new Timestamp(System.currentTimeMillis()));
		
		List<GroupAndUserPO> list = new ArrayList<>();
		list.add(groupAndUserPO);
		list.add(groupAndUserPO1);
		groupSettingsDao.saveMembers(list);
	}
} 



