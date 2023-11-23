package main

import (
	"context"
	pb "de_pract/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedDeLogicServer
}

func (s *server) Calc(ctx context.Context, in *pb.Input) (*pb.Result, error) {
	return &pb.Result{Val: in.Val1 * in.Val2}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDeLogicServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
