package main

import (
	"fmt"
)

func main() {
	chanOwn := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				// time.Sleep(time.Second)
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		fmt.Println(results)
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done Receiving!")
	}

	fmt.Println("initiated channel")
	result := chanOwn()

	fmt.Println("Consumer ready - ok")
	consumer(result)
}
