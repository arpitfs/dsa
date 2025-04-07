package concurrency

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func functionLiteral() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}

func functionImmediatelyLiteral() {
	add := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(add)
}

func functionLiteralWithClouse() {
	sum := func() func(int) int {
		s := 0
		return func(x int) int {
			s += x
			return s
		}
	}()
	fmt.Println(sum(10))
	fmt.Println(sum(10))
}

func closure() {
	// firstClosure := add()
	// fmt.Println(firstClosure(1))
	// fmt.Println(firstClosure(1))
	// fmt.Println(firstClosure(1))
	// sClosure := add()
	// fmt.Println(sClosure(1))
	// fmt.Println(sClosure(1))
	// fmt.Println(sClosure(1))
	//functionLiteralWithClouse()
	//typeAssertionUsingSwitch()
	//atomicCounterMain()
	contextCancelMain()

}

var val interface{} = 1.23

func typeAssertion() {
	x, ok := val.(string)
	if ok {
		fmt.Println("string", x)
	} else {
		fmt.Println("Not a strign")
	}
}

var counterForTest int64

func atomicCounterMain() {
	threads := 10
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counterForTest, 1)
		}()
	}

	wg.Wait()
	fmt.Println(counterForTest)

}

func typeAssertionUsingSwitch() {
	switch x := val.(type) {
	case int:
		fmt.Println("Int", x)
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Default")

	}

}

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func contextCancelMain() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Signaled")
					return
				default:
					fmt.Println(i)
					time.Sleep(1 * time.Second)
				}
			}

		}(ctx, i, &wg)
	}
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}
