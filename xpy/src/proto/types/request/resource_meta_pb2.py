# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/request/resource_meta.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from github.com.yu31.protoc_plugin.proto import defaults_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_defaults__pb2
from proto.types.model import resource_pb2 as proto_dot_types_dot_model_dot_resource__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/request/resource_meta.proto',
  package='request',
  syntax='proto3',
  serialized_options=b'\n$com.dataomnis.gproto.types.pbrequestB\025PBRequestResourceMetaP\000Z3github.com/DataWorkbench/gproto/xgo/types/pbrequest',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\'proto/types/request/resource_meta.proto\x12\x07request\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a\x32github.com/yu31/protoc-plugin/proto/defaults.proto\x1a proto/types/model/resource.proto\"\xae\x01\n\x11\x43reateFilePrepare\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12\x32\n\x03pid\x18\x02 \x01(\tB%\xe2\xdf\x1f\x0e\n\x0c\n\x03pid\x12\x05\xc2\x01\x02\"\x00\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04res-\x12\x1e\n\x04name\x18\x03 \x01(\tB\x10\xe2\xdf\x1f\x0c\x12\n\xc2\x01\x07\xc0\x01\x02\xc8\x01\x80\x01\x12\x1e\n\x04size\x18\x04 \x01(\x03\x42\x10\xe2\xdf\x1f\x0c\x12\n\xb2\x01\x07\x30\x00\x38\x80\x80\x80\x32\"\xf3\x02\n\x0e\x43reateFileMeta\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12(\n\x0bresource_id\x18\x02 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04res-\x12\x32\n\x03pid\x18\x03 \x01(\tB%\xe2\xdf\x1f\x0e\n\x0c\n\x03pid\x12\x05\xc2\x01\x02\"\x00\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04res-\x12\x1e\n\x04name\x18\x04 \x01(\tB\x10\xe2\xdf\x1f\x0c\x12\n\xc2\x01\x07\xc0\x01\x02\xc8\x01\x80\x01\x12\x1b\n\x04\x64\x65sc\x18\x05 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\xc8\x01\x80\x08\x12\x0c\n\x04size\x18\x06 \x01(\x03\x12\x1a\n\x04\x65tag\x18\x07 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01 \x12\x31\n\x04type\x18\x08 \x01(\x0e\x32\x14.model.Resource.TypeB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12\x1d\n\x07version\x18\t \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x10\x12#\n\ncreated_by\x18\n \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x80\x02\x00\x88\x02\x41\"\x86\x01\n\x13ReCreateFilePrepare\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12(\n\x0bresource_id\x18\x02 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04res-\x12\x1e\n\x04size\x18\x03 \x01(\x03\x42\x10\xe2\xdf\x1f\x0c\x12\n\xb2\x01\x07\x30\x00\x38\x80\x80\x80\x32\"\xac\x01\n\x10ReCreateFileMeta\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12(\n\x0bresource_id\x18\x02 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04res-\x12\x0c\n\x04size\x18\x03 \x01(\x03\x12\x1a\n\x04\x65tag\x18\x04 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01 \x12\x1d\n\x07version\x18\x05 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x10\"<\n\x10\x44\x65scribeFileMeta\x12(\n\x0bresource_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04res-\"\xbe\x01\n\x0eUpdateFileMeta\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12!\n\x0bresource_id\x18\x02 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x12\x1b\n\x04name\x18\x03 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\x98\x02\xf4\x03\x12\x12\n\x04\x64\x65sc\x18\x04 \x01(\tB\x04\xe2\xdf\x1f\x00\x12\x31\n\x04type\x18\x05 \x01(\x0e\x32\x14.model.Resource.TypeB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\"\xc3\x02\n\rListFileMetas\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12&\n\x05limit\x18\x02 \x01(\x05\x42\x17\xa2\xa1\x1f\x06\xaa\x06\x03\x31\x30\x30\xe2\xdf\x1f\t\x12\x07\xb2\x01\x04\x30\x00\x38\x64\x12\x1b\n\x06offset\x18\x03 \x01(\x05\x42\x0b\xe2\xdf\x1f\x07\x12\x05\xb2\x01\x02@\x00\x12/\n\x04type\x18\x04 \x01(\x0e\x32\x14.model.Resource.TypeB\x0b\xe2\xdf\x1f\x07\x12\x05\xda\x01\x02X\x01\x12\x12\n\x04name\x18\x05 \x01(\tB\x04\xe2\xdf\x1f\x00\x12\x1d\n\x06search\x18\x06 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\x98\x02\xf4\x03\x12>\n\x07sort_by\x18\x07 \x01(\tB-\xe2\xdf\x1f)\x12\'\xc2\x01$J\x00J\x02idJ\x07\x63reatedJ\x07updatedJ\x04nameJ\x04size\x12\x15\n\x07reverse\x18\x08 \x01(\x08\x42\x04\xe2\xdf\x1f\x00\x12\x0b\n\x03pid\x18\t \x01(\t\"]\n\x0f\x44\x65leteFileMetas\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12#\n\x0cresource_ids\x18\x02 \x03(\tB\r\xe2\xdf\x1f\t\x12\x07\xea\x01\x04\x30\x00\x38\x64\x42t\n$com.dataomnis.gproto.types.pbrequestB\x15PBRequestResourceMetaP\x00Z3github.com/DataWorkbench/gproto/xgo/types/pbrequestb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_defaults__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_resource__pb2.DESCRIPTOR,])




