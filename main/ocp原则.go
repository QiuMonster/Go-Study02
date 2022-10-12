package main

import "fmt"

//OCP原则 ：扩展是开放的 修改是关闭的

type Pet interface {
	eat()
	sleep()
}
type Dog1 struct {
}
type Cat1 struct {
}
type Person1 struct {
}

func (d Dog1) eat() {
	fmt.Println("小狗在吃饭")
}
func (c Cat1) eat() {
	fmt.Println("小猫在吃饭")
}

func (d Dog1) sleep() {
	fmt.Println("小狗在睡觉")
}

func (c Cat1) sleep() {
	fmt.Println("小猫在睡觉")
}
func (p Person1) care(pet Pet) {
	fmt.Println("人正在照顾宠物")
	pet.eat()
	pet.sleep()
}
func go_ocp() {
	p := Person1{}
	p.care(Dog1{})
	p.care(Cat1{})
}
