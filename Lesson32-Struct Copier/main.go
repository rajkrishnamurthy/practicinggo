package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type basicstruct struct {
	name        string
	description string
	tags        []string
	sub1struct
}

type sub1struct struct {
	subname string
	submap  map[string]string
	sub2struct
}

type sub2struct struct {
	subsqname string
	subsqtags []int
}

func main() {

	var sub2source = sub2struct{}
	sub2source.subsqname = "sub2 Name"
	sub2source.subsqtags = []int{1, 2, 3, 4, 5}
	fmt.Printf("Sub2Source \n ------ \n %#v \n", sub2source)

	var sub1source = sub1struct{}
	sub1source.subname = "sub1 Name"
	sub1source.submap = make(map[string]string)
	sub1source.submap["key1"] = "value1"
	sub1source.submap["key2"] = "value2"
	fmt.Printf("Sub1Source \n ------ \n %#v \n", sub1source)

	var target = basicstruct{}
	target.name = "target1"
	target.description = "some target"
	fmt.Printf("Target \n ------ \n %#v \n", target)

	var source = basicstruct{}
	sub1source.sub2struct = sub2source
	source.sub1struct = sub1source
	fmt.Printf("Source \n ------ \n %#v \n", source)

	copier.Copy(&target, &sub2source)
	fmt.Printf("New Target \n ------ \n %#v \n", target)

}
