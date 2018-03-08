package main

import (
	"fmt"
	"runtime"
)

type recstruct struct {
	input1string string
	input2string string
	input3number int64
}

func main() {
	trace()
	redirect()
	// //file, err := os.Open("inputtestfile")
	// file, err := os.Create("outputtestfile")
	// if err != nil {
	// 	log.Printf("Error : %v", err)
	// }
	// defer file.Close()

	// // var structer recstruct
	// var structer = recstruct{"\"Hello\"", "World", 0}
	// fmt.Fprintf(file, "%q %q \n", structer.input1string, structer.input2string)
	// // fmt.Fscanf(file, "%v %v \n", &structer.input1string, &structer.input2string)
	// fmt.Printf("Values = %v \n", structer)

	// structer = recstruct{"\"Hello\"", "世界", 0}
	// fmt.Fprintf(file, "%v %v \n", structer.input1string, structer.input2string)
	// // fmt.Fscanf(file, "%v %v \n", &structer.input1string, &structer.input2string)
	// fmt.Printf("Values = %v \n", structer)

	// structer = recstruct{"\"Hello\"", "உலகம்", 0}
	// fmt.Fprintf(file, "%v %v \n", structer.input1string, structer.input2string)
	// // fmt.Fscanf(file, "%v %v \n", &structer.input1string, &structer.input2string)
	// fmt.Printf("Values = %v \n", structer)
}

func redirect() {
	fmt.Printf("\n redirecting function \n")
	trace()
}

func trace() {
	pcs := make([]uintptr, 15)
	n := runtime.Callers(2, pcs)
	fmt.Printf("Printing runtime.Callers(2, pcs) = %d \n", n)
	frames := runtime.CallersFrames(pcs[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)

	for _, pc := range pcs {
		f := runtime.FuncForPC(pc)
		fmt.Printf("func.Name() = %s \n", f.Name())
	}
	fmt.Printf("Again func.Name() = %s \n", runtime.FuncForPC(pcs[0]).Name())
}
