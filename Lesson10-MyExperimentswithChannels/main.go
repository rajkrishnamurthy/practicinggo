package main

func main() {
	countervalue := 0
	channel1 := make(chan int)
	chsemaphore1 := make(chan bool)
	chsemaphore2 := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			println("Inside Producer: Counter Value = ", countervalue)
			countervalue++
			channel1 <- 1
		}
		chsemaphore1 <- true
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			println("Inside Consumer: Counter Value = ", countervalue)
			<-channel1
		}
		chsemaphore2 <- true
	}()
	<-chsemaphore1
	<-chsemaphore2
	close(channel1)
	close(chsemaphore1)
	close(chsemaphore2)
}
