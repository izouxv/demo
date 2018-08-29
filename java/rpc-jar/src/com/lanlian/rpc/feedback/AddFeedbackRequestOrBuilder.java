// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: feedback.proto

package com.lanlian.rpc.feedback;

public interface AddFeedbackRequestOrBuilder extends
		// @@protoc_insertion_point(interface_extends:feedback.AddFeedbackRequest)
		com.google.protobuf.MessageOrBuilder {

	/**
	 * <code>optional string source = 1;</code>
	 */
	java.lang.String getSource();

	/**
	 * <code>optional string source = 1;</code>
	 */
	com.google.protobuf.ByteString getSourceBytes();

	/**
	 * <code>optional string description = 2;</code>
	 */
	java.lang.String getDescription();

	/**
	 * <code>optional string description = 2;</code>
	 */
	com.google.protobuf.ByteString getDescriptionBytes();

	/**
	 * <code>optional string deviceInfo = 3;</code>
	 */
	java.lang.String getDeviceInfo();

	/**
	 * <code>optional string deviceInfo = 3;</code>
	 */
	com.google.protobuf.ByteString getDeviceInfoBytes();

	/**
	 * <code>optional string appInfo = 4;</code>
	 */
	java.lang.String getAppInfo();

	/**
	 * <code>optional string appInfo = 4;</code>
	 */
	com.google.protobuf.ByteString getAppInfoBytes();

	/**
	 * <code>optional string userInfo = 5;</code>
	 */
	java.lang.String getUserInfo();

	/**
	 * <code>optional string userInfo = 5;</code>
	 */
	com.google.protobuf.ByteString getUserInfoBytes();

	/**
	 * <code>optional string mobileInfo = 6;</code>
	 */
	java.lang.String getMobileInfo();

	/**
	 * <code>optional string mobileInfo = 6;</code>
	 */
	com.google.protobuf.ByteString getMobileInfoBytes();

	/**
	 * <code>optional string extendInfo = 7;</code>
	 */
	java.lang.String getExtendInfo();

	/**
	 * <code>optional string extendInfo = 7;</code>
	 */
	com.google.protobuf.ByteString getExtendInfoBytes();

	/**
	 * <code>repeated string files = 8;</code>
	 */
	com.google.protobuf.ProtocolStringList getFilesList();

	/**
	 * <code>repeated string files = 8;</code>
	 */
	int getFilesCount();

	/**
	 * <code>repeated string files = 8;</code>
	 */
	java.lang.String getFiles(int index);

	/**
	 * <code>repeated string files = 8;</code>
	 */
	com.google.protobuf.ByteString getFilesBytes(int index);

	/**
	 * <code>optional string contact = 9;</code>
	 */
	java.lang.String getContact();

	/**
	 * <code>optional string contact = 9;</code>
	 */
	com.google.protobuf.ByteString getContactBytes();

	/**
	 * <code>optional int64 createTime = 10;</code>
	 */
	long getCreateTime();

	/**
	 * <code>optional int64 updateTime = 11;</code>
	 */
	long getUpdateTime();
}