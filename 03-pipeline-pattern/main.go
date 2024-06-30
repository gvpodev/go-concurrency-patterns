package main

import (
	"fmt"
	"time"
)

func main() {
	//input
	nums := []int{1, 2, 3, 4, 5}

	//stage 1
	ch1 := sliceToChannel(nums)

	//stage 2
	sqChannel := sq(ch1)

	//stage 3
	for n := range sqChannel {
		fmt.Println(n)
	}
}

func sq(ch <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for n := range ch {
			fmt.Println("received sq", n)
			time.Sleep(10 * time.Second)
			out <- n * n
			fmt.Println("sent sq", n)
		}
		close(out)
	}()

	return out
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			fmt.Println("sending", n)
			out <- n
			fmt.Println("sent", n)
		}
		close(out)
	}()

	return out
}
