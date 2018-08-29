package com.lanlian.chatting.controller;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Scope;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.server.redis.RedisClient79;
import com.wordnik.swagger.annotations.Api;
import com.wordnik.swagger.annotations.ApiOperation;
import com.wordnik.swagger.annotations.ApiParam;

//@RequestMapping(value = "/v1.0/demo", produces = "application/json;charset=UTF-8")
@RestController
@Scope(value = "prototype")
@Api(value = "test的swaggerAPI", description = "demo", produces = MediaType.APPLICATION_JSON_VALUE)
public class Demo extends MyAbstractController {

	@Autowired
	RedisClient79 checkCode;

	/**
	 * demotest
	 */
	@RequestMapping(value = "/demo/demo")
	@ApiOperation(value = "demo Method", notes = "测试用例方法", httpMethod = "POST", produces = MediaType.APPLICATION_JSON_VALUE)
	public Object hello(@ApiParam(value = "路径后缀aaa", required = true) @RequestParam("aaa") String aaa,
			@RequestParam(value = "uid", required = true) int uid1, HttpServletRequest req, HttpServletResponse res)
			throws Parameter_Exception, InterruptedException {
		System.out.println("aaa:" + aaa);
		return new JSON<>();
	}

	// @Autowired
	// private Meter requestMeter;
	//
	// @Autowired
	// private Histogram responseSizes;
	//
	// @Autowired
	// private Counter pendingJobs;
	//
	// @Autowired
	// private Timer responses;
	//
	// @ResponseBody
	// @RequestMapping("/hello")
	// public String helloWorld() {
	// System.err.println("start------------");
	//
	// requestMeter.mark();
	//
	// pendingJobs.inc();
	//
	// responseSizes.update(new Random().nextInt(10));
	//
	// final Timer.Context context = responses.time();
	// try {
	// return "Hello World";
	// } finally {
	// context.stop();
	// }
	// }

}