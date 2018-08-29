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
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import com.lanlian.chatting.controller.MyAbstractController;
import com.lanlian.chatting.dao.GroupSettingsDao_1_1;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.GroupAndUserPO;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.DadaGroupService;
import com.lanlian.server.redis.RedisClient79;

/**
 * @Title LiveChatGroupServiceImpl.java
 * @Package com.lanlian.chatting.serviceImpl
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午5:03:34
 * @explain
 */
@Service("dadaGroupServiceImpl_1_1")
@Primary
public class DadaGroupServiceImpl_1_1 extends MyAbstractController implements DadaGroupService {

	@Resource
	GroupSettingsDao_1_1 groupSettingsDao_1_1;
	
	@Resource
	SsoClient ssoClient;
	
	@Resource
	RedisClient79 redisClient79;
	
	/**
	 * 查询群id是否存在
	 * 
	 * @param dadaGroupPo
	 * @return
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public DadaGroupPo findDadaGroup(DadaGroupPo dadaGroupPo) {
		dadaGroupPo = groupSettingsDao_1_1.findDadaGroup(dadaGroupPo);
		return dadaGroupPo;
	}

	/**
	 * 修改实时上报群开关
	 * 
	 * @param dadaGroupPo
	 */
	@Transactional
	@Override
	public DadaGroupPo updateDadaGroup(DadaGroupPo dadaGroupPo) {
		groupSettingsDao_1_1.updateDadaGroup(dadaGroupPo);
		DadaGroupPo dadaGroupPo2 = new DadaGroupPo();
		dadaGroupPo2.setGid(dadaGroupPo.getGid());
		return groupSettingsDao_1_1.findDadaGroup(dadaGroupPo2);
	}

	/**
	 * 创建实时上报群与开关
	 * 
	 * @param dadaGroupPo
	 * @throws Parameter_Exception
	 */
	@Transactional
	@Override
	public DadaGroupPo saveDadaGroup(DadaGroupPo dadaGroupPo, List<Integer> list, String source) throws Parameter_Exception {
		int flag1 = list.size();
		list = ssoClient.getBatchSsoInfos(list, source);
		int flag2 = list.size();
		if (flag1 != flag2) {
			throw new Parameter_Exception(20009);
		}
		try {
			dadaGroupPo.setGname(URLEncoder.encode(dadaGroupPo.getGname(), "UTF-8"));
		} catch (UnsupportedEncodingException e) {
			throw new Parameter_Exception(21002);
		}
		DadaGroupPo dadaGroupPo2 = new DadaGroupPo();
		dadaGroupPo2.setUid(dadaGroupPo.getUid());
		dadaGroupPo2.setUpid(dadaGroupPo.getUid());
		dadaGroupPo2.setGroupState(dadaGroupPo.getGroupState());
		dadaGroupPo2.setCreateTime(dadaGroupPo.getCreateTime());
		dadaGroupPo2.setModifyTime(dadaGroupPo.getModifyTime());
		dadaGroupPo2.setDataState(dadaGroupPo.getDataState());
		// 通过redis获取一个新的临时群id
		Integer newgid = redisClient79.upgradeOneTG(dadaGroupPo.getGid(), dadaGroupPo.getUid());
		// 记录群上传开关
		redisClient79.setGidUidSwitch(dadaGroupPo.getGid(), dadaGroupPo.getUid(), dadaGroupPo.getGroupState());
		try {
			// 添加群资料
			groupSettingsDao_1_1.saveDadaGroup(dadaGroupPo);
			//添加群成员
			List<GroupAndUserPO> list2 = new ArrayList<>();
			Timestamp timestamp = new Timestamp(System.currentTimeMillis());
			GroupAndUserPO groupAndUserPO = null;
			for (Integer uid : list) {
				groupAndUserPO = new GroupAndUserPO();
				groupAndUserPO.setUid(uid);
				groupAndUserPO.setGid(dadaGroupPo.getGid());
				groupAndUserPO.setCreatTime(timestamp);
				list2.add(groupAndUserPO);
			}
			// 添加群成员
			groupSettingsDao_1_1.saveMembers(list2);
			dadaGroupPo2.setGid(dadaGroupPo.getGid());
			//添加群邀请码
			groupSettingsDao_1_1.saveDadaGroupCode(dadaGroupPo);
			// 添加群上报开关
			groupSettingsDao_1_1.saveDadaGroupSwitching(dadaGroupPo2);
			//将gid存入 dadaGroupPo2 对象
			dadaGroupPo2.setNewgid(newgid);
			dadaGroupPo2.setInviteCode(dadaGroupPo.getInviteCode());
			return dadaGroupPo2;
		} catch (Exception e) {
			redisClient79.compensateRedis(dadaGroupPo.getUid(), dadaGroupPo.getGid(), newgid);
			redisClient79.delGIdUid(dadaGroupPo.getGid());
			throw e;
		}
	}
	
	/**
	 * 创建实时上报群开关
	 * 
	 * @param dadaGroupPo
	 * @throws Parameter_Exception
	 */
	@Transactional
	@Override
	public void saveDadaGroupSwitching(DadaGroupPo dadaGroupPo) throws Parameter_Exception {
		dadaGroupPo.setGroupState(1);
		dadaGroupPo.setCreateTime(getTime());
		dadaGroupPo.setModifyTime(getTime());
		dadaGroupPo.setDataState(1);
		groupSettingsDao_1_1.saveDadaGroupSwitching(dadaGroupPo);
	}

	/**
	 * 存入邀请码
	 * 
	 * @param dadaGroupPo
	 */
	@Transactional
	@Override
	public void saveDadaGroupCode(DadaGroupPo dadaGroupPo) {
		int aa =0;
		for (int i = 0; i < 3; i++) {
			aa = groupSettingsDao_1_1.saveDadaGroupCode(dadaGroupPo);
			if (aa == 1) {
				break;
			}
		}
	}

	/**
	 * 查询实时上报群与开关
	 * 
	 * @param dadaGroupPo
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public DadaGroupPo findDadaGroupSwitchingUser(DadaGroupPo dadaGroupPo) throws Parameter_Exception {
		dadaGroupPo = groupSettingsDao_1_1.findDadaGroupSwitchingUser(dadaGroupPo);
		return dadaGroupPo;
	}

}
