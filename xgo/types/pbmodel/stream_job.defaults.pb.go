// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.1
// source: proto/types/model/stream_job.proto

package pbmodel

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbflink"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message model.StreamJob
func (this *StreamJob) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.StreamJobProperty
func (this *StreamJobProperty) SetDefaults() {
	if this == nil {
		return
	}
	if this.Code != nil {
		if dt, ok := interface{}(this.Code).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Args != nil {
		if dt, ok := interface{}(this.Args).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Schedule != nil {
		if dt, ok := interface{}(this.Schedule).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message model.StreamJobCode
func (this *StreamJobCode) SetDefaults() {
	if this == nil {
		return
	}
	if this.Sql != nil {
		if dt, ok := interface{}(this.Sql).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Jar != nil {
		if dt, ok := interface{}(this.Jar).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Python != nil {
		if dt, ok := interface{}(this.Python).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Scala != nil {
		if dt, ok := interface{}(this.Scala).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message model.StreamJobArgs
func (this *StreamJobArgs) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.StreamJobSchedule
func (this *StreamJobSchedule) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.StreamJobRelease
func (this *StreamJobRelease) SetDefaults() {
	if this == nil {
		return
	}
	return
}