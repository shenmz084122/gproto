// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/service/spacemanager/space_manage.proto

package pbsvcspace

import (
	context "context"
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	pbrequest "github.com/DataWorkbench/gproto/xgo/types/pbrequest"
	pbresponse "github.com/DataWorkbench/gproto/xgo/types/pbresponse"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SpaceManageClient is the client API for SpaceManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceManageClient interface {
	// ListWorkspaces get a lists of workspaces.
	ListWorkspaces(ctx context.Context, in *pbrequest.ListWorkspaces, opts ...grpc.CallOption) (*pbresponse.ListWorkspaces, error)
	// ListMemberWorkspaces get a lists of workspaces that the specified user has be joined.
	ListMemberWorkspaces(ctx context.Context, in *pbrequest.ListWorkspaces, opts ...grpc.CallOption) (*pbresponse.ListWorkspaces, error)
	// DeleteWorkspaces allowed only invoke by space owner.
	DeleteWorkspaces(ctx context.Context, in *pbrequest.DeleteWorkspaces, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	// DisableWorkspaces allowed only invoke by space owner.
	DisableWorkspaces(ctx context.Context, in *pbrequest.DisableWorkspaces, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	// DisableWorkspaces allowed only invoke by space owner.
	EnableWorkspaces(ctx context.Context, in *pbrequest.EnableWorkspaces, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	CreateWorkspace(ctx context.Context, in *pbrequest.CreateWorkspace, opts ...grpc.CallOption) (*pbresponse.CreateWorkspace, error)
	UpdateWorkspace(ctx context.Context, in *pbrequest.UpdateWorkspace, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	DescribeWorkspace(ctx context.Context, in *pbrequest.DescribeWorkspace, opts ...grpc.CallOption) (*pbresponse.DescribeWorkspace, error)
	// Permission Check.
	// Notice: cannot check API includes: ListWorkspaces, DeleteWorkspaces, DisableWorkspaces, EnableWorkspaces.
	CheckPermission(ctx context.Context, in *pbrequest.CheckPermission, opts ...grpc.CallOption) (*pbresponse.CheckPermission, error)
	// Network config
	DescribeNetworkConfig(ctx context.Context, in *pbrequest.DescribeNetworkConfig, opts ...grpc.CallOption) (*pbresponse.DescribeNetworkConfig, error)
	AttachVPCToWorkspace(ctx context.Context, in *pbrequest.AttachVPCToWorkspace, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	DetachVPCFromWorkspace(ctx context.Context, in *pbrequest.DetachVPCFromWorkspace, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
}

type spaceManageClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceManageClient(cc grpc.ClientConnInterface) SpaceManageClient {
	return &spaceManageClient{cc}
}

func (c *spaceManageClient) ListWorkspaces(ctx context.Context, in *pbrequest.ListWorkspaces, opts ...grpc.CallOption) (*pbresponse.ListWorkspaces, error) {
	out := new(pbresponse.ListWorkspaces)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/ListWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) ListMemberWorkspaces(ctx context.Context, in *pbrequest.ListWorkspaces, opts ...grpc.CallOption) (*pbresponse.ListWorkspaces, error) {
	out := new(pbresponse.ListWorkspaces)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/ListMemberWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) DeleteWorkspaces(ctx context.Context, in *pbrequest.DeleteWorkspaces, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/DeleteWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) DisableWorkspaces(ctx context.Context, in *pbrequest.DisableWorkspaces, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/DisableWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) EnableWorkspaces(ctx context.Context, in *pbrequest.EnableWorkspaces, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/EnableWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) CreateWorkspace(ctx context.Context, in *pbrequest.CreateWorkspace, opts ...grpc.CallOption) (*pbresponse.CreateWorkspace, error) {
	out := new(pbresponse.CreateWorkspace)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/CreateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) UpdateWorkspace(ctx context.Context, in *pbrequest.UpdateWorkspace, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/UpdateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) DescribeWorkspace(ctx context.Context, in *pbrequest.DescribeWorkspace, opts ...grpc.CallOption) (*pbresponse.DescribeWorkspace, error) {
	out := new(pbresponse.DescribeWorkspace)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/DescribeWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) CheckPermission(ctx context.Context, in *pbrequest.CheckPermission, opts ...grpc.CallOption) (*pbresponse.CheckPermission, error) {
	out := new(pbresponse.CheckPermission)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) DescribeNetworkConfig(ctx context.Context, in *pbrequest.DescribeNetworkConfig, opts ...grpc.CallOption) (*pbresponse.DescribeNetworkConfig, error) {
	out := new(pbresponse.DescribeNetworkConfig)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/DescribeNetworkConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) AttachVPCToWorkspace(ctx context.Context, in *pbrequest.AttachVPCToWorkspace, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/AttachVPCToWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceManageClient) DetachVPCFromWorkspace(ctx context.Context, in *pbrequest.DetachVPCFromWorkspace, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/spacemanager.SpaceManage/DetachVPCFromWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceManageServer is the server API for SpaceManage service.
// All implementations must embed UnimplementedSpaceManageServer
// for forward compatibility
type SpaceManageServer interface {
	// ListWorkspaces get a lists of workspaces.
	ListWorkspaces(context.Context, *pbrequest.ListWorkspaces) (*pbresponse.ListWorkspaces, error)
	// ListMemberWorkspaces get a lists of workspaces that the specified user has be joined.
	ListMemberWorkspaces(context.Context, *pbrequest.ListWorkspaces) (*pbresponse.ListWorkspaces, error)
	// DeleteWorkspaces allowed only invoke by space owner.
	DeleteWorkspaces(context.Context, *pbrequest.DeleteWorkspaces) (*pbmodel.EmptyStruct, error)
	// DisableWorkspaces allowed only invoke by space owner.
	DisableWorkspaces(context.Context, *pbrequest.DisableWorkspaces) (*pbmodel.EmptyStruct, error)
	// DisableWorkspaces allowed only invoke by space owner.
	EnableWorkspaces(context.Context, *pbrequest.EnableWorkspaces) (*pbmodel.EmptyStruct, error)
	CreateWorkspace(context.Context, *pbrequest.CreateWorkspace) (*pbresponse.CreateWorkspace, error)
	UpdateWorkspace(context.Context, *pbrequest.UpdateWorkspace) (*pbmodel.EmptyStruct, error)
	DescribeWorkspace(context.Context, *pbrequest.DescribeWorkspace) (*pbresponse.DescribeWorkspace, error)
	// Permission Check.
	// Notice: cannot check API includes: ListWorkspaces, DeleteWorkspaces, DisableWorkspaces, EnableWorkspaces.
	CheckPermission(context.Context, *pbrequest.CheckPermission) (*pbresponse.CheckPermission, error)
	// Network config
	DescribeNetworkConfig(context.Context, *pbrequest.DescribeNetworkConfig) (*pbresponse.DescribeNetworkConfig, error)
	AttachVPCToWorkspace(context.Context, *pbrequest.AttachVPCToWorkspace) (*pbmodel.EmptyStruct, error)
	DetachVPCFromWorkspace(context.Context, *pbrequest.DetachVPCFromWorkspace) (*pbmodel.EmptyStruct, error)
	mustEmbedUnimplementedSpaceManageServer()
}

// UnimplementedSpaceManageServer must be embedded to have forward compatible implementations.
type UnimplementedSpaceManageServer struct {
}

func (UnimplementedSpaceManageServer) ListWorkspaces(context.Context, *pbrequest.ListWorkspaces) (*pbresponse.ListWorkspaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaces not implemented")
}
func (UnimplementedSpaceManageServer) ListMemberWorkspaces(context.Context, *pbrequest.ListWorkspaces) (*pbresponse.ListWorkspaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemberWorkspaces not implemented")
}
func (UnimplementedSpaceManageServer) DeleteWorkspaces(context.Context, *pbrequest.DeleteWorkspaces) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspaces not implemented")
}
func (UnimplementedSpaceManageServer) DisableWorkspaces(context.Context, *pbrequest.DisableWorkspaces) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableWorkspaces not implemented")
}
func (UnimplementedSpaceManageServer) EnableWorkspaces(context.Context, *pbrequest.EnableWorkspaces) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableWorkspaces not implemented")
}
func (UnimplementedSpaceManageServer) CreateWorkspace(context.Context, *pbrequest.CreateWorkspace) (*pbresponse.CreateWorkspace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedSpaceManageServer) UpdateWorkspace(context.Context, *pbrequest.UpdateWorkspace) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspace not implemented")
}
func (UnimplementedSpaceManageServer) DescribeWorkspace(context.Context, *pbrequest.DescribeWorkspace) (*pbresponse.DescribeWorkspace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeWorkspace not implemented")
}
func (UnimplementedSpaceManageServer) CheckPermission(context.Context, *pbrequest.CheckPermission) (*pbresponse.CheckPermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedSpaceManageServer) DescribeNetworkConfig(context.Context, *pbrequest.DescribeNetworkConfig) (*pbresponse.DescribeNetworkConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNetworkConfig not implemented")
}
func (UnimplementedSpaceManageServer) AttachVPCToWorkspace(context.Context, *pbrequest.AttachVPCToWorkspace) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachVPCToWorkspace not implemented")
}
func (UnimplementedSpaceManageServer) DetachVPCFromWorkspace(context.Context, *pbrequest.DetachVPCFromWorkspace) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachVPCFromWorkspace not implemented")
}
func (UnimplementedSpaceManageServer) mustEmbedUnimplementedSpaceManageServer() {}

// UnsafeSpaceManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceManageServer will
// result in compilation errors.
type UnsafeSpaceManageServer interface {
	mustEmbedUnimplementedSpaceManageServer()
}

func RegisterSpaceManageServer(s grpc.ServiceRegistrar, srv SpaceManageServer) {
	s.RegisterService(&SpaceManage_ServiceDesc, srv)
}

func _SpaceManage_ListWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.ListWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).ListWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/ListWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).ListWorkspaces(ctx, req.(*pbrequest.ListWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_ListMemberWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.ListWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).ListMemberWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/ListMemberWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).ListMemberWorkspaces(ctx, req.(*pbrequest.ListWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_DeleteWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.DeleteWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).DeleteWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/DeleteWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).DeleteWorkspaces(ctx, req.(*pbrequest.DeleteWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_DisableWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.DisableWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).DisableWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/DisableWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).DisableWorkspaces(ctx, req.(*pbrequest.DisableWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_EnableWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.EnableWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).EnableWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/EnableWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).EnableWorkspaces(ctx, req.(*pbrequest.EnableWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.CreateWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/CreateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).CreateWorkspace(ctx, req.(*pbrequest.CreateWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.UpdateWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/UpdateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).UpdateWorkspace(ctx, req.(*pbrequest.UpdateWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_DescribeWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.DescribeWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).DescribeWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/DescribeWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).DescribeWorkspace(ctx, req.(*pbrequest.DescribeWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.CheckPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).CheckPermission(ctx, req.(*pbrequest.CheckPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_DescribeNetworkConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.DescribeNetworkConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).DescribeNetworkConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/DescribeNetworkConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).DescribeNetworkConfig(ctx, req.(*pbrequest.DescribeNetworkConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_AttachVPCToWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.AttachVPCToWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).AttachVPCToWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/AttachVPCToWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).AttachVPCToWorkspace(ctx, req.(*pbrequest.AttachVPCToWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceManage_DetachVPCFromWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.DetachVPCFromWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceManageServer).DetachVPCFromWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemanager.SpaceManage/DetachVPCFromWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceManageServer).DetachVPCFromWorkspace(ctx, req.(*pbrequest.DetachVPCFromWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

// SpaceManage_ServiceDesc is the grpc.ServiceDesc for SpaceManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpaceManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemanager.SpaceManage",
	HandlerType: (*SpaceManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorkspaces",
			Handler:    _SpaceManage_ListWorkspaces_Handler,
		},
		{
			MethodName: "ListMemberWorkspaces",
			Handler:    _SpaceManage_ListMemberWorkspaces_Handler,
		},
		{
			MethodName: "DeleteWorkspaces",
			Handler:    _SpaceManage_DeleteWorkspaces_Handler,
		},
		{
			MethodName: "DisableWorkspaces",
			Handler:    _SpaceManage_DisableWorkspaces_Handler,
		},
		{
			MethodName: "EnableWorkspaces",
			Handler:    _SpaceManage_EnableWorkspaces_Handler,
		},
		{
			MethodName: "CreateWorkspace",
			Handler:    _SpaceManage_CreateWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _SpaceManage_UpdateWorkspace_Handler,
		},
		{
			MethodName: "DescribeWorkspace",
			Handler:    _SpaceManage_DescribeWorkspace_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _SpaceManage_CheckPermission_Handler,
		},
		{
			MethodName: "DescribeNetworkConfig",
			Handler:    _SpaceManage_DescribeNetworkConfig_Handler,
		},
		{
			MethodName: "AttachVPCToWorkspace",
			Handler:    _SpaceManage_AttachVPCToWorkspace_Handler,
		},
		{
			MethodName: "DetachVPCFromWorkspace",
			Handler:    _SpaceManage_DetachVPCFromWorkspace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/spacemanager/space_manage.proto",
}
