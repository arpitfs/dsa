package concurrency

import "fmt"

func pipelinePattern() {
	gen := generator(1, 2, 3, 4)
	s := sq(gen)
	for data := range s {
		fmt.Println("Squared:", data)
	}
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
