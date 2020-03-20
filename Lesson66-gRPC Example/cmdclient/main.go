package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "cmdexecutor"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	// address     = "cmdserver"
	defaultName = "world"
)

func main() {

	for {
		go func() {
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
			stream, err := c.Execute(ctx)

			for j := 0; j < 10; j++ {
				stream.Send(cmd)
				cmdOutput, err := stream.Recv()
				// fmt.Println("eof :", err)
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Fatalf("cannot connect: %v", err)
				}
				log.Printf("output: %s", cmdOutput.OutputType)
			}
		}()
		// stream.CloseSend()
	}

}
