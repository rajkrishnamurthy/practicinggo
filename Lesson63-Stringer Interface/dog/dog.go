package dog

import "fmt"

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) String() string {
	return fmt.Sprintf("My name is %s, I'm a %s! Woof!", d.Name, d.Breed)
}
