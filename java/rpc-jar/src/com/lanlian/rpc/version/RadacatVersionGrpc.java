package com.lanlian.rpc.version;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(value = "by gRPC proto compiler (version 1.0.0)", comments = "Source: radacat_version.proto")
public class RadacatVersionGrpc {

	private RadacatVersionGrpc() {
	}

	public static final String SERVICE_NAME = "setting.RadacatVersion";

	// Static method descriptors that strictly reflect the proto.
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.version.AddNewVersionRequest, com.lanlian.rpc.version.AddNewVersionResponse> METHOD_ADD_NEW_VERSION = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("setting.RadacatVersion", "AddNewVersion"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.AddNewVersionRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.AddNewVersionResponse.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.version.GetAllVersionsRequest, com.lanlian.rpc.version.GetAllVersionsResponse> METHOD_GET_ALL_VERSIONS = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("setting.RadacatVersion", "GetAllVersions"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.GetAllVersionsRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.GetAllVersionsResponse.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.version.GetLatestVersionRequest, com.lanlian.rpc.version.GetLatestVersionResponse> METHOD_GET_LATEST_VERSION = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("setting.RadacatVersion", "GetLatestVersion"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.GetLatestVersionRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.GetLatestVersionResponse.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.version.UpdateVersionRequest, com.lanlian.rpc.version.UpdateVersionResponse> METHOD_UPDATE_VERSION = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("setting.RadacatVersion", "UpdateVersion"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.UpdateVersionRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.UpdateVersionResponse.getDefaultInstance()));
	@io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
	public static final io.grpc.MethodDescriptor<com.lanlian.rpc.version.DeleteVersionRequest, com.lanlian.rpc.version.DeleteVersionResponse> METHOD_DELETE_VERSION = io.grpc.MethodDescriptor
			.create(io.grpc.MethodDescriptor.MethodType.UNARY,
					generateFullMethodName("setting.RadacatVersion", "DeleteVersion"),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.DeleteVersionRequest.getDefaultInstance()),
					io.grpc.protobuf.ProtoUtils
							.marshaller(com.lanlian.rpc.version.DeleteVersionResponse.getDefaultInstance()));

	/**
	 * Creates a new async stub that supports all call types for the service
	 */
	public static RadacatVersionStub newStub(io.grpc.Channel channel) {
		return new RadacatVersionStub(channel);
	}

	/**
	 * Creates a new blocking-style stub that supports unary and streaming output
	 * calls on the service
	 */
	public static RadacatVersionBlockingStub newBlockingStub(io.grpc.Channel channel) {
		return new RadacatVersionBlockingStub(channel);
	}

	/**
	 * Creates a new ListenableFuture-style stub that supports unary and streaming
	 * output calls on the service
	 */
	public static RadacatVersionFutureStub newFutureStub(io.grpc.Channel channel) {
		return new RadacatVersionFutureStub(channel);
	}

	/**
	 */
	public static abstract class RadacatVersionImplBase implements io.grpc.BindableService {

		/**
		 */
		public void addNewVersion(com.lanlian.rpc.version.AddNewVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.AddNewVersionResponse> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_ADD_NEW_VERSION, responseObserver);
		}

		/**
		 */
		public void getAllVersions(com.lanlian.rpc.version.GetAllVersionsRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.GetAllVersionsResponse> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_ALL_VERSIONS, responseObserver);
		}

