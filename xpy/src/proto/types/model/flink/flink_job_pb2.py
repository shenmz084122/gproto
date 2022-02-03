# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/flink/flink_job.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2
from proto.types.model.flink import flink_operator_pb2 as proto_dot_types_dot_model_dot_flink_dot_flink__operator__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/flink/flink_job.proto',
  package='flink',
  syntax='proto3',
  serialized_options=b'\n*com.dataomnis.gproto.types.pbmodel.pbflinkB\nPBFlinkJobP\000Z9github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbflink',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\'proto/types/model/flink/flink_job.proto\x12\x05\x66link\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\x1a,proto/types/model/flink/flink_operator.proto\"\xc2\x04\n\rFlinkOperator\x12\x36\n\x04type\x18\x01 \x01(\x0e\x32\x19.flink.FlinkOperator.TypeB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12\x18\n\x02id\x18\x02 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x12\x1a\n\x04name\x18\x03 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\x88\x02\x41\x12\x1e\n\x08upstream\x18\x04 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x12$\n\x0eupstream_right\x18\x05 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x12!\n\x0b\x64own_stream\x18\x06 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x15\x12\x1f\n\x07point_x\x18\x07 \x01(\x05\x42\x0e\xe2\xdf\x1f\n\x12\x08\xb2\x01\x05\x38\xc8\x01@\x01\x12\x1f\n\x07point_y\x18\x08 \x01(\x05\x42\x0e\xe2\xdf\x1f\n\x12\x08\xb2\x01\x05\x38\xc8\x01@\x01\x12/\n\x08property\x18\t \x01(\x0b\x32\x17.flink.OperatorPropertyB\x04\xe2\xdf\x1f\x00\"\xe6\x01\n\x04Type\x12\t\n\x05\x45mpty\x10\x00\x12\n\n\x06Values\x10\x01\x12\t\n\x05\x43onst\x10\x02\x12\n\n\x06Source\x10\x03\x12\r\n\tDimension\x10\x04\x12\x08\n\x04\x44\x65st\x10\x05\x12\x0b\n\x07OrderBy\x10\x06\x12\t\n\x05Limit\x10\x07\x12\n\n\x06Offset\x10\x08\x12\t\n\x05\x46\x65tch\x10\t\x12\n\n\x06\x46ilter\x10\n\x12\t\n\x05Union\x10\x0b\x12\n\n\x06\x45xcept\x10\x0c\x12\r\n\tIntersect\x10\r\x12\x0b\n\x07GroupBy\x10\x0e\x12\n\n\x06Having\x10\x0f\x12\x08\n\x04Join\x10\x10\x12\x08\n\x04UDTF\x10\x11\x12\t\n\x05UDTTF\x10\x12\"p\n\x08\x46linkJar\x12!\n\x0bresource_id\x18\x01 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x12\x1f\n\x08jar_args\x18\x02 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\x98\x02\x80\x08\x12 \n\tjar_entry\x18\x03 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\x98\x02\x80\x08\"-\n\nFlinkScala\x12\x1f\n\x04\x63ode\x18\x01 \x01(\tB\x11\xe2\xdf\x1f\r\x12\x0b\xc2\x01\x08\x98\x02\xc0\xb8\x02\x88\x05\x01\".\n\x0b\x46linkPython\x12\x1f\n\x04\x63ode\x18\x01 \x01(\tB\x11\xe2\xdf\x1f\r\x12\x0b\xc2\x01\x08\x98\x02\xc0\xb8\x02\x88\x05\x01\"+\n\x08\x46linkSQL\x12\x1f\n\x04\x63ode\x18\x01 \x01(\tB\x11\xe2\xdf\x1f\r\x12\x0b\xc2\x01\x08\x98\x02\xc0\xb8\x02\x88\x05\x01\x42u\n*com.dataomnis.gproto.types.pbmodel.pbflinkB\nPBFlinkJobP\x00Z9github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbflinkb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,proto_dot_types_dot_model_dot_flink_dot_flink__operator__pb2.DESCRIPTOR,])



_FLINKOPERATOR_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='flink.FlinkOperator.Type',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Empty', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Values', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Const', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Source', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Dimension', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Dest', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OrderBy', index=6, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Limit', index=7, number=7,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Offset', index=8, number=8,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Fetch', index=9, number=9,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Filter', index=10, number=10,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Union', index=11, number=11,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Except', index=12, number=12,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Intersect', index=13, number=13,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='GroupBy', index=14, number=14,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Having', index=15, number=15,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Join', index=16, number=16,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UDTF', index=17, number=17,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UDTTF', index=18, number=18,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=498,
  serialized_end=728,
)
_sym_db.RegisterEnumDescriptor(_FLINKOPERATOR_TYPE)


_FLINKOPERATOR = _descriptor.Descriptor(
  name='FlinkOperator',
  full_name='flink.FlinkOperator',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='flink.FlinkOperator.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='flink.FlinkOperator.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='flink.FlinkOperator.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\210\002A', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='upstream', full_name='flink.FlinkOperator.upstream', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='upstream_right', full_name='flink.FlinkOperator.upstream_right', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='down_stream', full_name='flink.FlinkOperator.down_stream', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\025', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='point_x', full_name='flink.FlinkOperator.point_x', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\n\022\010\262\001\0058\310\001@\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='point_y', full_name='flink.FlinkOperator.point_y', index=7,
      number=8, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\n\022\010\262\001\0058\310\001@\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='property', full_name='flink.FlinkOperator.property', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _FLINKOPERATOR_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=150,
  serialized_end=728,
)


