/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

package com.lanlian.server.redis;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
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
import com.lanlian.chatting.po.virtual.TemporaryGroup;
import com.lanlian.chatting.result.Parameter_Exception;
import com.lanlian.chatting.util.JsonBeanUtil;

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
@Service("redisClient79")
public class RedisClient79 implements ApplicationListener<ContextRefreshedEvent> {
	
	private static Logger logger = Logger.getLogger(RedisClient79.class);

	@Resource(name = "mutilJedisPoolMap")
	Map<String, JedisPool> mutilJedisPoolMap;
	
	private JedisPool jedisPool;

	/**
	 * sets group id 群的sets
	 */
	public String key3 = "chatting_groupid_sets";

	/**
	 * key tgid id 临时群id
	 */
	public String key5 = "chatting_tgid_channel";

	/**
	 * key tid id 临时嗒嗒id
	 */
	public String key6 = "chatting_tid_channel";

	/**
	 * key tgid id 临时群id未启用
	 */
	public String key7 = "chatting_tgid_state";

	/**
	 * key tgid id 临时群id未启用
	 */
	public String key8 = "chatting_temporary_group";

	/**
	 * key allTGid id 16个临时群id
	 */
	public String key9 = "chatting_tgid_all";
	
	/**
	 * 上报群信息key
	 */
	public String key10 = "chatting_gid_msg";
	
	/**
	 * key 上报群开关hash
	 */
	public String key11 = "chatting_gid_uid_switch";
	
	/**
	 * HTTP请求源的key
	 */
	public String key13 = "key13";
	
	/**
	 * 意见反馈邮箱地址
	 */
	public String key14 = "Ag:chatting:v1.0:feedback:mail";
	
	/**
	 * 用户心跳信息的key
	 */
	public String key15 = "key15";

	// 读取配置文件信息
	public static final ResourceBundle bundle = ResourceBundle.getBundle("rediskey");

	/**
	 * 初始化redis-key
	 * 
	 * @return
	 */
	@Override
	public void onApplicationEvent(ContextRefreshedEvent event) {
		try {
			if (event.getApplicationContext().getParent() == null) {
				jedisPool = mutilJedisPoolMap.get("jedisPool79");
				key3 = key3 + bundle.getString(key3);
				key5 = key5 + bundle.getString(key5);
				key6 = key6 + bundle.getString(key6);
				key7 = key7 + bundle.getString(key7);
				key8 = key8 + bundle.getString(key8);
				key9 = key9 + bundle.getString(key9);
				key10 = key10 + bundle.getString(key10);
				key11 = key11 + bundle.getString(key11);
				key13 = bundle.getString(key13);
				key15 = bundle.getString(key15);
				initId(0, key5, "268435457");
				initId(1, key6, "16777217");
				initAllTGId(0, key9, "10000");
			}
		} catch (Exception e) {
			logger.fatal("Redis连接错误---JedisConnectionException");
			logger.error("Redis连接错误---JedisConnectionException:" + e.getMessage());
		}
	}

	/**
	 * 获取一个redis连接
	 * 
	 * @param jedis
	 */
	public Jedis getConnection() {
		return jedisPool.getResource();
	}

