package main

import (
	"examples/patterns/behaviour"
	"examples/patterns/stream"
	"fmt"
	"log"
)

func IsTimeout(err error) bool {
	type timeout interface {
		IsTimeout() bool
	}
	t, ok := err.(timeout)
	return t.IsTimeout() && ok
}

func main() {
	if err := stream.ReadStreamExample1(); err != nil {
		log.Fatal(err)
	}

	err := behaviour.Connection()
	if ok := IsTimeout(err); ok {
		fmt.Println("err: ", err.Error())
	}
}
