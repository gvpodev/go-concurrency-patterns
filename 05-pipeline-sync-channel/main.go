package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator[T any, K any](done <-chan K, fn func() T) <-chan T {
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

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

func primeFinder(done <-chan bool, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}

		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case random := <-randIntStream:
				if isPrime(random) {
					primes <- random
				}
			}
		}
	}()

	return primes
}

func main() {
	start := time.Now()
	done := make(chan bool)
	defer close(done)

	getInt := func() int { return rand.Intn(10000000) }
	randIntStream := generator(done, getInt)
	primes := primeFinder(done, randIntStream)

	for random := range take(done, primes, 10) {
		fmt.Println(random)
	}

	fmt.Println("Time taken:", time.Since(start))
}
