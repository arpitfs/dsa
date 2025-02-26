package concurrency

import "fmt"

// Implement a function that sends numbers 1-10 into a channel and another function that reads from the channel and prints them.

// func channelOperation() {
// 	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	ch := make(chan int, len(input))
// 	var wg sync.WaitGroup
// 	for _, val := range input {
// 		wg.Add(1)
// 		ch <- val
// 		go printValue(ch, &wg)
// 	}
// 	wg.Wait()

// }

// func printValue(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(<-ch)
// }

func channelOperation() {
	ch := make(chan int)
	go sendNumber(ch)
	receiveNumber(ch)
}

func sendNumber(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func receiveNumber(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
