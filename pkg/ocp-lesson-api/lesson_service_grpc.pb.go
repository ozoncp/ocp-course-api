// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: lesson_service.proto

package ocp_lesson_api

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

// OcpLessonApiClient is the client API for OcpLessonApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpLessonApiClient interface {
	ListLessonsV1(ctx context.Context, in *ListLessonsV1Request, opts ...grpc.CallOption) (*ListLessonsV1Response, error)
	DescribeLessonV1(ctx context.Context, in *DescribeLessonV1Request, opts ...grpc.CallOption) (*DescribeLessonV1Response, error)
	CreateLessonV1(ctx context.Context, in *CreateLessonV1Request, opts ...grpc.CallOption) (*CreateLessonV1Response, error)
	RemoveLessonV1(ctx context.Context, in *RemoveLessonV1Request, opts ...grpc.CallOption) (*RemoveLessonV1Response, error)
	UpdateLessonV1(ctx context.Context, in *UpdateLessonV1Request, opts ...grpc.CallOption) (*UpdateLessonV1Response, error)
	MultiCreateLessonV1(ctx context.Context, in *MultiCreateLessonV1Request, opts ...grpc.CallOption) (*MultiCreateLessonV1Response, error)
}

type ocpLessonApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpLessonApiClient(cc grpc.ClientConnInterface) OcpLessonApiClient {
	return &ocpLessonApiClient{cc}
}

