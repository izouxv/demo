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
package com.lanlian.chatting.mongo;

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
import org.springframework.data.mongodb.core.MongoTemplate;

import com.mongodb.Block;
import com.mongodb.DBObject;
import com.mongodb.MongoClient;
import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import com.mongodb.client.gridfs.GridFSBucket;
import com.mongodb.client.gridfs.GridFSBuckets;
import com.mongodb.client.gridfs.GridFSFindIterable;
import com.mongodb.client.gridfs.model.GridFSFile;
import com.mongodb.client.model.Filters;
import com.mongodb.util.JSON;

/**
 * @Title Mongo.java
 * @Package testMgDB
 * @author 王东阳
 * @version V1.0
 * @date 2017年4月27日 下午6:23:43
 * @explain
 */

public class Mongo {

	@Resource
	private MongoTemplate mongoTemplate;

	private static MongoClient MONGOCLIENT;
	private static String MONGOHOST = "127.0.0.1";
	private static int MONGOPORT = 27017;
	private static String DBNAME = "test";

	// 使用spring整合的话, 就直接注入就可以了, 这是测试uanjing
//	@Before
//	public void testBefore() {
//		System.err.println("111");
//		ClassPathXmlApplicationContext context = new ClassPathXmlApplicationContext("spring-mongo.xml");
//		System.err.println("222");
//		mongoTemplate = (MongoTemplate) context.getBean("mongoTemplate");
//		System.err.println("333");
//	}



	public static void main(String[] args) throws Exception {
		long time1 = System.currentTimeMillis();
		String path = "E://a.txt";
		// 上传文件
		 for (int i = 0; i < 400; i++) {
		 upload(path);
		 }
		// 下载文件
		// download(objectId, path);
		// 删除文件
		// delete(objectId);
		// 查询单个文件
		// findOne(objectId);
		// 查询所有文件
		// find();
		json();
		long time2 = System.currentTimeMillis();
		System.out.println("time:" + (time2 - time1));
	}

	/**
	 * 连接服务器、端口、数据库
	 * 
	 * @param MONGOHOST
	 * @param MONGOPORT
	 * @param DBNAME
	 * @return
	 */
	public static MongoDatabase connect() {
		MONGOCLIENT = new MongoClient(MONGOHOST, MONGOPORT);
		return MONGOCLIENT.getDatabase(DBNAME);
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
		MongoCollection<Document> doc = connect().getCollection("fs.files");

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
			MONGOCLIENT = new MongoClient("localhost", 27017);
			System.out.println("连接server:" + MONGOCLIENT.getConnectPoint());
			// 连接到数据库
			MongoDatabase mongoDatabase = MONGOCLIENT.getDatabase("gridfs");
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

	static void json() {
		// 构造一个Json字符串
		String json = " {" + " 'school_code' : '111111', " + " 'school_name' : '汉法大学', "
				+ " 'teacher_idcard' : '0000001', " + " 'teacher_name' : 'test' " + " } ";
		MongoDatabase database = MONGOCLIENT.getDatabase("test");
		MongoCollection<DBObject> collection = database.getCollection("logs", DBObject.class);
		DBObject bson = (DBObject) JSON.parse(json);
		collection.insertOne(bson);
	}

}
