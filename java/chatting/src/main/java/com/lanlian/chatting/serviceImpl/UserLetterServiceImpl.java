package com.lanlian.chatting.serviceImpl;

import java.io.UnsupportedEncodingException;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import com.lanlian.chatting.controller.MyAbstractController;
import com.lanlian.chatting.dao.UserMessageDao;
import com.lanlian.chatting.po.PrivateMessageInfoPO;
import com.lanlian.chatting.po.PrivateMessagePO;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.UserLetterService;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.chatting.vo.UserLetter;
import com.lanlian.chatting.vo.UserMessageNum;

import net.sf.json.JSONArray;

/**
 * @author wdyqxx
 * @version 2017年1月4日 下午2:27:13
 * @explain 此类用于实现用户发送、接收私信的操作；调用接口UserMessageDao；
 *          1.sendPrivateLette：对于发送私信功能的业务处理； 传入参数：UserMessageInfo类型的用户永久pid；
 *          返回参数：无； 2.receiveLetteNum：对于接收私信数量功能的业务处理；
 *          传入参数：UserMessageNum类型的用户永久pid； 返回参数：UserMessageNum类型的List集合；
 *          3.receivePrivateMessageInfo：对于接收私信内容功能的业务处理；
 *          传入参数：UserMessageInfo类型的用户与联系人的永久pid； 返回参数：UserMessageInfo类型的List集合；
 */
@Service("userMessageService")
public class UserLetterServiceImpl extends MyAbstractController implements UserLetterService {

	@Autowired
	UserMessageDao privateDao;

	/**
	 * 业务层：发送私信
	 * 
	 * @throws Parameter_Exception
	 */
	@Transactional(rollbackFor = Exception.class)
	@Override
	public void sendPrivateLette(UserLetter letter) throws Parameter_Exception {

		PrivateMessageInfoPO pmipojo = new PrivateMessageInfoPO();
		pmipojo.setMessageInfo(letter.getLetter());
		privateDao.sendPrivateLetteInfoId(pmipojo);
		if (pmipojo == null || pmipojo.getMessageInfoId() <= 0) {
			// 插入私信内容失败
			throw new Parameter_Exception(10002);
		}
		/**
		 * 转换数据格式，存储数据；
		 */
		PrivateMessagePO pmpojo = sendMessageConver(letter, pmipojo);
		int send = privateDao.sendPrivateLette(pmpojo);
		if (send != 1) {
			// 插入私信记录失败
			throw new Parameter_Exception(10002);
		}
	}

	/**
	 * 业务层：接收私信数量；
	 * 
	 * @throws Parameter_Exception
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public JSONArray receiveLetteNum(UserMessageNum opid) throws Parameter_Exception {
		// 返回list集合
		PrivateMessagePO pmpojo = new PrivateMessagePO();
		pmpojo.setOuid(opid.getOpid());
		List<PrivateMessagePO> pojoList = privateDao.receiveLetteById(pmpojo);
//		System.out.println("pojoList:" + pojoList);
		if (pojoList.size() == 0) {
			// 无数据
			throw new Parameter_Exception(20023);
		}
		// 将pojoList集合中的数据遍历出来，统计数量，放入listnum中；
		// System.out.println("pojoList:"+pojoList);
		Map<String, Integer> map = letteNumCount(pojoList);
//		System.out.println("map:" + map);
		JSONArray jsonArray = PublicMethod.mapInt_Json(map);
//		System.out.println("jsonArray:" + jsonArray);
		return jsonArray;
	}

	/**
	 * 业务层：用户接收通知内容；
	 * 
	 * @throws Parameter_Exception
	 * @throws UnsupportedEncodingException 
	 */
	@Transactional(propagation = Propagation.SUPPORTS, readOnly = true)
	@Override
	public List<UserLetter> receiveLetter(UserLetter letter) throws Parameter_Exception, UnsupportedEncodingException {
		List<UserLetter> umlist = new ArrayList<>();
		// 查询数据库的数据，返回list集合的pmpojo对象；
		PrivateMessagePO pmpojo = new PrivateMessagePO();
		pmpojo.setUid(letter.getUid());
		pmpojo.setOuid(letter.getTouid());
		List<PrivateMessagePO> pmpos = privateDao.receiveLetteUidById(pmpojo);
		if (pmpos.size() == 0) {
			return umlist;
		}
		pmpojo.setStatus(2);
		pmpojo.setModifyTime(new Timestamp(System.currentTimeMillis()));
		int as = privateDao.modifyLetteStatus(pmpojo);
		if (as == 0) {
			// 修改私信状态失败
			throw new Parameter_Exception(10002);
		}
		// 遍历该list集合，将得到的数据放入umList集合中，返回给控制层；
		List<Long> ids = new ArrayList<>();
		for (PrivateMessagePO pmpo : pmpos) {
			// 将pmipo做为参数,查询数据库的数据；
			ids.add(Long.parseLong(pmpo.getMessageInfoId()));
		}
		List<PrivateMessageInfoPO> pms = privateDao.receiveLetteInfo(ids);
		if (pms.isEmpty()) {
			return umlist;
		}
		// 调用ConverInfo的letteConverInfo方法：将pmipo对象的数据，转换为umInfo对象的数据，放入list集合中；
		letteConverInfo(letter, umlist, pmpos, pms);
		return umlist;
	}

