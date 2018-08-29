///** 
// *<p>开发公司：	鹏联优思 <p>
// *<p>版权所有：	鹏联优思 <p>
// *<p>责任人：	王东阳    <p> 
// *<p>网址：www.penslink.com <p>
// */
//
//package com.lanlian.server.redis;
//
//import java.util.Map;
//
//import javax.annotation.Resource;
//
//import org.apache.log4j.Logger;
//import org.springframework.context.ApplicationListener;
//import org.springframework.context.event.ContextRefreshedEvent;
//
//import com.lanlian.chatting.result.Fatal_Exception;
//
//import redis.clients.jedis.Jedis;
//import redis.clients.jedis.JedisPool;
//
///**
// * @author wangdyq
// * @version V1.0
// * @email wangdy@radacat.com
// * @date 2018年4月17日 下午6:27:06
// * @explain
// */
//
//public class RedisConf implements ApplicationListener<ContextRefreshedEvent> {
//
//	private static Logger logger = Logger.getLogger(RedisConf.class);
//
//	@Resource(name = "mutilJedisPoolMap")
//	Map<String, JedisPool> mutilJedisPoolMap;
//
//	private JedisPool jedisPool79;
//	private JedisPool jedisPool80;
//	private Jedis jedis79;
//	private Jedis jedis80;
//
//	@Override
//	public void onApplicationEvent(ContextRefreshedEvent event) {
//		if (event.getApplicationContext().getParent() == null) {
//			jedisPool79 = mutilJedisPoolMap.get("jedisPool79");
//			jedisPool80 = mutilJedisPoolMap.get("jedisPool80");
//		}
//	}
//
//	public Jedis getRedisConn(Integer port) throws Fatal_Exception {
//		try {
//			return initRedisConn(port);
//		} catch (Exception e) {
//			e.printStackTrace();
//			logger.fatal("RedisConf-getConn:" + e.getMessage());
//			throw new Fatal_Exception("获取redis连接错误");
//		}
//	}
//
//	private Jedis initRedisConn(Integer port) throws Fatal_Exception {
//		switch (port) {
//		case 6379:
//			this.jedis79 = jedisPool79.getResource();
//			return jedis79;
//		case 6380:
//			this.jedis80 = jedisPool80.getResource();
//			return jedis80;
//		default:
//			return null;
//		}
//	}
//	
//	/**
//	 * 结束redis
//	 * 
//	 * @param jedis
//	 */
//	public void close() {
//		if (jedis79.isConnected()) {
//			jedis79.close();
//		}
//	}
//
//	static {
//		/**
//		 * sets group id 群的sets
//		 */
//		String key3 = null;
//		String KEY_3 = "chatting_groupid_sets";
//
//		/**
//		 * key tgid id 临时群id
//		 */
//		String key5 = null;
//		String KEY_5 = "chatting_tgid_channel";
//
//		/**
//		 * key tid id 临时嗒嗒id
//		 */
//		String key6 = null;
//		String KEY_6 = "chatting_tid_channel";
//
//		/**
//		 * key tgid id 临时群id未启用
//		 */
//		String key7 = null;
//		String KEY_7 = "chatting_tgid_state";
//
//		/**
//		 * key tgid id 临时群id未启用
//		 */
//		String key8 = null;
//		String KEY_8 = "chatting_temporary_group";
//
//		/**
//		 * key allTGid id 16个临时群id
//		 */
//		String key9 = null;
//		String KEY_9 = "chatting_tgid_all";
//
//		/**
//		 * 上报群信息key
//		 */
//		String key10 = null;
//		String KEY_10 = "chatting_gid_msg";
//
//		/**
//		 * key 上报群开关hash
//		 */
//		String key11 = null;
//		String KEY_11 = "chatting_gid_uid_switch";
//
//		/**
//		 * HTTP请求源的key
//		 */
//		String key13 = null;
//		String KEY13 = "key13";
//
//		/**
//		 * 意见反馈邮箱地址
//		 */
//		String KEY14 = "Ag:chatting:v1.0:feedback:mail";
//
//	}
//}
