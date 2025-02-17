// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.6
// source: agent-vault.proto

package vaultservpb

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

// VaultClient is the client API for Vault service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VaultClient interface {
	StoreCred(ctx context.Context, in *StoreCredRequest, opts ...grpc.CallOption) (*StoreCredResponse, error)
}

type vaultClient struct {
	cc grpc.ClientConnInterface
}

func NewVaultClient(cc grpc.ClientConnInterface) VaultClient {
	return &vaultClient{cc}
}

func (c *vaultClient) StoreCred(ctx context.Context, in *StoreCredRequest, opts ...grpc.CallOption) (*StoreCredResponse, error) {
	out := new(StoreCredResponse)
	err := c.cc.Invoke(ctx, "/vaultservpb.Vault/StoreCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VaultServer is the server API for Vault service.
// All implementations must embed UnimplementedVaultServer
// for forward compatibility
type VaultServer interface {
	StoreCred(context.Context, *StoreCredRequest) (*StoreCredResponse, error)
	mustEmbedUnimplementedVaultServer()
}

// UnimplementedVaultServer must be embedded to have forward compatible implementations.
type UnimplementedVaultServer struct {
}

func (UnimplementedVaultServer) StoreCred(context.Context, *StoreCredRequest) (*StoreCredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCred not implemented")
}
func (UnimplementedVaultServer) mustEmbedUnimplementedVaultServer() {}

// UnsafeVaultServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VaultServer will
// result in compilation errors.
type UnsafeVaultServer interface {
	mustEmbedUnimplementedVaultServer()
}

func RegisterVaultServer(s grpc.ServiceRegistrar, srv VaultServer) {
	s.RegisterService(&Vault_ServiceDesc, srv)
}

func _Vault_StoreCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).StoreCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vaultservpb.Vault/StoreCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).StoreCred(ctx, req.(*StoreCredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vault_ServiceDesc is the grpc.ServiceDesc for Vault service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vault_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vaultservpb.Vault",
	HandlerType: (*VaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreCred",
			Handler:    _Vault_StoreCred_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent-vault.proto",
}
