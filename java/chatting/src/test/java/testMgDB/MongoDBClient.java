/** 
 *<p>开发公司：		蓝涟科技 <p>
 *<p>版权所有：		蓝涟科技 <p>
 *<p>责任人：		      王东阳<p> 
 *<p>网址：www.radacat.com <p>
 * @author wangdyqxx
 * @email wangdy@radacat.com
 */

/**
 * 
 */
package testMgDB;

import java.util.ArrayList;
import java.util.List;

import com.mongodb.MongoClient;
import com.mongodb.MongoCredential;
import com.mongodb.ServerAddress;
import com.mongodb.client.MongoDatabase;

/**
 * @Title MongoDBClient.java
 * @Package testMgDB
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月27日 下午6:14:21
 * @explain
 */

public class MongoDBClient {
	
	private static MongoClient mongoClient;

	/**
	 * 需要验证用户名 密码的 MongoDB的连接方式 com.mongodb.MongoClient.getDatabase("数据库名")
	 * 
	 * @return
	 */
	public MongoDatabase getConnection() {
		try {
			// 连接到MongoDB服务 如果是远程连接可以替换“localhost”为服务器所在IP地址
			// ServerAddress()两个参数分别为 服务器地址 和 端口
			ServerAddress serverAddress = new ServerAddress("localhost", 27017);
			List<ServerAddress> addrs = new ArrayList<ServerAddress>();
			addrs.add(serverAddress);

			// MongoCredential.createScramSha1Credential()三个参数分别为 用户名 数据库名称 密码
			MongoCredential credential = MongoCredential.createScramSha1Credential("username", "databaseName",
					"password".toCharArray());
			List<MongoCredential> credentials = new ArrayList<MongoCredential>();
			credentials.add(credential);

			mongoClient = new MongoClient(addrs, credentials);
			// 连接到数据库
			MongoDatabase mongoDatabase = mongoClient.getDatabase("databaseName");
			System.out.println("连接成功");
			return mongoDatabase;
		} catch (Exception e) {
			System.err.println(e.getClass().getName() + ": " + e.getMessage());
		}
		return null;
	}

	/**
	 * 不需要验证 用户名+密码 的获取连接的方式 com.mongodb.MongoClient.getDatabase("数据库名")
	 * 
	 * @return
	 */
	public static MongoDatabase getConnectionBasis() {
		try {
			// 连接到mongodb服务
			mongoClient = new MongoClient("localhost", 27017);
			System.out.println("mongoClient"+mongoClient);
			MongoDatabase mongoDatabase = mongoClient.getDatabase("test");
			System.out.println("连接成功");
			return mongoDatabase;
		} catch (Exception e) {
			System.out.println(e.getClass().getName() + ":" + e.getMessage());
		}
		return null;
	}
	
	public static void main(String[] args) {
		MongoDatabase data = getConnectionBasis();
		System.out.println(data);
	}
}
