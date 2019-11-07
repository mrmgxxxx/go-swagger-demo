package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "go-swagger-demo/examples/protocol/echoserver"
)

const (
	// endpoint to listen.
	endpoint = ":9527"
)

// EchoServer is example echo server.
type EchoServer struct {
}

// NewEchoServer creates a new echo server instance.
func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

// Echo impls Echo gRPC interface.
func (s *EchoServer) Echo(ctx context.Context, req *pb.EchoReq) (*pb.EchoResp, error) {
	return &pb.EchoResp{Msg: req.Msg}, nil
}

func main() {
	// new echo server instance.
	server := NewEchoServer()

	// listen.
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	// gRPC server register.
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, server)

	// serve.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
