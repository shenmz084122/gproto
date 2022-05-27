// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/service/developer/flink_developer.proto

package pbsvcdeveloper

import (
	context "context"
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

// FlinkDeveloperClient is the client API for FlinkDeveloper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlinkDeveloperClient interface {
	SubmitFlinkJob(ctx context.Context, in *pbrequest.SubmitFlinkJob, opts ...grpc.CallOption) (*pbresponse.SubmitFlinkJob, error)
	SubmitFlinkJobInteractive(ctx context.Context, opts ...grpc.CallOption) (FlinkDeveloper_SubmitFlinkJobInteractiveClient, error)
	ValidateFlinkJob(ctx context.Context, in *pbrequest.ValidateFlinkJob, opts ...grpc.CallOption) (*pbresponse.ValidateFlinkJob, error)
	ValidateFlinkJobV2(ctx context.Context, in *pbrequest.ValidateFlinkJobV2, opts ...grpc.CallOption) (*pbresponse.ValidateFlinkJob, error)
}

type flinkDeveloperClient struct {
	cc grpc.ClientConnInterface
}

func NewFlinkDeveloperClient(cc grpc.ClientConnInterface) FlinkDeveloperClient {
	return &flinkDeveloperClient{cc}
}

func (c *flinkDeveloperClient) SubmitFlinkJob(ctx context.Context, in *pbrequest.SubmitFlinkJob, opts ...grpc.CallOption) (*pbresponse.SubmitFlinkJob, error) {
	out := new(pbresponse.SubmitFlinkJob)
	err := c.cc.Invoke(ctx, "/developer.FlinkDeveloper/SubmitFlinkJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flinkDeveloperClient) SubmitFlinkJobInteractive(ctx context.Context, opts ...grpc.CallOption) (FlinkDeveloper_SubmitFlinkJobInteractiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlinkDeveloper_ServiceDesc.Streams[0], "/developer.FlinkDeveloper/SubmitFlinkJobInteractive", opts...)
	if err != nil {
		return nil, err
	}
	x := &flinkDeveloperSubmitFlinkJobInteractiveClient{stream}
	return x, nil
}

type FlinkDeveloper_SubmitFlinkJobInteractiveClient interface {
	Send(*pbrequest.SubmitFlinkJobInteractive) error
	Recv() (*pbresponse.SubmitFlinkJobInteractive, error)
	grpc.ClientStream
}

type flinkDeveloperSubmitFlinkJobInteractiveClient struct {
	grpc.ClientStream
}

func (x *flinkDeveloperSubmitFlinkJobInteractiveClient) Send(m *pbrequest.SubmitFlinkJobInteractive) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flinkDeveloperSubmitFlinkJobInteractiveClient) Recv() (*pbresponse.SubmitFlinkJobInteractive, error) {
	m := new(pbresponse.SubmitFlinkJobInteractive)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flinkDeveloperClient) ValidateFlinkJob(ctx context.Context, in *pbrequest.ValidateFlinkJob, opts ...grpc.CallOption) (*pbresponse.ValidateFlinkJob, error) {
	out := new(pbresponse.ValidateFlinkJob)
	err := c.cc.Invoke(ctx, "/developer.FlinkDeveloper/ValidateFlinkJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flinkDeveloperClient) ValidateFlinkJobV2(ctx context.Context, in *pbrequest.ValidateFlinkJobV2, opts ...grpc.CallOption) (*pbresponse.ValidateFlinkJob, error) {
	out := new(pbresponse.ValidateFlinkJob)
	err := c.cc.Invoke(ctx, "/developer.FlinkDeveloper/ValidateFlinkJob_v2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlinkDeveloperServer is the server API for FlinkDeveloper service.
// All implementations must embed UnimplementedFlinkDeveloperServer
// for forward compatibility
type FlinkDeveloperServer interface {
	SubmitFlinkJob(context.Context, *pbrequest.SubmitFlinkJob) (*pbresponse.SubmitFlinkJob, error)
	SubmitFlinkJobInteractive(FlinkDeveloper_SubmitFlinkJobInteractiveServer) error
	ValidateFlinkJob(context.Context, *pbrequest.ValidateFlinkJob) (*pbresponse.ValidateFlinkJob, error)
	ValidateFlinkJobV2(context.Context, *pbrequest.ValidateFlinkJobV2) (*pbresponse.ValidateFlinkJob, error)
	mustEmbedUnimplementedFlinkDeveloperServer()
}

// UnimplementedFlinkDeveloperServer must be embedded to have forward compatible implementations.
type UnimplementedFlinkDeveloperServer struct {
}

func (UnimplementedFlinkDeveloperServer) SubmitFlinkJob(context.Context, *pbrequest.SubmitFlinkJob) (*pbresponse.SubmitFlinkJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitFlinkJob not implemented")
}
func (UnimplementedFlinkDeveloperServer) SubmitFlinkJobInteractive(FlinkDeveloper_SubmitFlinkJobInteractiveServer) error {
	return status.Errorf(codes.Unimplemented, "method SubmitFlinkJobInteractive not implemented")
}
func (UnimplementedFlinkDeveloperServer) ValidateFlinkJob(context.Context, *pbrequest.ValidateFlinkJob) (*pbresponse.ValidateFlinkJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateFlinkJob not implemented")
}
func (UnimplementedFlinkDeveloperServer) ValidateFlinkJobV2(context.Context, *pbrequest.ValidateFlinkJobV2) (*pbresponse.ValidateFlinkJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateFlinkJobV2 not implemented")
}
func (UnimplementedFlinkDeveloperServer) mustEmbedUnimplementedFlinkDeveloperServer() {}

// UnsafeFlinkDeveloperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlinkDeveloperServer will
// result in compilation errors.
type UnsafeFlinkDeveloperServer interface {
	mustEmbedUnimplementedFlinkDeveloperServer()
}

func RegisterFlinkDeveloperServer(s grpc.ServiceRegistrar, srv FlinkDeveloperServer) {
	s.RegisterService(&FlinkDeveloper_ServiceDesc, srv)
}

func _FlinkDeveloper_SubmitFlinkJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.SubmitFlinkJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlinkDeveloperServer).SubmitFlinkJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/developer.FlinkDeveloper/SubmitFlinkJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlinkDeveloperServer).SubmitFlinkJob(ctx, req.(*pbrequest.SubmitFlinkJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlinkDeveloper_SubmitFlinkJobInteractive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlinkDeveloperServer).SubmitFlinkJobInteractive(&flinkDeveloperSubmitFlinkJobInteractiveServer{stream})
}

type FlinkDeveloper_SubmitFlinkJobInteractiveServer interface {
	Send(*pbresponse.SubmitFlinkJobInteractive) error
	Recv() (*pbrequest.SubmitFlinkJobInteractive, error)
	grpc.ServerStream
}

type flinkDeveloperSubmitFlinkJobInteractiveServer struct {
	grpc.ServerStream
}

func (x *flinkDeveloperSubmitFlinkJobInteractiveServer) Send(m *pbresponse.SubmitFlinkJobInteractive) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flinkDeveloperSubmitFlinkJobInteractiveServer) Recv() (*pbrequest.SubmitFlinkJobInteractive, error) {
	m := new(pbrequest.SubmitFlinkJobInteractive)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlinkDeveloper_ValidateFlinkJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.ValidateFlinkJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlinkDeveloperServer).ValidateFlinkJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/developer.FlinkDeveloper/ValidateFlinkJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlinkDeveloperServer).ValidateFlinkJob(ctx, req.(*pbrequest.ValidateFlinkJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlinkDeveloper_ValidateFlinkJobV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbrequest.ValidateFlinkJobV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlinkDeveloperServer).ValidateFlinkJobV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/developer.FlinkDeveloper/ValidateFlinkJob_v2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlinkDeveloperServer).ValidateFlinkJobV2(ctx, req.(*pbrequest.ValidateFlinkJobV2))
	}
	return interceptor(ctx, in, info, handler)
}

// FlinkDeveloper_ServiceDesc is the grpc.ServiceDesc for FlinkDeveloper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlinkDeveloper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "developer.FlinkDeveloper",
	HandlerType: (*FlinkDeveloperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitFlinkJob",
			Handler:    _FlinkDeveloper_SubmitFlinkJob_Handler,
		},
		{
			MethodName: "ValidateFlinkJob",
			Handler:    _FlinkDeveloper_ValidateFlinkJob_Handler,
		},
		{
			MethodName: "ValidateFlinkJob_v2",
			Handler:    _FlinkDeveloper_ValidateFlinkJobV2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitFlinkJobInteractive",
			Handler:       _FlinkDeveloper_SubmitFlinkJobInteractive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service/developer/flink_developer.proto",
}