		/**
		 */
		public void getLatestVersion(com.lanlian.rpc.version.GetLatestVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.GetLatestVersionResponse> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_GET_LATEST_VERSION, responseObserver);
		}

		/**
		 */
		public void updateVersion(com.lanlian.rpc.version.UpdateVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.UpdateVersionResponse> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_UPDATE_VERSION, responseObserver);
		}

		/**
		 */
		public void deleteVersion(com.lanlian.rpc.version.DeleteVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.DeleteVersionResponse> responseObserver) {
			asyncUnimplementedUnaryCall(METHOD_DELETE_VERSION, responseObserver);
		}

		@java.lang.Override
		public io.grpc.ServerServiceDefinition bindService() {
			return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
					.addMethod(METHOD_ADD_NEW_VERSION, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.version.AddNewVersionRequest, com.lanlian.rpc.version.AddNewVersionResponse>(
									this, METHODID_ADD_NEW_VERSION)))
					.addMethod(METHOD_GET_ALL_VERSIONS, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.version.GetAllVersionsRequest, com.lanlian.rpc.version.GetAllVersionsResponse>(
									this, METHODID_GET_ALL_VERSIONS)))
					.addMethod(METHOD_GET_LATEST_VERSION, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.version.GetLatestVersionRequest, com.lanlian.rpc.version.GetLatestVersionResponse>(
									this, METHODID_GET_LATEST_VERSION)))
					.addMethod(METHOD_UPDATE_VERSION, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.version.UpdateVersionRequest, com.lanlian.rpc.version.UpdateVersionResponse>(
									this, METHODID_UPDATE_VERSION)))
					.addMethod(METHOD_DELETE_VERSION, asyncUnaryCall(
							new MethodHandlers<com.lanlian.rpc.version.DeleteVersionRequest, com.lanlian.rpc.version.DeleteVersionResponse>(
									this, METHODID_DELETE_VERSION)))
					.build();
		}
	}

	/**
	 */
	public static final class RadacatVersionStub extends io.grpc.stub.AbstractStub<RadacatVersionStub> {
		private RadacatVersionStub(io.grpc.Channel channel) {
			super(channel);
		}

		private RadacatVersionStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected RadacatVersionStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new RadacatVersionStub(channel, callOptions);
		}

		/**
		 */
		public void addNewVersion(com.lanlian.rpc.version.AddNewVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.AddNewVersionResponse> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_ADD_NEW_VERSION, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getAllVersions(com.lanlian.rpc.version.GetAllVersionsRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.GetAllVersionsResponse> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_ALL_VERSIONS, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void getLatestVersion(com.lanlian.rpc.version.GetLatestVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.GetLatestVersionResponse> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_GET_LATEST_VERSION, getCallOptions()), request,
					responseObserver);
		}

		/**
		 */
		public void updateVersion(com.lanlian.rpc.version.UpdateVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.UpdateVersionResponse> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_UPDATE_VERSION, getCallOptions()), request, responseObserver);
		}

		/**
		 */
		public void deleteVersion(com.lanlian.rpc.version.DeleteVersionRequest request,
				io.grpc.stub.StreamObserver<com.lanlian.rpc.version.DeleteVersionResponse> responseObserver) {
			asyncUnaryCall(getChannel().newCall(METHOD_DELETE_VERSION, getCallOptions()), request, responseObserver);
		}
	}

	/**
	 */
	public static final class RadacatVersionBlockingStub extends io.grpc.stub.AbstractStub<RadacatVersionBlockingStub> {
		private RadacatVersionBlockingStub(io.grpc.Channel channel) {
			super(channel);
		}

		private RadacatVersionBlockingStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected RadacatVersionBlockingStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new RadacatVersionBlockingStub(channel, callOptions);
		}

		/**
		 */
		public com.lanlian.rpc.version.AddNewVersionResponse addNewVersion(
				com.lanlian.rpc.version.AddNewVersionRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_ADD_NEW_VERSION, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.version.GetAllVersionsResponse getAllVersions(
				com.lanlian.rpc.version.GetAllVersionsRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_ALL_VERSIONS, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.version.GetLatestVersionResponse getLatestVersion(
				com.lanlian.rpc.version.GetLatestVersionRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_GET_LATEST_VERSION, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.version.UpdateVersionResponse updateVersion(
				com.lanlian.rpc.version.UpdateVersionRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_UPDATE_VERSION, getCallOptions(), request);
		}

		/**
		 */
		public com.lanlian.rpc.version.DeleteVersionResponse deleteVersion(
				com.lanlian.rpc.version.DeleteVersionRequest request) {
			return blockingUnaryCall(getChannel(), METHOD_DELETE_VERSION, getCallOptions(), request);
		}
	}

	/**
	 */
	public static final class RadacatVersionFutureStub extends io.grpc.stub.AbstractStub<RadacatVersionFutureStub> {
		private RadacatVersionFutureStub(io.grpc.Channel channel) {
			super(channel);
		}

		private RadacatVersionFutureStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			super(channel, callOptions);
		}

		@java.lang.Override
		protected RadacatVersionFutureStub build(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
			return new RadacatVersionFutureStub(channel, callOptions);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.version.AddNewVersionResponse> addNewVersion(
				com.lanlian.rpc.version.AddNewVersionRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_ADD_NEW_VERSION, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.version.GetAllVersionsResponse> getAllVersions(
				com.lanlian.rpc.version.GetAllVersionsRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_ALL_VERSIONS, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.version.GetLatestVersionResponse> getLatestVersion(
				com.lanlian.rpc.version.GetLatestVersionRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_GET_LATEST_VERSION, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.version.UpdateVersionResponse> updateVersion(
				com.lanlian.rpc.version.UpdateVersionRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_UPDATE_VERSION, getCallOptions()), request);
		}

		/**
		 */
		public com.google.common.util.concurrent.ListenableFuture<com.lanlian.rpc.version.DeleteVersionResponse> deleteVersion(
				com.lanlian.rpc.version.DeleteVersionRequest request) {
			return futureUnaryCall(getChannel().newCall(METHOD_DELETE_VERSION, getCallOptions()), request);
		}
	}

	private static final int METHODID_ADD_NEW_VERSION = 0;
	private static final int METHODID_GET_ALL_VERSIONS = 1;
	private static final int METHODID_GET_LATEST_VERSION = 2;
	private static final int METHODID_UPDATE_VERSION = 3;
	private static final int METHODID_DELETE_VERSION = 4;

	private static class MethodHandlers<Req, Resp> implements io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
			io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
		private final RadacatVersionImplBase serviceImpl;
		private final int methodId;

		public MethodHandlers(RadacatVersionImplBase serviceImpl, int methodId) {
			this.serviceImpl = serviceImpl;
			this.methodId = methodId;
		}

		@java.lang.Override
		@java.lang.SuppressWarnings("unchecked")
		public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
			switch (methodId) {
			case METHODID_ADD_NEW_VERSION:
				serviceImpl.addNewVersion((com.lanlian.rpc.version.AddNewVersionRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.version.AddNewVersionResponse>) responseObserver);
				break;
			case METHODID_GET_ALL_VERSIONS:
				serviceImpl.getAllVersions((com.lanlian.rpc.version.GetAllVersionsRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.version.GetAllVersionsResponse>) responseObserver);
				break;
			case METHODID_GET_LATEST_VERSION:
				serviceImpl.getLatestVersion((com.lanlian.rpc.version.GetLatestVersionRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.version.GetLatestVersionResponse>) responseObserver);
				break;
			case METHODID_UPDATE_VERSION:
				serviceImpl.updateVersion((com.lanlian.rpc.version.UpdateVersionRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.version.UpdateVersionResponse>) responseObserver);
				break;
			case METHODID_DELETE_VERSION:
				serviceImpl.deleteVersion((com.lanlian.rpc.version.DeleteVersionRequest) request,
						(io.grpc.stub.StreamObserver<com.lanlian.rpc.version.DeleteVersionResponse>) responseObserver);
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
		return new io.grpc.ServiceDescriptor(SERVICE_NAME, METHOD_ADD_NEW_VERSION, METHOD_GET_ALL_VERSIONS,
				METHOD_GET_LATEST_VERSION, METHOD_UPDATE_VERSION, METHOD_DELETE_VERSION);
	}

}
