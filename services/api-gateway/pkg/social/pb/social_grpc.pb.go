// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pb/social.proto

package pb

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

// SocialServiceClient is the client API for SocialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialServiceClient interface {
	AddFollower(ctx context.Context, in *AddFollowerRequest, opts ...grpc.CallOption) (*AddFollowerResponse, error)
	RemoveFollower(ctx context.Context, in *RemoveFollowerRequest, opts ...grpc.CallOption) (*RemoveFollowerResponse, error)
	GetFollowers(ctx context.Context, in *GetFollowersRequest, opts ...grpc.CallOption) (*GetFollowersResponse, error)
	CountFollowers(ctx context.Context, in *CountFollowersRequest, opts ...grpc.CallOption) (*CountFollowersResponse, error)
}

type socialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialServiceClient(cc grpc.ClientConnInterface) SocialServiceClient {
	return &socialServiceClient{cc}
}

func (c *socialServiceClient) AddFollower(ctx context.Context, in *AddFollowerRequest, opts ...grpc.CallOption) (*AddFollowerResponse, error) {
	out := new(AddFollowerResponse)
	err := c.cc.Invoke(ctx, "/pb.SocialService/AddFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) RemoveFollower(ctx context.Context, in *RemoveFollowerRequest, opts ...grpc.CallOption) (*RemoveFollowerResponse, error) {
	out := new(RemoveFollowerResponse)
	err := c.cc.Invoke(ctx, "/pb.SocialService/RemoveFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) GetFollowers(ctx context.Context, in *GetFollowersRequest, opts ...grpc.CallOption) (*GetFollowersResponse, error) {
	out := new(GetFollowersResponse)
	err := c.cc.Invoke(ctx, "/pb.SocialService/GetFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) CountFollowers(ctx context.Context, in *CountFollowersRequest, opts ...grpc.CallOption) (*CountFollowersResponse, error) {
	out := new(CountFollowersResponse)
	err := c.cc.Invoke(ctx, "/pb.SocialService/CountFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialServiceServer is the server API for SocialService service.
// All implementations must embed UnimplementedSocialServiceServer
// for forward compatibility
type SocialServiceServer interface {
	AddFollower(context.Context, *AddFollowerRequest) (*AddFollowerResponse, error)
	RemoveFollower(context.Context, *RemoveFollowerRequest) (*RemoveFollowerResponse, error)
	GetFollowers(context.Context, *GetFollowersRequest) (*GetFollowersResponse, error)
	CountFollowers(context.Context, *CountFollowersRequest) (*CountFollowersResponse, error)
	mustEmbedUnimplementedSocialServiceServer()
}

// UnimplementedSocialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSocialServiceServer struct {
}

func (UnimplementedSocialServiceServer) AddFollower(context.Context, *AddFollowerRequest) (*AddFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFollower not implemented")
}
func (UnimplementedSocialServiceServer) RemoveFollower(context.Context, *RemoveFollowerRequest) (*RemoveFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFollower not implemented")
}
func (UnimplementedSocialServiceServer) GetFollowers(context.Context, *GetFollowersRequest) (*GetFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedSocialServiceServer) CountFollowers(context.Context, *CountFollowersRequest) (*CountFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFollowers not implemented")
}
func (UnimplementedSocialServiceServer) mustEmbedUnimplementedSocialServiceServer() {}

// UnsafeSocialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialServiceServer will
// result in compilation errors.
type UnsafeSocialServiceServer interface {
	mustEmbedUnimplementedSocialServiceServer()
}

func RegisterSocialServiceServer(s grpc.ServiceRegistrar, srv SocialServiceServer) {
	s.RegisterService(&SocialService_ServiceDesc, srv)
}

func _SocialService_AddFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).AddFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SocialService/AddFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).AddFollower(ctx, req.(*AddFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_RemoveFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).RemoveFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SocialService/RemoveFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).RemoveFollower(ctx, req.(*RemoveFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_GetFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).GetFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SocialService/GetFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).GetFollowers(ctx, req.(*GetFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_CountFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).CountFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SocialService/CountFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).CountFollowers(ctx, req.(*CountFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SocialService_ServiceDesc is the grpc.ServiceDesc for SocialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SocialService",
	HandlerType: (*SocialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFollower",
			Handler:    _SocialService_AddFollower_Handler,
		},
		{
			MethodName: "RemoveFollower",
			Handler:    _SocialService_RemoveFollower_Handler,
		},
		{
			MethodName: "GetFollowers",
			Handler:    _SocialService_GetFollowers_Handler,
		},
		{
			MethodName: "CountFollowers",
			Handler:    _SocialService_CountFollowers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/social.proto",
}