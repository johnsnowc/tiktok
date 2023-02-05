// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: relation.proto

package relation

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

// RelationSrvClient is the client API for RelationSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationSrvClient interface {
	RelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error)
	RelationFollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error)
	RelationFollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error)
}

type relationSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationSrvClient(cc grpc.ClientConnInterface) RelationSrvClient {
	return &relationSrvClient{cc}
}

func (c *relationSrvClient) RelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error) {
	out := new(DouyinRelationActionResponse)
	err := c.cc.Invoke(ctx, "/relation.RelationSrv/RelationAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationSrvClient) RelationFollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error) {
	out := new(DouyinRelationFollowListResponse)
	err := c.cc.Invoke(ctx, "/relation.RelationSrv/RelationFollowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationSrvClient) RelationFollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error) {
	out := new(DouyinRelationFollowerListResponse)
	err := c.cc.Invoke(ctx, "/relation.RelationSrv/RelationFollowerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationSrvServer is the server API for RelationSrv service.
// All implementations must embed UnimplementedRelationSrvServer
// for forward compatibility
type RelationSrvServer interface {
	RelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error)
	RelationFollowList(context.Context, *DouyinRelationFollowListRequest) (*DouyinRelationFollowListResponse, error)
	RelationFollowerList(context.Context, *DouyinRelationFollowerListRequest) (*DouyinRelationFollowerListResponse, error)
	mustEmbedUnimplementedRelationSrvServer()
}

// UnimplementedRelationSrvServer must be embedded to have forward compatible implementations.
type UnimplementedRelationSrvServer struct {
}

func (UnimplementedRelationSrvServer) RelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedRelationSrvServer) RelationFollowList(context.Context, *DouyinRelationFollowListRequest) (*DouyinRelationFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationFollowList not implemented")
}
func (UnimplementedRelationSrvServer) RelationFollowerList(context.Context, *DouyinRelationFollowerListRequest) (*DouyinRelationFollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationFollowerList not implemented")
}
func (UnimplementedRelationSrvServer) mustEmbedUnimplementedRelationSrvServer() {}

// UnsafeRelationSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationSrvServer will
// result in compilation errors.
type UnsafeRelationSrvServer interface {
	mustEmbedUnimplementedRelationSrvServer()
}

func RegisterRelationSrvServer(s grpc.ServiceRegistrar, srv RelationSrvServer) {
	s.RegisterService(&RelationSrv_ServiceDesc, srv)
}

func _RelationSrv_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationSrvServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationSrv/RelationAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationSrvServer).RelationAction(ctx, req.(*DouyinRelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationSrv_RelationFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationSrvServer).RelationFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationSrv/RelationFollowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationSrvServer).RelationFollowList(ctx, req.(*DouyinRelationFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationSrv_RelationFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationSrvServer).RelationFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationSrv/RelationFollowerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationSrvServer).RelationFollowerList(ctx, req.(*DouyinRelationFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationSrv_ServiceDesc is the grpc.ServiceDesc for RelationSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.RelationSrv",
	HandlerType: (*RelationSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RelationAction",
			Handler:    _RelationSrv_RelationAction_Handler,
		},
		{
			MethodName: "RelationFollowList",
			Handler:    _RelationSrv_RelationFollowList_Handler,
		},
		{
			MethodName: "RelationFollowerList",
			Handler:    _RelationSrv_RelationFollowerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}
