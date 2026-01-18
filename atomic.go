package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicDemo() {
	var wg sync.WaitGroup
	var num int64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for n := 1; n <= 1000; n++ {
				atomic.AddInt64(&num, 1)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("计数器的值：%d", atomic.LoadInt64(&num))
}

func main() {
	AtomicDemo()
}
