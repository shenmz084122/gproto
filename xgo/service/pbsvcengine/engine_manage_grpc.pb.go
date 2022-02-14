// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/service/enginemanager/engine_manage.proto

package pbsvcengine

import (
	context "context"
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	pbrequest "github.com/DataWorkbench/gproto/xgo/types/pbrequest"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EngineManageClient is the client API for EngineManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngineManageClient interface {
	CreateFlinkClusterInK8S(ctx context.Context, in *pbrequest.CreateFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	//  rpc DeleteFlinkClusterInK8s(request.DeleteFlinkClusterInK8s) returns (model.EmptyStruct) {}
	StartFlinkClusterInK8S(ctx context.Context, in *pbrequest.StartFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	StopFlinkClusterInK8S(ctx context.Context, in *pbrequest.StopFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	CreateNetworkBrokerInK8S(ctx context.Context, in *pbrequest.CreateFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
	DeleteNetworkBrokerInK8S(ctx context.Context, in *pbrequest.DeleteNetworkBrokerInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error)
}

type engineManageClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineManageClient(cc grpc.ClientConnInterface) EngineManageClient {
	return &engineManageClient{cc}
}

func (c *engineManageClient) CreateFlinkClusterInK8S(ctx context.Context, in *pbrequest.CreateFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/enginecenter.EngineManage/CreateFlinkClusterInK8s", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineManageClient) StartFlinkClusterInK8S(ctx context.Context, in *pbrequest.StartFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/enginecenter.EngineManage/StartFlinkClusterInK8s", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineManageClient) StopFlinkClusterInK8S(ctx context.Context, in *pbrequest.StopFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/enginecenter.EngineManage/StopFlinkClusterInK8s", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineManageClient) CreateNetworkBrokerInK8S(ctx context.Context, in *pbrequest.CreateFlinkClusterInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/enginecenter.EngineManage/CreateNetworkBrokerInK8s", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineManageClient) DeleteNetworkBrokerInK8S(ctx context.Context, in *pbrequest.DeleteNetworkBrokerInK8S, opts ...grpc.CallOption) (*pbmodel.EmptyStruct, error) {
	out := new(pbmodel.EmptyStruct)
	err := c.cc.Invoke(ctx, "/enginecenter.EngineManage/DeleteNetworkBrokerInK8s", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineManageServer is the server API for EngineManage service.
// All implementations must embed UnimplementedEngineManageServer
// for forward compatibility
type EngineManageServer interface {
	CreateFlinkClusterInK8S(context.Context, *pbrequest.CreateFlinkClusterInK8S) (*pbmodel.EmptyStruct, error)
	//  rpc DeleteFlinkClusterInK8s(request.DeleteFlinkClusterInK8s) returns (model.EmptyStruct) {}
	StartFlinkClusterInK8S(context.Context, *pbrequest.StartFlinkClusterInK8S) (*pbmodel.EmptyStruct, error)
	StopFlinkClusterInK8S(context.Context, *pbrequest.StopFlinkClusterInK8S) (*pbmodel.EmptyStruct, error)
	CreateNetworkBrokerInK8S(context.Context, *pbrequest.CreateFlinkClusterInK8S) (*pbmodel.EmptyStruct, error)
	DeleteNetworkBrokerInK8S(context.Context, *pbrequest.DeleteNetworkBrokerInK8S) (*pbmodel.EmptyStruct, error)
	mustEmbedUnimplementedEngineManageServer()
}

// UnimplementedEngineManageServer must be embedded to have forward compatible implementations.
type UnimplementedEngineManageServer struct {
}

func (UnimplementedEngineManageServer) CreateFlinkClusterInK8S(context.Context, *pbrequest.CreateFlinkClusterInK8S) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlinkClusterInK8S not implemented")
}
func (UnimplementedEngineManageServer) StartFlinkClusterInK8S(context.Context, *pbrequest.StartFlinkClusterInK8S) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFlinkClusterInK8S not implemented")
}
func (UnimplementedEngineManageServer) StopFlinkClusterInK8S(context.Context, *pbrequest.StopFlinkClusterInK8S) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopFlinkClusterInK8S not implemented")
}
func (UnimplementedEngineManageServer) CreateNetworkBrokerInK8S(context.Context, *pbrequest.CreateFlinkClusterInK8S) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNetworkBrokerInK8S not implemented")
}
func (UnimplementedEngineManageServer) DeleteNetworkBrokerInK8S(context.Context, *pbrequest.DeleteNetworkBrokerInK8S) (*pbmodel.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetworkBrokerInK8S not implemented")
}
func (UnimplementedEngineManageServer) mustEmbedUnimplementedEngineManageServer() {}

// UnsafeEngineManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EngineManageServer will
// result in compilation errors.
type UnsafeEngineManageServer interface {
	mustEmbedUnimplementedEngineManageServer()
}

func RegisterEngineManageServer(s grpc.ServiceRegistrar, srv EngineManageServer) {
	s.RegisterService(&EngineManage_ServiceDesc, srv)
}

func _EngineManage_CreateFlinkClusterInK8S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.CreateFlinkClusterInK8S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineManageServer).CreateFlinkClusterInK8S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginecenter.EngineManage/CreateFlinkClusterInK8s",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineManageServer).CreateFlinkClusterInK8S(ctx, req.(*pbrequest.CreateFlinkClusterInK8S))
	}
	return interceptor(ctx, in, info, handler)
}

