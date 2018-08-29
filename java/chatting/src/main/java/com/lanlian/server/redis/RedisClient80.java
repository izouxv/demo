/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.server.redis;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.ResourceBundle;
import java.util.Set;

import javax.annotation.Resource;

import org.apache.log4j.Logger;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;
import com.lanlian.chatting.po.virtual.NearbyPo;
import com.lanlian.chatting.po.virtual.NearbyGroupInfoPo;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.vo.MessageInfoPage;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;

/**
 * @Title RedisClient.java
 * @Package com.lanlian.chatting.redis
 * @author 王东阳
 * @version V1.0
 * @date 2017年6月14日 上午9:37:08
 * @explain
 */
@Service("redisClient80")
public class RedisClient80 implements ApplicationListener<ContextRefreshedEvent> {
	
	private static Logger logger = Logger.getLogger(RedisClient79.class);

	@Resource(name = "mutilJedisPoolMap")
	Map<String, JedisPool> mutilJedisPoolMap;
	
	private JedisPool jedisPool;

	/**
	 * list group id list群的id
	 */
	private String key1 = null;
	private static final String KEY_1 = "chatting_groupid_list";

	/**
	 * list group msg 群信息的list
	 */
	@SuppressWarnings("unused")
	private String key2 = null;
	private static final String KEY_2 = "chatting_groupmsg_list";

	/**
	 * pub group id 发布订阅群id
	 */
	private String key4 = null;
	private static final String KEY_4 = "chatting_groupid_pub";

	// 读取配置文件信息
	private static final ResourceBundle bundle = ResourceBundle.getBundle("rediskey");

	/**
	 * 初始化redis
	 * 
	 * @return
	 */
	@Override
	public void onApplicationEvent(ContextRefreshedEvent arg0) {
		jedisPool = mutilJedisPoolMap.get("jedisPool80");
		key1 = KEY_1 + bundle.getString(KEY_1);
		key2 = KEY_2 + bundle.getString(KEY_2);
		key4 = KEY_4 + bundle.getString(KEY_4);
		//key改进之后
//			key1 = "AQ:chatting:setting_message_info:";
//			key2 = KEY_2 + bundle.getString(KEY_2);
//			key4 = "AQ:chatting:setting_publish_groupid";
	}
	
	/**
	 * 获取一个redis连接
	 * 
	 * @param mutilJedisPoolMap
	 * @param jedis
	 */
	public Jedis getConnection() {
		return jedisPool.getResource();
	}

	/**
	 * 结束redis
	 */
	public void close(Jedis jedis) {
		try {
			if (jedis.isConnected()) {
				jedis.close();
			}
		} catch (Exception e) {
			logger.fatal("Redis连接close错误:"+e.getMessage());
		}
	}

	/************************* 虚拟信道服务 *******************************************/

