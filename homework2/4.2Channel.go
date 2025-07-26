package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	go func() {
		for i := 0; i < 50; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 51; i < 100; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for num := range ch1 {
		fmt.Printf("Get ch1 %d ", num)
	}

	for num := range ch2 {
		fmt.Printf("Get ch2 %d ", num)
	}

	for ch1 != nil || ch2 != nil {
		select {
		case num1, ok := <-ch1:
			if ok {

				fmt.Printf("Get ch1 %d ", num1)
			}
		case num2, ok := <-ch2:
			if ok {

				fmt.Printf("Get ch2 %d ", num2)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
