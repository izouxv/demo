/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.chatting.serviceImpl;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.List;

import javax.annotation.Resource;

import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.controller.MyAbstractController;
import com.lanlian.chatting.dao.DadaGroupMsgDao;
import com.lanlian.chatting.dao.GroupSettingsDao;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.DadaGroupMsgPo;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.DadaGroupMsgService;
import com.lanlian.chatting.service.WeChatGroupSettingsService;
import com.lanlian.chatting.vo.MessageInfoPage;
import com.lanlian.server.redis.RedisClient79;
import com.lanlian.server.redis.RedisClient80;

/**
 * @Title DadaGroupMsgServiceImpl.java
 * @Package com.lanlian.chatting.service
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午4:36:12
 * @explain 信息实时上报群设置接口业务层
 */

@Service(value = "dadaGroupMsgServiceImpl")
public class DadaGroupMsgServiceImpl extends MyAbstractController implements DadaGroupMsgService {

	@Resource
	DadaGroupMsgDao dadaGroupMsgDao;

	@Resource
	GroupSettingsDao groupSettingsDao;

	@Resource
	RedisClient79 redisClient79;
	
	@Resource
	RedisClient80 redisClient80;

	@Resource
	WeChatGroupSettingsService weChatGroupSettingsService;

	/**
	 * 查询打开实时群信息上报开关的用户信息
	 * 
	 * @return
	 */
	@Override
	public DadaGroupPo findDadaGroupSwitchingUser(DadaGroupPo dadaGroupPo) {
		dadaGroupPo = groupSettingsDao.findDadaGroupSwitchingUser(dadaGroupPo);
		return dadaGroupPo;
	}

	/**
	 * 将实时群信息存储
	 * 
	 * @throws Parameter_Exception
	 */
	@Transactional
	@Override
	public void saveGroupMessage(JSONArray jsonArray, SsoPo ssoPo, DadaGroupPo dadaGroupPo) throws Parameter_Exception {
		List<DadaGroupMsgPo> list = new ArrayList<>();
		try {
			for (int i = 0; i < jsonArray.size(); i++) {
				JSONObject jsonObject = jsonArray.getJSONObject(i);
				DadaGroupMsgPo dadaGroupMsgPo = new DadaGroupMsgPo();
				dadaGroupMsgPo.setUpuid(ssoPo.getUid());
				dadaGroupMsgPo.setGid(dadaGroupPo.getGid());
				dadaGroupMsgPo.setUid(Integer.parseInt(jsonObject.get("uid").toString()));
				dadaGroupMsgPo.setType(Integer.parseInt(jsonObject.get("type").toString()));
				String info = jsonObject.get("info").toString().trim();
				if (info.isEmpty()) {
					throw new Parameter_Exception(21002);
				}
				dadaGroupMsgPo.setInfo(URLEncoder.encode(info, "UTF-8"));
				dadaGroupMsgPo.setSendTime(new Timestamp(Long.parseLong(jsonObject.get("sendTime").toString())));
				dadaGroupMsgPo.setCreateTime(getTime());
				dadaGroupMsgPo.setUpdateTime(getTime());
				dadaGroupMsgPo.setStateTable(1);
				list.add(dadaGroupMsgPo);
			}
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(21002);
		}
		if (list.isEmpty()) {
			throw new Parameter_Exception(21002);
		}
		dadaGroupMsgDao.saveGroupMsg(list);
		// System.out.println("list:" + list);
		//判断该群有无订阅
		boolean flag = redisClient79.findSubscribeGroupId(dadaGroupPo.getGid());
		//有订阅则存入list
		if (flag) {
			setRedis(dadaGroupPo, list, ssoPo.getSource());
		}
	}

	/**
	 * 
	 * @param dadaGroupPo
	 * @param list
	 * @throws Parameter_Exception
	 */
	@Async
	private void setRedis(DadaGroupPo dadaGroupPo, List<DadaGroupMsgPo> list, String source) {
		try {
			List<MessageInfoPage> messageInfoPages = new ArrayList<>();
			MessageInfoPage messageInfoPage = null;
			for (DadaGroupMsgPo dadaGroupMsgPo : list) {
				messageInfoPage = new MessageInfoPage();
				messageInfoPage.setMid(dadaGroupMsgPo.getMid());
				messageInfoPage.setUid(dadaGroupMsgPo.getUid());
				messageInfoPage.setTime(dadaGroupMsgPo.getSendTime());
				messageInfoPage.setType(dadaGroupMsgPo.getType());
				messageInfoPage.setInfo(dadaGroupMsgPo.getInfo());
				messageInfoPages.add(messageInfoPage);
			}
			// System.out.println("messageInfoPages:" + messageInfoPages);
			// 查询聊天记录中的用户信息
			weChatGroupSettingsService.findUserInfoByUidToAccount(messageInfoPages, source);
			// 向list中存入聊天信息
			redisClient80.setBundlingGroupInfo(dadaGroupPo.getGid().toString(), messageInfoPages);
		} catch (Exception e) {
			LogUtil.error(e);
		}
	}

}
