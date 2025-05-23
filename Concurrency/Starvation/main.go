package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go starvedWorker()

	for i := 0; i < 3; i++ {
		go greedyworker(i)
	}

	time.Sleep(2 * time.Second)

}

var lock sync.Mutex

func starvedWorker() {
	for {
		time.Sleep(100 * time.Millisecond)
		lock.Lock()
		fmt.Println("Starved worker got the lock")
		time.Sleep(10 * time.Millisecond)
		lock.Unlock()
	}
}

func greedyworker(id int) {
	for {
		lock.Lock()
		fmt.Printf("Greedy worker %d got the lock\n ", id)
		time.Sleep(10 * time.Millisecond)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
}
