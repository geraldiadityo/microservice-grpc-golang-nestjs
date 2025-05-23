// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: pengguna/pengguna.proto

package pengguna

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PenggunaService_CreatePengguna_FullMethodName = "/pengguna.PenggunaService/CreatePengguna"
	PenggunaService_GetPenggunas_FullMethodName   = "/pengguna.PenggunaService/GetPenggunas"
	PenggunaService_GetPengguna_FullMethodName    = "/pengguna.PenggunaService/GetPengguna"
	PenggunaService_UpdatePengguna_FullMethodName = "/pengguna.PenggunaService/UpdatePengguna"
	PenggunaService_DeletePengguna_FullMethodName = "/pengguna.PenggunaService/DeletePengguna"
)

// PenggunaServiceClient is the client API for PenggunaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PenggunaServiceClient interface {
	CreatePengguna(ctx context.Context, in *CreatePenggunaRequest, opts ...grpc.CallOption) (*PenggunaResponse, error)
	GetPenggunas(ctx context.Context, in *GetPenggunasRequest, opts ...grpc.CallOption) (*GetPenggunasResponse, error)
	GetPengguna(ctx context.Context, in *GetPenggunaRequest, opts ...grpc.CallOption) (*PenggunaResponse, error)
	UpdatePengguna(ctx context.Context, in *UpdatePenggunaRequest, opts ...grpc.CallOption) (*PenggunaResponse, error)
	DeletePengguna(ctx context.Context, in *DeletePenggunaRequest, opts ...grpc.CallOption) (*DeletePenggunaResponse, error)
}

type penggunaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPenggunaServiceClient(cc grpc.ClientConnInterface) PenggunaServiceClient {
	return &penggunaServiceClient{cc}
}

func (c *penggunaServiceClient) CreatePengguna(ctx context.Context, in *CreatePenggunaRequest, opts ...grpc.CallOption) (*PenggunaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PenggunaResponse)
	err := c.cc.Invoke(ctx, PenggunaService_CreatePengguna_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *penggunaServiceClient) GetPenggunas(ctx context.Context, in *GetPenggunasRequest, opts ...grpc.CallOption) (*GetPenggunasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPenggunasResponse)
	err := c.cc.Invoke(ctx, PenggunaService_GetPenggunas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *penggunaServiceClient) GetPengguna(ctx context.Context, in *GetPenggunaRequest, opts ...grpc.CallOption) (*PenggunaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PenggunaResponse)
	err := c.cc.Invoke(ctx, PenggunaService_GetPengguna_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *penggunaServiceClient) UpdatePengguna(ctx context.Context, in *UpdatePenggunaRequest, opts ...grpc.CallOption) (*PenggunaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PenggunaResponse)
	err := c.cc.Invoke(ctx, PenggunaService_UpdatePengguna_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *penggunaServiceClient) DeletePengguna(ctx context.Context, in *DeletePenggunaRequest, opts ...grpc.CallOption) (*DeletePenggunaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePenggunaResponse)
	err := c.cc.Invoke(ctx, PenggunaService_DeletePengguna_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PenggunaServiceServer is the server API for PenggunaService service.
// All implementations must embed UnimplementedPenggunaServiceServer
// for forward compatibility.
type PenggunaServiceServer interface {
	CreatePengguna(context.Context, *CreatePenggunaRequest) (*PenggunaResponse, error)
	GetPenggunas(context.Context, *GetPenggunasRequest) (*GetPenggunasResponse, error)
	GetPengguna(context.Context, *GetPenggunaRequest) (*PenggunaResponse, error)
	UpdatePengguna(context.Context, *UpdatePenggunaRequest) (*PenggunaResponse, error)
	DeletePengguna(context.Context, *DeletePenggunaRequest) (*DeletePenggunaResponse, error)
	mustEmbedUnimplementedPenggunaServiceServer()
}

// UnimplementedPenggunaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPenggunaServiceServer struct{}

func (UnimplementedPenggunaServiceServer) CreatePengguna(context.Context, *CreatePenggunaRequest) (*PenggunaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePengguna not implemented")
}
func (UnimplementedPenggunaServiceServer) GetPenggunas(context.Context, *GetPenggunasRequest) (*GetPenggunasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPenggunas not implemented")
}
func (UnimplementedPenggunaServiceServer) GetPengguna(context.Context, *GetPenggunaRequest) (*PenggunaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPengguna not implemented")
}
func (UnimplementedPenggunaServiceServer) UpdatePengguna(context.Context, *UpdatePenggunaRequest) (*PenggunaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePengguna not implemented")
}
func (UnimplementedPenggunaServiceServer) DeletePengguna(context.Context, *DeletePenggunaRequest) (*DeletePenggunaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePengguna not implemented")
}
func (UnimplementedPenggunaServiceServer) mustEmbedUnimplementedPenggunaServiceServer() {}
func (UnimplementedPenggunaServiceServer) testEmbeddedByValue()                         {}

// UnsafePenggunaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PenggunaServiceServer will
// result in compilation errors.
type UnsafePenggunaServiceServer interface {
	mustEmbedUnimplementedPenggunaServiceServer()
}

func RegisterPenggunaServiceServer(s grpc.ServiceRegistrar, srv PenggunaServiceServer) {
	// If the following call pancis, it indicates UnimplementedPenggunaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PenggunaService_ServiceDesc, srv)
}

func _PenggunaService_CreatePengguna_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePenggunaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).CreatePengguna(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PenggunaService_CreatePengguna_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).CreatePengguna(ctx, req.(*CreatePenggunaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PenggunaService_GetPenggunas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPenggunasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).GetPenggunas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PenggunaService_GetPenggunas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).GetPenggunas(ctx, req.(*GetPenggunasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PenggunaService_GetPengguna_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPenggunaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).GetPengguna(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PenggunaService_GetPengguna_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).GetPengguna(ctx, req.(*GetPenggunaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PenggunaService_UpdatePengguna_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePenggunaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).UpdatePengguna(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PenggunaService_UpdatePengguna_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).UpdatePengguna(ctx, req.(*UpdatePenggunaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PenggunaService_DeletePengguna_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePenggunaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).DeletePengguna(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PenggunaService_DeletePengguna_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).DeletePengguna(ctx, req.(*DeletePenggunaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PenggunaService_ServiceDesc is the grpc.ServiceDesc for PenggunaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PenggunaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pengguna.PenggunaService",
	HandlerType: (*PenggunaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePengguna",
			Handler:    _PenggunaService_CreatePengguna_Handler,
		},
		{
			MethodName: "GetPenggunas",
			Handler:    _PenggunaService_GetPenggunas_Handler,
		},
		{
			MethodName: "GetPengguna",
			Handler:    _PenggunaService_GetPengguna_Handler,
		},
		{
			MethodName: "UpdatePengguna",
			Handler:    _PenggunaService_UpdatePengguna_Handler,
		},
		{
			MethodName: "DeletePengguna",
			Handler:    _PenggunaService_DeletePengguna_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pengguna/pengguna.proto",
}
