# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/scheduler.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from proto.types.model import stream_job_pb2 as proto_dot_types_dot_model_dot_stream__job__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/scheduler.proto',
  package='model',
  syntax='proto3',
  serialized_options=b'\n\"com.dataomnis.gproto.types.pbmodelB\020PBModelSchedulerP\000Z1github.com/DataWorkbench/gproto/xgo/types/pbmodel',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n!proto/types/model/scheduler.proto\x12\x05model\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a\"proto/types/model/stream_job.proto\"\xb2\x02\n\x0eStreamJobEvent\x12\x37\n\x08property\x18\x01 \x01(\x0b\x32\x18.model.StreamJobPropertyB\x0b\xe2\xdf\x1f\x07\x12\x05\xe2\x01\x02\x10\x01\x12\x1e\n\tflink_url\x18\x02 \x01(\tB\x0b\xe2\xdf\x1f\x07\x12\x05\xc2\x01\x02\"\x00\x12\"\n\rflink_version\x18\x03 \x01(\tB\x0b\xe2\xdf\x1f\x07\x12\x05\xc2\x01\x02\"\x00\x12\x15\n\x07retries\x18\x04 \x01(\x05\x42\x04\xe2\xdf\x1f\x00\x12\x15\n\x07started\x18\x05 \x01(\x03\x42\x04\xe2\xdf\x1f\x00\",\n\x04Type\x12\r\n\tTypeUnset\x10\x00\x12\x07\n\x03Job\x10\x01\x12\x0c\n\x08Instance\x10\x02\"G\n\x06\x41\x63tion\x12\x0f\n\x0b\x41\x63tionUnset\x10\x00\x12\n\n\x06\x43reate\x10\x01\x12\n\n\x06Submit\x10\x02\x12\t\n\x05\x43heck\x10\x04\x12\t\n\x05Retry\x10\x03\x42k\n\"com.dataomnis.gproto.types.pbmodelB\x10PBModelSchedulerP\x00Z1github.com/DataWorkbench/gproto/xgo/types/pbmodelb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_stream__job__pb2.DESCRIPTOR,])



_STREAMJOBEVENT_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='model.StreamJobEvent.Type',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TypeUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Job', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Instance', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=323,
  serialized_end=367,
)
_sym_db.RegisterEnumDescriptor(_STREAMJOBEVENT_TYPE)

_STREAMJOBEVENT_ACTION = _descriptor.EnumDescriptor(
  name='Action',
  full_name='model.StreamJobEvent.Action',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ActionUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Create', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Submit', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Check', index=3, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Retry', index=4, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=369,
  serialized_end=440,
)
_sym_db.RegisterEnumDescriptor(_STREAMJOBEVENT_ACTION)


_STREAMJOBEVENT = _descriptor.Descriptor(
  name='StreamJobEvent',
  full_name='model.StreamJobEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='property', full_name='model.StreamJobEvent.property', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\342\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='flink_url', full_name='model.StreamJobEvent.flink_url', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\302\001\002\"\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='flink_version', full_name='model.StreamJobEvent.flink_version', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\302\001\002\"\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='retries', full_name='model.StreamJobEvent.retries', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='started', full_name='model.StreamJobEvent.started', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _STREAMJOBEVENT_TYPE,
    _STREAMJOBEVENT_ACTION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=134,
  serialized_end=440,
)

_STREAMJOBEVENT.fields_by_name['property'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOBPROPERTY
_STREAMJOBEVENT_TYPE.containing_type = _STREAMJOBEVENT
_STREAMJOBEVENT_ACTION.containing_type = _STREAMJOBEVENT
DESCRIPTOR.message_types_by_name['StreamJobEvent'] = _STREAMJOBEVENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

StreamJobEvent = _reflection.GeneratedProtocolMessageType('StreamJobEvent', (_message.Message,), {
  'DESCRIPTOR' : _STREAMJOBEVENT,
  '__module__' : 'proto.types.model.scheduler_pb2'
  # @@protoc_insertion_point(class_scope:model.StreamJobEvent)
  })
_sym_db.RegisterMessage(StreamJobEvent)


DESCRIPTOR._options = None
_STREAMJOBEVENT.fields_by_name['property']._options = None
_STREAMJOBEVENT.fields_by_name['flink_url']._options = None
_STREAMJOBEVENT.fields_by_name['flink_version']._options = None
_STREAMJOBEVENT.fields_by_name['retries']._options = None
_STREAMJOBEVENT.fields_by_name['started']._options = None
# @@protoc_insertion_point(module_scope)
