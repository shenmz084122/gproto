# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/datasource/datasource_url.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from github.com.yu31.protoc_plugin.proto import gosql_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/datasource/datasource_url.proto',
  package='datasource',
  syntax='proto3',
  serialized_options=b'\n/com.dataomnis.gproto.types.pbmodel.pbdatasourceB\017PBDataSourceURLP\000Z>github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbdatasource',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n1proto/types/model/datasource/datasource_url.proto\x12\ndatasource\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a/github.com/yu31/protoc-plugin/proto/gosql.proto\"D\n\x04Host\x12\x1d\n\x04host\x18\x01 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12\x1d\n\x04port\x18\x02 \x01(\x05\x42\x0f\xe2\xdf\x1f\x0b\x12\t\xb2\x01\x06(\xff\xff\x03\x30\x00\"\xb5\x01\n\x08MySQLURL\x12\x1d\n\x04host\x18\x01 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12\x1d\n\x04port\x18\x02 \x01(\x05\x42\x0f\xe2\xdf\x1f\x0b\x12\t\xb2\x01\x06\x38\x80\x80\x04@\x00\x12\x1d\n\x04user\x18\x03 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12!\n\x08password\x18\x04 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12!\n\x08\x64\x61tabase\x18\x05 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@:\x06\xca\xb2\x04\x02\n\x00\"\xba\x01\n\rPostgreSQLURL\x12\x1d\n\x04host\x18\x01 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12\x1d\n\x04port\x18\x02 \x01(\x05\x42\x0f\xe2\xdf\x1f\x0b\x12\t\xb2\x01\x06\x38\x80\x80\x04@\x00\x12\x1d\n\x04user\x18\x03 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12!\n\x08password\x18\x04 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12!\n\x08\x64\x61tabase\x18\x05 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@:\x06\xca\xb2\x04\x02\n\x00\"\xba\x01\n\rClickHouseURL\x12\x1d\n\x04host\x18\x01 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12\x1d\n\x04port\x18\x02 \x01(\x05\x42\x0f\xe2\xdf\x1f\x0b\x12\t\xb2\x01\x06\x38\x80\x80\x04@\x00\x12\x1d\n\x04user\x18\x03 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12!\n\x08password\x18\x04 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12!\n\x08\x64\x61tabase\x18\x05 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@:\x06\xca\xb2\x04\x02\n\x00\"\xe7\x03\n\x06\x46tpURL\x12<\n\x08protocol\x18\x01 \x01(\x0e\x32\x1b.datasource.FtpURL.ProtocolB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12`\n\x0f\x63onnection_mode\x18\x02 \x01(\x0e\x32!.datasource.FtpURL.ConnectionModeB$\xe2\xdf\x1f\x13\n\x11\n\x08protocol\x12\x05\xda\x01\x02\x18\x01\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12?\n\x0bprivate_key\x18\x03 \x01(\tB*\xe2\xdf\x1f\x13\n\x11\n\x08protocol\x12\x05\xda\x01\x02\x18\x02\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\x80\x02\x01\x98\x02\x80\x10\x88\x05\x01\x12\x1d\n\x04host\x18\x04 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12\x1d\n\x04port\x18\x05 \x01(\x05\x42\x0f\xe2\xdf\x1f\x0b\x12\t\xb2\x01\x06\x38\x80\x80\x04@\x00\x12\x1d\n\x04user\x18\x06 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12!\n\x08password\x18\x07 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\"0\n\x08Protocol\x12\x11\n\rProtocolUnset\x10\x00\x12\x07\n\x03\x46TP\x10\x01\x12\x08\n\x04SFTP\x10\x02\"B\n\x0e\x43onnectionMode\x12\x17\n\x13\x43onnectionModeUnset\x10\x00\x12\n\n\x06\x41\x63tive\x10\x01\x12\x0b\n\x07Passive\x10\x02:\x06\xca\xb2\x04\x02\n\x00\"\x0f\n\x05S3URL:\x06\xca\xb2\x04\x02\n\x00\"K\n\x08KafkaURL\x12\x37\n\rkafka_brokers\x18\x01 \x03(\x0b\x32\x10.datasource.HostB\x0e\xe2\xdf\x1f\n\x12\x08\xea\x01\x05\x38\x80\x01@\x01:\x06\xca\xb2\x04\x02\n\x00\"8\n\x08HBaseURL\x12$\n\x06\x63onfig\x18\x01 \x01(\tB\x14\xe2\xdf\x1f\x10\x12\x0e\xc2\x01\x0b\x90\x02\x01\x98\x02\x80\x80\x01\xe8\x08\x01:\x06\xca\xb2\x04\x02\n\x00\"\x8c\x01\n\x07HDFSURL\x12\"\n\tname_node\x18\x01 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x01\x98\x02@\x12\x1d\n\x04port\x18\x02 \x01(\x05\x42\x0f\xe2\xdf\x1f\x0b\x12\t\xb2\x01\x06\x38\x80\x80\x04@\x00\x12\x36\n\x06\x63onfig\x18\x03 \x01(\tB&\xe2\xdf\x1f\x11\n\x0f\n\x06\x63onfig\x12\x05\xc2\x01\x02\"\x00\xe2\xdf\x1f\r\x12\x0b\xc2\x01\x08\x98\x02\x80\x80\x01\xe8\x08\x01:\x06\xca\xb2\x04\x02\n\x00\x42\x84\x01\n/com.dataomnis.gproto.types.pbmodel.pbdatasourceB\x0fPBDataSourceURLP\x00Z>github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbdatasourceb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_gosql__pb2.DESCRIPTOR,])



