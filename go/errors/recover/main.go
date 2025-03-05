package main

import (
	"fmt"
)

func divide(i int) (int, error) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Inside recover: %s\n", v)
		}
	}()
	result := 60 / i
	fmt.Println(result)
	return result, nil
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		divide(val)
	}
}
