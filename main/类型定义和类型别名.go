package main

import "fmt"

//类型定义  创建了一个新的类型
type Qint int

func new_type() {
	var num Qint = 12
	fmt.Printf("num的类型:%T num: %v\n",num, num)
	//给类型起别名
	type newint =int
	var num1 newint=12
	fmt.Printf("num1的类型:%T num1: %v\n",num1, num1)
}
