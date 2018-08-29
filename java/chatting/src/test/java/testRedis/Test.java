/** 
 *<p>开发公司 :		          蓝涟科技 <p>
 *<p>版权所有 :		          蓝涟科技 <p>
 *<p>责任人     :		              王东阳 <p> 
 *<p>网址         :   www.radacat.com <p>
 *<p>邮箱         : wangdy@radact.com <p>
 */

package testRedis;

import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;

import com.lanlian.chatting.logger.LogUtil;
import com.lanlian.chatting.result.Parameter_Exception;

import junit.framework.Assert;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;
import redis.clients.jedis.Pipeline;
import redis.clients.jedis.ScanResult;
import redis.clients.jedis.Transaction;

/**
 * @author 王东阳
 * @version V1.0
 * @email wangdy@radacat.com
 * @date 2017年9月2日 下午4:15:33
 * @explain
 */

public class Test {

	// Redis服务器IP
	private static String ADDR = "192.168.1.6";

	// Redis的端口号
	private static int PORT = 6379;

	// 访问密码
	private static String AUTH = "radacat1234";

	// 可用连接实例的最大数目，默认值为8；
	// 如果赋值为-1，则表示不限制；如果pool已经分配了maxActive个jedis实例，则此时pool的状态为exhausted(耗尽)。
//	private static int MAX_ACTIVE = 1024;

	// 控制一个pool最多有多少个状态为idle(空闲的)的jedis实例，默认值也是8。
	private static int MAX_IDLE = 200;

	// 等待可用连接的最大时间，单位毫秒，默认值为-1，表示永不超时。如果超过等待时间，则直接抛出JedisConnectionException；
	private static int MAX_WAIT = 10000;

	private static int TIMEOUT = 10000;

	// 在borrow一个jedis实例时，是否提前进行validate操作；如果为true，则得到的jedis实例均是可用的；
	private static boolean TEST_ON_BORROW = true;

	private static JedisPool jedisPool = null;

	Jedis jedis;

