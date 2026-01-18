package main

import (
	"fmt"
	"sync"
	"time"
)

// 加10后的值
func IncrePointer(p *int) {
	fmt.Println("--------加10后的值--------")
	*p = *p + 10
	fmt.Printf("加10后的值：%d\n", *p)
}

// 乘以2后的值
func SignPointer(nums *[]int) {
	fmt.Println("--------乘以2后的值--------")
	for i, v := range *nums {
		(*nums)[i] = v * 2
	}
	fmt.Printf("乘以2后的值：%v\n", *nums)
}

// 1-10的奇偶数
func Parity() {
	fmt.Println("--------1-10的奇偶数--------")
	var wg sync.WaitGroup
	var nums1 []int
	var nums2 []int

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if id == 0 {
				for n := 1; n <= 10; n++ {
					if n%2 != 0 {
						nums1 = append(nums1, n)
					}
				}
			} else {
				for n := 2; n <= 10; n++ {
					if n%2 == 0 {
						nums2 = append(nums2, n)
					}
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("奇数：%v\n", nums1)
	fmt.Printf("偶数：%v\n", nums2)
}

// 任务调度
func Schedule() {
	fmt.Println("--------任务调度--------")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()

		var nums1 []int
		for n := 1; n <= 10; n++ {
			if n%2 != 0 {
				nums1 = append(nums1, n)
			}
		}

		times := time.Since(start)
		fmt.Printf("任务1执行耗时：%v\n", times)
		fmt.Printf("任务1内容：奇数：%v\n", nums1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()

		time.Sleep(1 * time.Second)
		sum := 0
		for i := 1; i <= 100; i++ {
			sum = sum + i
		}

		times := time.Since(start)
		fmt.Printf("任务2执行耗时：%v\n", times)
		fmt.Printf("任务2内容：1-100的和为：%d\n", sum)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()

		time.Sleep(2 * time.Second)
		sum := 0
		for i := 1; i <= 100; i++ {
			sum = sum + i
		}

		times := time.Since(start)
		fmt.Printf("任务3执行耗时：%v\n", times)
		fmt.Printf("任务3内容：1-100的和为：%d\n", sum)
	}()

	wg.Wait()
}

// 协程通信
func Channel01() {
	fmt.Println("--------协程通信--------")
	var wg sync.WaitGroup
	ch1 := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch1)
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch1 {
			fmt.Println(value)
		}
	}()

	wg.Wait()
}

// 协程通信-缓冲
func Channel02() {
	fmt.Println("--------协程通信-缓冲--------")
	var wg sync.WaitGroup
	ch1 := make(chan int, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch1)
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for value := range ch1 {
				fmt.Printf("消费者：%d，读取值：%d \n", id, value)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	//p := 10
	//IncrePointer(&p)
	//nums := []int{1, 2, 3}
	//SignPointer(&nums)
	//Parity()
	//Schedule()
	//Channel01()
	//Channel02()
}
