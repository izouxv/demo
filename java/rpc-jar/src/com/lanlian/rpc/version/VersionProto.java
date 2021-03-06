// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: radacat_version.proto

package com.lanlian.rpc.version;

public final class VersionProto {
	private VersionProto() {
	}

	public static void registerAllExtensions(com.google.protobuf.ExtensionRegistry registry) {
	}

	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_Version_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_Version_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_AddNewVersionRequest_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_AddNewVersionRequest_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_AddNewVersionResponse_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_AddNewVersionResponse_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_GetAllVersionsRequest_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_GetAllVersionsRequest_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_GetAllVersionsResponse_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_GetAllVersionsResponse_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_GetLatestVersionRequest_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_GetLatestVersionRequest_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_GetLatestVersionResponse_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_GetLatestVersionResponse_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_DeleteVersionRequest_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_DeleteVersionRequest_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_DeleteVersionResponse_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_DeleteVersionResponse_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_UpdateVersionRequest_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_UpdateVersionRequest_fieldAccessorTable;
	static com.google.protobuf.Descriptors.Descriptor internal_static_setting_UpdateVersionResponse_descriptor;
	static com.google.protobuf.GeneratedMessage.FieldAccessorTable internal_static_setting_UpdateVersionResponse_fieldAccessorTable;

	public static com.google.protobuf.Descriptors.FileDescriptor getDescriptor() {
		return descriptor;
	}

