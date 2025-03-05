package main

import (
	"errors"
	"fmt"

	my_err "demos/go/errors/my_sentinel_err"
)

func main() {
	const (
		ErrFoo    = my_err.Sentinel("foo error")
		ErrBar    = my_err.Sentinel("bar error")
		ErrFooTwo = my_err.Sentinel("foo error")
	)

	var ErrFooThree = errors.New("foo error")

	// This looks like a function call, but it’s actually casting a string literal to a type that implements the error interface.
	// Changing the values of ErrFoo and ErrBar would be impossible.

	fmt.Printf("Type of ErrFoo: %T\n", ErrFoo)
	fmt.Printf("ErrFoo: %v\n", ErrFoo)
	fmt.Printf("Type of ErrBar: %T\n", ErrBar)
	fmt.Printf("ErrBar: %v\n", ErrBar)
	fmt.Printf("ErrFoo == ErrBar: %v\n", ErrFoo == ErrBar)

	// If you used the same type to create constant errors across packages, two errors would be equal if their error strings are equal.
	// They’d also be equal to a string literal with the same value.
	println("---")
	fmt.Printf("Type of ErrFooTwo: %T\n", ErrFooTwo)
	fmt.Printf("ErrFooTwo: %v\n", ErrFooTwo)
	fmt.Printf("ErrFoo == ErrFooTwo: %v\n", ErrFoo == ErrFooTwo)
	fmt.Printf("ErrFoo == \"foo error\": %v\n", ErrFoo == "foo error")

	// Meanwhile, an error created with errors.New is equal only to itself or to variables explicitly assigned its value.
	println("---")
	fmt.Printf("Type of ErrFooThree: %T\n", ErrFooThree)
	fmt.Printf("ErrFooThree: %v\n", ErrFooThree)
	fmt.Printf("ErrFoo == ErrFooThree: %v\n", ErrFoo == ErrFooThree)
	// fmt.Printf("ErrFooThree == \"foo error\": %v\n", ErrFooThree == "foo error") // compiler error
}
