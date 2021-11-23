package main

import (
	"fmt"

	"./dog"
)

func main() {
	d := dog.Dog{"Rex", "poodle"}
	fmt.Print(d)
}