	/**
	 * 初始化Redis连接池
	 */
	static {
		try {
			JedisPoolConfig config = new JedisPoolConfig();
			config.setMaxIdle(MAX_IDLE);
			config.setMaxWaitMillis(MAX_WAIT);
			config.setTestOnBorrow(TEST_ON_BORROW);
			jedisPool = new JedisPool(config, ADDR, PORT, TIMEOUT, AUTH);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	/**
	 * 获取Jedis实例
	 * 
	 * @return
	 */
	public synchronized static Jedis getJedis() {
		try {
			if (jedisPool != null) {
				Jedis resource = jedisPool.getResource();
				return resource;
			} else {
				return null;
			}
		} catch (Exception e) {
			e.printStackTrace();
			return null;
		}
	}

	/**
	 * 释放jedis资源
	 * 
	 * @param jedis
	 */
	public static void close(Jedis jedis) {
		if (jedis != null) {
			jedis.close();
		}
	}

	public static void main(String[] args) throws InterruptedException, Parameter_Exception {
		// group msg 选取用户
//		 System.out.println(findGroupIdValue(123456, 123456));
		// Thread.sleep(4000);
//		 System.out.println(findGroupIdValue(123456, 123456));
		// Thread.sleep(6000);
//		 System.out.println(findGroupIdValue(123456, 123457));
		// 存储group switch
		switchGroupUid(123, 1234, 1);
	}

	/**
	 * group msg 选取用户
	 * 
	 * @param gid
	 * @param uid
	 * @return
	 * @throws Parameter_Exception
	 */
	public static Integer findGroupIdValue(Integer gid, Integer uid) throws Parameter_Exception {
		LogUtil.info("dadaGroupBackup-gid:" + gid + ",uid:" + uid);
		Jedis jedis = null;
		try {
			// 获取一个连接
			new Test();
			jedis = getJedis();
			jedis.select(15);
			String key = gid.toString();
			String oldUid = jedis.get(key);
			System.out.println("oldUid:" + oldUid);
			if (oldUid == null) {
				jedis.setex(key, 5, uid.toString());
				return uid;
			}
			Integer olduid = Integer.valueOf(oldUid);
			if (olduid.equals(uid)) {
				jedis.expire(key, 5);
				return olduid;
			}
			LogUtil.info("dadaGroupBackup-gid-uid");
			throw new Parameter_Exception(20031);
		} catch (Parameter_Exception e) {
			e.printStackTrace();
			throw e;
		} catch (Exception e) {
			e.printStackTrace();
			throw new Parameter_Exception(10002);
		} finally {
			close(jedis);
		}
	}

	public static Integer switchGroupUid(Integer gid, Integer uid, Integer state) throws Parameter_Exception {
		LogUtil.info("dadaGroupBackup-gid:" + gid + ",uid:" + uid + ",state:" + state);
		Jedis jedis = null;
		Transaction transaction = null;
		try {
			// 获取一个连接
			new Test();
			jedis = getJedis();
			jedis.select(15);
			transaction = jedis.multi();
			String key = gid.toString();
			transaction.hset(key, uid.toString(), state.toString());
			System.out.println(transaction.hget(key, uid.toString()));
			List<Object> list = transaction.exec();
			System.out.println(list);
			return null;
		} catch (Exception e) {
			e.printStackTrace();
			throw new Parameter_Exception(10002);
		} finally {
			try {
				transaction.close();
				close(jedis);
			} catch (IOException e) {
				e.printStackTrace();
			}
		}
	}

	public void testHSetAndHGetAndHDel() {
		// hset 将哈希表 key 中的域 field 的值设为 value 。
		// hget 返回哈希表 key 中给定域 field 的值。
		// hdel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略。
		Assert.assertTrue(jedis.hset("hash", "key", "value") == 1);
		Assert.assertTrue(jedis.hget("hash", "key").equals("value"));
		Assert.assertTrue(jedis.hdel("hash", "key") == 1);
		Assert.assertTrue(jedis.hget("hash", "key") == null);
	}

	public void testHExists() {
		// hexists 查看哈希表 key 中，给定域 field 是否存在。
		jedis.hset("hash", "key", "value");
		Assert.assertTrue(jedis.hexists("hash", "key"));
		jedis.hdel("hash", "key");
		Assert.assertTrue(!jedis.hexists("hash", "key"));
	}

	public void testHGetAll() {
		// 返回哈希表 key 中，所有的域和值。
		Pipeline pipeline = jedis.pipelined();// 流水线一次性提交
		for (int i = 0; i < 10; i++) {
			pipeline.hset("hash", "key" + i, String.valueOf(i));
		}
		pipeline.sync();
		Map<String, String> map = jedis.hgetAll("hash");
		System.out.println(map.toString());
		Assert.assertTrue(map.size() == 10); // 还有一个住的
		Assert.assertTrue(jedis.del("hash") == 1);
	}

	public void testHIncrbyAndHincrFloat() {
		// hincrBy 为哈希表 key 中的域 field 的值加上增量 increment。如果域 field 不存在，域的值先被初始化为 0 。
		// hincrByFloat 同上，支持浮点数
		jedis.hincrBy("hash", "key", 5);
		Assert.assertTrue(jedis.hincrBy("hash", "key", 5) == 10);
		Assert.assertTrue(jedis.hincrByFloat("hash", "key", 2.5) == 12.5);
		Assert.assertTrue(jedis.del("hash") == 1);
	}

	public void testHKeys() {
		// hkeys 返回哈希表 key 中的所有域。
		// hlen 返回哈希表 key 中域的数量。
		Pipeline pipeline = jedis.pipelined();// 流水线一次性提交
		for (int i = 0; i < 10; i++) {
			pipeline.hset("hash", "key" + i, String.valueOf(i));
		}
		pipeline.sync();
		Set<String> keys = jedis.hkeys("hash");
		for (String s : keys) {
			System.out.println(s);// 乱序的
		}
		Assert.assertTrue(jedis.hlen("hash") == 10);
		Assert.assertTrue(keys.size() == 10);
		Assert.assertTrue(jedis.del("hash") == 1);
	}

	public void testHMSetAndHMGet() {
		Map<String, String> map = new HashMap<>();
		for (int i = 0; i < 10; i++) {
			map.put("key" + i, String.valueOf(i));
		}
		jedis.hmset("hash", map);
		Assert.assertTrue(jedis.hlen("hash") == 10);
		List<String> list = jedis.hmget("hash", "key1", "key2", "key0");
		System.out.println(list.toString());
		Assert.assertTrue(list.size() == 3);
		Assert.assertTrue(jedis.del("hash") == 1);
	}

	public void testHSetNX() {
		// hsetnx 将哈希表 key 中的域 field 的值设置为 value ，当且仅当域 field 不存在。
		Assert.assertTrue(jedis.hsetnx("hash", "key", "value") == 1);
		Assert.assertTrue(jedis.hsetnx("hash", "key", "value") == 0);
		Assert.assertTrue(jedis.del("hash") == 1);
	}

	public void testHVals() {
		// hvals 返回哈希表 key 中所有域的值。
		Map<String, String> map = new HashMap<>();
		for (int i = 0; i < 10; i++) {
			map.put("key" + i, String.valueOf(i));
		}
		jedis.hmset("hash", map);
		List<String> list = jedis.hvals("hash");
		System.out.println(list.toString());
		Assert.assertTrue(list.size() == 10);
		Assert.assertTrue(jedis.del("hash") == 1);
	}

	@SuppressWarnings("deprecation")
	public void testHScan() {
		// HSCAN 命令用于迭代哈希键中的键值对。
		Map<String, String> data = new HashMap<>();
		for (int i = 0; i < 1000; i++) {
			data.put("key" + i, String.valueOf(i));
		}
		jedis.hmset("hash", data);
		ScanResult<Map.Entry<String, String>> result;// = jedis.hscan("hash",DATASOURCE_SELECT);
		int count = 0;
		int cursor = 0;
		do {
			result = jedis.hscan("hash", cursor);
			cursor = Integer.valueOf(result.getStringCursor());
			for (Map.Entry<String, String> map : result.getResult()) {
				System.out.println(map.getKey() + ":" + map.getValue());
				count++;
			}
		} while (cursor != 0);
		Assert.assertTrue(count == 1000);
		Assert.assertTrue(jedis.del("hash") == 1);
	}

	public void testHStrLen() {
		// 返回哈希表 key 中， 与给定域 field 相关联的值的字符串长度（string length）。
		System.out.println("jedis没有HSTRLEN命令！");
	}
}
