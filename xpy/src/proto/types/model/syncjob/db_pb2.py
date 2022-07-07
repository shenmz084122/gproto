# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/syncjob/db.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model.syncjob import column_pb2 as proto_dot_types_dot_model_dot_syncjob_dot_column__pb2
from proto.types.model.syncjob import baseenum_pb2 as proto_dot_types_dot_model_dot_syncjob_dot_baseenum__pb2
from github.com.yu31.protoc_plugin.proto import gosql_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/syncjob/db.proto',
  package='model',
  syntax='proto3',
  serialized_options=b'\n,com.dataomnis.gproto.types.pbmodel.pbsyncjobB\016PBRelationaldbP\000Z;github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbsyncjob',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\"proto/types/model/syncjob/db.proto\x12\x05model\x1a&proto/types/model/syncjob/column.proto\x1a(proto/types/model/syncjob/baseenum.proto\x1a/github.com/yu31/protoc-plugin/proto/gosql.proto\"\x97\x02\n\x08\x44\x42Source\x12\x1d\n\x06\x63olumn\x18\x01 \x03(\x0b\x32\r.model.Column\x12\r\n\x05table\x18\x02 \x03(\t\x12\x0e\n\x06schema\x18\x03 \x01(\t\x12\r\n\x05where\x18\x04 \x01(\t\x12\x10\n\x08split_pk\x18\x05 \x01(\t\x12\x33\n\x0cmapping_type\x18\x06 \x01(\x0e\x32\x1d.model.BaseEnum.ColumnMapping\x12\x35\n\x0e\x63ondition_type\x18\x07 \x01(\x0e\x32\x1d.model.BaseEnum.ConditionType\x12\'\n\rvisualization\x18\x08 \x01(\x0b\x32\x10.model.Condition\x12\x0f\n\x07\x65xpress\x18\t \x01(\t:\x06\xca\xb2\x04\x02\n\x00\"{\n\tCondition\x12\x0e\n\x06\x63olumn\x18\x01 \x01(\t\x12\x17\n\x0fstart_condition\x18\x02 \x01(\t\x12\x13\n\x0bstart_value\x18\x03 \x01(\t\x12\x15\n\rend_condition\x18\x04 \x01(\t\x12\x11\n\tend_value\x18\x05 \x01(\t:\x06\xca\xb2\x04\x02\n\x00\"\x93\x03\n\x08\x44\x42Target\x12\x1d\n\x06\x63olumn\x18\x01 \x03(\x0b\x32\r.model.Column\x12\r\n\x05table\x18\x02 \x03(\t\x12\x0e\n\x06schema\x18\x03 \x01(\t\x12\x0f\n\x07pre_sql\x18\x04 \x03(\t\x12\x10\n\x08post_sql\x18\x05 \x03(\t\x12-\n\nwrite_mode\x18\x06 \x01(\x0e\x32\x19.model.DBTarget.WriteMode\x12\x12\n\nbatch_size\x18\x07 \x01(\x05\x12\x12\n\nupdate_key\x18\x08 \x03(\t\x12*\n\x08semantic\x18\t \x01(\x0e\x32\x18.model.DBTarget.Semantic\x12\x14\n\x0cwith_no_lock\x18\n \x01(\t\"D\n\tWriteMode\x12\x12\n\x0eWriteModeUnset\x10\x00\x12\n\n\x06insert\x10\x01\x12\x0b\n\x07replace\x10\x02\x12\n\n\x06update\x10\x03\"?\n\x08Semantic\x12\x11\n\rSemanticUnset\x10\x00\x12\x0f\n\x0b\x41tLeastOnce\x10\x01\x12\x0f\n\x0b\x45xactlyOnce\x10\x02:\x06\xca\xb2\x04\x02\n\x00\x42}\n,com.dataomnis.gproto.types.pbmodel.pbsyncjobB\x0ePBRelationaldbP\x00Z;github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbsyncjobb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_syncjob_dot_column__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_syncjob_dot_baseenum__pb2.DESCRIPTOR,github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2.DESCRIPTOR,])



_DBTARGET_WRITEMODE = _descriptor.EnumDescriptor(
  name='WriteMode',
  full_name='model.DBTarget.WriteMode',
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
  serialized_start=846,
  serialized_end=914,
)
_sym_db.RegisterEnumDescriptor(_DBTARGET_WRITEMODE)

_DBTARGET_SEMANTIC = _descriptor.EnumDescriptor(
  name='Semantic',
  full_name='model.DBTarget.Semantic',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SemanticUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='AtLeastOnce', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ExactlyOnce', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=916,
  serialized_end=979,
)
_sym_db.RegisterEnumDescriptor(_DBTARGET_SEMANTIC)


