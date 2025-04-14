package round2

import "fmt"

func channels() {
	unbufferedCh := make(chan int)
	bufferedCh := make(chan int, 2)
	unbufferedChannel(unbufferedCh)
	bufferedChannel(bufferedCh)

}

func unbufferedChannel(ch chan int) {
	go func() {
		ch <- 1
	}()
	fmt.Println("Unbuffered CHannel", <-ch)
}

func bufferedChannel(ch chan int) {
	go func() {
		defer func() {
			close(ch)
		}()
		ch <- 2
		ch <- 3
	}()
	fmt.Println("Buffered CHannel")
	for c := range ch {
		fmt.Println(c)
	}
}
