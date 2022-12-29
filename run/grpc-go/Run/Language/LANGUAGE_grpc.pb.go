// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Language

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

// LANGUAGEClient is the client API for LANGUAGE service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LANGUAGEClient interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type lANGUAGEClient struct {
	cc grpc.ClientConnInterface
}

func NewLANGUAGEClient(cc grpc.ClientConnInterface) LANGUAGEClient {
	return &lANGUAGEClient{cc}
}

func (c *lANGUAGEClient) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/Run.Language.LANGUAGE/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lANGUAGEClient) Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Run.Language.LANGUAGE/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LANGUAGEServer is the server API for LANGUAGE service.
// All implementations must embed UnimplementedLANGUAGEServer
// for forward compatibility
type LANGUAGEServer interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedLANGUAGEServer()
}

// UnimplementedLANGUAGEServer must be embedded to have forward compatible implementations.
type UnimplementedLANGUAGEServer struct {
}

func (UnimplementedLANGUAGEServer) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedLANGUAGEServer) Helloworld(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedLANGUAGEServer) mustEmbedUnimplementedLANGUAGEServer() {}

// UnsafeLANGUAGEServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LANGUAGEServer will
// result in compilation errors.
type UnsafeLANGUAGEServer interface {
	mustEmbedUnimplementedLANGUAGEServer()
}

func RegisterLANGUAGEServer(s grpc.ServiceRegistrar, srv LANGUAGEServer) {
	s.RegisterService(&LANGUAGE_ServiceDesc, srv)
}

func _LANGUAGE_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LANGUAGEServer).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Run.Language.LANGUAGE/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LANGUAGEServer).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LANGUAGE_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LANGUAGEServer).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Run.Language.LANGUAGE/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LANGUAGEServer).Helloworld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LANGUAGE_ServiceDesc is the grpc.ServiceDesc for LANGUAGE service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LANGUAGE_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Run.Language.LANGUAGE",
	HandlerType: (*LANGUAGEServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _LANGUAGE_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _LANGUAGE_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LANGUAGE.proto",
}
