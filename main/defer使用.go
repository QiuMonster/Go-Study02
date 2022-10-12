package main

import "fmt"

//defer 关键字
//defer用于注册延时调用  直到return前才被执行 可以用来进行资源清理
//多个defer按照先进后出的方式 栈

func go_defer() {
	go_defer1()
}
func go_defer1() {
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")
	//当将所有的非defer修饰的语句跑完到函数return时   
	//才会执行defer修饰的语句 并按照栈的方式进行执行
}
