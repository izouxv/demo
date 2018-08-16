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
//import com.lanlian.chatting.po.SsoPo;
//import com.lanlian.chatting.result.Parameter_Exception;
//import com.lanlian.rpc.sso.AgentInfo;
//import com.lanlian.rpc.sso.MapSsoReply;
//import com.lanlian.rpc.sso.MultiSsoRequest;
//import com.lanlian.rpc.sso.SsoGrpc;
//import com.lanlian.rpc.sso.SsoReply;
//import com.lanlian.rpc.sso.SsoRequest;
//import com.lanlian.rpc.sso.SsoRequest.Builder;
//import com.lanlian.rpc.sso.SsoGrpc.SsoBlockingStub;
//
//import io.grpc.ManagedChannel;
//import io.grpc.ManagedChannelBuilder;
//
///**
// * @Title SsoClient.java
// * @Package com.lanlian.chatting.rpc.clientrpc
// * @author 王东阳
// * @version V1.0
// * @date 2017年6月7日 下午2:38:18
// * @explain
// */
//
//public class SsoClient {
//
//	/**
//	 * sso初始化请求通道
//	 */
//	private static ManagedChannel ssoChannel;
//
//	/**
//	 * sso信息
//	 */
//	private static SsoBlockingStub ssoBlockingStub;
//
//	/**
//	 * 地址
//	 */
//	private String host;
//
//	/**
//	 * 端口
//	 */
//	private int port;
//
//	/**
//	 * 构造
//	 * 
//	 * @param host
//	 * @param port
//	 */
//	private SsoClient(String host, int ssoPort) {
//		ssoChannel = ManagedChannelBuilder.forAddress(host, ssoPort).usePlaintext(true).build();
//		ssoBlockingStub = SsoGrpc.newBlockingStub(ssoChannel).withDeadlineAfter(60, TimeUnit.DAYS);
//	}
//
//	@Override
//	public String toString() {
//		return "SsoClient [host=" + host + ", port=" + port + "]";
//	}
//
//	/**
//	 * 获取一个builder
//	 * 
//	 * @return
//	 */
//	private Builder getBuilder() {
//		Builder builder = SsoRequest.newBuilder();
//		return builder;
//	}
//	
//	/**
//	 * 获取一个builder
//	 * 
//	 * @return
//	 */
//	private com.lanlian.rpc.sso.AgentInfo.Builder getABuilder() {
//		com.lanlian.rpc.sso.AgentInfo.Builder builder = AgentInfo.newBuilder();
//		return builder;
//	}
//
//	/**
//	 * 获取一个批量builder
//	 * 
//	 * @return
//	 */
//	private com.lanlian.rpc.sso.MultiSsoRequest.Builder getBuilders() {
//		com.lanlian.rpc.sso.MultiSsoRequest.Builder builder = MultiSsoRequest.newBuilder();
//		return builder;
//	}
//
//	/**
//	 * 发送短信
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo sendMobileCode(SsoPo ssoPo, int action) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setUsername(ssoPo.getUsername());
//			builder.setSource(ssoPo.getSource());
//			builder.setCodeType(action);
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("RPC-request-sendMobileCode:" + request);
//			SsoReply ssoReply = ssoBlockingStub.sendMobileCode(request);
//			LogUtil.info("RPC-ssoReply-sendMobileCode:" + ssoReply);
//			BeanUtils.copyProperties(ssoReply, ssoPo);
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//	
//	/**
//	 * 发送短信
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo sendMobile(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setUsername(ssoPo.getUsername());
//			builder.setSource(ssoPo.getSource());
//			builder.setCodeType(ssoPo.getCodeType());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("RPC-request-sendMobileCode:" + request);
//			SsoReply ssoReply = ssoBlockingStub.sendMobileCode(request);
//			LogUtil.info("RPC-ssoReply-sendMobileCode:" + ssoReply);
//			BeanUtils.copyProperties(ssoReply, ssoPo);
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 校验短信
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public void verifyMobileCode(String username, String code, int action, String source) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder = SsoRequest.newBuilder();
//			builder.setSource(source);
//			builder.setUsername(username);
//			builder.setCode(code);
//			builder.setCodeType(action);
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("RPC-request-verifyMobileCode:" + request);
//			SsoReply ssoReply = ssoBlockingStub.checkCode(request);
//			LogUtil.info("RPC-ssoReply-sendMobileCode:" + ssoReply);
//			if (ssoReply.getErrorCode() == 33005) {
//				throw new Parameter_Exception(20014);
//			}
//			if (ssoReply.getErrorCode() != 10000) {
//				throw new Parameter_Exception(ssoReply.getErrorCode());
//			}
//		} catch (Parameter_Exception e) {
//			throw e;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 通过用户名查询用户信息（验证账号重复）
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo getUserByName(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setUsername(ssoPo.getUsername());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("SsoRPC-request-getUserByName:" + request);
//			SsoReply ssoReply = ssoBlockingStub.getUserByName(request);
//			LogUtil.info("SsoRPC-ssoReply-getUserByName:" + ssoReply);
//			BeanUtils.copyProperties(ssoReply, ssoPo);
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 添加用户
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo add(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setUsername(ssoPo.getUsername());
//			builder.setPassword(ssoPo.getPassword());
//			builder.setNickname(ssoPo.getNickname());
//			builder.setSalt(ssoPo.getSalt());
//			builder.setState(ssoPo.getState());
//			com.lanlian.rpc.sso.AgentInfo.Builder abuilder = getABuilder();
//			abuilder.setIp(ssoPo.getIp());
//			abuilder.setDevInfo(ssoPo.getLoginDevice());
//			builder.setAgentInfo(abuilder.build());
//			abuilder.clear();
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("ssoPo-add-RPC:" + request);
//			SsoReply ssoReply = ssoBlockingStub.add(request);
//			LogUtil.info("ssoReply-add-RPC:" + ssoReply);
//			ssoPo.setErrorCode(ssoReply.getErrorCode());
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 用户登录
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo login(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setUsername(ssoPo.getUsername());
//			builder.setPassword(ssoPo.getPassword());
//			com.lanlian.rpc.sso.AgentInfo.Builder abuilder = getABuilder();
//			abuilder.setIp(ssoPo.getIp());
//			abuilder.setDevInfo(ssoPo.getLoginDevice());
//			builder.setAgentInfo(abuilder.build());
//			abuilder.clear();
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("ssoPo-login-RPC:" + request);
//			SsoReply ssoReply = ssoBlockingStub.login(request);
//			LogUtil.info("ssoReply-login-RPC:" + ssoReply);
//			BeanUtils.copyProperties(ssoReply, ssoPo);
//			return ssoPo;
//		} catch (Exception e) {
//			LogUtil.error(e);
//			return null;
//		}
//	}
//	
//	/**
//	 * 用户登录
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoReply loginAll(Builder builder) {
//		SsoRequest request = builder.build();
//		builder.clear();
//		LogUtil.info("ssoPo-login-RPC:" + request);
//		SsoReply ssoReply = null;
//		if (!request.getCode().trim().isEmpty()) {
//			ssoReply = ssoBlockingStub.loginWithCode(request);
//		} else {
//			ssoReply = ssoBlockingStub.login(request);
//		}
//		LogUtil.info("ssoReply-login-RPC:" + ssoReply);
//		return ssoReply;
//	}
//
//	/**
//	 * 用户退出
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo logout(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setSessionName(ssoPo.getSessionName());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("ssoPo-logout-RPC:" + request);
//			SsoReply ssoReply = ssoBlockingStub.logout(request);
//			LogUtil.info("ssoReply-logout-RPC:" + ssoReply);
//			BeanUtils.copyProperties(ssoReply, ssoPo);
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 校验密码
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo checkPassword(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setUid(ssoPo.getUid());
//			builder.setPassword(ssoPo.getPassword());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("ssoPo-checkPassword-RPC:" + request);
//			SsoReply ssoReply = ssoBlockingStub.checkPassword(request);
//			LogUtil.info("ssoReply-checkPassword-RPC:" + ssoReply);
//			ssoPo.setErrorCode(ssoReply.getErrorCode());
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 修改密码
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo updatePassword(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setUid(ssoPo.getUid());
//			builder.setPassword(ssoPo.getPassword());
//			builder.setSalt(ssoPo.getSalt());
//			builder.setSessionName(ssoPo.getSessionName());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("ssoPo-updatePassword-RPC:" + request);
//			SsoReply ssoReply = ssoBlockingStub.updatePassword(request);
//			LogUtil.info("ssoReply-updatePassword-RPC:" + ssoReply);
//			ssoPo.setErrorCode(ssoReply.getErrorCode());
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 忘记密码
//	 * 
//	 * @param ssoPo
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo updatePasswordByName(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setUsername(ssoPo.getUsername());
//			builder.setPassword(ssoPo.getPassword());
//			builder.setSalt(ssoPo.getSalt());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("ssoPo-updatePasswordByName-RPC:" + request);
//			SsoReply ssoReply = ssoBlockingStub.updatePasswordByName(request);
//			LogUtil.info("ssoReply-updatePasswordByName-RPC:" + ssoReply);
//			BeanUtils.copyProperties(ssoReply, ssoPo);
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 获取单个用户信息
//	 * 
//	 * @param sessionName
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo getSsoInfo(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setSessionName(ssoPo.getSessionName());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("session-getSsoInfo:" + request);
//			SsoReply ssoReply = ssoBlockingStub.getUserInfo(request);
//			LogUtil.info("ssoReply-getSsoInfo:" + ssoReply);
//			ssoPo.setUid(ssoReply.getUid());
//			ssoPo.setUsername(ssoReply.getUsername());
//			ssoPo.setLoginState(ssoReply.getLoginState());
//			ssoPo.setState(ssoReply.getState());
//			ssoPo.setErrorCode(ssoReply.getErrorCode());
//			ssoPo.setNickname(ssoReply.getNickname());
//			return ssoPo;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 批量查询用户信息
//	 * 
//	 * @param pids
//	 * @return
//	 * @throws Parameter_Exception
//	 */
//	public List<Integer> getBatchSsoInfos(List<Integer> uids, String source) throws Parameter_Exception {
//		try {
//			com.lanlian.rpc.sso.MultiSsoRequest.Builder builders = getBuilders();
//			Map<Integer, SsoRequest> values = new HashMap<>();
//			Builder builder = null;
//			SsoRequest request = null;
//			int size = uids.size();
//			for (int i = 1; i <= size; i++) {
//				builder = getBuilder();
//				request = builder.build();
//				builder.clear();
//				values.put(uids.get(i-1), request);
//			}
//			builders.setSource(source);
//			builders.putAllSsos(values);
//			MultiSsoRequest requests = builders.build();
//			builders.clear();
//			LogUtil.info("request-getBatchSsoInfos-RPC:" + request);
//			MapSsoReply ssoReplyMap = ssoBlockingStub.getBatchSsoInfos(requests);
//			LogUtil.info("ssoReplyMap-getBatchSsoInfos-RPC:" + ssoReplyMap);
//			int code = ssoReplyMap.getErrorCode();
//			if (code != 10000) {
//				throw new Parameter_Exception(code);
//			}
//			Map<Integer, SsoReply> map = ssoReplyMap.getSsos();
//			uids.clear();
//			for (Entry<Integer, SsoReply> entryMap : map.entrySet()) {
//				SsoReply ssoReply = entryMap.getValue();
//				// System.out.println("ssoReply:" + ssoReply);
//				if (3 == ssoReply.getState()) {
//					uids.add(entryMap.getKey());
//				}
//			}
//			return uids;
//		} catch (Parameter_Exception e) {
//			throw e;
//		} catch (Exception e) {
//			throw new Parameter_Exception(10002);
//		}
//	}
//
//	/**
//	 * 发送邮件
//	 * 
//	 * @param sessionName
//	 * @throws Parameter_Exception
//	 */
//	public SsoPo sendEmail(SsoPo ssoPo) throws Parameter_Exception {
//		try {
//			Builder builder = getBuilder();
//			builder.setSource(ssoPo.getSource());
//			builder.setUsername(ssoPo.getUsername());
//			SsoRequest request = builder.build();
//			builder.clear();
//			LogUtil.info("sendEmail-request:" + request);
//			SsoReply ssoReply = ssoBlockingStub.findPasswordByMail(request);
//			LogUtil.info("sendEmail-getSsoInfo:" + ssoReply);
//			ssoPo.setErrorCode(ssoReply.getErrorCode());
//			return ssoPo;
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
//			ssoChannel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
//		} catch (InterruptedException e) {
//			LogUtil.error(e);
//		} finally {
//			ssoChannel.shutdownNow();
//		}
//	}
//
//	public static void main(String[] args) throws Parameter_Exception, InterruptedException {
////		long time1 = System.currentTimeMillis();
//		SsoClient clientRpc = new SsoClient("192.168.1.178", 8003);
//		// SsoPo ssoPo = new SsoPo();
//
//		// 检验是否被注册
//		// ssoPo.setUsername("17600117963");
//		// System.out.println(clientRpc.getUserByName(ssoPo));
//
//		// 登录
//		// ssoPo.setUsername("17600117962");
//		// ssoPo.setPassword("e10adc3949ba59abbe56e057f20f883e");
//		// ssoPo = clientRpc.login(ssoPo);
//		// System.out.println("ssoPo:"+ssoPo);
//
//		// 获取信息
//		// ssoPo = clientRpc.getSsoInfo("1908f8cb36bec156d6a0e7ebeb08f2d2");
//		// System.out.println(ssoPo);
//
//		// 批量查询用户
//		 List<Integer> list = new ArrayList<>();
//		 list.add(1000001);
//		 list.add(1000002);
//		 list.add(536870915);
//		 list = clientRpc.getBatchSsoInfos(list, "BAIDAA==");
//		 System.out.println("list:" + list);
//		 clientRpc.shutdown();
//		
//		//发送验证码
////		SsoPo ssoPo = new SsoPo();
////		ssoPo.setUsername("17600117962");
////		ssoPo.setSource("BAIDAA==");
////		ssoPo.setCodeType(3);
////		clientRpc.sendMobile(ssoPo);
//		
//		//登录新
////		Builder builder = SsoRequest.newBuilder();
////		builder.setSource("BAIDAA==");
////		builder.setUsername("17600117962");
////		builder.setPassword("e10adc3949ba59abbe56e057f20f883e");
////		builder.setCode("999930");
////		clientRpc.loginAll(builder);
//	}
//
//}
