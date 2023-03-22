// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: todo/v1/todo.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TodoService_CreateTodoTask_FullMethodName = "/todo.v1.TodoService/CreateTodoTask"
	TodoService_ListTodoTasks_FullMethodName  = "/todo.v1.TodoService/ListTodoTasks"
	TodoService_GetTodoTask_FullMethodName    = "/todo.v1.TodoService/GetTodoTask"
	TodoService_UpdateTodoTask_FullMethodName = "/todo.v1.TodoService/UpdateTodoTask"
	TodoService_DeleteTodoTask_FullMethodName = "/todo.v1.TodoService/DeleteTodoTask"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodoTask(ctx context.Context, in *CreateTodoTaskRequest, opts ...grpc.CallOption) (*CreateTodoTaskResponse, error)
	ListTodoTasks(ctx context.Context, in *ListTodoTasksRequest, opts ...grpc.CallOption) (*ListTodoTasksResponse, error)
	GetTodoTask(ctx context.Context, in *GetTodoTaskRequest, opts ...grpc.CallOption) (*GetTodoTaskResponse, error)
	UpdateTodoTask(ctx context.Context, in *UpdateTodoTaskRequest, opts ...grpc.CallOption) (*UpdateTodoTaskResponse, error)
	DeleteTodoTask(ctx context.Context, in *DeleteTodoTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodoTask(ctx context.Context, in *CreateTodoTaskRequest, opts ...grpc.CallOption) (*CreateTodoTaskResponse, error) {
	out := new(CreateTodoTaskResponse)
	err := c.cc.Invoke(ctx, TodoService_CreateTodoTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodoTasks(ctx context.Context, in *ListTodoTasksRequest, opts ...grpc.CallOption) (*ListTodoTasksResponse, error) {
	out := new(ListTodoTasksResponse)
	err := c.cc.Invoke(ctx, TodoService_ListTodoTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodoTask(ctx context.Context, in *GetTodoTaskRequest, opts ...grpc.CallOption) (*GetTodoTaskResponse, error) {
	out := new(GetTodoTaskResponse)
	err := c.cc.Invoke(ctx, TodoService_GetTodoTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodoTask(ctx context.Context, in *UpdateTodoTaskRequest, opts ...grpc.CallOption) (*UpdateTodoTaskResponse, error) {
	out := new(UpdateTodoTaskResponse)
	err := c.cc.Invoke(ctx, TodoService_UpdateTodoTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodoTask(ctx context.Context, in *DeleteTodoTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TodoService_DeleteTodoTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateTodoTask(context.Context, *CreateTodoTaskRequest) (*CreateTodoTaskResponse, error)
	ListTodoTasks(context.Context, *ListTodoTasksRequest) (*ListTodoTasksResponse, error)
	GetTodoTask(context.Context, *GetTodoTaskRequest) (*GetTodoTaskResponse, error)
	UpdateTodoTask(context.Context, *UpdateTodoTaskRequest) (*UpdateTodoTaskResponse, error)
	DeleteTodoTask(context.Context, *DeleteTodoTaskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateTodoTask(context.Context, *CreateTodoTaskRequest) (*CreateTodoTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodoTask not implemented")
}
func (UnimplementedTodoServiceServer) ListTodoTasks(context.Context, *ListTodoTasksRequest) (*ListTodoTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodoTasks not implemented")
}
func (UnimplementedTodoServiceServer) GetTodoTask(context.Context, *GetTodoTaskRequest) (*GetTodoTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoTask not implemented")
}
func (UnimplementedTodoServiceServer) UpdateTodoTask(context.Context, *UpdateTodoTaskRequest) (*UpdateTodoTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodoTask not implemented")
}
func (UnimplementedTodoServiceServer) DeleteTodoTask(context.Context, *DeleteTodoTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodoTask not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateTodoTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodoTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_CreateTodoTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodoTask(ctx, req.(*CreateTodoTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodoTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodoTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_ListTodoTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodoTasks(ctx, req.(*ListTodoTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodoTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodoTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_GetTodoTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodoTask(ctx, req.(*GetTodoTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodoTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodoTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_UpdateTodoTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodoTask(ctx, req.(*UpdateTodoTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodoTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodoTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_DeleteTodoTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodoTask(ctx, req.(*DeleteTodoTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.v1.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodoTask",
			Handler:    _TodoService_CreateTodoTask_Handler,
		},
		{
			MethodName: "ListTodoTasks",
			Handler:    _TodoService_ListTodoTasks_Handler,
		},
		{
			MethodName: "GetTodoTask",
			Handler:    _TodoService_GetTodoTask_Handler,
		},
		{
			MethodName: "UpdateTodoTask",
			Handler:    _TodoService_UpdateTodoTask_Handler,
		},
		{
			MethodName: "DeleteTodoTask",
			Handler:    _TodoService_DeleteTodoTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/v1/todo.proto",
}
