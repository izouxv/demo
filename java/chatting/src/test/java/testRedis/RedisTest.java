/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package testRedis;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.server.redis.RedisClient79;

/** 
 * @Title RedisTest.java
 * @Package testRedis
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月14日 上午11:26:38
 * @explain 
 */

public class RedisTest {
	
	ClassPathXmlApplicationContext ctx;
//	JedisPool jedisPool;
	RedisClient79 checkCode;
	
	@Before
	public void init(){
		ctx = new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-mybatis.xml"
				);
//		jedisPool = ctx.getBean("jedisPool",JedisPool.class);
		checkCode = ctx.getBean("checkCode", RedisClient79.class);
	}
	
	@Test
	public void getRedis() throws Parameter_Exception{
		
		checkCode.checkCode("17600117962", "688821");
		
	}
	
	
}

