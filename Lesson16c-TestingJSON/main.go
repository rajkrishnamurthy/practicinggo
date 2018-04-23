package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Testing JSON Nesting

type Person struct {
	Parent       *Person `json:"Parent,omitempty"`
	Name         string  `json:"Name,omitempty"`
	Relationship string  `json:"Relationship,omitempty"`
	Sibling      *Person `json:"Sibling,omitempty"`
}

func main() {
	var jsonified []byte
	var err error

	shradha := Person{
		Name:         "shradha",
		Relationship: "daughter",
	}

	shruti := Person{
		Name:         "shruti",
		Relationship: "daughter",
	}

	loga := Person{
		Name:         "loga",
		Relationship: "mom",
	}

	fmt.Printf("%v \n", shradha)
	if shradha.Parent == nil {
		fmt.Printf("No Parent Exists \n")
	}

	jsonified, err = json.Marshal(shradha)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSONified \n %s \n", jsonified)

	shradha.Parent = &loga
	shruti.Parent = &loga
	shruti.Sibling = &shradha
	//shradha.Sibling = &shruti

	jsonified, err = json.Marshal(shruti)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSONified AGAIN \n %s \n", jsonified)

}
