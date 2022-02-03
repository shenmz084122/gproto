# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/request/monitor_manage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from github.com.yu31.protoc_plugin.proto import defaults_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_defaults__pb2
from proto.types.model import monitor_pb2 as proto_dot_types_dot_model_dot_monitor__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/request/monitor_manage.proto',
  package='request',
  syntax='proto3',
  serialized_options=b'\n$com.dataomnis.gproto.types.pbrequestB\026PBRequestMonitorManageP\000Z3github.com/DataWorkbench/gproto/xgo/types/pbrequest',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n(proto/types/request/monitor_manage.proto\x12\x07request\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a\x32github.com/yu31/protoc-plugin/proto/defaults.proto\x1a\x1fproto/types/model/monitor.proto\"~\n\x10ListMonitorRules\x12&\n\x05limit\x18\x01 \x01(\x05\x42\x17\xa2\xa1\x1f\x06\xaa\x06\x03\x31\x30\x30\xe2\xdf\x1f\t\x12\x07\xb2\x01\x04\x30\x00\x38\x64\x12\x1b\n\x06offset\x18\x02 \x01(\x05\x42\x0b\xe2\xdf\x1f\x07\x12\x05\xb2\x01\x02@\x00\x12%\n\x08space_id\x18\x03 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\"B\n\x11\x43reateMonitorRule\x12-\n\x04info\x18\x01 \x01(\x0b\x32\x12.model.MonitorRuleB\x0b\xe2\xdf\x1f\x07\x12\x05\xe2\x01\x02\x10\x01\"5\n\x12\x44\x65leteMonitorRules\x12\x1f\n\x08rule_ids\x18\x01 \x03(\tB\r\xe2\xdf\x1f\t\x12\x07\xea\x01\x04\x30\x00\x38\x64\"5\n\x12\x45nableMonitorRules\x12\x1f\n\x08rule_ids\x18\x01 \x03(\tB\r\xe2\xdf\x1f\t\x12\x07\xea\x01\x04\x30\x00\x38\x64\"6\n\x13\x44isableMonitorRules\x12\x1f\n\x08rule_ids\x18\x01 \x03(\tB\r\xe2\xdf\x1f\t\x12\x07\xea\x01\x04\x30\x00\x38\x64\"B\n\x11UpdateMonitorRule\x12-\n\x04info\x18\x01 \x01(\x0b\x32\x12.model.MonitorRuleB\x0b\xe2\xdf\x1f\x07\x12\x05\xe2\x01\x02\x10\x01\"4\n\x13\x44\x65scribeMonitorRule\x12\x1d\n\x07rule_id\x18\x01 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x42u\n$com.dataomnis.gproto.types.pbrequestB\x16PBRequestMonitorManageP\x00Z3github.com/DataWorkbench/gproto/xgo/types/pbrequestb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_defaults__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_monitor__pb2.DESCRIPTOR,])




_LISTMONITORRULES = _descriptor.Descriptor(
  name='ListMonitorRules',
  full_name='request.ListMonitorRules',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='request.ListMonitorRules.limit', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\242\241\037\006\252\006\003100\342\337\037\t\022\007\262\001\0040\0008d', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset', full_name='request.ListMonitorRules.offset', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\262\001\002@\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='space_id', full_name='request.ListMonitorRules.space_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=191,
  serialized_end=317,
)


_CREATEMONITORRULE = _descriptor.Descriptor(
  name='CreateMonitorRule',
  full_name='request.CreateMonitorRule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='info', full_name='request.CreateMonitorRule.info', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\342\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=319,
  serialized_end=385,
)


_DELETEMONITORRULES = _descriptor.Descriptor(
  name='DeleteMonitorRules',
  full_name='request.DeleteMonitorRules',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='rule_ids', full_name='request.DeleteMonitorRules.rule_ids', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\352\001\0040\0008d', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=387,
  serialized_end=440,
)


_ENABLEMONITORRULES = _descriptor.Descriptor(
  name='EnableMonitorRules',
  full_name='request.EnableMonitorRules',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='rule_ids', full_name='request.EnableMonitorRules.rule_ids', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\352\001\0040\0008d', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=442,
  serialized_end=495,
)


_DISABLEMONITORRULES = _descriptor.Descriptor(
  name='DisableMonitorRules',
  full_name='request.DisableMonitorRules',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='rule_ids', full_name='request.DisableMonitorRules.rule_ids', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\352\001\0040\0008d', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=497,
  serialized_end=551,
)


_UPDATEMONITORRULE = _descriptor.Descriptor(
  name='UpdateMonitorRule',
  full_name='request.UpdateMonitorRule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='info', full_name='request.UpdateMonitorRule.info', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\342\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=553,
  serialized_end=619,
)


_DESCRIBEMONITORRULE = _descriptor.Descriptor(
  name='DescribeMonitorRule',
  full_name='request.DescribeMonitorRule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='rule_id', full_name='request.DescribeMonitorRule.rule_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=621,
  serialized_end=673,
)

_CREATEMONITORRULE.fields_by_name['info'].message_type = proto_dot_types_dot_model_dot_monitor__pb2._MONITORRULE
_UPDATEMONITORRULE.fields_by_name['info'].message_type = proto_dot_types_dot_model_dot_monitor__pb2._MONITORRULE
DESCRIPTOR.message_types_by_name['ListMonitorRules'] = _LISTMONITORRULES
DESCRIPTOR.message_types_by_name['CreateMonitorRule'] = _CREATEMONITORRULE
DESCRIPTOR.message_types_by_name['DeleteMonitorRules'] = _DELETEMONITORRULES
DESCRIPTOR.message_types_by_name['EnableMonitorRules'] = _ENABLEMONITORRULES
DESCRIPTOR.message_types_by_name['DisableMonitorRules'] = _DISABLEMONITORRULES
DESCRIPTOR.message_types_by_name['UpdateMonitorRule'] = _UPDATEMONITORRULE
DESCRIPTOR.message_types_by_name['DescribeMonitorRule'] = _DESCRIBEMONITORRULE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListMonitorRules = _reflection.GeneratedProtocolMessageType('ListMonitorRules', (_message.Message,), {
  'DESCRIPTOR' : _LISTMONITORRULES,
  '__module__' : 'proto.types.request.monitor_manage_pb2'
  # @@protoc_insertion_point(class_scope:request.ListMonitorRules)
  })
_sym_db.RegisterMessage(ListMonitorRules)

CreateMonitorRule = _reflection.GeneratedProtocolMessageType('CreateMonitorRule', (_message.Message,), {
  'DESCRIPTOR' : _CREATEMONITORRULE,
  '__module__' : 'proto.types.request.monitor_manage_pb2'
  # @@protoc_insertion_point(class_scope:request.CreateMonitorRule)
  })
_sym_db.RegisterMessage(CreateMonitorRule)

DeleteMonitorRules = _reflection.GeneratedProtocolMessageType('DeleteMonitorRules', (_message.Message,), {
  'DESCRIPTOR' : _DELETEMONITORRULES,
  '__module__' : 'proto.types.request.monitor_manage_pb2'
  # @@protoc_insertion_point(class_scope:request.DeleteMonitorRules)
  })
_sym_db.RegisterMessage(DeleteMonitorRules)

EnableMonitorRules = _reflection.GeneratedProtocolMessageType('EnableMonitorRules', (_message.Message,), {
  'DESCRIPTOR' : _ENABLEMONITORRULES,
  '__module__' : 'proto.types.request.monitor_manage_pb2'
  # @@protoc_insertion_point(class_scope:request.EnableMonitorRules)
  })
_sym_db.RegisterMessage(EnableMonitorRules)

DisableMonitorRules = _reflection.GeneratedProtocolMessageType('DisableMonitorRules', (_message.Message,), {
  'DESCRIPTOR' : _DISABLEMONITORRULES,
  '__module__' : 'proto.types.request.monitor_manage_pb2'
  # @@protoc_insertion_point(class_scope:request.DisableMonitorRules)
  })
_sym_db.RegisterMessage(DisableMonitorRules)

UpdateMonitorRule = _reflection.GeneratedProtocolMessageType('UpdateMonitorRule', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEMONITORRULE,
  '__module__' : 'proto.types.request.monitor_manage_pb2'
  # @@protoc_insertion_point(class_scope:request.UpdateMonitorRule)
  })
_sym_db.RegisterMessage(UpdateMonitorRule)

DescribeMonitorRule = _reflection.GeneratedProtocolMessageType('DescribeMonitorRule', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEMONITORRULE,
  '__module__' : 'proto.types.request.monitor_manage_pb2'
  # @@protoc_insertion_point(class_scope:request.DescribeMonitorRule)
  })
_sym_db.RegisterMessage(DescribeMonitorRule)


DESCRIPTOR._options = None
_LISTMONITORRULES.fields_by_name['limit']._options = None
_LISTMONITORRULES.fields_by_name['offset']._options = None
_LISTMONITORRULES.fields_by_name['space_id']._options = None
_CREATEMONITORRULE.fields_by_name['info']._options = None
_DELETEMONITORRULES.fields_by_name['rule_ids']._options = None
_ENABLEMONITORRULES.fields_by_name['rule_ids']._options = None
_DISABLEMONITORRULES.fields_by_name['rule_ids']._options = None
_UPDATEMONITORRULE.fields_by_name['info']._options = None
_DESCRIBEMONITORRULE.fields_by_name['rule_id']._options = None
# @@protoc_insertion_point(module_scope)
