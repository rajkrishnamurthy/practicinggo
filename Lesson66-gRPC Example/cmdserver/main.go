package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	pb "cmdexecutor"

	"google.golang.org/grpc"
)

const (
	// port = ":50051"
	port = ":80"
	// SuccessMessage :
	SuccessMessage = "Successfully executed"
)

type server struct{}

// Execute implements cmdexecutor.Execute
// func (s *server) Execute(ctx context.Context, in *pb.Command) (*pb.CommandExecutor_ExecuteClient, error) {
// 	cmdOutput := &pb.CommandOutput{}
// 	err := runCmd(in, cmdOutput)
// 	return cmdOutput, err
// }

// Execute implements cmdexecutor.Execute
func (s *server) Execute(stream pb.CommandExecutor_ExecuteServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("eeor 1 :", err)
			return nil
		}
		if err != nil {
			fmt.Println("eeor 2 :", err)
			return err
		}

		cmdOutput := &pb.CommandOutput{}
		err = runCmd(in, cmdOutput)

		if err := stream.Send(cmdOutput); err != nil {
			return err
		}

	}
}

func runCmd(in *pb.Command, cmdOutput *pb.CommandOutput) error {

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdObject := exec.Command(in.Cmd, in.Args...)

	cmdObject.Stdout = &out
	cmdObject.Stderr = &stderr
	if err := cmdObject.Run(); err != nil {
		stderrString := stderr.String()
		outString := out.String()
		if stderrString != "" || outString != "" {
			log.Println("Error in executing ssh cmd:", stderrString)
			log.Println("Output of ssh cmd:", outString)

			if strings.Contains(stderrString, "command not found") {

				// cmdSplit := strings.Split(stderrString, ":")
				// cmdOutput.Remediation = "install" + cmdSplit[1]
			}
			return err
		}
	}

	fmt.Println("out.String() :", out.String())

	cmdOutput.Message = out.String()

	if hostname, err := os.Hostname(); err == nil {
		cmdOutput.OutputType = hostname
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCommandExecutorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
