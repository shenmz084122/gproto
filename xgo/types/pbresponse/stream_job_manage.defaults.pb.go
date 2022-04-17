// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/response/stream_job_manage.proto

package pbresponse

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message response.ListStreamJobs
func (this *ListStreamJobs) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.CreateStreamJob
func (this *CreateStreamJob) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.DescribeStreamJob
func (this *DescribeStreamJob) SetDefaults() {
	if this == nil {
		return
	}
	if this.Info != nil {
		if dt, ok := interface{}(this.Info).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message response.GetStreamJobCode
func (this *GetStreamJobCode) SetDefaults() {
	if this == nil {
		return
	}
	if this.Code != nil {
		if dt, ok := interface{}(this.Code).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message response.GetStreamJobArgs
func (this *GetStreamJobArgs) SetDefaults() {
	if this == nil {
		return
	}
	if this.Args != nil {
		if dt, ok := interface{}(this.Args).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message response.GetStreamJobSchedule
func (this *GetStreamJobSchedule) SetDefaults() {
	if this == nil {
		return
	}
	if this.Schedule != nil {
		if dt, ok := interface{}(this.Schedule).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message response.ListReleaseStreamJobs
func (this *ListReleaseStreamJobs) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.ListStreamJobVersions
func (this *ListStreamJobVersions) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.ListBuiltInConnectors
func (this *ListBuiltInConnectors) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.DescribeFlinkUIByInstanceId
func (this *DescribeFlinkUIByInstanceId) SetDefaults() {
	if this == nil {
		return
	}
	return
}
