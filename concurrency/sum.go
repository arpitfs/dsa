package concurrency

import (
	"fmt"
	"sync"
)

func calculateSum(input []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0

	for _, num := range input {
		sum += num
	}

	ch <- sum
}

func sum() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}

	ch := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go calculateSum(arr1, ch, &wg)
	go calculateSum(arr2, ch, &wg)

	wg.Wait()

	close(ch)

	sum1, sum2 := <-ch, <-ch
	fmt.Println(sum1, sum2)
}
