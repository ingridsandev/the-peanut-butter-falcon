package calculator

func SquareNumbers(numbers ...int) <-chan int {
	pipelineChannel := SetPipeline(numbers...) // Response should be 4 and 9
	squaredChannel := SquareChannel(pipelineChannel)

	return squaredChannel
}

func SetPipeline(numbers ...int) <-chan int {
	channel := make(chan int)
	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()
	return channel
}

func SquareChannel(in <-chan int) <-chan int {
	channel := make(chan int)
	go func() {
		for n := range in {
			channel <- n * n
		}
		close(channel)
	}()
	return channel
}