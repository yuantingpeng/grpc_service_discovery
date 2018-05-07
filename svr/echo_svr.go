package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../pb/echo"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.EchoReq) (*pb.EchoRsp, error) {
    return &pb.EchoRsp{Str: in.Str}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
             log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoSvrServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
