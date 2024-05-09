package base

import "fmt"

func ChanInit() {
	// 声明容量为2的通道,如果大小为0，表示为无缓冲通道，要求发送goroutine和接收goroutine必须同时准备好，否则发送goroutine会被阻塞
	c := make(chan int, 2)
	// 往通道中写入数据
	c <- 2
	// 从管道中读取数据，并赋值给x
	x := <-c
	fmt.Println(x)
	// 从管道中读取数据并丢弃
	//<-c
	x = <-c
	fmt.Println(x)
	// 关闭通道，关闭后就不能往通道中写数据了，但还是可以读数据
	close(c)
}

func ChanExecute() {
	c := make(chan int)
	go func() {
		var sum int
		for i := 0; i < 10; i++ {
			sum += i
		}
		c <- sum
	}()
	fmt.Println("--------")
	// 上面的 c <- sum 未执行之前，下面的 <- c会一直阻塞，替代了锁的作用
	fmt.Println(<-c)
}

func DoubleChan() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 1
	}()

	go func() {
		value := <-c1
		c2 <- value
	}()
	fmt.Println(<-c1, <-c2)
}

func ChanInfo() {
	c := make(chan int, 2)
	// 获取通道的容量
	fmt.Println(cap(c))
	// 获取通道的元素数量
	fmt.Println(len(c))
}

func mirrorQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("A") }()
	go func() { responses <- request("B") }()
	go func() { responses <- request("C") }()
	return <-responses
}

func request(s string) (response string) {
	return s
}

func singleChan() {
	var send chan<- int
	var recv <-chan int
	fmt.Println(send, recv)
}
