package main

import "fmt"

func Func1() { fmt.Println("Func1 call") }
func Func2() { fmt.Println("Func2 call") }
func Func3() { fmt.Println("Func3 call") }

func DeferExample1() {
	defer Func1()
	defer Func2()
	defer Func3()

	/*  Can be written as follows:
	defer func() {
		Func1()
		Func2()
		Func3()
	}()
	*/
}

func DeferExample2() {
	var s string
	defer fmt.Println(s)

	s = "Hello :-)"
	return
}

func main() {
	var s = DeferExample3()
	fmt.Println(s)
}

func DeferExample3() (s string) {
	defer func() {
		fmt.Println(s)
		s += " World"
	}()

	return "Hello"
}
