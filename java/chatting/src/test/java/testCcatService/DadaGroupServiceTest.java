package testCcatService;


import java.sql.Timestamp;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.DadaGroupService;

/** 
 * @Title DadaGroupServiceTest.java
 * @Package testCcatService
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午4:36:12
 * @explain 信息实时上报群设置接口业务层
 */
public class DadaGroupServiceTest {
	
	ClassPathXmlApplicationContext ctx;
	DadaGroupService dadaGroupService;
	
	@Before
	public void init(){
		ctx=new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-mybatis.xml"
//				,"spring-controller.xml"
				);
		dadaGroupService=ctx.getBean("dadaGroupServiceImpl",DadaGroupService.class);
	}
	
	
	@Test
	public void findDadaGroupTest() throws Parameter_Exception{
		DadaGroupPo dadaGroupPo = new DadaGroupPo();
		dadaGroupPo.setGid(100000008);
		dadaGroupPo = dadaGroupService.findDadaGroup(dadaGroupPo);
		System.out.println(dadaGroupPo);
	}
	
	@Test
	public void saveDadaGroupTest() throws Parameter_Exception {
		Timestamp timestamp = new Timestamp(System.currentTimeMillis());
		DadaGroupPo dadaGroupPo = new DadaGroupPo();
		dadaGroupPo.setUid(1000001);
		dadaGroupPo.setGname("蓝涟");
		dadaGroupPo.setAvatar(1);
		dadaGroupPo.setAnnouncement("");
		dadaGroupPo.setCreateTime(timestamp);
		dadaGroupPo.setLongitude(0.0);
		dadaGroupPo.setLatitude(0.0);
		dadaGroupPo.setGroupState(0);
		dadaGroupPo.setInviteCode("a23d132a");
		dadaGroupPo.setDataCreateTime(timestamp);
		dadaGroupPo.setModifyTime(timestamp);
		dadaGroupPo.setDataState(1);
		dadaGroupPo = dadaGroupService.saveDadaGroup(dadaGroupPo, null,"");
		System.out.println("dadaGroupPo:" + dadaGroupPo);
	}
	
	
	
	
	
}
