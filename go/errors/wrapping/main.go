package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker() error {
	f, err := os.Open("file.txt")
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func mergeMultipleErrorf() error {
	f, err := os.Open("file.txt")
	if err != nil {
		err1 := errors.New("first error")
		err2 := errors.New("second error")
		err3 := errors.New("third error")
		return fmt.Errorf("first: %w, second: %w, third: %w", err1, err2, err3)
	}

	f.Close()
	return nil
}

func joinMultiple() error {
	f, err := os.Open("file.txt")
	if err != nil {
		var errs []error
		errs = append(errs, errors.New("field FirstName cannot be empty"))
		errs = append(errs, errors.New("field LastName cannot be empty"))
		errs = append(errs, errors.New("field Age cannot be negative"))
		return errors.Join(errs...)
	}

	f.Close()
	return nil
}

func main() {
	err := fileChecker()
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	fmt.Println("---")

	// this is different than wrapping
	err = mergeMultipleErrorf()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---")

	// also different than wrapping
	err = joinMultiple()
	if err != nil {
		fmt.Println(err)
	}
}
