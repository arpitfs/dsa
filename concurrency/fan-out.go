package concurrency

import (
	"fmt"
	"math"
)

// Distribute tasks from one channel to multiple worker goroutines and collect results.
func fanOut() {

	ch := make(chan int, 5)
	result := make(map[int]int)
	go distributor(ch)
	go worker(ch, &result)
	//	time.Sleep(2 * time.Second)
	fmt.Println(result)

}

func distributor(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i + 1
	}
}

func worker(ch chan int, result *map[int]int) {
	for value := range ch {
		data := math.Pow(float64(value), 2)
		(*result)[value] = int(data)

	}
}
