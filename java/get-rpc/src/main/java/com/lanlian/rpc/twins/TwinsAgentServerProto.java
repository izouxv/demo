// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: twins-agent.proto

package com.lanlian.rpc.twins;

public final class TwinsAgentServerProto {
  private TwinsAgentServerProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
  }
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_agent_AddTwinsAgentRequest_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_agent_AddTwinsAgentRequest_fieldAccessorTable;
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_agent_AddTwinsAgentResponse_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_agent_AddTwinsAgentResponse_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\021twins-agent.proto\022\005agent\"(\n\024AddTwinsAg" +
      "entRequest\022\020\n\010reported\030\001 \001(\t\"\027\n\025AddTwins" +
      "AgentResponse2`\n\020TwinsAgentServer\022L\n\rAdd" +
      "TwinsAgent\022\033.agent.AddTwinsAgentRequest\032" +
      "\034.agent.AddTwinsAgentResponse\"\000B7\n\025com.l" +
      "anlian.rpc.twinsB\025TwinsAgentServerProtoP" +
      "\001\242\002\004Grpcb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        }, assigner);
    internal_static_agent_AddTwinsAgentRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_agent_AddTwinsAgentRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_agent_AddTwinsAgentRequest_descriptor,
        new java.lang.String[] { "Reported", });
    internal_static_agent_AddTwinsAgentResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_agent_AddTwinsAgentResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_agent_AddTwinsAgentResponse_descriptor,
        new java.lang.String[] { });
  }

  // @@protoc_insertion_point(outer_class_scope)
}