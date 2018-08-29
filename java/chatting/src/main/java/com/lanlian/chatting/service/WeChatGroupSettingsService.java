package com.lanlian.chatting.service;

import java.io.UnsupportedEncodingException;
import java.util.List;

import com.lanlian.chatting.bo.InviteBo;
import com.lanlian.chatting.bo.MessagePageBo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.vo.BundlingGroupVo;
import com.lanlian.chatting.vo.MessageInfoPage;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月19日 上午10:24:38
 * @explain
 */

public interface WeChatGroupSettingsService {
	
	BundlingGroupVo bundling(InviteBo inviteBo) throws Parameter_Exception, UnsupportedEncodingException;
	
	List<MessageInfoPage> getMessage(MessagePageBo messagePageBo, String source) throws Parameter_Exception;
	
	/**
	 * 根据群信息获取uid后account批量查询用户信息
	 * 
	 * @param list
	 * @throws Parameter_Exception
	 */
	void findUserInfoByUidToAccount(List<MessageInfoPage> list, String source) throws Parameter_Exception;
}
