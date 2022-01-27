package com.dataomnis.gproto.workspacepb;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.43.2)",
    comments = "Source: proto/workspace.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class StreamJobManageGrpc {

  private StreamJobManageGrpc() {}

  public static final String SERVICE_NAME = "wspb.StreamJobManage";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs,
      com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs> getListStreamJobsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListStreamJobs",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs,
      com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs> getListStreamJobsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs, com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs> getListStreamJobsMethod;
    if ((getListStreamJobsMethod = StreamJobManageGrpc.getListStreamJobsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getListStreamJobsMethod = StreamJobManageGrpc.getListStreamJobsMethod) == null) {
          StreamJobManageGrpc.getListStreamJobsMethod = getListStreamJobsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs, com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListStreamJobs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("ListStreamJobs"))
              .build();
        }
      }
    }
    return getListStreamJobsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDeleteStreamJobsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteStreamJobs",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDeleteStreamJobsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getDeleteStreamJobsMethod;
    if ((getDeleteStreamJobsMethod = StreamJobManageGrpc.getDeleteStreamJobsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getDeleteStreamJobsMethod = StreamJobManageGrpc.getDeleteStreamJobsMethod) == null) {
          StreamJobManageGrpc.getDeleteStreamJobsMethod = getDeleteStreamJobsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteStreamJobs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("DeleteStreamJobs"))
              .build();
        }
      }
    }
    return getDeleteStreamJobsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob,
      com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob> getCreateStreamJobMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateStreamJob",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob,
      com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob> getCreateStreamJobMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob, com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob> getCreateStreamJobMethod;
    if ((getCreateStreamJobMethod = StreamJobManageGrpc.getCreateStreamJobMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getCreateStreamJobMethod = StreamJobManageGrpc.getCreateStreamJobMethod) == null) {
          StreamJobManageGrpc.getCreateStreamJobMethod = getCreateStreamJobMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob, com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateStreamJob"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("CreateStreamJob"))
              .build();
        }
      }
    }
    return getCreateStreamJobMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getUpdateStreamJobMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateStreamJob",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getUpdateStreamJobMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getUpdateStreamJobMethod;
    if ((getUpdateStreamJobMethod = StreamJobManageGrpc.getUpdateStreamJobMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getUpdateStreamJobMethod = StreamJobManageGrpc.getUpdateStreamJobMethod) == null) {
          StreamJobManageGrpc.getUpdateStreamJobMethod = getUpdateStreamJobMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateStreamJob"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("UpdateStreamJob"))
              .build();
        }
      }
    }
    return getUpdateStreamJobMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob,
      com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> getDescribeStreamJobMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DescribeStreamJob",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob,
      com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> getDescribeStreamJobMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob, com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> getDescribeStreamJobMethod;
    if ((getDescribeStreamJobMethod = StreamJobManageGrpc.getDescribeStreamJobMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getDescribeStreamJobMethod = StreamJobManageGrpc.getDescribeStreamJobMethod) == null) {
          StreamJobManageGrpc.getDescribeStreamJobMethod = getDescribeStreamJobMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob, com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DescribeStreamJob"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("DescribeStreamJob"))
              .build();
        }
      }
    }
    return getDescribeStreamJobMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobCodeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetStreamJobCode",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobCodeMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobCodeMethod;
    if ((getSetStreamJobCodeMethod = StreamJobManageGrpc.getSetStreamJobCodeMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getSetStreamJobCodeMethod = StreamJobManageGrpc.getSetStreamJobCodeMethod) == null) {
          StreamJobManageGrpc.getSetStreamJobCodeMethod = getSetStreamJobCodeMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetStreamJobCode"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("SetStreamJobCode"))
              .build();
        }
      }
    }
    return getSetStreamJobCodeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getGetStreamJobCodeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStreamJobCode",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getGetStreamJobCodeMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getGetStreamJobCodeMethod;
    if ((getGetStreamJobCodeMethod = StreamJobManageGrpc.getGetStreamJobCodeMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getGetStreamJobCodeMethod = StreamJobManageGrpc.getGetStreamJobCodeMethod) == null) {
          StreamJobManageGrpc.getGetStreamJobCodeMethod = getGetStreamJobCodeMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStreamJobCode"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("GetStreamJobCode"))
              .build();
        }
      }
    }
    return getGetStreamJobCodeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax,
      com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax> getStreamJobCodeSyntaxMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StreamJobCodeSyntax",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax,
      com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax> getStreamJobCodeSyntaxMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax, com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax> getStreamJobCodeSyntaxMethod;
    if ((getStreamJobCodeSyntaxMethod = StreamJobManageGrpc.getStreamJobCodeSyntaxMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getStreamJobCodeSyntaxMethod = StreamJobManageGrpc.getStreamJobCodeSyntaxMethod) == null) {
          StreamJobManageGrpc.getStreamJobCodeSyntaxMethod = getStreamJobCodeSyntaxMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax, com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StreamJobCodeSyntax"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("StreamJobCodeSyntax"))
              .build();
        }
      }
    }
    return getStreamJobCodeSyntaxMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobArgsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetStreamJobArgs",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobArgsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobArgsMethod;
    if ((getSetStreamJobArgsMethod = StreamJobManageGrpc.getSetStreamJobArgsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getSetStreamJobArgsMethod = StreamJobManageGrpc.getSetStreamJobArgsMethod) == null) {
          StreamJobManageGrpc.getSetStreamJobArgsMethod = getSetStreamJobArgsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetStreamJobArgs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("SetStreamJobArgs"))
              .build();
        }
      }
    }
    return getSetStreamJobArgsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getGetStreamJobArgsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStreamJobArgs",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getGetStreamJobArgsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getGetStreamJobArgsMethod;
    if ((getGetStreamJobArgsMethod = StreamJobManageGrpc.getGetStreamJobArgsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getGetStreamJobArgsMethod = StreamJobManageGrpc.getGetStreamJobArgsMethod) == null) {
          StreamJobManageGrpc.getGetStreamJobArgsMethod = getGetStreamJobArgsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStreamJobArgs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("GetStreamJobArgs"))
              .build();
        }
      }
    }
    return getGetStreamJobArgsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct,
      com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors> getListBuiltInConnectorsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListBuiltInConnectors",
      requestType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct,
      com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors> getListBuiltInConnectorsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct, com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors> getListBuiltInConnectorsMethod;
    if ((getListBuiltInConnectorsMethod = StreamJobManageGrpc.getListBuiltInConnectorsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getListBuiltInConnectorsMethod = StreamJobManageGrpc.getListBuiltInConnectorsMethod) == null) {
          StreamJobManageGrpc.getListBuiltInConnectorsMethod = getListBuiltInConnectorsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct, com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListBuiltInConnectors"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("ListBuiltInConnectors"))
              .build();
        }
      }
    }
    return getListBuiltInConnectorsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobScheduleMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetStreamJobSchedule",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobScheduleMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSetStreamJobScheduleMethod;
    if ((getSetStreamJobScheduleMethod = StreamJobManageGrpc.getSetStreamJobScheduleMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getSetStreamJobScheduleMethod = StreamJobManageGrpc.getSetStreamJobScheduleMethod) == null) {
          StreamJobManageGrpc.getSetStreamJobScheduleMethod = getSetStreamJobScheduleMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetStreamJobSchedule"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("SetStreamJobSchedule"))
              .build();
        }
      }
    }
    return getSetStreamJobScheduleMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getGetStreamJobScheduleMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStreamJobSchedule",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getGetStreamJobScheduleMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getGetStreamJobScheduleMethod;
    if ((getGetStreamJobScheduleMethod = StreamJobManageGrpc.getGetStreamJobScheduleMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getGetStreamJobScheduleMethod = StreamJobManageGrpc.getGetStreamJobScheduleMethod) == null) {
          StreamJobManageGrpc.getGetStreamJobScheduleMethod = getGetStreamJobScheduleMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStreamJobSchedule"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("GetStreamJobSchedule"))
              .build();
        }
      }
    }
    return getGetStreamJobScheduleMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getReleaseStreamJobMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ReleaseStreamJob",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getReleaseStreamJobMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getReleaseStreamJobMethod;
    if ((getReleaseStreamJobMethod = StreamJobManageGrpc.getReleaseStreamJobMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getReleaseStreamJobMethod = StreamJobManageGrpc.getReleaseStreamJobMethod) == null) {
          StreamJobManageGrpc.getReleaseStreamJobMethod = getReleaseStreamJobMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ReleaseStreamJob"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("ReleaseStreamJob"))
              .build();
        }
      }
    }
    return getReleaseStreamJobMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getOfflineReleaseStreamJobMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "OfflineReleaseStreamJob",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getOfflineReleaseStreamJobMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getOfflineReleaseStreamJobMethod;
    if ((getOfflineReleaseStreamJobMethod = StreamJobManageGrpc.getOfflineReleaseStreamJobMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getOfflineReleaseStreamJobMethod = StreamJobManageGrpc.getOfflineReleaseStreamJobMethod) == null) {
          StreamJobManageGrpc.getOfflineReleaseStreamJobMethod = getOfflineReleaseStreamJobMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "OfflineReleaseStreamJob"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("OfflineReleaseStreamJob"))
              .build();
        }
      }
    }
    return getOfflineReleaseStreamJobMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSuspendReleaseStreamJobMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SuspendReleaseStreamJob",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSuspendReleaseStreamJobMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getSuspendReleaseStreamJobMethod;
    if ((getSuspendReleaseStreamJobMethod = StreamJobManageGrpc.getSuspendReleaseStreamJobMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getSuspendReleaseStreamJobMethod = StreamJobManageGrpc.getSuspendReleaseStreamJobMethod) == null) {
          StreamJobManageGrpc.getSuspendReleaseStreamJobMethod = getSuspendReleaseStreamJobMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SuspendReleaseStreamJob"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("SuspendReleaseStreamJob"))
              .build();
        }
      }
    }
    return getSuspendReleaseStreamJobMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getResumeReleaseStreamJobMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ResumeReleaseStreamJob",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob.class,
      responseType = com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob,
      com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getResumeReleaseStreamJobMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> getResumeReleaseStreamJobMethod;
    if ((getResumeReleaseStreamJobMethod = StreamJobManageGrpc.getResumeReleaseStreamJobMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getResumeReleaseStreamJobMethod = StreamJobManageGrpc.getResumeReleaseStreamJobMethod) == null) {
          StreamJobManageGrpc.getResumeReleaseStreamJobMethod = getResumeReleaseStreamJobMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob, com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ResumeReleaseStreamJob"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("ResumeReleaseStreamJob"))
              .build();
        }
      }
    }
    return getResumeReleaseStreamJobMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs,
      com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs> getListReleaseStreamJobsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListReleaseStreamJobs",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs,
      com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs> getListReleaseStreamJobsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs, com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs> getListReleaseStreamJobsMethod;
    if ((getListReleaseStreamJobsMethod = StreamJobManageGrpc.getListReleaseStreamJobsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getListReleaseStreamJobsMethod = StreamJobManageGrpc.getListReleaseStreamJobsMethod) == null) {
          StreamJobManageGrpc.getListReleaseStreamJobsMethod = getListReleaseStreamJobsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs, com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListReleaseStreamJobs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("ListReleaseStreamJobs"))
              .build();
        }
      }
    }
    return getListReleaseStreamJobsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions,
      com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions> getListStreamJobVersionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListStreamJobVersions",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions,
      com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions> getListStreamJobVersionsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions, com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions> getListStreamJobVersionsMethod;
    if ((getListStreamJobVersionsMethod = StreamJobManageGrpc.getListStreamJobVersionsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getListStreamJobVersionsMethod = StreamJobManageGrpc.getListStreamJobVersionsMethod) == null) {
          StreamJobManageGrpc.getListStreamJobVersionsMethod = getListStreamJobVersionsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions, com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListStreamJobVersions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("ListStreamJobVersions"))
              .build();
        }
      }
    }
    return getListStreamJobVersionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob,
      com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> getDescribeStreamJobVersionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DescribeStreamJobVersion",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob,
      com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> getDescribeStreamJobVersionMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob, com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> getDescribeStreamJobVersionMethod;
    if ((getDescribeStreamJobVersionMethod = StreamJobManageGrpc.getDescribeStreamJobVersionMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getDescribeStreamJobVersionMethod = StreamJobManageGrpc.getDescribeStreamJobVersionMethod) == null) {
          StreamJobManageGrpc.getDescribeStreamJobVersionMethod = getDescribeStreamJobVersionMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob, com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DescribeStreamJobVersion"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("DescribeStreamJobVersion"))
              .build();
        }
      }
    }
    return getDescribeStreamJobVersionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getGetStreamJobVersionCodeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStreamJobVersionCode",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getGetStreamJobVersionCodeMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getGetStreamJobVersionCodeMethod;
    if ((getGetStreamJobVersionCodeMethod = StreamJobManageGrpc.getGetStreamJobVersionCodeMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getGetStreamJobVersionCodeMethod = StreamJobManageGrpc.getGetStreamJobVersionCodeMethod) == null) {
          StreamJobManageGrpc.getGetStreamJobVersionCodeMethod = getGetStreamJobVersionCodeMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStreamJobVersionCode"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("GetStreamJobVersionCode"))
              .build();
        }
      }
    }
    return getGetStreamJobVersionCodeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getGetStreamJobVersionArgsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStreamJobVersionArgs",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getGetStreamJobVersionArgsMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getGetStreamJobVersionArgsMethod;
    if ((getGetStreamJobVersionArgsMethod = StreamJobManageGrpc.getGetStreamJobVersionArgsMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getGetStreamJobVersionArgsMethod = StreamJobManageGrpc.getGetStreamJobVersionArgsMethod) == null) {
          StreamJobManageGrpc.getGetStreamJobVersionArgsMethod = getGetStreamJobVersionArgsMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStreamJobVersionArgs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("GetStreamJobVersionArgs"))
              .build();
        }
      }
    }
    return getGetStreamJobVersionArgsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getGetStreamJobVersionScheduleMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStreamJobVersionSchedule",
      requestType = com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule.class,
      responseType = com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule,
      com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getGetStreamJobVersionScheduleMethod() {
    io.grpc.MethodDescriptor<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getGetStreamJobVersionScheduleMethod;
    if ((getGetStreamJobVersionScheduleMethod = StreamJobManageGrpc.getGetStreamJobVersionScheduleMethod) == null) {
      synchronized (StreamJobManageGrpc.class) {
        if ((getGetStreamJobVersionScheduleMethod = StreamJobManageGrpc.getGetStreamJobVersionScheduleMethod) == null) {
          StreamJobManageGrpc.getGetStreamJobVersionScheduleMethod = getGetStreamJobVersionScheduleMethod =
              io.grpc.MethodDescriptor.<com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule, com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStreamJobVersionSchedule"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule.getDefaultInstance()))
              .setSchemaDescriptor(new StreamJobManageMethodDescriptorSupplier("GetStreamJobVersionSchedule"))
              .build();
        }
      }
    }
    return getGetStreamJobVersionScheduleMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static StreamJobManageStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StreamJobManageStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StreamJobManageStub>() {
        @java.lang.Override
        public StreamJobManageStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StreamJobManageStub(channel, callOptions);
        }
      };
    return StreamJobManageStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static StreamJobManageBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StreamJobManageBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StreamJobManageBlockingStub>() {
        @java.lang.Override
        public StreamJobManageBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StreamJobManageBlockingStub(channel, callOptions);
        }
      };
    return StreamJobManageBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static StreamJobManageFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StreamJobManageFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StreamJobManageFutureStub>() {
        @java.lang.Override
        public StreamJobManageFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StreamJobManageFutureStub(channel, callOptions);
        }
      };
    return StreamJobManageFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class StreamJobManageImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * Interface for stream job dev.
     * ListStreamJobs to get a list of stream job of the workspace.
     * </pre>
     */
    public void listStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListStreamJobsMethod(), responseObserver);
    }

    /**
     * <pre>
     * DeleteStreamJobs delete stream job ant its related resources where in ids.
     * Resources includes:
     *  - History version and Released's job.
     *  - Node, env schedule and its history version.
     *  - Offline job and force stop all running instances. (By Scheduler-Server)
     *  - All instances records. (By Scheduler-Server)
     * </pre>
     */
    public void deleteStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteStreamJobsMethod(), responseObserver);
    }

    /**
     * <pre>
     * CreateStreamJob to create a new stream job.
     * </pre>
     */
    public void createStreamJob(com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateStreamJobMethod(), responseObserver);
    }

    /**
     * <pre>
     * UpdateStreamJob to update the info for the specified stream job.
     * </pre>
     */
    public void updateStreamJob(com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateStreamJobMethod(), responseObserver);
    }

    /**
     * <pre>
     * DescribeStreamJob to get the info of the specified stream job.
     * </pre>
     */
    public void describeStreamJob(com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDescribeStreamJobMethod(), responseObserver);
    }

    /**
     * <pre>
     * SetStreamJobCode to set the code properties of the specified stream job.
     * </pre>
     */
    public void setStreamJobCode(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetStreamJobCodeMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobCode to get the code properties of the specified stream job.
     * </pre>
     */
    public void getStreamJobCode(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStreamJobCodeMethod(), responseObserver);
    }

    /**
     */
    public void streamJobCodeSyntax(com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStreamJobCodeSyntaxMethod(), responseObserver);
    }

    /**
     * <pre>
     * SetStreamJobArgs to set the run parameters of the specified stream job.
     * </pre>
     */
    public void setStreamJobArgs(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetStreamJobArgsMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobArgs to get the run parameters of the specified stream job.
     * </pre>
     */
    public void getStreamJobArgs(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStreamJobArgsMethod(), responseObserver);
    }

    /**
     */
    public void listBuiltInConnectors(com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListBuiltInConnectorsMethod(), responseObserver);
    }

    /**
     * <pre>
     * SetStreamJobSchedule to set the schedule properties of the specified stream job.
     * </pre>
     */
    public void setStreamJobSchedule(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetStreamJobScheduleMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobSchedule to get the schedule properties of the specified stream job.
     * </pre>
     */
    public void getStreamJobSchedule(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStreamJobScheduleMethod(), responseObserver);
    }

    /**
     * <pre>
     * Interface for stream job release.
     * ReleaseStreamJob to publish the specified job to schedule system with a new version.
     * </pre>
     */
    public void releaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReleaseStreamJobMethod(), responseObserver);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJobs to suspend the specified job list in schedule system.
     * </pre>
     */
    public void offlineReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getOfflineReleaseStreamJobMethod(), responseObserver);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJob to suspend the specified job list in schedule system.
     * </pre>
     */
    public void suspendReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSuspendReleaseStreamJobMethod(), responseObserver);
    }

    /**
     * <pre>
     * ResumeReleaseStreamJob to resume the suspended job list in schedule system.
     * </pre>
     */
    public void resumeReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getResumeReleaseStreamJobMethod(), responseObserver);
    }

    /**
     * <pre>
     * ListReleaseStreamJobs for gets a list of all published job in the workspace.
     * </pre>
     */
    public void listReleaseStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListReleaseStreamJobsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Interface for stream job versions.
     * ListStreamJobVersions for gets a list of all versions of the specified job.
     * </pre>
     */
    public void listStreamJobVersions(com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListStreamJobVersionsMethod(), responseObserver);
    }

    /**
     * <pre>
     * DescribeStreamJobVersion for get the info of the job of the specified version.
     * </pre>
     */
    public void describeStreamJobVersion(com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDescribeStreamJobVersionMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobVersionCode for get the node properties of the job of the specified version.
     * </pre>
     */
    public void getStreamJobVersionCode(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStreamJobVersionCodeMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobVersionArgs for get the run parameters of the job of the specified version.
     * </pre>
     */
    public void getStreamJobVersionArgs(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStreamJobVersionArgsMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobVersionSchedule for get the schedule properties of the job of the specified version.
     * </pre>
     */
    public void getStreamJobVersionSchedule(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStreamJobVersionScheduleMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListStreamJobsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs,
                com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs>(
                  this, METHODID_LIST_STREAM_JOBS)))
          .addMethod(
            getDeleteStreamJobsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_DELETE_STREAM_JOBS)))
          .addMethod(
            getCreateStreamJobMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob,
                com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob>(
                  this, METHODID_CREATE_STREAM_JOB)))
          .addMethod(
            getUpdateStreamJobMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_UPDATE_STREAM_JOB)))
          .addMethod(
            getDescribeStreamJobMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob,
                com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob>(
                  this, METHODID_DESCRIBE_STREAM_JOB)))
          .addMethod(
            getSetStreamJobCodeMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_SET_STREAM_JOB_CODE)))
          .addMethod(
            getGetStreamJobCodeMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode,
                com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode>(
                  this, METHODID_GET_STREAM_JOB_CODE)))
          .addMethod(
            getStreamJobCodeSyntaxMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax,
                com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax>(
                  this, METHODID_STREAM_JOB_CODE_SYNTAX)))
          .addMethod(
            getSetStreamJobArgsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_SET_STREAM_JOB_ARGS)))
          .addMethod(
            getGetStreamJobArgsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs,
                com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs>(
                  this, METHODID_GET_STREAM_JOB_ARGS)))
          .addMethod(
            getListBuiltInConnectorsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct,
                com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors>(
                  this, METHODID_LIST_BUILT_IN_CONNECTORS)))
          .addMethod(
            getSetStreamJobScheduleMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_SET_STREAM_JOB_SCHEDULE)))
          .addMethod(
            getGetStreamJobScheduleMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule,
                com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule>(
                  this, METHODID_GET_STREAM_JOB_SCHEDULE)))
          .addMethod(
            getReleaseStreamJobMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_RELEASE_STREAM_JOB)))
          .addMethod(
            getOfflineReleaseStreamJobMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_OFFLINE_RELEASE_STREAM_JOB)))
          .addMethod(
            getSuspendReleaseStreamJobMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_SUSPEND_RELEASE_STREAM_JOB)))
          .addMethod(
            getResumeReleaseStreamJobMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob,
                com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>(
                  this, METHODID_RESUME_RELEASE_STREAM_JOB)))
          .addMethod(
            getListReleaseStreamJobsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs,
                com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs>(
                  this, METHODID_LIST_RELEASE_STREAM_JOBS)))
          .addMethod(
            getListStreamJobVersionsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions,
                com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions>(
                  this, METHODID_LIST_STREAM_JOB_VERSIONS)))
          .addMethod(
            getDescribeStreamJobVersionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob,
                com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob>(
                  this, METHODID_DESCRIBE_STREAM_JOB_VERSION)))
          .addMethod(
            getGetStreamJobVersionCodeMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode,
                com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode>(
                  this, METHODID_GET_STREAM_JOB_VERSION_CODE)))
          .addMethod(
            getGetStreamJobVersionArgsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs,
                com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs>(
                  this, METHODID_GET_STREAM_JOB_VERSION_ARGS)))
          .addMethod(
            getGetStreamJobVersionScheduleMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule,
                com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule>(
                  this, METHODID_GET_STREAM_JOB_VERSION_SCHEDULE)))
          .build();
    }
  }

  /**
   */
  public static final class StreamJobManageStub extends io.grpc.stub.AbstractAsyncStub<StreamJobManageStub> {
    private StreamJobManageStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StreamJobManageStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StreamJobManageStub(channel, callOptions);
    }

    /**
     * <pre>
     * Interface for stream job dev.
     * ListStreamJobs to get a list of stream job of the workspace.
     * </pre>
     */
    public void listStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListStreamJobsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * DeleteStreamJobs delete stream job ant its related resources where in ids.
     * Resources includes:
     *  - History version and Released's job.
     *  - Node, env schedule and its history version.
     *  - Offline job and force stop all running instances. (By Scheduler-Server)
     *  - All instances records. (By Scheduler-Server)
     * </pre>
     */
    public void deleteStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteStreamJobsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * CreateStreamJob to create a new stream job.
     * </pre>
     */
    public void createStreamJob(com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateStreamJobMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * UpdateStreamJob to update the info for the specified stream job.
     * </pre>
     */
    public void updateStreamJob(com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateStreamJobMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * DescribeStreamJob to get the info of the specified stream job.
     * </pre>
     */
    public void describeStreamJob(com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDescribeStreamJobMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * SetStreamJobCode to set the code properties of the specified stream job.
     * </pre>
     */
    public void setStreamJobCode(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetStreamJobCodeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobCode to get the code properties of the specified stream job.
     * </pre>
     */
    public void getStreamJobCode(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStreamJobCodeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void streamJobCodeSyntax(com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStreamJobCodeSyntaxMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * SetStreamJobArgs to set the run parameters of the specified stream job.
     * </pre>
     */
    public void setStreamJobArgs(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetStreamJobArgsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobArgs to get the run parameters of the specified stream job.
     * </pre>
     */
    public void getStreamJobArgs(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStreamJobArgsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listBuiltInConnectors(com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListBuiltInConnectorsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * SetStreamJobSchedule to set the schedule properties of the specified stream job.
     * </pre>
     */
    public void setStreamJobSchedule(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetStreamJobScheduleMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobSchedule to get the schedule properties of the specified stream job.
     * </pre>
     */
    public void getStreamJobSchedule(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStreamJobScheduleMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Interface for stream job release.
     * ReleaseStreamJob to publish the specified job to schedule system with a new version.
     * </pre>
     */
    public void releaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReleaseStreamJobMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJobs to suspend the specified job list in schedule system.
     * </pre>
     */
    public void offlineReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getOfflineReleaseStreamJobMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJob to suspend the specified job list in schedule system.
     * </pre>
     */
    public void suspendReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSuspendReleaseStreamJobMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * ResumeReleaseStreamJob to resume the suspended job list in schedule system.
     * </pre>
     */
    public void resumeReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getResumeReleaseStreamJobMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * ListReleaseStreamJobs for gets a list of all published job in the workspace.
     * </pre>
     */
    public void listReleaseStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListReleaseStreamJobsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Interface for stream job versions.
     * ListStreamJobVersions for gets a list of all versions of the specified job.
     * </pre>
     */
    public void listStreamJobVersions(com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListStreamJobVersionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * DescribeStreamJobVersion for get the info of the job of the specified version.
     * </pre>
     */
    public void describeStreamJobVersion(com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDescribeStreamJobVersionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobVersionCode for get the node properties of the job of the specified version.
     * </pre>
     */
    public void getStreamJobVersionCode(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStreamJobVersionCodeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobVersionArgs for get the run parameters of the job of the specified version.
     * </pre>
     */
    public void getStreamJobVersionArgs(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStreamJobVersionArgsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetStreamJobVersionSchedule for get the schedule properties of the job of the specified version.
     * </pre>
     */
    public void getStreamJobVersionSchedule(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request,
        io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStreamJobVersionScheduleMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class StreamJobManageBlockingStub extends io.grpc.stub.AbstractBlockingStub<StreamJobManageBlockingStub> {
    private StreamJobManageBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StreamJobManageBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StreamJobManageBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Interface for stream job dev.
     * ListStreamJobs to get a list of stream job of the workspace.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs listStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListStreamJobsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * DeleteStreamJobs delete stream job ant its related resources where in ids.
     * Resources includes:
     *  - History version and Released's job.
     *  - Node, env schedule and its history version.
     *  - Offline job and force stop all running instances. (By Scheduler-Server)
     *  - All instances records. (By Scheduler-Server)
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct deleteStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteStreamJobsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * CreateStreamJob to create a new stream job.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob createStreamJob(com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateStreamJobMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * UpdateStreamJob to update the info for the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct updateStreamJob(com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateStreamJobMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * DescribeStreamJob to get the info of the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob describeStreamJob(com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDescribeStreamJobMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * SetStreamJobCode to set the code properties of the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct setStreamJobCode(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetStreamJobCodeMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetStreamJobCode to get the code properties of the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode getStreamJobCode(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStreamJobCodeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax streamJobCodeSyntax(com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStreamJobCodeSyntaxMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * SetStreamJobArgs to set the run parameters of the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct setStreamJobArgs(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetStreamJobArgsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetStreamJobArgs to get the run parameters of the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs getStreamJobArgs(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStreamJobArgsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors listBuiltInConnectors(com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListBuiltInConnectorsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * SetStreamJobSchedule to set the schedule properties of the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct setStreamJobSchedule(com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetStreamJobScheduleMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetStreamJobSchedule to get the schedule properties of the specified stream job.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule getStreamJobSchedule(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStreamJobScheduleMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Interface for stream job release.
     * ReleaseStreamJob to publish the specified job to schedule system with a new version.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct releaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReleaseStreamJobMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJobs to suspend the specified job list in schedule system.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct offlineReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getOfflineReleaseStreamJobMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJob to suspend the specified job list in schedule system.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct suspendReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSuspendReleaseStreamJobMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * ResumeReleaseStreamJob to resume the suspended job list in schedule system.
     * </pre>
     */
    public com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct resumeReleaseStreamJob(com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getResumeReleaseStreamJobMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * ListReleaseStreamJobs for gets a list of all published job in the workspace.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs listReleaseStreamJobs(com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListReleaseStreamJobsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Interface for stream job versions.
     * ListStreamJobVersions for gets a list of all versions of the specified job.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions listStreamJobVersions(com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListStreamJobVersionsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * DescribeStreamJobVersion for get the info of the job of the specified version.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob describeStreamJobVersion(com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDescribeStreamJobVersionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetStreamJobVersionCode for get the node properties of the job of the specified version.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode getStreamJobVersionCode(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStreamJobVersionCodeMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetStreamJobVersionArgs for get the run parameters of the job of the specified version.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs getStreamJobVersionArgs(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStreamJobVersionArgsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetStreamJobVersionSchedule for get the schedule properties of the job of the specified version.
     * </pre>
     */
    public com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule getStreamJobVersionSchedule(com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStreamJobVersionScheduleMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class StreamJobManageFutureStub extends io.grpc.stub.AbstractFutureStub<StreamJobManageFutureStub> {
    private StreamJobManageFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StreamJobManageFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StreamJobManageFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Interface for stream job dev.
     * ListStreamJobs to get a list of stream job of the workspace.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs> listStreamJobs(
        com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListStreamJobsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * DeleteStreamJobs delete stream job ant its related resources where in ids.
     * Resources includes:
     *  - History version and Released's job.
     *  - Node, env schedule and its history version.
     *  - Offline job and force stop all running instances. (By Scheduler-Server)
     *  - All instances records. (By Scheduler-Server)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> deleteStreamJobs(
        com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteStreamJobsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * CreateStreamJob to create a new stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob> createStreamJob(
        com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateStreamJobMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * UpdateStreamJob to update the info for the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> updateStreamJob(
        com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateStreamJobMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * DescribeStreamJob to get the info of the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> describeStreamJob(
        com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDescribeStreamJobMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * SetStreamJobCode to set the code properties of the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> setStreamJobCode(
        com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetStreamJobCodeMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetStreamJobCode to get the code properties of the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getStreamJobCode(
        com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStreamJobCodeMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax> streamJobCodeSyntax(
        com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStreamJobCodeSyntaxMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * SetStreamJobArgs to set the run parameters of the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> setStreamJobArgs(
        com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetStreamJobArgsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetStreamJobArgs to get the run parameters of the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getStreamJobArgs(
        com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStreamJobArgsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors> listBuiltInConnectors(
        com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListBuiltInConnectorsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * SetStreamJobSchedule to set the schedule properties of the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> setStreamJobSchedule(
        com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetStreamJobScheduleMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetStreamJobSchedule to get the schedule properties of the specified stream job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getStreamJobSchedule(
        com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStreamJobScheduleMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Interface for stream job release.
     * ReleaseStreamJob to publish the specified job to schedule system with a new version.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> releaseStreamJob(
        com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReleaseStreamJobMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJobs to suspend the specified job list in schedule system.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> offlineReleaseStreamJob(
        com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getOfflineReleaseStreamJobMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * SuspendReleaseStreamJob to suspend the specified job list in schedule system.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> suspendReleaseStreamJob(
        com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSuspendReleaseStreamJobMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * ResumeReleaseStreamJob to resume the suspended job list in schedule system.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct> resumeReleaseStreamJob(
        com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getResumeReleaseStreamJobMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * ListReleaseStreamJobs for gets a list of all published job in the workspace.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs> listReleaseStreamJobs(
        com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListReleaseStreamJobsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Interface for stream job versions.
     * ListStreamJobVersions for gets a list of all versions of the specified job.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions> listStreamJobVersions(
        com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListStreamJobVersionsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * DescribeStreamJobVersion for get the info of the job of the specified version.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob> describeStreamJobVersion(
        com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDescribeStreamJobVersionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetStreamJobVersionCode for get the node properties of the job of the specified version.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode> getStreamJobVersionCode(
        com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStreamJobVersionCodeMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetStreamJobVersionArgs for get the run parameters of the job of the specified version.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs> getStreamJobVersionArgs(
        com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStreamJobVersionArgsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetStreamJobVersionSchedule for get the schedule properties of the job of the specified version.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule> getStreamJobVersionSchedule(
        com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStreamJobVersionScheduleMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_STREAM_JOBS = 0;
  private static final int METHODID_DELETE_STREAM_JOBS = 1;
  private static final int METHODID_CREATE_STREAM_JOB = 2;
  private static final int METHODID_UPDATE_STREAM_JOB = 3;
  private static final int METHODID_DESCRIBE_STREAM_JOB = 4;
  private static final int METHODID_SET_STREAM_JOB_CODE = 5;
  private static final int METHODID_GET_STREAM_JOB_CODE = 6;
  private static final int METHODID_STREAM_JOB_CODE_SYNTAX = 7;
  private static final int METHODID_SET_STREAM_JOB_ARGS = 8;
  private static final int METHODID_GET_STREAM_JOB_ARGS = 9;
  private static final int METHODID_LIST_BUILT_IN_CONNECTORS = 10;
  private static final int METHODID_SET_STREAM_JOB_SCHEDULE = 11;
  private static final int METHODID_GET_STREAM_JOB_SCHEDULE = 12;
  private static final int METHODID_RELEASE_STREAM_JOB = 13;
  private static final int METHODID_OFFLINE_RELEASE_STREAM_JOB = 14;
  private static final int METHODID_SUSPEND_RELEASE_STREAM_JOB = 15;
  private static final int METHODID_RESUME_RELEASE_STREAM_JOB = 16;
  private static final int METHODID_LIST_RELEASE_STREAM_JOBS = 17;
  private static final int METHODID_LIST_STREAM_JOB_VERSIONS = 18;
  private static final int METHODID_DESCRIBE_STREAM_JOB_VERSION = 19;
  private static final int METHODID_GET_STREAM_JOB_VERSION_CODE = 20;
  private static final int METHODID_GET_STREAM_JOB_VERSION_ARGS = 21;
  private static final int METHODID_GET_STREAM_JOB_VERSION_SCHEDULE = 22;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final StreamJobManageImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(StreamJobManageImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST_STREAM_JOBS:
          serviceImpl.listStreamJobs((com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobs) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobs>) responseObserver);
          break;
        case METHODID_DELETE_STREAM_JOBS:
          serviceImpl.deleteStreamJobs((com.dataomnis.gproto.requestpb.RequestPB.DeleteStreamJobs) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_CREATE_STREAM_JOB:
          serviceImpl.createStreamJob((com.dataomnis.gproto.requestpb.RequestPB.CreateStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.CreateStreamJob>) responseObserver);
          break;
        case METHODID_UPDATE_STREAM_JOB:
          serviceImpl.updateStreamJob((com.dataomnis.gproto.requestpb.RequestPB.UpdateStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_DESCRIBE_STREAM_JOB:
          serviceImpl.describeStreamJob((com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob>) responseObserver);
          break;
        case METHODID_SET_STREAM_JOB_CODE:
          serviceImpl.setStreamJobCode((com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobCode) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_GET_STREAM_JOB_CODE:
          serviceImpl.getStreamJobCode((com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode>) responseObserver);
          break;
        case METHODID_STREAM_JOB_CODE_SYNTAX:
          serviceImpl.streamJobCodeSyntax((com.dataomnis.gproto.requestpb.RequestPB.StreamJobCodeSyntax) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.StreamJobCodeSyntax>) responseObserver);
          break;
        case METHODID_SET_STREAM_JOB_ARGS:
          serviceImpl.setStreamJobArgs((com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobArgs) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_GET_STREAM_JOB_ARGS:
          serviceImpl.getStreamJobArgs((com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs>) responseObserver);
          break;
        case METHODID_LIST_BUILT_IN_CONNECTORS:
          serviceImpl.listBuiltInConnectors((com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListBuiltInConnectors>) responseObserver);
          break;
        case METHODID_SET_STREAM_JOB_SCHEDULE:
          serviceImpl.setStreamJobSchedule((com.dataomnis.gproto.requestpb.RequestPB.SetStreamJobSchedule) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_GET_STREAM_JOB_SCHEDULE:
          serviceImpl.getStreamJobSchedule((com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule>) responseObserver);
          break;
        case METHODID_RELEASE_STREAM_JOB:
          serviceImpl.releaseStreamJob((com.dataomnis.gproto.requestpb.RequestPB.ReleaseStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_OFFLINE_RELEASE_STREAM_JOB:
          serviceImpl.offlineReleaseStreamJob((com.dataomnis.gproto.requestpb.RequestPB.OfflineReleaseStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_SUSPEND_RELEASE_STREAM_JOB:
          serviceImpl.suspendReleaseStreamJob((com.dataomnis.gproto.requestpb.RequestPB.SuspendReleaseStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_RESUME_RELEASE_STREAM_JOB:
          serviceImpl.resumeReleaseStreamJob((com.dataomnis.gproto.requestpb.RequestPB.ResumeReleaseStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.modelpb.ModelPB.EmptyStruct>) responseObserver);
          break;
        case METHODID_LIST_RELEASE_STREAM_JOBS:
          serviceImpl.listReleaseStreamJobs((com.dataomnis.gproto.requestpb.RequestPB.ListReleaseStreamJobs) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListReleaseStreamJobs>) responseObserver);
          break;
        case METHODID_LIST_STREAM_JOB_VERSIONS:
          serviceImpl.listStreamJobVersions((com.dataomnis.gproto.requestpb.RequestPB.ListStreamJobVersions) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.ListStreamJobVersions>) responseObserver);
          break;
        case METHODID_DESCRIBE_STREAM_JOB_VERSION:
          serviceImpl.describeStreamJobVersion((com.dataomnis.gproto.requestpb.RequestPB.DescribeStreamJob) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.DescribeStreamJob>) responseObserver);
          break;
        case METHODID_GET_STREAM_JOB_VERSION_CODE:
          serviceImpl.getStreamJobVersionCode((com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobCode) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobCode>) responseObserver);
          break;
        case METHODID_GET_STREAM_JOB_VERSION_ARGS:
          serviceImpl.getStreamJobVersionArgs((com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobArgs) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobArgs>) responseObserver);
          break;
        case METHODID_GET_STREAM_JOB_VERSION_SCHEDULE:
          serviceImpl.getStreamJobVersionSchedule((com.dataomnis.gproto.requestpb.RequestPB.GetStreamJobSchedule) request,
              (io.grpc.stub.StreamObserver<com.dataomnis.gproto.responsepb.ResponsePB.GetStreamJobSchedule>) responseObserver);
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

  private static abstract class StreamJobManageBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    StreamJobManageBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.dataomnis.gproto.workspacepb.WorkspacePB.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("StreamJobManage");
    }
  }

  private static final class StreamJobManageFileDescriptorSupplier
      extends StreamJobManageBaseDescriptorSupplier {
    StreamJobManageFileDescriptorSupplier() {}
  }

  private static final class StreamJobManageMethodDescriptorSupplier
      extends StreamJobManageBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    StreamJobManageMethodDescriptorSupplier(String methodName) {
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
      synchronized (StreamJobManageGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new StreamJobManageFileDescriptorSupplier())
              .addMethod(getListStreamJobsMethod())
              .addMethod(getDeleteStreamJobsMethod())
              .addMethod(getCreateStreamJobMethod())
              .addMethod(getUpdateStreamJobMethod())
              .addMethod(getDescribeStreamJobMethod())
              .addMethod(getSetStreamJobCodeMethod())
              .addMethod(getGetStreamJobCodeMethod())
              .addMethod(getStreamJobCodeSyntaxMethod())
              .addMethod(getSetStreamJobArgsMethod())
              .addMethod(getGetStreamJobArgsMethod())
              .addMethod(getListBuiltInConnectorsMethod())
              .addMethod(getSetStreamJobScheduleMethod())
              .addMethod(getGetStreamJobScheduleMethod())
              .addMethod(getReleaseStreamJobMethod())
              .addMethod(getOfflineReleaseStreamJobMethod())
              .addMethod(getSuspendReleaseStreamJobMethod())
              .addMethod(getResumeReleaseStreamJobMethod())
              .addMethod(getListReleaseStreamJobsMethod())
              .addMethod(getListStreamJobVersionsMethod())
              .addMethod(getDescribeStreamJobVersionMethod())
              .addMethod(getGetStreamJobVersionCodeMethod())
              .addMethod(getGetStreamJobVersionArgsMethod())
              .addMethod(getGetStreamJobVersionScheduleMethod())
              .build();
        }
      }
    }
    return result;
  }
}