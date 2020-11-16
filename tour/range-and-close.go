package main

import "fmt"

func fibonacci(n int, channel chan int) {
	println("fibonacci started")
	x, y := 0, 1

	for i := 0; i < n; i++ {
		channel <- x
		x, y = y, x+y
	}

	close(channel)
	println("fibonworkeracci finished")
}

func main() {
	channel := make(chan int, 10)
	capacity := cap(channel)
	go fibonacci(capacity, channel)
	for i := range channel {
		fmt.Println(i)
	}
}
