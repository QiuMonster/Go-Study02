package main

import (
	"fmt"
	"time"
)

//定时器的使用 内部是通过管道来实现的
func go_timer() {
	go_timer1()
}
func go_timer1() {
	timer1 := time.NewTimer(time.Second * 2) //设置等待时间
	//输出当前时间
	t1 := time.Now()
	fmt.Println(t1)
	//使用一个2秒的定时器
	t2 := <-timer1.C //等待2秒
	<-timer1.C       //再次等待2秒
	fmt.Println(t2)

	//可以使用After来创建定时   不是一个真正意义上的定时器 只是能实现一个定时任务  获取到当前的时间
	timer2 := time.After(time.Second * 2)
	t3 := <-timer2 //等待2秒
	<-timer2       //等待2秒
	fmt.Println(t3)

	timer1.Stop()                 //停止定时器
	timer1.Reset(time.Second * 1) //重新修改定时器的定时时间

}
