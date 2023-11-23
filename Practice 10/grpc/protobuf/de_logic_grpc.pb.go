// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: grpc/protobuf/de_logic.proto

package de_pb

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

const (
	DeLogic_Calc_FullMethodName = "/de_mult.de_logic/Calc"
)

// DeLogicClient is the client API for DeLogic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeLogicClient interface {
	Calc(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Result, error)
}

type deLogicClient struct {
	cc grpc.ClientConnInterface
}

func NewDeLogicClient(cc grpc.ClientConnInterface) DeLogicClient {
	return &deLogicClient{cc}
}

func (c *deLogicClient) Calc(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, DeLogic_Calc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeLogicServer is the server API for DeLogic service.
// All implementations must embed UnimplementedDeLogicServer
// for forward compatibility
type DeLogicServer interface {
	Calc(context.Context, *Input) (*Result, error)
	mustEmbedUnimplementedDeLogicServer()
}

// UnimplementedDeLogicServer must be embedded to have forward compatible implementations.
type UnimplementedDeLogicServer struct {
}

func (UnimplementedDeLogicServer) Calc(context.Context, *Input) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calc not implemented")
}
func (UnimplementedDeLogicServer) mustEmbedUnimplementedDeLogicServer() {}

// UnsafeDeLogicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeLogicServer will
// result in compilation errors.
type UnsafeDeLogicServer interface {
	mustEmbedUnimplementedDeLogicServer()
}

func RegisterDeLogicServer(s grpc.ServiceRegistrar, srv DeLogicServer) {
	s.RegisterService(&DeLogic_ServiceDesc, srv)
}

func _DeLogic_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeLogicServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeLogic_Calc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeLogicServer).Calc(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// DeLogic_ServiceDesc is the grpc.ServiceDesc for DeLogic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeLogic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "de_mult.de_logic",
	HandlerType: (*DeLogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _DeLogic_Calc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/protobuf/de_logic.proto",
}