_FTPURL_PROTOCOL = _descriptor.EnumDescriptor(
  name='Protocol',
  full_name='datasource.FtpURL.Protocol',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ProtocolUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='FTP', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SFTP', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1163,
  serialized_end=1211,
)
_sym_db.RegisterEnumDescriptor(_FTPURL_PROTOCOL)

_FTPURL_CONNECTIONMODE = _descriptor.EnumDescriptor(
  name='ConnectionMode',
  full_name='datasource.FtpURL.ConnectionMode',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ConnectionModeUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Active', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Passive', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1213,
  serialized_end=1279,
)
_sym_db.RegisterEnumDescriptor(_FTPURL_CONNECTIONMODE)


_HOST = _descriptor.Descriptor(
  name='Host',
  full_name='datasource.Host',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host', full_name='datasource.Host.host', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port', full_name='datasource.Host.port', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\262\001\006(\377\377\0030\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=167,
  serialized_end=235,
)


_MYSQLURL = _descriptor.Descriptor(
  name='MySQLURL',
  full_name='datasource.MySQLURL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host', full_name='datasource.MySQLURL.host', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port', full_name='datasource.MySQLURL.port', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\262\001\0068\200\200\004@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user', full_name='datasource.MySQLURL.user', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='password', full_name='datasource.MySQLURL.password', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='database', full_name='datasource.MySQLURL.database', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=238,
  serialized_end=419,
)


_POSTGRESQLURL = _descriptor.Descriptor(
  name='PostgreSQLURL',
  full_name='datasource.PostgreSQLURL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host', full_name='datasource.PostgreSQLURL.host', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port', full_name='datasource.PostgreSQLURL.port', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\262\001\0068\200\200\004@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user', full_name='datasource.PostgreSQLURL.user', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='password', full_name='datasource.PostgreSQLURL.password', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='database', full_name='datasource.PostgreSQLURL.database', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=422,
  serialized_end=608,
)


_CLICKHOUSEURL = _descriptor.Descriptor(
  name='ClickHouseURL',
  full_name='datasource.ClickHouseURL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host', full_name='datasource.ClickHouseURL.host', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port', full_name='datasource.ClickHouseURL.port', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\262\001\0068\200\200\004@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user', full_name='datasource.ClickHouseURL.user', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='password', full_name='datasource.ClickHouseURL.password', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='database', full_name='datasource.ClickHouseURL.database', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=611,
  serialized_end=797,
)


_FTPURL = _descriptor.Descriptor(
  name='FtpURL',
  full_name='datasource.FtpURL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='protocol', full_name='datasource.FtpURL.protocol', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='connection_mode', full_name='datasource.FtpURL.connection_mode', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\023\n\021\n\010protocol\022\005\332\001\002\030\001\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='private_key', full_name='datasource.FtpURL.private_key', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\023\n\021\n\010protocol\022\005\332\001\002\030\002\342\337\037\017\022\r\302\001\n\200\002\001\230\002\200\020\210\005\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='host', full_name='datasource.FtpURL.host', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port', full_name='datasource.FtpURL.port', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\262\001\0068\200\200\004@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user', full_name='datasource.FtpURL.user', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='password', full_name='datasource.FtpURL.password', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _FTPURL_PROTOCOL,
    _FTPURL_CONNECTIONMODE,
  ],
  serialized_options=b'\312\262\004\002\n\000',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=800,
  serialized_end=1287,
)


