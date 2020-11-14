package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	startedAt := time.Now()

	fmt.Printf("Peanut test 1 started at %v \n", time.Now())
	fmt.Println("- Fire 2 go routines but setting only one delta")
	FireAndWaitForRoutine(1, &waitGroup)
	fmt.Printf("Peanut test 1 finished at %v and took %v to complete\n", time.Now(), time.Since(startedAt))

	fmt.Println("------------------------------------------------")
	fmt.Printf("Peanut test 2 started at %v \n", time.Now())
	fmt.Println("- Fire 2 go routines and set 2 deltas")
	FireAndWaitForRoutine(2, &waitGroup)
	fmt.Printf("Peanut test 2 finished at %v and took %v to complete\n", time.Now(), time.Since(startedAt))

	fmt.Println("------------------------------------------------")
}

func FireAndWaitForRoutine(delta int, wg *sync.WaitGroup) {

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