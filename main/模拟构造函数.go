package main

import "fmt"

type Person2 struct {
	name string
	age  int
}

//模拟构造函数
func newperson(name string, age int) (*Person2, error) {
	p := Person2{name, age}
	return &p, nil
}
func go_constructor() {
	//通过模拟可以使用 () 来代替 {}
	p, msg := newperson("lks", 22)
	fmt.Println(p, msg)
}
