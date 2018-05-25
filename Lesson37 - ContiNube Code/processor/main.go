package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	instance := new(TaskInstance)
	inputs := &Inputs{}
	inputs.Path = "F:/personal"
	inputs.Cmd = "cmd"
	// inputs.Params = []string{"/C", "wsl", "ls", "-alrt"}
	inputs.Params = []string{"/C", "wsl", "tree"}

	outputs := &Outputs{}

	if err := instance.ListProcesses(inputs, outputs); err != nil {
		fmt.Printf("%v", err)
		return
	}

	var byteString = "{"
	for _, b := range outputs.Output1 {
		byteString = byteString + strconv.Itoa(int(b)) + ","
	}
	byteString = byteString + "}"
	fmt.Printf("Printing Output \n %s \n", byteString)
}

// ListProcesses : Executes Windows Commands (wsl) with the parameters. For example, wsl ls -alrt
// This function is registered through a TCP handler
// RPC Client Call
func (inst *TaskInstance) ListProcesses(inputs *Inputs, outputs *Outputs) (err error) {

	cmd := inputs.Cmd
	params := inputs.Params
	var outBytes []byte
	outval := bytes.NewBuffer(outBytes)
	fmt.Printf("Params = %v \n Command = %v \n", cmd, params)

	if err := os.Chdir(inputs.Path); err != nil {
		fmt.Printf("%v", err)
	}

	cmdObject := exec.Command(cmd, params...)
	cmdObject.Stdout = outval
	if errObj := cmdObject.Run(); errObj != nil {
		fmt.Println(errObj)
		return errObj
	}
	outputs.Output1 = outval.Bytes()
	return nil
}

type TaskInstance int

type Inputs struct {
	Path   string   // Path where this needs to be executed
	Cmd    string   // Command Line
	Params []string // Parameters/Arguments to Command Line
}

type Outputs struct {
	Output1 []byte // Output Buffer
}
