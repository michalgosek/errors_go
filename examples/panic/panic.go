package main

import "fmt"

func Func1() { fmt.Println("Func1") }
func Func2() { fmt.Println("Func2") }
func Func3() { fmt.Println("Func3") }

func PanicExample() {
	defer func() {
		Func1()
		Func2()
		Func3()
	}()

	panic("Panic attack!")

	fmt.Println("This code will not be executed!")
}

func main() {
	PanicExample()
}
