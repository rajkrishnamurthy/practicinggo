package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type somestruct struct {
	Name        string
	Description string
	Tags        []string
	Address     Address
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	Zipcode      string
}

// type someotherstruct somestruct
type someotherstruct struct {
	ID      int
	Tags    []string
	Name    string
	Address *Address
}

func main() {

	var source = &somestruct{
		Name:        "testing struct 1",
		Description: "lorem ipsum dupsum kupsum ....",
		Tags:        []string{"one", "two", "three", "four"},
		Address: Address{
			AddressLine1: "1808 Kern Loop",
			Zipcode:      "94539",
		},
	}
	var target = &someotherstruct{}

	fmt.Printf("Source \n ------ \n %#v \n", source)

	copier.Copy(target, source)
	fmt.Printf("New Target \n ------ \n %#v \n", target)

}
