# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/service/spacemanager/monitor_manage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model import empty_pb2 as proto_dot_types_dot_model_dot_empty__pb2
from proto.types.request import monitor_manage_pb2 as proto_dot_types_dot_request_dot_monitor__manage__pb2
from proto.types.response import monitor_manage_pb2 as proto_dot_types_dot_response_dot_monitor__manage__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/service/spacemanager/monitor_manage.proto',
  package='spacemanager',
  syntax='proto3',
  serialized_options=b'\n\'com.dataomnis.gproto.service.pbsvcspaceB\022PBSvcMonitorManageP\000Z6github.com/DataWorkbench/gproto/xgo/service/pbsvcspace',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n/proto/service/spacemanager/monitor_manage.proto\x12\x0cspacemanager\x1a\x1dproto/types/model/empty.proto\x1a(proto/types/request/monitor_manage.proto\x1a)proto/types/response/monitor_manage.proto2\x9d\x04\n\rMonitorManage\x12K\n\x10ListMonitorRules\x12\x19.request.ListMonitorRules\x1a\x1a.response.ListMonitorRules\"\x00\x12G\n\x12\x44\x65leteMonitorRules\x12\x1b.request.DeleteMonitorRules\x1a\x12.model.EmptyStruct\"\x00\x12G\n\x12\x45nableMonitorRules\x12\x1b.request.EnableMonitorRules\x1a\x12.model.EmptyStruct\"\x00\x12I\n\x13\x44isableMonitorRules\x12\x1c.request.DisableMonitorRules\x1a\x12.model.EmptyStruct\"\x00\x12\x45\n\x11\x43reateMonitorRule\x12\x1a.request.CreateMonitorRule\x1a\x12.model.EmptyStruct\"\x00\x12\x45\n\x11UpdateMonitorRule\x12\x1a.request.UpdateMonitorRule\x1a\x12.model.EmptyStruct\"\x00\x12T\n\x13\x44\x65scribeMonitorRule\x12\x1c.request.DescribeMonitorRule\x1a\x1d.response.DescribeMonitorRule\"\x00\x42w\n\'com.dataomnis.gproto.service.pbsvcspaceB\x12PBSvcMonitorManageP\x00Z6github.com/DataWorkbench/gproto/xgo/service/pbsvcspaceb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_empty__pb2.DESCRIPTOR,proto_dot_types_dot_request_dot_monitor__manage__pb2.DESCRIPTOR,proto_dot_types_dot_response_dot_monitor__manage__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_MONITORMANAGE = _descriptor.ServiceDescriptor(
  name='MonitorManage',
  full_name='spacemanager.MonitorManage',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=182,
  serialized_end=723,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListMonitorRules',
    full_name='spacemanager.MonitorManage.ListMonitorRules',
    index=0,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_monitor__manage__pb2._LISTMONITORRULES,
    output_type=proto_dot_types_dot_response_dot_monitor__manage__pb2._LISTMONITORRULES,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteMonitorRules',
    full_name='spacemanager.MonitorManage.DeleteMonitorRules',
    index=1,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_monitor__manage__pb2._DELETEMONITORRULES,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='EnableMonitorRules',
    full_name='spacemanager.MonitorManage.EnableMonitorRules',
    index=2,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_monitor__manage__pb2._ENABLEMONITORRULES,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DisableMonitorRules',
    full_name='spacemanager.MonitorManage.DisableMonitorRules',
    index=3,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_monitor__manage__pb2._DISABLEMONITORRULES,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateMonitorRule',
    full_name='spacemanager.MonitorManage.CreateMonitorRule',
    index=4,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_monitor__manage__pb2._CREATEMONITORRULE,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateMonitorRule',
    full_name='spacemanager.MonitorManage.UpdateMonitorRule',
    index=5,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_monitor__manage__pb2._UPDATEMONITORRULE,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DescribeMonitorRule',
    full_name='spacemanager.MonitorManage.DescribeMonitorRule',
    index=6,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_monitor__manage__pb2._DESCRIBEMONITORRULE,
    output_type=proto_dot_types_dot_response_dot_monitor__manage__pb2._DESCRIBEMONITORRULE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MONITORMANAGE)

DESCRIPTOR.services_by_name['MonitorManage'] = _MONITORMANAGE

# @@protoc_insertion_point(module_scope)
