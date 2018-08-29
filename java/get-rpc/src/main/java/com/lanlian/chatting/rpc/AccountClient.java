///** 
// *<p>开发公司：		蓝涟科技 <p>
// *<p>版权所有：		蓝涟科技 <p>
// *<p>责任人：		      王东阳<p> 
// *<p>网址：www.radacat.com <p>
// * @author wangdyqxx
// * @email wangdy@radacat.com
// */
//
//package com.lanlian.chatting.rpc;
//
//import java.util.ArrayList;
//import java.util.HashMap;
//import java.util.List;
//import java.util.Map;
//import java.util.Map.Entry;
//import java.util.concurrent.TimeUnit;
//
//import org.springframework.beans.BeanUtils;
//
//import com.lanlian.chatting.logger.LogUtil;
//import com.lanlian.chatting.po.AccountPo;
//import com.lanlian.chatting.result.Parameter_Exception;
//import com.lanlian.rpc.account.AccountGrpc;
//import com.lanlian.rpc.account.AccountReply;
//import com.lanlian.rpc.account.AccountRequest;
//import com.lanlian.rpc.account.AccountRequest.Builder;
//import com.lanlian.rpc.account.MapAccountReply;
//import com.lanlian.rpc.account.MultiAccountRequest;
//import com.lanlian.rpc.account.AccountGrpc.AccountBlockingStub;
//
//import io.grpc.ManagedChannel;
//import io.grpc.ManagedChannelBuilder;
//
///**
// * @Title AccountClient.java
// * @Package com.lanlian.chatting.rpc.clientrpc
// * @author 王东阳
// * @version V1.0
// * @date 2017年6月7日 下午2:38:18
// * @explain
// */
//
//public class AccountClient {
//
//	/**
//	 * account初始化请求通道
//	 */
//	private ManagedChannel accountChannel;
//
//	/**
//	 * account信息
//	 */
//	private AccountBlockingStub accountBlockingStub;
//	
//	/**
//	 * 获取一个Builder
//	 * @return
//	 */
//	private Builder getBuilder() {
//		Builder builder = AccountRequest.newBuilder();
//		return builder;
//	}
//	
//	/**
//	 * 获取一个批量Builder
//	 * @return
//	 */
//	private com.lanlian.rpc.account.MultiAccountRequest.Builder getBuilders() {
//		com.lanlian.rpc.account.MultiAccountRequest.Builder builder = MultiAccountRequest.newBuilder();
//		return builder;
//	}
//
//	/**
//	 * 构造
//	 * 
//	 * @param host
//	 * @param port
//	 */
//	private AccountClient(String host, int accountPort) {
//		accountChannel = ManagedChannelBuilder.forAddress(host, accountPort).usePlaintext(true).build();
//		accountBlockingStub = AccountGrpc.newBlockingStub(accountChannel).withDeadlineAfter(60, TimeUnit.DAYS);;
//	}
//
//	/**
//	 * 获取单个用户信息
//	 * 
//	 * @param UID
//	 * @throws Parameter_Exception
//	 */
//	public AccountPo getUserInfo(AccountPo accountPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(accountPo.getSource());
//			builder.setUid(accountPo.getUid());
//			AccountRequest request = builder.build();
//			LogUtil.info("getUserInfo-AccountClient-RPC:" + request);
//			AccountReply accountReply = accountBlockingStub.getUserInfoAll(request);
//			LogUtil.info("getUserInfo-AccountClient-RPC:" + accountReply);
//			BeanUtils.copyProperties(accountReply, accountPo);
//			return accountPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 基于uid插入或修改用户基本信息(account表)
//	 * 
//	 * @param accountPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public AccountPo addAccountInfo(AccountPo accountPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder = AccountRequest.newBuilder();
//			builder.setSource(accountPo.getSource());
//			builder.setUid(accountPo.getUid());
//			builder.setNickname(accountPo.getNickname());
//			AccountRequest request = builder.build();
//			LogUtil.info("request-add-accountClient:" + request);
//			AccountReply accountReply = accountBlockingStub.updateAccountInfo(request);
//			LogUtil.info("accountReply-accountClient:" + accountReply);
//			accountPo.setErrorCode(accountReply.getErrorCode());
//			return accountPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 基于uid插入或修改用户扩展信息(ExInfo表)
//	 * @param accountPo
//	 * @return
//	 * @throws Parameter_Exception 
//	 */
//	public AccountPo updateExInfo(AccountPo accountPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder = AccountRequest.newBuilder();
//			builder.setSource(accountPo.getSource());
//			builder.setUid(accountPo.getUid());
//			builder.setRegIp(accountPo.getRegIp());
//			AccountRequest request = builder.build();
//			LogUtil.info("request-updateExInfo-RPC:" + request);
//			AccountReply accountReply = accountBlockingStub.updateExInfo(request);
//			LogUtil.info("accountReply-updateExInfo-RPC:" + accountReply);
//			accountPo.setErrorCode(accountReply.getErrorCode());
//			return accountPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 添加或修改个人信息(ExInfo表)
//	 * 
//	 * @param accountPo
//	 * @return
//	 * @throws Parameter_Exception 
//	 */
//	public AccountPo updateAccountInfo(AccountPo accountPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder = AccountRequest.newBuilder();
//			builder.setSource(accountPo.getSource());
//			builder.setUid(accountPo.getUid());
//			builder.setPhone(accountPo.getPhone());
//			builder.setEmail(accountPo.getEmail());
//			builder.setNickname(accountPo.getNickname());
//			builder.setGender(accountPo.getGender());
//			builder.setBirthday(accountPo.getBirthday());
//			builder.setAvatar(accountPo.getAvatar());
//			builder.setSignature(accountPo.getSignature());
//			builder.setProvince(accountPo.getProvince());
//			builder.setCity(accountPo.getCity());
//			builder.setUserAddress(accountPo.getUserAddress());
//			AccountRequest request = builder.build();
//			LogUtil.info("updateAccountInfo-request-RPC:" + request);
//			AccountReply accountReply = accountBlockingStub.updateAccountInfo(request);
//			LogUtil.info("updateAccountInfo-accountReply-RPC:" + accountReply);
//			BeanUtils.copyProperties(accountReply, accountPo);
//			return accountPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 获取单个用户扩展信息
//	 * 
//	 * @param pid
//	 * @throws Parameter_Exception
//	 */
//	public AccountPo getUserInfoById(AccountPo accountPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder = AccountRequest.newBuilder();
//			builder.setSource(accountPo.getSource());
//			builder.setUid(accountPo.getUid());
//			AccountRequest request = builder.build();
//			AccountReply accountReply = accountBlockingStub.getUserInfoById(request);
//			LogUtil.info("AccountPo:" + accountReply);
//			if (accountReply.getErrorCode() != 10000) {
//				throw new Parameter_Exception(accountReply.getErrorCode());
//			}
//			BeanUtils.copyProperties(accountReply, accountPo);
//			LogUtil.info("accountPo:" + accountPo);
//			return accountPo;
//		} catch (Parameter_Exception e) {
//			throw e;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 查询实名认证
//	 * 
//	 * @param uid
//	 * @throws Parameter_Exception
//	 */
//	public AccountPo getCertification(AccountPo accountPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder = AccountRequest.newBuilder();
//			builder.setSource(accountPo.getSource());
//			builder.setUid(accountPo.getUid());
//			AccountRequest request = builder.build();
//			LogUtil.info("AccountPo-getCertification-RPC:" + request);
//			AccountReply accountReply = accountBlockingStub.getCertification(request);
//			LogUtil.info("accountReply-getCertification-RPC:" + accountReply);
//			BeanUtils.copyProperties(accountReply, accountPo);
//			return accountPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 用户实名认证
//	 * 
//	 * @param uid
//	 * @throws Parameter_Exception
//	 */
//	public AccountPo updateCertification(AccountPo accountPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder = AccountRequest.newBuilder();
//			builder.setSource(accountPo.getSource());
//			builder.setUid(accountPo.getUid());
//			builder.setRealname(accountPo.getRealname());
//			builder.setIdentityCard(accountPo.getIdentityCard());
//			AccountRequest request = builder.build();
//			LogUtil.info("AccountPo-updateCertification-RPC:" + request);
//			AccountReply accountReply = accountBlockingStub.updateCertification(request);
//			LogUtil.info("accountReply-updateCertification-RPC:" + accountReply);
//			BeanUtils.copyProperties(accountReply, accountPo);
//			return accountPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 获取批量用户信息
//	 * 
//	 * @param uids
//	 * @throws Parameter_Exception
//	 */
//	public List<AccountPo> getBatchAllUserInfo(List<Integer> uids, String source) throws Parameter_Exception {
//		try {
//			com.lanlian.rpc.account.MultiAccountRequest.Builder builderMap = getBuilders();
//			List<AccountPo> AccountPos = new ArrayList<>();
//			builderMap = MultiAccountRequest.newBuilder();
//			Map<Integer, AccountRequest> values = new HashMap<>();
//			for (Integer uid : uids) {
//				values.put(uid, getBuilder().build());
//			}
//			builderMap.setSource(source);
//			builderMap.putAllAccounts(values);
//			MultiAccountRequest request = builderMap.build();
//			LogUtil.info("getBatchAllUserInfo-request:" + request);
//			MapAccountReply accountMap = accountBlockingStub.getBatchAllUserInfo(request);
//			if (accountMap.getErrorCode() != 10000) {
//				throw new Parameter_Exception(accountMap.getErrorCode());
//			}
//			Map<Integer, AccountReply> accountReplyMap = accountMap.getAccounts();
//			LogUtil.info("getBatchAllUserInfo-accountReplyMap:" + accountReplyMap);
//			for (Entry<Integer, AccountReply> accountReplyE : accountReplyMap.entrySet()) {
//				AccountReply accountReply = accountReplyE.getValue();
//				AccountPo accountPo = new AccountPo();
//				BeanUtils.copyProperties(accountReply, accountPo);
//				accountPo.setUid(accountReply.getUid());
//				AccountPos.add(accountPo);
//			}
//			return AccountPos;
//		} catch (Parameter_Exception e) {
//			throw e;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 关闭请求连接
//	 * 
//	 * @throws InterruptedException
//	 */
//	public void shutdown() {
//		try {
//			accountChannel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
//		} catch (InterruptedException e) {
//			LogUtil.error(e);
//		}
//	}
//
//	public static void main(String[] args) throws Parameter_Exception, InterruptedException {
//		long time1 = System.currentTimeMillis();
//		AccountClient clientRpc = new AccountClient("192.168.1.178", 8003);
//		List<Integer> pids = new ArrayList<>();
//		pids.add(1000001);
//		pids.add(1000002);
//		List<AccountPo> batchUserInfo = clientRpc.getBatchAllUserInfo(pids, "BAIDAA==");
//		System.out.println(batchUserInfo);
//		long time2 = System.currentTimeMillis();
//		System.out.println(time2 - time1);
//		clientRpc.shutdown();
//	}
//
//}
