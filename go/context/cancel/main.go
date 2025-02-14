package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] ---> [%s] Received cancellation signal: %v\n", time.Now().Format("2006-01-02 15:04:05"), name, ctx.Err())
			return
		default:
			fmt.Printf("[%s] ---> [%s] Working...\n", time.Now().Format("2006-01-02 15:04:05"), name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("Starting context with cancellation demo.")

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "Worker-A")
	go worker(ctx, "Worker-B")

	time.Sleep(2 * time.Second)

	fmt.Printf("[%s] ---> Cancelling context now...\n", time.Now().Format("2006-01-02 15:04:05"))
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Printf("[%s] ---> Main function exiting\n", time.Now().Format("2006-01-02 15:04:05"))
}
