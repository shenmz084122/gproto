# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/quota.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import gosql_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/quota.proto',
  package='model',
  syntax='proto3',
  serialized_options=b'\n\"com.dataomnis.gproto.types.pbmodelB\014PBModelQuotaP\000Z1github.com/DataWorkbench/gproto/xgo/types/pbmodel',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1dproto/types/model/quota.proto\x12\x05model\x1a/github.com/yu31/protoc-plugin/proto/gosql.proto\"\'\n\x0eQuotaWorkspace\x12\r\n\x05limit\x18\x01 \x01(\x03:\x06\xca\xb2\x04\x02\n\x00\"\'\n\x0eQuotaStreamJob\x12\r\n\x05limit\x18\x01 \x01(\x03:\x06\xca\xb2\x04\x02\n\x00\"(\n\x0fQuotaDataSource\x12\r\n\x05limit\x18\x01 \x01(\x03:\x06\xca\xb2\x04\x02\n\x00\"!\n\x08QuotaUDF\x12\r\n\x05limit\x18\x01 \x01(\x03:\x06\xca\xb2\x04\x02\n\x00\"D\n\tQuotaFile\x12\r\n\x05limit\x18\x01 \x01(\x03\x12\x0c\n\x04size\x18\x02 \x01(\x03\x12\x12\n\nsize_total\x18\x03 \x01(\x03:\x06\xca\xb2\x04\x02\n\x00\"H\n\x11QuotaFlinkCluster\x12\r\n\x05limit\x18\x01 \x01(\x03\x12\n\n\x02\x63u\x18\x02 \x01(\x02\x12\x10\n\x08\x63u_total\x18\x03 \x01(\x02:\x06\xca\xb2\x04\x02\n\x00\"%\n\x0cQuotaNetwork\x12\r\n\x05limit\x18\x01 \x01(\x03:\x06\xca\xb2\x04\x02\n\x00\"%\n\x0cQuotaSyncJob\x12\r\n\x05limit\x18\x01 \x01(\x03:\x06\xca\xb2\x04\x02\n\x00\"\xda\x02\n\tUserQuota\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12(\n\tworkspace\x18\x02 \x01(\x0b\x32\x15.model.QuotaWorkspace\x12)\n\nstream_job\x18\x03 \x01(\x0b\x32\x15.model.QuotaStreamJob\x12+\n\x0b\x64\x61ta_source\x18\x04 \x01(\x0b\x32\x16.model.QuotaDataSource\x12\x1c\n\x03udf\x18\x05 \x01(\x0b\x32\x0f.model.QuotaUDF\x12\x1e\n\x04\x66ile\x18\x06 \x01(\x0b\x32\x10.model.QuotaFile\x12/\n\rflink_cluster\x18\x07 \x01(\x0b\x32\x18.model.QuotaFlinkCluster\x12$\n\x07network\x18\x08 \x01(\x0b\x32\x13.model.QuotaNetwork\x12%\n\x08sync_job\x18\t \x01(\x0b\x32\x13.model.QuotaSyncJobBg\n\"com.dataomnis.gproto.types.pbmodelB\x0cPBModelQuotaP\x00Z1github.com/DataWorkbench/gproto/xgo/types/pbmodelb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2.DESCRIPTOR,])




_QUOTAWORKSPACE = _descriptor.Descriptor(
  name='QuotaWorkspace',
  full_name='model.QuotaWorkspace',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaWorkspace.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=89,
  serialized_end=128,
)


_QUOTASTREAMJOB = _descriptor.Descriptor(
  name='QuotaStreamJob',
  full_name='model.QuotaStreamJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaStreamJob.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=130,
  serialized_end=169,
)


_QUOTADATASOURCE = _descriptor.Descriptor(
  name='QuotaDataSource',
  full_name='model.QuotaDataSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaDataSource.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=171,
  serialized_end=211,
)


_QUOTAUDF = _descriptor.Descriptor(
  name='QuotaUDF',
  full_name='model.QuotaUDF',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaUDF.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=213,
  serialized_end=246,
)


_QUOTAFILE = _descriptor.Descriptor(
  name='QuotaFile',
  full_name='model.QuotaFile',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaFile.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='size', full_name='model.QuotaFile.size', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='size_total', full_name='model.QuotaFile.size_total', index=2,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=248,
  serialized_end=316,
)


_QUOTAFLINKCLUSTER = _descriptor.Descriptor(
  name='QuotaFlinkCluster',
  full_name='model.QuotaFlinkCluster',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaFlinkCluster.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cu', full_name='model.QuotaFlinkCluster.cu', index=1,
      number=2, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cu_total', full_name='model.QuotaFlinkCluster.cu_total', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=318,
  serialized_end=390,
)


_QUOTANETWORK = _descriptor.Descriptor(
  name='QuotaNetwork',
  full_name='model.QuotaNetwork',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaNetwork.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=392,
  serialized_end=429,
)


_QUOTASYNCJOB = _descriptor.Descriptor(
  name='QuotaSyncJob',
  full_name='model.QuotaSyncJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='model.QuotaSyncJob.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
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
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=431,
  serialized_end=468,
)


