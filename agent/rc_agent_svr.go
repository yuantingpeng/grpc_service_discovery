package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../pb/agent"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetRouteInfo(ctx context.Context, in *pb.RouteInfoReq) (*pb.RouteInfoRsp, error) {
	return &pb.RouteInfoRsp{SIp: "", IPort: 0}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
             log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAgentSvrServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
