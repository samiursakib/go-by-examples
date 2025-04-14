package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func g(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg, nil
}

func customErrors() {
	_, err := g(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("error doesn't match argError")
	}
}