_FLINKJAR = _descriptor.Descriptor(
  name='FlinkJar',
  full_name='flink.FlinkJar',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='resource_id', full_name='flink.FlinkJar.resource_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='jar_args', full_name='flink.FlinkJar.jar_args', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\230\002\200\010', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='jar_entry', full_name='flink.FlinkJar.jar_entry', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\230\002\200\010', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=730,
  serialized_end=842,
)


_FLINKSCALA = _descriptor.Descriptor(
  name='FlinkScala',
  full_name='flink.FlinkScala',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='flink.FlinkScala.code', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\r\022\013\302\001\010\230\002\300\270\002\210\005\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=844,
  serialized_end=889,
)


_FLINKPYTHON = _descriptor.Descriptor(
  name='FlinkPython',
  full_name='flink.FlinkPython',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='flink.FlinkPython.code', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\r\022\013\302\001\010\230\002\300\270\002\210\005\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=891,
  serialized_end=937,
)


_FLINKSQL = _descriptor.Descriptor(
  name='FlinkSQL',
  full_name='flink.FlinkSQL',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='flink.FlinkSQL.code', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\r\022\013\302\001\010\230\002\300\270\002\210\005\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=939,
  serialized_end=982,
)

_FLINKOPERATOR.fields_by_name['type'].enum_type = _FLINKOPERATOR_TYPE
_FLINKOPERATOR.fields_by_name['property'].message_type = proto_dot_types_dot_model_dot_flink_dot_flink__operator__pb2._OPERATORPROPERTY
_FLINKOPERATOR_TYPE.containing_type = _FLINKOPERATOR
DESCRIPTOR.message_types_by_name['FlinkOperator'] = _FLINKOPERATOR
DESCRIPTOR.message_types_by_name['FlinkJar'] = _FLINKJAR
DESCRIPTOR.message_types_by_name['FlinkScala'] = _FLINKSCALA
DESCRIPTOR.message_types_by_name['FlinkPython'] = _FLINKPYTHON
DESCRIPTOR.message_types_by_name['FlinkSQL'] = _FLINKSQL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FlinkOperator = _reflection.GeneratedProtocolMessageType('FlinkOperator', (_message.Message,), {
  'DESCRIPTOR' : _FLINKOPERATOR,
  '__module__' : 'proto.types.model.flink.flink_job_pb2'
  # @@protoc_insertion_point(class_scope:flink.FlinkOperator)
  })
_sym_db.RegisterMessage(FlinkOperator)

FlinkJar = _reflection.GeneratedProtocolMessageType('FlinkJar', (_message.Message,), {
  'DESCRIPTOR' : _FLINKJAR,
  '__module__' : 'proto.types.model.flink.flink_job_pb2'
  # @@protoc_insertion_point(class_scope:flink.FlinkJar)
  })
_sym_db.RegisterMessage(FlinkJar)

FlinkScala = _reflection.GeneratedProtocolMessageType('FlinkScala', (_message.Message,), {
  'DESCRIPTOR' : _FLINKSCALA,
  '__module__' : 'proto.types.model.flink.flink_job_pb2'
  # @@protoc_insertion_point(class_scope:flink.FlinkScala)
  })
_sym_db.RegisterMessage(FlinkScala)

FlinkPython = _reflection.GeneratedProtocolMessageType('FlinkPython', (_message.Message,), {
  'DESCRIPTOR' : _FLINKPYTHON,
  '__module__' : 'proto.types.model.flink.flink_job_pb2'
  # @@protoc_insertion_point(class_scope:flink.FlinkPython)
  })
_sym_db.RegisterMessage(FlinkPython)

FlinkSQL = _reflection.GeneratedProtocolMessageType('FlinkSQL', (_message.Message,), {
  'DESCRIPTOR' : _FLINKSQL,
  '__module__' : 'proto.types.model.flink.flink_job_pb2'
  # @@protoc_insertion_point(class_scope:flink.FlinkSQL)
  })
_sym_db.RegisterMessage(FlinkSQL)


DESCRIPTOR._options = None
_FLINKOPERATOR.fields_by_name['type']._options = None
_FLINKOPERATOR.fields_by_name['id']._options = None
_FLINKOPERATOR.fields_by_name['name']._options = None
_FLINKOPERATOR.fields_by_name['upstream']._options = None
_FLINKOPERATOR.fields_by_name['upstream_right']._options = None
_FLINKOPERATOR.fields_by_name['down_stream']._options = None
_FLINKOPERATOR.fields_by_name['point_x']._options = None
_FLINKOPERATOR.fields_by_name['point_y']._options = None
_FLINKOPERATOR.fields_by_name['property']._options = None
_FLINKJAR.fields_by_name['resource_id']._options = None
_FLINKJAR.fields_by_name['jar_args']._options = None
_FLINKJAR.fields_by_name['jar_entry']._options = None
_FLINKSCALA.fields_by_name['code']._options = None
_FLINKPYTHON.fields_by_name['code']._options = None
_FLINKSQL.fields_by_name['code']._options = None
# @@protoc_insertion_point(module_scope)
