// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/service/spacemanager/flink_manage.proto

package com.dataomnis.gproto.service.pbsvcspace;

public final class PBSvcFlinkManage {
  private PBSvcFlinkManage() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n-proto/service/spacemanager/flink_manag" +
      "e.proto\022\014spacemanager\032\035proto/types/model" +
      "/empty.proto\032&proto/types/request/flink_" +
      "manage.proto\032\'proto/types/response/flink" +
      "_manage.proto2\351\005\n\013FlinkManage\022X\n\032ListAva" +
      "ilableFlinkVersions\022\022.model.EmptyStruct\032" +
      "$.response.ListAvailableFlinkVersions\"\000\022" +
      "`\n\027DescribeFlinkClusterAPI\022 .request.Des" +
      "cribeFlinkClusterAPI\032!.response.Describe" +
      "FlinkClusterAPI\"\000\022N\n\021ListFlinkClusters\022\032" +
      ".request.ListFlinkClusters\032\033.response.Li" +
      "stFlinkClusters\"\000\022Q\n\022CreateFlinkCluster\022" +
      "\033.request.CreateFlinkCluster\032\034.response." +
      "CreateFlinkCluster\"\000\022W\n\024DescribeFlinkClu" +
      "ster\022\035.request.DescribeFlinkCluster\032\036.re" +
      "sponse.DescribeFlinkCluster\"\000\022G\n\022UpdateF" +
      "linkCluster\022\033.request.UpdateFlinkCluster" +
      "\032\022.model.EmptyStruct\"\000\022I\n\023DeleteFlinkClu" +
      "sters\022\034.request.DeleteFlinkClusters\032\022.mo" +
      "del.EmptyStruct\"\000\022G\n\022StartFlinkClusters\022" +
      "\033.request.StartFlinkClusters\032\022.model.Emp" +
      "tyStruct\"\000\022E\n\021StopFlinkClusters\022\032.reques" +
      "t.StopFlinkClusters\032\022.model.EmptyStruct\"" +
      "\000Bu\n\'com.dataomnis.gproto.service.pbsvcs" +
      "paceB\020PBSvcFlinkManageP\000Z6github.com/Dat" +
      "aWorkbench/gproto/xgo/service/pbsvcspace" +
      "b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor(),
          com.dataomnis.gproto.types.pbrequest.PBRequestFlinkManage.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseFlinkManage.getDescriptor(),
        });
    com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor();
    com.dataomnis.gproto.types.pbrequest.PBRequestFlinkManage.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseFlinkManage.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
