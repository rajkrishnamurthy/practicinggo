package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CalledFunction : Exported Function
func (inst *Taskinstance) CalledFunction(Param1 *Obj1, Param2 *Obj2) (Err error) {

	params := Param2.Input.Params
	cmd := Param2.Input.Cmd

	var outBytes []byte

	outval := bytes.NewBuffer(outBytes)
	//params = append(params, "pwd")

	cmdObject := exec.Command(cmd, params...)
	cmdObject.Stdout = outval
	if errObj := cmdObject.Run(); errObj != nil {
		fmt.Println(errObj)
		return errObj
	}
	Param2.Output.OutBytes = outval.Bytes()
	Param2.Output.OutString = outval.String()
	return nil
}
