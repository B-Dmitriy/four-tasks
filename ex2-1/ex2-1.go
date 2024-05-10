package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func(iteration int) {
			fmt.Println(iteration)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
