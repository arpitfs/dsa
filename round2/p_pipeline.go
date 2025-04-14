package round2

import "fmt"

func pipelinePattern() {
	nums := []int{1, 2, 3, 4}
	output := generator(nums)
	resultantOutput := squareRoot(output)
	for r := range resultantOutput {
		fmt.Println(r)
	}

}

func generator(nums []int) chan int {
	output := make(chan int)
	go func() {
		for n := range nums {
			output <- n
		}
		close(output)
	}()
	return output
}

func squareRoot(ch chan int) chan int {
	output := make(chan int)
	go func() {
		for n := range ch {
			output <- n * n
		}
		close(output)
	}()
	return output
}
