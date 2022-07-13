# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/service/spacemanager/stream_instance_manage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model import empty_pb2 as proto_dot_types_dot_model_dot_empty__pb2
from proto.types.request import stream_instance_manage_pb2 as proto_dot_types_dot_request_dot_stream__instance__manage__pb2
from proto.types.response import stream_instance_manage_pb2 as proto_dot_types_dot_response_dot_stream__instance__manage__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/service/spacemanager/stream_instance_manage.proto',
  package='scheduler',
  syntax='proto3',
  serialized_options=b'\n\'com.dataomnis.gproto.service.pbsvcspaceB\031PBSvcStreamInstanceManageP\000Z6github.com/DataWorkbench/gproto/xgo/service/pbsvcspace',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n7proto/service/spacemanager/stream_instance_manage.proto\x12\tscheduler\x1a\x1dproto/types/model/empty.proto\x1a\x30proto/types/request/stream_instance_manage.proto\x1a\x31proto/types/response/stream_instance_manage.proto2\xf0\x04\n\x14StreamInstanceManage\x12T\n\x13ListStreamInstances\x12\x1c.request.ListStreamInstances\x1a\x1d.response.ListStreamInstances\"\x00\x12S\n\x18TerminateStreamInstances\x12!.request.TerminateStreamInstances\x1a\x12.model.EmptyStruct\"\x00\x12O\n\x16SuspendStreamInstances\x12\x1f.request.SuspendStreamInstances\x1a\x12.model.EmptyStruct\"\x00\x12M\n\x15ResumeStreamInstances\x12\x1e.request.ResumeStreamInstances\x1a\x12.model.EmptyStruct\"\x00\x12]\n\x16\x44\x65scribeStreamInstance\x12\x1f.request.DescribeStreamInstance\x1a .response.DescribeStreamInstance\"\x00\x12W\n\x1a\x43reateStreamInstanceWithId\x12#.request.CreateStreamInstanceWithId\x1a\x12.model.EmptyStruct\"\x00\x12U\n\x19UpdateStreamInstanceState\x12\".request.UpdateStreamInstanceState\x1a\x12.model.EmptyStruct\"\x00\x42~\n\'com.dataomnis.gproto.service.pbsvcspaceB\x19PBSvcStreamInstanceManageP\x00Z6github.com/DataWorkbench/gproto/xgo/service/pbsvcspaceb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_empty__pb2.DESCRIPTOR,proto_dot_types_dot_request_dot_stream__instance__manage__pb2.DESCRIPTOR,proto_dot_types_dot_response_dot_stream__instance__manage__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_STREAMINSTANCEMANAGE = _descriptor.ServiceDescriptor(
  name='StreamInstanceManage',
  full_name='scheduler.StreamInstanceManage',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=203,
  serialized_end=827,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListStreamInstances',
    full_name='scheduler.StreamInstanceManage.ListStreamInstances',
    index=0,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_stream__instance__manage__pb2._LISTSTREAMINSTANCES,
    output_type=proto_dot_types_dot_response_dot_stream__instance__manage__pb2._LISTSTREAMINSTANCES,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='TerminateStreamInstances',
    full_name='scheduler.StreamInstanceManage.TerminateStreamInstances',
    index=1,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_stream__instance__manage__pb2._TERMINATESTREAMINSTANCES,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SuspendStreamInstances',
    full_name='scheduler.StreamInstanceManage.SuspendStreamInstances',
    index=2,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_stream__instance__manage__pb2._SUSPENDSTREAMINSTANCES,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ResumeStreamInstances',
    full_name='scheduler.StreamInstanceManage.ResumeStreamInstances',
    index=3,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_stream__instance__manage__pb2._RESUMESTREAMINSTANCES,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DescribeStreamInstance',
    full_name='scheduler.StreamInstanceManage.DescribeStreamInstance',
    index=4,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_stream__instance__manage__pb2._DESCRIBESTREAMINSTANCE,
    output_type=proto_dot_types_dot_response_dot_stream__instance__manage__pb2._DESCRIBESTREAMINSTANCE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateStreamInstanceWithId',
    full_name='scheduler.StreamInstanceManage.CreateStreamInstanceWithId',
    index=5,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_stream__instance__manage__pb2._CREATESTREAMINSTANCEWITHID,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateStreamInstanceState',
    full_name='scheduler.StreamInstanceManage.UpdateStreamInstanceState',
    index=6,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_stream__instance__manage__pb2._UPDATESTREAMINSTANCESTATE,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_STREAMINSTANCEMANAGE)

DESCRIPTOR.services_by_name['StreamInstanceManage'] = _STREAMINSTANCEMANAGE

# @@protoc_insertion_point(module_scope)