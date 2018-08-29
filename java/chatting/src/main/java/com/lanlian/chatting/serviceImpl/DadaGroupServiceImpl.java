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
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import com.lanlian.chatting.controller.MyAbstractController;
import com.lanlian.chatting.dao.GroupSettingsDao;
import com.lanlian.chatting.po.DadaGroupPo;
import com.lanlian.chatting.po.GroupAndUserPO;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.rpc.SsoClient;
import com.lanlian.chatting.service.DadaGroupService;

/**
 * @Title LiveChatGroupServiceImpl.java
 * @Package com.lanlian.chatting.serviceImpl
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月30日 下午5:03:34
 * @explain
 */
@Service("dadaGroupServiceImpl")
public class DadaGroupServiceImpl extends MyAbstractController implements DadaGroupService {

	@Resource
	GroupSettingsDao groupSettingsDao;
	
	@Resource
	SsoClient ssoClient;

	/**
	 * 查询群id是否存在
	 * 
	 * @param dadaGroupPo
	 * @return
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public DadaGroupPo findDadaGroup(DadaGroupPo dadaGroupPo) {
		dadaGroupPo = groupSettingsDao.findDadaGroup(dadaGroupPo);
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
		groupSettingsDao.updateDadaGroup(dadaGroupPo);
		DadaGroupPo dadaGroupPo2 = new DadaGroupPo();
		dadaGroupPo2.setGid(dadaGroupPo.getGid());
		return groupSettingsDao.findDadaGroup(dadaGroupPo2);
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
		groupSettingsDao.saveDadaGroup(dadaGroupPo);
		if (dadaGroupPo.getGid() <= 100000000) {
			throw new Parameter_Exception(10002);
		}
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
		groupSettingsDao.saveMembers(list2);
		dadaGroupPo2.setGid(dadaGroupPo.getGid());
		groupSettingsDao.saveDadaGroupSwitching(dadaGroupPo2);
		return dadaGroupPo2;
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
		groupSettingsDao.saveDadaGroupSwitching(dadaGroupPo);
	}

	/**
	 * 存入邀请码
	 * 
	 * @param dadaGroupPo
	 */
	@Async
	@Transactional
	@Override
	public void saveDadaGroupCode(DadaGroupPo dadaGroupPo) {
		int aa =0;
		for (int i = 0; i < 3; i++) {
			aa = groupSettingsDao.saveDadaGroupCode(dadaGroupPo);
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
		dadaGroupPo = groupSettingsDao.findDadaGroupSwitchingUser(dadaGroupPo);
		return dadaGroupPo;
	}

}
