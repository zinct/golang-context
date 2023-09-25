package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Counter(ctx context.Context) chan int {
	counter := make(chan int)

	go func() {
		defer close(counter)
		num := 0
		
		for {
			select {
			case <- ctx.Done():
				fmt.Println("Request Canceled")
				return
			default:
				counter <- num
				num++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return counter
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Initial Goroutine", runtime.NumGoroutine())

	counter := Counter(ctx)

	fmt.Println("Counter Goroutine", runtime.NumGoroutine())


	for i := range counter {
		fmt.Println("Counter ke", i)
		if i >= 10 {
			cancel()
			break
		}
	}

	// Cancel the context
	time.Sleep(time.Second * 2)
	// cancel()

	time.Sleep(time.Second * 2)

	fmt.Println("End Goroutine", runtime.NumGoroutine())
	fmt.Println("Selesai")
}