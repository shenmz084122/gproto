// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/model/scheduler.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message model.StreamJobEvent
func (this *StreamJobEvent) SetDefaults() {
	if this == nil {
		return
	}
	if this.Property != nil {
		if dt, ok := interface{}(this.Property).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message model.SyncJobEvent
func (this *SyncJobEvent) SetDefaults() {
	if this == nil {
		return
	}
	if this.Property != nil {
		if dt, ok := interface{}(this.Property).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message model.FlinkMonitorEvent
func (this *FlinkMonitorEvent) SetDefaults() {
	if this == nil {
		return
	}
	return
}
