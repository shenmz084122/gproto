# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/response/account.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model import account_pb2 as proto_dot_types_dot_model_dot_account__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/response/account.proto',
  package='response',
  syntax='proto3',
  serialized_options=b'\n%com.dataomnis.gproto.types.pbresponseB\021PBResponseAccountP\000Z4github.com/DataWorkbench/gproto/xgo/types/pbresponse',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\"proto/types/response/account.proto\x12\x08response\x1a\x1fproto/types/model/account.proto\"=\n\x11\x44\x65scribeAccessKey\x12\r\n\x05owner\x18\x01 \x01(\t\x12\x19\n\x11secret_access_key\x18\x02 \x01(\t\"d\n\rDescribeUsers\x12\x1d\n\x08user_set\x18\x01 \x03(\x0b\x32\x0b.model.User\x12\x13\n\x0btotal_count\x18\x02 \x01(\x03\x12\x0e\n\x06status\x18\x03 \x01(\x05\x12\x0f\n\x07message\x18\x04 \x01(\t\"L\n\x18ValidateRequestSignature\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\x0f\n\x07user_id\x18\x03 \x01(\t\"\'\n\nCreateUser\x12\x19\n\x04user\x18\x01 \x01(\x0b\x32\x0b.model.User\"\'\n\nUpdateUser\x12\x19\n\x04user\x18\x01 \x01(\x0b\x32\x0b.model.User\"\x1d\n\nDeleteUser\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"Q\n\x0c\x43heckSession\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x15\n\raccess_key_id\x18\x02 \x01(\t\x12\x19\n\x11secret_access_key\x18\x03 \x01(\t\" \n\rCreateSession\x12\x0f\n\x07session\x18\x01 \x01(\t\"?\n\x0bGetUserRole\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04role\x18\x02 \x01(\t\x12\x11\n\tprivilege\x18\x03 \x01(\x05\x42r\n%com.dataomnis.gproto.types.pbresponseB\x11PBResponseAccountP\x00Z4github.com/DataWorkbench/gproto/xgo/types/pbresponseb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_account__pb2.DESCRIPTOR,])




_DESCRIBEACCESSKEY = _descriptor.Descriptor(
  name='DescribeAccessKey',
  full_name='response.DescribeAccessKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner', full_name='response.DescribeAccessKey.owner', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='secret_access_key', full_name='response.DescribeAccessKey.secret_access_key', index=1,
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
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=81,
  serialized_end=142,
)


_DESCRIBEUSERS = _descriptor.Descriptor(
  name='DescribeUsers',
  full_name='response.DescribeUsers',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_set', full_name='response.DescribeUsers.user_set', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total_count', full_name='response.DescribeUsers.total_count', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='response.DescribeUsers.status', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.DescribeUsers.message', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=144,
  serialized_end=244,
)


_VALIDATEREQUESTSIGNATURE = _descriptor.Descriptor(
  name='ValidateRequestSignature',
  full_name='response.ValidateRequestSignature',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='response.ValidateRequestSignature.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='response.ValidateRequestSignature.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='response.ValidateRequestSignature.user_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=246,
  serialized_end=322,
)


_CREATEUSER = _descriptor.Descriptor(
  name='CreateUser',
  full_name='response.CreateUser',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='response.CreateUser.user', index=0,
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
  serialized_start=324,
  serialized_end=363,
)


_UPDATEUSER = _descriptor.Descriptor(
  name='UpdateUser',
  full_name='response.UpdateUser',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='response.UpdateUser.user', index=0,
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
  serialized_start=365,
  serialized_end=404,
)


_DELETEUSER = _descriptor.Descriptor(
  name='DeleteUser',
  full_name='response.DeleteUser',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='response.DeleteUser.user_id', index=0,
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
  serialized_start=406,
  serialized_end=435,
)


_CHECKSESSION = _descriptor.Descriptor(
  name='CheckSession',
  full_name='response.CheckSession',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='response.CheckSession.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='access_key_id', full_name='response.CheckSession.access_key_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='secret_access_key', full_name='response.CheckSession.secret_access_key', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=437,
  serialized_end=518,
)


_CREATESESSION = _descriptor.Descriptor(
  name='CreateSession',
  full_name='response.CreateSession',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='session', full_name='response.CreateSession.session', index=0,
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
  serialized_start=520,
  serialized_end=552,
)


_GETUSERROLE = _descriptor.Descriptor(
  name='GetUserRole',
  full_name='response.GetUserRole',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='response.GetUserRole.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='role', full_name='response.GetUserRole.role', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='privilege', full_name='response.GetUserRole.privilege', index=2,
      number=3, type=5, cpp_type=1, label=1,
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
  serialized_start=554,
  serialized_end=617,
)

_DESCRIBEUSERS.fields_by_name['user_set'].message_type = proto_dot_types_dot_model_dot_account__pb2._USER
_CREATEUSER.fields_by_name['user'].message_type = proto_dot_types_dot_model_dot_account__pb2._USER
_UPDATEUSER.fields_by_name['user'].message_type = proto_dot_types_dot_model_dot_account__pb2._USER
DESCRIPTOR.message_types_by_name['DescribeAccessKey'] = _DESCRIBEACCESSKEY
DESCRIPTOR.message_types_by_name['DescribeUsers'] = _DESCRIBEUSERS
DESCRIPTOR.message_types_by_name['ValidateRequestSignature'] = _VALIDATEREQUESTSIGNATURE
DESCRIPTOR.message_types_by_name['CreateUser'] = _CREATEUSER
DESCRIPTOR.message_types_by_name['UpdateUser'] = _UPDATEUSER
DESCRIPTOR.message_types_by_name['DeleteUser'] = _DELETEUSER
DESCRIPTOR.message_types_by_name['CheckSession'] = _CHECKSESSION
DESCRIPTOR.message_types_by_name['CreateSession'] = _CREATESESSION
DESCRIPTOR.message_types_by_name['GetUserRole'] = _GETUSERROLE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DescribeAccessKey = _reflection.GeneratedProtocolMessageType('DescribeAccessKey', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEACCESSKEY,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.DescribeAccessKey)
  })
