// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/service/spacemanager/cluster_manage.proto

package com.dataomnis.gproto.service.pbsvcspace;

public final class PBSvcClusterManage {
  private PBSvcClusterManage() {}
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
      "\n/proto/service/spacemanager/cluster_man" +
      "age.proto\022\014spacemanager\032\035proto/types/mod" +
      "el/empty.proto\032(proto/types/request/clus" +
      "ter_manage.proto\032)proto/types/response/c" +
      "luster_manage.proto2\353\005\n\rClusterManage\022X\n" +
      "\032ListAvailableFlinkVersions\022\022.model.Empt" +
      "yStruct\032$.response.ListAvailableFlinkVer" +
      "sions\"\000\022`\n\027DescribeFlinkClusterAPI\022 .req" +
      "uest.DescribeFlinkClusterAPI\032!.response." +
      "DescribeFlinkClusterAPI\"\000\022N\n\021ListFlinkCl" +
      "usters\022\032.request.ListFlinkClusters\032\033.res" +
      "ponse.ListFlinkClusters\"\000\022Q\n\022CreateFlink" +
      "Cluster\022\033.request.CreateFlinkCluster\032\034.r" +
      "esponse.CreateFlinkCluster\"\000\022W\n\024Describe" +
      "FlinkCluster\022\035.request.DescribeFlinkClus" +
      "ter\032\036.response.DescribeFlinkCluster\"\000\022G\n" +
      "\022UpdateFlinkCluster\022\033.request.UpdateFlin" +
      "kCluster\032\022.model.EmptyStruct\"\000\022I\n\023Delete" +
      "FlinkClusters\022\034.request.DeleteFlinkClust" +
      "ers\032\022.model.EmptyStruct\"\000\022G\n\022StartFlinkC" +
      "lusters\022\033.request.StartFlinkClusters\032\022.m" +
      "odel.EmptyStruct\"\000\022E\n\021StopFlinkClusters\022" +
      "\032.request.StopFlinkClusters\032\022.model.Empt" +
      "yStruct\"\000Bw\n\'com.dataomnis.gproto.servic" +
      "e.pbsvcspaceB\022PBSvcClusterManageP\000Z6gith" +
      "ub.com/DataWorkbench/gproto/xgo/service/" +
      "pbsvcspaceb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor(),
          com.dataomnis.gproto.types.pbrequest.PBRequestClusterManage.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseClusterManage.getDescriptor(),
        });
    com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor();
    com.dataomnis.gproto.types.pbrequest.PBRequestClusterManage.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseClusterManage.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