_CREATEFILEPREPARE = _descriptor.Descriptor(
  name='CreateFilePrepare',
  full_name='request.CreateFilePrepare',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.CreateFilePrepare.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pid', full_name='request.CreateFilePrepare.pid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\016\n\014\n\003pid\022\005\302\001\002\"\000\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004res-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='request.CreateFilePrepare.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\014\022\n\302\001\007\300\001\002\310\001\200\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='size', full_name='request.CreateFilePrepare.size', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\014\022\n\262\001\0070\0008\200\200\2002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=192,
  serialized_end=366,
)


_CREATEFILEMETA = _descriptor.Descriptor(
  name='CreateFileMeta',
  full_name='request.CreateFileMeta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.CreateFileMeta.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_id', full_name='request.CreateFileMeta.resource_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004res-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pid', full_name='request.CreateFileMeta.pid', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\016\n\014\n\003pid\022\005\302\001\002\"\000\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004res-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='request.CreateFileMeta.name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\014\022\n\302\001\007\300\001\002\310\001\200\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='desc', full_name='request.CreateFileMeta.desc', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\310\001\200\010', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='size', full_name='request.CreateFileMeta.size', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='etag', full_name='request.CreateFileMeta.etag', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001 ', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='request.CreateFileMeta.type', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='request.CreateFileMeta.version', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\020', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_by', full_name='request.CreateFileMeta.created_by', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\200\002\000\210\002A', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=369,
  serialized_end=740,
)


_RECREATEFILEPREPARE = _descriptor.Descriptor(
  name='ReCreateFilePrepare',
  full_name='request.ReCreateFilePrepare',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.ReCreateFilePrepare.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_id', full_name='request.ReCreateFilePrepare.resource_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004res-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='size', full_name='request.ReCreateFilePrepare.size', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\014\022\n\262\001\0070\0008\200\200\2002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=743,
  serialized_end=877,
)


_RECREATEFILEMETA = _descriptor.Descriptor(
  name='ReCreateFileMeta',
  full_name='request.ReCreateFileMeta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.ReCreateFileMeta.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_id', full_name='request.ReCreateFileMeta.resource_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004res-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='size', full_name='request.ReCreateFileMeta.size', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='etag', full_name='request.ReCreateFileMeta.etag', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001 ', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='request.ReCreateFileMeta.version', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\020', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=880,
  serialized_end=1052,
)


_DESCRIBEFILEMETA = _descriptor.Descriptor(
  name='DescribeFileMeta',
  full_name='request.DescribeFileMeta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='resource_id', full_name='request.DescribeFileMeta.resource_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004res-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1054,
  serialized_end=1114,
)


