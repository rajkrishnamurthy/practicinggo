package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Foo is a sample struct
type Foo struct {
	A  int `tag1:"First Tag" tag2:"Second Tag"`
	B  string
	fn func()
}

func test() {
	fmt.Printf("Nothing is really happening in this function \n")
}

func main() {
	var fnName = "test"
	// var fnObject func()
	sl := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting

	fmt.Printf("Reflect ValueOf: %v \t TypeOf: %v \n", reflect.ValueOf(fnName), reflect.TypeOf(fnName))
	f := Foo{A: 10, B: "Salutations", fn: test}
	f.fn()
	fp := &f

	slType := reflect.TypeOf(sl)
	gType := reflect.TypeOf(greeting)
	grpType := reflect.TypeOf(greetingPtr)
	fType := reflect.TypeOf(f)
	fpType := reflect.TypeOf(fp)

	examiner(slType, 0)
	examiner(gType, 0)
	examiner(grpType, 0)
	examiner(fType, 0)
	examiner(fpType, 0)
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
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}
