package main

import "fmt"

type Flyer interface {
	fly()
}
type Swimmer interface {
	swim()
}

//定义一个嵌套的接口
type FlyFish interface {
	//使用子接口
	Flyer
	Swimmer
}

//定义一个结构体 继承嵌套接口 FlyFish
type Fish struct {
	name string
}

//重写所继承接口的方法
func (f Fish) fly() {
	fmt.Println(f.name, "可以飞行")
}
func (f Fish) swim() {
	fmt.Println(f.name, "可以游泳")
}
func intface_intface() {
	//实现的类型的上转型
	//接口中的方法全部实现 才能通过嵌套接口定义对象
	var fish FlyFish
	fish = Fish{"big fish"}

	f := Fish{"little fish"}
	f.fly()
	fish.swim()
}
