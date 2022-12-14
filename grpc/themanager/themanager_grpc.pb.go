// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: themanager.proto

package themanager

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

// TheManagerClient is the client API for TheManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TheManagerClient interface {
	TestFunction(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*Message, error)
}

type theManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTheManagerClient(cc grpc.ClientConnInterface) TheManagerClient {
	return &theManagerClient{cc}
}

func (c *theManagerClient) TestFunction(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/themanager.TheManager/TestFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TheManagerServer is the server API for TheManager service.
// All implementations must embed UnimplementedTheManagerServer
// for forward compatibility
type TheManagerServer interface {
	TestFunction(context.Context, *TestMessage) (*Message, error)
	mustEmbedUnimplementedTheManagerServer()
}

// UnimplementedTheManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTheManagerServer struct {
}

func (UnimplementedTheManagerServer) TestFunction(context.Context, *TestMessage) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestFunction not implemented")
}
func (UnimplementedTheManagerServer) mustEmbedUnimplementedTheManagerServer() {}

// UnsafeTheManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TheManagerServer will
// result in compilation errors.
type UnsafeTheManagerServer interface {
	mustEmbedUnimplementedTheManagerServer()
}

func RegisterTheManagerServer(s grpc.ServiceRegistrar, srv TheManagerServer) {
	s.RegisterService(&TheManager_ServiceDesc, srv)
}

func _TheManager_TestFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TheManagerServer).TestFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/themanager.TheManager/TestFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TheManagerServer).TestFunction(ctx, req.(*TestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// TheManager_ServiceDesc is the grpc.ServiceDesc for TheManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TheManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "themanager.TheManager",
	HandlerType: (*TheManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestFunction",
			Handler:    _TheManager_TestFunction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "themanager.proto",
}
