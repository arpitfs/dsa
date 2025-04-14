package round2

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func cancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		go contextProgram(ctx, &wg, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Sending Cancellation")
	cancel()
	wg.Wait()

}

func contextProgram(ctx context.Context, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelling")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
	}
}
