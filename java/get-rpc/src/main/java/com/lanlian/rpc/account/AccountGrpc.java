package com.lanlian.rpc.account;

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
 * The Id service definition.
 * </pre>
 */
@javax.annotation.Generated(value = "by gRPC proto compiler (version 1.0.0)", comments = "Source: account.proto")
public class AccountGrpc {

	private AccountGrpc() {
	}

	public static final String SERVICE_NAME = "pb.Account";

	// Static method descriptors that strictly reflect the proto.
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_UPDATE_EX_INFO = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "UpdateExInfo"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_SHOW = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "Show"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_GET_USER_INFO_BY_ID = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "GetUserInfoById"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_GET_USER_INFO_ALL = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "GetUserInfoAll"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_UPDATE_ACCOUNT_INFO = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("pb.Account", "UpdateAccountInfo"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_UPDATE_CERTIFICATION = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("pb.Account", "UpdateCertification"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_GET_CERTIFICATION = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "GetCertification"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_UPDATE_USER_VALUES = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "UpdateUserValues"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply> METHOD_GET_USER_VALUES = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "GetUserValues"),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils.marshaller(com.lanlian.rpc.account.AccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.MultiAccountRequest, com.lanlian.rpc.account.MapAccountReply> METHOD_GET_BATCH_ACCOUNT_INFO = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("pb.Account", "GetBatchAccountInfo"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.account.MultiAccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.account.MapAccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.MultiAccountRequest, com.lanlian.rpc.account.MapAccountReply> METHOD_GET_BATCH_EX_INFO = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY, generateFullMethodName("pb.Account", "GetBatchExInfo"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.account.MultiAccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.account.MapAccountReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.account.MultiAccountRequest, com.lanlian.rpc.account.MapAccountReply> METHOD_GET_BATCH_ALL_USER_INFO = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("pb.Account", "GetBatchAllUserInfo"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.account.MultiAccountRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.account.MapAccountReply.getDefaultInstance()));

	/**
	 * Creates a new async stub that supports all call types for the service
	 */
	public static AccountStub newStub(io.grpc.Channel channel) {
		return new AccountStub(channel);
	}

	/**
	 * Creates a new blocking-style stub that supports unary and streaming output
	 * calls on the service
	 */
	public static AccountBlockingStub newBlockingStub(io.grpc.Channel channel) {
		return new AccountBlockingStub(channel);
	}

	/**
	 * Creates a new ListenableFuture-style stub that supports unary and streaming
	 * output calls on the service
	 */
	public static AccountFutureStub newFutureStub(io.grpc.Channel channel) {
		return new AccountFutureStub(channel);
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static abstract class AccountImplBase implements io.grpc.BindableService {

		/**
		 */
		public void updateExInfo(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_EX_INFO, responseObserver);
		}

		/**
		 */
		public void show(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_SHOW, responseObserver);
		}

		/**
		 */
		public void getUserInfoById(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_USER_INFO_BY_ID, responseObserver);
		}

		/**
		 */
		public void getUserInfoAll(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_USER_INFO_ALL, responseObserver);
		}

		/**
		 */
		public void updateAccountInfo(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_ACCOUNT_INFO, responseObserver);
		}

