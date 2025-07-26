package main

import (
	"fmt"
	"sync"
)

func first(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func second(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go first(&wg)
	go second(&wg)

	wg.Wait()

}
