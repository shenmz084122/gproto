// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/service/enginemanager/engine_manage.proto

package com.dataomnis.gproto.service.pbsvcengine;

public final class PBSvcEngineManage {
  private PBSvcEngineManage() {}
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
      "\n/proto/service/enginemanager/engine_man" +
      "age.proto\022\014enginecenter\032\035proto/types/mod" +
      "el/empty.proto\032&proto/types/request/spac" +
      "e_manage.proto\032(proto/types/request/clus" +
      "ter_manage.proto\032(proto/types/request/ne" +
      "twork_manage.proto\032)proto/types/response" +
      "/cluster_manage.proto\032)proto/types/respo" +
      "nse/network_manage.proto2\370\010\n\006Engine\022C\n\020D" +
      "eleteWorkspaces\022\031.request.DeleteWorkspac" +
      "es\032\022.model.EmptyStruct\"\000\022X\n\032ListAvailabl" +
      "eFlinkVersions\022\022.model.EmptyStruct\032$.res" +
      "ponse.ListAvailableFlinkVersions\"\000\022`\n\027De" +
      "scribeFlinkClusterAPI\022 .request.Describe" +
      "FlinkClusterAPI\032!.response.DescribeFlink" +
      "ClusterAPI\"\000\022Q\n\022CreateFlinkCluster\022\033.req" +
      "uest.CreateFlinkCluster\032\034.response.Creat" +
      "eFlinkCluster\"\000\022N\n\021ListFlinkClusters\022\032.r" +
      "equest.ListFlinkClusters\032\033.response.List" +
      "FlinkClusters\"\000\022W\n\024DescribeFlinkCluster\022" +
      "\035.request.DescribeFlinkCluster\032\036.respons" +
      "e.DescribeFlinkCluster\"\000\022G\n\022UpdateFlinkC" +
      "luster\022\033.request.UpdateFlinkCluster\032\022.mo" +
      "del.EmptyStruct\"\000\022I\n\023DeleteFlinkClusters" +
      "\022\034.request.DeleteFlinkClusters\032\022.model.E" +
      "mptyStruct\"\000\022G\n\022StartFlinkClusters\022\033.req" +
      "uest.StartFlinkClusters\032\022.model.EmptyStr" +
      "uct\"\000\022E\n\021StopFlinkClusters\022\032.request.Sto" +
      "pFlinkClusters\032\022.model.EmptyStruct\"\000\022B\n\r" +
      "CreateNetwork\022\026.request.CreateNetwork\032\027." +
      "response.CreateNetwork\"\000\022?\n\014ListNetworks" +
      "\022\025.request.ListNetworks\032\026.response.ListN" +
      "etworks\"\000\022H\n\017DescribeNetwork\022\030.request.D" +
      "escribeNetwork\032\031.response.DescribeNetwor" +
      "k\"\000\022=\n\rUpdateNetwork\022\026.request.UpdateNet" +
      "work\032\022.model.EmptyStruct\"\000\022?\n\016DeleteNetw" +
      "orks\022\027.request.DeleteNetworks\032\022.model.Em" +
      "ptyStruct\"\000Bx\n(com.dataomnis.gproto.serv" +
      "ice.pbsvcengineB\021PBSvcEngineManageP\000Z7gi" +
      "thub.com/DataWorkbench/gproto/xgo/servic" +
      "e/pbsvcengineb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor(),
          com.dataomnis.gproto.types.pbrequest.PBRequestWorkspaceManage.getDescriptor(),
          com.dataomnis.gproto.types.pbrequest.PBRequestClusterManage.getDescriptor(),
          com.dataomnis.gproto.types.pbrequest.PBRequestNetworkManage.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseClusterManage.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseNetworkManage.getDescriptor(),
        });
    com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor();
    com.dataomnis.gproto.types.pbrequest.PBRequestWorkspaceManage.getDescriptor();
    com.dataomnis.gproto.types.pbrequest.PBRequestClusterManage.getDescriptor();
    com.dataomnis.gproto.types.pbrequest.PBRequestNetworkManage.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseClusterManage.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseNetworkManage.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
