// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: pb/keyvaluestore.proto

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

const (
	Storage_SayHello_FullMethodName     = "/keyvaluestore.Storage/SayHello"
	Storage_SayGoodbye_FullMethodName   = "/keyvaluestore.Storage/SayGoodbye"
	Storage_Get_FullMethodName          = "/keyvaluestore.Storage/Get"
	Storage_Set_FullMethodName          = "/keyvaluestore.Storage/Set"
	Storage_Delete_FullMethodName       = "/keyvaluestore.Storage/Delete"
	Storage_GetAll_FullMethodName       = "/keyvaluestore.Storage/GetAll"
	Storage_Heartbeat_FullMethodName    = "/keyvaluestore.Storage/Heartbeat"
	Storage_Election_FullMethodName     = "/keyvaluestore.Storage/Election"
	Storage_GetPorts_FullMethodName     = "/keyvaluestore.Storage/GetPorts"
	Storage_RegisterPort_FullMethodName = "/keyvaluestore.Storage/RegisterPort"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayGoodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeResponse, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Set(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*Value, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeyValuePairList, error)
	Heartbeat(ctx context.Context, in *KeyValuePairList, opts ...grpc.CallOption) (*Empty, error)
	Election(ctx context.Context, in *RequestforVote, opts ...grpc.CallOption) (*Vote, error)
	GetPorts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PortList, error)
	RegisterPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*PortList, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, Storage_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SayGoodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeResponse, error) {
	out := new(GoodbyeResponse)
	err := c.cc.Invoke(ctx, Storage_SayGoodbye_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, Storage_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Set(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, Storage_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, Storage_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeyValuePairList, error) {
	out := new(KeyValuePairList)
	err := c.cc.Invoke(ctx, Storage_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Heartbeat(ctx context.Context, in *KeyValuePairList, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Storage_Heartbeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Election(ctx context.Context, in *RequestforVote, opts ...grpc.CallOption) (*Vote, error) {
	out := new(Vote)
	err := c.cc.Invoke(ctx, Storage_Election_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetPorts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PortList, error) {
	out := new(PortList)
	err := c.cc.Invoke(ctx, Storage_GetPorts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) RegisterPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*PortList, error) {
	out := new(PortList)
	err := c.cc.Invoke(ctx, Storage_RegisterPort_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayGoodbye(context.Context, *GoodbyeRequest) (*GoodbyeResponse, error)
	Get(context.Context, *Key) (*Value, error)
	Set(context.Context, *KeyValuePair) (*Value, error)
	Delete(context.Context, *Key) (*Value, error)
	GetAll(context.Context, *Empty) (*KeyValuePairList, error)
	Heartbeat(context.Context, *KeyValuePairList) (*Empty, error)
	Election(context.Context, *RequestforVote) (*Vote, error)
	GetPorts(context.Context, *Empty) (*PortList, error)
	RegisterPort(context.Context, *Port) (*PortList, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedStorageServer) SayGoodbye(context.Context, *GoodbyeRequest) (*GoodbyeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayGoodbye not implemented")
}
func (UnimplementedStorageServer) Get(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) Set(context.Context, *KeyValuePair) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedStorageServer) Delete(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStorageServer) GetAll(context.Context, *Empty) (*KeyValuePairList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedStorageServer) Heartbeat(context.Context, *KeyValuePairList) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedStorageServer) Election(context.Context, *RequestforVote) (*Vote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedStorageServer) GetPorts(context.Context, *Empty) (*PortList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPorts not implemented")
}
func (UnimplementedStorageServer) RegisterPort(context.Context, *Port) (*PortList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPort not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SayGoodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodbyeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SayGoodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SayGoodbye_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SayGoodbye(ctx, req.(*GoodbyeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Set(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePairList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Heartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Heartbeat(ctx, req.(*KeyValuePairList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestforVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Election_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Election(ctx, req.(*RequestforVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_GetPorts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetPorts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_RegisterPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).RegisterPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_RegisterPort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).RegisterPort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyvaluestore.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Storage_SayHello_Handler,
		},
		{
			MethodName: "SayGoodbye",
			Handler:    _Storage_SayGoodbye_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Storage_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Storage_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Storage_GetAll_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Storage_Heartbeat_Handler,
		},
		{
			MethodName: "Election",
			Handler:    _Storage_Election_Handler,
		},
		{
			MethodName: "GetPorts",
			Handler:    _Storage_GetPorts_Handler,
		},
		{
			MethodName: "RegisterPort",
			Handler:    _Storage_RegisterPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/keyvaluestore.proto",
}
