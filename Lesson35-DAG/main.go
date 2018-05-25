package main

import "fmt"

type funcList struct {
	callFunc []func()
}

func main() {

	funcer := func(inputString string) func() {
		return func() {
			fmt.Printf("Printing the input String %v \n", inputString)
		}
	}

	newlist := funcList{}
	newlist.callFunc = make([]func(), 10)
	newlist.callFunc[0] = funcer("testing")

	newlist.callFunc[0]()

}
