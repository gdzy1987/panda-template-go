package server

import (
	"github.com/XM-GO/panda-kit/transport/grpc"
)

// NewGRPCServer new a GRPC server.
func NewGRPCServer(addr string) *grpc.Server {
	s := grpc.NewServer(addr)
	return s
}
