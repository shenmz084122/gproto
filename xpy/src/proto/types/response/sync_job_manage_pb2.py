# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/response/sync_job_manage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from proto.types.model import sync_job_pb2 as proto_dot_types_dot_model_dot_sync__job__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/response/sync_job_manage.proto',
  package='response',
  syntax='proto3',
  serialized_options=b'\n%com.dataomnis.gproto.types.pbresponseB\027PBResponseSyncJobManageP\000Z4github.com/DataWorkbench/gproto/xgo/types/pbresponse',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n*proto/types/response/sync_job_manage.proto\x12\x08response\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a proto/types/model/sync_job.proto\"N\n\x0cListSyncJobs\x12\x1d\n\x05infos\x18\x01 \x03(\x0b\x32\x0e.model.SyncJob\x12\x10\n\x08has_more\x18\x02 \x01(\x08\x12\r\n\x05total\x18\x03 \x01(\x03\"\x1b\n\rCreateSyncJob\x12\n\n\x02id\x18\x01 \x01(\t\"/\n\x0f\x44\x65scribeSyncJob\x12\x1c\n\x04info\x18\x01 \x01(\x0b\x32\x0e.model.SyncJob\"?\n\x0eGetSyncJobConf\x12-\n\x04\x63onf\x18\x01 \x01(\x0b\x32\x12.model.SyncJobConfB\x0b\xe2\xdf\x1f\x07\x12\x05\xe2\x01\x02\x18\x01\"K\n\x12GetSyncJobSchedule\x12\x35\n\x08schedule\x18\x01 \x01(\x0b\x32\x16.model.SyncJobScheduleB\x0b\xe2\xdf\x1f\x07\x12\x05\xe2\x01\x02\x18\x01\"\\\n\x13ListReleaseSyncJobs\x12$\n\x05infos\x18\x01 \x03(\x0b\x32\x15.model.SyncJobRelease\x12\x10\n\x08has_more\x18\x02 \x01(\x08\x12\r\n\x05total\x18\x03 \x01(\x03\"U\n\x13ListSyncJobVersions\x12\x1d\n\x05infos\x18\x01 \x03(\x0b\x32\x0e.model.SyncJob\x12\x10\n\x08has_more\x18\x02 \x01(\x08\x12\r\n\x05total\x18\x03 \x01(\x03\"*\n\x0fGenerateJobJson\x12\x17\n\x0fsync_job_script\x18\x01 \x01(\t\"!\n\x12\x43onvertSyncJobMode\x12\x0b\n\x03job\x18\x01 \x01(\t\".\n\x1dLoadSyncJobScheduleParameters\x12\r\n\x05items\x18\x01 \x03(\tBx\n%com.dataomnis.gproto.types.pbresponseB\x17PBResponseSyncJobManageP\x00Z4github.com/DataWorkbench/gproto/xgo/types/pbresponseb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_sync__job__pb2.DESCRIPTOR,])




_LISTSYNCJOBS = _descriptor.Descriptor(
  name='ListSyncJobs',
  full_name='response.ListSyncJobs',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='infos', full_name='response.ListSyncJobs.infos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='has_more', full_name='response.ListSyncJobs.has_more', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total', full_name='response.ListSyncJobs.total', index=2,
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
  serialized_start=143,
  serialized_end=221,
)


_CREATESYNCJOB = _descriptor.Descriptor(
  name='CreateSyncJob',
  full_name='response.CreateSyncJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='response.CreateSyncJob.id', index=0,
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
  serialized_start=223,
  serialized_end=250,
)


_DESCRIBESYNCJOB = _descriptor.Descriptor(
  name='DescribeSyncJob',
  full_name='response.DescribeSyncJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='info', full_name='response.DescribeSyncJob.info', index=0,
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
  serialized_start=252,
  serialized_end=299,
)