_UPDATEFILEMETA = _descriptor.Descriptor(
  name='UpdateFileMeta',
  full_name='request.UpdateFileMeta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.UpdateFileMeta.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_id', full_name='request.UpdateFileMeta.resource_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='request.UpdateFileMeta.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\230\002\364\003', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='desc', full_name='request.UpdateFileMeta.desc', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='request.UpdateFileMeta.type', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1117,
  serialized_end=1307,
)


_LISTFILEMETAS = _descriptor.Descriptor(
  name='ListFileMetas',
  full_name='request.ListFileMetas',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.ListFileMetas.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='limit', full_name='request.ListFileMetas.limit', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\242\241\037\006\252\006\003100\342\337\037\t\022\007\262\001\0040\0008d', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset', full_name='request.ListFileMetas.offset', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\262\001\002@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='request.ListFileMetas.type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\332\001\002X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='request.ListFileMetas.name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='search', full_name='request.ListFileMetas.search', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\230\002\364\003', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sort_by', full_name='request.ListFileMetas.sort_by', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037)\022\'\302\001$J\000J\002idJ\007createdJ\007updatedJ\004nameJ\004size', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='reverse', full_name='request.ListFileMetas.reverse', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pid', full_name='request.ListFileMetas.pid', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1310,
  serialized_end=1633,
)


_DELETEFILEMETAS = _descriptor.Descriptor(
  name='DeleteFileMetas',
  full_name='request.DeleteFileMetas',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.DeleteFileMetas.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_ids', full_name='request.DeleteFileMetas.resource_ids', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\352\001\0040\0008d', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1635,
  serialized_end=1728,
)

_CREATEFILEMETA.fields_by_name['type'].enum_type = proto_dot_types_dot_model_dot_resource__pb2._RESOURCE_TYPE
_UPDATEFILEMETA.fields_by_name['type'].enum_type = proto_dot_types_dot_model_dot_resource__pb2._RESOURCE_TYPE
_LISTFILEMETAS.fields_by_name['type'].enum_type = proto_dot_types_dot_model_dot_resource__pb2._RESOURCE_TYPE
DESCRIPTOR.message_types_by_name['CreateFilePrepare'] = _CREATEFILEPREPARE
DESCRIPTOR.message_types_by_name['CreateFileMeta'] = _CREATEFILEMETA
DESCRIPTOR.message_types_by_name['ReCreateFilePrepare'] = _RECREATEFILEPREPARE
DESCRIPTOR.message_types_by_name['ReCreateFileMeta'] = _RECREATEFILEMETA
DESCRIPTOR.message_types_by_name['DescribeFileMeta'] = _DESCRIBEFILEMETA
DESCRIPTOR.message_types_by_name['UpdateFileMeta'] = _UPDATEFILEMETA
DESCRIPTOR.message_types_by_name['ListFileMetas'] = _LISTFILEMETAS
DESCRIPTOR.message_types_by_name['DeleteFileMetas'] = _DELETEFILEMETAS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateFilePrepare = _reflection.GeneratedProtocolMessageType('CreateFilePrepare', (_message.Message,), {
  'DESCRIPTOR' : _CREATEFILEPREPARE,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.CreateFilePrepare)
  })
_sym_db.RegisterMessage(CreateFilePrepare)

CreateFileMeta = _reflection.GeneratedProtocolMessageType('CreateFileMeta', (_message.Message,), {
  'DESCRIPTOR' : _CREATEFILEMETA,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.CreateFileMeta)
  })
_sym_db.RegisterMessage(CreateFileMeta)

ReCreateFilePrepare = _reflection.GeneratedProtocolMessageType('ReCreateFilePrepare', (_message.Message,), {
  'DESCRIPTOR' : _RECREATEFILEPREPARE,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.ReCreateFilePrepare)
  })
_sym_db.RegisterMessage(ReCreateFilePrepare)

