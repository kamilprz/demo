// In this example, we simulate an HTTP request handler where:

// A parent context is created when a request starts.
// The request handler launches multiple Goroutines (e.g., for DB queries, API calls, etc.).
// If the request times out or is manually canceled, all child operations should stop immediately.

package main

import (
	"context"
	"fmt"
	"time"
)

// Simulate a database query or an API call that takes time
func slowOperation(ctx context.Context, name string, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Printf("[%s] ---> [%s] Completed successfully!\n", time.Now().Format("2006-01-02 15:04:05"), name)
	case <-ctx.Done():
		fmt.Printf("[%s] ---> [%s] Canceled! Reason: %v\n", time.Now().Format("2006-01-02 15:04:05"), name, ctx.Err())
	}
}

// Simulates an HTTP request handling process
func handleRequest(parentCtx context.Context, id int) {
	fmt.Printf("[%s] ---> Handling request #%d\n", time.Now().Format("2006-01-02 15:04:05"), id)

	// Create a request-specific context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()

	// Simulate calling two slow operations (e.g., DB queries, external APIs)

	// Calling operation with 3 second timeout ---> This will TIMEOUT.
	go slowOperation(ctx, fmt.Sprintf("DB-Query-%d", id), 60*time.Second)
	// Calling operation with 1 second timeout ---> This will SUCCEED.
	go slowOperation(ctx, fmt.Sprintf("API-Call-%d", id), 1*time.Second)

	select {
	case <-ctx.Done():
		fmt.Printf("[%s] ---> Request #%d Timeout! Aborting...\n", time.Now().Format("2006-01-02 15:04:05"), id)
	}
}

func main() {
	// Create a root context (e.g., representing the server)
	rootCtx := context.Background()

	// Simulate handling multiple HTTP requests concurrently
	go handleRequest(rootCtx, 1)
	go handleRequest(rootCtx, 2)

	// Allow time for simulation
	time.Sleep(4 * time.Second)

	fmt.Printf("[%s] ---> Server shutting down.\n", time.Now().Format("2006-01-02 15:04:05"))
}
