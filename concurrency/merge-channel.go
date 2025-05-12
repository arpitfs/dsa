package concurrency

import "fmt"

func mergeChannel() {
	ch1 := generate("ping")
	ch2 := generate("pong")

	for m := range fanInGenerate(ch1, ch2) {
		fmt.Println(m)
	}
}

func generate(name string) <-chan string {
	ch := make(chan string)
	go func() {
		ch <- name
		close(ch)
	}()

	return ch
}

func fanInGenerate(ch1, ch2 <-chan string) <-chan string {
	merged := make(chan string)
	go func() {
		for {
			select {
			case msg1, ok := <-ch1:
				if ok {
					merged <- msg1
				} else {
					ch1 = nil
				}
			case msg2, ok1 := <-ch2:
				if ok1 {
					merged <- msg2
				} else {
					ch2 = nil
				}
			}

			if ch1 == nil && ch2 == nil {
				close(merged)
				return
			}
		}
	}()
	return merged
}
