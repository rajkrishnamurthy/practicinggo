package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := `processSpecification=Sp\u00e9cificati\u003con du \u003eprocessus`
	fmt.Println(input)
	// fmt.Println(strconv.Unquote("\"" + input + "\""))
	fmt.Println(strconv.Unquote(`"` + input + `"`))
}
