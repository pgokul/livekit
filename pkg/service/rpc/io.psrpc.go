// Code generated by protoc-gen-psrpc v0.2.3, DO NOT EDIT.
// source: pkg/service/rpc/io.proto

package rpc

import context "context"

import psrpc1 "github.com/livekit/psrpc"
import google_protobuf2 "google.golang.org/protobuf/types/known/emptypb"
import livekit "github.com/livekit/protocol/livekit"
import livekit3 "github.com/livekit/protocol/livekit"

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	UpdateEgressInfo(context.Context, *livekit.EgressInfo, ...psrpc1.RequestOption) (*google_protobuf2.Empty, error)

	GetIngressInfo(context.Context, *livekit3.GetIngressInfoRequest, ...psrpc1.RequestOption) (*livekit3.GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *livekit3.UpdateIngressStateRequest, ...psrpc1.RequestOption) (*google_protobuf2.Empty, error)
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	UpdateEgressInfo(context.Context, *livekit.EgressInfo) (*google_protobuf2.Empty, error)

	GetIngressInfo(context.Context, *livekit3.GetIngressInfoRequest) (*livekit3.GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *livekit3.UpdateIngressStateRequest) (*google_protobuf2.Empty, error)
}

// =======================
// IOInfo Server Interface
// =======================

type IOInfoServer interface {
	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// IOInfo Client
// =============

type iOInfoClient struct {
	client *psrpc1.RPCClient
}

// NewIOInfoClient creates a psrpc client that implements the IOInfoClient interface.
func NewIOInfoClient(clientID string, bus psrpc1.MessageBus, opts ...psrpc1.ClientOption) (IOInfoClient, error) {
	rpcClient, err := psrpc1.NewRPCClient("IOInfo", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) UpdateEgressInfo(ctx context.Context, req *livekit.EgressInfo, opts ...psrpc1.RequestOption) (*google_protobuf2.Empty, error) {
	return psrpc1.RequestSingle[*google_protobuf2.Empty](ctx, c.client, "UpdateEgressInfo", "", req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *livekit3.GetIngressInfoRequest, opts ...psrpc1.RequestOption) (*livekit3.GetIngressInfoResponse, error) {
	return psrpc1.RequestSingle[*livekit3.GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", "", req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *livekit3.UpdateIngressStateRequest, opts ...psrpc1.RequestOption) (*google_protobuf2.Empty, error) {
	return psrpc1.RequestSingle[*google_protobuf2.Empty](ctx, c.client, "UpdateIngressState", "", req, opts...)
}

// =============
// IOInfo Server
// =============

type iOInfoServer struct {
	svc IOInfoServerImpl
	rpc *psrpc1.RPCServer
}

// NewIOInfoServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIOInfoServer(serverID string, svc IOInfoServerImpl, bus psrpc1.MessageBus, opts ...psrpc1.ServerOption) (IOInfoServer, error) {
	s := psrpc1.NewRPCServer("IOInfo", serverID, bus, opts...)

	var err error
	err = psrpc1.RegisterHandler(s, "UpdateEgressInfo", "", svc.UpdateEgressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	err = psrpc1.RegisterHandler(s, "GetIngressInfo", "", svc.GetIngressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	err = psrpc1.RegisterHandler(s, "UpdateIngressState", "", svc.UpdateIngressState, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &iOInfoServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *iOInfoServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *iOInfoServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor2 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x11, 0x61, 0x16, 0x59, 0x88, 0x44, 0x11, 0x89, 0xa0, 0xe0, 0x52, 0x24, 0x01, 0x7d,
	0x00, 0x41, 0x18, 0xa4, 0x2b, 0xff, 0x70, 0xe3, 0x66, 0x68, 0xe3, 0x9d, 0x18, 0xa6, 0x93, 0x7b,
	0x4d, 0x6e, 0x07, 0x7c, 0x69, 0x9f, 0x41, 0x6c, 0xd2, 0x0e, 0x2a, 0x5d, 0x85, 0x9c, 0x8f, 0xf3,
	0x71, 0xb8, 0xe2, 0x98, 0x56, 0xce, 0x24, 0x88, 0x1b, 0x6f, 0xc1, 0x44, 0xb2, 0xc6, 0xa3, 0xa6,
	0x88, 0x8c, 0x72, 0x37, 0x92, 0x55, 0x87, 0xad, 0xdf, 0xc0, 0xca, 0xf3, 0x02, 0x5c, 0x84, 0x94,
	0x32, 0x52, 0x6a, 0x48, 0x23, 0xd9, 0x85, 0x0f, 0x0c, 0x31, 0xd4, 0x6d, 0x61, 0x27, 0x0e, 0xd1,
	0xb5, 0x60, 0xfa, 0x5f, 0xd3, 0x2d, 0x0d, 0xac, 0x89, 0x3f, 0x33, 0xbc, 0xfa, 0xda, 0x11, 0xb3,
	0xea, 0xbe, 0x0a, 0x4b, 0x94, 0x37, 0x62, 0xff, 0x85, 0xde, 0x6a, 0x86, 0x79, 0x6f, 0xee, 0xb3,
	0x03, 0x5d, 0xc4, 0x7a, 0x1b, 0xaa, 0x23, 0x9d, 0x8d, 0x7a, 0x30, 0xea, 0xf9, 0x8f, 0x51, 0x3e,
	0x8a, 0xbd, 0x3b, 0xe0, 0x2a, 0x6c, 0xeb, 0xa7, 0x63, 0xfd, 0x37, 0x78, 0x82, 0x8f, 0x0e, 0x12,
	0xab, 0xb3, 0x49, 0x9e, 0x08, 0x43, 0x02, 0xf9, 0x20, 0x64, 0xde, 0x54, 0xe0, 0x33, 0xd7, 0x0c,
	0xf2, 0x7c, 0xac, 0xfd, 0x87, 0x83, 0x7a, 0x62, 0xe4, 0xed, 0xe5, 0xeb, 0x85, 0xf3, 0xfc, 0xde,
	0x35, 0xda, 0xe2, 0xda, 0x14, 0xcf, 0xf8, 0xfe, 0xb9, 0x7d, 0x33, 0xeb, 0xdb, 0xd7, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x87, 0x58, 0x4d, 0x3b, 0x95, 0x01, 0x00, 0x00,
}
