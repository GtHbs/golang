package base

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	go的并发指让某个函数独立于其他函数运行的能力，一个goroutine就是一个独立的工作单元，
	go的runtime会在逻辑处理器上调度goroutine来运行，一个逻辑处理器绑定一个os线程，所以说goroutine不是线程，而是协程
	逻辑处理器：执行创建的goroutine，绑定一个线程
	调度器：go运行时中的，分配goroutine给不同的逻辑处理器
	全局运行队列：所有刚创建的goroutine会放在这里
	本地运行队列：逻辑处理器的goroutine队列
*/

func ExecuteMultiProcessor() {
	// 创建信号量
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// 函数执行完之前执行该语句，信号量-1
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("goroutineA:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("goroutineB:", i)
		}
	}()
	// 等待信号量归零
	wg.Wait()
}

func ExecuteOneProcessor() {
	// 强制指定使用一个逻辑处理器，可以看出，每次只会调度一个goroutine
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// 函数执行完之前执行该语句，信号量-1
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("goroutineA:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("goroutineB:", i)
		}
	}()
	// 等待信号量归零
	wg.Wait()
}