func _EngineManage_StartFlinkClusterInK8S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.StartFlinkClusterInK8S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineManageServer).StartFlinkClusterInK8S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginecenter.EngineManage/StartFlinkClusterInK8s",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineManageServer).StartFlinkClusterInK8S(ctx, req.(*pbrequest.StartFlinkClusterInK8S))
	}
	return interceptor(ctx, in, info, handler)
}

func _EngineManage_StopFlinkClusterInK8S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.StopFlinkClusterInK8S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineManageServer).StopFlinkClusterInK8S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginecenter.EngineManage/StopFlinkClusterInK8s",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineManageServer).StopFlinkClusterInK8S(ctx, req.(*pbrequest.StopFlinkClusterInK8S))
	}
	return interceptor(ctx, in, info, handler)
}

func _EngineManage_CreateNetworkBrokerInK8S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.CreateFlinkClusterInK8S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineManageServer).CreateNetworkBrokerInK8S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginecenter.EngineManage/CreateNetworkBrokerInK8s",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineManageServer).CreateNetworkBrokerInK8S(ctx, req.(*pbrequest.CreateFlinkClusterInK8S))
	}
	return interceptor(ctx, in, info, handler)
}

func _EngineManage_DeleteNetworkBrokerInK8S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.DeleteNetworkBrokerInK8S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineManageServer).DeleteNetworkBrokerInK8S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginecenter.EngineManage/DeleteNetworkBrokerInK8s",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineManageServer).DeleteNetworkBrokerInK8S(ctx, req.(*pbrequest.DeleteNetworkBrokerInK8S))
	}
	return interceptor(ctx, in, info, handler)
}

// EngineManage_ServiceDesc is the grpc.ServiceDesc for EngineManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EngineManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enginecenter.EngineManage",
	HandlerType: (*EngineManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFlinkClusterInK8s",
			Handler:    _EngineManage_CreateFlinkClusterInK8S_Handler,
		},
		{
			MethodName: "StartFlinkClusterInK8s",
			Handler:    _EngineManage_StartFlinkClusterInK8S_Handler,
		},
		{
			MethodName: "StopFlinkClusterInK8s",
			Handler:    _EngineManage_StopFlinkClusterInK8S_Handler,
		},
		{
			MethodName: "CreateNetworkBrokerInK8s",
			Handler:    _EngineManage_CreateNetworkBrokerInK8S_Handler,
		},
		{
			MethodName: "DeleteNetworkBrokerInK8s",
			Handler:    _EngineManage_DeleteNetworkBrokerInK8S_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/enginemanager/engine_manage.proto",
}