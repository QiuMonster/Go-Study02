package main

import (
	"fmt"
	"time"
)

//通过使用 Goroutines 是并发运行的函数
//使用 go关键字即可开启一个协程
//协程比线程更轻量级

func go_go() {
	//线程顺序执行
	// show("java")
	// show("golang")
	go_go1()

}

func go_go1() {
	//线程顺序执行
	// show("java")
	// show("golang")

	// go show("java...")
	// show("golang...") //如果它前面也添加go。程序是没有 go 程序无法执行  这个协程是在main中进行的

	go show("java...")   //协程 1
	go show("golang...") //协程 2
	//将主线程 等待2秒中 让协程运行
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("main end...") //协程 3 主函数退出了 程序就结束了
}

func show(msg string) {
	for i := 1; i < 5; i++ {
		fmt.Printf("msg:%v\n", msg)
		//睡眠100毫秒
		time.Sleep(time.Millisecond * 100)
	}
}
