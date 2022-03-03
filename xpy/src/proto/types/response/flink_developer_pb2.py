# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/response/flink_developer.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from proto.types.model.flink import flink_job_pb2 as proto_dot_types_dot_model_dot_flink_dot_flink__job__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/response/flink_developer.proto',
  package='response',
  syntax='proto3',
  serialized_options=b'\n%com.dataomnis.gproto.types.pbresponseB\030PBResponseFlinkDeveloperP\000Z4github.com/DataWorkbench/gproto/xgo/types/pbresponse',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n*proto/types/response/flink_developer.proto\x12\x08response\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a\'proto/types/model/flink/flink_job.proto\"\x8f\x01\n\x0eSubmitFlinkJob\x12+\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x1d.response.SubmitFlinkJob.Code\x12\x15\n\x07message\x18\x02 \x01(\tB\x04\xe2\xdf\x1f\x00\"9\n\x04\x43ode\x12\r\n\tCodeUnset\x10\x00\x12\x0b\n\x07Success\x10\x01\x12\n\n\x06\x46\x61iled\x10\x02\x12\t\n\x05Retry\x10\x03\"k\n\x11GetFlinkJobStatus\x12\x11\n\tis_exists\x18\x01 \x01(\x08\x12,\n\x06status\x18\x02 \x01(\x0e\x32\x16.flink.FlinkJob.StatusB\x04\xe2\xdf\x1f\x00\x12\x15\n\x07message\x18\x03 \x01(\tB\x04\xe2\xdf\x1f\x00\"\x8d\x01\n\x10ValidateFlinkJob\x12\x31\n\x06result\x18\x01 \x01(\x0e\x32!.response.ValidateFlinkJob.Result\x12\x0f\n\x07message\x18\x02 \x01(\t\"5\n\x06Result\x12\x0f\n\x0bResultUnset\x10\x00\x12\x0b\n\x07\x43orrect\x10\x01\x12\r\n\tIncorrect\x10\x02\x42y\n%com.dataomnis.gproto.types.pbresponseB\x18PBResponseFlinkDeveloperP\x00Z4github.com/DataWorkbench/gproto/xgo/types/pbresponseb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_flink_dot_flink__job__pb2.DESCRIPTOR,])



_SUBMITFLINKJOB_CODE = _descriptor.EnumDescriptor(
  name='Code',
  full_name='response.SubmitFlinkJob.Code',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CodeUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Success', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Failed', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Retry', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=237,
  serialized_end=294,
)
_sym_db.RegisterEnumDescriptor(_SUBMITFLINKJOB_CODE)

_VALIDATEFLINKJOB_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='response.ValidateFlinkJob.Result',
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
  serialized_start=494,
  serialized_end=547,
)
_sym_db.RegisterEnumDescriptor(_VALIDATEFLINKJOB_RESULT)


_SUBMITFLINKJOB = _descriptor.Descriptor(
  name='SubmitFlinkJob',
  full_name='response.SubmitFlinkJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='response.SubmitFlinkJob.code', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.SubmitFlinkJob.message', index=1,
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
    _SUBMITFLINKJOB_CODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=151,
  serialized_end=294,
)


_GETFLINKJOBSTATUS = _descriptor.Descriptor(
  name='GetFlinkJobStatus',
  full_name='response.GetFlinkJobStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='is_exists', full_name='response.GetFlinkJobStatus.is_exists', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='response.GetFlinkJobStatus.status', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.GetFlinkJobStatus.message', index=2,
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
  serialized_start=296,
  serialized_end=403,
)


_VALIDATEFLINKJOB = _descriptor.Descriptor(
  name='ValidateFlinkJob',
  full_name='response.ValidateFlinkJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='response.ValidateFlinkJob.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.ValidateFlinkJob.message', index=1,
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
    _VALIDATEFLINKJOB_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=406,
  serialized_end=547,
)

_SUBMITFLINKJOB.fields_by_name['code'].enum_type = _SUBMITFLINKJOB_CODE
_SUBMITFLINKJOB_CODE.containing_type = _SUBMITFLINKJOB
_GETFLINKJOBSTATUS.fields_by_name['status'].enum_type = proto_dot_types_dot_model_dot_flink_dot_flink__job__pb2._FLINKJOB_STATUS
_VALIDATEFLINKJOB.fields_by_name['result'].enum_type = _VALIDATEFLINKJOB_RESULT
_VALIDATEFLINKJOB_RESULT.containing_type = _VALIDATEFLINKJOB
DESCRIPTOR.message_types_by_name['SubmitFlinkJob'] = _SUBMITFLINKJOB
DESCRIPTOR.message_types_by_name['GetFlinkJobStatus'] = _GETFLINKJOBSTATUS
DESCRIPTOR.message_types_by_name['ValidateFlinkJob'] = _VALIDATEFLINKJOB
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SubmitFlinkJob = _reflection.GeneratedProtocolMessageType('SubmitFlinkJob', (_message.Message,), {
  'DESCRIPTOR' : _SUBMITFLINKJOB,
  '__module__' : 'proto.types.response.flink_developer_pb2'
  # @@protoc_insertion_point(class_scope:response.SubmitFlinkJob)
  })
_sym_db.RegisterMessage(SubmitFlinkJob)

GetFlinkJobStatus = _reflection.GeneratedProtocolMessageType('GetFlinkJobStatus', (_message.Message,), {
  'DESCRIPTOR' : _GETFLINKJOBSTATUS,
  '__module__' : 'proto.types.response.flink_developer_pb2'
  # @@protoc_insertion_point(class_scope:response.GetFlinkJobStatus)
  })
_sym_db.RegisterMessage(GetFlinkJobStatus)

ValidateFlinkJob = _reflection.GeneratedProtocolMessageType('ValidateFlinkJob', (_message.Message,), {
  'DESCRIPTOR' : _VALIDATEFLINKJOB,
  '__module__' : 'proto.types.response.flink_developer_pb2'
  # @@protoc_insertion_point(class_scope:response.ValidateFlinkJob)
  })
_sym_db.RegisterMessage(ValidateFlinkJob)


DESCRIPTOR._options = None
_SUBMITFLINKJOB.fields_by_name['message']._options = None
_GETFLINKJOBSTATUS.fields_by_name['status']._options = None
_GETFLINKJOBSTATUS.fields_by_name['message']._options = None
# @@protoc_insertion_point(module_scope)
