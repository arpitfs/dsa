package round2

import "fmt"

func panicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()
	panic("Panicked")
}
