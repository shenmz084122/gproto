// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/service/spacemanager/sync_job_manage.proto

package com.dataomnis.gproto.service.pbsvcspace;

public final class PBSvcSyncJobManage {
  private PBSvcSyncJobManage() {}
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
      "\n0proto/service/spacemanager/sync_job_ma" +
      "nage.proto\022\014spacemanager\032\035proto/types/mo" +
      "del/empty.proto\032)proto/types/request/syn" +
      "c_job_manage.proto\032*proto/types/response" +
      "/sync_job_manage.proto2\212\014\n\rSyncJobManage" +
      "\022?\n\014ListSyncJobs\022\025.request.ListSyncJobs\032" +
      "\026.response.ListSyncJobs\"\000\022?\n\016DeleteSyncJ" +
      "obs\022\027.request.DeleteSyncJobs\032\022.model.Emp" +
      "tyStruct\"\000\022;\n\014MoveSyncJobs\022\025.request.Mov" +
      "eSyncJobs\032\022.model.EmptyStruct\"\000\022B\n\rCreat" +
      "eSyncJob\022\026.request.CreateSyncJob\032\027.respo" +
      "nse.CreateSyncJob\"\000\022=\n\rUpdateSyncJob\022\026.r" +
      "equest.UpdateSyncJob\032\022.model.EmptyStruct" +
      "\"\000\022H\n\017DescribeSyncJob\022\030.request.Describe" +
      "SyncJob\032\031.response.DescribeSyncJob\"\000\022?\n\016" +
      "SetSyncJobArgs\022\027.request.SetSyncJobArgs\032" +
      "\022.model.EmptyStruct\"\000\022E\n\016GetSyncJobArgs\022" +
      "\027.request.GetSyncJobArgs\032\030.response.GetS" +
      "yncJobArgs\"\000\022G\n\022SetSyncJobSchedule\022\033.req" +
      "uest.SetSyncJobSchedule\032\022.model.EmptyStr" +
      "uct\"\000\022Q\n\022GetSyncJobSchedule\022\033.request.Ge" +
      "tSyncJobSchedule\032\034.response.GetSyncJobSc" +
      "hedule\"\000\022?\n\016ReleaseSyncJob\022\027.request.Rel" +
      "easeSyncJob\032\022.model.EmptyStruct\"\000\022M\n\025Off" +
      "lineReleaseSyncJob\022\036.request.OfflineRele" +
      "aseSyncJob\032\022.model.EmptyStruct\"\000\022M\n\025Susp" +
      "endReleaseSyncJob\022\036.request.SuspendRelea" +
      "seSyncJob\032\022.model.EmptyStruct\"\000\022K\n\024Resum" +
      "eReleaseSyncJob\022\035.request.ResumeReleaseS" +
      "yncJob\032\022.model.EmptyStruct\"\000\022T\n\023ListRele" +
      "aseSyncJobs\022\034.request.ListReleaseSyncJob" +
      "s\032\035.response.ListReleaseSyncJobs\"\000\022T\n\023Li" +
      "stSyncJobVersions\022\034.request.ListSyncJobV" +
      "ersions\032\035.response.ListSyncJobVersions\"\000" +
      "\022O\n\026DescribeSyncJobVersion\022\030.request.Des" +
      "cribeSyncJob\032\031.response.DescribeSyncJob\"" +
      "\000\022L\n\025GetSyncJobVersionArgs\022\027.request.Get" +
      "SyncJobArgs\032\030.response.GetSyncJobArgs\"\000\022" +
      "X\n\031GetSyncJobVersionSchedule\022\033.request.G" +
      "etSyncJobSchedule\032\034.response.GetSyncJobS" +
      "chedule\"\000\022x\n\037DescribeSyncFlinkUIByInstan" +
      "ceId\022(.request.DescribeSyncFlinkUIByInst" +
      "anceId\032).response.DescribeSyncFlinkUIByI" +
      "nstanceId\"\000Bw\n\'com.dataomnis.gproto.serv" +
      "ice.pbsvcspaceB\022PBSvcSyncJobManageP\000Z6gi" +
      "thub.com/DataWorkbench/gproto/xgo/servic" +
      "e/pbsvcspaceb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor(),
          com.dataomnis.gproto.types.pbrequest.PBRequestSyncJobManage.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseSyncJobManage.getDescriptor(),
        });
    com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor();
    com.dataomnis.gproto.types.pbrequest.PBRequestSyncJobManage.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseSyncJobManage.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}