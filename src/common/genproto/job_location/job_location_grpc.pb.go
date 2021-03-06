// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package job_location

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

// JobLocationServiceClient is the client API for JobLocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobLocationServiceClient interface {
	GetJobCity(ctx context.Context, in *GetJobCityRequest, opts ...grpc.CallOption) (*GetJobCityResponse, error)
	GetJobState(ctx context.Context, in *GetJobStateRequest, opts ...grpc.CallOption) (*GetJobStateResponse, error)
	GetJobLocations(ctx context.Context, in *GetJobLocationsRequest, opts ...grpc.CallOption) (*GetJobLocationsResponse, error)
}

type jobLocationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobLocationServiceClient(cc grpc.ClientConnInterface) JobLocationServiceClient {
	return &jobLocationServiceClient{cc}
}

func (c *jobLocationServiceClient) GetJobCity(ctx context.Context, in *GetJobCityRequest, opts ...grpc.CallOption) (*GetJobCityResponse, error) {
	out := new(GetJobCityResponse)
	err := c.cc.Invoke(ctx, "/job_location.JobLocationService/GetJobCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobLocationServiceClient) GetJobState(ctx context.Context, in *GetJobStateRequest, opts ...grpc.CallOption) (*GetJobStateResponse, error) {
	out := new(GetJobStateResponse)
	err := c.cc.Invoke(ctx, "/job_location.JobLocationService/GetJobState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobLocationServiceClient) GetJobLocations(ctx context.Context, in *GetJobLocationsRequest, opts ...grpc.CallOption) (*GetJobLocationsResponse, error) {
	out := new(GetJobLocationsResponse)
	err := c.cc.Invoke(ctx, "/job_location.JobLocationService/GetJobLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobLocationServiceServer is the server API for JobLocationService service.
// All implementations should embed UnimplementedJobLocationServiceServer
// for forward compatibility
type JobLocationServiceServer interface {
	GetJobCity(context.Context, *GetJobCityRequest) (*GetJobCityResponse, error)
	GetJobState(context.Context, *GetJobStateRequest) (*GetJobStateResponse, error)
	GetJobLocations(context.Context, *GetJobLocationsRequest) (*GetJobLocationsResponse, error)
}

// UnimplementedJobLocationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedJobLocationServiceServer struct {
}

func (UnimplementedJobLocationServiceServer) GetJobCity(context.Context, *GetJobCityRequest) (*GetJobCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobCity not implemented")
}
func (UnimplementedJobLocationServiceServer) GetJobState(context.Context, *GetJobStateRequest) (*GetJobStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobState not implemented")
}
func (UnimplementedJobLocationServiceServer) GetJobLocations(context.Context, *GetJobLocationsRequest) (*GetJobLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobLocations not implemented")
}

// UnsafeJobLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobLocationServiceServer will
// result in compilation errors.
type UnsafeJobLocationServiceServer interface {
	mustEmbedUnimplementedJobLocationServiceServer()
}

func RegisterJobLocationServiceServer(s grpc.ServiceRegistrar, srv JobLocationServiceServer) {
	s.RegisterService(&JobLocationService_ServiceDesc, srv)
}

func _JobLocationService_GetJobCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobLocationServiceServer).GetJobCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_location.JobLocationService/GetJobCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobLocationServiceServer).GetJobCity(ctx, req.(*GetJobCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobLocationService_GetJobState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobLocationServiceServer).GetJobState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_location.JobLocationService/GetJobState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobLocationServiceServer).GetJobState(ctx, req.(*GetJobStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobLocationService_GetJobLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobLocationServiceServer).GetJobLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_location.JobLocationService/GetJobLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobLocationServiceServer).GetJobLocations(ctx, req.(*GetJobLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobLocationService_ServiceDesc is the grpc.ServiceDesc for JobLocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobLocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job_location.JobLocationService",
	HandlerType: (*JobLocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobCity",
			Handler:    _JobLocationService_GetJobCity_Handler,
		},
		{
			MethodName: "GetJobState",
			Handler:    _JobLocationService_GetJobState_Handler,
		},
		{
			MethodName: "GetJobLocations",
			Handler:    _JobLocationService_GetJobLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_location.proto",
}
