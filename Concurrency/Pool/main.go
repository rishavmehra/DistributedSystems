package main

import (
	"fmt"
	"sync"
)

type SomeObject struct {
	Data []byte
}

func createObject() *SomeObject {
	return &SomeObject{
		Data: make([]byte, 1024*1024),
	}
}

func main() {
	var memoryPiece int

	objectPool := sync.Pool{
		New: func() any {
			memoryPiece++
			return createObject()
		},
	}
	const worker = 1024 * 1024
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			obj := objectPool.Get().(*SomeObject)
			objectPool.Put(obj)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done", memoryPiece)

}
