package base

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	c     int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func ConcurrentExecute() {
	//runtime.GOMAXPROCS(1)
	wg.Add(2)
	go intCount()
	go intCount()
	wg.Wait()
	fmt.Println("----------")
	fmt.Println(c)
}

func intCount() {
	defer wg.Done()
	for i := 0; i <= 2; i++ {
		// 使用锁
		mutex.Lock()
		// 使用原子类进行自增
		//value := atomic.LoadInt32(&c)
		cc := c
		// 这一句是让当前goroutine暂停的意思，退回执行队列
		runtime.Gosched()
		//value++
		//atomic.AddInt32(&c, 1)
		//atomic.StoreInt32(&c, value)
		cc++
		c = cc
		mutex.Unlock()
	}
}