_DBSOURCE = _descriptor.Descriptor(
  name='DBSource',
  full_name='model.DBSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='column', full_name='model.DBSource.column', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='table', full_name='model.DBSource.table', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='schema', full_name='model.DBSource.schema', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='where', full_name='model.DBSource.where', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='split_pk', full_name='model.DBSource.split_pk', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='mapping_type', full_name='model.DBSource.mapping_type', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='condition_type', full_name='model.DBSource.condition_type', index=6,
      number=7, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='visualization', full_name='model.DBSource.visualization', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='express', full_name='model.DBSource.express', index=8,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=177,
  serialized_end=456,
)


_CONDITION = _descriptor.Descriptor(
  name='Condition',
  full_name='model.Condition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='column', full_name='model.Condition.column', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_condition', full_name='model.Condition.start_condition', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_value', full_name='model.Condition.start_value', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_condition', full_name='model.Condition.end_condition', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_value', full_name='model.Condition.end_value', index=4,
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
  serialized_start=458,
  serialized_end=581,
)


_DBTARGET = _descriptor.Descriptor(
  name='DBTarget',
  full_name='model.DBTarget',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='column', full_name='model.DBTarget.column', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='table', full_name='model.DBTarget.table', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='schema', full_name='model.DBTarget.schema', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pre_sql', full_name='model.DBTarget.pre_sql', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='post_sql', full_name='model.DBTarget.post_sql', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='write_mode', full_name='model.DBTarget.write_mode', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='batch_size', full_name='model.DBTarget.batch_size', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='update_key', full_name='model.DBTarget.update_key', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='semantic', full_name='model.DBTarget.semantic', index=8,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='with_no_lock', full_name='model.DBTarget.with_no_lock', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _DBTARGET_WRITEMODE,
    _DBTARGET_SEMANTIC,
  ],
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=584,
  serialized_end=987,
)

_DBSOURCE.fields_by_name['column'].message_type = proto_dot_types_dot_model_dot_syncjob_dot_column__pb2._COLUMN
_DBSOURCE.fields_by_name['mapping_type'].enum_type = proto_dot_types_dot_model_dot_syncjob_dot_baseenum__pb2._BASEENUM_COLUMNMAPPING
_DBSOURCE.fields_by_name['condition_type'].enum_type = proto_dot_types_dot_model_dot_syncjob_dot_baseenum__pb2._BASEENUM_CONDITIONTYPE
_DBSOURCE.fields_by_name['visualization'].message_type = _CONDITION
_DBTARGET.fields_by_name['column'].message_type = proto_dot_types_dot_model_dot_syncjob_dot_column__pb2._COLUMN
_DBTARGET.fields_by_name['write_mode'].enum_type = _DBTARGET_WRITEMODE
_DBTARGET.fields_by_name['semantic'].enum_type = _DBTARGET_SEMANTIC
_DBTARGET_WRITEMODE.containing_type = _DBTARGET
_DBTARGET_SEMANTIC.containing_type = _DBTARGET
DESCRIPTOR.message_types_by_name['DBSource'] = _DBSOURCE
DESCRIPTOR.message_types_by_name['Condition'] = _CONDITION
DESCRIPTOR.message_types_by_name['DBTarget'] = _DBTARGET
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DBSource = _reflection.GeneratedProtocolMessageType('DBSource', (_message.Message,), {
  'DESCRIPTOR' : _DBSOURCE,
  '__module__' : 'proto.types.model.syncjob.db_pb2'
  # @@protoc_insertion_point(class_scope:model.DBSource)
  })
_sym_db.RegisterMessage(DBSource)

Condition = _reflection.GeneratedProtocolMessageType('Condition', (_message.Message,), {
  'DESCRIPTOR' : _CONDITION,
  '__module__' : 'proto.types.model.syncjob.db_pb2'
  # @@protoc_insertion_point(class_scope:model.Condition)
  })
_sym_db.RegisterMessage(Condition)

DBTarget = _reflection.GeneratedProtocolMessageType('DBTarget', (_message.Message,), {
  'DESCRIPTOR' : _DBTARGET,
  '__module__' : 'proto.types.model.syncjob.db_pb2'
  # @@protoc_insertion_point(class_scope:model.DBTarget)
  })
_sym_db.RegisterMessage(DBTarget)


DESCRIPTOR._options = None
_DBSOURCE._options = None
_CONDITION._options = None
_DBTARGET._options = None
# @@protoc_insertion_point(module_scope)
