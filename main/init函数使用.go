package main

import "fmt"
//不能使用快速赋值方式创建
// var s int = var_init()

//init函数没有输入参数 和返回值
//程序初始化的顺序  变量初始化->init->main
//可以创建多个init函数
func init() {
	// sin := 10
	// fmt.Println("init.....")

}

func go_init() {
	// fmt.Println("dsds")
}

//变量初始化函数
func var_init() int {
	fmt.Println("变量初始化...")
	return 100
}
