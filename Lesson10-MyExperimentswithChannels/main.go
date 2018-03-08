package main

func main() {
	countervalue := 0
	channel1 := make(chan int)
	chsemaphore := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			println("Inside Producer: Counter Value = ", countervalue)
			countervalue++
			channel1 <- 1
		}
		chsemaphore <- true
	}()

	go func() {
		for true {
			println("Inside Consumer: Counter Value = ", countervalue)
			<-channel1
		}
	}()
	<-chsemaphore
	close(channel1)
	close(chsemaphore)
}
