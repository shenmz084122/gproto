// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/alert_manage.proto

package pbrequest

import (
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.ListAlertPolicies
func (this *ListAlertPolicies) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.DeleteAlertPolicies
func (this *DeleteAlertPolicies) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.EnableAlertPolicies
func (this *EnableAlertPolicies) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DisableAlertPolicies
func (this *DisableAlertPolicies) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.CreateAlertPolicy
func (this *CreateAlertPolicy) SetDefaults() {
	if this == nil {
		return
	}
	if this.MonitorItem != nil {
		if dt, ok := interface{}(this.MonitorItem).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.TriggerRule == 0 {
		this.TriggerRule = pbmodel.AlertPolicy_TriggerRule(1)
	}
	if this.TriggerAction == 0 {
		this.TriggerAction = pbmodel.AlertPolicy_TriggerAction(1)
	}
	return
}

// Set default value for message request.UpdateAlertPolicy
func (this *UpdateAlertPolicy) SetDefaults() {
	if this == nil {
		return
	}
	if this.MonitorItem != nil {
		if dt, ok := interface{}(this.MonitorItem).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.TriggerRule == 0 {
		this.TriggerRule = pbmodel.AlertPolicy_TriggerRule(1)
	}
	if this.TriggerAction == 0 {
		this.TriggerAction = pbmodel.AlertPolicy_TriggerAction(1)
	}
	return
}

// Set default value for message request.DescribeAlertPolicy
func (this *DescribeAlertPolicy) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.JobBoundAlertPolicies
func (this *JobBoundAlertPolicies) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.JobUnboundAlertPolicies
func (this *JobUnboundAlertPolicies) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListAlertPoliciesByJob
func (this *ListAlertPoliciesByJob) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.AlertPolicyBoundJobs
func (this *AlertPolicyBoundJobs) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.AlertPolicyUnboundJobs
func (this *AlertPolicyUnboundJobs) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListJobsByAlertPolicy
func (this *ListJobsByAlertPolicy) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.ListAlertLogs
func (this *ListAlertLogs) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.SendAlert
func (this *SendAlert) SetDefaults() {
	if this == nil {
		return
	}
	return
}
