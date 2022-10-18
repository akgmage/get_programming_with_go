package main

import (
	"errors"
	"fmt"
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
// custom types as errors by implementing the Error() method
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// test out error-returning functions
	for _, i := range []int{6, 43} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
}