_sym_db.RegisterMessage(DescribeAccessKey)

DescribeUsers = _reflection.GeneratedProtocolMessageType('DescribeUsers', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEUSERS,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.DescribeUsers)
  })
_sym_db.RegisterMessage(DescribeUsers)

ValidateRequestSignature = _reflection.GeneratedProtocolMessageType('ValidateRequestSignature', (_message.Message,), {
  'DESCRIPTOR' : _VALIDATEREQUESTSIGNATURE,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.ValidateRequestSignature)
  })
_sym_db.RegisterMessage(ValidateRequestSignature)

CreateUser = _reflection.GeneratedProtocolMessageType('CreateUser', (_message.Message,), {
  'DESCRIPTOR' : _CREATEUSER,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.CreateUser)
  })
_sym_db.RegisterMessage(CreateUser)

UpdateUser = _reflection.GeneratedProtocolMessageType('UpdateUser', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEUSER,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.UpdateUser)
  })
_sym_db.RegisterMessage(UpdateUser)

DeleteUser = _reflection.GeneratedProtocolMessageType('DeleteUser', (_message.Message,), {
  'DESCRIPTOR' : _DELETEUSER,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.DeleteUser)
  })
_sym_db.RegisterMessage(DeleteUser)

CheckSession = _reflection.GeneratedProtocolMessageType('CheckSession', (_message.Message,), {
  'DESCRIPTOR' : _CHECKSESSION,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.CheckSession)
  })
_sym_db.RegisterMessage(CheckSession)

CreateSession = _reflection.GeneratedProtocolMessageType('CreateSession', (_message.Message,), {
  'DESCRIPTOR' : _CREATESESSION,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.CreateSession)
  })
_sym_db.RegisterMessage(CreateSession)

GetUserRole = _reflection.GeneratedProtocolMessageType('GetUserRole', (_message.Message,), {
  'DESCRIPTOR' : _GETUSERROLE,
  '__module__' : 'proto.types.response.account_pb2'
  # @@protoc_insertion_point(class_scope:response.GetUserRole)
  })
_sym_db.RegisterMessage(GetUserRole)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
