// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: sso.proto

package com.lanlian.rpc.sso;

/**
 * Protobuf type {@code pb.MapSsoReply}
 */
public final class MapSsoReply extends com.google.protobuf.GeneratedMessage implements
		// @@protoc_insertion_point(message_implements:pb.MapSsoReply)
		MapSsoReplyOrBuilder {
	// Use MapSsoReply.newBuilder() to construct.
	private MapSsoReply(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
		super(builder);
	}

	private MapSsoReply() {
		errorCode_ = 0;
	}

	@java.lang.Override
	public final com.google.protobuf.UnknownFieldSet getUnknownFields() {
		return com.google.protobuf.UnknownFieldSet.getDefaultInstance();
	}

	private MapSsoReply(com.google.protobuf.CodedInputStream input,
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
				case 10: {
					if (!((mutable_bitField0_ & 0x00000001) == 0x00000001)) {
						ssos_ = com.google.protobuf.MapField.newMapField(SsosDefaultEntryHolder.defaultEntry);
						mutable_bitField0_ |= 0x00000001;
					}
					com.google.protobuf.MapEntry<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> ssos = input
							.readMessage(SsosDefaultEntryHolder.defaultEntry.getParserForType(), extensionRegistry);
					ssos_.getMutableMap().put(ssos.getKey(), ssos.getValue());
					break;
				}
				case 16: {

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
		return com.lanlian.rpc.sso.SsoProto.internal_static_pb_MapSsoReply_descriptor;
	}

	@SuppressWarnings({ "rawtypes" })
	protected com.google.protobuf.MapField internalGetMapField(int number) {
		switch (number) {
		case 1:
			return internalGetSsos();
		default:
			throw new RuntimeException("Invalid map field number: " + number);
		}
	}

	protected com.google.protobuf.GeneratedMessage.FieldAccessorTable internalGetFieldAccessorTable() {
		return com.lanlian.rpc.sso.SsoProto.internal_static_pb_MapSsoReply_fieldAccessorTable
				.ensureFieldAccessorsInitialized(com.lanlian.rpc.sso.MapSsoReply.class,
						com.lanlian.rpc.sso.MapSsoReply.Builder.class);
	}

	private int bitField0_;
	public static final int SSOS_FIELD_NUMBER = 1;

	private static final class SsosDefaultEntryHolder {
		static final com.google.protobuf.MapEntry<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> defaultEntry = com.google.protobuf.MapEntry.<java.lang.Integer, com.lanlian.rpc.sso.SsoReply>newDefaultInstance(
				com.lanlian.rpc.sso.SsoProto.internal_static_pb_MapSsoReply_SsosEntry_descriptor,
				com.google.protobuf.WireFormat.FieldType.INT32, 0, com.google.protobuf.WireFormat.FieldType.MESSAGE,
				com.lanlian.rpc.sso.SsoReply.getDefaultInstance());
	}

	private com.google.protobuf.MapField<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> ssos_;

	private com.google.protobuf.MapField<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> internalGetSsos() {
		if (ssos_ == null) {
			return com.google.protobuf.MapField.emptyMapField(SsosDefaultEntryHolder.defaultEntry);
		}
		return ssos_;
	}

	/**
	 * <code>map&lt;int32, .pb.SsoReply&gt; ssos = 1;</code>
	 */

	public java.util.Map<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> getSsos() {
		return internalGetSsos().getMap();
	}

	public static final int ERRORCODE_FIELD_NUMBER = 2;
	private int errorCode_;

	/**
	 * <code>optional int32 errorCode = 2;</code>
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
		for (java.util.Map.Entry<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> entry : internalGetSsos().getMap()
				.entrySet()) {
			com.google.protobuf.MapEntry<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> ssos = SsosDefaultEntryHolder.defaultEntry
					.newBuilderForType().setKey(entry.getKey()).setValue(entry.getValue()).build();
			output.writeMessage(1, ssos);
		}
		if (errorCode_ != 0) {
			output.writeInt32(2, errorCode_);
		}
	}

	public int getSerializedSize() {
		int size = memoizedSize;
		if (size != -1)
			return size;

		size = 0;
		for (java.util.Map.Entry<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> entry : internalGetSsos().getMap()
				.entrySet()) {
			com.google.protobuf.MapEntry<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> ssos = SsosDefaultEntryHolder.defaultEntry
					.newBuilderForType().setKey(entry.getKey()).setValue(entry.getValue()).build();
			size += com.google.protobuf.CodedOutputStream.computeMessageSize(1, ssos);
		}
		if (errorCode_ != 0) {
			size += com.google.protobuf.CodedOutputStream.computeInt32Size(2, errorCode_);
		}
		memoizedSize = size;
		return size;
	}

	private static final long serialVersionUID = 0L;

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(com.google.protobuf.ByteString data)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(com.google.protobuf.ByteString data,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data, extensionRegistry);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(byte[] data)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(byte[] data,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry)
			throws com.google.protobuf.InvalidProtocolBufferException {
		return PARSER.parseFrom(data, extensionRegistry);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(java.io.InputStream input) throws java.io.IOException {
		return PARSER.parseFrom(input);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(java.io.InputStream input,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry) throws java.io.IOException {
		return PARSER.parseFrom(input, extensionRegistry);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseDelimitedFrom(java.io.InputStream input)
			throws java.io.IOException {
		return PARSER.parseDelimitedFrom(input);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseDelimitedFrom(java.io.InputStream input,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry) throws java.io.IOException {
		return PARSER.parseDelimitedFrom(input, extensionRegistry);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(com.google.protobuf.CodedInputStream input)
			throws java.io.IOException {
		return PARSER.parseFrom(input);
	}

	public static com.lanlian.rpc.sso.MapSsoReply parseFrom(com.google.protobuf.CodedInputStream input,
			com.google.protobuf.ExtensionRegistryLite extensionRegistry) throws java.io.IOException {
		return PARSER.parseFrom(input, extensionRegistry);
	}

	public Builder newBuilderForType() {
		return newBuilder();
	}

	public static Builder newBuilder() {
		return DEFAULT_INSTANCE.toBuilder();
	}

	public static Builder newBuilder(com.lanlian.rpc.sso.MapSsoReply prototype) {
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
	 * Protobuf type {@code pb.MapSsoReply}
	 */
	public static final class Builder extends com.google.protobuf.GeneratedMessage.Builder<Builder> implements
			// @@protoc_insertion_point(builder_implements:pb.MapSsoReply)
			com.lanlian.rpc.sso.MapSsoReplyOrBuilder {
		public static final com.google.protobuf.Descriptors.Descriptor getDescriptor() {
			return com.lanlian.rpc.sso.SsoProto.internal_static_pb_MapSsoReply_descriptor;
		}

		@SuppressWarnings({ "rawtypes" })
		protected com.google.protobuf.MapField internalGetMapField(int number) {
			switch (number) {
			case 1:
				return internalGetSsos();
			default:
				throw new RuntimeException("Invalid map field number: " + number);
			}
		}

		@SuppressWarnings({ "rawtypes" })
		protected com.google.protobuf.MapField internalGetMutableMapField(int number) {
			switch (number) {
			case 1:
				return internalGetMutableSsos();
			default:
				throw new RuntimeException("Invalid map field number: " + number);
			}
		}

		protected com.google.protobuf.GeneratedMessage.FieldAccessorTable internalGetFieldAccessorTable() {
			return com.lanlian.rpc.sso.SsoProto.internal_static_pb_MapSsoReply_fieldAccessorTable
					.ensureFieldAccessorsInitialized(com.lanlian.rpc.sso.MapSsoReply.class,
							com.lanlian.rpc.sso.MapSsoReply.Builder.class);
		}

		// Construct using com.lanlian.rpc.sso.MapSsoReply.newBuilder()
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
			internalGetMutableSsos().clear();
			errorCode_ = 0;

			return this;
		}

		public com.google.protobuf.Descriptors.Descriptor getDescriptorForType() {
			return com.lanlian.rpc.sso.SsoProto.internal_static_pb_MapSsoReply_descriptor;
		}

		public com.lanlian.rpc.sso.MapSsoReply getDefaultInstanceForType() {
			return com.lanlian.rpc.sso.MapSsoReply.getDefaultInstance();
		}

		public com.lanlian.rpc.sso.MapSsoReply build() {
			com.lanlian.rpc.sso.MapSsoReply result = buildPartial();
			if (!result.isInitialized()) {
				throw newUninitializedMessageException(result);
			}
			return result;
		}

		public com.lanlian.rpc.sso.MapSsoReply buildPartial() {
			com.lanlian.rpc.sso.MapSsoReply result = new com.lanlian.rpc.sso.MapSsoReply(this);
			int from_bitField0_ = bitField0_;
			int to_bitField0_ = 0;
			result.ssos_ = internalGetSsos();
			result.ssos_.makeImmutable();
			result.errorCode_ = errorCode_;
			result.bitField0_ = to_bitField0_;
			onBuilt();
			return result;
		}

		public Builder mergeFrom(com.google.protobuf.Message other) {
			if (other instanceof com.lanlian.rpc.sso.MapSsoReply) {
				return mergeFrom((com.lanlian.rpc.sso.MapSsoReply) other);
			} else {
				super.mergeFrom(other);
				return this;
			}
		}

		public Builder mergeFrom(com.lanlian.rpc.sso.MapSsoReply other) {
			if (other == com.lanlian.rpc.sso.MapSsoReply.getDefaultInstance())
				return this;
			internalGetMutableSsos().mergeFrom(other.internalGetSsos());
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
			com.lanlian.rpc.sso.MapSsoReply parsedMessage = null;
			try {
				parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
			} catch (com.google.protobuf.InvalidProtocolBufferException e) {
				parsedMessage = (com.lanlian.rpc.sso.MapSsoReply) e.getUnfinishedMessage();
				throw e;
			} finally {
				if (parsedMessage != null) {
					mergeFrom(parsedMessage);
				}
			}
			return this;
		}

		private int bitField0_;

		private com.google.protobuf.MapField<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> ssos_;

		private com.google.protobuf.MapField<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> internalGetSsos() {
			if (ssos_ == null) {
				return com.google.protobuf.MapField.emptyMapField(SsosDefaultEntryHolder.defaultEntry);
			}
			return ssos_;
		}

		private com.google.protobuf.MapField<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> internalGetMutableSsos() {
			onChanged();
			;
			if (ssos_ == null) {
				ssos_ = com.google.protobuf.MapField.newMapField(SsosDefaultEntryHolder.defaultEntry);
			}
			if (!ssos_.isMutable()) {
				ssos_ = ssos_.copy();
			}
			return ssos_;
		}

		/**
		 * <code>map&lt;int32, .pb.SsoReply&gt; ssos = 1;</code>
		 */
		public java.util.Map<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> getSsos() {
			return internalGetSsos().getMap();
		}

		/**
		 * <code>map&lt;int32, .pb.SsoReply&gt; ssos = 1;</code>
		 */
		public java.util.Map<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> getMutableSsos() {
			return internalGetMutableSsos().getMutableMap();
		}

		/**
		 * <code>map&lt;int32, .pb.SsoReply&gt; ssos = 1;</code>
		 */
		public Builder putAllSsos(java.util.Map<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> values) {
			getMutableSsos().putAll(values);
			return this;
		}

		private int errorCode_;

		/**
		 * <code>optional int32 errorCode = 2;</code>
		 */
		public int getErrorCode() {
			return errorCode_;
		}

		/**
		 * <code>optional int32 errorCode = 2;</code>
		 */
		public Builder setErrorCode(int value) {

			errorCode_ = value;
			onChanged();
			return this;
		}

		/**
		 * <code>optional int32 errorCode = 2;</code>
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

		// @@protoc_insertion_point(builder_scope:pb.MapSsoReply)
	}

	// @@protoc_insertion_point(class_scope:pb.MapSsoReply)
	private static final com.lanlian.rpc.sso.MapSsoReply DEFAULT_INSTANCE;
	static {
		DEFAULT_INSTANCE = new com.lanlian.rpc.sso.MapSsoReply();
	}

	public static com.lanlian.rpc.sso.MapSsoReply getDefaultInstance() {
		return DEFAULT_INSTANCE;
	}

	private static final com.google.protobuf.Parser<MapSsoReply> PARSER = new com.google.protobuf.AbstractParser<MapSsoReply>() {
		public MapSsoReply parsePartialFrom(com.google.protobuf.CodedInputStream input,
				com.google.protobuf.ExtensionRegistryLite extensionRegistry)
				throws com.google.protobuf.InvalidProtocolBufferException {
			try {
				return new MapSsoReply(input, extensionRegistry);
			} catch (RuntimeException e) {
				if (e.getCause() instanceof com.google.protobuf.InvalidProtocolBufferException) {
					throw (com.google.protobuf.InvalidProtocolBufferException) e.getCause();
				}
				throw e;
			}
		}
	};

	public static com.google.protobuf.Parser<MapSsoReply> parser() {
		return PARSER;
	}

	@java.lang.Override
	public com.google.protobuf.Parser<MapSsoReply> getParserForType() {
		return PARSER;
	}

	public com.lanlian.rpc.sso.MapSsoReply getDefaultInstanceForType() {
		return DEFAULT_INSTANCE;
	}

}
