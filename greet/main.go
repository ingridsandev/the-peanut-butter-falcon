package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	println("Peanut test 1 started")
	println("- Fire 2 go routines but setting only one delta")
	WaitForRoutine(1, &waitGroup)
	println("Peanut test 1 finished")

	println("------------------------------------------------")
	println("Peanut test 2 started")
	println("- Fire 2 go routines and set 2 deltas")
	WaitForRoutine(2, &waitGroup)
	println("Peanut test 2 finished")

	println("------------------------------------------------")
}

func WaitForRoutine(delta int, wg *sync.WaitGroup) {

	wg.Add(delta)

	go SayHello("peanut", wg)
	go SayHello("butter", wg)

	wg.Wait()
}

func SayHello(name string, wg *sync.WaitGroup){
	if name == "peanut" {
		time.Sleep(2 * time.Second)
	}
	fmt.Println("Hello", name)

	wg.Done()
}