		/**
		 */
		public void updateCertification(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_CERTIFICATION, responseObserver);
		}

		/**
		 */
		public void getCertification(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_CERTIFICATION, responseObserver);
		}

		/**
		 */
		public void updateUserValues(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_USER_VALUES, responseObserver);
		}

		/**
		 */
		public void getUserValues(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_USER_VALUES, responseObserver);
		}

		/**
		 */
		public void getBatchAccountInfo(com.lanlian.rpc.account.MultiAccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_BATCH_ACCOUNT_INFO, responseObserver);
		}

		/**
		 */
		public void getBatchExInfo(com.lanlian.rpc.account.MultiAccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_BATCH_EX_INFO, responseObserver);
		}

		/**
		 */
		public void getBatchAllUserInfo(com.lanlian.rpc.account.MultiAccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_BATCH_ALL_USER_INFO, responseObserver);
		}

		@java.lang.Override
		public io.grpc.ServerServiceDefinition bindService() {
			return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
					.addMethod(METHOD_UPDATE_EX_INFO, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_UPDATE_EX_INFO)))
					.addMethod(METHOD_SHOW, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_SHOW)))
					.addMethod(METHOD_GET_USER_INFO_BY_ID, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_GET_USER_INFO_BY_ID)))
					.addMethod(METHOD_GET_USER_INFO_ALL, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_GET_USER_INFO_ALL)))
					.addMethod(METHOD_UPDATE_ACCOUNT_INFO, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_UPDATE_ACCOUNT_INFO)))
					.addMethod(METHOD_UPDATE_CERTIFICATION, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_UPDATE_CERTIFICATION)))
					.addMethod(METHOD_GET_CERTIFICATION, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_GET_CERTIFICATION)))
					.addMethod(METHOD_UPDATE_USER_VALUES, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_UPDATE_USER_VALUES)))
					.addMethod(METHOD_GET_USER_VALUES, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.AccountRequest, com.lanlian.rpc.account.AccountReply>(
									this, METHODID_GET_USER_VALUES)))
					.addMethod(METHOD_GET_BATCH_ACCOUNT_INFO, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.MultiAccountRequest, com.lanlian.rpc.account.MapAccountReply>(
									this, METHODID_GET_BATCH_ACCOUNT_INFO)))
					.addMethod(METHOD_GET_BATCH_EX_INFO, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.MultiAccountRequest, com.lanlian.rpc.account.MapAccountReply>(
									this, METHODID_GET_BATCH_EX_INFO)))
					.addMethod(METHOD_GET_BATCH_ALL_USER_INFO, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.account.MultiAccountRequest, com.lanlian.rpc.account.MapAccountReply>(
									this, METHODID_GET_BATCH_ALL_USER_INFO)))
					.build();
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class AccountStub extends io.grpc.stub.AbstractStub<AccountStub> {
		private AccountStub(io.grpc.Channel channel) {
			super(channel);
		}

		private AccountStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected AccountStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new AccountStub(channel, callOptions);
		}

		/**
		 */
		public void updateExInfo(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_EX_INFO, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void show(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_SHOW, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getUserInfoById(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_USER_INFO_BY_ID, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void getUserInfoAll(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_USER_INFO_ALL, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void updateAccountInfo(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_ACCOUNT_INFO, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void updateCertification(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_CERTIFICATION, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void getCertification(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_CERTIFICATION, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void updateUserValues(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_USER_VALUES, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void getUserValues(com.lanlian.rpc.account.AccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_USER_VALUES, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getBatchAccountInfo(com.lanlian.rpc.account.MultiAccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_BATCH_ACCOUNT_INFO, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void getBatchExInfo(com.lanlian.rpc.account.MultiAccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_BATCH_EX_INFO, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getBatchAllUserInfo(com.lanlian.rpc.account.MultiAccountRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_BATCH_ALL_USER_INFO, getCallOptions()), request,
					responseObserver);
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class AccountBlockingStub extends io.grpc.stub.AbstractStub<AccountBlockingStub> {
		private AccountBlockingStub(io.grpc.Channel channel) {
			super(channel);
		}

		private AccountBlockingStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected AccountBlockingStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new AccountBlockingStub(channel, callOptions);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply updateExInfo(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_EX_INFO, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply show(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_SHOW, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply getUserInfoById(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_USER_INFO_BY_ID, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply getUserInfoAll(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_USER_INFO_ALL, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply updateAccountInfo(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_ACCOUNT_INFO, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply updateCertification(
				com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_CERTIFICATION, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply getCertification(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_CERTIFICATION, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply updateUserValues(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_USER_VALUES, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.AccountReply getUserValues(com.lanlian.rpc.account.AccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_USER_VALUES, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.MapAccountReply getBatchAccountInfo(
				com.lanlian.rpc.account.MultiAccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_BATCH_ACCOUNT_INFO, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.MapAccountReply getBatchExInfo(
				com.lanlian.rpc.account.MultiAccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_BATCH_EX_INFO, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.account.MapAccountReply getBatchAllUserInfo(
				com.lanlian.rpc.account.MultiAccountRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_BATCH_ALL_USER_INFO, getCallOptions(), request);
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class AccountFutureStub extends io.grpc.stub.AbstractStub<AccountFutureStub> {
		private AccountFutureStub(io.grpc.Channel channel) {
			super(channel);
		}

		private AccountFutureStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected AccountFutureStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new AccountFutureStub(channel, callOptions);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> updateExInfo(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_EX_INFO, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> show(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_SHOW, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> getUserInfoById(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_USER_INFO_BY_ID, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> getUserInfoAll(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_USER_INFO_ALL, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> updateAccountInfo(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_ACCOUNT_INFO, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> updateCertification(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_CERTIFICATION, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> getCertification(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_CERTIFICATION, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> updateUserValues(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_USER_VALUES, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.AccountReply> getUserValues(
				com.lanlian.rpc.account.AccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_USER_VALUES, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.MapAccountReply> getBatchAccountInfo(
				com.lanlian.rpc.account.MultiAccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_BATCH_ACCOUNT_INFO, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.MapAccountReply> getBatchExInfo(
				com.lanlian.rpc.account.MultiAccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_BATCH_EX_INFO, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.account.MapAccountReply> getBatchAllUserInfo(
				com.lanlian.rpc.account.MultiAccountRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_BATCH_ALL_USER_INFO, getCallOptions()), request);
		}
	}

	private static final int METHODID_UPDATE_EX_INFO = 0;
	private static final int METHODID_SHOW = 1;
	private static final int METHODID_GET_USER_INFO_BY_ID = 2;
	private static final int METHODID_GET_USER_INFO_ALL = 3;
	private static final int METHODID_UPDATE_ACCOUNT_INFO = 4;
	private static final int METHODID_UPDATE_CERTIFICATION = 5;
	private static final int METHODID_GET_CERTIFICATION = 6;
	private static final int METHODID_UPDATE_USER_VALUES = 7;
	private static final int METHODID_GET_USER_VALUES = 8;
	private static final int METHODID_GET_BATCH_ACCOUNT_INFO = 9;
	private static final int METHODID_GET_BATCH_EX_INFO = 10;
	private static final int METHODID_GET_BATCH_ALL_USER_INFO = 11;

	private static class MethodHandlers<Req, Resp> implements io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
		private final AccountImplBase serviceImpl;
		private final int methodId;

		public MethodHandlers(AccountImplBase serviceImpl, int methodId) {
			this.serviceImpl = serviceImpl;
			this.methodId = methodId;
		}

		@java.lang.Override
		@java.lang.SuppressWarnings("unchecked")
		public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			case METHODID_UPDATE_EX_INFO:
				serviceImpl.updateExInfo((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_SHOW:
				serviceImpl.show((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_GET_USER_INFO_BY_ID:
				serviceImpl.getUserInfoById((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_GET_USER_INFO_ALL:
				serviceImpl.getUserInfoAll((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_UPDATE_ACCOUNT_INFO:
				serviceImpl.updateAccountInfo((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_UPDATE_CERTIFICATION:
				serviceImpl.updateCertification((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_GET_CERTIFICATION:
				serviceImpl.getCertification((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_UPDATE_USER_VALUES:
				serviceImpl.updateUserValues((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_GET_USER_VALUES:
				serviceImpl.getUserValues((com.lanlian.rpc.account.AccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.AccountReply>) responseObserver);
				break;
			case METHODID_GET_BATCH_ACCOUNT_INFO:
				serviceImpl.getBatchAccountInfo((com.lanlian.rpc.account.MultiAccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply>) responseObserver);
				break;
			case METHODID_GET_BATCH_EX_INFO:
				serviceImpl.getBatchExInfo((com.lanlian.rpc.account.MultiAccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply>) responseObserver);
				break;
			case METHODID_GET_BATCH_ALL_USER_INFO:
				serviceImpl.getBatchAllUserInfo((com.lanlian.rpc.account.MultiAccountRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.account.MapAccountReply>) responseObserver);
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
		return new io.grpc.ServiceDescriptor(SERVICE_NAME, METHOD_UPDATE_EX_INFO, METHOD_SHOW,
				METHOD_GET_USER_INFO_BY_ID, METHOD_GET_USER_INFO_ALL, METHOD_UPDATE_ACCOUNT_INFO,
				METHOD_UPDATE_CERTIFICATION, METHOD_GET_CERTIFICATION, METHOD_UPDATE_USER_VALUES,
				METHOD_GET_USER_VALUES, METHOD_GET_BATCH_ACCOUNT_INFO, METHOD_GET_BATCH_EX_INFO,
				METHOD_GET_BATCH_ALL_USER_INFO);
	}

}
