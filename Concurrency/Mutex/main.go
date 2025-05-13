// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type Counter struct {
// 	mu    sync.Mutex
// 	value int
// }

// func (c *Counter) Increment() {
// 	fmt.Printf("Before: %+v \n", c)
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// 	time.Sleep(time.Second * 5)
// 	fmt.Printf("After: %+v \n", c)
// }

// func main() {
// 	counter := &Counter{}
// 	counter.Increment()
// }

package main

import "sync"

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func main()
