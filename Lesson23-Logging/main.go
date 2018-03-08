package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	logpath = flag.String("logpath", "logfile.log", "Log Path")
)

func main() {
	flag.Parse()
	NewLog(*logpath)
	Log.Println("hello")
	for i := 0; i < 10; i++ {
		Log.Println(i)
		//time.Sleep(3 * time.Second)
	}

	errObj := os.MkdirAll("F:\\Coding\\GoProg\\temp\\newdir\\", 0777)
	if errObj != nil {
		Log.Println(errObj)
	}

	var params, cmd = []string{}, string("wsl")
	var outBytes []byte
	outval := bytes.NewBuffer(outBytes)
	params = append(params, "pwd")

	cmdObject := exec.Command(cmd, params...)
	cmdObject.Stdout = outval
	if errObj := cmdObject.Run(); errObj != nil {
		Log.Println(errObj)
	}
	fmt.Printf("Output %s \n", outval.String())
}
