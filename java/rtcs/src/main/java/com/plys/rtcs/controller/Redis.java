package com.plys.rtcs.controller;

/// **
// *<p>开发公司 : 蓝涟科技<p>
// *<p>版权所有 : 蓝涟科技<p>
// *<p>责任人 : 王东阳<p>
// *<p>网址 : www.radacat.com<p>
// *<p>邮箱 : wangdy@radacat.com<p>
// */
//
// package controller;
//
// import java.util.Map;
// import java.util.ResourceBundle;
//
// import javax.annotation.Resource;
//
// import org.springframework.context.ApplicationListener;
// import org.springframework.context.event.ContextRefreshedEvent;
// import org.springframework.stereotype.Service;
//
// import redis.clients.jedis.Jedis;
// import redis.clients.jedis.JedisPool;
//
/// **
// * @author 王东阳
// * @version V1.0
// * @email wangdy@radacat.com
// * @date 2017年9月13日 下午4:57:49
// * @$
// * @Administrator
// * @explain
// */
//
// @Service("redis79")
// public class Redis implements ApplicationListener<ContextRefreshedEvent> {
//
// @Resource(name = "mutilJedisPoolMap")
// Map<String, JedisPool> mutilJedisPoolMap;
//
// JedisPool jedisPool;
//
// /**
// * sets group id 群的sets
// */
// private String key3 = null;
// private static final String KEY_3 = "rtcs.user.session";
//
// // 读取配置文件信息
// private static final ResourceBundle bundle =
/// ResourceBundle.getBundle("rediskey");
//
// /**
// * 初始化redis
// *
// * @return
// */
// @Override
// public void onApplicationEvent(ContextRefreshedEvent arg0) {
// try {
// jedisPool = mutilJedisPoolMap.get("jedisPool79");
// key3 = bundle.getString(KEY_3);
// } catch (Exception e) {
// e.printStackTrace();
// }
// }
//
// /**
// * 获取一个redis连接
// *
// * @param mutilJedisPoolMap
// * @param jedis79
// */
// public Jedis getConnection() {
// Jedis jedis79 = jedisPool.getResource();
// return jedis79;
// }
//
// /**
// * 结束redis
// *
// * @param jedis
// */
// public void close(Jedis jedis) {
// if (jedis.isConnected()) {
// jedis.close();
// }
// }
//
// public String setSession(Integer id, Object object) {
// Jedis redis = null;
// try {
// System.out.println("redis79-setSession-sessionid:" + id + ",obj:" + object);
// redis = getConnection();
// redis.select(14);
// redis.set(key3+id, object.toString());
// String session = redis.get(key3+id);
// System.out.println("redis79-setSession-session:"+session);
// return session;
// } catch (Exception e) {
// e.printStackTrace();
// return null;
// } finally {
// close(redis);
// }
// }
//
// public String getSession(Integer id) {
// Jedis redis = null;
// try {
// System.out.println("id:" + id);
// redis = getConnection();
// redis.select(14);
// String session = redis.get(key3+id);
// System.out.println("sess:"+session);
// return session;
// } catch (Exception e) {
// e.printStackTrace();
// return null;
// } finally {
// close(redis);
// }
// }
//
// }
//
