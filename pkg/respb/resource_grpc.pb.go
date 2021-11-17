// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package respb

import (
	context "context"
	model "github.com/DataWorkbench/gproto/pkg/model"
	request "github.com/DataWorkbench/gproto/pkg/request"
	response "github.com/DataWorkbench/gproto/pkg/response"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ResourceClient is the client API for Resource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceClient interface {
	ReUploadFile(ctx context.Context, opts ...grpc.CallOption) (Resource_ReUploadFileClient, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (Resource_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *request.DownloadFile, opts ...grpc.CallOption) (Resource_DownloadFileClient, error)
	DescribeFile(ctx context.Context, in *request.DescribeFile, opts ...grpc.CallOption) (*model.Resource, error)
	ListResources(ctx context.Context, in *request.ListResources, opts ...grpc.CallOption) (*response.ListResources, error)
	UpdateResource(ctx context.Context, in *request.UpdateResource, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DeleteResources(ctx context.Context, in *request.DeleteResources, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DeleteSpaces(ctx context.Context, in *request.DeleteWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	RenameFlinkStatePath(ctx context.Context, in *request.RenameFlinkStatePath, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DeleteFlinkState(ctx context.Context, in *request.DeleteFlinkState, opts ...grpc.CallOption) (*model.EmptyStruct, error)
}

type resourceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceClient(cc grpc.ClientConnInterface) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) ReUploadFile(ctx context.Context, opts ...grpc.CallOption) (Resource_ReUploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Resource_serviceDesc.Streams[0], "/resource.Resource/ReUploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceReUploadFileClient{stream}
	return x, nil
}

type Resource_ReUploadFileClient interface {
	Send(*ReUploadFileRequest) error
	CloseAndRecv() (*model.EmptyStruct, error)
	grpc.ClientStream
}

type resourceReUploadFileClient struct {
	grpc.ClientStream
}

func (x *resourceReUploadFileClient) Send(m *ReUploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *resourceReUploadFileClient) CloseAndRecv() (*model.EmptyStruct, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(model.EmptyStruct)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resourceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (Resource_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Resource_serviceDesc.Streams[1], "/resource.Resource/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceUploadFileClient{stream}
	return x, nil
}

type Resource_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*response.UploadFile, error)
	grpc.ClientStream
}

type resourceUploadFileClient struct {
	grpc.ClientStream
}

func (x *resourceUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *resourceUploadFileClient) CloseAndRecv() (*response.UploadFile, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(response.UploadFile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resourceClient) DownloadFile(ctx context.Context, in *request.DownloadFile, opts ...grpc.CallOption) (Resource_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Resource_serviceDesc.Streams[2], "/resource.Resource/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Resource_DownloadFileClient interface {
	Recv() (*response.DownloadFile, error)
	grpc.ClientStream
}

type resourceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *resourceDownloadFileClient) Recv() (*response.DownloadFile, error) {
	m := new(response.DownloadFile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resourceClient) DescribeFile(ctx context.Context, in *request.DescribeFile, opts ...grpc.CallOption) (*model.Resource, error) {
	out := new(model.Resource)
	err := c.cc.Invoke(ctx, "/resource.Resource/DescribeFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) ListResources(ctx context.Context, in *request.ListResources, opts ...grpc.CallOption) (*response.ListResources, error) {
	out := new(response.ListResources)
	err := c.cc.Invoke(ctx, "/resource.Resource/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) UpdateResource(ctx context.Context, in *request.UpdateResource, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/resource.Resource/UpdateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) DeleteResources(ctx context.Context, in *request.DeleteResources, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/resource.Resource/DeleteResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) DeleteSpaces(ctx context.Context, in *request.DeleteWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/resource.Resource/DeleteSpaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) RenameFlinkStatePath(ctx context.Context, in *request.RenameFlinkStatePath, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/resource.Resource/RenameFlinkStatePath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) DeleteFlinkState(ctx context.Context, in *request.DeleteFlinkState, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/resource.Resource/DeleteFlinkState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServer is the server API for Resource service.
// All implementations must embed UnimplementedResourceServer
// for forward compatibility
type ResourceServer interface {
	ReUploadFile(Resource_ReUploadFileServer) error
	UploadFile(Resource_UploadFileServer) error
	DownloadFile(*request.DownloadFile, Resource_DownloadFileServer) error
	DescribeFile(context.Context, *request.DescribeFile) (*model.Resource, error)
	ListResources(context.Context, *request.ListResources) (*response.ListResources, error)
	UpdateResource(context.Context, *request.UpdateResource) (*model.EmptyStruct, error)
	DeleteResources(context.Context, *request.DeleteResources) (*model.EmptyStruct, error)
	DeleteSpaces(context.Context, *request.DeleteWorkspaces) (*model.EmptyStruct, error)
	RenameFlinkStatePath(context.Context, *request.RenameFlinkStatePath) (*model.EmptyStruct, error)
	DeleteFlinkState(context.Context, *request.DeleteFlinkState) (*model.EmptyStruct, error)
	mustEmbedUnimplementedResourceServer()
}

// UnimplementedResourceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServer struct {
}

func (UnimplementedResourceServer) ReUploadFile(Resource_ReUploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method ReUploadFile not implemented")
}
func (UnimplementedResourceServer) UploadFile(Resource_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedResourceServer) DownloadFile(*request.DownloadFile, Resource_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedResourceServer) DescribeFile(context.Context, *request.DescribeFile) (*model.Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeFile not implemented")
}
func (UnimplementedResourceServer) ListResources(context.Context, *request.ListResources) (*response.ListResources, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedResourceServer) UpdateResource(context.Context, *request.UpdateResource) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (UnimplementedResourceServer) DeleteResources(context.Context, *request.DeleteResources) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResources not implemented")
}
func (UnimplementedResourceServer) DeleteSpaces(context.Context, *request.DeleteWorkspaces) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpaces not implemented")
}
func (UnimplementedResourceServer) RenameFlinkStatePath(context.Context, *request.RenameFlinkStatePath) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFlinkStatePath not implemented")
}
func (UnimplementedResourceServer) DeleteFlinkState(context.Context, *request.DeleteFlinkState) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFlinkState not implemented")
}
func (UnimplementedResourceServer) mustEmbedUnimplementedResourceServer() {}

// UnsafeResourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServer will
// result in compilation errors.
type UnsafeResourceServer interface {
	mustEmbedUnimplementedResourceServer()
}

func RegisterResourceServer(s grpc.ServiceRegistrar, srv ResourceServer) {
	s.RegisterService(&_Resource_serviceDesc, srv)
}

func _Resource_ReUploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResourceServer).ReUploadFile(&resourceReUploadFileServer{stream})
}

type Resource_ReUploadFileServer interface {
	SendAndClose(*model.EmptyStruct) error
	Recv() (*ReUploadFileRequest, error)
	grpc.ServerStream
}

type resourceReUploadFileServer struct {
	grpc.ServerStream
}

func (x *resourceReUploadFileServer) SendAndClose(m *model.EmptyStruct) error {
	return x.ServerStream.SendMsg(m)
}

func (x *resourceReUploadFileServer) Recv() (*ReUploadFileRequest, error) {
	m := new(ReUploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Resource_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResourceServer).UploadFile(&resourceUploadFileServer{stream})
}

type Resource_UploadFileServer interface {
	SendAndClose(*response.UploadFile) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type resourceUploadFileServer struct {
	grpc.ServerStream
}

func (x *resourceUploadFileServer) SendAndClose(m *response.UploadFile) error {
	return x.ServerStream.SendMsg(m)
}

func (x *resourceUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Resource_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(request.DownloadFile)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceServer).DownloadFile(m, &resourceDownloadFileServer{stream})
}

type Resource_DownloadFileServer interface {
	Send(*response.DownloadFile) error
	grpc.ServerStream
}

type resourceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *resourceDownloadFileServer) Send(m *response.DownloadFile) error {
	return x.ServerStream.SendMsg(m)
}

