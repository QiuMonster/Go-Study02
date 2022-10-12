package main

import (
	"fmt"
	"runtime"
	"time"
)

//runtime 与CPU的调度有关
//runtime.Gosched() 使用后 会先执行其他协程
func go_runtime() {
	// go_runtime1()
	// go_runtime2()
	// time.Sleep(time.Second)
	go_runtime3()
}

//runtime.GOMAXPROCS 修改cpu的核心数 runtime.NumCPU()查看cpu的最大核心数
func go_runtime3() {
	fmt.Println("CPU的核心数为:",runtime.NumCPU())
	runtime.GOMAXPROCS(1) //将CPU的核心数设置为1时只能执行一个协程  不能实现交替执行的现象
	go a()
	go b()
	time.Sleep(time.Second)
}
func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

//runtime.Goexit() 退出当前协程
func go_runtime2() {
	for i := 0; i < 10; i++ {
		fmt.Println("输出:", i)
		if i > 5 {
			runtime.Goexit() //退出当前的协程
		}

	}
}

//runtime.Gosched()
func go_runtime1() {
	go show1("java")
	for i := 0; i < 2; i++ {
		runtime.Gosched() //让出cpu的时间片 给协程使用
		fmt.Println("golang")
		// runtime.Goexit()
	}
	fmt.Println("end...")
}
func show1(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}
