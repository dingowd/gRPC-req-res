package main

import (
	"context"
	"fmt"
	rqrs "github.com/dingowd/gRPC-req-res/req-res"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var s string
	client := rqrs.NewBackClient(conn)
	for {
		fmt.Fscan(os.Stdin, &s)
		request := &rqrs.Request{
			Message: s,
		}
		response, err := client.Do(context.Background(), request)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		fmt.Fprintln(os.Stdout, response.Message)
	}
}
