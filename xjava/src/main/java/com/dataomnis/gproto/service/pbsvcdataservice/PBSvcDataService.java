// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/service/dataservice/dataservice.proto

package com.dataomnis.gproto.service.pbsvcdataservice;

public final class PBSvcDataService {
  private PBSvcDataService() {}
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
      "\n+proto/service/dataservice/dataservice." +
      "proto\022\013dataservice\032,proto/types/request/" +
      "dataservice_manage.proto\032-proto/types/re" +
      "sponse/dataservice_manage.proto\032\035proto/t" +
      "ypes/model/empty.proto\032#proto/types/mode" +
      "l/dataservice.proto2\344\r\n\013DataService\022c\n\030C" +
      "reateDataServiceCluster\022!.request.Create" +
      "DataServiceCluster\032\".response.CreateData" +
      "ServiceCluster\"\000\022`\n\027ListDataServiceClust" +
      "ers\022 .request.ListDataServiceClusters\032!." +
      "response.ListDataServiceClusters\"\000\022S\n\030Up" +
      "dateDataServiceCluster\022!.request.UpdateD" +
      "ataServiceCluster\032\022.model.EmptyStruct\"\000\022" +
      "^\n\032DescribeDataServiceCluster\022#.request." +
      "DescribeDataServiceCluster\032\031.model.DataS" +
      "erviceCluster\"\000\022U\n\031DeleteDataServiceClus" +
      "ters\022\".request.DeleteDataServiceClusters" +
      "\032\022.model.EmptyStruct\"\000\022S\n\030StartDataServi" +
      "ceClusters\022!.request.StartDataServiceClu" +
      "sters\032\022.model.EmptyStruct\"\000\022Q\n\027StopDataS" +
      "erviceClusters\022 .request.StopDataService" +
      "Clusters\032\022.model.EmptyStruct\"\000\022E\n\016Create" +
      "ApiGroup\022\027.request.CreateApiGroup\032\030.resp" +
      "onse.CreateApiGroup\"\000\022B\n\rListApiGroups\022\026" +
      ".request.ListApiGroups\032\027.response.ListAp" +
      "iGroups\"\000\022A\n\017DeleteApiGroups\022\030.request.D" +
      "eleteApiGroups\032\022.model.EmptyStruct\"\000\022E\n\016" +
      "ListApiConfigs\022\027.request.ListApiConfigs\032" +
      "\030.response.ListApiConfigs\"\000\022N\n\021DescribeA" +
      "piConfig\022\032.request.DescribeApiConfig\032\033.r" +
      "esponse.DescribeApiConfig\"\000\022H\n\017CreateApi" +
      "Config\022\030.request.CreateApiConfig\032\031.respo" +
      "nse.CreateApiConfig\"\000\022A\n\017UpdateApiConfig" +
      "\022\030.request.UpdateApiConfig\032\022.model.Empty" +
      "Struct\"\000\022C\n\020DeleteApiConfigs\022\031.request.D" +
      "eleteApiConfigs\032\022.model.EmptyStruct\"\000\022Q\n" +
      "\022TestDataServiceApi\022\033.request.TestDataSe" +
      "rviceApi\032\034.response.TestDataServiceApi\"\000" +
      "\022M\n\025PublishDataServiceApi\022\036.request.Publ" +
      "ishDataServiceApi\032\022.model.EmptyStruct\"\000\022" +
      "O\n\026AbolishDataServiceApis\022\037.request.Abol" +
      "ishDataServiceApis\032\022.model.EmptyStruct\"\000" +
      "\022i\n\032ListDataServiceApiVersions\022#.request" +
      ".ListDataServiceApiVersions\032$.response.L" +
      "istDataServiceApiVersions\"\000\022r\n\035DescribeD" +
      "ataServiceApiVersion\022&.request.DescribeD" +
      "ataServiceApiVersion\032\'.response.Describe" +
      "DataServiceApiVersion\"\000\022Q\n\027RepublishData" +
      "ServiceApi\022 .request.RepublishDataServic" +
      "eApi\032\022.model.EmptyStruct\"\000B\201\001\n-com.datao" +
      "mnis.gproto.service.pbsvcdataserviceB\020PB" +
      "SvcDataServiceP\000Z<github.com/DataWorkben" +
      "ch/gproto/xgo/service/pbsvcdataserviceb\006" +
      "proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.dataomnis.gproto.types.pbrequest.PBRequestDataServiceManage.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseDataServiceManage.getDescriptor(),
          com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor(),
          com.dataomnis.gproto.types.pbmodel.PBModelDataService.getDescriptor(),
        });
    com.dataomnis.gproto.types.pbrequest.PBRequestDataServiceManage.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseDataServiceManage.getDescriptor();
    com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor();
    com.dataomnis.gproto.types.pbmodel.PBModelDataService.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
