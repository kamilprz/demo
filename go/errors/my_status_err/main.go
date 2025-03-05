package main

import "fmt"

// ############## First define what is a Status

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

// ############## Then define a custom error type which uses Status

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

// ############## Sample program using the custom error type

func main() {
	// default  -> success
	// bad user -> InvalidLogin
	// bad file -> NotFound

	_, err := loginAndGetData("user", "pwd", "file")
	if err != nil {
		fmt.Println(err)
		if se, ok := err.(StatusErr); ok {
			switch se.Status {
			case InvalidLogin:
				fmt.Println("handle invalid login")
			case NotFound:
				fmt.Println("handle not found")
			}
		}
	} else {
		fmt.Println("Success")
	}

	fmt.Println("---")

	// Both of these calls return non-nil values, initializing the custom error. No surprises.

	err = GenerateErrorInitialized(true)
	fmt.Println("GenerateErrorInitialized(true) returns nil value:", err == nil)
	err2 := GenerateErrorInitialized2(true)
	fmt.Println("GenerateErrorInitialized2(true) returns nil value:", err2 == nil)

	fmt.Println("---")

	// The results differ however when the flag is false.
	err = GenerateErrorInitialized(false)
	fmt.Println("GenerateErrorInitialized(false) returns nil value:", err == nil)
	err2 = GenerateErrorInitialized2(false)
	fmt.Println("GenerateErrorInitialized2(false) returns nil value:", err2 == nil)

	// The reason err is non-nil is that `error` is an interface.
	// For an interface to be considered nil, both the underlying type and the underlying value must be nil.
	// In this case, the underlying type is StatusErr - meaning the interface is not nil.

	// Two ways to fix this:
	// 1. Explicitly return `nil` rather than an empty error - as in `GenerateErrorInitialized2`
	// 2. Ensure that you define the variable as type `error` not the custom defined type

	err = GenerateErrorInitialized3(false)
	fmt.Println("GenerateErrorInitialized3(false) returns nil value:", err == nil)

}

// Dummy function to simulate some work
func loginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

// Simulate a login
func login(uid, pwd string) (string, error) {
	if uid == "user" && pwd == "pwd" {
		return "token", nil
	}
	return "", fmt.Errorf("login failed %s", uid)
}

// Simulate getting data from a file
func getData(token, file string) ([]byte, error) {
	if file == "file" && token == "token" {
		return []byte("data"), nil
	}
	return nil, fmt.Errorf("getData failed %s", file)
}

// return genErr of type `StatusErr`
func GenerateErrorInitialized(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

// return genErr of type `error`
func GenerateErrorInitialized2(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

// return nil
func GenerateErrorInitialized3(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}
