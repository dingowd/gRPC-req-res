package main

import (
	"context"
	"fmt"
	rqrs "github.com/dingowd/gRPC-req-res/req-res"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
)

type server struct {
	rqrs.UnimplementedBackServer
}

func (s *server) Do(c context.Context, rq *rqrs.Request) (*rqrs.Response, error) {
	response := &rqrs.Response{
		Message: rq.Message + " Hello from server",
	}
	fmt.Fprintln(os.Stdout, "Received: ", rq.String())
	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	rqrs.RegisterBackServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		grpclog.Fatal(err)
	}
}
