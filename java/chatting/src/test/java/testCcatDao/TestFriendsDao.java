package testCcatDao;

import java.util.ArrayList;
import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.lanlian.chatting.dao.FriendsDao;
import com.lanlian.chatting.po.FriendsPo;

/**
 * @author wdyqxx
 * @version 2017年1月4日 上午11:23:25
 * @explain 此类用于测试UserContactsDao接口的所有方法；
 */
public class TestFriendsDao {

	ClassPathXmlApplicationContext ctx;
	FriendsDao friendsDao;

	@Before
	public void init() {
		ctx = new ClassPathXmlApplicationContext("spring-mybatis.xml");
		friendsDao = ctx.getBean("friendsDao", FriendsDao.class);
	}

	/**
	 * 查询好友
	 */
	@Test
	public void findMyFriends() {
		FriendsPo friendsPo = new FriendsPo();
		friendsPo.setUid1(1);
		List<FriendsPo> list = friendsDao.findFriends(friendsPo);
		System.out.println(list);
	}

	/**
	 * 添加好友
	 */
	@Test
	public void saveMyFriends() {
		int uid1 = 1000001;
		int uid2 = 1000002;
		int uid3 = 1000003;
		String note2 = "laohua";
		String note3 = "laobin";
		FriendsPo friendsPo = new FriendsPo();
		friendsPo.setUid1(uid1);
		friendsPo.setNote1(note2);
		friendsPo.setUid2(uid2);
		
		FriendsPo friendsPo1 = new FriendsPo();
		friendsPo1.setUid1(uid1);
		friendsPo1.setNote1(note3);
		friendsPo1.setUid2(uid3);

		List<FriendsPo> list = new ArrayList<>();
		list.add(friendsPo);
		list.add(friendsPo1);
		friendsDao.saveFriends(list);
	}

	/**
	 * 校验好友
	 */
	@Test
	public void checkFriends() {
		int uid1 = 1000001;
		int uid2 = 1000002;
		int uid3 = 1000003;
		FriendsPo friendsPo = new FriendsPo();
		friendsPo.setUid1(uid1);
		friendsPo.setUid2(uid2);
		
		FriendsPo friendsPo1 = new FriendsPo();
		friendsPo1.setUid1(uid2);
		friendsPo1.setUid2(uid3);

		List<FriendsPo> list = new ArrayList<>();
		list.add(friendsPo);
		list.add(friendsPo1);
		System.out.println("list:"+list);
		List<FriendsPo> list2 = friendsDao.verifyFriends(list);
		System.out.println(list2);
	}

	/**
	 * 删除好友
	 */
	@Test
	public void deleteMyFriends() {
		FriendsPo friendsPo1 = new FriendsPo();
		friendsPo1.setUid1(1);
		friendsPo1.setUid2(2);
		friendsPo1.setState(2);
		FriendsPo friendsPo2 = new FriendsPo();
		friendsPo2.setUid1(1);
		friendsPo2.setUid2(3);
		friendsPo2.setState(2);
		List<FriendsPo> list = new ArrayList<>();
		list.add(friendsPo1);
		list.add(friendsPo2);
		friendsDao.deleteFriends(list);
	}
	
	/**
	 * 修改好友信息
	 */
	@Test
	public void modifyFriends() {
//		FriendsPo friendsPo = new FriendsPo();
//		friendsPo.setUid1(1000001);
//		List<FriendsPo> list = friendsDao.findFriends(friendsPo);
//		System.out.println("list1:"+list);
//		
//		for (FriendsPo f1 : list) {
//			if (f1.getUid2() == 1000002) {				
//				f1.setNote1("qqqqq");
//			}
//			if (f1.getUid2() == 1000003) {				
//				f1.setNote1("wwwww");
//			}
//		}
		List<FriendsPo> list = new ArrayList<>();
		FriendsPo f1 = new FriendsPo();
		f1.setUid1(1000001);
		f1.setNote1("22222");
		f1.setUid2(1000002);
//		f1.setNote2("11111");
		
		FriendsPo f2 = new FriendsPo();
		f2.setUid1(1000001);
		f2.setNote1("33333");
		f2.setUid2(1000003);
//		f2.setNote2("2222222");
		list.add(f1);
		list.add(f2);
		System.out.println("list2:"+list);
		friendsDao.modifyFriends(list);
	}
	
	public static void main(String[] args) {
		
//		int gid;//群ID,参数在URI请求接口中后缀(http://ip/chatting/v1.0/message/groupid/{gid})
//		int mid;//消息ID
//		int uid;//用户ID
//		long sendTime;//发送时间
//		String info;//消息内容
//		String json = "[{\"mid\":\"1\","
//				+ "\"uid\":\"1\","
//				+ "\"sendTime\":\"1498618646625\","
//				+ "\"info\":\"大家好\"},"
//				
//				+ "{\"mid\":\"2\","
//				+ "\"uid\":\"2\","
//				+ "\"sendTime\":\"1498618647625\","
//				+ "\"info\":\"你好\"}]";
		
	}

}
