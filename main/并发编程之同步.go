package main

import (
	"fmt"
	"sync"
)

//实现同步的变量
var wg sync.WaitGroup

func go_witegroup() {
	//实现同步时 会输出全部的信息
	go_witegroup1()
	// go_witegroup2()
	fmt.Println("end...")
}
func go_witegroup2() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
	}
}
func showMsg(i int) {
	fmt.Println("输出:", i)
}

//实现协程之间的同步
func go_witegroup1() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //启动一个goroutine就登记+1
		go hello(i)
	}
	//使主协程进行等待 以便于将结果全部输出
	wg.Wait()
}
func hello(i int) {
	defer wg.Done() //goroutine结束就登记-1
	fmt.Println("hello Goroutine", i)
}
