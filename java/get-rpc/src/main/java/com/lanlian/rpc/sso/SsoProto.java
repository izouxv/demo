// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: sso.proto

package com.lanlian.rpc.sso;

public final class SsoProto {
	private SsoProto() {
	}

	public static void registerAllExtensions(com.google.protobuf.ExtensionRegistry registry) {
	}

	static com.google.protobuf.Descriptors.Descriptor internal_static_pb_SsoRequest_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_pb_SsoRequest_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_pb_AgentInfo_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_pb_AgentInfo_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_pb_SsoReply_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_pb_SsoReply_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_pb_MultiSsoRequest_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_pb_MultiSsoRequest_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_pb_MultiSsoRequest_SsosEntry_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_pb_MultiSsoRequest_SsosEntry_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_pb_MapSsoReply_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_pb_MapSsoReply_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_pb_MapSsoReply_SsosEntry_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_pb_MapSsoReply_SsosEntry_fieldAccessorTable;

	public static com.google.protobuf.Descriptors.FileDescriptor getDescriptor() {
		return descriptor;
	}

	private static com.google.protobuf.Descriptors.FileDescriptor descriptor;
	static {
		java.lang.String[] descriptorData = { "\n\tsso.proto\022\002pb\"\363\001\n\nSsoRequest\022\020\n\010userna"
				+ "me\030\001 \001(\t\022\020\n\010password\030\002 \001(\t\022\023\n\013sessionNam"
				+ "e\030\003 \001(\t\022\014\n\004salt\030\004 \001(\t\022\013\n\003uid\030\005 \001(\005\022\r\n\005st"
				+ "ate\030\006 \001(\005\022\017\n\007exptime\030\007 \001(\005\022\020\n\010nickname\030\010"
				+ " \001(\t\022\r\n\005token\030\t \001(\t\022\014\n\004code\030\n \001(\t\022\020\n\010cod"
				+ "eType\030\013 \001(\005\022\016\n\006source\030\014 \001(\t\022 \n\tagentInfo"
				+ "\030\r \001(\0132\r.pb.AgentInfo\"(\n\tAgentInfo\022\n\n\002ip"
				+ "\030\001 \001(\t\022\017\n\007devInfo\030\002 \001(\t\"\225\001\n\010SsoReply\022\013\n\003"
				+ "uid\030\001 \001(\005\022\020\n\010username\030\002 \001(\t\022\022\n\nloginStat"
				+ "e\030\003 \001(\005\022\r\n\005state\030\004 \001(\005\022\021\n\terrorCode\030\005 \001(",
				"\005\022\023\n\013sessionName\030\006 \001(\t\022\020\n\010nickname\030\007 \001(\t"
						+ "\022\r\n\005token\030\010 \001(\t\"\213\001\n\017MultiSsoRequest\022+\n\004s"
						+ "sos\030\001 \003(\0132\035.pb.MultiSsoRequest.SsosEntry"
						+ "\022\016\n\006source\030\002 \001(\t\032;\n\tSsosEntry\022\013\n\003key\030\001 \001"
						+ "(\005\022\035\n\005value\030\002 \001(\0132\016.pb.SsoRequest:\0028\001\"\204\001"
						+ "\n\013MapSsoReply\022\'\n\004ssos\030\001 \003(\0132\031.pb.MapSsoR"
						+ "eply.SsosEntry\022\021\n\terrorCode\030\002 \001(\005\0329\n\tSso"
						+ "sEntry\022\013\n\003key\030\001 \001(\005\022\033\n\005value\030\002 \001(\0132\014.pb."
						+ "SsoReply:\0028\0012\224\006\n\003Sso\022-\n\013GetUserInfo\022\016.pb"
						+ ".SsoRequest\032\014.pb.SsoReply\"\000\022\'\n\005Login\022\016.p",
				"b.SsoRequest\032\014.pb.SsoReply\"\000\022/\n\rGetUserB"
						+ "yName\022\016.pb.SsoRequest\032\014.pb.SsoReply\"\000\022%\n"
						+ "\003Add\022\016.pb.SsoRequest\032\014.pb.SsoReply\"\000\022/\n\r"
						+ "CheckPassword\022\016.pb.SsoRequest\032\014.pb.SsoRe"
						+ "ply\"\000\0220\n\016UpdatePassword\022\016.pb.SsoRequest\032"
						+ "\014.pb.SsoReply\"\000\022(\n\006Logout\022\016.pb.SsoReques"
						+ "t\032\014.pb.SsoReply\"\000\0226\n\024UpdatePasswordByNam"
						+ "e\022\016.pb.SsoRequest\032\014.pb.SsoReply\"\000\0224\n\022Fin"
						+ "dPasswordByMail\022\016.pb.SsoRequest\032\014.pb.Sso"
						+ "Reply\"\000\022/\n\rResetPassword\022\016.pb.SsoRequest",
				"\032\014.pb.SsoReply\"\000\022-\n\013UpdateState\022\016.pb.Sso"
						+ "Request\032\014.pb.SsoReply\"\000\022:\n\020GetBatchSsoIn"
						+ "fos\022\023.pb.MultiSsoRequest\032\017.pb.MapSsoRepl"
						+ "y\"\000\022+\n\tCheckCode\022\016.pb.SsoRequest\032\014.pb.Ss"
						+ "oReply\"\000\0226\n\024ResetPasswordByPhone\022\016.pb.Ss"
						+ "oRequest\032\014.pb.SsoReply\"\000\0220\n\016SendMobileCo"
						+ "de\022\016.pb.SsoRequest\032\014.pb.SsoReply\"\000\022/\n\rLo"
						+ "ginWithCode\022\016.pb.SsoRequest\032\014.pb.SsoRepl"
						+ "y\"\000B(\n\023com.lanlian.rpc.ssoB\010SsoProtoP\001\242\002" + "\004Grpcb\006proto3" };
		com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner = new com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner() {
			public com.google.protobuf.ExtensionRegistry assignDescriptors(
					com.google.protobuf.Descriptors.FileDescriptor root) {
				descriptor = root;
				return null;
			}
		};
		com.google.protobuf.Descriptors.FileDescriptor.internalBuildGeneratedFileFrom(descriptorData,
				new com.google.protobuf.Descriptors.FileDescriptor[] {}, assigner);
		internal_static_pb_SsoRequest_descriptor = getDescriptor().getMessageTypes().get(0);
		internal_static_pb_SsoRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_pb_SsoRequest_descriptor,
				new java.lang.String[] { "Username", "Password", "SessionName", "Salt", "Uid", "State", "Exptime",
						"Nickname", "Token", "Code", "CodeType", "Source", "AgentInfo", });
		internal_static_pb_AgentInfo_descriptor = getDescriptor().getMessageTypes().get(1);
		internal_static_pb_AgentInfo_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_pb_AgentInfo_descriptor, new java.lang.String[] { "Ip", "DevInfo", });
		internal_static_pb_SsoReply_descriptor = getDescriptor().getMessageTypes().get(2);
		internal_static_pb_SsoReply_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_pb_SsoReply_descriptor, new java.lang.String[] { "Uid", "Username", "LoginState",
						"State", "ErrorCode", "SessionName", "Nickname", "Token", });
		internal_static_pb_MultiSsoRequest_descriptor = getDescriptor().getMessageTypes().get(3);
		internal_static_pb_MultiSsoRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_pb_MultiSsoRequest_descriptor, new java.lang.String[] { "Ssos", "Source", });
		internal_static_pb_MultiSsoRequest_SsosEntry_descriptor = internal_static_pb_MultiSsoRequest_descriptor
				.getNestedTypes().get(0);
		internal_static_pb_MultiSsoRequest_SsosEntry_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_pb_MultiSsoRequest_SsosEntry_descriptor, new java.lang.String[] { "Key", "Value", });
		internal_static_pb_MapSsoReply_descriptor = getDescriptor().getMessageTypes().get(4);
		internal_static_pb_MapSsoReply_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_pb_MapSsoReply_descriptor, new java.lang.String[] { "Ssos", "ErrorCode", });
		internal_static_pb_MapSsoReply_SsosEntry_descriptor = internal_static_pb_MapSsoReply_descriptor.getNestedTypes()
				.get(0);
		internal_static_pb_MapSsoReply_SsosEntry_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_pb_MapSsoReply_SsosEntry_descriptor, new java.lang.String[] { "Key", "Value", });
	}

	// @@protoc_insertion_point(outer_class_scope)
}