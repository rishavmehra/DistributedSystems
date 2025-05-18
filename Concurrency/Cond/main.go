package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	wg.Add(2)

	go func() {
		fmt.Printf("Goroutine 1 is strated\n")
		defer wg.Done()

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Goroutine 1 is waiting for the condition")
		cond.Wait()
		fmt.Println("Goroutine 1 met the condition")

		fmt.Println("Goroutine 1 is Done")
	}()

	go func() {
		fmt.Println("Goroutine 2 is started")
		defer wg.Done()

		time.Sleep(5 * time.Second)

		cond.L.Lock()
		defer cond.L.Unlock()

		fmt.Println("Goroutine 2 is signaling for the condition")
		cond.Signal()
		fmt.Println("Goroutine 2 completed the signaling")

		fmt.Println("Goroutine 2 is Done")

	}()
	wg.Wait()
}
