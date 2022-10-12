package main

import (
	"fmt"
	"sync/atomic" //原子性 实现资源的正确请求使用
	"time"
)



//不同的线程对资源的使用 不加保护的时候 会使数据出错
func go_atomic() {
	// go_atomic1()
	go_atomic2()
}
//注意资源定义的类型 要符合原子操作的要求才行
var i2 int32=100
//使用原子性来解决并发线程同时请求资源 从而出错的问题
func go_atomic2() {
	for i := 0; i < 100; i++ {
		go add2()
		go sub2()
	}
	//防止主协程退出
	time.Sleep(time.Second * 2)
	fmt.Println("当前的 i2值是:", i2)
}
//实现原子性的添加操作
func add2(){
	//原子性实现 是使用了 cas compare  new 与 old进行比较 交换数据
	atomic.AddInt32(&i2,1)
}

func sub2(){
	atomic.AddInt32(&i2,-1)
}


var i1 int = 100

func go_atomic1() {
	for i := 0; i < 100; i++ {
		go add1()
		go sub1()
	}
	//防止主协程退出
	time.Sleep(time.Second * 2)
	fmt.Println("当前的 i1值是:", i1)
}
func add1() {
	lock.Lock()
	i1--
	lock.Unlock()
}
func sub1() {
	lock.Lock()
	i1++
	lock.Unlock()
}
