// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/apidashboard.proto

package pbrequest

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.CreateRoute
func (this *CreateRoute) SetDefaults() {
	if this == nil {
		return
	}
	if this.RouteInfo != nil {
		if dt, ok := interface{}(this.RouteInfo).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.UpdateRoute
func (this *UpdateRoute) SetDefaults() {
	if this == nil {
		return
	}
	if this.RouteInfo != nil {
		if dt, ok := interface{}(this.RouteInfo).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.RouteInfo
func (this *RouteInfo) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DeleteRoute
func (this *DeleteRoute) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListRoutes
func (this *ListRoutes) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.CreateUpstream
func (this *CreateUpstream) SetDefaults() {
	if this == nil {
		return
	}
	if this.UpstreamInfo != nil {
		if dt, ok := interface{}(this.UpstreamInfo).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.UpdateUpstream
func (this *UpdateUpstream) SetDefaults() {
	if this == nil {
		return
	}
	if this.UpstreamInfo != nil {
		if dt, ok := interface{}(this.UpstreamInfo).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.UpstreamInfo
func (this *UpstreamInfo) SetDefaults() {
	if this == nil {
		return
	}
	if this.Tls != nil {
		if dt, ok := interface{}(this.Tls).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Timeout != nil {
		if dt, ok := interface{}(this.Timeout).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.DeleteUpstream
func (this *DeleteUpstream) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListUpstreams
func (this *ListUpstreams) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.CreateSSL
func (this *CreateSSL) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DeleteSSL
func (this *DeleteSSL) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListSSLs
func (this *ListSSLs) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.CreateApiService
func (this *CreateApiService) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DeleteApiService
func (this *DeleteApiService) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.UpdateApiService
func (this *UpdateApiService) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListApiServices
func (this *ListApiServices) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.AddSvcReqCount
func (this *AddSvcReqCount) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.GetSvcReqCount
func (this *GetSvcReqCount) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DeleteProjectRoutes
func (this *DeleteProjectRoutes) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.CreateAuthKey
func (this *CreateAuthKey) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DeleteAuthKey
func (this *DeleteAuthKey) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.UpdateAuthKey
func (this *UpdateAuthKey) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListAuthKeys
func (this *ListAuthKeys) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.BindAuthKey
func (this *BindAuthKey) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.UnbindAuthKey
func (this *UnbindAuthKey) SetDefaults() {
	if this == nil {
		return
	}
	return
}
