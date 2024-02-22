// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HookServiceClient is the client API for HookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HookServiceClient interface {
	UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error)
	UpdateDevDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error)
}

type hookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHookServiceClient(cc grpc.ClientConnInterface) HookServiceClient {
	return &hookServiceClient{cc}
}

func (c *hookServiceClient) UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error) {
	out := new(UpdateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/hook_service.HookService/UpdateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) UpdateDevDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error) {
	out := new(UpdateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/hook_service.HookService/UpdateDevDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HookServiceServer is the server API for HookService service.
// All implementations must embed UnimplementedHookServiceServer
// for forward compatibility
type HookServiceServer interface {
	UpdateDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error)
	UpdateDevDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error)
	mustEmbedUnimplementedHookServiceServer()
}

// UnimplementedHookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHookServiceServer struct {
}

func (UnimplementedHookServiceServer) UpdateDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeployment not implemented")
}
func (UnimplementedHookServiceServer) UpdateDevDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevDeployment not implemented")
}
func (UnimplementedHookServiceServer) mustEmbedUnimplementedHookServiceServer() {}

// UnsafeHookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HookServiceServer will
// result in compilation errors.
type UnsafeHookServiceServer interface {
	mustEmbedUnimplementedHookServiceServer()
}

func RegisterHookServiceServer(s grpc.ServiceRegistrar, srv HookServiceServer) {
	s.RegisterService(&HookService_ServiceDesc, srv)
}

func _HookService_UpdateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).UpdateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hook_service.HookService/UpdateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).UpdateDeployment(ctx, req.(*UpdateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_UpdateDevDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).UpdateDevDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hook_service.HookService/UpdateDevDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).UpdateDevDeployment(ctx, req.(*UpdateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HookService_ServiceDesc is the grpc.ServiceDesc for HookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hook_service.HookService",
	HandlerType: (*HookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDeployment",
			Handler:    _HookService_UpdateDeployment_Handler,
		},
		{
			MethodName: "UpdateDevDeployment",
			Handler:    _HookService_UpdateDevDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/hook_service.proto",
}