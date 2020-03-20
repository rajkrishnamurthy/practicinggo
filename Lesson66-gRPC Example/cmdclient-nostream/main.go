package main

import (
	"context"
	"log"
	"time"

	pb "cmdexecutor"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCommandExecutorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 50000*time.Second)
	defer cancel()

	cmd := &pb.Command{}
	cmd.Cmd = "docker"
	cmd.Args = []string{"ps"}
	// cmd.IsFileToBeReturn = false
	cmdOutput, err := c.Execute(ctx, cmd)

	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	log.Printf("output: %s", cmdOutput.OutputType)

}
