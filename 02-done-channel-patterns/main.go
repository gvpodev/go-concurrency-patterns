package main

import "time"

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			println("default")
		}
	}
}

func main() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)
}
