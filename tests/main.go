package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1")
	}()

	go func() {
		nums := []int{1, 2, 3, 4, 5}
		defer close(ch)

		for _, n := range nums {
			fmt.Println("Goroutine 2. Sending: ", n)
			ch <- n
		}
	}()

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("Done")
}
