// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/model/observable.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message model.WorkspaceOverview
func (this *WorkspaceOverview) SetDefaults() {
	if this == nil {
		return
	}
	if this.StreamJobRelease != nil {
		if dt, ok := interface{}(this.StreamJobRelease).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.SyncJobRelease != nil {
		if dt, ok := interface{}(this.SyncJobRelease).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.StreamJobInstance != nil {
		if dt, ok := interface{}(this.StreamJobInstance).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.SyncJobInstance != nil {
		if dt, ok := interface{}(this.SyncJobInstance).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.FlinkCluster != nil {
		if dt, ok := interface{}(this.FlinkCluster).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message model.WorkspaceOverview.StreamJobRelease
func (this *WorkspaceOverview_StreamJobRelease) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.WorkspaceOverview.SyncJobRelease
func (this *WorkspaceOverview_SyncJobRelease) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.WorkspaceOverview.StreamInstance
func (this *WorkspaceOverview_StreamInstance) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.WorkspaceOverview.SyncInstance
func (this *WorkspaceOverview_SyncInstance) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.WorkspaceOverview.FlinkCluster
func (this *WorkspaceOverview_FlinkCluster) SetDefaults() {
	if this == nil {
		return
	}
	return
}
