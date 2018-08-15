// https://play.golang.org/p/9M5Gr2HDRN
// https://stackoverflow.com/questions/32673407/dynamic-function-call-in-go

package main

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
)

type A struct {
	Name  string
	Value int
}

type B struct {
	Name1 string
	Name2 string
	Value float64
}

func doA() *A {
	return &A{"Cats", 10}
}

func doB() *B {
	return &B{"Cats", "Dogs", 10.0}
}

func Generic(w io.Writer, fn interface{}) {
	result := reflect.ValueOf(fn).Call([]reflect.Value{})[0].Interface()
	json.NewEncoder(w).Encode(result)
}

func main() {
	Generic(os.Stdout, doA)
	Generic(os.Stdout, doB)
}
