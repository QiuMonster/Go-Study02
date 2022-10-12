package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println(a.name, "正在吃饭")
}

type Dog2 struct {
	Animal        //嵌套结构体 实现继承功能
	ani    Animal //只能通过属性调用Animal中的方法
}

//通过结构体嵌套实现类似继承的功能
func go_extend() {
	d := Dog2{Animal{"小黑", 3}, Animal{"小白", 3}}
	d.eat()
	d.ani.eat()
}
