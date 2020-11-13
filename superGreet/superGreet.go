package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go func(name string, ch chan string) {
		println("greet started")
		time.Sleep(time.Second * 2)
		ch <- fmt.Sprintf("hello %s", name)
		println("greet finished")
	}("Peanut", channel)

	go func(name string, ch chan string) {
		println("superGreet started")
		time.Sleep(time.Second * 10)
		ch <- fmt.Sprintf("HELLO %s", name)
		println("superGreet finished")
	}("butter", channel)

	fmt.Println("greet:", <-channel)
	fmt.Println("superGreet:", <-channel)
}