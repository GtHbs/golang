package base

import (
	"fmt"
	"os"
	"time"
)

func TimeAfterTest() {
	var timeout <-chan time.Time
	fmt.Println(time.Now())
	timeout = time.After(2 * time.Second)
	go func() {
		time.Sleep(3 * time.Second)
	}()
	select {
	case t := <-timeout:
		//fmt.Println("timeout")
		fmt.Println(t)
		os.Exit(1)
	}
}
