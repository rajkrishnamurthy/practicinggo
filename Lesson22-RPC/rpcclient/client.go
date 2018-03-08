package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}

	for {
		go func() {
			client := rpc.NewClient(conn)
			defer client.Close()

			param1 := new(Obj1)
			param1.Id = 94539
			param1.Name = "Kern Loop"

			param2 := new(Obj2)
			param2.Input.Cmd = "wsl"
			param2.Input.Params = []string{"pwd"}

			if err := client.Call("Instance.CalledFunction", param1, param2); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Main Function: Output = %s \n", param2.Output.OutString)
			wg.Done()
		}()
	}
	wg.Wait()

}
