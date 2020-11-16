package main

import (
	"fmt"
	"time"
)

func main() {
	const numTasks = 5
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	for w := 1; w <= 3; w++ {
		go func (id int, tasks <-chan int, results chan<- int) {
			for j := range tasks {
				fmt.Println("Task", j, "started on worker", id)
				time.Sleep(time.Second)
				fmt.Println("Task", j, "finished on worker", id)
				results <- j + 10
			}
		}(w, tasks, results)
	}

	for j := 1; j <= numTasks; j++ {
		tasks <- j
	}

	close(tasks)

	for r := range results {
		fmt.Println("result:", r)
	}
}


