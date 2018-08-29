package com.lanlian.chatting.controller;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.JSON;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.AccountService;
import com.lanlian.chatting.util.ParameterVerify;
import com.lanlian.chatting.util.RequestSetting;

/**
 * @author wdyqxx
 * @version 2017年1月4日 下午8:12:03
 * @explain 此类为控制层，用于用户获取联系人信息的操作；
 * 
 */
@Controller
@RequestMapping(value = RequestSetting.USER_PARENT, consumes = RequestSetting.CONSUMES, method = RequestMethod.POST, produces = RequestSetting.PRODUCES)
public class UserInfoController extends MyAbstractController {

	@Resource
	AccountService accountService;

	/**
	 * 编辑个人资料
	 * 
	 * @param info
	 * @return UserInfo
	 * @throws Parameter_Exception
	 * @throws InterruptedException
	 */
	@ResponseBody
	@RequestMapping(value = RequestSetting.USER_BODY_INFO)
	public String basicInfo(@RequestParam(value = "uid", required = true) int uid,
			@RequestParam(value = "nickname", required = true) String nickname,
			@RequestParam(value = "address", required = true) String address,
			@RequestParam(value = "signature", required = true) String signature,
			@RequestParam(value = "email", required = true) String email,
			@RequestParam(value = "phone", required = true) String phone,
			@RequestParam(value = "birthday", required = true) Long birthday,
			@RequestParam(value = "gender", required = true) int gender,
			@RequestParam(value = "avatar", required = true) int avatar,
			@RequestParam(value = "job", required = true) int job, HttpServletRequest request,
			HttpServletResponse response) throws Parameter_Exception, InterruptedException {
		LogUtil.info("basic_info-request-uid:" + uid + ",nickname:" + nickname + ",gender:" + gender + ",birthday:"
				+ birthday + ",avatar:" + avatar + ",address:" + address + ",job:" + job + ",signature:" + signature
				+ ",email:" + email + ",phone:" + phone);
		SsoPo ssoPo = getContextSsoPo(request, response);
		if (uid != ssoPo.getUid()) {
			throw new Parameter_Exception(21002);
		}
		// 对参数进行判空与格式校验
		ParameterVerify.verifyUid(ssoPo.getUid());
		ParameterVerify.nicknameVerify(nickname);
		ParameterVerify.genderVerify(gender);
		ParameterVerify.verifyIntegerPositive(avatar, job);
		// ParameterVerify.verifyNull(address,signature);
		if (!email.isEmpty()) {
			ParameterVerify.verifyEmail(email);
		}
		if (!phone.isEmpty()) {
			ParameterVerify.verifyPhone(phone);
		}
		// 对参数对象处理
		AccountPo accountPo = new AccountPo();

		accountPo = basicInfo(accountPo, ssoPo.getUid(), nickname, gender, birthday, avatar, address, job, signature,
				email, phone);
		accountPo.setSource(ssoPo.getSource());
		accountService.accountUserInfo(accountPo);
		return JSON.toJson();
	}

	/**
	 * =========== 私有方法===================
	 */

	/**
	 * 编辑个人资料；
	 * 
	 * @param accountPo
	 * @param uid
	 * @param nickname
	 * @param gender
	 * @param birthday
	 * @param avatar
	 * @param address
	 * @param jobId
	 * @param signature
	 * @param email
	 * @param phone
	 * @return
	 */
	public static AccountPo basicInfo(AccountPo accountPo, int uid, String nickname, int gender, long birthday,
			int avatar, String address, int job, String signature, String email, String phone) {
		accountPo.setUid(uid);
		accountPo.setNickname(nickname);
		accountPo.setGender(gender);
		accountPo.setBirthday(birthday);
		accountPo.setAvatar(avatar);
		accountPo.setProvince(address);
		accountPo.setCity(address);
		accountPo.setUserAddress(address);
		accountPo.setUserJobId(job);
		accountPo.setSignature(signature.trim());
		// accountPo.setSignature(PublicMethod.filterEmoji(signature.trim()));
		accountPo.setEmail(email);
		accountPo.setPhone(phone);
		return accountPo;
	}

}
