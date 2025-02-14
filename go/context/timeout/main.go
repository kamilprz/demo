package main

import (
	"context"
	"fmt"
	"time"
)

// simulate some work
func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Exiting: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("[%s] ---> [%s] Running...\n", time.Now().Format("2006-01-02 15:04:05"), name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("Starting context with timeout demo - timeout 2 seconds.")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker(ctx, "Worker-1")
	go worker(ctx, "Worker-2")

	time.Sleep(3 * time.Second)
	fmt.Printf("[%s] ---> Main function exiting\n", time.Now().Format("2006-01-02 15:04:05"))
}
