package com.lanlian.chatting.controller;

import java.io.UnsupportedEncodingException;
import java.util.List;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.UserLetterService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;
import com.lanlian.chatting.vo.UserLetter;

/**
 * @author wdyqxx
 * @version 2017年1月6日 下午13:12:03
 * @explain 此类为控制层，用于用户发送私信信息的操作；
 */
@Controller
@RequestMapping(value = RequestSetting.LETTER_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class UserLetterController extends MyAbstractController {

	@Resource
	UserLetterService letterService;

	/**
	 * 对好友发送私信
	 * 
	 * @param uid
	 * @param toUid
	 * @param type
	 * @param letter
	 * @return
	 * @throws Parameter_Exception
	 * @throws UnsupportedEncodingException
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.LETTER_BODY_SEND)
	public String sendLetter(@RequestParam(value = "uid", required = true) int uid,
			@PathVariable(value = "touid", required = true) int touid,
			@RequestParam(value = "type", required = true) String type,
			@RequestParam(value = "letter", required = true) String letter, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception, UnsupportedEncodingException {

		LogUtil.info("sendLetter-uid:" + uid + ",touid:" + touid + ",type:" + type + ",letter:" + letter + "}");
		// 进行参数判空校验
		ParameterVerify.verifyNull(type, letter);
		// 进行参数格式校验
		ParameterVerify.verifyUid(uid, touid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		UserLetter userLetter = new UserLetter();
		userLetter.setUid(ssoPo.getUid());
		userLetter.setTouid(touid);
		userLetter.setType(type);
		userLetter.setLetter(encode(letter));
		letterService.sendPrivateLette(userLetter);
		LogUtil.info(JSON.toJson());
		return JSON.toJson();
	}

	/**
	 * 接收指定好友的私信内容
	 * 
	 * @param uid
	 * @param toUid
	 * @return
	 * @throws Parameter_Exception
	 * @throws UnsupportedEncodingException
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.LETTER_BODY_GET)
	public Object getLetter(@RequestParam(value = "uid", required = true) int uid,
			@PathVariable(value = "touid", required = true) int touid, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception, UnsupportedEncodingException {
		LogUtil.info("getLetter-uid:" + uid + ",touid:" + touid);
		// 进行参数格式校验
		ParameterVerify.verifyUid(uid, touid);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (ssoPo.getUid() != uid) {
			return JSON.toJson(21002);
		}
		UserLetter info = new UserLetter();
		info.setUid(touid);
		info.setTouid(ssoPo.getUid());
		List<UserLetter> listnum = letterService.receiveLetter(info);
		LogUtil.info("返回参数{code:10000,listnum:" + listnum + "}");
		return JSON.toJson(listnum);
	}

}
