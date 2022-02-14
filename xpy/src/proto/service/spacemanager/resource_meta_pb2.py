# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/service/spacemanager/resource_meta.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model import empty_pb2 as proto_dot_types_dot_model_dot_empty__pb2
from proto.types.request import resource_meta_pb2 as proto_dot_types_dot_request_dot_resource__meta__pb2
from proto.types.response import resource_meta_pb2 as proto_dot_types_dot_response_dot_resource__meta__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/service/spacemanager/resource_meta.proto',
  package='spacemanager',
  syntax='proto3',
  serialized_options=b'\n\'com.dataomnis.gproto.service.pbsvcspaceB\021PBSvcResourceMetaP\000Z6github.com/DataWorkbench/gproto/xgo/service/pbsvcspace',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n.proto/service/spacemanager/resource_meta.proto\x12\x0cspacemanager\x1a\x1dproto/types/model/empty.proto\x1a\'proto/types/request/resource_meta.proto\x1a(proto/types/response/resource_meta.proto2\xcf\x04\n\x0cResourceMeta\x12N\n\x11\x43reateFilePrepare\x12\x1a.request.CreateFilePrepare\x1a\x1b.response.CreateFilePrepare\"\x00\x12?\n\x0e\x43reateFileMeta\x12\x17.request.CreateFileMeta\x1a\x12.model.EmptyStruct\"\x00\x12T\n\x13ReCreateFilePrepare\x12\x1c.request.ReCreateFilePrepare\x1a\x1d.response.ReCreateFilePrepare\"\x00\x12\x43\n\x10ReCreateFileMeta\x12\x19.request.ReCreateFileMeta\x1a\x12.model.EmptyStruct\"\x00\x12K\n\x10\x44\x65scribeFileMeta\x12\x19.request.DescribeFileMeta\x1a\x1a.response.DescribeFileMeta\"\x00\x12?\n\x0eUpdateFileMeta\x12\x17.request.UpdateFileMeta\x1a\x12.model.EmptyStruct\"\x00\x12\x42\n\rListFileMetas\x12\x16.request.ListFileMetas\x1a\x17.response.ListFileMetas\"\x00\x12\x41\n\x0f\x44\x65leteFileMetas\x12\x18.request.DeleteFileMetas\x1a\x12.model.EmptyStruct\"\x00\x42v\n\'com.dataomnis.gproto.service.pbsvcspaceB\x11PBSvcResourceMetaP\x00Z6github.com/DataWorkbench/gproto/xgo/service/pbsvcspaceb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_empty__pb2.DESCRIPTOR,proto_dot_types_dot_request_dot_resource__meta__pb2.DESCRIPTOR,proto_dot_types_dot_response_dot_resource__meta__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_RESOURCEMETA = _descriptor.ServiceDescriptor(
  name='ResourceMeta',
  full_name='spacemanager.ResourceMeta',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=179,
  serialized_end=770,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateFilePrepare',
    full_name='spacemanager.ResourceMeta.CreateFilePrepare',
    index=0,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._CREATEFILEPREPARE,
    output_type=proto_dot_types_dot_response_dot_resource__meta__pb2._CREATEFILEPREPARE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateFileMeta',
    full_name='spacemanager.ResourceMeta.CreateFileMeta',
    index=1,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._CREATEFILEMETA,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ReCreateFilePrepare',
    full_name='spacemanager.ResourceMeta.ReCreateFilePrepare',
    index=2,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._RECREATEFILEPREPARE,
    output_type=proto_dot_types_dot_response_dot_resource__meta__pb2._RECREATEFILEPREPARE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ReCreateFileMeta',
    full_name='spacemanager.ResourceMeta.ReCreateFileMeta',
    index=3,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._RECREATEFILEMETA,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DescribeFileMeta',
    full_name='spacemanager.ResourceMeta.DescribeFileMeta',
    index=4,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._DESCRIBEFILEMETA,
    output_type=proto_dot_types_dot_response_dot_resource__meta__pb2._DESCRIBEFILEMETA,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateFileMeta',
    full_name='spacemanager.ResourceMeta.UpdateFileMeta',
    index=5,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._UPDATEFILEMETA,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ListFileMetas',
    full_name='spacemanager.ResourceMeta.ListFileMetas',
    index=6,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._LISTFILEMETAS,
    output_type=proto_dot_types_dot_response_dot_resource__meta__pb2._LISTFILEMETAS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteFileMetas',
    full_name='spacemanager.ResourceMeta.DeleteFileMetas',
    index=7,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_resource__meta__pb2._DELETEFILEMETAS,
    output_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_RESOURCEMETA)

DESCRIPTOR.services_by_name['ResourceMeta'] = _RESOURCEMETA

# @@protoc_insertion_point(module_scope)
