# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/service/billing/billing.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.types.model import empty_pb2 as proto_dot_types_dot_model_dot_empty__pb2
from proto.types.request import billing_pb2 as proto_dot_types_dot_request_dot_billing__pb2
from proto.types.response import billing_pb2 as proto_dot_types_dot_response_dot_billing__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/service/billing/billing.proto',
  package='billing',
  syntax='proto3',
  serialized_options=b'\n)com.dataomnis.gproto.service.pbsvcbillingB\014PBSvcBillingP\000Z8github.com/DataWorkbench/gproto/xgo/service/pbsvcbilling',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n#proto/service/billing/billing.proto\x12\x07\x62illing\x1a\x1dproto/types/model/empty.proto\x1a!proto/types/request/billing.proto\x1a\"proto/types/response/billing.proto2\xee\x02\n\x07\x42illing\x12O\n\x19GetBillingPriceComponents\x12\x12.model.EmptyStruct\x1a\x1c.response.GetPriceComponents\"\x00\x12Z\n\x15\x43reateBillingInstance\x12\x1e.request.CreateBillingInstance\x1a\x1f.response.CreateBillingInstance\"\x00\x12T\n\x13StopBillingInstance\x12\x1c.request.StopBillingInstance\x1a\x1d.response.StopBillingInstance\"\x00\x12`\n\x17RecoveryBillingInstance\x12 .request.RecoveryBillingInstance\x1a!.response.RecoveryBillingInstance\"\x00\x42u\n)com.dataomnis.gproto.service.pbsvcbillingB\x0cPBSvcBillingP\x00Z8github.com/DataWorkbench/gproto/xgo/service/pbsvcbillingb\x06proto3'
  ,
  dependencies=[proto_dot_types_dot_model_dot_empty__pb2.DESCRIPTOR,proto_dot_types_dot_request_dot_billing__pb2.DESCRIPTOR,proto_dot_types_dot_response_dot_billing__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_BILLING = _descriptor.ServiceDescriptor(
  name='Billing',
  full_name='billing.Billing',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=151,
  serialized_end=517,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetBillingPriceComponents',
    full_name='billing.Billing.GetBillingPriceComponents',
    index=0,
    containing_service=None,
    input_type=proto_dot_types_dot_model_dot_empty__pb2._EMPTYSTRUCT,
    output_type=proto_dot_types_dot_response_dot_billing__pb2._GETPRICECOMPONENTS,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateBillingInstance',
    full_name='billing.Billing.CreateBillingInstance',
    index=1,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_billing__pb2._CREATEBILLINGINSTANCE,
    output_type=proto_dot_types_dot_response_dot_billing__pb2._CREATEBILLINGINSTANCE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='StopBillingInstance',
    full_name='billing.Billing.StopBillingInstance',
    index=2,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_billing__pb2._STOPBILLINGINSTANCE,
    output_type=proto_dot_types_dot_response_dot_billing__pb2._STOPBILLINGINSTANCE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RecoveryBillingInstance',
    full_name='billing.Billing.RecoveryBillingInstance',
    index=3,
    containing_service=None,
    input_type=proto_dot_types_dot_request_dot_billing__pb2._RECOVERYBILLINGINSTANCE,
    output_type=proto_dot_types_dot_response_dot_billing__pb2._RECOVERYBILLINGINSTANCE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_BILLING)

DESCRIPTOR.services_by_name['Billing'] = _BILLING

# @@protoc_insertion_point(module_scope)