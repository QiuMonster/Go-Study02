package main

import (
	"fmt"
	"time"
)

func go_select() {
	go_select1()
}

func go_select1() {
	ch := make(chan int)
	ch1 := make(chan string)

	//执行一个匿名函数
	go func() {
		ch <- 100
		ch1 <- "hello world"
		defer close(ch)
		defer  close(ch1)
	}()

	for {
		//与通道相关
		select { //选择一个可以输出的条件
		case c := <-ch:
			fmt.Println(c)
		case s := <-ch1:
			fmt.Println(s)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
