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

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)

}
