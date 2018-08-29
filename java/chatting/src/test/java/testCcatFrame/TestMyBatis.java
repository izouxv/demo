package testCcatFrame;

import javax.sql.DataSource;

import org.apache.ibatis.session.SqlSessionFactory;
import org.junit.Before;
import org.junit.Test;
import org.mybatis.spring.mapper.MapperScannerConfigurer;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class TestMyBatis {
	
	ClassPathXmlApplicationContext ctx;
	@Before
	public void init(){
		ctx=new ClassPathXmlApplicationContext(
				"spring-service.xml",
				"spring-controller.xml",
				"spring-mybatis.xml"
				);
	}
	
	 @Test
    public void testDataSource(){
        DataSource ds = ctx.getBean(
            "dataSource", DataSource.class);
        System.out.println(ds); 
    }
	
	@Test
	public void testSqlSessionFactory(){
	     SqlSessionFactory factory=
	         ctx.getBean("sqlSessionFactory",
	         SqlSessionFactory.class);
	     System.out.println(factory);
	}
	
	 @Test
	 public void testMapperScanner(){
	     MapperScannerConfigurer scanner=
	         ctx.getBean("mapperScanner",
	         MapperScannerConfigurer.class);
	     System.out.println(scanner);
	 }
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 
	 

	
}
