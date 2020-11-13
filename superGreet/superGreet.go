package main

import (
	"fmt"
	"time"
)

func main() {
	var sGreet string
	var sSuperGreet string

	channel := make(chan string)

	go greet("Peanut", channel)
	go superGreet("butter", channel)

	sGreet = <-channel
	sSuperGreet = <-channel

	fmt.Println(sGreet)
	fmt.Println(sSuperGreet)
}

func greet(name string, ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- fmt.Sprintf("hello %s", name)
}

func superGreet(name string, ch chan string) {
	time.Sleep(time.Second * 30)
	ch <- fmt.Sprintf("HELLO %s", name)
}