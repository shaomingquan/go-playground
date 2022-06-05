package goroutine 

import (
	"fmt"
	"time"
	"runtime"
)

func Goexit () {
	go func () {
		go func () {
			time.Sleep(1 * time.Second)
			fmt.Println(2)
		}()
		fmt.Println(1)
	}()

	runtime.Goexit()

	// 死锁，终止了当前goroutine，导致main未返回
	// main上面的协程未返回则继续执行其他goroutine
	// 当其他goroutine都执行完毕之后，程序无事可做，陷入死锁？？
}

// runtime.Goexit() vs return ?