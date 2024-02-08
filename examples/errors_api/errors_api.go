package main

import (
	"errors"
	"fmt"
)

// Creating an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
var (
	err1 = errors.New("this is the first error")
	err2 = errors.New("this is the second error")
	err3 error // nil err
	err4 = errors.New("this is the fourth error")
)

func main() {
	// Creates an error that wraps the given errors. Any nil error values are discarded.
	var mergedErr = errors.Join(err1, err2, err3, err4)
	fmt.Println(mergedErr.Error())

	// Is reports whether any error in err's tree matches target
	if errors.Is(mergedErr, err1) {
		fmt.Println("mergedErr contains the err1 in it's chain")
	}
	if errors.Is(err1, err2) {
		fmt.Println("err1 does not contain err2 in it's chain")
	}

	// An error is considered to match a target if it is equal to that target or if it implements a method Is(error)
	// bool such that Is(target) returns true.
	var err5 ExampleError
	if errors.Is(err5, err1) {
		fmt.Println("err5 is  a err1")
	}

	// An easy way to create wrapped errors is to call fmt.Errorf and apply the %w verb to the error argument
	err9 := errors.New("error9")
	err10 := fmt.Errorf("error10: [%w]", err9)

	// Unwrap returns the result of calling the Unwrap method on err, if err's type contains an Unwrap method returning error.
	// Otherwise, Unwrap returns nil.
	var err11 = errors.Unwrap(err10)
	fmt.Println("Unwrapped error from err11 is:", err11.Error())

	// Unwrap only calls a method of the form "Unwrap() error". In particular Unwrap does not unwrap errors returned by Join.
	var err12 = errors.Unwrap(mergedErr)
	fmt.Println(err12.Error())

	// As finds the first error in err's tree that matches target, and if one is found,
	// sets target to that error value and returns true. Otherwise, it returns false.
	var err6 ExampleError
	err7 := fmt.Errorf("error7: [%w]", err6)

	var err8 ExampleError
	if errors.As(err7, &err8) {
		fmt.Println("err7 found the Example Error node and assigned it to err8")
	}
}

type ExampleError string

func (e ExampleError) Error() string { return string(e) }

func (e ExampleError) Is(target error) bool { return target == err1 }
