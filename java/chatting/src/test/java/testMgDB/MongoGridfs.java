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

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import javax.annotation.Resource;

import org.bson.Document;
import org.bson.types.ObjectId;
import org.junit.Test;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;

import com.mongodb.Block;
import com.mongodb.MongoClient;
import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import com.mongodb.client.gridfs.GridFSBucket;
import com.mongodb.client.gridfs.GridFSBuckets;
import com.mongodb.client.gridfs.GridFSFindIterable;
import com.mongodb.client.gridfs.model.GridFSFile;
import com.mongodb.client.model.Filters;

/**
 * @Title MongoGridfs.java
 * @Package testMgDB
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月27日 下午6:23:43
 * @explain
 */

public class MongoGridfs {

	private static MongoClient mongoClient;
	private static String mongoHost = "192.168.0.114";
	private static int mongoPort = 27017;
	private static String dbName = "test";
	
	@Resource
	private MongoTemplate mongoTemplate;
	private static MongoClient mongoClie;

	/**
	 * 按条件删除数据
	 */
	@Test
	public void removeUser() {
		// 设置删除条件，如果条件内容为空则删除所有
		Query query = new Query();
		Criteria criteria = new Criteria("name");
		criteria.is("zhangguochen");
		query.addCriteria(criteria);
		mongoTemplate.remove(query, "user");
	}

	/**
	 * 修改数据
	 */
	@Test
	public void updateUser() {
		// 设置修改条件
		Query query = new Query();
		Criteria criteria = new Criteria("name");
		criteria.is("wangdy");
		query.addCriteria(criteria);
		// 设置修改内容
		Update update = Update.update("age", 2);
		// 参数：查询条件，更改结果，集合名
		mongoTemplate.updateFirst(query, update, "user");
	}


	public static void main(String[] args) throws Exception {
		long time1 = System.currentTimeMillis();
		String objectId = "590ad2aa42b0b82cd864f02f";
		String path = "E://a.txt";
		// 上传文件
//		for (int i = 0; i < 400; i++) {
//			upload(path);
//		}
		// 下载文件
		 download(objectId, path);
		// 删除文件
		// delete(objectId);
		// 查询单个文件
		// findOne(objectId);
		// 查询所有文件
		// find();
		long time2 = System.currentTimeMillis();
		System.out.println("time:" + (time2 - time1));
	}

	/**
	 * 连接服务器、端口、数据库
	 * 
	 * @param mongoHost
	 * @param mongoPort
	 * @param dbName
	 * @return
	 */
	public static MongoDatabase connect() {
		mongoClient = new MongoClient(mongoHost, mongoPort);
		return mongoClient.getDatabase(dbName);
	}

	/**
	 * 上传文件
	 * 
	 * @throws FileNotFoundException
	 */
	public static void upload(String path) throws FileNotFoundException {
		// 获取文件路径
		File file = new File(path);
		// 获取文件流
		InputStream in = new FileInputStream(file);
		// 创建容器，传入连接db
		GridFSBucket bucket = GridFSBuckets.create(connect());
		// 上传
		ObjectId fileId = bucket.uploadFromStream(UUID.randomUUID().toString(), in);
		System.out.println("上传完成，id：" + fileId);
	}

	/**
	 * 下载文件
	 * 
	 * @throws Exception
	 */
	public static void download(String objectId, String path) throws Exception {
		File file = new File(path);
		// 创建一个容器，传入一个`MongoDatabase`类实例db
		GridFSBucket bucket = GridFSBuckets.create(connect());
		// 创建输出流
		OutputStream os = new FileOutputStream(file);
		// 下载
		bucket.downloadToStream(new ObjectId(objectId), os);
		System.out.println("下载完成。");
	}

	/**
	 * 删除文件(根据文件id)
	 * 
	 * @param objectId
	 * @throws Exception
	 */
	public static void delete(String objectId) throws Exception {
		// 创建一个容器，传入一个`MongoDatabase`类实例db
		GridFSBucket bucket = GridFSBuckets.create(connect());
		// 删除
		bucket.delete(new ObjectId(objectId));
		System.out.println("删除完成。");
	}

	/**
	 * 查询文件
	 * 
	 * @param objectId
	 * @throws Exception
	 */
	public static void findOne(String objectId) throws Exception {
		// 创建一个容器，传入一个`MongoDatabase`类实例db
		GridFSBucket bucket = GridFSBuckets.create(connect());
		// 获取内容
		GridFSFindIterable gridFSFindIterable = bucket.find(Filters.eq("_id", new ObjectId(objectId)));
		GridFSFile gridFSFile = gridFSFindIterable.first();
		System.out.println("filename: " + gridFSFile.getFilename());
	}

	/**
	 * 查询文件
	 */
	public static void find() {
		mongoClie = new MongoClient("192.168.0.114", 27017);
		MongoDatabase db = mongoClie.getDatabase("test");
		MongoCollection<Document> doc = db.getCollection("fs.files");

		FindIterable<Document> iter = doc.find();
		iter.forEach(new Block<Document>() {
			@Override
			public void apply(Document doc) {
				System.out.println(doc.toJson());
				System.out.println(doc.get("_id"));
				System.out.println(doc.get("userId"));
				System.out.println(doc.get("chunkSize"));
				System.out.println(doc.get("md5"));
				System.out.println(doc.get("filename"));
				System.out.println(doc.get("contentType"));
				System.out.println(doc.get("uploadDate"));
				System.out.println(doc.get("aliases"));
			}
		});
	}

	/**
	 * 上传文档
	 */
	void document() {
		try {
			mongoClient = new MongoClient("localhost", 27017);
			System.out.println("连接server:" + mongoClient.getConnectPoint());
			// 连接到数据库
			MongoDatabase mongoDatabase = mongoClient.getDatabase("gridfs");
			System.out.println("连接数据库成功！");
			MongoCollection<Document> collection = mongoDatabase.getCollection("fs.files");
			System.out.println("集合 fs.files 选择成功");
			// 插入文档 1. 创建文档 org.bson.Document参数为key-value的格式 2.
			// 创建文档集合List<Document>
			// 3. 将文档集合插入数据库集合中mongoCollection.insertMany(List // <Document>)
			// 插入单个文档可以用mongoCollection.insertOne(Document)
			Document document = new Document("title", "file").append("description", "文件").append("likes", 100)
					.append("by", "Fly");
			List<Document> documents = new ArrayList<Document>();
			documents.add(document);
			collection.insertMany(documents);
			System.out.println("文档插入成功");
		} catch (Exception e) {
			System.err.println(e.getClass().getName() + ": " + e.getMessage());
		}
	}

}
