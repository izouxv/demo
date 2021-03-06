// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: radacat_version.proto

package com.lanlian.rpc.version;

/**
 * Protobuf type {@code setting.AddNewVersionResponse}
 */
public final class AddNewVersionResponse extends com.google.protobuf.GeneratedMessage implements
		// @@protoc_insertion_point(message_implements:setting.AddNewVersionResponse)
		AddNewVersionResponseOrBuilder {
	// Use AddNewVersionResponse.newBuilder() to construct.
	private AddNewVersionResponse(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
		super(builder);
	}

	private AddNewVersionResponse() {
		errorCode_ = 0;
	}

	@java.lang.Override
	public final com.google.protobuf.UnknownFieldSet getUnknownFields() {
		return com.google.protobuf.UnknownFieldSet.getDefaultInstance();
	}

	@SuppressWarnings("unused")
	private AddNewVersionResponse(com.google.protobuf.CodedInputStream input,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry) {
		this();
		int mutable_bitField0_ = 0;
		try {
			boolean done = false;
			while (!done) {
				int tag = input.readTag();
				switch (tag) {
				case 0:
					done = true;
					break;
				default: {
					if (!input.skipField(tag)) {
						done = true;
					}
					break;
				}
				case 8: {

					errorCode_ = input.readInt32();
					break;
				}
				}
			}
		} catch (com.google.protobuf.InvalidProtocolBufferException e) {
			throw new RuntimeException(e.setUnfinishedMessage(this));
		} catch (java.io.IOException e) {
			throw new RuntimeException(
					new com.google.protobuf.InvalidProtocolBufferException(e.getMessage()).setUnfinishedMessage(this));
		} finally {
			makeExtensionsImmutable();
		}
	}

	public static final com.google.protobuf.Descriptors.Descriptor getDescriptor() {
		return com.lanlian.rpc.version.VersionProto.internal_static_setting_AddNewVersionResponse_descriptor;
	}

	protected com.google.protobuf.GeneratedMessage.FieldAccessorTable internalGetFieldAccessorTable() {
		return com.lanlian.rpc.version.VersionProto.internal_static_setting_AddNewVersionResponse_fieldAccessorTable
				.ensureFieldAccessorsInitialized(com.lanlian.rpc.version.AddNewVersionResponse.class,
						com.lanlian.rpc.version.AddNewVersionResponse.Builder.class);
	}

	public static final int ERRORCODE_FIELD_NUMBER = 1;
	private int errorCode_;

	/**
	 * <code>optional int32 ErrorCode = 1;</code>
	 */
	public int getErrorCode() {
		return errorCode_;
	}

	private byte memoizedIsInitialized = -1;

	public final boolean isInitialized() {
		byte isInitialized = memoizedIsInitialized;
		if (isInitialized == 1)
			return true;
		if (isInitialized == 0)
			return false;

		memoizedIsInitialized = 1;
		return true;
	}

	public void writeTo(com.google.protobuf.CodedOutputStream output) throws java.io.IOException {
		if (errorCode_ != 0) {
			output.writeInt32(1, errorCode_);
		}
	}

	public int getSerializedSize() {
		int size = memoizedSize;
		if (size != -1)
			return size;

		size = 0;
		if (errorCode_ != 0) {
			size += com.google.protobuf.CodedOutputStream.computeInt32Size(1, errorCode_);
		}
		memoizedSize = size;
		return size;
	}

	private static final long serialVersionUID = 0L;

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(com.google.protobuf.ByteString data)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(com.google.protobuf.ByteString data,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data, extensionRegistry);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(byte[] data)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(byte[] data,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data, extensionRegistry);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(java.io.InputStream input)
			throws java.io.IOException {
		return PARSER.parseFrom(input);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(java.io.InputStream input,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry) throws java.io.IOException {
		return PARSER.parseFrom(input, extensionRegistry);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseDelimitedFrom(java.io.InputStream input)
			throws java.io.IOException {
		return PARSER.parseDelimitedFrom(input);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseDelimitedFrom(java.io.InputStream input,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry) throws java.io.IOException {
		return PARSER.parseDelimitedFrom(input, extensionRegistry);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(com.google.protobuf.CodedInputStream input)
			throws java.io.IOException {
		return PARSER.parseFrom(input);
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse parseFrom(com.google.protobuf.CodedInputStream input,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry) throws java.io.IOException {
		return PARSER.parseFrom(input, extensionRegistry);
	}

	public Builder newBuilderForType() {
		return newBuilder();
	}

	public static Builder newBuilder() {
		return DEFAULT_INSTANCE.toBuilder();
	}

	public static Builder newBuilder(com.lanlian.rpc.version.AddNewVersionResponse prototype) {
		return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
	}

	public Builder toBuilder() {
		return this == DEFAULT_INSTANCE ? new Builder() : new Builder().mergeFrom(this);
	}

	@java.lang.Override
	protected Builder newBuilderForType(com.google.protobuf.GeneratedMessage.BuilderParent parent) {
		Builder builder = new Builder(parent);
		return builder;
	}

	/**
	 * Protobuf type {@code setting.AddNewVersionResponse}
	 */
	public static final class Builder extends com.google.protobuf.GeneratedMessage.Builder<Builder> implements
			// @@protoc_insertion_point(builder_implements:setting.AddNewVersionResponse)
			com.lanlian.rpc.version.AddNewVersionResponseOrBuilder {
		public static final com.google.protobuf.Descriptors.Descriptor getDescriptor() {
			return com.lanlian.rpc.version.VersionProto.internal_static_setting_AddNewVersionResponse_descriptor;
		}

		protected com.google.protobuf.GeneratedMessage.FieldAccessorTable internalGetFieldAccessorTable() {
			return com.lanlian.rpc.version.VersionProto.internal_static_setting_AddNewVersionResponse_fieldAccessorTable
					.ensureFieldAccessorsInitialized(com.lanlian.rpc.version.AddNewVersionResponse.class,
							com.lanlian.rpc.version.AddNewVersionResponse.Builder.class);
		}

		// Construct using com.lanlian.rpc.version.AddNewVersionResponse.newBuilder()
		private Builder() {
			maybeForceBuilderInitialization();
		}

		private Builder(com.google.protobuf.GeneratedMessage.BuilderParent parent) {
			super(parent);
			maybeForceBuilderInitialization();
		}

		private void maybeForceBuilderInitialization() {
			if (com.google.protobuf.GeneratedMessage.alwaysUseFieldBuilders) {
			}
		}

		public Builder clear() {
			super.clear();
			errorCode_ = 0;

			return this;
		}

		public com.google.protobuf.Descriptors.Descriptor getDescriptorForType() {
			return com.lanlian.rpc.version.VersionProto.internal_static_setting_AddNewVersionResponse_descriptor;
		}

		public com.lanlian.rpc.version.AddNewVersionResponse getDefaultInstanceForType() {
			return com.lanlian.rpc.version.AddNewVersionResponse.getDefaultInstance();
		}

		public com.lanlian.rpc.version.AddNewVersionResponse build() {
			com.lanlian.rpc.version.AddNewVersionResponse result = buildPartial();
			if (!result.isInitialized()) {
				throw newUninitializedMessageException(result);
			}
			return result;
		}

		public com.lanlian.rpc.version.AddNewVersionResponse buildPartial() {
			com.lanlian.rpc.version.AddNewVersionResponse result = new com.lanlian.rpc.version.AddNewVersionResponse(
					this);
			result.errorCode_ = errorCode_;
			onBuilt();
			return result;
		}

		public Builder mergeFrom(com.google.protobuf.Message other) {
			if (other instanceof com.lanlian.rpc.version.AddNewVersionResponse) {
				return mergeFrom((com.lanlian.rpc.version.AddNewVersionResponse) other);
			} else {
				super.mergeFrom(other);
				return this;
			}
		}

		public Builder mergeFrom(com.lanlian.rpc.version.AddNewVersionResponse other) {
			if (other == com.lanlian.rpc.version.AddNewVersionResponse.getDefaultInstance())
				return this;
			if (other.getErrorCode() != 0) {
				setErrorCode(other.getErrorCode());
			}
			onChanged();
			return this;
		}

		public final boolean isInitialized() {
			return true;
		}

		public Builder mergeFrom(com.google.protobuf.CodedInputStream input,
				com.google.protobuf.ExtensionRegistryLite extensionRegistry) throws java.io.IOException {
			com.lanlian.rpc.version.AddNewVersionResponse parsedMessage = null;
			try {
				parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
			} catch (com.google.protobuf.InvalidProtocolBufferException e) {
				parsedMessage = (com.lanlian.rpc.version.AddNewVersionResponse) e.getUnfinishedMessage();
				throw e;
			} finally {
				if (parsedMessage != null) {
					mergeFrom(parsedMessage);
				}
			}
			return this;
		}

		private int errorCode_;

		/**
		 * <code>optional int32 ErrorCode = 1;</code>
		 */
		public int getErrorCode() {
			return errorCode_;
		}

		/**
		 * <code>optional int32 ErrorCode = 1;</code>
		 */
		public Builder setErrorCode(int value) {

			errorCode_ = value;
			onChanged();
			return this;
		}

		/**
		 * <code>optional int32 ErrorCode = 1;</code>
		 */
		public Builder clearErrorCode() {

			errorCode_ = 0;
			onChanged();
			return this;
		}

		public final Builder setUnknownFields(final com.google.protobuf.UnknownFieldSet unknownFields) {
			return this;
		}

		public final Builder mergeUnknownFields(final com.google.protobuf.UnknownFieldSet unknownFields) {
			return this;
		}

		// @@protoc_insertion_point(builder_scope:setting.AddNewVersionResponse)
	}

	// @@protoc_insertion_point(class_scope:setting.AddNewVersionResponse)
	private static final com.lanlian.rpc.version.AddNewVersionResponse DEFAULT_INSTANCE;
	static {
		DEFAULT_INSTANCE = new com.lanlian.rpc.version.AddNewVersionResponse();
	}

	public static com.lanlian.rpc.version.AddNewVersionResponse getDefaultInstance() {
		return DEFAULT_INSTANCE;
	}

	private static final com.google.protobuf.Parser<AddNewVersionResponse> PARSER = new com.google.protobuf.AbstractParser<AddNewVersionResponse>() {
		public AddNewVersionResponse parsePartialFrom(com.google.protobuf.CodedInputStream input,
				com.google.protobuf.ExtensionRegistryLite extensionRegistry)
				throws com.google.protobuf.InvalidProtocolBufferException {
			try {
				return new AddNewVersionResponse(input, extensionRegistry);
			} catch (RuntimeException e) {
				if (e.getCause() instanceof com.google.protobuf.InvalidProtocolBufferException) {
					throw (com.google.protobuf.InvalidProtocolBufferException) e.getCause();
				}
				throw e;
			}
		}
	};

	public static com.google.protobuf.Parser<AddNewVersionResponse> parser() {
		return PARSER;
	}

	@java.lang.Override
	public com.google.protobuf.Parser<AddNewVersionResponse> getParserForType() {
		return PARSER;
	}

	public com.lanlian.rpc.version.AddNewVersionResponse getDefaultInstanceForType() {
		return DEFAULT_INSTANCE;
	}

}
