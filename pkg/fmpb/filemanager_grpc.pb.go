// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fmpb

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

// FileManagerClient is the client API for FileManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileManagerClient interface {
	CreateDir(ctx context.Context, in *CreateDirRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileManager_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileManager_DownloadFileClient, error)
	DescribeFile(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*FileInfoResponse, error)
	ListFiles(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	CheckExist(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
}

type fileManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewFileManagerClient(cc grpc.ClientConnInterface) FileManagerClient {
	return &fileManagerClient{cc}
}

func (c *fileManagerClient) CreateDir(ctx context.Context, in *CreateDirRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/CreateDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileManager_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileManager_serviceDesc.Streams[0], "/fmpb.FileManager/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerUploadFileClient{stream}
	return x, nil
}

type FileManager_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*model.EmptyStruct, error)
	grpc.ClientStream
}

type fileManagerUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileManagerUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileManagerUploadFileClient) CloseAndRecv() (*model.EmptyStruct, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(model.EmptyStruct)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileManagerClient) DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileManager_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileManager_serviceDesc.Streams[1], "/fmpb.FileManager/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileManager_DownloadFileClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type fileManagerDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileManagerDownloadFileClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileManagerClient) DescribeFile(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*FileInfoResponse, error) {
	out := new(FileInfoResponse)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/DescribeFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) ListFiles(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/DeleteSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) CheckExist(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/CheckExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileManagerServer is the server API for FileManager service.
// All implementations must embed UnimplementedFileManagerServer
// for forward compatibility
type FileManagerServer interface {
	CreateDir(context.Context, *CreateDirRequest) (*model.EmptyStruct, error)
	UploadFile(FileManager_UploadFileServer) error
	DownloadFile(*DownloadRequest, FileManager_DownloadFileServer) error
	DescribeFile(context.Context, *DescribeRequest) (*FileInfoResponse, error)
	ListFiles(context.Context, *ListRequest) (*ListResponse, error)
	UpdateFile(context.Context, *UpdateFileRequest) (*model.EmptyStruct, error)
	Delete(context.Context, *DeleteRequest) (*model.EmptyStruct, error)
	DeleteSpace(context.Context, *DeleteSpaceRequest) (*model.EmptyStruct, error)
	CheckExist(context.Context, *CheckRequest) (*model.EmptyStruct, error)
	mustEmbedUnimplementedFileManagerServer()
}

// UnimplementedFileManagerServer must be embedded to have forward compatible implementations.
type UnimplementedFileManagerServer struct {
}

func (UnimplementedFileManagerServer) CreateDir(context.Context, *CreateDirRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDir not implemented")
}
func (UnimplementedFileManagerServer) UploadFile(FileManager_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileManagerServer) DownloadFile(*DownloadRequest, FileManager_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileManagerServer) DescribeFile(context.Context, *DescribeRequest) (*FileInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeFile not implemented")
}
func (UnimplementedFileManagerServer) ListFiles(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileManagerServer) UpdateFile(context.Context, *UpdateFileRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFileManagerServer) Delete(context.Context, *DeleteRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFileManagerServer) DeleteSpace(context.Context, *DeleteSpaceRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}
func (UnimplementedFileManagerServer) CheckExist(context.Context, *CheckRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExist not implemented")
}
func (UnimplementedFileManagerServer) mustEmbedUnimplementedFileManagerServer() {}

// UnsafeFileManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileManagerServer will
// result in compilation errors.
type UnsafeFileManagerServer interface {
	mustEmbedUnimplementedFileManagerServer()
}

func RegisterFileManagerServer(s grpc.ServiceRegistrar, srv FileManagerServer) {
	s.RegisterService(&_FileManager_serviceDesc, srv)
}

func _FileManager_CreateDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).CreateDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/CreateDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).CreateDir(ctx, req.(*CreateDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileManagerServer).UploadFile(&fileManagerUploadFileServer{stream})
}

type FileManager_UploadFileServer interface {
	SendAndClose(*model.EmptyStruct) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type fileManagerUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileManagerUploadFileServer) SendAndClose(m *model.EmptyStruct) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileManagerUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileManager_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileManagerServer).DownloadFile(m, &fileManagerDownloadFileServer{stream})
}

type FileManager_DownloadFileServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type fileManagerDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileManagerDownloadFileServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileManager_DescribeFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).DescribeFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/DescribeFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).DescribeFile(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).ListFiles(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).UpdateFile(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_DeleteSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).DeleteSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/DeleteSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).DeleteSpace(ctx, req.(*DeleteSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_CheckExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).CheckExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/CheckExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).CheckExist(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fmpb.FileManager",
	HandlerType: (*FileManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDir",
			Handler:    _FileManager_CreateDir_Handler,
		},
		{
			MethodName: "DescribeFile",
			Handler:    _FileManager_DescribeFile_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _FileManager_ListFiles_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _FileManager_UpdateFile_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FileManager_Delete_Handler,
		},
		{
			MethodName: "DeleteSpace",
			Handler:    _FileManager_DeleteSpace_Handler,
		},
		{
			MethodName: "CheckExist",
			Handler:    _FileManager_CheckExist_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileManager_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileManager_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/filemanager.proto",
}
