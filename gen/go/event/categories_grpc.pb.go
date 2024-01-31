// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: event/categories.proto

package event

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

// CategoriesServiceClient is the client API for CategoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoriesServiceClient interface {
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
}

type categoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoriesServiceClient(cc grpc.ClientConnInterface) CategoriesServiceClient {
	return &categoriesServiceClient{cc}
}

func (c *categoriesServiceClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/event.CategoriesService/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoriesServiceServer is the server API for CategoriesService service.
// All implementations must embed UnimplementedCategoriesServiceServer
// for forward compatibility
type CategoriesServiceServer interface {
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	mustEmbedUnimplementedCategoriesServiceServer()
}

// UnimplementedCategoriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoriesServiceServer struct {
}

func (UnimplementedCategoriesServiceServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedCategoriesServiceServer) mustEmbedUnimplementedCategoriesServiceServer() {}

// UnsafeCategoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoriesServiceServer will
// result in compilation errors.
type UnsafeCategoriesServiceServer interface {
	mustEmbedUnimplementedCategoriesServiceServer()
}

func RegisterCategoriesServiceServer(s grpc.ServiceRegistrar, srv CategoriesServiceServer) {
	s.RegisterService(&CategoriesService_ServiceDesc, srv)
}

func _CategoriesService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.CategoriesService/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoriesService_ServiceDesc is the grpc.ServiceDesc for CategoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.CategoriesService",
	HandlerType: (*CategoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategories",
			Handler:    _CategoriesService_GetCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event/categories.proto",
}