_S3URL = _descriptor.Descriptor(
  name='S3URL',
  full_name='datasource.S3URL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=1289,
  serialized_end=1304,
)


_KAFKAURL = _descriptor.Descriptor(
  name='KafkaURL',
  full_name='datasource.KafkaURL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='kafka_brokers', full_name='datasource.KafkaURL.kafka_brokers', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\n\022\010\352\001\0058\200\001@\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=1306,
  serialized_end=1381,
)


_HBASEURL = _descriptor.Descriptor(
  name='HBaseURL',
  full_name='datasource.HBaseURL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='config', full_name='datasource.HBaseURL.config', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\020\022\016\302\001\013\220\002\001\230\002\200\200\001\350\010\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=1383,
  serialized_end=1439,
)


_HDFSURL = _descriptor.Descriptor(
  name='HDFSURL',
  full_name='datasource.HDFSURL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name_node', full_name='datasource.HDFSURL.name_node', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\001\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port', full_name='datasource.HDFSURL.port', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\262\001\0068\200\200\004@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='config', full_name='datasource.HDFSURL.config', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\021\n\017\n\006config\022\005\302\001\002\"\000\342\337\037\r\022\013\302\001\010\230\002\200\200\001\350\010\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=1442,
  serialized_end=1582,
)

_FTPURL.fields_by_name['protocol'].enum_type = _FTPURL_PROTOCOL
_FTPURL.fields_by_name['connection_mode'].enum_type = _FTPURL_CONNECTIONMODE
_FTPURL_PROTOCOL.containing_type = _FTPURL
_FTPURL_CONNECTIONMODE.containing_type = _FTPURL
_KAFKAURL.fields_by_name['kafka_brokers'].message_type = _HOST
DESCRIPTOR.message_types_by_name['Host'] = _HOST
DESCRIPTOR.message_types_by_name['MySQLURL'] = _MYSQLURL
DESCRIPTOR.message_types_by_name['PostgreSQLURL'] = _POSTGRESQLURL
DESCRIPTOR.message_types_by_name['ClickHouseURL'] = _CLICKHOUSEURL
DESCRIPTOR.message_types_by_name['FtpURL'] = _FTPURL
DESCRIPTOR.message_types_by_name['S3URL'] = _S3URL
DESCRIPTOR.message_types_by_name['KafkaURL'] = _KAFKAURL
DESCRIPTOR.message_types_by_name['HBaseURL'] = _HBASEURL
DESCRIPTOR.message_types_by_name['HDFSURL'] = _HDFSURL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Host = _reflection.GeneratedProtocolMessageType('Host', (_message.Message,), {
  'DESCRIPTOR' : _HOST,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.Host)
  })
_sym_db.RegisterMessage(Host)

MySQLURL = _reflection.GeneratedProtocolMessageType('MySQLURL', (_message.Message,), {
  'DESCRIPTOR' : _MYSQLURL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.MySQLURL)
  })
_sym_db.RegisterMessage(MySQLURL)