_GETSYNCJOBCONF = _descriptor.Descriptor(
  name='GetSyncJobConf',
  full_name='response.GetSyncJobConf',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='conf', full_name='response.GetSyncJobConf.conf', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\342\001\002\030\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=301,
  serialized_end=364,
)


_GETSYNCJOBSCHEDULE = _descriptor.Descriptor(
  name='GetSyncJobSchedule',
  full_name='response.GetSyncJobSchedule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schedule', full_name='response.GetSyncJobSchedule.schedule', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\342\001\002\030\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=366,
  serialized_end=441,
)


_LISTRELEASESYNCJOBS = _descriptor.Descriptor(
  name='ListReleaseSyncJobs',
  full_name='response.ListReleaseSyncJobs',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='infos', full_name='response.ListReleaseSyncJobs.infos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='has_more', full_name='response.ListReleaseSyncJobs.has_more', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total', full_name='response.ListReleaseSyncJobs.total', index=2,
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
  serialized_start=443,
  serialized_end=535,
)


_LISTSYNCJOBVERSIONS = _descriptor.Descriptor(
  name='ListSyncJobVersions',
  full_name='response.ListSyncJobVersions',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='infos', full_name='response.ListSyncJobVersions.infos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='has_more', full_name='response.ListSyncJobVersions.has_more', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total', full_name='response.ListSyncJobVersions.total', index=2,
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
  serialized_start=537,
  serialized_end=622,
)


_GENERATEJOBJSON = _descriptor.Descriptor(
  name='GenerateJobJson',
  full_name='response.GenerateJobJson',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='sync_job_script', full_name='response.GenerateJobJson.sync_job_script', index=0,
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
  serialized_start=624,
  serialized_end=666,
)


_CONVERTSYNCJOBMODE = _descriptor.Descriptor(
  name='ConvertSyncJobMode',
  full_name='response.ConvertSyncJobMode',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='job', full_name='response.ConvertSyncJobMode.job', index=0,
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
  serialized_start=668,
  serialized_end=701,
)


_LOADSYNCJOBSCHEDULEPARAMETERS = _descriptor.Descriptor(
  name='LoadSyncJobScheduleParameters',
  full_name='response.LoadSyncJobScheduleParameters',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='response.LoadSyncJobScheduleParameters.items', index=0,
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
  serialized_start=703,
  serialized_end=749,
)

_LISTSYNCJOBS.fields_by_name['infos'].message_type = proto_dot_types_dot_model_dot_sync__job__pb2._SYNCJOB
_DESCRIBESYNCJOB.fields_by_name['info'].message_type = proto_dot_types_dot_model_dot_sync__job__pb2._SYNCJOB
_GETSYNCJOBCONF.fields_by_name['conf'].message_type = proto_dot_types_dot_model_dot_sync__job__pb2._SYNCJOBCONF
_GETSYNCJOBSCHEDULE.fields_by_name['schedule'].message_type = proto_dot_types_dot_model_dot_sync__job__pb2._SYNCJOBSCHEDULE
_LISTRELEASESYNCJOBS.fields_by_name['infos'].message_type = proto_dot_types_dot_model_dot_sync__job__pb2._SYNCJOBRELEASE
_LISTSYNCJOBVERSIONS.fields_by_name['infos'].message_type = proto_dot_types_dot_model_dot_sync__job__pb2._SYNCJOB
DESCRIPTOR.message_types_by_name['ListSyncJobs'] = _LISTSYNCJOBS
DESCRIPTOR.message_types_by_name['CreateSyncJob'] = _CREATESYNCJOB
DESCRIPTOR.message_types_by_name['DescribeSyncJob'] = _DESCRIBESYNCJOB
DESCRIPTOR.message_types_by_name['GetSyncJobConf'] = _GETSYNCJOBCONF
DESCRIPTOR.message_types_by_name['GetSyncJobSchedule'] = _GETSYNCJOBSCHEDULE
DESCRIPTOR.message_types_by_name['ListReleaseSyncJobs'] = _LISTRELEASESYNCJOBS
DESCRIPTOR.message_types_by_name['ListSyncJobVersions'] = _LISTSYNCJOBVERSIONS
DESCRIPTOR.message_types_by_name['GenerateJobJson'] = _GENERATEJOBJSON
DESCRIPTOR.message_types_by_name['ConvertSyncJobMode'] = _CONVERTSYNCJOBMODE
DESCRIPTOR.message_types_by_name['LoadSyncJobScheduleParameters'] = _LOADSYNCJOBSCHEDULEPARAMETERS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListSyncJobs = _reflection.GeneratedProtocolMessageType('ListSyncJobs', (_message.Message,), {
  'DESCRIPTOR' : _LISTSYNCJOBS,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ListSyncJobs)
  })
