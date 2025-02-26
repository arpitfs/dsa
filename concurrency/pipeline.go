package concurrency

import (
	"fmt"
	"math"
)

// Implement a pipeline where numbers are passed through multiple stages (e.g., square, double, filter even).

func pipeline() {

	inputChannel := make(chan float64, 5)
	ouptutChannel := make(chan float64, 5)
	go square(inputChannel)
	go double(inputChannel, ouptutChannel)
	print(ouptutChannel)

}

func square(ch chan float64) {
	for i := 0; i < 5; i++ {
		sq := math.Pow(float64(i), 2)
		ch <- sq
	}
	close(ch)
}

func double(input, output chan float64) {
	for val := range input {
		db := val * 2
		output <- db
	}
	close(output)
}

func print(ch chan float64) {
	for v := range ch {
		fmt.Println(v)
	}
}
