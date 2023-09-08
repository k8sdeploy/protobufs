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

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*CreateAgentResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*CreateAgentResponse, error) {
	out := new(CreateAgentResponse)
	err := c.cc.Invoke(ctx, "/agent_service.AgentService/CreateAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgent not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_CreateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).CreateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent_service.AgentService/CreateAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).CreateAgent(ctx, req.(*CreateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent_service.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgent",
			Handler:    _AgentService_CreateAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/agent_service.proto",
}
