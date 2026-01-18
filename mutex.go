package main

import (
	"fmt"
	"sync"
)

func WaitGroupMutex() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	num := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mu.Lock()
			for n := 1; n <= 1000; n++ {
				num++
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("计数器的值：%d", num)
}

func main() {
	WaitGroupMutex()
}
