package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var outBytes, output []byte
	var outval *bytes.Buffer
	var cmdObject *exec.Cmd
	var err error

	// cmd := "goreturns"
	cmd := "goreturns.exe"
	//params := []string{"-w", "./"}
	params := []string{"-w", "./"}
	sandboxpath := "f:\\Coding\\DONOTUSE\\somecode\\sandboxes\\c994edb0-521d-11e8-8d2b-28d244281b6a"

	outval = bytes.NewBuffer(outBytes)
	cmdObject = exec.Command(cmd, params...)
	cmdObject.Dir = sandboxpath
	cmdObject.Stdout = outval

	fmt.Printf("cmdObject \n %v \n", cmdObject)

	if err = cmdObject.Run(); err != nil {
		//log.Fatalf(err.Error())
		log.Printf("Error in Running Commmand in WorkerNode = %v ", err)
		log.Printf("goreturns -w Output =  %v \n", outval.Bytes())
	}

	output, err = cmdObject.Output()
	fmt.Printf("%v \n", output)

	log.Printf("goreturns -w Output =  %v \n", outval.Bytes())
	outval.Reset()

	return
}