func _Resource_DescribeFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DescribeFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DescribeFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/DescribeFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DescribeFile(ctx, req.(*request.DescribeFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListResources)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).ListResources(ctx, req.(*request.ListResources))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateResource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).UpdateResource(ctx, req.(*request.UpdateResource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_DeleteResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteResources)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DeleteResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/DeleteResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DeleteResources(ctx, req.(*request.DeleteResources))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_DeleteSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DeleteSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/DeleteSpaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DeleteSpaces(ctx, req.(*request.DeleteWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_RenameFlinkStatePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.RenameFlinkStatePath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).RenameFlinkStatePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/RenameFlinkStatePath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).RenameFlinkStatePath(ctx, req.(*request.RenameFlinkStatePath))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_DeleteFlinkState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteFlinkState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DeleteFlinkState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/DeleteFlinkState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DeleteFlinkState(ctx, req.(*request.DeleteFlinkState))
	}
	return interceptor(ctx, in, info, handler)
}

var _Resource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resource.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeFile",
			Handler:    _Resource_DescribeFile_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _Resource_ListResources_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _Resource_UpdateResource_Handler,
		},
		{
			MethodName: "DeleteResources",
			Handler:    _Resource_DeleteResources_Handler,
		},
		{
			MethodName: "DeleteSpaces",
			Handler:    _Resource_DeleteSpaces_Handler,
		},
		{
			MethodName: "RenameFlinkStatePath",
			Handler:    _Resource_RenameFlinkStatePath_Handler,
		},
		{
			MethodName: "DeleteFlinkState",
			Handler:    _Resource_DeleteFlinkState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReUploadFile",
			Handler:       _Resource_ReUploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _Resource_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _Resource_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/resource.proto",
}