// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/service/spacemanager/stream_job_manage.proto

package com.dataomnis.gproto.service.pbsvcspace;

public final class PBSvcStreamJobManage {
  private PBSvcStreamJobManage() {}
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
      "\n2proto/service/spacemanager/stream_job_" +
      "manage.proto\022\014spacemanager\032\035proto/types/" +
      "model/empty.proto\032+proto/types/request/s" +
      "tream_job_manage.proto\032,proto/types/resp" +
      "onse/stream_job_manage.proto2\363\017\n\017StreamJ" +
      "obManage\022E\n\016ListStreamJobs\022\027.request.Lis" +
      "tStreamJobs\032\030.response.ListStreamJobs\"\000\022" +
      "C\n\020DeleteStreamJobs\022\031.request.DeleteStre" +
      "amJobs\032\022.model.EmptyStruct\"\000\022?\n\016MoveStre" +
      "amJobs\022\027.request.MoveStreamJobs\032\022.model." +
      "EmptyStruct\"\000\022H\n\017CreateStreamJob\022\030.reque" +
      "st.CreateStreamJob\032\031.response.CreateStre" +
      "amJob\"\000\022A\n\017UpdateStreamJob\022\030.request.Upd" +
      "ateStreamJob\032\022.model.EmptyStruct\"\000\022N\n\021De" +
      "scribeStreamJob\022\032.request.DescribeStream" +
      "Job\032\033.response.DescribeStreamJob\"\000\022C\n\020Se" +
      "tStreamJobCode\022\031.request.SetStreamJobCod" +
      "e\032\022.model.EmptyStruct\"\000\022K\n\020GetStreamJobC" +
      "ode\022\031.request.GetStreamJobCode\032\032.respons" +
      "e.GetStreamJobCode\"\000\022C\n\020SetStreamJobArgs" +
      "\022\031.request.SetStreamJobArgs\032\022.model.Empt" +
      "yStruct\"\000\022K\n\020GetStreamJobArgs\022\031.request." +
      "GetStreamJobArgs\032\032.response.GetStreamJob" +
      "Args\"\000\022N\n\025ListBuiltInConnectors\022\022.model." +
      "EmptyStruct\032\037.response.ListBuiltInConnec" +
      "tors\"\000\022K\n\024SetStreamJobSchedule\022\035.request" +
      ".SetStreamJobSchedule\032\022.model.EmptyStruc" +
      "t\"\000\022W\n\024GetStreamJobSchedule\022\035.request.Ge" +
      "tStreamJobSchedule\032\036.response.GetStreamJ" +
      "obSchedule\"\000\022Z\n\025ListReleaseStreamJobs\022\036." +
      "request.ListReleaseStreamJobs\032\037.response" +
      ".ListReleaseStreamJobs\"\000\022C\n\020ReleaseStrea" +
      "mJob\022\031.request.ReleaseStreamJob\032\022.model." +
      "EmptyStruct\"\000\022Q\n\027OfflineReleaseStreamJob" +
      "\022 .request.OfflineReleaseStreamJob\032\022.mod" +
      "el.EmptyStruct\"\000\022Q\n\027SuspendReleaseStream" +
      "Job\022 .request.SuspendReleaseStreamJob\032\022." +
      "model.EmptyStruct\"\000\022O\n\026ResumeReleaseStre" +
      "amJob\022\037.request.ResumeReleaseStreamJob\032\022" +
      ".model.EmptyStruct\"\000\022[\n\034UpdateReleaseStr" +
      "eamJobStatus\022%.request.UpdateReleaseStre" +
      "amJobStatus\032\022.model.EmptyStruct\"\000\022Z\n\025Lis" +
      "tStreamJobVersions\022\036.request.ListStreamJ" +
      "obVersions\032\037.response.ListStreamJobVersi" +
      "ons\"\000\022U\n\030DescribeStreamJobVersion\022\032.requ" +
      "est.DescribeStreamJob\032\033.response.Describ" +
      "eStreamJob\"\000\022R\n\027GetStreamJobVersionCode\022" +
      "\031.request.GetStreamJobCode\032\032.response.Ge" +
      "tStreamJobCode\"\000\022R\n\027GetStreamJobVersionA" +
      "rgs\022\031.request.GetStreamJobArgs\032\032.respons" +
      "e.GetStreamJobArgs\"\000\022^\n\033GetStreamJobVers" +
      "ionSchedule\022\035.request.GetStreamJobSchedu" +
      "le\032\036.response.GetStreamJobSchedule\"\000\022l\n\033" +
      "DescribeFlinkUIByInstanceId\022$.request.De" +
      "scribeFlinkUIByInstanceId\032%.response.Des" +
      "cribeFlinkUIByInstanceId\"\000By\n\'com.dataom" +
      "nis.gproto.service.pbsvcspaceB\024PBSvcStre" +
      "amJobManageP\000Z6github.com/DataWorkbench/" +
      "gproto/xgo/service/pbsvcspaceb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor(),
          com.dataomnis.gproto.types.pbrequest.PBRequestStreamJobManage.getDescriptor(),
          com.dataomnis.gproto.types.pbresponse.PBResponseStreamJobManage.getDescriptor(),
        });
    com.dataomnis.gproto.types.pbmodel.PBModelEmpty.getDescriptor();
    com.dataomnis.gproto.types.pbrequest.PBRequestStreamJobManage.getDescriptor();
    com.dataomnis.gproto.types.pbresponse.PBResponseStreamJobManage.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