_USERQUOTA = _descriptor.Descriptor(
  name='UserQuota',
  full_name='model.UserQuota',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='model.UserQuota.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='workspace', full_name='model.UserQuota.workspace', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='stream_job', full_name='model.UserQuota.stream_job', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data_source', full_name='model.UserQuota.data_source', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='udf', full_name='model.UserQuota.udf', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='file', full_name='model.UserQuota.file', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='flink_cluster', full_name='model.UserQuota.flink_cluster', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='network', full_name='model.UserQuota.network', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sync_job', full_name='model.UserQuota.sync_job', index=8,
      number=9, type=11, cpp_type=10, label=1,
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
  serialized_start=471,
  serialized_end=817,
)

_USERQUOTA.fields_by_name['workspace'].message_type = _QUOTAWORKSPACE
_USERQUOTA.fields_by_name['stream_job'].message_type = _QUOTASTREAMJOB
_USERQUOTA.fields_by_name['data_source'].message_type = _QUOTADATASOURCE
_USERQUOTA.fields_by_name['udf'].message_type = _QUOTAUDF
_USERQUOTA.fields_by_name['file'].message_type = _QUOTAFILE
_USERQUOTA.fields_by_name['flink_cluster'].message_type = _QUOTAFLINKCLUSTER
_USERQUOTA.fields_by_name['network'].message_type = _QUOTANETWORK
_USERQUOTA.fields_by_name['sync_job'].message_type = _QUOTASYNCJOB
DESCRIPTOR.message_types_by_name['QuotaWorkspace'] = _QUOTAWORKSPACE
DESCRIPTOR.message_types_by_name['QuotaStreamJob'] = _QUOTASTREAMJOB
DESCRIPTOR.message_types_by_name['QuotaDataSource'] = _QUOTADATASOURCE
DESCRIPTOR.message_types_by_name['QuotaUDF'] = _QUOTAUDF
DESCRIPTOR.message_types_by_name['QuotaFile'] = _QUOTAFILE
DESCRIPTOR.message_types_by_name['QuotaFlinkCluster'] = _QUOTAFLINKCLUSTER
DESCRIPTOR.message_types_by_name['QuotaNetwork'] = _QUOTANETWORK
DESCRIPTOR.message_types_by_name['QuotaSyncJob'] = _QUOTASYNCJOB
DESCRIPTOR.message_types_by_name['UserQuota'] = _USERQUOTA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

QuotaWorkspace = _reflection.GeneratedProtocolMessageType('QuotaWorkspace', (_message.Message,), {
  'DESCRIPTOR' : _QUOTAWORKSPACE,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaWorkspace)
  })
_sym_db.RegisterMessage(QuotaWorkspace)

QuotaStreamJob = _reflection.GeneratedProtocolMessageType('QuotaStreamJob', (_message.Message,), {
  'DESCRIPTOR' : _QUOTASTREAMJOB,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaStreamJob)
  })
_sym_db.RegisterMessage(QuotaStreamJob)

QuotaDataSource = _reflection.GeneratedProtocolMessageType('QuotaDataSource', (_message.Message,), {
  'DESCRIPTOR' : _QUOTADATASOURCE,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaDataSource)
  })
_sym_db.RegisterMessage(QuotaDataSource)

QuotaUDF = _reflection.GeneratedProtocolMessageType('QuotaUDF', (_message.Message,), {
  'DESCRIPTOR' : _QUOTAUDF,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaUDF)
  })
_sym_db.RegisterMessage(QuotaUDF)

QuotaFile = _reflection.GeneratedProtocolMessageType('QuotaFile', (_message.Message,), {
  'DESCRIPTOR' : _QUOTAFILE,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaFile)
  })
_sym_db.RegisterMessage(QuotaFile)

QuotaFlinkCluster = _reflection.GeneratedProtocolMessageType('QuotaFlinkCluster', (_message.Message,), {
  'DESCRIPTOR' : _QUOTAFLINKCLUSTER,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaFlinkCluster)
  })
_sym_db.RegisterMessage(QuotaFlinkCluster)

QuotaNetwork = _reflection.GeneratedProtocolMessageType('QuotaNetwork', (_message.Message,), {
  'DESCRIPTOR' : _QUOTANETWORK,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaNetwork)
  })
_sym_db.RegisterMessage(QuotaNetwork)

QuotaSyncJob = _reflection.GeneratedProtocolMessageType('QuotaSyncJob', (_message.Message,), {
  'DESCRIPTOR' : _QUOTASYNCJOB,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.QuotaSyncJob)
  })
_sym_db.RegisterMessage(QuotaSyncJob)

UserQuota = _reflection.GeneratedProtocolMessageType('UserQuota', (_message.Message,), {
  'DESCRIPTOR' : _USERQUOTA,
  '__module__' : 'proto.types.model.quota_pb2'
  # @@protoc_insertion_point(class_scope:model.UserQuota)
  })
_sym_db.RegisterMessage(UserQuota)


DESCRIPTOR._options = None
_QUOTAWORKSPACE._options = None
_QUOTASTREAMJOB._options = None
_QUOTADATASOURCE._options = None
_QUOTAUDF._options = None
_QUOTAFILE._options = None
_QUOTAFLINKCLUSTER._options = None
_QUOTANETWORK._options = None
_QUOTASYNCJOB._options = None
# @@protoc_insertion_point(module_scope)
