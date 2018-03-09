package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {

	var servicename = "Instance"
	//Creating an instance of struct which implement Arith interface
	inst := new(Taskinstance)

	// Register a new rpc server (In most cases, you will use default server only)
	// And register struct we created above by name "Arith"
	// The wrapper method here ensures that only structs which implement Arith interface
	// are allowed to register themselves.
	server := rpc.NewServer()
	if err := server.RegisterName(servicename, inst); err != nil {
		log.Fatal(err)
		return
	}

	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", ":9876")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// This statement links rpc server to the socket, and allows rpc server to accept
	// rpc request coming from that socket.
	server.Accept(l)

}
