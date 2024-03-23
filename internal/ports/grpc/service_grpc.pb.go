// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: internal/ports/grpc/service.proto

package grpc

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

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	GetPrompts(ctx context.Context, in *GetPromptsRequest, opts ...grpc.CallOption) (*PromptsResponse, error)
	AddPrompts(ctx context.Context, in *AddPromptsRequest, opts ...grpc.CallOption) (*PromptsResponse, error)
	UpdatePrompt(ctx context.Context, in *UpdatePromptRequest, opts ...grpc.CallOption) (*SinglePromptResponse, error)
	ReorderPrompts(ctx context.Context, in *UpdatePromptsPositionsRequest, opts ...grpc.CallOption) (*PromptsResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetPrompts(ctx context.Context, in *GetPromptsRequest, opts ...grpc.CallOption) (*PromptsResponse, error) {
	out := new(PromptsResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/GetPrompts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) AddPrompts(ctx context.Context, in *AddPromptsRequest, opts ...grpc.CallOption) (*PromptsResponse, error) {
	out := new(PromptsResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/AddPrompts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UpdatePrompt(ctx context.Context, in *UpdatePromptRequest, opts ...grpc.CallOption) (*SinglePromptResponse, error) {
	out := new(SinglePromptResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/UpdatePrompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) ReorderPrompts(ctx context.Context, in *UpdatePromptsPositionsRequest, opts ...grpc.CallOption) (*PromptsResponse, error) {
	out := new(PromptsResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/ReorderPrompts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*ProfileResponse, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*ProfileResponse, error)
	GetPrompts(context.Context, *GetPromptsRequest) (*PromptsResponse, error)
	AddPrompts(context.Context, *AddPromptsRequest) (*PromptsResponse, error)
	UpdatePrompt(context.Context, *UpdatePromptRequest) (*SinglePromptResponse, error)
	ReorderPrompts(context.Context, *UpdatePromptsPositionsRequest) (*PromptsResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) CreateProfile(context.Context, *CreateProfileRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfileServiceServer) GetProfile(context.Context, *GetProfileRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfileServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedProfileServiceServer) GetPrompts(context.Context, *GetPromptsRequest) (*PromptsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrompts not implemented")
}
func (UnimplementedProfileServiceServer) AddPrompts(context.Context, *AddPromptsRequest) (*PromptsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPrompts not implemented")
}
func (UnimplementedProfileServiceServer) UpdatePrompt(context.Context, *UpdatePromptRequest) (*SinglePromptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrompt not implemented")
}
func (UnimplementedProfileServiceServer) ReorderPrompts(context.Context, *UpdatePromptsPositionsRequest) (*PromptsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReorderPrompts not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetPrompts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromptsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetPrompts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/GetPrompts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetPrompts(ctx, req.(*GetPromptsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_AddPrompts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPromptsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).AddPrompts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/AddPrompts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).AddPrompts(ctx, req.(*AddPromptsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UpdatePrompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UpdatePrompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/UpdatePrompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UpdatePrompt(ctx, req.(*UpdatePromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_ReorderPrompts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePromptsPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).ReorderPrompts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/ReorderPrompts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).ReorderPrompts(ctx, req.(*UpdatePromptsPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profiles.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _ProfileService_CreateProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _ProfileService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _ProfileService_UpdateProfile_Handler,
		},
		{
			MethodName: "GetPrompts",
			Handler:    _ProfileService_GetPrompts_Handler,
		},
		{
			MethodName: "AddPrompts",
			Handler:    _ProfileService_AddPrompts_Handler,
		},
		{
			MethodName: "UpdatePrompt",
			Handler:    _ProfileService_UpdatePrompt_Handler,
		},
		{
			MethodName: "ReorderPrompts",
			Handler:    _ProfileService_ReorderPrompts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/ports/grpc/service.proto",
}
