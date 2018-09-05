package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func greet(name string) {
	fmt.Printf("hello, %s!\n", name)
}

func main() {
	vm := otto.New()
	// firstTest(vm)
	secondTest(vm)
	// thirdTest(vm)

}

func firstTest(vm *otto.Otto) {
	if err := vm.Set("greetFromGo", greet); err != nil {
		panic(err)
	}

	// `hello, friends!`
	if _, err := vm.Run(`greetFromGo('friends')`); err != nil {
		panic(err)
	}

	if _, err := vm.Run(`function greetFromJS(name) {
    console.log('hello, ' + name + '!');
  }`); err != nil {
		panic(err)
	}

	// `hello, friends!`
	if _, err := vm.Call(`greetFromJS`, nil, "friends"); err != nil {
		panic(err)
	}

	if _, err := vm.Run("var x = 1 + 1"); err != nil {
		panic(err)
	}

	val, err := vm.Get("x")
	if err != nil {
		panic(err)
	}

	v, err := val.Export()
	if err != nil {
		panic(err)
	}

	// (all numbers in JavaScript are floats!)
	// `float64: 2`
	fmt.Printf("%T: %v\n", v, v)

	if _, err := vm.Run(`function add(a, b) {
    return a + b;
  }`); err != nil {
		panic(err)
	}

	r, err := vm.Call("add", nil, 2, 3)
	if err != nil {
		panic(err)
	}

	// `5`
	fmt.Printf("%s\n", r)

}

func secondTest(vm *otto.Otto) {

	fn := `function animalCount(animalExp, animalString) {
		var regExp = /\b\d+ (pig|cow|chicken)s?\b/;
		// var regExp = new RegExp(animalExp);
		console.log(animalString, "\t", animalExp);
		console.log(regExp)
		console.log(regExp.test(animalString));
		return regExp.test(animalString);
		// console.log("test");
	}`

	animalExp := `/\b\d+ (pig|cow|chicken)s?\b/`
	animalString := `15 pigs`

	if _, err := vm.Run(fn); err != nil {
		panic(err)
	}

	// `hello, friends!`
	val, err := vm.Call(`animalCount`, nil, animalExp, animalString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T : %v", val, val)
}

func thirdTest(vm *otto.Otto) {

	fn := `function animalCount(expString, inString) {
		var regExp = /\b\d+\b/g;
		console.log(inString, "\t", expString);
		console.log(regExp)
		console.log(regExp.exec(inString));
		return regExp.exec(inString);
		// console.log("test");
	}`

	expString := `/\b\d+\b/g`
	inString := `"A string with 3 numbers in it... 42 and 88.`

	if _, err := vm.Run(fn); err != nil {
		panic(err)
	}

	// `hello, friends!`
	val, err := vm.Call(`animalCount`, nil, expString, inString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T : %v", val, val)
}
