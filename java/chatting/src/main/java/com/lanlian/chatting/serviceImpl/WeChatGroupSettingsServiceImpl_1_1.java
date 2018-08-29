package com.lanlian.chatting.serviceImpl;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.List;

import javax.annotation.Resource;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.lanlian.chatting.bo.InviteBo;
import com.lanlian.chatting.bo.MessagePageBo;
import com.lanlian.chatting.dao.WeiChatGroupSettingsDao_1_1;
import com.lanlian.chatting.po.AccountPo;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.LiveChatGidUid;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.AccountClient;
import com.lanlian.chatting.service.WeChatGroupSettingsService;
import com.lanlian.chatting.vo.BundlingGroupVo;
import com.lanlian.chatting.vo.MessageInfoPage;
import com.lanlian.server.redis.RedisClient79;

/**
 * @author  王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年7月14日 下午5:08:33
 * @explain 
 */

@Service("weChatGroupSettingsServiceImpl_1_1")
@Primary
public class WeChatGroupSettingsServiceImpl_1_1 implements WeChatGroupSettingsService {

	@Autowired
	WeiChatGroupSettingsDao_1_1 weiChatGroupSettingsDao_1_1;

	@Autowired
	AccountClient accountClient;
	
	@Resource
	RedisClient79 redisClient;

	/**
	 * 绑定群id
	 * @throws UnsupportedEncodingException 
	 */
	@Transactional
	@Override
	public BundlingGroupVo bundling(InviteBo inviteBo) throws Parameter_Exception, UnsupportedEncodingException {
		LiveChatGidUid liveChatGidUid = new LiveChatGidUid();
		DadaGroupPo dadaGroupPo = new DadaGroupPo();
		dadaGroupPo.setInviteCode(inviteBo.getInviteCode());
		DadaGroupPo dadaGroupPo2 = weiChatGroupSettingsDao_1_1.findGroupInfoByinviteCode(dadaGroupPo);
		if (dadaGroupPo2 == null) {
			// 提示信息：邀请码错误，请查看
			throw new Parameter_Exception(20028);
		}
		liveChatGidUid.setGid(dadaGroupPo2.getGid());
		liveChatGidUid.setCreateTableTime(new Timestamp(System.currentTimeMillis()));
		liveChatGidUid.setModifyTableTime(new Timestamp(System.currentTimeMillis()));
		liveChatGidUid.setUid(inviteBo.getUid());
		liveChatGidUid.setStateTable(1);
//		int count = weiChatGroupSettingsDao.findBundilingGP(liveChatGidUid);
//		if (count > 0) {
//			// 已经绑定过了
//			throw new Parameter_Exception(20029);
//		}
		weiChatGroupSettingsDao_1_1.bundlingGP(liveChatGidUid);
		BundlingGroupVo bundlingGroupVo = new BundlingGroupVo();
		bundlingGroupVo.setGid(dadaGroupPo2.getGid());
		bundlingGroupVo.setGname(URLDecoder.decode(dadaGroupPo2.getGname(), "UTF-8"));
		bundlingGroupVo.setUid(dadaGroupPo2.getUid());
		bundlingGroupVo.setAvatar(dadaGroupPo2.getAvatar());
		bundlingGroupVo.setAnnouncement(dadaGroupPo2.getAnnouncement());
		bundlingGroupVo.setLongitude(dadaGroupPo2.getLongitude());
		bundlingGroupVo.setLatitude(dadaGroupPo2.getLatitude());
		return bundlingGroupVo;
	}

	/**
	 * 查询历史群信息
	 */
	@Transactional
	@Override
	public List<MessageInfoPage> getMessage(MessagePageBo messagePageBo, String source) throws Parameter_Exception {
		int endid = messagePageBo.getEndid();
		int startid = messagePageBo.getStartid();
		if (endid == 0 && startid == 0) {
			startid = weiChatGroupSettingsDao_1_1.findAllMsgCount() - messagePageBo.getCount();
			messagePageBo.setStartid(startid);
			List<MessageInfoPage> list = weiChatGroupSettingsDao_1_1.findMessageInfo(messagePageBo);
			findUserInfoByUidToAccount(list, source);
			return list;
		} 
		if (startid == 0) {
			// 向上查询历史
			List<MessageInfoPage> list = weiChatGroupSettingsDao_1_1.findMessageInfoUp(messagePageBo);
			findUserInfoByUidToAccount(list, source);
			return list;
		}
		if (endid == 0) {
			// 向下查询，新消息
			List<MessageInfoPage> list = weiChatGroupSettingsDao_1_1.findMessageInfoDown(messagePageBo);
			findUserInfoByUidToAccount(list, source);
			return list;
		}
		return null;
	}

	/**
	 * 根据群信息获取uid后account批量查询用户信息
	 * 
	 * @param list
	 * @throws Parameter_Exception
	 */
	@Override
	public void findUserInfoByUidToAccount(List<MessageInfoPage> list, String source) throws Parameter_Exception {
		if (list.isEmpty()) {
			return;
		}
		List<Integer> uids = new ArrayList<>();
		for (MessageInfoPage messageInfoPage : list) {
			uids.add(messageInfoPage.getUid());
		}
		List<AccountPo> accountPos = accountClient.getBatchAllUserInfo(uids, source);
		try {
			for (MessageInfoPage messageInfoPage : list) {
				for (AccountPo accountPo : accountPos) {
					if (messageInfoPage.getUid() == accountPo.getUid()) {
						messageInfoPage.setInfo(URLDecoder.decode(messageInfoPage.getInfo(), "UTF-8"));
						messageInfoPage.setAvatar(accountPo.getAvatar());
						messageInfoPage.setNickname(accountPo.getNickname());
					}
				}
			}
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(10002);
		}
	}

}
