package main

import (
	"fmt"
	"sync"
)

// A deadlock is a situation in which processes block each other due to resource acquisition and none of the processes makes any progress as they wait for the resource held by the other process.

// For a deadlock to occur, four conditions, known as the Coffman conditions, must be met simultaneously: mutual exclusion, hold and wait, no preemption, and circular wait.

func main() {
	var wg sync.WaitGroup
	var mu1, mu2, mu3 sync.Mutex

	wg.Add(3)

	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("goroutine Acquired lock 1")
		mu2.Lock()
		fmt.Println("goroutine Acquired lock 2")
		mu3.Lock()
		fmt.Println("goroutine Acquired lock 3")
	}()
	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("goroutine Acquired lock 2")
		mu3.Lock()
		fmt.Println("goroutine Acquired lock 3")
		mu1.Lock()
		fmt.Println("goroutine Acquired lock 1")
	}()
	go func() {
		defer wg.Done()
		mu3.Lock()
		fmt.Println("goroutine Acquired lock 3")
		mu1.Lock()
		fmt.Println("goroutine Acquired lock 1")
		mu2.Lock()
		fmt.Println("goroutine Acquired lock 2")
	}()
	wg.Wait()
}
