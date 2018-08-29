
package testCcatService;

import org.junit.Before;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.service.GroupSettingsService;

/** 
* @author  yfq
* @version 2017年3月23日 下午6:38:57
* @explain 此类用于测试UpgradeGroupService的所有方法
*  
*/

public class TestUpgradeGroupService {
	
	ClassPathXmlApplicationContext ctx;
	GroupSettingsService ugService;
	
	@Before
	public void init(){
		ctx = new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-controller.xml",
				"spring-mybatis.xml"
				);
		ugService = ctx.getBean("upgradeGroupService",GroupSettingsService.class);
	}

	
}


