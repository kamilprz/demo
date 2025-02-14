package main

import (
	"context"
	"fmt"
	"time"
)

// Context key type to avoid conflicts
type contextKey string

const requestIDKey contextKey = "requestID"

// Simulate a database query or an API call that takes time
func slowOperation(ctx context.Context, name string, duration time.Duration) {
	// Extract request ID from context
	reqID, ok := ctx.Value(requestIDKey).(int)
	if !ok {
		reqID = -1 // Default value if request ID is missing
	}

	select {
	case <-time.After(duration):
		fmt.Printf("[%s] ---> [Request-%d] [%s] Completed successfully!\n", time.Now().Format("2006-01-02 15:04:05"), reqID, name)
	case <-ctx.Done():
		fmt.Printf("[%s] ---> [Request-%d] [%s] Canceled! Reason: %v\n", time.Now().Format("2006-01-02 15:04:05"), reqID, name, ctx.Err())
	}
}

// Simulates an HTTP request handling process
func handleRequest(parentCtx context.Context, id int) {
	fmt.Printf("[%s] ---> Handling request #%d\n", time.Now().Format("2006-01-02 15:04:05"), id)

	// Create a request-specific context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()

	// Store request ID inside the context
	ctx = context.WithValue(ctx, requestIDKey, id)

	// Simulate calling two slow operations (e.g., DB queries, external APIs)
	go slowOperation(ctx, fmt.Sprintf("DB-Query-%d", id), 60*time.Second) // This will timeout
	go slowOperation(ctx, fmt.Sprintf("API-Call-%d", id), 1*time.Second)  // This will succeed

	select {
	case <-ctx.Done():
		fmt.Printf("[%s] ---> [Request-%d] Timeout! Aborting...\n", time.Now().Format("2006-01-02 15:04:05"), id)
	}
}

func main() {
	// Create a root context (e.g., representing the server)
	rootCtx := context.Background()

	// Simulate handling multiple HTTP requests concurrently
	go handleRequest(rootCtx, 101)
	go handleRequest(rootCtx, 202)

	// Allow time for simulation
	time.Sleep(4 * time.Second)

	fmt.Printf("[%s] ---> Server shutting down.\n", time.Now().Format("2006-01-02 15:04:05"))
}
