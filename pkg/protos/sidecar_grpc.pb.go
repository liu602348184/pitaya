// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// SidecarClient is the client API for Sidecar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SidecarClient interface {
	ListenRPC(ctx context.Context, opts ...grpc.CallOption) (Sidecar_ListenRPCClient, error)
	SendRPC(ctx context.Context, in *RequestTo, opts ...grpc.CallOption) (*Response, error)
	StartPitaya(ctx context.Context, in *StartPitayaRequest, opts ...grpc.CallOption) (*Error, error)
	StopPitaya(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Error, error)
}

type sidecarClient struct {
	cc grpc.ClientConnInterface
}

func NewSidecarClient(cc grpc.ClientConnInterface) SidecarClient {
	return &sidecarClient{cc}
}

func (c *sidecarClient) ListenRPC(ctx context.Context, opts ...grpc.CallOption) (Sidecar_ListenRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sidecar_ServiceDesc.Streams[0], "/protos.Sidecar/ListenRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &sidecarListenRPCClient{stream}
	return x, nil
}

type Sidecar_ListenRPCClient interface {
	Send(*RPCResponse) error
	Recv() (*SidecarRequest, error)
	grpc.ClientStream
}

type sidecarListenRPCClient struct {
	grpc.ClientStream
}

func (x *sidecarListenRPCClient) Send(m *RPCResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sidecarListenRPCClient) Recv() (*SidecarRequest, error) {
	m := new(SidecarRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sidecarClient) SendRPC(ctx context.Context, in *RequestTo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.Sidecar/SendRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidecarClient) StartPitaya(ctx context.Context, in *StartPitayaRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/protos.Sidecar/StartPitaya", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sidecarClient) StopPitaya(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/protos.Sidecar/StopPitaya", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SidecarServer is the server API for Sidecar service.
// All implementations should embed UnimplementedSidecarServer
// for forward compatibility
type SidecarServer interface {
	ListenRPC(Sidecar_ListenRPCServer) error
	SendRPC(context.Context, *RequestTo) (*Response, error)
	StartPitaya(context.Context, *StartPitayaRequest) (*Error, error)
	StopPitaya(context.Context, *emptypb.Empty) (*Error, error)
}

// UnimplementedSidecarServer should be embedded to have forward compatible implementations.
type UnimplementedSidecarServer struct {
}

func (UnimplementedSidecarServer) ListenRPC(Sidecar_ListenRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenRPC not implemented")
}
func (UnimplementedSidecarServer) SendRPC(context.Context, *RequestTo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRPC not implemented")
}
func (UnimplementedSidecarServer) StartPitaya(context.Context, *StartPitayaRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPitaya not implemented")
}
func (UnimplementedSidecarServer) StopPitaya(context.Context, *emptypb.Empty) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPitaya not implemented")
}

// UnsafeSidecarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SidecarServer will
// result in compilation errors.
type UnsafeSidecarServer interface {
	mustEmbedUnimplementedSidecarServer()
}

func RegisterSidecarServer(s grpc.ServiceRegistrar, srv SidecarServer) {
	s.RegisterService(&Sidecar_ServiceDesc, srv)
}

func _Sidecar_ListenRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SidecarServer).ListenRPC(&sidecarListenRPCServer{stream})
}

type Sidecar_ListenRPCServer interface {
	Send(*SidecarRequest) error
	Recv() (*RPCResponse, error)
	grpc.ServerStream
}

type sidecarListenRPCServer struct {
	grpc.ServerStream
}

func (x *sidecarListenRPCServer) Send(m *SidecarRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sidecarListenRPCServer) Recv() (*RPCResponse, error) {
	m := new(RPCResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sidecar_SendRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).SendRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sidecar/SendRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).SendRPC(ctx, req.(*RequestTo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sidecar_StartPitaya_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPitayaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).StartPitaya(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sidecar/StartPitaya",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).StartPitaya(ctx, req.(*StartPitayaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sidecar_StopPitaya_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SidecarServer).StopPitaya(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Sidecar/StopPitaya",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SidecarServer).StopPitaya(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Sidecar_ServiceDesc is the grpc.ServiceDesc for Sidecar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sidecar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Sidecar",
	HandlerType: (*SidecarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRPC",
			Handler:    _Sidecar_SendRPC_Handler,
		},
		{
			MethodName: "StartPitaya",
			Handler:    _Sidecar_StartPitaya_Handler,
		},
		{
			MethodName: "StopPitaya",
			Handler:    _Sidecar_StopPitaya_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenRPC",
			Handler:       _Sidecar_ListenRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sidecar.proto",
}
