package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	"google.golang.org/grpc"
	pb "grpc-hw/pkg/helloworld"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	tempFile := "srv.wav"
	err := exec.Command("flite", "-t", "Hello "+in.GetName(), "-o", tempFile).Run()
	if err != nil {
		return nil, fmt.Errorf("make audio failed: %v", err)
	}
	data, _ := ioutil.ReadFile(tempFile)
	return &pb.HelloReply{Message: data}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
