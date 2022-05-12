# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/syncjob/mongodb.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model.syncjob import column_pb2 as proto_dot_types_dot_model_dot_syncjob_dot_column__pb2
from github.com.yu31.protoc_plugin.proto import gosql_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/syncjob/mongodb.proto',
  package='model',
  syntax='proto3',
  serialized_options=b'\n,com.dataomnis.gproto.types.pbmodel.pbsyncjobB\tPBMongodbP\000Z;github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbsyncjob',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\'proto/types/model/syncjob/mongodb.proto\x12\x05model\x1a&proto/types/model/syncjob/column.proto\x1a/github.com/yu31/protoc-plugin/proto/gosql.proto\"\x85\x01\n\rMongodbSource\x12\x1d\n\x06\x63olumn\x18\x01 \x03(\x0b\x32\r.model.Column\x12\x10\n\x08\x64\x61tabase\x18\x02 \x01(\t\x12\x17\n\x0f\x63ollection_name\x18\x03 \x01(\t\x12\x12\n\nfetch_size\x18\x04 \x01(\x05\x12\x0e\n\x06\x66ilter\x18\x05 \x01(\t:\x06\xca\xb2\x04\x02\n\x00\"\xa2\x02\n\rMongodbTarget\x12\x1d\n\x06\x63olumn\x18\x01 \x03(\x0b\x32\r.model.Column\x12\x10\n\x08\x64\x61tabase\x18\x02 \x01(\t\x12\x17\n\x0f\x63ollection_name\x18\x03 \x01(\t\x12\x13\n\x0breplace_key\x18\x04 \x01(\t\x12\x32\n\nwrite_mode\x18\x05 \x01(\x0e\x32\x1e.model.MongodbTarget.WriteMode\x12\x12\n\nbatch_size\x18\x06 \x01(\x05\x12\x1c\n\x14\x66lush_interval_mills\x18\x07 \x01(\x05\"D\n\tWriteMode\x12\x12\n\x0eWriteModeUnset\x10\x00\x12\n\n\x06insert\x10\x01\x12\x0b\n\x07replace\x10\x02\x12\n\n\x06update\x10\x03:\x06\xca\xb2\x04\x02\n\x00\x42x\n,com.dataomnis.gproto.types.pbmodel.pbsyncjobB\tPBMongodbP\x00Z;github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbsyncjobb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_syncjob_dot_column__pb2.DESCRIPTOR,github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2.DESCRIPTOR,])



_MONGODBTARGET_WRITEMODE = _descriptor.EnumDescriptor(
  name='WriteMode',
  full_name='model.MongodbTarget.WriteMode',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='WriteModeUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='insert', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='replace', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='update', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=490,
  serialized_end=558,
)
_sym_db.RegisterEnumDescriptor(_MONGODBTARGET_WRITEMODE)


_MONGODBSOURCE = _descriptor.Descriptor(
  name='MongodbSource',
  full_name='model.MongodbSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='column', full_name='model.MongodbSource.column', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='database', full_name='model.MongodbSource.database', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='collection_name', full_name='model.MongodbSource.collection_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='fetch_size', full_name='model.MongodbSource.fetch_size', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='filter', full_name='model.MongodbSource.filter', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=140,
  serialized_end=273,
)


_MONGODBTARGET = _descriptor.Descriptor(
  name='MongodbTarget',
  full_name='model.MongodbTarget',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='column', full_name='model.MongodbTarget.column', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='database', full_name='model.MongodbTarget.database', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='collection_name', full_name='model.MongodbTarget.collection_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='replace_key', full_name='model.MongodbTarget.replace_key', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='write_mode', full_name='model.MongodbTarget.write_mode', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='batch_size', full_name='model.MongodbTarget.batch_size', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='flush_interval_mills', full_name='model.MongodbTarget.flush_interval_mills', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _MONGODBTARGET_WRITEMODE,
  ],
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=276,
  serialized_end=566,
)

_MONGODBSOURCE.fields_by_name['column'].message_type = proto_dot_types_dot_model_dot_syncjob_dot_column__pb2._COLUMN
_MONGODBTARGET.fields_by_name['column'].message_type = proto_dot_types_dot_model_dot_syncjob_dot_column__pb2._COLUMN
_MONGODBTARGET.fields_by_name['write_mode'].enum_type = _MONGODBTARGET_WRITEMODE
_MONGODBTARGET_WRITEMODE.containing_type = _MONGODBTARGET
DESCRIPTOR.message_types_by_name['MongodbSource'] = _MONGODBSOURCE
DESCRIPTOR.message_types_by_name['MongodbTarget'] = _MONGODBTARGET
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MongodbSource = _reflection.GeneratedProtocolMessageType('MongodbSource', (_message.Message,), {
  'DESCRIPTOR' : _MONGODBSOURCE,
  '__module__' : 'proto.types.model.syncjob.mongodb_pb2'
  # @@protoc_insertion_point(class_scope:model.MongodbSource)
  })
_sym_db.RegisterMessage(MongodbSource)

MongodbTarget = _reflection.GeneratedProtocolMessageType('MongodbTarget', (_message.Message,), {
  'DESCRIPTOR' : _MONGODBTARGET,
  '__module__' : 'proto.types.model.syncjob.mongodb_pb2'
  # @@protoc_insertion_point(class_scope:model.MongodbTarget)
  })
_sym_db.RegisterMessage(MongodbTarget)


DESCRIPTOR._options = None
_MONGODBSOURCE._options = None
_MONGODBTARGET._options = None
# @@protoc_insertion_point(module_scope)
