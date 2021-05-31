// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wfpb

import (
	context "context"
	model "github.com/DataWorkbench/gproto/pkg/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WorkflowClient is the client API for Workflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowClient interface {
	// Interface for stream workflow.
	//
	// ListStreamFlows to get a list of stream workflow of the workspace.
	ListStreamFlows(ctx context.Context, in *ListStreamFlowsRequest, opts ...grpc.CallOption) (*ListStreamFlowsReply, error)
	// CreateStreamFlow to create a new stream workflow.
	CreateStreamFlow(ctx context.Context, in *CreateStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// DeleteStreamFlow to delete the specified stream workflow.
	DeleteStreamFlow(ctx context.Context, in *DeleteStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// UpdateStreamFlow to update the info for the specified stream workflow.
	UpdateStreamFlow(ctx context.Context, in *UpdateStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// DescribeStreamFlow to get the info of the specified stream workflow.
	DescribeStreamFlow(ctx context.Context, in *DescribeStreamFlowRequest, opts ...grpc.CallOption) (*DescribeStreamFlowReply, error)
	// SetStreamFlowNode to set the node properties of the specified stream workflow.
	SetStreamFlowNode(ctx context.Context, in *SetStreamFlowNodeRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// GetStreamFlowNode to get the node properties of the specified stream workflow.
	GetStreamFlowNode(ctx context.Context, in *GetStreamFlowNodeRequest, opts ...grpc.CallOption) (*GetStreamFlowNodeReply, error)
	// SetStreamFlowEnv to set the environmental parameters of the specified stream workflow.
	SetStreamFlowEnv(ctx context.Context, in *SetStreamFlowEnvRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// GetStreamFlowEnv to get the environmental parameters of the specified stream workflow.
	GetStreamFlowEnv(ctx context.Context, in *GetStreamFlowEnvRequest, opts ...grpc.CallOption) (*GetStreamFlowEnvReply, error)
	// SetStreamFlowSchedule to set the schedule properties of the specified stream workflow.
	SetStreamFlowSchedule(ctx context.Context, in *SetStreamFlowScheduleRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// GetStreamFlowSchedule to get the schedule properties of the specified stream workflow.
	GetStreamFlowSchedule(ctx context.Context, in *GetStreamFlowScheduleRequest, opts ...grpc.CallOption) (*GetStreamFlowScheduleReply, error)
	// ExecuteStreamFlow to manual execution a stream workflow task.
	ExecuteStreamFlow(ctx context.Context, in *ExecuteStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// Interface for stream workflow release.
	//
	// ReleaseStreamFlow to publish the specified workflow to schedule system with a new version.
	ReleaseStreamFlow(ctx context.Context, in *ReleaseStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// SuspendStreamFlow to suspend the specified workflow in schedule system.
	SuspendReleaseStreamFlow(ctx context.Context, in *SuspendReleaseStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// ResumeStreamFlow to resume the suspended workflow in schedule system.
	ResumeReleaseStreamFlow(ctx context.Context, in *ResumeReleaseStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// ListStreamReleases for gets a list of all published workflow in the workspace.
	ListReleaseStreamFlows(ctx context.Context, in *ListReleaseStreamFlowsRequest, opts ...grpc.CallOption) (*ListReleaseStreamFlowsReply, error)
	// Interface for stream workflow versions.
	//
	// ListStreamVersions for gets a list of all versions of the specified workflow.
	ListStreamFlowVersions(ctx context.Context, in *ListStreamFlowVersionsRequest, opts ...grpc.CallOption) (*ListStreamFlowVersionsReply, error)
	// DescribeStreamVersion for get the info of the workflow of the specified version.
	DescribeStreamFlowVersion(ctx context.Context, in *DescribeStreamFlowVersionRequest, opts ...grpc.CallOption) (*DescribeStreamFlowVersionReply, error)
	// GetStreamVersionNode for get the node properties of the workflow of the specified version.
	GetStreamFlowVersionNode(ctx context.Context, in *GetStreamFlowVersionNodeRequest, opts ...grpc.CallOption) (*GetStreamFlowVersionNodeReply, error)
	// GetStreamVersionEnv for get the environmental parameters of the workflow of the specified version.
	GetStreamFlowVersionEnv(ctx context.Context, in *GetStreamFlowVersionEnvRequest, opts ...grpc.CallOption) (*GetStreamFlowVersionEnvReply, error)
	// GetStreamReleaseSchedule for get the schedule properties of the workflow of the specified version.
	GetStreamFlowVersionSchedule(ctx context.Context, in *GetStreamFlowVersionScheduleRequest, opts ...grpc.CallOption) (*GetStreamFlowVersionScheduleReply, error)
}

type workflowClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowClient(cc grpc.ClientConnInterface) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) ListStreamFlows(ctx context.Context, in *ListStreamFlowsRequest, opts ...grpc.CallOption) (*ListStreamFlowsReply, error) {
	out := new(ListStreamFlowsReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/ListStreamFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) CreateStreamFlow(ctx context.Context, in *CreateStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/CreateStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) DeleteStreamFlow(ctx context.Context, in *DeleteStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/DeleteStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) UpdateStreamFlow(ctx context.Context, in *UpdateStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/UpdateStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) DescribeStreamFlow(ctx context.Context, in *DescribeStreamFlowRequest, opts ...grpc.CallOption) (*DescribeStreamFlowReply, error) {
	out := new(DescribeStreamFlowReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/DescribeStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) SetStreamFlowNode(ctx context.Context, in *SetStreamFlowNodeRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/SetStreamFlowNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetStreamFlowNode(ctx context.Context, in *GetStreamFlowNodeRequest, opts ...grpc.CallOption) (*GetStreamFlowNodeReply, error) {
	out := new(GetStreamFlowNodeReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/GetStreamFlowNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) SetStreamFlowEnv(ctx context.Context, in *SetStreamFlowEnvRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/SetStreamFlowEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetStreamFlowEnv(ctx context.Context, in *GetStreamFlowEnvRequest, opts ...grpc.CallOption) (*GetStreamFlowEnvReply, error) {
	out := new(GetStreamFlowEnvReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/GetStreamFlowEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) SetStreamFlowSchedule(ctx context.Context, in *SetStreamFlowScheduleRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/SetStreamFlowSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetStreamFlowSchedule(ctx context.Context, in *GetStreamFlowScheduleRequest, opts ...grpc.CallOption) (*GetStreamFlowScheduleReply, error) {
	out := new(GetStreamFlowScheduleReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/GetStreamFlowSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ExecuteStreamFlow(ctx context.Context, in *ExecuteStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/ExecuteStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ReleaseStreamFlow(ctx context.Context, in *ReleaseStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/ReleaseStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) SuspendReleaseStreamFlow(ctx context.Context, in *SuspendReleaseStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/SuspendReleaseStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ResumeReleaseStreamFlow(ctx context.Context, in *ResumeReleaseStreamFlowRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/ResumeReleaseStreamFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ListReleaseStreamFlows(ctx context.Context, in *ListReleaseStreamFlowsRequest, opts ...grpc.CallOption) (*ListReleaseStreamFlowsReply, error) {
	out := new(ListReleaseStreamFlowsReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/ListReleaseStreamFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) ListStreamFlowVersions(ctx context.Context, in *ListStreamFlowVersionsRequest, opts ...grpc.CallOption) (*ListStreamFlowVersionsReply, error) {
	out := new(ListStreamFlowVersionsReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/ListStreamFlowVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) DescribeStreamFlowVersion(ctx context.Context, in *DescribeStreamFlowVersionRequest, opts ...grpc.CallOption) (*DescribeStreamFlowVersionReply, error) {
	out := new(DescribeStreamFlowVersionReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/DescribeStreamFlowVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetStreamFlowVersionNode(ctx context.Context, in *GetStreamFlowVersionNodeRequest, opts ...grpc.CallOption) (*GetStreamFlowVersionNodeReply, error) {
	out := new(GetStreamFlowVersionNodeReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/GetStreamFlowVersionNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetStreamFlowVersionEnv(ctx context.Context, in *GetStreamFlowVersionEnvRequest, opts ...grpc.CallOption) (*GetStreamFlowVersionEnvReply, error) {
	out := new(GetStreamFlowVersionEnvReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/GetStreamFlowVersionEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowClient) GetStreamFlowVersionSchedule(ctx context.Context, in *GetStreamFlowVersionScheduleRequest, opts ...grpc.CallOption) (*GetStreamFlowVersionScheduleReply, error) {
	out := new(GetStreamFlowVersionScheduleReply)
	err := c.cc.Invoke(ctx, "/wfpb.Workflow/GetStreamFlowVersionSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServer is the server API for Workflow service.
// All implementations must embed UnimplementedWorkflowServer
// for forward compatibility
type WorkflowServer interface {
	// Interface for stream workflow.
	//
	// ListStreamFlows to get a list of stream workflow of the workspace.
	ListStreamFlows(context.Context, *ListStreamFlowsRequest) (*ListStreamFlowsReply, error)
	// CreateStreamFlow to create a new stream workflow.
	CreateStreamFlow(context.Context, *CreateStreamFlowRequest) (*model.EmptyStruct, error)
	// DeleteStreamFlow to delete the specified stream workflow.
	DeleteStreamFlow(context.Context, *DeleteStreamFlowRequest) (*model.EmptyStruct, error)
	// UpdateStreamFlow to update the info for the specified stream workflow.
	UpdateStreamFlow(context.Context, *UpdateStreamFlowRequest) (*model.EmptyStruct, error)
	// DescribeStreamFlow to get the info of the specified stream workflow.
	DescribeStreamFlow(context.Context, *DescribeStreamFlowRequest) (*DescribeStreamFlowReply, error)
	// SetStreamFlowNode to set the node properties of the specified stream workflow.
	SetStreamFlowNode(context.Context, *SetStreamFlowNodeRequest) (*model.EmptyStruct, error)
	// GetStreamFlowNode to get the node properties of the specified stream workflow.
	GetStreamFlowNode(context.Context, *GetStreamFlowNodeRequest) (*GetStreamFlowNodeReply, error)
	// SetStreamFlowEnv to set the environmental parameters of the specified stream workflow.
	SetStreamFlowEnv(context.Context, *SetStreamFlowEnvRequest) (*model.EmptyStruct, error)
	// GetStreamFlowEnv to get the environmental parameters of the specified stream workflow.
	GetStreamFlowEnv(context.Context, *GetStreamFlowEnvRequest) (*GetStreamFlowEnvReply, error)
	// SetStreamFlowSchedule to set the schedule properties of the specified stream workflow.
	SetStreamFlowSchedule(context.Context, *SetStreamFlowScheduleRequest) (*model.EmptyStruct, error)
	// GetStreamFlowSchedule to get the schedule properties of the specified stream workflow.
	GetStreamFlowSchedule(context.Context, *GetStreamFlowScheduleRequest) (*GetStreamFlowScheduleReply, error)
	// ExecuteStreamFlow to manual execution a stream workflow task.
	ExecuteStreamFlow(context.Context, *ExecuteStreamFlowRequest) (*model.EmptyStruct, error)
	// Interface for stream workflow release.
	//
	// ReleaseStreamFlow to publish the specified workflow to schedule system with a new version.
	ReleaseStreamFlow(context.Context, *ReleaseStreamFlowRequest) (*model.EmptyStruct, error)
	// SuspendStreamFlow to suspend the specified workflow in schedule system.
	SuspendReleaseStreamFlow(context.Context, *SuspendReleaseStreamFlowRequest) (*model.EmptyStruct, error)
	// ResumeStreamFlow to resume the suspended workflow in schedule system.
	ResumeReleaseStreamFlow(context.Context, *ResumeReleaseStreamFlowRequest) (*model.EmptyStruct, error)
	// ListStreamReleases for gets a list of all published workflow in the workspace.
	ListReleaseStreamFlows(context.Context, *ListReleaseStreamFlowsRequest) (*ListReleaseStreamFlowsReply, error)
	// Interface for stream workflow versions.
	//
	// ListStreamVersions for gets a list of all versions of the specified workflow.
	ListStreamFlowVersions(context.Context, *ListStreamFlowVersionsRequest) (*ListStreamFlowVersionsReply, error)
	// DescribeStreamVersion for get the info of the workflow of the specified version.
	DescribeStreamFlowVersion(context.Context, *DescribeStreamFlowVersionRequest) (*DescribeStreamFlowVersionReply, error)
	// GetStreamVersionNode for get the node properties of the workflow of the specified version.
	GetStreamFlowVersionNode(context.Context, *GetStreamFlowVersionNodeRequest) (*GetStreamFlowVersionNodeReply, error)
	// GetStreamVersionEnv for get the environmental parameters of the workflow of the specified version.
	GetStreamFlowVersionEnv(context.Context, *GetStreamFlowVersionEnvRequest) (*GetStreamFlowVersionEnvReply, error)
	// GetStreamReleaseSchedule for get the schedule properties of the workflow of the specified version.
	GetStreamFlowVersionSchedule(context.Context, *GetStreamFlowVersionScheduleRequest) (*GetStreamFlowVersionScheduleReply, error)
	mustEmbedUnimplementedWorkflowServer()
}

// UnimplementedWorkflowServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowServer struct {
}

func (UnimplementedWorkflowServer) ListStreamFlows(context.Context, *ListStreamFlowsRequest) (*ListStreamFlowsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStreamFlows not implemented")
}
func (UnimplementedWorkflowServer) CreateStreamFlow(context.Context, *CreateStreamFlowRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) DeleteStreamFlow(context.Context, *DeleteStreamFlowRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) UpdateStreamFlow(context.Context, *UpdateStreamFlowRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) DescribeStreamFlow(context.Context, *DescribeStreamFlowRequest) (*DescribeStreamFlowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) SetStreamFlowNode(context.Context, *SetStreamFlowNodeRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStreamFlowNode not implemented")
}
func (UnimplementedWorkflowServer) GetStreamFlowNode(context.Context, *GetStreamFlowNodeRequest) (*GetStreamFlowNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamFlowNode not implemented")
}
func (UnimplementedWorkflowServer) SetStreamFlowEnv(context.Context, *SetStreamFlowEnvRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStreamFlowEnv not implemented")
}
func (UnimplementedWorkflowServer) GetStreamFlowEnv(context.Context, *GetStreamFlowEnvRequest) (*GetStreamFlowEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamFlowEnv not implemented")
}
func (UnimplementedWorkflowServer) SetStreamFlowSchedule(context.Context, *SetStreamFlowScheduleRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStreamFlowSchedule not implemented")
}
func (UnimplementedWorkflowServer) GetStreamFlowSchedule(context.Context, *GetStreamFlowScheduleRequest) (*GetStreamFlowScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamFlowSchedule not implemented")
}
func (UnimplementedWorkflowServer) ExecuteStreamFlow(context.Context, *ExecuteStreamFlowRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) ReleaseStreamFlow(context.Context, *ReleaseStreamFlowRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) SuspendReleaseStreamFlow(context.Context, *SuspendReleaseStreamFlowRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuspendReleaseStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) ResumeReleaseStreamFlow(context.Context, *ResumeReleaseStreamFlowRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeReleaseStreamFlow not implemented")
}
func (UnimplementedWorkflowServer) ListReleaseStreamFlows(context.Context, *ListReleaseStreamFlowsRequest) (*ListReleaseStreamFlowsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReleaseStreamFlows not implemented")
}
func (UnimplementedWorkflowServer) ListStreamFlowVersions(context.Context, *ListStreamFlowVersionsRequest) (*ListStreamFlowVersionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStreamFlowVersions not implemented")
}
func (UnimplementedWorkflowServer) DescribeStreamFlowVersion(context.Context, *DescribeStreamFlowVersionRequest) (*DescribeStreamFlowVersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStreamFlowVersion not implemented")
}
func (UnimplementedWorkflowServer) GetStreamFlowVersionNode(context.Context, *GetStreamFlowVersionNodeRequest) (*GetStreamFlowVersionNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamFlowVersionNode not implemented")
}
func (UnimplementedWorkflowServer) GetStreamFlowVersionEnv(context.Context, *GetStreamFlowVersionEnvRequest) (*GetStreamFlowVersionEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamFlowVersionEnv not implemented")
}
func (UnimplementedWorkflowServer) GetStreamFlowVersionSchedule(context.Context, *GetStreamFlowVersionScheduleRequest) (*GetStreamFlowVersionScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamFlowVersionSchedule not implemented")
}
func (UnimplementedWorkflowServer) mustEmbedUnimplementedWorkflowServer() {}

// UnsafeWorkflowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServer will
// result in compilation errors.
type UnsafeWorkflowServer interface {
	mustEmbedUnimplementedWorkflowServer()
}

func RegisterWorkflowServer(s grpc.ServiceRegistrar, srv WorkflowServer) {
	s.RegisterService(&_Workflow_serviceDesc, srv)
}

func _Workflow_ListStreamFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStreamFlowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ListStreamFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/ListStreamFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ListStreamFlows(ctx, req.(*ListStreamFlowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_CreateStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).CreateStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/CreateStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).CreateStreamFlow(ctx, req.(*CreateStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_DeleteStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).DeleteStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/DeleteStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).DeleteStreamFlow(ctx, req.(*DeleteStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_UpdateStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).UpdateStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/UpdateStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).UpdateStreamFlow(ctx, req.(*UpdateStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_DescribeStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).DescribeStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/DescribeStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).DescribeStreamFlow(ctx, req.(*DescribeStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_SetStreamFlowNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStreamFlowNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).SetStreamFlowNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/SetStreamFlowNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).SetStreamFlowNode(ctx, req.(*SetStreamFlowNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetStreamFlowNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamFlowNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetStreamFlowNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/GetStreamFlowNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetStreamFlowNode(ctx, req.(*GetStreamFlowNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_SetStreamFlowEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStreamFlowEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).SetStreamFlowEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/SetStreamFlowEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).SetStreamFlowEnv(ctx, req.(*SetStreamFlowEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetStreamFlowEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamFlowEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetStreamFlowEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/GetStreamFlowEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetStreamFlowEnv(ctx, req.(*GetStreamFlowEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_SetStreamFlowSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStreamFlowScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).SetStreamFlowSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/SetStreamFlowSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).SetStreamFlowSchedule(ctx, req.(*SetStreamFlowScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetStreamFlowSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamFlowScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetStreamFlowSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/GetStreamFlowSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetStreamFlowSchedule(ctx, req.(*GetStreamFlowScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ExecuteStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ExecuteStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/ExecuteStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ExecuteStreamFlow(ctx, req.(*ExecuteStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ReleaseStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ReleaseStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/ReleaseStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ReleaseStreamFlow(ctx, req.(*ReleaseStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_SuspendReleaseStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendReleaseStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).SuspendReleaseStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/SuspendReleaseStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).SuspendReleaseStreamFlow(ctx, req.(*SuspendReleaseStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ResumeReleaseStreamFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeReleaseStreamFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ResumeReleaseStreamFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/ResumeReleaseStreamFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ResumeReleaseStreamFlow(ctx, req.(*ResumeReleaseStreamFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ListReleaseStreamFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReleaseStreamFlowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ListReleaseStreamFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/ListReleaseStreamFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ListReleaseStreamFlows(ctx, req.(*ListReleaseStreamFlowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_ListStreamFlowVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStreamFlowVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).ListStreamFlowVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/ListStreamFlowVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).ListStreamFlowVersions(ctx, req.(*ListStreamFlowVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_DescribeStreamFlowVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeStreamFlowVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).DescribeStreamFlowVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/DescribeStreamFlowVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).DescribeStreamFlowVersion(ctx, req.(*DescribeStreamFlowVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetStreamFlowVersionNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamFlowVersionNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetStreamFlowVersionNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/GetStreamFlowVersionNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetStreamFlowVersionNode(ctx, req.(*GetStreamFlowVersionNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetStreamFlowVersionEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamFlowVersionEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetStreamFlowVersionEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/GetStreamFlowVersionEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetStreamFlowVersionEnv(ctx, req.(*GetStreamFlowVersionEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflow_GetStreamFlowVersionSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamFlowVersionScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).GetStreamFlowVersionSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wfpb.Workflow/GetStreamFlowVersionSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).GetStreamFlowVersionSchedule(ctx, req.(*GetStreamFlowVersionScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wfpb.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStreamFlows",
			Handler:    _Workflow_ListStreamFlows_Handler,
		},
		{
			MethodName: "CreateStreamFlow",
			Handler:    _Workflow_CreateStreamFlow_Handler,
		},
		{
			MethodName: "DeleteStreamFlow",
			Handler:    _Workflow_DeleteStreamFlow_Handler,
		},
		{
			MethodName: "UpdateStreamFlow",
			Handler:    _Workflow_UpdateStreamFlow_Handler,
		},
		{
			MethodName: "DescribeStreamFlow",
			Handler:    _Workflow_DescribeStreamFlow_Handler,
		},
		{
			MethodName: "SetStreamFlowNode",
			Handler:    _Workflow_SetStreamFlowNode_Handler,
		},
		{
			MethodName: "GetStreamFlowNode",
			Handler:    _Workflow_GetStreamFlowNode_Handler,
		},
		{
			MethodName: "SetStreamFlowEnv",
			Handler:    _Workflow_SetStreamFlowEnv_Handler,
		},
		{
			MethodName: "GetStreamFlowEnv",
			Handler:    _Workflow_GetStreamFlowEnv_Handler,
		},
		{
			MethodName: "SetStreamFlowSchedule",
			Handler:    _Workflow_SetStreamFlowSchedule_Handler,
		},
		{
			MethodName: "GetStreamFlowSchedule",
			Handler:    _Workflow_GetStreamFlowSchedule_Handler,
		},
		{
			MethodName: "ExecuteStreamFlow",
			Handler:    _Workflow_ExecuteStreamFlow_Handler,
		},
		{
			MethodName: "ReleaseStreamFlow",
			Handler:    _Workflow_ReleaseStreamFlow_Handler,
		},
		{
			MethodName: "SuspendReleaseStreamFlow",
			Handler:    _Workflow_SuspendReleaseStreamFlow_Handler,
		},
		{
			MethodName: "ResumeReleaseStreamFlow",
			Handler:    _Workflow_ResumeReleaseStreamFlow_Handler,
		},
		{
			MethodName: "ListReleaseStreamFlows",
			Handler:    _Workflow_ListReleaseStreamFlows_Handler,
		},
		{
			MethodName: "ListStreamFlowVersions",
			Handler:    _Workflow_ListStreamFlowVersions_Handler,
		},
		{
			MethodName: "DescribeStreamFlowVersion",
			Handler:    _Workflow_DescribeStreamFlowVersion_Handler,
		},
		{
			MethodName: "GetStreamFlowVersionNode",
			Handler:    _Workflow_GetStreamFlowVersionNode_Handler,
		},
		{
			MethodName: "GetStreamFlowVersionEnv",
			Handler:    _Workflow_GetStreamFlowVersionEnv_Handler,
		},
		{
			MethodName: "GetStreamFlowVersionSchedule",
			Handler:    _Workflow_GetStreamFlowVersionSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/workflow.proto",
}
