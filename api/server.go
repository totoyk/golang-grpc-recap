package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "golang-grpc-recap/helloworld/github.com/totoyk/golang-grpc-recap/proto/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type HelloworldHandler struct {
	pb.UnimplementedGreeterServer
}

func (h HelloworldHandler) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}
func (h HelloworldHandler) SayRepeatHello(*pb.RepeatHelloRequest, grpc.ServerStreamingServer[pb.HelloReply]) error {
	return status.Errorf(codes.Unimplemented, "method SayRepeatHello not implemented")
}

func main() {
	port := "50051"
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &HelloworldHandler{})
	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	server.Serve(lis)
}
