package main

import "fmt"

type Dog struct {
	name  string
	color string
	age   int
}
type Person struct {
	dog  Dog //使用嵌套的结构体
	age  int
	name string
}

func go_struct() {
	// d := Dog{"小白", "white", 12}
	p := Person{Dog{"小白", "white", 12}, 22, "qiu"}
	fmt.Println(p.dog.age)
	d := Dog{"aa", "black", 1}
	var p1 *Dog=&d
	fmt.Printf("p: %T\n", p1)
	fmt.Printf("d: %T\n", d)
}
