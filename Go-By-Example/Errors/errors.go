package main

import (
	"errors"
)

// By convention, errors are the last return value and have type error, a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message
		return -1, errors.New("can't work with 42")
	}
	// A nil value in the error position indicates that there was no error
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

