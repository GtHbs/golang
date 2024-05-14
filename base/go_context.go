package base

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func WaitGroupInit() {
	// WaitGroup 类似于信号量，和Java中的CountDownLatch功能一致
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("goroutineA")
		wg.Done()
	}()
	go func() {
		fmt.Println("goroutineB")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Done")
}

func GoroutineNotify() {
	notifier := make(chan bool)

	go func() {
		for {
			select {
			case <-notifier:
				fmt.Println("goroutineNotify")
				return
			default:
				fmt.Println("wait notify")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	notifier <- true
	time.Sleep(time.Second * 2)
}

func ContextInit() {
	// 给生成的context设置一个回调函数cancel，调用该函数，会让此context停止
	// context.Background() 返回一个空的context，该context一般用于整个context树的根节点
	// context.WithCancel() 创建一个可取消的context，该context可以追踪使用者
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			// 循环接收context的信号，判断是否需要结束
			case <-ctx.Done():
				fmt.Println("goroutineCancel")
				return
			default:
				fmt.Println("wait cancel")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 3)
	fmt.Println("Done")
	// CancelFunc类型，用于回调取消指令
	cancel()
	time.Sleep(time.Second * 1)
}

/*
MultiContextControl
Context接口
Deadline() (deadline time.Time, ok bool)	deadline为截止时间，到该时间后Context会自动取消。ok == false时表示没有设置截止时间，需要调用函数进行取消
Done() <-chan struct{}	信号量，用于发送取消信号
Err() error	返回取消的错误原因，context为什么被取消
Value(key any) any	在context上绑定一对值，该值一般是线程安全的

Context衍生继承
可以使用with系列函数，基于父Context创建子Context
withCancel函数，传递父Context，返回子Context，以及一个取消函数用来取消Context

Context使用规则
1. 不要把Context放在结构体中，要以参数的方式传递
2. 以Context作为参数的函数方法，Context应该位于参数第一位
3. 传递Context时入参不能为nil，用context.TODO代替
4. Context的Value相关方法应该传递必须的数据，不要什么数据都是用这个传递
5. Context是线程安全的，可以在多个goroutine中传递
*/
func MultiContextControl() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ctx.Deadline())
	go ContextWatch(ctx, "watcherA")
	go ContextWatch(ctx, "watcherB")
	go ContextWatch(ctx, "watcherC")

	time.Sleep(time.Second * 3)
	fmt.Println("Done")
	cancel()
	fmt.Println("finish")
}

func ContextWatch(context context.Context, name string) {
	for {
		select {
		case <-context.Done():
			fmt.Println("goroutineCancel", name)
			return
		default:
			fmt.Println("wait cancel", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func UsingContextDone(ctx context.Context, out chan<- string) error {
	for {
		str, err := DoSth(ctx)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- str:
		}
	}
}

func DoSth(ctx context.Context) (str string, err error) {
	return "", nil
}

func UsingWithFunction() {
	key := "name"
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "alone")
	// 设置了deadline后，超时后会自动触发取消
	deadlineCtx, deadlineCancel := context.WithDeadline(valueCtx, time.Now().Add(time.Second*2))
	fmt.Println(time.Now())
	fmt.Println(deadlineCtx.Deadline())
	go UsingWithValue(valueCtx)
	go UsingWithDeadline(deadlineCtx)
	fmt.Println("ctx", ctx.Value(key))
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	cancel()
	time.Sleep(time.Second * 3)
	fmt.Println("finish")
	deadlineCancel()
}

func UsingWithValue(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("name"), "canceled")
			return
		default:
			fmt.Println(ctx.Value("name"), "wait cancel")
			time.Sleep(2 * time.Second)
		}

	}
}

func UsingWithDeadline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(time.Now(), "canceled")
			fmt.Println("done")
			return
		default:
			fmt.Println("wait cancel")
			time.Sleep(1 * time.Second)
		}
	}
}
