package main

// Passing by Reference: 15_passing_by_value: 06_REFERENCE-TYPE

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	fmt.Println(&m[0])
	changeMe(m)
	fmt.Println(m) // [Todd]
}

func changeMe(z []string) {
	fmt.Printf("&z[0]: %0x \n", &z[0])
	z[0] = "Todd"
	fmt.Printf("z: %s \n", z) // [Todd]
	fmt.Println("*&z[0] :" + *&z[0])
}
