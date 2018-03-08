package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	cmdlist := [][]string{{"pwd"}, {"ls", "-alrt"}, {"echo", "raj"}, {"echo", "loga"}, {"echo", "shradha"}, {"echo", "shruti"}, {"echo", "continube"}}
	//cmdlist := [][]string{{"pwd"}, {"pwd"}, {"pwd"}, {"pwd"}, {"pwd"}, {"pwd"}, {"pwd"}, {"pwd"}, {"pwd"}}
	//fmt.Printf(cmdlist)
	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		log.Fatal("Connecting:", err)
	}
	client := rpc.NewClient(conn)
	defer client.Close()

	for i := 0; i < 10; i++ {
		go func() {
			param1 := new(Obj1)
			//param1.Input.Params = make([]string, 5, 5)
			param1.ID = 94539
			param1.Name = "Kern Loop"
			param1.Input.Cmd = "wsl"
			//param1.Input.Params = []string{"pwd"}
			//param1.Input.Params = append(param1.Input.Params, "pwd")

			picker := rand.Intn(5)

			//fmt.Printf("Picker Value = %v \n", picker)
			//fmt.Printf("CMD List Value = %v \n", cmdlist[picker])
			param1.Input.Params = func() []string {
				var temp []string
				for _, val := range cmdlist[picker] {
					temp = append(temp, val)
				}
				return temp
			}()
			fmt.Printf("param1.Input.Params = %v \n", param1.Input.Params)

			param2 := new(Obj2)

			//fmt.Printf("Params = %v \n Command = %v \n", param1.Input.Params, param1.Input.Cmd)
			if err := client.Call("Instance.CalledFunction", param1, param2); err != nil {
				fmt.Printf("Error in Client = %v \n", err.Error())
			} else {
				fmt.Printf("Main Function: Output = %s \n", param2.Output.OutString)
			}

			wg.Done()
		}()
	}
	wg.Wait()

}