	private static com.google.protobuf.Descriptors.FileDescriptor descriptor;
	static {
		java.lang.String[] descriptorData = { "\n\025radacat_version.proto\022\007setting\"\224\002\n\007Ver"
				+ "sion\022\016\n\006device\030\001 \001(\t\022\023\n\013versionName\030\002 \001("
				+ "\t\022\023\n\013versionCode\030\003 \001(\t\022\013\n\003md5\030\004 \001(\t\022\020\n\010f"
				+ "ileName\030\005 \001(\t\022\016\n\006length\030\006 \001(\003\022\014\n\004path\030\007 "
				+ "\001(\t\022\025\n\rdescriptionCn\030\010 \001(\t\022\025\n\rdescriptio"
				+ "nEn\030\t \001(\t\022\023\n\013uploaderUid\030\n \001(\003\022\016\n\006status"
				+ "\030\013 \001(\005\022\n\n\002id\030\014 \001(\005\022\022\n\ncreateTime\030\r \001(\003\022\022"
				+ "\n\nupdateTime\030\016 \001(\003\022\013\n\003tid\030\017 \001(\003\"9\n\024AddNe"
				+ "wVersionRequest\022!\n\007version\030\001 \001(\0132\020.setti"
				+ "ng.Version\"*\n\025AddNewVersionResponse\022\021\n\tE",
				"rrorCode\030\001 \001(\005\"A\n\025GetAllVersionsRequest\022"
						+ "\r\n\005count\030\001 \001(\005\022\014\n\004page\030\002 \001(\005\022\013\n\003tid\030\003 \001("
						+ "\003\"c\n\026GetAllVersionsResponse\022\021\n\tErrorCode"
						+ "\030\001 \001(\005\022\"\n\010versions\030\002 \003(\0132\020.setting.Versi"
						+ "on\022\022\n\ntotalCount\030\003 \001(\005\"9\n\027GetLatestVersi"
						+ "onRequest\022\016\n\006device\030\001 \001(\t\022\016\n\006source\030\002 \001("
						+ "\t\"P\n\030GetLatestVersionResponse\022\021\n\terrorCo"
						+ "de\030\001 \001(\005\022!\n\007version\030\002 \001(\0132\020.setting.Vers"
						+ "ion\"/\n\024DeleteVersionRequest\022\n\n\002id\030\001 \001(\005\022"
						+ "\013\n\003tid\030\002 \001(\003\"*\n\025DeleteVersionResponse\022\021\n",
				"\terrorCode\030\001 \001(\005\"9\n\024UpdateVersionRequest"
						+ "\022!\n\007version\030\001 \001(\0132\020.setting.Version\"*\n\025U"
						+ "pdateVersionResponse\022\021\n\tErrorCode\030\001 \001(\0052"
						+ "\266\003\n\016RadacatVersion\022P\n\rAddNewVersion\022\035.se"
						+ "tting.AddNewVersionRequest\032\036.setting.Add"
						+ "NewVersionResponse\"\000\022S\n\016GetAllVersions\022\036"
						+ ".setting.GetAllVersionsRequest\032\037.setting"
						+ ".GetAllVersionsResponse\"\000\022Y\n\020GetLatestVe"
						+ "rsion\022 .setting.GetLatestVersionRequest\032"
						+ "!.setting.GetLatestVersionResponse\"\000\022P\n\r",
				"UpdateVersion\022\035.setting.UpdateVersionReq"
						+ "uest\032\036.setting.UpdateVersionResponse\"\000\022P"
						+ "\n\rDeleteVersion\022\035.setting.DeleteVersionR"
						+ "equest\032\036.setting.DeleteVersionResponse\"\000"
						+ "B0\n\027com.lanlian.rpc.versionB\014VersionProt" + "oP\001\242\002\004Grpcb\006proto3" };
		com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner = new com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner() {
			public com.google.protobuf.ExtensionRegistry assignDescriptors(
					com.google.protobuf.Descriptors.FileDescriptor root) {
				descriptor = root;
				return null;
			}
		};
		com.google.protobuf.Descriptors.FileDescriptor.internalBuildGeneratedFileFrom(descriptorData,
				new com.google.protobuf.Descriptors.FileDescriptor[] {}, assigner);
		internal_static_setting_Version_descriptor = getDescriptor().getMessageTypes().get(0);
		internal_static_setting_Version_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_Version_descriptor,
				new java.lang.String[] { "Device", "VersionName", "VersionCode", "Md5", "FileName", "Length", "Path",
						"DescriptionCn", "DescriptionEn", "UploaderUid", "Status", "Id", "CreateTime", "UpdateTime",
						"Tid", });
		internal_static_setting_AddNewVersionRequest_descriptor = getDescriptor().getMessageTypes().get(1);
		internal_static_setting_AddNewVersionRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_AddNewVersionRequest_descriptor, new java.lang.String[] { "Version", });
		internal_static_setting_AddNewVersionResponse_descriptor = getDescriptor().getMessageTypes().get(2);
		internal_static_setting_AddNewVersionResponse_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_AddNewVersionResponse_descriptor, new java.lang.String[] { "ErrorCode", });
		internal_static_setting_GetAllVersionsRequest_descriptor = getDescriptor().getMessageTypes().get(3);
		internal_static_setting_GetAllVersionsRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_GetAllVersionsRequest_descriptor,
				new java.lang.String[] { "Count", "Page", "Tid", });
		internal_static_setting_GetAllVersionsResponse_descriptor = getDescriptor().getMessageTypes().get(4);
		internal_static_setting_GetAllVersionsResponse_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_GetAllVersionsResponse_descriptor,
				new java.lang.String[] { "ErrorCode", "Versions", "TotalCount", });
		internal_static_setting_GetLatestVersionRequest_descriptor = getDescriptor().getMessageTypes().get(5);
		internal_static_setting_GetLatestVersionRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_GetLatestVersionRequest_descriptor,
				new java.lang.String[] { "Device", "Source", });
		internal_static_setting_GetLatestVersionResponse_descriptor = getDescriptor().getMessageTypes().get(6);
		internal_static_setting_GetLatestVersionResponse_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_GetLatestVersionResponse_descriptor,
				new java.lang.String[] { "ErrorCode", "Version", });
		internal_static_setting_DeleteVersionRequest_descriptor = getDescriptor().getMessageTypes().get(7);
		internal_static_setting_DeleteVersionRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_DeleteVersionRequest_descriptor, new java.lang.String[] { "Id", "Tid", });
		internal_static_setting_DeleteVersionResponse_descriptor = getDescriptor().getMessageTypes().get(8);
		internal_static_setting_DeleteVersionResponse_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_DeleteVersionResponse_descriptor, new java.lang.String[] { "ErrorCode", });
		internal_static_setting_UpdateVersionRequest_descriptor = getDescriptor().getMessageTypes().get(9);
		internal_static_setting_UpdateVersionRequest_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_UpdateVersionRequest_descriptor, new java.lang.String[] { "Version", });
		internal_static_setting_UpdateVersionResponse_descriptor = getDescriptor().getMessageTypes().get(10);
		internal_static_setting_UpdateVersionResponse_fieldAccessorTable = new com.google.protobuf.GeneratedMessage.FieldAccessorTable(
				internal_static_setting_UpdateVersionResponse_descriptor, new java.lang.String[] { "ErrorCode", });
	}

	// @@protoc_insertion_point(outer_class_scope)
}
