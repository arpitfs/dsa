package concurrency

import (
	"fmt"
	"strconv"
	"sync"
)

func workingModel() {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go goOne(ch, &wg)
	go goTwo(ch, &wg)
	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i <= 3; i++ {
	// 		ch <- "From one goroutine" + strconv.Itoa(i)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i <= 3; i++ {
	// 		ch <- "From two goroutine" + strconv.Itoa(i)
	// 	}
	// }()

	go func() {
		wg.Wait()
		close(ch)

	}()

	for m := range ch {
		fmt.Println(m)
	}
}

func goOne(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		ch <- "From one goroutine" + strconv.Itoa(i)
	}
}

func goTwo(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		ch <- "From two goroutine" + strconv.Itoa(i)
	}
}
