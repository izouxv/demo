// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: twins-agent.proto

package com.lanlian.rpc.twins;

public interface AddTwinsAgentRequestOrBuilder extends
		// @@protoc_insertion_point(interface_extends:api.AddTwinsAgentRequest)
		com.google.protobuf.MessageOrBuilder {

	/**
	 * <code>optional string reported = 1;</code>
	 *
	 * <pre>
	 *上报临时设备影子的信息，json中必须包含sn字段
	 * </pre>
	 */
	java.lang.String getReported();

	/**
	 * <code>optional string reported = 1;</code>
	 *
	 * <pre>
	 *上报临时设备影子的信息，json中必须包含sn字段
	 * </pre>
	 */
	com.google.protobuf.ByteString getReportedBytes();
}
