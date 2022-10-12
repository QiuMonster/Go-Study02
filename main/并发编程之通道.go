package main

import (
	"fmt"
	"math/rand"
	"time"
)

//通道实现了不同协程之间的通信
//通道类型应该与发送的数据类型相同
func go_chan() {
	// go go_chan1()
	defer close(values)
	go send()
	fmt.Println("wait...")
	value:=<-values
	fmt.Println("receive:",value)
	fmt.Println("end...")
}

var values = make(chan int)

//通道发送数据的方法
func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Println("send:",value)
	// time.Sleep(time.Second*5)
	//将值存入通道
	values<-value
}

//使用go进行发送接收
func go_chan1() {
	//无缓存通道
	ch1 := make(chan int)
	ch1 <- 10
	//有缓存通道
	ch2 := make(chan int, 10)
	ch2 <- 100
	num := <-ch1
	fmt.Println(num)
}
