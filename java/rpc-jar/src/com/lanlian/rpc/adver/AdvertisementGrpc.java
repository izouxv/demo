package com.lanlian.rpc.adver;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 * <pre>
 * The Id service definition.
 * </pre>
 */
@javax.annotation.Generated(value = "by gRPC proto compiler (version 1.0.0)", comments = "Source: advertisement.proto")
public class AdvertisementGrpc {

	private AdvertisementGrpc() {
	}

	public static final String SERVICE_NAME = "adv.Advertisement";

	// Static method descriptors that strictly reflect the proto.
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply> METHOD_NEW_ADVERTISEMENT = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("adv.Advertisement", "NewAdvertisement"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply> METHOD_UPDATE_ADVERTISEMENT = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("adv.Advertisement", "UpdateAdvertisement"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply> METHOD_GET_ADVERTISEMENT = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("adv.Advertisement", "GetAdvertisement"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.MapAdvertisementReply> METHOD_GET_ALL_ADVERTISEMENT = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("adv.Advertisement", "GetAllAdvertisement"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.MapAdvertisementReply.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply> METHOD_DEL_ADVERTISEMENT = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("adv.Advertisement", "DelAdvertisement"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.adver.AdvertisementReply.getDefaultInstance()));

	/**
	 * Creates a new async stub that supports all call types for the service
	 */
	public static AdvertisementStub newStub(io.grpc.Channel channel) {
		return new AdvertisementStub(channel);
	}

	/**
	 * Creates a new blocking-style stub that supports unary and streaming output
	 * calls on the service
	 */
	public static AdvertisementBlockingStub newBlockingStub(io.grpc.Channel channel) {
		return new AdvertisementBlockingStub(channel);
	}

	/**
	 * Creates a new ListenableFuture-style stub that supports unary and streaming
	 * output calls on the service
	 */
	public static AdvertisementFutureStub newFutureStub(io.grpc.Channel channel) {
		return new AdvertisementFutureStub(channel);
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static abstract class AdvertisementImplBase implements io.grpc.BindableService {

		/**
		 */
		public void newAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_NEW_ADVERTISEMENT, responseObserver);
		}

		/**
		 */
		public void updateAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_ADVERTISEMENT, responseObserver);
		}

		/**
		 */
		public void getAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_ADVERTISEMENT, responseObserver);
		}

		/**
		 */
		public void getAllAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.MapAdvertisementReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_ALL_ADVERTISEMENT, responseObserver);
		}

		/**
		 */
		public void delAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_DEL_ADVERTISEMENT, responseObserver);
		}

		@java.lang.Override
		public io.grpc.ServerServiceDefinition bindService() {
			return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
					.addMethod(METHOD_NEW_ADVERTISEMENT, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply>(
									this, METHODID_NEW_ADVERTISEMENT)))
					.addMethod(METHOD_UPDATE_ADVERTISEMENT, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply>(
									this, METHODID_UPDATE_ADVERTISEMENT)))
					.addMethod(METHOD_GET_ADVERTISEMENT, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply>(
									this, METHODID_GET_ADVERTISEMENT)))
					.addMethod(METHOD_GET_ALL_ADVERTISEMENT, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.MapAdvertisementReply>(
									this, METHODID_GET_ALL_ADVERTISEMENT)))
					.addMethod(METHOD_DEL_ADVERTISEMENT, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.adver.AdvertisementRequest, com.lanlian.rpc.adver.AdvertisementReply>(
									this, METHODID_DEL_ADVERTISEMENT)))
					.build();
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class AdvertisementStub extends io.grpc.stub.AbstractStub<AdvertisementStub> {
		private AdvertisementStub(io.grpc.Channel channel) {
			super(channel);
		}

		private AdvertisementStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected AdvertisementStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new AdvertisementStub(channel, callOptions);
		}

		/**
		 */
		public void newAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_NEW_ADVERTISEMENT, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void updateAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_ADVERTISEMENT, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void getAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_ADVERTISEMENT, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getAllAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.MapAdvertisementReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_ALL_ADVERTISEMENT, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void delAdvertisement(com.lanlian.rpc.adver.AdvertisementRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_DEL_ADVERTISEMENT, getCallOptions()), request, responseObserver);
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class AdvertisementBlockingStub extends io.grpc.stub.AbstractStub<AdvertisementBlockingStub> {
		private AdvertisementBlockingStub(io.grpc.Channel channel) {
			super(channel);
		}

		private AdvertisementBlockingStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected AdvertisementBlockingStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new AdvertisementBlockingStub(channel, callOptions);
		}

		/**
		 */
		public com.lanlian.rpc.adver.AdvertisementReply newAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_NEW_ADVERTISEMENT, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.adver.AdvertisementReply updateAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_ADVERTISEMENT, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.adver.AdvertisementReply getAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_ADVERTISEMENT, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.adver.MapAdvertisementReply getAllAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_ALL_ADVERTISEMENT, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.adver.AdvertisementReply delAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_DEL_ADVERTISEMENT, getCallOptions(), request);
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class AdvertisementFutureStub extends io.grpc.stub.AbstractStub<AdvertisementFutureStub> {
		private AdvertisementFutureStub(io.grpc.Channel channel) {
			super(channel);
		}

		private AdvertisementFutureStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected AdvertisementFutureStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new AdvertisementFutureStub(channel, callOptions);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.adver.AdvertisementReply> newAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_NEW_ADVERTISEMENT, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.adver.AdvertisementReply> updateAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_ADVERTISEMENT, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.adver.AdvertisementReply> getAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_ADVERTISEMENT, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.adver.MapAdvertisementReply> getAllAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_ALL_ADVERTISEMENT, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.adver.AdvertisementReply> delAdvertisement(
				com.lanlian.rpc.adver.AdvertisementRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_DEL_ADVERTISEMENT, getCallOptions()), request);
		}
	}

	private static final int METHODID_NEW_ADVERTISEMENT = 0;
	private static final int METHODID_UPDATE_ADVERTISEMENT = 1;
	private static final int METHODID_GET_ADVERTISEMENT = 2;
	private static final int METHODID_GET_ALL_ADVERTISEMENT = 3;
	private static final int METHODID_DEL_ADVERTISEMENT = 4;

	private static class MethodHandlers<Req, Resp> implements io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
		private final AdvertisementImplBase serviceImpl;
		private final int methodId;

		public MethodHandlers(AdvertisementImplBase serviceImpl, int methodId) {
			this.serviceImpl = serviceImpl;
			this.methodId = methodId;
		}

		@java.lang.Override
		@java.lang.SuppressWarnings("unchecked")
		public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			case METHODID_NEW_ADVERTISEMENT:
				serviceImpl.newAdvertisement((com.lanlian.rpc.adver.AdvertisementRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply>) responseObserver);
				break;
			case METHODID_UPDATE_ADVERTISEMENT:
				serviceImpl.updateAdvertisement((com.lanlian.rpc.adver.AdvertisementRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply>) responseObserver);
				break;
			case METHODID_GET_ADVERTISEMENT:
				serviceImpl.getAdvertisement((com.lanlian.rpc.adver.AdvertisementRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply>) responseObserver);
				break;
			case METHODID_GET_ALL_ADVERTISEMENT:
				serviceImpl.getAllAdvertisement((com.lanlian.rpc.adver.AdvertisementRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.MapAdvertisementReply>) responseObserver);
				break;
			case METHODID_DEL_ADVERTISEMENT:
				serviceImpl.delAdvertisement((com.lanlian.rpc.adver.AdvertisementRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.adver.AdvertisementReply>) responseObserver);
				break;
			default:
				throw new AssertionError();
			}
		}

		@java.lang.Override
		public io.grpc.stub.StreamObserver<Req> invoke(io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			default:
				throw new AssertionError();
			}
		}
	}

	public static io.grpc.ServiceDescriptor getServiceDescriptor() {
		return new io.grpc.ServiceDescriptor(SERVICE_NAME, METHOD_NEW_ADVERTISEMENT, METHOD_UPDATE_ADVERTISEMENT,
				METHOD_GET_ADVERTISEMENT, METHOD_GET_ALL_ADVERTISEMENT, METHOD_DEL_ADVERTISEMENT);
	}

}
