package main

import (
	"fmt"
	"sync"
	"time"
)

// the processes still depend on each other and can never finish their tasks.

func main() {
	c1 := Chopstick{}
	c2 := Chopstick{}

	go dine("Rishav", &c1, &c2)
	go dine("Harsh", &c2, &c1)
	time.Sleep(time.Second * 1)
}

type Chopstick struct {
	sync.Mutex
}

func dine(name string, left, right *Chopstick) {
	for {
		left.Lock()
		fmt.Println(name, "pick up left chopstick")

		time.Sleep(1 * time.Millisecond)

		locked := tryLock(right)
		if !locked {
			left.Unlock()
			fmt.Println(name, "put down the chopstick and try it again")
			continue
		}

		fmt.Println(name, "is eating 🍽️")
		time.Sleep(10 * time.Millisecond)

		right.Unlock()
		left.Unlock()
		fmt.Println("finished eating put down both chopstick")
		break

	}
}

func tryLock(c *Chopstick) bool {
	locked := make(chan struct{}, 1)
	go func() {
		c.Lock()
		locked <- struct{}{}
	}()

	select {
	case <-locked:
		return true
	case <-time.After(2 * time.Millisecond):
		return false
	}

}
