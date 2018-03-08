package main

// Implement 100! with concurrency using the channels/pipeline model
// Key questions to answer:
//	What are the realizations working on concurrency?
//	Post answer at https://goo.gl/uJa99G

var incrementor int

//var inputchannel, outputchannel chan int

func main() {
	factorialnumber := 3
	inputchannel := make(chan int)
	outputchannel := make(chan int)
	mainstartflag := make(chan int)

	go func() {
		mainstartflag <- 1
	}()
	go func() {
		// Assign the factorial value to input channel
		<-mainstartflag
		println("Factorial Assigned")
		inputchannel <- factorialnumber
	}()

	go func() {
		//<-startflag
		println("Inside First Call to Computefactorial")
		outputchannel = computefactorial(inputchannel)
	}()

	println("Final output value of factorial ", factorialnumber, " = ", <-outputchannel)
	//close(outputchannel)
}

func computefactorial(p chan int) (q chan int) {
	incrementor++
	println("Inside Computefactorial Function. Iteration #", incrementor)
	outchannel := make(chan int)
	tempchannel := make(chan int)
	funcstartflag := make(chan int)

	factorialvalue := <-p
	println("Inside Computefactorial Function. Factorial Value", factorialvalue)
	if factorialvalue > 1 {
		go func() {
			println("Assigning (factorialvalue -1): factorialvalue =", factorialvalue)
			tempchannel <- (factorialvalue - 1)
			funcstartflag <- 1
		}()
	}

	go func() {
		<-funcstartflag
		if factorialvalue > 1 {
			println("Inside Computefactorial Function. Recursive Thread")
			nextfactorialvalue := <-computefactorial(tempchannel)
			outchannel <- (factorialvalue * nextfactorialvalue)
		}
		<-tempchannel
	}()

	return outchannel
}
