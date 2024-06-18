// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: pubsub.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PubSubService_Subscribe_FullMethodName   = "/pubsub.PubSubService/Subscribe"
	PubSubService_Unsubscribe_FullMethodName = "/pubsub.PubSubService/Unsubscribe"
)

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSubServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
}

type pubSubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubServiceClient(cc grpc.ClientConnInterface) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSubService_SubscribeClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PubSubService_ServiceDesc.Streams[0], PubSubService_Subscribe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServiceSubscribeClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSubService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type pubSubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, PubSubService_Unsubscribe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSubServiceServer is the server API for PubSubService service.
// All implementations must embed UnimplementedPubSubServiceServer
// for forward compatibility
type PubSubServiceServer interface {
	Subscribe(*SubscribeRequest, PubSubService_SubscribeServer) error
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	mustEmbedUnimplementedPubSubServiceServer()
}

// UnimplementedPubSubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPubSubServiceServer struct {
}

func (UnimplementedPubSubServiceServer) Subscribe(*SubscribeRequest, PubSubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedPubSubServiceServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedPubSubServiceServer) mustEmbedUnimplementedPubSubServiceServer() {}

// UnsafePubSubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSubServiceServer will
// result in compilation errors.
type UnsafePubSubServiceServer interface {
	mustEmbedUnimplementedPubSubServiceServer()
}

func RegisterPubSubServiceServer(s grpc.ServiceRegistrar, srv PubSubServiceServer) {
	s.RegisterService(&PubSubService_ServiceDesc, srv)
}

func _PubSubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServiceServer).Subscribe(m, &pubSubServiceSubscribeServer{ServerStream: stream})
}

type PubSubService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type pubSubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PubSubService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSubService_Unsubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PubSubService_ServiceDesc is the grpc.ServiceDesc for PubSubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _PubSubService_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pubsub.proto",
}
