package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()
	// time.Sleep(time.Second * 3)
	if data == 0 {
		fmt.Printf("this value is %d./n", data)
	}
}
