package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	"google.golang.org/grpc"
	pb "grpc-hw/pkg/helloworld"
)

var (
	address     = "localhost:50051"
	name        = ""
	defaultName = "world"
)

func main() {
	flag.StringVar(&address, "addr", address, "server address")
	flag.StringVar(&name, "name", "world", "name")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	tempFile := "cli.wav"
	ioutil.WriteFile(tempFile, r.Message, 0666)
	exec.Command("afplay", tempFile).Run()
}
