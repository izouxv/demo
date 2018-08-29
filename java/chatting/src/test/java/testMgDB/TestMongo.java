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

import com.mongodb.MongoClient;
import com.mongodb.client.MongoDatabase;

/**
 * @Title TestMongo.java
 * @Package Demo
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月27日 上午11:46:51
 * @explain
 */

public class TestMongo {

	private static MongoClient mongoClient;

	public static void main(String[] args) {
		try {
			mongoClient = new MongoClient("localhost", 27017);
			System.out.println("mongoClient:"+mongoClient.toString());
			// 连接到数据库
			MongoDatabase mongoDatabase = mongoClient.getDatabase("test");
			System.out.println("Connect to database successfully"+mongoDatabase.toString());

		} catch (Exception e) {
			System.err.println(e.getClass().getName() + ": " + e.getMessage());
		}
	}

}
