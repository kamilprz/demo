package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

// ###

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("HTTP error codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return slices.Equal(me.Codes, me2.Codes)
	}
	return false
}

// ###

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

// ###

// Dummy function to open a non-existing file and crash
func openFakeFile() error {
	f, err := os.Open("file.txt")
	if err != nil {
		return fmt.Errorf("in openFakeFile: %w", err)
	}
	f.Close()
	return nil
}

// Simulate an API request
func simulateAPIRequest() ([]byte, error) {
	return nil, MyErr{Codes: []int{401, 403, 404}}
}

// Simulate a database request
func simulateDbCall(id int) ([]byte, error) {
	if id == 123 {
		return nil, ResourceErr{Resource: "Database", Code: 123}
	} else if id == 456 {
		return nil, ResourceErr{Resource: "Database", Code: 456}
	}
	return nil, ResourceErr{Code: 789}
}

func main() {
	err := openFakeFile()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}

	// MyErr is not a comparable type, so we defined a custom Is method
	_, myErr := simulateAPIRequest()
	if errors.Is(myErr, MyErr{Codes: []int{401, 403, 404}}) {
		fmt.Println("The error is:", myErr)
	}

	// Match all Database errors regardless of their code
	_, resErr := simulateDbCall(123)
	if errors.Is(resErr, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", resErr)
	}

	_, resErr = simulateDbCall(456)
	if errors.Is(resErr, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", resErr)
	}

	_, resErr = simulateDbCall(789)
	if errors.Is(resErr, ResourceErr{}) {
		fmt.Println("The database is broken:", resErr)
	}
}
