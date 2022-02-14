# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/service/enginemanager/engine_manage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model import empty_pb2 as proto_dot_types_dot_model_dot_empty__pb2
from proto.types.request import engine_manage_pb2 as proto_dot_types_dot_request_dot_engine__manage__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/service/enginemanager/engine_manage.proto',
  package='enginecenter',
  syntax='proto3',
  serialized_options=b'\n(com.dataomnis.gproto.service.pbsvcengineB\021PBSvcEngineManageP\000Z7github.com/DataWorkbench/gproto/xgo/service/pbsvcengine',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n/proto/service/enginemanager/engine_manage.proto\x12\x0c\x65nginecenter\x1a\x1dproto/types/model/empty.proto\x1a\'proto/types/request/engine_manage.proto2\xaa\x03\n\x0c\x45ngineManage\x12Q\n\x17\x43reateFlinkClusterInK8s\x12 .request.CreateFlinkClusterInK8s\x1a\x12.model.EmptyStruct\"\x00\x12O\n\x16StartFlinkClusterInK8s\x12\x1f.request.StartFlinkClusterInK8s\x1a\x12.model.EmptyStruct\"\x00\x12M\n\x15StopFlinkClusterInK8s\x12\x1e.request.StopFlinkClusterInK8s\x1a\x12.model.EmptyStruct\"\x00\x12R\n\x18\x43reateNetworkBrokerInK8s\x12 .request.CreateFlinkClusterInK8s\x1a\x12.model.EmptyStruct\"\x00\x12S\n\x18\x44\x65leteNetworkBrokerInK8s\x12!.request.DeleteNetworkBrokerInK8s\x1a\x12.model.EmptyStruct\"\x00\x42x\n(com.dataomnis.gproto.service.pbsvcengineB\x11PBSvcEngineManageP\x00Z7github.com/DataWorkbench/gproto/xgo/service/pbsvcengineb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_empty__pb2.DESCRIPTOR,proto_dot_types_dot_request_dot_engine__manage__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_ENGINEMANAGE = _descriptor.ServiceDescriptor(
  name='EngineManage',
  full_name='enginecenter.EngineManage',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=138,
  serialized_end=564,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateFlinkClusterInK8s',
    full_name='enginecenter.EngineManage.CreateFlinkClusterInK8s',
    index=0,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_engine__manage__pb2._CREATEFLINKCLUSTERINK8S,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='StartFlinkClusterInK8s',
    full_name='enginecenter.EngineManage.StartFlinkClusterInK8s',
    index=1,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_engine__manage__pb2._STARTFLINKCLUSTERINK8S,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='StopFlinkClusterInK8s',
    full_name='enginecenter.EngineManage.StopFlinkClusterInK8s',
    index=2,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_engine__manage__pb2._STOPFLINKCLUSTERINK8S,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateNetworkBrokerInK8s',
    full_name='enginecenter.EngineManage.CreateNetworkBrokerInK8s',
    index=3,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_engine__manage__pb2._CREATEFLINKCLUSTERINK8S,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteNetworkBrokerInK8s',
    full_name='enginecenter.EngineManage.DeleteNetworkBrokerInK8s',
    index=4,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_engine__manage__pb2._DELETENETWORKBROKERINK8S,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ENGINEMANAGE)

DESCRIPTOR.services_by_name['EngineManage'] = _ENGINEMANAGE

# @@protoc_insertion_point(module_scope)