	/**
	 * 私有方法=========================================================
	 */
	/**
	 * UserLetter -> PrivateMessageInfoPO 将信息转换对象；
	 * 
	 * @param letter
	 * @param pmipojo
	 * @return
	 */
	private PrivateMessagePO sendMessageConver(UserLetter letter, PrivateMessageInfoPO pmipojo) {
		PrivateMessagePO pmpojo = new PrivateMessagePO();
		pmpojo.setUid(letter.getUid());
		pmpojo.setOuid(letter.getTouid());
		pmpojo.setType(letter.getType());
		pmpojo.setMessageInfoId(Long.toString(pmipojo.getMessageInfoId()));
		Timestamp now = new Timestamp(System.currentTimeMillis());
		pmpojo.setCreateTime(now);
		pmpojo.setModifyTime(now);
		return pmpojo;
	}

	/**
	 * 计算每个联系人给自己发送私信的未读数量；
	 * 
	 * @param opid
	 * @param map
	 * @param listnum
	 */
	private Map<String, Integer> letteNumCount(List<PrivateMessagePO> poList) {
		Map<String, Integer> map = new HashMap<>();
		String uid = null;
		for (PrivateMessagePO pmpo : poList) {
			uid = pmpo.getUid().toString();
			Integer value = map.get(pmpo.getUid().toString());
			if (value == null) {
				map.put(uid, 1);
			} else {				
				map.put(uid, map.get(uid) + 1);
			}
		}
		return map;
	}

	/**
	 * UserMessageServiceImpl的receivePrivateMessageInfo方法调用； 将数据信息转换对象；
	 * 
	 * @param pids
	 * @param umlist
	 * @param privateId
	 * @param pmipojo
	 * @throws UnsupportedEncodingException 
	 */
	private void letteConverInfo(UserLetter letter, List<UserLetter> umlist, List<PrivateMessagePO> pms,
			List<PrivateMessageInfoPO> pmis) throws UnsupportedEncodingException {
		Map<Long, String> map = new HashMap<>();
		for (PrivateMessageInfoPO pmi : pmis) {
			map.put(pmi.getMessageInfoId(), decode(pmi.getMessageInfo()));
		}
		UserLetter umInfo = null;
		for (PrivateMessagePO pm : pms) {
			umInfo = new UserLetter();
			umInfo.setUid(pm.getOuid());
			umInfo.setTouid(pm.getUid());
			umInfo.setSendTime(pm.getCreateTime().getTime());
			umInfo.setType(pm.getType());
			umInfo.setLetter(map.get(Long.parseLong(pm.getMessageInfoId())));
			umlist.add(umInfo);
		}
	}

}