_sym_db.RegisterMessage(ListSyncJobs)

CreateSyncJob = _reflection.GeneratedProtocolMessageType('CreateSyncJob', (_message.Message,), {
  'DESCRIPTOR' : _CREATESYNCJOB,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.CreateSyncJob)
  })
_sym_db.RegisterMessage(CreateSyncJob)

DescribeSyncJob = _reflection.GeneratedProtocolMessageType('DescribeSyncJob', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBESYNCJOB,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.DescribeSyncJob)
  })
_sym_db.RegisterMessage(DescribeSyncJob)

GetSyncJobConf = _reflection.GeneratedProtocolMessageType('GetSyncJobConf', (_message.Message,), {
  'DESCRIPTOR' : _GETSYNCJOBCONF,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.GetSyncJobConf)
  })
_sym_db.RegisterMessage(GetSyncJobConf)

GetSyncJobSchedule = _reflection.GeneratedProtocolMessageType('GetSyncJobSchedule', (_message.Message,), {
  'DESCRIPTOR' : _GETSYNCJOBSCHEDULE,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.GetSyncJobSchedule)
  })
_sym_db.RegisterMessage(GetSyncJobSchedule)

ListReleaseSyncJobs = _reflection.GeneratedProtocolMessageType('ListReleaseSyncJobs', (_message.Message,), {
  'DESCRIPTOR' : _LISTRELEASESYNCJOBS,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ListReleaseSyncJobs)
  })
_sym_db.RegisterMessage(ListReleaseSyncJobs)

ListSyncJobVersions = _reflection.GeneratedProtocolMessageType('ListSyncJobVersions', (_message.Message,), {
  'DESCRIPTOR' : _LISTSYNCJOBVERSIONS,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ListSyncJobVersions)
  })
_sym_db.RegisterMessage(ListSyncJobVersions)

GenerateJobJson = _reflection.GeneratedProtocolMessageType('GenerateJobJson', (_message.Message,), {
  'DESCRIPTOR' : _GENERATEJOBJSON,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.GenerateJobJson)
  })
_sym_db.RegisterMessage(GenerateJobJson)

ConvertSyncJobMode = _reflection.GeneratedProtocolMessageType('ConvertSyncJobMode', (_message.Message,), {
  'DESCRIPTOR' : _CONVERTSYNCJOBMODE,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.ConvertSyncJobMode)
  })
_sym_db.RegisterMessage(ConvertSyncJobMode)

LoadSyncJobScheduleParameters = _reflection.GeneratedProtocolMessageType('LoadSyncJobScheduleParameters', (_message.Message,), {
  'DESCRIPTOR' : _LOADSYNCJOBSCHEDULEPARAMETERS,
  '__module__' : 'proto.types.response.sync_job_manage_pb2'
  # @@protoc_insertion_point(class_scope:response.LoadSyncJobScheduleParameters)
  })
_sym_db.RegisterMessage(LoadSyncJobScheduleParameters)


DESCRIPTOR._options = None
_GETSYNCJOBCONF.fields_by_name['conf']._options = None
_GETSYNCJOBSCHEDULE.fields_by_name['schedule']._options = None
# @@protoc_insertion_point(module_scope)
