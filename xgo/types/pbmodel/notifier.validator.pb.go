// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/model/notifier.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

func (this *NotificationList) _xxx_xxx_Validator_Validate_items() error {
	for _, item := range this.Items {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message model.NotificationList
func (this *NotificationList) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_items(); err != nil {
		return err
	}
	return nil
}

// Set default value for message model.NotificationListItem
func (this *NotificationListItem) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}