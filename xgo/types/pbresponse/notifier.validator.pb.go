// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/response/notifier.proto

package pbresponse

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
)

func (this *DescNotificationListResponse) _xxx_xxx_Validator_Validate_notification_lists() error {
	for _, item := range this.NotificationLists {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message response.DescNotificationListResponse
func (this *DescNotificationListResponse) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_notification_lists(); err != nil {
		return err
	}
	return nil
}
