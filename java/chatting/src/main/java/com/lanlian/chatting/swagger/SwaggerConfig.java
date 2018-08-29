/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.swagger;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.mangofactory.swagger.configuration.SpringSwaggerConfig;
import com.mangofactory.swagger.models.dto.ApiInfo;
import com.mangofactory.swagger.plugin.EnableSwagger;
import com.mangofactory.swagger.plugin.SwaggerSpringMvcPlugin;

/**
 * @Title SwaggerConfig.java
 * @Package com.lanlian.chatting.swagger
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月27日 上午10:19:41
 * @explain
 */

@Configuration
@EnableSwagger
//@ComponentScan("com.lanlian.chatting.controller")
public class SwaggerConfig {

	private SpringSwaggerConfig springSwaggerConfig;

	/**
	 * 自动装配的springSwaggerConfig
	 * 
	 * @param springSwaggerConfig
	 */
	@Autowired
	public void setSpringSwaggerConfig(SpringSwaggerConfig springSwaggerConfig) {
		this.springSwaggerConfig = springSwaggerConfig;
	}

	/**
	 * 每一个SwaggerSpringMvcPlugin bean是由swagger-mvc提供框架，允许多个swagger组i.e，即相同的代码基
	 */
	@Bean
	public SwaggerSpringMvcPlugin customImplementation() {
		return new SwaggerSpringMvcPlugin(this.springSwaggerConfig).apiInfo(apiInfo()).includePatterns(".*?");
	}

	private ApiInfo apiInfo() {
		ApiInfo apiInfo = new ApiInfo(
				"我的程序 API 标题", 
				"我的程序 API 描述", 
				"我的程序 API 服务条件", 
				"我的程序 API 邮箱", 
				"我的程序 API 允许类型",
				"我的程序 API 应用URI");
		return apiInfo;
	}

}
