# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/response/stream_job_manage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model import stream_job_pb2 as proto_dot_types_dot_model_dot_stream__job__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/response/stream_job_manage.proto',
  package='response',
  syntax='proto3',
  serialized_options=b'\n%com.dataomnis.gproto.types.pbresponseB\031PBResponseStreamJobManageP\000Z4github.com/DataWorkbench/gproto/xgo/types/pbresponse',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n,proto/types/response/stream_job_manage.proto\x12\x08response\x1a\"proto/types/model/stream_job.proto\"R\n\x0eListStreamJobs\x12\x1f\n\x05infos\x18\x01 \x03(\x0b\x32\x10.model.StreamJob\x12\x10\n\x08has_more\x18\x02 \x01(\x08\x12\r\n\x05total\x18\x03 \x01(\x03\"\x1d\n\x0f\x43reateStreamJob\x12\n\n\x02id\x18\x01 \x01(\t\"3\n\x11\x44\x65scribeStreamJob\x12\x1e\n\x04info\x18\x01 \x01(\x0b\x32\x10.model.StreamJob\"6\n\x10GetStreamJobCode\x12\"\n\x04\x63ode\x18\x01 \x01(\x0b\x32\x14.model.StreamJobCode\"\x93\x01\n\x13StreamJobCodeSyntax\x12\x34\n\x06result\x18\x01 \x01(\x0e\x32$.response.StreamJobCodeSyntax.Result\x12\x0f\n\x07message\x18\x02 \x01(\t\"5\n\x06Result\x12\x0f\n\x0bResultUnset\x10\x00\x12\x0b\n\x07\x43orrect\x10\x01\x12\r\n\tIncorrect\x10\x02\"6\n\x10GetStreamJobArgs\x12\"\n\x04info\x18\x01 \x01(\x0b\x32\x14.model.StreamJobArgs\">\n\x14GetStreamJobSchedule\x12&\n\x04info\x18\x01 \x01(\x0b\x32\x18.model.StreamJobSchedule\"`\n\x15ListReleaseStreamJobs\x12&\n\x05infos\x18\x01 \x03(\x0b\x32\x17.model.StreamJobRelease\x12\x10\n\x08has_more\x18\x02 \x01(\x08\x12\r\n\x05total\x18\x03 \x01(\x03\"J\n\x15ListStreamJobVersions\x12\x1f\n\x05infos\x18\x01 \x03(\x0b\x32\x10.model.StreamJob\x12\x10\n\x08has_more\x18\x02 \x01(\x08\"&\n\x15ListBuiltInConnectors\x12\r\n\x05items\x18\x01 \x03(\tBz\n%com.dataomnis.gproto.types.pbresponseB\x19PBResponseStreamJobManageP\x00Z4github.com/DataWorkbench/gproto/xgo/types/pbresponseb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_stream__job__pb2.DESCRIPTOR,])



_STREAMJOBCODESYNTAX_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='response.StreamJobCodeSyntax.Result',
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
  serialized_start=413,
  serialized_end=466,
)
_sym_db.RegisterEnumDescriptor(_STREAMJOBCODESYNTAX_RESULT)


_LISTSTREAMJOBS = _descriptor.Descriptor(
  name='ListStreamJobs',
  full_name='response.ListStreamJobs',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='infos', full_name='response.ListStreamJobs.infos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='has_more', full_name='response.ListStreamJobs.has_more', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total', full_name='response.ListStreamJobs.total', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=94,
  serialized_end=176,
)


_CREATESTREAMJOB = _descriptor.Descriptor(
  name='CreateStreamJob',
  full_name='response.CreateStreamJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='response.CreateStreamJob.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=178,
  serialized_end=207,
)


_DESCRIBESTREAMJOB = _descriptor.Descriptor(
  name='DescribeStreamJob',
  full_name='response.DescribeStreamJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='info', full_name='response.DescribeStreamJob.info', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=209,
  serialized_end=260,
)


_GETSTREAMJOBCODE = _descriptor.Descriptor(
  name='GetStreamJobCode',
  full_name='response.GetStreamJobCode',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='response.GetStreamJobCode.code', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=262,
  serialized_end=316,
)


_STREAMJOBCODESYNTAX = _descriptor.Descriptor(
  name='StreamJobCodeSyntax',
  full_name='response.StreamJobCodeSyntax',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='response.StreamJobCodeSyntax.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.StreamJobCodeSyntax.message', index=1,
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
    _STREAMJOBCODESYNTAX_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=319,
  serialized_end=466,
)


_GETSTREAMJOBARGS = _descriptor.Descriptor(
  name='GetStreamJobArgs',
  full_name='response.GetStreamJobArgs',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='info', full_name='response.GetStreamJobArgs.info', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=468,
  serialized_end=522,
)


_GETSTREAMJOBSCHEDULE = _descriptor.Descriptor(
  name='GetStreamJobSchedule',
  full_name='response.GetStreamJobSchedule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='info', full_name='response.GetStreamJobSchedule.info', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=524,
  serialized_end=586,
)


_LISTRELEASESTREAMJOBS = _descriptor.Descriptor(
  name='ListReleaseStreamJobs',
  full_name='response.ListReleaseStreamJobs',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='infos', full_name='response.ListReleaseStreamJobs.infos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='has_more', full_name='response.ListReleaseStreamJobs.has_more', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total', full_name='response.ListReleaseStreamJobs.total', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=588,
  serialized_end=684,
)


