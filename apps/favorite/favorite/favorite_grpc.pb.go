// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: favorite.proto

package favorite

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

// FavoriteSrvClient is the client API for FavoriteSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteSrvClient interface {
	FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error)
	FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error)
}

type favoriteSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteSrvClient(cc grpc.ClientConnInterface) FavoriteSrvClient {
	return &favoriteSrvClient{cc}
}

func (c *favoriteSrvClient) FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error) {
	out := new(DouyinFavoriteActionResponse)
	err := c.cc.Invoke(ctx, "/favorite.FavoriteSrv/FavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteSrvClient) FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error) {
	out := new(DouyinFavoriteListResponse)
	err := c.cc.Invoke(ctx, "/favorite.FavoriteSrv/FavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteSrvServer is the server API for FavoriteSrv service.
// All implementations must embed UnimplementedFavoriteSrvServer
// for forward compatibility
type FavoriteSrvServer interface {
	FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error)
	FavoriteList(context.Context, *DouyinFavoriteListRequest) (*DouyinFavoriteListResponse, error)
	mustEmbedUnimplementedFavoriteSrvServer()
}

// UnimplementedFavoriteSrvServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteSrvServer struct {
}

func (UnimplementedFavoriteSrvServer) FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteSrvServer) FavoriteList(context.Context, *DouyinFavoriteListRequest) (*DouyinFavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedFavoriteSrvServer) mustEmbedUnimplementedFavoriteSrvServer() {}

// UnsafeFavoriteSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteSrvServer will
// result in compilation errors.
type UnsafeFavoriteSrvServer interface {
	mustEmbedUnimplementedFavoriteSrvServer()
}

func RegisterFavoriteSrvServer(s grpc.ServiceRegistrar, srv FavoriteSrvServer) {
	s.RegisterService(&FavoriteSrv_ServiceDesc, srv)
}

func _FavoriteSrv_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteSrvServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.FavoriteSrv/FavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteSrvServer).FavoriteAction(ctx, req.(*DouyinFavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteSrv_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteSrvServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.FavoriteSrv/FavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteSrvServer).FavoriteList(ctx, req.(*DouyinFavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteSrv_ServiceDesc is the grpc.ServiceDesc for FavoriteSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "favorite.FavoriteSrv",
	HandlerType: (*FavoriteSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _FavoriteSrv_FavoriteAction_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _FavoriteSrv_FavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
