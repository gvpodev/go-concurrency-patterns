package main

import "time"

//func main() {
//	charCh := make(chan string, 3)
//	chars := []string{"a", "b", "c"}
//
//	for _, c := range chars {
//		select {
//		case charCh <- c:
//		}
//	}
//
//	close(charCh)
//
//	for result := range charCh {
//		println(result)
//	}
//}

func main() {
	go func() {
		for {
			select {
			default:
				println("default")
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
