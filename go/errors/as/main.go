package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	Message string
}

func (e MyErr) Error() string {
	return e.Message
}

func foo() error {
	return fmt.Errorf("an error occurred: %w", &MyErr{Message: "something went wrong"})
}

func bar() error {
	return fmt.Errorf("another error occurred: %w", fmt.Errorf("yet another error occurred"))
}

func main() {
	err := foo()
	var myErr *MyErr
	if errors.As(err, &myErr) {
		// myErr is assigned the value of the error matched by errors.As
		fmt.Printf("Inside As - Error: %s\n", myErr.Message)
	}

	err = bar()
	var myErr2 *MyErr
	if errors.As(err, &myErr2) {
		fmt.Printf("Inside As - Error: %s\n", myErr2.Message)
	} else {
		fmt.Println("Outisde As - An error occurred which is not of type *MyError")
	}
}
