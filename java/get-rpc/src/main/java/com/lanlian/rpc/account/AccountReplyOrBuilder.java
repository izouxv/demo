// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: account.proto

package com.lanlian.rpc.account;

public interface AccountReplyOrBuilder extends
		// @@protoc_insertion_point(interface_extends:pb.AccountReply)
		com.google.protobuf.MessageOrBuilder {

	/**
	 * <code>optional int32 uid = 1;</code>
	 */
	int getUid();

	/**
	 * <code>optional string username = 2;</code>
	 */
	java.lang.String getUsername();

	/**
	 * <code>optional string username = 2;</code>
	 */
	com.google.protobuf.ByteString getUsernameBytes();

	/**
	 * <code>optional string email = 3;</code>
	 */
	java.lang.String getEmail();

	/**
	 * <code>optional string email = 3;</code>
	 */
	com.google.protobuf.ByteString getEmailBytes();

	/**
	 * <code>optional string phone = 4;</code>
	 */
	java.lang.String getPhone();

	/**
	 * <code>optional string phone = 4;</code>
	 */
	com.google.protobuf.ByteString getPhoneBytes();

	/**
	 * <code>optional int32 state = 5;</code>
	 */
	int getState();

	/**
	 * <code>optional int64 lastLoginTime = 6;</code>
	 */
	long getLastLoginTime();

	/**
	 * <code>optional int64 createTime = 7;</code>
	 */
	long getCreateTime();

	/**
	 * <code>optional string nickname = 8;</code>
	 */
	java.lang.String getNickname();

	/**
	 * <code>optional string nickname = 8;</code>
	 */
	com.google.protobuf.ByteString getNicknameBytes();

	/**
	 * <code>optional string realname = 9;</code>
	 */
	java.lang.String getRealname();

	/**
	 * <code>optional string realname = 9;</code>
	 */
	com.google.protobuf.ByteString getRealnameBytes();

	/**
	 * <code>optional int32 isCertification = 10;</code>
	 */
	int getIsCertification();

	/**
	 * <code>optional string identityCard = 11;</code>
	 */
	java.lang.String getIdentityCard();

	/**
	 * <code>optional string identityCard = 11;</code>
	 */
	com.google.protobuf.ByteString getIdentityCardBytes();

	/**
	 * <code>optional int32 errorCode = 12;</code>
	 */
	int getErrorCode();

	/**
	 * <code>optional int32 gender = 13;</code>
	 */
	int getGender();

	/**
	 * <code>optional int64 birthday = 14;</code>
	 */
	long getBirthday();

	/**
	 * <code>optional int32 avatar = 15;</code>
	 */
	int getAvatar();

	/**
	 * <code>optional string province = 16;</code>
	 */
	java.lang.String getProvince();

	/**
	 * <code>optional string province = 16;</code>
	 */
	com.google.protobuf.ByteString getProvinceBytes();

	/**
	 * <code>optional string city = 17;</code>
	 */
	java.lang.String getCity();

	/**
	 * <code>optional string city = 17;</code>
	 */
	com.google.protobuf.ByteString getCityBytes();

	/**
	 * <code>optional string signature = 18;</code>
	 */
	java.lang.String getSignature();

	/**
	 * <code>optional string signature = 18;</code>
	 */
	com.google.protobuf.ByteString getSignatureBytes();

	/**
	 * <code>optional string userAddress = 19;</code>
	 */
	java.lang.String getUserAddress();

	/**
	 * <code>optional string userAddress = 19;</code>
	 */
	com.google.protobuf.ByteString getUserAddressBytes();

	/**
	 * <code>optional int32 userJobId = 20;</code>
	 */
	int getUserJobId();

	/**
	 * <code>optional int32 creditValues = 21;</code>
	 */
	int getCreditValues();

	/**
	 * <code>optional int32 userPoint = 22;</code>
	 */
	int getUserPoint();

	/**
	 * <code>optional int32 userGradeId = 23;</code>
	 */
	int getUserGradeId();

	/**
	 * <code>optional int64 regTime = 24;</code>
	 */
	long getRegTime();

	/**
	 * <code>optional int64 regIp = 25;</code>
	 */
	long getRegIp();

	/**
	 * <code>optional int64 lastLoginIp = 26;</code>
	 */
	long getLastLoginIp();

	/**
	 * <code>optional int64 lastActive = 27;</code>
	 */
	long getLastActive();

	/**
	 * <code>optional int64 userModify = 28;</code>
	 */
	long getUserModify();

	/**
	 * <code>optional int32 isFirstLogin = 29;</code>
	 */
	int getIsFirstLogin();
}
