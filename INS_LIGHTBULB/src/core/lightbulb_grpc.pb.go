// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: lightbulb.proto

package lightblub

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

// LightBulbClient is the client API for LightBulb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LightBulbClient interface {
	HandleFrame(ctx context.Context, in *LightBulbFrame, opts ...grpc.CallOption) (*LightBulbRes, error)
}

type lightBulbClient struct {
	cc grpc.ClientConnInterface
}

func NewLightBulbClient(cc grpc.ClientConnInterface) LightBulbClient {
	return &lightBulbClient{cc}
}

func (c *lightBulbClient) HandleFrame(ctx context.Context, in *LightBulbFrame, opts ...grpc.CallOption) (*LightBulbRes, error) {
	out := new(LightBulbRes)
	err := c.cc.Invoke(ctx, "/lightblub.LightBulb/HandleFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LightBulbServer is the server API for LightBulb service.
// All implementations must embed UnimplementedLightBulbServer
// for forward compatibility
type LightBulbServer interface {
	HandleFrame(context.Context, *LightBulbFrame) (*LightBulbRes, error)
	mustEmbedUnimplementedLightBulbServer()
}

// UnimplementedLightBulbServer must be embedded to have forward compatible implementations.
type UnimplementedLightBulbServer struct {
}

func (UnimplementedLightBulbServer) HandleFrame(context.Context, *LightBulbFrame) (*LightBulbRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleFrame not implemented")
}
func (UnimplementedLightBulbServer) mustEmbedUnimplementedLightBulbServer() {}

// UnsafeLightBulbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LightBulbServer will
// result in compilation errors.
type UnsafeLightBulbServer interface {
	mustEmbedUnimplementedLightBulbServer()
}

func RegisterLightBulbServer(s grpc.ServiceRegistrar, srv LightBulbServer) {
	s.RegisterService(&LightBulb_ServiceDesc, srv)
}

func _LightBulb_HandleFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LightBulbFrame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightBulbServer).HandleFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightblub.LightBulb/HandleFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightBulbServer).HandleFrame(ctx, req.(*LightBulbFrame))
	}
	return interceptor(ctx, in, info, handler)
}

// LightBulb_ServiceDesc is the grpc.ServiceDesc for LightBulb service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LightBulb_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lightblub.LightBulb",
	HandlerType: (*LightBulbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleFrame",
			Handler:    _LightBulb_HandleFrame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lightbulb.proto",
}