package base

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	tasks     []func(int)      // 要执行的任务
	complete  chan error       // 用于通知任务全部完成
	timeout   <-chan time.Time // 任务可以在多久内完成
	interrupt chan os.Signal   // 控制强制终止的信号
}

func NewFactory(t time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),        // 通知通道
		timeout:   time.After(t),           // 表示在t时间之后该任务必须完成，超时后会往timeout chan中写入数据数据为当前时间
		interrupt: make(chan os.Signal, 1), // 终止信号
	}
}

func (runner *Runner) AddTask(task ...func(int)) {
	runner.tasks = append(runner.tasks, task...)
}

var errTimeOut = errors.New("执行者执行超时")
var errInterrupt = errors.New("执行者被中断")

func (r *Runner) Run() error {
	for i, task := range r.tasks {
		if r.IsInterrupt() {
			return errInterrupt
		}
		task(i)
	}
	return nil
}

// IsInterrupt 检测是否收到了中断信号
func (r *Runner) IsInterrupt() bool {
	select {
	// 判断是否中断
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Start() error {
	// 设置希望接收哪些信号，如果接收到改信号后发给r.interrupt
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.Run()
	}()
	// select只能对chan进行操作，switch可以对所有类型操作
	// select中多个可以操作，则随机找一个进行操作
	// 如果select中没有default则会被阻塞，循环查找可执行的
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return errTimeOut
	}
}

func RunnerExecute() {
	log.Println("任务开始...")
	timeout := 3 * time.Second
	runner := NewFactory(timeout)
	runner.AddTask(CreateTask(), CreateTask(), CreateTask())
	if err := runner.Start(); err != nil {
		switch err {
		case errTimeOut:
			log.Println(err)
			os.Exit(1)
		case errInterrupt:
			log.Println(err)
			os.Exit(1)
		}
	}
	log.Println("任务执行完成....")
}

func CreateTask() func(int) {
	return func(id int) {
		log.Println("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
