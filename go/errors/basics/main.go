package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := divide(10, 0)
	if err != nil {
		println("#1: " + err.Error())
	}

	_, err2 := divide(10, 0)
	if err2 != nil {
		// If you pass an error to fmt.Println, it calls the `Error` method automatically
		fmt.Println(err2)
	}

	_, err3 := divideFmt(10, 0)
	if err3 != nil {
		fmt.Printf("#3: %s\n", err3)
	}
}

func divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("denominator is 0")
	}
	return numerator / denominator, nil
}

func divideFmt(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("cannot divide by 0")
	}
	return numerator / denominator, nil
}
