package main

import "fmt"

func TypeAssertionFunc1() {
	var x any = 1337
	v := x.(int)
	fmt.Println(v)
}

func TypeAssertionFunc2() {
	var x any = 1337
	v := x.(string)
	fmt.Println(v)
}

func TypeAssertionFunc3() {
	var x any = 1337
	if v2, ok := x.(int); ok {
		fmt.Println(v2)
	}
}

func main() {
	TypeAssertionFunc3()
}
