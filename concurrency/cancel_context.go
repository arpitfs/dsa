package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func contextCancel() {

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	threads := 3
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go contextWorker(ctx, i, &wg)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Sending cancellation")
	cancel()

	wg.Wait()

}

func contextWorker(ctx context.Context, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received Signal")
			return
		default:
			fmt.Println("Worker Id", i)
			//time.Sleep(1 * time.Second)
		}
	}

}
