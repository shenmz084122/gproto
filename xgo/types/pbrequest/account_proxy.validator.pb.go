// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/request/account_proxy.proto

package pbrequest

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	strconv "strconv"
)

// Set default value for message request.DescribeAccessKeyByProxy
func (this *DescribeAccessKeyByProxy) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}

func (this *ListUsersByProxy) _xxx_xxx_Validator_Validate_limit() error {
	if !(this.Limit > 0) {
		return protovalidator.FieldError1("ListUsersByProxy", "the value of field 'limit' must be greater than '0'", protovalidator.Int32ToString(this.Limit))
	}
	if !(this.Limit <= 100) {
		return protovalidator.FieldError1("ListUsersByProxy", "the value of field 'limit' must be less than or equal to '100'", protovalidator.Int32ToString(this.Limit))
	}
	return nil
}

// Set default value for message request.ListUsersByProxy
func (this *ListUsersByProxy) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_limit(); err != nil {
		return err
	}
	return nil
}

func (this *ListNotificationsByProxy) _xxx_xxx_Validator_Validate_user_id() error {
	if !(this.UserId != "") {
		return protovalidator.FieldError1("ListNotificationsByProxy", "the value of field 'user_id' must be not equal to ''", this.UserId)
	}
	return nil
}

func (this *ListNotificationsByProxy) _xxx_xxx_Validator_Validate_limit() error {
	if !(this.Limit > 0) {
		return protovalidator.FieldError1("ListNotificationsByProxy", "the value of field 'limit' must be greater than '0'", protovalidator.Int32ToString(this.Limit))
	}
	if !(this.Limit <= 100) {
		return protovalidator.FieldError1("ListNotificationsByProxy", "the value of field 'limit' must be less than or equal to '100'", protovalidator.Int32ToString(this.Limit))
	}
	return nil
}

func (this *ListNotificationsByProxy) _xxx_xxx_Validator_Validate_nf_ids() error {
	if !(len(this.NfIds) >= 0) {
		return protovalidator.FieldError1("ListNotificationsByProxy", "the length of field 'nf_ids' must be greater than or equal to '0'", strconv.Itoa(len(this.NfIds)))
	}
	if !(len(this.NfIds) <= 100) {
		return protovalidator.FieldError1("ListNotificationsByProxy", "the length of field 'nf_ids' must be less than or equal to '100'", strconv.Itoa(len(this.NfIds)))
	}
	return nil
}

// Set default value for message request.ListNotificationsByProxy
func (this *ListNotificationsByProxy) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_user_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_limit(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_nf_ids(); err != nil {
		return err
	}
	return nil
}