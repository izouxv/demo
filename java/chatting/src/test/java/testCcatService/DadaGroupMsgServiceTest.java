package testCcatService;

import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.alibaba.fastjson.JSONArray;
import com.lanlian.chatting.bo.MessagePageBo;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.DadaGroupMsgService;
import com.lanlian.chatting.service.WeChatGroupSettingsService;
import com.lanlian.chatting.vo.MessageInfoPage;

/** 
 * @Title DadaGroupMsgServiceTest.java
 * @Package testCcatService
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午4:36:12
 * @explain 信息实时上报群设置接口业务层
 */
public class DadaGroupMsgServiceTest {
	
	ClassPathXmlApplicationContext ctx;
	
	DadaGroupMsgService dadaGroupMsgService;
	
	WeChatGroupSettingsService weChatGroupSettingsService;
	
	@Before
	public void init(){
		ctx=new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-mybatis.xml"
				);
		dadaGroupMsgService=ctx.getBean("dadaGroupMsgServiceImpl",DadaGroupMsgService.class);
		weChatGroupSettingsService=ctx.getBean("weChatGroupSettingsService",WeChatGroupSettingsService.class);
	}
	
	@Test
	public void findMsgTest() throws Parameter_Exception{
		MessagePageBo messagePageBo = new MessagePageBo();
		messagePageBo.setGid(100000001);
		messagePageBo.setCount(100);
		List<MessageInfoPage> list = weChatGroupSettingsService.getMessage(messagePageBo,"");
		System.out.println(list);
	}
	
	@Test
	public void dadaGroupMsgTest() throws Parameter_Exception{
		DadaGroupPo dadaGroupPo = new DadaGroupPo();
		dadaGroupPo.setGid(100000001);
		SsoPo ssoPo = new SsoPo();
		ssoPo.setUid(1000001);
		String json = "[{\"uid\":\"1000001\","
				+ "\"type\":\"1\",\"info\":\"你好啊啊啊啊啊\",\"sendTime\":\"1234567892132\"},"
				+ "{\"uid\":\"1000002\","
				+ "\"type\":\"1\",\"info\":\"你好啊啊啊啊啊\",\"sendTime\":\"1234567893132\"}]";
		JSONArray jsonArray = JSONArray.parseArray(json);
		System.err.println(jsonArray);
		dadaGroupMsgService.saveGroupMessage(jsonArray, ssoPo, dadaGroupPo);
	}
	
}
