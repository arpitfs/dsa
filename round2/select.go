package round2

import "fmt"

func selectProgram() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producerOne(ch1, "Channel One")
	go producerTwo(ch2, "Channel Two")

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}

func producerOne(ch chan string, data string) {
	ch <- data
}

func producerTwo(ch chan string, data string) {
	ch <- data
}
