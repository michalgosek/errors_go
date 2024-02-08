package main

import "fmt"

func RecoverExample() {
	var x = recover()
	if x == nil {
		fmt.Println("nil value after calling recover in normal execution flow")
	}
}

func main() {
	RecoverExample2()
	fmt.Println("The main function continues")
}

func PanicFunc(s string) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("Panic attack!")
	panic(s)
}

func RecoverExample2() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Recovered: %s\n", v)
		}
	}()
	PanicFunc("Hello!")
}
