package com.lanlian.rpc.feedback;

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
@javax.annotation.Generated(value = "by gRPC proto compiler (version 1.0.0)", comments = "Source: feedback.proto")
public class FeedBackGrpc {

	private FeedBackGrpc() {
	}

	public static final String SERVICE_NAME = "feedback.FeedBack";

	// Static method descriptors that strictly reflect the proto.
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.feedback.AddFeedbackRequest, com.lanlian.rpc.feedback.AddFeedbackReply> METHOD_ADD_FEEDBACK = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("feedback.FeedBack", "AddFeedback"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.feedback.AddFeedbackRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.feedback.AddFeedbackReply.getDefaultInstance()));

	/**
	 * Creates a new async stub that supports all call types for the service
	 */
	public static FeedBackStub newStub(io.grpc.Channel channel) {
		return new FeedBackStub(channel);
	}

	/**
	 * Creates a new blocking-style stub that supports unary and streaming output
	 * calls on the service
	 */
	public static FeedBackBlockingStub newBlockingStub(io.grpc.Channel channel) {
		return new FeedBackBlockingStub(channel);
	}

	/**
	 * Creates a new ListenableFuture-style stub that supports unary and streaming
	 * output calls on the service
	 */
	public static FeedBackFutureStub newFutureStub(io.grpc.Channel channel) {
		return new FeedBackFutureStub(channel);
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static abstract class FeedBackImplBase implements io.grpc.BindableService {

		/**
		 */
		public void addFeedback(com.lanlian.rpc.feedback.AddFeedbackRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.feedback.AddFeedbackReply> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_ADD_FEEDBACK, responseObserver);
		}

		@java.lang.Override
		public io.grpc.ServerServiceDefinition bindService() {
			return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
					.addMethod(METHOD_ADD_FEEDBACK, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.feedback.AddFeedbackRequest, com.lanlian.rpc.feedback.AddFeedbackReply>(
									this, METHODID_ADD_FEEDBACK)))
					.build();
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class FeedBackStub extends io.grpc.stub.AbstractStub<FeedBackStub> {
		private FeedBackStub(io.grpc.Channel channel) {
			super(channel);
		}

		private FeedBackStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected FeedBackStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new FeedBackStub(channel, callOptions);
		}

		/**
		 */
		public void addFeedback(com.lanlian.rpc.feedback.AddFeedbackRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.feedback.AddFeedbackReply> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_ADD_FEEDBACK, getCallOptions()), request, responseObserver);
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class FeedBackBlockingStub extends io.grpc.stub.AbstractStub<FeedBackBlockingStub> {
		private FeedBackBlockingStub(io.grpc.Channel channel) {
			super(channel);
		}

		private FeedBackBlockingStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected FeedBackBlockingStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new FeedBackBlockingStub(channel, callOptions);
		}

		/**
		 */
		public com.lanlian.rpc.feedback.AddFeedbackReply addFeedback(
				com.lanlian.rpc.feedback.AddFeedbackRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_ADD_FEEDBACK, getCallOptions(), request);
		}
	}

	/**
	 * <pre>
	 * The Id service definition.
	 * </pre>
	 */
	public static final class FeedBackFutureStub extends io.grpc.stub.AbstractStub<FeedBackFutureStub> {
		private FeedBackFutureStub(io.grpc.Channel channel) {
			super(channel);
		}

		private FeedBackFutureStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected FeedBackFutureStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new FeedBackFutureStub(channel, callOptions);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.feedback.AddFeedbackReply> addFeedback(
				com.lanlian.rpc.feedback.AddFeedbackRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_ADD_FEEDBACK, getCallOptions()), request);
		}
	}

	private static final int METHODID_ADD_FEEDBACK = 0;

	private static class MethodHandlers<Req, Resp> implements io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
		private final FeedBackImplBase serviceImpl;
		private final int methodId;

		public MethodHandlers(FeedBackImplBase serviceImpl, int methodId) {
			this.serviceImpl = serviceImpl;
			this.methodId = methodId;
		}

		@java.lang.Override
		@java.lang.SuppressWarnings("unchecked")
		public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			case METHODID_ADD_FEEDBACK:
				serviceImpl.addFeedback((com.lanlian.rpc.feedback.AddFeedbackRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.feedback.AddFeedbackReply>) responseObserver);
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
		return new io.grpc.ServiceDescriptor(SERVICE_NAME, METHOD_ADD_FEEDBACK);
	}

}
