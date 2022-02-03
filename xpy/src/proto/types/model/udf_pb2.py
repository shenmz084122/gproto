# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/types/model/udf.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.yu31.protoc_plugin.proto import validator_pb2 as github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/types/model/udf.proto',
  package='model',
  syntax='proto3',
  serialized_options=b'\n\"com.dataomnis.gproto.types.pbmodelB\nPBModelUDFP\000Z1github.com/DataWorkbench/gproto/xgo/types/pbmodel',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1bproto/types/model/udf.proto\x12\x05model\x1a\x33github.com/yu31/protoc-plugin/proto/validator.proto\"\x91\x05\n\x07UDFInfo\x12\x1c\n\x06udf_id\x18\x01 \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\xf0\x01\x14\x12%\n\x08space_id\x18\x02 \x01(\tB\x13\xe2\xdf\x1f\x0f\x12\r\xc2\x01\n\xf0\x01\x14\xca\x02\x04wks-\x12\x34\n\x08udf_type\x18\x03 \x01(\x0e\x32\x13.model.UDFInfo.TypeB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12<\n\x0cudf_language\x18\x04 \x01(\x0e\x32\x17.model.UDFInfo.LanguageB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12\x1d\n\x04name\x18\x05 \x01(\tB\x0f\xe2\xdf\x1f\x0b\x12\t\xc2\x01\x06\x90\x02\x02\x98\x02@\x12\x1e\n\x07\x63omment\x18\x06 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\xc8\x01\x80\x02\x12\x1e\n\x06\x64\x65\x66ine\x18\x07 \x01(\tB\x0e\xe2\xdf\x1f\n\x12\x08\xc2\x01\x05\xc8\x01\xa0\x9c\x01\x12#\n\x0cusage_sample\x18\x08 \x01(\tB\r\xe2\xdf\x1f\t\x12\x07\xc2\x01\x04\xc8\x01\xd0\x0f\x12\x1c\n\x07\x63reated\x18\t \x01(\x03\x42\x0b\xe2\xdf\x1f\x07\x12\x05\xb2\x01\x02\x30\x00\x12\x1c\n\x07updated\x18\n \x01(\x03\x42\x0b\xe2\xdf\x1f\x07\x12\x05\xb2\x01\x02\x30\x00\x12\x34\n\x06status\x18\x0b \x01(\x0e\x32\x15.model.UDFInfo.StatusB\r\xe2\xdf\x1f\t\x12\x07\xda\x01\x04\x30\x00X\x01\x12\x1f\n\tcreate_by\x18\x0c \x01(\tB\x0c\xe2\xdf\x1f\x08\x12\x06\xc2\x01\x03\x98\x02@\"3\n\x04Type\x12\r\n\tTypeUnset\x10\x00\x12\x07\n\x03UDF\x10\x01\x12\x08\n\x04UDTF\x10\x02\x12\t\n\x05UDTTF\x10\x03\">\n\x08Language\x12\x11\n\rLanguageUnset\x10\x00\x12\t\n\x05Scala\x10\x01\x12\x08\n\x04Java\x10\x02\x12\n\n\x06Python\x10\x03\"A\n\x06Status\x12\x0f\n\x0bStatusUnset\x10\x00\x12\x0b\n\x07\x45nabled\x10\x01\x12\x0c\n\x08\x44isabled\x10\x02\x12\x0b\n\x07\x44\x65leted\x10\x03\x42\x65\n\"com.dataomnis.gproto.types.pbmodelB\nPBModelUDFP\x00Z1github.com/DataWorkbench/gproto/xgo/types/pbmodelb\x06proto3'
  ,
  dependencies=[github_dot_com_dot_yu31_dot_protoc__plugin_dot_proto_dot_validator__pb2.DESCRIPTOR,])



_UDFINFO_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='model.UDFInfo.Type',
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
      name='UDF', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UDTF', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UDTTF', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=567,
  serialized_end=618,
)
_sym_db.RegisterEnumDescriptor(_UDFINFO_TYPE)

_UDFINFO_LANGUAGE = _descriptor.EnumDescriptor(
  name='Language',
  full_name='model.UDFInfo.Language',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='LanguageUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Scala', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Java', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Python', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=620,
  serialized_end=682,
)
_sym_db.RegisterEnumDescriptor(_UDFINFO_LANGUAGE)

_UDFINFO_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='model.UDFInfo.Status',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='StatusUnset', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Enabled', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Disabled', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='Deleted', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=684,
  serialized_end=749,
)
_sym_db.RegisterEnumDescriptor(_UDFINFO_STATUS)


_UDFINFO = _descriptor.Descriptor(
  name='UDFInfo',
  full_name='model.UDFInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='udf_id', full_name='model.UDFInfo.udf_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\360\001\024', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='space_id', full_name='model.UDFInfo.space_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\017\022\r\302\001\n\360\001\024\312\002\004wks-', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='udf_type', full_name='model.UDFInfo.udf_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='udf_language', full_name='model.UDFInfo.udf_language', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='model.UDFInfo.name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\013\022\t\302\001\006\220\002\002\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='comment', full_name='model.UDFInfo.comment', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\310\001\200\002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='define', full_name='model.UDFInfo.define', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\n\022\010\302\001\005\310\001\240\234\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='usage_sample', full_name='model.UDFInfo.usage_sample', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\302\001\004\310\001\320\017', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created', full_name='model.UDFInfo.created', index=8,
      number=9, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\262\001\0020\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated', full_name='model.UDFInfo.updated', index=9,
      number=10, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\007\022\005\262\001\0020\000', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='model.UDFInfo.status', index=10,
      number=11, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\t\022\007\332\001\0040\000X\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='create_by', full_name='model.UDFInfo.create_by', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\342\337\037\010\022\006\302\001\003\230\002@', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _UDFINFO_TYPE,
    _UDFINFO_LANGUAGE,
    _UDFINFO_STATUS,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=92,
  serialized_end=749,
)

_UDFINFO.fields_by_name['udf_type'].enum_type = _UDFINFO_TYPE
_UDFINFO.fields_by_name['udf_language'].enum_type = _UDFINFO_LANGUAGE
_UDFINFO.fields_by_name['status'].enum_type = _UDFINFO_STATUS
_UDFINFO_TYPE.containing_type = _UDFINFO
_UDFINFO_LANGUAGE.containing_type = _UDFINFO
_UDFINFO_STATUS.containing_type = _UDFINFO
DESCRIPTOR.message_types_by_name['UDFInfo'] = _UDFINFO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UDFInfo = _reflection.GeneratedProtocolMessageType('UDFInfo', (_message.Message,), {
  'DESCRIPTOR' : _UDFINFO,
  '__module__' : 'proto.types.model.udf_pb2'
  # @@protoc_insertion_point(class_scope:model.UDFInfo)
  })
_sym_db.RegisterMessage(UDFInfo)


DESCRIPTOR._options = None
_UDFINFO.fields_by_name['udf_id']._options = None
_UDFINFO.fields_by_name['space_id']._options = None
_UDFINFO.fields_by_name['udf_type']._options = None
_UDFINFO.fields_by_name['udf_language']._options = None
_UDFINFO.fields_by_name['name']._options = None
_UDFINFO.fields_by_name['comment']._options = None
_UDFINFO.fields_by_name['define']._options = None
_UDFINFO.fields_by_name['usage_sample']._options = None
_UDFINFO.fields_by_name['created']._options = None
_UDFINFO.fields_by_name['updated']._options = None
_UDFINFO.fields_by_name['status']._options = None
_UDFINFO.fields_by_name['create_by']._options = None
# @@protoc_insertion_point(module_scope)