PostgreSQLURL = _reflection.GeneratedProtocolMessageType('PostgreSQLURL', (_message.Message,), {
  'DESCRIPTOR' : _POSTGRESQLURL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.PostgreSQLURL)
  })
_sym_db.RegisterMessage(PostgreSQLURL)

ClickHouseURL = _reflection.GeneratedProtocolMessageType('ClickHouseURL', (_message.Message,), {
  'DESCRIPTOR' : _CLICKHOUSEURL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.ClickHouseURL)
  })
_sym_db.RegisterMessage(ClickHouseURL)

FtpURL = _reflection.GeneratedProtocolMessageType('FtpURL', (_message.Message,), {
  'DESCRIPTOR' : _FTPURL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.FtpURL)
  })
_sym_db.RegisterMessage(FtpURL)

S3URL = _reflection.GeneratedProtocolMessageType('S3URL', (_message.Message,), {
  'DESCRIPTOR' : _S3URL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.S3URL)
  })
_sym_db.RegisterMessage(S3URL)

KafkaURL = _reflection.GeneratedProtocolMessageType('KafkaURL', (_message.Message,), {
  'DESCRIPTOR' : _KAFKAURL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.KafkaURL)
  })
_sym_db.RegisterMessage(KafkaURL)

HBaseURL = _reflection.GeneratedProtocolMessageType('HBaseURL', (_message.Message,), {
  'DESCRIPTOR' : _HBASEURL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.HBaseURL)
  })
_sym_db.RegisterMessage(HBaseURL)

HDFSURL = _reflection.GeneratedProtocolMessageType('HDFSURL', (_message.Message,), {
  'DESCRIPTOR' : _HDFSURL,
  '__module__' : 'proto.types.model.datasource.datasource_url_pb2'
  # @@protoc_insertion_point(class_scope:datasource.HDFSURL)
  })
_sym_db.RegisterMessage(HDFSURL)


DESCRIPTOR._options = None
_HOST.fields_by_name['host']._options = None
_HOST.fields_by_name['port']._options = None
_MYSQLURL.fields_by_name['host']._options = None
_MYSQLURL.fields_by_name['port']._options = None
_MYSQLURL.fields_by_name['user']._options = None
_MYSQLURL.fields_by_name['password']._options = None
_MYSQLURL.fields_by_name['database']._options = None
_MYSQLURL._options = None
_POSTGRESQLURL.fields_by_name['host']._options = None
_POSTGRESQLURL.fields_by_name['port']._options = None
_POSTGRESQLURL.fields_by_name['user']._options = None
_POSTGRESQLURL.fields_by_name['password']._options = None
_POSTGRESQLURL.fields_by_name['database']._options = None
_POSTGRESQLURL._options = None
_CLICKHOUSEURL.fields_by_name['host']._options = None
_CLICKHOUSEURL.fields_by_name['port']._options = None
_CLICKHOUSEURL.fields_by_name['user']._options = None
_CLICKHOUSEURL.fields_by_name['password']._options = None
_CLICKHOUSEURL.fields_by_name['database']._options = None
_CLICKHOUSEURL._options = None
_FTPURL.fields_by_name['protocol']._options = None
_FTPURL.fields_by_name['connection_mode']._options = None
_FTPURL.fields_by_name['private_key']._options = None
_FTPURL.fields_by_name['host']._options = None
_FTPURL.fields_by_name['port']._options = None
_FTPURL.fields_by_name['user']._options = None
_FTPURL.fields_by_name['password']._options = None
_FTPURL._options = None
_S3URL._options = None
_KAFKAURL.fields_by_name['kafka_brokers']._options = None
_KAFKAURL._options = None
_HBASEURL.fields_by_name['config']._options = None
_HBASEURL._options = None
_HDFSURL.fields_by_name['name_node']._options = None
_HDFSURL.fields_by_name['port']._options = None
_HDFSURL.fields_by_name['config']._options = None
_HDFSURL._options = None
# @@protoc_insertion_point(module_scope)
