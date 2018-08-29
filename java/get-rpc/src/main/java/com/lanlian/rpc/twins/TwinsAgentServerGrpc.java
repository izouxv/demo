package com.lanlian.rpc.twins;

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
 */
@javax.annotation.Generated(value = "by gRPC proto compiler (version 1.0.0)", comments = "Source: twins-agent.proto")
public class TwinsAgentServerGrpc {

	private TwinsAgentServerGrpc() {
	}

	public static final String SERVICE_NAME = "api.TwinsAgentServer";

	// Static method descriptors that strictly reflect the proto.
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.twins.AddTwinsAgentRequest, com.lanlian.rpc.twins.AddTwinsAgentResponse> METHOD_ADD_TWINS_AGENT = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("api.TwinsAgentServer", "AddTwinsAgent"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.twins.AddTwinsAgentRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.twins.AddTwinsAgentResponse.getDefaultInstance()));

	/**
	 * Creates a new async stub that supports all call types for the service
	 */
	public static TwinsAgentServerStub newStub(io.grpc.Channel channel) {
		return new TwinsAgentServerStub(channel);
	}

	/**
	 * Creates a new blocking-style stub that supports unary and streaming output
	 * calls on the service
	 */
	public static TwinsAgentServerBlockingStub newBlockingStub(io.grpc.Channel channel) {
		return new TwinsAgentServerBlockingStub(channel);
	}

	/**
	 * Creates a new ListenableFuture-style stub that supports unary and streaming
	 * output calls on the service
	 */
	public static TwinsAgentServerFutureStub newFutureStub(io.grpc.Channel channel) {
		return new TwinsAgentServerFutureStub(channel);
	}

	/**
	 */
	public static abstract class TwinsAgentServerImplBase implements io.grpc.BindableService {

		/**
		 * <pre>
		 *添加设备影子,给radacat设备临时异步调用
		 * </pre>
		 */
		public void addTwinsAgent(com.lanlian.rpc.twins.AddTwinsAgentRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.twins.AddTwinsAgentResponse> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_ADD_TWINS_AGENT, responseObserver);
		}

		@java.lang.Override
		public io.grpc.ServerServiceDefinition bindService() {
			return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
					.addMethod(METHOD_ADD_TWINS_AGENT, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.twins.AddTwinsAgentRequest, com.lanlian.rpc.twins.AddTwinsAgentResponse>(
									this, METHODID_ADD_TWINS_AGENT)))
					.build();
		}
	}

	/**
	 */
	public static final class TwinsAgentServerStub extends io.grpc.stub.AbstractStub<TwinsAgentServerStub> {
		private TwinsAgentServerStub(io.grpc.Channel channel) {
			super(channel);
		}

		private TwinsAgentServerStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected TwinsAgentServerStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new TwinsAgentServerStub(channel, callOptions);
		}

		/**
		 * <pre>
		 *添加设备影子,给radacat设备临时异步调用
		 * </pre>
		 */
		public void addTwinsAgent(com.lanlian.rpc.twins.AddTwinsAgentRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.twins.AddTwinsAgentResponse> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_ADD_TWINS_AGENT, getCallOptions()), request, responseObserver);
		}
	}

	/**
	 */
	public static final class TwinsAgentServerBlockingStub
			extends io.grpc.stub.AbstractStub<TwinsAgentServerBlockingStub> {
		private TwinsAgentServerBlockingStub(io.grpc.Channel channel) {
			super(channel);
		}

		private TwinsAgentServerBlockingStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected TwinsAgentServerBlockingStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new TwinsAgentServerBlockingStub(channel, callOptions);
		}

		/**
		 * <pre>
		 *添加设备影子,给radacat设备临时异步调用
		 * </pre>
		 */
		public com.lanlian.rpc.twins.AddTwinsAgentResponse addTwinsAgent(
				com.lanlian.rpc.twins.AddTwinsAgentRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_ADD_TWINS_AGENT, getCallOptions(), request);
		}
	}

	/**
	 */
	public static final class TwinsAgentServerFutureStub extends io.grpc.stub.AbstractStub<TwinsAgentServerFutureStub> {
		private TwinsAgentServerFutureStub(io.grpc.Channel channel) {
			super(channel);
		}

		private TwinsAgentServerFutureStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected TwinsAgentServerFutureStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new TwinsAgentServerFutureStub(channel, callOptions);
		}

		/**
		 * <pre>
		 *添加设备影子,给radacat设备临时异步调用
		 * </pre>
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.twins.AddTwinsAgentResponse> addTwinsAgent(
				com.lanlian.rpc.twins.AddTwinsAgentRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_ADD_TWINS_AGENT, getCallOptions()), request);
		}
	}

	private static final int METHODID_ADD_TWINS_AGENT = 0;

	private static class MethodHandlers<Req, Resp> implements io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
		private final TwinsAgentServerImplBase serviceImpl;
		private final int methodId;

		public MethodHandlers(TwinsAgentServerImplBase serviceImpl, int methodId) {
			this.serviceImpl = serviceImpl;
			this.methodId = methodId;
		}

		@java.lang.Override
		@java.lang.SuppressWarnings("unchecked")
		public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			case METHODID_ADD_TWINS_AGENT:
				serviceImpl.addTwinsAgent((com.lanlian.rpc.twins.AddTwinsAgentRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.twins.AddTwinsAgentResponse>) responseObserver);
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
		return new io.grpc.ServiceDescriptor(SERVICE_NAME, METHOD_ADD_TWINS_AGENT);
	}

}
