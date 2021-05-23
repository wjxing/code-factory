package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Printf("[%s]监控退出，停止了...\n", time.Now())
				return
			default:
				fmt.Printf("[%s]goroutine监控中...\n", time.Now())
				time.Sleep(4 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Printf("[%s]可以了，通知监控停止\n", time.Now())
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(2 * time.Second)
}
