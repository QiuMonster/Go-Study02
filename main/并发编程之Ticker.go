package main

import (
	"fmt"
	"time"
)

//可以周期性执行
func go_ticker() {
	// go_ticker1()
	ticker := time.NewTicker(time.Second * 1)
	chanint := make(chan int)

	//写数据的时候 使用协程实现
	go func() {
		//使用定时任务进行通道信息发送
		// counter := 0
		for _ = range ticker.C {
			// if counter > 10 {
			// 	break
			// }
			// counter++
			select { //直接选择一个case进行数据的写入
			case chanint <- 1:
			case chanint <- 2:
			case chanint <- 3:
				// fmt.Println("send 1...")
			}
		}
	}()

	sum := 0
	for val := range chanint {
		fmt.Println("收到数据", val)
		sum += val
		if sum > 10 {
			break
		}
	}

}
func go_ticker1() {
	ticker := time.NewTicker(time.Second * 2)
	counter := 0
	for _ = range ticker.C {
		fmt.Println("ticker...")
		counter++
		if counter > 5 {
			ticker.Stop()
			break
		}
	}
}
