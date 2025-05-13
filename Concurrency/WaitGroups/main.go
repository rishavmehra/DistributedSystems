package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d Started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d Done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()

}