	/**
	 * 结束redis
	 * 
	 * @param jedis
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
	
	/**
	 * 初始化id
	 * 
	 * @throws Parameter_Exception
	 */
	private void initId(Integer flag, String key, String value) {
		logger.info("init-gid-key:" + key + ",value:" + value);
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			// 初始化tid,tgid
			switch (flag) {
			case 0:
				if (!jedis.exists(key)) {
					jedis.set(key, value);
				}
				break;
			case 1:
				String as = jedis.set(key, value);
				logger.info("initId: " + as);
				break;
			default:
				break;
			}
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 存入string
	 * @param key
	 * @return
	 * @throws Parameter_Exception 
	 */
	public void setString(String key,String value) throws Parameter_Exception {
		logger.info("redisClien79 setString key:"+key+",value:"+value);
		Jedis jedis = getConnection();
		try {
			jedis.select(0);
			String isOk = jedis.set(key, value);
			if (isOk.equals("OK")) {
				return;
			}
			logger.error("redisClien79 setString key:"+key+",value:"+value);
			throw new Parameter_Exception(10001);
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 获取一个string
	 * @param key
	 * @return
	 */
	public String getString(String key) {
		Jedis jedis = getConnection();
		try {
			jedis.select(0);
			if (jedis.exists(key)) {
				String value = jedis.get(key);
				return value;
			} else {
				return null;
			}			
		} finally {
			close(jedis);
		}
	}

	/********************** 回传位置信息 ******************************************/

	/**
	 * 将获取位置信息存入list
	 * 
	 * @param data
	 * @throws Parameter_Exception
	 */
	public void setlocation(String key, String data) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(0);
			jedis.rpush(key, data);
		} finally {
			close(jedis);
		}
	}

	/********************** 临时群同步信息 **********************************/

	/**
	 * 将临时群信息存入
	 * 
	 * @param data
	 * @param temporaryGroup
	 * @throws Parameter_Exception
	 */
	public void setTemporaryGroup(TemporaryGroup temporaryGroup) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			net.sf.json.JSONObject jsonObject = net.sf.json.JSONObject.fromObject(temporaryGroup);
			jedis.set(key8 + temporaryGroup.getUid().toString(), jsonObject.toString());
			logger.info("setTemporaryGroup-key:" + key8 + temporaryGroup.getUid() + ",temporaryGroup:" + jsonObject);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 查询临时群信息
	 * 
	 * @param key
	 * @param temporaryGroup
	 * @throws Parameter_Exception
	 */
	public void findTemporaryGroup(TemporaryGroup temporaryGroup) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			String data = jedis.get(key8 + temporaryGroup.getUid().toString());
			JsonBeanUtil.jsonToBean(data, temporaryGroup.getClass());
			logger.info("findTemporaryGroup-key:" + key8 + temporaryGroup.getUid() + ",temporaryGroup:" + temporaryGroup);
		} finally {
			close(jedis);
		}
	}

	/************************* 虚拟信道服务 *******************************************/
	/**
	 * 获取一个临时群id
	 * 
	 * @return
	 * @throws Parameter_Exception
	 */
	public Integer getGid() throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			if (!jedis.exists(key5)) {
				initId(1, key5, "268435457");
			}
			String value = jedis.get(key5);
			Integer tgid = Integer.valueOf(value);
			jedis.set(key5, (--tgid).toString());
			if (tgid.equals(2)) {
				initId(1, key5, "268435457");
			}
			return tgid;
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 获取一个临时id
	 * 
	 * @return
	 * @throws Parameter_Exception
	 */
	public Integer getTid() throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			if (!jedis.exists(key6)) {
				initId(1, key6, "16777217");
			}
			String value = jedis.get(key6);
			Integer tid = Integer.valueOf(value);
			jedis.set(key6, (--tid).toString());
			if (tid.intValue() <= 1) {
				initId(1, key6, "16777217");
			}
			return tid;
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 将获取的tgid存入已建群sets
	 * 
	 * @param tgid
	 * @throws Parameter_Exception
	 */
	public void setGroupId(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			jedis.sadd(key7, nearbyPo.getTgid().toString());
			jedis.select(14);
			JSONObject jsonObject = new JSONObject();
			jsonObject.put("ownerid", nearbyPo.getOwnerid());
			jsonObject.put("notice", "");
			jedis.set(nearbyPo.getTgid().toString(), jsonObject.toJSONString());
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 获取已建群tgid的资料
	 * 
	 * @param tgid
	 * @throws Parameter_Exception
	 */
	public void findGroupInfo(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(14);
			String str = jedis.get(nearbyPo.getTgid().toString());
			JSONObject jsonObject = JSONObject.parseObject(str);
			Integer ownerid = jsonObject.getInteger("ownerid");
			logger.info("ownerid:" + ownerid);
			String notice = jsonObject.getString("notice");
			nearbyPo.setOwnerid(ownerid);
			nearbyPo.setNotice(notice);
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 获取已建群tgid的ownerid
	 * 
	 * @param tgid
	 * @throws Parameter_Exception
	 */
	public void findGroupId(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(14);
			String str = jedis.get(nearbyPo.getTgid().toString());
			JSONObject jsonObject = JSONObject.parseObject(str);
			Integer ownerid = jsonObject.getInteger("ownerid");
			nearbyPo.setOwnerid(ownerid);
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 修改已建群tgid的资料
	 * 
	 * @param tgid
	 * @throws Parameter_Exception
	 */
	public void updateGroupId(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			logger.info("updateGroupId-nearbyPo" + nearbyPo);
			jedis.select(14);
			String str = jedis.get(nearbyPo.getTgid().toString());
			JSONObject jsonObject = JSONObject.parseObject(str);
			jsonObject.put("notice", nearbyPo.getNotice());
			jedis.set(nearbyPo.getTgid().toString(), jsonObject.toJSONString());
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 解散群
	 * 
	 * @param tgid
	 * @throws Parameter_Exception
	 */
	public void deleteGroupId(Integer tgid) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			jedis.del(tgid.toString());
			jedis.select(14);
			jedis.del(tgid.toString());
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 向群sets添加群成员
	 * 
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	public void setGroupMember(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			logger.info(nearbyPo);
			jedis.select(15);
			boolean flag = jedis.sismember(key7, nearbyPo.getTgid().toString());
			if (!flag) {
				throw new Parameter_Exception(21005);
			}
			List<Integer> tuids = nearbyPo.getTuids();
			String[] uids = new String[tuids.size()];
			for (int i = 0; i < uids.length; i++) {
				uids[i] = tuids.get(i).toString();
			}
			logger.info(Arrays.asList(uids) + nearbyPo.toString());
			jedis.sadd(nearbyPo.getTgid().toString(), uids);
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 获取sets群成员
	 * 
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	public List<String> findGroupMember(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			Set<String> sets = jedis.smembers(nearbyPo.getTgid().toString());
			List<String> list = new ArrayList<>();
			list.addAll(sets);
			return list;
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * 群sets删除群成员
	 * 
	 * @param nearbyPo
	 * @throws Parameter_Exception
	 */
	public void deleteGroupMember(NearbyPo nearbyPo) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			logger.info(nearbyPo);
			jedis.select(15);
			List<Integer> tuids = nearbyPo.getTuids();
			String[] uids = new String[tuids.size()];
			for (int i = 0; i < uids.length; i++) {
				uids[i] = tuids.get(i).toString();
			}
			jedis.srem(nearbyPo.getTgid().toString(), uids);
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**************************** 群缓存 ******************************************/

	/**
	 * sets 存入被订阅的群id
	 * 
	 * @param groupid
	 * @throws Parameter_Exception
	 */
	public void saveSubscribeGroupId(Integer... groupids) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			logger.info("saveSubscribeGroupId-key:" + Arrays.asList(groupids));
			jedis.select(15);
			String[] gids = new String[groupids.length];
			for (int i = 0; i < groupids.length; i++) {
				gids[i] = groupids[i].toString();
			}
			jedis.sadd(key3, gids);
			logger.info("saveSubscribeGroupId-key3:" + key3 + ",ids:" + Arrays.asList(groupids));
		} catch (Exception e) {
			logger.error(e);
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	/**
	 * sets 删除被订阅的群id
	 * 
	 * @param groupid
	 * @throws Parameter_Exception
	 */
	public void deleteSubscribeGroupId(Integer... groupids) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			for (Integer groupid : groupids) {
				// 删除数据
				jedis.srem(key3, groupid.toString());
			}
			logger.info("deleteSubscribeGroupId-key3:" + key3 + ",ids:" + Arrays.asList(groupids));
		} finally {
			close(jedis);
		}
	}

	/**
	 * sets 查询被订阅的群id
	 * 
	 * @param groupid
	 * @throws Parameter_Exception
	 */
	public boolean findSubscribeGroupId(Integer groupid) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			logger.info("findSubscribeGroupId-key:" + groupid);
			jedis.select(15);
			// 判断是否包含该数据
			boolean flag = jedis.sismember(key3, groupid.toString());
			logger.info("findSubscribeGroupId-key3:" + key3 + ",gid:" + groupid);
			return flag;
		} finally {
			close(jedis);
		}
	}

	/**
	 * 查询手机验证码
	 * 
	 * @param username
	 * @param code
	 * 
	 * @throws Parameter_Exception
	 */
	public void checkCode(String username, String code) throws Parameter_Exception {
		String number = "";
		Jedis jedis = getConnection();
		try {
			jedis.select(0);
			logger.info("checkCode-jedis:" + jedis.get(username));
			number = jedis.get(username);
			if (number == null || !code.trim().equals(number)) {
				throw new Parameter_Exception(20013);
			}
		} finally {
			close(jedis);
		}
	}

	/**************************** 群id ******************************************/

	/**
	 * 分配给用户16个临时群id
	 * 
	 * @param flag
	 *            标志
	 * @param key
	 *            值为常量key9
	 * @param setValue
	 *            "10000000"
	 * @throws Parameter_Exception
	 */
	public void initAllTGId(Integer flag, String key, String setValue) {
		logger.info("initAllTGId-check-gid：" + flag + ", " + key);
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			switch (flag) {
			case 0:
				if (!jedis.exists(key)) {
					jedis.set(key, setValue);
					logger.info("initAllTGId: init 16 tempGroup id:" + jedis.smembers(key));
				}
				return;
			case 1:
				if (!jedis.exists(key)) {
					jedis.set(key, setValue);
					logger.info("initAllTGId: resetting 16 tempGroup id" + jedis.smembers(key));
				}
				return;	
			default:
				logger.info("initAllTGId-gid: flag:" + flag + ",parameter error");
				return;
			}
		} finally {
			close(jedis);
		}
	}

	/**
	 * 
	 * @param 嗒嗒id
	 * 将获取的16个tgid存入已建群sets
	 * @return 此id的16个临时群id的set集合元素
	 * @throws Parameter_Exception
	 */
	public Set<Integer> getAllTGId(Integer uid) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			Set<Integer> setTGId = new HashSet<>();
			String keyUid = key9 + uid.toString();
			logger.info("keyUid:" + keyUid);
			if (!jedis.exists(keyUid)) {
				// 不存在就开始生成16个临时群id
				setTGId = generateAllGId(jedis, 16, keyUid);
				logger.info("getAllTGId: generate 16 tempGroup id  " + setTGId);
				return setTGId;
			}
			// 取出16个并返回
			Set<String> setId = jedis.smembers(keyUid);
			for (String sgid : setId) {
				setTGId.add(Integer.valueOf(sgid));
			}
			logger.info("getAllTGId: getting 16 tempGroup id  " + setId);
			return setTGId;
		} finally {
			close(jedis);
		}
	}

	/**
	 * 获取新的群id并存入用户群集合，去除用户的该未激活群id
	 * 
	 * @param 升级的群id
	 * @param 嗒嗒id
	 * @return 此id的16个临时群id的set集合元素
	 * @throws Parameter_Exception
	 */
	public Integer upgradeOneTG(Integer gId, Integer uid) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			String key = key9 + uid.toString();
			//生成一个群id
			Set<Integer> newTGId = generateAllGId(jedis, 1, key);
			//移除gid
			jedis.srem(key, gId.toString());
			//取出新的gid
			Integer[] idArr = (Integer[]) newTGId.toArray(new Integer[1]);
			//新的gid添加到redis里面
			jedis.sadd(key, idArr[0].toString());
			logger.info("upgradeOneTG: setting gId:" + gId + " a perpetual group");
			return idArr[0];
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 生成群id,不需要重新获取redis
	 * 
	 * @param 判断位：生成16个临时群id还是1个临时群id
	 * @param 已经加修饰的用户的嗒嗒id
	 *            key
	 * @param 传过来redis里保存的临时群id的值
	 * @throws Parameter_Exception
	 */
	public synchronized Set<Integer> generateAllGId(Jedis jedis, Integer flag, String key) throws Parameter_Exception {
		logger.info("gid-key:" + key9 + ",key:" + key + ",flag:" + flag);
		Set<Integer> setId = new HashSet<>();
		String lastTGId = jedis.get(key9);
		Integer tgid = Integer.valueOf(lastTGId);
		switch (flag) {
		case 16:
			//遍历生成16个gid放入Set集合和redis中
			String [] strArr = new String[16];
			for (int i = 0; i < 16; i++) {
				setId.add(tgid);
				strArr[i] = tgid.toString();
				tgid++;
			}
			//16个群id放入redis
			jedis.sadd(key, strArr);
			//更新下redis里面的群id的已生成数
			jedis.set(key9, tgid.toString());
			logger.info("lastTGId:" + lastTGId + ",tgid" + tgid);
			break;
		case 1:
			setId.add(tgid);
			String tgidAdd = (++tgid).toString();
			jedis.set(key9, tgidAdd);
			logger.info("tgidAdd: " + tgidAdd);
			break;
		default:
			logger.info("flag = " + flag + "parameter error");
			throw new Parameter_Exception(10002);
		}
		return setId;
	}

	/**
	 * 判断用户的用户的未激活群里面是否包含这个群
	 * 
	 * @param 未激活的群id
	 * @param 用户id
	 * @throws Parameter_Exception
	 */
	public Boolean isExistAndSismember(String key, Integer gid) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			logger.info("isExistAndSismember-key:"+key+",gid:"+gid);
			if (!jedis.exists(key)) {
				return false;
			}
			Boolean bool = jedis.sismember(key, gid.toString());
			logger.info("isExistAndSismember-bool:" + bool);
			return bool;
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 记录群用户上报开关
	 * @param gid
	 * @param uid
	 * @param state
	 * @return
	 * @throws Parameter_Exception
	 */
	public void setGidUidSwitch(Integer gid, Integer uid, Integer state) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			logger.info("switchGidUid-state:" + state + ",gid:" + gid + ",uid:" + uid);
			jedis.select(15);
			String key = key11 + gid.toString();
			jedis.hset(key, uid.toString(), state.toString());
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 获取群用户上报开关
	 * @param gid
	 * @param uid
	 * @return
	 * @throws Parameter_Exception
	 */
	public Integer getGidUidSwitchs(Integer gid, Integer uid) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			logger.info("switchGidUid-gid:" + gid + ",uid:" + uid);
			jedis.select(15);
			String key = key11 + gid.toString();
			boolean flag = jedis.hexists(key, uid.toString());
			if (flag) {
				String state = jedis.hget(key, uid.toString());
				return Integer.valueOf(state);				
			}
			throw new Parameter_Exception(21005);
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 选取上报群聊用户
	 * @param gid
	 * @param uid
	 * @return
	 * @throws Parameter_Exception
	 */
	public Boolean groupMsgUid(Integer gid, Integer uid) throws Parameter_Exception {
		logger.info("dadaGroupBackup-gid:"+gid+",uid:"+uid);
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			String key = key10 + gid.toString();
			String oldUid = jedis.get(key);
			if (oldUid == null) {
				jedis.setex(key, 10, uid.toString());
				return true;
			}
			Integer olduid2 = Integer.valueOf(oldUid);
			if (olduid2.equals(uid)) {
				jedis.expire(key, 10);
				return true;
			}
			return false;
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 上报消息
	 * @param gid
	 * @param uid
	 * @return
	 * @throws Parameter_Exception
	 */
	public int groupGroupSwitchsMsg(Integer gid, Integer uid) throws Parameter_Exception {
		logger.info("dadaGroupBackup-gid:"+gid+",uid:"+uid);
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			String key = key11 + gid.toString();
			boolean flag = jedis.hexists(key, uid.toString());
			if (flag) {
				String state = jedis.hget(key, uid.toString());
				if (Integer.valueOf(state).intValue() == 0) {
					return 3;
				}
				key = key10 + gid.toString();
				String openUid = jedis.get(key);
				if (openUid == null) {
					jedis.setex(key, 10, uid.toString());
					return 1;
				}
				Integer olduid2 = Integer.valueOf(openUid);
				if (olduid2.equals(uid)) {
					jedis.expire(key, 10);
					return 1;
				}
				return 2;
			}
			return 4;
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * @param key
	 * @param hashKey
	 * @param hashValue
	 * @return 向hash中存入键值对
	 * @throws Parameter_Exception 
	 * @throws IOException 
	 */
	public void setHashKV(String key, Map<String, String> map) throws Parameter_Exception {
		System.out.println("setHashKV Key:"+key+",map"+map);
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			String str = jedis.hmset(key, map);
			System.out.println("aaa:"+str);
			if (str.equals("OK")) {
				return;
			}
			throw new Parameter_Exception(10001);
		} catch (Parameter_Exception e) {
			throw e;
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * @param key
	 * @param hashKey
	 * @return 从hash中获取所有键值对
	 * @throws Parameter_Exception
	 */
	public Map<String, String> getHashKV(String key) throws Parameter_Exception {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			return jedis.hgetAll(key);
		} finally {
			close(jedis);
		}
	}
	
	/********************************   redis操作失败，补偿事物    ***************************************/

	/**
	 * 群操作失败，补偿redis事物
	 * 
	 * @param gid
	 */
	public synchronized void compensateRedis(Integer uid, Integer oldgid, Integer newgid) {
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			String key = key9 + uid.toString();
			logger.info("redis-compensateRedis-key:" + key + ",uid:" + uid + ",oldgid:" + oldgid + ",newgid:" + newgid);
			// 群id生成器回滚id,id生成器减一
			String lastTgid = jedis.get(key9);
			Integer tgid = Integer.valueOf(lastTgid);
			jedis.set(key9, (--tgid).toString());
			// 对用户群id-sets回滚,删除新群id，添加旧群id
			jedis.srem(key, newgid.toString());
			jedis.sadd(key, oldgid.toString());
			logger.info("redis-compensateRedis-key:" + key + ",uid:" + uid + ",oldgid:" + oldgid + ",newgid:" + newgid);
		} finally {
			close(jedis);
		}
	}
	
	/**
	 * 群操作失败，模拟回滚事物
	 * @param gid
	 * @return
	 * @throws Parameter_Exception
	 */
	public void delGIdUid(Integer gid) throws Parameter_Exception  {
		logger.info("dadaGroupBackup-gid:"+gid);
		Jedis jedis = getConnection();
		try {
			jedis.select(15);
			String key = key10 + gid.toString();
			Long num = jedis.del(key);
			logger.info("dadaGroupBackup-num:"+num);
		} finally {
			close(jedis);
		}
	}
	
}
