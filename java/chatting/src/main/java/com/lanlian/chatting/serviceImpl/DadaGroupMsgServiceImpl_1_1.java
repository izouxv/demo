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

import org.springframework.context.annotation.Primary;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.controller.MyAbstractController;
import com.lanlian.chatting.dao.DadaGroupMsgDao_1_1;
import com.lanlian.chatting.dao.GroupSettingsDao_1_1;
import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.po.DadaGroupMsgPo;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.DadaGroupMsgService;
import com.lanlian.chatting.service.WeChatGroupSettingsService;
import com.lanlian.chatting.vo.GroupMessageInfo;
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

@Service(value = "dadaGroupMsgServiceImpl_1_1")
@Primary
public class DadaGroupMsgServiceImpl_1_1 extends MyAbstractController implements DadaGroupMsgService {

	@Resource
	DadaGroupMsgDao_1_1 dadaGroupMsgDao_1_1;

	@Resource
	GroupSettingsDao_1_1 groupSettingsDao_1_1;

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
		dadaGroupPo = groupSettingsDao_1_1.findDadaGroupSwitchingUser(dadaGroupPo);
		return dadaGroupPo;
	}
	
	@Transactional
	@Override
	public void saveGroupMessage(JSONArray jsonArray, SsoPo ssoPo, DadaGroupPo dadaGroupPo)
			throws Parameter_Exception {
		List<DadaGroupMsgPo> msgPoList = new ArrayList<>();
		List<GroupMessageInfo> msgInfoList = new ArrayList<>();
		try {
			DadaGroupMsgPo dadaGroupMsgPo = null;
			GroupMessageInfo groupMessageInfo = null;
			JSONObject jsonObject = null;
			String info = "";
			int size = jsonArray.size();
			for (int i = 0; i < size; i++) {
				dadaGroupMsgPo = new DadaGroupMsgPo();
				groupMessageInfo = new GroupMessageInfo();
				jsonObject = jsonArray.getJSONObject(i);
				dadaGroupMsgPo.setUpuid(ssoPo.getUid());
				dadaGroupMsgPo.setGid(dadaGroupPo.getGid());
				dadaGroupMsgPo.setUid(jsonObject.getIntValue("uid"));
				dadaGroupMsgPo.setType(jsonObject.getIntValue("type"));
				info = jsonObject.get("info").toString();
				groupMessageInfo.setMessageInfo(URLEncoder.encode(info,"UTF-8"));
				dadaGroupMsgPo.setInfo(URLEncoder.encode(info, "UTF-8"));
				dadaGroupMsgPo.setSendTime(new Timestamp(jsonObject.getLongValue("sendTime")));
				dadaGroupMsgPo.setCreateTime(getTime());
				dadaGroupMsgPo.setUpdateTime(getTime());
				dadaGroupMsgPo.setStateTable(1);
				msgPoList.add(dadaGroupMsgPo);
				msgInfoList.add(groupMessageInfo);
			}
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(21002);
		}
		if (msgPoList.isEmpty()) {
			throw new Parameter_Exception(21002);
		}
		dadaGroupMsgDao_1_1.saveGroupMsgInfoId(msgInfoList);
		sendGroupMessageConver(msgPoList, msgInfoList);
//		System.out.println("开始插入msgPoList对象");
		dadaGroupMsgDao_1_1.saveGroupMessage(msgPoList);
		// System.out.println("list:" + list);
		// 判断该群有无订阅
		boolean flag = redisClient79.findSubscribeGroupId(dadaGroupPo.getGid());
		// 有订阅则存入list
		if (flag) {
			setRedis(dadaGroupPo, msgPoList, ssoPo.getSource());
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
			// 向list中存入聊天信息并通知
			redisClient80.setBundlingGroupInfo(dadaGroupPo.getGid().toString(), messageInfoPages);
		} catch (Exception e) {
			LogUtil.error(e);
		}
	}
	
	/**
	 * 将GroupMessageInfo对象的messageInfoId赋值给DadaGroupMsgPo对象的mid
	 * @param list   DadaGroupMsgPo对象
	 * @param list2  GroupMessageInfo对象
	 * @return
	 * @throws Parameter_Exception
	 */
	private List<DadaGroupMsgPo> sendGroupMessageConver(List<DadaGroupMsgPo> list, List<GroupMessageInfo> list2 ) throws Parameter_Exception {
		try {
			//大小不相等数据错误
			if (list.size() != list2.size()) {
				throw new Parameter_Exception(10002);
			}
			int size = list.size();
			for (int i = 0; i < size; i++) {
				list.get(i).setMid(list2.get(i).getMessageInfoId());
//				System.out.println("sendGroupMessageConver:" + list2.get(i).getMessageInfoId());
			}
		} catch (Exception e) {
			throw new Parameter_Exception(10002);
		}
		return list;
	}

//	/**
//	 * json转为Bean
//	 * @param json
//	 * @return
//	 */
//	private DadaGroupMsgPo jsonToBean(String json) {
//		return JSON.parseObject(json, DadaGroupMsgPo.class);
//	}
}