	/**
	 * 存储用户的当前空间信息到redis的key-value
	 * 
	 * @param geoHash
	 * @throws Parameter_Exception
	 */
	public void saveNearbyToKey(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			JSONObject jsonObject = new JSONObject();
			jsonObject.put("tuid", nearbyPo.getTuid());
			jsonObject.put("nickname", nearbyPo.getNickname());
			jsonObject.put("avatar", nearbyPo.getAvatar());
			jsonObject.put("gender", nearbyPo.getGender());
			jsonObject.put("age", nearbyPo.getAge());
			jsonObject.put("signature", nearbyPo.getSignature());
			jsonObject.put("imei", nearbyPo.getImei());
			jsonObject.put("longitude", nearbyPo.getLongitude());
			jsonObject.put("latitude", nearbyPo.getLatitude());
			jedis.set(nearbyPo.getTuid().toString(), jsonObject.toJSONString());
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 删除用户的当前空间信息-redis的key-value
	 * 
	 * @param geoHash
	 * @throws Parameter_Exception
	 */
	public void deleteNearbyToKey(Integer tuid) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			jedis.del(tuid.toString());
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 获取用户的当前空间信息到redis的key-value
	 * 
	 * @param geoHash
	 * @throws Parameter_Exception
	 */
	public List<NearbyPo> findNearbyToKey(List<String> tuids) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			List<NearbyPo> list = new ArrayList<>();
			NearbyPo nearbyPo = null;
			for (String tuid : tuids) {
				String str = jedis.get(tuid);
				if (str == null) {
					continue;
				}
				System.out.println("jedis:" + str);
				nearbyPo = new NearbyPo();
				JSONObject jsonObject = JSONObject.parseObject(str);
				System.out.println(jsonObject);
				nearbyPo.setTuid(jsonObject.getInteger("tuid"));
				nearbyPo.setNickname(jsonObject.getString("nickname"));
				nearbyPo.setAvatar(jsonObject.getInteger("avatar"));
				nearbyPo.setGender(jsonObject.getInteger("gender"));
				nearbyPo.setSignature(jsonObject.getString("signature"));
				nearbyPo.setImei(jsonObject.getString("imei"));
				nearbyPo.setLongitude(jsonObject.getDouble("longitude"));
				nearbyPo.setLatitude(jsonObject.getDouble("latitude"));
				list.add(nearbyPo);
			}
			return list;
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 存储用户的当前geohash到sets
	 * 
	 * @param geoHash
	 * @throws Parameter_Exception
	 */
	public void saveNearbyToSets(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			jedis.sadd(nearbyPo.getGeoHash(), nearbyPo.getTuid().toString());
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 根据用户的当前geohash获取九块sets的并集
	 * 
	 * @param geoHash
	 * @throws Parameter_Exception
	 */
	public Set<String> findNearbyToSets(String... geohash) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			Set<String> set = jedis.sunion(geohash);
			return set;
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 向附近的人发送信息到对方的list
	 * 
	 * @param nearbyGroupInfoPo
	 * @throws Parameter_Exception
	 */
	public void sendNearbyInfoToKey(NearbyGroupInfoPo nearbyGroupInfoPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(14);
			JSONObject jsonObject = new JSONObject();
			jsonObject.put("types", nearbyGroupInfoPo.getTypes());
			jsonObject.put("tuid", nearbyGroupInfoPo.getTuid());
			jsonObject.put("toid", nearbyGroupInfoPo.getToid());
			jsonObject.put("time", nearbyGroupInfoPo.getTime());
			jsonObject.put("type", nearbyGroupInfoPo.getType());
			jsonObject.put("info", nearbyGroupInfoPo.getInfo());
			jedis.rpush(nearbyGroupInfoPo.getToid().toString(), jsonObject.toJSONString());
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 收取自己的list中的信息
	 * 
	 * @param geoHash
	 * @throws Parameter_Exception
	 */
	public List<NearbyGroupInfoPo> findNearbyInfoToKey(NearbyGroupInfoPo nearbyGroupInfoPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(14);
			NearbyGroupInfoPo nearbyGroupInfoPo1 = null;
			Long len = jedis.llen(nearbyGroupInfoPo.getTuid().toString());
			List<String> list2 = new ArrayList<>();
			for (int i = 0; i < len; i++) {
				String pop = jedis.lpop(nearbyGroupInfoPo.getTuid().toString());
				list2.add(pop);
			}
			List<NearbyGroupInfoPo> list = new ArrayList<>();
			for (String object : list2) {
				JSONObject jsonObject = JSONObject.parseObject(object);
				nearbyGroupInfoPo1 = new NearbyGroupInfoPo();
				nearbyGroupInfoPo1.setTypes(jsonObject.getInteger("types"));
				nearbyGroupInfoPo1.setTuid(jsonObject.getInteger("toid"));
				nearbyGroupInfoPo1.setToid(jsonObject.getInteger("tuid"));
				nearbyGroupInfoPo1.setTime(jsonObject.getLong("time"));
				nearbyGroupInfoPo1.setType(jsonObject.getInteger("type"));
				nearbyGroupInfoPo1.setInfo(jsonObject.getString("info"));
				list.add(nearbyGroupInfoPo1);
			}
			return list;
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**************************** 群缓存 ******************************************/

	/**
	 * 向存在的list中存入聊天数据
	 * 
	 * @param groupId
	 */
	public void setBundlingGroupInfo(String groupid, List<MessageInfoPage> list) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			// 存入数据
			for (MessageInfoPage messageInfoPage : list) {
				Object value = JSONObject.toJSON(messageInfoPage);
				// 向已存在的list中存入数据：jedis.rpushx(key, value.toString());
				jedis.rpush(key1 + groupid, value.toString());
			}
			//修剪list
			long len = jedis.llen(key1 + groupid);
			if (len > 15) {
				jedis.ltrim(key1 + groupid, len-15, -1);
			}
			//pub通知
			jedis.publish(key4, groupid.toString());
			logger.info("setBundlingGroupInfo-key1:" + key1 + ",list:" + list);
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

}
