package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const servicename = "ListProcesses"

// const servicename = "{{function_name}}"

func main() {

	//Creating an instance of struct which implement dummy TaskInstance interface
	inst := new(TaskInstance)

	// Register a new rpc server (In most cases, you will use default server only)
	// And register struct we created above by service name same as function name = task name
	// The wrapper method here ensures that only structs which implement TaskInstance interface
	// are allowed to register themselves.
	server := rpc.NewServer()
	if err := server.RegisterName(servicename, inst); err != nil {
		log.Fatal(err)
		return
	}

	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", (":" + "46005"))
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// This statement links rpc server to the socket, and allows rpc server to accept
	// rpc request coming from that socket.
	server.Accept(l)

	// alternateMain()
}

//ExecuteTask : Exported
func (inst *TaskInstance) ExecuteTask(TaskInputvalues *InputValues, TaskOutputvalues *OutputValues) (err error) {

	tempOutputvalues := initOutputValues()

	// Outputvalues.Output.Names = []CNTaskIOName{"name1", "name2", "name3"}
	// outputvalues.Output.Types["name1"] = CNTaskIOType("some random value for name1")
	// outputvalues.Output.Output["name1"] = []byte("this is some test")
	// Outputvalues.Output.Code = []byte("this is some test")

	// Outputvalues.Names = []string{"name1", "name2", "name3"}
	// Outputvalues.Code = []byte("this is some test")

	tempOutputvalues.Name = "Will this atleast work?"
	TaskOutputvalues.Name = tempOutputvalues.Name

	fmt.Printf("InputValues = %v \t OutputValues = %v \n", TaskInputvalues, TaskOutputvalues)

	return nil

}

func initOutputValues() (outputvalues *OutputValues) {
	outputvalues = &OutputValues{}

	// outputvalues.Output.Names = make([]CNTaskIOName, 5)
	// outputvalues.Output.Descriptions = make(map[CNTaskIOName]string, 5)
	// outputvalues.Output.Types = make(map[CNTaskIOName]CNTaskIOType, 5)

	// outputvalues.Output.Values = make(map[CNTaskIOName]string, 5)
	// outputvalues.Output.Output = make(map[CNTaskIOName][]byte, 5)

	// outputvalues.Output.Code = make([]byte, 100)

	// outputvalues.Names = make([]string, 5)
	// outputvalues.Code = make([]byte, 100)

	return outputvalues
}
