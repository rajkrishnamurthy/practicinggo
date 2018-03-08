package main

import (
	"fmt"
)

// What is the objective:
// Can you encapsulate 2 interfaces: color and shape; under 1 common interface called supershape
// so that when the applications call they dynamically call the relevant receiver/method
// Findings: Yes

// Interface and Receiver Pointers
// http://kunalkushwaha.github.io/2015/09/11/understanding-golang-interfaces/

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (z square) color() string {
	return "Red Square"
}

func (z circle) color() string {
	return "Green Circle"
}

func (z circle) area() float64 {
	return 3.14 * z.radius * z.radius
}

type shape interface {
	area() float64
}

type color interface {
	color() string
}

type supershape interface {
	shape
	color
	perimeter() float64
}

func getarea(z supershape) {
	fmt.Printf("Inside getarea() Type : %T & Value : %v \n", z, z)
	fmt.Println(z.area())
}

func getcolor(z supershape) {
	fmt.Printf("Inside getcolor() Type : %T & Value : %v \n", z, z)
	fmt.Println(z.color())
}

func main() {
	s := square{10}
	fmt.Printf("Inside main(). Type : %T \n", s)
	getarea(s)
	getcolor(s)
	//fmt.Println(s.area())

}