ReCreateFileMeta = _reflection.GeneratedProtocolMessageType('ReCreateFileMeta', (_message.Message,), {
  'DESCRIPTOR' : _RECREATEFILEMETA,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.ReCreateFileMeta)
  })
_sym_db.RegisterMessage(ReCreateFileMeta)

DescribeFileMeta = _reflection.GeneratedProtocolMessageType('DescribeFileMeta', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEFILEMETA,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.DescribeFileMeta)
  })
_sym_db.RegisterMessage(DescribeFileMeta)

UpdateFileMeta = _reflection.GeneratedProtocolMessageType('UpdateFileMeta', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEFILEMETA,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.UpdateFileMeta)
  })
_sym_db.RegisterMessage(UpdateFileMeta)

ListFileMetas = _reflection.GeneratedProtocolMessageType('ListFileMetas', (_message.Message,), {
  'DESCRIPTOR' : _LISTFILEMETAS,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.ListFileMetas)
  })
_sym_db.RegisterMessage(ListFileMetas)

DeleteFileMetas = _reflection.GeneratedProtocolMessageType('DeleteFileMetas', (_message.Message,), {
  'DESCRIPTOR' : _DELETEFILEMETAS,
  '__module__' : 'proto.types.request.resource_meta_pb2'
  # @@protoc_insertion_point(class_scope:request.DeleteFileMetas)
  })
_sym_db.RegisterMessage(DeleteFileMetas)


DESCRIPTOR._options = None
_CREATEFILEPREPARE.fields_by_name['space_id']._options = None
_CREATEFILEPREPARE.fields_by_name['pid']._options = None
_CREATEFILEPREPARE.fields_by_name['name']._options = None
_CREATEFILEPREPARE.fields_by_name['size']._options = None
_CREATEFILEMETA.fields_by_name['space_id']._options = None
_CREATEFILEMETA.fields_by_name['resource_id']._options = None
_CREATEFILEMETA.fields_by_name['pid']._options = None
_CREATEFILEMETA.fields_by_name['name']._options = None
_CREATEFILEMETA.fields_by_name['desc']._options = None
_CREATEFILEMETA.fields_by_name['etag']._options = None
_CREATEFILEMETA.fields_by_name['type']._options = None
_CREATEFILEMETA.fields_by_name['version']._options = None
_CREATEFILEMETA.fields_by_name['created_by']._options = None
_RECREATEFILEPREPARE.fields_by_name['space_id']._options = None
_RECREATEFILEPREPARE.fields_by_name['resource_id']._options = None
_RECREATEFILEPREPARE.fields_by_name['size']._options = None
_RECREATEFILEMETA.fields_by_name['space_id']._options = None
_RECREATEFILEMETA.fields_by_name['resource_id']._options = None
_RECREATEFILEMETA.fields_by_name['etag']._options = None
_RECREATEFILEMETA.fields_by_name['version']._options = None
_DESCRIBEFILEMETA.fields_by_name['resource_id']._options = None
_UPDATEFILEMETA.fields_by_name['space_id']._options = None
_UPDATEFILEMETA.fields_by_name['resource_id']._options = None
_UPDATEFILEMETA.fields_by_name['name']._options = None
_UPDATEFILEMETA.fields_by_name['desc']._options = None
_UPDATEFILEMETA.fields_by_name['type']._options = None
_LISTFILEMETAS.fields_by_name['space_id']._options = None
_LISTFILEMETAS.fields_by_name['limit']._options = None
_LISTFILEMETAS.fields_by_name['offset']._options = None
_LISTFILEMETAS.fields_by_name['type']._options = None
_LISTFILEMETAS.fields_by_name['name']._options = None
_LISTFILEMETAS.fields_by_name['search']._options = None
_LISTFILEMETAS.fields_by_name['sort_by']._options = None
_LISTFILEMETAS.fields_by_name['reverse']._options = None
_DELETEFILEMETAS.fields_by_name['space_id']._options = None
_DELETEFILEMETAS.fields_by_name['resource_ids']._options = None
# @@protoc_insertion_point(module_scope)
