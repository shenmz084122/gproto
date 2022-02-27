# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/response/developer.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from proto.types.model import stream_instance_pb2 as proto_dot_types_dot_model_dot_stream__instance__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/response/developer.proto',
  package='response',
  syntax='proto3',
  serialized_options=b'\n%com.dataomnis.gproto.types.pbresponseB\023PBResponseDeveloperP\000Z4github.com/DataWorkbench/gproto/xgo/types/pbresponse',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n$proto/types/response/developer.proto\x12\x08response\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a\'proto/types/model/stream_instance.proto\"n\n\x11SubmitFlinkJob_v2\x12*\n\x05State\x18\x01 \x01(\x0e\x32\x1b.model.StreamInstance.State\x12\x15\n\x07message\x18\x02 \x01(\tB\x04\xe2\xdf\x1f\x00\x12\x16\n\x08\x66link_id\x18\x03 \x01(\tB\x04\xe2\xdf\x1f\x00\"Y\n\x0eGetFlinkJob_v2\x12\x30\n\x05State\x18\x01 \x01(\x0e\x32\x1b.model.StreamInstance.StateB\x04\xe2\xdf\x1f\x00\x12\x15\n\x07message\x18\x02 \x01(\tB\x04\xe2\xdf\x1f\x00\"\x93\x01\n\x13ValidateFlinkJob_v2\x12\x34\n\x06result\x18\x01 \x01(\x0e\x32$.response.ValidateFlinkJob_v2.Result\x12\x0f\n\x07message\x18\x02 \x01(\t\"5\n\x06Result\x12\x0f\n\x0bResultUnset\x10\x00\x12\x0b\n\x07\x43orrect\x10\x01\x12\r\n\tIncorrect\x10\x02\x42t\n%com.dataomnis.gproto.types.pbresponseB\x13PBResponseDeveloperP\x00Z4github.com/DataWorkbench/gproto/xgo/types/pbresponseb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_stream__instance__pb2.DESCRIPTOR,])



_VALIDATEFLINKJOB_V2_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='response.ValidateFlinkJob_v2.Result',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ResultUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Correct', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Incorrect', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=442,
  serialized_end=495,
)
_sym_db.RegisterEnumDescriptor(_VALIDATEFLINKJOB_V2_RESULT)


_SUBMITFLINKJOB_V2 = _descriptor.Descriptor(
  name='SubmitFlinkJob_v2',
  full_name='response.SubmitFlinkJob_v2',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='State', full_name='response.SubmitFlinkJob_v2.State', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.SubmitFlinkJob_v2.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='flink_id', full_name='response.SubmitFlinkJob_v2.flink_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=144,
  serialized_end=254,
)


_GETFLINKJOB_V2 = _descriptor.Descriptor(
  name='GetFlinkJob_v2',
  full_name='response.GetFlinkJob_v2',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='State', full_name='response.GetFlinkJob_v2.State', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.GetFlinkJob_v2.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=256,
  serialized_end=345,
)


_VALIDATEFLINKJOB_V2 = _descriptor.Descriptor(
  name='ValidateFlinkJob_v2',
  full_name='response.ValidateFlinkJob_v2',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='response.ValidateFlinkJob_v2.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.ValidateFlinkJob_v2.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _VALIDATEFLINKJOB_V2_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=348,
  serialized_end=495,
)

_SUBMITFLINKJOB_V2.fields_by_name['State'].enum_type = proto_dot_types_dot_model_dot_stream__instance__pb2._STREAMINSTANCE_STATE
_GETFLINKJOB_V2.fields_by_name['State'].enum_type = proto_dot_types_dot_model_dot_stream__instance__pb2._STREAMINSTANCE_STATE
_VALIDATEFLINKJOB_V2.fields_by_name['result'].enum_type = _VALIDATEFLINKJOB_V2_RESULT
_VALIDATEFLINKJOB_V2_RESULT.containing_type = _VALIDATEFLINKJOB_V2
DESCRIPTOR.message_types_by_name['SubmitFlinkJob_v2'] = _SUBMITFLINKJOB_V2
DESCRIPTOR.message_types_by_name['GetFlinkJob_v2'] = _GETFLINKJOB_V2
DESCRIPTOR.message_types_by_name['ValidateFlinkJob_v2'] = _VALIDATEFLINKJOB_V2
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SubmitFlinkJob_v2 = _reflection.GeneratedProtocolMessageType('SubmitFlinkJob_v2', (_message.Message,), {
  'DESCRIPTOR' : _SUBMITFLINKJOB_V2,
  '__module__' : 'proto.types.response.developer_pb2'
  # @@protoc_insertion_point(class_scope:response.SubmitFlinkJob_v2)
  })
_sym_db.RegisterMessage(SubmitFlinkJob_v2)

GetFlinkJob_v2 = _reflection.GeneratedProtocolMessageType('GetFlinkJob_v2', (_message.Message,), {
  'DESCRIPTOR' : _GETFLINKJOB_V2,
  '__module__' : 'proto.types.response.developer_pb2'
  # @@protoc_insertion_point(class_scope:response.GetFlinkJob_v2)
  })
_sym_db.RegisterMessage(GetFlinkJob_v2)

ValidateFlinkJob_v2 = _reflection.GeneratedProtocolMessageType('ValidateFlinkJob_v2', (_message.Message,), {
  'DESCRIPTOR' : _VALIDATEFLINKJOB_V2,
  '__module__' : 'proto.types.response.developer_pb2'
  # @@protoc_insertion_point(class_scope:response.ValidateFlinkJob_v2)
  })
_sym_db.RegisterMessage(ValidateFlinkJob_v2)


DESCRIPTOR._options = None
_SUBMITFLINKJOB_V2.fields_by_name['message']._options = None
_SUBMITFLINKJOB_V2.fields_by_name['flink_id']._options = None
_GETFLINKJOB_V2.fields_by_name['State']._options = None
_GETFLINKJOB_V2.fields_by_name['message']._options = None
# @@protoc_insertion_point(module_scope)