_LISTSTREAMJOBVERSIONS = _descriptor.Descriptor(
  name='ListStreamJobVersions',
  full_name='response.ListStreamJobVersions',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='infos', full_name='response.ListStreamJobVersions.infos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='has_more', full_name='response.ListStreamJobVersions.has_more', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=686,
  serialized_end=760,
)


_LISTBUILTINCONNECTORS = _descriptor.Descriptor(
  name='ListBuiltInConnectors',
  full_name='response.ListBuiltInConnectors',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='response.ListBuiltInConnectors.items', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=762,
  serialized_end=800,
)

_LISTSTREAMJOBS.fields_by_name['infos'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOB
_DESCRIBESTREAMJOB.fields_by_name['info'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOB
_GETSTREAMJOBCODE.fields_by_name['code'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOBCODE
_STREAMJOBCODESYNTAX.fields_by_name['result'].enum_type = _STREAMJOBCODESYNTAX_RESULT
_STREAMJOBCODESYNTAX_RESULT.containing_type = _STREAMJOBCODESYNTAX
_GETSTREAMJOBARGS.fields_by_name['info'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOBARGS
_GETSTREAMJOBSCHEDULE.fields_by_name['info'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOBSCHEDULE
_LISTRELEASESTREAMJOBS.fields_by_name['infos'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOBRELEASE
_LISTSTREAMJOBVERSIONS.fields_by_name['infos'].message_type = proto_dot_types_dot_model_dot_stream__job__pb2._STREAMJOB
DESCRIPTOR.message_types_by_name['ListStreamJobs'] = _LISTSTREAMJOBS
DESCRIPTOR.message_types_by_name['CreateStreamJob'] = _CREATESTREAMJOB
DESCRIPTOR.message_types_by_name['DescribeStreamJob'] = _DESCRIBESTREAMJOB
DESCRIPTOR.message_types_by_name['GetStreamJobCode'] = _GETSTREAMJOBCODE
DESCRIPTOR.message_types_by_name['StreamJobCodeSyntax'] = _STREAMJOBCODESYNTAX
DESCRIPTOR.message_types_by_name['GetStreamJobArgs'] = _GETSTREAMJOBARGS
DESCRIPTOR.message_types_by_name['GetStreamJobSchedule'] = _GETSTREAMJOBSCHEDULE
DESCRIPTOR.message_types_by_name['ListReleaseStreamJobs'] = _LISTRELEASESTREAMJOBS
DESCRIPTOR.message_types_by_name['ListStreamJobVersions'] = _LISTSTREAMJOBVERSIONS
DESCRIPTOR.message_types_by_name['ListBuiltInConnectors'] = _LISTBUILTINCONNECTORS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListStreamJobs = _reflection.GeneratedProtocolMessageType('ListStreamJobs', (_message.Message,), {
  'DESCRIPTOR' : _LISTSTREAMJOBS,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ListStreamJobs)
  })
_sym_db.RegisterMessage(ListStreamJobs)

CreateStreamJob = _reflection.GeneratedProtocolMessageType('CreateStreamJob', (_message.Message,), {
  'DESCRIPTOR' : _CREATESTREAMJOB,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.CreateStreamJob)
  })
_sym_db.RegisterMessage(CreateStreamJob)

DescribeStreamJob = _reflection.GeneratedProtocolMessageType('DescribeStreamJob', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBESTREAMJOB,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.DescribeStreamJob)
  })
_sym_db.RegisterMessage(DescribeStreamJob)

GetStreamJobCode = _reflection.GeneratedProtocolMessageType('GetStreamJobCode', (_message.Message,), {
  'DESCRIPTOR' : _GETSTREAMJOBCODE,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.GetStreamJobCode)
  })
_sym_db.RegisterMessage(GetStreamJobCode)

StreamJobCodeSyntax = _reflection.GeneratedProtocolMessageType('StreamJobCodeSyntax', (_message.Message,), {
  'DESCRIPTOR' : _STREAMJOBCODESYNTAX,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.StreamJobCodeSyntax)
  })
_sym_db.RegisterMessage(StreamJobCodeSyntax)

GetStreamJobArgs = _reflection.GeneratedProtocolMessageType('GetStreamJobArgs', (_message.Message,), {
  'DESCRIPTOR' : _GETSTREAMJOBARGS,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.GetStreamJobArgs)
  })
_sym_db.RegisterMessage(GetStreamJobArgs)

GetStreamJobSchedule = _reflection.GeneratedProtocolMessageType('GetStreamJobSchedule', (_message.Message,), {
  'DESCRIPTOR' : _GETSTREAMJOBSCHEDULE,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.GetStreamJobSchedule)
  })
_sym_db.RegisterMessage(GetStreamJobSchedule)

ListReleaseStreamJobs = _reflection.GeneratedProtocolMessageType('ListReleaseStreamJobs', (_message.Message,), {
  'DESCRIPTOR' : _LISTRELEASESTREAMJOBS,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ListReleaseStreamJobs)
  })
_sym_db.RegisterMessage(ListReleaseStreamJobs)

ListStreamJobVersions = _reflection.GeneratedProtocolMessageType('ListStreamJobVersions', (_message.Message,), {
  'DESCRIPTOR' : _LISTSTREAMJOBVERSIONS,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ListStreamJobVersions)
  })
_sym_db.RegisterMessage(ListStreamJobVersions)

ListBuiltInConnectors = _reflection.GeneratedProtocolMessageType('ListBuiltInConnectors', (_message.Message,), {
  'DESCRIPTOR' : _LISTBUILTINCONNECTORS,
  '__module__' : 'proto.types.response.stream_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ListBuiltInConnectors)
  })
_sym_db.RegisterMessage(ListBuiltInConnectors)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
