package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	cmd := "cmd"
	paramline := "/C, wsl, ls, -alrt"
	params := strings.Split(paramline, ",")
	params = func() []string {
		var tmpA []string
		for _, v := range params {
			tmpS := strings.TrimSpace(v)
			tmpA = append(tmpA, tmpS)
		}
		return tmpA
	}()
	// paramline2 := strings.Join([]string{"/C", "wsl", "ls", "-alrt"}, ",")

	// fmt.Printf("paramline = %v \n paramline2 = %v \n", paramline, paramline2)
	// fmt.Printf("Stirng Compare %v \n", strings.Compare(paramline, paramline2))

	var outBytes []byte
	outval := bytes.NewBuffer(outBytes)
	fmt.Printf("Params = %v \n Command = %v \n", cmd, params)
	cmdObject := exec.Command(cmd, params...)
	cmdObject.Args = params
	cmdObject.Stdout = outval
	if errObj := cmdObject.Run(); errObj != nil {
		fmt.Println(errObj)
	}

	fmt.Printf("Print outVal = %v \n", outval)

}
