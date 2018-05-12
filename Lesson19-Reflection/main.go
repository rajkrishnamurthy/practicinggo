package main

import (
	"fmt"
	"reflect"
	"strings"
)

// // Foo is a sample struct
// type Foo struct {
// 	A  int `tag1:"First Tag" tag2:"Second Tag"`
// 	B  string
// 	fn func()
// }

// func test() {
// 	fmt.Printf("Nothing is really happening in this function \n")
// }
type Inputs struct {
	Cmd    string   // Command Line
	Params []string // Parameters/Arguments to Command Line
	test   struct {
		test1 string
		test2 []byte
	}
}

func main() {

	// type s []int
	// sl := s{1, 2, 3}
	// fmt.Printf("slType := reflect.TypeOf(sl); where sl := []int{1, 2, 3}  \n")

	sl := Inputs{}
	slType := reflect.TypeOf(sl)
	examiner(slType, 0)

}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			examiner(f.Type, depth+2)
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}

func valuer(v reflect.Value, depth int) {

}

func createSlice(t reflect.Type) reflect.Value {
	var sliceType reflect.Type
	sliceType = reflect.SliceOf(t)
	return reflect.Zero(sliceType)
}
