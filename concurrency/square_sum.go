package concurrency

import (
	"fmt"
	"math"
	"sync"
)

// Use sync.WaitGroup to launch multiple goroutines that compute the square of numbers in a slice and print the results.

type Result struct {
	mu  sync.Mutex
	set map[int]int
}

func squareSum() {
	var wg sync.WaitGroup
	var result Result
	result.set = make(map[int]int)

	input := []int{1, 2, 3}
	for _, val := range input {
		wg.Add(1)
		go result.findSquare(val, &wg)
	}

	wg.Wait()
	fmt.Println(result.set)
}

func (r *Result) findSquare(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := math.Pow(float64(value), 2)
	r.mu.Lock()
	r.set[value] = int(square)
	r.mu.Unlock()

}
