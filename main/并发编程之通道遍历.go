package main

import (
	"fmt"
	// "time"
)

//将通道中的数据遍历出来
func go_for_chan() {
	// go_for_chan1()
	// time.Sleep(time.Second*2)
	go_for_chan2()
}

//法一 通过for和if进行遍历
func go_for_chan1() {
	ch := make(chan int)
	//使用内联的函数 向通道中填充数据
	myfun := func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}
	//必须开启协程进行数据的存入
	go myfun()

	//创建一个死的for循环
	for {
		//if可以使用多条件进行判断
		if val, ok := <-ch; ok {
			fmt.Println(val)
		} else {
			break //退出循环
		}

	}
	//关闭通道

}

//法2 使用 range来读取管道中的数据
func go_for_chan2() {
	ch := make(chan int)
	myfun := func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//写完后就需要关闭通道了
		close(ch)
	}
	go myfun()
	for val := range ch {
		fmt.Println(val)
	}

}
