package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

func Schedule(tasks []Task) {
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)

		go func(idx int, task Task) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Println("Task", idx, "completed in", duration)
		}(i, task)
		task()
	}

	wg.Wait()
}

func main() {
	tasks := []Task{
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Task 1 executed")
		},

		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Task 2 executed")
		},

		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("Task 3 executed")
		},
	}
	Schedule(tasks)
}
