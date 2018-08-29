package com.lanlian.rpc.sso;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 * <pre>
 * The Sso service definition.
 * </pre>
 */
@javax.annotation.Generated(value = "by gRPC proto compiler (version 1.0.0)", comments = "Source: sso.proto")
public class SsoGrpc {

	private SsoGrpc() {
	}

	public static final String SERVICE_NAME = "pb.Sso";

	// Static method descriptors that strictly reflect the proto.
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_GET_USER_INFO = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "GetUserInfo"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_LOGIN = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "Login"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_GET_USER_BY_NAME = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "GetUserByName"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_ADD = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "Add"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_CHECK_PASSWORD = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "CheckPassword"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_UPDATE_PASSWORD = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "UpdatePassword"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_LOGOUT = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "Logout"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_UPDATE_PASSWORD_BY_NAME = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "UpdatePasswordByName"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_FIND_PASSWORD_BY_MAIL = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "FindPasswordByMail"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_RESET_PASSWORD = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "ResetPassword"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_UPDATE_STATE = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "UpdateState"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.MultiSsoRequest, com.lanlian.rpc.sso.MapSsoReply> METHOD_GET_BATCH_SSO_INFOS = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "GetBatchSsoInfos"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.MultiSsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.MapSsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_CHECK_CODE = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "CheckCode"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_RESET_PASSWORD_BY_PHONE = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "ResetPasswordByPhone"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_SEND_MOBILE_CODE = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "SendMobileCode"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply> METHOD_LOGIN_WITH_CODE = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Sso", "LoginWithCode"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.sso.SsoReply.getDefaultInstance()));

	/**
	 * Creates a new async stub that supports all call types for the service
	 */
	public static SsoStub newStub(io.grpc.Channel channel) {
		return new SsoStub(channel);
	}

	/**
	 * Creates a new blocking-style stub that supports unary and streaming output
	 * calls on the service
	 */
	public static SsoBlockingStub newBlockingStub(io.grpc.Channel channel) {
		return new SsoBlockingStub(channel);
	}

	/**
	 * Creates a new ListenableFuture-style stub that supports unary and streaming
	 * output calls on the service
	 */
	public static SsoFutureStub newFutureStub(io.grpc.Channel channel) {
		return new SsoFutureStub(channel);
	}

	/**
	 * <pre>
	 * The Sso service definition.
	 * </pre>
	 */
	public static abstract class SsoImplBase implements io.grpc.BindableService {

		/**
		 */
		public void getUserInfo(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_USER_INFO, responseObserver);
		}