func (c *ocpLessonApiClient) ListLessonsV1(ctx context.Context, in *ListLessonsV1Request, opts ...grpc.CallOption) (*ListLessonsV1Response, error) {
	out := new(ListLessonsV1Response)
	err := c.cc.Invoke(ctx, "/ocp.lesson.api.OcpLessonApi/ListLessonsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpLessonApiClient) DescribeLessonV1(ctx context.Context, in *DescribeLessonV1Request, opts ...grpc.CallOption) (*DescribeLessonV1Response, error) {
	out := new(DescribeLessonV1Response)
	err := c.cc.Invoke(ctx, "/ocp.lesson.api.OcpLessonApi/DescribeLessonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpLessonApiClient) CreateLessonV1(ctx context.Context, in *CreateLessonV1Request, opts ...grpc.CallOption) (*CreateLessonV1Response, error) {
	out := new(CreateLessonV1Response)
	err := c.cc.Invoke(ctx, "/ocp.lesson.api.OcpLessonApi/CreateLessonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpLessonApiClient) RemoveLessonV1(ctx context.Context, in *RemoveLessonV1Request, opts ...grpc.CallOption) (*RemoveLessonV1Response, error) {
	out := new(RemoveLessonV1Response)
	err := c.cc.Invoke(ctx, "/ocp.lesson.api.OcpLessonApi/RemoveLessonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpLessonApiClient) UpdateLessonV1(ctx context.Context, in *UpdateLessonV1Request, opts ...grpc.CallOption) (*UpdateLessonV1Response, error) {
	out := new(UpdateLessonV1Response)
	err := c.cc.Invoke(ctx, "/ocp.lesson.api.OcpLessonApi/UpdateLessonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpLessonApiClient) MultiCreateLessonV1(ctx context.Context, in *MultiCreateLessonV1Request, opts ...grpc.CallOption) (*MultiCreateLessonV1Response, error) {
	out := new(MultiCreateLessonV1Response)
	err := c.cc.Invoke(ctx, "/ocp.lesson.api.OcpLessonApi/MultiCreateLessonV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpLessonApiServer is the server API for OcpLessonApi service.
// All implementations must embed UnimplementedOcpLessonApiServer
// for forward compatibility
type OcpLessonApiServer interface {
	ListLessonsV1(context.Context, *ListLessonsV1Request) (*ListLessonsV1Response, error)
	DescribeLessonV1(context.Context, *DescribeLessonV1Request) (*DescribeLessonV1Response, error)
	CreateLessonV1(context.Context, *CreateLessonV1Request) (*CreateLessonV1Response, error)
	RemoveLessonV1(context.Context, *RemoveLessonV1Request) (*RemoveLessonV1Response, error)
	UpdateLessonV1(context.Context, *UpdateLessonV1Request) (*UpdateLessonV1Response, error)
	MultiCreateLessonV1(context.Context, *MultiCreateLessonV1Request) (*MultiCreateLessonV1Response, error)
	mustEmbedUnimplementedOcpLessonApiServer()
}

// UnimplementedOcpLessonApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpLessonApiServer struct {
}

func (UnimplementedOcpLessonApiServer) ListLessonsV1(context.Context, *ListLessonsV1Request) (*ListLessonsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLessonsV1 not implemented")
}
func (UnimplementedOcpLessonApiServer) DescribeLessonV1(context.Context, *DescribeLessonV1Request) (*DescribeLessonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeLessonV1 not implemented")
}
func (UnimplementedOcpLessonApiServer) CreateLessonV1(context.Context, *CreateLessonV1Request) (*CreateLessonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLessonV1 not implemented")
}
func (UnimplementedOcpLessonApiServer) RemoveLessonV1(context.Context, *RemoveLessonV1Request) (*RemoveLessonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLessonV1 not implemented")
}
func (UnimplementedOcpLessonApiServer) UpdateLessonV1(context.Context, *UpdateLessonV1Request) (*UpdateLessonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLessonV1 not implemented")
}
func (UnimplementedOcpLessonApiServer) MultiCreateLessonV1(context.Context, *MultiCreateLessonV1Request) (*MultiCreateLessonV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateLessonV1 not implemented")
}
func (UnimplementedOcpLessonApiServer) mustEmbedUnimplementedOcpLessonApiServer() {}

// UnsafeOcpLessonApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpLessonApiServer will
// result in compilation errors.
type UnsafeOcpLessonApiServer interface {
	mustEmbedUnimplementedOcpLessonApiServer()
}

func RegisterOcpLessonApiServer(s grpc.ServiceRegistrar, srv OcpLessonApiServer) {
	s.RegisterService(&OcpLessonApi_ServiceDesc, srv)
}

func _OcpLessonApi_ListLessonsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLessonsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpLessonApiServer).ListLessonsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.lesson.api.OcpLessonApi/ListLessonsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpLessonApiServer).ListLessonsV1(ctx, req.(*ListLessonsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpLessonApi_DescribeLessonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeLessonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpLessonApiServer).DescribeLessonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.lesson.api.OcpLessonApi/DescribeLessonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpLessonApiServer).DescribeLessonV1(ctx, req.(*DescribeLessonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpLessonApi_CreateLessonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLessonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpLessonApiServer).CreateLessonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.lesson.api.OcpLessonApi/CreateLessonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpLessonApiServer).CreateLessonV1(ctx, req.(*CreateLessonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpLessonApi_RemoveLessonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLessonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpLessonApiServer).RemoveLessonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.lesson.api.OcpLessonApi/RemoveLessonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpLessonApiServer).RemoveLessonV1(ctx, req.(*RemoveLessonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpLessonApi_UpdateLessonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLessonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpLessonApiServer).UpdateLessonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.lesson.api.OcpLessonApi/UpdateLessonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpLessonApiServer).UpdateLessonV1(ctx, req.(*UpdateLessonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpLessonApi_MultiCreateLessonV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateLessonV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpLessonApiServer).MultiCreateLessonV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.lesson.api.OcpLessonApi/MultiCreateLessonV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpLessonApiServer).MultiCreateLessonV1(ctx, req.(*MultiCreateLessonV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpLessonApi_ServiceDesc is the grpc.ServiceDesc for OcpLessonApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpLessonApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.lesson.api.OcpLessonApi",
	HandlerType: (*OcpLessonApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLessonsV1",
			Handler:    _OcpLessonApi_ListLessonsV1_Handler,
		},
		{
			MethodName: "DescribeLessonV1",
			Handler:    _OcpLessonApi_DescribeLessonV1_Handler,
		},
		{
			MethodName: "CreateLessonV1",
			Handler:    _OcpLessonApi_CreateLessonV1_Handler,
		},
		{
			MethodName: "RemoveLessonV1",
			Handler:    _OcpLessonApi_RemoveLessonV1_Handler,
		},
		{
			MethodName: "UpdateLessonV1",
			Handler:    _OcpLessonApi_UpdateLessonV1_Handler,
		},
		{
			MethodName: "MultiCreateLessonV1",
			Handler:    _OcpLessonApi_MultiCreateLessonV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lesson_service.proto",
}
