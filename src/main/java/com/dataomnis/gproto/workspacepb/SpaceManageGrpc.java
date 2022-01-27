package com.dataomnis.gproto.workspacepb;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * SpaceManage declares API for manage workspace.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.43.2)",
    comments = "Source: proto/workspace.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class SpaceManageGrpc {

  private SpaceManageGrpc() {}

  public static final String SERVICE_NAME = "wspb.SpaceManage";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces,
      com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces> getListWorkspacesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListWorkspaces",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces,
      com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces> getListWorkspacesMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces, com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces> getListWorkspacesMethod;
    if ((getListWorkspacesMethod = SpaceManageGrpc.getListWorkspacesMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getListWorkspacesMethod = SpaceManageGrpc.getListWorkspacesMethod) == null) {
          SpaceManageGrpc.getListWorkspacesMethod = getListWorkspacesMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces, com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListWorkspaces"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("ListWorkspaces"))
              .build();
        }
      }
    }
    return getListWorkspacesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDeleteWorkspacesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteWorkspaces",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDeleteWorkspacesMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDeleteWorkspacesMethod;
    if ((getDeleteWorkspacesMethod = SpaceManageGrpc.getDeleteWorkspacesMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getDeleteWorkspacesMethod = SpaceManageGrpc.getDeleteWorkspacesMethod) == null) {
          SpaceManageGrpc.getDeleteWorkspacesMethod = getDeleteWorkspacesMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteWorkspaces"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("DeleteWorkspaces"))
              .build();
        }
      }
    }
    return getDeleteWorkspacesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDisableWorkspacesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DisableWorkspaces",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDisableWorkspacesMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDisableWorkspacesMethod;
    if ((getDisableWorkspacesMethod = SpaceManageGrpc.getDisableWorkspacesMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getDisableWorkspacesMethod = SpaceManageGrpc.getDisableWorkspacesMethod) == null) {
          SpaceManageGrpc.getDisableWorkspacesMethod = getDisableWorkspacesMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DisableWorkspaces"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("DisableWorkspaces"))
              .build();
        }
      }
    }
    return getDisableWorkspacesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getEnableWorkspacesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "EnableWorkspaces",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getEnableWorkspacesMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getEnableWorkspacesMethod;
    if ((getEnableWorkspacesMethod = SpaceManageGrpc.getEnableWorkspacesMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getEnableWorkspacesMethod = SpaceManageGrpc.getEnableWorkspacesMethod) == null) {
          SpaceManageGrpc.getEnableWorkspacesMethod = getEnableWorkspacesMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "EnableWorkspaces"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("EnableWorkspaces"))
              .build();
        }
      }
    }
    return getEnableWorkspacesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace,
      com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace> getCreateWorkspaceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateWorkspace",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace,
      com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace> getCreateWorkspaceMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace, com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace> getCreateWorkspaceMethod;
    if ((getCreateWorkspaceMethod = SpaceManageGrpc.getCreateWorkspaceMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getCreateWorkspaceMethod = SpaceManageGrpc.getCreateWorkspaceMethod) == null) {
          SpaceManageGrpc.getCreateWorkspaceMethod = getCreateWorkspaceMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace, com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateWorkspace"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("CreateWorkspace"))
              .build();
        }
      }
    }
    return getCreateWorkspaceMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getUpdateWorkspaceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateWorkspace",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getUpdateWorkspaceMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getUpdateWorkspaceMethod;
    if ((getUpdateWorkspaceMethod = SpaceManageGrpc.getUpdateWorkspaceMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getUpdateWorkspaceMethod = SpaceManageGrpc.getUpdateWorkspaceMethod) == null) {
          SpaceManageGrpc.getUpdateWorkspaceMethod = getUpdateWorkspaceMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateWorkspace"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("UpdateWorkspace"))
              .build();
        }
      }
    }
    return getUpdateWorkspaceMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace,
      com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace> getDescribeWorkspaceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DescribeWorkspace",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace,
      com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace> getDescribeWorkspaceMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace, com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace> getDescribeWorkspaceMethod;
    if ((getDescribeWorkspaceMethod = SpaceManageGrpc.getDescribeWorkspaceMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getDescribeWorkspaceMethod = SpaceManageGrpc.getDescribeWorkspaceMethod) == null) {
          SpaceManageGrpc.getDescribeWorkspaceMethod = getDescribeWorkspaceMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace, com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DescribeWorkspace"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("DescribeWorkspace"))
              .build();
        }
      }
    }
    return getDescribeWorkspaceMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CheckPermission,
      com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission> getCheckPermissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CheckPermission",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.CheckPermission.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CheckPermission,
      com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission> getCheckPermissionMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CheckPermission, com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission> getCheckPermissionMethod;
    if ((getCheckPermissionMethod = SpaceManageGrpc.getCheckPermissionMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getCheckPermissionMethod = SpaceManageGrpc.getCheckPermissionMethod) == null) {
          SpaceManageGrpc.getCheckPermissionMethod = getCheckPermissionMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.CheckPermission, com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CheckPermission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.CheckPermission.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("CheckPermission"))
              .build();
        }
      }
    }
    return getCheckPermissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits,
      com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits> getListOpAuditsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListOpAudits",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits,
      com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits> getListOpAuditsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits, com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits> getListOpAuditsMethod;
    if ((getListOpAuditsMethod = SpaceManageGrpc.getListOpAuditsMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getListOpAuditsMethod = SpaceManageGrpc.getListOpAuditsMethod) == null) {
          SpaceManageGrpc.getListOpAuditsMethod = getListOpAuditsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits, com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListOpAudits"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("ListOpAudits"))
              .build();
        }
      }
    }
    return getListOpAuditsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getAddOpAuditMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AddOpAudit",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getAddOpAuditMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getAddOpAuditMethod;
    if ((getAddOpAuditMethod = SpaceManageGrpc.getAddOpAuditMethod) == null) {
      synchronized (SpaceManageGrpc.class) {
        if ((getAddOpAuditMethod = SpaceManageGrpc.getAddOpAuditMethod) == null) {
          SpaceManageGrpc.getAddOpAuditMethod = getAddOpAuditMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AddOpAudit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new SpaceManageMethodDescriptorSupplier("AddOpAudit"))
              .build();
        }
      }
    }
    return getAddOpAuditMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SpaceManageStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SpaceManageStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SpaceManageStub>() {
        @java.lang.Override
        public SpaceManageStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SpaceManageStub(channel, callOptions);
        }
      };
    return SpaceManageStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SpaceManageBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SpaceManageBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SpaceManageBlockingStub>() {
        @java.lang.Override
        public SpaceManageBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SpaceManageBlockingStub(channel, callOptions);
        }
      };
    return SpaceManageBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SpaceManageFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SpaceManageFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SpaceManageFutureStub>() {
        @java.lang.Override
        public SpaceManageFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SpaceManageFutureStub(channel, callOptions);
        }
      };
    return SpaceManageFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * SpaceManage declares API for manage workspace.
   * </pre>
   */
  public static abstract class SpaceManageImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * API of workspace manager.
     * </pre>
     */
    public void listWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListWorkspacesMethod(), responseObserver);
    }

    /**
     */
    public void deleteWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteWorkspacesMethod(), responseObserver);
    }

    /**
     */
    public void disableWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDisableWorkspacesMethod(), responseObserver);
    }

    /**
     */
    public void enableWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getEnableWorkspacesMethod(), responseObserver);
    }

    /**
     */
    public void createWorkspace(com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateWorkspaceMethod(), responseObserver);
    }

    /**
     */
    public void updateWorkspace(com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateWorkspaceMethod(), responseObserver);
    }

    /**
     */
    public void describeWorkspace(com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDescribeWorkspaceMethod(), responseObserver);
    }

    /**
     * <pre>
     * Permission Check.
     * Notice: cannot check API includes: ListWorkspaces, DeleteWorkspaces, DisableWorkspaces, EnableWorkspaces.
     * </pre>
     */
    public void checkPermission(com.dataomnis.gproto.requestpb.RequestPB.CheckPermission request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCheckPermissionMethod(), responseObserver);
    }

    /**
     * <pre>
     * API of workspace operation audit log.
     * </pre>
     */
    public void listOpAudits(com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListOpAuditsMethod(), responseObserver);
    }

    /**
     */
    public void addOpAudit(com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAddOpAuditMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListWorkspacesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces,
                com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces>(
                  this, METHODID_LIST_WORKSPACES)))
          .addMethod(
            getDeleteWorkspacesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_DELETE_WORKSPACES)))
          .addMethod(
            getDisableWorkspacesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_DISABLE_WORKSPACES)))
          .addMethod(
            getEnableWorkspacesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_ENABLE_WORKSPACES)))
          .addMethod(
            getCreateWorkspaceMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace,
                com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace>(
                  this, METHODID_CREATE_WORKSPACE)))
          .addMethod(
            getUpdateWorkspaceMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_UPDATE_WORKSPACE)))
          .addMethod(
            getDescribeWorkspaceMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace,
                com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace>(
                  this, METHODID_DESCRIBE_WORKSPACE)))
          .addMethod(
            getCheckPermissionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.CheckPermission,
                com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission>(
                  this, METHODID_CHECK_PERMISSION)))
          .addMethod(
            getListOpAuditsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits,
                com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits>(
                  this, METHODID_LIST_OP_AUDITS)))
          .addMethod(
            getAddOpAuditMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_ADD_OP_AUDIT)))
          .build();
    }
  }

  /**
   * <pre>
   * SpaceManage declares API for manage workspace.
   * </pre>
   */
  public static final class SpaceManageStub extends io.grpc.stub.AbstractAsyncStub<SpaceManageStub> {
    private SpaceManageStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SpaceManageStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SpaceManageStub(channel, callOptions);
    }

    /**
     * <pre>
     * API of workspace manager.
     * </pre>
     */
    public void listWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListWorkspacesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteWorkspacesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void disableWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDisableWorkspacesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void enableWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getEnableWorkspacesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createWorkspace(com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateWorkspaceMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateWorkspace(com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateWorkspaceMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void describeWorkspace(com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDescribeWorkspaceMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Permission Check.
     * Notice: cannot check API includes: ListWorkspaces, DeleteWorkspaces, DisableWorkspaces, EnableWorkspaces.
     * </pre>
     */
    public void checkPermission(com.dataomnis.gproto.requestpb.RequestPB.CheckPermission request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCheckPermissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * API of workspace operation audit log.
     * </pre>
     */
    public void listOpAudits(com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListOpAuditsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void addOpAudit(com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAddOpAuditMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * SpaceManage declares API for manage workspace.
   * </pre>
   */
  public static final class SpaceManageBlockingStub extends io.grpc.stub.AbstractBlockingStub<SpaceManageBlockingStub> {
    private SpaceManageBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SpaceManageBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SpaceManageBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * API of workspace manager.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces listWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListWorkspacesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct deleteWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteWorkspacesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct disableWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDisableWorkspacesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct enableWorkspaces(com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getEnableWorkspacesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace createWorkspace(com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateWorkspaceMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct updateWorkspace(com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateWorkspaceMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace describeWorkspace(com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDescribeWorkspaceMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Permission Check.
     * Notice: cannot check API includes: ListWorkspaces, DeleteWorkspaces, DisableWorkspaces, EnableWorkspaces.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission checkPermission(com.dataomnis.gproto.requestpb.RequestPB.CheckPermission request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCheckPermissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * API of workspace operation audit log.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits listOpAudits(com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListOpAuditsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct addOpAudit(com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAddOpAuditMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * SpaceManage declares API for manage workspace.
   * </pre>
   */
  public static final class SpaceManageFutureStub extends io.grpc.stub.AbstractFutureStub<SpaceManageFutureStub> {
    private SpaceManageFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SpaceManageFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SpaceManageFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * API of workspace manager.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces> listWorkspaces(
        com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListWorkspacesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> deleteWorkspaces(
        com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteWorkspacesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> disableWorkspaces(
        com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDisableWorkspacesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> enableWorkspaces(
        com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getEnableWorkspacesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace> createWorkspace(
        com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateWorkspaceMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> updateWorkspace(
        com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateWorkspaceMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace> describeWorkspace(
        com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDescribeWorkspaceMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Permission Check.
     * Notice: cannot check API includes: ListWorkspaces, DeleteWorkspaces, DisableWorkspaces, EnableWorkspaces.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission> checkPermission(
        com.dataomnis.gproto.requestpb.RequestPB.CheckPermission request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCheckPermissionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * API of workspace operation audit log.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits> listOpAudits(
        com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListOpAuditsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> addOpAudit(
        com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAddOpAuditMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_WORKSPACES = 0;
  private static final int METHODID_DELETE_WORKSPACES = 1;
  private static final int METHODID_DISABLE_WORKSPACES = 2;
  private static final int METHODID_ENABLE_WORKSPACES = 3;
  private static final int METHODID_CREATE_WORKSPACE = 4;
  private static final int METHODID_UPDATE_WORKSPACE = 5;
  private static final int METHODID_DESCRIBE_WORKSPACE = 6;
  private static final int METHODID_CHECK_PERMISSION = 7;
  private static final int METHODID_LIST_OP_AUDITS = 8;
  private static final int METHODID_ADD_OP_AUDIT = 9;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SpaceManageImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SpaceManageImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST_WORKSPACES:
          serviceImpl.listWorkspaces((com.dataomnis.gproto.requestpb.RequestPB.ListWorkspaces) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListWorkspaces>) responseObserver);
          break;
        case METHODID_DELETE_WORKSPACES:
          serviceImpl.deleteWorkspaces((com.dataomnis.gproto.requestpb.RequestPB.DeleteWorkspaces) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_DISABLE_WORKSPACES:
          serviceImpl.disableWorkspaces((com.dataomnis.gproto.requestpb.RequestPB.DisableWorkspaces) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_ENABLE_WORKSPACES:
          serviceImpl.enableWorkspaces((com.dataomnis.gproto.requestpb.RequestPB.EnableWorkspaces) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_CREATE_WORKSPACE:
          serviceImpl.createWorkspace((com.dataomnis.gproto.requestpb.RequestPB.CreateWorkspace) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CreateWorkspace>) responseObserver);
          break;
        case METHODID_UPDATE_WORKSPACE:
          serviceImpl.updateWorkspace((com.dataomnis.gproto.requestpb.RequestPB.UpdateWorkspace) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_DESCRIBE_WORKSPACE:
          serviceImpl.describeWorkspace((com.dataomnis.gproto.requestpb.RequestPB.DescribeWorkspace) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeWorkspace>) responseObserver);
          break;
        case METHODID_CHECK_PERMISSION:
          serviceImpl.checkPermission((com.dataomnis.gproto.requestpb.RequestPB.CheckPermission) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CheckPermission>) responseObserver);
          break;
        case METHODID_LIST_OP_AUDITS:
          serviceImpl.listOpAudits((com.dataomnis.gproto.requestpb.RequestPB.ListOpAudits) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListOpAudits>) responseObserver);
          break;
        case METHODID_ADD_OP_AUDIT:
          serviceImpl.addOpAudit((com.dataomnis.gproto.requestpb.RequestPB.AddOpAudit) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class SpaceManageBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SpaceManageBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.dataomnis.gproto.workspacepb.WorkspacePB.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("SpaceManage");
    }
  }

  private static final class SpaceManageFileDescriptorSupplier
      extends SpaceManageBaseDescriptorSupplier {
    SpaceManageFileDescriptorSupplier() {}
  }

  private static final class SpaceManageMethodDescriptorSupplier
      extends SpaceManageBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    SpaceManageMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (SpaceManageGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SpaceManageFileDescriptorSupplier())
              .addMethod(getListWorkspacesMethod())
              .addMethod(getDeleteWorkspacesMethod())
              .addMethod(getDisableWorkspacesMethod())
              .addMethod(getEnableWorkspacesMethod())
              .addMethod(getCreateWorkspaceMethod())
              .addMethod(getUpdateWorkspaceMethod())
              .addMethod(getDescribeWorkspaceMethod())
              .addMethod(getCheckPermissionMethod())
              .addMethod(getListOpAuditsMethod())
              .addMethod(getAddOpAuditMethod())
              .build();
        }
      }
    }
    return result;
  }
}