package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go greet("Peanut", channel)
	go superGreet("butter", channel)

	fmt.Println("greet:", <-channel)
	fmt.Println("superGreet:", <-channel)
}

func greet(name string, ch chan string) {
	println("greet started")
	time.Sleep(time.Second * 2)
	ch <- fmt.Sprintf("hello %s", name)
	println("greet finished")
}

func superGreet(name string, ch chan string) {
	println("superGreet started")
	time.Sleep(time.Second * 20)
	ch <- fmt.Sprintf("HELLO %s", name)
	println("superGreet finished")
}