		/**
		 */
		public void login(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_LOGIN, responseObserver);
		}

		/**
		 */
		public void getUserByName(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_USER_BY_NAME, responseObserver);
		}

		/**
		 */
		public void add(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_ADD, responseObserver);
		}

		/**
		 */
		public void checkPassword(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_CHECK_PASSWORD, responseObserver);
		}

		/**
		 */
		public void updatePassword(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_PASSWORD, responseObserver);
		}

		/**
		 */
		public void logout(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_LOGOUT, responseObserver);
		}

		/**
		 */
		public void updatePasswordByName(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_PASSWORD_BY_NAME, responseObserver);
		}

		/**
		 */
		public void findPasswordByMail(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_FIND_PASSWORD_BY_MAIL, responseObserver);
		}

		/**
		 */
		public void resetPassword(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_RESET_PASSWORD, responseObserver);
		}

		/**
		 */
		public void updateState(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_STATE, responseObserver);
		}

		/**
		 */
		public void getBatchSsoInfos(com.lanlian.rpc.sso.MultiSsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.MapSsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_BATCH_SSO_INFOS, responseObserver);
		}

		/**
		 */
		public void checkCode(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_CHECK_CODE, responseObserver);
		}

		/**
		 */
		public void resetPasswordByPhone(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_RESET_PASSWORD_BY_PHONE, responseObserver);
		}

		/**
		 */
		public void sendMobileCode(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_SEND_MOBILE_CODE, responseObserver);
		}

		/**
		 */
		public void loginWithCode(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_LOGIN_WITH_CODE, responseObserver);
		}

		@java.lang.Override
		public io.grpc.ServerServiceDefinition bindService() {
			return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor()).addMethod(METHOD_GET_USER_INFO,
					asyncUnaryCall(new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
							this, METHODID_GET_USER_INFO)))
					.addMethod(METHOD_LOGIN,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_LOGIN)))
					.addMethod(METHOD_GET_USER_BY_NAME,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_GET_USER_BY_NAME)))
					.addMethod(METHOD_ADD,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_ADD)))
					.addMethod(METHOD_CHECK_PASSWORD,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_CHECK_PASSWORD)))
					.addMethod(METHOD_UPDATE_PASSWORD,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_UPDATE_PASSWORD)))
					.addMethod(METHOD_LOGOUT,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_LOGOUT)))
					.addMethod(METHOD_UPDATE_PASSWORD_BY_NAME,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_UPDATE_PASSWORD_BY_NAME)))
					.addMethod(METHOD_FIND_PASSWORD_BY_MAIL,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_FIND_PASSWORD_BY_MAIL)))
					.addMethod(METHOD_RESET_PASSWORD,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_RESET_PASSWORD)))
					.addMethod(METHOD_UPDATE_STATE,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_UPDATE_STATE)))
					.addMethod(METHOD_GET_BATCH_SSO_INFOS, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.sso.MultiSsoRequest, com.lanlian.rpc.sso.MapSsoReply>(
									this, METHODID_GET_BATCH_SSO_INFOS)))
					.addMethod(METHOD_CHECK_CODE,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_CHECK_CODE)))
					.addMethod(METHOD_RESET_PASSWORD_BY_PHONE,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_RESET_PASSWORD_BY_PHONE)))
					.addMethod(METHOD_SEND_MOBILE_CODE,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_SEND_MOBILE_CODE)))
					.addMethod(METHOD_LOGIN_WITH_CODE,
							asyncUnaryCall(
									new MethodHandlers<com.lanlian.rpc.sso.SsoRequest, com.lanlian.rpc.sso.SsoReply>(
											this, METHODID_LOGIN_WITH_CODE)))
					.build();
		}
	}

	/**
	 * <pre>
	 * The Sso service definition.
	 * </pre>
	 */
	public static final class SsoStub extends io.grpc.stub.AbstractStub<SsoStub> {
		private SsoStub(io.grpc.Channel channel) {
			super(channel);
		}

		private SsoStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected SsoStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new SsoStub(channel, callOptions);
		}

		/**
		 */
		public void getUserInfo(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_USER_INFO, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void login(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_LOGIN, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getUserByName(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_USER_BY_NAME, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void add(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_ADD, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void checkPassword(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_CHECK_PASSWORD, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void updatePassword(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_PASSWORD, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void logout(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_LOGOUT, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void updatePasswordByName(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_PASSWORD_BY_NAME, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void findPasswordByMail(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_FIND_PASSWORD_BY_MAIL, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void resetPassword(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_RESET_PASSWORD, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void updateState(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_STATE, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getBatchSsoInfos(com.lanlian.rpc.sso.MultiSsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.MapSsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_BATCH_SSO_INFOS, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void checkCode(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_CHECK_CODE, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void resetPasswordByPhone(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_RESET_PASSWORD_BY_PHONE, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void sendMobileCode(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_SEND_MOBILE_CODE, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void loginWithCode(com.lanlian.rpc.sso.SsoRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_LOGIN_WITH_CODE, getCallOptions()), request, responseObserver);
		}
	}

	/**
	 * <pre>
	 * The Sso service definition.
	 * </pre>
	 */
	public static final class SsoBlockingStub extends io.grpc.stub.AbstractStub<SsoBlockingStub> {
		private SsoBlockingStub(io.grpc.Channel channel) {
			super(channel);
		}

		private SsoBlockingStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected SsoBlockingStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new SsoBlockingStub(channel, callOptions);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply getUserInfo(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_USER_INFO, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply login(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_LOGIN, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply getUserByName(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_USER_BY_NAME, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply add(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_ADD, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply checkPassword(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_CHECK_PASSWORD, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply updatePassword(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_PASSWORD, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply logout(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_LOGOUT, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply updatePasswordByName(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_PASSWORD_BY_NAME, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply findPasswordByMail(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_FIND_PASSWORD_BY_MAIL, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply resetPassword(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_RESET_PASSWORD, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply updateState(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_STATE, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.MapSsoReply getBatchSsoInfos(com.lanlian.rpc.sso.MultiSsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_BATCH_SSO_INFOS, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply checkCode(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_CHECK_CODE, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply resetPasswordByPhone(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_RESET_PASSWORD_BY_PHONE, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply sendMobileCode(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_SEND_MOBILE_CODE, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.sso.SsoReply loginWithCode(com.lanlian.rpc.sso.SsoRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_LOGIN_WITH_CODE, getCallOptions(), request);
		}
	}

	/**
	 * <pre>
	 * The Sso service definition.
	 * </pre>
	 */
	public static final class SsoFutureStub extends io.grpc.stub.AbstractStub<SsoFutureStub> {
		private SsoFutureStub(io.grpc.Channel channel) {
			super(channel);
		}

		private SsoFutureStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected SsoFutureStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new SsoFutureStub(channel, callOptions);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> getUserInfo(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_USER_INFO, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> login(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_LOGIN, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> getUserByName(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_USER_BY_NAME, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> add(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_ADD, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> checkPassword(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_CHECK_PASSWORD, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> updatePassword(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_PASSWORD, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> logout(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_LOGOUT, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> updatePasswordByName(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_PASSWORD_BY_NAME, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> findPasswordByMail(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_FIND_PASSWORD_BY_MAIL, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> resetPassword(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_RESET_PASSWORD, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> updateState(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_STATE, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.MapSsoReply> getBatchSsoInfos(
				com.lanlian.rpc.sso.MultiSsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_BATCH_SSO_INFOS, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> checkCode(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_CHECK_CODE, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> resetPasswordByPhone(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_RESET_PASSWORD_BY_PHONE, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> sendMobileCode(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_SEND_MOBILE_CODE, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.sso.SsoReply> loginWithCode(
				com.lanlian.rpc.sso.SsoRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_LOGIN_WITH_CODE, getCallOptions()), request);
		}
	}

	private static final int METHODID_GET_USER_INFO = 0;
	private static final int METHODID_LOGIN = 1;
	private static final int METHODID_GET_USER_BY_NAME = 2;
	private static final int METHODID_ADD = 3;
	private static final int METHODID_CHECK_PASSWORD = 4;
	private static final int METHODID_UPDATE_PASSWORD = 5;
	private static final int METHODID_LOGOUT = 6;
	private static final int METHODID_UPDATE_PASSWORD_BY_NAME = 7;
	private static final int METHODID_FIND_PASSWORD_BY_MAIL = 8;
	private static final int METHODID_RESET_PASSWORD = 9;
	private static final int METHODID_UPDATE_STATE = 10;
	private static final int METHODID_GET_BATCH_SSO_INFOS = 11;
	private static final int METHODID_CHECK_CODE = 12;
	private static final int METHODID_RESET_PASSWORD_BY_PHONE = 13;
	private static final int METHODID_SEND_MOBILE_CODE = 14;
	private static final int METHODID_LOGIN_WITH_CODE = 15;

	private static class MethodHandlers<Req, Resp> implements io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
		private final SsoImplBase serviceImpl;
		private final int methodId;

		public MethodHandlers(SsoImplBase serviceImpl, int methodId) {
			this.serviceImpl = serviceImpl;
			this.methodId = methodId;
		}

		@java.lang.Override
		@java.lang.SuppressWarnings("unchecked")
		public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			case METHODID_GET_USER_INFO:
				serviceImpl.getUserInfo((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_LOGIN:
				serviceImpl.login((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_GET_USER_BY_NAME:
				serviceImpl.getUserByName((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_ADD:
				serviceImpl.add((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_CHECK_PASSWORD:
				serviceImpl.checkPassword((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_UPDATE_PASSWORD:
				serviceImpl.updatePassword((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_LOGOUT:
				serviceImpl.logout((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_UPDATE_PASSWORD_BY_NAME:
				serviceImpl.updatePasswordByName((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_FIND_PASSWORD_BY_MAIL:
				serviceImpl.findPasswordByMail((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_RESET_PASSWORD:
				serviceImpl.resetPassword((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_UPDATE_STATE:
				serviceImpl.updateState((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_GET_BATCH_SSO_INFOS:
				serviceImpl.getBatchSsoInfos((com.lanlian.rpc.sso.MultiSsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.MapSsoReply>) responseObserver);
				break;
			case METHODID_CHECK_CODE:
				serviceImpl.checkCode((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_RESET_PASSWORD_BY_PHONE:
				serviceImpl.resetPasswordByPhone((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_SEND_MOBILE_CODE:
				serviceImpl.sendMobileCode((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			case METHODID_LOGIN_WITH_CODE:
				serviceImpl.loginWithCode((com.lanlian.rpc.sso.SsoRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.sso.SsoReply>) responseObserver);
				break;
			default:
				throw new AssertionError();
			}
		}

		@java.lang.Override
		@java.lang.SuppressWarnings("unchecked")
		public io.grpc.stub.StreamObserver<Req> invoke(io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			default:
				throw new AssertionError();
			}
		}
	}

	public static io.grpc.ServiceDescriptor getServiceDescriptor() {
		return new io.grpc.ServiceDescriptor(SERVICE_NAME, METHOD_GET_USER_INFO, METHOD_LOGIN, METHOD_GET_USER_BY_NAME,
				METHOD_ADD, METHOD_CHECK_PASSWORD, METHOD_UPDATE_PASSWORD, METHOD_LOGOUT,
				METHOD_UPDATE_PASSWORD_BY_NAME, METHOD_FIND_PASSWORD_BY_MAIL, METHOD_RESET_PASSWORD,
				METHOD_UPDATE_STATE, METHOD_GET_BATCH_SSO_INFOS, METHOD_CHECK_CODE, METHOD_RESET_PASSWORD_BY_PHONE,
				METHOD_SEND_MOBILE_CODE, METHOD_LOGIN_WITH_CODE);
	}

}
