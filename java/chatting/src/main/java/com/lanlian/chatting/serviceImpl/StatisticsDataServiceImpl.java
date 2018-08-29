/** 
 *<p>开发公司：	鹏联优思 <p>
 *<p>版权所有：	鹏联优思 <p>
 *<p>责任人：	王东阳    <p> 
 *<p>网址：www.penslink.com <p>
 */

package com.lanlian.chatting.serviceImpl;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

import javax.annotation.Resource;

import org.apache.log4j.Logger;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.model.UserHeartbeatModel;
import com.lanlian.chatting.model.UserHeartbeatsModel;
import com.lanlian.chatting.po.SsoPo;
import com.lanlian.chatting.po.UserHeartbeatPO;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.service.StatisticsDataService;
import com.lanlian.chatting.util.PublicMethod;
import com.lanlian.server.redis.RedisClient79;

/**
 * @author  wangdyq
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2018年8月2日 下午2:14:48
 * @explain 统计信息业务
 */

@Service(value = "statisticsDataService")
public class StatisticsDataServiceImpl implements StatisticsDataService {
	
	private static Logger logger = Logger.getLogger(StatisticsDataServiceImpl.class);
	
	@Resource
	RedisClient79 redisClient79;
	
	private String groupHashKey = ":info";
	
	/**
	 * 上报心与获取心跳
	 * @param userHeartbeatPO
	 * @return
	 * @throws Parameter_Exception 
	 * @throws IOException 
	 */
	public UserHeartbeatsModel heartbeat(SsoPo ssoPo,UserHeartbeatPO userHeartbeatPO,String info) throws Parameter_Exception {
		//检查群与群主id
		if (!redisClient79.isExistAndSismember(redisClient79.key9+userHeartbeatPO.getOwnerId(), userHeartbeatPO.getGroupId())) {
			throw new Parameter_Exception(21002);
		}
		String json = JSONObject.toJSONString(userHeartbeatPO);
		writerLogFile(ssoPo.getUid()+"|||"+json);
		String redisKey = redisClient79.key15+userHeartbeatPO.getGroupId()+":"+userHeartbeatPO.getChanId()+":"+userHeartbeatPO.getRoomId();
		if (ssoPo.getUid().intValue() == userHeartbeatPO.getOwnerId().intValue() && info != null) {
			redisClient79.setString(redisKey+groupHashKey, info);
		}
		Map<String, String> map = new HashMap<>();
		map.put(ssoPo.getUid().toString(), json);
		redisClient79.setHashKV(redisKey, map);
		//查询群与成员信息
		UserHeartbeatsModel userHeartbeatsModel = new UserHeartbeatsModel();
		userHeartbeatsModel.setGroupInfo(redisClient79.getString(redisKey+groupHashKey));
		map = redisClient79.getHashKV(redisKey);
		List<UserHeartbeatModel> list = new ArrayList<>();
		Iterator<Map.Entry<String, String>> entries = map.entrySet().iterator();
		while (entries.hasNext()) {
			Map.Entry<String, String> entry = entries.next();
			UserHeartbeatModel userHeartbeatModel = JSONObject.parseObject(entry.getValue(), UserHeartbeatModel.class);
			userHeartbeatModel.setUid(Integer.parseInt(entry.getKey()));
			list.add(userHeartbeatModel);
		}
		userHeartbeatsModel.setHeartbeats(list);
		return userHeartbeatsModel;
	}
	
	@Async
	public void writerLogFile(String context) {
		try {
			PublicMethod.printWriterFile("/var/log/chatting/heartbeat/user_heartbeat.log", context);
		} catch (Exception e) {
			logger.fatal("printWriterFile err:"+e);
			logger.error("printWriterFile context:"+context);
		}
		
	}
	
}
