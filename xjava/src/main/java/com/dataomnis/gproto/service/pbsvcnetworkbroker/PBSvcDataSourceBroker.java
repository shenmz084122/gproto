// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/service/networkbroker/dataousrce_broker.proto

package com.dataomnis.gproto.service.pbsvcnetworkbroker;

public final class PBSvcDataSourceBroker {
  private PBSvcDataSourceBroker() {}
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
      "\n3proto/service/networkbroker/dataousrce" +
      "_broker.proto\022\rnetworkbroker\032+proto/type" +
      "s/request/dataousrce_broker.proto\032,proto" +
      "/types/response/dataousrce_broker.proto\032" +
      ",proto/types/response/datasource_manage." +
      "proto2\211\003\n\020DataSourceBroker\022{\n PingDataSo" +
      "urceConnectionByBroker\022).request.PingDat" +
      "aSourceConnectionByBroker\032*.response.Pin" +
      "gDataSourceConnectionByBroker\"\000\022s\n Descr" +
      "ibeDataSourceTablesByBroker\022).request.De" +
      "scribeDataSourceTablesByBroker\032\".respons" +
      "e.DescribeDataSourceTables\"\000\022\202\001\n%Describ" +
      "eDataSourceTableSchemaByBroker\022..request" +
      ".DescribeDataSourceTableSchemaByBroker\032\'" +
      ".response.DescribeDataSourceTableSchema\"" +
      "\000B\212\001\n/com.dataomnis.gproto.service.pbsvc" +
      "networkbrokerB\025PBSvcDataSourceBrokerP\000Z>" +
      "github.com/DataWorkbench/gproto/xgo/serv" +
      "ice/pbsvcnetworkbrokerb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.dataomnis.gproto.types.pbrequest.PBRequestDataSourceBroker.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseDataSourceBroker.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseDataSourceManage.getDescriptor(),
        });
    com.dataomnis.gproto.types.pbrequest.PBRequestDataSourceBroker.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseDataSourceBroker.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseDataSourceManage.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
