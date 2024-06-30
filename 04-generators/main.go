package main

import (
	"fmt"
	"math/rand"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func main() {
	done := make(chan bool)
	defer close(done)

	randNumFetcher := func() int { return rand.Intn(100000) }
	for r := range repeatFunc(done, randNumFetcher) {
		fmt.Println(r)
	}
}
