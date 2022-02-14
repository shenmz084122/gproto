# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/resource.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/resource.proto',
  package='model',
  syntax='proto3',
  serialized_options=b'\n\"com.dataomnis.gproto.types.pbmodelB\017PBModelResourceP\000Z1github.com/DataWorkbench/gproto/xgo/types/pbmodel',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n proto/types/model/resource.proto\x12\x05model\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\"\xca\x04\n\x08Resource\x12%\n\x08space_id\x18\x01 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12!\n\x0bresource_id\x18\x02 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x12\x0b\n\x03pid\x18\x03 \x01(\t\x12\x14\n\x0cis_directory\x18\x04 \x01(\x08\x12\x1e\n\x04name\x18\x05 \x01(\tB\x10\xe2\xdf\x1f\x0c\x12\n\xc2\x01\x07\x90\x02\x02\x98\x02\x80\x01\x12\"\n\x0b\x64\x65scription\x18\x06 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\xc8\x01\x80\x08\x12\x19\n\x04size\x18\x07 \x01(\x03\x42\x0b\xe2\xdf\x1f\x07\x12\x05\xb2\x01\x02@\x00\x12\x1a\n\x04\x65tag\x18\x08 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01 \x12\x31\n\x04type\x18\t \x01(\x0e\x32\x14.model.Resource.TypeB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12\x35\n\x06status\x18\n \x01(\x0e\x32\x16.model.Resource.StatusB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12\x1d\n\x07version\x18\x0b \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x10\x12 \n\ncreated_by\x18\x0c \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\x80\x02\x01\x12\x1c\n\x07\x63reated\x18\r \x01(\x03\x42\x0b\xe2\xdf\x1f\x07\x12\x05\xb2\x01\x02\x30\x00\x12\x1c\n\x07updated\x18\x0e \x01(\x03\x42\x0b\xe2\xdf\x1f\x07\x12\x05\xb2\x01\x02\x30\x00\":\n\x04Type\x12\x11\n\rResourceUnset\x10\x00\x12\x07\n\x03Jar\x10\x01\x12\x07\n\x03Udf\x10\x02\x12\r\n\tCONNECTOR\x10\x03\"3\n\x06Status\x12\x0f\n\x0bStatusUnset\x10\x00\x12\x0b\n\x07\x44\x65leted\x10\x01\x12\x0b\n\x07\x45nabled\x10\x02\x42j\n\"com.dataomnis.gproto.types.pbmodelB\x0fPBModelResourceP\x00Z1github.com/DataWorkbench/gproto/xgo/types/pbmodelb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,])



_RESOURCE_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='model.Resource.Type',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ResourceUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Jar', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Udf', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='CONNECTOR', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=572,
  serialized_end=630,
)
_sym_db.RegisterEnumDescriptor(_RESOURCE_TYPE)

_RESOURCE_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='model.Resource.Status',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='StatusUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Deleted', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Enabled', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=632,
  serialized_end=683,
)
_sym_db.RegisterEnumDescriptor(_RESOURCE_STATUS)


_RESOURCE = _descriptor.Descriptor(
  name='Resource',
  full_name='model.Resource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='space_id', full_name='model.Resource.space_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_id', full_name='model.Resource.resource_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pid', full_name='model.Resource.pid', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='is_directory', full_name='model.Resource.is_directory', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='model.Resource.name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\014\022\n\302\001\007\220\002\002\230\002\200\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='description', full_name='model.Resource.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\310\001\200\010', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='size', full_name='model.Resource.size', index=6,
      number=7, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\262\001\002@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='etag', full_name='model.Resource.etag', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001 ', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='model.Resource.type', index=8,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='model.Resource.status', index=9,
      number=10, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='model.Resource.version', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\020', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_by', full_name='model.Resource.created_by', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\200\002\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created', full_name='model.Resource.created', index=12,
      number=13, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\262\001\0020\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated', full_name='model.Resource.updated', index=13,
      number=14, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\262\001\0020\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _RESOURCE_TYPE,
    _RESOURCE_STATUS,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=97,
  serialized_end=683,
)

_RESOURCE.fields_by_name['type'].enum_type = _RESOURCE_TYPE
_RESOURCE.fields_by_name['status'].enum_type = _RESOURCE_STATUS
_RESOURCE_TYPE.containing_type = _RESOURCE
_RESOURCE_STATUS.containing_type = _RESOURCE
DESCRIPTOR.message_types_by_name['Resource'] = _RESOURCE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Resource = _reflection.GeneratedProtocolMessageType('Resource', (_message.Message,), {
  'DESCRIPTOR' : _RESOURCE,
  '__module__' : 'proto.types.model.resource_pb2'
  # @@protoc_insertion_point(class_scope:model.Resource)
  })
_sym_db.RegisterMessage(Resource)


DESCRIPTOR._options = None
_RESOURCE.fields_by_name['space_id']._options = None
_RESOURCE.fields_by_name['resource_id']._options = None
_RESOURCE.fields_by_name['name']._options = None
_RESOURCE.fields_by_name['description']._options = None
_RESOURCE.fields_by_name['size']._options = None
_RESOURCE.fields_by_name['etag']._options = None
_RESOURCE.fields_by_name['type']._options = None
_RESOURCE.fields_by_name['status']._options = None
_RESOURCE.fields_by_name['version']._options = None
_RESOURCE.fields_by_name['created_by']._options = None
_RESOURCE.fields_by_name['created']._options = None
_RESOURCE.fields_by_name['updated']._options = None
# @@protoc_insertion_point(module_scope)
