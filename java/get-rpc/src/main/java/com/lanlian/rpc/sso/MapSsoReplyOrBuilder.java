// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: sso.proto

package com.lanlian.rpc.sso;

public interface MapSsoReplyOrBuilder extends
		// @@protoc_insertion_point(interface_extends:pb.MapSsoReply)
		com.google.protobuf.MessageOrBuilder {

	/**
	 * <code>map&lt;int32, .pb.SsoReply&gt; ssos = 1;</code>
	 */
	java.util.Map<java.lang.Integer, com.lanlian.rpc.sso.SsoReply> getSsos();

	/**
	 * <code>optional int32 errorCode = 2;</code>
	 */
	int getErrorCode();
}
